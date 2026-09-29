package opuscc

import "testing"

func TestEntropyPatchPointers(t *testing.T) {
	buf := [3]byte{77, 0xab, 88}
	enc := OpusT_ec_enc{Fbuf: &buf[1], Foffs: 1, Frem: 3, Fval: 123, Ferror1: 7}
	want := enc
	Opus_ec_enc_patch_initial_bits(nil, &enc, 5, 3)
	if enc != want || buf != [3]byte{77, 0xab, 88} {
		t.Fatal("finalized", enc, buf)
	}
	Opus_ec_enc_patch_initial_bits(nil, &enc, 0, 3)
	if buf[1] != 0xb || buf[0] != 77 || buf[2] != 88 {
		t.Fatal("mask", buf)
	}
	enc = OpusT_ec_enc{Frem: 0xab, Fval: 123, Ferror1: 7}
	Opus_ec_enc_patch_initial_bits(nil, &enc, 0, 3)
	if enc.Frem != 0xb || enc.Fval != 123 || enc.Ferror1 != 7 {
		t.Fatal("pending", enc)
	}
	enc = OpusT_ec_enc{Frem: -1, Frng: 1 << 28, Fval: 0x7fffffff}
	Opus_ec_enc_patch_initial_bits(nil, &enc, 0, 3)
	if enc.Fval != 0xfffffff || enc.Ferror1 != 0 {
		t.Fatal("interval", enc)
	}
	enc = OpusT_ec_enc{Frem: -1, Frng: 1<<28 + 1, Fval: 123}
	Opus_ec_enc_patch_initial_bits(nil, &enc, 0, 3)
	if enc.Fval != 123 || enc.Ferror1 != -1 {
		t.Fatal("error", enc)
	}
}
