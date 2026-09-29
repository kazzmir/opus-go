//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
)

func TestEntropyDoneAgainstC(t *testing.T) {
	for _, capacity := range []uint32{0, 1, 2, 3, 8} {
		for _, rng := range []uint32{1<<23 + 1, 1 << 24, 1 << 30, 1 << 31} {
			for _, value := range []uint32{0, 1, 0x1fffffff, 0x7fffffff} {
				for _, used := range []int32{0, 1, 2, 3, 7, 8, 15, 24, 32} {
					for _, front := range []uint32{0, capacity} {
						for _, rem := range []int32{-1, 0, 255} {
							g := make([]byte, capacity+2)
							for i := range g {
								g[i] = 0xa5
							}
							c := slices.Clone(g)
							mask := uint32(0xffffffff)
							if used < 32 {
								mask = (uint32(1) << used) - 1
							}
							ge := opuscc.OpusT_ec_enc{Fbuf: &g[1], Fstorage: capacity, Foffs: front, Frng: rng, Fval: value, Fext: 2, Frem: rem, Fend_window: 0xb7ab91e5 & mask, Fnend_bits: used, Fnbits_total: 71}
							ce := ge
							opuscc.Opus_ec_enc_done(nil, &ge)
							nativeEncoderStep(&ce, c[1:], 10, 0, 0)
							ce.Fbuf = ge.Fbuf
							if ge != ce || !slices.Equal(g, c) {
								t.Fatal(capacity, rng, value, used, front, rem, ge, ce)
							}
						}
					}
				}
			}
		}
	}
}
