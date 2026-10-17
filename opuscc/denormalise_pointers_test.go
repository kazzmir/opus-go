package opuscc

import (
	"math"
	"runtime"
	"testing"
)

var denormalisePointerBands = [4]int16{0, 2, 4, 6}

func TestBandEnergiesPointers(t *testing.T) {
	Opus_compute_band_energies(nil, nil, 0, 0, nil, nil, 0, 0, 0, 0)
	bands := [4]int16{1, 3, 3, 4}
	input := [8]float32{99, 3, 4, 2, 99, 6, 8, 3}
	out := [10]float32{77, 9, 9, 9, 9, 9, 9, 9, 9, 88}
	entropyInitGrowStack(12)
	runtime.GC()
	Opus_compute_band_energies(nil, &bands[0], 4, 4, &input[0], &out[1], 3, 2, 0, 0)
	tiny := float32(math.Sqrt(float64(float32(1e-27))))
	if out != [10]float32{77, 5, tiny, 2, 9, 10, tiny, 3, 9, 88} {
		t.Fatal(out)
	}
	// A scalar rounded-square oracle catches inadvertent fused accumulation.
	samples := [7]float32{.12345, -2.3456, 31.2345, .01234, -.76543, 4.56789, -13.987}
	oneBand := [2]int16{0, 7}
	sum := float32(0)
	for _, v := range samples {
		sum += float32(v * v)
	}
	want := float32(math.Sqrt(float64(float32(1e-27) + sum)))
	result := float32(0)
	Opus_compute_band_energies(nil, &oneBand[0], 7, 1, &samples[0], &result, 1, 0, 0, 0)
	if math.Float32bits(result) != math.Float32bits(want) {
		t.Fatal(result, want)
	}
}

func TestNormaliseBandsPointers(t *testing.T) {
	Opus_normalise_bands(nil, nil, 0, 0, nil, nil, nil, 0, 0, 0)
	bands := [3]int16{1, 2, 3}
	in := [8]float32{1, 2, 3, 4, 5, 6, 7, 8}
	energy := [4]float32{2, 4, 2, 4}
	out := [10]float32{77, 9, 9, 9, 9, 9, 9, 9, 9, 88}
	entropyInitGrowStack(12)
	runtime.GC()
	Opus_normalise_bands(nil, &bands[0], 4, 2, &in[0], &out[1], &energy[0], 2, 2, 1)
	if out != [10]float32{77, 9, 1, .75, 9, 9, 3, 1.75, 9, 88} {
		t.Fatal(out)
	}
	// C also visits channel zero for C=0; only selected bands are overwritten.
	Opus_normalise_bands(nil, &bands[0], 4, 2, &in[0], &out[1], &energy[0], 2, 0, 1)
	if out[2] != 1 || out[3] != .75 || out[0] != 77 || out[9] != 88 {
		t.Fatal("zero channel contract", out)
	}
}

func TestDenormaliseOffsetPointers(t *testing.T) {
	for _, start := range []int32{0, 1} {
		for _, down := range []int32{1, 2} {
			for _, energy := range []int32{7, 13} {
				var want, got, numeric, span, owned [18]float32
				for i := range want {
					want[i] = float32(i%7-3) / 17
				}
				got = want
				numeric = want
				span = want
				owned = want
				// Output overlaps both spectrum and (in one case) energy.
				Opus_denormalise_bands(nil, &denormalisePointerBands[0], 8, &want[2], &want[4], &want[energy], start, 3, 1, down, 0)
				entropyInitGrowStack(12)
				runtime.GC()
				celtDenormaliseBandsAt(nil, &denormalisePointerBands[0], 8, &got[0], 2, &got[4], &got[0], energy, start, 3, 1, down, 0)
				celtDenormaliseBandsOffsets(nil, &denormalisePointerBands[0], 8, &numeric[0], 2, &numeric[0], 4, &numeric[0], energy, start, 3, 1, down, 0)
				celtDenormaliseBandsTo(nil, &denormalisePointerBands[0], 8, &owned[0], 2, celtDenormaliseOutput{owner: owned[:], owned: true}, 4, &owned[0], energy, start, 3, 1, down, 0)
				bound := int32(6)
				if down != 1 {
					bound = min(bound, 8/down)
				}
				celtDenormaliseSpans(nil, denormalisePointerBands[:], span[2:8], span[4:12], span[energy:energy+3], start, 3, 1, bound)
				for i := range got {
					if math.Float32bits(span[i]) != math.Float32bits(want[i]) || math.Float32bits(owned[i]) != math.Float32bits(want[i]) {
						t.Fatal("denormalise span live aliases/guards", i)
					}
					if math.Float32bits(got[i]) != math.Float32bits(want[i]) || math.Float32bits(numeric[i]) != math.Float32bits(want[i]) {
						t.Fatal("offset bits/live aliases/guards", start, down, energy, i)
					}
				}
			}
		}
	}
	if celtDenormaliseSpan(nil, 2147483647, 0) != nil {
		t.Fatal("unused displaced span")
	}
	owner := [3]float32{.25, .5, 1}
	if span := celtDenormaliseSpan(&owner[1], -1, 2); span[0] != .25 || span[1] != .5 {
		t.Fatal("signed history")
	}
	out := [4]float32{77, 1, 2, 88}
	bands := [2]int16{0, 2}
	celtDenormaliseBandsAt(nil, &bands[0], 2, nil, 2147483647, &out[1], nil, 2147483647, 0, 1, 1, 1, 1)
	celtDenormaliseBandsTo(nil, &bands[0], 2, nil, 2147483647, celtDenormaliseOutput{owner: out[:], owned: true}, 1, nil, 2147483647, 0, 1, 1, 1, 1)
	if out != [4]float32{77, 0, 0, 88} {
		t.Fatal("silence skips offset formation", out)
	}
	zeroBand := [1]int16{0}
	celtDenormaliseBandsOffsets(nil, &zeroBand[0], 0, nil, 2147483647, nil, 2147483647, nil, 2147483647, 0, 0, 1, 1, 1)
	celtDenormaliseBandsAt(nil, &zeroBand[0], 2, nil, 2147483647, &out[1], nil, 2147483647, 0, 0, 1, 1, 0)
	if out != [4]float32{77, 0, 0, 88} {
		t.Fatal("empty spans/guards", out)
	}
}

func TestDenormalisePointers(t *testing.T) {
	in := [6]float32{1, 2, 3, 4, 5, 6}
	logE := [3]float32{0, -100, 100}
	out := [10]float32{123, 9, 9, 9, 9, 9, 9, 9, 9, 456}
	Opus_denormalise_bands(nil, &denormalisePointerBands[0], 8, &in[0], &out[1], &logE[0], 1, 3, 1, 1, 0)
	if out[0] != 123 || out[9] != 456 || out[1] != 0 || out[2] != 0 || out[3] != 0 || out[4] != 0 || out[5] <= 0 || out[6] <= 0 || out[7] != 0 || out[8] != 0 {
		t.Fatalf("out=%v", out)
	}
	Opus_denormalise_bands(nil, &denormalisePointerBands[0], 8, nil, &out[1], nil, 1, 3, 1, 1, 1)
	for _, v := range out[1:9] {
		if v != 0 {
			t.Fatal("silence did not clear output")
		}
	}
	if out[0] != 123 || out[9] != 456 {
		t.Fatal("silence changed guards")
	}
}
