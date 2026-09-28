package opuscc

import (
	"math"
	"slices"
	"testing"
)

func TestStereoMergePointers(t *testing.T) {
	scale := float32(1) / float32(math.Sqrt(2))
	for _, tc := range []struct {
		x, y         []float32
		mid          float32
		wantX, wantY []float32
	}{
		{[]float32{1, 0}, []float32{0, 1}, 1, []float32{scale, -scale}, []float32{scale, scale}},
		{[]float32{1, 0}, []float32{0, 1}, 0, []float32{0, -1}, []float32{0, 1}},
		{[]float32{1, 0}, []float32{1, 0}, 1, []float32{1, 0}, []float32{1, 0}},
		{[]float32{1, 0}, []float32{-1, 0}, 1, []float32{1, 0}, []float32{1, 0}},
		{[]float32{0.25, 0.5}, []float32{0, 0}, 0.01, []float32{0.25, 0.5}, []float32{0.25, 0.5}},
	} {
		x := append([]float32{111}, tc.x...)
		x = append(x, 222)
		y := append([]float32{333}, tc.y...)
		y = append(y, 444)
		stereo_merge(nil, &x[1], &y[1], tc.mid, int32(len(tc.x)), 0)
		if !slices.Equal(x[1:len(x)-1], tc.wantX) || !slices.Equal(y[1:len(y)-1], tc.wantY) {
			t.Fatalf("mid=%g X=%v Y=%v", tc.mid, x, y)
		}
		if x[0] != 111 || x[len(x)-1] != 222 || y[0] != 333 || y[len(y)-1] != 444 {
			t.Fatal("sentinels changed")
		}
	}
	stereo_merge(nil, nil, nil, 1, 0, 0)
}
