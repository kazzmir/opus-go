//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
)

func TestNLSFDecodeAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, cb := range []*opuscc.OpusT_silk_NLSF_CB_struct{&opuscc.Opus_silk_NLSF_CB_NB_MB, &opuscc.Opus_silk_NLSF_CB_WB} {
		n := int(cb.Forder)
		for index := 0; index < int(cb.FnVectors); index++ {
			for trial := 0; trial < 100; trial++ {
				indices := make([]int8, n+1)
				indices[0] = int8(index)
				for i := 1; i <= n; i++ {
					indices[i] = int8(rng.Intn(21) - 10)
					if trial == 0 {
						indices[i] = 0
					}
					if trial == 1 {
						indices[i] = -10
					}
					if trial == 2 {
						indices[i] = 10
					}
				}
				before := slices.Clone(indices)
				g, c := make([]int16, n), make([]int16, n)
				opuscc.Opus_silk_NLSF_decode(nil, &g[0], &indices[0], cb)
				nativeNLSFDecode(c, indices, n == 16)
				if !slices.Equal(g, c) || !slices.Equal(indices, before) {
					t.Fatalf("order=%d index=%d trial=%d Go=%v C=%v", n, index, trial, g, c)
				}
			}
		}
	}
}
