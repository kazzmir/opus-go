//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
)

func TestEntropyNormalizeAgainstC(t *testing.T) {
	for _, capacity := range []uint32{0, 1, 3, 16} {
		for _, rng := range []uint32{1, 255, 256, 65535, 65536, 1<<23 - 1, 1 << 23, 1<<23 + 1, 1 << 31} {
			for _, value := range []uint32{0, 0x7f800000, 0x7fffffff, 0x80000000, 0xffffffff} {
				for _, rem := range []int32{-1, 0, 254} {
					g := make([]byte, capacity+2)
					for i := range g {
						g[i] = 77
					}
					c := slices.Clone(g)
					ge := opuscc.OpusT_ec_enc{Fbuf: &g[1], Fstorage: capacity, Frem: rem, Fext: 2, Frng: rng, Fval: value, Fnbits_total: 33}
					ce := ge
					opuscc.CompareEntropyNormalize(&ge)
					nativeEncoderStep(&ce, c[1:], 4, 0, 0)
					ce.Fbuf = ge.Fbuf
					if ge != ce || !slices.Equal(g, c) {
						t.Fatal(capacity, rng, value, rem, ge, ce)
					}
				}
			}
		}
	}
}
