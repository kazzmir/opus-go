package opuscc

import (
	"testing"
	"unsafe"
)

func TestVADGetNoiseLevelsFieldAccesses(t *testing.T) {
	energies := [4]OpusT_opus_int32{187000, 59000, 1210000, 318000}
	state := OpusT_silk_VAD_state{
		FNL:             [4]OpusT_opus_int32{430000, 290000, 880000, 510000},
		Finv_NL:         [4]OpusT_opus_int32{4994, 7405, 2440, 4209},
		FNoiseLevelBias: [4]OpusT_opus_int32{1200, 3400, 5600, 7800},
		Fcounter:        33,
	}

	silk_VAD_GetNoiseLevels(nil, uintptr(unsafe.Pointer(&energies[0])), uintptr(unsafe.Pointer(&state)))

	wantNoiseLevels := [4]OpusT_opus_int32{354194, 180369, 922855, 466337}
	if state.FNL != wantNoiseLevels {
		t.Fatalf("noise levels: got %v, want %v", state.FNL, wantNoiseLevels)
	}
	wantInverseLevels := [4]OpusT_opus_int32{6063, 11906, 2327, 4605}
	if state.Finv_NL != wantInverseLevels {
		t.Fatalf("inverse noise levels: got %v, want %v", state.Finv_NL, wantInverseLevels)
	}
	if state.Fcounter != 34 {
		t.Fatalf("counter: got %d, want 34", state.Fcounter)
	}
}
