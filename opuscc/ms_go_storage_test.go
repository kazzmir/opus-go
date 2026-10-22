package opuscc

import (
	"runtime"
	"slices"
	"testing"
	"unsafe"
)

func TestMSDecoderStaticGoAlignmentPointers(t *testing.T) {
	if unsafe.Sizeof(opusGoMonoAllocation{}) != 32768 || unsafe.Offsetof(opusGoMonoAllocation{}.Cell) != 0 || unsafe.Offsetof(opusGoMonoCell{}.Body) != 0 {
		t.Fatal("fixed mono allocation geometry")
	}
	if unsafe.Sizeof(opusGoStereoAligned{})%16 != unsafe.Alignof(opusGoStereoCell{}) {
		t.Fatal("typed stereo alignment arrays must visit every alignment residue")
	}
	mono := opusGoMonoStorage(nil)
	stereo := new([16 / unsafe.Alignof(opusGoStereoCell{})]opusGoStereoAligned)
	foundMono, foundStereo := uintptr(unsafe.Pointer(mono.Decoder))&15 == 0, false
	for i := range stereo {
		foundStereo = foundStereo || uintptr(unsafe.Pointer(&stereo[i].Cell.Body.Decoder))&15 == 0
	}
	if !foundMono || !foundStereo {
		t.Fatal("mono allocation/stereo array must provide aligned child fields")
	}
}

func TestMSDecoderMonoFixedAllocationPointers(t *testing.T) {
	owners := make([]opusDecoderStorageViews, 32)
	for i := range owners {
		owners[i] = opusGoMonoStorage(nil)
		owners[i].Decoder.Fdecode_gain = int32(i + 1)
	}
	entropyInitGrowStack(12)
	runtime.GC()
	for i, owner := range owners {
		if uintptr(unsafe.Pointer(owner.Decoder))&15 != 0 || owner.Decoder.Fdecode_gain != int32(i+1) {
			t.Fatal("mono alignment/independent lifetime", i)
		}
		owner.Clear()
		if owner.Decoder.Fdecode_gain != 0 {
			t.Fatal("mono typed clear")
		}
	}
}

func TestMSDecoderStaticGoStoragePointers(t *testing.T) {
	for _, shape := range [][2]int32{{1, 0}, {1, 1}, {2, 0}, {2, 1}, {3, 2}} {
		streams, coupled := shape[0], shape[1]
		mapping := make([]byte, streams+coupled)
		for i := range mapping {
			mapping[i] = byte(i)
		}
		st, err := Opus_opus_multistream_decoder_create_typed(48000, int32(len(mapping)), streams, coupled, &mapping[0])
		if err != nil {
			t.Fatal(err)
		}
		flat := flatMSComparisonStorage(streams, coupled)
		if opusMSDecoderInitWithCodecs(nil, flat.Decoder, 48000, int32(len(mapping)), streams, coupled, &mapping[0], flat.Children, flat.Codecs) != 0 {
			t.Fatal("legacy oracle init")
		}
		if st.Flayout != flat.Decoder.Flayout || len(st.codecs) != int(streams) {
			t.Fatal("ordinary Go owner")
		}
		entropyInitGrowStack(12)
		runtime.GC()
		for i, codec := range st.codecs {
			channels := int32(1)
			if int32(i) < coupled {
				channels = 2
			}
			if uintptr(unsafe.Pointer(codec.Decoder))&15 != 0 {
				t.Fatal("child alignment", shape, i)
			}
			n := Opus_opus_decoder_get_size(nil, channels)
			if !slices.Equal(unsafe.Slice((*byte)(unsafe.Pointer(codec.Decoder)), n), unsafe.Slice((*byte)(unsafe.Pointer(flat.Codecs[i].Decoder)), n)) {
				t.Fatal("static child ABI fields/extent", shape, i)
			}
			// Typed Clear must never erase the back reference outside the ABI body.
			if opusDecoderInitWithStorage(nil, codec.Decoder, 48000, channels, &codec) != 0 {
				t.Fatal("static child clear/reinit")
			}
		}
		for _, request := range []int32{OPUS_SET_GAIN_REQUEST, OPUS_SET_COMPLEXITY_REQUEST, OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST, OPUS_GET_GAIN_REQUEST, OPUS_GET_COMPLEXITY_REQUEST, OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST, OPUS_GET_FINAL_RANGE_REQUEST, OPUS_RESET_STATE} {
			var a, b int32
			var ua, ub uint32
			argsA, argsB := OpusDecoderCtlArgs{Value: 1, I32: &a, U32: &ua}, OpusDecoderCtlArgs{Value: 1, I32: &b, U32: &ub}
			if ra, rb := Opus_opus_multistream_decoder_ctl_typed(nil, st, request, argsA), Opus_opus_multistream_decoder_ctl_typed(nil, flat.Decoder, request, argsB); ra != rb || a != b || ua != ub {
				t.Fatal("Go/flat CTL parity", shape, request, ra, rb, a, b, ua, ub)
			}
		}
	}
}

func TestMSDecoderStaticGoLiveCountPointers(t *testing.T) {
	mapping := [3]byte{0, 1, 2}
	st, err := Opus_opus_multistream_decoder_create_typed(48000, 3, 2, 1, &mapping[0])
	if err != nil {
		t.Fatal(err)
	}
	// Store-zero must precede the live loop condition, even in normal Go storage.
	if Opus_opus_multistream_decoder_ctl_typed(nil, st, OPUS_GET_FINAL_RANGE_REQUEST, OpusDecoderCtlArgs{U32: (*uint32)(unsafe.Pointer(&st.Flayout.Fnb_streams))}) != 0 || st.Flayout.Fnb_streams != 0 {
		t.Fatal("live Go count alias")
	}
}
