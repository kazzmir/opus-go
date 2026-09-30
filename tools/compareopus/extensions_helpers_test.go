//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
	"unsafe"
)

func TestSkipPayloadAgainstC(t *testing.T) {
	for id := int32(0); id < 256; id++ {
		for _, length := range []int32{0, 1, 2, 20, 256, 257, 600} {
			for _, trailing := range []int32{0, 1, 5} {
				for _, lace := range []byte{0, 1, 254, 255} {
					data := make([]byte, 601)
					for i := range data {
						data[i] = byte(11*i + 3)
					}
					data[0] = lace
					before := slices.Clone(data)
					p := &data[0]
					h := int32(77)
					g := opuscc.CompareSkipPayload(&p, length, &h, id, trailing)
					c, offset, ch := nativeSkipPayload(data, length, id, trailing, 77, 0)
					if g != c || h != ch || unsafe.Pointer(p) != unsafe.Add(unsafe.Pointer(&data[0]), offset) || !slices.Equal(data, before) {
						t.Fatal(id, length, trailing, lace, g, c, h, ch, offset)
					}
				}
			}
		}
	}
}
