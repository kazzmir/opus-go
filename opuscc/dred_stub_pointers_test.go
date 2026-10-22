package opuscc

import (
	"runtime"
	"testing"
	"unsafe"
)

// These test the existing unimplemented branches, not enabled native DRED layouts.
func TestDREDStubPointers(t *testing.T) {
	if Opus_opus_dred_get_size() != 0 {
		t.Fatal("unimplemented DRED size")
	}
	if p, err := Opus_opus_dred_alloc(); p != nil || err == nil || err.Error() != opusErrorFromCode(-5).Error() {
		t.Fatal("unimplemented allocation", p, err)
	}
	Opus_opus_dred_free(nil)
	if Opus_opus_dred_parse(nil, nil, nil, -1, -1, 0, nil, -1) != -5 || Opus_opus_dred_process(nil, nil, nil) != -5 {
		t.Fatal("nil DRED arguments")
	}
	if Opus_opus_decoder_dred_decode(nil, nil, -1, nil, -1) != -5 || Opus_opus_decoder_dred_decode24(nil, nil, -1, nil, -1) != -5 || Opus_opus_decoder_dred_decode_float(nil, nil, -1, nil, -1) != -5 {
		t.Fatal("nil decoder/PCM arguments")
	}
	dec := OpusT_OpusDREDDecoder{Floaded: -3, Farch: 7, Fmagic: 0x12345678}
	before := dec
	state := OpusT_OpusDecoder{FFs: 48000, Fchannels: 1, Fdecode_gain: 17}
	beforeState := state
	src, dst := [3]byte{77, 37, 88}, [3]byte{99, 41, 111}
	packet := [3]byte{17, 19, 23}
	end := int32(-41)
	pcm16, pcm24, pcmFloat := [3]int16{77, -7, 88}, [3]int32{77, -9, 88}, [3]float32{77, .25, 88}
	entropyInitGrowStack(12)
	runtime.GC()
	for _, tc := range []struct{ length, max, rate, deferProcessing int32 }{{1, 480, 48000, 0}, {-1, -1, 0, -1}, {2147483647, -2147483648, 1, 2147483647}} {
		if Opus_opus_dred_parse(&dec, unsafe.Pointer(&dst[1]), &packet[1], tc.length, tc.max, tc.rate, &end, tc.deferProcessing) != -5 {
			t.Fatal("parse stub")
		}
	}
	if Opus_opus_dred_process(&dec, unsafe.Pointer(&src[1]), unsafe.Pointer(&dst[1])) != -5 || Opus_opus_dred_process(&dec, unsafe.Pointer(&src[1]), unsafe.Pointer(&src[1])) != -5 {
		t.Fatal("process stub/alias")
	}
	if Opus_opus_decoder_dred_decode(&state, unsafe.Pointer(&src[1]), -1, &pcm16[1], 1) != -5 || Opus_opus_decoder_dred_decode24(&state, unsafe.Pointer(&src[1]), 2147483647, &pcm24[1], 1) != -5 || Opus_opus_decoder_dred_decode_float(&state, unsafe.Pointer(&src[1]), -2147483648, &pcmFloat[1], 1) != -5 {
		t.Fatal("decode stubs")
	}
	Opus_opus_dred_free(unsafe.Pointer(&dst[1]))
	if dec != before || state != beforeState || src != ([3]byte{77, 37, 88}) || dst != ([3]byte{99, 41, 111}) || packet != ([3]byte{17, 19, 23}) || end != -41 || pcm16 != ([3]int16{77, -7, 88}) || pcm24 != ([3]int32{77, -9, 88}) || pcmFloat != ([3]float32{77, .25, 88}) {
		t.Fatal("unimplemented DRED must not access/mutate arguments")
	}
}
