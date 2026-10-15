package opuscc

import (
	"math"
	"runtime"
	"slices"
	"testing"
)

func TestIIROwnedOutputPointers(t *testing.T) {
	for _, ord := range []int32{0, 4, 8} {
		for _, n := range []int32{0, 1, 3, 4, 5, 8, 17} {
			if ord == 0 && n >= 4 {
				continue
			} // Existing four-lane assertion.
			for _, inplace := range []bool{false, true} {
				start := ord + 1
				got := make([]float32, ord+n+2)
				for i := range got {
					got[i] = float32(i%7-3) * .125
				}
				got[0], got[len(got)-1] = 77, 88
				want := slices.Clone(got)
				memory := make([]float32, ord+2)
				for i := range memory {
					memory[i] = float32(i%3-1) * .03125
				}
				memory[0], memory[len(memory)-1] = 77, 88
				wantMemory := slices.Clone(memory)
				spanOutput, spanMemory := slices.Clone(got), slices.Clone(memory)
				ownedOutput, ownedMemory := slices.Clone(got), slices.Clone(memory)
				coeff := make([]float32, ord+1)
				for i := range coeff {
					coeff[i] = float32(i+1) * .001
				}
				input := make([]float32, n+1)
				for i := range input {
					input[i] = float32(i%5-2) * .0625
				}
				gx, wx := &input[0], &input[0]
				if inplace {
					gx, wx = &got[start], &want[start]
				}
				coefficientOwner := append([]float32{77}, coeff...)
				wantDen, gotDen, coefficientOffset := &coeff[0], &coefficientOwner[0], int32(1)
				entropyInitGrowStack(12)
				runtime.GC()
				Opus_celt_iir(nil, wx, wantDen, &want[start], n, ord, &wantMemory[1], 0)
				spanInput := input[:n]
				if inplace {
					spanInput = spanOutput[start : start+n]
				}
				celtIIRSpans(nil, spanInput, coeff[:ord], spanOutput[start:start+n], spanMemory[1:1+ord], n, ord)
				for i := int32(0); i < ord; i++ {
					spanMemory[1+i] = spanOutput[start+n-1-i]
				}
				ownedInput, inputOffset := input, int32(0)
				if inplace {
					ownedInput, inputOffset = ownedOutput, start
				}
				celtIIRInputOwned(nil, ownedInput, inputOffset, gotDen, coefficientOffset, n, ord, &ownedMemory[1], 0, ownedOutput, start)
				if inplace {
					celtIIRAt(nil, &got[0], start, gotDen, coefficientOffset, nil, n, ord, &memory[1], 0, got, start)
				} else {
					celtIIROutputOwned(nil, gx, &coeff[0], nil, n, ord, &memory[1], 0, got, start)
				}
				for i := range got {
					if math.Float32bits(got[i]) != math.Float32bits(want[i]) || math.Float32bits(spanOutput[i]) != math.Float32bits(want[i]) || math.Float32bits(ownedOutput[i]) != math.Float32bits(want[i]) {
						t.Fatal("owned IIR output", ord, n, inplace, i)
					}
				}
				for i := range memory {
					if math.Float32bits(memory[i]) != math.Float32bits(wantMemory[i]) || math.Float32bits(spanMemory[i]) != math.Float32bits(wantMemory[i]) || math.Float32bits(ownedMemory[i]) != math.Float32bits(wantMemory[i]) {
						t.Fatal("owned IIR memory", ord, n, inplace, i)
					}
				}
			}
		}
	}
}

func TestIIROwnedEmptyInputPointers(t *testing.T) {
	output := []float32{77, 1, 2, 3, 4, 88}
	memory := []float32{77, 0, 0, 0, 0, 88}
	coeff := [4]float32{}
	celtIIRInputOwned(nil, nil, math.MaxInt32, &coeff[0], 0, 0, 4, &memory[1], 0, output, 5)
	if !slices.Equal(memory, []float32{77, 4, 3, 2, 1, 88}) || !slices.Equal(output, []float32{77, 1, 2, 3, 4, 88}) {
		t.Fatal("unused input keeps numeric output history/guards")
	}
	want := "Fatal (internal) error in " + opusDiagnosticString(3330) + ", line 225: " + opusDiagnosticString(3390) + "\n"
	checkDiagnosticFatal(t, func() { celtIIRInputOwned(nil, nil, math.MaxInt32, nil, 0, 1, 3, nil, 0, nil, 0) }, want)
}

func TestIIROwnedCoefficientAliasPointers(t *testing.T) {
	for _, ord := range []int32{4, 8} {
		for _, n := range []int32{8, 17} {
			start := ord + 1
			got := make([]float32, ord+n+2)
			for i := range got {
				got[i] = float32(i%5-2) * .0625
			}
			got[0], got[len(got)-1] = 77, 88
			want := slices.Clone(got)
			owned := slices.Clone(got)
			entropyInitGrowStack(12)
			runtime.GC()
			Opus_celt_iir(nil, &want[start], &want[start], &want[start], n, ord, &want[1], 0)
			celtIIRAt(nil, &got[0], start, &got[0], start, nil, n, ord, &got[1], 0, got, start)
			celtIIRInputOwned(nil, owned, start, &owned[0], start, n, ord, &owned[1], 0, owned, start)
			for i := range got {
				if math.Float32bits(got[i]) != math.Float32bits(want[i]) || math.Float32bits(owned[i]) != math.Float32bits(want[i]) {
					t.Fatal("IIR coefficient/input/output/memory aliases", ord, n, i)
				}
			}
			if got[0] != 77 || got[len(got)-1] != 88 {
				t.Fatal("aliased IIR guards")
			}
		}
	}
}

func TestPLCPitchSearchPointers(t *testing.T) {
	input := [DEC_PITCH_BUF_SIZE + 2]float32{}
	input[0] = 77
	input[len(input)-1] = 88
	before := input
	entropyInitGrowStack(12)
	runtime.GC()
	got := celt_plc_pitch_search(nil, nil, &input[1], nil, 1, 0)
	if got != PLC_PITCH_LAG_MAX || input != before {
		t.Fatal("silence/guards", got)
	}
	got = celt_plc_pitch_search(nil, nil, &input[1], &input[1], 2, 0)
	if got != PLC_PITCH_LAG_MAX || input != before {
		t.Fatal("shared channels", got)
	}
}

func TestRemoveDoublingPointers(t *testing.T) {
	input := [16]float32{}
	period := [3]int32{77, 100, 88}
	entropyInitGrowStack(12)
	runtime.GC()
	gain := Opus_remove_doubling(nil, &input[0], 16, 4, 16, &period[1], 9, 0.46, 0)
	if gain != 0 || period != [3]int32{77, 14, 88} {
		t.Fatal(gain, period)
	}
	period[1] = 4
	gain = Opus_remove_doubling(nil, &input[0], 16, 5, 16, &period[1], 0, 0, 0)
	if gain != 0 || period != [3]int32{77, 5, 88} {
		t.Fatal("minimum", gain, period)
	}
}

func TestPitchSearchPointers(t *testing.T) {
	x := [16]float32{}
	y := [32]float32{}
	pitch := [3]int32{77, -1, 88}
	entropyInitGrowStack(12)
	runtime.GC()
	Opus_pitch_search(nil, &x[0], &y[0], 16, 16, &pitch[1], 0)
	if pitch != [3]int32{77, 0, 88} {
		t.Fatal(pitch)
	}
	x[0] = 1
	y[3] = 1
	Opus_pitch_search(nil, &x[0], &y[0], 16, 16, &pitch[1], 0)
	if pitch[0] != 77 || pitch[2] != 88 || pitch[1] < 0 || pitch[1] >= 16 {
		t.Fatal(pitch)
	}
}

func TestPitchDownsamplePointers(t *testing.T) {
	left := [16]float32{}
	right := [16]float32{}
	out := [10]float32{}
	out[0] = 77
	out[9] = 88
	entropyInitGrowStack(12)
	runtime.GC()
	Opus_pitch_downsample(nil, &left[0], &right[0], &out[1], 8, 2, 2, 0)
	if out[0] != 77 || out[9] != 88 {
		t.Fatal("guards", out)
	}
	for _, x := range out[1:9] {
		if x != 0 {
			t.Fatal("silence", out)
		}
	}
	// Any channel count other than 2 must leave the optional right input untouched.
	Opus_pitch_downsample(nil, &left[0], nil, &out[1], 8, 0, 2, 0)
}

func TestAutocorrPointers(t *testing.T) {
	input := [8]float32{1, 2, 3, 4, 5, 6, 7, 8}
	out := [6]float32{77, 0, 0, 0, 0, 88}
	window := [4]float32{0.25, 0.5, 0.75, 1}
	entropyInitGrowStack(12)
	runtime.GC()
	r := Opus__celt_autocorr(nil, &input[0], &out[1], &window[0], 4, 3, 8, 0)
	signal := [8]float32{0.25, 1, 2.25, 4, 5, 4.5, 3.5, 2}
	if r != 0 || out[0] != 77 || out[5] != 88 {
		t.Fatal(r, out)
	}
	for k := 0; k <= 3; k++ {
		sum := float32(0)
		for i := k; i < 8; i++ {
			sum += float32(signal[i] * signal[i-k])
		}
		if out[k+1] != sum {
			t.Fatal(k, out, sum)
		}
	}
	one := float32(3)
	r = Opus__celt_autocorr(nil, &one, &out[1], nil, 0, 0, 1, 0)
	if r != 0 || out[1] != 9 {
		t.Fatal("single sample")
	}
}

func TestAutocorrSpanPointers(t *testing.T) {
	for _, n := range []int32{8, 32} {
		for _, lag := range []int32{0, 1, 3} {
			for _, overlap := range []int32{0, 2, n / 2, n} {
				for _, offset := range []int32{0, 1, n + 1} {
					a := make([]float32, n+6)
					for i := range a {
						a[i] = float32(i-7) / 37
					}
					b := slices.Clone(a)
					entropyInitGrowStack(12)
					runtime.GC()
					r := Opus__celt_autocorr(nil, &a[1], &a[offset], &a[1], overlap, lag, n, 0)
					r2 := celtAutocorrSpan(nil, b[1:], &b[offset], &b[1], overlap, lag, n, 0)
					if r != r2 {
						t.Fatal("autocorr span return")
					}
					for i := range a {
						if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
							t.Fatal("autocorr span/window/output aliases/guards", n, lag, overlap, offset, i)
						}
					}
				}
			}
		}
	}
	for _, n := range []int32{0, 1} {
		line, text := "302", opusDiagnosticString(3419)
		if n == 1 {
			line, text = "303", opusDiagnosticString(3441)
		}
		want := "Fatal (internal) error in " + opusDiagnosticString(3330) + ", line " + line + ": " + text + "\n"
		checkDiagnosticFatal(t, func() { Opus__celt_autocorr(nil, nil, nil, nil, -1, 0, n, 0) }, want)
		checkDiagnosticFatal(t, func() { celtAutocorrSpan(nil, nil, nil, nil, -1, 0, n, 0) }, want)
	}
}

func TestIIRPointers(t *testing.T) {
	input := [8]float32{1, 2, 3, 4, 5, 6, 7, 8}
	coeff := [4]float32{0, 0, 0, 0}
	out := [10]float32{}
	out[0] = 77
	out[9] = 88
	mem := [6]float32{77, 1, 2, 3, 4, 88}
	entropyInitGrowStack(12)
	runtime.GC()
	Opus_celt_iir(nil, &input[0], &coeff[0], &out[1], 8, 4, &mem[1], 0)
	for i := range input {
		if out[i+1] != input[i] {
			t.Fatal("identity", out)
		}
	}
	if mem != [6]float32{77, 8, 7, 6, 5, 88} || out[0] != 77 || out[9] != 88 {
		t.Fatal("history or guards", mem, out)
	}
	// A zero-length call still updates memory from the preceding output history.
	history := [5]float32{1, 2, 3, 4, 99}
	Opus_celt_iir(nil, nil, &coeff[0], &history[4], 0, 4, &mem[1], 0)
	if mem != [6]float32{77, 4, 3, 2, 1, 88} {
		t.Fatal(mem)
	}
	singleton := float32(17)
	Opus_celt_iir(nil, &singleton, &coeff[0], &history[4], 1, 4, &mem[1], 0)
	if mem != [6]float32{77, 17, 4, 3, 2, 88} || history != [5]float32{1, 2, 3, 4, 17} {
		t.Fatal("short backward tail", mem, history)
	}
	// The empty order-zero tail consumes no pointers.
	Opus_celt_iir(nil, nil, nil, nil, 0, 0, nil, 0)
}

func TestCombHistoryOwnerPointers(t *testing.T) {
	for _, n := range []int32{0, 1, 4, 12} {
		for _, overlap := range []int32{0, 4} {
			if overlap > n {
				continue
			}
			for _, gains := range [][2]float32{{0, 0}, {0, 0.5}, {0.4, 0}, {0.3, 0.6}} {
				var left [200]float32
				for i := range left {
					left[i] = float32(i-99) / 113
				}
				right := left
				numeric := left
				window := [4]float32{0.1, 0.3, 0.6, 0.9}
				Opus_comb_filter(nil, &left[128], &left[128], 20, 31, n, gains[0], gains[1], 0, 1, &window[0], overlap, 0)
				entropyInitGrowStack(12)
				runtime.GC()
				combFilterWithHistory(nil, &right[128], &right[128], 20, 31, n, gains[0], gains[1], 0, 1, &window[0], overlap, 0, right[:], 128)
				combFilterOwned(nil, &numeric[128], numeric[:], 128, 20, 31, n, gains[0], gains[1], 0, 1, &window[0], overlap, 0)
				if left != right || left != numeric {
					t.Fatal("owned comb live alias/guards", n, overlap, gains)
				}
			}
		}
	}
}

func TestCombSourceCopyPointers(t *testing.T) {
	owner := [6]float32{77, .25, .5, 1, 2, 88}
	reference := owner
	Opus_comb_filter(nil, &reference[2], &reference[1], 0, 0, 3, 0, 0, -1, -1, nil, 0, 0)
	combFilterOwned(nil, &owner[2], owner[:], 1, 0, 0, 3, 0, 0, -1, -1, nil, 0, 0)
	if owner != reference {
		t.Fatal("owned comb overlapping copy/guards")
	}
	legacy := [2]float32{.75, -.25}
	var output [2]float32
	combFilterWithHistory(nil, &output[0], &legacy[0], 0, 0, 2, 0, 0, -1, -1, nil, 0, 0, owner[:], 1)
	if output != legacy {
		t.Fatal("legacy zero-gain copy must still use explicit x")
	}
	combFilterOwned(nil, nil, nil, 0, 0, 0, 0, 0, 0, -1, -1, nil, 0, 0)
	if !combFilterOwnerMatches(owner[:], 1, &owner[1]) || combFilterOwnerMatches(owner[:], 1, &legacy[0]) || combFilterOwnerMatches(owner[:], int32(len(owner)), nil) || combFilterOwnerMatches(nil, 0, nil) {
		t.Fatal("readable exact owner matching")
	}
	mode := OpusT_OpusCustomMode{FshortMdctSize: 1}
	state := OpusT_OpusCustomDecoder{Fpostfilter_tapset_old: -1, Fpostfilter_tapset: -1}
	before := output
	celtDecodePostfilterFirstWithHistory(nil, &state, &mode, &output[0], 0, owner[:], 1)
	celtDecodePostfilterTailWithHistory(nil, &state, &mode, &output[0], 2, 0, 0, -1, 0, owner[:], 1)
	if output != before {
		t.Fatal("foreign history retains legacy in-place zero-gain skip")
	}
}

func TestFIRHistoryOwnerPointers(t *testing.T) {
	for _, ord := range []int32{0, 4, 8} {
		for _, n := range []int32{0, 1, 4, 12} {
			if ord == 0 && n >= 4 {
				continue // existing unrolled correlation asserts ord >= 3
			}
			input := make([]float32, ord+n+2)
			for i := range input {
				input[i] = float32(i-9) / 7
			}
			coeff := make([]float32, ord+1)
			for i := range coeff {
				coeff[i] = float32(i+1) / 31
			}
			want, got := make([]float32, n+2), make([]float32, n+2)
			want[0], want[n+1], got[0], got[n+1] = 77, 88, 77, 88
			span := slices.Clone(got)
			owned := slices.Clone(got)
			coefficientOwner := append([]float32{77}, coeff...)
			reversed := make([]float32, ord)
			for i := int32(0); i < ord; i++ {
				reversed[i] = coeff[ord-i-1]
			}
			Opus_celt_fir_c(nil, &input[ord], &coeff[0], &want[1], n, ord, 0)
			entropyInitGrowStack(12)
			runtime.GC()
			celtFIRWithHistory(nil, &input[ord], &coeff[0], &got[1], n, ord, 0, input[:ord+n])
			celtFIRSpans(nil, input[:ord+n], reversed, span[1:n+1], n, ord)
			celtFIROwned(nil, input, ord, &coefficientOwner[0], 1, owned[1:], n, ord, 0)
			if !slices.Equal(got, want) || !slices.Equal(span, want) || !slices.Equal(owned, want) {
				t.Fatal("typed FIR history/guards", ord, n)
			}
		}
	}
}

func TestFIROwnedOverlapPointers(t *testing.T) {
	for _, ord := range []int32{4, 8} {
		for _, n := range []int32{1, 4, 8, 12} {
			owned := make([]float32, ord+n+4)
			for i := range owned {
				owned[i] = float32(i%9-4) * .0625
			}
			owned[0], owned[len(owned)-1] = 77, 88
			want := slices.Clone(owned)
			coeff := make([]float32, ord+2)
			for i := range coeff {
				coeff[i] = float32(i) * .00390625
			}
			entropyInitGrowStack(12)
			runtime.GC()
			Opus_celt_fir_c(nil, &want[ord+1], &coeff[1], &want[2], n, ord, 0)
			celtFIROwned(nil, owned, ord+1, &coeff[0], 1, owned[2:], n, ord, 0)
			for i := range owned {
				if math.Float32bits(owned[i]) != math.Float32bits(want[i]) {
					t.Fatal("owned FIR live overlap", ord, n, i)
				}
			}
			if owned[0] != 77 || owned[len(owned)-1] != 88 {
				t.Fatal("owned FIR overlap guards")
			}
		}
	}
}

func TestFIRPointers(t *testing.T) {
	input := [12]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	coeff := [4]float32{1, 2, 3, 4}
	out := [10]float32{}
	out[0] = 77
	out[9] = 88
	entropyInitGrowStack(12)
	runtime.GC()
	Opus_celt_fir_c(nil, &input[4], &coeff[0], &out[1], 8, 4, 0)
	for i := 0; i < 8; i++ {
		want := input[4+i]
		for j := 0; j < 4; j++ {
			want += float32(coeff[3-j] * input[i+j])
		}
		if out[i+1] != want {
			t.Fatal(i, out, want)
		}
	}
	if out[0] != 77 || out[9] != 88 {
		t.Fatal("guard")
	}
	// Tail-only, zero-order filtering requires no coefficients or input history.
	Opus_celt_fir_c(nil, &input[0], nil, &out[1], 3, 0, 0)
	for i := 0; i < 3; i++ {
		if out[i+1] != input[i] {
			t.Fatal("identity")
		}
	}
}

func TestLPCPointers(t *testing.T) {
	for _, tc := range []struct{ ac, want []float32 }{
		{[]float32{0, 1, 2}, []float32{0, 0}},
		{[]float32{1e-10, 1, 2}, []float32{0, 0}},
		{[]float32{float32(math.NaN()), 1, 2}, []float32{0, 0}},
		{[]float32{1, 0.5}, []float32{-0.5}},
		{[]float32{1, 0.5, 0.25}, []float32{-0.5, 0}},
		{[]float32{1, 1, 1, 1}, []float32{-1, 0, 0}},
	} {
		out := make([]float32, len(tc.want)+2)
		for i := range out {
			out[i] = 123
		}
		Opus__celt_lpc(nil, &out[1], &tc.ac[0], int32(len(tc.want)))
		if !slices.Equal(out[1:len(out)-1], tc.want) || out[0] != 123 || out[len(out)-1] != 123 {
			t.Fatalf("ac=%v got=%v want=%v", tc.ac, out, tc.want)
		}
	}
	ac := float32(1)
	Opus__celt_lpc(nil, nil, &ac, 0)
}
