//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
)

func TestIIRInterpolAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(401))
	for _, limit := range []int32{1, 65536, 37 << 16, 480 << 16} {
		for _, step := range []int32{1, 21846, 32768, 65536, 98304, 131072, 196608} {
			if step == 1 && limit > 65536 {
				continue
			}
			count := 1 + (limit-1)/step
			length := ((count - 1) * step >> 16) + 8
			for trial := 0; trial < 8; trial++ {
				in := make([]int16, length)
				for i := range in {
					in[i] = int16(r.Uint32())
					if trial == 0 {
						in[i] = -32768
					}
					if trial == 1 {
						in[i] = 32767
					}
				}
				g, c := make([]int16, count+2), make([]int16, count+2)
				g[0] = 123
				c[0] = 123
				g[count+1] = 456
				c[count+1] = 456
				gn := opuscc.CompareIIRFIRInterpol(&g[1], &in[0], limit, step)
				cn := nativeIIRInterpol(c[1:], in, limit, step)
				if gn != cn || !slices.Equal(g, c) {
					t.Fatalf("limit=%d step=%d trial=%d", limit, step, trial)
				}
			}
		}
	}
	// Output/input overlap: preserve the C read-before-store order.
	for _, step := range []int32{65536, 98304, 131072} {
		g := make([]int16, 48)
		for i := range g {
			g[i] = int16(r.Uint32())
		}
		c := slices.Clone(g)
		gn := opuscc.CompareIIRFIRInterpol(&g[1], &g[0], 32<<16, step)
		cn := nativeIIRInterpol(c[1:], c, 32<<16, step)
		if gn != cn || !slices.Equal(g, c) {
			t.Fatal("overlap", step)
		}
	}
}
