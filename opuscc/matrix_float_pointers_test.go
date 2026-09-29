package opuscc

import "testing"

func TestMatrixFloatPointers(t *testing.T) {
	var owner mappingTestBuffer
	coefs := [6]int16{16384, -32768, 32767, 8192, -16384, 4096}
	Opus_mapping_matrix_init(nil, &owner.Header, 2, 3, 999, &coefs[0], 12)
	Opus_mapping_matrix_multiply_channel_out_float(nil, &owner.Header, nil, 0, 1, nil, 2, 0)
	in := [3]float32{0.5, 123, -0.25}
	out := [6]float32{77, 1, 2, 3, 4, 88}
	Opus_mapping_matrix_multiply_channel_out_float(nil, &owner.Header, &in[0], 2, 2, &out[1], 2, 2)
	if out != [6]float32{77, 0.75, 2.0625, 3.125, 3.96875, 88} {
		t.Fatal(out)
	}
	// Read each frame's sample before writing any rows, even when aliased.
	alias := [4]float32{0.5, 0.25, 1, 2}
	Opus_mapping_matrix_multiply_channel_out_float(nil, &owner.Header, &alias[0], 0, 1, &alias[0], 2, 2)
	if alias != [4]float32{0.75, -0.25, 0.875, 2.25} {
		t.Fatal("alias", alias)
	}
}
