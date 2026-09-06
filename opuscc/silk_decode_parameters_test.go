package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestDecodeParametersFieldAccesses(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	var decoder OpusT_silk_decoder_state
	decoder.Fnb_subfr = MAX_NB_SUBFR
	if got := Opus_silk_decoder_set_fs(tls, uintptr(unsafe.Pointer(&decoder)), 8, 8000); got != OPUS_OK {
		t.Fatalf("set decoder sample rate: got %d", got)
	}
	decoder.Findices.FsignalType = TYPE_UNVOICED
	decoder.Findices.FNLSFInterpCoef_Q2 = 4
	decoder.Findices.FGainsIndices = [4]OpusT_opus_int8{9, 3, 5, 7}
	decoder.Findices.FNLSFIndices[0] = 0
	decoder.FLastGainIndex = 12
	var control OpusT_silk_decoder_control

	Opus_silk_decode_parameters(tls, uintptr(unsafe.Pointer(&decoder)), uintptr(unsafe.Pointer(&control)), CODE_INDEPENDENTLY)
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
	if got := Opus_silk_decoder_set_fs(tls, uintptr(unsafe.Pointer(&decoder)), 8, 8000); got != OPUS_OK {
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

	Opus_silk_decode_parameters(tls, uintptr(unsafe.Pointer(&decoder)), uintptr(unsafe.Pointer(&control)), CODE_INDEPENDENTLY)

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
	if got := Opus_silk_decoder_set_fs(tls, uintptr(unsafe.Pointer(&decoder2)), 8, 8000); got != OPUS_OK {
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

	Opus_silk_decode_parameters(tls, uintptr(unsafe.Pointer(&decoder2)), uintptr(unsafe.Pointer(&control2)), CODE_INDEPENDENTLY)

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
