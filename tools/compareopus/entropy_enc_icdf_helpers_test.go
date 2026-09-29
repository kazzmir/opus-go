//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
	"unsafe"
)

func TestEntropyEncICDFAgainstC(t *testing.T) {
	for _, table := range [][]byte{{0}, {255, 1, 0}, {240, 180, 100, 0}, {3, 2, 1, 0}} {
		ftb := uint32(8)
		if table[0] < 4 {
			ftb = 2
		}
		for _, capacity := range []uint32{0, 1, 3, 64} {
			g := make([]byte, capacity+2)
			for i := range g {
				g[i] = 77
			}
			c := slices.Clone(g)
			var ge opuscc.OpusT_ec_enc
			opuscc.Opus_ec_enc_init(nil, &ge, &g[1], capacity)
			ce := ge
			original := slices.Clone(table)
			for step := 0; step < 300; step++ {
				s := int32(step % len(table))
				opuscc.Opus_ec_enc_icdf(nil, &ge, s, &table[0], ftb)
				nativeEncoderStepPointer(&ce, c[1:], 12, uint32(s), ftb, 0, unsafe.Pointer(&table[0]))
				ce.Fbuf = ge.Fbuf
				if ge != ce || !slices.Equal(g, c) || !slices.Equal(table, original) {
					t.Fatal(table, capacity, step, ge, ce)
				}
			}
		}
	}
}
