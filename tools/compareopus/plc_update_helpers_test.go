//go:build compareopus && cgo

package main

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
	"unsafe"
)

func TestPLCLTPHistoryAgainstC(t *testing.T) {
	h := []int32{-2147483648, 2147483647, -1, 0, 100000003, 12345, -54321, 2147483647}
	for _, b := range [][5]int16{{-32768, 32767, -1, 16384, 12345}, {300, -150, 1200, -100, 75}, {0, 0, 0, 0, 0}} {
		for index := int32(4); index < int32(len(h)); index++ {
			g := opuscc.ComparePLCLTPPrediction(h, index, &b)
			c := nativePLCLTPPrediction(h, index, &b)
			if g != c {
				t.Fatal("LTP", index, b, g, c)
			}
		}
	}
}

func TestPLCWhiteningAgainstC(t *testing.T) {
	for _, length := range []int32{160, 240, 320} {
		for _, order := range []int32{10, 16} {
			for _, index := range []int32{1, 17, length - order - 1} {
				d := opuscc.OpusT_silk_decoder_state{Fltp_mem_length: length, FLPC_order: order}
				for i := range d.FoutBuf {
					d.FoutBuf[i] = int16((i*71)%60000 - 30000)
				}
				var a [16]int16
				for i := range a {
					a[i] = int16(i*137 - 700)
				}
				samples := make([]int16, length+2)
				for i := range samples {
					samples[i] = 123
				}
				samples[0], samples[len(samples)-1] = 77, 88
				want := slices.Clone(samples)
				before := d
				opuscc.ComparePLCWhiten(&d, samples[1:len(samples)-1], &a, index)
				nativeLPCAnalysis(want[1+index:len(want)-1], d.FoutBuf[index:length], a[:order])
				if !slices.Equal(samples, want) || d != before {
					t.Fatal("whitening", length, order, index)
				}
			}
		}
	}
}

func TestPLCLPCAgainstC(t *testing.T) {
	for _, length := range []int32{0, 1, 10, 16, 17, 80, 320} {
		for _, order := range []int32{10, 16} {
			for _, gain := range []int32{0, 1024, 65536, 2147483647} {
				for _, extreme := range []bool{false, true} {
					d := opuscc.OpusT_silk_decoder_state{Fframe_length: length, FLPC_order: order}
					for i := range d.FsLPC_Q14_buf {
						d.FsLPC_Q14_buf[i] = int32(i*1000003) - 8000000
					}
					state := d.FsLPC_Q14_buf
					var a [16]int16
					for i := range a {
						a[i] = int16(i*137 - 700)
						if extreme {
							a[i] = int16(i*12345 - 32768)
						}
					}
					history := make([]int32, 18+length)
					history[0], history[len(history)-1] = 77, 88
					for i := int32(0); i < length; i++ {
						history[17+i] = int32(int64(i) * 1000000003)
					}
					ch := slices.Clone(history)
					pcm := make([]int16, length+2)
					pcm[0], pcm[len(pcm)-1] = 77, 88
					cp := slices.Clone(pcm)
					opuscc.ComparePLCLPC(&d, history[1:len(history)-1], &a, unsafe.SliceData(pcm[1:len(pcm)-1]), gain)
					nativePLCLPC(&state, ch[1:len(ch)-1], &a, cp[1:len(cp)-1], order, gain)
					if d.FsLPC_Q14_buf != state || !slices.Equal(history, ch) || !slices.Equal(pcm, cp) {
						t.Fatal("LPC history", length, order, gain, extreme, d.FsLPC_Q14_buf, state, pcm, cp)
					}
				}
			}
		}
	}
}

func TestPLCPCMAgainstC(t *testing.T) {
	for _, sample := range []int32{-2147483648, -8388609, -128, -1, 0, 127, 128, 8388608, 2147483647} {
		for _, gain := range []int32{-2147483648, -65536, -1, 0, 1, 1024, 65536, 2147483647} {
			g := opuscc.ComparePLCPCM(sample, gain)
			c := nativePLCPCM(sample, gain)
			if g != c {
				t.Fatal("PCM", sample, gain, g, c)
			}
		}
	}
}

func TestPLCNoiseAgainstC(t *testing.T) {
	random := make([]int32, 128)
	for i := range random {
		random[i] = int32(i*1000003) - 160000000
	}
	for i := int32(0); i < 128; i++ {
		for _, scale := range []int16{-32768, -1, 0, 16384, 32767} {
			for _, p := range []int32{-1000000, 0, 1000000} {
				g := opuscc.ComparePLCNoise(p, random, i, scale)
				c := nativePLCNoise(p, random, i, scale)
				if g != c {
					t.Fatal("noise", i, scale, p, g, c)
				}
			}
		}
	}
}

func TestPLCDecayAgainstC(t *testing.T) {
	for _, gain := range []int32{-32768, -32767, 0, 16384, 32767, 32768, 65535} {
		g := [5]int16{-32768, 32767, -1, 0, 12345}
		c := g
		opuscc.ComparePLCDecay(&g, gain)
		nativePLCDecay(&c, gain)
		if g != c {
			t.Fatal("decay", gain, g, c)
		}
	}
}

func TestPLCDispatchLossAgainstC(t *testing.T) {
	for _, rate := range []int32{8, 12, 16} {
		for _, nb := range []int32{2, 4} {
			for _, signal := range []int32{1, 2} {
				for _, lossCnt := range []int32{0, 1, 3} {
					for _, lost := range []int32{1, -1, 7} {
						for _, reset := range []bool{false, true} {
							order := int32(10)
							if rate == 16 {
								order = 16
							}
							g := opuscc.OpusT_silk_decoder_state{Ffs_kHz: rate, Fframe_length: rate * 5 * nb, Fsubfr_length: rate * 5, Fnb_subfr: nb, Fltp_mem_length: rate * 20, FLPC_order: order, FprevSignalType: signal, FlossCnt: lossCnt}
							g.FsPLC = opuscc.OpusT_silk_PLC_struct{Ffs_kHz: rate, FpitchL_Q8: rate * 5 << 8, FLTPCoef_Q14: [5]int16{300, -150, 1200, -100, 75}, FprevLPC_Q12: [16]int16{120, -80, 60, -45, 30, -20, 15, -10, 8, -5}, FprevGain_Q16: [2]int32{65536, 65536}, FprevLTP_scale_Q14: 13000, FrandScale_Q14: 11000, Frand_seed: 12345, Fsubfr_length: rate * 5, Fnb_subfr: nb}
							if reset {
								g.FsPLC.Ffs_kHz = 0
								g.Ffirst_frame_after_reset = 1
							}
							for i := range g.Fexc_Q14 {
								g.Fexc_Q14[i] = int32((i*71)%3000000 - 1500000)
							}
							for i := range g.FoutBuf {
								g.FoutBuf[i] = int16((i*37)%1000 - 500)
							}
							control := opuscc.OpusT_silk_decoder_control{FGains_Q16: [4]int32{77, 88, 99, 111}, FpitchL: [4]int32{7, 8, 9, 10}}
							c, cc := g, control
							frame := make([]int16, int(g.Fframe_length)+2)
							for i := range frame {
								frame[i] = int16(i*13 + 77)
							}
							cf := slices.Clone(frame)
							tls := libc.NewTLS()
							ps := libc.XmallocPointer(tls, uint64(unsafe.Sizeof(opuscc.OpusT_opus_ccgo_pseudostack_state{})))
							scratch := libc.XmallocPointer(tls, opuscc.GLOBAL_STACK_SIZE)
							*(*opuscc.OpusT_opus_ccgo_pseudostack_state)(ps) = opuscc.OpusT_opus_ccgo_pseudostack_state{Fscratch_ptr: uintptr(scratch), Fglobal_stack: uintptr(scratch)}
							libc.Xpthread_setspecific(tls, 0x6f707573, uintptr(ps))
							opuscc.ComparePLCDispatch(tls, &g, &control, &frame[1], lost, 0)
							r := nativePLCDispatch(&c, &cc, cf[1:len(cf)-1], lost, 0)
							tls.Close()
							if r != 0 || g != c || control != cc || !slices.Equal(frame, cf) {
								t.Fatal("PLC lost dispatch", rate, nb, signal, lossCnt, lost, reset, r, g.FsPLC, c.FsPLC, control.FpitchL, cc.FpitchL, frame[:9], cf[:9])
							}
							if g.FlossCnt != lossCnt+1 {
								t.Fatal("increment after conceal", g.FlossCnt)
							}
						}
					}
				}
			}
		}
	}
}

func TestPLCDispatchUpdateAgainstC(t *testing.T) {
	for _, rate := range []int32{8, 12, 16} {
		for _, oldRate := range []int32{0, 8, 12, 16} {
			for _, nb := range []int32{2, 4} {
				for _, order := range []int32{10, 16} {
					for signal := int8(0); signal <= 2; signal++ {
						g := opuscc.OpusT_silk_decoder_state{Ffs_kHz: rate, Fframe_length: rate * 5 * nb, Fsubfr_length: rate * 5, Fnb_subfr: nb, FLPC_order: order, FlossCnt: 3, FprevSignalType: 7}
						g.Findices.FsignalType = signal
						g.FsPLC = opuscc.OpusT_silk_PLC_struct{Ffs_kHz: oldRate, FpitchL_Q8: 12345, FprevGain_Q16: [2]int32{77, 88}, Frand_seed: 123, Fsubfr_length: 33, Fnb_subfr: 3}
						for i := range g.FsPLC.FprevLPC_Q12 {
							g.FsPLC.FprevLPC_Q12[i] = 123
						}
						ctrl := opuscc.OpusT_silk_decoder_control{FGains_Q16: [4]int32{65536, 123456, 76543, 2147483647}, FpitchL: [4]int32{40, 80, 120, 160}, FLTP_scale_Q14: 8192}
						for i := range ctrl.FLTPCoef_Q14 {
							ctrl.FLTPCoef_Q14[i] = int16(i*103 - 700)
						}
						for i := range ctrl.FPredCoef_Q12[1] {
							ctrl.FPredCoef_Q12[1][i] = int16(i * 137)
						}
						c, cc := g, ctrl
						before := ctrl
						opuscc.ComparePLCDispatch(nil, &g, &ctrl, nil, 0, 0)
						if r := nativePLCDispatch(&c, &cc, nil, 0, 0); r != 0 || g != c || ctrl != cc || ctrl != before {
							t.Fatal("PLC dispatch", rate, oldRate, nb, order, signal, r, g.FsPLC, c.FsPLC)
						}
					}
				}
			}
		}
	}
}

func TestPLCUpdateAgainstC(t *testing.T) {
	for _, nb := range []int32{2, 4} {
		for _, order := range []int32{10, 16} {
			for signal := int8(0); signal <= 2; signal++ {
				for _, gain := range []int16{-100, 0, 500, 3000, 5000, 15000} {
					g := opuscc.OpusT_silk_decoder_state{Ffs_kHz: 16, Fsubfr_length: 80, Fnb_subfr: nb, FLPC_order: order, FprevSignalType: 7}
					g.Findices.FsignalType = signal
					g.FsPLC.FpitchL_Q8 = 12345
					g.FsPLC.Frand_seed = 777
					for i := range g.FsPLC.FprevLPC_Q12 {
						g.FsPLC.FprevLPC_Q12[i] = 123
					}
					ctrl := opuscc.OpusT_silk_decoder_control{FGains_Q16: [4]int32{65536, 123456, 2147483647, 76543}, FLTP_scale_Q14: 23456, FpitchL: [4]int32{100, 180, 250, 300}}
					for i := range ctrl.FLTPCoef_Q14 {
						ctrl.FLTPCoef_Q14[i] = gain + int16(i%5)
					}
					for i := range ctrl.FPredCoef_Q12[1] {
						ctrl.FPredCoef_Q12[1][i] = int16(i * 137)
					}
					c := g
					before := ctrl
					opuscc.ComparePLCUpdate(&g, &ctrl)
					nativePLCUpdate(&c, &ctrl)
					if g != c || ctrl != before {
						t.Fatal(nb, order, signal, gain, g.FsPLC, c.FsPLC)
					}
				}
			}
		}
	}
}
