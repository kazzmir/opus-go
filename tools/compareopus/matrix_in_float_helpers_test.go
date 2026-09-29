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

func TestMatrixInFloatAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(902))
	for _, shape := range [][2]int32{{2, 3}, {6, 6}, {11, 11}, {18, 18}} {
		rows, cols := shape[0], shape[1]
		data := make([]int16, rows*cols)
		for i := range data {
			data[i] = int16(r.Uint32())
		}
		m := newTestMapping(rows, cols, 0, data)
		for row := int32(0); row < rows; row++ {
			for _, istride := range []int32{0, 1, cols} {
				for _, ostride := range []int32{0, 1, rows} {
					for _, frames := range []int32{0, 1, 7} {
						for _, alias := range []bool{false, true} {
							size := int(frames*max(istride, ostride) + 4)
							gi := make([]float32, size)
							for i := range gi {
								gi[i] = (r.Float32() - 0.5) * 8
							}
							ci := slices.Clone(gi)
							g := make([]float32, size)
							for i := range g {
								g[i] = 12345
							}
							c := slices.Clone(g)
							if alias {
								g = gi
								c = ci
							}
							opuscc.Opus_mapping_matrix_multiply_channel_in_float(nil, m, unsafe.SliceData(gi), istride, &g[1], row, ostride, frames)
							nativeMatrixInFloat(rows, cols, data, ci, istride, c[1:], row, ostride, frames)
							for i := range g {
								if math.Float32bits(g[i]) != math.Float32bits(c[i]) {
									t.Fatalf("shape=%v row=%d strides=%d/%d frames=%d alias=%v i=%d", shape, row, istride, ostride, frames, alias, i)
								}
							}
						}
					}
				}
			}
		}
	}
}
