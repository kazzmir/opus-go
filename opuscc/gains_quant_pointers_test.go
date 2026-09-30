package opuscc

import "testing"

func TestGainsQuantPointers(t *testing.T) {
	Opus_silk_gains_quant(nil, nil, nil, nil, 0, 0)
	for previous := int8(0); previous < 64; previous++ {
		for cond := int32(0); cond < 2; cond++ {
			gains := [4]int32{1, 65536, 1 << 24, 2147483647}
			var indices [4]int8
			p := previous
			Opus_silk_gains_quant(nil, &indices[0], &gains[0], &p, cond, 4)
			var decoded [4]int32
			dp := previous
			Opus_silk_gains_dequant(nil, &decoded[0], &indices[0], &dp, cond, 4)
			if gains != decoded || p != dp {
				t.Fatal(previous, cond, gains, decoded, p, dp)
			}
		}
	}
}
