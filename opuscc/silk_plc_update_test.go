package opuscc

import (
	"testing"
	"unsafe"
)

func TestPLCUpdateFieldAccesses(t *testing.T) {
	dec := OpusT_silk_decoder_state{
		Ffs_kHz:       16,
		Fnb_subfr:     4,
		Fsubfr_length: 20,
		FLPC_order:    10,
		Findices:      OpusT_SideInfoIndices{FsignalType: TYPE_VOICED},
		FsPLC: OpusT_silk_PLC_struct{
			FLTPCoef_Q14: [5]OpusT_opus_int16{99, -88, 77, -66, 55},
			FprevLPC_Q12: [16]OpusT_opus_int16{12, -23, 34, -45, 56, -67, 78, -89, 90, -101},
		},
	}
	control := OpusT_silk_decoder_control{
		FpitchL:        [4]int32{35, 57, 79, 101},
		FGains_Q16:     [4]OpusT_opus_int32{300000, 500000, 700000, 900000},
		FPredCoef_Q12:  [2][16]OpusT_opus_int16{{}, {100, -200, 300, -400, 500, -600, 700, -800, 900, -1000}},
		FLTPCoef_Q14:   [20]OpusT_opus_int16{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1200, 1300, 1400, 1500, 1600, 1700, 1800, 1900, 2000},
		FLTP_scale_Q14: 12345,
	}

	silk_PLC_update(nil, uintptr(unsafe.Pointer(&dec)), uintptr(unsafe.Pointer(&control)))

	if got, want := dec.FprevSignalType, int32(TYPE_VOICED); got != want {
		t.Fatalf("previous signal type: got %d, want %d", got, want)
	}
	if got, want := dec.FsPLC.FpitchL_Q8, int32(25856); got != want {
		t.Fatalf("pitch: got %d, want %d", got, want)
	}
	if got, want := dec.FsPLC.FLTPCoef_Q14, [5]OpusT_opus_int16{0, 0, 11460, 0, 0}; got != want {
		t.Fatalf("LTP coefficients: got %v, want %v", got, want)
	}
	if got, want := dec.FsPLC.FprevLPC_Q12, [16]OpusT_opus_int16{100, -200, 300, -400, 500, -600, 700, -800, 900, -1000}; got != want {
		t.Fatalf("previous LPC coefficients: got %v, want %v", got, want)
	}
	if got, want := dec.FsPLC.FprevLTP_scale_Q14, OpusT_opus_int16(12345); got != want {
		t.Fatalf("previous LTP scale: got %d, want %d", got, want)
	}
	if got, want := dec.FsPLC.FprevGain_Q16, [2]OpusT_opus_int32{700000, 900000}; got != want {
		t.Fatalf("previous gains: got %v, want %v", got, want)
	}
	if got, want := dec.FsPLC.Fsubfr_length, int32(20); got != want {
		t.Fatalf("subframe length: got %d, want %d", got, want)
	}
	if got, want := dec.FsPLC.Fnb_subfr, int32(4); got != want {
		t.Fatalf("subframes: got %d, want %d", got, want)
	}
}
