package opuscc

import (
	"runtime"
	"testing"
	"unsafe"
)

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
