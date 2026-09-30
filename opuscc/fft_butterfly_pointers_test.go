package opuscc

import (
	"math"
	"runtime"
	"slices"
	"testing"
)

// Explicit intermediate conversions are the oracle: multiply then add,
// not an ARM64 fused multiply-add. Both radix-3 implementations need this.
func TestRadix3ScratchRounding(t *testing.T) {
	tw := [2]OpusT_kiss_twiddle_cpx{{Fr: 1}, {Fr: -.5, Fi: -.8660254}}
	seed := uint32(1234567)
	for trial := 0; trial < 128; trial++ {
		var input [3]OpusT_kiss_fft_cpx
		for i := range input {
			seed = seed*1664525 + 1013904223
			input[i].Fr = float32(int32(seed)) / 1234567
			seed = seed*1664525 + 1013904223
			input[i].Fi = float32(int32(seed)) / 7654321
		}
		s1 := fftMul(input[1], tw[0])
		s2 := fftMul(input[2], tw[0])
		sumR := float32(s1.Fr + s2.Fr)
		sumI := float32(s1.Fi + s2.Fi)
		baseR := float32(input[0].Fr - float32(sumR*.5))
		baseI := float32(input[0].Fi - float32(sumI*.5))
		pR := float32(float32(s1.Fr-s2.Fr) * tw[1].Fi)
		pI := float32(float32(s1.Fi-s2.Fi) * tw[1].Fi)
		want := [3]OpusT_kiss_fft_cpx{{Fr: input[0].Fr + sumR, Fi: input[0].Fi + sumI}, {Fr: baseR - pI, Fi: baseI + pR}, {Fr: baseR + pI, Fi: baseI - pR}}
		for _, mini := range []bool{false, true} {
			got := input
			if mini {
				kf_bfly31(nil, &got[0], 1, &tw[0], 1)
			} else {
				kf_bfly3(nil, &got[0], 1, &tw[0], 1, 1, 3)
			}
			for i := range got {
				if math.Float32bits(got[i].Fr) != math.Float32bits(want[i].Fr) || math.Float32bits(got[i].Fi) != math.Float32bits(want[i].Fi) {
					t.Fatalf("trial=%d mini=%v bin=%d got=%+v want=%+v", trial, mini, i, got[i], want[i])
				}
			}
		}
	}
}

func TestMiniButterfly5Pointers(t *testing.T) {
	kf_bfly51(nil, nil, 1, nil, 0)
	tw := [3]OpusT_mini_kiss_fft_cpx{{Fr: 1}, {Fr: .30901699, Fi: -.95105652}, {Fr: -.80901699, Fi: -.58778525}}
	out := [7]OpusT_mini_kiss_fft_cpx{{Fr: 77}, {Fr: 1}, {}, {}, {}, {}, {Fr: 88}}
	entropyInitGrowStack(12)
	runtime.GC()
	kf_bfly51(nil, &out[1], 1, &tw[0], 1)
	if out[0].Fr != 77 || out[6].Fr != 88 {
		t.Fatal("guards")
	}
	for i := 1; i <= 5; i++ {
		if out[i].Fr != 1 || out[i].Fi != 0 {
			t.Fatal(out)
		}
	}
}

func TestMiniButterfly3Pointers(t *testing.T) {
	kf_bfly31(nil, nil, 1, nil, 0)
	tw := [2]OpusT_mini_kiss_fft_cpx{{Fr: 1}, {Fr: -.5, Fi: -.8660254}}
	out := [5]OpusT_mini_kiss_fft_cpx{{Fr: 77}, {Fr: 1}, {}, {}, {Fr: 88}}
	entropyInitGrowStack(12)
	runtime.GC()
	kf_bfly31(nil, &out[1], 1, &tw[0], 1)
	if out[0].Fr != 77 || out[4].Fr != 88 {
		t.Fatal("guards")
	}
	for i := 1; i <= 3; i++ {
		if out[i].Fr != 1 || out[i].Fi != 0 {
			t.Fatal(out)
		}
	}
}

func TestMiniButterfly4Pointers(t *testing.T) {
	kf_bfly41(nil, nil, 1, nil, 0, 0)
	tw := [1]OpusT_mini_kiss_fft_cpx{{Fr: 1}}
	for _, inverse := range []int32{0, 1, -3} {
		out := [6]OpusT_mini_kiss_fft_cpx{{Fr: 77}, {}, {Fr: 1}, {}, {}, {Fr: 88}}
		entropyInitGrowStack(12)
		runtime.GC()
		kf_bfly41(nil, &out[1], 1, &tw[0], 1, inverse)
		sign := float32(-1)
		if inverse != 0 {
			sign = 1
		}
		if out[0].Fr != 77 || out[5].Fr != 88 || out[1].Fr != 1 || out[2].Fi != sign || out[3].Fr != -1 || out[4].Fi != -sign {
			t.Fatal(out)
		}
	}
}

func TestMiniButterfly2Pointers(t *testing.T) {
	kf_bfly21(nil, nil, 1, nil, 0)
	out := [6]OpusT_mini_kiss_fft_cpx{{Fr: 77}, {Fr: 1}, {Fr: 2}, {Fr: 3}, {Fr: 4}, {Fr: 88}}
	tw := [2]OpusT_mini_kiss_fft_cpx{{Fr: 1}, {Fr: 1}}
	entropyInitGrowStack(12)
	runtime.GC()
	kf_bfly21(nil, &out[1], 1, &tw[0], 2)
	if out != [6]OpusT_mini_kiss_fft_cpx{{Fr: 77}, {Fr: 4}, {Fr: 6}, {Fr: -2}, {Fr: -2}, {Fr: 88}} {
		t.Fatal(out)
	}
}

func TestFFTInversePointers(t *testing.T) {
	st := OpusT_kiss_fft_state{Fnfft: 4, Fscale: .25, Fshift: -1, Ffactors: [16]int16{4, 1}}
	rev := [4]int16{0, 1, 2, 3}
	in := [4]OpusT_kiss_fft_cpx{{Fr: 1}}
	out := [6]OpusT_kiss_fft_cpx{{Fr: 77}, {}, {}, {}, {}, {Fr: 88}}
	entropyInitGrowStack(12)
	runtime.GC()
	Opus_opus_ifft_c(nil, &st, &rev[0], nil, &in[0], &out[1])
	if out[0].Fr != 77 || out[5].Fr != 88 {
		t.Fatal("guards")
	}
	for i := 1; i <= 4; i++ {
		if out[i].Fr != 1 || out[i].Fi != 0 {
			t.Fatal(out)
		}
	}
	// Forward/inverse scales cancel for the full four-point chain.
	round := [4]OpusT_kiss_fft_cpx{}
	Opus_opus_fft_c(nil, &st, &rev[0], nil, &out[1], &round[0])
	if round != in {
		t.Fatal("round trip", round)
	}
}

func TestFFTForwardPointers(t *testing.T) {
	st := OpusT_kiss_fft_state{Fnfft: 4, Fscale: .25, Fshift: -1, Ffactors: [16]int16{4, 1}}
	rev := [4]int16{0, 1, 2, 3}
	in := [4]OpusT_kiss_fft_cpx{{Fr: 1}}
	out := [6]OpusT_kiss_fft_cpx{{Fr: 77}, {}, {}, {}, {}, {Fr: 88}}
	entropyInitGrowStack(12)
	runtime.GC()
	Opus_opus_fft_c(nil, &st, &rev[0], nil, &in[0], &out[1])
	if out[0].Fr != 77 || out[5].Fr != 88 {
		t.Fatal("guards")
	}
	for i := 1; i <= 4; i++ {
		if out[i].Fr != .25 || out[i].Fi != 0 {
			t.Fatal(out)
		}
	}
}

func TestFFTImplPointers(t *testing.T) {
	state := OpusT_kiss_fft_state{Fnfft: 4, Fshift: -1, Ffactors: [16]int16{4, 1}}
	out := [6]OpusT_kiss_fft_cpx{{Fr: 77}, {Fr: 1}, {Fr: 1}, {Fr: 1}, {Fr: 1}, {Fr: 88}}
	entropyInitGrowStack(12)
	runtime.GC()
	Opus_opus_fft_impl(nil, &state, nil, &out[1])
	if out[0].Fr != 77 || out[5].Fr != 88 || out[1].Fr != 4 || out[2].Fr != 0 {
		t.Fatal(out)
	}
}

func TestFFTButterfly5Pointers(t *testing.T) {
	kf_bfly5(nil, nil, 1, nil, 4, 0, 20)
	tw := make([]OpusT_kiss_twiddle_cpx, 13)
	for i := range tw {
		tw[i].Fr = 1
	}
	out := make([]OpusT_kiss_fft_cpx, 22)
	out[0].Fr = 77
	out[21].Fr = 88
	for i := 1; i <= 4; i++ {
		out[i].Fr = 1
	}
	entropyInitGrowStack(12)
	runtime.GC()
	kf_bfly5(nil, &out[1], 1, &tw[0], 4, 1, 20)
	if out[0].Fr != 77 || out[21].Fr != 88 {
		t.Fatal("guards")
	}
	for i := 1; i <= 20; i++ {
		if out[i].Fr != 1 || out[i].Fi != 0 {
			t.Fatal(out)
		}
	}
}

func TestFFTButterfly3Pointers(t *testing.T) {
	kf_bfly3(nil, nil, 1, nil, 4, 0, 12)
	tw := make([]OpusT_kiss_twiddle_cpx, 7)
	for i := range tw {
		tw[i].Fr = 1
	}
	out := make([]OpusT_kiss_fft_cpx, 14)
	out[0].Fr = 77
	out[13].Fr = 88
	for i := 1; i < 13; i++ {
		out[i].Fr = 1
	}
	entropyInitGrowStack(12)
	runtime.GC()
	kf_bfly3(nil, &out[1], 1, &tw[0], 4, 1, 12)
	if out[0].Fr != 77 || out[13].Fr != 88 {
		t.Fatal("guards")
	}
	for i := 1; i <= 4; i++ {
		if out[i].Fr != 3 || out[i+4].Fr != 0 || out[i+8].Fr != 0 {
			t.Fatal(out)
		}
	}
}

func TestFFTButterfly4Pointers(t *testing.T) {
	kf_bfly4(nil, nil, 1, nil, 1, 0, 4)
	out := make([]OpusT_kiss_fft_cpx, 10)
	out[0].Fr = 77
	out[9].Fr = 88
	for i := 1; i < 9; i++ {
		out[i].Fr = 1
	}
	entropyInitGrowStack(12)
	runtime.GC()
	kf_bfly4(nil, &out[1], 1, nil, 1, 2, 999)
	if out[0].Fr != 77 || out[9].Fr != 88 || out[1].Fr != 4 || out[5].Fr != 4 || out[2].Fr != 0 {
		t.Fatal(out)
	}
	tw := make([]OpusT_kiss_twiddle_cpx, 10)
	for i := range tw {
		tw[i].Fr = 1
	}
	out = make([]OpusT_kiss_fft_cpx, 18)
	for i := 1; i < 17; i++ {
		out[i].Fr = 1
	}
	kf_bfly4(nil, &out[1], 1, &tw[0], 4, 1, 16)
	for i := 1; i <= 4; i++ {
		if out[i].Fr != 4 {
			t.Fatal(out)
		}
	}
}

func TestFFTButterfly2Pointers(t *testing.T) {
	kf_bfly2(nil, nil, 4, 0)
	data := make([]OpusT_kiss_fft_cpx, 18)
	for i := range data {
		data[i] = OpusT_kiss_fft_cpx{Fr: float32(i), Fi: float32(-i)}
	}
	before := slices.Clone(data)
	kf_bfly2(nil, &data[1], 4, 2)
	if data[0] != before[0] || data[17] != before[17] || data[1].Fr != 6 || data[5].Fr != -4 {
		t.Fatal(data)
	}
}
