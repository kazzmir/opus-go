package opuscc

import (
	"slices"
	"testing"
)

func TestNLSF2APointers(t *testing.T) {
	for _, d := range []int{10, 16} {
		for _, tight := range []bool{false, true} {
			in := make([]int16, d)
			for i := range in {
				in[i] = int16((i + 1) * 32768 / (d + 1))
				if tight {
					in[i] = int16(2000 + 100*i)
				}
			}
			before := slices.Clone(in)
			out := make([]int16, d+2)
			out[0] = 123
			out[d+1] = 456
			Opus_silk_NLSF2A(nil, &out[1], &in[0], int32(d), 0)
			if out[0] != 123 || out[d+1] != 456 || !slices.Equal(in, before) {
				t.Fatal("buffer boundary changed")
			}
			Opus_silk_NLSF2A(nil, &in[0], &in[0], int32(d), 0)
			if !slices.Equal(in, out[1:d+1]) {
				t.Fatal("in-place output differs")
			}
		}
	}
}
