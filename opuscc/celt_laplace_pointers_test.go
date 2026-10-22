package opuscc

import (
	"testing"
)

// From laplace_decode_pointers_test.go.
func TestLaplaceDecodePointers(t *testing.T) {
	const fs = uint32(16000)
	const decay = int32(8000)
	freq1 := ((32768 - 32 - fs) * uint32(16384-decay) >> 15) + 1
	for _, tc := range []struct {
		fm   uint32
		want int32
	}{{0, 0}, {fs - 1, 0}, {fs, -1}, {fs + freq1 - 1, -1}, {fs + freq1, 1}, {fs + 2*freq1 - 1, 1}} {
		dec := OpusT_ec_dec{Frng: 1 << 31, Fval: (32767 - tc.fm) << 16, Fnbits_total: 33, Fend_window: 123, Fnend_bits: 7, Ferror1: 9}
		got := Opus_ec_laplace_decode(nil, &dec, fs, decay)
		if got != tc.want || dec.Frng <= 1<<23 || dec.Fval >= dec.Frng || dec.Fend_window != 123 || dec.Fnend_bits != 7 || dec.Ferror1 != 9 {
			t.Fatalf("fm=%d got=%d want=%d state=%+v", tc.fm, got, tc.want, dec)
		}
	}
}

// From laplace_encode_pointers_test.go.
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

// From laplace_encode_p0_pointers_test.go.
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

// From laplace_p0_pointers_test.go.
func TestLaplaceP0Pointers(t *testing.T) {
	for _, tc := range []struct {
		value uint32
		want  int32
	}{{32767 << 16, 0}, {12000 << 16, 1}, {4000 << 16, -1}} {
		dec := OpusT_ec_dec{Frng: 1 << 31, Fval: tc.value, Fnbits_total: 33, Fext: 123, Fend_window: 456, Fnend_bits: 7, Ferror1: 9}
		if got := Opus_ec_laplace_decode_p0(nil, &dec, 16000, 0); got != tc.want || dec.Fext != 123 || dec.Fend_window != 456 || dec.Fnend_bits != 7 || dec.Ferror1 != 9 {
			t.Fatalf("value=%d got=%d want=%d state=%+v", tc.value, got, tc.want, dec)
		}
	}
	dec := OpusT_ec_dec{Frng: 1 << 31, Fval: 1234567, Fnbits_total: 33}
	before := dec
	if got := Opus_ec_laplace_decode_p0(nil, &dec, 32768, 12000); got != 0 || dec != before {
		t.Fatal("certain-zero distribution changed context")
	}
}
