//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"math/rand"
	"slices"
	"testing"
)

func TestInterleaveAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(504))
	for _, stride := range []int32{1, 2, 3, 4, 8, 16} {
		for _, n0 := range []int32{0, 1, 2, 3, 7, 22, 37, 64} {
			for h := int32(0); h <= 1; h++ {
				if h != 0 && (stride == 1 || stride == 3) {
					continue
				}
				n := n0 * stride
				g := make([]float32, n+2)
				for i := range g {
					g[i] = math.Float32frombits(r.Uint32())
				}
				special := []uint32{0, 0x80000000, 0x7f800001, 0x7fc12345, 0xff800001, 0x7f800000, 0xff800000, 1, 0x80000001}
				for i, bits := range special {
					if i < int(n) {
						g[i+1] = math.Float32frombits(bits)
					}
				}
				c := slices.Clone(g)
				original := slices.Clone(g)
				opuscc.CompareInterleaveHadamard(&g[1], n0, stride, h)
				nativeInterleaveHadamard(c[1:], n0, stride, h)
				for i := range g {
					if math.Float32bits(g[i]) != math.Float32bits(c[i]) {
						t.Fatalf("n0=%d stride=%d h=%d i=%d", n0, stride, h, i)
					}
				}
				opuscc.CompareDeinterleaveHadamard(&g[1], n0, stride, h)
				nativeDeinterleaveHadamard(c[1:], n0, stride, h)
				for i := range original {
					bits := math.Float32bits(original[i])
					if math.Float32bits(g[i]) != bits || math.Float32bits(c[i]) != bits {
						t.Fatalf("round trip n0=%d stride=%d h=%d i=%d", n0, stride, h, i)
					}
				}
			}
		}
	}
}
