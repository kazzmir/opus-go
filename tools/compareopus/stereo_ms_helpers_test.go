//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
)

func TestStereoMSAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, fs := range []int32{8, 12, 16} {
		for _, ms := range []int32{8, 10, 20} {
			n := int(fs * ms)
			for trial := 0; trial < 100; trial++ {
				var gs opuscc.OpusT_stereo_dec_state
				for i := 0; i < 2; i++ {
					gs.Fpred_prev_Q13[i] = int16(rng.Uint32())
					gs.FsMid[i] = int16(rng.Uint32())
					gs.FsSide[i] = int16(rng.Uint32())
				}
				cs := gs
				for frame := 0; frame < 2; frame++ {
					gm, gy := make([]int16, n+2), make([]int16, n+2)
					for i := range gm {
						gm[i] = int16(rng.Uint32())
						gy[i] = int16(rng.Uint32())
					}
					cm, cy := slices.Clone(gm), slices.Clone(gy)
					pred := [2]int32{int32(int16(rng.Uint32())), int32(int16(rng.Uint32()))}
					before := pred
					opuscc.Opus_silk_stereo_MS_to_LR(nil, &gs, &gm[0], &gy[0], &pred, fs, int32(n))
					nativeStereoMS(&cs, cm, cy, &pred, fs)
					if !slices.Equal(gm, cm) || !slices.Equal(gy, cy) || gs != cs || pred != before {
						t.Fatalf("fs=%d ms=%d trial=%d frame=%d: PCM or state differs", fs, ms, trial, frame)
					}
				}
			}
		}
	}
}
