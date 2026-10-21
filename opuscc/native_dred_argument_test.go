package opuscc

import (
	"runtime"
	"testing"
	"unsafe"
)

// Bounded Go cases for the port's ignored argument, not native DRED/model claims.
func TestOpusNativeIgnoredDREDPointers(t *testing.T) {
	for _, offset := range []int32{-2147483648, -1, 0, 2147483647} {
		left, right := newOpusFrameOwnerDecoder(t, 1), newOpusFrameOwnerDecoder(t, 1)
		a, b := make([]float32, 122), make([]float32, 122)
		a[0], a[121], b[0], b[121] = 77, 88, 77, 88
		ignored := byte(37)
		entropyInitGrowStack(12)
		runtime.GC()
		got := Opus_opus_decode_native(nil, &left.Decoder, nil, 0, &a[1], 120, 0, 0, nil, 0, unsafe.Pointer(&ignored), offset)
		want := opusDecodeNative(nil, &right.Decoder, nil, 0, &b[1], 120, 0, 0, nil, 0)
		if got != 120 || got != want || left.Decoder != right.Decoder || ignored != 37 {
			t.Fatal("ignored opaque DRED argument/offset", offset, got, want)
		}
		for i := range a {
			if a[i] != b[i] {
				t.Fatal("ignored DRED PCM/guards", offset, i)
			}
		}
		if a[0] != 77 || a[121] != 88 {
			t.Fatal("output bounds")
		}
	}
}
