//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"runtime"
	"testing"
)

func TestStereoPredAgainstC(t *testing.T) {
	for combo := 0; combo < 25*3*5*3*5; combo++ {
		v := combo
		var ix [5]int32
		for i, radix := range []int{25, 3, 5, 3, 5} {
			ix[i] = int32(v % radix)
			v /= radix
		}
		mid := int32(combo % 2)
		data := make([]byte, 32)
		if err := nativeStereoPredEncode(data, &ix, mid); err != 0 {
			t.Fatalf("C encoder error=%d", err)
		}
		var g opuscc.OpusT_ec_dec
		opuscc.Opus_ec_dec_init(nil, &g, &data[0], uint32(len(data)))
		c := g
		var pred [2]int32
		opuscc.Opus_silk_stereo_decode_pred(nil, &g, &pred)
		_, cp := nativeEntropyStepOutput(&c, data, 7, 0, 0, 0)
		if pred != cp || g != c {
			t.Fatalf("ix=%v Go=%v C=%v state=%+v C=%+v", ix, pred, cp, g, c)
		}
		var gm int32
		opuscc.Opus_silk_stereo_decode_mid_only(nil, &g, &gm)
		cm := int32(nativeEntropyStep(&c, data, 8, 0, 0, 0))
		if gm != mid || gm != cm || g != c {
			t.Fatalf("ix=%v mid Go=%d C=%d want=%d", ix, gm, cm, mid)
		}
		runtime.KeepAlive(data)
	}
	// Exhausted-packet behavior still normalizes and produces valid symbols.
	for _, value := range []uint32{0, (1 << 31) - 1} {
		g := opuscc.OpusT_ec_dec{Frng: 1 << 31, Fval: value, Fnbits_total: 33}
		c := g
		var pred [2]int32
		opuscc.Opus_silk_stereo_decode_pred(nil, &g, &pred)
		_, cp := nativeEntropyStepOutput(&c, nil, 7, 0, 0, 0)
		if pred != cp || g != c {
			t.Fatal("exhausted-packet predictor differs")
		}
	}
}
