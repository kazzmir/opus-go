package opuscc

import (
	"math"
	"slices"
	"testing"
	"unsafe"
)

func TestBestPitchOwnedAliasesPointers(t *testing.T) {
	for _, origin := range []int{1, 2, 8, 9} {
		a := [15]float32{77, 1, 9, 3, 8, 0, 0, 0, 1, 2, 3, 4, 5, 6, 88}
		b := a
		oldBest := (*[2]int32)(unsafe.Pointer(&a[origin]))
		ownedBest := (*[2]int32)(unsafe.Pointer(&b[origin]))
		find_best_pitch(nil, &a[1], &a[8], 2, 4, oldBest)
		findBestPitchOwned(nil, b[1:5], b[8:14], 2, 4, ownedBest)
		for i := range a {
			if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
				t.Fatal("owned best pitch/live numeric aliases/guards", origin, i)
			}
		}
	}
	corr, short := [4]float32{}, [5]float32{}
	best := [2]int32{77, 88}
	if !validationPanics(func() { findBestPitchOwned(nil, corr[:], short[:], 2, 4, &best) }) || best != [2]int32{77, 88} {
		t.Fatal("terminal energy sample extent before best initialization")
	}
}

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
		spanResult, ownedResult := result, result
		find_best_pitch(nil, unsafe.SliceData(corr), unsafe.SliceData(y), tc.window, int32(len(corr)), &result.pitch)
		findBestPitchSpans(nil, corr, y, tc.window, &spanResult.pitch)
		findBestPitchOwned(nil, corr, y, tc.window, int32(len(corr)), &ownedResult.pitch)
		if spanResult != result || ownedResult != result {
			t.Fatal("best-pitch slice kernel and guards")
		}
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
