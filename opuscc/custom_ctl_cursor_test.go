package opuscc

import (
	"runtime"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestCustomCtlCursorPointers(t *testing.T) {
	parent := newOpusCtlTestStorage()
	state := &parent.Celt.State
	heapMode := new(OpusT_OpusCustomMode)
	*heapMode = mode48000_960_120
	state.Fmode = heapMode
	heapMode = nil
	value := int(9)
	scalar := libc.NewVaCursor(unsafe.Pointer(&value))
	var out int32
	outputSlot := &out
	output := libc.NewVaCursor(unsafe.Pointer(&outputSlot))
	entropyInitGrowStack(12)
	runtime.GC()
	if Opus_opus_custom_decoder_ctl(nil, state, OPUS_SET_COMPLEXITY_REQUEST, scalar) != 0 || state.Fcomplexity != 9 {
		t.Fatal("pointer-backed setter")
	}
	if Opus_opus_custom_decoder_ctl(nil, state, OPUS_GET_COMPLEXITY_REQUEST, output) != 0 || out != 9 {
		t.Fatal("pointer-backed getter")
	}
	value = 11
	if Opus_opus_custom_decoder_ctl(nil, state, OPUS_SET_COMPLEXITY_REQUEST, scalar) != -1 || state.Fcomplexity != 9 {
		t.Fatal("invalid complexity before stores")
	}
	value = 3
	if Opus_opus_custom_decoder_ctl(nil, state, OPUS_SET_COMPLEXITY_REQUEST, scalar) != 0 || Opus_opus_custom_decoder_ctl(nil, state, OPUS_GET_COMPLEXITY_REQUEST, output) != 0 || out != 3 {
		t.Fatal("cursor copied per call/live scalar reload")
	}
	var nilSlot *int32
	if Opus_opus_custom_decoder_ctl(nil, state, OPUS_GET_COMPLEXITY_REQUEST, libc.NewVaCursor(unsafe.Pointer(&nilSlot))) != -1 {
		t.Fatal("null output")
	}
	if Opus_opus_custom_decoder_ctl(nil, state, -1234, libc.VaCursor{}) != -5 {
		t.Fatal("unknown request does not consume arguments")
	}
	state.Ferror1 = 37
	errorSlot := &state.Ferror1
	if Opus_opus_custom_decoder_ctl(nil, state, CELT_GET_AND_CLEAR_ERROR_REQUEST, libc.NewVaCursor(unsafe.Pointer(&errorSlot))) != 0 || state.Ferror1 != 0 {
		t.Fatal("aliased error output keeps store/clear order")
	}
	state.Frng = 0xfedcba98
	var rng uint32
	rangeSlot := &rng
	if Opus_opus_custom_decoder_ctl(nil, state, OPUS_GET_FINAL_RANGE_REQUEST, libc.NewVaCursor(unsafe.Pointer(&rangeSlot))) != 0 || rng != 0xfedcba98 {
		t.Fatal("unsigned range slot")
	}
	var mode *OpusT_OpusCustomMode
	modeSlot := &mode
	if Opus_opus_custom_decoder_ctl(nil, state, CELT_GET_MODE_REQUEST, libc.NewVaCursor(unsafe.Pointer(&modeSlot))) != 0 || mode != state.Fmode {
		t.Fatal("scanned mode output slot")
	}
	if Opus_opus_custom_decoder_ctl(nil, state, OPUS_RESET_STATE, libc.VaCursor{}) != 0 || state.Fskip_plc != 1 || state.Fcomplexity != 3 {
		t.Fatal("reset with no varargs")
	}
	parent = nil
	state = nil
	runtime.GC()
	if mode == nil || mode.FnbEBands != 21 {
		t.Fatal("mode output retains actual Go owner")
	}
}
