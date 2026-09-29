//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
)

func TestEntropyEncICDF16AgainstC(t *testing.T) {
	for _, table := range [][]uint16{{0}, {255, 1, 0}, {240, 180, 100, 0}, {32767, 16000, 1, 0}, {60000, 40000, 20000, 0}} {
		ftb := uint32(16)
		if table[0] < 256 {
			ftb = 8
		} else if table[0] < 32768 {
			ftb = 15
		}
		for _, capacity := range []int{0, 1, 3, 64} {
			g := make([]byte, capacity+2)
			for i := range g {
				g[i] = 77
			}
			c := slices.Clone(g)
			var ge opuscc.OpusT_ec_enc
			opuscc.Opus_ec_enc_init(nil, &ge, &g[1], uint32(capacity))
			ce := ge
			original := slices.Clone(table)
			for step := 0; step < 300; step++ {
				s := int32(step % len(table))
				opuscc.Opus_ec_enc_icdf16(nil, &ge, s, &table[0], ftb)
				nativeEncoderStepTable(&ce, c[1:], 8, uint32(s), ftb, 0, table)
				ce.Fbuf = ge.Fbuf
				if ge != ce || !slices.Equal(g, c) || !slices.Equal(table, original) {
					t.Fatal(table, capacity, step, ge, ce)
				}
			}
		}
	}
}
