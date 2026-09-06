package opuscc

import (
	"math"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestPitchDownsampleLocalArrays(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)

	left := []OpusT_celt_sig{0.12, -0.35, 0.56, -0.21, 0.73, -0.44, 0.18, -0.67, 0.39, -0.15, 0.82, -0.51, 0.27, -0.76, 0.48, -0.09}
	right := []OpusT_celt_sig{-0.28, 0.41, -0.63, 0.24, -0.57, 0.36, -0.19, 0.69, -0.47, 0.13, -0.74, 0.52, -0.31, 0.64, -0.42, 0.08}
	input := [2]uintptr{uintptr(unsafe.Pointer(&left[0])), uintptr(unsafe.Pointer(&right[0]))}
	output := make([]OpusT_opus_val16, 8)

	Opus_pitch_downsample(tls, uintptr(unsafe.Pointer(&input[0])), uintptr(unsafe.Pointer(&output[0])), int32(len(output)), 2, 2, 0)

	want := []OpusT_opus_val16{-0.065, -0.0821798, 0.0133065805, 0.02715645, -0.025118109, -0.0051225154, -0.035841793, -0.036521006}
	for i, value := range output {
		if value != want[i] {
			t.Fatalf("output[%d]: got %.8f, want %.8f", i, value, want[i])
		}
	}
}

func TestPitchSearchLocalBestPitch(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)

	x := []OpusT_opus_val16{0.31, -0.48, 0.72, -0.19, 0.55, -0.63, 0.28, 0.81, -0.36, 0.14, -0.67, 0.43, -0.24, 0.59, -0.71, 0.38}
	y := []OpusT_opus_val16{-0.22, 0.47, -0.61, 0.33, 0.31, -0.48, 0.72, -0.19, 0.55, -0.63, 0.28, 0.81, -0.36, 0.14, -0.67, 0.43, -0.24, 0.59, -0.71, 0.38, 0.16, -0.52, 0.44, -0.29}
	var pitch int32

	Opus_pitch_search(tls, uintptr(unsafe.Pointer(&x[0])), uintptr(unsafe.Pointer(&y[0])), int32(len(x)), 8, uintptr(unsafe.Pointer(&pitch)), 0)

	if got, want := pitch, int32(1); got != want {
		t.Fatalf("pitch: got %d, want %d", got, want)
	}
}

func TestRemoveDoublingLocalCorrelations(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)

	samples := []OpusT_opus_val16{0.18, -0.42, 0.67, -0.23, 0.51, -0.76, 0.34, 0.82, -0.39, 0.12, -0.58, 0.45, -0.16, 0.71, -0.64, 0.29, 0.18, -0.42, 0.67, -0.23, 0.51, -0.76, 0.34, 0.82, -0.39, 0.12, -0.58, 0.45, -0.16, 0.71, -0.64, 0.29}
	period := int32(10)

	gain := Opus_remove_doubling(tls, uintptr(unsafe.Pointer(&samples[0])), 16, 4, 16, uintptr(unsafe.Pointer(&period)), 9, 0.46, 0)

	if got, want := period, int32(9); got != want {
		t.Fatalf("period: got %d, want %d", got, want)
	}
	if got, want := gain, OpusT_opus_val16(0.02879144); math.Abs(float64(got-want)) > 1e-7 {
		t.Fatalf("gain: got %.8f, want %.8f", got, want)
	}
}

func TestCeltFIRLocalSums(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)

	input := []OpusT_opus_val16{0.16, -0.27, 0.38, -0.49, 0.61, -0.72, 0.83, -0.94, 0.25, -0.36, 0.47, -0.58}
	coefficients := []OpusT_opus_val16{0.11, -0.23, 0.37, -0.41}
	output := make([]OpusT_opus_val16, 8)

	Opus_celt_fir_c(tls, uintptr(unsafe.Pointer(&input[4])), uintptr(unsafe.Pointer(&coefficients[0])), uintptr(unsafe.Pointer(&output[0])), int32(len(output)), 4, 0)

	want := []OpusT_opus_val16{0.30320004, -0.28890002, 0.2734, -0.25649995, -0.5608, 0.48600003, -0.31520003, 0.032400023}
	for i, value := range output {
		if value != want[i] {
			t.Fatalf("output[%d]: got %.8f, want %.8f", i, value, want[i])
		}
	}
}
