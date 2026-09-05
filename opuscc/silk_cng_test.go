package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestCNGUpdatesSmoothedGain(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	dec := OpusT_silk_decoder_state{
		Ffs_kHz:         16,
		Fnb_subfr:       4,
		Fsubfr_length:   2,
		FLPC_order:      10,
		FprevSignalType: TYPE_NO_VOICE_ACTIVITY,
		FprevNLSF_Q15:   [16]OpusT_opus_int16{1300, 2800, 4700, 6900, 9400, 12200, 15300, 18700, 22100, 25700},
		FsCNG: OpusT_silk_CNG_struct{
			Ffs_kHz:            16,
			FCNG_smth_Gain_Q16: 300000,
			FCNG_smth_NLSF_Q15: [16]OpusT_opus_int16{1000, 2500, 4300, 6600, 9000, 12000, 15100, 18500, 22000, 25600},
			FCNG_exc_buf_Q14:   [320]OpusT_opus_int32{7, -11, 19, -23, 29, -31, 37, -41},
		},
	}

	control := OpusT_silk_decoder_control{
		FGains_Q16: [4]OpusT_opus_int32{800000, 500000, 1100000, 400000},
	}
	frame := []int16{23, -17, 41, -53, 67, -71, 89, -97}

	Opus_silk_CNG(
		tls,
		uintptr(unsafe.Pointer(&dec)),
		uintptr(unsafe.Pointer(&control)),
		uintptr(unsafe.Pointer(&frame[0])),
		int32(len(frame)),
	)

	if got, want := dec.FsCNG.FCNG_smth_Gain_Q16, OpusT_opus_int32(400222); got != want {
		t.Fatalf("smoothed gain: got %d, want %d", got, want)
	}
	wantNLSF := [10]OpusT_opus_int16{1074, 2574, 4399, 6674, 9099, 12049, 15149, 18549, 22024, 25624}
	for i, want := range wantNLSF {
		if got := dec.FsCNG.FCNG_smth_NLSF_Q15[i]; got != want {
			t.Fatalf("smoothed NLSF[%d]: got %d, want %d", i, got, want)
		}
	}

	wantExcitation := [8]OpusT_opus_int32{0, 0, 7, -11, 19, -23, 29, -31}
	for i, want := range wantExcitation {
		if got := dec.FsCNG.FCNG_exc_buf_Q14[i]; got != want {
			t.Fatalf("excitation[%d]: got %d, want %d", i, got, want)
		}
	}

	wantFrame := [8]int16{23, -17, 41, -53, 67, -71, 89, -97}
	for i, want := range wantFrame {
		if got := frame[i]; got != want {
			t.Fatalf("frame[%d]: got %d, want %d", i, got, want)
		}
	}

}

func TestCNGResetFieldAccesses(t *testing.T) {
	dec := OpusT_silk_decoder_state{
		FLPC_order: 10,
		FsCNG: OpusT_silk_CNG_struct{
			FCNG_smth_Gain_Q16: 918273,
			Frand_seed:         -123456,
			FCNG_smth_NLSF_Q15: [16]OpusT_opus_int16{11, 22, 33, 44, 55, 66, 77, 88, 99, 111, 122, 133},
		},
	}

	Opus_silk_CNG_Reset(nil, uintptr(unsafe.Pointer(&dec)))

	if dec.FsCNG.FCNG_smth_Gain_Q16 != 0 || dec.FsCNG.Frand_seed != 3176576 {
		t.Fatalf("unexpected reset state: gain=%d seed=%d", dec.FsCNG.FCNG_smth_Gain_Q16, dec.FsCNG.Frand_seed)
	}
	wantNLSF := [10]OpusT_opus_int16{2978, 5956, 8934, 11912, 14890, 17868, 20846, 23824, 26802, 29780}
	for i, want := range wantNLSF {
		if got := dec.FsCNG.FCNG_smth_NLSF_Q15[i]; got != want {
			t.Fatalf("reset NLSF[%d]: got %d, want %d", i, got, want)
		}
	}
	if got, want := dec.FsCNG.FCNG_smth_NLSF_Q15[10], OpusT_opus_int16(122); got != want {
		t.Fatalf("reset NLSF beyond LPC order: got %d, want %d", got, want)
	}
}
