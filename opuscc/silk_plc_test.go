package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestPLCResetFieldAccesses(t *testing.T) {
	dec := OpusT_silk_decoder_state{Fframe_length: 319, FsPLC: OpusT_silk_PLC_struct{FpitchL_Q8: -111, FprevGain_Q16: [2]OpusT_opus_int32{222222, -333333}, Fsubfr_length: 47, Fnb_subfr: 4, Ffs_kHz: 16}}
	Opus_silk_PLC_Reset(nil, uintptr(unsafe.Pointer(&dec)))
	if got, want := dec.FsPLC.FpitchL_Q8, int32(40832); got != want {
		t.Fatalf("pitchL_Q8: got %d, want %d", got, want)
	}
	if got, want := dec.FsPLC.FprevGain_Q16, [2]OpusT_opus_int32{65536, 65536}; got != want {
		t.Fatalf("previous gains: got %v, want %v", got, want)
	}
	if got, want := dec.FsPLC.Fsubfr_length, int32(20); got != want {
		t.Fatalf("subframe length: got %d, want %d", got, want)
	}
	if got, want := dec.FsPLC.Fnb_subfr, int32(2); got != want {
		t.Fatalf("subframes: got %d, want %d", got, want)
	}
	if got, want := dec.FsPLC.Ffs_kHz, int32(16); got != want {
		t.Fatalf("unrelated PLC state changed: got %d, want %d", got, want)
	}
}

func TestPLCConcealFieldAccesses(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)
	decoder := OpusT_silk_decoder_state{Ffs_kHz: 8, Fframe_length: 80, Fsubfr_length: 20, Fnb_subfr: 4, Fltp_mem_length: 200, FLPC_order: 10, FprevSignalType: TYPE_VOICED, FsPLC: OpusT_silk_PLC_struct{FLTPCoef_Q14: [5]OpusT_opus_int16{300, -150, 1200, -100, 75}, FprevLPC_Q12: [16]OpusT_opus_int16{120, -80, 60, -45, 30, -20, 15, -10, 8, -5}, FprevGain_Q16: [2]OpusT_opus_int32{65536, 65536}, FprevLTP_scale_Q14: 13000, FpitchL_Q8: 20 << 8, FrandScale_Q14: 11000, Frand_seed: 12345, Fsubfr_length: 20, Fnb_subfr: 4}}
	for i := range decoder.Fexc_Q14 {
		decoder.Fexc_Q14[i] = int32((i*71)%3000 - 1500)
	}
	for i := range decoder.FoutBuf {
		decoder.FoutBuf[i] = int16((i*37)%1000 - 500)
	}
	var control OpusT_silk_decoder_control
	frame := make([]int16, decoder.Fframe_length)
	silk_PLC_conceal(tls, uintptr(unsafe.Pointer(&decoder)), uintptr(unsafe.Pointer(&control)), uintptr(unsafe.Pointer(&frame[0])), 0)
	if got, want := control.FpitchL, [4]OpusT_opus_int32{21, 21, 21, 21}; got != want {
		t.Fatalf("pitch lags: got %v, want %v", got, want)
	}
	if got, want := decoder.FsPLC.Frand_seed, int32(-1769093111); got != want {
		t.Fatalf("random seed: got %d, want %d", got, want)
	}
	if got, want := frame[:8], []int16{14, 17, 20, 23, 26, 29, 32, 35}; !equalInt16s(got, want) {
		t.Fatalf("frame prefix: got %v, want %v", got, want)
	}
}

func equalInt16s(got, want []int16) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}
