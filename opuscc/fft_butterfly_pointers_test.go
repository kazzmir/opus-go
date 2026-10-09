package opuscc

import (
	"math"
	"runtime"
	"slices"
	"testing"
	"unsafe"
	"weak"
)

func TestMDCTConsumedStridePointers(t *testing.T) {
	singleton := float32(7)
	for _, stride := range []int32{-2147483648, -2, 0, 2, 2147483647} {
		view, origin := mdctStridedSlice(&singleton, 1, stride)
		if len(view) != 1 || origin != 0 || &view[0] != &singleton {
			t.Fatal("singleton extent/identity", stride)
		}
	}
	owner := []float32{11, 12, 13, 14, 15, 16, 17}
	forward, f := mdctStridedSlice(&owner[2], 3, 2)
	backward, b := mdctStridedSlice(&owner[6], 3, -2)
	zero, z := mdctStridedSlice(&owner[2], 3, 0)
	entropyInitGrowStack(12)
	runtime.GC()
	if len(forward) != 5 || len(backward) != 5 || len(zero) != 1 || f != 0 || b != 4 || z != 0 {
		t.Fatal("consumed stride geometry")
	}
	if &forward[f] != &owner[2] || &backward[b] != &owner[6] || &zero[z] != &owner[2] {
		t.Fatal("live owner identity")
	}
}

func TestFFTTablePointers(t *testing.T) {
	makeLookup := func() *OpusT_mdct_lookup {
		bitrev := []int16{0, 1, 2, 3}
		tw := []OpusT_kiss_twiddle_cpx{{Fr: 1}, {Fi: -1}, {Fr: -1}, {Fi: 1}}
		trig := make([]float32, 8)
		for i := range trig {
			trig[i] = float32(math.Cos(2 * math.Pi * (float64(i) + .125) / 16))
		}
		st := &OpusT_kiss_fft_state{Fnfft: 4, Fscale: .25, Fshift: -1, Ffactors: [16]int16{4, 1}, Fbitrev: &bitrev[0], Ftwiddles: &tw[0]}
		return &OpusT_mdct_lookup{Fn: 16, Fkfft: [4]*OpusT_kiss_fft_state{st}, Ftrig: &trig[0]}
	}
	l := makeLookup()
	entropyInitGrowStack(12)
	for i := 0; i < 3; i++ {
		runtime.GC()
	}
	st := l.Fkfft[0]
	if unsafe.Slice(st.Fbitrev, 4)[3] != 3 || unsafe.Slice(st.Ftwiddles, 4)[3].Fi != 1 {
		t.Fatal("table backing lifetime")
	}
	input := [4]OpusT_kiss_fft_cpx{{Fr: 1}}
	out := [6]OpusT_kiss_fft_cpx{}
	out[0].Fr = 77
	out[5].Fr = 88
	Opus_opus_fft_c(nil, st, st.Fbitrev, st.Ftwiddles, &input[0], &out[1])
	for _, v := range out[1:5] {
		if v != (OpusT_kiss_fft_cpx{Fr: .25}) {
			t.Fatal(out)
		}
	}
	if out[0].Fr != 77 || out[5].Fr != 88 {
		t.Fatal("FFT guards")
	}
	pcm := [12]float32{}
	spectrum := [8]float32{}
	Opus_clt_mdct_forward_c(nil, l, st.Fbitrev, st.Ftwiddles, &pcm[0], &spectrum[0], nil, 0, 0, 1, 0)
	Opus_clt_mdct_backward_c(nil, l, st.Fbitrev, st.Ftwiddles, &spectrum[0], &pcm[0], nil, 0, 0, 1, 0)
	var reversed [17]float32
	reversed[0] = 77
	reversed[16] = 88
	pcm[3] = 1
	Opus_clt_mdct_forward_c(nil, l, st.Fbitrev, st.Ftwiddles, &pcm[0], &spectrum[0], nil, 0, 0, 1, 0)
	Opus_clt_mdct_forward_c(nil, l, st.Fbitrev, st.Ftwiddles, &pcm[0], &reversed[15], nil, 0, 0, -2, 0)
	entropyInitGrowStack(12)
	runtime.GC()
	for i := 0; i < 8; i++ {
		if reversed[15-2*i] != spectrum[i] {
			t.Fatal("negative stride")
		}
	}
	if reversed[0] != 77 || reversed[16] != 88 {
		t.Fatal("reverse guards")
	}
	var expected, back [8]float32
	Opus_clt_mdct_backward_c(nil, l, st.Fbitrev, st.Ftwiddles, &spectrum[0], &expected[0], nil, 0, 0, 1, 0)
	Opus_clt_mdct_backward_c(nil, l, st.Fbitrev, st.Ftwiddles, &reversed[15], &back[0], nil, 0, 0, -2, 0)
	if back != expected {
		t.Fatal("negative input stride")
	}
	var collapsed float32
	Opus_clt_mdct_forward_c(nil, l, st.Fbitrev, st.Ftwiddles, &pcm[0], &collapsed, nil, 0, 0, 0, 0)
	if collapsed != spectrum[1] {
		t.Fatal("zero stride store order", collapsed, spectrum)
	}
}

func TestMDCTBackwardPointers(t *testing.T) {
	l := &mode48000_960_120.Fmdct
	st := l.Fkfft[3]
	input := [239]float32{}
	output := [182]float32{}
	output[0] = 77
	output[181] = 88
	entropyInitGrowStack(12)
	runtime.GC()
	Opus_clt_mdct_backward_c(nil, l, (*int16)(unsafe.Pointer(st.Fbitrev)), (*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(st.Ftwiddles)), &input[0], &output[1], &window120[0], 120, 3, 2, 0)
	if output[0] != 77 || output[181] != 88 {
		t.Fatal("guards")
	}
	for _, v := range output[1:181] {
		if v != 0 {
			t.Fatal("zero synthesis")
		}
	}
	Opus_clt_mdct_backward_c(nil, l, (*int16)(unsafe.Pointer(st.Fbitrev)), (*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(st.Ftwiddles)), &input[0], &output[1], nil, 0, 3, 2, 0)
}

func TestMDCTForwardPointers(t *testing.T) {
	l := &mode48000_960_120.Fmdct
	st := l.Fkfft[3]
	input := [240]float32{}
	output := [241]float32{}
	output[0] = 77
	output[240] = 88
	entropyInitGrowStack(12)
	runtime.GC()
	Opus_clt_mdct_forward_c(nil, l, (*int16)(unsafe.Pointer(st.Fbitrev)), (*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(st.Ftwiddles)), &input[0], &output[1], &window120[0], 120, 3, 2, 0)
	if output[0] != 77 || output[240] != 88 {
		t.Fatal("guards")
	}
	for _, v := range output[1:240] {
		if v != 0 {
			t.Fatal("zero spectrum")
		}
	}
	for i := range input {
		input[i] = float32(i%13) - 6
	}
	var expected [120]float32
	Opus_clt_mdct_forward_c(nil, l, (*int16)(unsafe.Pointer(st.Fbitrev)), (*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(st.Ftwiddles)), &input[0], &expected[0], &window120[0], 120, 3, 1, 0)
	Opus_clt_mdct_forward_c(nil, l, (*int16)(unsafe.Pointer(st.Fbitrev)), (*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(st.Ftwiddles)), &input[0], &input[0], &window120[0], 120, 3, 1, 0)
	if !slices.Equal(input[:120], expected[:]) {
		t.Fatal("fold-before-output alias")
	}
}

func TestMDCTLookupPointers(t *testing.T) {
	makeLookup := func() *OpusT_mdct_lookup {
		trig := make([]float32, 12)
		for i := range trig {
			trig[i] = float32(i) + .25
		}
		l := &OpusT_mdct_lookup{Fn: 16, Fmaxshift: 1, Ftrig: &trig[0]}
		l.Fkfft[0] = &OpusT_kiss_fft_state{Fnfft: 4}
		l.Fkfft[1] = &OpusT_kiss_fft_state{Fnfft: 2}
		return l
	}
	l := makeLookup()
	entropyInitGrowStack(12)
	for i := 0; i < 3; i++ {
		runtime.GC()
	}
	if l.Fkfft[0].Fnfft != 4 || l.Fkfft[1].Fnfft != 2 || unsafe.Slice(l.Ftrig, 12)[11] != 11.25 {
		t.Fatal("lookup backing lifetime")
	}
	if unsafe.Offsetof(l.Fkfft) != 8 || unsafe.Offsetof(l.Ftrig) != 8+4*unsafe.Sizeof(l.Ftrig) {
		t.Fatal("C layout")
	}
}

func TestMiniFFTRPointers(t *testing.T) {
	st := Opus_mini_kiss_fftr_alloc(nil, 16, 0, nil, nil)
	input := [16]float32{1}
	out := [11]OpusT_mini_kiss_fft_cpx{}
	out[0].Fr = 77
	out[10].Fr = 88
	entropyInitGrowStack(12)
	runtime.GC()
	Opus_mini_kiss_fftr(nil, st, &input[0], &out[1])
	for _, v := range out[1:10] {
		if v.Fr != 1 || v.Fi != 0 {
			t.Fatal("impulse spectrum", out)
		}
	}
	if out[0].Fr != 77 || out[10].Fr != 88 || input != [16]float32{1} {
		t.Fatal("guards/input")
	}
}

//go:noinline
func miniFFTRForeignTemporary(st *OpusT_mini_kiss_fftr_state) weak.Pointer[[8]OpusT_mini_kiss_fft_cpx] {
	tmp := new([8]OpusT_mini_kiss_fft_cpx)
	tmp[0].Fr = 123
	st.Ftmpbuf = &tmp[0]
	return weak.Make(tmp)
}

func TestMiniFFTRScannedHeaderPointers(t *testing.T) {
	st := Opus_mini_kiss_fftr_alloc(nil, 16, 0, nil, nil)
	foreign := miniFFTRForeignTemporary(st)
	entropyInitGrowStack(12)
	runtime.GC()
	runtime.GC()
	if foreign.Value() == nil || st.Ftmpbuf.Fr != 123 {
		t.Fatal("Go-created FFT header must scan its typed pointer fields")
	}
	runtime.KeepAlive(st)
}

func TestMiniFFTRTemporaryGeometryPointers(t *testing.T) {
	if unsafe.Offsetof(OpusT_mini_kiss_fft_state{}.Ftwiddles)+unsafe.Sizeof(OpusT_mini_kiss_fft_cpx{}) != unsafe.Sizeof(OpusT_mini_kiss_fft_state{}) {
		t.Fatal("mini FFT flexible tail geometry")
	}
	for _, n := range []int32{2, 4, 6, 16} {
		st := Opus_mini_kiss_fftr_alloc(nil, n, 0, nil, nil)
		var subsize OpusT_size_t
		Opus_mini_kiss_fft_alloc(nil, n/2, 0, nil, &subsize)
		want := (*OpusT_mini_kiss_fft_cpx)(unsafe.Add(unsafe.Pointer(st.Fsubstate), subsize))
		if st.Ftmpbuf != want {
			t.Fatal("mini FFT temporary displacement", n)
		}
		st.Ftmpbuf.Fr = 77
		entropyInitGrowStack(12)
		runtime.GC()
		if st.Ftmpbuf.Fr != 77 {
			t.Fatal("mini FFT temporary owner")
		}
	}
}

func TestMiniFFTRSuperPrefixPointers(t *testing.T) {
	for _, n := range []int32{2, 3, 8} {
		owner := make([]OpusT_mini_kiss_fft_cpx, n+n/2)
		p := miniFFTRSuperTwiddles(&owner[0], n)
		entropyInitGrowStack(12)
		runtime.GC()
		if p != &owner[n] {
			t.Fatal("consumed super-twiddle owner")
		}
		p.Fr = 77
		if owner[n].Fr != 77 || owner[n-1].Fr != 0 {
			t.Fatal("super-twiddle alias/guard")
		}
	}
	var owner [2]OpusT_mini_kiss_fft_cpx
	if miniFFTRSuperTwiddles(nil, 0) != nil || miniFFTRSuperTwiddles(&owner[0], 1) != &owner[1] || miniFFTRSuperTwiddles(&owner[1], -1) != &owner[0] {
		t.Fatalf("empty/backward native super-twiddle boundary: nil=%p forward=%p want=%p backward=%p want=%p", miniFFTRSuperTwiddles(nil, 0), miniFFTRSuperTwiddles(&owner[0], 1), &owner[1], miniFFTRSuperTwiddles(&owner[1], -1), &owner[0])
	}
}

func TestMiniFFTRAllocPointers(t *testing.T) {
	var needed OpusT_size_t
	Opus_mini_kiss_fftr_alloc(nil, 8, 0, nil, &needed)
	header := unsafe.Sizeof(OpusT_mini_kiss_fftr_state{})
	if needed != uint64(header)+296+48 {
		t.Fatal("size query", needed)
	}
	backing := make([]uint64, (needed+7)/8+2)
	for i := range backing {
		backing[i] = 0xa5a5a5a5a5a5a5a5
	}
	mem := (*byte)(unsafe.Pointer(&backing[1]))
	capacity := needed - 1
	before := slices.Clone(backing)
	if Opus_mini_kiss_fftr_alloc(nil, 8, 0, mem, &capacity) != nil || capacity != needed || !slices.Equal(backing, before) {
		t.Fatal("undersized")
	}
	st := Opus_mini_kiss_fftr_alloc(nil, 8, 0, mem, &capacity)
	if st != (*OpusT_mini_kiss_fftr_state)(unsafe.Pointer(mem)) || backing[0] != before[0] || backing[len(backing)-1] != before[len(backing)-1] {
		t.Fatal("caller storage")
	}
	owned := Opus_mini_kiss_fftr_alloc(nil, 16, 1, nil, nil)
	entropyInitGrowStack(12)
	runtime.GC()
	if owned.Fsubstate.Fnfft != 8 || owned.Fsubstate.Finverse != 1 || owned.Fsuper_twiddles.Fi != 0.9238795 {
		t.Fatal("owned interiors", owned.Fsubstate, *owned.Fsuper_twiddles)
	}
	if uintptr(unsafe.Pointer(owned.Fsubstate))-uintptr(unsafe.Pointer(owned)) != header {
		t.Fatal("header size")
	}
}

func TestMiniFFTStorageTailPointers(t *testing.T) {
	for _, extra := range []uint64{0, 8, 24} {
		header, tail := miniFFTStorage[OpusT_mini_kiss_fft_state](uint64(unsafe.Sizeof(OpusT_mini_kiss_fft_state{})) + extra)
		if uint64(len(tail)) != extra {
			t.Fatal("typed mini FFT tail length")
		}
		if extra != 0 && unsafe.SliceData(tail) != (*byte)(unsafe.Add(unsafe.Pointer(header), unsafe.Sizeof(*header))) {
			t.Fatal("typed mini FFT tail displacement")
		}
		entropyInitGrowStack(12)
		runtime.GC()
		header.Fnfft = 123
		if extra != 0 {
			tail[0] = 77
		}
		runtime.KeepAlive(header)
	}
}

func TestMiniFFTAllocPointers(t *testing.T) {
	var needed OpusT_size_t
	if Opus_mini_kiss_fft_alloc(nil, 8, 0, nil, &needed) != nil || needed != 328 {
		t.Fatal("size query", needed)
	}
	backing := make([]uint64, (needed+7)/8+2)
	for i := range backing {
		backing[i] = 0xa5a5a5a5a5a5a5a5
	}
	mem := (*byte)(unsafe.Pointer(&backing[1]))
	capacity := needed - 1
	before := slices.Clone(backing)
	if Opus_mini_kiss_fft_alloc(nil, 8, 0, mem, &capacity) != nil || capacity != needed || !slices.Equal(backing, before) {
		t.Fatal("undersized")
	}
	st := Opus_mini_kiss_fft_alloc(nil, 8, 0, mem, &capacity)
	if st != (*OpusT_mini_kiss_fft_state)(unsafe.Pointer(mem)) || backing[0] != before[0] || backing[len(backing)-1] != before[len(backing)-1] {
		t.Fatal("caller buffer")
	}
	owned := Opus_mini_kiss_fft_alloc(nil, 16, 1, nil, nil)
	entropyInitGrowStack(12)
	runtime.GC()
	tw := unsafe.Slice(&owned.Ftwiddles[0], 16)
	if owned.Fnfft != 16 || owned.Finverse != 1 || tw[0].Fr != 1 || tw[4].Fi != 1 {
		t.Fatal("owned state")
	}
}

func miniPointerFixture(n, inverse int32) *OpusT_mini_kiss_fft_state {
	backing := make([]uint64, (264+8*int(n)+7)/8)
	st := (*OpusT_mini_kiss_fft_state)(unsafe.Pointer(&backing[0]))
	st.Fnfft = n
	st.Finverse = inverse
	kf_factor(nil, n, &st.Ffactors)
	tw := unsafe.Slice(&st.Ftwiddles[0], n)
	for i := range tw {
		phase := -2 * math.Pi * float64(i) / float64(n)
		if inverse != 0 {
			phase = -phase
		}
		tw[i] = OpusT_mini_kiss_fft_cpx{Fr: float32(math.Cos(phase)), Fi: float32(math.Sin(phase))}
	}
	return st
}

func TestMiniFFTPointers(t *testing.T) {
	st := miniPointerFixture(4, 0)
	inverse := miniPointerFixture(4, 1)
	in := [4]OpusT_mini_kiss_fft_cpx{{Fr: 1, Fi: 2}, {Fr: -3, Fi: 4}, {Fr: 5, Fi: -6}, {Fr: 7, Fi: 8}}
	out := [6]OpusT_mini_kiss_fft_cpx{{Fr: 77}, {}, {}, {}, {}, {Fr: 88}}
	round := [4]OpusT_mini_kiss_fft_cpx{}
	entropyInitGrowStack(12)
	runtime.GC()
	Opus_mini_kiss_fft(nil, st, &in[0], &out[1])
	Opus_mini_kiss_fft(nil, inverse, &out[1], &round[0])
	if out[0].Fr != 77 || out[5].Fr != 88 {
		t.Fatal("guards")
	}
	for i := range in {
		if round[i].Fr != 4*in[i].Fr || round[i].Fi != 4*in[i].Fi {
			t.Fatal("unnormalized round trip", round)
		}
	}
}

func TestMiniFFTStridePointers(t *testing.T) {
	for _, n := range []int32{2, 5, 12, 60} {
		for _, stride := range []int32{1, 3} {
			st := miniPointerFixture(n, 0)
			in := make([]OpusT_mini_kiss_fft_cpx, (n-1)*stride+1)
			in[0].Fr = 1
			before := slices.Clone(in)
			out := make([]OpusT_mini_kiss_fft_cpx, n+2)
			out[0].Fr = 77
			out[n+1].Fr = 88
			entropyInitGrowStack(12)
			runtime.GC()
			Opus_mini_kiss_fft_stride(nil, st, &in[0], &out[1], stride)
			if out[0].Fr != 77 || out[n+1].Fr != 88 || !slices.Equal(in, before) {
				t.Fatal("guards/input")
			}
			for _, v := range out[1 : n+1] {
				if v.Fr != 1 || v.Fi != 0 {
					t.Fatal(n, stride, v)
				}
			}
		}
	}
}

func TestMiniFFTWorkPointers(t *testing.T) {
	for _, n := range []int32{2, 3, 4, 5, 8, 12, 60, 120} {
		for _, inverse := range []int32{0, 1} {
			st := miniPointerFixture(n, inverse)
			in := make([]OpusT_mini_kiss_fft_cpx, n)
			in[0].Fr = 1
			out := make([]OpusT_mini_kiss_fft_cpx, n+2)
			out[0].Fr = 77
			out[n+1].Fr = 88
			entropyInitGrowStack(12)
			runtime.GC()
			kf_work(nil, &out[1], &in[0], 1, 1, st.Ffactors[:], st)
			if out[0].Fr != 77 || out[n+1].Fr != 88 {
				t.Fatal("guards")
			}
			for _, v := range out[1 : n+1] {
				if v.Fr != 1 || v.Fi != 0 {
					t.Fatal(n, inverse, v)
				}
			}
		}
	}
}

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
