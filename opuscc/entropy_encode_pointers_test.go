package opuscc

import "testing"

func TestEntropyEncodePointers(t *testing.T) {
	var e OpusT_ec_enc
	Opus_ec_enc_init(nil, &e, nil, 0)
	before := e
	Opus_ec_encode(nil, &e, 0, 13, 13)
	if e != before {
		t.Fatal("full interval")
	}
	Opus_ec_encode(nil, &e, 3, 7, 13)
	if e.Fval != 495573158 || e.Frng != 660764196 {
		t.Fatal(e)
	}
	var b [4]byte
	Opus_ec_enc_init(nil, &e, &b[0], 4)
	for i := 0; i < 64; i++ {
		Opus_ec_encode(nil, &e, 1, 2, 256)
	}
	if e.Ferror1 != -1 || e.Foffs != 4 || e.Frng == 0 {
		t.Fatal("exhausted", e)
	}
}
