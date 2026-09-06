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
