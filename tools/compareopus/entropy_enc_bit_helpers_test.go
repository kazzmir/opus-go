//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
)

func TestEntropyEncBitAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(1201))
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
			logp := []uint32{1, 2, 3, 8, 15}[step%5]
			value := []int32{0, 1, -1, 7}[r.Intn(4)]
			opuscc.Opus_ec_enc_bit_logp(nil, &ge, value, logp)
			nativeEncoderStep(&ce, c[1:], 7, uint32(value), logp)
			ce.Fbuf = ge.Fbuf
			if ge != ce || !slices.Equal(g, c) {
				t.Fatal(capacity, step, value, logp, ge, ce)
			}
		}
	}
}
