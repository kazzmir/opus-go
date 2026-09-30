//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"slices"
	"testing"
	"unsafe"
)

func TestAmpLogAgainstC(t *testing.T) {
	const bands = 21
	values := []uint32{0, 1, 0x007fffff, 0x00800000, 0x3f7fffff, 0x3f800000, 0x3f800001, 0x7f7fffff, 0x7f800000, 0x7fc01234, 0x80000000, 0xbf800000, 0xff800000}
	for exponent := uint32(1); exponent < 255; exponent += 19 {
		for bin := uint32(0); bin < 8; bin++ {
			values = append(values, exponent<<23|bin<<20, exponent<<23|bin<<20|0xfffff)
		}
	}
	for _, channels := range []int32{1, 2} {
		for _, eff := range []int32{0, 1, 13, 21} {
			for _, shift := range []int{0, 1, 2} {
				for offset := 0; offset < len(values); offset += int(bands * channels) {
					g := make([]float32, int(bands*channels)+shift+2)
					for i := range g {
						g[i] = math.Float32frombits(values[(offset+i)%len(values)])
					}
					c := slices.Clone(g)
					outG, outC := g[1+shift:], c[1+shift:]
					inG, inC := g[1:], c[1:]
					m := opuscc.OpusT_OpusCustomMode{FnbEBands: bands}
					opuscc.Opus_amp2Log2(nil, &m, eff, bands, unsafe.SliceData(inG), unsafe.SliceData(outG), channels)
					nativeAmpLog(inC, outC, bands, eff, bands, channels)
					if !sameFloatBits(g, c) {
						t.Fatal(channels, eff, shift, offset, g, c)
					}
				}
			}
		}
	}
}
