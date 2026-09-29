package opuscc

import "testing"

func TestEntropyEncICDF16Pointers(t *testing.T) {
	table := [5]uint16{77, 60000, 40000, 20000, 0}
	for s := int32(0); s < 4; s++ {
		var e OpusT_ec_enc
		Opus_ec_enc_init(nil, &e, nil, 0)
		want := e
		r := want.Frng >> 16
		if s == 0 {
			want.Frng -= r * uint32(table[1])
		} else {
			want.Fval += want.Frng - r*uint32(table[s])
			want.Frng = r * uint32(table[s]-table[s+1])
		}
		Opus_ec_enc_icdf16(nil, &e, s, &table[1], 16)
		if e != want || table != [5]uint16{77, 60000, 40000, 20000, 0} {
			t.Fatal(s, e, want)
		}
	}
	singleton := uint16(0)
	var e OpusT_ec_enc
	Opus_ec_enc_init(nil, &e, nil, 0)
	before := e
	Opus_ec_enc_icdf16(nil, &e, 0, &singleton, 16)
	if e != before {
		t.Fatal("single entry")
	}
}
