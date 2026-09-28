package opuscc

import (
	"testing"
)

var denormalisePointerBands = [4]int16{0, 2, 4, 6}

func TestDenormalisePointers(t *testing.T) {
	in := [6]float32{1, 2, 3, 4, 5, 6}
	logE := [3]float32{0, -100, 100}
	out := [10]float32{123, 9, 9, 9, 9, 9, 9, 9, 9, 456}
	Opus_denormalise_bands(nil, &denormalisePointerBands[0], 8, &in[0], &out[1], &logE[0], 1, 3, 1, 1, 0)
	if out[0] != 123 || out[9] != 456 || out[1] != 0 || out[2] != 0 || out[3] != 0 || out[4] != 0 || out[5] <= 0 || out[6] <= 0 || out[7] != 0 || out[8] != 0 {
		t.Fatalf("out=%v", out)
	}
	Opus_denormalise_bands(nil, &denormalisePointerBands[0], 8, nil, &out[1], nil, 1, 3, 1, 1, 1)
	for _, v := range out[1:9] {
		if v != 0 {
			t.Fatal("silence did not clear output")
		}
	}
	if out[0] != 123 || out[9] != 456 {
		t.Fatal("silence changed guards")
	}
}
