package opuscc

import "testing"

func TestShellDecodePointers(t *testing.T) {
	for total := int32(0); total <= 16; total++ {
		state := OpusT_ec_dec{Frng: 1 << 31, Fval: (1 << 31) - 1, Fnbits_total: 33}
		var output [18]int16
		output[0], output[17] = 111, 222
		for i := 1; i <= 16; i++ {
			output[i] = -1
		}
		Opus_silk_shell_decoder(nil, (*[16]int16)(output[1:17]), &state, total)
		for i := 1; i <= 16; i++ {
			want := int16(0)
			if i == 16 {
				want = int16(total)
			}
			if output[i] != want {
				t.Fatalf("total=%d output=%v", total, output)
			}
		}
		if output[0] != 111 || output[17] != 222 {
			t.Fatal("sentinels changed")
		}
		if total == 0 && state != (OpusT_ec_dec{Frng: 1 << 31, Fval: (1 << 31) - 1, Fnbits_total: 33}) {
			t.Fatal("zero pulses changed entropy state")
		}
	}
	var child int16 = 123
	decode_split(nil, &child, &child, nil, 0, nil)
	if child != 0 {
		t.Fatal("zero split did not clear aliased output")
	}
}

func TestICDFCorePointers(t *testing.T) {
	table := []uint8{250, 180, 100, 0}
	for symbol := int32(0); symbol < 4; symbol++ {
		dec := OpusT_ec_dec{Frng: 1 << 31, Fval: uint32(table[symbol]) * (1 << 23), Fnbits_total: 33}
		if got := ec_dec_icdf(nil, &dec, &table[0], 8); got != symbol || dec.Fval != 0 {
			t.Fatalf("symbol=%d got=%d state=%+v", symbol, got, dec)
		}
	}
}
