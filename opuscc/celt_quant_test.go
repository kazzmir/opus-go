package opuscc

import (
	"testing"
	"unsafe"
)

func TestQuantBandN1FieldAccesses(t *testing.T) {
	buffer := make([]byte, 16)
	var encoder OpusT_ec_enc
	Opus_ec_enc_init(nil, uintptr(unsafe.Pointer(&encoder)), entropyBufferPointer(buffer), uint32(len(buffer)))

	context := band_ctx{
		Fencode:         1,
		Fresynth:        1,
		Fec:             uintptr(unsafe.Pointer(&encoder)),
		Fremaining_bits: 24,
	}
	x := OpusT_celt_norm(-0.375)
	y := OpusT_celt_norm(0.625)
	lowband := OpusT_celt_norm(0)

	if got, want := quant_band_n1(nil, uintptr(unsafe.Pointer(&context)), uintptr(unsafe.Pointer(&x)), uintptr(unsafe.Pointer(&y)), uintptr(unsafe.Pointer(&lowband))), uint32(1); got != want {
		t.Fatalf("coded dimensions: got %d, want %d", got, want)
	}
	Opus_ec_enc_done(nil, uintptr(unsafe.Pointer(&encoder)))

	if got, want := x, OpusT_celt_norm(-1); got != want {
		t.Fatalf("resynthesized X: got %v, want %v", got, want)
	}
	if got, want := y, OpusT_celt_norm(1); got != want {
		t.Fatalf("resynthesized Y: got %v, want %v", got, want)
	}
	if got, want := lowband, OpusT_celt_norm(-1); got != want {
		t.Fatalf("lowband output: got %v, want %v", got, want)
	}
	if got, want := context.Fremaining_bits, OpusT_opus_int32(8); got != want {
		t.Fatalf("remaining bits: got %d, want %d", got, want)
	}
	if got, want := encoder.Fnbits_total, int32(35); got != want {
		t.Fatalf("encoded sign bit count: got %d, want %d", got, want)
	}
}
