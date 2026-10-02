package opuscc

import (
	"runtime"
	"testing"
	"unsafe"
	"weak"
)

func TestPLCLTPHistoryPointers(t *testing.T) {
	h := []int32{-2147483648, 2147483647, -1, 0, 100000003, 12345, -54321, 2147483647}
	b := &[5]int16{-32768, 32767, -1, 16384, 12345}
	entropyInitGrowStack(12)
	runtime.GC()
	for index := int32(4); index < int32(len(h)); index++ {
		want := int32(2)
		for j := int32(0); j < 5; j++ {
			want = int32(int64(want) + (int64(h[index-j]) * int64(b[j]) >> 16))
		}
		if got := silkPLCLTPPrediction(h, index, b); got != want {
			t.Fatal("LTP history", index, got, want)
		}
	}
	b[0] = 0
	h[4] = 1
	if got := silkPLCLTPPrediction(h, 4, b); got != int32(int64(2)+(int64(h[3])*int64(b[1])>>16)+(int64(h[2])*int64(b[2])>>16)+(int64(h[1])*int64(b[3])>>16)+(int64(h[0])*int64(b[4])>>16)) {
		t.Fatal("live taps/history", got)
	}
}

func TestPLCWhiteningPointers(t *testing.T) {
	for _, length := range []int32{160, 240, 320} {
		for _, order := range []int32{10, 16} {
			for _, index := range []int32{1, 17, length - order - 1} {
				d := &OpusT_silk_decoder_state{Fltp_mem_length: length, FLPC_order: order, FpsNLSF_CB: cloneTestNLSFCodebook(&Opus_silk_NLSF_CB_WB)}
				for i := range d.FoutBuf {
					d.FoutBuf[i] = int16((i*71)%60000 - 30000)
				}
				a := new([16]int16)
				for i := range a {
					a[i] = int16(i*137 - 700)
				}
				samples := make([]int16, length+2)
				for i := range samples {
					samples[i] = 123
				}
				samples[0], samples[len(samples)-1] = 77, 88
				want := append([]int16(nil), samples...)
				Opus_silk_LPC_analysis_filter(nil, &want[1+index], &d.FoutBuf[index], &a[0], length-index, order, 0)
				before := *d
				entropyInitGrowStack(12)
				runtime.GC()
				silkPLCWhiten(nil, d, samples[1:len(samples)-1], a, index, 0)
				for i := range samples {
					if samples[i] != want[i] {
						t.Fatal("whitening cursor", length, order, index, i, samples[i], want[i])
					}
				}
				if *d != before {
					t.Fatal("whitening changed decoder")
				}
			}
		}
	}
}

func TestPLCLPCHistoryPointers(t *testing.T) {
	for _, length := range []int32{0, 1, 10, 16, 17, 80, 320} {
		for _, order := range []int32{10, 16} {
			d := &OpusT_silk_decoder_state{Fframe_length: length, FLPC_order: order, FpsNLSF_CB: cloneTestNLSFCodebook(&Opus_silk_NLSF_CB_WB)}
			for i := range d.FsLPC_Q14_buf {
				d.FsLPC_Q14_buf[i] = int32(i*1000003) - 8000000
			}
			a := new([MAX_LPC_ORDER]int16)
			for i := range a {
				a[i] = int16(i*137 - 700)
			}
			history := make([]int32, 18+length)
			history[0], history[len(history)-1] = 77, 88
			for i := int32(0); i < length; i++ {
				history[17+i] = int32(int64(i) * 1000000003)
			}
			want := append([]int32(nil), history...)
			copy(want[1:17], d.FsLPC_Q14_buf[:])
			wf := make([]int16, length)
			for i := int32(0); i < length; i++ {
				pred := order >> 1
				for j := int32(0); j < order; j++ {
					pred = int32(int64(pred) + (int64(want[17+i-j-1]) * int64(a[j]) >> 16))
				}
				p := int64(int32(uint32(min(max(pred, -2147483648>>4), 2147483647>>4)) << 4))
				want[17+i] = int32(min(max(int64(want[17+i])+p, int64(-2147483648)), int64(2147483647)))
				wf[i] = int16(min(max(((int32(int64(want[17+i])*1024>>16)>>7)+1)>>1, -32768), 32767))
			}
			pcm := make([]int16, length+2)
			pcm[0], pcm[len(pcm)-1] = 77, 88
			entropyInitGrowStack(12)
			runtime.GC()
			silkPLCLPC(nil, d, history[1:len(history)-1], a, unsafe.SliceData(pcm[1:len(pcm)-1]), 1024)
			for i := range history {
				if history[i] != want[i] {
					t.Fatal("LPC working history", length, order, i)
				}
			}
			for i := range wf {
				if pcm[i+1] != wf[i] {
					t.Fatal("LPC PCM", length, order, i)
				}
			}
			for i, v := range d.FsLPC_Q14_buf {
				if v != want[1+int(length)+i] {
					t.Fatal("history saved after PCM", length, order, i)
				}
			}
			if pcm[0] != 77 || pcm[len(pcm)-1] != 88 || d.FpsNLSF_CB.Forder != 16 {
				t.Fatal("history guard/owner")
			}
		}
	}
	d := OpusT_silk_decoder_state{FLPC_order: 8, Fframe_length: 1, FsLPC_Q14_buf: [16]int32{123}}
	h := make([]int32, 17)
	panicked := false
	func() {
		defer func() { panicked = recover() != nil }()
		silkPLCLPC(nil, &d, h, new([16]int16), nil, 1024)
	}()
	if !panicked || h[0] != 123 || d.FsLPC_Q14_buf[0] != 123 {
		t.Fatal("copy-before-LPC assertion order")
	}
	for _, f := range []struct{ x, p, w int32 }{{2147483647, 1, 2147483647}, {-2147483648, -1, -2147483648}, {0, 2147483647, 2147483632}, {0, -2147483648, -2147483648}} {
		if got := silkPLCAddPrediction(f.x, f.p); got != f.w {
			t.Fatal("saturating prediction", f, got)
		}
	}
}

// This alias is Go-only: short PCM into int32 state violates C effective types.
func TestPLCLPCStatePCMOrderPointers(t *testing.T) {
	d := OpusT_silk_decoder_state{FLPC_order: 16, Fframe_length: 16}
	h := make([]int32, 32)
	for i := 0; i < 16; i++ {
		d.FsLPC_Q14_buf[i] = int32(i + 77)
		h[16+i] = int32(i-8) << 16
	}
	silkPLCLPC(nil, &d, h, new([16]int16), (*int16)(unsafe.Pointer(&d.FsLPC_Q14_buf[0])), 1024)
	for i := 0; i < 16; i++ {
		want := (int32(i-8) << 16) + 128
		if h[16+i] != want || d.FsLPC_Q14_buf[i] != want {
			t.Fatal("PCM-before-history-save order", i, h[16+i], d.FsLPC_Q14_buf[i], want)
		}
	}
}

func TestPLCConcealPCMPointers(t *testing.T) {
	pcm := make([]int16, 3)
	pcm[0], pcm[2] = 77, 88
	entropyInitGrowStack(12)
	runtime.GC()
	for _, sample := range []int32{-2147483648, -8388609, -8388608, -128, -1, 0, 127, 128, 8388352, 8388608, 2147483647} {
		for _, gain := range []int32{-2147483648, -65536, -1, 0, 1, 1024, 65536, 2147483647} {
			scaled := int32(int64(sample) * int64(gain) >> 16)
			want := int16(min(max(((scaled>>7)+1)>>1, -32768), 32767))
			silkPLCPCMStore(pcm[1:2], 0, sample, gain)
			if pcm[1] != want || pcm[0] != 77 || pcm[2] != 88 {
				t.Fatal("PLC PCM", sample, gain, pcm, want)
			}
		}
	}
	for _, f := range []struct {
		s, g int32
		want int16
	}{{128, 65536, 1}, {-128, 65536, 0}, {-129, 65536, -1}, {8388608, 65536, 32767}, {-8388609, 65536, -32768}, {2147483647, 2147483647, -256}} {
		if got := silkPLCPCM(f.s, f.g); got != f.want {
			t.Fatal("PCM narrowing/rounding/saturation", f, got)
		}
	}
}

func TestPLCRandomPointers(t *testing.T) {
	random, owner := func() ([]int32, weak.Pointer[OpusT_silk_NLSF_CB_struct]) {
		cb := cloneTestNLSFCodebook(&Opus_silk_NLSF_CB_WB)
		d := &OpusT_silk_decoder_state{FpsNLSF_CB: cb}
		for i := range d.Fexc_Q14 {
			d.Fexc_Q14[i] = int32(i*1000003) - 160000000
		}
		return silkPLCRandom(d, 192), weak.Make(cb)
	}()
	entropyInitGrowStack(12)
	runtime.GC()
	if len(random) != 128 || owner.Value() == nil {
		t.Fatal("random interior owner")
	}
	for i := int32(0); i < 128; i++ {
		for _, scale := range []int16{-32768, -1, 0, 16384, 32767} {
			for _, prediction := range []int32{-2147483648, -1, 0, 2147483647} {
				want := int32(uint32(int32(int64(prediction)+(int64(int32((i+192)*1000003)-160000000)*int64(scale)>>16))) << 2)
				if got := silkPLCNoise(prediction, random, i, scale); got != want {
					t.Fatal("random Q14", i, scale, prediction, got, want)
				}
			}
		}
	}
	random[0] = 1 << 20
	if got := silkPLCNoise(2, random, 0, 16384); got != (2+(1<<18))<<2 {
		t.Fatal("live random load", got)
	}
}

func TestPLCLTPCoefficientPointers(t *testing.T) {
	b, owner := func() (*[LTP_ORDER]int16, weak.Pointer[OpusT_silk_NLSF_CB_struct]) {
		cb := cloneTestNLSFCodebook(&Opus_silk_NLSF_CB_WB)
		d := &OpusT_silk_decoder_state{FpsNLSF_CB: cb}
		d.FsPLC.FLTPCoef_Q14 = [5]int16{-32768, 32767, -1, 0, 12345}
		return &d.FsPLC.FLTPCoef_Q14, weak.Make(cb)
	}()
	entropyInitGrowStack(12)
	runtime.GC()
	if owner.Value() == nil {
		t.Fatal("coefficient interior owner")
	}
	for _, gain := range []int32{-32768, -32767, 0, 16384, 32767, 32768, 65535} {
		before := *b
		silkPLCDecayLTP(b, gain)
		for j := range b {
			want := int16(int32(int16(gain)) * int32(before[j]) >> 15)
			if b[j] != want {
				t.Fatal("LTP decay", gain, j, b[j], want)
			}
		}
	}
	runtime.KeepAlive(b)
}

func TestPLCConcealStateOwnersPointers(t *testing.T) {
	plc, owner := func() (*OpusT_silk_PLC_struct, weak.Pointer[OpusT_silk_NLSF_CB_struct]) {
		cb := cloneTestNLSFCodebook(&Opus_silk_NLSF_CB_WB)
		d := &OpusT_silk_decoder_state{FpsNLSF_CB: cb, FsPLC: OpusT_silk_PLC_struct{Frand_seed: 123, FprevGain_Q16: [2]int32{65536, 131072}}}
		return silkPLCConcealState(d), weak.Make(cb)
	}()
	entropyInitGrowStack(12)
	runtime.GC()
	runtime.GC()
	cb := owner.Value()
	if cb == nil || cb.Forder != 16 || plc.Frand_seed != 123 || plc.FprevGain_Q16 != [2]int32{65536, 131072} {
		t.Fatal("sole PLC interior lost scanned decoder owner")
	}
	plc.FpitchL_Q8 = 77
	runtime.KeepAlive(plc)
}

func TestPLCDispatchFramePointers(t *testing.T) {
	d := OpusT_silk_decoder_state{Ffs_kHz: 16, Fframe_length: 320, Fsubfr_length: 80, Fnb_subfr: 4, FLPC_order: 16}
	d.Findices.FsignalType = TYPE_UNVOICED
	c := OpusT_silk_decoder_control{FGains_Q16: [4]int32{1, 2, 3, 4}}
	frame := make([]int16, 322)
	for i := range frame {
		frame[i] = int16(i*17 - 300)
	}
	before := append([]int16(nil), frame...)
	entropyInitGrowStack(12)
	runtime.GC()
	silk_PLC(nil, &d, &c, &frame[1], 0, 0)
	for i := range frame {
		if frame[i] != before[i] {
			t.Fatal("nonloss touched PCM", i)
		}
	}
	// Frame aliases are unconsumed on nonloss: even a numeric state/control
	// interior must not be read or written, and nil is equally valid.
	wd, wc := d, c
	silk_PLC(nil, &d, &c, (*int16)(unsafe.Pointer(&c.FpitchL[0])), 0, 0)
	silk_PLC(nil, &wd, &wc, nil, 0, 0)
	if d != wd || c != wc {
		t.Fatal("unused frame alias")
	}
}

func TestPLCDispatchControlPointers(t *testing.T) {
	for _, rate := range []int32{8, 12, 16} {
		for _, nb := range []int32{2, 4} {
			for _, order := range []int32{10, 16} {
				for _, signal := range []int8{0, 1, 2} {
					d := &OpusT_silk_decoder_state{Ffs_kHz: rate, Fframe_length: rate * 5 * nb, Fsubfr_length: rate * 5, Fnb_subfr: nb, FLPC_order: order, FlossCnt: 3}
					d.Findices.FsignalType = signal
					c := &OpusT_silk_decoder_control{FGains_Q16: [4]int32{1, 2, 3, 4}, FLTP_scale_Q14: 8192, FpitchL: [4]int32{40, 80, 120, 160}}
					for i := range c.FPredCoef_Q12[1] {
						c.FPredCoef_Q12[1][i] = int16(i * 13)
					}
					for i := range c.FLTPCoef_Q14 {
						c.FLTPCoef_Q14[i] = int16(i * 17)
					}
					w, wc := *d, *c
					silkPLCRate(nil, &w)
					silk_PLC_update(nil, &w, &wc)
					entropyInitGrowStack(12)
					runtime.GC()
					silk_PLC(nil, d, c, nil, 0, -37)
					if *d != w || *c != wc || d.FlossCnt != 3 {
						t.Fatal("typed PLC update", rate, nb, order, signal)
					}
				}
			}
		}
	}
	// Go-only invalid-control fixture: rate reset and the preceding update stores
	// are committed before control is first dereferenced.
	d := OpusT_silk_decoder_state{Ffs_kHz: 16, Fframe_length: 320, Fnb_subfr: 4, Fsubfr_length: 80, FLPC_order: 16}
	d.Findices.FsignalType = TYPE_UNVOICED
	panicked := false
	func() { defer func() { panicked = recover() != nil }(); silk_PLC(nil, &d, nil, nil, 0, 0) }()
	if !panicked || d.FsPLC.Ffs_kHz != 16 || d.FsPLC.FpitchL_Q8 != 16*18*256 || d.FsPLC.FprevGain_Q16 != [2]int32{65536, 65536} || d.FprevSignalType != TYPE_UNVOICED {
		t.Fatal("reset/update-before-control order", panicked, d.FsPLC)
	}
}

func TestPLCDispatchStatePointers(t *testing.T) {
	for _, rate := range []int32{8, 12, 16} {
		d := func() *OpusT_silk_decoder_state {
			return &OpusT_silk_decoder_state{Fframe_length: rate * 20, Ffs_kHz: rate, FpsNLSF_CB: cloneTestNLSFCodebook(&Opus_silk_NLSF_CB_WB), FlossCnt: 7, FsPLC: OpusT_silk_PLC_struct{Ffs_kHz: 0, Frand_seed: 123}}
		}()
		entropyInitGrowStack(12)
		runtime.GC()
		silkPLCRate(nil, d)
		want := d.FsPLC
		if want.Ffs_kHz != rate || want.FpitchL_Q8 != rate*20*128 || want.FprevGain_Q16 != [2]int32{65536, 65536} || want.Fsubfr_length != 20 || want.Fnb_subfr != 2 || want.Frand_seed != 123 || d.FlossCnt != 7 || d.FpsNLSF_CB.Forder != 16 {
			t.Fatal("PLC state owner/reset", rate, want)
		}
		d.FsPLC.FpitchL_Q8 = 77
		silkPLCRate(nil, d)
		if d.FsPLC.FpitchL_Q8 != 77 {
			t.Fatal("matching rate reset")
		}
	}
}

func TestPLCUpdatePointers(t *testing.T) {
	d := OpusT_silk_decoder_state{Ffs_kHz: 16, Fnb_subfr: 4, Fsubfr_length: 80, FLPC_order: 10}
	d.Findices.FsignalType = TYPE_UNVOICED
	for i := range d.FsPLC.FprevLPC_Q12 {
		d.FsPLC.FprevLPC_Q12[i] = 77
	}
	c := OpusT_silk_decoder_control{FGains_Q16: [4]int32{1, 2, 3, 4}, FLTP_scale_Q14: 8192}
	for i := range c.FPredCoef_Q12[1] {
		c.FPredCoef_Q12[1][i] = int16(i)
	}
	before := c
	silk_PLC_update(nil, &d, &c)
	if c != before || d.FprevSignalType != TYPE_UNVOICED || d.FsPLC.FpitchL_Q8 != 16*18*256 || d.FsPLC.FprevGain_Q16 != [2]int32{3, 4} {
		t.Fatal(d.FsPLC)
	}
	for i, v := range d.FsPLC.FprevLPC_Q12 {
		want := int16(77)
		if i < 10 {
			want = int16(i)
		}
		if v != want {
			t.Fatal(i, v, want)
		}
	}
}
