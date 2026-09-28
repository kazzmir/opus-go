package opuscc

import "testing"

func TestRangeLookupPointers(t *testing.T) {
	for _, tc := range []struct{ rng, val, ft, ext, symbol uint32 }{
		{100, 0, 10, 10, 9}, {100, 99, 10, 10, 0},
		{101, 100, 10, 10, 0}, {100, 42, 1, 100, 0},
		{0x80000000, 0x7fffffff, 32768, 65536, 0},
	} {
		dec := OpusT_ec_dec{Frng: tc.rng, Fval: tc.val, Fext: 99, Foffs: 7, Ferror1: -1}
		want := dec
		want.Fext = tc.ext
		got := Opus_ec_decode(nil, &dec, tc.ft)
		if got != tc.symbol || dec != want {
			t.Fatalf("%+v: symbol=%d state=%+v", tc, got, dec)
		}
	}
	for bits := uint32(0); bits <= 15; bits++ {
		for _, value := range []uint32{0, 123456, 0x7ffffffe} {
			dec := OpusT_ec_dec{Frng: 0x7fffffff, Fval: value}
			other := dec
			got := Opus_ec_decode_bin(nil, &dec, bits)
			want := Opus_ec_decode(nil, &other, 1<<bits)
			if got != want || dec != other {
				t.Fatalf("bits=%d val=%d: binary and uniform lookup differ", bits, value)
			}
		}
	}
}
