package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestQuantBandStereoOneSampleFieldAccesses(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	buffer := make([]byte, 16)
	var encoder OpusT_ec_enc
	Opus_ec_enc_init(nil, uintptr(unsafe.Pointer(&encoder)), entropyBufferPointer(buffer), uint32(len(buffer)))
	context := band_ctx{
		Fencode:         1,
		Fresynth:        1,
		Fec:             uintptr(unsafe.Pointer(&encoder)),
		Fremaining_bits: 16,
	}
	x := OpusT_celt_norm(-0.75)
	y := OpusT_celt_norm(0.5)
	lowband := OpusT_celt_norm(0)

	if got, want := quant_band_stereo(tls, uintptr(unsafe.Pointer(&context)), uintptr(unsafe.Pointer(&x)), uintptr(unsafe.Pointer(&y)), 1, 0, 1, 0, 0, uintptr(unsafe.Pointer(&lowband)), 0, 3), uint32(1); got != want {
		t.Fatalf("coded dimensions: got %d, want %d", got, want)
	}

	if got, want := x, OpusT_celt_norm(-1); got != want {
		t.Fatalf("resynthesized X: got %v, want %v", got, want)
	}
	if got, want := y, OpusT_celt_norm(1); got != want {
		t.Fatalf("resynthesized Y: got %v, want %v", got, want)
	}
	if got, want := lowband, OpusT_celt_norm(-1); got != want {
		t.Fatalf("lowband output: got %v, want %v", got, want)
	}
	if got, want := context.Fremaining_bits, OpusT_opus_int32(0); got != want {
		t.Fatalf("remaining bits: got %d, want %d", got, want)
	}
}
