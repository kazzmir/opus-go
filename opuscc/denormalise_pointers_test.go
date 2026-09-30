package opuscc

import (
	"runtime"
	"testing"
)

var denormalisePointerBands = [4]int16{0, 2, 4, 6}

func TestNormaliseBandsPointers(t *testing.T) {
	Opus_normalise_bands(nil, nil, 0, 0, nil, nil, nil, 0, 0, 0)
	bands := [3]int16{1, 2, 3}
	in := [8]float32{1, 2, 3, 4, 5, 6, 7, 8}
	energy := [4]float32{2, 4, 2, 4}
	out := [10]float32{77, 9, 9, 9, 9, 9, 9, 9, 9, 88}
	entropyInitGrowStack(12)
	runtime.GC()
	Opus_normalise_bands(nil, &bands[0], 4, 2, &in[0], &out[1], &energy[0], 2, 2, 1)
	if out != [10]float32{77, 9, 1, .75, 9, 9, 3, 1.75, 9, 88} {
		t.Fatal(out)
	}
	// C also visits channel zero for C=0; only selected bands are overwritten.
	Opus_normalise_bands(nil, &bands[0], 4, 2, &in[0], &out[1], &energy[0], 2, 0, 1)
	if out[2] != 1 || out[3] != .75 || out[0] != 77 || out[9] != 88 {
		t.Fatal("zero channel contract", out)
	}
}

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
