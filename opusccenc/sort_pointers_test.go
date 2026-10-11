package opusccenc

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"runtime"
	"slices"
	"testing"
)

func TestEncoderGainsIDPointers(t *testing.T) {
	for _, count := range []int32{1, 2, 4, 8, 16} {
		indices := make([]int8, count+2)
		for i := range indices {
			indices[i] = int8(i*73 - 128)
		}
		before := slices.Clone(indices)
		want := opuscc.Opus_silk_gains_ID(nil, &indices[1], count)
		runtime.GC()
		if got := silkGainsID(&indices[1], count); got != want || !slices.Equal(indices, before) {
			t.Fatal("signed/wrapping gains ID/guards", count, got, want)
		}
	}
	for _, count := range []int32{0, -1} {
		if silkGainsID(nil, count) != 0 {
			t.Fatal("empty gains ID", count)
		}
	}
}

func TestEncoderSortFloatPointers(t *testing.T) {
	cases := []struct {
		values, want []float32
		idx          []int32
		k            int32
	}{
		{[]float32{3, 1, 3, -2, 4, 4, 0}, []float32{4, 4, 3, -2, 4, 4, 0}, []int32{4, 5, 0}, 3},
		{[]float32{math.Float32frombits(0x7fc12345), 3, 1, 4}, []float32{math.Float32frombits(0x7fc12345), 4, 3, 1}, []int32{0, 3, 1, 2}, 4},
		{[]float32{0, float32(math.Copysign(0, -1)), float32(math.Inf(1)), float32(math.Inf(-1)), math.Float32frombits(0x7fc12345)}, []float32{float32(math.Inf(1)), 0, float32(math.Copysign(0, -1)), float32(math.Inf(-1)), math.Float32frombits(0x7fc12345)}, []int32{2, 0, 1, 3, 4}, 5},
	}
	for _, tc := range cases {
		data := append([]float32{77}, tc.values...)
		data = append(data, 88)
		idx := make([]int32, tc.k+2)
		idx[0], idx[len(idx)-1] = 99, 111
		runtime.GC()
		silkInsertionSortDecreasingFLP(nil, &data[1], &idx[1], int32(len(tc.values)), tc.k)
		if !flpVectorBitsEqual(data[1:len(data)-1], tc.want) || !slices.Equal(idx[1:len(idx)-1], tc.idx) || data[0] != 77 || data[len(data)-1] != 88 || idx[0] != 99 || idx[len(idx)-1] != 111 {
			t.Fatal("float sort/ties/IEEE/guards")
		}
	}
	index := int32(99)
	silkInsertionSortDecreasingFLP(nil, nil, &index, 1, 1)
	if index != 0 {
		t.Fatal("float singleton index")
	}
}

func TestEncoderSort16Pointers(t *testing.T) {
	for _, n := range []int32{1, 6, 16, 32} {
		data := make([]int16, n+2)
		for i := range data {
			data[i] = int16(i*7919 - 32768)
		}
		want := slices.Clone(data)
		opuscc.Opus_silk_insertion_sort_increasing_all_values_int16(nil, &want[1], n)
		runtime.GC()
		silkInsertionSortInt16(nil, &data[1], n)
		if !slices.Equal(data, want) {
			t.Fatal("sort16/guards", n)
		}
	}
	silkInsertionSortInt16(nil, nil, 1)
}

func TestEncoderSort32Pointers(t *testing.T) {
	for _, n := range []int32{1, 6, 16, 32} {
		for k := int32(1); k <= n; k++ {
			for _, shift := range []int32{-1, 0, 1} {
				data := make([]int32, n+3)
				for i := range data {
					data[i] = int32(i%7 - 3)
				}
				want := slices.Clone(data)
				opuscc.Opus_silk_insertion_sort_increasing(nil, &want[1], &want[1+shift], n, k)
				runtime.GC()
				silkInsertionSortIncreasing(nil, &data[1], &data[1+shift], n, k)
				if !slices.Equal(data, want) {
					t.Fatal("sort32 live aliases/guards", n, k, shift)
				}
			}
		}
	}
	index := int32(99)
	silkInsertionSortIncreasing(nil, nil, &index, 1, 1)
	if index != 0 {
		t.Fatal("singleton must initialize index without consuming values")
	}
}
