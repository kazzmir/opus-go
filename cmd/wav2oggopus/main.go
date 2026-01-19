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

	"opusgo/ogg"
	"opusgo/oggopus"
	"opusgo/opus"
	"opusgo/wav"
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

	frameSize, err := frameSizeFromMS(*frameMS)
	if err != nil {
		fatal(err)
	}

	inF, err := os.Open(inPath)
	if err != nil {
		fatal(err)
	}
	defer inF.Close()

	wr, err := wav.NewReader(inF)
	if err != nil {
		fatal(err)
	}
	if wr.SampleRate() != 48000 {
		fatal(fmt.Errorf("only 48kHz WAV supported currently (got %d)", wr.SampleRate()))
	}

	enc, err := opus.NewEncoder(wr.SampleRate(), wr.Channels(), app)
	if err != nil {
		fatal(err)
	}
	defer enc.Close()

	if *bitrate > 0 {
		if err := enc.SetBitrate(*bitrate); err != nil {
			fatal(err)
		}
	}
	if err := enc.SetVBR(*vbr); err != nil {
		fatal(err)
	}
	if *complexity >= 0 {
		if err := enc.SetComplexity(*complexity); err != nil {
			fatal(err)
		}
	}

	lookahead, err := enc.Lookahead()
	if err != nil {
		fatal(err)
	}

	outF, err := os.Create(*outPath)
	if err != nil {
		fatal(err)
	}
	defer func() {
		_ = outF.Close()
	}()
	outBW := bufio.NewWriterSize(outF, 1<<20)
	defer func() {
		_ = outBW.Flush()
	}()

	serial := randomSerial()
	pw := ogg.NewPacketWriter(outBW, serial)

	head := oggopus.OpusHead{
		Version:         1,
		Channels:        uint8(wr.Channels()),
		PreSkip:         uint16(lookahead),
		InputSampleRate: 48000,
		OutputGainQ8:    0,
		// ChannelMappingFamily=0 covers mono/stereo and lets decoders infer mapping.
		ChannelMappingFamily: 0,
	}
	headPkt, err := oggopus.BuildOpusHeadPacket(head)
	if err != nil {
		fatal(err)
	}

	tags := oggopus.OpusTags{
		Vendor:   *vendor,
		Comments: []string{"ENCODER=opusgo", "ENCODED=" + time.Now().UTC().Format(time.RFC3339)},
	}
	tagsPkt, err := oggopus.BuildOpusTagsPacket(tags)
	if err != nil {
		fatal(err)
	}

	if err := pw.WritePacket(headPkt, 0, true, false); err != nil {
		fatal(err)
	}
	if err := pw.WritePacket(tagsPkt, 0, false, false); err != nil {
		fatal(err)
	}

	pcm := make([]int16, frameSize*wr.Channels())
	packet := make([]byte, 4000)
	var totalSamplesPerCh uint64

	for {
		n, rerr := wr.ReadInt16PCM(pcm)
		if rerr != nil && !errors.Is(rerr, io.EOF) {
			fatal(rerr)
		}
		if n == 0 {
			break
		}
		// n is in samples (interleaved). Ensure we have whole frames.
		if n%wr.Channels() != 0 {
			fatal(fmt.Errorf("wav: sample count not multiple of channels"))
		}
		framesRead := n / wr.Channels()
		isLast := errors.Is(rerr, io.EOF)
		if framesRead < frameSize {
			// Pad to a full Opus frame.
			for i := n; i < len(pcm); i++ {
				pcm[i] = 0
			}
			isLast = true
		}

		nBytes, err := enc.Encode(pcm, frameSize, packet)
		if err != nil {
			fatal(err)
		}

		totalSamplesPerCh += uint64(framesRead)
		granule := uint64(head.PreSkip) + totalSamplesPerCh

		if err := pw.WritePacket(packet[:nBytes], granule, false, isLast); err != nil {
			fatal(err)
		}

		if isLast {
			break
		}
	}

	if err := pw.Flush(); err != nil {
		fatal(err)
	}
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
