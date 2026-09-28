package opuscc

import (
	"slices"
	"testing"
)

func TestBiquad2Pointers(t *testing.T) {
	b := [3]int32{1 << 27, 1 << 26, -(1 << 25)}
	a := [2]int32{-(1 << 26) + 123, (1 << 24) + 456}
	for _, n := range []int{2, 16, 64} {
		input := make([]int16, 2*n)
		for i := range input {
			input[i] = int16(32767 - i*1700)
		}
		initial := [4]int32{10000, -20000, -30000, 40000}
		state := initial
		out := make([]int16, len(input)+2)
		out[0], out[len(out)-1] = 111, 222
		Opus_silk_biquad_alt_stride2_c(nil, &input[0], &b, &a, &state, &out[1], int32(n))
		for channel := 0; channel < 2; channel++ {
			mono := make([]int16, n)
			for i := range mono {
				mono[i] = input[2*i+channel]
			}
			monoState := [2]int32{initial[2*channel], initial[2*channel+1]}
			Opus_silk_biquad_alt_stride1(nil, &mono[0], &b, &a, &monoState, &mono[0], int32(n))
			for i := range mono {
				if mono[i] != out[1+2*i+channel] {
					t.Fatalf("n=%d channel=%d sample=%d: interleaved output differs", n, channel, i)
				}
			}
			if monoState != [2]int32{state[2*channel], state[2*channel+1]} {
				t.Fatal("channel state differs")
			}
		}
		if out[0] != 111 || out[len(out)-1] != 222 {
			t.Fatal("sentinel changed")
		}
		inPlace := slices.Clone(input)
		splitState := initial
		Opus_silk_biquad_alt_stride2_c(nil, &inPlace[0], &b, &a, &splitState, &inPlace[0], int32(n/2))
		Opus_silk_biquad_alt_stride2_c(nil, &inPlace[n], &b, &a, &splitState, &inPlace[n], int32(n/2))
		if !slices.Equal(inPlace, out[1:len(out)-1]) || splitState != state {
			t.Fatal("in-place/chunked processing differs")
		}
	}
	state := [4]int32{1, 2, 3, 4}
	Opus_silk_biquad_alt_stride2_c(nil, nil, &b, &a, &state, nil, 0)
	if state != [4]int32{1, 2, 3, 4} {
		t.Fatal("empty input changed state")
	}
}
