//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"math/rand"
	"testing"
)

func TestLPCAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, p := range []int{1, 4, 16, 24} {
		for trial := 0; trial < 100; trial++ {
			signal := make([]float32, 128)
			for i := range signal {
				signal[i] = float32(rng.NormFloat64())
			}
			ac := make([]float32, p+1)
			for lag := range ac {
				for i := lag; i < len(signal); i++ {
					ac[lag] += float32(signal[i] * signal[i-lag])
				}
			}
			if trial == 0 {
				clear(ac)
			}
			if trial == 1 {
				for i := range ac {
					ac[i] = 1
				}
			}
			g, c := make([]float32, p), make([]float32, p)
			opuscc.Opus__celt_lpc(nil, &g[0], &ac[0], int32(p))
			nativeLPC(c, ac)
			for i := range g {
				if math.Float32bits(g[i]) != math.Float32bits(c[i]) {
					t.Fatalf("p=%d trial=%d tap=%d Go=%g C=%g", p, trial, i, g[i], c[i])
				}
			}
		}
	}
}
