package opuscc

import "testing"

func TestNoiseLevelsPointers(t *testing.T) {
	for _, tc := range []struct {
		counter        int32
		noise, inverse [4]int32
	}{
		{0, [4]int32{99, 2512, 2426, 2399}, [4]int32{21688935, 854740, 885165, 894812}},
		{15, [4]int32{99, 2512, 2426, 2399}, [4]int32{21688935, 854740, 885165, 894812}},
		{16, [4]int32{194, 2506, 1928, 1599}, [4]int32{11058891, 856866, 1113678, 1342204}},
		{999, [4]int32{1963, 2500, 1608, 1209}, [4]int32{1093873, 858861, 1334924, 1775369}},
		{1000, [4]int32{1963, 2500, 1605, 1202}, [4]int32{1093873, 858861, 1337630, 1786073}},
		{1001, [4]int32{1963, 2500, 1605, 1202}, [4]int32{1093873, 858861, 1337630, 1786073}},
	} {
		var state OpusT_silk_VAD_state
		Opus_silk_VAD_Init(nil, &state)
		state.Fcounter = tc.counter
		state.FAnaState = [2]int32{123, -456}
		state.FHPstate = 789
		want := state
		want.FNL, want.Finv_NL = tc.noise, tc.inverse
		if tc.counter < 1000 {
			want.Fcounter++
		}
		energies := [4]int32{0, 2500, 5000, 2147483647}
		before := energies
		silk_VAD_GetNoiseLevels(nil, &energies, &state)
		if state != want || energies != before {
			t.Fatalf("counter=%d got=%+v want=%+v", tc.counter, state, want)
		}
	}
}
