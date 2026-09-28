package opuscc

import "testing"

func TestNLSFStabilizePointers(t *testing.T) {
	for _, input := range [][4]int16{{1000, 2000, 3000, 4000}, {0, 0, 0, 0}, {32767, 32767, 32767, 32767}, {30000, 20000, 10000, -32768}} {
		values := [6]int16{111, input[0], input[1], input[2], input[3], 222}
		delta := [5]int16{50, 100, 200, 300, 400}
		before := delta
		Opus_silk_NLSF_stabilize(nil, &values[1], &delta[0], 4)
		if values[0] != 111 || values[5] != 222 || delta != before || values[1] < 50 || values[4] > 32768-400 {
			t.Fatal("bounds or sentinels changed")
		}
		for i := 1; i < 4; i++ {
			if int32(values[i+1])-int32(values[i]) < int32(delta[i]) {
				t.Fatal("spacing violated")
			}
		}
		stable := values
		Opus_silk_NLSF_stabilize(nil, &values[1], &delta[0], 4)
		if values != stable {
			t.Fatal("stable vector changed")
		}
	}
}
