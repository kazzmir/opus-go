package opuscc

import "testing"

func TestMatrixInt24Pointers(t *testing.T) {
	var owner mappingTestBuffer
	coef := int16(-32768)
	Opus_mapping_matrix_init(nil, &owner.Header, 1, 1, 0, &coef, 2)
	Opus_mapping_matrix_multiply_channel_out_int24(nil, &owner.Header, nil, 0, 1, nil, 1, 0)
	in := []float32{1.25, -1.25, 0.5 / 8388608, 1.5 / 8388608, -1.5 / 8388608, 255, -256}
	samples := []int32{10485760, -10485760, 0, 2, -2, 2139095040, -2147483648}
	out := make([]int32, len(in)+2)
	for i := range out {
		out[i] = 2147483640
	}
	out[0] = 77
	out[len(out)-1] = 88
	Opus_mapping_matrix_multiply_channel_out_int24(nil, &owner.Header, &in[0], 0, 1, &out[1], 1, int32(len(in)))
	for i, sample := range samples {
		want := int32(int64(2147483640) - int64(sample))
		if out[i+1] != want {
			t.Fatalf("i=%d got=%d want=%d", i, out[i+1], want)
		}
	}
	if out[0] != 77 || out[len(out)-1] != 88 {
		t.Fatal("guards")
	}
	coef = 16384
	Opus_mapping_matrix_init(nil, &owner.Header, 1, 1, 0, &coef, 2)
	in = []float32{1.0 / 8388608, -1.0 / 8388608}
	var rounded [2]int32
	Opus_mapping_matrix_multiply_channel_out_int24(nil, &owner.Header, &in[0], 0, 1, &rounded[0], 1, 2)
	if rounded != [2]int32{1, 0} {
		t.Fatal("Q15 half rounding", rounded)
	}
}
