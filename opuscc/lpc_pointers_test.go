package opuscc

import (
	"math"
	"runtime"
	"slices"
	"testing"
	"unsafe"
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
				celtIIROwned(nil, ownedInput, inputOffset, coefficientOwner, coefficientOffset, n, ord, ownedMemory[1:], 0, ownedOutput, start)
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

func TestIIROwnedMemoryAliasPointers(t *testing.T) {
	for _, ord := range []int32{4, 8} {
		for _, n := range []int32{0, 1, 3, 4, 5, 8, 17} {
			start := ord + 1
			got := make([]float32, ord+n+2)
			for i := range got {
				got[i] = float32(i%7-3) * .125
			}
			got[0], got[len(got)-1] = 77, 88
			want := slices.Clone(got)
			coeff := make([]float32, ord)
			for i := range coeff {
				coeff[i] = float32(i+1) * .001
			}
			entropyInitGrowStack(12)
			runtime.GC()
			Opus_celt_iir(nil, &want[start], &coeff[0], &want[start], n, ord, &want[1], 0)
			celtIIROwned(nil, got, start, coeff, 0, n, ord, got[1:], 0, got, start)
			for i := range got {
				if math.Float32bits(got[i]) != math.Float32bits(want[i]) {
					t.Fatal("live IIR output/memory/history aliases", ord, n, i)
				}
			}
		}
	}
}

func TestIIROwnedEmptyInputPointers(t *testing.T) {
	output := []float32{77, 1, 2, 3, 4, 88}
	memory := []float32{77, 0, 0, 0, 0, 88}
	coeff := [4]float32{}
	celtIIROwned(nil, nil, math.MaxInt32, coeff[:], 0, 0, 4, memory[1:], 0, output, 5)
	if !slices.Equal(memory, []float32{77, 4, 3, 2, 1, 88}) || !slices.Equal(output, []float32{77, 1, 2, 3, 4, 88}) {
		t.Fatal("unused input keeps numeric output history/guards")
	}
	celtIIROwned(nil, nil, math.MaxInt32, nil, math.MaxInt32, 0, 0, nil, 0, nil, 0)
	zeroOrder := []float32{77, 0, 88}
	celtIIROwned(nil, []float32{.25}, 0, nil, math.MaxInt32, 1, 0, nil, 0, zeroOrder, 1)
	if !slices.Equal(zeroOrder, []float32{77, .25, 88}) {
		t.Fatal("unused zero-order coefficient owner")
	}
	want := "Fatal (internal) error in " + opusDiagnosticString(3330) + ", line 225: " + opusDiagnosticString(3390) + "\n"
	checkDiagnosticFatal(t, func() { celtIIROwned(nil, nil, math.MaxInt32, nil, 0, 1, 3, nil, 0, nil, 0) }, want)
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
			celtIIROwned(nil, owned, start, owned, start, n, ord, owned[1:], 0, owned, start)
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
	owned := celtPLCPitchSearchOwned(nil, nil, [2][]float32{input[1 : len(input)-1], nil}, 1, 0)
	if got != PLC_PITCH_LAG_MAX || owned != got || input != before {
		t.Fatal("silence/guards", got, owned)
	}
	got = celt_plc_pitch_search(nil, nil, &input[1], &input[1], 2, 0)
	owned = celtPLCPitchSearchOwned(nil, nil, [2][]float32{input[1 : len(input)-1], input[1 : len(input)-1]}, 2, 0)
	if got != PLC_PITCH_LAG_MAX || owned != got || input != before {
		t.Fatal("shared channels", got, owned)
	}
	for i := 1; i < len(input)-1; i++ {
		input[i] = float32(math.Sin(float64(i) * .17))
	}
	before = input
	for _, C := range []int32{0, 1, 2} {
		var right *float32
		var rightOwner []float32
		if C == 2 {
			right, rightOwner = &input[1], input[1:len(input)-1]
		}
		entropyInitGrowStack(12)
		runtime.GC()
		got = celt_plc_pitch_search(nil, nil, &input[1], right, C, 0)
		owned = celtPLCPitchSearchOwned(nil, nil, [2][]float32{input[1 : len(input)-1], rightOwner}, C, 0)
		if got != owned || input != before {
			t.Fatal("periodic PLC pitch histories/guards", C, got, owned)
		}
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

func TestPitchSearchSpansPointers(t *testing.T) {
	for _, n := range []int32{16, 17, 32} {
		for _, maxPitch := range []int32{8, 9, 16, 18} {
			x, y := make([]float32, n>>1), make([]float32, (n+maxPitch)>>1)
			for i := range x {
				x[i] = float32(i%7-3) / 17
			}
			for i := range y {
				y[i] = float32(i%13-6) / 23
			}
			a, b := [3]int32{77, -1, 88}, [3]int32{77, -1, 88}
			Opus_pitch_search(nil, &x[0], &y[0], n, maxPitch, &a[1], 0)
			pitchSearchSpans(nil, x, y, n, maxPitch, &b[1], 0)
			if a != b {
				t.Fatal("pitch search span/result guards", n, maxPitch, a, b)
			}
		}
	}
}

func TestPitchSearchOwnedPointers(t *testing.T) {
	for _, n := range []int32{16, 17, 32} {
		for _, maxPitch := range []int32{8, 9, 16, 18} {
			for _, historyOffset := range []int32{1, 5, 7, n + 4} {
				for _, aliasOutput := range []bool{false, true} {
					a := make([]float32, 2*n+maxPitch+12)
					for i := range a {
						a[i] = float32(i%13-6) / 37
					}
					a[0], a[len(a)-1] = 77, 88
					b := slices.Clone(a)
					old, owned := [3]int32{77, -1, 88}, [3]int32{77, -1, 88}
					oldResult, ownedResult := &old[1], &owned[1]
					if aliasOutput {
						oldResult, ownedResult = (*int32)(unsafe.Pointer(&a[5])), (*int32)(unsafe.Pointer(&b[5]))
					}
					entropyInitGrowStack(12)
					runtime.GC()
					Opus_pitch_search(nil, &a[5], &a[historyOffset], n, maxPitch, oldResult, 0)
					pitchSearchOwned(nil, b, 5, historyOffset, n, maxPitch, ownedResult, 0)
					if old != owned {
						t.Fatal("owned search result/guards", n, maxPitch, historyOffset)
					}
					for i := range a {
						if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
							t.Fatal("owned search full buffer/live result alias", n, maxPitch, historyOffset, aliasOutput, i)
						}
					}
				}
			}
		}
	}
	for _, origin := range []int32{math.MinInt32, math.MaxInt32} {
		history := [4]float32{}
		a, b := int32(-1), int32(-1)
		Opus_pitch_search(nil, nil, &history[0], 1, 8, &a, 0)
		pitchSearchOwned(nil, history[:], origin, 0, 1, 8, &b, 0)
		if a != 0 || a != b {
			t.Fatal("unused zero-span input origin", origin)
		}
	}
	want := "Fatal (internal) error in " + opusDiagnosticString(4659) + ", line 325: " + opusDiagnosticString(4675) + "\n"
	checkDiagnosticFatal(t, func() { pitchSearchOwned(nil, nil, math.MaxInt32, math.MinInt32, 0, 0, nil, 0) }, want)
	want = "Fatal (internal) error in " + opusDiagnosticString(4659) + ", line 326: " + opusDiagnosticString(4629) + "\n"
	checkDiagnosticFatal(t, func() { pitchSearchOwned(nil, nil, math.MaxInt32, math.MinInt32, 16, 0, nil, 0) }, want)
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

func TestPitchDownsampleKernelsPointers(t *testing.T) {
	for _, n := range []int32{8, 16} {
		for _, factor := range []int32{2, 4} {
			for _, C := range []int32{0, 1, 2} {
				left, right := make([]float32, n*factor), make([]float32, n*factor)
				for i := range left {
					left[i], right[i] = float32(i%13-6)/37, float32(i%7-3)/19
				}
				a := make([]float32, n+2)
				a[0], a[len(a)-1] = 77, 88
				b := slices.Clone(a)
				Opus_pitch_downsample(nil, &left[0], &right[0], &a[1], n, C, factor, 0)
				pitchDownsampleLeft(left, b[1:len(b)-1], n, factor)
				if C == 2 {
					pitchDownsampleRight(right, b[1:len(b)-1], n, factor)
				}
				pitchDownsampleFilter(nil, b[1:len(b)-1], n, 0)
				for i := range a {
					if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
						t.Fatal("downsample kernels/guards", n, factor, C, i)
					}
				}
			}
		}
	}
}

func TestPitchDownsampleOwnedPointers(t *testing.T) {
	for _, n := range []int32{8, 16} {
		for _, factor := range []int32{2, 4} {
			for _, C := range []int32{0, 1, 2} {
				for alias := 0; alias < 5; alias++ {
					a := make([]float32, 3*n*factor+n+4)
					for i := range a {
						a[i] = float32(i%13-6) / 37
					}
					a[0], a[len(a)-1] = 77, 88
					b := slices.Clone(a)
					l, r, o := int32(1), 1+n*factor+1, 1+2*n*factor+2
					switch alias {
					case 1:
						o = l
					case 2:
						o = r
					case 3:
						r, o = l+1, l+2
					case 4:
						r, o = l, l
					}
					var oldRight *float32
					var ownedRight []float32
					if C == 2 {
						oldRight, ownedRight = &a[r], b[r:r+n*factor]
					}
					entropyInitGrowStack(12)
					runtime.GC()
					Opus_pitch_downsample(nil, &a[l], oldRight, &a[o], n, C, factor, 0)
					pitchDownsampleOwned(nil, b[l:l+n*factor], ownedRight, b[o:o+n], n, C, factor, 0)
					for i := range a {
						if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
							t.Fatal("owned downsample full buffer/aliases/guards", n, factor, C, alias, i)
						}
					}
				}
			}
		}
	}
}

func TestPitchDownsampleRightOrderPointers(t *testing.T) {
	left := [16]float32{}
	for i := range left {
		left[i] = float32(i+1) / 17
	}
	a, b := [10]float32{77, 0, 0, 0, 0, 0, 0, 0, 0, 88}, [10]float32{77, 0, 0, 0, 0, 0, 0, 0, 0, 88}
	pitchDownsampleLeft(left[:], a[1:9], 8, 2)
	if !validationPanics(func() { pitchDownsampleOwned(nil, left[:], nil, b[1:9], 8, 2, 2, 0) }) {
		t.Fatal("missing stereo history must fail")
	}
	if a != b {
		t.Fatal("all left stores precede right view and whitening", a, b)
	}
	for _, C := range []int32{0, 1, -1, 3} {
		a, b = [10]float32{77, 0, 0, 0, 0, 0, 0, 0, 0, 88}, [10]float32{77, 0, 0, 0, 0, 0, 0, 0, 0, 88}
		Opus_pitch_downsample(nil, &left[0], nil, &a[1], 8, C, 2, 0)
		pitchDownsampleOwned(nil, left[:], nil, b[1:9], 8, C, 2, 0)
		if a != b {
			t.Fatal("nonstereo ignores absent right history", C)
		}
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

func TestAutocorrCanonicalWindowPointers(t *testing.T) {
	saved := window120
	defer func() { window120 = saved }()
	bound := celtAutocorrKnownWindow(&window120[0], 120)
	if !bound.owned || len(bound.owner) != len(window120) || &bound.owner[0] != &window120[0] {
		t.Fatal("canonical actual window owner")
	}
	for _, candidate := range []struct {
		pointer *float32
		overlap int32
	}{{&saved[0], 120}, {&window120[1], 119}, {&window120[0], 121}, {&window120[0], -1}, {nil, 0}} {
		if celtAutocorrKnownWindow(candidate.pointer, candidate.overlap).owned {
			t.Fatal("noncanonical/shifted/unsupported windows must stay legacy")
		}
	}
	// Go-only live window/input/output aliases: preserve sequential stores, not
	// a new C restrict-domain claim. Restore the canonical table after EACH pair.
	for _, overlap := range []int32{0, 2, 60, 120} {
		for _, lag := range []int32{0, 1, 3, 24} {
			for _, origin := range []int32{0, 1, 80, 95} {
				window120 = saved
				entropyInitGrowStack(12)
				runtime.GC()
				r := Opus__celt_autocorr(nil, &window120[0], &window120[origin], &window120[0], overlap, lag, 120, 0)
				want := window120
				window120 = saved
				r2 := celtAutocorrOwned(nil, window120[:], window120[origin:], &window120[0], overlap, lag, 120, 0)
				if r != r2 {
					t.Fatal("canonical autocorr return")
				}
				for i := range want {
					if math.Float32bits(window120[i]) != math.Float32bits(want[i]) {
						t.Fatal("canonical live window/input/output aliases", overlap, lag, origin, i)
					}
				}
				window120 = saved
			}
		}
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
					c := slices.Clone(a)
					entropyInitGrowStack(12)
					runtime.GC()
					r := Opus__celt_autocorr(nil, &a[1], &a[offset], &a[1], overlap, lag, n, 0)
					r2 := celtAutocorrSpan(nil, b[1:], &b[offset], &b[1], overlap, lag, n, 0)
					r3 := celtAutocorrOwned(nil, c[1:], c[offset:], &c[1], overlap, lag, n, 0)
					if r != r2 || r != r3 {
						t.Fatal("autocorr span return")
					}
					for i := range a {
						if math.Float32bits(a[i]) != math.Float32bits(b[i]) || math.Float32bits(a[i]) != math.Float32bits(c[i]) {
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
		checkDiagnosticFatal(t, func() { celtAutocorrOwned(nil, nil, nil, nil, -1, 0, n, 0) }, want)
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

func TestCombOwnedOutputPointers(t *testing.T) {
	window := [4]float32{.1, .25, .6, .9}
	for _, n := range []int32{0, 1, 3, 8, 17, 64} {
		for _, overlap := range []int32{0, min(n, 4)} {
			for _, outputOffset := range []int32{1, 40, 41, 130} {
				for _, gains := range [][2]float32{{0, 0}, {.3, 0}, {0, .4}, {.3, .4}, {-.4, -.4}} {
					a := make([]float32, 3*n+150)
					for i := range a {
						a[i] = float32(i%17-8) / 31
					}
					a[0], a[len(a)-1] = 77, 88
					b := slices.Clone(a)
					entropyInitGrowStack(12)
					runtime.GC()
					combFilterOwned(nil, &a[outputOffset], a, 40, 15, 17, n, gains[0], gains[1], 0, 1, &window[0], overlap, 0)
					combFilterBuffersOwned(nil, b, outputOffset, b, 40, 15, 17, n, gains[0], gains[1], 0, 1, &window[0], overlap, 0)
					for i := range a {
						if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
							t.Fatal("owned comb output/full-buffer aliases/guards", n, overlap, outputOffset, gains, i)
						}
					}
				}
			}
		}
	}
	combFilterBuffersOwned(nil, nil, 0, nil, 0, 0, 0, 0, 0, 0, -1, -1, nil, 0, 0)
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
			celtFIROwned(nil, input, ord, coefficientOwner, 1, owned[1:], n, ord, 0)
			if !slices.Equal(got, want) || !slices.Equal(span, want) || !slices.Equal(owned, want) {
				t.Fatal("typed FIR history/guards", ord, n)
			}
		}
	}
}

func TestFIROwnedOverlapPointers(t *testing.T) {
	for _, ord := range []int32{4, 8} {
		for _, n := range []int32{1, 4, 8, 12} {
			for _, aliasCoefficients := range []bool{false, true} {
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
				referenceCoefficients, ownedCoefficients := &coeff[1], coeff[:]
				if aliasCoefficients {
					referenceCoefficients, ownedCoefficients = &want[1], owned
				}
				Opus_celt_fir_c(nil, &want[ord+1], referenceCoefficients, &want[2], n, ord, 0)
				celtFIROwned(nil, owned, ord+1, ownedCoefficients, 1, owned[2:], n, ord, 0)
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

func TestFIROwnedUnusedCoefficientsPointers(t *testing.T) {
	input, output := []float32{77, .25, 88}, []float32{99, 0, 66}
	celtFIROwned(nil, input, 1, nil, math.MaxInt32, output[1:], 1, 0, 0)
	if !slices.Equal(input, []float32{77, .25, 88}) || !slices.Equal(output, []float32{99, .25, 66}) {
		t.Fatal("FIR zero-order coefficient owner/guards")
	}
	want := "Fatal (internal) error in " + opusDiagnosticString(3330) + ", line 157: " + opusDiagnosticString(3305) + "\n"
	checkDiagnosticFatal(t, func() { celtFIROwned(nil, input, 1, nil, math.MaxInt32, input[1:], 1, 4, 0) }, want)
}

func TestLPCOwnedPointers(t *testing.T) {
	for _, p := range []int32{0, 1, 2, 4, 8, 24} {
		for _, origin := range []int32{1, 2, p + 3} {
			for _, energy := range []float32{0, 1e-10, float32(math.NaN()), 1} {
				a := make([]float32, 2*p+5)
				for i := range a {
					a[i] = 77
				}
				a[1] = energy
				for i := int32(1); i <= p; i++ {
					a[1+i] = float32(math.Ldexp(1, -int(i)))
				}
				b := slices.Clone(a)
				entropyInitGrowStack(12)
				runtime.GC()
				Opus__celt_lpc(nil, &a[origin], &a[1], p)
				celtLPCOwned(nil, b, origin, b[1:], p)
				for i := range a {
					if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
						t.Fatal("LPC live correlation/output alias and guards", p, origin, energy, i)
					}
				}
			}
		}
	}
	celtLPCOwned(nil, nil, math.MaxInt32, []float32{1}, 0)
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
		span := slices.Clone(out)
		Opus__celt_lpc(nil, &out[1], &tc.ac[0], int32(len(tc.want)))
		celtLPCSpans(nil, span[1:len(span)-1], tc.ac)
		if !slices.Equal(span, out) {
			t.Fatal("LPC span full output/guards")
		}
		if !slices.Equal(out[1:len(out)-1], tc.want) || out[0] != 123 || out[len(out)-1] != 123 {
			t.Fatalf("ac=%v got=%v want=%v", tc.ac, out, tc.want)
		}
	}
	ac := float32(1)
	Opus__celt_lpc(nil, nil, &ac, 0)
	celtLPCSpans(nil, nil, []float32{ac})
}
