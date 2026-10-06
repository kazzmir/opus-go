package opuscc

import (
	"runtime"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

type opusFrameOwnerTestStorage struct {
	Decoder        OpusT_OpusDecoder
	DecoderPadding [(8 - unsafe.Sizeof(OpusT_OpusDecoder{})%8) % 8]byte
	Silk           OpusT_silk_decoder
	SilkPadding    [(8 - unsafe.Sizeof(OpusT_silk_decoder{})%8) % 8]byte
	Celt           celtStateTestStorage
}

func newOpusFrameOwnerDecoder(t *testing.T, C int32) *opusFrameOwnerTestStorage {
	t.Helper()
	storage := new(opusFrameOwnerTestStorage)
	if Opus_opus_decoder_init(nil, &storage.Decoder, 48000, C) != 0 {
		t.Fatal("frame init")
	}
	if storage.Decoder.Fsilk_dec_offset != int32(unsafe.Offsetof(storage.Silk)) || storage.Decoder.Fcelt_dec_offset != int32(unsafe.Offsetof(storage.Celt)) {
		t.Fatal("scanned frame geometry")
	}
	return storage
}

func TestOpusFrameNumericPCMOffsetPointers(t *testing.T) {
	if opusFramePCMAtBytes(nil, 0) != nil {
		t.Fatal("nil zero byte displacement")
	}
	samples := []float32{77, 11, 22, 88}
	base := &samples[1]
	middle := opusFramePCMAtBytes(base, 4)
	samples = nil
	entropyInitGrowStack(12)
	runtime.GC()
	if *middle != 22 || *opusFramePCMAtBytes(base, 0) != 11 {
		t.Fatal("typed numeric PCM owner")
	}
	*middle = 33
	if *opusFramePCMAtBytes(base, 4) != 33 {
		t.Fatal("PCM alias")
	}
	single := float32(19)
	if *opusFramePCMAtBytes(&single, 0) != 19 {
		t.Fatal("single consumed PCM")
	}
}

func TestOpusFrameFecPointers(t *testing.T) {
	storage := newOpusFrameOwnerDecoder(t, 2)
	decoder := &storage.Decoder
	packet := mustHex(t, "18007523c11e84d40a7ed0075134da9ffc0529ef9f410157b57c1f843e40")
	decoder.Fmode = MODE_SILK_ONLY
	decoder.Fbandwidth = OPUS_BANDWIDTH_NARROWBAND
	decoder.Fstream_channels = 1
	decoder.Fframe_size = 2880
	pcm := make([]float32, 5762)
	pcm[0], pcm[5761] = 77, 88
	for _, fec := range []int32{0, 1} {
		entropyInitGrowStack(12)
		runtime.GC()
		if opusDecodeFrame(nil, decoder, &packet[1], int32(len(packet)-1), &pcm[1], 2880, fec) != 2880 || pcm[0] != 77 || pcm[5761] != 88 || decoder.Fprev_mode != MODE_SILK_ONLY {
			t.Fatal("typed frame FEC", fec)
		}
	}
}

func TestOpusFrameWholeArgumentsPointers(t *testing.T) {
	storage := newOpusFrameOwnerDecoder(t, 1)
	decoder := &storage.Decoder
	sample := float32(77)
	entropyInitGrowStack(12)
	runtime.GC()
	if opusDecodeFrame(nil, decoder, nil, 0, &sample, 119, 0) != -2 || sample != 77 {
		t.Fatal("typed frame minimum size")
	}
	decoder.Fmode = MODE_CELT_ONLY
	decoder.Fframe_size = 240
	decoder.Fbandwidth = OPUS_BANDWIDTH_FULLBAND
	data := [2]byte{165, 73}
	if opusDecodeFrame(nil, decoder, &data[0], 2, &sample, 120, 0) != -1 || sample != 77 {
		t.Fatal("typed frame payload size gate")
	}
}

func TestOpusFrameWholeCursorPointers(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	cursor := &OpusT_opus_ccgo_pseudostack_state{Fscratch_ptr: 123, Fglobal_stack: 456}
	before := *cursor
	libc.Xpthread_setspecific(tls, 0x6f707573, uintptr(unsafe.Pointer(cursor)))
	storage := newOpusFrameOwnerDecoder(t, 1)
	decoder := &storage.Decoder
	decoder.Fmode = MODE_CELT_ONLY
	decoder.Fframe_size = 120
	decoder.Fbandwidth = OPUS_BANDWIDTH_FULLBAND
	data := make([]byte, 128)
	for i := range data {
		data[i] = byte(i*73 + 165)
	}
	pcm := make([]float32, 120)
	entropyInitGrowStack(12)
	runtime.GC()
	if opusDecodeFrame(tls, decoder, &data[0], 128, &pcm[0], 120, 0) != 120 || opusDecodeFrame(tls, decoder, nil, 0, &pcm[0], 120, 0) != 120 || *cursor != before {
		t.Fatal("typed whole frame touched cursor")
	}
}

// Reuse unchanged scalar-C/Go frame goldens from the existing reference fixtures.
func TestOpusFrameActualPacketsPointers(t *testing.T) {
	for _, sequence := range []struct {
		packets        []string
		hashes, ranges []uint32
	}{{[]string{"18007523c11e84d40a7ed0075134da9ffc0529ef9f410157b57c1f843e40", "182312cf4040d200ea1335b36ad4d1a12853dd70b1861253119131ec38"}, []uint32{0x20a8ba55, 0x579c5571}, []uint32{0x50373c71, 0x0671027e}}, {[]string{"7c8cb723a4e954f30817690d8021804d5c6f7c7a79cdaeeeda68b9cc67aab183653ff229912863fc3f7a335205cb0e033bed80eb1a0cfd3f5f", "7c8cc676dcad40dbc8d29d8c3742932526cf5cb73422f4484e23294abafbd6f43cf4f0be8f037456e8bf8a941b1cc6bb", "7c89304a6e5771290464a49a7b9ae7faad938f4cfa4f7e87b076d0fd29b8148c9bba0b085e8d4a8fc70fe4ba126ee8e11a80", "7c89169498b4acaa71ef3c3036299ca0842e1c72362e835f826fcac1eeb088470759f85df05cd98914a64cd3f7972fb7e5023815"}, []uint32{0x53ba9704, 0x5ff5cd0b, 0xc65da4a1, 0xa052e9c5}, []uint32{0x01ad2800, 0x02879800, 0x03449d00, 0x096a4200}}} {
		storage := newOpusFrameOwnerDecoder(t, 2)
		decoder := &storage.Decoder
		for step, text := range sequence.packets {
			packet := mustHex(t, text)
			if packet[0]&3 != 0 {
				t.Fatal("fixture must be single-frame")
			}
			decoder.Fmode = opus_packet_get_mode(nil, &packet[0])
			decoder.Fbandwidth = Opus_opus_packet_get_bandwidth(nil, &packet[0])
			decoder.Fstream_channels = Opus_opus_packet_get_nb_channels(nil, &packet[0])
			decoder.Fframe_size = Opus_opus_packet_get_samples_per_frame(nil, &packet[0], 48000)
			N := decoder.Fframe_size
			pcm := make([]float32, 2*N+2)
			pcm[0], pcm[len(pcm)-1] = 77, 88
			entropyInitGrowStack(12)
			runtime.GC()
			if opusDecodeFrame(nil, decoder, &packet[1], int32(len(packet)-1), &pcm[1], N, 0) != N || fnv1aFloats(pcm[1:1+2*N]) != sequence.hashes[step] || decoder.FrangeFinal != sequence.ranges[step] || pcm[0] != 77 || pcm[len(pcm)-1] != 88 {
				t.Fatal("typed actual frame golden", decoder.Fmode, step)
			}
		}
	}
}

func TestOpusFrameTypedPCMViews(t *testing.T) {
	for _, gain := range []int32{-32768, 0, 256, 32767} {
		left, right := newOpusFrameOwnerDecoder(t, 1), newOpusFrameOwnerDecoder(t, 1)
		for _, decoder := range []*OpusT_OpusDecoder{&left.Decoder, &right.Decoder} {
			decoder.Fmode = MODE_CELT_ONLY
			decoder.Fframe_size = 120
			decoder.Fbandwidth = OPUS_BANDWIDTH_FULLBAND
			decoder.Fdecode_gain = gain
		}
		data := make([]byte, 128)
		for i := range data {
			data[i] = byte(i*73 + 165)
		}
		a, b := make([]float32, 122), make([]float32, 122)
		a[0], a[121], b[0], b[121] = 77, 88, 77, 88
		entropyInitGrowStack(12)
		runtime.GC()
		ra := opusDecodeFrame(nil, &left.Decoder, &data[0], 128, &a[1], 120, 0)
		rb := opus_decode_frame(nil, uintptr(unsafe.Pointer(&right.Decoder)), uintptr(unsafe.Pointer(&data[0])), 128, uintptr(unsafe.Pointer(&b[1])), 120, 0)
		if ra != 120 || ra != rb || left.Decoder != right.Decoder || a[0] != 77 || a[121] != 88 {
			t.Fatal("PCM view gain/guards", gain)
		}
		for i := range a {
			if a[i] != b[i] {
				t.Fatal("gain PCM", gain, i)
			}
		}
	}
}

func TestOpusFrameTypedPayloadEntry(t *testing.T) {
	for _, length := range []int32{0, 1, 128} {
		left, right := newOpusFrameOwnerDecoder(t, 1), newOpusFrameOwnerDecoder(t, 1)
		for _, decoder := range []*OpusT_OpusDecoder{&left.Decoder, &right.Decoder} {
			decoder.Fmode = MODE_CELT_ONLY
			decoder.Fframe_size = 120
			decoder.Fbandwidth = OPUS_BANDWIDTH_FULLBAND
		}
		data := make([]byte, 128)
		for i := range data {
			data[i] = byte(i*73 + 165)
		}
		a, b := make([]float32, 122), make([]float32, 122)
		a[0], a[121], b[0], b[121] = 77, 88, 77, 88
		entropyInitGrowStack(12)
		runtime.GC()
		ra := opusDecodeFrame(nil, &left.Decoder, &data[0], length, &a[1], 120, 0)
		rb := opus_decode_frame(nil, uintptr(unsafe.Pointer(&right.Decoder)), uintptr(unsafe.Pointer(&data[0])), length, uintptr(unsafe.Pointer(&b[1])), 120, 0)
		if ra != 120 || rb != ra || left.Decoder != right.Decoder {
			t.Fatal("typed payload gates", length)
		}
		for i := range a {
			if a[i] != b[i] {
				t.Fatal("typed payload PCM/guards", length, i)
			}
		}
	}
}

func TestOpusFrameTypedDecoderEntry(t *testing.T) {
	left, right := newOpusFrameOwnerDecoder(t, 1), newOpusFrameOwnerDecoder(t, 1)
	a, b := make([]float32, 122), make([]float32, 122)
	a[0], a[121], b[0], b[121] = 77, 88, 77, 88
	entropyInitGrowStack(12)
	runtime.GC()
	ra := opusDecodeFrame(nil, &left.Decoder, nil, 0, &a[1], 120, 0)
	rb := opus_decode_frame(nil, uintptr(unsafe.Pointer(&right.Decoder)), 0, 0, uintptr(unsafe.Pointer(&b[1])), 120, 0)
	if ra != 120 || rb != ra || left.Decoder != right.Decoder {
		t.Fatal("typed/legacy decoder entry")
	}
	for i := range a {
		if a[i] != b[i] {
			t.Fatal("typed decoder PCM/guards")
		}
	}
}

// The complete private frame entry and active consumers are typed.
func TestOpusFrameWholePointers(t *testing.T) {
	for LM := int32(0); LM <= 3; LM++ {
		for _, C := range []int32{1, 2} {
			storage := newOpusFrameOwnerDecoder(t, C)
			decoder := &storage.Decoder
			N := int32(120) << LM
			decoder.Fmode = MODE_CELT_ONLY
			decoder.Fframe_size = N
			decoder.Fbandwidth = OPUS_BANDWIDTH_FULLBAND
			data := make([]byte, 128)
			for i := range data {
				data[i] = byte(i*73 + 165)
			}
			pcm := make([]float32, N*C+2)
			pcm[0], pcm[len(pcm)-1] = 77, 88
			for step := 0; step < 4; step++ {
				var packet *byte
				length := int32(0)
				if step == 0 || step == 3 {
					packet = &data[0]
					length = 128
				}
				entropyInitGrowStack(12)
				runtime.GC()
				if got := opusDecodeFrame(nil, decoder, packet, length, &pcm[1], N, 0); got != N {
					t.Fatal("nil TLS frame", LM, C, step, got)
				}
				if pcm[0] != 77 || pcm[len(pcm)-1] != 88 || decoder.Fprev_mode != MODE_CELT_ONLY {
					t.Fatal("nil TLS frame guards/finalization")
				}
			}
			decoder.Fframe_size = 1920
			long := make([]float32, 1920*C+2)
			long[0], long[len(long)-1] = 77, 88
			entropyInitGrowStack(12)
			runtime.GC()
			if opusDecodeFrame(nil, decoder, nil, 0, &long[1], 1920, 0) != 1920 || long[0] != 77 || long[len(long)-1] != 88 || decoder.FrangeFinal != 0 {
				t.Fatal("nil TLS recursive PLC")
			}
		}
	}
}

func TestOpusFrameSilkWholePointers(t *testing.T) {
	storage := newOpusFrameOwnerDecoder(t, 1)
	decoder := &storage.Decoder
	decoder.Fprev_mode = MODE_SILK_ONLY
	decoder.Fframe_size = 480
	decoder.FDecControl.FnChannelsInternal = 1
	decoder.FDecControl.FinternalSampleRate = 16000
	for _, N := range []int32{120, 480, 960} {
		decoder.Fframe_size = N
		pcm := make([]float32, N+2)
		pcm[0], pcm[N+1] = 77, 88
		entropyInitGrowStack(12)
		runtime.GC()
		if opusDecodeFrame(nil, decoder, nil, 0, &pcm[1], N, 0) != N || pcm[0] != 77 || pcm[N+1] != 88 {
			t.Fatal("nil TLS SILK PLC/short scratch", N)
		}
	}
}

func TestOpusFrameNormalReturnCursor(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	raw := libc.Xmalloc(tls, 16)
	defer libc.Xfree(tls, raw)
	cursor := (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(raw))
	*cursor = OpusT_opus_ccgo_pseudostack_state{Fscratch_ptr: 123, Fglobal_stack: 456}
	before := *cursor
	libc.Xpthread_setspecific(tls, 0x6f707573, raw)
	storage := newOpusFrameOwnerDecoder(t, 1)
	decoder := &storage.Decoder
	decoder.Fmode = MODE_CELT_ONLY
	decoder.Fframe_size = 120
	decoder.Fbandwidth = OPUS_BANDWIDTH_FULLBAND
	data := make([]byte, 128)
	for i := range data {
		data[i] = byte(i*73 + 165)
	}
	pcm := make([]float32, 122)
	pcm[0], pcm[121] = 77, 88
	entropyInitGrowStack(12)
	runtime.GC()
	if opus_decode_frame(tls, uintptr(unsafe.Pointer(decoder)), uintptr(unsafe.Pointer(&data[0])), 128, uintptr(unsafe.Pointer(&pcm[1])), 120, 0) != 120 || decoder.Fprev_mode != MODE_CELT_ONLY || *cursor != before || pcm[0] != 77 || pcm[121] != 88 {
		t.Fatal("normal frame cursor/finalization")
	}
}

func TestOpusFrameEarlyReturnCursor(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	raw := libc.Xmalloc(tls, 16)
	defer libc.Xfree(tls, raw)
	cursor := (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(raw))
	*cursor = OpusT_opus_ccgo_pseudostack_state{Fscratch_ptr: 123, Fglobal_stack: 456}
	before := *cursor
	libc.Xpthread_setspecific(tls, 0x6f707573, raw)
	storage := newOpusFrameOwnerDecoder(t, 1)
	pcm := make([]float32, 122)
	pcm[0], pcm[121] = 77, 88
	if opus_decode_frame(tls, uintptr(unsafe.Pointer(&storage.Decoder)), 0, 0, uintptr(unsafe.Pointer(&pcm[1])), 120, 0) != 120 || *cursor != before || pcm[0] != 77 || pcm[121] != 88 {
		t.Fatal("initial no-packet return")
	}
	storage.Decoder.Fmode = MODE_CELT_ONLY
	storage.Decoder.Fframe_size = 240
	storage.Decoder.Fbandwidth = OPUS_BANDWIDTH_FULLBAND
	data := make([]byte, 64)
	data[0] = 165
	if opus_decode_frame(tls, uintptr(unsafe.Pointer(&storage.Decoder)), uintptr(unsafe.Pointer(&data[0])), 64, uintptr(unsafe.Pointer(&pcm[1])), 120, 0) != -1 || *cursor != before || pcm[0] != 77 || pcm[121] != 88 {
		t.Fatal("bad frame-size return")
	}
}

func TestOpusFrameNoScratchInitialization(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	raw := libc.Xmalloc(tls, 16)
	defer libc.Xfree(tls, raw)
	cursor := (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(raw))
	*cursor = OpusT_opus_ccgo_pseudostack_state{}
	libc.Xpthread_setspecific(tls, 0x6f707573, raw)
	storage := newOpusFrameOwnerDecoder(t, 1)
	sample := float32(77)
	if opus_decode_frame(tls, uintptr(unsafe.Pointer(&storage.Decoder)), 0, 0, uintptr(unsafe.Pointer(&sample)), 119, 0) != -2 || sample != 77 || cursor.Fscratch_ptr != 0 || cursor.Fglobal_stack != 0 {
		t.Fatal("frame scratch/early validation")
	}
}

func TestOpusFrameRedundancyStoragePointers(t *testing.T) {
	for _, C := range []int32{1, 2} {
		temporary := unsafe.SliceData(opusFrameAudioStorage(240 * C))
		state := new(celtStateTestStorage)
		Opus_celt_decoder_init(nil, &state.State, 48000, C)
		data := make([]byte, 64)
		for i := range data {
			data[i] = byte(i*73 + 165)
		}
		entropyInitGrowStack(12)
		runtime.GC()
		if opusFrameCeltRedundant(nil, &state.State, &data[0], 0, 64, temporary, 240) != 240 {
			t.Fatal("redundant Go scratch decode")
		}
		output := make([]float32, 240*C+2)
		output[0], output[len(output)-1] = 77, 88
		for channel := int32(0); channel < C; channel++ {
			for i := int32(0); i < 120; i++ {
				output[1+C*i+channel] = *opusFrameSilkPCM(temporary, uintptr(C*i+channel)*4)
			}
		}
		smooth_fade(nil, opusFrameSilkPCM(temporary, uintptr(120*C)*4), &output[1+120*C], &output[1+120*C], 120, C, mode48000_960_120.Fwindow, 48000)
		if output[0] != 77 || output[len(output)-1] != 88 {
			t.Fatal("redundancy prefix/fade guards")
		}
	}
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
