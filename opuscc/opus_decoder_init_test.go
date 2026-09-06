package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestDecoderInitFieldAccesses(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	memory := libc.Xmalloc(tls, uint64(Opus_opus_decoder_get_size(tls, 2)))
	if got := Opus_opus_decoder_init(tls, memory, 24000, 2); got != OPUS_OK {
		t.Fatalf("initialization result: got %d", got)
	}
	decoder := (*OpusT_OpusDecoder)(unsafe.Pointer(memory))
	if decoder.Fchannels != 2 || decoder.Fstream_channels != 2 || decoder.FFs != 24000 || decoder.Fframe_size != 60 {
		t.Fatalf("decoder fields were not initialized: %+v", decoder)
	}
	if decoder.Fsilk_dec_offset <= 0 || decoder.Fcelt_dec_offset <= decoder.Fsilk_dec_offset {
		t.Fatalf("invalid component offsets")
	}
	if decoder.FDecControl.FAPI_sampleRate != 24000 || decoder.FDecControl.FnChannelsAPI != 2 {
		t.Fatalf("invalid SILK control state")
	}
}

func TestDecoderCtlResetFieldAccesses(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	memory := libc.Xmalloc(tls, uint64(Opus_opus_decoder_get_size(tls, 2)))
	if got := Opus_opus_decoder_init(tls, memory, 24000, 2); got != OPUS_OK {
		t.Fatalf("decoder initialization: got %d", got)
	}
	decoder := (*OpusT_OpusDecoder)(unsafe.Pointer(memory))
	decoder.Fstream_channels, decoder.Fbandwidth, decoder.Fprev_mode, decoder.Fframe_size = 1, OPUS_BANDWIDTH_FULLBAND, MODE_CELT_ONLY, 999
	if got := Opus_opus_decoder_ctl(tls, memory, OPUS_RESET_STATE, 0); got != OPUS_OK {
		t.Fatalf("reset result: got %d", got)
	}
	if decoder.Fstream_channels != 2 || decoder.Fframe_size != 60 || decoder.Fbandwidth != 0 {
		t.Fatalf("reset fields: channels=%d frame=%d bandwidth=%d", decoder.Fstream_channels, decoder.Fframe_size, decoder.Fbandwidth)
	}
}
