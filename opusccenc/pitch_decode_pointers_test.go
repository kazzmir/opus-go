package opusccenc

import "testing"

func TestPitchDecodePointers(t *testing.T) {
	for _, fs := range []int32{8, 12, 16} {
		for _, n := range []int32{2, 4} {
			count := 34
			if n == 2 {
				count = 12
			}
			if fs == 8 {
				count = 11
				if n == 2 {
					count = 3
				}
			}
			for contour := 0; contour < count; contour++ {
				for _, lag := range []int16{-32768, 0, 1, 60, 32767} {
					out := [6]int32{111, 222, 333, 444, 555, 666}
					before := out
					Opus_silk_decode_pitch(nil, lag, int8(contour), &out[1], fs, n)
					for k := 0; k < int(n); k++ {
						var offset int8
						if fs == 8 {
							if n == 2 {
								offset = Opus_silk_CB_lags_stage2_10_ms[k][contour]
							} else {
								offset = Opus_silk_CB_lags_stage2[k][contour]
							}
						} else {
							if n == 2 {
								offset = Opus_silk_CB_lags_stage3_10_ms[k][contour]
							} else {
								offset = Opus_silk_CB_lags_stage3[k][contour]
							}
						}
						want := min(max(2*fs+int32(lag)+int32(offset), 2*fs), 18*fs)
						if out[k+1] != want {
							t.Fatalf("fs=%d n=%d contour=%d lag=%d output=%v", fs, n, contour, lag, out)
						}
					}
					if out[0] != before[0] {
						t.Fatal("leading sentinel changed")
					}
					for k := int(n) + 1; k < len(out); k++ {
						if out[k] != before[k] {
							t.Fatal("tail changed")
						}
					}
				}
			}
		}
	}
}
