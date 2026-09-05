package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestMultistreamDecoderInitFieldAccesses(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	const streams, coupledStreams, channels = 2, 1, 3
	size := Opus_opus_multistream_decoder_get_size(tls, streams, coupledStreams)
	decoder := libc.Xmalloc(tls, uint64(size))
	mapping := []uint8{0, 1, 2}

	if got, want := Opus_opus_multistream_decoder_init(tls, decoder, 48000, channels, streams, coupledStreams, uintptr(unsafe.Pointer(&mapping[0]))), int32(OPUS_OK); got != want {
		t.Fatalf("initialization result: got %d, want %d", got, want)
	}

	layout := (*OpusT_OpusMSDecoder)(unsafe.Pointer(decoder)).Flayout
	if got, want := layout.Fnb_channels, int32(channels); got != want {
		t.Fatalf("channels: got %d, want %d", got, want)
	}
	if got, want := layout.Fnb_streams, int32(streams); got != want {
		t.Fatalf("streams: got %d, want %d", got, want)
	}
	if got, want := layout.Fnb_coupled_streams, int32(coupledStreams); got != want {
		t.Fatalf("coupled streams: got %d, want %d", got, want)
	}
	if got, want := layout.Fmapping[:channels], mapping; string(got) != string(want) {
		t.Fatalf("mapping: got %v, want %v", got, want)
	}
}
