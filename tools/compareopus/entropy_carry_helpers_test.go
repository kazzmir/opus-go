//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
)

func TestEntropyCarryAgainstC(t *testing.T) {
	for _, capacity := range []uint32{0, 1, 3, 16} {
		for _, rem := range []int32{-1, 0, 127, 254, 255} {
			for _, ext := range []uint32{0, 1, 2, 17} {
				for _, value := range []int32{0, 1, 127, 254, 255, 256, 257, 510, 511} {
					g := make([]byte, capacity+2)
					for i := range g {
						g[i] = 77
					}
					c := slices.Clone(g)
					ge := opuscc.OpusT_ec_enc{Fbuf: &g[1], Fstorage: capacity, Frem: rem, Fext: ext, Frng: 123, Fval: 456}
					ce := ge
					opuscc.CompareEntropyCarry(&ge, value)
					nativeEncoderStep(&ce, c[1:], 3, uint32(value), 0)
					ce.Fbuf = ge.Fbuf
					if ge != ce || !slices.Equal(g, c) {
						t.Fatal(capacity, rem, ext, value, ge, ce)
					}
				}
			}
		}
	}
}
