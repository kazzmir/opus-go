package opuscc

import (
	"testing"
	"unsafe"

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

	state := make([]byte, stereo)
	if got := Opus_opus_decoder_init(tls, uintptr(unsafe.Pointer(&state[0])), 48000, 2); got != OPUS_OK {
		t.Fatalf("decoder initialization: got %d, want %d", got, OPUS_OK)
	}
	decoder := (*OpusT_OpusDecoder)(unsafe.Pointer(&state[0]))
	if decoder.Fsilk_dec_offset <= 0 || decoder.Fcelt_dec_offset <= decoder.Fsilk_dec_offset || decoder.Fframe_size != 120 {
		t.Fatalf("decoder layout: silk=%d celt=%d frame=%d", decoder.Fsilk_dec_offset, decoder.Fcelt_dec_offset, decoder.Fframe_size)
	}
}
