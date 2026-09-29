package opuscc

import "testing"

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
