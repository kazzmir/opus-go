package opuscc

import (
	"runtime"
	"testing"
	"unsafe"
)

// This live numeric-state alias checks the existing Go cached-load contract,
// not parity with upstream C's intervening table rereads.
func TestEntropyEncICDFCachedPointers(t *testing.T) {
	for _, wide := range []bool{false, true} {
		enc := OpusT_ec_enc{Frng: 1 << 31, Fval: 0x006400c8, Frem: -1}
		table8 := (*byte)(unsafe.Pointer(&enc.Fval))
		table16 := (*uint16)(unsafe.Pointer(&enc.Fval))
		var previous, current uint32
		if wide {
			row := unsafe.Slice(table16, 2)
			previous, current = uint32(row[0]), uint32(row[1])
		} else {
			row := unsafe.Slice(table8, 2)
			previous, current = uint32(row[0]), uint32(row[1])
		}
		want := enc
		r := want.Frng >> 24
		want.Fval += want.Frng - r*previous
		want.Frng = r * (previous - current)
		ec_enc_normalize(nil, &want)
		entropyInitGrowStack(12)
		runtime.GC()
		if wide {
			Opus_ec_enc_icdf16(nil, &enc, 1, table16, 24)
		} else {
			Opus_ec_enc_icdf(nil, &enc, 1, table8, 24)
		}
		if enc != want {
			t.Fatal("cached rate row", wide, enc, want)
		}
	}
}

func TestEntropyEncICDFPointers(t *testing.T) {
	table := [5]byte{77, 240, 180, 100, 0}
	var e OpusT_ec_enc
	var b [128]byte
	Opus_ec_enc_init(nil, &e, &b[0], 128)
	for i := 0; i < 100; i++ {
		Opus_ec_enc_icdf(nil, &e, int32(i%4), &table[1], 8)
	}
	Opus_ec_enc_done(nil, &e)
	if e.Ferror1 != 0 || table != [5]byte{77, 240, 180, 100, 0} {
		t.Fatal(e, table)
	}
	var d OpusT_ec_dec
	Opus_ec_dec_init(nil, &d, &b[0], 128)
	for i := 0; i < 100; i++ {
		if s := ec_dec_icdf(nil, &d, &table[1], 8); s != int32(i%4) {
			t.Fatal(i, s)
		}
	}
	single := byte(0)
	Opus_ec_enc_init(nil, &e, nil, 0)
	before := e
	Opus_ec_enc_icdf(nil, &e, 0, &single, 8)
	if e != before {
		t.Fatal("singleton")
	}
}
