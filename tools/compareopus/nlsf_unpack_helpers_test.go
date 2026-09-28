//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
)

func TestNLSFUnpackAgainstC(t *testing.T) {
	for _, cb := range []*opuscc.OpusT_silk_NLSF_CB_struct{&opuscc.Opus_silk_NLSF_CB_NB_MB, &opuscc.Opus_silk_NLSF_CB_WB} {
		n := int(cb.Forder)
		for index := int32(0); index < int32(cb.FnVectors); index++ {
			g, c := make([]int16, n), make([]int16, n)
			gp, cp := make([]uint8, n), make([]uint8, n)
			opuscc.Opus_silk_NLSF_unpack(nil, &g[0], &gp[0], cb, index)
			nativeNLSFUnpack(c, cp, n == 16, index)
			if !slices.Equal(g, c) || !slices.Equal(gp, cp) {
				t.Fatalf("order=%d index=%d Go=%v/%v C=%v/%v", n, index, g, gp, c, cp)
			}
		}
	}
}
