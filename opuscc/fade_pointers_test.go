package opuscc

import "testing"

func TestFadePointers(t *testing.T) {
	smooth_fade(nil, nil, nil, nil, 0, 2, nil, 48000)
	a := [4]float32{1, 2, 3, 4}
	b := [4]float32{5, 6, 7, 8}
	w := [3]float32{0, 99, 0.5}
	out := [6]float32{77, 0, 0, 0, 0, 88}
	smooth_fade(nil, &a[0], &b[0], &out[1], 2, 2, &w[0], 24000)
	if out != [6]float32{77, 1, 2, 4, 5, 88} {
		t.Fatal(out)
	}
	smooth_fade(nil, &a[0], &b[0], &a[0], 2, 2, &w[0], 24000)
	if a != [4]float32{1, 2, 4, 5} {
		t.Fatal("alias", a)
	}
}
