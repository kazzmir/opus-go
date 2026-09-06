package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestDecodeIndicesFieldAccesses(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	var decoder OpusT_silk_decoder_state
	decoder.Fnb_subfr = MAX_NB_SUBFR
	if got := Opus_silk_decoder_set_fs(tls, uintptr(unsafe.Pointer(&decoder)), 8, 8000); got != OPUS_OK {
		t.Fatalf("set decoder sample rate: got %d", got)
	}
	data := []byte{0x93, 0x57, 0xc1, 0x2a, 0xee, 0x44, 0x18, 0xb7, 0x6d, 0x09, 0xfa, 0x35, 0x81, 0x62, 0xdc, 0x4e}
	var rangeDecoder OpusT_ec_dec
	Opus_ec_dec_init(tls, uintptr(unsafe.Pointer(&rangeDecoder)), uintptr(unsafe.Pointer(&data[0])), uint32(len(data)))
	Opus_silk_decode_indices(tls, uintptr(unsafe.Pointer(&decoder)), uintptr(unsafe.Pointer(&rangeDecoder)), 0, 0, CODE_INDEPENDENTLY)
	if got, want := decoder.Findices.FsignalType, int8(0); got != want {
		t.Fatalf("signal type: got %d, want %d", got, want)
	}
	if got, want := decoder.Findices.FGainsIndices, [4]OpusT_opus_int8{15, 4, 4, 3}; got != want {
		t.Fatalf("gain indices: got %v, want %v", got, want)
	}
	if got, want := decoder.Findices.FNLSFIndices[:10], []int8{12, -1, -2, 1, -2, -2, 0, -1, 0, 0}; !equalInt8s(got, want) {
		t.Fatalf("NLSF indices: got %v, want %v", got, want)
	}
	if got, want := decoder.Findices.FSeed, int8(1); got != want {
		t.Fatalf("seed: got %d, want %d", got, want)
	}
}

func equalInt8s(got, want []int8) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}
