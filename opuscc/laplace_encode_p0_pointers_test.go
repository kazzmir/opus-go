package opuscc

import "testing"

func TestLaplaceEncodeP0Pointers(t *testing.T) {
	for _, p0 := range []uint16{1, 16000, 32766} {
		for _, decay := range []uint16{0, 1, 7, 12000, 32767} {
			var b [1024]byte
			var e OpusT_ec_enc
			Opus_ec_enc_init(nil, &e, &b[0], 1024)
			inputs := []int32{0, 1, -1, 6, -6, 7, -7, 8, -8, 14, -14, 15, -15, 127, -127}
			for _, v := range inputs {
				Opus_ec_laplace_encode_p0(nil, &e, v, p0, decay)
			}
			Opus_ec_enc_done(nil, &e)
			if e.Ferror1 != 0 {
				t.Fatal(e)
			}
			var d OpusT_ec_dec
			Opus_ec_dec_init(nil, &d, &b[0], 1024)
			for i, want := range inputs {
				if got := Opus_ec_laplace_decode_p0(nil, &d, p0, decay); got != want {
					t.Fatal(p0, decay, i, got, want)
				}
			}
		}
	}
}
