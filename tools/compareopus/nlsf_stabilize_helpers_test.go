//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
)

func TestNLSFStabilizeAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, n := range []int{2, 10, 16} {
		for trial := 0; trial < 2000; trial++ {
			g := make([]int16, n)
			delta := make([]int16, n+1)
			for i := range g {
				g[i] = int16(rng.Uint32())
				if trial == 0 {
					g[i] = 32767
				}
				if trial == 1 {
					g[i] = int16(32767 - i*1000)
				}
				if trial == 2 {
					g[i] = 0
				}
			}
			for i := range delta {
				delta[i] = int16(1 + rng.Intn(1000))
			}
			c := slices.Clone(g)
			before := slices.Clone(delta)
			opuscc.Opus_silk_NLSF_stabilize(nil, &g[0], &delta[0], int32(n))
			nativeNLSFStabilize(c, delta)
			if !slices.Equal(g, c) || !slices.Equal(delta, before) {
				t.Fatalf("n=%d trial=%d Go=%v C=%v", n, trial, g, c)
			}
		}
	}
}
