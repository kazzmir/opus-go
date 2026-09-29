package opuscc

import "testing"

func TestEntropyCarryPointers(t *testing.T) {
	b := [6]byte{77, 0, 0, 0, 0, 88}
	e := OpusT_ec_enc{Fbuf: &b[1], Fstorage: 4, Frem: -1}
	ec_enc_carry_out(nil, &e, 255)
	ec_enc_carry_out(nil, &e, 255)
	ec_enc_carry_out(nil, &e, 256)
	if e.Fext != 0 || e.Frem != 0 || e.Foffs != 2 || b != [6]byte{77, 0, 0, 0, 0, 88} {
		t.Fatal(e, b)
	}
	e.Fext = 3
	ec_enc_carry_out(nil, &e, 1)
	if e.Foffs != 4 || e.Ferror1 != -1 || e.Fext != 0 || b[0] != 77 || b[5] != 88 {
		t.Fatal("full", e, b)
	}
	e = OpusT_ec_enc{Frem: -1, Fext: 0xffffffff}
	ec_enc_carry_out(nil, &e, 255)
	if e.Fext != 0 || e.Frem != -1 {
		t.Fatal("wrap", e)
	}
}
