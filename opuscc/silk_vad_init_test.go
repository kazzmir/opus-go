package opuscc

import (
	"testing"
	"unsafe"
)

func TestVADInitFieldAccesses(t *testing.T) {
	state := OpusT_silk_VAD_state{
		FAnaState:        [2]OpusT_opus_int32{111, -222},
		FNrgRatioSmth_Q8: [4]OpusT_opus_int32{333, 444, 555, 666},
		FNL:              [4]OpusT_opus_int32{777, 888, 999, 1111},
		Finv_NL:          [4]OpusT_opus_int32{2222, 3333, 4444, 5555},
		FNoiseLevelBias:  [4]OpusT_opus_int32{6666, 7777, 8888, 9999},
		Fcounter:         42,
	}

	if got := Opus_silk_VAD_Init(nil, uintptr(unsafe.Pointer(&state))); got != 0 {
		t.Fatalf("VAD init: got %d, want 0", got)
	}

	if got, want := state.FNoiseLevelBias, [4]OpusT_opus_int32{50, 25, 16, 12}; got != want {
		t.Fatalf("noise bias: got %v, want %v", got, want)
	}
	if got, want := state.FNL, [4]OpusT_opus_int32{5000, 2500, 1600, 1200}; got != want {
		t.Fatalf("noise levels: got %v, want %v", got, want)
	}
	if got, want := state.Finv_NL, [4]OpusT_opus_int32{429496, 858993, 1342177, 1789569}; got != want {
		t.Fatalf("inverse noise levels: got %v, want %v", got, want)
	}
	if got, want := state.FNrgRatioSmth_Q8, [4]OpusT_opus_int32{25600, 25600, 25600, 25600}; got != want {
		t.Fatalf("noise ratios: got %v, want %v", got, want)
	}
	if state.Fcounter != 15 {
		t.Fatalf("counter: got %d, want 15", state.Fcounter)
	}
	if state.FAnaState != [2]OpusT_opus_int32{} {
		t.Fatalf("analysis state was not reset: got %v", state.FAnaState)
	}
}
