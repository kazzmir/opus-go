package opuscc

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"math"
	"runtime"
	"slices"
	"testing"
	"unsafe"
)

// No byte-backed objects or pins: the mode is the sole owner of its transform tables.
func newSynthesisTestMode() *OpusT_OpusCustomMode {
	m := mode48000_960_120
	bands := slices.Clone(unsafe.Slice(m.FeBands, m.FnbEBands+1))
	m.FeBands = unsafe.SliceData(bands)
	window := slices.Clone(unsafe.Slice(m.Fwindow, m.Foverlap))
	m.Fwindow = unsafe.SliceData(window)
	trig := slices.Clone(unsafe.Slice(m.Fmdct.Ftrig, m.Fmdct.Fn-((m.Fmdct.Fn/2)>>m.Fmdct.Fmaxshift)))
	m.Fmdct.Ftrig = unsafe.SliceData(trig)
	for i, original := range m.Fmdct.Fkfft {
		if original == nil {
			continue
		}
		st := *original
		rev := slices.Clone(unsafe.Slice(st.Fbitrev, st.Fnfft))
		tw := slices.Clone(unsafe.Slice(st.Ftwiddles, mode48000_960_120.Fmdct.Fkfft[0].Fnfft))
		st.Fbitrev = unsafe.SliceData(rev)
		st.Ftwiddles = unsafe.SliceData(tw)
		m.Fmdct.Fkfft[i] = &st
	}
	return &m
}

func TestSynthesisScratchPointers(t *testing.T) {
	m := newSynthesisTestMode()
	for LM := int32(0); LM <= 3; LM++ {
		_, bands, N := celtSynthesisGeometry(m, LM)
		for _, layout := range [][2]int32{{1, 1}, {2, 2}, {1, 2}, {2, 1}, {0, 0}, {0, 1}} {
			owners := func() []*float32 {
				x, e := make([]float32, N*max(layout[0], layout[1], 1)+2), make([]float32, bands*max(layout[0], layout[1], 1)+2)
				l, r := make([]float32, N+62), make([]float32, N+62)
				for i := range x {
					x[i] = float32(i%19-9) / 32
				}
				for i := range e {
					e[i] = []float32{-28, -9, .25, 8, 32}[i%5]
				}
				for i := range l {
					l[i] = float32(i%11-5) / 31
					r[i] = float32(i%13-6) / 37
				}
				l[0], l[len(l)-1], r[0], r[len(r)-1] = 77, 88, 99, 111
				return []*float32{&x[1], &e[1], &l[0], &r[0]}
			}()
			g0, g1 := unsafe.Slice(owners[2], N+62), unsafe.Slice(owners[3], N+62)
			w0, w1 := slices.Clone(g0), slices.Clone(g1)
			actual, expected := make([]*float32, max(layout[1], 1)), make([]*float32, max(layout[1], 1))
			actual[0] = &g0[1]
			expected[0] = &w0[1]
			if layout[1] == 2 {
				actual[1] = &g1[1]
				expected[1] = &w1[1]
			}
			entropyInitGrowStack(12)
			runtime.GC()
			for _, transient := range []int32{0, 1} {
				for _, down := range []int32{1, 2, 3, 4, 6} {
					for _, silence := range []int32{0, 1} {
						for _, limits := range [][2]int32{{0, 21}, {1, 19}, {21, 21}} {
							celt_synthesis(nil, m, owners[0], &actual[0], owners[1], limits[0], limits[1], layout[0], layout[1], transient, LM, down, silence, 0)
							celt_synthesis(nil, &mode48000_960_120, owners[0], &expected[0], owners[1], limits[0], limits[1], layout[0], layout[1], transient, LM, down, silence, 0)
							for _, pair := range [][2][]float32{{g0, w0}, {g1, w1}} {
								for i := range pair[0] {
									if math.Float32bits(pair[0][i]) != math.Float32bits(pair[1][i]) {
										t.Fatal("owned synthesis", LM, layout, transient, down, silence, limits, i, pair[0][i], pair[1][i])
									}
								}
							}
							if g0[0] != 77 || g0[len(g0)-1] != 88 || g1[0] != 99 || g1[len(g1)-1] != 111 {
								t.Fatal("synthesis guards", LM, layout)
							}
						}
					}
				}
			}
		}
	}
	// This path must neither consult nor update a TLS pseudostack.
	tls := libc.NewTLS()
	defer tls.Close()
	libc.Xpthread_setspecific(tls, 0x6f707573, 123)
	x, e, out := make([]float32, 120), make([]float32, 21), make([]float32, 180)
	heads := [1]*float32{&out[0]}
	celt_synthesis(tls, m, &x[0], &heads[0], &e[0], 0, 21, 1, 1, 0, 0, 1, 1, 0)
	if libc.Xpthread_getspecific(tls, 0x6f707573) != 123 {
		t.Fatal("TLS changed")
	}
	// Silent mono synthesis ignores spectral/energy data but still consumes TDAC history.
	celt_synthesis(nil, m, nil, &heads[0], nil, 0, 21, 1, 1, 0, 0, 1, 1, 0)
}

func TestSynthesisAliasesPointers(t *testing.T) {
	m := newSynthesisTestMode()
	for LM := int32(0); LM <= 3; LM++ {
		overlap, bands, N := celtSynthesisGeometry(m, LM)
		for _, transient := range []int32{0, 1} {
			for _, offset := range []int32{0, 17, 60} {
				x, e := make([]float32, 2*N), make([]float32, 2*bands)
				for i := range x {
					x[i] = float32(i%19-9) / 32
				}
				g := make([]float32, N+overlap/2+offset+2)
				for i := range g {
					g[i] = float32(i%11-5) / 31
				}
				g[0], g[len(g)-1] = 77, 88
				want := slices.Clone(g)
				heads := [2]*float32{&g[1], &g[1+offset]}
				expected := [2]*float32{&want[1], &want[1+offset]}
				entropyInitGrowStack(12)
				runtime.GC()
				celt_synthesis(nil, m, &x[0], &heads[0], &e[0], 0, 21, 2, 2, transient, LM, 1, 0, 0)
				celt_synthesis(nil, &mode48000_960_120, &x[0], &expected[0], &e[0], 0, 21, 2, 2, transient, LM, 1, 0, 0)
				if !slices.Equal(g, want) || g[0] != 77 || g[len(g)-1] != 88 {
					t.Fatal("shared output lifetime", LM, transient, offset)
				}
			}
		}
	}
	// Upmix outputs may overlap its staging input: a Go-only ordering test,
	// not a C oracle input (the MDCT output restrict contract forbids this).
	x, e := make([]float32, 120), make([]float32, 21)
	for i := range x {
		x[i] = float32(i%19-9) / 32
	}
	g := make([]float32, 182)
	for i := range g {
		g[i] = float32(i%11-5) / 31
	}
	g[0], g[181] = 77, 88
	want := slices.Clone(g)
	freq := make([]float32, 120)
	Opus_denormalise_bands(nil, m.FeBands, 120, &x[0], &freq[0], &e[0], 0, 21, 1, 1, 0)
	copy(want[61:181], freq)
	celtSynthesisIMDCT(nil, m, &want[61], &want[1], 120, 3, 1, 0)
	celtSynthesisIMDCT(nil, m, &freq[0], &want[1], 120, 3, 1, 0)
	heads := [2]*float32{&g[1], &g[1]}
	celt_synthesis(nil, m, &x[0], &heads[0], &e[0], 0, 21, 1, 2, 0, 0, 1, 0, 0)
	if !slices.Equal(g, want) || g[0] != 77 || g[181] != 88 {
		t.Fatal("upmix copy/left/right order")
	}
}

func TestSynthesisOutputPointers(t *testing.T) {
	m := newSynthesisTestMode()
	for LM := int32(0); LM <= 3; LM++ {
		overlap, _, N := celtSynthesisGeometry(m, LM)
		outputs := func() []*float32 {
			l, r := make([]float32, N+overlap/2+2), make([]float32, N+overlap/2+2)
			for i := range l {
				l[i] = float32(i%11-5) / 31
				r[i] = float32(i%13-6) / 37
			}
			l[0], l[len(l)-1], r[0], r[len(r)-1] = 77, 88, 99, 111
			return []*float32{unsafe.SliceData(l), unsafe.SliceData(r)}
		}()
		entropyInitGrowStack(12)
		runtime.GC()
		in := make([]float32, N)
		for i := range in {
			in[i] = float32(i%19-9) / 32
		}
		for _, transient := range []bool{false, true} {
			g0, g1 := unsafe.Slice(outputs[0], N+overlap/2+2), unsafe.Slice(outputs[1], N+overlap/2+2)
			w0, w1 := slices.Clone(g0), slices.Clone(g1)
			shift := m.FmaxLM - LM
			B, NB := int32(1), N
			if transient {
				shift = m.FmaxLM
				B = 1 << LM
				NB = m.FshortMdctSize
			}
			temp := celtNormAdd(outputs[1], 1+overlap/2)
			celtSynthesisOutputCopy(temp, &in[0], N)
			copy(w1[1+overlap/2:], in)
			for b := int32(0); b < B; b++ {
				celtSynthesisIMDCT(nil, m, celtNormAdd(temp, b), celtNormAdd(outputs[0], 1+NB*b), overlap, shift, B, 0)
				celtSynthesisIMDCT(nil, m, &w1[1+overlap/2+b], &w0[1+NB*b], overlap, shift, B, 0)
			}
			for b := int32(0); b < B; b++ {
				celtSynthesisIMDCT(nil, m, &in[b], celtNormAdd(outputs[1], 1+NB*b), overlap, shift, B, 0)
				celtSynthesisIMDCT(nil, m, &in[b], &w1[1+NB*b], overlap, shift, B, 0)
			}
			if !slices.Equal(g0, w0) || !slices.Equal(g1, w1) || g0[0] != 77 || g0[len(g0)-1] != 88 || g1[0] != 99 || g1[len(g1)-1] != 111 {
				t.Fatal("output owners/staging", LM, transient)
			}
		}
	}
}

// Early migration rounds exercise the typed input chain through the exact
// denormalizer used by synthesis. The completed driver is checked above.
func TestSynthesisSpectrumPointers(t *testing.T) {
	for LM := int32(0); LM <= 3; LM++ {
		m := newSynthesisTestMode()
		_, bands, N := celtSynthesisGeometry(m, LM)
		owners := func() []*float32 {
			x, e := make([]float32, 2*N), make([]float32, 2*bands)
			for i := range x {
				x[i] = float32(i%19-9) / 32
			}
			for i := range e {
				e[i] = float32(i%5 - 3)
			}
			return []*float32{unsafe.SliceData(x), unsafe.SliceData(e)}
		}()
		entropyInitGrowStack(12)
		runtime.GC()
		for _, limits := range [][2]int32{{0, 21}, {1, 19}, {21, 21}} {
			for _, down := range []int32{1, 2, 3, 4, 6} {
				for _, silence := range []int32{0, 1} {
					for c := int32(0); c < 2; c++ {
						g := make([]float32, N+2)
						g[0] = 77
						g[N+1] = 88
						want := slices.Clone(g)
						x, e := celtNormAdd(owners[0], c*N), celtNormAdd(owners[1], c*bands)
						Opus_denormalise_bands(nil, m.FeBands, m.FshortMdctSize, x, &g[1], e, limits[0], limits[1], 1<<LM, down, silence)
						Opus_denormalise_bands(nil, mode48000_960_120.FeBands, 120, x, &want[1], e, limits[0], limits[1], 1<<LM, down, silence)
						if !slices.Equal(g, want) || g[0] != 77 || g[N+1] != 88 {
							t.Fatal("spectral owners", LM, limits, down, silence, c)
						}
					}
				}
			}
		}
	}
}

func TestSynthesisModePointers(t *testing.T) {
	m := newSynthesisTestMode()
	entropyInitGrowStack(12)
	runtime.GC()
	for LM := int32(0); LM <= 3; LM++ {
		overlap, bands, N := celtSynthesisGeometry(m, LM)
		if overlap != 120 || bands != 21 || N != 120<<LM {
			t.Fatal("geometry", overlap, bands, N)
		}
		for _, transient := range []bool{false, true} {
			shift := m.FmaxLM - LM
			B := int32(1)
			NB := N
			if transient {
				shift = m.FmaxLM
				B = 1 << LM
				NB = m.FshortMdctSize
			}
			in := make([]float32, N)
			for i := range in {
				in[i] = float32(i%13-6) / 8
			}
			g := make([]float32, N+overlap/2+2)
			g[0] = 77
			g[len(g)-1] = 88
			for i := 1; i < len(g)-1; i++ {
				g[i] = float32(i%7-3) / 31
			}
			want := slices.Clone(g)
			for b := int32(0); b < B; b++ {
				celtSynthesisIMDCT(nil, m, &in[b], &g[1+NB*b], overlap, shift, B, 0)
				celtSynthesisIMDCT(nil, &mode48000_960_120, &in[b], &want[1+NB*b], overlap, shift, B, 0)
			}
			if !slices.Equal(g, want) || g[0] != 77 || g[len(g)-1] != 88 {
				t.Fatal("owned transform", LM, transient)
			}
		}
	}
}
