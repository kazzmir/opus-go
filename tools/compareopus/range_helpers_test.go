//go:build compareopus && cgo

package main

import (
	"math/rand"
	"testing"

	"github.com/kazzmir/opus-go/opuscc"
)

func TestRangeLookupsAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(2026))
	for trial := 0; trial < 10000; trial++ {
		rng := uint32(1<<24) + r.Uint32()%(1<<30)
		val := r.Uint32() % rng
		if trial%2 == 0 {
			val = rng - 1
		} // Includes clamping of division remainder.
		for _, binary := range []bool{false, true} {
			param := uint32(1) + r.Uint32()%65536
			if binary {
				param = r.Uint32() % 16
			}
			dec := opuscc.OpusT_ec_dec{Frng: rng, Fval: val, Foffs: 13}
			want := dec
			symbol, ext := nativeRangeLookup(rng, val, param, binary)
			want.Fext = ext
			var got uint32
			if binary {
				got = opuscc.Opus_ec_decode_bin(nil, &dec, param)
			} else {
				got = opuscc.Opus_ec_decode(nil, &dec, param)
			}
			if got != symbol || dec != want {
				t.Fatalf("trial=%d binary=%t param=%d: symbol Go=%d C=%d state=%+v want=%+v", trial, binary, param, got, symbol, dec, want)
			}
		}
	}
}
