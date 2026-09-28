//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
)

func TestNLSFVQAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, order := range []int{2, 10, 16} {
		for _, k := range []int{1, 2, 32} {
			for trial := 0; trial < 100; trial++ {
				input := make([]int16, order)
				cb := make([]uint8, k*order)
				weights := make([]int16, len(cb))
				for i := range input {
					input[i] = int16(rng.Uint32())
				}
				for i := range cb {
					cb[i] = uint8(rng.Uint32())
					weights[i] = int16(rng.Intn(1024) + 1)
				}
				g, c := make([]int32, k), make([]int32, k)
				opuscc.Opus_silk_NLSF_VQ(nil, &g[0], &input[0], &cb[0], &weights[0], int32(k), int32(order))
				nativeNLSFVQ(c, input, cb, weights)
				if !slices.Equal(g, c) {
					t.Fatalf("order=%d K=%d trial=%d Go=%v C=%v", order, k, trial, g, c)
				}
			}
		}
	}
}
