package opuscc

import "testing"

func TestLaplaceP0Pointers(t *testing.T) {
	for _, tc := range []struct {
		value uint32
		want  int32
	}{{32767 << 16, 0}, {12000 << 16, 1}, {4000 << 16, -1}} {
		dec := OpusT_ec_dec{Frng: 1 << 31, Fval: tc.value, Fnbits_total: 33, Fext: 123, Fend_window: 456, Fnend_bits: 7, Ferror1: 9}
		if got := Opus_ec_laplace_decode_p0(nil, &dec, 16000, 0); got != tc.want || dec.Fext != 123 || dec.Fend_window != 456 || dec.Fnend_bits != 7 || dec.Ferror1 != 9 {
			t.Fatalf("value=%d got=%d want=%d state=%+v", tc.value, got, tc.want, dec)
		}
	}
	dec := OpusT_ec_dec{Frng: 1 << 31, Fval: 1234567, Fnbits_total: 33}
	before := dec
	if got := Opus_ec_laplace_decode_p0(nil, &dec, 32768, 12000); got != 0 || dec != before {
		t.Fatal("certain-zero distribution changed context")
	}
}
