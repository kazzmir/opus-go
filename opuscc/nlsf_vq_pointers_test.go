package opuscc

import "testing"

func TestNLSFVQPointers(t *testing.T) {
	input := [2]int16{100, 200}
	codebook := [4]uint8{0, 1, 0, 0}
	weights := [4]int16{2, 3, 4, 5}
	beforeInput, beforeCB, beforeWeights := input, codebook, weights
	output := [4]int32{111, 0, 0, 222}
	Opus_silk_NLSF_VQ(nil, &output[1], &input[0], &codebook[0], &weights[0], 2, 2)
	if output != [4]int32{111, 308, 1100, 222} {
		t.Fatalf("output=%v", output)
	}
	if input != beforeInput || codebook != beforeCB || weights != beforeWeights {
		t.Fatal("input changed")
	}
	input = [2]int16{-32768, 32767}
	codebook = [4]uint8{255, 0}
	weights = [4]int16{1, 1}
	Opus_silk_NLSF_VQ(nil, &output[1], &input[0], &codebook[0], &weights[0], 1, 2)
	if output[1] != 49022 {
		t.Fatalf("signed narrowing: got %d", output[1])
	}
	Opus_silk_NLSF_VQ(nil, nil, nil, nil, nil, 0, 16)
	Opus_silk_NLSF_VQ(nil, &output[1], nil, nil, nil, 2, 0)
	if output != [4]int32{111, 0, 0, 222} {
		t.Fatal("zero-order output differs")
	}
}
