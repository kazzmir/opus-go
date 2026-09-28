package opuscc

import "testing"

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
