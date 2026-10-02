package opuscc

import (
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
