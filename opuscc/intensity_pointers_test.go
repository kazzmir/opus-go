package opuscc

import (
	"math"
	"testing"
)

func TestIntensityPointers(t *testing.T) {
	mode := OpusT_OpusCustomMode{FnbEBands: 3}
	for _, energy := range [][2]float32{{3, 4}, {1, 0}, {0, 1}, {0, 0}, {1e-12, 2e-12}} {
		bands := [5]float32{99, energy[0], 88, 77, energy[1]}
		before := bands
		x := [5]float32{111, 1, -2, 0.5, 222}
		y := [3]float32{2, 1, -0.5}
		oldX, oldY := x, y
		norm := float32(1e-15) + float32(math.Sqrt(float64(float32(1e-15)+float32(energy[0]*energy[0])+float32(energy[1]*energy[1]))))
		a, b := energy[0]/norm, energy[1]/norm
		intensity_stereo(nil, &mode, &x[1], &y[0], &bands[0], 1, 3)
		for i := range y {
			want := float32(a*oldX[i+1]) + float32(b*oldY[i])
			if math.Float32bits(x[i+1]) != math.Float32bits(want) {
				t.Fatalf("energy=%v i=%d got=%g want=%g", energy, i, x[i+1], want)
			}
		}
		if x[0] != 111 || x[4] != 222 || y != oldY || bands != before {
			t.Fatal("read-only input or sentinel changed")
		}
		intensity_stereo(nil, &mode, nil, nil, &bands[0], 1, 0)
	}
}
