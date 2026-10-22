package opuscc

import (
	"math"
	"runtime"
	"slices"
	"testing"
)

// From interleave_pointers_test.go.
func TestInterleavePointers(t *testing.T) {
	interleave_hadamard(nil, nil, 0, 2, 1)
	interleave_hadamard(nil, nil, 2, 0, 0)
	for _, stride := range []int32{1, 2, 3, 4, 8, 16} {
		for _, n0 := range []int32{1, 2, 7, 22, 37} {
			for h := int32(0); h <= 1; h++ {
				if h != 0 && hadamardTestOrders[stride] == nil {
					continue
				}
				n := n0 * stride
				data := make([]float32, n+2)
				data[0] = 123
				data[n+1] = 456
				original := make([]uint32, n)
				want := make([]uint32, n)
				for i := range original {
					bits := uint32(i) * 0x9e3779b9
					if i%7 == 0 {
						bits = 0xff800001 + uint32(i)
					}
					original[i] = bits
					data[i+1] = math.Float32frombits(bits)
				}
				for i := int32(0); i < stride; i++ {
					row := i
					if h != 0 {
						row = hadamardTestOrders[stride][i]
					}
					for j := int32(0); j < n0; j++ {
						want[j*stride+i] = original[row*n0+j]
					}
				}
				interleave_hadamard(nil, &data[1], n0, stride, h)
				for i, bits := range want {
					if math.Float32bits(data[i+1]) != bits {
						t.Fatalf("n0=%d stride=%d h=%d i=%d", n0, stride, h, i)
					}
				}
				deinterleave_hadamard(nil, &data[1], n0, stride, h)
				for i, bits := range original {
					if math.Float32bits(data[i+1]) != bits {
						t.Fatal("round trip", n0, stride, h, i)
					}
				}
				if data[0] != 123 || data[n+1] != 456 {
					t.Fatal("guards")
				}
			}
		}
	}
}

// From deinterleave_pointers_test.go.
var hadamardTestOrders = map[int32][]int32{
	2: {1, 0}, 4: {3, 0, 2, 1}, 8: {7, 0, 4, 3, 6, 1, 5, 2}, 16: {15, 0, 8, 7, 12, 3, 11, 4, 14, 1, 9, 6, 13, 2, 10, 5},
}

func TestDeinterleavePointers(t *testing.T) {
	deinterleave_hadamard(nil, nil, 0, 2, 1)
	for _, offset := range []int32{math.MinInt32, 0, math.MaxInt32} {
		deinterleaveFrom(nil, quantBandScratch{offset: offset, owned: true}, 0, 2, 1)
	}
	wantFatal := "Fatal (internal) error in " + opusDiagnosticString(5312) + ", line 582: " + opusDiagnosticString(5405) + "\n"
	checkDiagnosticFatal(t, func() { deinterleaveFrom(nil, quantBandScratch{owned: true}, 0, 0, 0) }, wantFatal)
	for _, stride := range []int32{1, 2, 3, 4, 8, 16} {
		for _, n0 := range []int32{1, 2, 7, 22, 37} {
			for h := int32(0); h <= 1; h++ {
				if h != 0 && hadamardTestOrders[stride] == nil {
					continue
				}
				n := n0 * stride
				data := make([]float32, n+2)
				data[0] = 123
				data[n+1] = 456
				want := make([]uint32, n)
				for i := int32(0); i < stride; i++ {
					row := i
					if h != 0 {
						row = hadamardTestOrders[stride][i]
					}
					for j := int32(0); j < n0; j++ {
						idx := j*stride + i
						bits := uint32(idx) * 0x9e3779b9
						if idx%7 == 0 {
							bits = 0x7f800001 + uint32(idx)
						}
						data[idx+1] = math.Float32frombits(bits)
						want[row*n0+j] = bits
					}
				}
				span := append([]float32(nil), data...)
				owned := append([]float32(nil), data...)
				entropyInitGrowStack(12)
				runtime.GC()
				deinterleave_hadamard(nil, &data[1], n0, stride, h)
				deinterleaveSpans(span[1:n+1], n0, stride, h)
				deinterleaveFrom(nil, quantBandScratch{owner: owned, offset: 1, owned: true}, n0, stride, h)
				for i := range data {
					if math.Float32bits(data[i]) != math.Float32bits(span[i]) || math.Float32bits(data[i]) != math.Float32bits(owned[i]) {
						t.Fatal("deinterleave span full bits/NaN payloads/guards", n0, stride, h, i)
					}
				}
				for i, bits := range want {
					if math.Float32bits(data[i+1]) != bits {
						t.Fatalf("n0=%d stride=%d h=%d i=%d", n0, stride, h, i)
					}
				}
				if data[0] != 123 || data[n+1] != 456 {
					t.Fatal("guards")
				}
			}
		}
	}
}

// From exp_rotation_pointers_test.go.
func TestExpRotationOuterPointers(t *testing.T) {
	for _, n := range []int{8, 16, 64, 128} {
		for _, stride := range []int32{1, 2, 4} {
			for spread := int32(1); spread <= 3; spread++ {
				x := make([]float32, n+2)
				x[0], x[n+1] = 111, 222
				for i := 1; i <= n; i++ {
					x[i] = float32(math.Sin(float64(i)))
				}
				before := append([]float32(nil), x...)
				Opus_exp_rotation(nil, &x[1], int32(n), 1, stride, 1, spread)
				Opus_exp_rotation(nil, &x[1], int32(n), -1, stride, 1, spread)
				for i := 1; i <= n; i++ {
					if math.Abs(float64(x[i]-before[i])) > 2e-6 {
						t.Fatalf("round trip n=%d stride=%d spread=%d index=%d got=%g want=%g", n, stride, spread, i, x[i], before[i])
					}
				}
				if x[0] != 111 || x[n+1] != 222 {
					t.Fatal("sentinels changed")
				}
			}
		}
	}
	Opus_exp_rotation(nil, nil, 16, 1, 1, 8, 3)
	Opus_exp_rotation(nil, nil, 16, 1, 1, 1, SPREAD_NONE)
}

// From normalise_pointers_test.go.
func TestNormalisePointers(t *testing.T) {
	input := [4]int32{3, -4, 0, 12}
	for _, gain := range []float32{0, 0.5, 1, 2} {
		out := [6]float32{111, 0, 0, 0, 0, 222}
		g := float32(1) / float32(13) * gain
		for _, shift := range []int32{0, 3} {
			owned := out
			entropyInitGrowStack(12)
			runtime.GC()
			normalise_residual(nil, &input[0], &out[1], 4, 169, gain, shift)
			normaliseResidualOwned(nil, input[:], &owned[1], 4, 169, gain, shift)
			for i := range out {
				if math.Float32bits(out[i]) != math.Float32bits(owned[i]) {
					t.Fatal("residual pulse-owner full bits/guards", gain, shift, i)
				}
			}
			if input != [4]int32{3, -4, 0, 12} {
				t.Fatal("residual pulse input changed")
			}
			for i, value := range input {
				if out[i+1] != float32(value)*g {
					t.Fatalf("gain=%g shift=%d index=%d: %g", gain, shift, i, out[i+1])
				}
			}
			if out[0] != 111 || out[5] != 222 {
				t.Fatal("sentinels changed")
			}
		}
	}
	normalise_residual(nil, nil, nil, 0, 1, 1, 0)
	normaliseResidualOwned(nil, nil, nil, 0, 1, 1, 0)
	for _, input := range [][]float32{{0}, {3, 4}, {0, 0, 0}, {-12, 3, -4}, {1e-10, -1e-10}} {
		for _, gain := range []float32{0, 0.5, 1, 2} {
			x := append([]float32{111}, input...)
			x = append(x, 222)
			want := slices.Clone(input)
			var energy float32
			for _, v := range input {
				energy += float32(v * v)
			}
			scale := float32(1) / float32(math.Sqrt(float64(float32(1e-15)+energy))) * gain
			for i := range want {
				want[i] *= scale
			}
			owned := slices.Clone(x)
			entropyInitGrowStack(12)
			runtime.GC()
			Opus_renormalise_vector(nil, &x[1], int32(len(input)), gain, 0)
			renormaliseVectorOwned(nil, owned, 1, int32(len(input)), gain, 0)
			for i := range x {
				if math.Float32bits(x[i]) != math.Float32bits(owned[i]) {
					t.Fatal("owned renormalization full bits/guards", input, gain, i)
				}
			}
			if !slices.Equal(x[1:len(x)-1], want) || x[0] != 111 || x[len(x)-1] != 222 {
				t.Fatalf("input=%v gain=%g: got=%v want=%v", input, gain, x, want)
			}
		}
	}
	Opus_renormalise_vector(nil, nil, 0, 1, 0)
	for _, offset := range []int32{math.MinInt32, 0, math.MaxInt32} {
		renormaliseVectorOwned(nil, nil, offset, 0, 1, 0)
		renormaliseVectorOwned(nil, nil, offset, -1, 1, 0)
	}
}

// From denormalise_pointers_test.go.
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
				var want, got, numeric, span, owned, full [18]float32
				for i := range want {
					want[i] = float32(i%7-3) / 17
				}
				got = want
				numeric = want
				span = want
				owned = want
				full = want
				// Output overlaps both spectrum and (in one case) energy.
				Opus_denormalise_bands(nil, &denormalisePointerBands[0], 8, &want[2], &want[4], &want[energy], start, 3, 1, down, 0)
				entropyInitGrowStack(12)
				runtime.GC()
				celtDenormaliseBandsAt(nil, &denormalisePointerBands[0], 8, &got[0], 2, &got[4], &got[0], energy, start, 3, 1, down, 0)
				celtDenormaliseBandsOffsets(nil, &denormalisePointerBands[0], 8, &numeric[0], 2, &numeric[0], 4, &numeric[0], energy, start, 3, 1, down, 0)
				celtDenormaliseBandsTo(nil, &denormalisePointerBands[0], 8, &owned[0], 2, celtDenormaliseOutput{owner: owned[:], owned: true}, 4, &owned[0], energy, start, 3, 1, down, 0)
				celtDenormaliseBandsFromTo(nil, &denormalisePointerBands[0], 8, celtDenormaliseInput{owner: full[:], owned: true}, 2, celtDenormaliseOutput{owner: full[:], owned: true}, 4, &full[0], energy, start, 3, 1, down, 0)
				bound := int32(6)
				if down != 1 {
					bound = min(bound, 8/down)
				}
				celtDenormaliseSpans(nil, denormalisePointerBands[:], span[2:8], span[4:12], span[energy:energy+3], start, 3, 1, bound)
				for i := range got {
					if math.Float32bits(span[i]) != math.Float32bits(want[i]) || math.Float32bits(owned[i]) != math.Float32bits(want[i]) || math.Float32bits(full[i]) != math.Float32bits(want[i]) {
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
