//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
)

func TestResamplerDriverAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(802))
	for _, enc := range []int32{0, 1} {
		for _, inRate := range []int32{8000, 12000, 16000, 24000, 48000} {
			for _, outRate := range []int32{8000, 12000, 16000, 24000, 48000} {
				if enc == 0 && inRate > 16000 || enc != 0 && outRate > 16000 {
					continue
				}
				var g opuscc.OpusT_silk_resampler_state_struct
				opuscc.Opus_silk_resampler_init(nil, &g, inRate, outRate, enc)
				c, _ := nativeResamplerInit(inRate, outRate, enc)
				for _, ms := range []int32{1, 2, 10, 11, 21, 30, 1} {
					in := make([]int16, ms*inRate/1000)
					for i := range in {
						in[i] = int16(r.Uint32())
					}
					original := slices.Clone(in)
					gout := make([]int16, ms*outRate/1000+2)
					for i := range gout {
						gout[i] = -12345
					}
					cout := slices.Clone(gout)
					gr := opuscc.Opus_silk_resampler(nil, &g, &gout[1], &in[0], int32(len(in)))
					cr := nativeResamplerDriver(&c, cout[1:], in)
					if gr != cr || g != c || !slices.Equal(gout, cout) || !slices.Equal(in, original) {
						t.Fatalf("%d -> %d enc=%d ms=%d status=%d/%d state=%v pcm=%v", inRate, outRate, enc, ms, gr, cr, g != c, !slices.Equal(gout, cout))
					}
				}
			}
		}
	}
}
