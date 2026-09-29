//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
	"unsafe"
)

func TestMatrixDataAgainstC(t *testing.T) {
	for _, shape := range [][2]int32{{1, 1}, {2, 3}, {6, 6}, {11, 11}, {18, 18}, {255, 127}, {127, 255}} {
		rows, cols := shape[0], shape[1]
		data := make([]int16, rows*cols)
		for i := range data {
			data[i] = int16(i * 7919)
		}
		m := newTestMapping(rows, cols, 0, data)
		p := opuscc.Opus_mapping_matrix_get_data(nil, m)
		want, meta := nativeMatrixInit(rows, cols, 0, data)
		if uintptr(unsafe.Pointer(p))-uintptr(unsafe.Pointer(m)) != uintptr(meta[3]) || !slices.Equal(unsafe.Slice(p, len(data)), want) {
			t.Fatal(shape)
		}
	}
}
