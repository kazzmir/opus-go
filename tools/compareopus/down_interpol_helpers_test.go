//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
)

var downFIRCases = []struct {
	order, fracs, step int32
	coefs              []int16
}{
	{18, 3, 87382, opuscc.Opus_silk_Resampler_3_4_COEFS[:]},
	{18, 2, 98304, opuscc.Opus_silk_Resampler_2_3_COEFS[:]},
	{24, 1, 131072, opuscc.Opus_silk_Resampler_1_2_COEFS[:]},
	{36, 1, 196608, opuscc.Opus_silk_Resampler_1_3_COEFS[:]},
	{36, 1, 262144, opuscc.Opus_silk_Resampler_1_4_COEFS[:]},
	{36, 1, 393216, opuscc.Opus_silk_Resampler_1_6_COEFS[:]},
}

func TestDownInterpolAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(403))
	for _, tc := range downFIRCases {
		for _, limit := range []int32{1, 65536, 37 << 16, 480 << 16} {
			for _, step := range []int32{1, tc.step} {
				if step == 1 && limit > 65536 {
					continue
				}
				count := 1 + (limit-1)/step
				length := ((count - 1) * step >> 16) + tc.order
				for trial := 0; trial < 12; trial++ {
					in := make([]int32, length)
					for i := range in {
						in[i] = int32(r.Uint32())
						if trial < 4 {
							in[i] >>= 8
						}
						if trial == 4 {
							in[i] = 2147483647
						}
						if trial == 5 {
							in[i] = -2147483648
						}
					}
					g, c := make([]int16, count+2), make([]int16, count+2)
					g[0] = 123
					c[0] = 123
					g[count+1] = 456
					c[count+1] = 456
					gn := opuscc.CompareDownFIRInterpol(&g[1], &in[0], &tc.coefs[2], tc.order, tc.fracs, limit, step)
					cn := nativeDownInterpol(c[1:], in, tc.coefs[2:], tc.order, tc.fracs, limit, step)
					if gn != cn || !slices.Equal(g, c) {
						t.Fatalf("order=%d fracs=%d limit=%d step=%d trial=%d", tc.order, tc.fracs, limit, step, trial)
					}
				}
			}
		}
	}
}
