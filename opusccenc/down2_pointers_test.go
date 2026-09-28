package opusccenc

import (
	"slices"
	"testing"
)

func TestDown2Pointers(t *testing.T) {
	input := []int16{-32768, 32767, 0, 1, -1, 100, 20000, -30000, 0, 0, 1234}
	before := slices.Clone(input)
	want := []int16{-7523, 5678, 3876, 412, -6094}
	initial := [2]int32{12345, -98765}
	wantState := [2]int32{-27341915, 5347117}
	for _, n := range []int32{10, 11} {
		state := initial
		out := [7]int16{111, 0, 0, 0, 0, 0, 222}
		Opus_silk_resampler_down2(nil, &state, &out[1], &input[0], n)
		if !slices.Equal(out[1:6], want) || state != wantState || out[0] != 111 || out[6] != 222 || !slices.Equal(input, before) {
			t.Fatalf("n=%d output=%v state=%v", n, out, state)
		}
	}
	state := initial
	inPlace := slices.Clone(input)
	Opus_silk_resampler_down2(nil, &state, &inPlace[0], &inPlace[0], 4)
	Opus_silk_resampler_down2(nil, &state, &inPlace[2], &inPlace[4], 6)
	if !slices.Equal(inPlace[:5], want) || state != wantState {
		t.Fatal("in-place/chunked downsampling differs")
	}
	for _, n := range []int32{0, 1} {
		Opus_silk_resampler_down2(nil, &state, nil, nil, n)
		if state != wantState {
			t.Fatal("incomplete pair changed state")
		}
	}
}
