//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"testing"
	"unsafe"
)

func TestPacketDurationAgainstC(t *testing.T) {
	for _, rate := range []int32{8000, 12000, 16000, 24000, 48000} {
		for _, n := range []int32{-1, 0} {
			if g, c := opuscc.Opus_opus_packet_get_nb_samples(nil, nil, n, rate), nativePacketDuration(nil, n, rate); g != c {
				t.Fatal(rate, n, g, c)
			}
		}
		for toc := 0; toc < 256; toc++ {
			for count := 0; count < 256; count++ {
				packet := []byte{byte(toc), byte(count)}
				for _, n := range []int32{1, 2} {
					g := opuscc.Opus_opus_packet_get_nb_samples(nil, unsafe.SliceData(packet), n, rate)
					c := nativePacketDuration(packet, n, rate)
					if g != c {
						t.Fatalf("rate=%d toc=%x count=%x n=%d Go=%d C=%d", rate, toc, count, n, g, c)
					}
				}
			}
		}
	}
}
