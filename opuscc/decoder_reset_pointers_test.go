package opuscc

import (
	"runtime"
	"slices"
	"testing"
	"unsafe"
)

func seedDecoderNumericFixture(s *OpusT_silk_decoder_state, value byte) {
	*s = OpusT_silk_decoder_state{}
	raw := unsafe.Slice((*byte)(unsafe.Pointer(s)), int(unsafe.Sizeof(*s)))
	ptrSize := int(unsafe.Sizeof(s.FpsNLSF_CB))
	coefs := int(unsafe.Offsetof(s.Fresampler_state)) + int(unsafe.Offsetof(s.Fresampler_state.FCoefs))
	cb := int(unsafe.Offsetof(s.FpsNLSF_CB))
	lag := int(unsafe.Offsetof(s.Fpitch_lag_low_bits_iCDF))
	contour := int(unsafe.Offsetof(s.Fpitch_contour_iCDF))
	for i := range raw {
		if (i >= coefs && i < coefs+ptrSize) || (i >= cb && i < cb+ptrSize) || (i >= lag && i < lag+ptrSize) || (i >= contour && i < contour+ptrSize) {
			continue
		}
		raw[i] = value
	}
	s.FpsNLSF_CB = &Opus_silk_NLSF_CB_WB
	s.Fpitch_lag_low_bits_iCDF = &Opus_silk_uniform8_iCDF[0]
	s.Fpitch_contour_iCDF = &Opus_silk_pitch_contour_iCDF[0]
	s.Fresampler_state.FCoefs = &Opus_silk_Resampler_1_2_COEFS[0]
}

func TestPitchLagTablePointers(t *testing.T) {
	for _, source := range [][]byte{Opus_silk_uniform4_iCDF[:], Opus_silk_uniform6_iCDF[:], Opus_silk_uniform8_iCDF[:]} {
		makeDecoder := func() *OpusT_silk_decoder_state {
			owned := slices.Clone(source)
			return &OpusT_silk_decoder_state{Fpitch_lag_low_bits_iCDF: &owned[0]}
		}
		makeEncoder := func() *OpusT_silk_encoder_state {
			owned := slices.Clone(source)
			return &OpusT_silk_encoder_state{Fpitch_lag_low_bits_iCDF: &owned[0]}
		}
		dec, enc := makeDecoder(), makeEncoder()
		entropyInitGrowStack(12)
		runtime.GC()
		data := [32]byte{17, 255, 88, 1, 192, 0, 77}
		var gd, ge, reference OpusT_ec_dec
		Opus_ec_dec_init(nil, &gd, &data[0], 32)
		ge = gd
		reference = gd
		for i := 0; i < 100; i++ {
			want := Opus_ec_dec_icdf(nil, &reference, &source[0], 8)
			if Opus_ec_dec_icdf(nil, &gd, dec.Fpitch_lag_low_bits_iCDF, 8) != want || Opus_ec_dec_icdf(nil, &ge, enc.Fpitch_lag_low_bits_iCDF, 8) != want || gd != reference || ge != reference {
				t.Fatal("owned lag table", len(source), i)
			}
		}
		Opus_silk_reset_decoder(nil, dec)
		if dec.Fpitch_lag_low_bits_iCDF != nil {
			t.Fatal("lag reset")
		}
	}
	st := OpusT_silk_decoder_state{Fnb_subfr: 4}
	for _, fs := range []int32{8, 12, 16, 8} {
		Opus_silk_decoder_set_fs(nil, &st, fs, 16000)
		want := &Opus_silk_uniform4_iCDF[0]
		if fs == 12 {
			want = &Opus_silk_uniform6_iCDF[0]
		}
		if fs == 16 {
			want = &Opus_silk_uniform8_iCDF[0]
		}
		if st.Fpitch_lag_low_bits_iCDF != want {
			t.Fatal("lag selection", fs)
		}
	}
}

func TestPitchContourTablePointers(t *testing.T) {
	for _, source := range [][]byte{Opus_silk_pitch_contour_NB_iCDF[:], Opus_silk_pitch_contour_10_ms_NB_iCDF[:], Opus_silk_pitch_contour_iCDF[:], Opus_silk_pitch_contour_10_ms_iCDF[:]} {
		makeDecoder := func() *OpusT_silk_decoder_state {
			owned := slices.Clone(source)
			return &OpusT_silk_decoder_state{Fpitch_contour_iCDF: &owned[0]}
		}
		makeEncoder := func() *OpusT_silk_encoder_state {
			owned := slices.Clone(source)
			return &OpusT_silk_encoder_state{Fpitch_contour_iCDF: &owned[0]}
		}
		dec, enc := makeDecoder(), makeEncoder()
		entropyInitGrowStack(12)
		runtime.GC()
		data := [32]byte{17, 255, 88, 1, 192, 0, 77}
		var gd, ge, reference OpusT_ec_dec
		Opus_ec_dec_init(nil, &gd, &data[0], 32)
		ge = gd
		reference = gd
		for i := 0; i < 100; i++ {
			want := Opus_ec_dec_icdf(nil, &reference, &source[0], 8)
			if Opus_ec_dec_icdf(nil, &gd, dec.Fpitch_contour_iCDF, 8) != want || Opus_ec_dec_icdf(nil, &ge, enc.Fpitch_contour_iCDF, 8) != want || gd != reference || ge != reference {
				t.Fatal("owned contour table", len(source), i)
			}
		}
		Opus_silk_reset_decoder(nil, dec)
		if dec.Fpitch_contour_iCDF != nil {
			t.Fatal("contour reset")
		}
	}
	st := OpusT_silk_decoder_state{}
	for _, fs := range []int32{8, 12, 16, 8} {
		for _, sub := range []int32{4, 2, 4} {
			st.Fnb_subfr = sub
			Opus_silk_decoder_set_fs(nil, &st, fs, 16000)
			want := &Opus_silk_pitch_contour_iCDF[0]
			if fs == 8 {
				want = &Opus_silk_pitch_contour_NB_iCDF[0]
			}
			if sub == 2 {
				want = &Opus_silk_pitch_contour_10_ms_iCDF[0]
				if fs == 8 {
					want = &Opus_silk_pitch_contour_10_ms_NB_iCDF[0]
				}
			}
			if st.Fpitch_contour_iCDF != want {
				t.Fatal("contour selection", fs, sub)
			}
		}
	}
}

func TestSilkCodebookReferencePointers(t *testing.T) {
	makeDecoder := func() *OpusT_silk_decoder_state {
		return &OpusT_silk_decoder_state{FpsNLSF_CB: cloneTestNLSFCodebook(&Opus_silk_NLSF_CB_NB_MB)}
	}
	makeEncoder := func() *OpusT_silk_encoder_state {
		return &OpusT_silk_encoder_state{FpsNLSF_CB: cloneTestNLSFCodebook(&Opus_silk_NLSF_CB_WB)}
	}
	dec, enc := makeDecoder(), makeEncoder()
	entropyInitGrowStack(12)
	runtime.GC()
	indices := [17]int8{3}
	g, c := [16]int16{}, [16]int16{}
	Opus_silk_NLSF_decode(nil, &g[0], &indices[0], dec.FpsNLSF_CB)
	Opus_silk_NLSF_decode(nil, &c[0], &indices[0], &Opus_silk_NLSF_CB_NB_MB)
	if g != c {
		t.Fatal("decoder-owned codebook")
	}
	Opus_silk_NLSF_decode(nil, &g[0], &indices[0], enc.FpsNLSF_CB)
	Opus_silk_NLSF_decode(nil, &c[0], &indices[0], &Opus_silk_NLSF_CB_WB)
	if g != c {
		t.Fatal("encoder-owned codebook")
	}
	dec.Fnb_subfr = 4
	Opus_silk_decoder_set_fs(nil, dec, 16, 16000)
	if dec.FpsNLSF_CB != &Opus_silk_NLSF_CB_WB {
		t.Fatal("set_fs codebook")
	}
	Opus_silk_reset_decoder(nil, dec)
	if dec.FpsNLSF_CB != nil || dec.Fresampler_state.FCoefs != nil {
		t.Fatal("reset table ownership")
	}
}

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
		s.FpsNLSF_CB = &Opus_silk_NLSF_CB_WB
		s.FsCNG.Frand_seed = 99
		s.FsPLC.FpitchL_Q8 = 777
		for i := 0; i < 2; i++ {
			if result := reset(s); result != 0 || *s != want || wrapped.before != 123 || wrapped.after != 456 {
				t.Fatalf("reset=%d state differs=%v guards=%d,%d", result, *s != want, wrapped.before, wrapped.after)
			}
		}
	}
}
