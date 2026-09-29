//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
)

func TestLaplaceEncodeP0AgainstC(t *testing.T) {
	for _, p0 := range []uint16{0, 1, 16000, 32766, 32767, 32768} {
		for _, decay := range []uint16{0, 1, 7, 8, 12000, 32767} {
			for _, capacity := range []uint32{0, 1, 3, 1024} {
				g := make([]byte, capacity+2)
				for i := range g {
					g[i] = 77
				}
				c := slices.Clone(g)
				var ge opuscc.OpusT_ec_enc
				opuscc.Opus_ec_enc_init(nil, &ge, &g[1], capacity)
				ce := ge
				for _, v := range []int32{0, 1, -1, 6, -6, 7, -7, 8, -8, 14, -14, 15, -15, 127, -127} {
					// A zero-probability symbol would leave range=0 and normalization would never end.
					if p0 == 0 && v == 0 || p0 == 32768 && v != 0 || p0 == 32767 && v < 0 {
						continue
					}
					opuscc.Opus_ec_laplace_encode_p0(nil, &ge, v, p0, decay)
					nativeEncoderStepPointer(&ce, c[1:], 14, uint32(v), uint32(p0), uint32(decay), nil)
					ce.Fbuf = ge.Fbuf
					if ge != ce || !slices.Equal(g, c) {
						t.Fatal(p0, decay, capacity, v, ge, ce)
					}
				}
			}
		}
	}
}
