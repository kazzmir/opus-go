package opuscc

import (
	"runtime"
	"slices"
	"testing"
	"unsafe"
)

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
