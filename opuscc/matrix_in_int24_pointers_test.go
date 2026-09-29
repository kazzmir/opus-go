package opuscc

import "testing"

func TestMatrixInInt24Pointers(t *testing.T) {
	var owner mappingTestBuffer
	data := [6]int16{16384, 8192, -16384, 4096, 0, 32767}
	Opus_mapping_matrix_init(nil, &owner.Header, 2, 3, 0, &data[0], 12)
	in := [6]int32{10000 * 256, 20000 * 256, 30000 * 256, 4000 * 256, 5000 * 256, 6000 * 256}
	out := [5]float32{77, 9, 8, 7, 88}
	Opus_mapping_matrix_multiply_channel_in_int24(nil, &owner.Header, &in[0], 3, &out[1], 0, 2, 2)
	if out != [5]float32{77, -5000.0 / 32768, 8, -500.0 / 32768, 88} {
		t.Fatal(out)
	}
	Opus_mapping_matrix_multiply_channel_in_int24(nil, &owner.Header, nil, 3, nil, 0, 2, 0)
	Opus_mapping_matrix_multiply_channel_in_int24(nil, &owner.Header, nil, 0, &out[1], 0, 0, 3)
	if out[1] != 0 || out[0] != 77 || out[4] != 88 {
		t.Fatal("empty dot product", out)
	}
}
