package opuscc

import (
	"runtime"
	"testing"
	"unsafe"
	"weak"
)

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
