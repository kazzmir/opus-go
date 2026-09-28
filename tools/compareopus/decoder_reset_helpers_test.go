//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"testing"
	"unsafe"
)

func TestDecoderResetAgainstC(t *testing.T) {
	for _, init := range []bool{false, true} {
		var g opuscc.OpusT_silk_decoder_state
		// All fields are numeric, fixed arrays, or legacy uintptr values.
		bytes := unsafe.Slice((*byte)(unsafe.Pointer(&g)), int(unsafe.Sizeof(g)))
		for i := range bytes {
			bytes[i] = 0xa5
		}
		var result int32
		if init {
			result = opuscc.Opus_silk_init_decoder(nil, &g)
		} else {
			result = opuscc.Opus_silk_reset_decoder(nil, &g)
		}
		c, cr, zero := nativeDecoderReset(init)
		if result != cr || g != c || !zero {
			t.Fatalf("init=%v Go status=%d C=%d state differs=%v C remainder zero=%v", init, result, cr, g != c, zero)
		}
	}
}
