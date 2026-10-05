package opuscc

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

const msCoupledTailWords = (2*(DEC_PITCH_BUF_SIZE+120) - 1) + 21*8 + 2*CELT_LPC_ORDER

type msCoupledTestChild struct {
	Decoder        OpusT_OpusDecoder
	DecoderPadding [(8 - unsafe.Sizeof(OpusT_OpusDecoder{})%8) % 8]byte
	Silk           OpusT_silk_decoder
	SilkPadding    [(8 - unsafe.Sizeof(OpusT_silk_decoder{})%8) % 8]byte
	Celt           OpusT_OpusCustomDecoder
	Tail           [msCoupledTailWords]float32
	Padding        [(8 - (unsafe.Sizeof(OpusT_OpusCustomDecoder{})+msCoupledTailWords*4)%8) % 8]byte
}
type msTwoStreamTestOwner struct {
	MS       OpusT_OpusMSDecoder
	Padding  [(8 - unsafe.Sizeof(OpusT_OpusMSDecoder{})%8) % 8]byte
	Children [2]msCoupledTestChild
}

func newTwoStreamMSOwner(t *testing.T) *msTwoStreamTestOwner {
	t.Helper()
	owner := new(msTwoStreamTestOwner)
	size := uintptr((uint32(Opus_opus_decoder_get_size(nil, 2)) + 7) &^ 7)
	if unsafe.Sizeof(owner.Children[0]) != size || unsafe.Offsetof(owner.Children) != 272 {
		t.Fatal("MS wrapper geometry")
	}
	owner.MS.Flayout.Fnb_channels = 6
	owner.MS.Flayout.Fnb_streams = 2
	owner.MS.Flayout.Fnb_coupled_streams = 2
	copy(owner.MS.Flayout.Fmapping[:], []byte{2, 0, 1, 2, 3, 255})
	for i := range owner.Children {
		if Opus_opus_decoder_init(nil, &owner.Children[i].Decoder, 48000, 2) != 0 {
			t.Fatal("MS wrapper init")
		}
	}
	return owner
}

func TestMultistreamInt24WrapperPointers(t *testing.T) {
	owner, baseline := newTwoStreamMSOwner(t), newTwoStreamMSOwner(t)
	packet := make([]byte, 129)
	packet[0] = 252
	for i := 1; i < len(packet); i++ {
		packet[i] = byte((i-1)*73 + 165)
	}
	combined := append([]byte{packet[0], 128}, packet[1:]...)
	combined = append(combined, packet...)
	out, want := make([]int32, 5760*6+2), make([]int32, 5760*6+2)
	out[0], out[len(out)-1] = 77, 88
	for _, step := range []int{0, 1, 2} {
		var data *byte
		length, fec := int32(0), int32(0)
		if step != 1 {
			data = &combined[0]
			length = int32(len(combined))
		}
		if step == 2 {
			fec = 1
		}
		entropyInitGrowStack(12)
		runtime.GC()
		got := opusMSDecodeInt24(nil, &owner.MS, data, length, &out[1], 5760, fec)
		expected := opusMSDecodeNative(nil, &baseline.MS, data, length, unsafe.Pointer(&want[1]), opusMSCopyInt24, 5760, fec, 0)
		if got != expected || got <= 0 || owner.Children[0].Decoder != baseline.Children[0].Decoder || owner.Children[1].Decoder != baseline.Children[1].Decoder || out[0] != 77 || out[len(out)-1] != 88 {
			t.Fatal("int24 wrapper dispatch/state")
		}
		for i := int32(0); i < got*6; i++ {
			if out[i+1] != want[i+1] {
				t.Fatal("int24 wrapper PCM", i)
			}
		}
	}
	if opusMSDecodeInt24(nil, &owner.MS, nil, 0, nil, 0, 0) != -1 {
		t.Fatal("int24 wrapper validation")
	}
}

func TestMultistreamShortWrapperPointers(t *testing.T) {
	owner, baseline := newTwoStreamMSOwner(t), newTwoStreamMSOwner(t)
	packet := make([]byte, 129)
	packet[0] = 252
	for i := 1; i < len(packet); i++ {
		packet[i] = byte((i-1)*73 + 165)
	}
	combined := append([]byte{packet[0], 128}, packet[1:]...)
	combined = append(combined, packet...)
	for i := range owner.Children {
		owner.Children[i].Decoder.Fdecode_gain = 32767
		baseline.Children[i].Decoder.Fdecode_gain = 32767
	}
	out, want := make([]int16, 5760*6+2), make([]int16, 5760*6+2)
	out[0], out[len(out)-1] = 77, 88
	for _, step := range []int{0, 1, 2} {
		var data *byte
		length, fec := int32(0), int32(0)
		if step != 1 {
			data = &combined[0]
			length = int32(len(combined))
		}
		if step == 2 {
			fec = 1
		}
		entropyInitGrowStack(12)
		runtime.GC()
		got := opusMSDecodeShort(nil, &owner.MS, data, length, &out[1], 5760, fec)
		expected := opusMSDecodeNative(nil, &baseline.MS, data, length, unsafe.Pointer(&want[1]), opusMSCopyShort, 5760, fec, OPTIONAL_CLIP)
		if got != expected || got <= 0 || owner.Children[0].Decoder != baseline.Children[0].Decoder || owner.Children[1].Decoder != baseline.Children[1].Decoder || out[0] != 77 || out[len(out)-1] != 88 {
			t.Fatal("short wrapper dispatch/clipping/state")
		}
		for i := int32(0); i < got*6; i++ {
			if out[i+1] != want[i+1] {
				t.Fatal("short wrapper PCM", i)
			}
		}
	}
	if opusMSDecodeShort(nil, &owner.MS, nil, 0, nil, 0, 0) != -1 {
		t.Fatal("short wrapper validation")
	}
}

func TestMultistreamFloatWrapperPointers(t *testing.T) {
	owner, baseline := newTwoStreamMSOwner(t), newTwoStreamMSOwner(t)
	packet := mustHex(t, "18007523c11e84d40a7ed0075134da9ffc0529ef9f410157b57c1f843e40")
	combined := append([]byte{packet[0], byte(len(packet) - 1)}, packet[1:]...)
	combined = append(combined, packet...)
	out, want := make([]float32, 5760*6+2), make([]float32, 5760*6+2)
	out[0], out[len(out)-1] = 77, 88
	for _, step := range []int{0, 1, 2} {
		var data *byte
		length, fec := int32(0), int32(0)
		if step != 1 {
			data = &combined[0]
			length = int32(len(combined))
		}
		if step == 2 {
			fec = 1
		}
		entropyInitGrowStack(12)
		runtime.GC()
		got := opusMSDecodeFloat(nil, &owner.MS, data, length, &out[1], 5760, fec)
		expected := opusMSDecodeNative(nil, &baseline.MS, data, length, unsafe.Pointer(&want[1]), opusMSCopyFloat, 5760, fec, 0)
		if got != expected || got <= 0 || owner.Children[0].Decoder != baseline.Children[0].Decoder || owner.Children[1].Decoder != baseline.Children[1].Decoder || out[0] != 77 || out[len(out)-1] != 88 {
			t.Fatal("float wrapper dispatch/state")
		}
		for i := int32(0); i < got*6; i++ {
			if out[i+1] != want[i+1] {
				t.Fatal("float wrapper PCM", i)
			}
		}
	}
	if opusMSDecodeFloat(nil, &owner.MS, nil, 0, nil, 0, 0) != -1 {
		t.Fatal("float wrapper validation")
	}
}

func TestMultistreamWholePointers(t *testing.T) {
	for _, format := range []int{16, 24, 32} {
		owner := new(msTwoStreamTestOwner)
		size := uintptr((uint32(Opus_opus_decoder_get_size(nil, 2)) + 7) &^ 7)
		if unsafe.Sizeof(owner.Children[0]) != size || unsafe.Offsetof(owner.Children) != 272 {
			t.Fatal("scanned MS stream geometry", unsafe.Sizeof(owner.Children[0]), size)
		}
		owner.MS.Flayout.Fnb_channels = 6
		owner.MS.Flayout.Fnb_streams = 2
		owner.MS.Flayout.Fnb_coupled_streams = 2
		copy(owner.MS.Flayout.Fmapping[:], []byte{2, 0, 1, 2, 3, 255})
		for i := range owner.Children {
			if Opus_opus_decoder_init(nil, &owner.Children[i].Decoder, 48000, 2) != 0 {
				t.Fatal("scanned child init")
			}
		}
		var dst unsafe.Pointer
		var copyOut opusMSChannelCopy
		short := make([]int16, 5760*6+2)
		wide := make([]int32, 5760*6+2)
		floats := make([]float32, 5760*6+2)
		short[0], short[len(short)-1] = 77, 88
		wide[0], wide[len(wide)-1] = 77, 88
		floats[0], floats[len(floats)-1] = 77, 88
		switch format {
		case 16:
			dst = unsafe.Pointer(&short[1])
			copyOut = opusMSBindCopy(__ccgo_fp(opus_copy_channel_out_short_legacy), 0)
		case 24:
			dst = unsafe.Pointer(&wide[1])
			copyOut = opusMSBindCopy(__ccgo_fp(opus_copy_channel_out_int24_legacy), 0)
		default:
			dst = unsafe.Pointer(&floats[1])
			copyOut = opusMSBindCopy(__ccgo_fp(opus_copy_channel_out_float_legacy), 0)
		}
		silk := mustHex(t, "18007523c11e84d40a7ed0075134da9ffc0529ef9f410157b57c1f843e40")
		combined := append([]byte{silk[0], byte(len(silk) - 1)}, silk[1:]...)
		combined = append(combined, silk...)
		entropyInitGrowStack(12)
		runtime.GC()
		if opusMSDecodeNative(nil, &owner.MS, &combined[0], int32(len(combined)), dst, copyOut, 5760, 0, 0) != 2880 {
			t.Fatal("whole typed MS decode")
		}
		for i := range owner.Children {
			if owner.Children[i].Decoder.FrangeFinal != 0x50373c71 {
				t.Fatal("whole MS C range")
			}
		}
		for _, fec := range []int32{0, 1} {
			var data *byte
			length := int32(0)
			if fec != 0 {
				data = &combined[0]
				length = int32(len(combined))
			}
			runtime.GC()
			if opusMSDecodeNative(nil, &owner.MS, data, length, dst, copyOut, 5760, fec, 0) != 5760 {
				t.Fatal("whole typed MS PLC/FEC")
			}
		}
		if short[0] != 77 || short[len(short)-1] != 88 || wide[0] != 77 || wide[len(wide)-1] != 88 || floats[0] != 77 || floats[len(floats)-1] != 88 {
			t.Fatal("whole MS guards")
		}
	}
}

func TestMultistreamArgumentsPointers(t *testing.T) {
	owner := newMultistreamOwner(t)
	copyOut := opusMSBindCopy(__ccgo_fp(opus_copy_channel_out_float_legacy), 0)
	value := float32(77)
	for _, item := range []struct {
		packet              []byte
		length, frame, want int32
	}{{nil, 0, 0, -1}, {nil, -1, 480, -1}, {[]byte{3, 0}, 2, 480, -4}, {[]byte{4}, 1, 479, -2}} {
		var data *byte
		if len(item.packet) > 0 {
			data = &item.packet[0]
		}
		before := owner.Child.Decoder
		if result := opusMSDecodeNative(nil, &owner.MS, data, item.length, unsafe.Pointer(&value), copyOut, item.frame, 0, 0); result != item.want || value != 77 || owner.Child.Decoder != before {
			t.Fatal("MS error order/outputs", result, item.want)
		}
	}
}

type multistreamOwnerTestStorage struct {
	MS      OpusT_OpusMSDecoder
	Padding [(8 - unsafe.Sizeof(OpusT_OpusMSDecoder{})%8) % 8]byte
	Child   opusFrameOwnerTestStorage
}

func newMultistreamOwner(t *testing.T) *multistreamOwnerTestStorage {
	owner := new(multistreamOwnerTestStorage)
	owner.Child = *newOpusFrameOwnerDecoder(t, 2)
	owner.MS.Flayout.Fnb_channels = 4
	owner.MS.Flayout.Fnb_streams = 1
	owner.MS.Flayout.Fnb_coupled_streams = 1
	copy(owner.MS.Flayout.Fmapping[:], []byte{0, 1, 0, 255})
	return owner
}

func TestMultistreamNoPseudostack(t *testing.T) {
	owner := newMultistreamOwner(t)
	packet := []byte{4}
	pcm := make([]float32, 5760*4+2)
	pcm[0], pcm[len(pcm)-1] = 77, 88
	callback := __ccgo_fp(opus_copy_channel_out_float_legacy)
	entropyInitGrowStack(12)
	runtime.GC()
	if Opus_opus_multistream_decode_native(nil, uintptr(unsafe.Pointer(&owner.MS)), uintptr(unsafe.Pointer(&packet[0])), 1, uintptr(unsafe.Pointer(&pcm[1])), callback, 10000, 0, 0, 0) != 480 || pcm[0] != 77 || pcm[len(pcm)-1] != 88 {
		t.Fatal("nil TLS MS entry")
	}
	if Opus_opus_multistream_decode_native(nil, uintptr(unsafe.Pointer(&owner.MS)), 0, 0, uintptr(unsafe.Pointer(&pcm[1])), callback, 10000, 0, 0, 0) != 5760 {
		t.Fatal("nil TLS MS capped PLC")
	}
}

func TestMultistreamScalarScratchPointers(t *testing.T) {
	storage := newOpusFrameOwnerDecoder(t, 1)
	var slots struct{ Before, Fs, Offset, After int32 }
	slots.Before, slots.After = 77, 88
	packet := []byte{0, 0}
	pcm := make([]float32, 480)
	entropyInitGrowStack(12)
	runtime.GC()
	if Opus_opus_decoder_ctl_typed(nil, &storage.Decoder, OPUS_GET_SAMPLE_RATE_REQUEST, OpusDecoderCtlArgs{I32: &slots.Fs}) != 0 || opusMSDecodeChild(nil, &storage.Decoder, &packet[0], 2, &pcm[0], 480, 0, 1, &slots.Offset, 0) != 480 || slots.Fs != 48000 || slots.Offset != 2 || slots.Before != 77 || slots.After != 88 {
		t.Fatal("Go MS scalar slots")
	}
}

func TestMultistreamSampleRatePointers(t *testing.T) {
	type owner struct {
		MS      OpusT_OpusMSDecoder
		Padding [(8 - unsafe.Sizeof(OpusT_OpusMSDecoder{})%8) % 8]byte
		Child   opusFrameOwnerTestStorage
	}
	storage := new(owner)
	storage.Child = *newOpusFrameOwnerDecoder(t, 2)
	storage.MS.Flayout.Fnb_streams = 1
	storage.MS.Flayout.Fnb_coupled_streams = 1
	slots := []int32{77, -1, 88}
	entropyInitGrowStack(12)
	runtime.GC()
	if Opus_opus_multistream_decoder_ctl_typed(nil, &storage.MS, OPUS_GET_SAMPLE_RATE_REQUEST, OpusDecoderCtlArgs{I32: &slots[1]}) != 0 || slots[1] != 48000 || slots[0] != 77 || slots[2] != 88 {
		t.Fatal("MS typed sample rate")
	}
}

func TestMultistreamCopyBindingPointers(t *testing.T) {
	src := []float32{0.25, -0.25, 0.5, -0.5, 0.75, -0.75}
	floats := make([]float32, 11)
	shorts := make([]int16, 11)
	wide := make([]int32, 11)
	floats[0], floats[10] = 77, 88
	shorts[0], shorts[10] = 77, 88
	wide[0], wide[10] = 77, 88
	cases := []struct {
		callback uintptr
		dst      unsafe.Pointer
	}{{__ccgo_fp(opus_copy_channel_out_float_legacy), unsafe.Pointer(&floats[1])}, {__ccgo_fp(opus_copy_channel_out_short_legacy), unsafe.Pointer(&shorts[1])}, {__ccgo_fp(opus_copy_channel_out_int24_legacy), unsafe.Pointer(&wide[1])}}
	for _, item := range cases {
		copyOut := opusMSBindCopy(item.callback, 123) // Standard callbacks ignore legacy user-data.
		entropyInitGrowStack(12)
		runtime.GC()
		copyOut(nil, item.dst, 3, 0, &src[0], 2, 3)
		copyOut(nil, item.dst, 3, 2, &src[1], 2, 3)
		copyOut(nil, item.dst, 3, 1, nil, 0, 3)
	}
	for i := 0; i < 3; i++ {
		if floats[1+3*i] != src[2*i] || floats[3+3*i] != src[2*i+1] || shorts[1+3*i] != int16(src[2*i]*32768) || shorts[3+3*i] != int16(src[2*i+1]*32768) || wide[1+3*i] != int32(src[2*i]*8388608) || wide[3+3*i] != int32(src[2*i+1]*8388608) || floats[2+3*i] != 0 || shorts[2+3*i] != 0 || wide[2+3*i] != 0 {
			t.Fatal("bound copy arithmetic/muting")
		}
	}
	if floats[0] != 77 || floats[10] != 88 || shorts[0] != 77 || shorts[10] != 88 || wide[0] != 77 || wide[10] != 88 {
		t.Fatal("copy binding guards")
	}
	called := false
	callback := func(tls *libc.TLS, dst uintptr, ds, dc int32, input uintptr, ss, n int32, user uintptr) {
		entropyInitGrowStack(12)
		runtime.GC()
		called = dst == uintptr(unsafe.Pointer(&floats[1])) && input == uintptr(unsafe.Pointer(&src[0])) && ds == 3 && dc == 1 && ss == 2 && n == 3 && user == 123
	}
	copyOut := opusMSBindLegacyCopy(callback, 123)
	callback = nil
	runtime.GC()
	copyOut(nil, unsafe.Pointer(&floats[1]), 3, 1, &src[0], 2, 3)
	if !called {
		t.Fatal("custom legacy fallback arguments/owner")
	}
}

func TestMultistreamAudioScratchLegacy(t *testing.T) {
	audio := opusFrameAudioStorage(10)
	for i := range audio {
		audio[i] = float32(i+1) / 8
	}
	output := make([]float32, 17)
	output[0], output[16] = 77, 88
	callback := opus_copy_channel_out_float_legacy
	entropyInitGrowStack(12)
	runtime.GC()
	opusMSInvokeLegacy(nil, callback, uintptr(unsafe.Pointer(&output[1])), 3, 0, uintptr(unsafe.Pointer(&audio[0])), 2, 5, 0)
	opusMSInvokeLegacy(nil, callback, uintptr(unsafe.Pointer(&output[1])), 3, 2, uintptr(unsafe.Pointer(&audio[1])), 2, 5, 0)
	opusMSInvokeLegacy(nil, callback, uintptr(unsafe.Pointer(&output[1])), 3, 1, 0, 0, 5, 0)
	for i := 0; i < 5; i++ {
		if output[1+3*i] != audio[2*i] || output[2+3*i] != 0 || output[3+3*i] != audio[2*i+1] {
			t.Fatal("Go stereo/mono/muted scratch")
		}
	}
	if output[0] != 77 || output[16] != 88 {
		t.Fatal("scratch copy guards")
	}
}

func TestMultistreamAudioScratchPointers(t *testing.T) {
	storage := newOpusFrameOwnerDecoder(t, 2)
	audio := opusFrameAudioStorage(960)
	out := make([]float32, 1442)
	out[0], out[1441] = 77, 88
	packet := []byte{4, 0}
	offset := int32(-1)
	entropyInitGrowStack(12)
	runtime.GC()
	if opusMSDecodeChild(nil, &storage.Decoder, &packet[0], 2, unsafe.SliceData(audio), 480, 0, 1, &offset, 0) != 480 {
		t.Fatal("Go scratch native decode")
	}
	opus_copy_channel_out_float(nil, &out[1], 3, 0, &audio[0], 2, 480)
	opus_copy_channel_out_float(nil, &out[1], 3, 2, &audio[1], 2, 480)
	opus_copy_channel_out_float(nil, &out[1], 3, 1, nil, 0, 480)
	if out[0] != 77 || out[1441] != 88 || offset != 2 {
		t.Fatal("typed Go scratch consumers")
	}
}

func TestMultistreamChildOwnerPointers(t *testing.T) {
	type owner struct {
		MS      OpusT_OpusMSDecoder
		Padding [(8 - unsafe.Sizeof(OpusT_OpusMSDecoder{})%8) % 8]byte
		Child   opusFrameOwnerTestStorage
	}
	storage := new(owner)
	storage.Child = *newOpusFrameOwnerDecoder(t, 2)
	offset := uintptr((uint32(268) + 7) / 8 * 8)
	if offset != unsafe.Offsetof(storage.Child) {
		t.Fatal("multistream header geometry")
	}
	decoder := opusMSDecoderAt(&storage.MS, offset)
	storage = nil
	entropyInitGrowStack(12)
	runtime.GC()
	pcm := make([]float32, 962)
	pcm[0], pcm[961] = 77, 88
	packet := []byte{4, 0}
	consumed := int32(-1)
	if opusMSDecodeChild(nil, decoder, &packet[0], 2, &pcm[1], 480, 0, 1, &consumed, 0) != 480 || consumed != 2 || pcm[0] != 77 || pcm[961] != 88 {
		t.Fatal("scanned MS child owner")
	}
}

func TestMultistreamNativeChildPointers(t *testing.T) {
	storage := newOpusFrameOwnerDecoder(t, 1)
	packet := []byte{0, 0}
	pcm := make([]float32, 482)
	pcm[0], pcm[481] = 77, 88
	slot := []int32{77, 99, 88}
	entropyInitGrowStack(12)
	runtime.GC()
	if opusMSDecodeChild(nil, &storage.Decoder, &packet[0], 2, &pcm[1], 480, 0, 1, &slot[1], 0) != 480 || slot[1] != 2 || slot[0] != 77 || slot[2] != 88 || pcm[0] != 77 || pcm[481] != 88 {
		t.Fatal("typed multistream child offsets")
	}
	slot[1] = 99
	if opusMSDecodeChild(nil, &storage.Decoder, nil, 0, &pcm[1], 480, 0, 0, &slot[1], 0) != 480 || slot[1] != 99 {
		t.Fatal("PLC leaves packet offset untouched")
	}
}

func TestMultistreamDecodeNativeCReference(t *testing.T) {
	// SILK NB packets and all expectations are generated by the retained C
	// harness: silent coupled stream plus nonzero mono stream, reordered and
	// duplicated channels, and a muted channel. Both public wrappers exercise
	// the native multistream loop. See the generator's noisy-stream reproducer
	// and summary.txt for the separate pre-existing PCM mismatch.
	f, err := os.Open("testdata/multistream_decode_ref.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	tls := libc.NewTLS()
	defer tls.Close()
	mapping := libc.Xmalloc(tls, 5)
	defer libc.Xfree(tls, mapping)
	copy(unsafe.Slice((*byte)(unsafe.Pointer(mapping)), 5), []byte{2, 0, 1, 2, 255})
	dec, err := Opus_opus_multistream_decoder_create(tls, 48000, 5, 2, 1, mapping)
	if err != nil {
		t.Fatal(err)
	}
	defer Opus_opus_multistream_decoder_destroy(tls, dec)
	data := libc.Xmalloc(tls, 4000)
	defer libc.Xfree(tls, data)
	pcm := libc.Xmalloc(tls, 6400*5*4)
	defer libc.Xfree(tls, pcm)
	va := libc.Xmalloc(tls, uint64(unsafe.Sizeof(uintptr(0))))
	defer libc.Xfree(tls, va)
	rng := libc.Xmalloc(tls, 4)
	defer libc.Xfree(tls, rng)
	scanner := bufio.NewScanner(f)
	for line := 1; scanner.Scan(); line++ {
		var name, packet string
		var length, frame, fec, shortOut, wantRet int32
		var wantHash uint64
		var wantRange uint32
		if n, err := fmt.Sscanf(scanner.Text(), "%s %d %d %d %d %d %d %d %s", &name, &length, &frame, &fec, &shortOut, &wantRet, &wantHash, &wantRange, &packet); err != nil || n != 9 {
			t.Fatalf("line %d: %v", line, err)
		}
		input := uintptr(0)
		if packet != "-" {
			bytes, err := hex.DecodeString(packet)
			if err != nil {
				t.Fatal(err)
			}
			if len(bytes) != int(length) {
				t.Fatal("fixture length mismatch")
			}
			copy(unsafe.Slice((*byte)(unsafe.Pointer(data)), 4000), bytes)
			input = data
		}
		libc.Xmemset(tls, pcm, 0xa5, 6400*5*4)
		var ret int32
		if shortOut != 0 {
			ret = Opus_opus_multistream_decode(tls, dec, input, length, pcm, frame, fec)
		} else {
			ret = Opus_opus_multistream_decode_float(tls, dec, input, length, pcm, frame, fec)
		}
		if ret != wantRet {
			t.Fatalf("%s line %d: return %d, want %d", name, line, ret, wantRet)
		}
		hash := uint64(14695981039346656037)
		if ret > 0 {
			width := 4
			if shortOut != 0 {
				width = 2
			}
			for _, b := range unsafe.Slice((*byte)(unsafe.Pointer(pcm)), int(ret)*5*width) {
				hash ^= uint64(b)
				hash *= 1099511628211
			}
		}
		if hash != wantHash {
			t.Fatalf("%s line %d: PCM hash %d, want %d", name, line, hash, wantHash)
		}
		if ret := Opus_opus_multistream_decoder_ctl(tls, dec, OPUS_GET_FINAL_RANGE_REQUEST, libc.VaList(va, rng)); ret != 0 {
			t.Fatalf("range CTL = %d", ret)
		}
		if got := *(*uint32)(unsafe.Pointer(rng)); got != wantRange {
			t.Fatalf("%s: range %d, want %d", name, got, wantRange)
		}
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
}
