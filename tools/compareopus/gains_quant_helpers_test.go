//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
)

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
