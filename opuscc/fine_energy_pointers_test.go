package opuscc

import (
	"math"
	"runtime"
	"testing"
	"unsafe"
)

func TestFineEnergyOwnedPointers(t *testing.T) {
	mode := OpusT_OpusCustomMode{FnbEBands: 5}
	extra := [7]int32{77, -1, 0, 1, 2, 14, 88}
	previous := [7]int32{77, 0, 1, 2, 0, 0, 88}
	for _, C := range []int32{1, 2} {
		for _, size := range []uint32{0, 1, 2, 8, 64} {
			for _, usePrevious := range []bool{false, true} {
				data := make([]byte, 64)
				for i := range data {
					data[i] = byte(i*37 + 19)
				}
				var ec OpusT_ec_dec
				Opus_ec_dec_init(nil, &ec, &data[0], uint32(len(data)))
				ec.Fstorage = size
				ownedEC := ec
				a := [12]float32{77, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 88}
				b := a
				var prev *int32
				var prevSpan []int32
				if usePrevious {
					prev, prevSpan = &previous[1], previous[1:6]
				}
				entropyInitGrowStack(12)
				runtime.GC()
				Opus_unquant_fine_energy(nil, &mode, 0, 5, &a[1], prev, &extra[1], &ec, C)
				unquantFineEnergyOwned(nil, &mode, 0, 5, &b[1], prevSpan, extra[1:6], &ownedEC, C)
				if ec != ownedEC {
					t.Fatal("fine energy full entropy context", C, size, usePrevious)
				}
				for i := range a {
					if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
						t.Fatal("fine energy owned samples/guards", C, size, usePrevious, i)
					}
				}
			}
		}
	}
	mode.FnbEBands = 1
	a := OpusT_ec_dec{Fstorage: 8, Frng: 1 << 31, Fnbits_total: 8, Fend_window: 0xa5, Fnend_bits: 32}
	b := a
	x, y := [3]float32{77, 0, 88}, [3]float32{77, 0, 88}
	Opus_unquant_fine_energy(nil, &mode, 0, 1, &x[1], nil, &a.Fnbits_total, &a, 1)
	unquantFineEnergyOwned(nil, &mode, 0, 1, &y[1], nil, unsafe.Slice(&b.Fnbits_total, 1), &b, 1)
	if x != y || a != b {
		t.Fatal("fine energy captures extra before aliased entropy mutation")
	}
}

func TestFineEnergyPointers(t *testing.T) {
	mode := OpusT_OpusCustomMode{FnbEBands: 3}
	extra := [3]int32{0, 2, 14}
	prev := [3]int32{0, 1, 0}
	for _, previous := range []*int32{nil, &prev[0]} {
		energy := [8]float32{123, 1, 2, 3, 4, 5, 6, 456}
		span := energy
		// Four available bits permit only the middle band's stereo refinement.
		dec := OpusT_ec_dec{Fstorage: 1, Frng: 1 << 31, Fnbits_total: 36, Fend_window: 15, Fnend_bits: 32}
		spanDec := dec
		Opus_unquant_fine_energy(nil, &mode, 0, 3, &energy[1], previous, &extra[0], &dec, 2)
		var prevSpan []int32
		if previous != nil {
			prevSpan = prev[:]
		}
		unquantFineEnergySpans(nil, &mode, 0, 3, span[1:7], prevSpan, extra[:], &spanDec, 2)
		if span != energy || spanDec != dec {
			t.Fatal("fine energy span/context/guards")
		}
		delta := float32(0.375)
		if previous != nil {
			delta /= 2
		}
		if energy != [8]float32{123, 1, 2 + delta, 3, 4, 5 + delta, 6, 456} || dec.Fnbits_total != 40 || dec.Fnend_bits != 28 {
			t.Fatalf("energy=%v state=%+v", energy, dec)
		}
	}
}
