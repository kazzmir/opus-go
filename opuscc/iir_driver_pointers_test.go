package opuscc

import (
	"slices"
	"testing"
	"unsafe"
)

func TestIIRHistoryWordsPointers(t *testing.T) {
	var state OpusT_silk_resampler_state_struct
	state.FsFIR.Fi32[0] = -2147483648
	state.FsFIR.Fi32[1] = -1
	state.FsFIR.Fi32[2] = 0x12345678
	state.FsFIR.Fi32[3] = -0x1234567
	state.FsFIR.Fi32[4] = 77
	before := state
	want := *(*[8]int16)(unsafe.Pointer(&state.FsFIR.Fi32[0]))
	if got := silkResamplerIIRHistory(&state); got != want || state != before {
		t.Fatal("native-endian signed history", got, want)
	}
	values := [8]int16{-32768, -1, 0, 1, 32767, -12345, 23456, -2}
	wantState := state
	copy((*[8]int16)(unsafe.Pointer(&wantState.FsFIR.Fi32[0]))[:], values[:])
	silkResamplerStoreIIRHistory(&state, values[:])
	if state != wantState || silkResamplerIIRHistory(&state) != values {
		t.Fatal("native-endian store/tail/roundtrip")
	}
}

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
	exactState := original
	exact := make([]int16, len(want))
	Opus_silk_resampler_private_IIR_FIR(nil, &exactState, &exact[0], &in[0], 161)
	if exactState != s || !slices.Equal(exact, want) {
		t.Fatal("exact output extent/state")
	}
	if !slices.Equal(s.FsFIR.Fi32[4:], original.FsFIR.Fi32[4:]) {
		t.Fatal("union tail modified")
	}
}
