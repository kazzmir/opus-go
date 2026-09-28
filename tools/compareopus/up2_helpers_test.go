//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
)

func TestUp2AgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, n := range []int{1, 2, 17, 120, 480} {
		for trial := 0; trial < 100; trial++ {
			input := make([]int16, n)
			for i := range input {
				input[i] = int16(rng.Uint32())
				if trial == 0 {
					input[i] = -32768
				}
				if trial == 1 {
					input[i] = 32767
				}
			}
			var gs [6]int32
			for i := range gs {
				gs[i] = rng.Int31n(1<<27) - (1 << 26)
			}
			cs := gs
			g, c := make([]int16, 2*n), make([]int16, 2*n)
			opuscc.Opus_silk_resampler_private_up2_HQ(nil, &gs, &g[0], &input[0], int32(n))
			nativeUp2(&cs, c, input)
			if !slices.Equal(g, c) || gs != cs {
				t.Fatalf("n=%d trial=%d: output matches=%v Go state=%v C state=%v", n, trial, slices.Equal(g, c), gs, cs)
			}
		}
	}
}
