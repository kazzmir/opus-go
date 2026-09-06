package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestNLSFDelayedDecisionQuantizerLocalStates(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	indices := make([]OpusT_opus_int8, 10)
	xQ10 := []OpusT_opus_int16{218, -147, 364, -291, 108, -75, 239, -186, 51, -32}
	wQ5 := []OpusT_opus_int16{31, 28, 35, 30, 33, 29, 36, 32, 27, 34}
	predictors := []OpusT_opus_uint8{0, 14, 242, 19, 237, 11, 248, 17, 241, 9}
	ecIndices := make([]OpusT_opus_int16, len(indices))
	ecRates := []OpusT_opus_uint8{12, 17, 21, 25, 29, 34, 39, 44, 50, 56, 63, 71}

	if got, want := Opus_silk_NLSF_del_dec_quant(tls, uintptr(unsafe.Pointer(&indices[0])), uintptr(unsafe.Pointer(&xQ10[0])), uintptr(unsafe.Pointer(&wQ5[0])), uintptr(unsafe.Pointer(&predictors[0])), uintptr(unsafe.Pointer(&ecIndices[0])), uintptr(unsafe.Pointer(&ecRates[0])), 65536, 64, 8192, 10), OpusT_opus_int32(15284434); got != want {
		t.Fatalf("rate-distortion: got %d, want %d", got, want)
	}
	wantIndices := []OpusT_opus_int8{0, -1, 0, -1, 0, -1, 0, -1, 0, -1}
	for i, want := range wantIndices {
		if got := indices[i]; got != want {
			t.Fatalf("index[%d]: got %d, want %d", i, got, want)
		}
	}
}
