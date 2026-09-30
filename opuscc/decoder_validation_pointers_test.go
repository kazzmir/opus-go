package opuscc

import (
	"runtime"
	"testing"
)

func validationPanics(f func()) (panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	f()
	return false
}

func TestOpusValidationPointers(t *testing.T) {
	st := OpusT_OpusDecoder{Fchannels: 2, FFs: 48000, Fstream_channels: 1}
	st.FDecControl.FAPI_sampleRate = 48000
	st.FDecControl.FnChannelsAPI = 2
	before := st
	entropyInitGrowStack(12)
	runtime.GC()
	validate_opus_decoder(nil, &st)
	if st != before {
		t.Fatal("state changed")
	}
	for _, bad := range []int32{0, 3, -1} {
		st = before
		st.Fchannels = bad
		if !validationPanics(func() { validate_opus_decoder(nil, &st) }) {
			t.Fatal("channel accepted", bad)
		}
	}
	st = before
	st.Farch = -1
	if !validationPanics(func() { validate_opus_decoder(nil, &st) }) {
		t.Fatal("arch accepted")
	}
}
