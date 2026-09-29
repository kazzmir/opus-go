//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
)

func TestEntropyEncBitsAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(1203))
	for used := int32(0); used <= 32; used++ {
		for width := uint32(1); width <= 25; width++ {
			for _, capacity := range []uint32{0, 1, 3, 8} {
				g := make([]byte, capacity+2)
				for i := range g {
					g[i] = 77
				}
				c := slices.Clone(g)
				mask := uint32(0xffffffff)
				if used < 32 {
					mask = (uint32(1) << used) - 1
				}
				ge := opuscc.OpusT_ec_enc{Fbuf: &g[1], Fstorage: capacity, Fend_window: r.Uint32() & mask, Fnend_bits: used, Fnbits_total: 123, Frng: 1 << 31, Frem: -1}
				ce := ge
				value := r.Uint32() & ((uint32(1) << width) - 1)
				opuscc.Opus_ec_enc_bits(nil, &ge, value, width)
				nativeEncoderStep(&ce, c[1:], 9, value, width)
				ce.Fbuf = ge.Fbuf
				if ge != ce || !slices.Equal(g, c) {
					t.Fatal(used, width, capacity, ge, ce)
				}
			}
		}
	}
}
