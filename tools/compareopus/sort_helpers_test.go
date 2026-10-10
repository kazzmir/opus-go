//go:build compareopus && cgo

package main

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/kazzmir/opus-go/opuscc"
	"github.com/kazzmir/opus-go/opusccenc"
	"unsafe"
)

func TestEncoderSort32AliasAgainstC(t *testing.T) {
	for _, n := range []int{1, 6, 16, 32} {
		for k := 1; k <= n; k++ {
			for _, shift := range []int{-1, 0, 1} {
				g := make([]int32, n+3)
				for i := range g {
					g[i] = int32(i%7 - 3)
				}
				c, decoder := slices.Clone(g), slices.Clone(g)
				opuscc.Opus_silk_insertion_sort_increasing(nil, &decoder[1], &decoder[1+shift], int32(n), int32(k))
				opusccenc.Opus_silk_insertion_sort_increasing(nil, uintptr(unsafe.Pointer(&g[1])), uintptr(unsafe.Pointer(&g[1+shift])), int32(n), int32(k))
				nativeSort32(c[1:1+n], c[1+shift:1+shift+k])
				if !slices.Equal(g, c) || !slices.Equal(decoder, c) {
					t.Fatal("sort32 value/index alias", n, k, shift)
				}
			}
		}
	}
}

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
				e, e16 := slices.Clone(a), slices.Clone(a16)
				opusccenc.Opus_silk_insertion_sort_increasing_all_values_int16(nil, uintptr(unsafe.Pointer(&e16[0])), int32(n))
				eidx := make([]int32, k)
				opusccenc.Opus_silk_insertion_sort_increasing(nil, uintptr(unsafe.Pointer(&e[0])), uintptr(unsafe.Pointer(&eidx[0])), int32(n), int32(k))
				idx, cidx := make([]int32, k), make([]int32, k)
				opuscc.Opus_silk_insertion_sort_increasing(nil, &a[0], &idx[0], int32(n), int32(k))
				opuscc.Opus_silk_insertion_sort_increasing_all_values_int16(nil, &a16[0], int32(n))
				nativeSort32(c, cidx)
				nativeSort16(c16)
				if !slices.Equal(a, c) || !slices.Equal(idx, cidx) || !slices.Equal(a16, c16) || !slices.Equal(e, c) || !slices.Equal(eidx, cidx) || !slices.Equal(e16, c16) {
					t.Fatalf("n=%d k=%d trial=%d: sorting differs from C", n, k, trial)
				}
			}
		}
	}
}
