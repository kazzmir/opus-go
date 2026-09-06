package opuscc

import (
	"testing"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestOpusDecoderGetSizeLocalSilkSize(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	mono := Opus_opus_decoder_get_size(tls, 1)
	stereo := Opus_opus_decoder_get_size(tls, 2)
	if mono <= 0 || stereo <= mono {
		t.Fatalf("decoder sizes: mono=%d stereo=%d", mono, stereo)
	}
	if got := Opus_opus_decoder_get_size(tls, 0); got != 0 {
		t.Fatalf("zero-channel size: got %d, want 0", got)
	}
	if got := Opus_opus_decoder_get_size(tls, 3); got != 0 {
		t.Fatalf("three-channel size: got %d, want 0", got)
	}
}
