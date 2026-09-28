package opuscc

import "testing"

func TestCoarseEnergyPointers(t *testing.T) {
	mode := OpusT_OpusCustomMode{FnbEBands: 4}
	energy := [10]float32{123, 9, -20, 2, 8, 9, -20, 2, 8, 456}
	dec := OpusT_ec_dec{Frng: 1 << 31, Fval: (1 << 31) - 1, Fnbits_total: 33}
	before := dec
	Opus_unquant_coarse_energy(nil, &mode, 1, 3, &energy[1], 1, &dec, 2, 0)
	if energy[0] != 123 || energy[9] != 456 || energy[1] != 9 || energy[4] != 8 || energy[5] != 9 || energy[8] != 8 || energy[2] != -1 || energy[6] != -1 || energy[3] != energy[7] || dec != before {
		t.Fatalf("energy=%v state=%+v", energy, dec)
	}
}
