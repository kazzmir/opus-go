//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"testing"
	"unsafe"
)

func TestProjectionMatrixAgainstC(t *testing.T) {
	for _, size := range []int32{0, 16, 128, 65536} {
		for _, shape := range [][3]int32{{0, 0, 0}, {2, 3, 77}, {255, 127, -32768}} {
			backing := make([]uint64, 64)
			header := (*opuscc.OpusT_OpusProjectionDecoder)(unsafe.Pointer(&backing[0]))
			header.Fdemixing_matrix_size_in_bytes = size
			g := opuscc.CompareProjectionMatrix(header)
			g.Frows = shape[0]
			g.Fcols = shape[1]
			g.Fgain = shape[2]
			offset, fields := nativeProjectionMatrix(unsafe.Pointer(header))
			if unsafe.Pointer(g) != unsafe.Add(unsafe.Pointer(header), offset) || fields != shape || header.Fdemixing_matrix_size_in_bytes != size {
				t.Fatal(size, shape, offset, fields)
			}
		}
	}
}
