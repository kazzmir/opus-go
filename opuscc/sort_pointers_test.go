package opuscc

import (
	"slices"
	"sort"
	"testing"
)

func TestSortPointers(t *testing.T) {
	for _, input := range [][]int32{{7}, {3, -1, 3, -1, 0, 2147483647, -2147483648}, {4, 3, 2, 1}, {1, 1, 1, 1}} {
		order := make([]int, len(input))
		for i := range order {
			order[i] = i
		}
		sort.SliceStable(order, func(i, j int) bool { return input[order[i]] < input[order[j]] })
		for k := 1; k <= len(input); k++ {
			a := append([]int32{111}, input...)
			a = append(a, 222)
			idx := make([]int32, k+2)
			idx[0], idx[k+1] = 333, 444
			Opus_silk_insertion_sort_increasing(nil, &a[1], &idx[1], int32(len(input)), int32(k))
			for i := 0; i < k; i++ {
				if a[i+1] != input[order[i]] || idx[i+1] != int32(order[i]) {
					t.Fatalf("input=%v k=%d i=%d values=%v indices=%v", input, k, i, a, idx)
				}
			}
			if !slices.Equal(a[k+1:len(a)-1], input[k:]) || a[0] != 111 || a[len(a)-1] != 222 || idx[0] != 333 || idx[k+1] != 444 {
				t.Fatal("tail or sentinel changed")
			}
		}
	}
	for _, input := range [][]int16{{7}, {-32768, 32767, -1, 0, -1}, {3, 2, 1}} {
		want := slices.Clone(input)
		slices.Sort(want)
		a := append([]int16{111}, input...)
		a = append(a, 222)
		Opus_silk_insertion_sort_increasing_all_values_int16(nil, &a[1], int32(len(input)))
		if !slices.Equal(a[1:len(a)-1], want) || a[0] != 111 || a[len(a)-1] != 222 {
			t.Fatalf("input=%v output=%v", input, a)
		}
	}
}
