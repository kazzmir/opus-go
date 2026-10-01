package opuscc

import (
	"runtime"
	"testing"
	"unsafe"
)

func TestQuantBandN1Pointers(t *testing.T) {
	buffer := [16]byte{}
	var encoder OpusT_ec_enc
	Opus_ec_enc_init(nil, &encoder, &buffer[0], 16)
	ctx := band_ctx{Fencode: 1, Fresynth: 1, Fremaining_bits: 8}
	x := float32(-.5)
	low := float32(77)
	entropyInitGrowStack(12)
	runtime.GC()
	if quant_band_n1(nil, &ctx, &encoder, &x, &x, &low) != 1 || x != 1 || low != 1 || ctx.Fremaining_bits != 0 {
		t.Fatal("alias/budget/store order", x, low, ctx.Fremaining_bits)
	}
	ctx = band_ctx{Fencode: 1, Fremaining_bits: 0}
	if quant_band_n1(nil, &ctx, nil, nil, nil, nil) != 1 {
		t.Fatal("unused nil inputs")
	}
}

func TestQuantBandN1FieldAccesses(t *testing.T) {
	buffer := make([]byte, 16)
	var encoder OpusT_ec_enc
	Opus_ec_enc_init(nil, &encoder, unsafe.SliceData(buffer), uint32(len(buffer)))

	context := band_ctx{
		Fencode:         1,
		Fresynth:        1,
		Fec:             uintptr(unsafe.Pointer(&encoder)),
		Fremaining_bits: 24,
	}
	x := OpusT_celt_norm(-0.375)
	y := OpusT_celt_norm(0.625)
	lowband := OpusT_celt_norm(0)

	if got, want := quant_band_n1(nil, &context, &encoder, &x, &y, &lowband), uint32(1); got != want {
		t.Fatalf("coded dimensions: got %d, want %d", got, want)
	}
	Opus_ec_enc_done(nil, &encoder)

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
