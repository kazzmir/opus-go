package opuscc

import (
	"runtime"
	"slices"
	"testing"
	"unsafe"
)

// From down2_pointers_test.go.
func TestDown2Pointers(t *testing.T) {
	input := []int16{-32768, 32767, 0, 1, -1, 100, 20000, -30000, 0, 0, 1234}
	want := []int16{-7523, 5678, 3876, 412, -6094}
	initial := [2]int32{12345, -98765}
	wantState := [2]int32{-27341915, 5347117}
	for _, n := range []int32{10, 11} {
		state := initial
		out := [7]int16{111, 0, 0, 0, 0, 0, 222}
		Opus_silk_resampler_down2(nil, &state, &out[1], &input[0], n)
		if !slices.Equal(out[1:6], want) || state != wantState || out[0] != 111 || out[6] != 222 {
			t.Fatalf("n=%d output=%v state=%v", n, out, state)
		}
	}
	state := initial
	inPlace := slices.Clone(input)
	Opus_silk_resampler_down2(nil, &state, &inPlace[0], &inPlace[0], 4)
	Opus_silk_resampler_down2(nil, &state, &inPlace[2], &inPlace[4], 6)
	if !slices.Equal(inPlace[:5], want) || state != wantState {
		t.Fatal("in-place/chunked downsampling differs")
	}
	for _, n := range []int32{0, 1} {
		Opus_silk_resampler_down2(nil, &state, nil, nil, n)
		if state != wantState {
			t.Fatal("incomplete pair changed state")
		}
	}
}

// From down2_3_pointers_test.go.
func TestDownTwoThirdsPointers(t *testing.T) {
	s := [6]int32{1, 2, 3, 4, 5, 6}
	before := s
	Opus_silk_resampler_down2_3(nil, &s, nil, nil, 0)
	if s != before {
		t.Fatal("empty state")
	}
	for _, n := range []int{1, 2, 3, 4, 5, 479, 480, 481, 482, 960, 962} {
		in := make([]int16, n)
		for i := range in {
			in[i] = int16(i * 1999)
		}
		out := make([]int16, 2*(n/3)+2)
		out[0] = 123
		out[len(out)-1] = 456
		Opus_silk_resampler_down2_3(nil, &s, &out[1], unsafe.SliceData(in), int32(n))
		if out[0] != 123 || out[len(out)-1] != 456 {
			t.Fatal("output guards", n)
		}
	}
	var input [2]int16
	Opus_silk_resampler_down2_3(nil, &s, nil, &input[0], 2)
}

// From down_driver_pointers_test.go.
func TestDownDriverPointers(t *testing.T) {
	coefs := Opus_silk_Resampler_1_2_COEFS // Local owner; FCoefs remains zero.
	s := OpusT_silk_resampler_state_struct{FbatchSize: 80, FinvRatio_Q16: 131072, FFIR_Order: 24, FFIR_Fracs: 1}
	for i := range s.FsFIR.Fi32 {
		s.FsFIR.Fi32[i] = int32(i * 71)
	}
	original := s
	Opus_silk_resampler_private_down_FIR(nil, &s, &coefs[0], nil, nil, 0)
	if s != original {
		t.Fatal("empty mutated state")
	}
	// Zero work only consumes the two AR coefficients, not a FIR row.
	arOnly := [2]int16{coefs[0], coefs[1]}
	if silkResamplerAdvanceOutput(nil, 0) != nil || silkResamplerAdvanceOutput(&arOnly[0], 0) != &arOnly[0] || silkResamplerAdvanceOutput(&arOnly[0], 1) != &arOnly[1] {
		t.Fatal("typed output cursor identity/next sample")
	}
	Opus_silk_resampler_private_down_FIR(nil, &s, &arOnly[0], nil, nil, 0)
	if s != original {
		t.Fatal("AR-only empty call mutated state")
	}
	ref := s
	in := make([]int16, 161)
	for i := range in {
		in[i] = int16(i * 537)
	}
	out := make([]int16, 82)
	out[0] = 123
	out[81] = 456
	want := make([]int16, 80)
	Opus_silk_resampler_private_down_FIR(nil, &s, &coefs[0], &out[1], &in[0], 161)
	Opus_silk_resampler_private_down_FIR(nil, &ref, &coefs[0], &want[0], &in[0], 80)
	Opus_silk_resampler_private_down_FIR(nil, &ref, &coefs[0], &want[40], &in[80], 80)
	if s != ref || !slices.Equal(out[1:81], want) || out[0] != 123 || out[81] != 456 {
		t.Fatal("batch/remainder/state/guards")
	}
	exactState := original
	exact := make([]int16, len(want))
	Opus_silk_resampler_private_down_FIR(nil, &exactState, &coefs[0], &exact[0], &in[0], 161)
	if exactState != s || !slices.Equal(exact, want) {
		t.Fatal("exact output extent/state")
	}
	if !slices.Equal(s.FsFIR.Fi32[24:], original.FsFIR.Fi32[24:]) {
		t.Fatal("history tail modified")
	}
}

// From down_interpol_pointers_test.go.
func TestDownInterpolPointers(t *testing.T) {
	for _, order := range []int32{18, 24, 36} {
		if silk_resampler_private_down_FIR_INTERPOL(nil, nil, nil, nil, order, 1, 0, 1) != 0 {
			t.Fatal("empty")
		}
		in := make([]int32, order)
		coefs := make([]int16, order/2)
		out := [3]int16{123, 0, 456}
		in[0] = 2147483647
		in[order-1] = 1
		coefs[0] = 1
		n := silk_resampler_private_down_FIR_INTERPOL(nil, &out[1], &in[0], &coefs[0], order, 1, 65536, 65536)
		want := int16(-512)
		if order == 18 {
			want = 512
		}
		if n != 1 || out != [3]int16{123, want, 456} {
			t.Fatalf("order %d: %v (%d)", order, out, n)
		}
	}
}

// From up2_pointers_test.go.
func TestUp2Pointers(t *testing.T) {
	input := []int16{-32768, 32767, 0, 1, -1, 100, 20000, -30000}
	before := slices.Clone(input)
	initial := [6]int32{12345, -98765, 65432, -11111, 22222, -33333}
	want := []int16{-136, -1139, -5013, -12497, -18418, -11580, 11619, 32718, 25806, -5450, -23750, -7408, 16187, 12092, -6977, -4300}
	wantState := [6]int32{-32098525, 23918734, 23502913, -36297401, 21703549, 88885463}
	for split := 0; split < len(input); split++ {
		state := initial
		out := make([]int16, len(want)+2)
		out[0], out[len(out)-1] = 111, 222
		Opus_silk_resampler_private_up2_HQ(nil, &state, &out[1], &input[0], int32(split))
		Opus_silk_resampler_private_up2_HQ(nil, &state, &out[1+2*split], &input[split], int32(len(input)-split))
		if !slices.Equal(out[1:len(out)-1], want) || state != wantState {
			t.Fatalf("split=%d output=%v state=%v", split, out, state)
		}
		if out[0] != 111 || out[len(out)-1] != 222 || !slices.Equal(input, before) {
			t.Fatal("input or sentinels changed")
		}
	}
	state := initial
	Opus_silk_resampler_private_up2_HQ(nil, &state, nil, nil, 0)
	if state != initial {
		t.Fatal("empty frame changed state")
	}
	// A one-sample expansion reads the input before writing either output.
	overlap := []int16{input[0], 0}
	state = initial
	Opus_silk_resampler_private_up2_HQ(nil, &state, &overlap[0], &overlap[0], 1)
	if !slices.Equal(overlap, want[:2]) {
		t.Fatal("one-sample overlap differs")
	}
}

// From up2_wrapper_pointers_test.go.
func TestUp2WrapperPointers(t *testing.T) {
	coef := int16(123)
	state := OpusT_silk_resampler_state_struct{FsIIR: [6]int32{1, -2, 3, -4, 5, -6}, FCoefs: &coef, FinputDelay: 7}
	state.FdelayBuf[95] = -123
	state.FsFIR.Fi32[35] = 456
	before := state
	Opus_silk_resampler_private_up2_HQ_wrapper(nil, &state, nil, nil, 0)
	if state != before {
		t.Fatal("empty call changed state")
	}
	in := []int16{-32768, 32767, 0, -1, 1, 20000, -20000}
	out := make([]int16, 2*len(in)+2)
	out[0] = 123
	out[len(out)-1] = 456
	want := make([]int16, 2*len(in))
	iir := state.FsIIR
	Opus_silk_resampler_private_up2_HQ(nil, &iir, &want[0], &in[0], int32(len(in)))
	Opus_silk_resampler_private_up2_HQ_wrapper(nil, &state, &out[1], &in[0], int32(len(in)))
	before.FsIIR = iir
	if state != before || !slices.Equal(out[1:len(out)-1], want) || out[0] != 123 || out[len(out)-1] != 456 {
		t.Fatal("wrapper output, state, or guard differs")
	}
}

// From resampler_init_pointers_test.go.
func TestResamplerCoefficientPointers(t *testing.T) {
	makeState := func() *OpusT_silk_resampler_state_struct {
		s := new(OpusT_silk_resampler_state_struct)
		Opus_silk_resampler_init(nil, s, 16000, 8000, 0)
		coefs := slices.Clone(Opus_silk_Resampler_1_2_COEFS[:])
		s.FCoefs = &coefs[0]
		return s
	}
	s := makeState()
	var ref OpusT_silk_resampler_state_struct
	Opus_silk_resampler_init(nil, &ref, 16000, 8000, 0)
	entropyInitGrowStack(12)
	runtime.GC()
	input := [160]int16{}
	for i := range input {
		input[i] = int16(i*791 + 81)
	}
	g, c := [82]int16{}, [82]int16{}
	g[0] = 77
	g[81] = 88
	c = g
	Opus_silk_resampler(nil, s, &g[1], &input[0], 160)
	Opus_silk_resampler(nil, &ref, &c[1], &input[0], 160)
	if g != c {
		t.Fatal("owned table output")
	}
	s.FCoefs = ref.FCoefs
	if *s != ref {
		t.Fatal("owned table state")
	}
	Opus_silk_resampler_init(nil, s, 8000, 16000, 0)
	if s.FCoefs != nil {
		t.Fatal("reinit retained old coefficients")
	}
}

func TestResamplerInitPointers(t *testing.T) {
	for _, enc := range []int32{0, 1} {
		for _, in := range []int32{8000, 12000, 16000, 24000, 48000} {
			for _, out := range []int32{8000, 12000, 16000, 24000, 48000} {
				if enc == 0 && in > 16000 || enc != 0 && out > 16000 {
					continue
				}
				var owner struct {
					before uint32
					state  OpusT_silk_resampler_state_struct
					after  uint32
				}
				owner.before = 123
				owner.after = 456
				// Do not place arbitrary integer patterns in a GC pointer slot.
				coef := int16(12)
				owner.state.FCoefs = &coef
				raw := unsafe.Slice((*byte)(unsafe.Pointer(&owner.state)), int(unsafe.Offsetof(owner.state.FCoefs)))
				for i := range raw {
					raw[i] = 0xa5
				}
				if ret := Opus_silk_resampler_init(nil, &owner.state, in, out, enc); ret != 0 {
					t.Fatal(ret)
				}
				s := owner.state
				if owner.before != 123 || owner.after != 456 {
					t.Fatalf("overwrite %d -> %d enc=%d", in, out, enc)
				}
				if s.FFs_in_kHz != in/1000 || s.FFs_out_kHz != out/1000 || s.FbatchSize != in/100 {
					t.Fatal("rates")
				}
				if s.FsIIR != [6]int32{} || s.FsFIR.Fi32 != [36]int32{} || s.FdelayBuf != [96]int16{} {
					t.Fatal("history not reset")
				}
				if s.FinputDelay < 0 || s.FinputDelay > s.FFs_in_kHz {
					t.Fatal("delay")
				}
			}
		}
	}
}

// From resampler_driver_pointers_test.go.
func TestResamplerPartialConsumedPointers(t *testing.T) {
	for _, rates := range [][2]int32{{8000, 8000}, {8000, 16000}, {12000, 48000}, {16000, 8000}, {48000, 8000}} {
		var state OpusT_silk_resampler_state_struct
		forEncoder := int32(0)
		if rates[0] == 48000 {
			forEncoder = 1
		}
		if Opus_silk_resampler_init(nil, &state, rates[0], rates[1], forEncoder) != 0 {
			t.Fatal("partial init", rates)
		}
		length := state.FFs_in_kHz + 1
		input := make([]int16, length)
		for i := range input {
			input[i] = int16(i*719 + 123)
		}
		count := (length*state.FFs_out_kHz + state.FFs_in_kHz - 1) / state.FFs_in_kHz
		output := make([]int16, count+2)
		output[0], output[len(output)-1] = 12345, -23456
		entropyInitGrowStack(12)
		runtime.GC()
		if Opus_silk_resampler(nil, &state, &output[1], &input[0], length) != 0 || output[0] != 12345 || output[len(output)-1] != -23456 {
			t.Fatal("partial consumed output guards", rates, output)
		}
	}
}

func TestResamplerDriverPointers(t *testing.T) {
	for _, rates := range [][2]int32{{8000, 8000}, {8000, 16000}, {12000, 48000}, {16000, 8000}} {
		var s OpusT_silk_resampler_state_struct
		Opus_silk_resampler_init(nil, &s, rates[0], rates[1], 0)
		for _, ms := range []int32{1, 2, 10, 21} {
			in := make([]int16, ms*s.FFs_in_kHz)
			for i := range in {
				in[i] = int16(i*719 + 123)
			}
			out := make([]int16, ms*s.FFs_out_kHz+2)
			out[0] = 12345
			out[len(out)-1] = -23456
			entropyInitGrowStack(12)
			runtime.GC()
			if ret := Opus_silk_resampler(nil, &s, &out[1], &in[0], int32(len(in))); ret != 0 {
				t.Fatal(ret)
			}
			if out[0] != 12345 || out[len(out)-1] != -23456 {
				t.Fatal("guards", rates, ms)
			}
			for i := int32(0); i < s.FinputDelay; i++ {
				if s.FdelayBuf[i] != in[int32(len(in))-s.FinputDelay+i] {
					t.Fatal("delay history", rates, ms)
				}
			}
		}
	}
}
