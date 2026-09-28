package opuscc

import (
	"slices"
	"testing"
)

func TestLPCFitPointers(t *testing.T) {
	input := []int32{111, -7, -5, -3, -1, 1, 3, 5, 7, 222}
	before := slices.Clone(input)
	output := [10]int16{123, 0, 0, 0, 0, 0, 0, 0, 0, 234}
	Opus_silk_LPC_fit(nil, &output[1], &input[1], 12, 13, 8)
	if !slices.Equal(output[:], []int16{123, -3, -2, -1, 0, 1, 2, 3, 4, 234}) || !slices.Equal(input, before) {
		t.Fatal("rounding or no-expansion input differs")
	}
	// A large late coefficient reaches C's ten-iteration clipping fallback.
	for _, sign := range []int32{-1, 1} {
		var high [26]int32
		high[0], high[25] = 111, 222
		high[24] = sign * 2000000000
		var out [26]int16
		out[0], out[25] = 123, 234
		Opus_silk_LPC_fit(nil, &out[1], &high[1], 12, 13, 24)
		expected := int16(32767)
		if sign < 0 {
			expected = -32768
		}
		if out[24] != expected || high[24] != int32(expected)<<1 || high[0] != 111 || high[25] != 222 || out[0] != 123 || out[25] != 234 {
			t.Fatalf("sign=%d out=%v input=%v", sign, out, high)
		}
	}
	Opus_silk_LPC_fit(nil, nil, nil, 12, 13, 0)
}
