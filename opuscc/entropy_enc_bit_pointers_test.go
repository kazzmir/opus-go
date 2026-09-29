package opuscc

import "testing"

func TestEntropyEncBitPointers(t *testing.T) {
	for _, value := range []int32{0, 1, -1, 7} {
		var e OpusT_ec_enc
		Opus_ec_enc_init(nil, &e, nil, 0)
		Opus_ec_enc_bit_logp(nil, &e, value, 3)
		if value == 0 {
			if e.Fval != 0 || e.Frng != 0x70000000 {
				t.Fatal(e)
			}
		} else {
			if e.Fval != 0x70000000 || e.Frng != 0x10000000 {
				t.Fatal(e)
			}
		}
	}
	var b [3]byte
	e := OpusT_ec_enc{}
	Opus_ec_enc_init(nil, &e, &b[0], 3)
	for i := 0; i < 100; i++ {
		Opus_ec_enc_bit_logp(nil, &e, int32(i%2), 15)
	}
	if e.Foffs != 3 || e.Ferror1 != -1 || e.Frng == 0 {
		t.Fatal(e)
	}
}
