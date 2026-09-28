//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
)

func TestPitchDecodeAgainstC(t *testing.T) {
	for _, fs := range []int32{8, 12, 16} {
		for _, n := range []int{2, 4} {
			count := 34
			if n == 2 {
				count = 12
			}
			if fs == 8 {
				count = 11
				if n == 2 {
					count = 3
				}
			}
			for contour := 0; contour < count; contour++ {
				for lag := int16(0); lag <= int16(16*fs); lag++ {
					g, c := make([]int32, n), make([]int32, n)
					opuscc.Opus_silk_decode_pitch(nil, lag, int8(contour), &g[0], fs, int32(n))
					nativePitchDecode(lag, int8(contour), c, fs)
					if !slices.Equal(g, c) {
						t.Fatalf("fs=%d n=%d contour=%d lag=%d Go=%v C=%v", fs, n, contour, lag, g, c)
					}
				}
			}
		}
	}
}
