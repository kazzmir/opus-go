package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestSignCodingLocalICDF(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	pulses := []OpusT_opus_int8{2, -1, 0, 3, -2, 1, 0, -1, 2, 0, -3, 1, -1, 2, 0, -2, 1, 0, -2, 3, -1, 2, 0, -1, 2, -3, 1, 0, -2, 1, -1, 2}
	sumPulses := []int32{21, 22}
	buffer := make([]byte, 32)
	var encoder OpusT_ec_enc
	Opus_ec_enc_init(tls, uintptr(unsafe.Pointer(&encoder)), uintptr(unsafe.Pointer(&buffer[0])), uint32(len(buffer)))
	Opus_silk_encode_signs(tls, uintptr(unsafe.Pointer(&encoder)), uintptr(unsafe.Pointer(&pulses[0])), 32, 1, 1, uintptr(unsafe.Pointer(&sumPulses[0])))
	Opus_ec_enc_done(tls, uintptr(unsafe.Pointer(&encoder)))

	decoded := make([]OpusT_opus_int16, len(pulses))
	for i, pulse := range pulses {
		if pulse < 0 {
			decoded[i] = OpusT_opus_int16(-pulse)
		} else {
			decoded[i] = OpusT_opus_int16(pulse)
		}
	}
	var decoder OpusT_ec_dec
	Opus_ec_dec_init(tls, uintptr(unsafe.Pointer(&decoder)), uintptr(unsafe.Pointer(&buffer[0])), uint32(len(buffer)))
	Opus_silk_decode_signs(tls, uintptr(unsafe.Pointer(&decoder)), uintptr(unsafe.Pointer(&decoded[0])), 32, 1, 1, uintptr(unsafe.Pointer(&sumPulses[0])))

	for i, want := range pulses {
		if got := decoded[i]; got != OpusT_opus_int16(want) {
			t.Fatalf("pulse[%d]: got %d, want %d", i, got, want)
		}
	}
}
