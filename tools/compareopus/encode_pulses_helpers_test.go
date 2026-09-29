//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
	"unsafe"
)

func TestEncodePulsesAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(1402))
	for _, pair := range pvqTestPairs {
		n, k := pair[0], pair[1]
		for _, capacity := range []uint32{0, 1, 3, 1024} {
			g := make([]byte, capacity+2)
			for i := range g {
				g[i] = 77
			}
			c := slices.Clone(g)
			var ge opuscc.OpusT_ec_enc
			opuscc.Opus_ec_enc_init(nil, &ge, &g[1], capacity)
			ce := ge
			for trial := 0; trial < 30; trial++ {
				y := make([]int32, n)
				for p := int32(0); p < k; p++ {
					pos := r.Intn(int(n))
					sign := int32(1)
					if y[pos] < 0 || y[pos] == 0 && r.Intn(2) == 0 {
						sign = -1
					}
					y[pos] += sign
				}
				original := slices.Clone(y)
				opuscc.Opus_encode_pulses(nil, &y[0], n, k, &ge)
				nativeEncoderStepPointer(&ce, c[1:], 15, 0, uint32(n), uint32(k), unsafe.Pointer(&y[0]))
				ce.Fbuf = ge.Fbuf
				if ge != ce || !slices.Equal(g, c) || !slices.Equal(y, original) {
					t.Fatal(n, k, capacity, trial, ge, ce)
				}
			}
		}
	}
}
