//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"math/rand"
	"slices"
	"testing"
	"unsafe"
)

func TestMatrixInShortAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(903))
	for _, shape := range [][2]int32{{2, 3}, {6, 6}, {11, 11}, {18, 18}, {2, 255}} {
		rows, cols := shape[0], shape[1]
		data := make([]int16, rows*cols)
		for i := range data {
			data[i] = int16(r.Uint32())
		}
		data[0] = -32768
		data[1] = 32767
		m := newTestMapping(rows, cols, 0, data)
		for row := int32(0); row < rows; row++ {
			for _, istride := range []int32{0, 1, cols} {
				for _, ostride := range []int32{0, 1, rows} {
					for _, frames := range []int32{0, 1, 7} {
						input := make([]int16, frames*istride)
						for i := range input {
							if i%3 == 0 {
								input[i] = -32768
							} else if i%3 == 1 {
								input[i] = 32767
							} else {
								input[i] = int16(r.Uint32())
							}
						}
						original := slices.Clone(input)
						g := make([]float32, frames*ostride+3)
						for i := range g {
							g[i] = 12345
						}
						c := slices.Clone(g)
						opuscc.Opus_mapping_matrix_multiply_channel_in_short(nil, m, unsafe.SliceData(input), istride, &g[1], row, ostride, frames)
						nativeMatrixInShort(rows, cols, data, input, istride, c[1:], row, ostride, frames)
						for i := range g {
							if math.Float32bits(g[i]) != math.Float32bits(c[i]) {
								t.Fatalf("shape=%v row=%d strides=%d/%d frames=%d i=%d", shape, row, istride, ostride, frames, i)
							}
						}
						if !slices.Equal(input, original) {
							t.Fatal("input modified")
						}
					}
				}
			}
		}
	}
}
