package opuscc

import (
	"math"
	"testing"
)

func TestAmpLogPointers(t *testing.T) {
	Opus_amp2Log2(nil, nil, 0, 0, nil, nil, 2)
	m := OpusT_OpusCustomMode{FnbEBands: 4}
	in := [8]float32{1, 2, 4, 8, .5, .25, .125, .0625}
	out := [8]float32{77, 77, 77, 77, 77, 77, 77, 77}
	Opus_amp2Log2(nil, &m, 2, 3, &in[0], &out[0], 2)
	if out[2] != -14 || out[6] != -14 || out[3] != 77 || out[7] != 77 {
		t.Fatal(out)
	}
	for _, pos := range []int{0, 1, 4, 5} {
		i := pos % 4
		want := float32(math.Log2(float64(in[pos]))) - Opus_eMeans[i]
		if math.Abs(float64(out[pos]-want)) > 1e-4 {
			t.Fatal(pos, out[pos], want)
		}
	}
	alias := in
	Opus_amp2Log2(nil, &m, 2, 3, &alias[0], &alias[0], 2)
	for _, pos := range []int{0, 1, 2, 4, 5, 6} {
		if math.Float32bits(alias[pos]) != math.Float32bits(out[pos]) {
			t.Fatal("alias", alias, out)
		}
	}
}
