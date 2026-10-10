package opusccenc

import (
	"github.com/kazzmir/opus-go/opuscc"
	"runtime"
	"slices"
	"testing"
)

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
