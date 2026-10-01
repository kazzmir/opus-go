//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"testing"
	"unsafe"
)

func TestResamplerCoefficientLayoutAgainstC(t *testing.T) {
	var s opuscc.OpusT_silk_resampler_state_struct
	want := [2]uint64{uint64(unsafe.Sizeof(s)), uint64(unsafe.Offsetof(s.FCoefs))}
	if nativeResamplerLayout() != want {
		t.Fatal(nativeResamplerLayout(), want)
	}
}

func TestResamplerInitAgainstC(t *testing.T) {
	for _, enc := range []int32{0, 1} {
		for _, in := range []int32{8000, 12000, 16000, 24000, 48000} {
			for _, out := range []int32{8000, 12000, 16000, 24000, 48000} {
				if enc == 0 && in > 16000 || enc != 0 && out > 16000 {
					continue
				}
				var g opuscc.OpusT_silk_resampler_state_struct
				coef := int16(123)
				g.FCoefs = &coef
				raw := unsafe.Slice((*byte)(unsafe.Pointer(&g)), int(unsafe.Offsetof(g.FCoefs)))
				for i := range raw {
					raw[i] = 0xa5
				}
				gr := opuscc.Opus_silk_resampler_init(nil, &g, in, out, enc)
				c, cr := nativeResamplerInit(in, out, enc)
				if gr != cr || g != c {
					t.Fatalf("%d -> %d enc=%d status=%d/%d\nGo=%+v\nC=%+v", in, out, enc, gr, cr, g, c)
				}
			}
		}
	}
}
