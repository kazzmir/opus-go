package opuscc

import (
	"math"
	"runtime"
	"slices"
	"testing"
)

func TestFIRPointers(t *testing.T) {
	input := [12]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	coeff := [4]float32{1, 2, 3, 4}
	out := [10]float32{}
	out[0] = 77
	out[9] = 88
	entropyInitGrowStack(12)
	runtime.GC()
	Opus_celt_fir_c(nil, &input[4], &coeff[0], &out[1], 8, 4, 0)
	for i := 0; i < 8; i++ {
		want := input[4+i]
		for j := 0; j < 4; j++ {
			want += float32(coeff[3-j] * input[i+j])
		}
		if out[i+1] != want {
			t.Fatal(i, out, want)
		}
	}
	if out[0] != 77 || out[9] != 88 {
		t.Fatal("guard")
	}
	// Tail-only, zero-order filtering requires no coefficients or input history.
	Opus_celt_fir_c(nil, &input[0], nil, &out[1], 3, 0, 0)
	for i := 0; i < 3; i++ {
		if out[i+1] != input[i] {
			t.Fatal("identity")
		}
	}
}

func TestLPCPointers(t *testing.T) {
	for _, tc := range []struct{ ac, want []float32 }{
		{[]float32{0, 1, 2}, []float32{0, 0}},
		{[]float32{1e-10, 1, 2}, []float32{0, 0}},
		{[]float32{float32(math.NaN()), 1, 2}, []float32{0, 0}},
		{[]float32{1, 0.5}, []float32{-0.5}},
		{[]float32{1, 0.5, 0.25}, []float32{-0.5, 0}},
		{[]float32{1, 1, 1, 1}, []float32{-1, 0, 0}},
	} {
		out := make([]float32, len(tc.want)+2)
		for i := range out {
			out[i] = 123
		}
		Opus__celt_lpc(nil, &out[1], &tc.ac[0], int32(len(tc.want)))
		if !slices.Equal(out[1:len(out)-1], tc.want) || out[0] != 123 || out[len(out)-1] != 123 {
			t.Fatalf("ac=%v got=%v want=%v", tc.ac, out, tc.want)
		}
	}
	ac := float32(1)
	Opus__celt_lpc(nil, nil, &ac, 0)
}
