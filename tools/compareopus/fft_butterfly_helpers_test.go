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

func TestFFTButterfly3AgainstC(t *testing.T) {
	compareButterflyStages(t, 3, opuscc.CompareFFTButterfly3)
}

func TestFFTButterfly4AgainstC(t *testing.T) {
	compareButterflyStages(t, 4, opuscc.CompareFFTButterfly4)
}

func compareButterflyStages(t *testing.T, radix int32, goButterfly func(*opuscc.OpusT_kiss_fft_cpx, uint64, *opuscc.OpusT_kiss_twiddle_cpx, int32, int32, int32)) {
	t.Helper()
	for _, m := range []int32{1, 4, 8, 16} {
		for _, N := range []int32{0, 1, 3} {
			for _, stride := range []uint64{0, 1, 3} {
				for _, gap := range []int32{0, 5} {
					for variant := 0; variant < 3; variant++ {
						mm := radix*m + gap
						extent := N*mm + radix*m + 2
						g := make([]opuscc.OpusT_kiss_fft_cpx, extent)
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
						tw := butterflyTwiddles(int(4*uint64(m)*stride + 1))
						before := slices.Clone(tw)
						goButterfly(&g[1], stride, &tw[0], m, N, mm)
						nativeFFTButterfly(radix, c[1:], tw, stride, m, N, mm)
						if !sameComplexBits(g, c) || !slices.Equal(tw, before) {
							t.Fatal(radix, m, N, stride, gap, variant)
						}
					}
				}
			}
		}
	}
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
