package opuscc

import (
	"math/bits"
	"runtime"
	"testing"
	"unsafe"
)

// From entropy_bit_pointers_test.go.
func TestTFDecodePointers(t *testing.T) {
	dec := OpusT_ec_dec{}
	Opus_ec_dec_init(nil, &dec, nil, 0)
	out := [6]int32{77, 9, 9, 9, 9, 88}
	tf_decode(nil, 1, 5, 0, &out[0], 0, &dec)
	if out != [6]int32{77, 0, 0, 0, 0, 88} {
		t.Fatal(out)
	}
	dec = OpusT_ec_dec{}
	Opus_ec_dec_init(nil, &dec, nil, 0)
	tf_decode(nil, 0, 0, 0, nil, 0, &dec)
}

func TestTFDecodeSpanPointers(t *testing.T) {
	for LM := int32(0); LM <= 3; LM++ {
		for transient := int32(0); transient <= 1; transient++ {
			for _, size := range []int{0, 1, 32} {
				for _, bounds := range [][2]int32{{0, 0}, {1, 5}, {5, 5}} {
					data := make([]byte, size)
					for i := range data {
						data[i] = byte(i*37 + 19)
					}
					var dec OpusT_ec_dec
					Opus_ec_dec_init(nil, &dec, unsafe.SliceData(data), uint32(size))
					owned := dec
					a := [6]int32{77, 9, 9, 9, 9, 88}
					b := a
					entropyInitGrowStack(12)
					runtime.GC()
					tf_decode(nil, bounds[0], bounds[1], transient, &a[0], LM, &dec)
					tfDecodeSpan(nil, bounds[0], bounds[1], transient, b[:], LM, &owned)
					if a != b || dec != owned || a[0] != 77 || a[5] != 88 {
						t.Fatal("TF span/context/guards", LM, transient, size, bounds)
					}
				}
			}
		}
	}
	var a OpusT_ec_dec
	Opus_ec_dec_init(nil, &a, nil, 0)
	b := a
	tf_decode(nil, 0, 1, 0, &a.Fnbits_total, 0, &a)
	tfDecodeSpan(nil, 0, 1, 0, unsafe.Slice(&b.Fnbits_total, 1), 0, &b)
	if a != b {
		t.Fatal("TF scalar entropy-field alias")
	}
	Opus_ec_dec_init(nil, &a, nil, 0)
	b = a
	tf_decode(nil, 0, 0, 0, nil, 0, &a)
	tfDecodeSpan(nil, 0, 0, 0, nil, 0, &b)
	if a != b {
		t.Fatal("nil TF zero extent")
	}
}

func TestEntropyBitPointers(t *testing.T) {
	for _, logp := range []uint32{1, 2, 7, 8, 15, 24, 31} {
		const rng = uint32(1 << 31)
		split := rng >> logp
		for _, value := range []uint32{0, split - 1, split, split + 1, rng - 1} {
			dec := OpusT_ec_dec{Frng: rng, Fval: value, Fnbits_total: 33, Fext: 123, Fend_window: 456, Fnend_bits: 9, Ferror1: 7}
			want := dec
			var bit int32
			if value < split {
				bit = 1
				want.Frng = split
			} else {
				want.Frng = rng - split
				want.Fval = value - split
			}
			for want.Frng <= 1<<23 {
				want.Frng <<= 8
				want.Fval = ((want.Fval << 8) + 255) & 0x7fffffff
				want.Fnbits_total += 8
			}
			if got := Opus_ec_dec_bit_logp(nil, &dec, logp); got != bit || dec != want {
				t.Fatalf("logp=%d value=%d bit=%d want=%d state=%+v want=%+v", logp, value, got, bit, dec, want)
			}
		}
	}
}

// From entropy_bits_pointers_test.go.
func TestEntropyBitsPointers(t *testing.T) {
	for available := int32(0); available <= 32; available++ {
		for width := uint32(0); width <= 25; width++ {
			window := uint32(0xd6b59a73) & ((uint32(1) << available) - 1)
			dec := OpusT_ec_dec{Frng: 1 << 31, Fval: 123, Fext: 456, Frem: 78, Ferror1: 9, Fend_window: window, Fnend_bits: available, Fnbits_total: 33}
			want := dec
			if uint32(want.Fnend_bits) < width {
				for {
					want.Fnend_bits += 8
					if want.Fnend_bits > 24 {
						break
					}
				}
			}
			want.Fnend_bits -= int32(width)
			want.Fnbits_total += int32(width)
			want.Fend_window >>= width
			expected := window & ((uint32(1) << width) - 1)
			if got := Opus_ec_dec_bits(nil, &dec, width); got != expected || dec != want {
				t.Fatalf("available=%d width=%d got=%x want=%x state=%+v want=%+v", available, width, got, expected, dec, want)
			}
		}
	}
}

// From entropy_uint_pointers_test.go.
func TestEntropyUintPointers(t *testing.T) {
	for _, ft := range []uint32{2, 3, 255, 256, 257, 258, 65535, 65536, 65537, 1 << 24, 0xffffffff} {
		dec := OpusT_ec_dec{Frng: 1 << 31, Fval: (1 << 31) - 1, Fnbits_total: 33, Fnend_bits: 32, Ferror1: 7}
		got := Opus_ec_dec_uint(nil, &dec, ft)
		consumed := max(0, bits.Len32(ft-1)-8)
		if got != 0 || dec.Ferror1 != 7 || dec.Fnend_bits != 32-int32(consumed) || dec.Frng <= 1<<23 {
			t.Fatalf("ft=%d got=%d state=%+v", ft, got, dec)
		}
	}
	// Highest range symbol plus a nonzero tail can represent an invalid value.
	for _, tail := range []uint32{0, 1} {
		dec := OpusT_ec_dec{Frng: 1 << 31, Fval: 0, Fnbits_total: 33, Fnend_bits: 1, Fend_window: tail, Ferror1: 7}
		got := Opus_ec_dec_uint(nil, &dec, 257)
		wantError := int32(7)
		if tail == 1 {
			wantError = 1
		}
		if got != 256 || dec.Ferror1 != wantError || dec.Fnend_bits != 0 || dec.Fend_window != 0 || dec.Fnbits_total != 34 {
			t.Fatalf("tail=%d got=%d state=%+v", tail, got, dec)
		}
	}
}

// From entropy_update_pointers_test.go.
func TestEntropyUpdatePointers(t *testing.T) {
	for _, fl := range []uint32{0, 1} {
		dec := OpusT_ec_dec{Frng: 100000003, Fval: 80000000, Fext: 10000000, Fnbits_total: 33, Frem: 123, Fend_window: 456, Fnend_bits: 9, Ferror1: 7}
		want := dec
		want.Fval = 10000000
		if fl == 0 {
			want.Frng = 30000003
		} else {
			want.Frng = 20000000
		}
		Opus_ec_dec_update(nil, &dec, fl, 3, 10)
		if dec != want {
			t.Fatalf("fl=%d got=%+v want=%+v", fl, dec, want)
		}
	}
	// An exhausted packet supplies zero bytes while normalization continues.
	dec := OpusT_ec_dec{Frng: 1 << 24, Fval: 123, Fext: 1, Fnbits_total: 33, Frem: 0}
	Opus_ec_dec_update(nil, &dec, 1, 2, 2)
	if dec.Frng != 1<<24 || dec.Fval != 2080374783 || dec.Fnbits_total != 57 || dec.Foffs != 0 || dec.Frem != 0 {
		t.Fatalf("normalized=%+v", dec)
	}
}

// From entropy_normalize_pointers_test.go.
func TestEntropyContextAliasPointers(t *testing.T) {
	packet := []byte{11, 22}
	context := new(OpusT_ec_ctx)
	context.Fbuf = &packet[0]
	context.Fstorage = 2
	var encoder *OpusT_ec_enc = context
	var decoder *OpusT_ec_dec = context
	saved := *context
	encoder.Foffs = 1
	decoder.Fend_offs = 1
	context.Fval = 77
	*context = saved
	packet = nil
	entropyInitGrowStack(12)
	runtime.GC()
	if encoder != decoder || encoder != context || ec_read_byte(decoder) != 11 || ec_read_byte_from_end(decoder) != 22 || encoder.Foffs != 1 || encoder.Fend_offs != 1 || context.Fval != 0 {
		t.Fatal("typed context identity/snapshot")
	}
}

func TestEntropyByteWritePointers(t *testing.T) {
	data := [5]byte{77, 0, 0, 0, 88}
	enc := OpusT_ec_enc{Fbuf: &data[1], Fstorage: 3}
	entropyInitGrowStack(12)
	runtime.GC()
	if ec_write_byte(nil, &enc, 0x113) != 0 || ec_write_byte_at_end(nil, &enc, 0x122) != 0 || ec_write_byte(nil, &enc, 0x131) != 0 {
		t.Fatal("byte writes")
	}
	before := enc
	image := data
	if ec_write_byte(nil, &enc, 9) != -1 || ec_write_byte_at_end(nil, &enc, 9) != -1 || enc != before || data != image {
		t.Fatal("full buffer mutation")
	}
	if data != [5]byte{77, 19, 49, 34, 88} {
		t.Fatal("write guards", data)
	}
	empty := OpusT_ec_enc{}
	if ec_write_byte(nil, &empty, 1) != -1 || ec_write_byte_at_end(nil, &empty, 1) != -1 {
		t.Fatal("empty")
	}
	single := byte(0)
	prefix := OpusT_ec_enc{Fbuf: &single, Fstorage: 100}
	if ec_write_byte(nil, &prefix, 19) != 0 || single != 19 {
		t.Fatal("consumed write prefix")
	}
	for _, tail := range []bool{false, true} {
		alias := OpusT_ec_enc{Fstorage: 4}
		want := uint32(1)
		wantBytes := unsafe.Slice((*byte)(unsafe.Pointer(&want)), 4)
		if tail {
			alias.Fbuf = (*byte)(unsafe.Pointer(&alias.Fend_offs))
			wantBytes[3] = 0xa5
			if ec_write_byte_at_end(nil, &alias, 0xa5) != 0 || alias.Fend_offs != want {
				t.Fatal("live tail counter")
			}
		} else {
			alias.Fbuf = (*byte)(unsafe.Pointer(&alias.Foffs))
			wantBytes[0] = 0xa5
			if ec_write_byte(nil, &alias, 0xa5) != 0 || alias.Foffs != want {
				t.Fatal("live front counter")
			}
		}
	}
}

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

// From entropy_init_pointers_test.go.
//
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

// From icdf8_pointers_test.go.
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

// From icdf16_pointers_test.go.
func TestICDF16Pointers(t *testing.T) {
	table := []uint16{60000, 40000, 20000, 0}
	for symbol := int32(0); symbol < 4; symbol++ {
		dec := OpusT_ec_dec{Frng: 1 << 31, Fval: uint32(table[symbol]) * (1 << 15), Fnbits_total: 33, Fext: 123, Fend_window: 456, Fnend_bits: 7}
		before := dec
		got := Opus_ec_dec_icdf16(nil, &dec, &table[0], 16)
		upper := uint32(65536)
		if symbol > 0 {
			upper = uint32(table[symbol-1])
		}
		want := before
		want.Fval = 0
		want.Frng = (upper - uint32(table[symbol])) * (1 << 15)
		if got != symbol || dec != want {
			t.Fatalf("symbol=%d got=%d state=%+v want=%+v", symbol, got, dec, want)
		}
	}
	// Probability 1/65536 requires two normalization steps on an exhausted packet.
	tiny := []uint16{65535, 0}
	dec := OpusT_ec_dec{Frng: 1 << 31, Fval: (1 << 31) - 1, Fnbits_total: 33}
	if got := Opus_ec_dec_icdf16(nil, &dec, &tiny[0], 16); got != 0 || dec.Frng != 1<<31 || dec.Fval != (1<<31)-1 || dec.Fnbits_total != 49 {
		t.Fatalf("tiny symbol=%d state=%+v", got, dec)
	}
}

// From range_lookup_pointers_test.go.
func TestRangeLookupPointers(t *testing.T) {
	for _, tc := range []struct{ rng, val, ft, ext, symbol uint32 }{
		{100, 0, 10, 10, 9}, {100, 99, 10, 10, 0},
		{101, 100, 10, 10, 0}, {100, 42, 1, 100, 0},
		{0x80000000, 0x7fffffff, 32768, 65536, 0},
	} {
		dec := OpusT_ec_dec{Frng: tc.rng, Fval: tc.val, Fext: 99, Foffs: 7, Ferror1: -1}
		want := dec
		want.Fext = tc.ext
		got := Opus_ec_decode(nil, &dec, tc.ft)
		if got != tc.symbol || dec != want {
			t.Fatalf("%+v: symbol=%d state=%+v", tc, got, dec)
		}
	}
	for bits := uint32(0); bits <= 15; bits++ {
		for _, value := range []uint32{0, 123456, 0x7ffffffe} {
			dec := OpusT_ec_dec{Frng: 0x7fffffff, Fval: value}
			other := dec
			got := Opus_ec_decode_bin(nil, &dec, bits)
			want := Opus_ec_decode(nil, &other, 1<<bits)
			if got != want || dec != other {
				t.Fatalf("bits=%d val=%d: binary and uniform lookup differ", bits, value)
			}
		}
	}
}

// From tell_frac_pointers_test.go.
func TestTellFracPointers(t *testing.T) {
	for exponent := uint(15); exponent <= 31; exponent++ {
		ctx := OpusT_ec_ctx{Frng: 1 << exponent, Fnbits_total: 100, Fval: 123}
		before := ctx
		if got, want := Opus_ec_tell_frac(nil, &ctx), uint32(100-int32(exponent)-1)*8; got != want {
			t.Fatalf("exponent=%d got=%d want=%d", exponent, got, want)
		}
		if ctx != before {
			t.Fatal("context changed")
		}
	}
	thresholds := []uint32{35733, 38967, 42495, 46340, 50535, 55109, 60097}
	for i, r := range thresholds {
		ctx := OpusT_ec_ctx{Frng: r << 8, Fnbits_total: 100}
		at := Opus_ec_tell_frac(nil, &ctx)
		ctx.Frng = (r + 1) << 8
		after := Opus_ec_tell_frac(nil, &ctx)
		if at != uint32(800-192-i) || after != at-1 {
			t.Fatalf("threshold=%d at=%d after=%d", r, at, after)
		}
	}
}
