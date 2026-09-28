package opuscc

import (
	"slices"
	"testing"
)

func TestGainsDequantPointers(t *testing.T) {
	for _, tc := range []struct {
		index, previous int8
		conditional     int32
		want            int8
	}{
		{0, 63, 0, 47}, {0, 0, 1, 0}, {4, 20, 1, 20}, {127, 63, 1, 0}, {63, 0, 0, 63},
	} {
		output := [3]int32{111, 0, 222}
		previous := tc.previous
		Opus_silk_gains_dequant(nil, &output[1], &tc.index, &previous, tc.conditional, 1)
		if previous != tc.want || output[1] <= 0 || output[0] != 111 || output[2] != 222 {
			t.Fatalf("case=%+v state=%d output=%v", tc, previous, output)
		}
	}
	input := []int8{20, 3, 15, 0}
	whole, chunks := [4]int32{}, [4]int32{}
	a, b := int8(30), int8(30)
	Opus_silk_gains_dequant(nil, &whole[0], &input[0], &a, 0, 4)
	Opus_silk_gains_dequant(nil, &chunks[0], &input[0], &b, 0, 2)
	Opus_silk_gains_dequant(nil, &chunks[2], &input[2], &b, 1, 2)
	if whole != chunks || a != b || !slices.Equal(input, []int8{20, 3, 15, 0}) {
		t.Fatal("chunked result or input differs")
	}
	Opus_silk_gains_dequant(nil, nil, nil, nil, 0, 0)
}
