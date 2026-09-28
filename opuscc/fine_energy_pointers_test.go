package opuscc

import "testing"

func TestFineEnergyPointers(t *testing.T) {
	mode := OpusT_OpusCustomMode{FnbEBands: 3}
	extra := [3]int32{0, 2, 14}
	prev := [3]int32{0, 1, 0}
	for _, previous := range []*int32{nil, &prev[0]} {
		energy := [8]float32{123, 1, 2, 3, 4, 5, 6, 456}
		// Four available bits permit only the middle band's stereo refinement.
		dec := OpusT_ec_dec{Fstorage: 1, Frng: 1 << 31, Fnbits_total: 36, Fend_window: 15, Fnend_bits: 32}
		Opus_unquant_fine_energy(nil, &mode, 0, 3, &energy[1], previous, &extra[0], &dec, 2)
		delta := float32(0.375)
		if previous != nil {
			delta /= 2
		}
		if energy != [8]float32{123, 1, 2 + delta, 3, 4, 5 + delta, 6, 456} || dec.Fnbits_total != 40 || dec.Fnend_bits != 28 {
			t.Fatalf("energy=%v state=%+v", energy, dec)
		}
	}
}
