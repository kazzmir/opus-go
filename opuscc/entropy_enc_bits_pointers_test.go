package opuscc

import "testing"

func TestEntropyEncBitsPointers(t *testing.T) {
	b := [6]byte{77, 0, 0, 0, 0, 88}
	var e OpusT_ec_enc
	Opus_ec_enc_init(nil, &e, &b[1], 4)
	Opus_ec_enc_bits(nil, &e, 0xabcdef, 24)
	Opus_ec_enc_bits(nil, &e, 0x12345, 20)
	if b != [6]byte{77, 0, 0xab, 0xcd, 0xef, 88} || e.Fend_offs != 3 || e.Fend_window != 0x12345 || e.Fnend_bits != 20 || e.Fnbits_total != 77 {
		t.Fatal(e, b)
	}
	Opus_ec_enc_bits(nil, &e, 0x1ffffff, 25)
	if e.Fend_offs != 4 || e.Ferror1 != -1 || b[0] != 77 || b[5] != 88 {
		t.Fatal("exhausted", e, b)
	}
	e = OpusT_ec_enc{Fnbits_total: 2147483647}
	Opus_ec_enc_bits(nil, &e, 1, 1)
	if e.Fnbits_total != -2147483648 || e.Fend_window != 1 {
		t.Fatal("wrap", e)
	}
}
