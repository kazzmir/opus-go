package main

import (
	"bufio"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime/pprof"
	"strings"
	"time"

	"github.com/kazzmir/opus-go/ogg"
	"github.com/kazzmir/opus-go/opus"
	"github.com/kazzmir/opus-go/wav"
)

func main() {
	var (
		outPath     = flag.String("out", "out.opus", "output .opus file")
		cpuProfile  = flag.String("cpuprofile", "cpu.pprof", "write CPU profile to file (set to empty to disable)")
		bitrate     = flag.Int("bitrate", 64000, "target bitrate in bits/sec")
		vbr         = flag.Bool("vbr", true, "enable variable bitrate")
		complexity  = flag.Int("complexity", 10, "encoder complexity (0-10)")
		application = flag.String("application", "audio", "opus application: audio|voip|lowdelay")
		frameMS     = flag.Int("frame-ms", 20, "frame duration in ms (5|10|20|40|60)")
		vendor      = flag.String("vendor", "opusgo", "OpusTags vendor string")
	)
	flag.Parse()

	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			fatal(err)
		}
		defer func() {
			_ = f.Close()
		}()
		if err := pprof.StartCPUProfile(f); err != nil {
			fatal(err)
		}
		defer pprof.StopCPUProfile()
	}

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "usage: %s [flags] input.wav\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}
	inPath := flag.Arg(0)

	app, err := parseApplication(*application)
	if err != nil {
		fatal(err)
	}

	if _, err := frameSizeFromMS(*frameMS); err != nil {
		fatal(err)
	}

	inF, err := os.Open(inPath)
	if err != nil {
		fatal(err)
	}
	defer inF.Close()

	outF, err := os.Create(*outPath)
	if err != nil {
		fatal(err)
	}
	defer func() {
		_ = outF.Close()
	}()

	err = encode(inF, outF, options{
		bitrate:     *bitrate,
		vbr:         *vbr,
		complexity:  *complexity,
		application: app,
		frameMS:     *frameMS,
		vendor:      *vendor,
		serial:      randomSerial(),
	})
	if err != nil {
		fatal(err)
	}
}

type options struct {
	bitrate     int // <= 0 keeps the encoder default
	vbr         bool
	complexity  int // < 0 keeps the encoder default
	application int
	frameMS     int
	vendor      string
	serial      uint32
}

// encode reads a 16-bit PCM WAV from in and writes it to out as Ogg Opus.
func encode(in io.Reader, out io.Writer, opt options) error {
	frameSize, err := frameSizeFromMS(opt.frameMS)
	if err != nil {
		return err
	}

	wr, err := wav.NewReader(in)
	if err != nil {
		return err
	}
	if wr.SampleRate() != 48000 {
		return fmt.Errorf("only 48kHz WAV supported currently (got %d)", wr.SampleRate())
	}

	enc, err := opus.NewEncoder(wr.SampleRate(), wr.Channels(), opt.application)
	if err != nil {
		return err
	}
	defer enc.Close()

	if opt.bitrate > 0 {
		if err := enc.SetBitrate(opt.bitrate); err != nil {
			return err
		}
	}
	if err := enc.SetVBR(opt.vbr); err != nil {
		return err
	}
	if opt.complexity >= 0 {
		if err := enc.SetComplexity(opt.complexity); err != nil {
			return err
		}
	}

	lookahead, err := enc.Lookahead()
	if err != nil {
		return err
	}
	preSkip, err := enc.PreSkip()
	if err != nil {
		return err
	}

	outBW := bufio.NewWriterSize(out, 1<<20)
	pw := ogg.NewPacketWriter(outBW, opt.serial)

	head := ogg.OpusHead{
		Version:         1,
		Channels:        uint8(wr.Channels()),
		PreSkip:         uint16(preSkip),
		InputSampleRate: 48000,
		OutputGainQ8:    0,
		// ChannelMappingFamily=0 covers mono/stereo and lets decoders infer mapping.
		ChannelMappingFamily: 0,
	}
	headPkt, err := ogg.BuildOpusHeadPacket(head)
	if err != nil {
		return err
	}

	tags := ogg.OpusTags{
		Vendor:   opt.vendor,
		Comments: []string{"ENCODER=opusgo", "ENCODED=" + time.Now().UTC().Format(time.RFC3339)},
	}
	tagsPkt, err := ogg.BuildOpusTagsPacket(tags)
	if err != nil {
		return err
	}

	if err := pw.WritePacket(headPkt, 0, true, false); err != nil {
		return err
	}
	if err := pw.WritePacket(tagsPkt, 0, false, false); err != nil {
		return err
	}

	// The encoder's output lags its input by lookahead samples, so after
	// the input runs out it's fed silence until the real audio's last
	// sample has been pushed through. Each packet is held back one step so
	// the final one - which carries EOS and the end-trimming granule - is
	// known when it's written.
	channels := wr.Channels()
	src := &pcmSource{wr: wr, channels: channels}
	pcm := make([]int16, frameSize*channels)
	packet := make([]byte, 4000)
	var held []byte
	var heldGranule uint64
	fed, realFrames := 0, 0
	for !src.exhausted || fed < realFrames+lookahead {
		n, err := src.fill(pcm)
		if err != nil {
			return err
		}
		realFrames += n / channels
		nBytes, err := enc.Encode(pcm, frameSize, packet)
		if err != nil {
			return err
		}
		fed += frameSize
		if held != nil {
			if err := pw.WritePacket(held, heldGranule, false, false); err != nil {
				return err
			}
		}
		held = append(held[:0], packet[:nBytes]...)
		// A page's granule counts every sample decodable through it,
		// pre-skip included (RFC 7845 section 4).
		heldGranule = uint64(fed)
	}
	// End trimming: the last granule marks where the real audio stops.
	if err := pw.WritePacket(held, uint64(preSkip+realFrames), false, true); err != nil {
		return err
	}
	return pw.Flush()
}

// pcmSource yields the input's samples a frame at a time, zero-padded
// once it runs out.
type pcmSource struct {
	wr       *wav.Reader
	channels int

	pending   []int16
	eof       bool // the WAV reader is drained
	exhausted bool // eof, and every pending sample handed out
}

// fill copies the next len(dst) samples into dst, zeroing whatever the
// input can't cover, and returns how many real samples it copied.
func (s *pcmSource) fill(dst []int16) (int, error) {
	for len(s.pending) < len(dst) && !s.eof {
		chunk := make([]int16, 4096*s.channels)
		n, err := s.wr.ReadInt16PCM(chunk)
		if err != nil && !errors.Is(err, io.EOF) {
			return 0, err
		}
		s.pending = append(s.pending, chunk[:n-n%s.channels]...)
		if n == 0 || errors.Is(err, io.EOF) {
			s.eof = true
		}
	}
	n := copy(dst, s.pending)
	clear(dst[n:])
	s.pending = s.pending[n:]
	if s.eof && len(s.pending) == 0 {
		s.exhausted = true
	}
	return n, nil
}

func parseApplication(s string) (int, error) {
	s = strings.ToLower(strings.TrimSpace(s))
	s = strings.ReplaceAll(s, "_", "-")
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, "restricted-", "")
	if strings.HasPrefix(s, "app-") {
		s = strings.TrimPrefix(s, "app-")
	}

	switch s {
	case "audio":
		return opus.ApplicationAudio, nil
	case "voip":
		return opus.ApplicationVoIP, nil
	case "lowdelay", "low-delay", "restricted-lowdelay":
		return opus.ApplicationRestrictedLowDelay, nil
	default:
		return 0, fmt.Errorf("unknown application: %q", s)
	}
}

func frameSizeFromMS(ms int) (int, error) {
	switch ms {
	case 5:
		return 240, nil
	case 10:
		return 480, nil
	case 20:
		return 960, nil
	case 40:
		return 1920, nil
	case 60:
		return 2880, nil
	default:
		return 0, fmt.Errorf("unsupported frame-ms %d (use 5|10|20|40|60)", ms)
	}
}

func randomSerial() uint32 {
	var b [4]byte
	if _, err := rand.Read(b[:]); err == nil {
		return binary.LittleEndian.Uint32(b[:])
	}
	// Fallback: time-based.
	return uint32(time.Now().UnixNano())
}

func fatal(err error) {
	if err == nil {
		return
	}
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}
