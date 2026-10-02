//go:build compareopus && cgo

package main

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"runtime"
	"slices"
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

func TestDecodeIndicesAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(8917))
	for _, fs := range []int32{8, 12, 16} {
		for _, sub := range []int32{2, 4} {
			for _, cond := range []int32{0, 1, 2} {
				for trial := 0; trial < 80; trial++ {
					var g opuscc.OpusT_silk_decoder_state
					g.Fnb_subfr = sub
					opuscc.Opus_silk_decoder_set_fs(nil, &g, fs, 16000)
					frame := int32(trial % 3)
					lbrr := int32((trial / 3) % 2)
					g.FVAD_flags[frame] = int32((trial / 6) % 2)
					g.Fec_prevSignalType = int32(trial % 3)
					g.Fec_prevLagIndex = 100
					g.Findices.FLTP_scaleIndex = 2
					g.Findices.FPERIndex = 1
					c := g
					data := make([]byte, []int{0, 1, 32}[trial%3])
					rng.Read(data)
					var gd opuscc.OpusT_ec_dec
					gd.Fext = uint32(trial)
					opuscc.Opus_ec_dec_init(nil, &gd, unsafe.SliceData(data), uint32(len(data)))
					cd := gd
					opuscc.Opus_silk_decode_indices(nil, &g, &gd, frame, lbrr, cond)
					nativeSilkIndices(&c, &cd, data, frame, lbrr, cond)
					g.FpsNLSF_CB = nil
					gd.Fbuf = nil
					cd.Fbuf = nil
					if g != c || gd != cd {
						t.Fatal(fs, sub, cond, trial, "state/entropy", g.Findices, c.Findices, gd, cd)
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
						g.FpsNLSF_CB = nil
						if g != c || gc != cc {
							t.Fatal(fs, sub, signal, cond, trial, "state/control")
						}
					}
				}
			}
		}
	}
}

func TestSilkCodebookReferenceLayoutAgainstC(t *testing.T) {
	var d opuscc.OpusT_silk_decoder_state
	var e opuscc.OpusT_silk_encoder_state
	want := [4]uint64{uint64(unsafe.Sizeof(d)), uint64(unsafe.Offsetof(d.FpsNLSF_CB)), uint64(unsafe.Sizeof(e)), uint64(unsafe.Offsetof(e.FpsNLSF_CB))}
	if nativeSilkCBReferenceLayout() != want {
		t.Fatal(nativeSilkCBReferenceLayout(), want)
	}
}

func TestPitchLagReferenceLayoutAgainstC(t *testing.T) {
	var d opuscc.OpusT_silk_decoder_state
	var e opuscc.OpusT_silk_encoder_state
	want := [4]uint64{uint64(unsafe.Sizeof(d)), uint64(unsafe.Offsetof(d.Fpitch_lag_low_bits_iCDF)), uint64(unsafe.Sizeof(e)), uint64(unsafe.Offsetof(e.Fpitch_lag_low_bits_iCDF))}
	if nativeSilkPitchReferenceLayout(false) != want {
		t.Fatal(nativeSilkPitchReferenceLayout(false), want)
	}
}

func TestPitchContourReferenceLayoutAgainstC(t *testing.T) {
	var d opuscc.OpusT_silk_decoder_state
	var e opuscc.OpusT_silk_encoder_state
	want := [4]uint64{uint64(unsafe.Sizeof(d)), uint64(unsafe.Offsetof(d.Fpitch_contour_iCDF)), uint64(unsafe.Sizeof(e)), uint64(unsafe.Offsetof(e.Fpitch_contour_iCDF))}
	if nativeSilkPitchReferenceLayout(true) != want {
		t.Fatal(nativeSilkPitchReferenceLayout(true), want)
	}
}

func TestWholeCNGAgainstC(t *testing.T) {
	for _, order := range []int32{10, 16} {
		for _, sub := range []int32{2, 4} {
			for _, length := range []int{0, 1, 8, 80, 160} {
				for _, loss := range []int32{0, 1, 3} {
					for _, gain := range []int32{0, 1100000, 11950000} {
						tls := libc.NewTLS()
						ps := libc.XmallocPointer(tls, uint64(unsafe.Sizeof(opuscc.OpusT_opus_ccgo_pseudostack_state{})))
						scratch := libc.XmallocPointer(tls, opuscc.GLOBAL_STACK_SIZE)
						*(*opuscc.OpusT_opus_ccgo_pseudostack_state)(ps) = opuscc.OpusT_opus_ccgo_pseudostack_state{Fscratch_ptr: uintptr(scratch), Fglobal_stack: uintptr(scratch)}
						libc.Xpthread_setspecific(tls, 0x6f707573, uintptr(ps))
						g := opuscc.OpusT_silk_decoder_state{Ffs_kHz: 16, FLPC_order: order, Fnb_subfr: sub, Fsubfr_length: 40, FlossCnt: loss, FprevSignalType: 0}
						g.FsCNG.Ffs_kHz = 16
						g.FsCNG.FCNG_smth_Gain_Q16 = gain
						g.FsCNG.Frand_seed = 24681357
						g.FsPLC.FrandScale_Q14 = 12000
						g.FsPLC.FprevGain_Q16 = [2]int32{650000, 900000}
						for i := int32(0); i < order; i++ {
							g.FprevNLSF_Q15[i] = int16((i + 1) * 32767 / (order + 1))
							g.FsCNG.FCNG_smth_NLSF_Q15[i] = g.FprevNLSF_Q15[i]
						}
						for i := range g.Fexc_Q14 {
							g.Fexc_Q14[i] = int32(i*101 - 17000)
							g.FsCNG.FCNG_exc_buf_Q14[i] = int32(i*37 - 7000)
						}
						for i := range g.FsCNG.FCNG_synth_state {
							g.FsCNG.FCNG_synth_state[i] = int32(i*19 - 149)
						}
						control := opuscc.OpusT_silk_decoder_control{FGains_Q16: [4]int32{800000, 500000, 1100000, 400000}}
						cc := control
						c := g
						frame, cf := make([]int16, length+2), make([]int16, length+2)
						for i := range frame {
							frame[i] = int16(i*131 - 2000)
							cf[i] = frame[i]
						}
						var pins runtime.Pinner
						pins.Pin(&g)
						pins.Pin(unsafe.SliceData(frame))
						opuscc.Opus_silk_CNG(tls, &g, &control, uintptr(unsafe.Pointer(&frame[1])), int32(length))
						nativeWholeCNG(&c, &cc, cf[1:len(cf)-1])
						pins.Unpin()
						tls.Close()
						if g != c || control != cc || !slices.Equal(frame, cf) {
							t.Fatal(order, sub, length, loss, gain, "CNG state/frame", g.FsCNG, c.FsCNG, frame, cf)
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
		// Poison numeric storage only; keep GC pointer slots valid.
		bytes := unsafe.Slice((*byte)(unsafe.Pointer(&g)), int(unsafe.Sizeof(g)))
		ptrSize := int(unsafe.Sizeof(g.FpsNLSF_CB))
		coefs := int(unsafe.Offsetof(g.Fresampler_state)) + int(unsafe.Offsetof(g.Fresampler_state.FCoefs))
		cb := int(unsafe.Offsetof(g.FpsNLSF_CB))
		lag := int(unsafe.Offsetof(g.Fpitch_lag_low_bits_iCDF))
		contour := int(unsafe.Offsetof(g.Fpitch_contour_iCDF))
		for i := range bytes {
			if (i >= coefs && i < coefs+ptrSize) || (i >= cb && i < cb+ptrSize) || (i >= lag && i < lag+ptrSize) || (i >= contour && i < contour+ptrSize) {
				continue
			}
			bytes[i] = 0xa5
		}
		g.FpsNLSF_CB = &opuscc.Opus_silk_NLSF_CB_WB
		g.Fpitch_lag_low_bits_iCDF = &opuscc.Opus_silk_uniform8_iCDF[0]
		g.Fpitch_contour_iCDF = &opuscc.Opus_silk_pitch_contour_iCDF[0]
		g.Fresampler_state.FCoefs = &opuscc.Opus_silk_Resampler_1_2_COEFS[0]
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
