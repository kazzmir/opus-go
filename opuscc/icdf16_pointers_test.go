package opuscc

import "testing"

func TestICDF16Pointers(t *testing.T) {
	table := []uint16{60000, 40000, 20000, 0}
	for symbol := int32(0); symbol < 4; symbol++ {
		dec := OpusT_ec_dec{Frng: 1 << 31, Fval: uint32(table[symbol]) * (1 << 15), Fnbits_total: 33, Fext: 123, Fend_window: 456, Fnend_bits: 7}
		before := dec
		got := Opus_ec_dec_icdf16(nil, &dec, &table[0], 16)
		upper := uint32(65536)
		if symbol > 0 {
			upper = uint32(table[symbol-1])
		}
		want := before
		want.Fval = 0
		want.Frng = (upper - uint32(table[symbol])) * (1 << 15)
		if got != symbol || dec != want {
			t.Fatalf("symbol=%d got=%d state=%+v want=%+v", symbol, got, dec, want)
		}
	}
	// Probability 1/65536 requires two normalization steps on an exhausted packet.
	tiny := []uint16{65535, 0}
	dec := OpusT_ec_dec{Frng: 1 << 31, Fval: (1 << 31) - 1, Fnbits_total: 33}
	if got := Opus_ec_dec_icdf16(nil, &dec, &tiny[0], 16); got != 0 || dec.Frng != 1<<31 || dec.Fval != (1<<31)-1 || dec.Fnbits_total != 49 {
		t.Fatalf("tiny symbol=%d state=%+v", got, dec)
	}
}
