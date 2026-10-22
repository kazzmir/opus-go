package opuscc

import (
	"math"
	"runtime"
	"testing"
	"unsafe"
)

// From coarse_energy_pointers_test.go.
func TestCoarseEnergyPointers(t *testing.T) {
	mode := OpusT_OpusCustomMode{FnbEBands: 4}
	energy := [10]float32{123, 9, -20, 2, 8, 9, -20, 2, 8, 456}
	dec := OpusT_ec_dec{Frng: 1 << 31, Fval: (1 << 31) - 1, Fnbits_total: 33}
	before := dec
	Opus_unquant_coarse_energy(nil, &mode, 1, 3, &energy[1], 1, &dec, 2, 0)
	if energy[0] != 123 || energy[9] != 456 || energy[1] != 9 || energy[4] != 8 || energy[5] != 9 || energy[8] != 8 || energy[2] != -1 || energy[6] != -1 || energy[3] != energy[7] || dec != before {
		t.Fatalf("energy=%v state=%+v", energy, dec)
	}
}

// From fine_energy_pointers_test.go.
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

// From final_energy_pointers_test.go.
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

// From noise_levels_pointers_test.go.
func TestNoiseLevelsPointers(t *testing.T) {
	for _, tc := range []struct {
		counter        int32
		noise, inverse [4]int32
	}{
		{0, [4]int32{99, 2512, 2426, 2399}, [4]int32{21688935, 854740, 885165, 894812}},
		{15, [4]int32{99, 2512, 2426, 2399}, [4]int32{21688935, 854740, 885165, 894812}},
		{16, [4]int32{194, 2506, 1928, 1599}, [4]int32{11058891, 856866, 1113678, 1342204}},
		{999, [4]int32{1963, 2500, 1608, 1209}, [4]int32{1093873, 858861, 1334924, 1775369}},
		{1000, [4]int32{1963, 2500, 1605, 1202}, [4]int32{1093873, 858861, 1337630, 1786073}},
		{1001, [4]int32{1963, 2500, 1605, 1202}, [4]int32{1093873, 858861, 1337630, 1786073}},
	} {
		var state OpusT_silk_VAD_state
		Opus_silk_VAD_Init(nil, &state)
		state.Fcounter = tc.counter
		state.FAnaState = [2]int32{123, -456}
		state.FHPstate = 789
		want := state
		want.FNL, want.Finv_NL = tc.noise, tc.inverse
		if tc.counter < 1000 {
			want.Fcounter++
		}
		energies := [4]int32{0, 2500, 5000, 2147483647}
		before := energies
		silk_VAD_GetNoiseLevels(nil, &energies, &state)
		if state != want || energies != before {
			t.Fatalf("counter=%d got=%+v want=%+v", tc.counter, state, want)
		}
	}
}

// From amp_log_pointers_test.go.
func TestAmpLogPointers(t *testing.T) {
	Opus_amp2Log2(nil, nil, 0, 0, nil, nil, 2)
	m := OpusT_OpusCustomMode{FnbEBands: 4}
	in := [8]float32{1, 2, 4, 8, .5, .25, .125, .0625}
	out := [8]float32{77, 77, 77, 77, 77, 77, 77, 77}
	Opus_amp2Log2(nil, &m, 2, 3, &in[0], &out[0], 2)
	if out[2] != -14 || out[6] != -14 || out[3] != 77 || out[7] != 77 {
		t.Fatal(out)
	}
	for _, pos := range []int{0, 1, 4, 5} {
		i := pos % 4
		want := float32(math.Log2(float64(in[pos]))) - Opus_eMeans[i]
		if math.Abs(float64(out[pos]-want)) > 1e-4 {
			t.Fatal(pos, out[pos], want)
		}
	}
	alias := in
	Opus_amp2Log2(nil, &m, 2, 3, &alias[0], &alias[0], 2)
	for _, pos := range []int{0, 1, 2, 4, 5, 6} {
		if math.Float32bits(alias[pos]) != math.Float32bits(out[pos]) {
			t.Fatal("alias", alias, out)
		}
	}
}

// From quant_final_pointers_test.go.
func TestQuantFinalPointers(t *testing.T) {
	Opus_quant_energy_finalise(nil, nil, 0, 3, nil, nil, nil, nil, 1, nil, 2)
	m := OpusT_OpusCustomMode{FnbEBands: 3}
	old := [6]float32{}
	err := [6]float32{-.5, 0, .5, -.25, .25, 0}
	quant := [3]int32{0, 1, 8}
	priority := [3]int32{1, 0, 0}
	var b [8]byte
	var e OpusT_ec_enc
	Opus_ec_enc_init(nil, &e, &b[0], 8)
	Opus_quant_energy_finalise(nil, &m, 0, 3, &old[0], &err[0], &quant[0], &priority[0], 2, &e, 2)
	if old != [6]float32{0, .125, 0, 0, .125, 0} || e.Fnbits_total != 35 {
		t.Fatal(old, e)
	}
	before := e
	Opus_quant_energy_finalise(nil, &m, 0, 3, nil, &err[0], &quant[0], &priority[0], 2, &e, 2)
	if e.Fnbits_total != before.Fnbits_total+2 {
		t.Fatal("nil old must still encode", e)
	}
}

// From quant_fine_pointers_test.go.
func TestQuantFinePointers(t *testing.T) {
	Opus_quant_fine_energy(nil, nil, 2, 2, nil, nil, nil, nil, nil, 2)
	m := OpusT_OpusCustomMode{FnbEBands: 3}
	old := [6]float32{}
	err := [6]float32{-.6, 0, .6, -.25, .25, 0}
	extra := [3]int32{2, 0, 2}
	var b [8]byte
	var e OpusT_ec_enc
	Opus_ec_enc_init(nil, &e, &b[0], 8)
	Opus_quant_fine_energy(nil, &m, 0, 3, &old[0], &err[0], nil, &extra[0], &e, 2)
	if old != [6]float32{-.375, 0, .375, -.125, 0, .125} || e.Fnbits_total != 41 {
		t.Fatal(old, err, e)
	}
	// C updates old before subtracting error, including when they alias.
	a := [3]float32{.2, .3, .4}
	before := a
	Opus_quant_fine_energy(nil, &m, 0, 3, &a[0], &a[0], nil, &extra[0], &e, 1)
	for i := range a {
		offset := []float32{.125, 0, .375}[i]
		if a[i] != float32(float32(before[i]+offset)-offset) {
			t.Fatal("alias", a, before)
		}
	}
}
