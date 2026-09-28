package opuscc

import "testing"

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
