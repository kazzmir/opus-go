//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"slices"
	"testing"
	"unsafe"
)

func TestProjectionFloatAgainstC(t *testing.T) {
	matrix := newTestMapping(4, 3, 77, []int16{16384, -8192, 32767, 8192, 4096, -32768, 8192, 4096, 12345, -23456, 777, 888})
	for _, n := range []int32{0, 1, 3, 17} {
		for _, ds := range []int32{2, 4} {
			for _, ss := range []int32{0, 1, 3} {
				for variant := 0; variant < 3; variant++ {
					g := make([]float32, 256)
					for i := range g {
						g[i] = []float32{.5, -1, 2, math.Float32frombits(0x80000000)}[i%4]
					}
					c := slices.Clone(g)
					var sg, sc []float32
					if variant == 0 {
						sg = g[1:]
						sc = c[1:]
					} else if variant == 1 {
						sg = g[150:]
						sc = c[150:]
					}
					for dc := int32(0); dc < 3; dc++ {
						opuscc.CompareProjectionFloat(&g[2], ds, dc, unsafe.SliceData(sg), ss, n, matrix)
						nativeProjectionOutput(unsafe.Pointer(&c[2]), ds, dc, sc, ss, n, unsafe.Pointer(matrix), 0)
						if !sameFloatBits(g, c) {
							t.Fatal(n, ds, ss, variant, dc)
						}
					}
				}
			}
		}
	}
}
