//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
)

func TestDecodeCoreResidualPCMAgainstC(t *testing.T) {
	for _, sample := range []int32{-2147483648, -100000003, -1, 0, 1, 100000003, 2147483647} {
		for _, gain := range []int32{-2147483648, -65536, 0, 1024, 131072, 2147483647} {
			g := opuscc.CompareDecodeCorePCM(sample, gain)
			n := nativePLCPCM(sample, gain)
			if g != n {
				t.Fatal("PCM", sample, gain, g, n)
			}
		}
	}
	for _, excitation := range []int32{-16000000, -1, 0, 1, 16000000} {
		for _, prediction := range []int32{-1000000, -1, 0, 1, 1000000} {
			g := opuscc.CompareDecodeCoreResidual(excitation, prediction)
			n := nativeDecodeCoreResidual(excitation, prediction)
			if g != n {
				t.Fatal("residual", excitation, prediction, g, n)
			}
		}
	}
}
func TestDecodeCoreLTPStorageAgainstC(t *testing.T) {
	for _, shape := range [][3]int32{{160, 160, 240}, {240, 360, 480}, {320, 640, 640}} {
		for _, lag := range []int32{0, 1, 40, 80} {
			for _, gain := range []int32{-2147483648, -65536, 0, 65536, 2147483647} {
				samples := make([]int16, shape[0])
				for i := range samples {
					samples[i] = int16(i*137 - 20000)
				}
				h := make([]int32, shape[2]+2)
				for i := range h {
					h[i] = int32(i*17000003 - 2100000000)
				}
				want := slices.Clone(h)
				opuscc.CompareDecodeCoreLTPWhiten(h[1:len(h)-1], samples, shape[1], shape[0], lag, gain)
				nativeDecodeCoreLTP(want[1:len(want)-1], samples, shape[1], shape[0], lag, gain, 0)
				if !slices.Equal(h, want) {
					t.Fatal("LTP rewhitening", shape, lag, gain)
				}
				opuscc.CompareDecodeCoreLTPScale(h[1:len(h)-1], shape[1], lag, gain)
				nativeDecodeCoreLTP(want[1:len(want)-1], nil, shape[1], shape[0], lag, gain, 1)
				if !slices.Equal(h, want) {
					t.Fatal("LTP scaling", shape, lag, gain)
				}
			}
		}
	}
	h := []int32{-2147483648, 2147483647, -1, 0, 100000003, 12345, -54321, 2147483647}
	for _, b := range [][5]int16{{-32768, 32767, -1, 16384, 12345}, {300, -150, 1200, -100, 75}, {}} {
		for index := int32(4); index < int32(len(h)); index++ {
			g := opuscc.CompareDecodeCoreLTPPrediction(h, index, &b)
			n := nativePLCLTPPrediction(h, index, &b)
			if g != n {
				t.Fatal("Q13 MAC narrowing", index, b, g, n)
			}
		}
	}
}
func TestDecodeCoreWhiteningAgainstC(t *testing.T) {
	for _, rate := range []int32{8, 12, 16} {
		for _, order := range []int32{10, 16} {
			for _, k := range []int32{0, 2} {
				d := opuscc.OpusT_silk_decoder_state{Fltp_mem_length: rate * 20, Fsubfr_length: rate * 5, FLPC_order: order}
				for i := range d.FoutBuf {
					d.FoutBuf[i] = int16(i*71 - 20000)
				}
				for _, start := range []int32{1, 17, d.Fltp_mem_length - order - 1} {
					a := &[16]int16{120, -80, 60, -45, 30, -20, 15, -10, 8, -5, 4, -3, 2, -2, 1, -1}
					samples := make([]int16, d.Fltp_mem_length+2)
					for i := range samples {
						samples[i] = 123
					}
					samples[0], samples[len(samples)-1] = 77, 88
					want := slices.Clone(samples)
					before := d
					opuscc.CompareDecodeCoreWhiten(&d, samples[1:len(samples)-1], a, start, k)
					nativeLPCAnalysis(want[start+1:len(want)-1], d.FoutBuf[start+k*d.Fsubfr_length:d.Fltp_mem_length+k*d.Fsubfr_length], a[:order])
					if !slices.Equal(samples, want) || d != before {
						t.Fatal("whitening", rate, order, k, start)
					}
				}
			}
		}
	}
}
func TestDecodeCoreCoefficientsAgainstC(t *testing.T) {
	for _, order := range []int32{0, 10, 16} {
		for k := int32(0); k < 4; k++ {
			d := opuscc.OpusT_silk_decoder_state{FLPC_order: order}
			c := opuscc.OpusT_silk_decoder_control{}
			for row := range c.FPredCoef_Q12 {
				for i := range c.FPredCoef_Q12[row] {
					c.FPredCoef_Q12[row][i] = int16(row*1000 + i*71 - 900)
				}
			}
			for i := range c.FLTPCoef_Q14 {
				c.FLTPCoef_Q14[i] = int16(i*31 - 400)
			}
			before := c
			var snapshot [16]int16
			for i := range snapshot {
				snapshot[i] = 123
			}
			want := snapshot
			a, b := opuscc.CompareDecodeCoreCoefficients(&d, &c, k, &snapshot)
			na, nb := nativeDecodeCoreCoefficients(&c, order, k, &want)
			if *a != na || *b != nb || snapshot != want || c != before {
				t.Fatal("coefficients", order, k, *a, na, *b, nb, snapshot, want)
			}
		}
	}
}
func TestDecodeCoreExcitationAgainstC(t *testing.T) {
	for _, length := range []int32{0, 1, 17, 80, 160, 320} {
		for _, seed := range []int8{-128, -1, 0, 1, 17, 127} {
			for _, signal := range []int8{0, 1, 2} {
				for _, q := range []int8{0, 1} {
					d := opuscc.OpusT_silk_decoder_state{Fframe_length: length}
					d.Findices.FSeed = seed
					for i := range d.Fexc_Q14 {
						d.Fexc_Q14[i] = 1234567
					}
					c := d
					pulses := make([]int16, length+2)
					pulses[0], pulses[len(pulses)-1] = 77, 88
					edges := []int16{-32768, 32767, -1, 0, 1, 13, -14}
					for i := int32(0); i < length; i++ {
						pulses[i+1] = edges[i%int32(len(edges))]
					}
					before := slices.Clone(pulses)
					offset := int32(opuscc.Opus_silk_Quantization_Offsets_Q10[signal>>1][q])
					g := opuscc.CompareDecodeCoreExcitation(&d, &pulses[1], offset)
					n := nativeDecodeCoreExcitation(&c, pulses[1:len(pulses)-1], offset)
					if g != n || d != c || !slices.Equal(pulses, before) {
						t.Fatal("excitation", length, seed, signal, q, g, n)
					}
				}
			}
		}
	}
}
func TestDecodeCoreTransitionAgainstC(t *testing.T) {
	for _, loss := range []int32{0, 1, -1} {
		for _, prev := range []int32{1, 2} {
			for _, signal := range []int8{0, 1, 2} {
				for k := int32(0); k < 4; k++ {
					d := opuscc.OpusT_silk_decoder_state{FlossCnt: loss, FprevSignalType: prev, FlagPrev: 77}
					d.Findices.FsignalType = signal
					ctrl := opuscc.OpusT_silk_decoder_control{FpitchL: [4]int32{101, 102, 103, 104}}
					for i := range ctrl.FLTPCoef_Q14 {
						ctrl.FLTPCoef_Q14[i] = int16(i*71 - 900)
					}
					c, cc := d, ctrl
					g := opuscc.CompareDecodeCoreTransition(&d, &ctrl, k)
					n := nativeDecodeCoreTransition(&c, &cc, k)
					if g != n || d != c || ctrl != cc {
						t.Fatal("transition", loss, prev, signal, k, g, n)
					}
				}
			}
		}
	}
}
func TestDecodeFrameFinishAgainstC(t *testing.T) {
	for _, nb := range []int32{2, 4} {
		for _, alias := range []int32{0, 1, 2} {
			d := opuscc.OpusT_silk_decoder_state{Fnb_subfr: nb, FlagPrev: 17}
			ctrl := opuscc.OpusT_silk_decoder_control{FpitchL: [4]int32{41, 42, 43, 44}}
			c, cc := d, ctrl
			count, expected := int32(-77), int32(-77)
			p := &count
			if alias == 1 {
				p = &d.FlagPrev
			} else if alias == 2 {
				p = &ctrl.FpitchL[nb-1]
			}
			opuscc.CompareDecodeFrameFinish(&d, &ctrl, p, 160)
			nativeDecodeFrameFinish(&c, &cc, &expected, 160, alias)
			if d != c || ctrl != cc || *p != expected {
				t.Fatal("finish", nb, alias, d.FlagPrev, c.FlagPrev, *p, expected)
			}
		}
	}
}
func TestDecodeFrameHistoryAgainstC(t *testing.T) {
	for _, rate := range []int32{8, 12, 16} {
		for _, nb := range []int32{2, 4} {
			d := opuscc.OpusT_silk_decoder_state{Fltp_mem_length: rate * 20, Fframe_length: rate * 5 * nb}
			for i := range d.FoutBuf {
				d.FoutBuf[i] = int16(i*37 - 1000)
			}
			c := d
			frame := make([]int16, d.Fframe_length)
			for i := range frame {
				frame[i] = int16(i*71 - 9000)
			}
			opuscc.CompareDecodeFrameHistory(&d, frame)
			nativeDecodeFrameHistory(&c, frame)
			if d != c {
				t.Fatal("frame history", rate, nb)
			}
		}
	}
}
func TestDecodeCoreHistoryAgainstC(t *testing.T) {
	for _, rate := range []int32{8, 12, 16} {
		d := opuscc.OpusT_silk_decoder_state{Fltp_mem_length: rate * 20, Fsubfr_length: rate * 5}
		for i := range d.FoutBuf {
			d.FoutBuf[i] = int16(i*3 - 900)
		}
		c := d
		frame := make([]int16, 2*d.Fsubfr_length)
		for i := range frame {
			frame[i] = int16(i*31 - 900)
		}
		opuscc.CompareDecodeCoreHistory(&d, frame)
		nativeDecodeCoreHistory(&c, frame)
		if d != c {
			t.Fatal("history", rate)
		}
	}
}
func TestDecodeCoreOutputAliasesAgainstC(t *testing.T) {
	for _, rate := range []int32{8, 12, 16} {
		for _, signal := range []int8{1, 2} {
			for _, alias := range []int32{1, 2} {
				order := int32(10)
				if rate == 16 {
					order = 16
				}
				d := opuscc.OpusT_silk_decoder_state{Ffs_kHz: rate, Fframe_length: rate * 20, Fsubfr_length: rate * 5, Fnb_subfr: 4, Fltp_mem_length: rate * 20, FLPC_order: order, Fprev_gain_Q16: 65536}
				d.Findices.FsignalType = signal
				d.Findices.FquantOffsetType = 1
				d.Findices.FNLSFInterpCoef_Q2 = 0
				d.Findices.FSeed = 17
				for i := range d.FoutBuf {
					d.FoutBuf[i] = int16((i*37)%1000 - 500)
				}
				ctrl := opuscc.OpusT_silk_decoder_control{FGains_Q16: [4]int32{65536, 72000, 68000, 76000}, FLTP_scale_Q14: 12288}
				for k := range ctrl.FPredCoef_Q12 {
					ctrl.FPredCoef_Q12[k] = [16]int16{120, -80, 60, -45, 30, -20, 15, -10, 8, -5, 4, -3, 2, -2, 1, -1}
				}
				for k := int32(0); k < 4; k++ {
					ctrl.FpitchL[k] = rate * 5
					copy(ctrl.FLTPCoef_Q14[k*5:k*5+5], []int16{300, -150, 1200, -100, 75})
				}
				c, cc := d, ctrl
				pulses := make([]int16, d.Fframe_length)
				for i := range pulses {
					pulses[i] = int16((i*7)%9 - 4)
				}
				want := slices.Clone(pulses)
				output := &d.FoutBuf[0]
				if alias == 2 {
					output = &pulses[0]
				}
				opuscc.CompareDecodeCore(nil, &d, &ctrl, output, &pulses[0])
				r := nativeDecodeCoreAlias(&c, &cc, nil, want, alias)
				if r != 0 || d != c || ctrl != cc || !slices.Equal(pulses, want) {
					t.Fatal("output alias", rate, signal, alias, r)
				}
			}
		}
	}
}
func TestDecodeCoreAgainstC(t *testing.T) {
	for _, rate := range []int32{8, 12, 16} {
		for _, nb := range []int32{2, 4} {
			for _, signal := range []int8{0, 1, 2} {
				for _, interp := range []int8{0, 4} {
					for _, loss := range []int32{0, 1} {
						order := int32(10)
						if rate == 16 {
							order = 16
						}
						d := opuscc.OpusT_silk_decoder_state{Ffs_kHz: rate, Fframe_length: rate * 5 * nb, Fsubfr_length: rate * 5, Fnb_subfr: nb, Fltp_mem_length: rate * 20, FLPC_order: order, Fprev_gain_Q16: 65536, FprevSignalType: 2, FlagPrev: rate * 5, FlossCnt: loss}
						d.Findices.FsignalType = signal
						d.Findices.FquantOffsetType = 1
						d.Findices.FNLSFInterpCoef_Q2 = interp
						d.Findices.FSeed = 17
						for i := range d.FoutBuf {
							d.FoutBuf[i] = int16((i*37)%1000 - 500)
						}
						for i := range d.FsLPC_Q14_buf {
							d.FsLPC_Q14_buf[i] = int32(i*97 - 800)
						}
						ctrl := opuscc.OpusT_silk_decoder_control{FGains_Q16: [4]int32{65536, 72000, 68000, 76000}, FLTP_scale_Q14: 12288}
						for k := range ctrl.FPredCoef_Q12 {
							ctrl.FPredCoef_Q12[k] = [16]int16{120, -80, 60, -45, 30, -20, 15, -10, 8, -5, 4, -3, 2, -2, 1, -1}
						}
						for k := int32(0); k < nb; k++ {
							ctrl.FpitchL[k] = rate * 5
							copy(ctrl.FLTPCoef_Q14[k*5:k*5+5], []int16{300, -150, 1200, -100, 75})
						}
						c, cc := d, ctrl
						pulses := make([]int16, d.Fframe_length)
						for i := range pulses {
							pulses[i] = int16((i*7)%9 - 4)
						}
						before := slices.Clone(pulses)
						out := make([]int16, d.Fframe_length+2)
						out[0], out[len(out)-1] = 77, 88
						want := slices.Clone(out)
						for frame := 0; frame < 3; frame++ {
							opuscc.CompareDecodeCore(nil, &d, &ctrl, &out[1], &pulses[0])
							r := nativeDecodeCore(&c, &cc, want[1:len(want)-1], pulses)
							if r != 0 || d != c || ctrl != cc || !slices.Equal(out, want) || !slices.Equal(pulses, before) {
								t.Fatal("decode core", rate, nb, signal, interp, loss, frame, r, d.Fprev_gain_Q16, c.Fprev_gain_Q16)
							}
						}
					}
				}
			}
		}
	}
}
