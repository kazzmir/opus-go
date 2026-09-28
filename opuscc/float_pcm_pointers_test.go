package opuscc

import (
	"math"
	"slices"
	"testing"
)

func TestFloatPCMPointers(t *testing.T) {
	input := []float32{-2, -1, -32767.5 / 32768, -2.5 / 32768, -1.5 / 32768, -0.5 / 32768, 0, 0.5 / 32768, 1.5 / 32768, 2.5 / 32768, 32766.5 / 32768, 1, 2, float32(math.Inf(-1)), float32(math.Inf(1))}
	want := []int16{-32768, -32768, -32768, -2, -2, 0, 0, 0, 2, 2, 32766, 32767, 32767, -32768, 32767}
	before := slices.Clone(input)
	output := make([]int16, len(input)+2)
	output[0], output[len(output)-1] = 111, 222
	Opus_celt_float2int16_c(nil, &input[0], &output[1], int32(len(input)))
	if !slices.Equal(output[1:len(output)-1], want) || output[0] != 111 || output[len(output)-1] != 222 {
		t.Fatalf("got=%v want=%v", output, want)
	}
	if !slices.Equal(input, before) {
		t.Fatal("input changed")
	}
	Opus_celt_float2int16_c(nil, nil, nil, 0)
	Opus_celt_float2int16_c(nil, nil, nil, -1)
}
