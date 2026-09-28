//go:build compareopus && cgo

package main

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/kazzmir/opus-go/opuscc"
	"github.com/kazzmir/opus-go/opusccenc"
)

func TestInterpolateAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, n := range []int{1, 10, 16} {
		for factor := int32(0); factor <= 4; factor++ {
			for trial := 0; trial < 100; trial++ {
				a, b := make([]int16, n), make([]int16, n)
				for i := range a {
					a[i], b[i] = int16(rng.Uint32()), int16(rng.Uint32())
				}
				want, dec, enc := make([]int16, n), make([]int16, n), make([]int16, n)
				nativeInterpolate(want, a, b, factor)
				opuscc.Opus_silk_interpolate(nil, &dec[0], &a[0], &b[0], factor, int32(n))
				opusccenc.Opus_silk_interpolate(nil, &enc[0], &a[0], &b[0], factor, int32(n))
				if !slices.Equal(want, dec) || !slices.Equal(want, enc) {
					t.Fatalf("n=%d factor=%d trial=%d: C=%v decoder=%v encoder=%v", n, factor, trial, want, dec, enc)
				}
			}
		}
	}
}
