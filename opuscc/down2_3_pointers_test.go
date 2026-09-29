package opuscc

import (
	"testing"
	"unsafe"
)

func TestDownTwoThirdsPointers(t *testing.T) {
	s := [6]int32{1, 2, 3, 4, 5, 6}
	before := s
	Opus_silk_resampler_down2_3(nil, &s, nil, nil, 0)
	if s != before {
		t.Fatal("empty state")
	}
	for _, n := range []int{1, 2, 3, 4, 5, 479, 480, 481, 482, 960, 962} {
		in := make([]int16, n)
		for i := range in {
			in[i] = int16(i * 1999)
		}
		out := make([]int16, 2*(n/3)+2)
		out[0] = 123
		out[len(out)-1] = 456
		Opus_silk_resampler_down2_3(nil, &s, &out[1], unsafe.SliceData(in), int32(n))
		if out[0] != 123 || out[len(out)-1] != 456 {
			t.Fatal("output guards", n)
		}
	}
	var input [2]int16
	Opus_silk_resampler_down2_3(nil, &s, nil, &input[0], 2)
}
