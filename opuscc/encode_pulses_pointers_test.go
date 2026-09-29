package opuscc

import (
	"slices"
	"testing"
)

func TestEncodePulsesPointers(t *testing.T) {
	for _, pair := range [][2]int32{{2, 128}, {3, 32}, {8, 8}, {176, 1}} {
		n, k := pair[0], pair[1]
		var b [512]byte
		var e OpusT_ec_enc
		Opus_ec_enc_init(nil, &e, &b[0], 512)
		vectors := make([][]int32, 20)
		for i := range vectors {
			y := make([]int32, n)
			y[i%int(n)] = k
			if i%2 != 0 {
				y[i%int(n)] = -k
			}
			vectors[i] = y
			Opus_encode_pulses(nil, &y[0], n, k, &e)
		}
		Opus_ec_enc_done(nil, &e)
		if e.Ferror1 != 0 {
			t.Fatal(e)
		}
		var d OpusT_ec_dec
		Opus_ec_dec_init(nil, &d, &b[0], 512)
		for _, want := range vectors {
			got := make([]int32, n)
			Opus_decode_pulses(nil, &got[0], n, k, &d)
			if !slices.Equal(got, want) {
				t.Fatal(n, k, got, want)
			}
		}
	}
}
