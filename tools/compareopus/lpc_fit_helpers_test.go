//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
)

func TestLPCFitAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, n := range []int{1, 10, 16, 24} {
		for _, shift := range []int32{1, 4, 5, 8, 12, 16} {
			for trial := 0; trial < 200; trial++ {
				g := make([]int32, n)
				for i := range g {
					g[i] = int32(rng.Uint32()) >> 1
					if trial%3 == 0 {
						g[i] >>= 16
					}
				}
				if trial == 1 {
					clear(g)
					g[n-1] = 2000000000
				}
				if trial == 2 {
					clear(g)
					g[n-1] = -2000000000
				}
				c := slices.Clone(g)
				goOut, cOut := make([]int16, n), make([]int16, n)
				opuscc.Opus_silk_LPC_fit(nil, &goOut[0], &g[0], 12, 12+shift, int32(n))
				nativeLPCFit(cOut, c, 12, 12+shift)
				if !slices.Equal(g, c) || !slices.Equal(goOut, cOut) {
					t.Fatalf("n=%d shift=%d trial=%d Go=%v/%v C=%v/%v", n, shift, trial, goOut, g, cOut, c)
				}
			}
		}
	}
}
