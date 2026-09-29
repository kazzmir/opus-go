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

func TestMatrixInt24AgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(604))
	// Keep float-to-int conversion within C's representable int32 domain.
	special := []float32{-256, math.Nextafter32(256, 0), -2, -1, 0, 1, 2, 0.5 / 8388608, 1.5 / 8388608, 2.5 / 8388608, -0.5 / 8388608, -1.5 / 8388608, -2.5 / 8388608}
	for _, shape := range [][2]int32{{2, 3}, {6, 6}, {11, 11}, {18, 18}} {
		rows, cols := shape[0], shape[1]
		data := make([]int16, rows*cols)
		for i := range data {
			data[i] = int16(r.Uint32())
		}
		data[0] = -32768
		data[1] = 32767
		m := newTestMapping(rows, cols, 31, data)
		for _, frames := range []int32{0, 1, 7, 120} {
			for _, istride := range []int32{1, 2, cols} {
				for _, ostride := range []int32{1, rows} {
					for _, col := range []int32{0, cols / 2, cols - 1} {
						n := int32(0)
						if frames > 0 {
							n = (frames-1)*istride + 1
						}
						in := make([]float32, n)
						for i := range in {
							in[i] = (r.Float32() - 0.5) * 6
							if i%2 == 0 {
								in[i] = special[(i/2)%len(special)]
							}
						}
						g := make([]int32, frames*ostride+2)
						for i := range g {
							g[i] = int32(r.Uint32())
						}
						c := slices.Clone(g)
						opuscc.Opus_mapping_matrix_multiply_channel_out_int24(nil, m, unsafe.SliceData(in), col, istride, &g[1], ostride, frames)
						nativeMatrixInt24(rows, cols, data, in, col, istride, c[1:], ostride, frames)
						if !slices.Equal(g, c) {
							t.Fatalf("shape=%v frames=%d in=%d out=%d col=%d", shape, frames, istride, ostride, col)
						}
					}
				}
			}
		}
	}
}
