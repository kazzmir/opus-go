//go:build compareopus && cgo

package main

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/kazzmir/opus-go/opuscc"
)

func TestDown2AgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, n := range []int{2, 3, 16, 321} {
		for trial := 0; trial < 100; trial++ {
			g := make([]int16, n)
			for i := range g {
				g[i] = int16(rng.Uint32())
			}
			c := slices.Clone(g)
			gs := [2]int32{rng.Int31n(1<<27) - (1 << 26), rng.Int31n(1<<27) - (1 << 26)}
			cs := gs
			opuscc.Opus_silk_resampler_down2(nil, &gs, &g[0], &g[0], int32(n))
			nativeDown2(&cs, c, c)
			if !slices.Equal(g, c) || gs != cs {
				t.Fatalf("n=%d trial=%d: Go=%v C=%v Go state=%v C state=%v", n, trial, g, c, gs, cs)
			}
		}
	}
}
