package opuscc

import (
	"testing"
	"unsafe"
)

func TestStereoMSToLRFieldAccesses(t *testing.T) {
	state := OpusT_stereo_dec_state{
		Fpred_prev_Q13: [2]OpusT_opus_int16{1200, -900},
		FsMid:          [2]OpusT_opus_int16{90, -130},
		FsSide:         [2]OpusT_opus_int16{-70, 110},
	}
	mid := []int16{17, -29, 43, -57, 71, -83, 97, -109, 127, -149}
	side := []int16{-19, 31, -47, 59, -73, 89, -101, 113, -131, 151}
	predictors := []int32{2600, -1700}

	Opus_silk_stereo_MS_to_LR(nil, uintptr(unsafe.Pointer(&state)), uintptr(unsafe.Pointer(&mid[0])), uintptr(unsafe.Pointer(&side[0])), uintptr(unsafe.Pointer(&predictors[0])), 1, 8)

	wantMid := []int16{90, 5491, -15, 6910, -13, 8320, -22, 9726, -31, -149}
	wantSide := []int16{-70, -5751, 101, -7024, 155, -8486, 216, -9944, 285, 151}
	for i, want := range wantMid {
		if got := mid[i]; got != want {
			t.Fatalf("mid[%d]: got %d, want %d", i, got, want)
		}
		if got := side[i]; got != wantSide[i] {
			t.Fatalf("side[%d]: got %d, want %d", i, got, wantSide[i])
		}
	}
	if got, want := state.Fpred_prev_Q13, [2]OpusT_opus_int16{2600, -1700}; got != want {
		t.Fatalf("predictor state: got %v, want %v", got, want)
	}
	if got, want := state.FsMid, [2]OpusT_opus_int16{127, -149}; got != want {
		t.Fatalf("mid history: got %v, want %v", got, want)
	}
	if got, want := state.FsSide, [2]OpusT_opus_int16{-131, 151}; got != want {
		t.Fatalf("side history: got %v, want %v", got, want)
	}
}
