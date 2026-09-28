//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"math/rand"
	"runtime"
	"slices"
	"testing"
)

func TestDenormaliseAgainstC(t *testing.T) {
	bands := []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 10, 12, 14, 16, 20, 24, 28, 34, 40, 48, 60, 78, 100}
	rng := rand.New(rand.NewSource(3604))
	levels := []float32{-100, -50.01, -50, -49.9, -9, 0, 0.1, 0.9999, 31.9, 32, 32.01, 100}
	for _, M := range []int32{1, 2, 4, 8} {
		for _, down := range []int32{1, 2, 3, 4, 6} {
			for trial := 0; trial < 60; trial++ {
				start, end := int32(trial%18), int32(18+trial%4)
				silence := int32(0)
				if trial%7 == 0 {
					silence = 1
				}
				x := make([]float32, M*100)
				for i := range x {
					x[i] = float32(rng.Intn(257)-128) / 128
				}
				x[0] = math.Float32frombits(1 << 31)
				energy := make([]float32, 21)
				for i := range energy {
					energy[i] = levels[(i+trial)%len(levels)] - opuscc.Opus_eMeans[i]
				}
				xb, eb := slices.Clone(x), slices.Clone(energy)
				g := make([]float32, M*120+2)
				for i := range g {
					g[i] = 123
				}
				c := slices.Clone(g)
				opuscc.Opus_denormalise_bands(nil, &bands[0], 120, &x[0], &g[1], &energy[0], start, end, M, down, silence)
				nativeDenormalise(bands, 120, x, c[1:len(c)-1], energy, start, end, M, down, silence)
				for i := range g {
					if math.Float32bits(g[i]) != math.Float32bits(c[i]) {
						t.Fatalf("M=%d down=%d trial=%d i=%d Go=%g C=%g", M, down, trial, i, g[i], c[i])
					}
				}
				if !slices.Equal(x, xb) || !slices.Equal(energy, eb) || g[0] != 123 || g[len(g)-1] != 123 {
					t.Fatal("input or guard changed")
				}
			}
		}
	}
	runtime.KeepAlive(bands)
}
