package opuscc

import (
	"runtime"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

type ctlCursorDecoder func(int32, func() libc.VaCursor) OpusDecoderCtlArgs

func TestCtlLegacyCursorUnusedPointers(t *testing.T) {
	unused := func() libc.VaCursor { t.Fatal("unused varargs opened"); return libc.VaCursor{} }
	for _, request := range []int32{OPUS_RESET_STATE, -1234} {
		if opusCtlLegacyArgs(request, unused) != (OpusDecoderCtlArgs{}) || customCtlLegacyArgs(request, unused) != (OpusDecoderCtlArgs{}) || msCtlLegacyArgs(nil, request, unused) != (OpusDecoderCtlArgs{}) {
			t.Fatal("unused arguments")
		}
	}
	// These requests are deliberately not forwarded by multistream dispatch.
	for _, request := range []int32{OPUS_GET_PITCH_REQUEST, OPUS_SET_IGNORE_EXTENSIONS_REQUEST, CELT_GET_MODE_REQUEST} {
		if msCtlLegacyArgs(nil, request, unused) != (OpusDecoderCtlArgs{}) {
			t.Fatal("unsupported multistream arguments")
		}
	}
}

func TestCtlLegacyCursorScalarPointers(t *testing.T) {
	for _, tc := range []struct {
		decode   ctlCursorDecoder
		requests []int32
	}{
		{opusCtlLegacyArgs, []int32{OPUS_SET_COMPLEXITY_REQUEST, OPUS_SET_GAIN_REQUEST, OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST, OPUS_SET_IGNORE_EXTENSIONS_REQUEST}},
		{customCtlLegacyArgs, []int32{OPUS_SET_COMPLEXITY_REQUEST, CELT_SET_START_BAND_REQUEST, CELT_SET_END_BAND_REQUEST, CELT_SET_CHANNELS_REQUEST, CELT_SET_SIGNALLING_REQUEST, OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST}},
		{func(r int32, ap func() libc.VaCursor) OpusDecoderCtlArgs { return msCtlLegacyArgs(nil, r, ap) }, []int32{OPUS_SET_GAIN_REQUEST, OPUS_SET_COMPLEXITY_REQUEST, OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST}},
	} {
		for _, request := range tc.requests {
			for _, value := range []int32{-2147483648, -7, 0, 2147483647} {
				word := int(value)
				opened := 0
				ap := func() libc.VaCursor { opened++; return libc.NewVaCursor(unsafe.Pointer(&word)) }
				entropyInitGrowStack(12)
				runtime.GC()
				if got := tc.decode(request, ap); got != (OpusDecoderCtlArgs{Value: value}) || opened != 1 {
					t.Fatal("one scalar cursor", request, value, got, opened)
				}
			}
		}
	}
}

func TestCtlLegacyCursorOutputPointers(t *testing.T) {
	for _, tc := range []struct {
		decode   ctlCursorDecoder
		requests []int32
	}{
		{opusCtlLegacyArgs, []int32{OPUS_GET_BANDWIDTH_REQUEST, OPUS_GET_COMPLEXITY_REQUEST, OPUS_GET_SAMPLE_RATE_REQUEST, OPUS_GET_PITCH_REQUEST, OPUS_GET_GAIN_REQUEST, OPUS_GET_LAST_PACKET_DURATION_REQUEST, OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST, OPUS_GET_IGNORE_EXTENSIONS_REQUEST}},
		{customCtlLegacyArgs, []int32{OPUS_GET_COMPLEXITY_REQUEST, CELT_GET_AND_CLEAR_ERROR_REQUEST, OPUS_GET_LOOKAHEAD_REQUEST, OPUS_GET_PITCH_REQUEST, OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST}},
		{func(r int32, ap func() libc.VaCursor) OpusDecoderCtlArgs { return msCtlLegacyArgs(nil, r, ap) }, []int32{OPUS_GET_BANDWIDTH_REQUEST, OPUS_GET_SAMPLE_RATE_REQUEST, OPUS_GET_GAIN_REQUEST, OPUS_GET_LAST_PACKET_DURATION_REQUEST, OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST, OPUS_GET_COMPLEXITY_REQUEST}},
	} {
		for _, request := range tc.requests {
			for _, present := range []bool{false, true} {
				value := int32(37)
				var slot *int32
				if present {
					slot = &value
				}
				opened := 0
				ap := func() libc.VaCursor { opened++; return libc.NewVaCursor(unsafe.Pointer(&slot)) }
				runtime.GC()
				if got := tc.decode(request, ap); got != (OpusDecoderCtlArgs{I32: slot}) || opened != 1 {
					t.Fatal("typed output slot", request, got, opened)
				}
			}
		}
		rng := uint32(0xfedcba98)
		slot := &rng
		ap := func() libc.VaCursor { return libc.NewVaCursor(unsafe.Pointer(&slot)) }
		if got := tc.decode(OPUS_GET_FINAL_RANGE_REQUEST, ap); got != (OpusDecoderCtlArgs{U32: &rng}) {
			t.Fatal("unsigned output slot", got)
		}
	}
	var mode *OpusT_OpusCustomMode
	slot := &mode
	ap := func() libc.VaCursor { return libc.NewVaCursor(unsafe.Pointer(&slot)) }
	if got := customCtlLegacyArgs(CELT_GET_MODE_REQUEST, ap); got != (OpusDecoderCtlArgs{Mode: &mode}) {
		t.Fatal("mode output slot", got)
	}
}

func TestCtlLegacyCursorStreamPointers(t *testing.T) {
	state := OpusT_OpusMSDecoder{}
	state.Flayout.Fnb_streams = 2
	// Only one physical native word: invalid IDs MUST NOT read a second slot.
	for _, value := range []int{-1, 2, 2147483647} {
		ap := func() libc.VaCursor { return libc.NewVaCursor(unsafe.Pointer(&value)) }
		if got := msCtlLegacyArgs(&state, OPUS_MULTISTREAM_GET_DECODER_STATE_REQUEST, ap); got != (OpusDecoderCtlArgs{Value: int32(value)}) {
			t.Fatal("invalid stream", got)
		}
	}
	for _, present := range []bool{false, true} {
		var decoder *OpusT_OpusDecoder
		slots := struct {
			Value  int
			Output **OpusT_OpusDecoder
		}{Value: 1}
		if present {
			slots.Output = &decoder
		}
		opened := 0
		ap := func() libc.VaCursor { opened++; return libc.NewVaCursor(unsafe.Pointer(&slots)) }
		entropyInitGrowStack(12)
		runtime.GC()
		if got := msCtlLegacyArgs(&state, OPUS_MULTISTREAM_GET_DECODER_STATE_REQUEST, ap); got != (OpusDecoderCtlArgs{Value: 1, Decoder: slots.Output}) || opened != 1 {
			t.Fatal("valid stream pointer slot", got, opened)
		}
	}
}
