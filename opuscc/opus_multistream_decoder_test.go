package opuscc

import (
	"runtime"
	"slices"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestMultistreamDecoderInitPointers(t *testing.T) {
	size := int(Opus_opus_multistream_decoder_get_size(nil, 2, 1))
	backing := make([]uint64, (size+7)/8+2)
	st := (*OpusT_OpusMSDecoder)(unsafe.Pointer(&backing[0]))
	image := unsafe.Slice((*byte)(unsafe.Pointer(st)), size+16)
	for i := size; i < len(image); i++ {
		image[i] = 0xa5
	}
	mapping := [3]byte{0, 1, 2}
	entropyInitGrowStack(12)
	runtime.GC()
	if Opus_opus_multistream_decoder_init(nil, st, 48000, 3, 2, 1, &mapping[0]) != 0 || st.Flayout.Fmapping[2] != 2 {
		t.Fatal("init")
	}
	for _, b := range image[size:] {
		if b != 0xa5 {
			t.Fatal("guard")
		}
	}
	before := slices.Clone(image)
	if Opus_opus_multistream_decoder_init(nil, st, 48000, 0, 2, 1, nil) != OPUS_BAD_ARG || !slices.Equal(before, image) {
		t.Fatal("argument failure")
	}
	mapping[2] = 3
	if Opus_opus_multistream_decoder_init(nil, st, 48000, 3, 2, 1, &mapping[0]) != OPUS_BAD_ARG || st.Flayout.Fmapping[2] != 3 {
		t.Fatal("layout failure writes")
	}
}

func TestMultistreamDecoderInitFieldAccesses(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	const streams, coupledStreams, channels = 2, 1, 3
	size := Opus_opus_multistream_decoder_get_size(tls, streams, coupledStreams)
	decoder := libc.Xmalloc(tls, uint64(size))
	mapping := []uint8{0, 1, 2}

	if got, want := Opus_opus_multistream_decoder_init(tls, (*OpusT_OpusMSDecoder)(unsafe.Pointer(decoder)), 48000, channels, streams, coupledStreams, &mapping[0]), int32(OPUS_OK); got != want {
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
