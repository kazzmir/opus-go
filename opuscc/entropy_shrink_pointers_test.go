package opuscc

import (
	"slices"
	"testing"
)

func TestEntropyShrinkPointers(t *testing.T) {
	buf := [10]byte{99, 1, 2, 3, 4, 5, 6, 7, 8, 88}
	enc := OpusT_ec_enc{Fbuf: &buf[1], Fstorage: 8, Foffs: 2, Fend_offs: 4, Frng: 123, Fval: 456}
	want := enc
	want.Fstorage = 6
	Opus_ec_enc_shrink(nil, &enc, 6)
	if enc != want || buf != [10]byte{99, 1, 2, 5, 6, 7, 8, 7, 8, 88} {
		t.Fatal(enc, buf)
	}
	before := slices.Clone(buf[:])
	Opus_ec_enc_shrink(nil, &enc, 6)
	if !slices.Equal(buf[:], before) {
		t.Fatal("identity")
	}
	empty := OpusT_ec_enc{Fstorage: 32}
	Opus_ec_enc_shrink(nil, &empty, 0)
	if empty != (OpusT_ec_enc{}) {
		t.Fatal("empty")
	}
}
