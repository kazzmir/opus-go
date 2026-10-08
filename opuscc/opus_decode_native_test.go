package opuscc

import (
	"runtime"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestOpusInt24WholePointers(t *testing.T) {
	if opusDecodeInt24(nil, nil, nil, 0, nil, 0, 2) != -1 {
		t.Fatal("int24 typed minimum frame")
	}
	for _, C := range []int32{1, 2} {
		for _, packet := range [][]byte{{0}, {3, 2}, mustHex(t, "18007523c11e84d40a7ed0075134da9ffc0529ef9f410157b57c1f843e40"), mustHex(t, "7c8cb723a4e954f30817690d8021804d5c6f7c7a79cdaeeeda68b9cc67aab183653ff229912863fc3f7a335205cb0e033bed80eb1a0cfd3f5f")} {
			storage := newOpusFrameOwnerDecoder(t, C)
			pcm := make([]int32, 5760*C+2)
			pcm[0], pcm[len(pcm)-1] = 77, 88
			entropyInitGrowStack(12)
			runtime.GC()
			count := opusDecodeInt24(nil, &storage.Decoder, &packet[0], int32(len(packet)), &pcm[1], 5760, 0)
			if count <= 0 || storage.Decoder.Flast_packet_duration != count || pcm[0] != 77 || pcm[len(pcm)-1] != 88 {
				t.Fatal("whole int24 normal", C, packet[0], count)
			}
			if packet[0] == 24 && storage.Decoder.FrangeFinal != 0x50373c71 {
				t.Fatal("int24 C range golden")
			}
			for _, fec := range []int32{0, 1} {
				var data *byte
				length := int32(0)
				if fec != 0 {
					data = &packet[0]
					length = int32(len(packet))
				}
				entropyInitGrowStack(12)
				runtime.GC()
				if opusDecodeInt24(nil, &storage.Decoder, data, length, &pcm[1], 5760, fec) != 5760 || pcm[0] != 77 || pcm[len(pcm)-1] != 88 {
					t.Fatal("whole int24 PLC/FEC")
				}
			}
		}
	}
}

func TestOpusIntegerRatesPointers(t *testing.T) {
	packet := mustHex(t, "18007523c11e84d40a7ed0075134da9ffc0529ef9f410157b57c1f843e40")
	for _, rate := range []int32{8000, 12000, 16000, 24000, 48000} {
		for _, C := range []int32{1, 2} {
			left, right := newOpusFrameOwnerDecoder(t, C), newOpusFrameOwnerDecoder(t, C)
			Opus_opus_decoder_init(nil, &left.Decoder, rate, C)
			Opus_opus_decoder_init(nil, &right.Decoder, rate, C)
			N := rate * 60 / 1000
			a, b := make([]int16, N*C+2), make([]int32, N*C+2)
			a[0], a[len(a)-1], b[0], b[len(b)-1] = 77, 88, 77, 88
			entropyInitGrowStack(12)
			runtime.GC()
			if Opus_opus_decode_typed(nil, &left.Decoder, &packet[0], int32(len(packet)), &a[1], N, 0) != N || opusDecodeInt24(nil, &right.Decoder, &packet[0], int32(len(packet)), &b[1], N, 0) != N || left.Decoder.FrangeFinal != 0x50373c71 || right.Decoder.FrangeFinal != 0x50373c71 || a[0] != 77 || a[len(a)-1] != 88 || b[0] != 77 || b[len(b)-1] != 88 {
				t.Fatal("integer API rate/guards", rate, C)
			}
		}
	}
}

func TestOpusIntegerArgumentsPointers(t *testing.T) {
	invalid := []byte{3, 0}
	for _, format := range []int{16, 24} {
		storage := newOpusFrameOwnerDecoder(t, 1)
		storage.Decoder.Fchannels = 0
		before := storage.Decoder
		short := int16(77)
		wide := int32(88)
		var result int32
		if format == 16 {
			result = opusDecodeInt16(nil, &storage.Decoder, &invalid[0], 2, &short, 120, 0)
		} else {
			result = opusDecodeInt24(nil, &storage.Decoder, &invalid[0], 2, &wide, 120, 0)
		}
		if result != -4 || short != 77 || wide != 88 || storage.Decoder != before {
			t.Fatal("packet validation precedes channel assert")
		}
		storage = newOpusFrameOwnerDecoder(t, 1)
		packet := []byte{0}
		if format == 16 {
			result = opusDecodeInt16(nil, &storage.Decoder, &packet[0], 1, &short, 479, 0)
		} else {
			result = opusDecodeInt24(nil, &storage.Decoder, &packet[0], 1, &wide, 479, 0)
		}
		if result != -2 || short != 77 || wide != 88 {
			t.Fatal("failed integer decode wrote PCM")
		}
	}
}

func TestOpusInt16WholePointers(t *testing.T) {
	if opusDecodeInt16(nil, nil, nil, 0, nil, 0, 2) != -1 {
		t.Fatal("int16 typed minimum frame")
	}
	for _, C := range []int32{1, 2} {
		for _, packet := range [][]byte{{0}, {3, 2}, mustHex(t, "18007523c11e84d40a7ed0075134da9ffc0529ef9f410157b57c1f843e40"), mustHex(t, "7c8cb723a4e954f30817690d8021804d5c6f7c7a79cdaeeeda68b9cc67aab183653ff229912863fc3f7a335205cb0e033bed80eb1a0cfd3f5f")} {
			storage := newOpusFrameOwnerDecoder(t, C)
			pcm := make([]int16, 5760*C+2)
			pcm[0], pcm[len(pcm)-1] = 77, 88
			entropyInitGrowStack(12)
			runtime.GC()
			count := opusDecodeInt16(nil, &storage.Decoder, &packet[0], int32(len(packet)), &pcm[1], 5760, 0)
			if count <= 0 || storage.Decoder.Flast_packet_duration != count || pcm[0] != 77 || pcm[len(pcm)-1] != 88 {
				t.Fatal("whole int16 normal", C, packet[0], count)
			}
			if packet[0] == 24 && storage.Decoder.FrangeFinal != 0x50373c71 {
				t.Fatal("int16 C range golden")
			}
			for _, fec := range []int32{0, 1} {
				var data *byte
				length := int32(0)
				if fec != 0 {
					data = &packet[0]
					length = int32(len(packet))
				}
				entropyInitGrowStack(12)
				runtime.GC()
				if opusDecodeInt16(nil, &storage.Decoder, data, length, &pcm[1], 5760, fec) != 5760 || pcm[0] != 77 || pcm[len(pcm)-1] != 88 {
					t.Fatal("whole int16 PLC/FEC")
				}
			}
		}
	}
}

func TestOpusInt24NoPseudostack(t *testing.T) {
	if Opus_opus_decode24(nil, 0, 0, 0, 0, 0, 2) != -1 {
		t.Fatal("int24 validation order")
	}
	for _, C := range []int32{1, 2} {
		storage := newOpusFrameOwnerDecoder(t, C)
		packet := mustHex(t, "18007523c11e84d40a7ed0075134da9ffc0529ef9f410157b57c1f843e40")
		pcm := make([]int32, 5760*C+2)
		pcm[0], pcm[len(pcm)-1] = 77, 88
		entropyInitGrowStack(12)
		runtime.GC()
		if Opus_opus_decode24(nil, uintptr(unsafe.Pointer(&storage.Decoder)), uintptr(unsafe.Pointer(&packet[0])), int32(len(packet)), uintptr(unsafe.Pointer(&pcm[1])), 5760, 0) != 2880 || pcm[0] != 77 || pcm[len(pcm)-1] != 88 {
			t.Fatal("int24 nil TLS", C)
		}
		if Opus_opus_decode24(nil, uintptr(unsafe.Pointer(&storage.Decoder)), 0, 0, uintptr(unsafe.Pointer(&pcm[1])), 120, 0) != 120 || pcm[0] != 77 || pcm[len(pcm)-1] != 88 {
			t.Fatal("int24 short PLC")
		}
	}
}

func TestOpusInt16NoPseudostack(t *testing.T) {
	if Opus_opus_decode(nil, 0, 0, 0, 0, 0, 2) != -1 {
		t.Fatal("int16 validation order")
	}
	for _, C := range []int32{1, 2} {
		storage := newOpusFrameOwnerDecoder(t, C)
		packet := mustHex(t, "18007523c11e84d40a7ed0075134da9ffc0529ef9f410157b57c1f843e40")
		pcm := make([]int16, 5760*C+2)
		pcm[0], pcm[len(pcm)-1] = 77, 88
		entropyInitGrowStack(12)
		runtime.GC()
		if Opus_opus_decode(nil, uintptr(unsafe.Pointer(&storage.Decoder)), uintptr(unsafe.Pointer(&packet[0])), int32(len(packet)), uintptr(unsafe.Pointer(&pcm[1])), 5760, 0) != 2880 || pcm[0] != 77 || pcm[len(pcm)-1] != 88 {
			t.Fatal("int16 nil TLS", C)
		}
		if Opus_opus_decode(nil, uintptr(unsafe.Pointer(&storage.Decoder)), 0, 0, uintptr(unsafe.Pointer(&pcm[1])), 120, 0) != 120 || pcm[0] != 77 || pcm[len(pcm)-1] != 88 {
			t.Fatal("int16 short PLC")
		}
	}
}

func TestOpusInt24ScratchPointers(t *testing.T) {
	storage := newOpusFrameOwnerDecoder(t, 2)
	packet := mustHex(t, "18007523c11e84d40a7ed0075134da9ffc0529ef9f410157b57c1f843e40")
	out := opusFrameAudioStorage(2880 * 2)
	pcm := make([]int32, 5762)
	pcm[0], pcm[5761] = 77, 88
	entropyInitGrowStack(12)
	runtime.GC()
	ret := opusDecodeNative(nil, &storage.Decoder, &packet[0], int32(len(packet)), unsafe.SliceData(out), 2880, 0, 0, nil, 0)
	if ret != 2880 {
		t.Fatal("int24 scratch decode")
	}
	opusDecodeInt24PCM(nil, unsafe.SliceData(out), &pcm[1], ret*2)
	if pcm[0] != 77 || pcm[5761] != 88 || storage.Decoder.FrangeFinal != 0x50373c71 {
		t.Fatal("int24 scratch guards/range")
	}
}

func TestOpusIntegerDecodeGoScratch(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	raw := libc.Xmalloc(tls, 16)
	defer libc.Xfree(tls, raw)
	cursor := (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(raw))
	*cursor = OpusT_opus_ccgo_pseudostack_state{Fscratch_ptr: 123, Fglobal_stack: 456}
	before := *cursor
	libc.Xpthread_setspecific(tls, 0x6f707573, raw)
	packet := mustHex(t, "18007523c11e84d40a7ed0075134da9ffc0529ef9f410157b57c1f843e40")
	for _, C := range []int32{1, 2} {
		left, right := newOpusFrameOwnerDecoder(t, C), newOpusFrameOwnerDecoder(t, C)
		a, b := make([]int16, 5760*C+2), make([]int32, 5760*C+2)
		a[0], a[len(a)-1], b[0], b[len(b)-1] = 77, 88, 77, 88
		entropyInitGrowStack(12)
		runtime.GC()
		if Opus_opus_decode(tls, uintptr(unsafe.Pointer(&left.Decoder)), uintptr(unsafe.Pointer(&packet[0])), int32(len(packet)), uintptr(unsafe.Pointer(&a[1])), 5760, 0) != 2880 || Opus_opus_decode24(tls, uintptr(unsafe.Pointer(&right.Decoder)), uintptr(unsafe.Pointer(&packet[0])), int32(len(packet)), uintptr(unsafe.Pointer(&b[1])), 5760, 0) != 2880 || a[0] != 77 || a[len(a)-1] != 88 || b[0] != 77 || b[len(b)-1] != 88 || *cursor != before {
			t.Fatal("integer Go scratch/cursors", C)
		}
	}
}

func TestOpusInt16ScratchPointers(t *testing.T) {
	storage := newOpusFrameOwnerDecoder(t, 2)
	packet := mustHex(t, "18007523c11e84d40a7ed0075134da9ffc0529ef9f410157b57c1f843e40")
	out := opusFrameAudioStorage(2880 * 2)
	pcm := make([]int16, 5762)
	pcm[0], pcm[5761] = 77, 88
	entropyInitGrowStack(12)
	runtime.GC()
	ret := opusDecodeNative(nil, &storage.Decoder, &packet[0], int32(len(packet)), unsafe.SliceData(out), 2880, 0, 0, nil, OPTIONAL_CLIP)
	if ret != 2880 {
		t.Fatal("int16 scratch decode")
	}
	Opus_celt_float2int16_c(nil, unsafe.SliceData(out), &pcm[1], ret*2)
	if pcm[0] != 77 || pcm[5761] != 88 || storage.Decoder.FrangeFinal != 0x50373c71 {
		t.Fatal("int16 scratch guards/range")
	}
}

func TestOpusInt24PCMPointers(t *testing.T) {
	opusDecodeInt24PCM(nil, nil, nil, 0)
	opusDecodeInt24PCM(nil, nil, nil, -1)
	input := []float32{-1, -.5, 0, .5, 1, 1.5 / 8388608, 2.5 / 8388608, -1.5 / 8388608, -2.5 / 8388608}
	output := make([]int32, len(input)+2)
	output[0], output[len(output)-1] = 77, 88
	entropyInitGrowStack(12)
	runtime.GC()
	opusDecodeInt24PCM(nil, &input[0], &output[1], int32(len(input)))
	want := []int32{-8388608, -4194304, 0, 4194304, 8388608, 2, 2, -2, -2}
	for i := range want {
		if output[1+i] != want[i] {
			t.Fatal("int24 rounding", i, output[1+i])
		}
	}
	if output[0] != 77 || output[len(output)-1] != 88 {
		t.Fatal("int24 guards")
	} /* Typed Go-only live float/int alias: C effective-type rules differ. */
	alias := []int32{0, 0}
	floats := unsafe.Slice((*float32)(unsafe.Pointer(&alias[0])), 2)
	floats[0], floats[1] = .5, 1
	opusDecodeInt24PCM(nil, &floats[0], &alias[0], 2)
	if alias[0] != 4194304 || alias[1] != 8388608 {
		t.Fatal("live int24 alias")
	}
}

func TestOpusFloatDecodeWholePointers(t *testing.T) {
	if opusDecodeFloat(nil, nil, nil, 0, nil, 0, 2) != -1 {
		t.Fatal("float validation order")
	}
	for _, tc := range []struct {
		text      string
		n         int32
		hash, rng uint32
	}{{"18007523c11e84d40a7ed0075134da9ffc0529ef9f410157b57c1f843e40", 2880, 0x20a8ba55, 0x50373c71}, {"7c8cb723a4e954f30817690d8021804d5c6f7c7a79cdaeeeda68b9cc67aab183653ff229912863fc3f7a335205cb0e033bed80eb1a0cfd3f5f", 960, 0x53ba9704, 0x01ad2800}} {
		storage := newOpusFrameOwnerDecoder(t, 2)
		packet := mustHex(t, tc.text)
		pcm := make([]float32, 11522)
		pcm[0], pcm[11521] = 77, 88
		entropyInitGrowStack(12)
		runtime.GC()
		if Opus_opus_decode_float_typed(nil, &storage.Decoder, &packet[0], int32(len(packet)), &pcm[1], 5760, 0) != tc.n || fnv1aFloats(pcm[1:1+2*tc.n]) != tc.hash || storage.Decoder.FrangeFinal != tc.rng || pcm[0] != 77 || pcm[11521] != 88 {
			t.Fatal("float decode typed golden")
		}
		entropyInitGrowStack(12)
		runtime.GC()
		if opusDecodeFloat(nil, &storage.Decoder, nil, 0, &pcm[1], 5760, 0) != 5760 {
			t.Fatal("float PLC")
		}
	}
}

func TestOpusNativePacketOffsetPointers(t *testing.T) {
	for _, frame := range []int32{479, 480} {
		storage := newOpusFrameOwnerDecoder(t, 1)
		packet := []byte{0, 0, 99}
		output := make([]float32, 482)
		output[0], output[481] = 77, 88
		slot := []int32{77, 99, 88}
		before := storage.Decoder.Fmode
		entropyInitGrowStack(12)
		runtime.GC()
		result := opusDecodeNative(nil, &storage.Decoder, &packet[0], 3, &output[1], frame, 0, 1, &slot[1], 0)
		want := int32(480)
		if frame == 479 {
			want = -2
			if storage.Decoder.Fmode != before {
				t.Fatal("premature packet metadata commit")
			}
		}
		if result != want || slot[1] != 2 || slot[0] != 77 || slot[2] != 88 || output[0] != 77 || output[481] != 88 {
			t.Fatal("typed packet offset/error order", frame, result, slot)
		}
	}
}

func TestOpusNativeTypedPayloadEntry(t *testing.T) {
	for _, packet := range [][]byte{{16}, mustHex(t, "18007523c11e84d40a7ed0075134da9ffc0529ef9f410157b57c1f843e40")} {
		left, right := newOpusFrameOwnerDecoder(t, 2), newOpusFrameOwnerDecoder(t, 2)
		a, b := make([]float32, 5762), make([]float32, 5762)
		a[0], a[5761], b[0], b[5761] = 77, 88, 77, 88
		entropyInitGrowStack(12)
		runtime.GC()
		ra := opusDecodeNative(nil, &left.Decoder, &packet[0], int32(len(packet)), &a[1], 2880, 0, 0, nil, 0)
		rb := Opus_opus_decode_native(nil, uintptr(unsafe.Pointer(&right.Decoder)), uintptr(unsafe.Pointer(&packet[0])), int32(len(packet)), uintptr(unsafe.Pointer(&b[1])), 2880, 0, 0, 0, 0, 0, 0)
		if ra <= 0 || rb != ra || left.Decoder != right.Decoder || a[0] != 77 || a[5761] != 88 {
			t.Fatal("native typed payload")
		}
		for i := range a {
			if a[i] != b[i] {
				t.Fatal("native typed payload PCM")
			}
		}
	}
}

func TestOpusNativeEmptyPaddingPointers(t *testing.T) {
	for _, packet := range [][]byte{{0}, {1}, {3, 2}, {0, 0, 99}, {255, 65, 1, 10, 11, 12, 99}} {
		var toc byte
		var size [48]int16
		var offset, consumed, padLength int32
		var padding *byte
		selfDelimited := int32(0)
		if len(packet) == 3 && packet[0] == 0 {
			selfDelimited = 1
		}
		count := opusNativeParsePacket(nil, &packet[0], int32(len(packet)), selfDelimited, &toc, &size, &offset, &consumed, &padding, &padLength)
		if count < 0 {
			t.Fatal("padding fixture")
		}
		entropyInitGrowStack(12)
		runtime.GC()
		if packet[0] == 255 {
			if padLength != 1 || padding != &packet[6] || *padding != 99 {
				t.Fatal("consumed padding view")
			}
		} else if padLength != 0 || padding != nil {
			t.Fatal("unused EOF padding view")
		}
	}
}

func TestOpusNativeArgumentsPointers(t *testing.T) {
	for _, tc := range []struct {
		data               *byte
		length, frame, fec int32
		want               int32
	}{{nil, 0, 119, 0, -1}, {nil, 0, 120, 2, -1}, {nil, 0, 120, -1, -1}} {
		storage := newOpusFrameOwnerDecoder(t, 1)
		sample := float32(77)
		if opusDecodeNative(nil, &storage.Decoder, tc.data, tc.length, &sample, tc.frame, tc.fec, 0, nil, 0) != tc.want || sample != 77 {
			t.Fatal("native argument order")
		}
	}
	storage := newOpusFrameOwnerDecoder(t, 1)
	sample := float32(77)
	invalid := []byte{3, 0}
	if opusDecodeNative(nil, &storage.Decoder, &invalid[0], 2, &sample, 120, 0, 0, nil, 0) != -4 || sample != 77 {
		t.Fatal("native invalid packet")
	}
}

func TestOpusNativeWholePointers(t *testing.T) {
	celtMono, celtStereo := make([]byte, 129), make([]byte, 129)
	celtMono[0], celtStereo[0] = 248, 252
	for i := 1; i < 129; i++ {
		celtMono[i] = byte((i-1)*73 + 165)
		celtStereo[i] = celtMono[i]
	}
	for _, C := range []int32{1, 2} {
		for _, packet := range [][]byte{{0}, {1}, {3, 2}, {3, 65, 1, 99}, celtMono, celtStereo, mustHex(t, "18007523c11e84d40a7ed0075134da9ffc0529ef9f410157b57c1f843e40"), mustHex(t, "7c8cb723a4e954f30817690d8021804d5c6f7c7a79cdaeeeda68b9cc67aab183653ff229912863fc3f7a335205cb0e033bed80eb1a0cfd3f5f")} {
			storage := newOpusFrameOwnerDecoder(t, C)
			pcm := make([]float32, 5760*C+2)
			pcm[0], pcm[len(pcm)-1] = 77, 88
			var offset int32
			entropyInitGrowStack(12)
			runtime.GC()
			n := opusDecodeNative(nil, &storage.Decoder, &packet[0], int32(len(packet)), &pcm[1], 5760, 0, 0, &offset, 0)
			if n <= 0 || offset != int32(len(packet)) || storage.Decoder.Flast_packet_duration != n || pcm[0] != 77 || pcm[len(pcm)-1] != 88 {
				t.Fatal("whole native packet descriptors/duration", C, packet[0], n, offset)
			}
			if C == 2 && packet[0] == 24 && (fnv1aFloats(pcm[1:1+2*n]) != 0x20a8ba55 || storage.Decoder.FrangeFinal != 0x50373c71) {
				t.Fatal("whole native SILK C golden")
			}
			if C == 2 && packet[0] == 124 && (fnv1aFloats(pcm[1:1+2*n]) != 0x53ba9704 || storage.Decoder.FrangeFinal != 0x01ad2800) {
				t.Fatal("whole native hybrid golden")
			}
			for _, fec := range []int32{0, 1} {
				entropyInitGrowStack(12)
				runtime.GC()
				var input *byte
				length := int32(0)
				if fec != 0 {
					input = &packet[0]
					length = int32(len(packet))
				}
				result := opusDecodeNative(nil, &storage.Decoder, input, length, &pcm[1], 5760, fec, 0, nil, 1)
				if result != 5760 || pcm[0] != 77 || pcm[len(pcm)-1] != 88 {
					t.Fatal("native PLC/FEC recursion", C, packet[0], fec, result)
				}
			}
		}
	}
}

func TestOpusNativePayloadPointers(t *testing.T) {
	if opusNativePayload(nil, 99, 0) != nil || opusNativePayload(nil, 99, 1) != nil {
		t.Fatal("unused payload view")
	}
	packet := []byte{77, 88, 10, 11, 12}
	view := opusNativePayload(&packet[0], 2, 3)
	entropyInitGrowStack(12)
	runtime.GC()
	if view != &packet[2] || unsafe.Slice(view, 3)[2] != 12 {
		t.Fatal("numeric packet cursor")
	}
	if opusNativePayload(&packet[0], 5, 0) != nil || opusNativePayload(&packet[0], ^uint(0), -1) != nil {
		t.Fatal("unused exact EOF view")
	}
}

func TestOpusNativeTypedDecoderEntry(t *testing.T) {
	left, right := newOpusFrameOwnerDecoder(t, 1), newOpusFrameOwnerDecoder(t, 1)
	a, b := make([]float32, 122), make([]float32, 122)
	a[0], a[121], b[0], b[121] = 77, 88, 77, 88
	entropyInitGrowStack(12)
	runtime.GC()
	ra := opusDecodeNative(nil, &left.Decoder, nil, 0, &a[1], 120, 0, 0, nil, 0)
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
