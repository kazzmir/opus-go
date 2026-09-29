//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
	"unsafe"
)

func TestIIRDriverAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(402))
	for _, batch := range []int{80, 120, 160, 480} {
		for _, step := range []int32{21846, 32768, 49152, 65536, 98304, 131072} {
			g := opuscc.OpusT_silk_resampler_state_struct{FbatchSize: int32(batch), FinvRatio_Q16: step}
			for i := range g.FsIIR {
				g.FsIIR[i] = int32(r.Intn(1<<25) - (1 << 24))
			}
			for i := range g.FsFIR.Fi32 {
				g.FsFIR.Fi32[i] = int32(r.Uint32())
			}
			c := g
			for _, n := range []int{0, 1, 7, batch - 1, batch, batch + 1, 3*batch + 7} {
				in := make([]int16, n)
				for i := range in {
					in[i] = int16(r.Uint32())
				}
				count := 0
				for left := n; left > 0; {
					take := min(left, batch)
					count += int(1 + (int32(take)<<17-1)/step)
					left -= take
				}
				goOut, cOut := make([]int16, count+2), make([]int16, count+2)
				goOut[0] = 123
				cOut[0] = 123
				goOut[count+1] = 456
				cOut[count+1] = 456
				opuscc.Opus_silk_resampler_private_IIR_FIR(nil, &g, &goOut[1], unsafe.SliceData(in), int32(n))
				nativeIIRDriver(&c, cOut[1:], in)
				if g != c || !slices.Equal(goOut, cOut) {
					t.Fatalf("batch=%d step=%d n=%d stateEqual=%v", batch, step, n, g == c)
				}
			}
		}
	}
}
