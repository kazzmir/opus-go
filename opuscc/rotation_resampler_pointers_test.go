package opuscc

import "testing"

func TestExpRotationPointers(t *testing.T) {
	for _, tc := range []struct {
		stride int32
		want   [6]float32
	}{
		{1, [6]float32{-6, -2, -3, -4, -5, 1}},
		{2, [6]float32{5, 6, -3, -4, 1, 2}},
		{3, [6]float32{-4, -5, -6, 1, 2, 3}},
		{6, [6]float32{1, 2, 3, 4, 5, 6}},
	} {
		x := [8]float32{99, 1, 2, 3, 4, 5, 6, 88}
		exp_rotation1(nil, &x[1], 6, tc.stride, 0, 1)
		if [6]float32(x[1:7]) != tc.want || x[0] != 99 || x[7] != 88 {
			t.Fatalf("stride %d: got %v, want %v with sentinels", tc.stride, x, tc.want)
		}
	}
	exp_rotation1(nil, nil, 0, 1, 0, 1)
}

func TestHysteresisPointers(t *testing.T) {
	thresholds := [2]float32{10, 20}
	margins := [2]float32{2, 3}
	for _, tc := range []struct {
		value      float32
		prev, want int32
	}{
		{5, 0, 0}, {11, 0, 0}, {12, 0, 1},
		{9, 1, 1}, {8, 1, 0}, {22, 1, 1}, {23, 1, 2},
		{18, 2, 2}, {17, 2, 1}, {30, 0, 2}, {0, 2, 0},
	} {
		if got := Opus_hysteresis_decision(nil, tc.value, &thresholds[0], &margins[0], 2, tc.prev); got != tc.want {
			t.Fatalf("value %v prev %d: got %d, want %d", tc.value, tc.prev, got, tc.want)
		}
	}
	if got := Opus_hysteresis_decision(nil, 0, nil, nil, 0, 0); got != 0 {
		t.Fatalf("empty thresholds: %d", got)
	}
}

func TestResamplerAR2Pointers(t *testing.T) {
	input := [...]int16{32767, -32768, 0, 1, -1, 12000, -16000}
	coefs := [2]int16{-12345, 23456}
	initial := [2]int32{123456, -98765}
	want := [7]int32{8511808, -14900842, 23413321, -38973841, 62885175, -100107195, 161361640}
	wantState := [2]int32{-264900136, 231011879}
	// Chunked processing must produce exactly the same state and samples.
	for _, split := range []int{0, 1, 3, 6} {
		state := initial
		out := [9]int32{123, 0, 0, 0, 0, 0, 0, 0, 456}
		Opus_silk_resampler_private_AR2(nil, &state[0], &out[1], &input[0], &coefs[0], int32(split))
		Opus_silk_resampler_private_AR2(nil, &state[0], &out[split+1], &input[split], &coefs[0], int32(len(input)-split))
		if [7]int32(out[1:8]) != want || state != wantState || out[0] != 123 || out[8] != 456 {
			t.Fatalf("split %d: out=%v state=%v", split, out, state)
		}
	}
	Opus_silk_resampler_private_AR2(nil, nil, nil, nil, nil, 0)
}
