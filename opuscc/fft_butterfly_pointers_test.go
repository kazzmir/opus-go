package opuscc

import (
	"runtime"
	"slices"
	"testing"
)

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
