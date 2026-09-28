package opuscc

import (
	"testing"
	"unsafe"
)

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
