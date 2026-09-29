package opuscc

import (
	"testing"
	"unsafe"
)

func TestCWRSIndexPointers(t *testing.T) {
	for _, pair := range [][2]int32{{2, 1}, {2, 128}, {3, 8}, {5, 4}, {8, 2}, {176, 1}} {
		n, k := pair[0], pair[1]
		total := celtPVQU(min(n, k), max(n, k)) + celtPVQU(min(n, k+1), max(n, k+1))
		step := max(uint32(1), total/1000)
		for index := uint32(0); index < total; index += step {
			y := make([]int32, n+2)
			y[0] = 77
			y[n+1] = 88
			cwrsi(nil, n, k, index, &y[1])
			if got := icwrs(nil, n, &y[1]); got != index || y[0] != 77 || y[n+1] != 88 {
				t.Fatal(n, k, index, got, y)
			}
		}
	}
	// Small exact stack-owned input.
	y := [2]int32{-1, 0}
	if got := icwrs(nil, 2, unsafe.SliceData(y[:])); got != 3 {
		t.Fatal(y, got)
	}
}
