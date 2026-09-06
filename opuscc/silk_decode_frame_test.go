package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestDecodeFrameFieldAccesses(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)
	var decoder OpusT_silk_decoder_state
	decoder.Fnb_subfr = MAX_NB_SUBFR
	if got := Opus_silk_decoder_set_fs(tls, uintptr(unsafe.Pointer(&decoder)), 8, 8000); got != OPUS_OK {
		t.Fatalf("set fs: got %d", got)
	}
	Opus_silk_PLC_Reset(tls, uintptr(unsafe.Pointer(&decoder)))
	decoder.FsPLC.FprevGain_Q16 = [2]OpusT_opus_int32{65536, 65536}
	decoder.FsPLC.FpitchL_Q8 = 20 << 8
	decoder.FsPLC.FLTPCoef_Q14 = [5]OpusT_opus_int16{300, -150, 1200, -100, 75}
	decoder.FsPLC.FprevLPC_Q12 = [16]OpusT_opus_int16{120, -80, 60, -45, 30, -20, 15, -10, 8, -5}
	decoder.FsPLC.FprevLTP_scale_Q14 = 13000
	decoder.FsPLC.FrandScale_Q14 = 11000
	decoder.FsPLC.Frand_seed = 12345
	for i := range decoder.Fexc_Q14 {
		decoder.Fexc_Q14[i] = int32((i*71)%3000 - 1500)
	}
	for i := range decoder.FoutBuf {
		decoder.FoutBuf[i] = int16((i*37)%1000 - 500)
	}
	output := make([]int16, decoder.Fframe_length)
	var samples int32
	if got := Opus_silk_decode_frame(tls, uintptr(unsafe.Pointer(&decoder)), 0, uintptr(unsafe.Pointer(&output[0])), uintptr(unsafe.Pointer(&samples)), 1, CODE_INDEPENDENTLY, 0); got != OPUS_OK {
		t.Fatalf("decode result: got %d", got)
	}
	if got, want := samples, int32(160); got != want {
		t.Fatalf("sample count: got %d, want %d", got, want)
	}
	if got, want := output[:8], []int16{20, 32, -38, -29, -31, -28, -25, -22}; !equalInt16s(got, want) {
		t.Fatalf("output prefix: got %v, want %v", got, want)
	}
	if got, want := decoder.FoutBuf[decoder.Fltp_mem_length-8:decoder.Fltp_mem_length], []int16{0, 1, 1, 1, 1, 1, 2, 2}; !equalInt16s(got, want) {
		t.Fatalf("output history suffix: got %v, want %v", got, want)
	}
}
