package opuscc

import (
	"testing"
)

func TestEntropyEncICDFPointers(t *testing.T) {
	table := [5]byte{77, 240, 180, 100, 0}
	var e OpusT_ec_enc
	var b [128]byte
	Opus_ec_enc_init(nil, &e, &b[0], 128)
	for i := 0; i < 100; i++ {
		Opus_ec_enc_icdf(nil, &e, int32(i%4), &table[1], 8)
	}
	Opus_ec_enc_done(nil, &e)
	if e.Ferror1 != 0 || table != [5]byte{77, 240, 180, 100, 0} {
		t.Fatal(e, table)
	}
	var d OpusT_ec_dec
	Opus_ec_dec_init(nil, &d, &b[0], 128)
	for i := 0; i < 100; i++ {
		if s := ec_dec_icdf(nil, &d, &table[1], 8); s != int32(i%4) {
			t.Fatal(i, s)
		}
	}
	single := byte(0)
	Opus_ec_enc_init(nil, &e, nil, 0)
	before := e
	Opus_ec_enc_icdf(nil, &e, 0, &single, 8)
	if e != before {
		t.Fatal("singleton")
	}
}
