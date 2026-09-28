package opusccenc

import (
	"math"
	"slices"
	"testing"
)

func TestFloatPCMPointers(t *testing.T) {
	input := []float32{-2, -1, -2.5 / 32768, -1.5 / 32768, -0.5 / 32768, 0, 0.5 / 32768, 1.5 / 32768, 2.5 / 32768, 1, 2, float32(math.NaN()), float32(math.Inf(-1)), float32(math.Inf(1))}
	want := []int16{-32768, -32768, -2, -2, 0, 0, 0, 2, 2, 32767, 32767, -32768, -32768, 32767}
	out := make([]int16, len(input)+2)
	out[0], out[len(out)-1] = 111, 222
	Opus_celt_float2int16_c(nil, &input[0], &out[1], int32(len(input)))
	if !slices.Equal(out[1:len(out)-1], want) || out[0] != 111 || out[len(out)-1] != 222 {
		t.Fatalf("got=%v want=%v", out, want)
	}
	Opus_celt_float2int16_c(nil, nil, nil, 0)
	Opus_celt_float2int16_c(nil, nil, nil, -1)
}
