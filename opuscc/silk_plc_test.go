package opuscc

import (
	"testing"
	"unsafe"
)

func TestPLCResetFieldAccesses(t *testing.T) {
	dec := OpusT_silk_decoder_state{
		Fframe_length: 319,
		FsPLC: OpusT_silk_PLC_struct{
			FpitchL_Q8:    -111,
			FprevGain_Q16: [2]OpusT_opus_int32{222222, -333333},
			Fsubfr_length: 47,
			Fnb_subfr:     4,
			Ffs_kHz:       16,
		},
	}

	Opus_silk_PLC_Reset(nil, uintptr(unsafe.Pointer(&dec)))

	if got, want := dec.FsPLC.FpitchL_Q8, int32(40832); got != want {
		t.Fatalf("pitchL_Q8: got %d, want %d", got, want)
	}
	if got, want := dec.FsPLC.FprevGain_Q16, [2]OpusT_opus_int32{65536, 65536}; got != want {
		t.Fatalf("previous gains: got %v, want %v", got, want)
	}
	if got, want := dec.FsPLC.Fsubfr_length, int32(20); got != want {
		t.Fatalf("subframe length: got %d, want %d", got, want)
	}
	if got, want := dec.FsPLC.Fnb_subfr, int32(2); got != want {
		t.Fatalf("subframes: got %d, want %d", got, want)
	}
	if got, want := dec.FsPLC.Ffs_kHz, int32(16); got != want {
		t.Fatalf("unrelated PLC state changed: got %d, want %d", got, want)
	}
}
