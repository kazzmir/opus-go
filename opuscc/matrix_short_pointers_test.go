package opuscc

import (
	"math"
	"slices"
	"testing"
)

func TestMatrixShortPointers(t *testing.T) {
	var owner mappingTestBuffer
	coef := int16(-32768)
	Opus_mapping_matrix_init(nil, &owner.Header, 1, 1, 0, &coef, 2)
	Opus_mapping_matrix_multiply_channel_out_short(nil, &owner.Header, nil, 0, 1, nil, 1, 0)
	in := []float32{float32(math.NaN()), float32(math.Inf(1)), float32(math.Inf(-1)), 0.5 / 32768, 1.5 / 32768, -1.5 / 32768, 1.2, -1.2}
	out := make([]int16, len(in)+2)
	for i := range out {
		out[i] = 10
	}
	out[0] = 77
	out[len(out)-1] = 88
	Opus_mapping_matrix_multiply_channel_out_short(nil, &owner.Header, &in[0], 0, 1, &out[1], 1, int32(len(in)))
	want := []int16{77, -32758, -32757, -32758, 10, 8, 12, -32757, -32758, 88}
	if !slices.Equal(out, want) {
		t.Fatal(out)
	}
}
