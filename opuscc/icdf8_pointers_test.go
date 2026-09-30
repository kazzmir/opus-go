package opuscc

import (
	"runtime"
	"testing"
	"unsafe"
)

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
