//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
)

func TestEntropyEncodeBinAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(1104))
	for _, capacity := range []int{0, 1, 3, 64} {
		g := make([]byte, capacity+2)
		for i := range g {
			g[i] = 77
		}
		c := slices.Clone(g)
		var ge opuscc.OpusT_ec_enc
		opuscc.Opus_ec_enc_init(nil, &ge, &g[1], uint32(capacity))
		ce := ge
		for step := 0; step < 500; step++ {
			bits := []uint32{0, 1, 2, 4, 8, 15}[step%6]
			ft := uint32(1) << bits
			fl := r.Uint32() % ft
			fh := fl + 1 + r.Uint32()%(ft-fl)
			if step%7 == 0 {
				fl = 0
			}
			if step%11 == 0 {
				fh = ft
			}
			opuscc.Opus_ec_encode_bin(nil, &ge, fl, fh, bits)
			nativeEncoderStep(&ce, c[1:], 6, fl, fh, bits)
			ce.Fbuf = ge.Fbuf
			if ge != ce || !slices.Equal(g, c) {
				t.Fatal(capacity, step, fl, fh, bits, ge, ce)
			}
		}
	}
}
