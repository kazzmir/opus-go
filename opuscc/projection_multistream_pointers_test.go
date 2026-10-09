package opuscc

import (
	"runtime"
	"testing"
	"unsafe"
)

func TestProjectionNumericOffsetPointers(t *testing.T) {
	for _, tc := range []struct{ size, want int32 }{
		{0, 8}, {4, 8}, {8, 16}, {12, 16},
		{0x7ffffff0, 0x7ffffff8}, {0x7fffffff, -2147483640}, {-1, 8}, {-16, -8},
	} {
		state := OpusT_OpusProjectionDecoder{Fdemixing_matrix_size_in_bytes: tc.size}
		if got := opusProjectionMSOffset(&state); got != uint(tc.want) {
			t.Fatal("projection numeric narrowing/wrapping", tc.size, got, uint(tc.want))
		}
	}
}

func TestProjectionMultistreamPointers(t *testing.T) {
	var ms *OpusT_OpusMSDecoder
	func() {
		b := make([]uint64, 128)
		h := (*OpusT_OpusProjectionDecoder)(unsafe.Pointer(&b[0]))
		h.Fdemixing_matrix_size_in_bytes = 129
		ms = get_multistream_decoder(nil, h)
		if unsafe.Pointer(ms) != unsafe.Add(unsafe.Pointer(h), 136) {
			t.Fatal("alignment")
		}
		ms.Flayout.Fnb_channels = 7
		ms.Flayout.Fmapping[6] = 77
	}()
	entropyInitGrowStack(12)
	runtime.GC()
	if ms.Flayout.Fnb_channels != 7 || ms.Flayout.Fmapping[6] != 77 {
		t.Fatal(ms)
	}
}
