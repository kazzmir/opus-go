//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
)

func TestEntropyShrinkAgainstC(t *testing.T) {
	for _, cap := range []uint32{0, 1, 2, 8, 32} {
		for size := uint32(0); size <= cap; size++ {
			for tail := uint32(0); tail <= size; tail++ {
				for _, front := range []uint32{0, size - tail} {
					g := make([]byte, cap+2)
					for i := range g {
						g[i] = byte(i*97 + 53)
					}
					c := slices.Clone(g)
					ge := opuscc.OpusT_ec_enc{Fbuf: &g[1], Fstorage: cap, Foffs: front, Fend_offs: tail, Fend_window: 0xabcdef, Fnend_bits: 7, Fnbits_total: 53, Frng: 0x87654321, Fval: 0x12345678, Fext: 3, Frem: 77, Ferror1: -1}
					ce := ge
					opuscc.Opus_ec_enc_shrink(nil, &ge, size)
					nativeEncoderStep(&ce, c[1:], 1, size, 0)
					ce.Fbuf = ge.Fbuf
					if ge != ce || !slices.Equal(g, c) {
						t.Fatal(cap, size, tail, front, ge, ce)
					}
				}
			}
		}
	}
}
