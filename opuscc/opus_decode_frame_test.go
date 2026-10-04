package opuscc

import (
	"runtime"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

type opusFrameOwnerTestStorage struct {
	Decoder OpusT_OpusDecoder
	Silk    OpusT_silk_decoder
	Celt    celtStateTestStorage
}

func TestOpusFrameSilkPCMPointers(t *testing.T) {
	if opusFrameSilkPCM(nil, 0) != nil {
		t.Fatal("unused nil PCM")
	}
	data := make([]float32, 242)
	data[0], data[241] = 77, 88
	base := &data[1]
	offset := uintptr(0)
	for chunk := 0; chunk < 2; chunk++ {
		pointer := opusFrameSilkPCM(base, offset)
		entropyInitGrowStack(12)
		runtime.GC()
		for i := 0; i < 120; i++ {
			*opusFrameSilkPCM(pointer, uintptr(i)*4) = float32(chunk + 1)
		}
		offset += 120 * 4
	}
	if offset != 240*4 || data[0] != 77 || data[241] != 88 || data[120] != 1 || data[121] != 2 {
		t.Fatal("numeric PCM cursor/guards")
	} /* The terminal offset is numeric only: do not form an unused EOF view. */
}

func TestOpusFrameSilkOwnerPointers(t *testing.T) {
	storage := new(opusFrameOwnerTestStorage)
	storage.Decoder.Fsilk_dec_offset = int32(unsafe.Offsetof(storage.Silk))
	pointer := opusFrameSilkState(&storage.Decoder)
	if pointer != &storage.Silk {
		t.Fatal("SILK interior offset")
	}
	pointer.Fchannel_state[0].FnFramesDecoded = 7
	entropyInitGrowStack(12)
	runtime.GC()
	if pointer.Fchannel_state[0].FnFramesDecoded != 7 || storage.Silk.Fchannel_state[0].FnFramesDecoded != 7 {
		t.Fatal("SILK scanned owner")
	}
	Opus_silk_ResetDecoder(nil, pointer)
	if pointer.Fchannel_state[0].FnFramesDecoded != 0 {
		t.Fatal("SILK reset through typed frame interior")
	}
}

func TestOpusDecodeFrameFieldAccesses(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)

	memory := libc.Xmalloc(tls, uint64(Opus_opus_decoder_get_size(tls, 1)))
	if got := Opus_opus_decoder_init(tls, (*OpusT_OpusDecoder)(unsafe.Pointer(memory)), 24000, 1); got != OPUS_OK {
		t.Fatalf("decoder initialization: got %d", got)
	}
	decoder := (*OpusT_OpusDecoder)(unsafe.Pointer(memory))
	output := make([]float32, 120)
	if got := opus_decode_frame(tls, memory, 0, 0, uintptr(unsafe.Pointer(&output[0])), 120, 0); got != 60 {
		t.Fatalf("decoded samples: got %d, want 60", got)
	}
	if got, want := decoder.Fframe_size, int32(60); got != want {
		t.Fatalf("frame size: got %d, want %d", got, want)
	}
	if got, want := decoder.Fprev_mode, int32(0); got != want {
		t.Fatalf("previous mode: got %d, want %d", got, want)
	}
}
