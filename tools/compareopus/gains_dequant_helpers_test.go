//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"testing"
)

func TestGainsDequantAgainstC(t *testing.T) {
	for previous := 0; previous < 64; previous++ {
		for index := -128; index < 128; index++ {
			for _, conditional := range []int32{0, 1} {
				indices := [4]int8{int8(index), 0, 40, 4}
				g, c := [4]int32{}, [4]int32{}
				gs, cs := int8(previous), int8(previous)
				opuscc.Opus_silk_gains_dequant(nil, &g[0], &indices[0], &gs, conditional, 4)
				nativeGainsDequant(c[:], indices[:], &cs, conditional)
				if g != c || gs != cs {
					t.Fatalf("prev=%d index=%d conditional=%d Go=%v/%d C=%v/%d", previous, index, conditional, g, gs, c, cs)
				}
			}
		}
	}
}
