//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"testing"
	"unsafe"
)

func TestDecoderSetFSAgainstC(t *testing.T) {
	for _, initial := range []int32{0, 8, 12, 16} {
		for _, subframes := range []int32{2, 4} {
			for _, api := range []int32{8000, 12000, 16000, 24000, 48000} {
				for _, rate := range []int32{8, 12, 16} {
					g := opuscc.OpusT_silk_decoder_state{Fnb_subfr: subframes, Fprev_gain_Q16: 123456}
					g.Fexc_Q14[0] = 98765
					if initial != 0 {
						opuscc.Opus_silk_decoder_set_fs(nil, &g, initial, api)
					}
					for i := range g.FoutBuf {
						g.FoutBuf[i] = int16(i + 1)
					}
					for i := range g.FsLPC_Q14_buf {
						g.FsLPC_Q14_buf[i] = int32(i + 1)
					}
					g.Fresampler_state.FsIIR[0] = 777
					g.Ffirst_frame_after_reset = 7
					g.FlagPrev = 300
					g.FLastGainIndex = -11
					g.FprevSignalType = 2
					c := g
					steps := []struct{ rate, api, sf int32 }{{rate, api, subframes}, {rate, api, subframes}, {rate, 48000, 6 - subframes}, {rate, 16000, subframes}, {16, 48000, 4}}
					for _, step := range steps {
						g.Fnb_subfr = step.sf
						c.Fnb_subfr = step.sf
						gr := opuscc.Opus_silk_decoder_set_fs(nil, &g, step.rate, step.api)
						cr, unchanged := nativeDecoderSetFS(&c, step.rate, step.api)
						if gr != cr || g != c || !unchanged {
							t.Fatal(initial, subframes, api, rate, step, gr, cr, "state differs", g != c, "C remainder", unchanged)
						}
					}
				}
			}
		}
	}
}

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
