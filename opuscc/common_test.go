package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestCompositeSizeAlignment(t *testing.T) {
	for _, n := range []int32{-2147483648, -2147483647, -9, -8, -7, -1, 0, 1, 7, 8, 9, 2147483640, 2147483647} {
		want := int32((uint32(n) + uint32(8) - 1) / 8 * 8)
		if got := opusAlignSize8(n); got != want {
			t.Fatal("unsigned size alignment", n, got, want)
		}
	}
	for _, channels := range []int32{-2147483648, -1, 0, 3, 2147483647} {
		if Opus_opus_decoder_get_size(nil, channels) != 0 {
			t.Fatal("invalid channel size", channels)
		}
	}
	for _, channels := range []int32{1, 2} {
		var silk int32
		if Opus_silk_Get_Decoder_Size(nil, &silk) != 0 {
			t.Fatal("silk size")
		}
		want := int32(104) + int32((uint32(silk)+7)/8*8) + Opus_celt_decoder_get_size(nil, channels)
		if got := Opus_opus_decoder_get_size(nil, channels); got != want {
			t.Fatal("decoder composition", channels, got, want)
		}
	}
}

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
	if got := Opus_opus_decoder_init(tls, (*OpusT_OpusDecoder)(unsafe.Pointer(&state[0])), 48000, 2); got != OPUS_OK {
		t.Fatalf("decoder initialization: got %d, want %d", got, OPUS_OK)
	}
	decoder := (*OpusT_OpusDecoder)(unsafe.Pointer(&state[0]))
	if decoder.Fsilk_dec_offset <= 0 || decoder.Fcelt_dec_offset <= decoder.Fsilk_dec_offset || decoder.Fframe_size != 120 {
		t.Fatalf("decoder layout: silk=%d celt=%d frame=%d", decoder.Fsilk_dec_offset, decoder.Fcelt_dec_offset, decoder.Fframe_size)
	}
}
