package opuscc_test

import (
	"slices"
	"testing"

	"github.com/kazzmir/opus-go/opuscc"
	"github.com/kazzmir/opus-go/opusccenc"
)

func TestInterpolatePointers(t *testing.T) {
	for name, interpolate := range map[string]func(*int16, *int16, *int16, int32, int32){
		"decoder": func(out, a, b *int16, f, n int32) { opuscc.Opus_silk_interpolate(nil, out, a, b, f, n) },
		"encoder": func(out, a, b *int16, f, n int32) { opusccenc.Opus_silk_interpolate(nil, out, a, b, f, n) },
	} {
		t.Run(name, func(t *testing.T) {
			left := []int16{32767, -32768, -10, 10, 0, 1234}
			right := []int16{-32768, 32767, 10, -10, 3, -4321}
			for factor := int32(0); factor <= 4; factor++ {
				want := make([]int16, len(left))
				for i := range want {
					delta := int64(int16(int64(right[i]) - int64(left[i])))
					want[i] = int16(int64(left[i]) + ((delta * int64(factor)) >> 2))
				}
				out := make([]int16, len(left)+2)
				out[0], out[len(out)-1] = 111, 222
				interpolate(&out[1], &left[0], &right[0], factor, int32(len(left)))
				if !slices.Equal(out[1:len(out)-1], want) || out[0] != 111 || out[len(out)-1] != 222 {
					t.Fatalf("factor=%d output=%v want=%v", factor, out, want)
				}
				for _, aliasLeft := range []bool{true, false} {
					a, b := slices.Clone(left), slices.Clone(right)
					dst := b
					if aliasLeft {
						dst = a
					}
					interpolate(&dst[0], &a[0], &b[0], factor, int32(len(dst)))
					if !slices.Equal(dst, want) {
						t.Fatal("aliased result differs")
					}
				}
			}
			interpolate(nil, nil, nil, 0, 0)
		})
	}
}
