package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestLaplaceP0LocalArrays(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	buffer := make([]byte, 32)
	var encoder OpusT_ec_enc
	Opus_ec_enc_init(tls, &encoder, &buffer[0], uint32(len(buffer)))
	values := []int32{0, 3, -5, 11}
	for _, value := range values {
		Opus_ec_laplace_encode_p0(tls, uintptr(unsafe.Pointer(&encoder)), value, 16000, 12000)
	}
	Opus_ec_enc_done(tls, &encoder)

	var decoder OpusT_ec_dec
	Opus_ec_dec_init(tls, &decoder, &buffer[0], uint32(len(buffer)))
	for i, want := range values {
		if got := Opus_ec_laplace_decode_p0(tls, &decoder, 16000, 12000); got != want {
			t.Fatalf("value[%d]: got %d, want %d", i, got, want)
		}
	}
}
