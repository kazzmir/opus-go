package opuscc

import "testing"

func TestShellEncodePointers(t *testing.T) {
	var b [1024]byte
	var e OpusT_ec_enc
	Opus_ec_enc_init(nil, &e, &b[0], 1024)
	var empty [16]int32
	before := e
	Opus_silk_shell_encoder(nil, &e, &empty)
	if e != before {
		t.Fatal("zero tree")
	}
	for total := int32(1); total <= 16; total++ {
		for pos := 0; pos < 16; pos++ {
			var p [16]int32
			p[pos] = total
			Opus_silk_shell_encoder(nil, &e, &p)
		}
	}
	Opus_ec_enc_done(nil, &e)
	if e.Ferror1 != 0 {
		t.Fatal(e)
	}
	var d OpusT_ec_dec
	Opus_ec_dec_init(nil, &d, &b[0], 1024)
	for total := int32(1); total <= 16; total++ {
		for pos := 0; pos < 16; pos++ {
			var got [16]int16
			Opus_silk_shell_decoder(nil, &got, &d, total)
			for j, v := range got {
				want := int16(0)
				if j == pos {
					want = int16(total)
				}
				if v != want {
					t.Fatal(total, pos, j, v)
				}
			}
		}
	}
}
