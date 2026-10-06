package opuscc

import (
	"runtime"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestVADAnalysisTypedOwnersPointers(t *testing.T) {
	state := new(OpusT_silk_encoder_state)
	state.Fframe_length = 160
	Opus_silk_VAD_Init(nil, &state.FsVAD)
	input := make([]int16, 160)
	for i := range input {
		input[i] = int16(i*31 - 2000)
	}
	entropyInitGrowStack(12)
	runtime.GC()
	if silkVADAnalysis(nil, state, &input[0]) != 0 || state.FsVAD.Fcounter != 16 {
		t.Fatal("typed VAD owners")
	}
}

func TestVADAnalysisWholePointers(t *testing.T) {
	for _, N := range []int{80, 120, 160, 240, 320} {
		state := new(OpusT_silk_encoder_state)
		state.Fframe_length = int32(N)
		Opus_silk_VAD_Init(nil, &state.FsVAD)
		for step := 0; step < 4; step++ {
			input := make([]int16, N+2)
			input[0], input[N+1] = 77, 88
			for i := 1; i <= N; i++ {
				input[i] = int16((i*137+step*997)%12000 - 6000)
			}
			entropyInitGrowStack(12)
			runtime.GC()
			if silkVADAnalysis(nil, state, &input[1]) != 0 || input[0] != 77 || input[N+1] != 88 || state.FsVAD.Fcounter != int32(16+step) {
				t.Fatal("whole VAD analysis", N, step)
			}
		}
	}
	state := new(OpusT_silk_encoder_state)
	state.Fframe_length = 321
	Opus_silk_VAD_Init(nil, &state.FsVAD)
	before := state.FsVAD
	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("VAD assertion missing")
			}
		}()
		silkVADAnalysis(nil, state, nil)
	}()
	if state.FsVAD != before {
		t.Fatal("VAD assertion mutated state")
	}
}

func TestVADAnalysisVADStateBase(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)

	encoder := OpusT_silk_encoder_state{
		Fframe_length: 160,
		Ffs_kHz:       8,
		FsVAD: OpusT_silk_VAD_state{
			FNL:              [4]OpusT_opus_int32{900, 1100, 1300, 1500},
			Finv_NL:          [4]OpusT_opus_int32{2386092, 1952257, 1651910, 1431655},
			FNoiseLevelBias:  [4]OpusT_opus_int32{50, 60, 70, 80},
			FNrgRatioSmth_Q8: [4]OpusT_opus_int32{25600, 23000, 21000, 19000},
			Fcounter:         1000,
			FHPstate:         -31,
		},
	}
	input := make([]int16, encoder.Fframe_length)
	for i := range input {
		input[i] = int16((i*137)%4000 - 2000)
	}

	if got := Opus_silk_VAD_GetSA_Q8_c(tls, uintptr(unsafe.Pointer(&encoder)), uintptr(unsafe.Pointer(&input[0]))); got != 0 {
		t.Fatalf("VAD result: got %d, want 0", got)
	}
	if got, want := encoder.Fspeech_activity_Q8, int32(255); got != want {
		t.Fatalf("speech activity: got %d, want %d", got, want)
	}
	if got, want := encoder.Finput_tilt_Q15, int32(27728); got != want {
		t.Fatalf("input tilt: got %d, want %d", got, want)
	}
	if got, want := encoder.FsVAD.FXnrgSubfr, [4]OpusT_opus_int32{39107, 11552, 28728, 22251}; got != want {
		t.Fatalf("subframe energies: got %v, want %v", got, want)
	}
	if got, want := encoder.FsVAD.FNL, [4]OpusT_opus_int32{901, 1102, 1302, 1502}; got != want {
		t.Fatalf("noise levels: got %v, want %v", got, want)
	}
	if got, want := encoder.Finput_quality_bands_Q15, [4]OpusT_opus_int32{23731, 22783, 21835, 21124}; got != want {
		t.Fatalf("input quality: got %v, want %v", got, want)
	}
}
