package opusccenc

import (
	"slices"
	"testing"
)

func TestLaroiaPointers(t *testing.T) {
	for _, input := range [][]int16{{8192, 16384}, {0, 32767}, {-32768, 32767}, {100, 100, 99, 100}, {1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000}} {
		want := make([]int16, len(input))
		for i, v := range input {
			left, right := int64(0), int64(32768)
			if i > 0 {
				left = int64(input[i-1])
			}
			if i+1 < len(input) {
				right = int64(input[i+1])
			}
			want[i] = int16(min(131072/max(int64(v)-left, 1)+131072/max(right-int64(v), 1), 32767))
		}
		out := make([]int16, len(input)+2)
		out[0], out[len(out)-1] = 111, 222
		before := slices.Clone(input)
		Opus_silk_NLSF_VQ_weights_laroia(nil, &out[1], &input[0], int32(len(input)))
		if !slices.Equal(out[1:len(out)-1], want) || out[0] != 111 || out[len(out)-1] != 222 || !slices.Equal(input, before) {
			t.Fatalf("input=%v got=%v want=%v", input, out, want)
		}
		alias := slices.Clone(input)
		Opus_silk_NLSF_VQ_weights_laroia(nil, &alias[0], &alias[0], int32(len(alias)))
		if !slices.Equal(alias, want) {
			t.Fatal("in-place result differs")
		}
	}
}
