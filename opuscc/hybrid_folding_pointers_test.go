package opuscc

import "testing"

func TestHybridFoldingPointers(t *testing.T) {
	bands := [4]int16{0, 2, 6, 12}
	for _, dual := range []int32{0, 1} {
		first := [8]float32{123, 1, 2, 3, 4, 5, 6, 456}
		second := [8]float32{123, -1, -2, -3, -4, -5, -6, 456}
		want2 := second
		var p2 *float32
		if dual != 0 {
			p2 = &second[1]
			want2[5] = -3
			want2[6] = -4
		}
		special_hybrid_folding(nil, &bands[0], &first[1], p2, 1, 1, dual)
		if first != [8]float32{123, 1, 2, 3, 4, 3, 4, 456} || second != want2 {
			t.Fatalf("dual=%d first=%v second=%v", dual, first, second)
		}
	}
	equal := [3]int16{0, 2, 4}
	special_hybrid_folding(nil, &equal[0], nil, nil, 0, 1, 1)
}
