package opuscc

import (
	"math"
	"testing"
)

func TestExpRotationOuterPointers(t *testing.T) {
	for _, n := range []int{8, 16, 64, 128} {
		for _, stride := range []int32{1, 2, 4} {
			for spread := int32(1); spread <= 3; spread++ {
				x := make([]float32, n+2)
				x[0], x[n+1] = 111, 222
				for i := 1; i <= n; i++ {
					x[i] = float32(math.Sin(float64(i)))
				}
				before := append([]float32(nil), x...)
				Opus_exp_rotation(nil, &x[1], int32(n), 1, stride, 1, spread)
				Opus_exp_rotation(nil, &x[1], int32(n), -1, stride, 1, spread)
				for i := 1; i <= n; i++ {
					if math.Abs(float64(x[i]-before[i])) > 2e-6 {
						t.Fatalf("round trip n=%d stride=%d spread=%d index=%d got=%g want=%g", n, stride, spread, i, x[i], before[i])
					}
				}
				if x[0] != 111 || x[n+1] != 222 {
					t.Fatal("sentinels changed")
				}
			}
		}
	}
	Opus_exp_rotation(nil, nil, 16, 1, 1, 8, 3)
	Opus_exp_rotation(nil, nil, 16, 1, 1, 1, SPREAD_NONE)
}
