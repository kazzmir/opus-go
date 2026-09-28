package opuscc

import "testing"

func TestEntropyUpdatePointers(t *testing.T) {
	for _, fl := range []uint32{0, 1} {
		dec := OpusT_ec_dec{Frng: 100000003, Fval: 80000000, Fext: 10000000, Fnbits_total: 33, Frem: 123, Fend_window: 456, Fnend_bits: 9, Ferror1: 7}
		want := dec
		want.Fval = 10000000
		if fl == 0 {
			want.Frng = 30000003
		} else {
			want.Frng = 20000000
		}
		Opus_ec_dec_update(nil, &dec, fl, 3, 10)
		if dec != want {
			t.Fatalf("fl=%d got=%+v want=%+v", fl, dec, want)
		}
	}
	// An exhausted packet supplies zero bytes while normalization continues.
	dec := OpusT_ec_dec{Frng: 1 << 24, Fval: 123, Fext: 1, Fnbits_total: 33, Frem: 0}
	Opus_ec_dec_update(nil, &dec, 1, 2, 2)
	if dec.Frng != 1<<24 || dec.Fval != 2080374783 || dec.Fnbits_total != 57 || dec.Foffs != 0 || dec.Frem != 0 {
		t.Fatalf("normalized=%+v", dec)
	}
}
