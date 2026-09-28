//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"github.com/kazzmir/opus-go/opusccenc"
	"math/rand"
	"slices"
	"testing"
)

func TestLaroiaAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, n := range []int{2, 10, 16} {
		for trial := 0; trial < 1000; trial++ {
			input := make([]int16, n)
			for i := range input {
				input[i] = int16(rng.Intn(32768))
			}
			slices.Sort(input)
			if trial == 0 {
				clear(input)
			}
			if trial == 1 {
				for i := range input {
					input[i] = 32767
				}
			}
			for _, alias := range []bool{false, true} {
				g, c := make([]int16, n), make([]int16, n)
				gi, ci := input, input
				if alias {
					g = slices.Clone(input)
					c = slices.Clone(input)
					gi, ci = g, c
				}
				enc, encInput := make([]int16, n), input
				if alias {
					enc = slices.Clone(input)
					encInput = enc
				}
				opusccenc.Opus_silk_NLSF_VQ_weights_laroia(nil, &enc[0], &encInput[0], int32(n))
				opuscc.Opus_silk_NLSF_VQ_weights_laroia(nil, &g[0], &gi[0], int32(n))
				nativeLaroia(c, ci)
				if !slices.Equal(g, c) || !slices.Equal(enc, c) {
					t.Fatalf("n=%d trial=%d alias=%v input=%v decoder=%v encoder=%v C=%v", n, trial, alias, input, g, enc, c)
				}
			}
		}
	}
}
