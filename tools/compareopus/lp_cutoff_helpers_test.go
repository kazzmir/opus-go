//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"github.com/kazzmir/opus-go/opusccenc"
	"math/rand"
	"slices"
	"testing"
	"unsafe"
)

func TestLPCutoffAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for pos := int32(0); pos <= 256; pos++ {
		for _, mode := range []int32{-2, -1, 0, 1, 2} {
			gs := opuscc.OpusT_silk_LP_state{FIn_LP_State: [2]int32{rng.Int31n(1<<24) - (1 << 23), rng.Int31n(1<<24) - (1 << 23)}, Ftransition_frame_no: pos, Fmode: mode, Fsaved_fs_kHz: 16}
			cs := gs
			es := opusccenc.OpusT_silk_LP_state(gs)
			// Consecutive frames test state continuity, endpoints, and both directions.
			for _, n := range []int{1, 20, 120, 320} {
				g := make([]int16, n)
				for i := range g {
					g[i] = int16(rng.Uint32())
				}
				c := slices.Clone(g)
				e := slices.Clone(g)
				opusccenc.Opus_silk_LP_variable_cutoff(nil, &es, unsafe.SliceData(e), int32(n))
				opuscc.Opus_silk_LP_variable_cutoff(nil, &gs, unsafe.SliceData(g), int32(n))
				nativeLPCutoff(&cs, c)
				if !slices.Equal(e, c) || es != opusccenc.OpusT_silk_LP_state(cs) {
					t.Fatalf("encoder differs: position=%d mode=%d n=%d state=%+v C=%+v", pos, mode, n, es, cs)
				}
				if !slices.Equal(g, c) || gs != cs {
					t.Fatalf("position=%d mode=%d n=%d: Go state=%+v C state=%+v PCM matches=%v", pos, mode, n, gs, cs, slices.Equal(g, c))
				}
			}
		}
	}
}
