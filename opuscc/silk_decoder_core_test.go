package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestDecodeCoreFieldAccesses(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)
	var decoder OpusT_silk_decoder_state
	decoder.Fnb_subfr = MAX_NB_SUBFR
	if got := Opus_silk_decoder_set_fs(tls, uintptr(unsafe.Pointer(&decoder)), 8, 8000); got != OPUS_OK {
		t.Fatalf("set fs: got %d", got)
	}
	decoder.Fprev_gain_Q16 = 65536
	decoder.Findices.FsignalType = TYPE_UNVOICED
	decoder.Findices.FquantOffsetType = 1
	decoder.Findices.FSeed = 17
	var control OpusT_silk_decoder_control
	control.FGains_Q16 = [4]OpusT_opus_int32{65536, 72000, 68000, 76000}
	for i := range control.FPredCoef_Q12 {
		control.FPredCoef_Q12[i] = [16]OpusT_opus_int16{120, -80, 60, -45, 30, -20, 15, -10, 8, -5}
	}
	pulses := make([]int16, decoder.Fframe_length)
	for i := range pulses {
		pulses[i] = int16((i*7)%9 - 4)
	}
	output := make([]int16, decoder.Fframe_length)
	Opus_silk_decode_core(tls, uintptr(unsafe.Pointer(&decoder)), uintptr(unsafe.Pointer(&control)), uintptr(unsafe.Pointer(&output[0])), uintptr(unsafe.Pointer(&pulses[0])), 0)
	/* expected values verified against the C reference implementation (silk/decode_core.c) */
	if got, want := output[:8], []int16{4, 3, 1, 1, -3, -4, 2, 0}; !equalInt16s(got, want) {
		t.Fatalf("output prefix: got %v, want %v", got, want)
	}
	if got, want := decoder.Fexc_Q14[:8], []int32{60416, 51712, 18944, 11264, -44032, -68096, 35328, 3840}; !equalInt32s(got, want) {
		t.Fatalf("excitation prefix: got %v, want %v", got, want)
	}
	if got, want := decoder.Fprev_gain_Q16, int32(76000); got != want {
		t.Fatalf("previous gain: got %d, want %d", got, want)
	}
}

func TestDecodeCoreLocalLPCArray(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)
	var decoder OpusT_silk_decoder_state
	decoder.Fnb_subfr = MAX_NB_SUBFR
	if got := Opus_silk_decoder_set_fs(tls, uintptr(unsafe.Pointer(&decoder)), 16, 16000); got != OPUS_OK {
		t.Fatalf("set fs: got %d", got)
	}
	decoder.Fprev_gain_Q16 = 65536
	decoder.Findices.FsignalType = TYPE_UNVOICED
	decoder.Findices.FquantOffsetType = 1
	decoder.Findices.FSeed = 17
	var control OpusT_silk_decoder_control
	control.FGains_Q16 = [4]OpusT_opus_int32{65536, 72000, 68000, 76000}
	for i := range control.FPredCoef_Q12 {
		control.FPredCoef_Q12[i] = [16]OpusT_opus_int16{120, -80, 60, -45, 30, -20, 15, -10, 8, -5, 4, -3, 2, -2, 1, -1}
	}
	pulses := make([]int16, decoder.Fframe_length)
	for i := range pulses {
		pulses[i] = int16((i*7)%9 - 4)
	}
	output := make([]int16, decoder.Fframe_length)
	Opus_silk_decode_core(tls, uintptr(unsafe.Pointer(&decoder)), uintptr(unsafe.Pointer(&control)), uintptr(unsafe.Pointer(&output[0])), uintptr(unsafe.Pointer(&pulses[0])), 0)

	/* expected values from the C reference implementation (silk/decode_core.c) */
	if got, want := output[:16], []int16{4, 3, 1, 1, -3, -4, 2, 0, 2, -4, -3, -1, -1, -3, -4, -2}; !equalInt16s(got, want) {
		t.Fatalf("output prefix: got %v, want %v", got, want)
	}
	if got, want := decoder.Fexc_Q14[:8], []int32{60416, 51712, 18944, 11264, -44032, -68096, 35328, 3840}; !equalInt32s(got, want) {
		t.Fatalf("excitation prefix: got %v, want %v", got, want)
	}
	wantSLPC := []int32{-3536, 27216, 61472, -50592, -21152, 12160, -44400, -69136, 34048, -2080, -28944, 60576, -50080, 16336, 13584, 42736}
	if got := decoder.FsLPC_Q14_buf[:]; !equalInt32s(got, wantSLPC) {
		t.Fatalf("LPC state: got %v, want %v", got, wantSLPC)
	}
	if got, want := decoder.Fprev_gain_Q16, int32(76000); got != want {
		t.Fatalf("previous gain: got %d, want %d", got, want)
	}
}

func equalInt32s(got, want []int32) bool {
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
