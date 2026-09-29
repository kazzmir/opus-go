package opuscc

import (
	"runtime"
	"testing"
	"unsafe"
)

func TestEntropyEncUintPointers(t *testing.T) {
	totals := []uint32{2, 255, 256, 257, 511, 512, 513, 65535, 65536, 65537, 1 << 31, 0xffffffff}
	var enc OpusT_ec_enc
	// The context alone retains the Go output allocation through collection/stack growth.
	func() { b := make([]byte, 512); Opus_ec_enc_init(nil, &enc, unsafe.SliceData(b), 512) }()
	entropyInitGrowStack(12)
	runtime.GC()
	for _, total := range totals {
		for _, v := range []uint32{0, 1, total / 2, total - 1} {
			Opus_ec_enc_uint(nil, &enc, v, total)
		}
	}
	Opus_ec_enc_done(nil, &enc)
	if enc.Ferror1 != 0 {
		t.Fatal(enc)
	}
	var dec OpusT_ec_dec
	Opus_ec_dec_init(nil, &dec, enc.Fbuf, enc.Fstorage)
	for _, total := range totals {
		for _, want := range []uint32{0, 1, total / 2, total - 1} {
			if got := Opus_ec_dec_uint(nil, &dec, total); got != want {
				t.Fatal(total, got, want)
			}
		}
	}
}
