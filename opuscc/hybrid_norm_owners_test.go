package opuscc

import (
	"math"
	"runtime"
	"testing"
)

func TestHybridFoldingOwnedBoundariesPointers(t *testing.T) {
	bands := [4]int16{0, 2, 6, 12}
	first := [8]float32{123, 1, 2, 3, 4, 5, 6, 456}
	absent := quantAllBandsNormRight(first[:], 1, false)
	if absent.owned || absent.present() || absent.span(0) != nil {
		t.Fatal("absent right norm must stay absent")
	}
	specialHybridFoldingFrom(&bands[0], quantBandFoldOutput{owner: first[:], offset: 1, owned: true}, quantBandFoldOutput{offset: math.MaxInt64, owned: true}, 1, 1, 0)
	if first != [8]float32{123, 1, 2, 3, 4, 3, 4, 456} {
		t.Fatal("unused right must not be viewed", first)
	}
	first = [8]float32{123, 1, 2, 3, 4, 5, 6, 456}
	var caught any
	func() {
		defer func() { caught = recover() }()
		specialHybridFoldingFrom(&bands[0], quantBandFoldOutput{owner: first[:], offset: 1, owned: true}, quantBandFoldOutput{owned: true}, 1, 1, 1)
	}()
	if _, ok := caught.(runtime.Error); !ok || first != [8]float32{123, 1, 2, 3, 4, 3, 4, 456} {
		t.Fatal("first copy must finish before bounded right-view failure", caught, first)
	}
}

func TestHybridFoldingOwnedSequentialAliasPointers(t *testing.T) {
	bands := [4]int16{0, 2, 6, 12}
	raw := [10]float32{77, 1, 2, 3, 4, 5, 6, 7, 8, 88}
	owned := raw
	entropyInitGrowStack(12)
	runtime.GC()
	special_hybrid_folding(nil, &bands[0], &raw[1], &raw[3], 1, 1, 1)
	specialHybridFoldingFrom(&bands[0], quantBandFoldOutput{owner: owned[:], offset: 1, owned: true}, quantBandFoldOutput{owner: owned[:], offset: 3, owned: true}, 1, 1, 1)
	want := [10]float32{77, 1, 2, 3, 4, 3, 4, 3, 4, 88}
	if raw != want || owned != want {
		t.Fatal("second copy must see completed first copy", raw, owned)
	}
}
