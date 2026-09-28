package opusccenc

import (
	"slices"
	"testing"
)

func TestBiquad1Pointers(t *testing.T) {
	input := []int16{-32768, 32767, 0, 1, -1, 100, -200, 300}
	before := slices.Clone(input)
	var a [2]int32
	for _, tc := range []struct {
		b    [3]int32
		want []int16
	}{
		{[3]int32{1 << 28}, input},
		{[3]int32{0, 1 << 28}, []int16{0, -32768, 32767, 0, 1, -1, 100, -200}},
		{[3]int32{1 << 29}, []int16{-32768, 32767, 0, 2, -2, 200, -400, 600}},
	} {
		var state [2]int32
		out := make([]int16, len(input)+2)
		out[0], out[len(out)-1] = 111, 222
		bBefore := tc.b
		Opus_silk_biquad_alt_stride1(nil, &input[0], &tc.b, &a, &state, &out[1], int32(len(input)))
		if !slices.Equal(out[1:len(out)-1], tc.want) || out[0] != 111 || out[len(out)-1] != 222 || !slices.Equal(input, before) || tc.b != bBefore || a != [2]int32{} {
			t.Fatalf("b=%v output=%v want=%v", tc.b, out, tc.want)
		}
		inPlace := slices.Clone(input)
		var splitState [2]int32
		Opus_silk_biquad_alt_stride1(nil, &inPlace[0], &tc.b, &a, &splitState, &inPlace[0], 4)
		Opus_silk_biquad_alt_stride1(nil, &inPlace[4], &tc.b, &a, &splitState, &inPlace[4], 4)
		if !slices.Equal(inPlace, tc.want) || splitState != state {
			t.Fatal("in-place/chunked result differs")
		}
	}
	state := [2]int32{12, 34}
	b := [3]int32{1 << 28}
	Opus_silk_biquad_alt_stride1(nil, nil, &b, &a, &state, nil, 0)
	if state != [2]int32{12, 34} {
		t.Fatal("empty input changed state")
	}
}
