package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestPLCResetFieldAccesses(t *testing.T) {
	dec := OpusT_silk_decoder_state{Fframe_length: 319, FsPLC: OpusT_silk_PLC_struct{FpitchL_Q8: -111, FprevGain_Q16: [2]OpusT_opus_int32{222222, -333333}, Fsubfr_length: 47, Fnb_subfr: 4, Ffs_kHz: 16}}
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

func TestPLCConcealFieldAccesses(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)
	decoder := OpusT_silk_decoder_state{Ffs_kHz: 8, Fframe_length: 80, Fsubfr_length: 20, Fnb_subfr: 4, Fltp_mem_length: 200, FLPC_order: 10, FprevSignalType: TYPE_VOICED, FsPLC: OpusT_silk_PLC_struct{FLTPCoef_Q14: [5]OpusT_opus_int16{300, -150, 1200, -100, 75}, FprevLPC_Q12: [16]OpusT_opus_int16{120, -80, 60, -45, 30, -20, 15, -10, 8, -5}, FprevGain_Q16: [2]OpusT_opus_int32{65536, 65536}, FprevLTP_scale_Q14: 13000, FpitchL_Q8: 20 << 8, FrandScale_Q14: 11000, Frand_seed: 12345, Fsubfr_length: 20, Fnb_subfr: 4}}
	for i := range decoder.Fexc_Q14 {
		decoder.Fexc_Q14[i] = int32((i*71)%3000 - 1500)
	}
	for i := range decoder.FoutBuf {
		decoder.FoutBuf[i] = int16((i*37)%1000 - 500)
	}
	var control OpusT_silk_decoder_control
	frame := make([]int16, decoder.Fframe_length)
	silk_PLC_conceal(tls, uintptr(unsafe.Pointer(&decoder)), uintptr(unsafe.Pointer(&control)), uintptr(unsafe.Pointer(&frame[0])), 0)
	if got, want := control.FpitchL, [4]OpusT_opus_int32{21, 21, 21, 21}; got != want {
		t.Fatalf("pitch lags: got %v, want %v", got, want)
	}
	if got, want := decoder.FsPLC.Frand_seed, int32(-1769093111); got != want {
		t.Fatalf("random seed: got %d, want %d", got, want)
	}
	if got, want := frame[:8], []int16{14, 17, 20, 23, 26, 29, 32, 35}; !equalInt16s(got, want) {
		t.Fatalf("frame prefix: got %v, want %v", got, want)
	}
}

func equalInt16s(got, want []int16) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}

func fnv1aInt16s(v []int16) uint32 {
	h := uint32(2166136261)
	for _, x := range v {
		u := uint16(x)
		h ^= uint32(u & 0xff)
		h *= 16777619
		h ^= uint32(u >> 8)
		h *= 16777619
	}
	return h
}

type plcCallState struct {
	lossCnt      int32
	randSeed     int32
	randScale    int16
	pitchL_Q8    int32
	pitchLCtrl   [4]int32
	LTPCoef      [5]int16
	prevLPC      [16]int16
	sLPC_Q14_buf [16]int32
}

func plcConcealState(dec *OpusT_silk_decoder_state, ctrl *OpusT_silk_decoder_control) plcCallState {
	return plcCallState{
		lossCnt:      dec.FlossCnt,
		randSeed:     dec.FsPLC.Frand_seed,
		randScale:    dec.FsPLC.FrandScale_Q14,
		pitchL_Q8:    dec.FsPLC.FpitchL_Q8,
		pitchLCtrl:   ctrl.FpitchL,
		LTPCoef:      dec.FsPLC.FLTPCoef_Q14,
		prevLPC:      dec.FsPLC.FprevLPC_Q12,
		sLPC_Q14_buf: dec.FsLPC_Q14_buf,
	}
}

// silk_PLC concealment compared against the C reference harness
// (/tmp/opencode/plc_ref.c, calls silk_PLC with lost=1 twice per scenario:
// first lost frame and a follow-up concealment). Covers the voiced and
// unvoiced first-frame paths, attenuation tables, pitch drift and state
// carry-over between concealed frames.
func TestPLCConcealCReference(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)

	/* Scenario V: voiced, 16 kHz, 20 ms frame, LPC order 16. */
	decoder := OpusT_silk_decoder_state{
		Ffs_kHz: 16, Fframe_length: 320, Fsubfr_length: 80, Fnb_subfr: 4, Fltp_mem_length: 320,
		FLPC_order: 16, FprevSignalType: TYPE_VOICED,
		FsPLC: OpusT_silk_PLC_struct{
			Ffs_kHz:            16,
			FLTPCoef_Q14:       [5]OpusT_opus_int16{2800, -1400, 900, -600, 300},
			FprevLPC_Q12:       [16]OpusT_opus_int16{1800, -1200, 900, -700, 500, -300, 250, -150, 90, -60, 40, -30, 20, -15, 10, -7},
			FprevGain_Q16:      [2]OpusT_opus_int32{98304, 131072},
			FprevLTP_scale_Q14: 12288,
			FpitchL_Q8:         144 << 8,
			FrandScale_Q14:     11000,
			Frand_seed:         12345,
			Fsubfr_length:      80,
			Fnb_subfr:          4,
		},
	}
	for i := range decoder.Fexc_Q14 {
		decoder.Fexc_Q14[i] = int32((i*71)%6000 - 3000)
	}
	for i := range decoder.FoutBuf {
		decoder.FoutBuf[i] = int16((i*37)%1000 - 500)
	}
	for i := range decoder.FsLPC_Q14_buf {
		decoder.FsLPC_Q14_buf[i] = int32(i*97 - 800)
	}
	var control OpusT_silk_decoder_control
	frame := make([]int16, 320)

	Opus_silk_PLC(tls, uintptr(unsafe.Pointer(&decoder)), uintptr(unsafe.Pointer(&control)), uintptr(unsafe.Pointer(&frame[0])), 1, 0)
	if got, want := fnv1aInt16s(frame), uint32(0xe552e96e); got != want {
		t.Fatalf("V.call1 fnv: got %08x, want %08x", got, want)
	}
	if got, want := frame[:12], []int16{10, 16, 20, 25, 29, 33, 38, 42, 47, 52, 56, 61}; !equalInt16s(got, want) {
		t.Fatalf("V.call1 frame prefix: got %v, want %v", got, want)
	}
	if got, want := plcConcealState(&decoder, &control), (plcCallState{
		lossCnt: 1, randSeed: 30906233, randScale: 8785, pitchL_Q8: 38358,
		pitchLCtrl:   [4]int32{150, 150, 150, 150},
		LTPCoef:      [5]int16{2687, -1347, 863, -579, 287},
		prevLPC:      [16]int16{1782, -1176, 873, -672, 476, -282, 233, -138, 82, -54, 36, -27, 18, -13, 9, -6},
		sLPC_Q14_buf: [16]int32{29796, -28848, 23936, -26316, 12196, -13456, 3632, -2868, -452, -1756, -348, -496, -180, 1736, 1768, 808},
	}); got != want {
		t.Fatalf("V.call1 state: got %+v, want %+v", got, want)
	}

	Opus_silk_PLC(tls, uintptr(unsafe.Pointer(&decoder)), uintptr(unsafe.Pointer(&control)), uintptr(unsafe.Pointer(&frame[0])), 1, 0)
	if got, want := fnv1aInt16s(frame), uint32(0xad44a286); got != want {
		t.Fatalf("V.call2 fnv: got %08x, want %08x", got, want)
	}
	if got, want := frame[:12], []int16{-5, -13, -6, -2, 2, 7, 11, 15, 19, 23, 28, 32}; !equalInt16s(got, want) {
		t.Fatalf("V.call2 frame prefix: got %v, want %v", got, want)
	}
	if got, want := plcConcealState(&decoder, &control), (plcCallState{
		lossCnt: 2, randSeed: 1245822649, randScale: 3596, pitchL_Q8: 39913,
		pitchLCtrl:   [4]int32{156, 156, 156, 156},
		LTPCoef:      [5]int16{2186, -1100, 702, -474, 232},
		prevLPC:      [16]int16{1764, -1153, 847, -646, 453, -266, 217, -127, 75, -49, 32, -24, 16, -11, 8, -5},
		sLPC_Q14_buf: [16]int32{4608, 5276, -564, 2864, 428, 2172, 2876, 2660, 2244, 2192, 2564, 3932, 4084, 3892, 4252, 4756},
	}); got != want {
		t.Fatalf("V.call2 state: got %+v, want %+v", got, want)
	}

	/* Scenario U: unvoiced, 16 kHz, 10 ms frame, LPC order 10; exercises the
	   invGain/down_scale rand_Gain path on the first lost frame. */
	decoder = OpusT_silk_decoder_state{
		Ffs_kHz: 16, Fframe_length: 160, Fsubfr_length: 80, Fnb_subfr: 2, Fltp_mem_length: 160,
		FLPC_order: 10, FprevSignalType: TYPE_UNVOICED,
		FsPLC: OpusT_silk_PLC_struct{
			Ffs_kHz:            16,
			FLTPCoef_Q14:       [5]OpusT_opus_int16{1000, -800, 600, -400, 200},
			FprevLPC_Q12:       [16]OpusT_opus_int16{3600, -2600, 1900, -1400, 1000, -700, 450, -300, 180, -110},
			FprevGain_Q16:      [2]OpusT_opus_int32{45875, 65536},
			FprevLTP_scale_Q14: 8192,
			FpitchL_Q8:         40 << 8,
			FrandScale_Q14:     9000,
			Frand_seed:         987654321,
			Fsubfr_length:      80,
			Fnb_subfr:          2,
		},
	}
	for i := range decoder.Fexc_Q14 {
		decoder.Fexc_Q14[i] = int32((i*71)%6000 - 3000)
	}
	for i := range decoder.FoutBuf {
		decoder.FoutBuf[i] = int16((i*53)%800 - 400)
	}
	for i := range decoder.FsLPC_Q14_buf {
		decoder.FsLPC_Q14_buf[i] = int32(i*131 - 1000)
	}
	control = OpusT_silk_decoder_control{}
	frame = frame[:160]

	Opus_silk_PLC(tls, uintptr(unsafe.Pointer(&decoder)), uintptr(unsafe.Pointer(&control)), uintptr(unsafe.Pointer(&frame[0])), 1, 0)
	if got, want := fnv1aInt16s(frame), uint32(0x5b9922fd); got != want {
		t.Fatalf("U.call1 fnv: got %08x, want %08x", got, want)
	}
	if got, want := frame[:12], []int16{42, -15, 2, -6, -5, -2, -1, 1, 3, 5, 7, 8}; !equalInt16s(got, want) {
		t.Fatalf("U.call1 frame prefix: got %v, want %v", got, want)
	}
	if got, want := plcConcealState(&decoder, &control), (plcCallState{
		lossCnt: 1, randSeed: -1162450863, randScale: 16057, pitchL_Q8: 10445,
		pitchLCtrl:   [4]int32{41, 41, 41, 41},
		LTPCoef:      [5]int16{979, -785, 587, -393, 195},
		prevLPC:      [16]int16{3564, -2548, 1844, -1345, 951, -659, 419, -277, 164, -99},
		sLPC_Q14_buf: [16]int32{-396, 1216, 2740, -308, 500, -1444, 3600, 916, 448, 440, -1312, -2264, -3200, -804, 1884, 692},
	}); got != want {
		t.Fatalf("U.call1 state: got %+v, want %+v", got, want)
	}

	Opus_silk_PLC(tls, uintptr(unsafe.Pointer(&decoder)), uintptr(unsafe.Pointer(&control)), uintptr(unsafe.Pointer(&frame[0])), 1, 0)
	if got, want := fnv1aInt16s(frame), uint32(0xdbae42f4); got != want {
		t.Fatalf("U.call2 fnv: got %08x, want %08x", got, want)
	}
	if got, want := frame[:12], []int16{-40, 6, -20, 2, -6, -4, -2, 0, 2, 4, 5, 7}; !equalInt16s(got, want) {
		t.Fatalf("U.call2 frame prefix: got %v, want %v", got, want)
	}
	if got, want := plcConcealState(&decoder, &control), (plcCallState{
		lossCnt: 2, randSeed: 1636374513, randScale: 13005, pitchL_Q8: 10654,
		pitchLCtrl:   [4]int32{42, 42, 42, 42},
		LTPCoef:      [5]int16{883, -709, 529, -356, 175},
		prevLPC:      [16]int16{3528, -2497, 1789, -1292, 904, -620, 391, -256, 150, -90},
		sLPC_Q14_buf: [16]int32{740, 1016, -96, -648, 1824, 2668, 3744, 236, 1340, -620, 2944, 492, 660, 468, 280, 2908},
	}); got != want {
		t.Fatalf("U.call2 state: got %+v, want %+v", got, want)
	}
}
