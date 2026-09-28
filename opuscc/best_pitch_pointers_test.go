package opuscc

import (
	"math"
	"slices"
	"testing"
	"unsafe"
)

func TestBestPitchPointers(t *testing.T) {
	for _, tc := range []struct {
		corr, y []float32
		window  int32
		want    [2]int32
	}{
		{nil, nil, 0, [2]int32{0, 1}},
		{[]float32{-1, 0, -2}, make([]float32, 5), 2, [2]int32{0, 1}},
		{[]float32{1, 1, 1}, []float32{1, 1, 1, 1, 1}, 2, [2]int32{0, 1}},
		{[]float32{1, 9, 3, 8}, make([]float32, 6), 2, [2]int32{1, 3}},
		{[]float32{10, 2, 3, 1}, []float32{100, 0, 0, 0, 0}, 1, [2]int32{2, 1}},
		{[]float32{float32(math.NaN()), 4, 2}, make([]float32, 4), 1, [2]int32{1, 2}},
		{[]float32{5}, []float32{1, 2}, 1, [2]int32{0, 0}},
	} {
		result := struct {
			before int32
			pitch  [2]int32
			after  int32
		}{before: 111, after: 222}
		corr, y := slices.Clone(tc.corr), slices.Clone(tc.y)
		find_best_pitch(nil, unsafe.SliceData(corr), unsafe.SliceData(y), tc.window, int32(len(corr)), &result.pitch)
		if result.pitch != tc.want || result.before != 111 || result.after != 222 {
			t.Fatalf("corr=%v y=%v got=%v want=%v", corr, y, result.pitch, tc.want)
		}
		for i := range corr {
			if math.Float32bits(corr[i]) != math.Float32bits(tc.corr[i]) {
				t.Fatal("correlation input changed")
			}
		}
		if !slices.Equal(y, tc.y) {
			t.Fatal("signal input changed")
		}
	}
}
