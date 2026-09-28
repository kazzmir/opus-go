//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"testing"
)

func TestTellFracAgainstC(t *testing.T) {
	// Exhaust every normalized mantissa, including all correction thresholds.
	for _, shift := range []uint{0, 8, 15, 16} {
		for r := uint32(32768); r < 65536; r++ {
			for _, nbits := range []int32{33, 10000} {
				ctx := opuscc.OpusT_ec_ctx{Frng: r << shift, Fnbits_total: nbits}
				g, c := opuscc.Opus_ec_tell_frac(nil, &ctx), nativeTellFrac(ctx.Frng, nbits)
				if g != c {
					t.Fatalf("range=%d nbits=%d Go=%d C=%d", ctx.Frng, nbits, g, c)
				}
			}
		}
	}
}
