package opuscc

import (
	"runtime"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestCommonCtlCursorPointers(t *testing.T) {
	for kind := 0; kind < 4; kind++ {
		parent := new(projectionCtlTestStorage)
		parent.MS.State.Flayout.Fnb_streams = 1
		parent.MS.State.Flayout.Fnb_channels = 1
		parent.MS.Decoder = *newOpusCtlTestStorage()
		decoder := &parent.MS.Decoder.State
		call := func(request int32, ap libc.VaCursor) int32 {
			switch kind {
			case 0:
				return Opus_opus_decoder_ctl(nil, decoder, request, ap)
			case 1:
				return Opus_opus_multistream_decoder_ctl(nil, &parent.MS.State, request, ap)
			case 2:
				return Opus_opus_multistream_decoder_ctl_va_list(nil, &parent.MS.State, request, ap)
			default:
				return Opus_opus_projection_decoder_ctl(nil, &parent.State, request, ap)
			}
		}
		value := int(-19)
		scalar := libc.NewVaCursor(unsafe.Pointer(&value))
		var out int32
		outputSlot := &out
		output := libc.NewVaCursor(unsafe.Pointer(&outputSlot))
		entropyInitGrowStack(12)
		runtime.GC()
		if call(OPUS_SET_GAIN_REQUEST, scalar) != 0 || decoder.Fdecode_gain != -19 {
			t.Fatal("pointer-backed setter", kind)
		}
		if call(OPUS_GET_GAIN_REQUEST, output) != 0 || out != -19 {
			t.Fatal("pointer-backed getter", kind)
		}
		value = -32769
		if call(OPUS_SET_GAIN_REQUEST, scalar) != -1 || decoder.Fdecode_gain != -19 {
			t.Fatal("validation before stores/cursor copied per call", kind)
		}
		value = 27
		if call(OPUS_SET_GAIN_REQUEST, scalar) != 0 || call(OPUS_GET_GAIN_REQUEST, output) != 0 || out != 27 {
			t.Fatal("reused cursor reloads live word", kind)
		}
		var nilSlot *int32
		if call(OPUS_GET_GAIN_REQUEST, libc.NewVaCursor(unsafe.Pointer(&nilSlot))) != -1 {
			t.Fatal("null output", kind)
		}
		if call(-1234, libc.VaCursor{}) != -5 {
			t.Fatal("unknown request consumes no arguments", kind)
		}
		decoder.FrangeFinal = 0xfedcba98
		rangeSlot := &decoder.FrangeFinal
		if call(OPUS_GET_FINAL_RANGE_REQUEST, libc.NewVaCursor(unsafe.Pointer(&rangeSlot))) != 0 {
			t.Fatal("range output", kind)
		}
		want := uint32(0xfedcba98)
		if kind != 0 {
			want = 0
		}
		if decoder.FrangeFinal != want {
			t.Fatal("live aliased range: multistream clears before child query", kind)
		}
		if kind != 0 {
			// Exactly one physical word, invalid stream: the second slot is absent.
			stream := int(-1)
			if call(OPUS_MULTISTREAM_GET_DECODER_STATE_REQUEST, libc.NewVaCursor(unsafe.Pointer(&stream))) != -1 {
				t.Fatal("invalid stream before output", kind)
			}
			var child *OpusT_OpusDecoder
			slots := struct {
				Stream int
				Output **OpusT_OpusDecoder
			}{0, &child}
			if call(OPUS_MULTISTREAM_GET_DECODER_STATE_REQUEST, libc.NewVaCursor(unsafe.Pointer(&slots))) != 0 || child != decoder {
				t.Fatal("scanned child output slot", kind)
			}
		}
		if call(OPUS_RESET_STATE, libc.VaCursor{}) != 0 || decoder.Fframe_size != 120 || decoder.Fdecode_gain != 27 {
			t.Fatal("reset with no varargs preserves gain", kind)
		}
	}
}

func TestCommonDREDCtlCursorPointers(t *testing.T) {
	for _, request := range []int32{-1234, OPUS_RESET_STATE, OPUS_GET_FINAL_RANGE_REQUEST} {
		if Opus_opus_dred_decoder_ctl(nil, nil, request, libc.VaCursor{}) != -5 {
			t.Fatal("disabled DRED never consumes arguments")
		}
	}
}
