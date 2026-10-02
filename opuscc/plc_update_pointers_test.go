package opuscc

import (
	"runtime"
	"testing"
)

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
