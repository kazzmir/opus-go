//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
	"unsafe"
)

func TestWriteExtensionAgainstC(t *testing.T) {
	for _, id := range []int32{3, 31, 32, 127} {
		for _, length := range []int32{-1, 0, 1, 2, 254, 255, 256, 510, 511} {
			for _, last := range []int32{0, 1, -1} {
				for _, capacity := range []int32{0, 1, 2, 255, 257, 600} {
					for _, pos := range []int32{0, 1} {
						for _, sizeOnly := range []bool{false, true} {
							payload := make([]byte, 512)
							for i := range payload {
								payload[i] = byte(7*i + 11)
							}
							before := slices.Clone(payload)
							g := make([]byte, 602)
							for i := range g {
								g[i] = 77
							}
							c := slices.Clone(g)
							gb, cb := g, c
							if sizeOnly {
								gb = nil
								cb = nil
							}
							r := opuscc.CompareWriteExtension(unsafe.SliceData(gb), capacity, pos, id, length, &payload[0], last)
							want := nativeWriteExtension(cb, capacity, pos, id, length, payload, last)
							if r != want || !slices.Equal(g, c) || !slices.Equal(payload, before) {
								t.Fatal(id, length, last, capacity, pos, sizeOnly, r, want)
							}
						}
					}
				}
			}
		}
	}
}

func TestWritePayloadAgainstC(t *testing.T) {
	for _, id := range []int32{3, 31, 32, 127} {
		for _, length := range []int32{-1, 0, 1, 2, 254, 255, 256, 510, 511} {
			for _, last := range []int32{0, 1, -1} {
				for _, capacity := range []int32{0, 1, 2, 255, 257, 600} {
					for _, pos := range []int32{0, 1} {
						for _, sizeOnly := range []bool{false, true} {
							payload := make([]byte, 512)
							for i := range payload {
								payload[i] = byte(7*i + 11)
							}
							before := slices.Clone(payload)
							g := make([]byte, 602)
							for i := range g {
								g[i] = 77
							}
							c := slices.Clone(g)
							gb, cb := g, c
							if sizeOnly {
								gb = nil
								cb = nil
							}
							r := opuscc.CompareWritePayload(unsafe.SliceData(gb), capacity, pos, id, length, &payload[0], last)
							want := nativeWritePayload(cb, capacity, pos, id, length, payload, last)
							if r != want || !slices.Equal(g, c) || !slices.Equal(payload, before) {
								t.Fatal(id, length, last, capacity, pos, sizeOnly, r, want)
							}
						}
					}
				}
			}
		}
	}
}

func TestSkipExtensionAgainstC(t *testing.T) {
	for id := 0; id < 256; id++ {
		for _, length := range []int32{-1, 0, 1, 2, 3, 20, 257, 600} {
			for _, lace := range []byte{0, 1, 254, 255} {
				data := make([]byte, 601)
				for i := range data {
					data[i] = byte(11*i + 3)
				}
				data[0] = byte(id)
				data[1] = lace
				before := slices.Clone(data)
				p := &data[0]
				h := int32(77)
				g := opuscc.CompareSkipExtension(&p, length, &h)
				c, offset, ch := nativeSkipPayload(data, length, 0, 0, 77, 1)
				if g != c || h != ch || unsafe.Pointer(p) != unsafe.Add(unsafe.Pointer(&data[0]), offset) || !slices.Equal(data, before) {
					t.Fatal(id, length, lace, g, c, h, ch, offset)
				}
			}
		}
	}
}

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
