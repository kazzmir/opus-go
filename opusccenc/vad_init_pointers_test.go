package opusccenc

import "testing"

func TestVADInitPointers(t *testing.T) {
	state := OpusT_silk_VAD_state{FAnaState: [2]int32{1, 2}, FAnaState1: [2]int32{3, 4}, FAnaState2: [2]int32{5, 6}, FXnrgSubfr: [4]int32{7, 8, 9, 10}, FNrgRatioSmth_Q8: [4]int32{11, 12, 13, 14}, FHPstate: 15, FNL: [4]int32{16, 17, 18, 19}, Finv_NL: [4]int32{20, 21, 22, 23}, FNoiseLevelBias: [4]int32{24, 25, 26, 27}, Fcounter: 28}
	want := OpusT_silk_VAD_state{FNoiseLevelBias: [4]int32{50, 25, 16, 12}, FNL: [4]int32{5000, 2500, 1600, 1200}, Finv_NL: [4]int32{429496, 858993, 1342177, 1789569}, FNrgRatioSmth_Q8: [4]int32{25600, 25600, 25600, 25600}, Fcounter: 15}
	for pass := 0; pass < 2; pass++ {
		if ret := Opus_silk_VAD_Init(nil, &state); ret != 0 || state != want {
			t.Fatalf("state=%+v return=%d", state, ret)
		}
	}
}
