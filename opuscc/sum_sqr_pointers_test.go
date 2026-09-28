package opuscc

import (
	"slices"
	"testing"
	"unsafe"
)

func TestSumSqrPointers(t *testing.T) {
	for _, tc := range []struct {
		input         []int16
		energy, shift int32
	}{
		{nil, 0, 0}, {[]int16{0}, 0, 0}, {[]int16{3, -4}, 25, 0}, {[]int16{3, -4, 12}, 169, 0},
		{[]int16{-32768}, 268435456, 2}, {[]int16{-32768, -32768}, 268435456, 3},
		{[]int16{-32768, -32768, -32768}, 402653184, 3}, {[]int16{-32768, 32767}, 536854528, 2},
	} {
		before := slices.Clone(tc.input)
		e, s := [3]int32{111, 0, 222}, [3]int32{333, 0, 444}
		Opus_silk_sum_sqr_shift(nil, &e[1], &s[1], unsafe.SliceData(tc.input), int32(len(tc.input)))
		if e[1] != tc.energy || s[1] != tc.shift {
			t.Fatalf("input=%v got=(%d,%d) want=(%d,%d)", tc.input, e[1], s[1], tc.energy, tc.shift)
		}
		if e[0] != 111 || e[2] != 222 || s[0] != 333 || s[2] != 444 || !slices.Equal(before, tc.input) {
			t.Fatal("input or sentinels changed")
		}
		var shared int32
		Opus_silk_sum_sqr_shift(nil, &shared, &shared, unsafe.SliceData(tc.input), int32(len(tc.input)))
		if shared != tc.energy {
			t.Fatal("output write order changed")
		}
	}
}
