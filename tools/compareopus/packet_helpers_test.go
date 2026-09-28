//go:build compareopus && cgo

package main

import (
	"testing"

	"github.com/kazzmir/opus-go/opuscc"
)

func TestPacketHelpersAgainstC(t *testing.T) {
	for toc := 0; toc < 256; toc++ {
		b := byte(toc)
		if got, want := opuscc.Opus_opus_packet_get_bandwidth(nil, &b), nativePacketBandwidth(&b); got != want {
			t.Fatalf("TOC %#x bandwidth: Go=%d C=%d", toc, got, want)
		}
		if got, want := opuscc.Opus_opus_packet_get_nb_channels(nil, &b), nativePacketChannels(&b); got != want {
			t.Fatalf("TOC %#x channels: Go=%d C=%d", toc, got, want)
		}
		for _, rate := range []int32{8000, 12000, 16000, 24000, 48000} {
			if got, want := opuscc.Opus_opus_packet_get_samples_per_frame(nil, &b, rate), nativePacketSamplesPerFrame(&b, rate); got != want {
				t.Fatalf("TOC %#x rate %d samples/frame: Go=%d C=%d", toc, rate, got, want)
			}
		}
		if got, want := opuscc.Opus_opus_packet_get_nb_frames(nil, &b, 1), nativePacketFrames(&b, 1); got != want {
			t.Fatalf("TOC %#x length 1 frames: Go=%d C=%d", toc, got, want)
		}
		for second := 0; second < 256; second++ {
			packet := [2]byte{b, byte(second)}
			if got, want := opuscc.Opus_opus_packet_get_nb_frames(nil, &packet[0], 2), nativePacketFrames(&packet[0], 2); got != want {
				t.Fatalf("packet %x frames: Go=%d C=%d", packet, got, want)
			}
		}
	}
	for _, length := range []int32{-1, 0} {
		if got, want := opuscc.Opus_opus_packet_get_nb_frames(nil, nil, length), nativePacketFrames(nil, length); got != want {
			t.Fatalf("length %d frames: Go=%d C=%d", length, got, want)
		}
	}
}
