package opuscc

import (
	"slices"
	"testing"
)

func TestLPCAnalysisPointers(t *testing.T) {
	for _, d := range []int{6, 10, 16, 24} {
		for _, extra := range []int{0, 1, 19} {
			n := d + extra
			in := make([]int16, n)
			b := make([]int16, d)
			for i := range in {
				in[i] = int16(32767)
				if i%2 == 0 {
					in[i] = -32768
				}
			}
			for i := range b {
				b[i] = -32768
			}
			before := slices.Clone(in)
			out := make([]int16, n+2)
			out[0] = 123
			out[n+1] = 456
			Opus_silk_LPC_analysis_filter(nil, &out[1], &in[0], &b[0], int32(n), int32(d), 0)
			if out[0] != 123 || out[n+1] != 456 || !slices.Equal(in, before) {
				t.Fatal("buffer boundary changed")
			}
			for i := 0; i < n; i++ {
				var want int16
				if i >= d {
					var sum int64
					for j := range b {
						sum += int64(in[i-j-1]) * int64(b[j])
					}
					r := int32((int64(in[i]) << 12) - sum)
					rounded := (int64(r) + 2048) >> 12
					want = int16(max(-32768, min(32767, rounded)))
				}
				if out[i+1] != want {
					t.Fatalf("d=%d n=%d i=%d got=%d want=%d", d, n, i, out[i+1], want)
				}
			}
		}
	}
}
