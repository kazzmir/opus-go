package opuscc

import (
	"testing"
	"unsafe"
)

func TestEntropyDonePartialPointers(t *testing.T) {
	for _, full := range []bool{false, true} {
		buffer := [3]byte{77, 0xaa, 88}
		var enc OpusT_ec_enc
		Opus_ec_enc_init(nil, &enc, &buffer[1], 1)
		enc.Fend_window = 5
		enc.Fnend_bits = 3
		wantByte, wantError := byte(5), int32(0)
		if full {
			enc.Foffs = 1
			wantByte = 0xaa
			wantError = -1
		}
		before := enc
		Opus_ec_enc_done(nil, &enc)
		if buffer != [3]byte{77, wantByte, 88} || enc.Ferror1 != wantError || enc.Fend_window != before.Fend_window || enc.Fnend_bits != before.Fnend_bits || enc.Foffs != before.Foffs || enc.Fend_offs != before.Fend_offs {
			t.Fatal("partial byte/collision", full, enc, buffer)
		}
	}
	var empty OpusT_ec_enc
	Opus_ec_enc_init(nil, &empty, nil, 0)
	empty.Fend_window = 5
	empty.Fnend_bits = 3
	Opus_ec_enc_done(nil, &empty)
	if empty.Ferror1 != -1 || empty.Fend_window != 5 || empty.Fnend_bits != 3 {
		t.Fatal("partial empty", empty)
	}
}

func TestEntropyDonePointers(t *testing.T) {
	b := [8]byte{77, 99, 99, 99, 99, 99, 99, 88}
	var e OpusT_ec_enc
	Opus_ec_enc_init(nil, &e, &b[1], 6)
	Opus_ec_enc_bits(nil, &e, 0xabc, 12)
	window, used := e.Fend_window, e.Fnend_bits
	Opus_ec_enc_done(nil, &e)
	if b != [8]byte{77, 0, 0, 0, 0, 0xa, 0xbc, 88} || e.Ferror1 != 0 || e.Fend_window != window || e.Fnend_bits != used {
		t.Fatal(e, b)
	}
	empty := OpusT_ec_enc{}
	Opus_ec_enc_init(nil, &empty, nil, 0)
	Opus_ec_enc_done(nil, &empty)
	if empty.Ferror1 != 0 {
		t.Fatal("empty", empty)
	}
	// Typed encoder-to-decoder round trip including both range and raw bits.
	packet := make([]byte, 64)
	Opus_ec_enc_init(nil, &e, unsafe.SliceData(packet), 64)
	for i := uint32(0); i < 20; i++ {
		Opus_ec_encode_bin(nil, &e, i%16, i%16+1, 4)
		Opus_ec_enc_bits(nil, &e, i%8, 3)
	}
	Opus_ec_enc_done(nil, &e)
	if e.Ferror1 != 0 {
		t.Fatal(e)
	}
	var d OpusT_ec_dec
	Opus_ec_dec_init(nil, &d, unsafe.SliceData(packet), 64)
	for i := uint32(0); i < 20; i++ {
		s := Opus_ec_decode_bin(nil, &d, 4)
		if s != i%16 {
			t.Fatal("range", i, s)
		}
		Opus_ec_dec_update(nil, &d, s, s+1, 16)
		if v := Opus_ec_dec_bits(nil, &d, 3); v != i%8 {
			t.Fatal("tail", i, v)
		}
	}
}
