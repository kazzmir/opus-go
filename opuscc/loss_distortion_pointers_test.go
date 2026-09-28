package opuscc

import (
	"math"
	"slices"
	"testing"
)

func TestLossDistortionPointers(t *testing.T) {
	current := []float32{999, 3, 4, 999, 999, 999, 6, 8}
	previous := make([]float32, len(current))
	before := slices.Clone(current)
	for _, tc := range []struct {
		start, end, stride, channels int32
		want                         float32
	}{
		{1, 3, 5, 1, 25}, {1, 3, 5, 2, 125}, {1, 2, 5, 2, 45}, {1, 3, 5, 0, 25},
	} {
		got := loss_distortion(nil, &current[0], &previous[0], tc.start, tc.end, tc.stride, tc.channels)
		if got != tc.want {
			t.Fatalf("case=%+v got=%g", tc, got)
		}
	}
	if !slices.Equal(current, before) {
		t.Fatal("input changed")
	}
	if got := loss_distortion(nil, nil, nil, 2, 2, 5, 2); got != 0 {
		t.Fatal("nonzero empty result")
	}
	for _, v := range []float32{20, float32(math.Inf(1)), float32(math.NaN())} {
		zero := float32(0)
		got := loss_distortion(nil, &v, &zero, 0, 1, 1, 1)
		if math.IsNaN(float64(v)) {
			if !math.IsNaN(float64(got)) {
				t.Fatal("NaN was lost")
			}
		} else if got != 200 {
			t.Fatalf("cap: got %g", got)
		}
	}
	// Float32 addition must round after each term, rather than accumulate in float64.
	a := []float32{8, 0.001, 0.001}
	b := make([]float32, len(a))
	var want float32
	for _, v := range a {
		want = want + float32(v*v)
	}
	if got := loss_distortion(nil, &a[0], &b[0], 0, 3, 3, 1); math.Float32bits(got) != math.Float32bits(want) {
		t.Fatal("accumulation order changed")
	}
}
