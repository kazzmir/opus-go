package opuscc

import (
	"slices"
	"testing"
)

// From gains_quant_pointers_test.go.
func TestGainsQuantPointers(t *testing.T) {
	Opus_silk_gains_quant(nil, nil, nil, nil, 0, 0)
	for previous := int8(0); previous < 64; previous++ {
		for cond := int32(0); cond < 2; cond++ {
			gains := [4]int32{1, 65536, 1 << 24, 2147483647}
			var indices [4]int8
			p := previous
			Opus_silk_gains_quant(nil, &indices[0], &gains[0], &p, cond, 4)
			var decoded [4]int32
			dp := previous
			Opus_silk_gains_dequant(nil, &decoded[0], &indices[0], &dp, cond, 4)
			if gains != decoded || p != dp {
				t.Fatal(previous, cond, gains, decoded, p, dp)
			}
		}
	}
}

// From gains_dequant_pointers_test.go.
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
