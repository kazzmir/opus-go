package opuscc

import (
	"testing"
	"unsafe"
)

func TestAmp2Log2LocalUnion(t *testing.T) {
	mode := OpusT_OpusCustomMode{FnbEBands: 3}
	bandE := []OpusT_celt_ener{0.75, 1.5, 3.25, 0.625, 2.25, 4.5}
	bandLogE := make([]OpusT_celt_glog, len(bandE))

	Opus_amp2Log2(nil, uintptr(unsafe.Pointer(&mode)), 2, 3, uintptr(unsafe.Pointer(&bandE[0])), uintptr(unsafe.Pointer(&bandLogE[0])), 2)

	want := []OpusT_celt_glog{-6.8525376, -5.6650376, -14, -7.115572, -5.0800753, -14}
	for i, expected := range want {
		if got := bandLogE[i]; got != expected {
			t.Fatalf("bandLogE[%d]: got %v, want %v", i, got, expected)
		}
	}
}
