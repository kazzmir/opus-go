//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"runtime"
	"testing"
)

func TestLaplaceP0AgainstC(t *testing.T) {
	values := []int32{0, 1, -1, 6, -6, 7, -7, 8, -8, 14, -14, 15, -15, 49, -49, 100, -100}
	for _, p0 := range []uint16{1, 8000, 16000, 32766} {
		for _, decay := range []uint16{0, 1, 7, 8000, 16000, 30000} {
			data := make([]byte, 4096)
			if err := nativeLaplaceP0Encode(data, values, p0, decay); err != 0 {
				t.Fatalf("C encoder error=%d", err)
			}
			var g opuscc.OpusT_ec_dec
			opuscc.Opus_ec_dec_init(nil, &g, &data[0], uint32(len(data)))
			c := g
			for i, want := range values {
				gv := opuscc.Opus_ec_laplace_decode_p0(nil, &g, p0, decay)
				cv := int32(nativeEntropyStep(&c, data, 6, uint32(p0), uint32(decay), 0))
				if gv != want || gv != cv || g != c {
					t.Fatalf("p0=%d decay=%d i=%d Go=%d C=%d want=%d state=%+v C=%+v", p0, decay, i, gv, cv, want, g, c)
				}
			}
			runtime.KeepAlive(data)
		}
	}
}
