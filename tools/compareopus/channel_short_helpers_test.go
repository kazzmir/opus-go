//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"slices"
	"testing"
	"unsafe"
)

func TestChannelShortAgainstC(t *testing.T) {
	src := make([]float32, 100)
	for i := range src {
		src[i] = []float32{-2, -1, -.5, -.5 / 32768, .5 / 32768, 1.5 / 32768, 2.5 / 32768, 1, 2, math.Float32frombits(0x7fc01234), float32(math.Inf(-1)), float32(math.Inf(1))}[i%12]
	}
	for _, n := range []int32{0, 1, 2, 17} {
		for _, ds := range []int32{0, 1, 3, 8} {
			for _, ss := range []int32{0, 1, 2} {
				for _, dc := range []int32{0, 2} {
					for _, zero := range []bool{false, true} {
						g := make([]int16, 256)
						for i := range g {
							g[i] = 77
						}
						c := slices.Clone(g)
						input := src
						if zero {
							input = nil
						}
						opuscc.CompareChannelShort(&g[1], ds, dc, unsafe.SliceData(input), ss, n)
						nativeChannelOutput(unsafe.Pointer(&c[1]), ds, dc, input, ss, n, 1)
						if !slices.Equal(g, c) {
							t.Fatal(n, ds, ss, dc, zero, g, c)
						}
					}
				}
			}
		}
	}
}
