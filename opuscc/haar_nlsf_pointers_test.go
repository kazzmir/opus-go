package opuscc

import (
	"math"
	"testing"
)

func TestHaarPointers(t *testing.T) {
	for _, n := range []int{0, 1, 2, 3, 4, 7, 16} {
		for _, stride := range []int{1, 2, 4} {
			x := make([]float32, n*stride+2)
			for i := range x {
				x[i] = float32(i-5) * 0.125
			}
			want := append([]float32(nil), x...)
			// Each adjacent pair of rows is transformed independently.
			for row := 0; row+1 < n; row += 2 {
				for col := 0; col < stride; col++ {
					a := 1 + row*stride + col
					b := a + stride
					left := float32(float32(0.70710678) * want[a])
					right := float32(float32(0.70710678) * want[b])
					want[a], want[b] = left+right, left-right
				}
			}
			Opus_haar1(nil, &x[1], int32(n), int32(stride))
			for i := range x {
				if x[i] != want[i] {
					t.Fatalf("n=%d stride=%d index=%d: got %g want %g", n, stride, i, x[i], want[i])
				}
			}
		}
	}
	Opus_haar1(nil, nil, 1, 2)
	Opus_haar1(nil, nil, 4, 0)
}

func TestChannelWeightsPointers(t *testing.T) {
	for _, tc := range []struct {
		x, y float32
		want [2]float32
	}{
		{0, 0, [2]float32{0, 0}},
		{3, 9, [2]float32{4, 10}},
		{9, 3, [2]float32{10, 4}},
		{6, 6, [2]float32{8, 8}},
		{0, 12, [2]float32{0, 12}},
	} {
		var got [2]float32
		compute_channel_weights(nil, tc.x, tc.y, &got)
		if got != tc.want {
			t.Fatalf("energies %g,%g: got %v want %v", tc.x, tc.y, got, tc.want)
		}
	}
	// Preserve MIN32's comparison behavior, rather than Go min's NaN rules.
	var got [2]float32
	compute_channel_weights(nil, float32(math.NaN()), 3, &got)
	if !math.IsNaN(float64(got[0])) || got[1] != 4 {
		t.Fatalf("NaN left energy: %v", got)
	}
}

func TestNLSFResidualPointers(t *testing.T) {
	indices := [8]int8{-128, 127, 0, -1, 1, 7, -8, 2}
	pred := [8]uint8{255, 254, 128, 1, 0, 200, 80, 255}
	// Golden vectors use the C silk_SMULBB/SMLAWB signed narrowing and
	// arithmetic shifts, including the negative low-word quantization step.
	for _, tc := range []struct {
		step int32
		want [8]int16
	}{
		{16384, [8]int16{-500, 32370, -116, -231, 230, 303, -1872, 486}},
		{65535, [8]int16{-2, -3, -1, -1, -1, -2, -1, -1}},
	} {
		out := [10]int16{123, 0, 0, 0, 0, 0, 0, 0, 0, 456}
		silk_NLSF_residual_dequant(nil, &out[1], &indices[0], &pred[0], tc.step, 8)
		if [8]int16(out[1:9]) != tc.want || out[0] != 123 || out[9] != 456 {
			t.Fatalf("step %d: got %v want %v with sentinels", tc.step, out, tc.want)
		}
	}
	silk_NLSF_residual_dequant(nil, nil, nil, nil, 16384, 0)
}
