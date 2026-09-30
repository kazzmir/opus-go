//go:build compareopus && cgo

package main

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"slices"
	"testing"
	"unsafe"
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

func TestMiniFFTStrideAgainstC(t *testing.T) {
	compareMiniTransforms(t, 1, opuscc.CompareMiniFFTStride)
}

func TestMiniFFTWorkAgainstC(t *testing.T) { compareMiniTransforms(t, 0, opuscc.CompareMiniFFTWork) }

func compareMiniTransforms(t *testing.T, op int32, transform func(*opuscc.OpusT_mini_kiss_fft_state, *opuscc.OpusT_mini_kiss_fft_cpx, *opuscc.OpusT_mini_kiss_fft_cpx, int32)) {
	t.Helper()
	for _, n := range []int32{2, 3, 4, 5, 8, 12, 16, 20, 60, 120, 240, 480} {
		for _, inverse := range []int32{0, 1, -3} {
			for _, stride := range []int32{0, 1, 2, 5} {
				if op == 2 && stride != 1 {
					continue
				}
				st := nativeMiniFixture(n, inverse)
				stateBytes := unsafe.Slice((*byte)(unsafe.Pointer(st)), 264+8*n)
				before := slices.Clone(stateBytes)
				input := make([]opuscc.OpusT_mini_kiss_fft_cpx, (n-1)*stride+1)
				for i := range input {
					input[i] = opuscc.OpusT_mini_kiss_fft_cpx{Fr: float32(i-13) / 7, Fi: float32(19-i) / 11}
				}
				ib := slices.Clone(input)
				g := make([]opuscc.OpusT_mini_kiss_fft_cpx, n+2)
				g[0].Fr = 77
				g[n+1].Fr = 88
				c := slices.Clone(g)
				transform(st, &input[0], &g[1], stride)
				nativeMiniTransform(st, input, c[1:], stride, op)
				if !sameComplexBits(g, c) || !sameComplexBits(input, ib) || !slices.Equal(stateBytes, before) {
					t.Fatal(op, n, inverse, stride)
				}
			}
		}
	}
}

func TestMiniButterfly5AgainstC(t *testing.T) {
	compareMiniButterflies(t, 5, opuscc.CompareMiniFFTButterfly5)
}

func TestMiniButterfly3AgainstC(t *testing.T) {
	compareMiniButterflies(t, 3, opuscc.CompareMiniFFTButterfly3)
}

func TestMiniButterfly4AgainstC(t *testing.T) {
	compareMiniButterflies(t, 4, opuscc.CompareMiniFFTButterfly4)
}

func TestMiniButterfly2AgainstC(t *testing.T) {
	compareMiniButterflies(t, 2, opuscc.CompareMiniFFTButterfly2)
}

func compareMiniButterflies(t *testing.T, radix int32, transform func(*opuscc.OpusT_mini_kiss_fft_cpx, *opuscc.OpusT_mini_kiss_fft_cpx, uint64, uint64, int32)) {
	t.Helper()
	for _, m := range []uint64{1, 2, 4, 7, 16} {
		for _, stride := range []uint64{0, 1, 3} {
			for _, inverse := range []int32{0, 1, -3} {
				for variant := 0; variant < 4; variant++ {
					g := make([]opuscc.OpusT_mini_kiss_fft_cpx, uint64(radix)*m+2)
					for i := range g {
						r := float32(i-13) / 7
						if variant == 1 {
							r = math.Float32frombits(uint32(i)*123457 + 1)
						}
						if variant == 2 {
							r = math.Float32frombits(0x80000000)
						}
						if variant == 3 {
							r = []float32{1e20, -1e20, 1, -1}[i%4]
						}
						g[i] = opuscc.OpusT_mini_kiss_fft_cpx{Fr: r, Fi: -r}
					}
					c := slices.Clone(g)
					tw := butterflyTwiddles(int(4*m*stride + 1))
					before := slices.Clone(tw)
					transform(&g[1], &tw[0], stride, m, inverse)
					nativeMiniButterfly(radix, c[1:], tw, stride, m, inverse)
					if !sameComplexBits(g, c) || !slices.Equal(tw, before) {
						t.Fatal(radix, m, stride, inverse, variant)
					}
				}
			}
		}
	}
}

func TestFFTInverseAgainstC(t *testing.T) { compareFullFFT(t, 2, opuscc.Opus_opus_ifft_c) }

func TestFFTForwardAgainstC(t *testing.T) {
	compareFullFFT(t, 1, opuscc.Opus_opus_fft_c)
}

func compareFullFFT(t *testing.T, op int32, transform func(*libc.TLS, *opuscc.OpusT_kiss_fft_state, *int16, *opuscc.OpusT_kiss_twiddle_cpx, *opuscc.OpusT_kiss_fft_cpx, *opuscc.OpusT_kiss_fft_cpx)) {
	t.Helper()
	for _, n := range []int32{4, 8, 16, 32, 60, 120, 240, 480} {
		for _, shift := range []int32{-1, 0, 1, 2} {
			for variant := 0; variant < 3; variant++ {
				st, br, tw := nativeFFTFixture(n, shift)
				g := make([]opuscc.OpusT_kiss_fft_cpx, n+4)
				for i := range g {
					g[i] = opuscc.OpusT_kiss_fft_cpx{Fr: float32(i-13) / 7, Fi: float32(19-i) / 11}
				}
				c := slices.Clone(g)
				input := slices.Clone(g[1 : n+1])
				ci := slices.Clone(input)
				if variant == 1 {
					input = g[:n]
					ci = c[:n]
				} else if variant == 2 {
					input = g[2 : n+2]
					ci = c[2 : n+2]
				}
				before := st
				brBefore := slices.Clone(br)
				twBefore := slices.Clone(tw)
				transform(nil, &st, &br[0], &tw[0], &input[0], &g[1])
				nativeFFTTransform(&st, br, tw, ci, c[1:], op)
				if !sameComplexBits(g, c) || st != before || !slices.Equal(br, brBefore) || !slices.Equal(tw, twBefore) {
					t.Fatal(op, n, shift, variant)
				}
			}
		}
	}
}

func TestFFTImplAgainstC(t *testing.T) {
	for _, n := range []int32{4, 8, 16, 32, 60, 120, 240, 480} {
		for _, shift := range []int32{-1, 0, 1, 2} {
			st, br, tw := nativeFFTFixture(n, shift)
			g := make([]opuscc.OpusT_kiss_fft_cpx, n+2)
			for i := range g {
				g[i] = opuscc.OpusT_kiss_fft_cpx{Fr: float32(i-13) / 7, Fi: float32(19-i) / 11}
			}
			c := slices.Clone(g)
			before := st
			twBefore := slices.Clone(tw)
			opuscc.Opus_opus_fft_impl(nil, &st, &tw[0], &g[1])
			nativeFFTTransform(&st, br, tw, nil, c[1:], 0)
			if !sameComplexBits(g, c) || st != before || !slices.Equal(tw, twBefore) {
				t.Fatal(n, shift)
			}
		}
	}
}

func TestFFTButterfly5AgainstC(t *testing.T) {
	compareButterflyStages(t, 5, opuscc.CompareFFTButterfly5)
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
