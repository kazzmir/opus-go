package opuscc

import (
	"slices"
	"testing"
)

func TestUp2WrapperPointers(t *testing.T) {
	state := OpusT_silk_resampler_state_struct{FsIIR: [6]int32{1, -2, 3, -4, 5, -6}, FCoefs: 123, FinputDelay: 7}
	state.FdelayBuf[95] = -123
	state.FsFIR.Fi32[35] = 456
	before := state
	Opus_silk_resampler_private_up2_HQ_wrapper(nil, &state, nil, nil, 0)
	if state != before {
		t.Fatal("empty call changed state")
	}
	in := []int16{-32768, 32767, 0, -1, 1, 20000, -20000}
	out := make([]int16, 2*len(in)+2)
	out[0] = 123
	out[len(out)-1] = 456
	want := make([]int16, 2*len(in))
	iir := state.FsIIR
	Opus_silk_resampler_private_up2_HQ(nil, &iir, &want[0], &in[0], int32(len(in)))
	Opus_silk_resampler_private_up2_HQ_wrapper(nil, &state, &out[1], &in[0], int32(len(in)))
	before.FsIIR = iir
	if state != before || !slices.Equal(out[1:len(out)-1], want) || out[0] != 123 || out[len(out)-1] != 456 {
		t.Fatal("wrapper output, state, or guard differs")
	}
}
