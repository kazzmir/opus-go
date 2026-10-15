package opuscc

import (
	"runtime"
	"testing"
	"unsafe"
)

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

func TestTFDecodeSpanPointers(t *testing.T) {
	for LM := int32(0); LM <= 3; LM++ {
		for transient := int32(0); transient <= 1; transient++ {
			for _, size := range []int{0, 1, 32} {
				for _, bounds := range [][2]int32{{0, 0}, {1, 5}, {5, 5}} {
					data := make([]byte, size)
					for i := range data {
						data[i] = byte(i*37 + 19)
					}
					var dec OpusT_ec_dec
					Opus_ec_dec_init(nil, &dec, unsafe.SliceData(data), uint32(size))
					owned := dec
					a := [6]int32{77, 9, 9, 9, 9, 88}
					b := a
					entropyInitGrowStack(12)
					runtime.GC()
					tf_decode(nil, bounds[0], bounds[1], transient, &a[0], LM, &dec)
					tfDecodeSpan(nil, bounds[0], bounds[1], transient, b[:], LM, &owned)
					if a != b || dec != owned || a[0] != 77 || a[5] != 88 {
						t.Fatal("TF span/context/guards", LM, transient, size, bounds)
					}
				}
			}
		}
	}
	var a OpusT_ec_dec
	Opus_ec_dec_init(nil, &a, nil, 0)
	b := a
	tf_decode(nil, 0, 1, 0, &a.Fnbits_total, 0, &a)
	tfDecodeSpan(nil, 0, 1, 0, unsafe.Slice(&b.Fnbits_total, 1), 0, &b)
	if a != b {
		t.Fatal("TF scalar entropy-field alias")
	}
	Opus_ec_dec_init(nil, &a, nil, 0)
	b = a
	tf_decode(nil, 0, 0, 0, nil, 0, &a)
	tfDecodeSpan(nil, 0, 0, 0, nil, 0, &b)
	if a != b {
		t.Fatal("nil TF zero extent")
	}
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
