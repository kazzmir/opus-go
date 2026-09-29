package opuscc

import "testing"

func TestResamplerDriverPointers(t *testing.T) {
	for _, rates := range [][2]int32{{8000, 8000}, {8000, 16000}, {12000, 48000}, {16000, 8000}} {
		var s OpusT_silk_resampler_state_struct
		Opus_silk_resampler_init(nil, &s, rates[0], rates[1], 0)
		for _, ms := range []int32{1, 2, 10, 21} {
			in := make([]int16, ms*s.FFs_in_kHz)
			for i := range in {
				in[i] = int16(i*719 + 123)
			}
			out := make([]int16, ms*s.FFs_out_kHz+2)
			out[0] = 12345
			out[len(out)-1] = -23456
			if ret := Opus_silk_resampler(nil, &s, &out[1], &in[0], int32(len(in))); ret != 0 {
				t.Fatal(ret)
			}
			if out[0] != 12345 || out[len(out)-1] != -23456 {
				t.Fatal("guards", rates, ms)
			}
			for i := int32(0); i < s.FinputDelay; i++ {
				if s.FdelayBuf[i] != in[int32(len(in))-s.FinputDelay+i] {
					t.Fatal("delay history", rates, ms)
				}
			}
		}
	}
}
