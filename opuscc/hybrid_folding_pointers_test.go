package opuscc

import (
	"math"
	"runtime"
	"testing"
)

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
		ownedFirst, ownedSecond := first, second
		allFirst, allSecond := first, second
		var ownedP2 *float32
		if dual != 0 {
			ownedP2 = &ownedSecond[1]
		}
		entropyInitGrowStack(12)
		runtime.GC()
		special_hybrid_folding(nil, &bands[0], &first[1], p2, 1, 1, dual)
		specialHybridFoldingFrom(&bands[0], quantBandFoldOutput{owner: ownedFirst[:], offset: 1, owned: true}, quantBandFoldOutput{legacy: ownedP2}, 1, 1, dual)
		specialHybridFoldingFrom(&bands[0], quantBandFoldOutput{owner: allFirst[:], offset: 1, owned: true}, quantAllBandsNormRight(allSecond[:], 1, dual != 0), 1, 1, dual)
		for i := range first {
			if math.Float32bits(first[i]) != math.Float32bits(ownedFirst[i]) || math.Float32bits(second[i]) != math.Float32bits(ownedSecond[i]) || math.Float32bits(first[i]) != math.Float32bits(allFirst[i]) || math.Float32bits(second[i]) != math.Float32bits(allSecond[i]) {
				t.Fatal("primary hybrid concrete owner/full bits", dual, i)
			}
		}
		if first != [8]float32{123, 1, 2, 3, 4, 3, 4, 456} || second != want2 {
			t.Fatalf("dual=%d first=%v second=%v", dual, first, second)
		}
	}
	equal := [3]int16{0, 2, 4}
	special_hybrid_folding(nil, &equal[0], nil, nil, 0, 1, 1)
	specialHybridFoldingFrom(&equal[0], quantBandFoldOutput{offset: math.MaxInt64, owned: true}, quantBandFoldOutput{offset: math.MaxInt64, owned: true}, 0, 1, 1)
}
