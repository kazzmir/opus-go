package opuscc

import (
	"math"
	"slices"
	"testing"
)

// From intensity_pointers_test.go.
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

// From stereo_itheta_pointers_test.go.
func TestStereoIThetaPointers(t *testing.T) {
	for _, stereo := range []int32{0, 1} {
		if Opus_stereo_itheta(nil, nil, nil, stereo, 0, 0) != 0 {
			t.Fatal("empty")
		}
	}
	x := [3]float32{123, 0.75, 456}
	y := [3]float32{789, 0.75, 321}
	if got := Opus_stereo_itheta(nil, &x[1], &y[1], 1, 1, 0); got != 0 {
		t.Fatal("mid", got)
	}
	y[1] = -x[1]
	if got := Opus_stereo_itheta(nil, &x[1], &y[1], 1, 1, 0); got != 1<<30 {
		t.Fatal("side", got)
	}
	y[1] = 0
	if got := Opus_stereo_itheta(nil, &x[1], &y[1], 0, 1, 0); got != 0 {
		t.Fatal("X only", got)
	}
	x[1] = 0
	y[1] = 0.75
	if got := Opus_stereo_itheta(nil, &x[1], &y[1], 0, 1, 0); got != 1<<30 {
		t.Fatal("Y only", got)
	}
	x[1] = 1e-12
	y[1] = 1e-12
	if got := Opus_stereo_itheta(nil, &x[1], &y[1], 1, 1, 0); got != 0 {
		t.Fatal("tiny", got)
	}
	if x != [3]float32{123, 1e-12, 456} || y != [3]float32{789, 1e-12, 321} {
		t.Fatal("modified inputs")
	}
}

// From stereo_merge_pointers_test.go.
func TestStereoMergePointers(t *testing.T) {
	scale := float32(1) / float32(math.Sqrt(2))
	for _, tc := range []struct {
		x, y         []float32
		mid          float32
		wantX, wantY []float32
	}{
		{[]float32{1, 0}, []float32{0, 1}, 1, []float32{scale, -scale}, []float32{scale, scale}},
		{[]float32{1, 0}, []float32{0, 1}, 0, []float32{0, -1}, []float32{0, 1}},
		{[]float32{1, 0}, []float32{1, 0}, 1, []float32{1, 0}, []float32{1, 0}},
		{[]float32{1, 0}, []float32{-1, 0}, 1, []float32{1, 0}, []float32{1, 0}},
		{[]float32{0.25, 0.5}, []float32{0, 0}, 0.01, []float32{0.25, 0.5}, []float32{0.25, 0.5}},
	} {
		x := append([]float32{111}, tc.x...)
		x = append(x, 222)
		y := append([]float32{333}, tc.y...)
		y = append(y, 444)
		stereo_merge(nil, &x[1], &y[1], tc.mid, int32(len(tc.x)), 0)
		if !slices.Equal(x[1:len(x)-1], tc.wantX) || !slices.Equal(y[1:len(y)-1], tc.wantY) {
			t.Fatalf("mid=%g X=%v Y=%v", tc.mid, x, y)
		}
		if x[0] != 111 || x[len(x)-1] != 222 || y[0] != 333 || y[len(y)-1] != 444 {
			t.Fatal("sentinels changed")
		}
	}
	stereo_merge(nil, nil, nil, 1, 0, 0)
}
