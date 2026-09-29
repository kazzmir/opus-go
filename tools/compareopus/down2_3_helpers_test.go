//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
	"unsafe"
)

func TestDownTwoThirdsAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(803))
	for _, alias := range []bool{false, true} {
		for trial := 0; trial < 8; trial++ {
			var gs [6]int32
			for i := range gs {
				gs[i] = int32(r.Uint32())
			}
			cs := gs
			for _, n := range []int{0, 1, 2, 3, 4, 5, 6, 479, 480, 481, 482, 483, 959, 960, 961, 962, 1441} {
				gi := make([]int16, n+2)
				for i := range gi {
					gi[i] = int16(r.Uint32())
				}
				ci := slices.Clone(gi)
				g := make([]int16, 2*(n/3)+2)
				for i := range g {
					g[i] = 12345
				}
				c := slices.Clone(g)
				if alias {
					g = gi
					c = ci
				}
				opuscc.Opus_silk_resampler_down2_3(nil, &gs, &g[1], unsafe.SliceData(gi[1:1+n]), int32(n))
				nativeDownTwoThirds(&cs, c[1:], ci[1:1+n])
				if gs != cs || !slices.Equal(g, c) || !slices.Equal(gi, ci) {
					t.Fatalf("alias=%v trial=%d N=%d state=%v PCM=%v", alias, trial, n, gs != cs, !slices.Equal(g, c))
				}
			}
		}
	}
}
