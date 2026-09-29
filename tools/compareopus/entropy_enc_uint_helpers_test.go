//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
)

func TestEntropyEncUintAgainstC(t *testing.T) {
	totals := []uint32{2, 3, 255, 256, 257, 511, 512, 513, 65535, 65536, 65537, 1 << 31, 0xffffffff}
	for _, capacity := range []uint32{0, 1, 3, 256} {
		g := make([]byte, capacity+2)
		for i := range g {
			g[i] = 77
		}
		c := slices.Clone(g)
		var ge opuscc.OpusT_ec_enc
		opuscc.Opus_ec_enc_init(nil, &ge, &g[1], capacity)
		ce := ge
		for repeat := 0; repeat < 5; repeat++ {
			for _, total := range totals {
				for _, v := range []uint32{0, 1, total / 2, total - 1} {
					opuscc.Opus_ec_enc_uint(nil, &ge, v, total)
					nativeEncoderStep(&ce, c[1:], 11, v, total)
					ce.Fbuf = ge.Fbuf
					if ge != ce || !slices.Equal(g, c) {
						t.Fatal(capacity, repeat, total, v, ge, ce)
					}
				}
			}
		}
		opuscc.Opus_ec_enc_done(nil, &ge)
		nativeEncoderStep(&ce, c[1:], 10, 0, 0)
		ce.Fbuf = ge.Fbuf
		if ge != ce || !slices.Equal(g, c) {
			t.Fatal("final", ge, ce)
		}
	}
}
