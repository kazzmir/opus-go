//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
	"unsafe"
)

func TestSignEncodeAgainstC(t *testing.T) {
	for signal := int32(0); signal < 3; signal++ {
		for quant := int32(0); quant < 2; quant++ {
			for _, length := range []int32{0, 7, 8, 16, 24, 32, 120, 160, 320} {
				for _, capacity := range []uint32{0, 1, 3, 128} {
					blocks := (length + 8) >> 4
					p := make([]int8, blocks*16)
					sums := make([]int32, blocks)
					for i := range sums {
						sums[i] = []int32{-1, 0, 1, 5, 6, 7, 31, 32, 33, 63, 64, 65}[i%12]
					}
					for i := range p {
						p[i] = []int8{0, 1, -1, 127, -128, 2, -2}[i%7]
					}
					original := slices.Clone(p)
					originalSums := slices.Clone(sums)
					g := make([]byte, capacity+2)
					for i := range g {
						g[i] = 77
					}
					c := slices.Clone(g)
					var ge opuscc.OpusT_ec_enc
					opuscc.Opus_ec_enc_init(nil, &ge, &g[1], capacity)
					ce := ge
					opuscc.Opus_silk_encode_signs(nil, &ge, unsafe.SliceData(p), length, signal, quant, unsafe.SliceData(sums))
					nativeEncoderStepPointer(&ce, c[1:], 17, uint32(length), uint32(signal), uint32(quant), unsafe.Pointer(unsafe.SliceData(p)), unsafe.SliceData(sums))
					ce.Fbuf = ge.Fbuf
					if ge != ce || !slices.Equal(g, c) || !slices.Equal(p, original) || !slices.Equal(sums, originalSums) {
						t.Fatal(signal, quant, length, capacity, ge, ce)
					}
				}
			}
		}
	}
}
