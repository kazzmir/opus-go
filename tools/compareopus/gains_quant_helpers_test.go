//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"github.com/kazzmir/opus-go/opusccenc"
	"slices"
	"testing"
	"unsafe"
)

func TestEncoderGainsIDAgainstC(t *testing.T) {
	for _, count := range []int32{1, 2, 4, 8, 16} {
		for first := -128; first < 128; first++ {
			indices := make([]int8, count+2)
			for i := range indices {
				indices[i] = int8(first + i*73)
			}
			before := slices.Clone(indices)
			g := opusccenc.Opus_silk_gains_ID(nil, uintptr(unsafe.Pointer(&indices[1])), count)
			c := nativeGainsID(indices[1:], count)
			if g != c || !slices.Equal(indices, before) {
				t.Fatal("signed/wrapping gains ID", first, count, g, c)
			}
		}
	}
	for _, count := range []int32{0, -1} {
		if nativeGainsID(nil, count) != 0 || opusccenc.Opus_silk_gains_ID(nil, 0, count) != 0 {
			t.Fatal("empty gains ID", count)
		}
	}
}

func TestGainsQuantAgainstC(t *testing.T) {
	for prev := -128; prev < 128; prev++ {
		for cond := int32(0); cond < 2; cond++ {
			for _, nb := range []int{1, 2, 4} {
				for _, alias := range []bool{false, true} {
					g := []int32{1, 65536, 1 << 24, 2147483647}[:nb]
					c := slices.Clone(g)
					gi := []int8{77, 0, 0, 0, 0, 88}
					ci := slices.Clone(gi)
					gp, cp := int8(prev), int8(prev)
					pg, pc := &gp, &cp
					if alias {
						pg = &gi[1]
						pc = &ci[1]
						*pg = int8(prev)
						*pc = int8(prev)
					}
					opuscc.Opus_silk_gains_quant(nil, &gi[1], &g[0], pg, cond, int32(nb))
					nativeGainsQuant(ci[1:1+nb], c, pc, cond)
					if !slices.Equal(g, c) || !slices.Equal(gi, ci) || *pg != *pc {
						t.Fatal(prev, cond, nb, alias, g, c, gi, ci, *pg, *pc)
					}
				}
			}
		}
	}
}
