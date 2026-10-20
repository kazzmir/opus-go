package opuscc

import (
	"runtime"
	"testing"
)

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
