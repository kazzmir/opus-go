package opuscc

import (
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
