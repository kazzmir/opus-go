//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"testing"
	"unsafe"
)

func TestDecoderDurationAgainstC(t *testing.T) {
	for _, rate := range []int32{8000, 12000, 16000, 24000, 48000} {
		d := opuscc.OpusT_OpusDecoder{FFs: rate}
		for _, n := range []int32{-1, 0} {
			if g, c := opuscc.Opus_opus_decoder_get_nb_samples(nil, &d, nil, n), nativeDecoderDuration(nil, n, rate); g != c {
				t.Fatal(g, c)
			}
		}
		for toc := 0; toc < 256; toc++ {
			for _, count := range []byte{0, 1, 2, 3, 6, 24, 48, 49, 63, 255} {
				p := []byte{byte(toc), count}
				for _, n := range []int32{1, 2} {
					g := opuscc.Opus_opus_decoder_get_nb_samples(nil, &d, unsafe.SliceData(p), n)
					c := nativeDecoderDuration(p, n, rate)
					if g != c {
						t.Fatal(rate, toc, count, n, g, c)
					}
				}
			}
		}
	}
}
