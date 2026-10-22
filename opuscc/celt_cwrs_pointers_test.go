package opuscc

import (
	"runtime"
	"slices"
	"testing"
	"unsafe"
)

// From cwrs_pointers_test.go.
func TestCWRSPointers(t *testing.T) {
	for n := int32(2); n <= 6; n++ {
		for k := int32(1); k <= 6; k++ {
			count := celtPVQU(min(n, k), max(n, k)) + celtPVQU(min(n, k+1), max(n, k+1))
			for index := uint32(0); index < count; index++ {
				out := make([]int32, n+2)
				out[0] = 123
				out[n+1] = 456
				owned := slices.Clone(out)
				got := cwrsi(nil, n, k, index, &out[1])
				ownedEnergy := cwrsiVector(nil, n, k, index, celtPulseVector{owner: owned[1 : n+1], owned: true})
				if ownedEnergy != got || !slices.Equal(owned, out) {
					t.Fatal("CWRS owned full vector/energy/guards", n, k, index)
				}
				var total int32
				var energy float32
				for _, v := range out[1 : n+1] {
					total += max(v, -v)
					energy += float32(v) * float32(v)
				}
				if total != k || got != energy || out[0] != 123 || out[n+1] != 456 {
					t.Fatalf("n=%d k=%d index=%d out=%v energy=%g", n, k, index, out, got)
				}
			}
		}
	}
	dec := OpusT_ec_dec{Frng: 1 << 31, Fval: (1 << 31) - 1, Fnbits_total: 33}
	out := [4]int32{123, 0, 0, 456}
	owned, ownedDec := out, dec
	energy := Opus_decode_pulses(nil, &out[1], 2, 1, &dec)
	entropyInitGrowStack(12)
	runtime.GC()
	ownedEnergy := decodePulsesOwned(nil, owned[1:3], 2, 1, &ownedDec)
	if ownedEnergy != energy || owned != out || ownedDec != dec {
		t.Fatal("CWRS owned complete decoder/vector/energy")
	}
	if energy != 1 || out[0] != 123 || out[3] != 456 {
		t.Fatalf("energy=%g out=%v", energy, out)
	}
}

// From cwrs_index_pointers_test.go.
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
			before := slices.Clone(y)
			if index == 0 {
				entropyInitGrowStack(12)
				runtime.GC()
			}
			if got := icwrsVector(nil, n, celtPulseVector{owner: y[1 : n+1], owned: true}); got != index || !slices.Equal(y, before) {
				t.Fatal("owned CWRS rank/root/immutability", n, k, index, got)
			}
			if got := icwrs(nil, n, &y[1]); got != index || y[0] != 77 || y[n+1] != 88 {
				t.Fatal(n, k, index, got, y)
			}
		}
	}
	// Small exact stack-owned input.
	y := [2]int32{-1, 0}
	if got := icwrs(nil, 2, unsafe.SliceData(y[:])); got != 3 || icwrsVector(nil, 2, celtPulseVector{owner: y[:], owned: true}) != 3 {
		t.Fatal(y, got)
	}
}

// From cwrs_owned_boundaries_test.go.
func TestCWRSOwnedBoundariesPointers(t *testing.T) {
	legacy := int32(7)
	if (celtPulseVector{legacy: &legacy, owned: true}).span(0) != nil {
		t.Fatal("no owned-to-legacy fallback")
	}
	var caught any
	func() { defer func() { caught = recover() }(); cwrsiVector(nil, 1, 0, 0, celtPulseVector{owned: true}) }()
	if _, ok := caught.(runtime.Error); !ok {
		t.Fatal("CWRS full output view before assertions", caught)
	}
	want := "Fatal (internal) error in " + opusDiagnosticString(3634) + ", line 473: " + opusDiagnosticString(3649) + "\n"
	checkDiagnosticFatal(t, func() { cwrsiVector(nil, 1, 0, 0, celtPulseVector{owner: []int32{0}, owned: true}) }, want)
	want = "Fatal (internal) error in " + opusDiagnosticString(3634) + ", line 474: " + opusDiagnosticString(3672) + "\n"
	checkDiagnosticFatal(t, func() { cwrsiVector(nil, 0, 1, 0, celtPulseVector{owned: true}) }, want)
	want = "Fatal (internal) error in " + opusDiagnosticString(3634) + ", line 448: " + opusDiagnosticString(3610) + "\n"
	checkDiagnosticFatal(t, func() { icwrsVector(nil, 1, celtPulseVector{owned: true}) }, want)
	want = "Fatal (internal) error in " + opusDiagnosticString(3634) + ", line 463: " + opusDiagnosticString(3649) + "\n"
	checkDiagnosticFatal(t, func() { encodePulsesOwned(nil, nil, 2, 0, nil) }, want)
	var bytes [16]byte
	var enc OpusT_ec_enc
	Opus_ec_enc_init(nil, &enc, &bytes[0], 16)
	before := enc
	caught = nil
	func() { defer func() { caught = recover() }(); encodePulsesOwned(nil, nil, 2, 1, &enc) }()
	if _, ok := caught.(runtime.Error); !ok || enc != before || bytes != [16]byte{} {
		t.Fatal("pulse rank view before entropy writes", caught)
	}
	// Bounded invalid Go storage: entropy must be consumed BEFORE output clipping.
	dec := OpusT_ec_dec{Frng: 1 << 31, Fval: (1 << 31) - 1, Fnbits_total: 33}
	expected := dec
	total := celtPVQU(1, 2) + celtPVQU(2, 2)
	Opus_ec_dec_uint(nil, &expected, total)
	short := []int32{123}
	caught = nil
	func() { defer func() { caught = recover() }(); decodePulsesOwned(nil, short, 2, 1, &dec) }()
	if _, ok := caught.(runtime.Error); !ok || dec != expected || short[0] != 123 {
		t.Fatal("CWRS entropy-before-output view", caught, dec, expected, short)
	}
}
