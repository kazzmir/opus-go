package opuscc

import (
	"runtime"
	"testing"
	"unsafe"
)

func TestICDF8OffsetPointers(t *testing.T) {
	owner := [6]byte{77, 88, 240, 180, 100, 0}
	buffer := make([]byte, 128)
	var encoder OpusT_ec_enc
	Opus_ec_enc_init(nil, &encoder, &buffer[0], 128)
	for i := 0; i < 200; i++ {
		Opus_ec_enc_icdf(nil, &encoder, int32(i%4), &owner[2], 8)
	}
	Opus_ec_enc_done(nil, &encoder)
	var expected OpusT_ec_dec
	Opus_ec_dec_init(nil, &expected, &buffer[0], 128)
	positive, negative := expected, expected
	entropyInitGrowStack(12)
	runtime.GC()
	for i := 0; i < 200; i++ {
		want := Opus_ec_dec_icdf(nil, &expected, &owner[2], 8)
		got := ecDecICDFAt(nil, &positive, &owner[0], 2, 8)
		signed := ecDecICDFAt(nil, &negative, &owner[4], -2, 8)
		if want != int32(i%4) || got != want || signed != want || positive != expected || negative != expected {
			t.Fatal("numeric ICDF symbol/state", i, got, signed, want)
		}
	}
	if owner != [6]byte{77, 88, 240, 180, 100, 0} {
		t.Fatal("table guards")
	}
}

func TestICDF8Pointers(t *testing.T) {
	table := [4]byte{240, 180, 100, 0}
	var e OpusT_ec_enc
	func() { b := make([]byte, 128); Opus_ec_enc_init(nil, &e, unsafe.SliceData(b), 128) }()
	for i := 0; i < 200; i++ {
		Opus_ec_enc_icdf(nil, &e, int32(i%4), &table[0], 8)
	}
	Opus_ec_enc_done(nil, &e)
	var d OpusT_ec_dec
	Opus_ec_dec_init(nil, &d, e.Fbuf, e.Fstorage)
	e = OpusT_ec_enc{}
	entropyInitGrowStack(12)
	runtime.GC()
	for i := 0; i < 200; i++ {
		if got := Opus_ec_dec_icdf(nil, &d, &table[0], 8); got != int32(i%4) {
			t.Fatal(i, got)
		}
	}
	single := byte(0)
	before := d
	Opus_ec_dec_icdf(nil, &d, &single, 8)
	if d != before {
		t.Fatal("singleton", d, before)
	}
}
