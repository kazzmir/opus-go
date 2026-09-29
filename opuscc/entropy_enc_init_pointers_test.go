package opuscc

import (
	"runtime"
	"testing"
)

//go:noinline
func ownedEntropyEncoder() OpusT_ec_enc {
	var enc OpusT_ec_enc
	buf := make([]byte, 7)
	for i := range buf {
		buf[i] = 0xa5
	}
	Opus_ec_enc_init(nil, &enc, &buf[0], uint32(len(buf)))
	return enc
}

func TestEntropyEncInitPointers(t *testing.T) {
	enc := OpusT_ec_enc{Fend_offs: 3, Fval: 4, Fext: 5, Ferror1: -1}
	Opus_ec_enc_init(nil, &enc, nil, 0)
	if enc != (OpusT_ec_enc{Fnbits_total: 33, Frng: 0x80000000, Frem: -1}) {
		t.Fatal(enc)
	}
	enc = ownedEntropyEncoder()
	entropyInitGrowStack(128)
	runtime.GC()
	if enc.Fbuf == nil || *enc.Fbuf != 0xa5 {
		t.Fatal("owner lost")
	}
	if ec_write_byte(nil, &enc, 123) != 0 || *enc.Fbuf != 123 {
		t.Fatal("write")
	}
}
