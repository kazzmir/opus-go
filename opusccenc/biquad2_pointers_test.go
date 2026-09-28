package opusccenc

import (
	"slices"
	"testing"
)

func TestBiquad2Pointers(t *testing.T) {
	input := []int16{-32768, 32767, 0, 1, -1, 100, -200, 300}
	before := slices.Clone(input)
	for _, b := range [][3]int32{{1 << 28}, {1 << 29}, {1 << 27, 1 << 26, -(1 << 25)}} {
		a := [2]int32{-(1 << 26) + 123, (1 << 24) + 456}
		bs, as := b, a
		initial := [4]int32{1234, -5678, -9876, 5432}
		state := initial
		out := [10]int16{111, 0, 0, 0, 0, 0, 0, 0, 0, 222}
		Opus_silk_biquad_alt_stride2_c(nil, &input[0], &b, &a, &state, &out[1], 4)
		for ch := 0; ch < 2; ch++ {
			mono := []int16{input[ch], input[ch+2], input[ch+4], input[ch+6]}
			ms := [2]int32{initial[2*ch], initial[2*ch+1]}
			Opus_silk_biquad_alt_stride1(nil, &mono[0], &b, &a, &ms, &mono[0], 4)
			for i, v := range mono {
				if v != out[1+2*i+ch] {
					t.Fatal("channel result differs from mono")
				}
			}
			if ms != [2]int32{state[2*ch], state[2*ch+1]} {
				t.Fatal("channel state differs")
			}
		}
		inPlace := slices.Clone(input)
		ss := initial
		Opus_silk_biquad_alt_stride2_c(nil, &inPlace[0], &b, &a, &ss, &inPlace[0], 2)
		Opus_silk_biquad_alt_stride2_c(nil, &inPlace[4], &b, &a, &ss, &inPlace[4], 2)
		if !slices.Equal(inPlace, out[1:9]) || ss != state || out[0] != 111 || out[9] != 222 || !slices.Equal(input, before) || b != bs || a != as {
			t.Fatal("chunking, aliasing, or read-only input changed")
		}
		Opus_silk_biquad_alt_stride2_c(nil, nil, &b, &a, &ss, nil, 0)
		if ss != state {
			t.Fatal("empty input changed state")
		}
	}
}
