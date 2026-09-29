package opuscc

import "testing"

func TestLaplaceEncodePointers(t *testing.T) {
	for _, fs := range []uint32{1, 1024, 16000, 32700} {
		for _, decay := range []int32{1, 6000, 11456} {
			var b [256]byte
			var e OpusT_ec_enc
			Opus_ec_enc_init(nil, &e, &b[0], 256)
			inputs := []int32{0, 1, -1, 2, -2, 10, -10, 16384, -16384, 2147483647, -2147483647}
			clipped := make([]int32, len(inputs))
			for i, v := range inputs {
				clipped[i] = v
				Opus_ec_laplace_encode(nil, &e, &clipped[i], fs, decay)
				if v < 0 && clipped[i] >= 0 || v > 0 && clipped[i] <= 0 {
					t.Fatal("sign", v, clipped[i])
				}
			}
			Opus_ec_enc_done(nil, &e)
			if e.Ferror1 != 0 {
				t.Fatal(e)
			}
			var d OpusT_ec_dec
			Opus_ec_dec_init(nil, &d, &b[0], 256)
			for i, want := range clipped {
				if got := Opus_ec_laplace_decode(nil, &d, fs, decay); got != want {
					t.Fatal(fs, decay, i, got, want)
				}
			}
		}
	}
}
