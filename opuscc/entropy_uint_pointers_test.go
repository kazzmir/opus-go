package opuscc

import (
	"math/bits"
	"testing"
)

func TestEntropyUintPointers(t *testing.T) {
	for _, ft := range []uint32{2, 3, 255, 256, 257, 258, 65535, 65536, 65537, 1 << 24, 0xffffffff} {
		dec := OpusT_ec_dec{Frng: 1 << 31, Fval: (1 << 31) - 1, Fnbits_total: 33, Fnend_bits: 32, Ferror1: 7}
		got := Opus_ec_dec_uint(nil, &dec, ft)
		consumed := max(0, bits.Len32(ft-1)-8)
		if got != 0 || dec.Ferror1 != 7 || dec.Fnend_bits != 32-int32(consumed) || dec.Frng <= 1<<23 {
			t.Fatalf("ft=%d got=%d state=%+v", ft, got, dec)
		}
	}
	// Highest range symbol plus a nonzero tail can represent an invalid value.
	for _, tail := range []uint32{0, 1} {
		dec := OpusT_ec_dec{Frng: 1 << 31, Fval: 0, Fnbits_total: 33, Fnend_bits: 1, Fend_window: tail, Ferror1: 7}
		got := Opus_ec_dec_uint(nil, &dec, 257)
		wantError := int32(7)
		if tail == 1 {
			wantError = 1
		}
		if got != 256 || dec.Ferror1 != wantError || dec.Fnend_bits != 0 || dec.Fend_window != 0 || dec.Fnbits_total != 34 {
			t.Fatalf("tail=%d got=%d state=%+v", tail, got, dec)
		}
	}
}
