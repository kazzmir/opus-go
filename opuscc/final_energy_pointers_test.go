package opuscc

import (
	"math"
	"runtime"
	"testing"
	"unsafe"
)

func TestFinalEnergyOwnedPointers(t *testing.T) {
	mode := OpusT_OpusCustomMode{FnbEBands: 3}
	quant := [5]int32{77, 0, 1, MAX_FINE_BITS, 88}
	priority := [5]int32{77, 1, 0, 0, 88}
	for _, C := range []int32{1, 2} {
		for _, budget := range []int32{0, 1, 2, 3, 4, 9} {
			a := [8]float32{77, 1, 2, 3, 4, 5, 6, 88}
			b := a
			ec := OpusT_ec_dec{Frng: 1 << 31, Fnbits_total: 33, Fend_window: 0xa5, Fnend_bits: 32}
			ownedEC, nilEC := ec, ec
			entropyInitGrowStack(12)
			runtime.GC()
			Opus_unquant_energy_finalise(nil, &mode, 0, 3, &a[1], &quant[1], &priority[1], budget, &ec, C)
			unquantFinalEnergyOwned(nil, &mode, 0, 3, &b[1], quant[1:4], priority[1:4], budget, &ownedEC, C)
			unquantFinalEnergyOwned(nil, &mode, 0, 3, nil, quant[1:4], priority[1:4], budget, &nilEC, C)
			if ec != ownedEC || ec != nilEC {
				t.Fatal("final energy full context/nil-output consumption", C, budget)
			}
			for i := range a {
				if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
					t.Fatal("final energy owned samples/guards", C, budget, i)
				}
			}
		}
	}
	mode.FnbEBands = 1
	a := OpusT_ec_dec{Fstorage: 8, Frng: 1 << 31, Fnbits_total: 0, Fend_window: 3, Fnend_bits: 32}
	b := a
	x, y := [4]float32{77, 0, 0, 88}, [4]float32{77, 0, 0, 88}
	prio := [1]int32{0}
	Opus_unquant_energy_finalise(nil, &mode, 0, 1, &x[1], &a.Fnbits_total, &prio[0], 2, &a, 2)
	unquantFinalEnergyOwned(nil, &mode, 0, 1, &y[1], unsafe.Slice(&b.Fnbits_total, 1), prio[:], 2, &b, 2)
	if x != y || a != b || y != [4]float32{77, .125, .0625, 88} {
		t.Fatal("live quantizer reload after each entropy bit", x, y)
	}
	mode.FnbEBands = 3
	for _, length := range []int32{0, 1, 2, math.MaxInt32} {
		a = OpusT_ec_dec{Frng: 1 << 31, Fnbits_total: 33, Fend_window: 0xa5, Fnend_bits: 32}
		b = a
		celtDecodeFinalEnergy(nil, &mode, nil, &quant[1], &priority[1], 0, 3, length, 2, &a)
		celtDecodeFinalEnergyOwned(nil, &mode, nil, quant[1:4], priority[1:4], 0, 3, length, 2, &b)
		if a != b {
			t.Fatal("decode final-energy tell/budget wrapping", length)
		}
	}
}

func TestFinalEnergyPointers(t *testing.T) {
	mode := OpusT_OpusCustomMode{FnbEBands: 3}
	quant := [3]int32{0, 0, MAX_FINE_BITS}
	priority := [3]int32{1, 0, 0}
	for _, budget := range []int32{1, 3, 4} {
		out := [8]float32{123, 0, 0, 0, 0, 0, 0, 456}
		dec := OpusT_ec_dec{Frng: 1 << 31, Fnbits_total: 33, Fend_window: 15, Fnend_bits: 32}
		skip := dec
		span, spanDec, nilSpanDec := out, dec, dec
		Opus_unquant_energy_finalise(nil, &mode, 0, 3, &out[1], &quant[0], &priority[0], budget, &dec, 2)
		Opus_unquant_energy_finalise(nil, &mode, 0, 3, nil, &quant[0], &priority[0], budget, &skip, 2)
		unquantFinalEnergySpans(nil, &mode, 0, 3, span[1:7], quant[:], priority[:], budget, &spanDec, 2)
		unquantFinalEnergySpans(nil, &mode, 0, 3, nil, quant[:], priority[:], budget, &nilSpanDec, 2)
		if span != out || spanDec != dec || nilSpanDec != dec {
			t.Fatal("final energy span/context/nil-output consumption")
		}
		want := [8]float32{123, 0, 0, 0, 0, 0, 0, 456}
		if budget >= 2 {
			want[2] = 0.25
			want[5] = 0.25
		}
		if budget >= 4 {
			want[1] = 0.25
			want[4] = 0.25
		}
		if out != want || dec != skip || dec.Fnbits_total != 33+budget/2*2 {
			t.Fatalf("budget=%d out=%v state=%+v skip=%+v", budget, out, dec, skip)
		}
	}
}
