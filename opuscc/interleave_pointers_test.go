package opuscc

import (
	"math"
	"testing"
)

func TestInterleavePointers(t *testing.T) {
	interleave_hadamard(nil, nil, 0, 2, 1)
	interleave_hadamard(nil, nil, 2, 0, 0)
	for _, stride := range []int32{1, 2, 3, 4, 8, 16} {
		for _, n0 := range []int32{1, 2, 7, 22, 37} {
			for h := int32(0); h <= 1; h++ {
				if h != 0 && hadamardTestOrders[stride] == nil {
					continue
				}
				n := n0 * stride
				data := make([]float32, n+2)
				data[0] = 123
				data[n+1] = 456
				original := make([]uint32, n)
				want := make([]uint32, n)
				for i := range original {
					bits := uint32(i) * 0x9e3779b9
					if i%7 == 0 {
						bits = 0xff800001 + uint32(i)
					}
					original[i] = bits
					data[i+1] = math.Float32frombits(bits)
				}
				for i := int32(0); i < stride; i++ {
					row := i
					if h != 0 {
						row = hadamardTestOrders[stride][i]
					}
					for j := int32(0); j < n0; j++ {
						want[j*stride+i] = original[row*n0+j]
					}
				}
				interleave_hadamard(nil, &data[1], n0, stride, h)
				for i, bits := range want {
					if math.Float32bits(data[i+1]) != bits {
						t.Fatalf("n0=%d stride=%d h=%d i=%d", n0, stride, h, i)
					}
				}
				deinterleave_hadamard(nil, &data[1], n0, stride, h)
				for i, bits := range original {
					if math.Float32bits(data[i+1]) != bits {
						t.Fatal("round trip", n0, stride, h, i)
					}
				}
				if data[0] != 123 || data[n+1] != 456 {
					t.Fatal("guards")
				}
			}
		}
	}
}
