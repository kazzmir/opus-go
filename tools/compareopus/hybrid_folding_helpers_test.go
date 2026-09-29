//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"slices"
	"testing"
)

func TestHybridFoldingAgainstC(t *testing.T) {
	for _, M := range []int32{1, 2, 4, 8} {
		for width := int16(1); width <= 24; width++ {
			for next := width; next <= 2*width; next++ {
				for dual := int32(0); dual <= 1; dual++ {
					bands := []int16{0, 3, 3 + width, 3 + width + next}
					n := int(M) * int(next)
					g1, g2 := make([]float32, n+2), make([]float32, n+2)
					for i := range g1 {
						g1[i] = math.Float32frombits(uint32(i)*99991 + 0x7fc00000)
						g2[i] = -float32(i)
					}
					c1, c2 := slices.Clone(g1), slices.Clone(g2)
					opuscc.CompareHybridFolding(bands, g1[1:n+1], g2[1:n+1], 1, M, dual)
					nativeHybridFolding(bands, c1[1:n+1], c2[1:n+1], 1, M, dual)
					for i := range g1 {
						if math.Float32bits(g1[i]) != math.Float32bits(c1[i]) || math.Float32bits(g2[i]) != math.Float32bits(c2[i]) {
							t.Fatalf("M=%d width=%d next=%d dual=%d i=%d", M, width, next, dual, i)
						}
					}
				}
			}
		}
	}
}
