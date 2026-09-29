package opuscc

import (
	"slices"
	"testing"
)

func TestDownDriverPointers(t *testing.T) {
	coefs := Opus_silk_Resampler_1_2_COEFS // Local owner; FCoefs remains zero.
	s := OpusT_silk_resampler_state_struct{FbatchSize: 80, FinvRatio_Q16: 131072, FFIR_Order: 24, FFIR_Fracs: 1}
	for i := range s.FsFIR.Fi32 {
		s.FsFIR.Fi32[i] = int32(i * 71)
	}
	original := s
	Opus_silk_resampler_private_down_FIR(nil, &s, &coefs[0], nil, nil, 0)
	if s != original {
		t.Fatal("empty mutated state")
	}
	ref := s
	in := make([]int16, 161)
	for i := range in {
		in[i] = int16(i * 537)
	}
	out := make([]int16, 82)
	out[0] = 123
	out[81] = 456
	want := make([]int16, 80)
	Opus_silk_resampler_private_down_FIR(nil, &s, &coefs[0], &out[1], &in[0], 161)
	Opus_silk_resampler_private_down_FIR(nil, &ref, &coefs[0], &want[0], &in[0], 80)
	Opus_silk_resampler_private_down_FIR(nil, &ref, &coefs[0], &want[40], &in[80], 80)
	if s != ref || !slices.Equal(out[1:81], want) || out[0] != 123 || out[81] != 456 {
		t.Fatal("batch/remainder/state/guards")
	}
	if !slices.Equal(s.FsFIR.Fi32[24:], original.FsFIR.Fi32[24:]) {
		t.Fatal("history tail modified")
	}
}
