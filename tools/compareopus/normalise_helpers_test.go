//go:build compareopus && cgo

package main

import (
	"math"
	"math/rand"
	"slices"
	"testing"

	"github.com/kazzmir/opus-go/opuscc"
)

func TestRenormaliseAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	var maxRelativeError float64
	for _, n := range []int{1, 2, 16, 64, 120} {
		for _, gain := range []float32{0, 0.5, 1, 2} {
			for trial := 0; trial < 100; trial++ {
				g := make([]float32, n)
				for i := range g {
					g[i] = float32(rng.NormFloat64())
				}
				if trial == 0 {
					clear(g)
				}
				c, scalar := slices.Clone(g), slices.Clone(g)
				opuscc.Opus_renormalise_vector(nil, &g[0], int32(n), gain, 0)
				nativeRenormalise(c, gain)
				scalarCRenormalise(scalar, gain)
				for i := range g {
					if math.Float32bits(g[i]) != math.Float32bits(scalar[i]) {
						t.Fatalf("n=%d gain=%g trial=%d i=%d: Go=%g scalar C=%g", n, gain, trial, i, g[i], scalar[i])
					}
					// Native SSE sums in a different order. Bound that difference
					// by eight float32 machine epsilons; scalar agreement is exact.
					scale := math.Max(math.Abs(float64(g[i])), math.Abs(float64(c[i])))
					error := math.Abs(float64(g[i]) - float64(c[i]))
					if error > (8.0/(1<<23))*scale || math.IsNaN(error) {
						t.Fatalf("n=%d gain=%g trial=%d i=%d: Go=%g native C=%g", n, gain, trial, i, g[i], c[i])
					}
					if scale > 0 {
						maxRelativeError = math.Max(maxRelativeError, error/scale)
					}
				}
			}
		}
	}
	t.Logf("scalar C: exact; maximum relative error vs native C: %.9g", maxRelativeError)
}
