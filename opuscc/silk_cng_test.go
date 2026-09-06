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

func TestCNGLossPathFieldAccesses(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	pseudostack := libc.Xmalloc(tls, 16)
	scratch := libc.Xmalloc(tls, GLOBAL_STACK_SIZE)
	*(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(pseudostack)) = OpusT_opus_ccgo_pseudostack_state{
		Fscratch_ptr:  scratch,
		Fglobal_stack: scratch,
	}
	libc.Xpthread_setspecific(tls, 0x6f707573, pseudostack)

	dec := OpusT_silk_decoder_state{
		Ffs_kHz:    16,
		FLPC_order: 10,
		FlossCnt:   1,
		FsPLC: OpusT_silk_PLC_struct{
			FrandScale_Q14: 12000,
			FprevGain_Q16:  [2]OpusT_opus_int32{650000, 900000},
		},
		FsCNG: OpusT_silk_CNG_struct{
			Ffs_kHz:            16,
			FCNG_smth_Gain_Q16: 1100000,
			Frand_seed:         24681357,
			FCNG_smth_NLSF_Q15: [16]OpusT_opus_int16{1800, 4300, 7200, 10500, 13900, 17100, 20100, 22900, 25500, 28000},
			FCNG_exc_buf_Q14:   [320]OpusT_opus_int32{170, -290, 410, -530, 650, -770, 890, -1010},
			FCNG_synth_state:   [16]OpusT_opus_int32{19, -31, 47, -59, 71, -83, 97, -109, 127, -149},
		},
	}
	frame := []int16{150, -230, 310, -390, 470, -550, 630, -710}

	Opus_silk_CNG(tls, uintptr(unsafe.Pointer(&dec)), 0, uintptr(unsafe.Pointer(&frame[0])), int32(len(frame)))

	wantFrame := [8]int16{150, -229, 311, -390, 471, -549, 630, -709}
	for i, want := range wantFrame {
		if got := frame[i]; got != want {
			t.Fatalf("frame[%d]: got %d, want %d", i, got, want)
		}
	}
	if got, want := dec.FsCNG.Frand_seed, OpusT_opus_int32(1318865493); got != want {
		t.Fatalf("random seed: got %d, want %d", got, want)
	}
	wantSynthState := [16]OpusT_opus_int32{127, -149, 0, 0, 0, 0, 0, 0, 890, 1674, 1578, 830, 1450, 1022, 730, 1818}
	if got := dec.FsCNG.FCNG_synth_state; got != wantSynthState {
		t.Fatalf("synthesis state: got %v, want %v", got, wantSynthState)
	}
}

func TestCNGLossPathHighGainLocalArrays(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	pseudostack := libc.Xmalloc(tls, 16)
	scratch := libc.Xmalloc(tls, GLOBAL_STACK_SIZE)
	*(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(pseudostack)) = OpusT_opus_ccgo_pseudostack_state{
		Fscratch_ptr:  scratch,
		Fglobal_stack: scratch,
	}
	libc.Xpthread_setspecific(tls, 0x6f707573, pseudostack)

	/* loss path with a smoothed gain above 1<<23 (takes the high-gain SQRT_APPROX
	   branch, exercising the lz and frac_Q7 locals) and LPC_order 16 (exercising
	   all 16 A_Q12 taps in the synthesis filter) */
	dec := OpusT_silk_decoder_state{
		Ffs_kHz:    16,
		FLPC_order: 16,
		FlossCnt:   1,
		FsPLC: OpusT_silk_PLC_struct{
			FrandScale_Q14: 1000,
			FprevGain_Q16:  [2]OpusT_opus_int32{0, 100000},
		},
		FsCNG: OpusT_silk_CNG_struct{
			Ffs_kHz:            16,
			FCNG_smth_Gain_Q16: 11950000,
			Frand_seed:         24681357,
			FCNG_smth_NLSF_Q15: [16]OpusT_opus_int16{1800, 4300, 7200, 10500, 13900, 17100, 20100, 22900, 25500, 28000, 30000, 31000, 31800, 32300, 32600, 32700},
			FCNG_exc_buf_Q14:   [320]OpusT_opus_int32{170, -290, 410, -530, 650, -770, 890, -1010},
			FCNG_synth_state:   [16]OpusT_opus_int32{19, -31, 47, -59, 71, -83, 97, -109, 127, -149, 151, -163, 167, -179, 181, -191},
		},
	}
	frame := []int16{150, -230, 310, -390, 470, -550, 630, -710}

	Opus_silk_CNG(tls, uintptr(unsafe.Pointer(&dec)), 0, uintptr(unsafe.Pointer(&frame[0])), int32(len(frame)))

	/* expected values from the C reference implementation (silk/CNG.c) */
	wantFrame := [8]int16{162, -257, 358, -464, 585, -725, 882, -1038}
	for i, want := range wantFrame {
		if got := frame[i]; got != want {
			t.Fatalf("frame[%d]: got %d, want %d", i, got, want)
		}
	}
	if got, want := dec.FsCNG.Frand_seed, OpusT_opus_int32(1318865493); got != want {
		t.Fatalf("random seed: got %d, want %d", got, want)
	}
	wantSynthState := [16]OpusT_opus_int32{127, -149, 151, -163, 167, -179, 181, -191, 1114, -2502, 4330, -6738, 10474, -15906, 22938, -29878}
	if got := dec.FsCNG.FCNG_synth_state; got != wantSynthState {
		t.Fatalf("synthesis state: got %v, want %v", got, wantSynthState)
	}
}
