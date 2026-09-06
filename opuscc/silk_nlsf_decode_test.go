package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestNLSFDecodeLocalArrays(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	indices := make([]OpusT_opus_int8, 17)
	var decoded [16]OpusT_opus_int16
	Opus_silk_NLSF_decode(tls, uintptr(unsafe.Pointer(&decoded[0])), uintptr(unsafe.Pointer(&indices[0])), uintptr(unsafe.Pointer(&Opus_silk_NLSF_CB_WB)))

	want := [16]OpusT_opus_int16{896, 2944, 4864, 6912, 8832, 10880, 12800, 14848, 16768, 18816, 20736, 22784, 24704, 26624, 28544, 30592}
	if decoded != want {
		t.Fatalf("decoded NLSFs: got %v, want %v", decoded, want)
	}
}
