package opuscc

import (
	"math"
	"slices"
	"testing"
)

func TestNormalisePointers(t *testing.T) {
	input := [4]int32{3, -4, 0, 12}
	for _, gain := range []float32{0, 0.5, 1, 2} {
		out := [6]float32{111, 0, 0, 0, 0, 222}
		g := float32(1) / float32(13) * gain
		for _, shift := range []int32{0, 3} {
			normalise_residual(nil, &input[0], &out[1], 4, 169, gain, shift)
			for i, value := range input {
				if out[i+1] != float32(value)*g {
					t.Fatalf("gain=%g shift=%d index=%d: %g", gain, shift, i, out[i+1])
				}
			}
			if out[0] != 111 || out[5] != 222 {
				t.Fatal("sentinels changed")
			}
		}
	}
	for _, input := range [][]float32{{0}, {3, 4}, {0, 0, 0}, {-12, 3, -4}, {1e-10, -1e-10}} {
		for _, gain := range []float32{0, 0.5, 1, 2} {
			x := append([]float32{111}, input...)
			x = append(x, 222)
			want := slices.Clone(input)
			var energy float32
			for _, v := range input {
				energy += float32(v * v)
			}
			scale := float32(1) / float32(math.Sqrt(float64(float32(1e-15)+energy))) * gain
			for i := range want {
				want[i] *= scale
			}
			Opus_renormalise_vector(nil, &x[1], int32(len(input)), gain, 0)
			if !slices.Equal(x[1:len(x)-1], want) || x[0] != 111 || x[len(x)-1] != 222 {
				t.Fatalf("input=%v gain=%g: got=%v want=%v", input, gain, x, want)
			}
		}
	}
	Opus_renormalise_vector(nil, nil, 0, 1, 0)
}
