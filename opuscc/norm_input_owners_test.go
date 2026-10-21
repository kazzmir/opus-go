package opuscc

import (
	"math"
	"runtime"
	"testing"
)

func TestQuantBandOwnedInputPreparationPointers(t *testing.T) {
	raw := [5]float32{77, 1, 2, 3, 88}
	owned := raw
	legacy := quantBandScratch{legacy: &raw[1]}.prepare(&raw[0], 3)
	prepared := (quantBandScratch{owner: owned[:], offset: 1, owned: true}).prepareFrom(quantBandScratch{owner: owned[:], owned: true}, 3)
	entropyInitGrowStack(12)
	runtime.GC()
	if raw != [5]float32{77, 77, 1, 2, 88} || owned != raw || legacy.legacy != &raw[1] || !prepared.owned || &prepared.owner[0] != &owned[0] || prepared.offset != 1 {
		t.Fatal("owner-to-owner copy/overlap/root", raw, owned)
	}
	unused := prepared.prepareFrom(quantBandScratch{offset: math.MaxInt32, owned: true}, 0)
	if !unused.owned || unused.offset != prepared.offset {
		t.Fatal("zero owned preparation must not view source")
	}
	var caught any
	before := owned
	func() {
		defer func() { caught = recover() }()
		prepared.prepareFrom(quantBandScratch{owner: []float32{1}, owned: true}, 3)
	}()
	if _, ok := caught.(runtime.Error); !ok || owned != before {
		t.Fatal("source prefix before copy stores", caught, owned)
	}
}
