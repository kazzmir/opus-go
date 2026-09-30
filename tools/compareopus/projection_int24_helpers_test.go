//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
	"unsafe"
)

func TestProjectionInt24AgainstC(t *testing.T) {
	matrix := newTestMapping(4, 3, 77, []int16{16384, -8192, 32767, 8192, 4096, -32768, 8192, 4096, 12345, -23456, 777, 888})
	// C float-to-int conversion is compared only in the representable int32 domain.
	src := make([]float32, 100)
	for i := range src {
		src[i] = []float32{0, .5, -.5, 2, -2, .5 / 8388608, 1.5 / 8388608, -256, 255.99998}[i%9]
	}
	for _, n := range []int32{0, 1, 3, 17} {
		for _, ds := range []int32{2, 4} {
			for _, ss := range []int32{0, 1, 3} {
				for _, zero := range []bool{false, true} {
					g := make([]int32, 256)
					for i := range g {
						g[i] = 77
					}
					c := slices.Clone(g)
					input := src
					if zero {
						input = nil
					}
					for dc := int32(0); dc < 3; dc++ {
						opuscc.CompareProjectionInt24(&g[1], ds, dc, unsafe.SliceData(input), ss, n, matrix)
						nativeProjectionOutput(unsafe.Pointer(&c[1]), ds, dc, input, ss, n, unsafe.Pointer(matrix), 2)
						if !slices.Equal(g, c) {
							t.Fatal(n, ds, ss, zero, dc)
						}
					}
				}
			}
		}
	}
}
