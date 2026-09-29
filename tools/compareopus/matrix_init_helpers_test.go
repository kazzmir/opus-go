//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
	"unsafe"
)

func newTestMapping(rows, cols, gain int32, data []int16) *opuscc.OpusT_MappingMatrix {
	size := opuscc.Opus_mapping_matrix_get_size(nil, rows, cols)
	memory := make([]uint64, (size+7)/8)
	m := (*opuscc.OpusT_MappingMatrix)(unsafe.Pointer(&memory[0]))
	opuscc.Opus_mapping_matrix_init(nil, m, rows, cols, gain, unsafe.SliceData(data), int32(2*len(data)))
	return m
}

func TestMatrixInitAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(601))
	for _, shape := range [][2]int32{{0, 0}, {1, 1}, {2, 3}, {6, 6}, {11, 11}, {18, 18}, {255, 127}, {127, 255}} {
		rows, cols := shape[0], shape[1]
		data := make([]int16, rows*cols)
		for i := range data {
			data[i] = int16(r.Uint32())
		}
		for _, gain := range []int32{-32768, 0, 32767} {
			m := newTestMapping(rows, cols, gain, data)
			want, meta := nativeMatrixInit(rows, cols, gain, data)
			got := unsafe.Slice((*int16)(unsafe.Add(unsafe.Pointer(m), 16)), len(data))
			if m.Frows != meta[0] || m.Fcols != meta[1] || m.Fgain != meta[2] || meta[3] != 16 || opuscc.Opus_mapping_matrix_get_size(nil, rows, cols) != meta[4] || !slices.Equal(got, want) {
				t.Fatal(shape, gain, meta)
			}
		}
	}
}
