package opuscc

import "testing"

func TestStereoPredPointers(t *testing.T) {
	dec := OpusT_ec_dec{Frng: 1 << 31, Fval: (1 << 31) - 1, Fnbits_total: 33, Fext: 123, Ferror1: 7, Fend_window: 456, Fnend_bits: 9}
	out := [4]int32{111, 0, 0, 222}
	Opus_silk_stereo_decode_pred(nil, &dec, (*[2]int32)(out[1:3]))
	if out != [4]int32{111, 0, -13364, 222} || dec.Fext != 123 || dec.Ferror1 != 7 || dec.Fend_window != 456 || dec.Fnend_bits != 9 {
		t.Fatalf("out=%v state=%+v", out, dec)
	}
	for _, tc := range []struct {
		value uint32
		want  int32
	}{{0, 1}, {(1 << 31) - 1, 0}} {
		dec = OpusT_ec_dec{Frng: 1 << 31, Fval: tc.value, Fnbits_total: 33}
		flag := [3]int32{111, -1, 222}
		Opus_silk_stereo_decode_mid_only(nil, &dec, &flag[1])
		if flag != [3]int32{111, tc.want, 222} {
			t.Fatalf("value=%d flag=%v", tc.value, flag)
		}
	}
}
