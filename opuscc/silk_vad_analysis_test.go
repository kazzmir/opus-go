package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

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
}
