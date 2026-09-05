package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestDownFIRResamplerFieldAccesses(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	pseudostack := libc.Xmalloc(tls, 16)
	scratch := libc.Xmalloc(tls, GLOBAL_STACK_SIZE)
	*(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(pseudostack)) = OpusT_opus_ccgo_pseudostack_state{
		Fscratch_ptr:  scratch,
		Fglobal_stack: scratch,
	}
	libc.Xpthread_setspecific(tls, 0x6f707573, pseudostack)

	coefs := [14]OpusT_opus_int16{900, -700, 2800, -2300, 1900, -1500, 1100, -900, 700, -500, 300, -200, 100, -50}
	state := OpusT_silk_resampler_state_struct{
		FsIIR: [6]OpusT_opus_int32{240, -350},
		FsFIR: struct {
			Fi16 [0][36]OpusT_opus_int16
			Fi32 [36]OpusT_opus_int32
		}{Fi32: [36]OpusT_opus_int32{31, -47, 59, -71, 83, -97, 109, -127}},
		FbatchSize:    4,
		FinvRatio_Q16: 65536,
		FFIR_Order:    RESAMPLER_DOWN_ORDER_FIR1,
		FFIR_Fracs:    1,
		FCoefs:        uintptr(unsafe.Pointer(&coefs[0])),
	}
	input := []int16{1300, -2400, 3600, -4700}
	output := make([]int16, len(input))

	Opus_silk_resampler_private_down_FIR(tls, uintptr(unsafe.Pointer(&state)), uintptr(unsafe.Pointer(&output[0])), uintptr(unsafe.Pointer(&input[0])), int32(len(input)))

	wantOutput := []int16{0, 222, -581, 1062}
	for i, want := range wantOutput {
		if got := output[i]; got != want {
			t.Fatalf("output[%d]: got %d, want %d", i, got, want)
		}
	}
	if got, want := state.FsIIR, [6]OpusT_opus_int32{-99423, 48264}; got != want {
		t.Fatalf("IIR state: got %v, want %v", got, want)
	}
	wantFIR := [24]OpusT_opus_int32{83, -97, 109, -127, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 333040, -596456, 874605, -1129674}
	for i, want := range wantFIR {
		if got := state.FsFIR.Fi32[i]; got != want {
			t.Fatalf("FIR state[%d]: got %d, want %d", i, got, want)
		}
	}
}
