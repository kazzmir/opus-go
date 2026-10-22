package opuscc

import (
	"slices"
	"testing"
	"unsafe"
)

// From iir_driver_pointers_test.go.
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

// From iir_interpol_pointers_test.go.
func TestIIRInterpolPointers(t *testing.T) {
	if silk_resampler_private_IIR_FIR_INTERPOL(nil, nil, nil, 0, 1) != 0 {
		t.Fatal("empty")
	}
	in := [8]int16{32767, -32768, 17, -45, 32000, -31500, 99, 71}
	out := make([]int16, 65538)
	out[0] = 123
	out[len(out)-1] = 456
	if n := silk_resampler_private_IIR_FIR_INTERPOL(nil, &out[1], &in[0], 65536, 1); n != 65536 {
		t.Fatal(n)
	}
	for i := int32(0); i < 65536; i++ {
		phase := i * 12 >> 16
		var sum int32
		for j := 0; j < 8; j++ {
			c := int16(0)
			if j < 4 {
				c = Opus_silk_resampler_frac_FIR_12[phase][j]
			} else {
				c = Opus_silk_resampler_frac_FIR_12[11-phase][7-j]
			}
			sum += int32(in[j]) * int32(c)
		}
		want := int16(max(-32768, min(32767, ((sum>>14)+1)>>1)))
		if out[i+1] != want {
			t.Fatalf("phase %d", i)
		}
	}
	if out[0] != 123 || out[len(out)-1] != 456 {
		t.Fatal("guard")
	}
}

// From rotation_resampler_pointers_test.go.
func TestExpRotationPointers(t *testing.T) {
	for _, tc := range []struct {
		stride int32
		want   [6]float32
	}{
		{1, [6]float32{-6, -2, -3, -4, -5, 1}},
		{2, [6]float32{5, 6, -3, -4, 1, 2}},
		{3, [6]float32{-4, -5, -6, 1, 2, 3}},
		{6, [6]float32{1, 2, 3, 4, 5, 6}},
	} {
		x := [8]float32{99, 1, 2, 3, 4, 5, 6, 88}
		exp_rotation1(nil, &x[1], 6, tc.stride, 0, 1)
		if [6]float32(x[1:7]) != tc.want || x[0] != 99 || x[7] != 88 {
			t.Fatalf("stride %d: got %v, want %v with sentinels", tc.stride, x, tc.want)
		}
	}
	exp_rotation1(nil, nil, 0, 1, 0, 1)
}

func TestHysteresisPointers(t *testing.T) {
	thresholds := [2]float32{10, 20}
	margins := [2]float32{2, 3}
	for _, tc := range []struct {
		value      float32
		prev, want int32
	}{
		{5, 0, 0}, {11, 0, 0}, {12, 0, 1},
		{9, 1, 1}, {8, 1, 0}, {22, 1, 1}, {23, 1, 2},
		{18, 2, 2}, {17, 2, 1}, {30, 0, 2}, {0, 2, 0},
	} {
		if got := Opus_hysteresis_decision(nil, tc.value, &thresholds[0], &margins[0], 2, tc.prev); got != tc.want {
			t.Fatalf("value %v prev %d: got %d, want %d", tc.value, tc.prev, got, tc.want)
		}
	}
	if got := Opus_hysteresis_decision(nil, 0, nil, nil, 0, 0); got != 0 {
		t.Fatalf("empty thresholds: %d", got)
	}
}

func TestResamplerAR2Pointers(t *testing.T) {
	input := [...]int16{32767, -32768, 0, 1, -1, 12000, -16000}
	coefs := [2]int16{-12345, 23456}
	initial := [2]int32{123456, -98765}
	want := [7]int32{8511808, -14900842, 23413321, -38973841, 62885175, -100107195, 161361640}
	wantState := [2]int32{-264900136, 231011879}
	// Chunked processing must produce exactly the same state and samples.
	for _, split := range []int{0, 1, 3, 6} {
		state := initial
		out := [9]int32{123, 0, 0, 0, 0, 0, 0, 0, 456}
		Opus_silk_resampler_private_AR2(nil, &state[0], &out[1], &input[0], &coefs[0], int32(split))
		Opus_silk_resampler_private_AR2(nil, &state[0], &out[split+1], &input[split], &coefs[0], int32(len(input)-split))
		if [7]int32(out[1:8]) != want || state != wantState || out[0] != 123 || out[8] != 456 {
			t.Fatalf("split %d: out=%v state=%v", split, out, state)
		}
	}
	Opus_silk_resampler_private_AR2(nil, nil, nil, nil, nil, 0)
}

// From biquad1_pointers_test.go.
func TestBiquad1Pointers(t *testing.T) {
	input := []int16{-32768, 32767, 0, 1, -1, 100, -200, 300}
	var a [2]int32
	for _, tc := range []struct {
		b    [3]int32
		want []int16
	}{
		{[3]int32{1 << 28, 0, 0}, input},
		{[3]int32{0, 1 << 28, 0}, []int16{0, -32768, 32767, 0, 1, -1, 100, -200}},
		{[3]int32{1 << 29, 0, 0}, []int16{-32768, 32767, 0, 2, -2, 200, -400, 600}},
	} {
		var state [2]int32
		out := make([]int16, len(input)+2)
		out[0], out[len(out)-1] = 111, 222
		Opus_silk_biquad_alt_stride1(nil, &input[0], &tc.b, &a, &state, &out[1], int32(len(input)))
		if !slices.Equal(out[1:len(out)-1], tc.want) || out[0] != 111 || out[len(out)-1] != 222 {
			t.Fatalf("b=%v output=%v want=%v", tc.b, out, tc.want)
		}
		inPlace := slices.Clone(input)
		var splitState [2]int32
		Opus_silk_biquad_alt_stride1(nil, &inPlace[0], &tc.b, &a, &splitState, &inPlace[0], 4)
		Opus_silk_biquad_alt_stride1(nil, &inPlace[4], &tc.b, &a, &splitState, &inPlace[4], 4)
		if !slices.Equal(inPlace, tc.want) || splitState != state {
			t.Fatal("in-place/chunked processing differs")
		}
	}
	state := [2]int32{12, 34}
	b := [3]int32{1 << 28}
	Opus_silk_biquad_alt_stride1(nil, nil, &b, &a, &state, nil, 0)
	if state != [2]int32{12, 34} {
		t.Fatal("empty input changes state")
	}
}

// From biquad2_pointers_test.go.
func TestBiquad2Pointers(t *testing.T) {
	b := [3]int32{1 << 27, 1 << 26, -(1 << 25)}
	a := [2]int32{-(1 << 26) + 123, (1 << 24) + 456}
	for _, n := range []int{2, 16, 64} {
		input := make([]int16, 2*n)
		for i := range input {
			input[i] = int16(32767 - i*1700)
		}
		initial := [4]int32{10000, -20000, -30000, 40000}
		state := initial
		out := make([]int16, len(input)+2)
		out[0], out[len(out)-1] = 111, 222
		Opus_silk_biquad_alt_stride2_c(nil, &input[0], &b, &a, &state, &out[1], int32(n))
		for channel := 0; channel < 2; channel++ {
			mono := make([]int16, n)
			for i := range mono {
				mono[i] = input[2*i+channel]
			}
			monoState := [2]int32{initial[2*channel], initial[2*channel+1]}
			Opus_silk_biquad_alt_stride1(nil, &mono[0], &b, &a, &monoState, &mono[0], int32(n))
			for i := range mono {
				if mono[i] != out[1+2*i+channel] {
					t.Fatalf("n=%d channel=%d sample=%d: interleaved output differs", n, channel, i)
				}
			}
			if monoState != [2]int32{state[2*channel], state[2*channel+1]} {
				t.Fatal("channel state differs")
			}
		}
		if out[0] != 111 || out[len(out)-1] != 222 {
			t.Fatal("sentinel changed")
		}
		inPlace := slices.Clone(input)
		splitState := initial
		Opus_silk_biquad_alt_stride2_c(nil, &inPlace[0], &b, &a, &splitState, &inPlace[0], int32(n/2))
		Opus_silk_biquad_alt_stride2_c(nil, &inPlace[n], &b, &a, &splitState, &inPlace[n], int32(n/2))
		if !slices.Equal(inPlace, out[1:len(out)-1]) || splitState != state {
			t.Fatal("in-place/chunked processing differs")
		}
	}
	state := [4]int32{1, 2, 3, 4}
	Opus_silk_biquad_alt_stride2_c(nil, nil, &b, &a, &state, nil, 0)
	if state != [4]int32{1, 2, 3, 4} {
		t.Fatal("empty input changed state")
	}
}

// From lp_taps_pointers_test.go.
func TestLPTapsPointers(t *testing.T) {
	// The full Q16 interpolation range includes both signed-multiplier branches.
	for ind := int32(0); ind < 5; ind++ {
		for factor := int32(0); factor <= 65536; factor++ {
			b := struct {
				before int32
				taps   [3]int32
				after  int32
			}{before: 111, after: 222}
			a := struct {
				before int32
				taps   [2]int32
				after  int32
			}{before: 333, after: 444}
			silk_LP_interpolate_filter_taps(nil, &b.taps, &a.taps, ind, factor)
			for i, got := range b.taps {
				want := int64(Opus_silk_Transition_LP_B_Q28[ind][i])
				if ind < 4 {
					want += (int64(Opus_silk_Transition_LP_B_Q28[ind+1][i]) - want) * int64(factor) >> 16
				}
				if got != int32(want) {
					t.Fatalf("B ind=%d factor=%d tap=%d got=%d want=%d", ind, factor, i, got, want)
				}
			}
			for i, got := range a.taps {
				want := int64(Opus_silk_Transition_LP_A_Q28[ind][i])
				if ind < 4 {
					want += (int64(Opus_silk_Transition_LP_A_Q28[ind+1][i]) - want) * int64(factor) >> 16
				}
				if got != int32(want) {
					t.Fatalf("A ind=%d factor=%d tap=%d got=%d want=%d", ind, factor, i, got, want)
				}
			}
			if b.before != 111 || b.after != 222 || a.before != 333 || a.after != 444 {
				t.Fatal("sentinels changed")
			}
		}
	}
	state := OpusT_silk_LP_state{FIn_LP_State: [2]int32{123, -456}, Ftransition_frame_no: 123, Fsaved_fs_kHz: 16}
	before := state
	Opus_silk_LP_variable_cutoff(nil, &state, nil, 120)
	if state != before {
		t.Fatal("disabled filter changed state")
	}
}
