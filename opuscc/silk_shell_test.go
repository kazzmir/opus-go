package opuscc

import (
	"testing"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestShellEncoderLocalPulseTrees(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	pulses := []int32{2, 0, 1, 1, 0, 2, 1, 0, 2, 0, 1, 1, 1, 0, 2, 1}
	buffer := make([]byte, 32)
	var encoder OpusT_ec_enc
	Opus_ec_enc_init(tls, &encoder, &buffer[0], uint32(len(buffer)))
	Opus_silk_shell_encoder(tls, &encoder, (*[16]int32)(pulses))
	Opus_ec_enc_done(tls, &encoder)

	var decoder OpusT_ec_dec
	Opus_ec_dec_init(tls, &decoder, &buffer[0], uint32(len(buffer)))
	var decoded [16]OpusT_opus_int16
	Opus_silk_shell_decoder(tls, &decoded, &decoder, 15)

	for i, want := range pulses {
		if got := int32(decoded[i]); got != want {
			t.Fatalf("pulse[%d]: got %d, want %d", i, got, want)
		}
	}
}
