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

func TestOpusFrameSilkTransitionStoragePointers(t *testing.T) {
	transition, celt := opusFrameAudioStorage(240), opusFrameAudioStorage(240)
	temporary := opusFrameAudioStorage(480)
	state := new(OpusT_silk_decoder)
	Opus_silk_InitDecoder(nil, state)
	control := OpusT_silk_DecControlStruct{FnChannelsAPI: 1, FnChannelsInternal: 1, FAPI_sampleRate: 48000, FinternalSampleRate: 16000, FpayloadSize_ms: 10}
	var count int32
	entropyInitGrowStack(12)
	runtime.GC()
	if silk_Decode(nil, state, &control, 1, 1, nil, unsafe.SliceData(temporary), &count, 0) != 0 || count != 480 {
		t.Fatal("SILK transition PLC")
	}
	copy(transition, temporary[:240])
	transition[0] = 3
	if celt[0] != 0 || transition[0] != 3 {
		t.Fatal("transition owners alias")
	}
	output := make([]float32, 242)
	output[0], output[241] = 77, 88
	copy(output[1:121], transition[:120])
	smooth_fade(nil, &transition[120], &output[121], &output[121], 120, 1, mode48000_960_120.Fwindow, 48000)
	if output[0] != 77 || output[241] != 88 || output[1] != 3 {
		t.Fatal("SILK transition prefix/fade guards")
	}
}

func TestOpusFrameCeltTransitionStoragePointers(t *testing.T) {
	for _, C := range []int32{1, 2} {
		temporary := opusFrameAudioStorage(240 * C)
		pointer := unsafe.SliceData(temporary)
		state := new(celtStateTestStorage)
		Opus_celt_decoder_init(nil, &state.State, 48000, C)
		entropyInitGrowStack(12)
		runtime.GC()
		if opusFrameCelt(nil, &state.State, nil, 0, pointer, 240, nil, 0, 0) != 240 {
			t.Fatal("CELT transition scratch")
		}
		output := make([]float32, 240*C+2)
		output[0], output[len(output)-1] = 77, 88
		copy(output[1:1+120*C], temporary[:120*C])
		smooth_fade(nil, opusFrameSilkPCM(pointer, uintptr(120*C)*4), &output[1+120*C], &output[1+120*C], 120, C, mode48000_960_120.Fwindow, 48000)
		if output[0] != 77 || output[len(output)-1] != 88 || pointer != &temporary[0] {
			t.Fatal("transition storage guards/owner")
		}
	}
}

func TestOpusFrameSilkScratchPointers(t *testing.T) {
	if opusFrameAudioStorage(0) != nil {
		t.Fatal("empty scratch")
	}
	temporary := opusFrameAudioStorage(480)
	state := new(OpusT_silk_decoder)
	Opus_silk_InitDecoder(nil, state)
	control := OpusT_silk_DecControlStruct{FnChannelsAPI: 1, FnChannelsInternal: 1, FAPI_sampleRate: 48000, FinternalSampleRate: 16000, FpayloadSize_ms: 10}
	var count int32
	entropyInitGrowStack(12)
	runtime.GC()
	if silk_Decode(nil, state, &control, 1, 1, nil, unsafe.SliceData(temporary), &count, 0) != 0 || count != 480 {
		t.Fatal("short SILK scratch", count)
	}
	output := make([]float32, 242)
	output[0], output[241] = 77, 88
	copy(output[1:241], temporary[:240])
	if output[0] != 77 || output[241] != 88 || len(temporary) != 480 {
		t.Fatal("short frame truncation/guards")
	}
}

func TestOpusFrameCeltOwnerPointers(t *testing.T) {
	storage := new(opusFrameOwnerTestStorage)
	storage.Decoder.Fcelt_dec_offset = int32(unsafe.Offsetof(storage.Celt))
	state := opusFrameCeltState(&storage.Decoder)
	if state != &storage.Celt.State {
		t.Fatal("CELT interior offset")
	}
	if Opus_celt_decoder_init(nil, state, 48000, 2) != 0 {
		t.Fatal("CELT init")
	}
	data := make([]byte, 64)
	for i := range data {
		data[i] = byte(i*73 + 165)
	}
	pcm := make([]float32, 242)
	pcm[0], pcm[241] = 77, 88
	entropyInitGrowStack(12)
	runtime.GC()
	if opusFrameCelt(nil, state, &data[0], 64, &pcm[1], 120, nil, 0, 0) != 120 {
		t.Fatal("CELT typed interior decode")
	}
	var mode *OpusT_OpusCustomMode
	var rng uint32
	if Opus_opus_custom_decoder_ctl_typed(nil, state, CELT_GET_MODE_REQUEST, OpusDecoderCtlArgs{Mode: &mode}) != 0 || mode != &mode48000_960_120 || Opus_opus_custom_decoder_ctl_typed(nil, state, OPUS_GET_FINAL_RANGE_REQUEST, OpusDecoderCtlArgs{U32: &rng}) != 0 || rng != state.Frng || pcm[0] != 77 || pcm[241] != 88 {
		t.Fatal("CELT interior range/mode/guards")
	}
}

func TestOpusFrameSilkDispatchPointers(t *testing.T) {
	storage := new(opusFrameOwnerTestStorage)
	storage.Decoder.Fsilk_dec_offset = int32(unsafe.Offsetof(storage.Silk))
	decoder := &storage.Decoder
	decoder.FDecControl = OpusT_silk_DecControlStruct{FnChannelsAPI: 1, FnChannelsInternal: 1, FAPI_sampleRate: 48000, FinternalSampleRate: 16000, FpayloadSize_ms: 20}
	state := opusFrameSilkState(decoder)
	Opus_silk_InitDecoder(nil, state)
	var ec OpusT_ec_ctx
	var count int32
	pcm := make([]float32, 962)
	pcm[0], pcm[961] = 77, 88
	for call := 0; call < 3; call++ {
		entropyInitGrowStack(12)
		runtime.GC()
		if silk_Decode(nil, state, &decoder.FDecControl, 1, 1, &ec, &pcm[1], &count, decoder.Farch) != 0 || count != 960 || pcm[0] != 77 || pcm[961] != 88 {
			t.Fatal("frame typed SILK dispatch", call, count)
		}
		if ec != (OpusT_ec_ctx{}) {
			t.Fatal("PLC consumed frame entropy")
		}
	}
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
