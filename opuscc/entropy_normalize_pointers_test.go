package opuscc

import (
	"runtime"
	"testing"
	"unsafe"
)

func TestEntropyByteReadPointers(t *testing.T) {
	data := []byte{77, 13, 21, 34, 88}
	dec := OpusT_ec_dec{Fbuf: &data[1], Fstorage: 3}
	entropyInitGrowStack(12)
	runtime.GC()
	for _, want := range []int32{13, 21, 34, 0, 0} {
		if got := ec_read_byte(&dec); got != want {
			t.Fatal("front byte", got, want)
		}
	}
	for _, want := range []int32{34, 21, 13, 0, 0} {
		if got := ec_read_byte_from_end(&dec); got != want {
			t.Fatal("tail byte", got, want)
		}
	}
	if dec.Foffs != 3 || dec.Fend_offs != 3 || data[0] != 77 || data[4] != 88 {
		t.Fatal("counters/guards")
	}
	empty := OpusT_ec_dec{}
	if ec_read_byte(&empty) != 0 || ec_read_byte_from_end(&empty) != 0 {
		t.Fatal("empty")
	}
	single := byte(19)
	prefix := OpusT_ec_dec{Fbuf: &single, Fstorage: 100}
	if ec_read_byte(&prefix) != 19 {
		t.Fatal("only consumed prefix")
	}
	alias := OpusT_ec_dec{Fstorage: 4}
	alias.Fbuf = (*byte)(unsafe.Pointer(&alias.Foffs))
	got := ec_read_byte(&alias)
	want := int32(unsafe.Slice((*byte)(unsafe.Pointer(&alias.Foffs)), 4)[0])
	if alias.Foffs != 1 || got != want {
		t.Fatal("counter mutation before byte read")
	}
}

func TestEntropyNormalizePointers(t *testing.T) {
	e := OpusT_ec_enc{Frng: 1<<23 + 1, Fval: 123, Frem: -1}
	before := e
	ec_enc_normalize(nil, &e)
	if e != before {
		t.Fatal("identity")
	}
	b := [5]byte{77, 0, 0, 0, 88}
	e = OpusT_ec_enc{Fbuf: &b[1], Fstorage: 3, Frng: 1, Fval: 0, Frem: -1, Fnbits_total: 33}
	ec_enc_normalize(nil, &e)
	if e.Frng != 1<<24 || e.Fnbits_total != 57 || e.Foffs != 2 || e.Frem != 0 || b[0] != 77 || b[4] != 88 {
		t.Fatal(e, b)
	}
	e = OpusT_ec_enc{Frng: 1, Fval: 0, Frem: 0}
	ec_enc_normalize(nil, &e)
	if e.Ferror1 != -1 || e.Frng != 1<<24 {
		t.Fatal("exhausted", e)
	}
}
