package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestMiniFFTAllocUsesFields(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	config := Opus_mini_kiss_fft_alloc(tls, 8, 1, 0, 0)
	if config == 0 {
		t.Fatal("mini FFT allocation returned nil")
	}
	state := (*mini_kiss_fft_state)(unsafe.Pointer(config))
	secondTwiddle := *(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(uintptr(unsafe.Pointer(&state.Ftwiddles[0])) + 8))
	if state.Fnfft != 8 || state.Finverse != 1 {
		t.Fatalf("FFT state dimensions: nfft=%d inverse=%d, want 8 and 1", state.Fnfft, state.Finverse)
	}
	wantFactors := [4]int32{4, 2, 2, 1}
	for i, want := range wantFactors {
		if got := state.Ffactors[i]; got != want {
			t.Fatalf("FFT factor[%d]: got %d, want %d", i, got, want)
		}
	}
	if state.Ftwiddles[0] != (OpusT_mini_kiss_fft_cpx{Fr: 1}) {
		t.Fatalf("first twiddle: got %+v, want {Fr:1 Fi:0}", state.Ftwiddles[0])
	}
	if secondTwiddle != (OpusT_mini_kiss_fft_cpx{Fr: 0.70710677, Fi: 0.70710677}) {
		t.Fatalf("second twiddle: got %+v, want {Fr:0.70710677 Fi:0.70710677}", secondTwiddle)
	}
}

func TestMiniFFTRAllocUsesLocalSubsize(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	config := Opus_mini_kiss_fftr_alloc(tls, 8, 0, 0, 0)
	if config == 0 {
		t.Fatal("mini real FFT allocation returned nil")
	}
	state := (*OpusT_mini_kiss_fftr_state)(unsafe.Pointer(config))
	substate := (*OpusT_mini_kiss_fft_state)(unsafe.Pointer(state.Fsubstate))
	if got, want := substate.Fnfft, int32(4); got != want {
		t.Fatalf("substate size: got %d, want %d", got, want)
	}
	if got, want := state.Ftmpbuf-state.Fsubstate, uintptr(296); got != want {
		t.Fatalf("temporary buffer offset: got %d, want %d", got, want)
	}
	if got, want := state.Fsuper_twiddles-state.Ftmpbuf, uintptr(32); got != want {
		t.Fatalf("super twiddle offset: got %d, want %d", got, want)
	}
	if got, want := *(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(state.Fsuper_twiddles)), (OpusT_mini_kiss_fft_cpx{Fr: -0.70710677, Fi: -0.70710677}); got != want {
		t.Fatalf("first super twiddle: got %+v, want %+v", got, want)
	}
}

func TestFFTImplUsesFactorsField(t *testing.T) {
	state := OpusT_kiss_fft_state{
		Fnfft:    4,
		Fshift:   1,
		Ffactors: [16]OpusT_opus_int16{4, 1},
	}
	values := []OpusT_kiss_fft_cpx{
		{Fr: 1, Fi: 2},
		{Fr: -3, Fi: 4},
		{Fr: 5, Fi: -6},
		{Fr: 7, Fi: 8},
	}

	Opus_opus_fft_impl(nil, uintptr(unsafe.Pointer(&state)), uintptr(unsafe.Pointer(&values[0])))

	want := [4]OpusT_kiss_fft_cpx{
		{Fr: 10, Fi: 8},
		{Fr: -8, Fi: 18},
		{Fr: 2, Fi: -16},
		{Fr: 0, Fi: -2},
	}
	for i, expected := range want {
		if got := values[i]; got != expected {
			t.Fatalf("FFT output[%d]: got %+v, want %+v", i, got, expected)
		}
	}

}

func TestRadix3ButterflyUsesTwiddlesField(t *testing.T) {
	state := struct {
		OpusT_mini_kiss_fft_state
		twiddles [2]OpusT_mini_kiss_fft_cpx
	}{
		OpusT_mini_kiss_fft_state: OpusT_mini_kiss_fft_state{
			Ftwiddles: [1]OpusT_mini_kiss_fft_cpx{{Fr: 1}},
		},
		twiddles: [2]OpusT_mini_kiss_fft_cpx{
			{Fr: -0.5, Fi: -0.8660254},
		},
	}
	values := []OpusT_mini_kiss_fft_cpx{
		{Fr: 3, Fi: -2},
		{Fr: -5, Fi: 7},
		{Fr: 11, Fi: 13},
	}

	kf_bfly31(nil, uintptr(unsafe.Pointer(&values[0])), 1, uintptr(unsafe.Pointer(&state.OpusT_mini_kiss_fft_state)), 1)

	want := [3]OpusT_mini_kiss_fft_cpx{
		{Fr: 9, Fi: 18},
		{Fr: -5.196152, Fi: 1.8564062},
		{Fr: 5.196152, Fi: -25.856407},
	}
	for i, expected := range want {
		if got := values[i]; got != expected {
			t.Fatalf("radix-3 output[%d]: got %+v, want %+v", i, got, expected)
		}
	}
}
