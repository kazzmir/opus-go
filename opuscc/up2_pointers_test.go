package opuscc

import (
	"slices"
	"testing"
)

func TestUp2Pointers(t *testing.T) {
	input := []int16{-32768, 32767, 0, 1, -1, 100, 20000, -30000}
	before := slices.Clone(input)
	initial := [6]int32{12345, -98765, 65432, -11111, 22222, -33333}
	want := []int16{-136, -1139, -5013, -12497, -18418, -11580, 11619, 32718, 25806, -5450, -23750, -7408, 16187, 12092, -6977, -4300}
	wantState := [6]int32{-32098525, 23918734, 23502913, -36297401, 21703549, 88885463}
	for split := 0; split < len(input); split++ {
		state := initial
		out := make([]int16, len(want)+2)
		out[0], out[len(out)-1] = 111, 222
		Opus_silk_resampler_private_up2_HQ(nil, &state, &out[1], &input[0], int32(split))
		Opus_silk_resampler_private_up2_HQ(nil, &state, &out[1+2*split], &input[split], int32(len(input)-split))
		if !slices.Equal(out[1:len(out)-1], want) || state != wantState {
			t.Fatalf("split=%d output=%v state=%v", split, out, state)
		}
		if out[0] != 111 || out[len(out)-1] != 222 || !slices.Equal(input, before) {
			t.Fatal("input or sentinels changed")
		}
	}
	state := initial
	Opus_silk_resampler_private_up2_HQ(nil, &state, nil, nil, 0)
	if state != initial {
		t.Fatal("empty frame changed state")
	}
	// A one-sample expansion reads the input before writing either output.
	overlap := []int16{input[0], 0}
	state = initial
	Opus_silk_resampler_private_up2_HQ(nil, &state, &overlap[0], &overlap[0], 1)
	if !slices.Equal(overlap, want[:2]) {
		t.Fatal("one-sample overlap differs")
	}
}
