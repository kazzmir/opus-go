package opuscc

import "testing"

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
