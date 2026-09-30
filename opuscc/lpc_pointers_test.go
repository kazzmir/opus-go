package opuscc

import (
	"math"
	"runtime"
	"slices"
	"testing"
)

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
