package opuscc

import "testing"

func TestEntropyNormalizePointers(t *testing.T) {
	e := OpusT_ec_enc{Frng: 1<<23 + 1, Fval: 123, Frem: -1}
	before := e
	ec_enc_normalize(nil, &e)
	if e != before {
		t.Fatal("identity")
	}
	b := [5]byte{77, 0, 0, 0, 88}
	e = OpusT_ec_enc{Fbuf: &b[1], Fstorage: 3, Frng: 1, Fval: 0, Frem: -1, Fnbits_total: 33}
	ec_enc_normalize(nil, &e)
	if e.Frng != 1<<24 || e.Fnbits_total != 57 || e.Foffs != 2 || e.Frem != 0 || b[0] != 77 || b[4] != 88 {
		t.Fatal(e, b)
	}
	e = OpusT_ec_enc{Frng: 1, Fval: 0, Frem: 0}
	ec_enc_normalize(nil, &e)
	if e.Ferror1 != -1 || e.Frng != 1<<24 {
		t.Fatal("exhausted", e)
	}
}
