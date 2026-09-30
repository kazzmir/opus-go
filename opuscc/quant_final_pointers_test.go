package opuscc

import "testing"

func TestQuantFinalPointers(t *testing.T) {
	Opus_quant_energy_finalise(nil, nil, 0, 3, nil, nil, nil, nil, 1, nil, 2)
	m := OpusT_OpusCustomMode{FnbEBands: 3}
	old := [6]float32{}
	err := [6]float32{-.5, 0, .5, -.25, .25, 0}
	quant := [3]int32{0, 1, 8}
	priority := [3]int32{1, 0, 0}
	var b [8]byte
	var e OpusT_ec_enc
	Opus_ec_enc_init(nil, &e, &b[0], 8)
	Opus_quant_energy_finalise(nil, &m, 0, 3, &old[0], &err[0], &quant[0], &priority[0], 2, &e, 2)
	if old != [6]float32{0, .125, 0, 0, .125, 0} || e.Fnbits_total != 35 {
		t.Fatal(old, e)
	}
	before := e
	Opus_quant_energy_finalise(nil, &m, 0, 3, nil, &err[0], &quant[0], &priority[0], 2, &e, 2)
	if e.Fnbits_total != before.Fnbits_total+2 {
		t.Fatal("nil old must still encode", e)
	}
}
