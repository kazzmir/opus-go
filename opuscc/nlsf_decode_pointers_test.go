package opuscc

import (
	"testing"
	"unsafe"
)

func TestNLSFDecodePointers(t *testing.T) {
	for _, cb := range []*OpusT_silk_NLSF_CB_struct{&Opus_silk_NLSF_CB_NB_MB, &Opus_silk_NLSF_CB_WB} {
		n := int(cb.Forder)
		delta := unsafe.Slice((*int16)(unsafe.Pointer(cb.FdeltaMin_Q15)), n+1)
		for index := 0; index < int(cb.FnVectors); index++ {
			for _, residual := range []int8{-10, 0, 10} {
				var indices [17]int8
				indices[0] = int8(index)
				for i := 1; i <= n; i++ {
					indices[i] = residual
				}
				before := indices
				var output [18]int16
				output[0], output[n+1] = 111, 222
				Opus_silk_NLSF_decode(nil, &output[1], &indices[0], cb)
				if output[0] != 111 || output[n+1] != 222 || indices != before {
					t.Fatal("sentinels or indices changed")
				}
				if output[1] < delta[0] || 32768-int32(output[n]) < int32(delta[n]) {
					t.Fatal("endpoint spacing violated")
				}
				for i := 1; i < n; i++ {
					if int32(output[i+1])-int32(output[i]) < int32(delta[i]) {
						t.Fatalf("order=%d index=%d output=%v", n, index, output)
					}
				}
			}
		}
	}
}
