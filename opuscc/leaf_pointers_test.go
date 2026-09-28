package opuscc

import "testing"

func TestCombinePulsesPointers(t *testing.T) {
	in := [...]int32{1, 2, -3, 4, 5, 6, 7, 8}
	out := [...]int32{0, 0, 0, 0, 99}
	combine_pulses(nil, &out[0], &in[0], 4)
	want := [...]int32{3, 1, 11, 15, 99}
	if out != want {
		t.Fatalf("got %v, want %v", out, want)
	}
	// The shell encoder combines successive levels of this tree.
	combine_pulses(nil, &out[0], &out[0], 2)
	if out[0] != 4 || out[1] != 26 {
		t.Fatalf("in-place combination: %v", out)
	}
	combine_pulses(nil, nil, nil, 0)
}

func TestGainsIDPointers(t *testing.T) {
	for _, tc := range []struct {
		indices [4]int8
		want    int32
	}{
		{[4]int8{1, 2, 3, 4}, 0x01020304},
		{[4]int8{-1, 0, 0, 0}, -16777216},
		{[4]int8{0, 0, 0, -1}, -1},
	} {
		if got := Opus_silk_gains_ID(nil, &tc.indices[0], 4); got != tc.want {
			t.Fatalf("indices %v: got %d, want %d", tc.indices, got, tc.want)
		}
	}
	if got := Opus_silk_gains_ID(nil, nil, 0); got != 0 {
		t.Fatalf("empty vector: %d", got)
	}
}

func TestInnerProductPointers(t *testing.T) {
	x := [...]int16{3, -3, 32767, -32768}
	y := [...]int16{3, 3, 32767, -32768}
	for _, scale := range []int32{0, 1, 8, 16} {
		var want int32
		for i := range x {
			want += (int32(x[i]) * int32(y[i])) >> scale
		}
		if got := Opus_silk_inner_prod_aligned_scale(nil, &x[0], &y[0], scale, 4); got != want {
			t.Fatalf("scale %d: got %d, want %d", scale, got, want)
		}
	}
	if got := Opus_silk_inner_prod_aligned_scale(nil, nil, nil, 0, 0); got != 0 {
		t.Fatalf("empty vectors: %d", got)
	}
}
