package opuscc

import "testing"

func TestCapsPointers(t *testing.T) {
	bands := [4]int16{0, 1, 3, 7}
	var cache [24]uint8
	for i := range cache {
		cache[i] = uint8(i * 7)
	}
	for lm := int32(0); lm < 4; lm++ {
		for channels := int32(1); channels <= 2; channels++ {
			out := [5]int32{123, 0, 0, 0, 456}
			Opus_init_caps(nil, &bands[0], &cache[0], &out[1], 3, lm, channels)
			for i := int32(0); i < 3; i++ {
				want := (int32(cache[3*(2*lm+channels-1)+i]) + 64) * channels * (int32(bands[i+1]-bands[i]) << lm) / 4
				if out[i+1] != want {
					t.Fatalf("LM=%d C=%d i=%d got=%d want=%d", lm, channels, i, out[i+1], want)
				}
			}
			if out[0] != 123 || out[4] != 456 {
				t.Fatal("guard changed")
			}
		}
	}
}
