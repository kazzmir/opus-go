//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
	"unsafe"
)

func TestPLCGlueAgainstC(t *testing.T) {
	for _, n := range []int{1, 16, 120, 320} {
		for _, amp := range []int16{0, 1, 100, 10000, 32767, -32768} {
			for _, energy := range []int32{0, 1, 16000, 1 << 24, 1 << 29} {
				for _, shift := range []int32{0, 2, 8, 16} {
					for _, loss := range []int32{0, 1} {
						g := make([]int16, n+2)
						g[0] = 77
						g[n+1] = 88
						for i := 1; i <= n; i++ {
							g[i] = amp
						}
						c := slices.Clone(g)
						gd := opuscc.OpusT_silk_decoder_state{FlossCnt: loss}
						gd.FsPLC.Fconc_energy = energy
						gd.FsPLC.Fconc_energy_shift = shift
						gd.FsPLC.Flast_frame_lost = 1
						cd := gd
						opuscc.Opus_silk_PLC_glue_frames(nil, &gd, unsafe.SliceData(g[1:]), int32(n))
						nativePLCGlue(&cd, c[1:n+1])
						if gd != cd || !slices.Equal(g, c) {
							t.Fatal(n, amp, energy, shift, loss, gd.FsPLC, cd.FsPLC, g, c)
						}
					}
				}
			}
		}
	}
}
