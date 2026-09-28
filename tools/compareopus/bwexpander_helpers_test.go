//go:build compareopus && cgo

package main

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/kazzmir/opus-go/opuscc"
)

func TestBWExpandersAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, n := range []int{1, 2, 10, 16} {
		for _, chirp := range []int32{0, 1, 32768, 60000, 64881, 65535, 65536} {
			for trial := 0; trial < 100; trial++ {
				g16, g32 := make([]int16, n), make([]int32, n)
				for i := range g16 {
					g16[i] = int16(rng.Uint32())
					g32[i] = int32(rng.Uint32())
				}
				c16, c32 := slices.Clone(g16), slices.Clone(g32)
				opuscc.Opus_silk_bwexpander(nil, &g16[0], int32(n), chirp)
				opuscc.Opus_silk_bwexpander_32(nil, &g32[0], int32(n), chirp)
				nativeBWExpander(c16, chirp)
				nativeBWExpander32(c32, chirp)
				if !slices.Equal(g16, c16) || !slices.Equal(g32, c32) {
					t.Fatalf("n=%d chirp=%d trial=%d: Go16=%v C16=%v Go32=%v C32=%v", n, chirp, trial, g16, c16, g32, c32)
				}
			}
		}
	}
}
