package opuscc

import (
	"runtime"
	"slices"
	"testing"
	"unsafe"
)

func cloneTestNLSFCodebook(src *OpusT_silk_NLSF_CB_struct) *OpusT_silk_NLSF_CB_struct {
	cb := *src
	n, v := int(cb.Forder), int(cb.FnVectors)
	cb.FCB1_NLSF_Q8 = unsafe.SliceData(slices.Clone(unsafe.Slice(src.FCB1_NLSF_Q8, n*v)))
	cb.FCB1_Wght_Q9 = unsafe.SliceData(slices.Clone(unsafe.Slice(src.FCB1_Wght_Q9, n*v)))
	cb.FCB1_iCDF = unsafe.SliceData(slices.Clone(unsafe.Slice(src.FCB1_iCDF, 2*v)))
	cb.Fpred_Q8 = unsafe.SliceData(slices.Clone(unsafe.Slice(src.Fpred_Q8, 2*(n-1))))
	cb.Fec_sel = unsafe.SliceData(slices.Clone(unsafe.Slice(src.Fec_sel, n*v/2)))
	cb.Fec_iCDF = unsafe.SliceData(slices.Clone(unsafe.Slice(src.Fec_iCDF, 72)))
	cb.Fec_Rates_Q5 = unsafe.SliceData(slices.Clone(unsafe.Slice(src.Fec_Rates_Q5, 72)))
	cb.FdeltaMin_Q15 = unsafe.SliceData(slices.Clone(unsafe.Slice(src.FdeltaMin_Q15, n+1)))
	return &cb
}

func TestNLSFCodebookPointers(t *testing.T) {
	for _, src := range []*OpusT_silk_NLSF_CB_struct{&Opus_silk_NLSF_CB_NB_MB, &Opus_silk_NLSF_CB_WB} {
		cb := cloneTestNLSFCodebook(src)
		entropyInitGrowStack(12)
		runtime.GC()
		for index := int32(0); index < int32(cb.FnVectors); index++ {
			g, c := [16]int16{}, [16]int16{}
			gp, cp := [16]byte{}, [16]byte{}
			Opus_silk_NLSF_unpack(nil, &g[0], &gp[0], cb, index)
			Opus_silk_NLSF_unpack(nil, &c[0], &cp[0], src, index)
			if g != c || gp != cp {
				t.Fatal("owned unpack", index)
			}
			indices := [17]int8{}
			indices[0] = int8(index)
			Opus_silk_NLSF_decode(nil, &g[0], &indices[0], cb)
			Opus_silk_NLSF_decode(nil, &c[0], &indices[0], src)
			if g != c {
				t.Fatal("owned decode", index)
			}
		}
		if !slices.Equal(unsafe.Slice(cb.Fec_iCDF, 72), unsafe.Slice(src.Fec_iCDF, 72)) || !slices.Equal(unsafe.Slice(cb.Fec_Rates_Q5, 72), unsafe.Slice(src.Fec_Rates_Q5, 72)) || !slices.Equal(unsafe.Slice(cb.FCB1_iCDF, 2*int(cb.FnVectors)), unsafe.Slice(src.FCB1_iCDF, 2*int(src.FnVectors))) {
			t.Fatal("owned entropy tables")
		}
	}
}

func TestNLSFUnpackPointers(t *testing.T) {
	for _, cb := range []*OpusT_silk_NLSF_CB_struct{&Opus_silk_NLSF_CB_NB_MB, &Opus_silk_NLSF_CB_WB} {
		n := int(cb.Forder)
		before := *cb
		selectors := unsafe.Slice((*uint8)(unsafe.Pointer(cb.Fec_sel)), int(cb.FnVectors)*n/2)
		pred := unsafe.Slice((*uint8)(unsafe.Pointer(cb.Fpred_Q8)), 2*(n-1))
		for index := 0; index < int(cb.FnVectors); index++ {
			ix := make([]int16, n+2)
			p := make([]uint8, n+2)
			ix[0], ix[n+1] = 111, 222
			p[0], p[n+1] = 123, 234
			Opus_silk_NLSF_unpack(nil, &ix[1], &p[1], cb, int32(index))
			for i := 0; i < n; i++ {
				nibble := selectors[index*n/2+i/2] >> uint(4*(i%2))
				if ix[i+1] != int16((nibble&14)/2)*9 || p[i+1] != pred[i+int(nibble&1)*(n-1)] {
					t.Fatalf("order=%d index=%d element=%d", n, index, i)
				}
			}
			if ix[0] != 111 || ix[n+1] != 222 || p[0] != 123 || p[n+1] != 234 || *cb != before {
				t.Fatal("sentinels or codebook changed")
			}
		}
	}
}
