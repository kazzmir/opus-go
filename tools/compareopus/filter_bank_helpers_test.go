//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
)

func TestFilterBankAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, n := range []int{2, 3, 16, 321, 960} {
		for trial := 0; trial < 100; trial++ {
			g := make([]int16, n)
			for i := range g {
				g[i] = int16(rng.Uint32())
			}
			c := slices.Clone(g)
			gh, ch := make([]int16, n/2), make([]int16, n/2)
			gs := [2]int32{rng.Int31n(1<<27) - (1 << 26), rng.Int31n(1<<27) - (1 << 26)}
			cs := gs
			opuscc.Opus_silk_ana_filt_bank_1(nil, &g[0], &gs, &g[0], &gh[0], int32(n))
			nativeFilterBank(c, &cs, c, ch)
			if !slices.Equal(g, c) || !slices.Equal(gh, ch) || gs != cs {
				t.Fatalf("n=%d trial=%d: output or state differs from C", n, trial)
			}
		}
	}
}
