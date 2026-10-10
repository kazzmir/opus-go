//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opusccenc"
	"math"
	"math/rand"
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

func TestInnerProductFLPAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(9541))
	for _, size := range []int32{0, 1, 3, 4, 5, 8, 31, 320, 2048} {
		for trial := 0; trial < 100; trial++ {
			a, b := make([]float32, size+2), make([]float32, size+2)
			for i := range a {
				a[i] = math.Float32frombits(rng.Uint32() & 0xfeffffff)
				b[i] = math.Float32frombits(rng.Uint32() & 0xfeffffff)
			}
			beforeA, beforeB := slices.Clone(a), slices.Clone(b)
			g := opusccenc.Opus_silk_inner_product_FLP_c(nil, uintptr(unsafe.Pointer(&a[1])), uintptr(unsafe.Pointer(&b[1])), size)
			c := nativeInnerProductFLP(a[1:], b[1:], size)
			if math.Float64bits(g) != math.Float64bits(c) || !flpSameBits(a, beforeA) || !flpSameBits(b, beforeB) {
				t.Fatal("inner-product", size, trial, g, c)
			}
		}
	}
	for _, size := range []int32{0, -1, -65536} {
		if nativeInnerProductFLP(nil, nil, size) != 0 || opusccenc.Opus_silk_inner_product_FLP_c(nil, 0, 0, size) != 0 {
			t.Fatal("empty product", size)
		}
	}
}

func TestScaleFLPAgainstC(t *testing.T) {
	for _, size := range []int32{0, 1, 3, 4, 5, 8, 31, 65537, -1, -65536} {
		for _, gain := range []float32{0, -.75, 1, 1.125} {
			length := max(size, size&0xFFFC)
			goData := make([]float32, length+2)
			for i := range goData {
				goData[i] = float32(i%19-9) / 7
			}
			cData := slices.Clone(goData)
			opusccenc.Opus_silk_scale_vector_FLP(nil, uintptr(unsafe.Pointer(&goData[1])), gain, size)
			nativeScaleFLP(cData[1:], gain, size)
			if !flpSameBits(goData, cData) {
				t.Fatal("scale", size, gain)
			}
		}
	}
	nativeScaleFLP(nil, 1, 0)
	nativeScaleFLP(nil, 1, -65536)
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
