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

func TestDecodeParametersAgainstC(t *testing.T) {
	for _, fs := range []int32{8, 12, 16} {
		for _, sub := range []int32{2, 4} {
			for _, signal := range []int8{0, 1, 2} {
				for _, cond := range []int32{0, 1, 2} {
					for trial := 0; trial < 12; trial++ {
						var g opuscc.OpusT_silk_decoder_state
						g.Fnb_subfr = sub
						opuscc.Opus_silk_decoder_set_fs(nil, &g, fs, 16000)
						g.Ffirst_frame_after_reset = int32(trial % 2)
						g.FlossCnt = int32(trial % 3)
						g.Findices.FsignalType = signal
						g.Findices.FNLSFInterpCoef_Q2 = int8(trial % 5)
						g.Findices.FGainsIndices = [4]int8{9, 3, 5, 7}
						g.FLastGainIndex = 12
						g.Findices.FlagIndex = 100
						g.Findices.FPERIndex = int8(trial % 3)
						g.Findices.FLTPIndex = [4]int8{2, 5, 1, 3}
						g.Findices.FLTP_scaleIndex = int8(trial % 3)
						for i := int32(0); i < g.FLPC_order; i++ {
							g.FprevNLSF_Q15[i] = int16((i + 1) * 32768 / (g.FLPC_order + 1))
						}
						c := g
						var gc opuscc.OpusT_silk_decoder_control
						for i := range gc.FLTPCoef_Q14 {
							gc.FLTPCoef_Q14[i] = 77
						}
						cc := gc
						opuscc.Opus_silk_decode_parameters(nil, &g, &gc, cond)
						nativeSilkParameters(&c, &cc, cond)
						g.FpsNLSF_CB = 0
						if g != c || gc != cc {
							t.Fatal(fs, sub, signal, cond, trial, "state/control")
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
