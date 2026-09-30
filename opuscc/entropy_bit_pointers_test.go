package opuscc

import "testing"

func TestTFDecodePointers(t *testing.T) {
	dec := OpusT_ec_dec{}
	Opus_ec_dec_init(nil, &dec, nil, 0)
	out := [6]int32{77, 9, 9, 9, 9, 88}
	tf_decode(nil, 1, 5, 0, &out[0], 0, &dec)
	if out != [6]int32{77, 0, 0, 0, 0, 88} {
		t.Fatal(out)
	}
	dec = OpusT_ec_dec{}
	Opus_ec_dec_init(nil, &dec, nil, 0)
	tf_decode(nil, 0, 0, 0, nil, 0, &dec)
}

func TestEntropyBitPointers(t *testing.T) {
	for _, logp := range []uint32{1, 2, 7, 8, 15, 24, 31} {
		const rng = uint32(1 << 31)
		split := rng >> logp
		for _, value := range []uint32{0, split - 1, split, split + 1, rng - 1} {
			dec := OpusT_ec_dec{Frng: rng, Fval: value, Fnbits_total: 33, Fext: 123, Fend_window: 456, Fnend_bits: 9, Ferror1: 7}
			want := dec
			var bit int32
			if value < split {
				bit = 1
				want.Frng = split
			} else {
				want.Frng = rng - split
				want.Fval = value - split
			}
			for want.Frng <= 1<<23 {
				want.Frng <<= 8
				want.Fval = ((want.Fval << 8) + 255) & 0x7fffffff
				want.Fnbits_total += 8
			}
			if got := Opus_ec_dec_bit_logp(nil, &dec, logp); got != bit || dec != want {
				t.Fatalf("logp=%d value=%d bit=%d want=%d state=%+v want=%+v", logp, value, got, bit, dec, want)
			}
		}
	}
}
