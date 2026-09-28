package opuscc

import "testing"

func TestTellFracPointers(t *testing.T) {
	for exponent := uint(15); exponent <= 31; exponent++ {
		ctx := OpusT_ec_ctx{Frng: 1 << exponent, Fnbits_total: 100, Fval: 123}
		before := ctx
		if got, want := Opus_ec_tell_frac(nil, &ctx), uint32(100-int32(exponent)-1)*8; got != want {
			t.Fatalf("exponent=%d got=%d want=%d", exponent, got, want)
		}
		if ctx != before {
			t.Fatal("context changed")
		}
	}
	thresholds := []uint32{35733, 38967, 42495, 46340, 50535, 55109, 60097}
	for i, r := range thresholds {
		ctx := OpusT_ec_ctx{Frng: r << 8, Fnbits_total: 100}
		at := Opus_ec_tell_frac(nil, &ctx)
		ctx.Frng = (r + 1) << 8
		after := Opus_ec_tell_frac(nil, &ctx)
		if at != uint32(800-192-i) || after != at-1 {
			t.Fatalf("threshold=%d at=%d after=%d", r, at, after)
		}
	}
}
