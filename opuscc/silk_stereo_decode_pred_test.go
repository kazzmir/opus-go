package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestStereoDecodePredFieldAccesses(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	data := []byte{0x93, 0x57, 0xc1, 0x2a, 0xee, 0x44, 0x18, 0xb7}
	var rangeDecoder OpusT_ec_dec
	Opus_ec_dec_init(tls, uintptr(unsafe.Pointer(&rangeDecoder)), uintptr(unsafe.Pointer(&data[0])), uint32(len(data)))
	pred := [2]OpusT_opus_int32{}

	Opus_silk_stereo_decode_pred(tls, uintptr(unsafe.Pointer(&rangeDecoder)), uintptr(unsafe.Pointer(&pred[0])))

	want := [2]OpusT_opus_int32{377, 656}
	if pred != want {
		t.Fatalf("predictors: got %v, want %v", pred, want)
	}
	if got, want := rangeDecoder.Foffs, uint32(5); got != want {
		t.Fatalf("range decoder offset: got %d, want %d", got, want)
	}
}
