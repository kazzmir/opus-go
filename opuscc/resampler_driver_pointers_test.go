package opuscc

import (
	"runtime"
	"testing"
)

func TestResamplerPartialConsumedPointers(t *testing.T) {
	for _, rates := range [][2]int32{{8000, 8000}, {8000, 16000}, {12000, 48000}, {16000, 8000}, {48000, 8000}} {
		var state OpusT_silk_resampler_state_struct
		forEncoder := int32(0)
		if rates[0] == 48000 {
			forEncoder = 1
		}
		if Opus_silk_resampler_init(nil, &state, rates[0], rates[1], forEncoder) != 0 {
			t.Fatal("partial init", rates)
		}
		length := state.FFs_in_kHz + 1
		input := make([]int16, length)
		for i := range input {
			input[i] = int16(i*719 + 123)
		}
		count := (length*state.FFs_out_kHz + state.FFs_in_kHz - 1) / state.FFs_in_kHz
		output := make([]int16, count+2)
		output[0], output[len(output)-1] = 12345, -23456
		entropyInitGrowStack(12)
		runtime.GC()
		if Opus_silk_resampler(nil, &state, &output[1], &input[0], length) != 0 || output[0] != 12345 || output[len(output)-1] != -23456 {
			t.Fatal("partial consumed output guards", rates, output)
		}
	}
}

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
			entropyInitGrowStack(12)
			runtime.GC()
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
