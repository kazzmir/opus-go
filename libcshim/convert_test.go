package libcshim

import "testing"

// C converts a signed integer to a wider unsigned type by value modulo
// 2^N - i.e. it sign-extends first. Transpiled expressions like
// (opus_uint32)idx << 10 on a negative opus_int8 depend on it.
func TestSignedToUnsignedSignExtends(t *testing.T) {
	cases := []struct {
		name      string
		got, want uint64
	}{
		{"Uint32FromInt8(-4)", uint64(Uint32FromInt8(-4)), 0xFFFFFFFC},
		{"Uint32FromInt16(-2)", uint64(Uint32FromInt16(-2)), 0xFFFFFFFE},
		{"Uint64FromInt16(-3)", Uint64FromInt16(-3), 0xFFFFFFFFFFFFFFFD},
		{"Uint64FromInt32(-1)", Uint64FromInt32(-1), 0xFFFFFFFFFFFFFFFF},
		{"Uint32FromInt8(5)", uint64(Uint32FromInt8(5)), 5},
	}
	for _, c := range cases {
		if c.got != c.want {
			t.Errorf("%s = %#x, want %#x", c.name, c.got, c.want)
		}
	}
	// (uintptr)(int32)-1 is all ones on every architecture, so p + that
	// offset steps backwards instead of ~4 GiB forwards on 64-bit.
	if got := UintptrFromInt32(-1); got != ^uintptr(0) {
		t.Errorf("UintptrFromInt32(-1) = %#x, want %#x", got, ^uintptr(0))
	}
	if got := Int32FromUint32(Uint32FromInt8(-4) << 10); got != -4096 {
		t.Errorf("(int32)((uint32)(int8)-4 << 10) = %d, want -4096", got)
	}
}
