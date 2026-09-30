//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"slices"
	"testing"
)

func sameComplexBits(a, b []opuscc.OpusT_kiss_fft_cpx) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Float32bits(a[i].Fr) != math.Float32bits(b[i].Fr) || math.Float32bits(a[i].Fi) != math.Float32bits(b[i].Fi) {
			return false
		}
	}
	return true
}

func TestFFTButterfly2AgainstC(t *testing.T) {
	for _, n := range []int32{0, 1, 2, 17} {
		for variant := 0; variant < 3; variant++ {
			g := make([]opuscc.OpusT_kiss_fft_cpx, 8*n+2)
			for i := range g {
				r := float32(i-13) / 7
				if variant == 1 {
					r = math.Float32frombits(uint32(i)*123457 + 1)
				}
				if variant == 2 {
					r = math.Float32frombits(0x80000000)
				}
				g[i] = opuscc.OpusT_kiss_fft_cpx{Fr: r, Fi: -r}
			}
			c := slices.Clone(g)
			opuscc.CompareFFTButterfly2(&g[1], 4, n)
			nativeFFTButterfly(2, c[1:], nil, 0, 4, n, 0)
			if !sameComplexBits(g, c) {
				t.Fatal(n, variant)
			}
		}
	}
}

// Shared fixtures used by the remaining radix tests.
func butterflyTwiddles(n int) []opuscc.OpusT_kiss_twiddle_cpx {
	tw := make([]opuscc.OpusT_kiss_twiddle_cpx, n)
	for i := range tw {
		phase := -2 * math.Pi * float64(i) / float64(n)
		tw[i] = opuscc.OpusT_kiss_twiddle_cpx{Fr: float32(math.Cos(phase)), Fi: float32(math.Sin(phase))}
	}
	return tw
}

