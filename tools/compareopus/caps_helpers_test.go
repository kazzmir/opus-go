//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
)

func TestCapsAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(3701))
	for _, nb := range []int{1, 5, 21} {
		for trial := 0; trial < 100; trial++ {
			bands := make([]int16, nb+1)
			for i := 1; i < len(bands); i++ {
				bands[i] = bands[i-1] + int16(1+rng.Intn(100))
			}
			cache := make([]byte, 8*nb)
			rng.Read(cache)
			bb, cb := slices.Clone(bands), slices.Clone(cache)
			for lm := int32(0); lm < 4; lm++ {
				for channels := int32(1); channels <= 2; channels++ {
					g, c := make([]int32, nb), make([]int32, nb)
					opuscc.Opus_init_caps(nil, &bands[0], &cache[0], &g[0], int32(nb), lm, channels)
					nativeCaps(bands, cache, c, lm, channels)
					if !slices.Equal(g, c) || !slices.Equal(bands, bb) || !slices.Equal(cache, cb) {
						t.Fatalf("nb=%d trial=%d LM=%d C=%d Go=%v C=%v", nb, trial, lm, channels, g, c)
					}
				}
			}
		}
	}
}
