package main

import (
	"bufio"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
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
		channels   = 2
		frameMS    = 20
		bitrate    = 64000
		vbr        = true
		complexity = 10
	)

	sampleRate := dec.SampleRate()
	if sampleRate <= 0 {
		fatal(fmt.Errorf("mp3: invalid sample rate: %d", sampleRate))
	}

	frameSize := (sampleRate * frameMS) / 1000
	if frameSize <= 0 {
		fatal(fmt.Errorf("mp3: computed invalid frame size: %d", frameSize))
	}

	enc, err := opus.NewEncoder(sampleRate, channels, opus.ApplicationAudio)
	if err != nil {
		fatal(err)
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
		InputSampleRate:      uint32(sampleRate),
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
		totalPCMFrames := decodedLenBytes / int64(channels*2)
		if totalPCMFrames > 0 {
			totalFramesEstimate = (totalPCMFrames + int64(frameSize) - 1) / int64(frameSize)
		}
	}
	progress := newProgressReporter(decodedLenBytes, totalFramesEstimate)

	pcm := make([]int16, frameSize*channels)
	raw := make([]byte, len(pcm)*2)
	packet := make([]byte, 4000)
	var totalSamplesPerCh48k uint64
	var rem48k uint64
	var bytesRead int64
	var framesDone int64

	for {
		n, rerr := io.ReadFull(dec, raw)
		if rerr != nil && !errors.Is(rerr, io.EOF) && !errors.Is(rerr, io.ErrUnexpectedEOF) {
			fatal(rerr)
		}
		if n == 0 {
			break
		}

		bytesPerFrame := channels * 2
		if n%bytesPerFrame != 0 {
			n -= n % bytesPerFrame
		}
		framesRead := n / bytesPerFrame
		bytesRead += int64(n)

		isLast := errors.Is(rerr, io.EOF) || errors.Is(rerr, io.ErrUnexpectedEOF) || framesRead < frameSize

		for i := n; i < len(raw); i++ {
			raw[i] = 0
		}
		for i := 0; i < len(raw); i += 2 {
			pcm[i/2] = int16(binary.LittleEndian.Uint16(raw[i : i+2]))
		}

		nBytes, err := enc.Encode(pcm, frameSize, packet)
		if err != nil {
			fatal(err)
		}

		framesDone++
		progress.Print(framesDone, bytesRead, false)

		num := uint64(framesRead)*uint64(ogg.OpusSampleRateHz) + rem48k
		inc48k := num / uint64(sampleRate)
		rem48k = num % uint64(sampleRate)
		totalSamplesPerCh48k += inc48k
		granule := uint64(head.PreSkip) + totalSamplesPerCh48k

		if err := pw.WritePacket(packet[:nBytes], granule, false, isLast); err != nil {
			fatal(err)
		}
		if isLast {
			break
		}
	}

	progress.Print(framesDone, bytesRead, true)
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
