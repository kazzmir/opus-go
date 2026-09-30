//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"testing"
	"unsafe"
)

func TestProjectionMultistreamAgainstC(t *testing.T) {
	for size := int32(0); size <= 1024; size++ {
		b := make([]uint64, 256)
		h := (*opuscc.OpusT_OpusProjectionDecoder)(unsafe.Pointer(&b[0]))
		h.Fdemixing_matrix_size_in_bytes = size
		g := opuscc.CompareProjectionMultistream(h)
		offset := nativeProjectionMultistream(unsafe.Pointer(h))
		if unsafe.Pointer(g) != unsafe.Add(unsafe.Pointer(h), offset) || offset%8 != 0 {
			t.Fatal(size, offset)
		}
		g.Flayout.Fnb_channels = 7
		if h.Fdemixing_matrix_size_in_bytes != size {
			t.Fatal("header clobber")
		}
	}
}
