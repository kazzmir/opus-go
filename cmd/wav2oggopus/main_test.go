package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"math"
	"testing"
	"unsafe"

	"github.com/kazzmir/opus-go/ogg"
	"github.com/kazzmir/opus-go/opus"
)

const testHz = 440.0

func sineWAV(rate, channels, frames int) []byte {
	data := make([]byte, frames*channels*2)
	for i := range frames {
		v := int16(16000 * math.Sin(2*math.Pi*testHz*float64(i)/float64(rate)))
		for c := range channels {
			binary.LittleEndian.PutUint16(data[(i*channels+c)*2:], uint16(v))
		}
	}
	h := make([]byte, 44)
	copy(h, "RIFF")
	binary.LittleEndian.PutUint32(h[4:], uint32(36+len(data)))
	copy(h[8:], "WAVEfmt ")
	binary.LittleEndian.PutUint32(h[16:], 16)
	binary.LittleEndian.PutUint16(h[20:], 1)
	binary.LittleEndian.PutUint16(h[22:], uint16(channels))
	binary.LittleEndian.PutUint32(h[24:], uint32(rate))
	binary.LittleEndian.PutUint32(h[28:], uint32(rate*channels*2))
	binary.LittleEndian.PutUint16(h[32:], uint16(channels*2))
	binary.LittleEndian.PutUint16(h[34:], 16)
	copy(h[36:], "data")
	binary.LittleEndian.PutUint32(h[40:], uint32(len(data)))
	return append(h, data...)
}

// played is an encoded stream as a player sees it: decoded at 48 kHz,
// pre-skip dropped, end-trimmed to the final granule.
type played struct {
	samples  []int16 // interleaved
	channels int
	pages    int
	eos      bool
}

func play(t *testing.T, data []byte) played {
	t.Helper()
	r, err := ogg.NewOpusReader(bytes.NewReader(data))
	if err != nil {
		t.Fatalf("read headers: %v", err)
	}
	r.SetVerifyCRC(true)
	dec, err := opus.NewDecoderFromHead(r.Head)
	if err != nil {
		t.Fatal(err)
	}
	defer dec.Close()
	ch := int(r.Head.Channels)
	var pcm []int16
	var last *ogg.OpusAudioPacket
	for {
		p, err := r.ReadAudioPacket()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			t.Fatalf("read packet: %v", err)
		}
		// Every page's granule must be reachable by the packets through it.
		if p.GranuleValid && !p.EOS && p.GranulePos > uint64(len(pcm)/ch)+5760 {
			t.Fatalf("page granule %d is ahead of the audio decodable through it", p.GranulePos)
		}
		out, _, err := dec.DecodePacket(p, nil)
		if err != nil {
			t.Fatalf("decode: %v", err)
		}
		pcm = append(pcm, out...)
		last = p
	}
	if last == nil {
		t.Fatal("no audio packets")
	}
	pre, end := int(r.Head.PreSkip), int(last.GranulePos)
	if end > len(pcm)/ch {
		t.Fatalf("final granule %d is past the %d samples the stream decodes to", end, len(pcm)/ch)
	}
	return played{
		samples:  pcm[pre*ch : end*ch],
		channels: ch,
		pages:    bytes.Count(data, []byte("OggS")),
		eos:      last.EOS,
	}
}

func TestEncode(t *testing.T) {
	if unsafe.Sizeof(uintptr(0)) == 4 {
		// The transpiled encoder aborts inside celt's FFT on 32-bit
		// targets - pre-existing, and unrelated to this command.
		t.Skip("opusccenc aborts on 32-bit targets (pre-existing)")
	}
	for _, tc := range []struct {
		rate, channels, frames, frameMS int
	}{
		{48000, 1, 48000, 20},       // exact multiple of the frame size (no EOS page before)
		{48000, 1, 48800, 20},       // final frame padded less than the pre-skip
		{48000, 1, 48100, 20},       // final frame padded more than the pre-skip
		{48000, 2, 48000 * 2, 20},   // stereo
		{48000, 1, 48000 + 7, 60},   // long frames
		{24000, 1, 24000, 20},       // native, below 48 kHz
		{16000, 1, 16000 + 7, 10},   // native, other frame size
		{44100, 2, 44100 * 2, 20},   // resampled
		{22050, 1, 22050 + 100, 60}, // resampled, long frames
	} {
		var out bytes.Buffer
		err := encode(bytes.NewReader(sineWAV(tc.rate, tc.channels, tc.frames)), &out, options{
			bitrate:     64000,
			vbr:         true,
			complexity:  10,
			application: opus.ApplicationAudio,
			frameMS:     tc.frameMS,
			vendor:      "test",
			serial:      1,
		})
		if err != nil {
			t.Fatalf("%+v: encode: %v", tc, err)
		}
		p := play(t, out.Bytes())

		want := int(math.Ceil(float64(tc.frames) * 48000 / float64(tc.rate)))
		if got := len(p.samples) / p.channels; got != want {
			t.Errorf("%+v: plays %d samples, want %d", tc, got, want)
		}
		if !p.eos {
			t.Errorf("%+v: last page isn't marked EOS", tc)
		}
		seconds := float64(tc.frames) / float64(tc.rate)
		if max := int(math.Ceil(seconds)) + 3; p.pages > max {
			t.Errorf("%+v: %d pages, want at most %d (about one per second plus headers)", tc, p.pages, max)
		}
		// The audio's own last 5 ms must survive, not be lost to the
		// encoder's lookahead.
		var peak int16
		for _, v := range p.samples[len(p.samples)-240*p.channels:] {
			peak = max(peak, v)
		}
		if peak < 8000 {
			t.Errorf("%+v: last 5 ms peak %d - the end of the audio was lost", tc, peak)
		}
	}
}
