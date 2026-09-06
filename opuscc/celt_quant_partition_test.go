package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestQuantPartitionRemainingBitsField(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	cacheIndex := [1]int16{0}
	cacheBits := [1]byte{0}
	mode := OpusT_OpusCustomMode{FnbEBands: 1}
	mode.Fcache.Findex = uintptr(unsafe.Pointer(&cacheIndex[0]))
	mode.Fcache.Fbits = uintptr(unsafe.Pointer(&cacheBits[0]))
	context := band_ctx{Fm: uintptr(unsafe.Pointer(&mode)), Fi: 0, Fresynth: 1, Fremaining_bits: 23, Fseed: 123456}
	x := [2]OpusT_celt_norm{0.3, -0.4}

	if got, want := quant_partition(tls, uintptr(unsafe.Pointer(&context)), uintptr(unsafe.Pointer(&x[0])), 2, 1, 1, 0, -1, 1, 1), uint32(1); got != want {
		t.Fatalf("collapse mask: got %d, want %d", got, want)
	}

	if got, want := x, [2]OpusT_celt_norm{0.37370443, 0.9275479}; got != want {
		t.Fatalf("resynthesized spectrum: got %v, want %v", got, want)
	}
	if got, want := context.Fremaining_bits, int32(23); got != want {
		t.Fatalf("remaining bits: got %d, want %d", got, want)
	}
	if got, want := context.Fseed, uint32(870155634); got != want {
		t.Fatalf("random seed: got %d, want %d", got, want)
	}
}

func TestQuantPartitionLocalSplitState(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)

	cacheIndex := [2]int16{0, 0}
	cacheBits := [2]byte{1, 0}
	logN := [1]int16{8}
	mode := OpusT_OpusCustomMode{FnbEBands: 1}
	mode.Fcache.Findex = uintptr(unsafe.Pointer(&cacheIndex[0]))
	mode.Fcache.Fbits = uintptr(unsafe.Pointer(&cacheBits[0]))
	mode.FlogN = uintptr(unsafe.Pointer(&logN[0]))
	buffer := make([]byte, 16)
	var encoder OpusT_ec_enc
	Opus_ec_enc_init(tls, uintptr(unsafe.Pointer(&encoder)), entropyBufferPointer(buffer), uint32(len(buffer)))
	context := band_ctx{
		Fm:              uintptr(unsafe.Pointer(&mode)),
		Fencode:         1,
		Fresynth:        1,
		Fec:             uintptr(unsafe.Pointer(&encoder)),
		Fremaining_bits: 80,
		Fseed:           987654,
	}
	x := [4]OpusT_celt_norm{0.2, -0.4, 0.6, -0.8}

	mask := quant_partition(tls, uintptr(unsafe.Pointer(&context)), uintptr(unsafe.Pointer(&x[0])), 4, 30, 1, 0, 0, 1, 3)
	Opus_ec_enc_done(tls, uintptr(unsafe.Pointer(&encoder)))

	if got, want := mask, uint32(1); got != want {
		t.Fatalf("collapse mask: got %d, want %d", got, want)
	}
	if got, want := x, [4]OpusT_celt_norm{0, -0.4083252, 0, -0.9128723}; got != want {
		t.Fatalf("resynthesized spectrum: got %v, want %v", got, want)
	}
	if got, want := context.Fremaining_bits, int32(78); got != want {
		t.Fatalf("remaining bits: got %d, want %d", got, want)
	}
	if got, want := context.Fseed, uint32(987654); got != want {
		t.Fatalf("random seed: got %d, want %d", got, want)
	}
	if got, want := buffer[:encoder.Foffs], []byte{0xa0}; string(got) != string(want) {
		t.Fatalf("encoded bytes: got % x, want % x", got, want)
	}
}
