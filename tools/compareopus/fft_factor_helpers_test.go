//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"testing"
)

func TestFFTFactorAgainstC(t *testing.T) {
	inputs := []int32{65535, 65536, 65537, 1 << 30, 2147483647}
	for n := int32(1); n <= 4096; n++ {
		inputs = append(inputs, n)
	}
	for _, n := range inputs {
		var g, c [64]int32
		for i := range g {
			g[i] = -9
			c[i] = -9
		}
		gu := opuscc.CompareFFTFactor(n, &g)
		cu := nativeFFTFactor(n, &c)
		if gu != cu || g != c {
			t.Fatal(n, gu, cu, g, c)
		}
	}
}
