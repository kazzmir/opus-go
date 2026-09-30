package opuscc

import (
	"runtime"
	"testing"
)

func TestDecoderSetFSPointers(t *testing.T) {
	wrapped := struct {
		before uint64
		state  OpusT_silk_decoder_state
		after  uint64
	}{before: 77, after: 88}
	st := &wrapped.state
	st.Fnb_subfr = 4
	for i := range st.FoutBuf {
		st.FoutBuf[i] = 123
	}
	for i := range st.FsLPC_Q14_buf {
		st.FsLPC_Q14_buf[i] = 456
	}
	entropyInitGrowStack(12)
	runtime.GC()
	if r := Opus_silk_decoder_set_fs(nil, st, 16, 48000); r != 0 {
		t.Fatal(r)
	}
	if st.Fframe_length != 320 || st.Fsubfr_length != 80 || st.FLPC_order != 16 || wrapped.before != 77 || wrapped.after != 88 {
		t.Fatal(wrapped)
	}
	for _, x := range st.FoutBuf {
		if x != 0 {
			t.Fatal("outBuf clear")
		}
	}
	for _, x := range st.FsLPC_Q14_buf {
		if x != 0 {
			t.Fatal("LPC clear")
		}
	}
	st.FoutBuf[0] = 123
	st.FsLPC_Q14_buf[0] = 456
	st.Fresampler_state.FsIIR[0] = 789
	before := *st
	if r := Opus_silk_decoder_set_fs(nil, st, 16, 48000); r != 0 || *st != before {
		t.Fatal("unchanged setup", r)
	}
	st.Fnb_subfr = 2
	if r := Opus_silk_decoder_set_fs(nil, st, 16, 24000); r != 0 || st.Fframe_length != 160 || st.FoutBuf[0] != 123 || st.FsLPC_Q14_buf[0] != 456 {
		t.Fatal("API/frame-only change", r)
	}
	if r := Opus_silk_decoder_set_fs(nil, st, 8, 8000); r != 0 || st.FoutBuf[0] != 0 || st.FsLPC_Q14_buf[0] != 0 || st.FLPC_order != 10 {
		t.Fatal("internal rate change", r)
	}
}

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
