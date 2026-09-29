//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
)

func TestLaplaceEncodeAgainstC(t *testing.T) {
	for _, fs := range []uint32{1, 2, 1024, 16000, 32700, 32735} {
		for _, decay := range []int32{1, 64, 6000, 11456} {
			for _, capacity := range []uint32{0, 1, 3, 256} {
				g := make([]byte, capacity+2)
				for i := range g {
					g[i] = 77
				}
				c := slices.Clone(g)
				var ge opuscc.OpusT_ec_enc
				opuscc.Opus_ec_enc_init(nil, &ge, &g[1], capacity)
				ce := ge
				for _, v := range []int32{0, 1, -1, 2, -2, 7, -7, 16, -16, 128, -128, 16384, -16384, 2147483647, -2147483647} {
					gv := v
					opuscc.Opus_ec_laplace_encode(nil, &ge, &gv, fs, decay)
					cv := nativeEncoderStepPointer(&ce, c[1:], 13, uint32(v), fs, uint32(decay), nil)
					ce.Fbuf = ge.Fbuf
					if ge != ce || gv != cv || !slices.Equal(g, c) {
						t.Fatal(fs, decay, capacity, v, gv, cv, ge, ce)
					}
				}
			}
		}
	}
}
