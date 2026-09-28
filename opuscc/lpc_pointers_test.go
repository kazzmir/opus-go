package opuscc

import (
	"math"
	"slices"
	"testing"
)

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
