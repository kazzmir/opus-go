package opuscc

import (
	"runtime"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestLTPICDFTablePointers(t *testing.T) {
	original := Opus_silk_LTP_gain_iCDF_ptrs
	defer func() { Opus_silk_LTP_gain_iCDF_ptrs = original }()
	for i := range original {
		N := 8 << i
		want := append([]byte(nil), unsafe.Slice(original[i], N)...)
		clone := append([]byte(nil), want...)
		Opus_silk_LTP_gain_iCDF_ptrs[i] = &clone[0]
		clone = nil
		entropyInitGrowStack(12)
		runtime.GC()
		data := []byte{0x73, 0x15, 0x98, 0x52, 0xa3, 0x7b, 0x66, 0x91}
		var got, expected OpusT_ec_ctx
		Opus_ec_dec_init(nil, &got, &data[0], uint32(len(data)))
		Opus_ec_dec_init(nil, &expected, &data[0], uint32(len(data)))
		for j := 0; j < 12; j++ {
			if Opus_ec_dec_icdf(nil, &got, Opus_silk_LTP_gain_iCDF_ptrs[i], 8) != Opus_ec_dec_icdf(nil, &expected, &want[0], 8) || got != expected {
				t.Fatal("typed LTP ICDF owner/state", i, j)
			}
		}
	}
}

func TestLTPVectorGainTablePointers(t *testing.T) {
	original := Opus_silk_LTP_vq_gain_ptrs_Q7
	defer func() { Opus_silk_LTP_vq_gain_ptrs_Q7 = original }()
	for i := range original {
		N := 8 << i
		want := append([]byte(nil), unsafe.Slice(original[i], N)...)
		clone := append([]byte(nil), want...)
		Opus_silk_LTP_vq_gain_ptrs_Q7[i] = &clone[0]
		clone = nil
		entropyInitGrowStack(12)
		runtime.GC()
		got := unsafe.Slice(Opus_silk_LTP_vq_gain_ptrs_Q7[i], N)
		for j := range want {
			if got[j] != want[j] {
				t.Fatal("typed LTP vector gain table owner", i, j)
			}
		}
	}
}

func TestLTPVectorTablePointers(t *testing.T) {
	original := Opus_silk_LTP_vq_ptrs_Q7
	defer func() { Opus_silk_LTP_vq_ptrs_Q7 = original }()
	for i := range original {
		N := 8 << i
		want := append([][LTP_ORDER]int8(nil), unsafe.Slice(original[i], N)...)
		clone := append([][LTP_ORDER]int8(nil), want...)
		Opus_silk_LTP_vq_ptrs_Q7[i] = &clone[0]
		clone = nil
		entropyInitGrowStack(12)
		runtime.GC()
		got := unsafe.Slice(Opus_silk_LTP_vq_ptrs_Q7[i], N)
		for j := range want {
			if got[j] != want[j] {
				t.Fatal("typed LTP vector table owner", i, j)
			}
		}
		for row := 0; row < N; row++ {
			var st OpusT_silk_decoder_state
			st.Fnb_subfr = 4
			Opus_silk_decoder_set_fs(nil, &st, 8, 8000)
			st.Findices.FsignalType = TYPE_VOICED
			st.Findices.FNLSFInterpCoef_Q2 = 4
			st.Findices.FPERIndex = int8(i)
			st.Findices.FGainsIndices = [4]int8{9, 3, 5, 7}
			st.FLastGainIndex = 12
			for k := range st.Findices.FLTPIndex {
				st.Findices.FLTPIndex[k] = int8(row)
			}
			var control OpusT_silk_decoder_control
			Opus_silk_decode_parameters(nil, &st, &control, CODE_INDEPENDENTLY)
			for k := 0; k < 4; k++ {
				for tap := 0; tap < LTP_ORDER; tap++ {
					if control.FLTPCoef_Q14[k*LTP_ORDER+tap] != int16(int32(want[row][tap])<<7) {
						t.Fatal("typed LTP vector decode", i, row, k, tap)
					}
				}
			}
		}
	}
}

func TestLTPBitTablePointers(t *testing.T) {
	original := Opus_silk_LTP_gain_BITS_Q5_ptrs
	defer func() { Opus_silk_LTP_gain_BITS_Q5_ptrs = original }()
	for i := range original {
		N := 8 << i
		want := append([]byte(nil), unsafe.Slice(original[i], N)...)
		clone := append([]byte(nil), want...)
		Opus_silk_LTP_gain_BITS_Q5_ptrs[i] = &clone[0]
		clone = nil
		entropyInitGrowStack(12)
		runtime.GC()
		got := unsafe.Slice(Opus_silk_LTP_gain_BITS_Q5_ptrs[i], N)
		for j := range want {
			if got[j] != want[j] {
				t.Fatal("typed LTP bit table owner", i, j)
			}
		}
	}
}

func TestDecodeParametersPointers(t *testing.T) {
	for _, fs := range []int32{8, 12, 16} {
		for _, sub := range []int32{2, 4} {
			var st OpusT_silk_decoder_state
			st.Fnb_subfr = sub
			Opus_silk_decoder_set_fs(nil, &st, fs, 16000)
			st.Findices.FsignalType = TYPE_UNVOICED
			st.Findices.FNLSFInterpCoef_Q2 = 4
			st.Findices.FGainsIndices = [4]int8{9, 3, 5, 7}
			st.FLastGainIndex = 12
			guarded := struct {
				Before  int32
				Control OpusT_silk_decoder_control
				After   int32
			}{Before: 77, After: 88}
			entropyInitGrowStack(12)
			runtime.GC()
			Opus_silk_decode_parameters(nil, &st, &guarded.Control, CODE_INDEPENDENTLY)
			if guarded.Before != 77 || guarded.After != 88 || st.Findices.FNLSFInterpCoef_Q2 != 4 || guarded.Control.FGains_Q16[0] == 0 {
				t.Fatal("state/guards")
			}
		}
	}
}

func TestDecodeParametersFieldAccesses(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	var decoder OpusT_silk_decoder_state
	decoder.Fnb_subfr = MAX_NB_SUBFR
	if got := Opus_silk_decoder_set_fs(tls, &decoder, 8, 8000); got != OPUS_OK {
		t.Fatalf("set decoder sample rate: got %d", got)
	}
	decoder.Findices.FsignalType = TYPE_UNVOICED
	decoder.Findices.FNLSFInterpCoef_Q2 = 4
	decoder.Findices.FGainsIndices = [4]OpusT_opus_int8{9, 3, 5, 7}
	decoder.Findices.FNLSFIndices[0] = 0
	decoder.FLastGainIndex = 12
	var control OpusT_silk_decoder_control

	Opus_silk_decode_parameters(tls, &decoder, &control, CODE_INDEPENDENTLY)
	if got, want := control.FGains_Q16, [4]OpusT_opus_int32{335872, 286720, 335872, 540672}; got != want {
		t.Fatalf("gains: got %v, want %v", got, want)
	}
	if got, want := decoder.FprevNLSF_Q15[:10], []int16{1536, 4480, 7680, 10624, 13824, 16896, 20096, 23040, 26368, 29184}; !equalInt16s(got, want) {
		t.Fatalf("NLSFs: got %v, want %v", got, want)
	}
	if got, want := control.FPredCoef_Q12[0][:10], []int16{2681, 30, 471, -120, 288, -82, 268, -184, 234, -164}; !equalInt16s(got, want) {
		t.Fatalf("first LPC coefficients: got %v, want %v", got, want)
	}
}

func TestDecodeParametersLocalNLSFs(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	/* unvoiced with NLSF interpolation: exercises both the pNLSF_Q15 and pNLSF0_Q15 locals */
	var decoder OpusT_silk_decoder_state
	decoder.Fnb_subfr = MAX_NB_SUBFR
	if got := Opus_silk_decoder_set_fs(tls, &decoder, 8, 8000); got != OPUS_OK {
		t.Fatalf("set decoder sample rate: got %d", got)
	}
	decoder.Findices.FsignalType = TYPE_UNVOICED
	decoder.Findices.FNLSFInterpCoef_Q2 = 2
	decoder.Ffirst_frame_after_reset = 0
	decoder.Findices.FGainsIndices = [4]OpusT_opus_int8{9, 3, 5, 7}
	decoder.Findices.FNLSFIndices[0] = 0
	decoder.FLastGainIndex = 12
	decoder.FprevNLSF_Q15 = [16]OpusT_opus_int16{1000, 3000, 5000, 7000, 9000, 11000, 13000, 15000, 17000, 19000}
	var control OpusT_silk_decoder_control

	Opus_silk_decode_parameters(tls, &decoder, &control, CODE_INDEPENDENTLY)

	if got, want := control.FPredCoef_Q12[0][:10], []int16{10969, -16800, 20578, -21315, 19297, -15196, 10325, -5752, 2412, -566}; !equalInt16s(got, want) {
		t.Fatalf("interpolated LPC coefficients: got %v, want %v", got, want)
	}
	if got, want := control.FPredCoef_Q12[1][:10], []int16{2681, 30, 471, -120, 288, -82, 268, -184, 234, -164}; !equalInt16s(got, want) {
		t.Fatalf("final LPC coefficients: got %v, want %v", got, want)
	}
	if got, want := decoder.FprevNLSF_Q15[:10], []int16{1536, 4480, 7680, 10624, 13824, 16896, 20096, 23040, 26368, 29184}; !equalInt16s(got, want) {
		t.Fatalf("previous NLSFs: got %v, want %v", got, want)
	}

	/* voiced after a packet loss: exercises the LTP codebook, pitch decode and bwexpander */
	var decoder2 OpusT_silk_decoder_state
	decoder2.Fnb_subfr = MAX_NB_SUBFR
	if got := Opus_silk_decoder_set_fs(tls, &decoder2, 8, 8000); got != OPUS_OK {
		t.Fatalf("set decoder sample rate: got %d", got)
	}
	decoder2.Findices.FsignalType = TYPE_VOICED
	decoder2.Findices.FNLSFInterpCoef_Q2 = 4
	decoder2.Findices.FGainsIndices = [4]OpusT_opus_int8{9, 3, 5, 7}
	decoder2.Findices.FNLSFIndices[0] = 0
	decoder2.FLastGainIndex = 12
	decoder2.FlossCnt = 1
	decoder2.Findices.FlagIndex = 100
	decoder2.Findices.FcontourIndex = 3
	decoder2.Findices.FPERIndex = 1
	decoder2.Findices.FLTPIndex = [4]OpusT_opus_int8{2, 5, 1, 3}
	decoder2.Findices.FLTP_scaleIndex = 2
	var control2 OpusT_silk_decoder_control

	Opus_silk_decode_parameters(tls, &decoder2, &control2, CODE_INDEPENDENTLY)

	if got, want := control2.FpitchL, [4]int32{115, 116, 116, 117}; got != want {
		t.Fatalf("pitch lags: got %v, want %v", got, want)
	}
	if got, want := control2.FLTPCoef_Q14[:], []int16{-896, 1280, 7040, 5504, 2176, -1536, 7040, 9728, -1536, 1024, -128, 4608, 8192, 3456, -768, 128, 128, 1024, 128, 128}; !equalInt16s(got, want) {
		t.Fatalf("LTP coefficients: got %v, want %v", got, want)
	}
	if got, want := control2.FLTP_scale_Q14, int32(8192); got != want {
		t.Fatalf("LTP scale: got %d, want %d", got, want)
	}
	if got, want := control2.FPredCoef_Q12[0][:10], []int16{2601, 28, 430, -106, 247, -68, 217, -144, 178, -121}; !equalInt16s(got, want) {
		t.Fatalf("loss-filtered LPC coefficients: got %v, want %v", got, want)
	}
}
