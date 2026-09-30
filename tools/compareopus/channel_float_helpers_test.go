//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"slices"
	"testing"
	"unsafe"
)

func TestChannelFloatAgainstC(t *testing.T) {
	for _, n := range []int32{0, 1, 2, 17} {
		for _, ds := range []int32{0, 1, 3, 8} {
			for _, ss := range []int32{0, 1, 2} {
				for _, dc := range []int32{0, 2} {
					for variant := 0; variant < 3; variant++ {
						g := make([]float32, 256)
						for i := range g {
							g[i] = math.Float32frombits([]uint32{0, 0x80000000, 0x7fc01234, 0x7f800000, 0xff800000, 0x3f800000, 1}[i%7])
						}
						c := slices.Clone(g)
						var srcG, srcC []float32
						if variant == 0 {
							srcG = g[1:]
							srcC = c[1:]
						} else if variant == 1 {
							srcG = g[150:]
							srcC = c[150:]
						}
						opuscc.CompareChannelFloat(&g[2], ds, dc, unsafe.SliceData(srcG), ss, n)
						nativeChannelOutput(unsafe.Pointer(&c[2]), ds, dc, srcC, ss, n)
						if !sameFloatBits(g, c) {
							t.Fatal(n, ds, ss, dc, variant)
						}
					}
				}
			}
		}
	}
}
