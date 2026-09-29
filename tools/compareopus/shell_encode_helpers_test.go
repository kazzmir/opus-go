//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
	"unsafe"
)

func TestShellEncodeAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(1403))
	for _, capacity := range []uint32{0, 1, 3, 1024} {
		g := make([]byte, capacity+2)
		for i := range g {
			g[i] = 77
		}
		c := slices.Clone(g)
		var ge opuscc.OpusT_ec_enc
		opuscc.Opus_ec_enc_init(nil, &ge, &g[1], capacity)
		ce := ge
		for total := 0; total <= 16; total++ {
			for trial := 0; trial < 50; trial++ {
				var p [16]int32
				for i := 0; i < total; i++ {
					pos := trial % 16
					if trial >= 16 {
						pos = r.Intn(16)
					}
					p[pos]++
				}
				original := p
				opuscc.Opus_silk_shell_encoder(nil, &ge, &p)
				nativeEncoderStepPointer(&ce, c[1:], 16, 0, 0, 0, unsafe.Pointer(&p[0]))
				ce.Fbuf = ge.Fbuf
				if ge != ce || !slices.Equal(g, c) || p != original {
					t.Fatal(capacity, total, trial, ge, ce)
				}
			}
		}
	}
}
