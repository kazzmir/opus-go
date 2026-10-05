package opuscc

import (
	"runtime"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestOpusNativeTypedDecoderEntry(t *testing.T) {
	left, right := newOpusFrameOwnerDecoder(t, 1), newOpusFrameOwnerDecoder(t, 1)
	a, b := make([]float32, 122), make([]float32, 122)
	a[0], a[121], b[0], b[121] = 77, 88, 77, 88
	entropyInitGrowStack(12)
	runtime.GC()
	ra := opusDecodeNative(nil, &left.Decoder, 0, 0, uintptr(unsafe.Pointer(&a[1])), 120, 0, 0, 0, 0)
	rb := Opus_opus_decode_native(nil, uintptr(unsafe.Pointer(&right.Decoder)), 0, 0, uintptr(unsafe.Pointer(&b[1])), 120, 0, 0, 0, 0, 0, 0)
	if ra != 120 || rb != ra || left.Decoder != right.Decoder || a[0] != 77 || a[121] != 88 {
		t.Fatal("native typed decoder entry")
	}
	for i := range a {
		if a[i] != b[i] {
			t.Fatal("native decoder PCM")
		}
	}
}

func TestOpusNativePacketFramePointers(t *testing.T) {
	storage := newOpusFrameOwnerDecoder(t, 2)
	decoder := &storage.Decoder
	decoder.Fmode = MODE_SILK_ONLY
	decoder.Fbandwidth = OPUS_BANDWIDTH_NARROWBAND
	decoder.Fframe_size = 2880
	decoder.Fstream_channels = 1
	packet := mustHex(t, "18007523c11e84d40a7ed0075134da9ffc0529ef9f410157b57c1f843e40")
	pcm := make([]float32, 6002)
	for i := range pcm {
		pcm[i] = 3
	}
	pcm[0], pcm[6001] = 77, 88
	entropyInitGrowStack(12)
	runtime.GC()
	if opusNativePacketFrame(nil, decoder, &packet[1], int32(len(packet)-1), &pcm[1], 120, 3000) != 2880 || fnv1aFloats(pcm[241:6001]) != 0x20a8ba55 || decoder.FrangeFinal != 0x50373c71 || pcm[0] != 77 || pcm[6001] != 88 {
		t.Fatal("native packet frame golden/offset")
	}
	for i := 1; i <= 240; i++ {
		if pcm[i] != 3 {
			t.Fatal("packet dispatch changed prefix")
		}
	}
}

func TestOpusNativeFECFramePointers(t *testing.T) {
	for _, C := range []int32{1, 2} {
		storage := newOpusFrameOwnerDecoder(t, C)
		decoder := &storage.Decoder
		decoder.Fmode = MODE_SILK_ONLY
		decoder.Fbandwidth = OPUS_BANDWIDTH_NARROWBAND
		decoder.Fframe_size = 2880
		decoder.Fstream_channels = 1
		packet := mustHex(t, "18007523c11e84d40a7ed0075134da9ffc0529ef9f410157b57c1f843e40")
		pcm := make([]float32, 5760*C+2)
		for i := range pcm {
			pcm[i] = 3
		}
		pcm[0], pcm[len(pcm)-1] = 77, 88
		entropyInitGrowStack(12)
		runtime.GC()
		if opusNativeFECFrame(nil, decoder, &packet[1], int32(len(packet)-1), &pcm[1], 5760, 2880) != 2880 || pcm[0] != 77 || pcm[len(pcm)-1] != 88 {
			t.Fatal("native FEC suffix/guards", C)
		}
		for i := int32(1); i <= 2880*C; i++ {
			if pcm[i] != 3 {
				t.Fatal("FEC changed PLC prefix", C, i)
			}
		}
	}
}

func TestOpusNativePLCFramePointers(t *testing.T) {
	for _, C := range []int32{1, 2} {
		storage := newOpusFrameOwnerDecoder(t, C)
		pcm := make([]float32, 240*C+2)
		for i := range pcm {
			pcm[i] = 3
		}
		pcm[0], pcm[len(pcm)-1] = 77, 88
		entropyInitGrowStack(12)
		runtime.GC()
		if opusNativePLCFrame(nil, &storage.Decoder, &pcm[1], 120, 240) != 120 || pcm[0] != 77 || pcm[len(pcm)-1] != 88 {
			t.Fatal("native PLC frame offset")
		}
		for i := int32(1); i <= 240*C; i++ {
			want := float32(0)
			if i <= 120*C {
				want = 3
			}
			if pcm[i] != want {
				t.Fatal("PLC consumption window", C, i)
			}
		}
	}
}

func TestOpusNativePacketDescriptorsPointers(t *testing.T) {
	packet := []byte{255, 65, 1, 10, 11, 12, 99}
	var toc byte
	var sizes [48]int16
	var offset, packetOffset, paddingLength int32
	var padding *byte
	if Opus_opus_packet_parse_impl(nil, &packet[0], 7, 0, &toc, nil, &sizes, &offset, &packetOffset, &padding, &paddingLength) != 1 || toc != 255 || sizes[0] != 3 || offset != 3 || packetOffset != 7 || paddingLength != 1 || padding != &packet[6] {
		t.Fatal("native packet descriptors")
	}
	entropyInitGrowStack(12)
	runtime.GC()
	if *padding != 99 || sizes[0] != 3 {
		t.Fatal("retained typed descriptors")
	}
	iter := OpusT_OpusExtensionIterator{}
	Opus_opus_extension_iterator_init(nil, &iter, padding, paddingLength, 1)
	if iter.Fdata != padding {
		t.Fatal("typed extension padding")
	}
}

func TestOpusDecodeNativeLossFrame(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)

	memory := libc.Xmalloc(tls, uint64(Opus_opus_decoder_get_size(tls, 1)))
	if got := Opus_opus_decoder_init(tls, (*OpusT_OpusDecoder)(unsafe.Pointer(memory)), 24000, 1); got != OPUS_OK {
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
