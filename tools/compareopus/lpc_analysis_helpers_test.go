//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
)

func TestLPCAnalysisAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(3403))
	for _, d := range []int{6, 8, 10, 16, 24} {
		for _, extra := range []int{0, 1, 32, 320} {
			for trial := 0; trial < 100; trial++ {
				n := d + extra
				in := make([]int16, n)
				b := make([]int16, d)
				for i := range in {
					in[i] = int16(rng.Uint32())
					if trial == 0 {
						in[i] = -32768
					}
				}
				for i := range b {
					b[i] = int16(rng.Uint32())
					if trial == 0 {
						b[i] = -32768
					}
				}
				before, bc := slices.Clone(in), slices.Clone(b)
				g, c := make([]int16, n), make([]int16, n)
				opuscc.Opus_silk_LPC_analysis_filter(nil, &g[0], &in[0], &b[0], int32(n), int32(d), 0)
				nativeLPCAnalysis(c, in, b)
				if !slices.Equal(g, c) || !slices.Equal(in, before) || !slices.Equal(b, bc) {
					t.Fatalf("d=%d n=%d trial=%d Go=%v C=%v", d, n, trial, g, c)
				}
				// The C API has no restrict qualifier: preserve forward store order when overlapping.
				for _, shift := range []int{-1, 0, 1} {
					g = make([]int16, n+2)
					copy(g[1:], in)
					c = slices.Clone(g)
					opuscc.Opus_silk_LPC_analysis_filter(nil, &g[1+shift], &g[1], &b[0], int32(n), int32(d), 0)
					nativeLPCAnalysis(c[1+shift:1+shift+n], c[1:1+n], b)
					if !slices.Equal(g, c) {
						t.Fatalf("overlap d=%d n=%d trial=%d shift=%d", d, n, trial, shift)
					}
				}
			}
		}
	}
}
