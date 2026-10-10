//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opusccenc"
	"math"
	"slices"
	"testing"
	"unsafe"
)

func flpSameBits(a, b []float32) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
			return false
		}
	}
	return true
}

func TestScaleCopyFLPAgainstC(t *testing.T) {
	for _, size := range []int32{0, 1, 3, 4, 5, 8, 31, 65537, -1, -65536} {
		for _, shift := range []int32{-1, 0, 1} {
			for _, gain := range []float32{0, -.75, 1, 1.125} {
				length := max(size, size&0xFFFC)
				goData := make([]float32, length+3)
				for i := range goData {
					goData[i] = float32(i%19-9) / 7
				}
				cData := slices.Clone(goData)
				opusccenc.Opus_silk_scale_copy_vector_FLP(nil, uintptr(unsafe.Pointer(&goData[1+shift])), uintptr(unsafe.Pointer(&goData[1])), gain, size)
				nativeScaleCopyFLP(cData[1+shift:], cData[1:], gain, size)
				if !flpSameBits(goData, cData) {
					t.Fatal("scale-copy", size, shift, gain)
				}
			}
		}
	}
	nativeScaleCopyFLP(nil, nil, 1, 0)
	nativeScaleCopyFLP(nil, nil, 1, -65536)
}
