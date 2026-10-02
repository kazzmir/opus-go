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

func TestAllocationOutputsPointers(t *testing.T) {
	m := newSynthesisTestMode()
	for _, C := range []int32{1, 2} {
		for LM := int32(0); LM <= 3; LM++ {
			for _, budget := range []int32{0, 8, 100, 512, 4096, 12000} {
				for _, encode := range []int32{0, 1} {
					a := new([7][23]int32)
					for k := range a {
						a[k][0], a[k][22] = 77, 88
					}
					for j := int32(0); j < 21; j++ {
						width := int32(modeBand(m, j+1) - modeBand(m, j))
						a[1][j+1] = width * C << LM << 3
						a[2][j+1] = max(C<<3, width<<LM<<2)
						a[3][j+1] = width * C << LM << 6
					}
					b := *a
					s, w := [3]int32{77, 21, 1}, [3]int32{77, 21, 1}
					g := func() *OpusT_ec_ctx {
						data := make([]byte, 256)
						for i := range data {
							data[i] = byte(i*17 + 31)
						}
						ctx := new(OpusT_ec_ctx)
						if encode != 0 {
							Opus_ec_enc_init(nil, ctx, &data[0], 256)
						} else {
							Opus_ec_dec_init(nil, ctx, &data[0], 256)
						}
						return ctx
					}()
					e := *g
					buf := slices.Clone(unsafe.Slice(e.Fbuf, 256))
					e.Fbuf = &buf[0]
					entropyInitGrowStack(12)
					runtime.GC()
					r := interp_bits2pulses(nil, m, 0, 21, 0, &a[0][1], &a[1][1], &a[2][1], &a[3][1], budget, &s[0], 0, &s[1], 0, &s[2], 0, &a[4][1], &a[5][1], &a[6][1], C, LM, g, encode, 20, 20)
					wr := interp_bits2pulses(nil, &mode48000_960_120, 0, 21, 0, &b[0][1], &b[1][1], &b[2][1], &b[3][1], budget, &w[0], 0, &w[1], 0, &w[2], 0, &b[4][1], &b[5][1], &b[6][1], C, LM, &e, encode, 20, 20)
					gc := *g
					gc.Fbuf = nil
					e.Fbuf = nil
					if r != wr || *a != b || s != w || gc != e || !slices.Equal(unsafe.Slice(g.Fbuf, 256), buf) {
						t.Fatal("allocation owners", C, LM, budget, encode, r, wr)
					}
					for k := range a {
						if a[k][0] != 77 || a[k][22] != 88 {
							t.Fatal("allocation guard", k)
						}
					}
				}
			}
		}
	}
	// Unused entropy may be nil. Scalar outputs may alias (last dual store wins).
	var arrays [7][1]int32
	var shared, balance int32
	tls := libc.NewTLS()
	defer tls.Close()
	libc.Xpthread_setspecific(tls, 0x6f707573, 123)
	r := interp_bits2pulses(tls, m, 0, 1, 0, &arrays[0][0], &arrays[1][0], &arrays[2][0], &arrays[3][0], 0, &balance, 0, &shared, 0, &shared, 0, &arrays[4][0], &arrays[5][0], &arrays[6][0], 1, 0, nil, 0, 0, 0)
	if r != 1 || balance != 0 || shared != 0 || arrays[6][0] != 1 || libc.Xpthread_getspecific(tls, 0x6f707573) != 123 {
		t.Fatal("allocation nil entropy/scalar alias/TLS", r, balance, shared, arrays)
	}
}

func TestAllocationBandInputPointers(t *testing.T) {
	views := func() [2][]int32 {
		a, b := make([]int32, 21), make([]int32, 21)
		for i := range a {
			a[i] = int32(i) * 17
			b[i] = int32(i) * 31
		}
		return [2][]int32{a, b}
	}()
	entropyInitGrowStack(12)
	runtime.GC()
	for j := int32(0); j < 21; j++ {
		for mid := int32(0); mid <= 1<<ALLOC_STEPS; mid++ {
			want := int32(j*17 + int32(mid*j*31)>>ALLOC_STEPS)
			if got := allocationInterpBit(views[0], views[1], j, mid); got != want {
				t.Fatal("allocation input", j, mid, got, want)
			}
		}
	}
	views[0][0] = 1 << 30
	views[1][0] = 1 << 30
	if got := allocationInterpBit(views[0], views[1], 0, 3); got != 1056964608 {
		t.Fatal("int32 product/shift wrapping", got)
	}
	views[1][0] = 32
	if got := allocationInterpBit(views[0], views[1], 0, 32); got != (1<<30)+16 {
		t.Fatal("live array load", got)
	}
}

func TestAllocationModeEntropyPointers(t *testing.T) {
	mode := newSynthesisTestMode()
	logs := slices.Clone(unsafe.Slice(mode.FlogN, mode.FnbEBands))
	mode.FlogN = unsafe.SliceData(logs)
	ctx := func() *OpusT_ec_ctx {
		b := make([]byte, 64)
		e := new(OpusT_ec_ctx)
		Opus_ec_enc_init(nil, e, &b[0], 64)
		return e
	}()
	entropyInitGrowStack(12)
	runtime.GC()
	for i := int32(0); i < 21; i++ {
		if modeBand(mode, i) != eband5ms[i] || modeLogN(mode, i) != logN400[i] {
			t.Fatal("allocation table owner", i)
		}
	}
	Opus_ec_enc_uint(nil, ctx, 3, 7)
	Opus_ec_enc_done(nil, ctx)
	if ctx.Fbuf == nil || ctx.Ferror1 != 0 {
		t.Fatal("allocation entropy owner")
	}
}

func TestPrefilterScratchPointers(t *testing.T) {
	for _, channels := range []int32{0, 1, 2} {
		for _, overlap := range []int32{0, 1, 3, 119, 120} {
			for _, N := range []int32{0, 120, 240, 480, 960} {
				for _, gains := range [][2]float32{{0, 0}, {.25, .5}, {-.25, .75}} {
					storage, image, size := celtStateTestBuffer(newSynthesisTestMode(), max(channels, 1))
					st := &storage.State
					st.Fchannels = channels
					st.Foverlap = overlap
					st.Fpostfilter_period_old = 31
					st.Fpostfilter_period = 128
					st.Fpostfilter_gain_old, st.Fpostfilter_gain = gains[0], gains[1]
					st.Fpostfilter_tapset_old, st.Fpostfilter_tapset = 1, 2
					st.Farch = 0
					history := unsafe.Slice(&st.F_decode_mem[0], (DEC_PITCH_BUF_SIZE+overlap)*max(channels, 1))
					for i := range history {
						history[i] = float32(i%29-14) * 173
					}
					reference := new(celtStateTestStorage)
					*reference = *storage
					want := unsafe.Slice((*byte)(unsafe.Pointer(&reference.State)), len(image))
					window := unsafe.Slice(st.Fmode.Fwindow, 120)
					if overlap > 0 {
						for c := int32(0); c < max(channels, 1); c++ {
							input := prefilterFoldHistory(&reference.State, overlap, c)
							input = celtNormAdd(input, DEC_PITCH_BUF_SIZE-N)
							scratch := make([]float32, overlap)
							Opus_comb_filter(nil, &scratch[0], input, 31, 128, overlap, -gains[0], -gains[1], 1, 2, nil, 0, 0)
							out := unsafe.Slice(input, overlap)
							for i := int32(0); i < overlap/2; i++ {
								out[i] = float32(window[i]*scratch[overlap-1-i]) + float32(window[overlap-i-1]*scratch[i])
							}
						}
					}
					entropyInitGrowStack(12)
					runtime.GC()
					prefilter_and_fold(nil, st, N)
					if !slices.Equal(image, want) {
						for i := range image {
							if image[i] != want[i] {
								t.Fatal("prefilter full state/order", channels, overlap, N, gains, i, image[i], want[i])
							}
						}
					}
					for _, b := range image[size:] {
						if b != 165 {
							t.Fatal("prefilter guard")
						}
					}
				}
			}
		}
	}
	// A genuinely scanned, exact-sized mono history; no trailing scratch/padding
	// workaround and no mode is needed for overlap zero, even with N zero.
	exact := new(struct {
		State   OpusT_OpusCustomDecoder
		History [DEC_PITCH_BUF_SIZE - 1]float32
	})
	exact.State.Fchannels = 1
	prefilter_and_fold(nil, &exact.State, 0)
	// Unused TLS must remain untouched on the active scratch path.
	storage, _, _ := celtStateTestBuffer(newSynthesisTestMode(), 1)
	storage.State.Fpostfilter_gain_old = 0
	storage.State.Fpostfilter_gain = 0
	tls := libc.NewTLS()
	defer tls.Close()
	libc.Xpthread_setspecific(tls, 0x6f707573, 123)
	prefilter_and_fold(tls, &storage.State, 120)
	if libc.Xpthread_getspecific(tls, 0x6f707573) != 123 {
		t.Fatal("prefilter TLS changed")
	}
}

func TestPrefilterTDACRounding(t *testing.T) {
	// The scalar C/original generated fold rounds products separately. A fused
	// rewrite changed the ARM PLC golden; this overlap-three case differs by 1 ULP.
	m := newSynthesisTestMode()
	filtered := []float32{-2422, -2249, -2076}
	var out float32
	prefilterFoldTDAC(m, &out, &filtered[0], 3)
	if bits := math.Float32bits(out); bits != 0xc086cced {
		t.Fatalf("fold product rounding: %08x want c086cced", bits)
	}
}

func TestPrefilterTDACPointers(t *testing.T) {
	for _, overlap := range []int32{0, 1, 2, 3, 119, 120} {
		m := newSynthesisTestMode()
		filtered := make([]float32, overlap)
		for i := range filtered {
			filtered[i] = float32(i%29-14) * 173
		}
		g := make([]float32, overlap+2)
		for i := range g {
			g[i] = float32(i + 77)
		}
		want := slices.Clone(g)
		entropyInitGrowStack(12)
		runtime.GC()
		prefilterFoldTDAC(m, &g[1], unsafe.SliceData(filtered), overlap)
		window := unsafe.Slice(m.Fwindow, 120)
		for i := int32(0); i < overlap/2; i++ {
			want[1+i] = float32(window[i]*filtered[overlap-1-i]) + float32(window[overlap-i-1]*filtered[i])
		}
		for i := range g {
			if math.Float32bits(g[i]) != math.Float32bits(want[i]) {
				t.Fatal("fold scratch", overlap, i, g[i], want[i])
			}
		}
	}
	// Sequential alias stores are Go-only tests; the enclosing C decoder restrict
	// contract does not permit arbitrary history/window/scratch aliases.
	for _, aliasWindow := range []bool{false, true} {
		data := []float32{.1, .2, .3, .4, .5, .6, .7, .8, 77}
		want := slices.Clone(data)
		filtered := []float32{13, 11, 7, 5, 3, 2}
		window := []float32{.1, .2, .3, .4, .5, .6}
		m := OpusT_OpusCustomMode{Fwindow: &window[0]}
		if aliasWindow {
			m.Fwindow = &data[0]
			for i := 0; i < 3; i++ {
				want[1+i] = float32(want[i]*filtered[5-i]) + float32(want[5-i]*filtered[i])
			}
		} else {
			for i := 0; i < 3; i++ {
				want[1+i] = float32(window[i]*want[5-i]) + float32(window[5-i]*want[i])
			}
		}
		p := &filtered[0]
		if !aliasWindow {
			p = &data[0]
		}
		prefilterFoldTDAC(&m, &data[1], p, 6)
		if !slices.Equal(data, want) {
			t.Fatal("fold sequential aliases", aliasWindow, data, want)
		}
	}
	// No window/scratch/output access when there are no folded samples.
	prefilterFoldTDAC(nil, nil, nil, 0)
	prefilterFoldTDAC(nil, nil, nil, 1)
}

func TestPrefilterHistoryPointers(t *testing.T) {
	for _, overlap := range []int32{0, 1, 3, 119, 120} {
		for c := int32(0); c < 2; c++ {
			history := func() *float32 {
				storage, _, _ := celtStateTestBuffer(newSynthesisTestMode(), 2)
				st := &storage.State
				st.Foverlap = overlap
				st.Fpostfilter_period_old = 31
				st.Fpostfilter_period = 128
				st.Fpostfilter_gain_old = .25
				st.Fpostfilter_gain = .5
				st.Fpostfilter_tapset_old = 1
				st.Fpostfilter_tapset = 2
				p := prefilterFoldHistory(st, overlap, c)
				h := unsafe.Slice(p, DEC_PITCH_BUF_SIZE+overlap)
				for i := range h {
					h[i] = float32(i%29-14) * 173
				}
				return p
			}()
			entropyInitGrowStack(12)
			runtime.GC()
			st := (*OpusT_OpusCustomDecoder)(unsafe.Add(unsafe.Pointer(history), -int64(unsafe.Offsetof(OpusT_OpusCustomDecoder{}.F_decode_mem))-int64(c)*(DEC_PITCH_BUF_SIZE+int64(overlap))*4))
			if st.Fmode == nil || unsafe.Slice(st.Fmode.Fwindow, 120)[119] != window120[119] || st.Foverlap != overlap {
				t.Fatal("history retains scanned owner", overlap, c)
			}
			for _, N := range []int32{0, 120, 960} {
				g, w := make([]float32, overlap+2), make([]float32, overlap+2)
				g[0], g[len(g)-1] = 77, 88
				copy(w, g)
				input := celtNormAdd(history, DEC_PITCH_BUF_SIZE-N)
				Opus_comb_filter(nil, &g[1], input, st.Fpostfilter_period_old, st.Fpostfilter_period, overlap, -st.Fpostfilter_gain_old, -st.Fpostfilter_gain, st.Fpostfilter_tapset_old, st.Fpostfilter_tapset, nil, 0, st.Farch)
				Opus_comb_filter(nil, &w[1], input, 31, 128, overlap, -.25, -.5, 1, 2, nil, 0, 0)
				if !slices.Equal(g, w) || g[0] != 77 || g[len(g)-1] != 88 {
					t.Fatal("typed prefilter history", overlap, c, N)
				}
			}
		}
	}
}

func TestPrefilterStatePointers(t *testing.T) {
	for _, channels := range []int32{0, 1, 2} {
		m := newSynthesisTestMode()
		storage, _, _ := celtStateTestBuffer(m, max(channels, 1))
		st := &storage.State
		st.Fchannels = channels
		for _, overlap := range []int32{0, 1, 3, 119, 120} {
			st.Foverlap = overlap
			entropyInitGrowStack(12)
			runtime.GC()
			owned, o, c := prefilterFoldState(st)
			if owned != m || o != overlap || c != channels || unsafe.Slice(owned.Fwindow, 120)[119] != window120[119] {
				t.Fatal("prefilter state owner", overlap, channels)
			}
		}
	}
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
