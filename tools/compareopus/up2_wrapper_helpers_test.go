//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
	"unsafe"
)

func TestUp2WrapperAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(3502))
	for trial := 0; trial < 100; trial++ {
		var g opuscc.OpusT_silk_resampler_state_struct
		for i := range g.FsIIR {
			g.FsIIR[i] = int32(rng.Uint32())
		}
		g.FCoefs = 123
		g.FdelayBuf[95] = -456
		g.FsFIR.Fi32[35] = 789
		before := g
		c := g.FsIIR
		for _, n := range []int{0, 1, 7, 80, 160, 480} {
			in := make([]int16, n)
			for i := range in {
				in[i] = int16(rng.Uint32())
			}
			original := slices.Clone(in)
			goOut, cOut := make([]int16, 2*n), make([]int16, 2*n)
			opuscc.Opus_silk_resampler_private_up2_HQ_wrapper(nil, &g, unsafe.SliceData(goOut), unsafe.SliceData(in), int32(n))
			if !nativeUp2Wrapper(&c, cOut, in) {
				t.Fatal("C changed non-IIR state")
			}
			before.FsIIR = c
			if !slices.Equal(goOut, cOut) || g != before || !slices.Equal(in, original) {
				t.Fatalf("trial=%d n=%d output/state differs", trial, n)
			}
		}
	}
}
