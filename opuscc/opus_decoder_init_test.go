package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestDecoderInitFieldAccesses(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	size := Opus_opus_decoder_get_size(tls, 2)
	decoderMemory := libc.Xmalloc(tls, uint64(size))
	if got, want := Opus_opus_decoder_init(tls, decoderMemory, 24000, 2), int32(OPUS_OK); got != want {
		t.Fatalf("initialization result: got %d, want %d", got, want)
	}

	decoder := (*OpusT_OpusDecoder)(unsafe.Pointer(decoderMemory))
	if got, want := decoder.Fchannels, int32(2); got != want {
		t.Fatalf("channels: got %d, want %d", got, want)
	}
	if got, want := decoder.Fstream_channels, int32(2); got != want {
		t.Fatalf("stream channels: got %d, want %d", got, want)
	}
	if got, want := decoder.FFs, OpusT_opus_int32(24000); got != want {
		t.Fatalf("sample rate: got %d, want %d", got, want)
	}
	if got, want := decoder.Fframe_size, int32(60); got != want {
		t.Fatalf("frame size: got %d, want %d", got, want)
	}
	if decoder.Fsilk_dec_offset <= 0 || decoder.Fcelt_dec_offset <= decoder.Fsilk_dec_offset {
		t.Fatalf("invalid component offsets: silk=%d celt=%d", decoder.Fsilk_dec_offset, decoder.Fcelt_dec_offset)
	}
	if got, want := decoder.FDecControl.FAPI_sampleRate, OpusT_opus_int32(24000); got != want {
		t.Fatalf("SILK API sample rate: got %d, want %d", got, want)
	}
	if got, want := decoder.FDecControl.FnChannelsAPI, int32(2); got != want {
		t.Fatalf("SILK API channels: got %d, want %d", got, want)
	}
}
