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
