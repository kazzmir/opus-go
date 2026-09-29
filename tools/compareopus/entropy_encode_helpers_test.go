//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
)

func TestEntropyEncodeAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(1103))
	for _, capacity := range []int{0, 1, 3, 64} {
		g := make([]byte, capacity+2)
		for i := range g {
			g[i] = 77
		}
		c := slices.Clone(g)
		var ge opuscc.OpusT_ec_enc
		opuscc.Opus_ec_enc_init(nil, &ge, &g[1], uint32(capacity))
		ce := ge
		for step := 0; step < 400; step++ {
			ft := []uint32{2, 13, 256, 32768, 65535}[step%5]
			fl := r.Uint32() % ft
			fh := fl + 1 + r.Uint32()%(ft-fl)
			if step%7 == 0 {
				fl = 0
			}
			if step%11 == 0 {
				fh = ft
			}
			opuscc.Opus_ec_encode(nil, &ge, fl, fh, ft)
			nativeEncoderStep(&ce, c[1:], 5, fl, fh, ft)
			ce.Fbuf = ge.Fbuf
			if ge != ce || !slices.Equal(g, c) {
				t.Fatal(capacity, step, fl, fh, ft, ge, ce)
			}
		}
	}
}
