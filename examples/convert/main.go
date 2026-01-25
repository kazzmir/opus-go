package main

import (
	"bufio"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/kazzmir/opus-go/ogg"
	"github.com/kazzmir/opus-go/opus"
)

func formatDur(d time.Duration) string {
	if d < 0 {
		d = 0
	}
	total := int64(d.Round(time.Second) / time.Second)
	h := total / 3600
	m := (total % 3600) / 60
	s := total % 60
	if h > 0 {
		return fmt.Sprintf("%d:%02d:%02d", h, m, s)
	}
	return fmt.Sprintf("%d:%02d", m, s)
}

type progressReporter struct {
	decodedLenBytes     int64
	totalFramesEstimate int64
	start               time.Time
	lastPrint           time.Time
	spinner             []byte
	spin                int
}

func newProgressReporter(decodedLenBytes, totalFramesEstimate int64) *progressReporter {
	now := time.Now()
	return &progressReporter{
		decodedLenBytes:     decodedLenBytes,
		totalFramesEstimate: totalFramesEstimate,
		start:               now,
		lastPrint:           now,
		spinner:             []byte{'|', '/', '-', '\\'},
	}
}

func (p *progressReporter) Print(framesDone, bytesRead int64, force bool) {
	if p == nil {
		return
	}
	if !force {
		if time.Since(p.lastPrint) < 100*time.Millisecond {
			return
		}
	}
	p.lastPrint = time.Now()

	elapsed := time.Since(p.start)
	ch := p.spinner[p.spin%len(p.spinner)]
	p.spin++

	if p.totalFramesEstimate > 0 {
		pct := float64(framesDone) / float64(p.totalFramesEstimate) * 100
		if pct > 100 {
			pct = 100
		}
		etaStr := "--:--"
		if framesDone > 0 && elapsed > 0 {
			rate := float64(framesDone) / elapsed.Seconds()
			if rate > 0 {
				remaining := float64(p.totalFramesEstimate - framesDone)
				if remaining < 0 {
					remaining = 0
				}
				eta := time.Duration(remaining / rate * float64(time.Second))
				etaStr = formatDur(eta)
			}
		}
		fmt.Fprintf(os.Stderr, "\r%c %6.2f%% (%d/%d frames) elapsed %s eta %s", ch, pct, framesDone, p.totalFramesEstimate, formatDur(elapsed), etaStr)
		return
	}

	if p.decodedLenBytes > 0 {
		pct := float64(bytesRead) / float64(p.decodedLenBytes) * 100
		if pct > 100 {
			pct = 100
		}
		etaStr := "--:--"
		if bytesRead > 0 && elapsed > 0 {
			rate := float64(bytesRead) / elapsed.Seconds()
			if rate > 0 {
				remaining := float64(p.decodedLenBytes - bytesRead)
				if remaining < 0 {
					remaining = 0
				}
				eta := time.Duration(remaining / rate * float64(time.Second))
				etaStr = formatDur(eta)
			}
		}
		fmt.Fprintf(os.Stderr, "\r%c %6.2f%% (%d bytes) elapsed %s eta %s", ch, pct, bytesRead, formatDur(elapsed), etaStr)
		return
	}

	fmt.Fprintf(os.Stderr, "\r%c %d frames elapsed %s", ch, framesDone, formatDur(elapsed))
}

type linearResampler struct {
	r        io.Reader
	inRate   int
	outRate  int
	channels int
	step     float64

	pos  float64 // absolute input frame position for the next output frame
	base int64   // absolute frame index of buf[0]
	buf  []int16 // interleaved PCM frames

	eof       bool
	bytesRead int64
}

func newLinearResampler(r io.Reader, inRate, outRate, channels int) *linearResampler {
	return &linearResampler{
		r:        r,
		inRate:   inRate,
		outRate:  outRate,
		channels: channels,
		step:     float64(inRate) / float64(outRate),
		pos:      0,
		base:     0,
		buf:      nil,
	}
}

func (rs *linearResampler) BytesRead() int64 {
	if rs == nil {
		return 0
	}
	return rs.bytesRead
}

func (rs *linearResampler) totalFramesAvailable() int64 {
	if rs == nil {
		return 0
	}
	return rs.base + int64(len(rs.buf))/int64(rs.channels)
}

func (rs *linearResampler) readMore() error {
	if rs == nil || rs.eof {
		return nil
	}

	// Read a chunk of decoded PCM bytes (s16le, interleaved).
	const chunkBytes = 32 * 1024
	tmp := make([]byte, chunkBytes)
	n, err := rs.r.Read(tmp)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}
	if n == 0 {
		if errors.Is(err, io.EOF) {
			rs.eof = true
		}
		return nil
	}
	rs.bytesRead += int64(n)
	if n%2 != 0 {
		// Truncate to whole int16.
		n--
	}
	if n <= 0 {
		if errors.Is(err, io.EOF) {
			rs.eof = true
		}
		return nil
	}

	// Convert bytes to int16 samples.
	oldLen := len(rs.buf)
	rs.buf = append(rs.buf, make([]int16, n/2)...)
	for i := 0; i < n; i += 2 {
		rs.buf[oldLen+i/2] = int16(binary.LittleEndian.Uint16(tmp[i : i+2]))
	}

	if errors.Is(err, io.EOF) {
		rs.eof = true
	}
	return nil
}

// Fill fills outPCM with exactly outFrames frames at outRate.
// It returns how many of those frames contain actual (non-padded) audio,
// and whether this call produced the final output (done=true).
func (rs *linearResampler) Fill(outPCM []int16, outFrames int) (actualFrames int, done bool, err error) {
	if rs == nil {
		return 0, true, errors.New("resampler: nil")
	}
	if outFrames <= 0 || len(outPCM) < outFrames*rs.channels {
		return 0, false, errors.New("resampler: invalid output buffer")
	}
	if rs.inRate <= 0 || rs.outRate <= 0 || rs.channels <= 0 {
		return 0, false, errors.New("resampler: invalid rates/channels")
	}

	// Ensure we have some input buffered.
	for len(rs.buf) == 0 && !rs.eof {
		if err := rs.readMore(); err != nil {
			return 0, false, err
		}
	}

	totalAvail := rs.totalFramesAvailable()
	for i := 0; i < outFrames; i++ {
		idx := int64(math.Floor(rs.pos))
		frac := rs.pos - float64(idx)

		// Try to buffer enough for interpolation when not EOF.
		for !rs.eof {
			need := idx + 1
			if need < rs.totalFramesAvailable() {
				break
			}
			if err := rs.readMore(); err != nil {
				return actualFrames, false, err
			}
		}
		totalAvail = rs.totalFramesAvailable()

		if idx >= totalAvail {
			// Out of input; pad with silence.
			for c := 0; c < rs.channels; c++ {
				outPCM[i*rs.channels+c] = 0
			}
			rs.pos += rs.step
			continue
		}

		// s0 at idx always exists here.
		o0 := int((idx - rs.base) * int64(rs.channels))
		// s1 at idx+1 if available; otherwise hold s0 (EOF tail).
		o1 := o0
		if idx+1 < totalAvail {
			o1 = int((idx + 1 - rs.base) * int64(rs.channels))
		}

		for c := 0; c < rs.channels; c++ {
			s0 := float64(rs.buf[o0+c])
			s1 := float64(rs.buf[o1+c])
			s := s0 + (s1-s0)*frac
			if s > 32767 {
				s = 32767
			} else if s < -32768 {
				s = -32768
			}
			outPCM[i*rs.channels+c] = int16(s)
		}

		actualFrames++
		rs.pos += rs.step
	}

	// Drop old input we can no longer reference (keep 1 frame of history).
	minNeeded := int64(math.Floor(rs.pos)) - 1
	if minNeeded > rs.base {
		dropFrames := minNeeded - rs.base
		dropSamples := int(dropFrames) * rs.channels
		if dropSamples > 0 && dropSamples < len(rs.buf) {
			rs.buf = rs.buf[dropSamples:]
			rs.base = minNeeded
		} else if dropSamples >= len(rs.buf) {
			rs.buf = nil
			rs.base = minNeeded
		}
	}

	if rs.eof {
		// Done when the next output would start beyond available input.
		nextIdx := int64(math.Floor(rs.pos))
		if nextIdx >= rs.totalFramesAvailable() {
			done = true
		}
	}

	return actualFrames, done, nil
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Fprintf(os.Stderr, "usage: %s input.mp3 [output.opus]\n", filepath.Base(os.Args[0]))
		os.Exit(2)
	}

	inPath := os.Args[1]
	outPath := ""
	if len(os.Args) == 3 {
		outPath = os.Args[2]
	} else {
		outPath = defaultOutPath(inPath)
	}

	inF, err := os.Open(inPath)
	if err != nil {
		fatal(err)
	}
	defer inF.Close()

	dec, err := mp3.NewDecoder(inF)
	if err != nil {
		fatal(err)
	}

	const (
		channels       = 2
		frameSize48k   = 960 // 20ms @ 48kHz
		bitrate        = 64000
		vbr            = true
		complexity     = 10
		opusSampleRate = ogg.OpusSampleRateHz
	)

	inSampleRate := dec.SampleRate()
	if inSampleRate <= 0 {
		fatal(fmt.Errorf("mp3: invalid sample rate: %d", inSampleRate))
	}

	enc, err := opus.NewEncoder(opusSampleRate, channels, opus.ApplicationAudio)
	if err != nil {
		fatal(fmt.Errorf("unable to create opus encoder: %v", err))
	}
	defer enc.Close()

	if err := enc.SetBitrate(bitrate); err != nil {
		fatal(err)
	}
	if err := enc.SetVBR(vbr); err != nil {
		fatal(err)
	}
	if err := enc.SetComplexity(complexity); err != nil {
		fatal(err)
	}

	lookahead, err := enc.Lookahead()
	if err != nil {
		fatal(err)
	}

	outF, err := os.Create(outPath)
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

	head := ogg.OpusHead{
		Version:              1,
		Channels:             uint8(channels),
		PreSkip:              uint16(lookahead),
		InputSampleRate:      uint32(inSampleRate),
		OutputGainQ8:         0,
		ChannelMappingFamily: 0,
	}
	headPkt, err := ogg.BuildOpusHeadPacket(head)
	if err != nil {
		fatal(err)
	}

	tags := ogg.OpusTags{
		Vendor:   "opusgo",
		Comments: []string{"ENCODER=examples/convert", "ENCODED=" + time.Now().UTC().Format(time.RFC3339)},
	}
	tagsPkt, err := ogg.BuildOpusTagsPacket(tags)
	if err != nil {
		fatal(err)
	}

	if err := pw.WritePacket(headPkt, 0, true, false); err != nil {
		fatal(err)
	}
	if err := pw.WritePacket(tagsPkt, 0, false, false); err != nil {
		fatal(err)
	}

	decodedLenBytes := dec.Length()
	var totalFramesEstimate int64
	if decodedLenBytes > 0 {
		inFrames := decodedLenBytes / int64(channels*2)
		if inFrames > 0 {
			outFrames := (inFrames*int64(opusSampleRate) + int64(inSampleRate) - 1) / int64(inSampleRate)
			totalFramesEstimate = (outFrames + int64(frameSize48k) - 1) / int64(frameSize48k)
		}
	}
	progress := newProgressReporter(decodedLenBytes, totalFramesEstimate)

	resampler := newLinearResampler(dec, inSampleRate, opusSampleRate, channels)

	pcm := make([]int16, frameSize48k*channels)
	packet := make([]byte, 4000)
	var totalSamplesPerCh48k uint64
	var framesDone int64

	for {
		framesActual, done, err := resampler.Fill(pcm, frameSize48k)
		if err != nil {
			fatal(err)
		}
		if framesActual == 0 && done {
			break
		}

		nBytes, err := enc.Encode(pcm, frameSize48k, packet)
		if err != nil {
			fatal(err)
		}

		framesDone++
		progress.Print(framesDone, resampler.BytesRead(), false)

		totalSamplesPerCh48k += uint64(framesActual)
		granule := uint64(head.PreSkip) + totalSamplesPerCh48k

		if err := pw.WritePacket(packet[:nBytes], granule, false, done); err != nil {
			fatal(err)
		}
		if done {
			break
		}
	}

	progress.Print(framesDone, resampler.BytesRead(), true)
	fmt.Fprintln(os.Stderr)
	fmt.Printf("Wrote %s\n", outPath)

	if err := pw.Flush(); err != nil {
		fatal(err)
	}
}

func defaultOutPath(inPath string) string {
	base := strings.TrimSuffix(inPath, filepath.Ext(inPath))
	if base == "" {
		return inPath + ".opus"
	}
	return base + ".opus"
}

func randomSerial() uint32 {
	var b [4]byte
	if _, err := rand.Read(b[:]); err == nil {
		return binary.LittleEndian.Uint32(b[:])
	}
	return uint32(time.Now().UnixNano())
}

func fatal(err error) {
	if err == nil {
		return
	}
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}
