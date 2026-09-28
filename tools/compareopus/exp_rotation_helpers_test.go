//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"math/rand"
	"slices"
	"testing"
)

func TestExpRotationAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, n := range []int{8, 16, 64, 128} {
		for _, stride := range []int32{1, 2, 4} {
			for spread := int32(0); spread <= 3; spread++ {
				for _, dir := range []int32{-1, 1} {
					for _, k := range []int32{1, 3, int32(n / 2)} {
						g := make([]float32, n)
						for i := range g {
							g[i] = float32(rng.NormFloat64())
						}
						c := slices.Clone(g)
						opuscc.Opus_exp_rotation(nil, &g[0], int32(n), dir, stride, k, spread)
						nativeExpRotation(c, dir, stride, k, spread)
						for i := range g {
							if math.Float32bits(g[i]) != math.Float32bits(c[i]) {
								t.Fatalf("n=%d stride=%d spread=%d dir=%d K=%d index=%d Go=%g C=%g", n, stride, spread, dir, k, i, g[i], c[i])
							}
						}
					}
				}
			}
		}
	}
}
