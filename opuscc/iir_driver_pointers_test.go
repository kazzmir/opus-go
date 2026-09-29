package opuscc

import (
	"slices"
	"testing"
	"unsafe"
)

func TestIIRDriverPointers(t *testing.T) {
	s := OpusT_silk_resampler_state_struct{FbatchSize: 80, FinvRatio_Q16: 65536}
	for i := range s.FsFIR.Fi32 {
		s.FsFIR.Fi32[i] = int32(i * 71)
	}
	original := s
	Opus_silk_resampler_private_IIR_FIR(nil, &s, nil, nil, 0)
	if s != original {
		t.Fatal("empty mutated state")
	}
	ref := s
	in := make([]int16, 161)
	for i := range in {
		in[i] = int16(i * 537)
	}
	out := make([]int16, 324)
	out[0] = 123
	out[323] = 456
	want := make([]int16, 322)
	Opus_silk_resampler_private_IIR_FIR(nil, &s, &out[1], &in[0], 161)
	for offset := 0; offset < len(in); {
		n := min(80, len(in)-offset)
		Opus_silk_resampler_private_IIR_FIR(nil, &ref, &want[2*offset], unsafe.SliceData(in[offset:]), int32(n))
		offset += n
	}
	if s != ref || !slices.Equal(out[1:323], want) || out[0] != 123 || out[323] != 456 {
		t.Fatal("batch/state/guards")
	}
	if !slices.Equal(s.FsFIR.Fi32[4:], original.FsFIR.Fi32[4:]) {
		t.Fatal("union tail modified")
	}
}
