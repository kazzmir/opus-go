package test

import (
	"io"
	"os"
	"testing"

	"github.com/kazzmir/opus-go/ogg"
	"github.com/kazzmir/opus-go/opus"
)

func TestDecodeMusic64kbpsSampleCount(t *testing.T) {
	f, err := os.Open("music_64kbps.opus")
	if err != nil {
		t.Fatalf("open fixture: %v", err)
	}
	defer f.Close()

	r, err := ogg.NewOpusReader(f)
	if err != nil {
		t.Fatalf("ogg opus reader: %v", err)
	}

	if r.Head.Channels != 2 {
		t.Fatalf("expected 2 channels, got %d", r.Head.Channels)
	}

	dec, err := opus.NewDecoderFromHead(r.Head)
	if err != nil {
		t.Fatalf("new decoder: %v", err)
	}
	defer dec.Close()

	// Max Opus frame size at 48kHz is 120ms = 5760 samples/channel.
	pcm := make([]int16, 5760*int(r.Head.Channels))

	preSkipRemaining := int(r.Head.PreSkip)
	totalInt16 := int64(0)

	for {
		pkt, err := r.ReadAudioPacket()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("read audio packet: %v", err)
		}

		nPerCh, err := dec.Decode(pkt.Data, pcm, 5760, false)
		if err != nil {
			t.Fatalf("decode: %v", err)
		}
		out := pcm[:nPerCh*dec.Channels()]

		if preSkipRemaining > 0 {
			skip := min(len(out), preSkipRemaining * dec.Channels())
			out = out[skip:]
			preSkipRemaining -= skip / dec.Channels()
		}

		totalInt16 += int64(len(out))

		if pkt.EOS {
			break
		}
	}

	// The expected total is approximate.
	//
	// Note: this is a count of int16 values (interleaved), not bytes.
	// For ~90s at 48kHz stereo, that's ~90 * 48000 * 2 = 8.64M.
	const expected = int64(8640000)
    tolerance := int64(ogg.OpusSampleRateHz * dec.Channels()) // 1 second tolerance
	if totalInt16 < expected-tolerance || totalInt16 > expected+tolerance {
		t.Fatalf("decoded int16 sample count out of range: got %d, want %d±%d", totalInt16, expected, tolerance)
	}
}
