//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
	"unsafe"
)

func TestChannelInt24AgainstC(t *testing.T) {
	// Keep float-to-int conversion within C's representable int32 domain; no 24-bit clip.
	src := make([]float32, 100)
	for i := range src {
		src[i] = []float32{-256, -2, -1, -.5, -.5 / 8388608, .5 / 8388608, 1.5 / 8388608, 2.5 / 8388608, 1, 2, 255, 255.99998}[i%12]
	}
	for _, n := range []int32{0, 1, 2, 17} {
		for _, ds := range []int32{0, 1, 3, 8} {
			for _, ss := range []int32{0, 1, 2} {
				for _, dc := range []int32{0, 2} {
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
						opuscc.CompareChannelInt24(&g[1], ds, dc, unsafe.SliceData(input), ss, n)
						nativeChannelOutput(unsafe.Pointer(&c[1]), ds, dc, input, ss, n, 2)
						if !slices.Equal(g, c) {
							t.Fatal(n, ds, ss, dc, zero, g, c)
						}
					}
				}
			}
		}
	}
}
