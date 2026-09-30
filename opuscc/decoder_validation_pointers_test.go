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

func TestMSValidationPointers(t *testing.T) {
	st := OpusT_OpusMSDecoder{Flayout: OpusT_ChannelLayout{Fnb_channels: 2, Fnb_streams: 1, Fnb_coupled_streams: 1, Fmapping: [256]uint8{0, 1}}}
	before := st
	entropyInitGrowStack(12)
	runtime.GC()
	validate_ms_decoder(nil, &st)
	if st != before {
		t.Fatal("state changed")
	}
	// The original validator intentionally ignores an invalid layout return value.
	st.Flayout.Fmapping[0] = 7
	before = st
	validate_ms_decoder(nil, &st)
	if st != before || Opus_validate_layout(nil, &st.Flayout) != 0 {
		t.Fatal("invalid-layout behavior changed")
	}
}

func TestCeltValidationPointers(t *testing.T) {
	mode, _ := Opus_opus_custom_mode_create(nil, 48000, 960)
	st := OpusT_OpusCustomDecoder{Fmode: mode, Foverlap: 120, Fchannels: 2, Fstream_channels: 1, Fdownsample: 1, Fend: 21}
	before := st
	entropyInitGrowStack(12)
	runtime.GC()
	Opus_validate_celt_decoder(nil, &st)
	if st != before {
		t.Fatal("state changed")
	}
	for _, pitch := range []int32{0, PLC_PITCH_LAG_MIN, PLC_PITCH_LAG_MAX} {
		st = before
		st.Flast_pitch_index = pitch
		Opus_validate_celt_decoder(nil, &st)
	}
	st = before
	st.Fpostfilter_period = MAX_PERIOD
	if !validationPanics(func() { Opus_validate_celt_decoder(nil, &st) }) {
		t.Fatal("period accepted")
	}
	st = before
	st.Fmode = 0
	if !validationPanics(func() { Opus_validate_celt_decoder(nil, &st) }) {
		t.Fatal("mode accepted")
	}
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
