//go:build compareopus && cgo

package main

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/kazzmir/opus-go/opuscc"
)

func TestSortAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, n := range []int{1, 10, 16, 32} {
		for k := 1; k <= n; k++ {
			for trial := 0; trial < 50; trial++ {
				a := make([]int32, n)
				a16 := make([]int16, n)
				for i := range a {
					a[i] = int32(rng.Uint32())
					a16[i] = int16(a[i])
					if trial%2 == 0 {
						a[i] %= 5
						a16[i] %= 5
					} // Repeated values exercise stable indices.
				}
				c, c16 := slices.Clone(a), slices.Clone(a16)
				idx, cidx := make([]int32, k), make([]int32, k)
				opuscc.Opus_silk_insertion_sort_increasing(nil, &a[0], &idx[0], int32(n), int32(k))
				opuscc.Opus_silk_insertion_sort_increasing_all_values_int16(nil, &a16[0], int32(n))
				nativeSort32(c, cidx)
				nativeSort16(c16)
				if !slices.Equal(a, c) || !slices.Equal(idx, cidx) || !slices.Equal(a16, c16) {
					t.Fatalf("n=%d k=%d trial=%d: sorting differs from C", n, k, trial)
				}
			}
		}
	}
}
