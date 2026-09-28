//go:build compareopus && cgo

package main

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/kazzmir/opus-go/opuscc"
	"github.com/kazzmir/opus-go/opusccenc"
)

func TestBiquad1AgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, n := range []int{2, 16, 320} {
		for _, coefs := range []struct {
			b [3]int32
			a [2]int32
		}{
			{[3]int32{1 << 28}, [2]int32{}},
			{[3]int32{1 << 29}, [2]int32{}},
			{[3]int32{1 << 27, 1 << 26, -(1 << 25)}, [2]int32{-(1 << 26) + 123, (1 << 24) + 456}},
		} {
			for trial := 0; trial < 100; trial++ {
				input := make([]int16, n)
				for i := range input {
					input[i] = int16(rng.Uint32())
				}
				g, c := slices.Clone(input), slices.Clone(input)
				gs := [2]int32{rng.Int31n(1<<21) - (1 << 20), rng.Int31n(1<<21) - (1 << 20)}
				cs := gs
				e, es := slices.Clone(input), gs
				opusccenc.Opus_silk_biquad_alt_stride1(nil, &e[0], &coefs.b, &coefs.a, &es, &e[0], int32(n))
				opuscc.Opus_silk_biquad_alt_stride1(nil, &g[0], &coefs.b, &coefs.a, &gs, &g[0], int32(n))
				nativeBiquad1(c, c, &coefs.b, &coefs.a, &cs)
				if !slices.Equal(e, c) || es != cs {
					t.Fatalf("encoder differs: n=%d trial=%d state=%v C=%v", n, trial, es, cs)
				}
				if !slices.Equal(g, c) || gs != cs {
					t.Fatalf("n=%d trial=%d: Go state=%v C state=%v Go PCM=%v C PCM=%v", n, trial, gs, cs, g, c)
				}
			}
		}
	}
}
