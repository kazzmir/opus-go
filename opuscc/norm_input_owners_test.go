package opuscc

import (
	"math"
	"runtime"
	"testing"
)

func TestQuantBandOwnedInputUnusedPointers(t *testing.T) {
	ctx := band_ctx{}
	x := float32(1)
	for _, offset := range []int32{math.MinInt32, 0, math.MaxInt32} {
		if quantBandLanes(nil, &ctx, &x, 1, 0, 1, quantBandScratch{offset: offset, owned: true}, 0, quantBandFoldOutput{}, 1, quantBandScratch{}, 1) != 1 {
			t.Fatal("N1 must not touch unused lowband owner")
		}
		y := float32(.25)
		if quantBandStereoLanes(nil, &ctx, &x, &y, 1, 0, 1, quantBandScratch{offset: offset, owned: true}, 0, quantBandFoldOutput{}, quantBandScratch{}, 1) != 1 || y != .25 {
			t.Fatal("stereo N1 must ignore lowband owner")
		}
	}
}

func TestQuantAllBandsOwnedInputBindingPointers(t *testing.T) {
	memory := make([]float32, 12)
	for _, lane := range []int64{0, 6} {
		input := quantAllBandsFoldInputOwned(memory, lane, true, 2, 2)
		if !input.owned || &input.owner[0] != &memory[lane] || input.offset != 2 || &input.span(2)[0] != &memory[lane+2] {
			t.Fatal("real normalization suffix owner/origin", lane)
		}
	}
	if quantAllBandsFoldInputOwned(nil, math.MaxInt64, false, -1, 4).present() || quantAllBandsFoldInputOwned(nil, math.MaxInt64, false, math.MaxInt32, 0).present() {
		t.Fatal("unused input must not touch lane/storage")
	}
	input, output := quantAllBandsBandLaneViews(nil, math.MaxInt64, false, nil, 99, 1, 0, 4, -1, 1)
	if input.present() || output.present() {
		t.Fatal("unused lane views")
	}
	var caught any
	func() {
		defer func() { caught = recover() }()
		quantAllBandsBandLaneViews([]float32{1}, 0, true, nil, 0, 1, 0, 2, 0, 0)
	}()
	if _, ok := caught.(runtime.Error); !ok {
		t.Fatal("input full prefix must precede output/table view", caught)
	}
}

func TestQuantBandOwnedUncopiedAliasPointers(t *testing.T) {
	raw := [5]float32{.2, .3, .4, .5, .6}
	owned := raw
	ec := OpusT_ec_ctx{Frng: 1 << 31, Fval: (1 << 31) - 1, Fnbits_total: 33}
	oec := ec
	ctx := band_ctx{Fm: &mode48000_960_120, Fec: &ec, Fresynth: 1, Fseed: 123456}
	oc := ctx
	oc.Fec = &oec
	entropyInitGrowStack(12)
	runtime.GC()
	mask := quant_band(nil, &ctx, &raw[1], 4, 0, 1, &raw[0], 2, nil, 1, nil, 1)
	om := quantBandLanes(nil, &oc, &owned[1], 4, 0, 1, quantBandScratch{owner: owned[:], owned: true}, 2, quantBandFoldOutput{}, 1, quantBandScratch{}, 1)
	oc.Fec = ctx.Fec
	if mask != om || ec != oec || ctx != oc {
		t.Fatal("uncopied owned source full state")
	}
	for i := range raw {
		if math.Float32bits(raw[i]) != math.Float32bits(owned[i]) {
			t.Fatal("uncopied source live sequential alias", i)
		}
	}
}

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
