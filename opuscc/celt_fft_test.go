package opuscc

import (
	"testing"
	"unsafe"
)

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
