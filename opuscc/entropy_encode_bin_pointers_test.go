package opuscc

import "testing"

func TestEntropyEncodeBinPointers(t *testing.T) {
	var e OpusT_ec_enc
	Opus_ec_enc_init(nil, &e, nil, 0)
	before := e
	Opus_ec_encode_bin(nil, &e, 0, 1, 0)
	if e != before {
		t.Fatal("full interval")
	}
	Opus_ec_encode_bin(nil, &e, 5, 10, 4)
	if e.Fval != 671088640 || e.Frng != 671088640 {
		t.Fatal(e)
	}
	var b [4]byte
	Opus_ec_enc_init(nil, &e, &b[0], 4)
	for i := 0; i < 64; i++ {
		Opus_ec_encode_bin(nil, &e, 1, 2, 15)
	}
	if e.Ferror1 != -1 || e.Foffs != 4 || e.Frng == 0 {
		t.Fatal("exhausted", e)
	}
}
