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

func TestMDCTBackwardAgainstC(t *testing.T) { compareMDCTTransforms(t, 1) }
func TestMDCTForwardAgainstC(t *testing.T)  { compareMDCTTransforms(t, 0) }
func compareMDCTTransforms(t *testing.T, op int32) {
	for _, n := range []int32{20, 32, 1920} {
		for shift := int32(0); shift <= 3; shift++ {
			effective := n >> shift
			if effective%4 != 0 || effective < 16 {
				continue
			}
			trig, _, _ := nativeMDCTLookup(n, shift)
			state, bitrev, tw := nativeFFTFixture(effective/4, func() int32 {
				if shift == 0 {
					return -1
				}
				return shift
			}())
			l := opuscc.OpusT_mdct_lookup{Fn: n, Fmaxshift: shift, Ftrig: &trig[0]}
			l.Fkfft[shift] = &state
			for _, overlap := range []int32{0, 4, min(120, (effective/2)&^3)} {
				for _, stride := range []int32{1, 2, 3} {
					for trial := 0; trial < 3; trial++ {
						count := (effective/2-1)*stride + 1
						input := make([]float32, max(count, effective/2+overlap))
						for i := range input {
							input[i] = float32((i*17+trial*3)%51-25) * .03125
						}
						if trial == 0 {
							clear(input)
						}
						window := make([]float32, overlap)
						for i := range window {
							window[i] = float32(math.Sin(float64(i+1) * .5 / float64(overlap) * math.Pi))
						}
						outLen := count
						if op != 0 {
							outLen = effective/2 + overlap/2
						}
						g := make([]float32, outLen+2)
						for i := range g {
							g[i] = .375
						}
						g[0] = 77
						g[len(g)-1] = 88
						c := slices.Clone(g)
						ci := slices.Clone(input)
						if op == 0 {
							opuscc.Opus_clt_mdct_forward_c(nil, &l, &bitrev[0], &tw[0], &input[0], &g[1], unsafe.SliceData(window), overlap, shift, stride, 0)
						} else {
							opuscc.Opus_clt_mdct_backward_c(nil, &l, &bitrev[0], &tw[0], &input[0], &g[1], unsafe.SliceData(window), overlap, shift, stride, 0)
						}
						nativeMDCTTransform(&l, bitrev, tw, trig, ci, c[1:len(c)-1], window, overlap, shift, stride, op)
						if !sameFloatBits(g, c) || !sameFloatBits(input, ci) {
							for i := range g {
								if math.Float32bits(g[i]) != math.Float32bits(c[i]) {
									t.Fatal(op, n, shift, overlap, stride, trial, i, g[i], c[i])
								}
							}
							t.Fatal("input mutation")
						}
					}
				}
			}
		}
	}
}

func TestMDCTLookupAgainstC(t *testing.T) {
	for _, n := range []int32{32, 240, 1920} {
		for shifts := int32(0); shifts <= 3; shifts++ {
			if (n>>shifts)%16 != 0 {
				continue
			}
			trig, sizes, layout := nativeMDCTLookup(n, shifts)
			l := opuscc.OpusT_mdct_lookup{Fn: n, Fmaxshift: shifts, Ftrig: &trig[0]}
			for i := int32(0); i <= shifts; i++ {
				l.Fkfft[i] = &opuscc.OpusT_kiss_fft_state{Fnfft: n >> 2 >> i}
				if l.Fkfft[i].Fnfft != sizes[i] {
					t.Fatal(n, shifts, sizes)
				}
			}
			want := [3]uint64{uint64(unsafe.Sizeof(l)), uint64(unsafe.Offsetof(l.Fkfft)), uint64(unsafe.Offsetof(l.Ftrig))}
			if layout != want {
				t.Fatal(layout, want)
			}
			if len(unsafe.Slice(l.Ftrig, n-((n/2)>>shifts))) != len(trig) {
				t.Fatal("trig span")
			}
		}
	}
}

func TestMiniFFTRAgainstC(t *testing.T) {
	for _, n := range []int32{4, 6, 8, 10, 12, 16, 24, 30, 60, 120, 240, 480} {
		for trial := 0; trial < 12; trial++ {
			g := opuscc.Opus_mini_kiss_fftr_alloc(nil, n, 0, nil, nil)
			c := nativeMiniRFixture(n)
			input := make([]float32, n)
			for i := range input {
				input[i] = float32(math.Sin(float64(i*17+trial) * 0.31))
			}
			if trial == 0 {
				clear(input)
				input[0] = 1
			}
			if trial == 1 {
				for i := range input {
					input[i] = math.Float32frombits(uint32(i%2) << 31)
				}
			}
			before := slices.Clone(input)
			subBefore := slices.Clone(unsafe.Slice((*byte)(unsafe.Pointer(g.Fsubstate)), 264+8*int(n/2)))
			twBefore := slices.Clone(unsafe.Slice(g.Fsuper_twiddles, n/4))
			out := make([]opuscc.OpusT_mini_kiss_fft_cpx, n/2+3)
			for i := range out {
				out[i] = opuscc.OpusT_mini_kiss_fft_cpx{Fr: 77, Fi: 88}
			}
			native := slices.Clone(out)
			opuscc.Opus_mini_kiss_fftr(nil, g, &input[0], &out[1])
			nativeMiniRTransform(c, &input[0], &native[1])
			if !sameComplexBits(out, native) || !sameFloatBits(input, before) || !sameComplexBits(unsafe.Slice(g.Ftmpbuf, n/2), unsafe.Slice(c.Ftmpbuf, n/2)) || !slices.Equal(subBefore, unsafe.Slice((*byte)(unsafe.Pointer(g.Fsubstate)), len(subBefore))) || !sameComplexBits(twBefore, unsafe.Slice(g.Fsuper_twiddles, n/4)) {
				t.Fatal(n, trial, "transform/state")
			}
		}
	}
	// The complete time input is read into scratch before aliased frequency output stores.
	for _, n := range []int32{8, 12, 30, 120} {
		g := opuscc.Opus_mini_kiss_fftr_alloc(nil, n, 0, nil, nil)
		c := nativeMiniRFixture(n)
		goBuffer := make([]float32, n+4)
		for i := range goBuffer {
			goBuffer[i] = float32(math.Sin(float64(i)))
		}
		cBuffer := slices.Clone(goBuffer)
		opuscc.Opus_mini_kiss_fftr(nil, g, &goBuffer[1], (*opuscc.OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(&goBuffer[1])))
		nativeMiniRTransform(c, &cBuffer[1], (*opuscc.OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(&cBuffer[1])))
		if !sameFloatBits(goBuffer, cBuffer) {
			t.Fatal("alias", n, goBuffer, cBuffer)
		}
	}
}

func TestMiniFFTRAllocAgainstC(t *testing.T) {
	for _, n := range []int32{2, 4, 6, 8, 10, 16, 24, 60, 120, 240, 480} {
		for _, inverse := range []int32{0, 1, -1, 2} {
			var needed uint64
			opuscc.Opus_mini_kiss_fftr_alloc(nil, n, inverse, nil, &needed)
			cn := uint64(0)
			if ok, _ := nativeMiniRAlloc(n, inverse, nil, &cn); ok || cn != needed {
				t.Fatal("query", n, inverse, needed, cn)
			}
			for _, capacity := range []uint64{0, needed - 1, needed, needed + 64} {
				for _, noMem := range []bool{false, true} {
					g := make([]uint64, (needed+64+7)/8+2)
					for i := range g {
						g[i] = 0xa5a5a5a5a5a5a5a5
					}
					c := slices.Clone(g)
					gp := (*byte)(unsafe.Pointer(&g[1]))
					cp := (*byte)(unsafe.Pointer(&c[1]))
					if noMem {
						gp = nil
						cp = nil
					}
					gs, cs := capacity, capacity
					st := opuscc.Opus_mini_kiss_fftr_alloc(nil, n, inverse, gp, &gs)
					success, offsets := nativeMiniRAlloc(n, inverse, cp, &cs)
					if (st != nil) != success || gs != cs {
						t.Fatal(n, inverse, capacity, noMem, "status/size")
					}
					if st != nil {
						base := uintptr(unsafe.Pointer(st))
						goOffsets := [3]uint64{uint64(uintptr(unsafe.Pointer(st.Fsubstate)) - base), uint64(uintptr(unsafe.Pointer(st.Ftmpbuf)) - base), uint64(uintptr(unsafe.Pointer(st.Fsuper_twiddles)) - base)}
						header := int(unsafe.Sizeof(*st)) / 8
						if goOffsets != offsets || g[0] != c[0] || !slices.Equal(g[1+header:], c[1+header:]) {
							t.Fatal(n, inverse, capacity, noMem, "layout/tables", goOffsets, offsets)
						}
					} else if !slices.Equal(g, c) {
						t.Fatal("failed allocation wrote storage")
					}
				}
			}
		}
	}
}

func TestMiniFFTAllocAgainstC(t *testing.T) {
	for _, n := range []int32{1, 2, 3, 4, 5, 7, 8, 11, 16, 31, 60, 120, 240, 480} {
		for _, inverse := range []int32{0, 1, -1, 2} {
			var needed uint64
			opuscc.Opus_mini_kiss_fft_alloc(nil, n, inverse, nil, &needed)
			cn := uint64(0)
			if nativeMiniAlloc(n, inverse, nil, &cn) || cn != needed {
				t.Fatal("query", n, inverse, needed, cn)
			}
			for _, capacity := range []uint64{0, needed - 1, needed, needed + 64} {
				for _, noMem := range []bool{false, true} {
					g := make([]uint64, (needed+64+7)/8+2)
					for i := range g {
						g[i] = 0xa5a5a5a5a5a5a5a5
					}
					c := slices.Clone(g)
					gp := (*byte)(unsafe.Pointer(&g[1]))
					cp := (*byte)(unsafe.Pointer(&c[1]))
					if noMem {
						gp = nil
						cp = nil
					}
					gs, cs := capacity, capacity
					st := opuscc.Opus_mini_kiss_fft_alloc(nil, n, inverse, gp, &gs)
					success := nativeMiniAlloc(n, inverse, cp, &cs)
					if (st != nil) != success || gs != cs || !slices.Equal(g, c) {
						t.Fatal(n, inverse, capacity, noMem, "success", st != nil, success, "size", gs, cs)
					}
				}
			}
		}
	}
}

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

func TestMiniFFTAgainstC(t *testing.T) { compareMiniTransforms(t, 2, opuscc.CompareMiniFFT) }

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
