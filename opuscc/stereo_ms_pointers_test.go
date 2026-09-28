package opuscc

import "testing"

func TestStereoMSPointers(t *testing.T) {
	const n = 80
	mid, side := make([]int16, n+4), make([]int16, n+4)
	mid[0], mid[n+3] = 111, 222
	side[0], side[n+3] = 333, 444
	for i := 2; i < n+2; i++ {
		if i%2 == 0 {
			mid[i+1], side[i+1] = 32767, 32767
		} else {
			mid[i+1], side[i+1] = -32768, -32768
		}
	}
	state := OpusT_stereo_dec_state{FsMid: [2]int16{200, -300}, FsSide: [2]int16{-100, 400}}
	var predictors [2]int32
	wantMid, wantSide := [2]int16{mid[n+1], mid[n+2]}, [2]int16{side[n+1], side[n+2]}
	Opus_silk_stereo_MS_to_LR(nil, &state, &mid[1], &side[1], &predictors, 8, n)
	if mid[0] != 111 || mid[n+3] != 222 || side[0] != 333 || side[n+3] != 444 {
		t.Fatal("sentinels changed")
	}
	if mid[1] != 200 || side[1] != -100 || mid[2] != 100 || side[2] != -700 {
		t.Fatal("history buffering differs")
	}
	for i := 2; i <= n; i++ {
		want := int16(-32768)
		if i%2 == 0 {
			want = 32767
		}
		if mid[i+1] != want || side[i+1] != 0 {
			t.Fatalf("sample=%d left=%d right=%d", i, mid[i+1], side[i+1])
		}
	}
	if state.FsMid != wantMid || state.FsSide != wantSide || state.Fpred_prev_Q13 != [2]int16{} {
		t.Fatalf("state=%+v", state)
	}
}
