package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestOpusDecodeNativeLossFrame(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)

	memory := libc.Xmalloc(tls, uint64(Opus_opus_decoder_get_size(tls, 1)))
	if got := Opus_opus_decoder_init(tls, memory, 24000, 1); got != OPUS_OK {
		t.Fatalf("decoder initialization: got %d", got)
	}
	decoder := (*OpusT_OpusDecoder)(unsafe.Pointer(memory))
	output := make([]float32, 120)
	if got, want := Opus_opus_decode_native(tls, memory, 0, 0, uintptr(unsafe.Pointer(&output[0])), 120, 0, 0, 0, 0, 0, 0), int32(120); got != want {
		t.Fatalf("decoded samples: got %d, want %d", got, want)
	}
	if got, want := decoder.Flast_packet_duration, int32(120); got != want {
		t.Fatalf("last packet duration: got %d, want %d", got, want)
	}
	if got, want := decoder.Fframe_size, int32(60); got != want {
		t.Fatalf("frame size: got %d, want %d", got, want)
	}
}
