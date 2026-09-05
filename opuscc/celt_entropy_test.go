package opuscc

import (
	"testing"
	"unsafe"
)

func entropyBufferPointer(buffer []byte) uintptr {
	return uintptr(unsafe.Pointer(&buffer[0]))
}

func TestEntropyEncoderFieldAccesses(t *testing.T) {
	buffer := make([]byte, 32)
	var enc OpusT_ec_enc
	Opus_ec_enc_init(nil, uintptr(unsafe.Pointer(&enc)), entropyBufferPointer(buffer), uint32(len(buffer)))

	if got := ec_write_byte(nil, uintptr(unsafe.Pointer(&enc)), 0x3b); got != 0 {
		t.Fatalf("write first byte: got %d, want 0", got)
	}
	if got := ec_write_byte_at_end(nil, uintptr(unsafe.Pointer(&enc)), 0xc4); got != 0 {
		t.Fatalf("write final byte: got %d, want 0", got)
	}
	if enc.Foffs != 1 || enc.Fend_offs != 1 || buffer[0] != 0x3b || buffer[len(buffer)-1] != 0xc4 {
		t.Fatalf("unexpected byte writes: offs=%d end_offs=%d buffer=% x", enc.Foffs, enc.Fend_offs, buffer)
	}

	enc = OpusT_ec_enc{
		Fbuf:     entropyBufferPointer(buffer),
		Fstorage: uint32(len(buffer)),
		Fext:     2,
		Frem:     0x44,
	}
	ec_enc_carry_out(nil, uintptr(unsafe.Pointer(&enc)), 0x100)
	if enc.Foffs != 3 || enc.Fext != 0 || enc.Frem != 0 || enc.Ferror1 != 0 {
		t.Fatalf("unexpected carry state: %+v", enc)
	}
	if got, want := buffer[:3], []byte{0x45, 0x00, 0x00}; string(got) != string(want) {
		t.Fatalf("carry output: got % x, want % x", got, want)
	}

	buffer = make([]byte, 32)
	Opus_ec_enc_init(nil, uintptr(unsafe.Pointer(&enc)), entropyBufferPointer(buffer), uint32(len(buffer)))
	icdf8 := []byte{240, 180, 100, 0}
	icdf16 := []uint16{60000, 40000, 20000, 0}
	Opus_ec_encode(nil, uintptr(unsafe.Pointer(&enc)), 3, 7, 13)
	Opus_ec_encode_bin(nil, uintptr(unsafe.Pointer(&enc)), 5, 10, 4)
	Opus_ec_enc_icdf(nil, uintptr(unsafe.Pointer(&enc)), 2, entropyBufferPointer(icdf8), 8)
	Opus_ec_enc_icdf16(nil, uintptr(unsafe.Pointer(&enc)), 1, uintptr(unsafe.Pointer(&icdf16[0])), 16)
	Opus_ec_enc_bits(nil, uintptr(unsafe.Pointer(&enc)), 0xbeef, 16)
	Opus_ec_enc_done(nil, uintptr(unsafe.Pointer(&enc)))

	if got, want := buffer[0], byte(0x5c); got != want {
		t.Fatalf("encoded first byte: got %#x, want %#x", got, want)
	}
	if got, want := buffer[len(buffer)-2:], []byte{0xbe, 0xef}; string(got) != string(want) {
		t.Fatalf("encoded tail: got % x, want % x", got, want)
	}
	for i, value := range buffer[1 : len(buffer)-2] {
		if value != 0 {
			t.Fatalf("encoded padding at %d: got %#x, want 0", i+1, value)
		}
	}
	if enc.Foffs != 1 || enc.Fend_offs != 2 || enc.Fend_window != 0xbeef ||
		enc.Fnend_bits != 16 || enc.Fnbits_total != 49 || enc.Frng != 19680000 ||
		enc.Fval != 768851182 || enc.Fext != 0 || enc.Frem != 0 || enc.Ferror1 != 0 {
		t.Fatalf("unexpected encoder state: %+v", enc)
	}
}

func TestEntropyDecoderFieldAccesses(t *testing.T) {
	buffer := []byte{0x3b, 0xa1, 0x7d, 0xc4, 0x19}
	dec := OpusT_ec_dec{
		Fbuf:     entropyBufferPointer(buffer),
		Fstorage: uint32(len(buffer)),
		Foffs:    1,
	}
	if got := ec_read_byte(&dec); got != 0xa1 {
		t.Fatalf("read byte: got %#x, want %#x", got, 0xa1)
	}
	if got := ec_read_byte(&dec); got != 0x7d {
		t.Fatalf("read byte: got %#x, want %#x", got, 0x7d)
	}
	if got := ec_read_byte_from_end(&dec); got != 0x19 {
		t.Fatalf("read final byte: got %#x, want %#x", got, 0x19)
	}
	if got := ec_read_byte_from_end(&dec); got != 0xc4 {
		t.Fatalf("read penultimate byte: got %#x, want %#x", got, 0xc4)
	}
	if dec.Foffs != 3 || dec.Fend_offs != 2 {
		t.Fatalf("unexpected read offsets: offs=%d end_offs=%d", dec.Foffs, dec.Fend_offs)
	}

	dec = OpusT_ec_dec{
		Fbuf:         entropyBufferPointer(buffer),
		Fstorage:     uint32(len(buffer)),
		Frng:         12345,
		Fval:         0x12345,
		Frem:         0x5a,
		Fnbits_total: 17,
	}
	ec_dec_normalize(nil, &dec)
	if dec.Foffs != 2 || dec.Fnbits_total != 33 || dec.Frng != 809041920 ||
		dec.Fval != 591782447 || dec.Frem != 0xa1 {
		t.Fatalf("unexpected normalized decoder state: %+v", dec)
	}

	encoded := []byte{0xd7, 0x4a, 0x91, 0x2e, 0xbc, 0x63}
	Opus_ec_dec_init(nil, uintptr(unsafe.Pointer(&dec)), entropyBufferPointer(encoded), uint32(len(encoded)))
	symbol := Opus_ec_decode(nil, uintptr(unsafe.Pointer(&dec)), 13)
	Opus_ec_dec_update(nil, uintptr(unsafe.Pointer(&dec)), symbol, symbol+1, 13)
	bits := Opus_ec_dec_bits(nil, uintptr(unsafe.Pointer(&dec)), 11)
	if symbol != 10 || bits != 0x463 {
		t.Fatalf("decoded values: symbol=%d bits=%#x, want symbol=10 bits=0x463", symbol, bits)
	}
	if dec.Fend_offs != 4 || dec.Fend_window != 1189335 || dec.Fnend_bits != 21 ||
		dec.Fnbits_total != 44 || dec.Foffs != 4 || dec.Frng != 165191049 ||
		dec.Fval != 11107414 || dec.Fext != 165191049 || dec.Frem != 46 || dec.Ferror1 != 0 {
		t.Fatalf("unexpected decoded state: %+v", dec)
	}

	dec = OpusT_ec_dec{
		Fbuf:         entropyBufferPointer(buffer),
		Fstorage:     uint32(len(buffer)),
		Frng:         1000000000,
		Fval:         500000000,
		Fnbits_total: 9,
	}
	icdf8 := []byte{250, 180, 100, 0}
	icdf16 := []uint16{60000, 40000, 20000, 0}
	bit := Opus_ec_dec_bit_logp(nil, uintptr(unsafe.Pointer(&dec)), 4)
	symbol8 := Opus_ec_dec_icdf(nil, uintptr(unsafe.Pointer(&dec)), entropyBufferPointer(icdf8), 8)
	symbol16 := Opus_ec_dec_icdf16(nil, uintptr(unsafe.Pointer(&dec)), uintptr(unsafe.Pointer(&icdf16[0])), 16)
	if bit != 0 || symbol8 != 2 || symbol16 != 3 {
		t.Fatalf("decoded coding results: bit=%d symbol8=%d symbol16=%d, want 0, 2, 3", bit, symbol8, symbol16)
	}
	if dec.Fnbits_total != 9 || dec.Frng != 89400000 || dec.Fval != 71289100 {
		t.Fatalf("unexpected decoder coding state: %+v", dec)
	}
}
