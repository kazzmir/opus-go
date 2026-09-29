package opuscc

import (
	"runtime"
	"testing"
)

//go:noinline
func entropyInitGrowStack(depth int) byte {
	var pad [1024]byte
	for i := range pad {
		pad[i] = byte(i + depth)
	}
	var result byte
	if depth > 0 {
		result = entropyInitGrowStack(depth - 1)
	}
	result += pad[depth%len(pad)]
	runtime.KeepAlive(&pad)
	return result
}

//go:noinline
func entropyOwnedPacket() OpusT_ec_dec {
	packet := []byte{3, 57, 127, 255, 1, 0xa7}
	var dec OpusT_ec_dec
	Opus_ec_dec_init(nil, &dec, &packet[0], uint32(len(packet)))
	return dec
}

func TestEntropyInitOwnership(t *testing.T) {
	// Only the context retains the packet after the helper returns.
	for i := 0; i < 8; i++ {
		dec := entropyOwnedPacket()
		entropyInitGrowStack(128)
		runtime.GC()
		if got := Opus_ec_dec_bits(nil, &dec, 8); got != 0xa7 {
			t.Fatal(got)
		}
	}
}

func TestEntropyInitPointers(t *testing.T) {
	empty := OpusT_ec_dec{Fext: 0xcafebabe}
	Opus_ec_dec_init(nil, &empty, nil, 0)
	want := OpusT_ec_dec{Fext: 0xcafebabe, Fnbits_total: 33, Frng: 0x80000000, Fval: 0x7fffffff}
	if empty != want {
		t.Fatalf("empty: %+v", empty)
	}
	packet := [6]byte{3, 57, 127, 255, 1, 0xa7}
	dec := OpusT_ec_dec{Fext: 0xcafebabe, Fend_offs: 123, Fend_window: 456, Fnend_bits: 7, Ferror1: -1}
	Opus_ec_dec_init(nil, &dec, &packet[0], uint32(len(packet)))
	if dec.Fext != 0xcafebabe || dec.Fnbits_total != 33 || dec.Foffs != 4 || dec.Frem != 255 || dec.Frng != 0x80000000 || dec.Ferror1 != 0 || dec.Fend_offs != 0 || dec.Fend_window != 0 || dec.Fnend_bits != 0 {
		t.Fatalf("state: %+v", dec)
	}
	entropyInitGrowStack(128)
	runtime.GC()
	if got := Opus_ec_dec_bits(nil, &dec, 8); got != 0xa7 {
		t.Fatal("packet address", got)
	}
	runtime.KeepAlive(&packet)
}
