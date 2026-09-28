package opuscc

import "testing"

func TestFinalEnergyPointers(t *testing.T) {
	mode := OpusT_OpusCustomMode{FnbEBands: 3}
	quant := [3]int32{0, 0, MAX_FINE_BITS}
	priority := [3]int32{1, 0, 0}
	for _, budget := range []int32{1, 3, 4} {
		out := [8]float32{123, 0, 0, 0, 0, 0, 0, 456}
		dec := OpusT_ec_dec{Frng: 1 << 31, Fnbits_total: 33, Fend_window: 15, Fnend_bits: 32}
		skip := dec
		Opus_unquant_energy_finalise(nil, &mode, 0, 3, &out[1], &quant[0], &priority[0], budget, &dec, 2)
		Opus_unquant_energy_finalise(nil, &mode, 0, 3, nil, &quant[0], &priority[0], budget, &skip, 2)
		want := [8]float32{123, 0, 0, 0, 0, 0, 0, 456}
		if budget >= 2 {
			want[2] = 0.25
			want[5] = 0.25
		}
		if budget >= 4 {
			want[1] = 0.25
			want[4] = 0.25
		}
		if out != want || dec != skip || dec.Fnbits_total != 33+budget/2*2 {
			t.Fatalf("budget=%d out=%v state=%+v skip=%+v", budget, out, dec, skip)
		}
	}
}
