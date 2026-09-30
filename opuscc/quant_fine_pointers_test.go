package opuscc

import "testing"

func TestQuantFinePointers(t *testing.T) {
	Opus_quant_fine_energy(nil, nil, 2, 2, nil, nil, nil, nil, nil, 2)
	m := OpusT_OpusCustomMode{FnbEBands: 3}
	old := [6]float32{}
	err := [6]float32{-.6, 0, .6, -.25, .25, 0}
	extra := [3]int32{2, 0, 2}
	var b [8]byte
	var e OpusT_ec_enc
	Opus_ec_enc_init(nil, &e, &b[0], 8)
	Opus_quant_fine_energy(nil, &m, 0, 3, &old[0], &err[0], nil, &extra[0], &e, 2)
	if old != [6]float32{-.375, 0, .375, -.125, 0, .125} || e.Fnbits_total != 41 {
		t.Fatal(old, err, e)
	}
	// C updates old before subtracting error, including when they alias.
	a := [3]float32{.2, .3, .4}
	before := a
	Opus_quant_fine_energy(nil, &m, 0, 3, &a[0], &a[0], nil, &extra[0], &e, 1)
	for i := range a {
		offset := []float32{.125, 0, .375}[i]
		if a[i] != float32(float32(before[i]+offset)-offset) {
			t.Fatal("alias", a, before)
		}
	}
}
