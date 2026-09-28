package opusccenc

import "testing"

func TestLPTapsPointers(t *testing.T) {
	for ind := int32(0); ind < 5; ind++ {
		for _, factor := range []int32{-1, 0, 1, 32767, 32768, 32769, 65535} {
			var b [3]int32
			var a [2]int32
			silk_LP_interpolate_filter_taps(nil, &b, &a, ind, factor)
			for i, v := range b {
				want := Opus_silk_Transition_LP_B_Q28[ind][i]
				if ind < 4 && factor > 0 {
					want += int32((int64(Opus_silk_Transition_LP_B_Q28[ind+1][i]-want) * int64(factor)) >> 16)
				}
				if v != want {
					t.Fatalf("B ind=%d factor=%d got=%d want=%d", ind, factor, v, want)
				}
			}
			for i, v := range a {
				want := Opus_silk_Transition_LP_A_Q28[ind][i]
				if ind < 4 && factor > 0 {
					want += int32((int64(Opus_silk_Transition_LP_A_Q28[ind+1][i]-want) * int64(factor)) >> 16)
				}
				if v != want {
					t.Fatalf("A ind=%d factor=%d got=%d want=%d", ind, factor, v, want)
				}
			}
		}
	}
}

func TestLPCutoffPointers(t *testing.T) {
	for _, pos := range []int32{0, 1, 63, 64, 127, 128, 255, 256} {
		for _, mode := range []int32{-2, -1, 0, 1, 2} {
			state := OpusT_silk_LP_state{FIn_LP_State: [2]int32{1234, -5678}, Ftransition_frame_no: pos, Fmode: mode, Fsaved_fs_kHz: 16}
			before := state
			frame := [6]int16{111, -32768, 32767, -1, 1, 222}
			original := frame
			Opus_silk_LP_variable_cutoff(nil, &state, &frame[1], 4)
			if mode == 0 && (state != before || frame != original) {
				t.Fatal("disabled filter changed data")
			}
			if state.Ftransition_frame_no != min(max(pos+mode, 0), 256) || state.Fmode != mode || state.Fsaved_fs_kHz != 16 || frame[0] != 111 || frame[5] != 222 {
				t.Fatal("state or sentinels differ")
			}
			// No samples means no IIR updates, but active transitions still advance.
			saved := state.FIn_LP_State
			Opus_silk_LP_variable_cutoff(nil, &state, nil, 0)
			if state.FIn_LP_State != saved {
				t.Fatal("empty frame changed IIR state")
			}
		}
	}
}
