package opuscc

import (
	"math"
	"testing"
)

var hadamardTestOrders = map[int32][]int32{
	2: {1, 0}, 4: {3, 0, 2, 1}, 8: {7, 0, 4, 3, 6, 1, 5, 2}, 16: {15, 0, 8, 7, 12, 3, 11, 4, 14, 1, 9, 6, 13, 2, 10, 5},
}

func TestDeinterleavePointers(t *testing.T) {
	deinterleave_hadamard(nil, nil, 0, 2, 1)
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
				want := make([]uint32, n)
				for i := int32(0); i < stride; i++ {
					row := i
					if h != 0 {
						row = hadamardTestOrders[stride][i]
					}
					for j := int32(0); j < n0; j++ {
						idx := j*stride + i
						bits := uint32(idx) * 0x9e3779b9
						if idx%7 == 0 {
							bits = 0x7f800001 + uint32(idx)
						}
						data[idx+1] = math.Float32frombits(bits)
						want[row*n0+j] = bits
					}
				}
				deinterleave_hadamard(nil, &data[1], n0, stride, h)
				for i, bits := range want {
					if math.Float32bits(data[i+1]) != bits {
						t.Fatalf("n0=%d stride=%d h=%d i=%d", n0, stride, h, i)
					}
				}
				if data[0] != 123 || data[n+1] != 456 {
					t.Fatal("guards")
				}
			}
		}
	}
}
