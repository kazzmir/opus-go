package opuscc

import "testing"

func TestDecoderResetPointers(t *testing.T) {
	want := OpusT_silk_decoder_state{Ffirst_frame_after_reset: 1, Fprev_gain_Q16: 65536}
	want.FsCNG.Frand_seed = 3176576
	want.FsPLC.FprevGain_Q16 = [2]int32{65536, 65536}
	want.FsPLC.Fsubfr_length = 20
	want.FsPLC.Fnb_subfr = 2
	for _, reset := range []func(*OpusT_silk_decoder_state) int32{
		func(s *OpusT_silk_decoder_state) int32 { return Opus_silk_reset_decoder(nil, s) },
		func(s *OpusT_silk_decoder_state) int32 { return Opus_silk_init_decoder(nil, s) },
	} {
		wrapped := struct {
			before int64
			state  OpusT_silk_decoder_state
			after  int64
		}{before: 123, after: 456}
		s := &wrapped.state
		s.Fexc_Q14[319] = 123
		s.FoutBuf[479] = -123
		s.Farch = 99
		s.Ffs_kHz = 16
		s.FLPC_order = 16
		s.FpsNLSF_CB = 1234
		s.FsCNG.Frand_seed = 99
		s.FsPLC.FpitchL_Q8 = 777
		for i := 0; i < 2; i++ {
			if result := reset(s); result != 0 || *s != want || wrapped.before != 123 || wrapped.after != 456 {
				t.Fatalf("reset=%d state differs=%v guards=%d,%d", result, *s != want, wrapped.before, wrapped.after)
			}
		}
	}
}
