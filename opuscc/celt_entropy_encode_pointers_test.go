package opuscc

import (
	"runtime"
	"slices"
	"testing"
	"unsafe"
)

// From entropy_enc_bit_pointers_test.go.
func TestEntropyEncBitPointers(t *testing.T) {
	for _, value := range []int32{0, 1, -1, 7} {
		var e OpusT_ec_enc
		Opus_ec_enc_init(nil, &e, nil, 0)
		Opus_ec_enc_bit_logp(nil, &e, value, 3)
		if value == 0 {
			if e.Fval != 0 || e.Frng != 0x70000000 {
				t.Fatal(e)
			}
		} else {
			if e.Fval != 0x70000000 || e.Frng != 0x10000000 {
				t.Fatal(e)
			}
		}
	}
	var b [3]byte
	e := OpusT_ec_enc{}
	Opus_ec_enc_init(nil, &e, &b[0], 3)
	for i := 0; i < 100; i++ {
		Opus_ec_enc_bit_logp(nil, &e, int32(i%2), 15)
	}
	if e.Foffs != 3 || e.Ferror1 != -1 || e.Frng == 0 {
		t.Fatal(e)
	}
}

// From entropy_enc_bits_pointers_test.go.
func TestEntropyEncBitsPointers(t *testing.T) {
	b := [6]byte{77, 0, 0, 0, 0, 88}
	var e OpusT_ec_enc
	Opus_ec_enc_init(nil, &e, &b[1], 4)
	Opus_ec_enc_bits(nil, &e, 0xabcdef, 24)
	Opus_ec_enc_bits(nil, &e, 0x12345, 20)
	if b != [6]byte{77, 0, 0xab, 0xcd, 0xef, 88} || e.Fend_offs != 3 || e.Fend_window != 0x12345 || e.Fnend_bits != 20 || e.Fnbits_total != 77 {
		t.Fatal(e, b)
	}
	Opus_ec_enc_bits(nil, &e, 0x1ffffff, 25)
	if e.Fend_offs != 4 || e.Ferror1 != -1 || b[0] != 77 || b[5] != 88 {
		t.Fatal("exhausted", e, b)
	}
	e = OpusT_ec_enc{Fnbits_total: 2147483647}
	Opus_ec_enc_bits(nil, &e, 1, 1)
	if e.Fnbits_total != -2147483648 || e.Fend_window != 1 {
		t.Fatal("wrap", e)
	}
}

// From entropy_enc_icdf_pointers_test.go.
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

// From entropy_enc_icdf16_pointers_test.go.
func TestEntropyEncICDF16Pointers(t *testing.T) {
	table := [5]uint16{77, 60000, 40000, 20000, 0}
	for s := int32(0); s < 4; s++ {
		var e OpusT_ec_enc
		Opus_ec_enc_init(nil, &e, nil, 0)
		want := e
		r := want.Frng >> 16
		if s == 0 {
			want.Frng -= r * uint32(table[1])
		} else {
			want.Fval += want.Frng - r*uint32(table[s])
			want.Frng = r * uint32(table[s]-table[s+1])
		}
		Opus_ec_enc_icdf16(nil, &e, s, &table[1], 16)
		if e != want || table != [5]uint16{77, 60000, 40000, 20000, 0} {
			t.Fatal(s, e, want)
		}
	}
	singleton := uint16(0)
	var e OpusT_ec_enc
	Opus_ec_enc_init(nil, &e, nil, 0)
	before := e
	Opus_ec_enc_icdf16(nil, &e, 0, &singleton, 16)
	if e != before {
		t.Fatal("single entry")
	}
}

// From entropy_enc_init_pointers_test.go.
//
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

// From entropy_enc_uint_pointers_test.go.
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

// From entropy_encode_pointers_test.go.
func TestEntropyEncodePointers(t *testing.T) {
	var e OpusT_ec_enc
	Opus_ec_enc_init(nil, &e, nil, 0)
	before := e
	Opus_ec_encode(nil, &e, 0, 13, 13)
	if e != before {
		t.Fatal("full interval")
	}
	Opus_ec_encode(nil, &e, 3, 7, 13)
	if e.Fval != 495573158 || e.Frng != 660764196 {
		t.Fatal(e)
	}
	var b [4]byte
	Opus_ec_enc_init(nil, &e, &b[0], 4)
	for i := 0; i < 64; i++ {
		Opus_ec_encode(nil, &e, 1, 2, 256)
	}
	if e.Ferror1 != -1 || e.Foffs != 4 || e.Frng == 0 {
		t.Fatal("exhausted", e)
	}
}

// From entropy_encode_bin_pointers_test.go.
func TestEntropyEncodeBinPointers(t *testing.T) {
	var e OpusT_ec_enc
	Opus_ec_enc_init(nil, &e, nil, 0)
	before := e
	Opus_ec_encode_bin(nil, &e, 0, 1, 0)
	if e != before {
		t.Fatal("full interval")
	}
	Opus_ec_encode_bin(nil, &e, 5, 10, 4)
	if e.Fval != 671088640 || e.Frng != 671088640 {
		t.Fatal(e)
	}
	var b [4]byte
	Opus_ec_enc_init(nil, &e, &b[0], 4)
	for i := 0; i < 64; i++ {
		Opus_ec_encode_bin(nil, &e, 1, 2, 15)
	}
	if e.Ferror1 != -1 || e.Foffs != 4 || e.Frng == 0 {
		t.Fatal("exhausted", e)
	}
}

// From entropy_carry_pointers_test.go.
func TestEntropyCarryPointers(t *testing.T) {
	b := [6]byte{77, 0, 0, 0, 0, 88}
	e := OpusT_ec_enc{Fbuf: &b[1], Fstorage: 4, Frem: -1}
	ec_enc_carry_out(nil, &e, 255)
	ec_enc_carry_out(nil, &e, 255)
	ec_enc_carry_out(nil, &e, 256)
	if e.Fext != 0 || e.Frem != 0 || e.Foffs != 2 || b != [6]byte{77, 0, 0, 0, 0, 88} {
		t.Fatal(e, b)
	}
	e.Fext = 3
	ec_enc_carry_out(nil, &e, 1)
	if e.Foffs != 4 || e.Ferror1 != -1 || e.Fext != 0 || b[0] != 77 || b[5] != 88 {
		t.Fatal("full", e, b)
	}
	e = OpusT_ec_enc{Frem: -1, Fext: 0xffffffff}
	ec_enc_carry_out(nil, &e, 255)
	if e.Fext != 0 || e.Frem != -1 {
		t.Fatal("wrap", e)
	}
}

// From entropy_done_pointers_test.go.
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

// From entropy_patch_pointers_test.go.
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

// From entropy_shrink_pointers_test.go.
func TestEntropyShrinkPointers(t *testing.T) {
	buf := [10]byte{99, 1, 2, 3, 4, 5, 6, 7, 8, 88}
	enc := OpusT_ec_enc{Fbuf: &buf[1], Fstorage: 8, Foffs: 2, Fend_offs: 4, Frng: 123, Fval: 456}
	want := enc
	want.Fstorage = 6
	Opus_ec_enc_shrink(nil, &enc, 6)
	if enc != want || buf != [10]byte{99, 1, 2, 5, 6, 7, 8, 7, 8, 88} {
		t.Fatal(enc, buf)
	}
	before := slices.Clone(buf[:])
	Opus_ec_enc_shrink(nil, &enc, 6)
	if !slices.Equal(buf[:], before) {
		t.Fatal("identity")
	}
	empty := OpusT_ec_enc{Fstorage: 32}
	Opus_ec_enc_shrink(nil, &empty, 0)
	if empty != (OpusT_ec_enc{}) {
		t.Fatal("empty")
	}
}
