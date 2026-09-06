package opuscc

import (
	"encoding/hex"
	"math"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func mustHex(t *testing.T, s string) []byte {
	t.Helper()
	b, err := hex.DecodeString(s)
	if err != nil {
		t.Fatalf("bad hex literal: %v", err)
	}
	return b
}

func fnv1aFloats(samples []float32) uint32 {
	h := uint32(2166136261)
	var buf [4]byte
	for i := range samples {
		bits := math.Float32bits(samples[i])
		buf[0] = byte(bits)
		buf[1] = byte(bits >> 8)
		buf[2] = byte(bits >> 16)
		buf[3] = byte(bits >> 24)
		for _, b := range buf {
			h ^= uint32(b)
			h *= 16777619
		}
	}
	return h
}

// silkDecodeCall records everything needed to compare one Opus_silk_Decode
// call against the C reference harness (/tmp/opencode/silkdec_ref.c).
type silkDecodeCall struct {
	ret          int32
	nSamplesOut  int32
	prevPitchLag int32
	fnv          uint32
	first8       [8]uint32
}

func runSilkPacket(t *testing.T, tls *libc.TLS, memory uintptr, control *OpusT_silk_DecControlStruct, pcm []float32, pkt []byte) []silkDecodeCall {
	t.Helper()
	payload := pkt[1:]
	var rangeDec OpusT_ec_dec
	Opus_ec_dec_init(tls, uintptr(unsafe.Pointer(&rangeDec)), uintptr(unsafe.Pointer(&payload[0])), uint32(len(payload)))
	var calls []silkDecodeCall
	for f := 0; f < 3; f++ {
		var n int32
		base := f * 960 * int(control.FnChannelsAPI)
		ret := Opus_silk_Decode(tls, memory, uintptr(unsafe.Pointer(control)), 0, libc.BoolInt32(f == 0), uintptr(unsafe.Pointer(&rangeDec)), uintptr(unsafe.Pointer(&pcm[base])), uintptr(unsafe.Pointer(&n)), 0)
		var first8 [8]uint32
		for i := range first8 {
			first8[i] = math.Float32bits(pcm[base+i])
		}
		calls = append(calls, silkDecodeCall{ret, n, control.FprevPitchLag, fnv1aFloats(pcm[base : base+int(n)*int(control.FnChannelsAPI)]), first8})
	}
	return calls
}

func checkSilkCalls(t *testing.T, tag string, got, want []silkDecodeCall) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("%s: got %d calls, want %d", tag, len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("%s call %d: got %+v, want %+v", tag, i, got[i], want[i])
		}
	}
}

// Decodes real SILK packets from opus_newvectors/testvector02.bit (packets 0
// and 1: mono NB 60 ms) through Opus_silk_Decode the way opus_decode_frame
// does, then exercises the packet-loss path, and compares against values
// produced by the C reference implementation.
func TestSilkDecodeRealPacketsMonoAndPLC(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)

	pkt0 := mustHex(t, "18007523c11e84d40a7ed0075134da9ffc0529ef9f410157b57c1f843e40")
	pkt1 := mustHex(t, "182312cf4040d200ea1335b36ad4d1a12853dd70b1861253119131ec38")

	var size int32
	Opus_silk_Get_Decoder_Size(tls, uintptr(unsafe.Pointer(&size)))
	memory := libc.Xmalloc(tls, uint64(size))
	libc.Xmemset(tls, memory, 0, uint64(size))
	if got := Opus_silk_InitDecoder(tls, memory); got != SILK_NO_ERROR {
		t.Fatalf("decoder initialization: got %d", got)
	}
	decoder := (*OpusT_silk_decoder)(unsafe.Pointer(memory))

	var control OpusT_silk_DecControlStruct
	control.FnChannelsAPI = 2
	control.FnChannelsInternal = 1
	control.FAPI_sampleRate = 48000
	control.FinternalSampleRate = 8000
	control.FpayloadSize_ms = 60

	pcm := make([]float32, 5760*2)

	checkSilkCalls(t, "pkt0", runSilkPacket(t, tls, memory, &control, pcm, pkt0), []silkDecodeCall{
		{0, 960, 0, 0x8725c6f5, [8]uint32{}},
		{0, 960, 0, 0x05a82535, [8]uint32{0xb9d00000, 0xb9d00000, 0xb9d00000, 0xb9d00000, 0xb9d00000, 0xb9d00000, 0xb9d00000, 0xb9d00000}},
		{0, 960, 0, 0x3aec7d35, [8]uint32{0xb9d00000, 0xb9d00000, 0xb9d00000, 0xb9d00000, 0xb9c00000, 0xb9c00000, 0xb9c00000, 0xb9c00000}},
	})
	checkSilkCalls(t, "pkt1", runSilkPacket(t, tls, memory, &control, pcm, pkt1), []silkDecodeCall{
		{0, 960, 0, 0x08df9055, [8]uint32{0xb9900000, 0xb9900000, 0xb9900000, 0xb9900000, 0xb9900000, 0xb9900000, 0xb9900000, 0xb9900000}},
		{0, 960, 0, 0x3946b6c5, [8]uint32{0xb9400000, 0xb9400000, 0xb9400000, 0xb9400000, 0xb9200000, 0xb9200000, 0xb9200000, 0xb9200000}},
		{0, 960, 0, 0x0721dd61, [8]uint32{0x39400000, 0x39400000, 0x39400000, 0x39400000, 0x39400000, 0x39400000, 0x39400000, 0x39400000}},
	})

	/* Packet loss concealment of one 20 ms frame, matching opus_decode_frame's
	   PLC call (payloadSize_ms 20, new packet, no range decoder). */
	control.FpayloadSize_ms = 20
	var n int32
	plc := silkDecodeCall{
		ret: Opus_silk_Decode(tls, memory, uintptr(unsafe.Pointer(&control)), 1, 1, 0, uintptr(unsafe.Pointer(&pcm[0])), uintptr(unsafe.Pointer(&n)), 0),
	}
	plc.nSamplesOut = n
	plc.prevPitchLag = control.FprevPitchLag
	plc.fnv = fnv1aFloats(pcm[:int(n)*2])
	for i := range plc.first8 {
		plc.first8[i] = math.Float32bits(pcm[i])
	}
	if want := (silkDecodeCall{0, 960, 0, 0xe4303d55, [8]uint32{0xba000000, 0xba000000, 0xb9c00000, 0xb9c00000, 0xb9900000, 0xb9900000, 0xb9800000, 0xb9800000}}); plc != want {
		t.Fatalf("plc: got %+v, want %+v", plc, want)
	}

	ch := &decoder.Fchannel_state[0]
	if got, want := ch.FnFramesDecoded, int32(1); got != want {
		t.Fatalf("nFramesDecoded: got %d, want %d", got, want)
	}
	if got, want := ch.FnFramesPerPacket, int32(1); got != want {
		t.Fatalf("nFramesPerPacket: got %d, want %d", got, want)
	}
	if got, want := ch.FVAD_flags, [3]int32{0, 0, 1}; got != want {
		t.Fatalf("VAD flags: got %v, want %v", got, want)
	}
	if got, want := ch.FLBRR_flag, int32(0); got != want {
		t.Fatalf("LBRR flag: got %d, want %d", got, want)
	}
	if got, want := ch.FLBRR_flags, [3]int32{0, 0, 0}; got != want {
		t.Fatalf("LBRR flags: got %v, want %v", got, want)
	}
	if got, want := ch.Findices.FsignalType, int8(1); got != want {
		t.Fatalf("signal type: got %d, want %d", got, want)
	}
	if got, want := ch.Findices.FquantOffsetType, int8(0); got != want {
		t.Fatalf("quant offset type: got %d, want %d", got, want)
	}
	if got, want := ch.Findices.FGainsIndices, [4]OpusT_opus_int8{4, 17, 11, 1}; got != want {
		t.Fatalf("gain indices: got %v, want %v", got, want)
	}
	if got, want := ch.Findices.FNLSFInterpCoef_Q2, int8(0); got != want {
		t.Fatalf("NLSF interp coef: got %d, want %d", got, want)
	}
	if got, want := ch.Findices.FSeed, int8(0); got != want {
		t.Fatalf("seed: got %d, want %d", got, want)
	}
	if got, want := ch.FlagPrev, int32(144); got != want {
		t.Fatalf("lagPrev: got %d, want %d", got, want)
	}
	if got, want := ch.FprevSignalType, int32(1); got != want {
		t.Fatalf("prevSignalType: got %d, want %d", got, want)
	}
	if got, want := ch.FLastGainIndex, int8(10); got != want {
		t.Fatalf("LastGainIndex: got %d, want %d", got, want)
	}
	if got, want := ch.Fprev_gain_Q16, OpusT_opus_int32(3080192); got != want {
		t.Fatalf("prev_gain_Q16: got %d, want %d", got, want)
	}
	if got, want := ch.FlossCnt, int32(1); got != want {
		t.Fatalf("lossCnt: got %d, want %d", got, want)
	}
	if got, want := ch.Ffs_kHz, int32(8); got != want {
		t.Fatalf("fs_kHz: got %d, want %d", got, want)
	}
	if got, want := ch.Fframe_length, int32(160); got != want {
		t.Fatalf("frame_length: got %d, want %d", got, want)
	}
	if got, want := ch.FLPC_order, int32(10); got != want {
		t.Fatalf("LPC order: got %d, want %d", got, want)
	}
	if got, want := ch.Ffirst_frame_after_reset, int32(0); got != want {
		t.Fatalf("first_frame_after_reset: got %d, want %d", got, want)
	}
	if got, want := ch.FprevNLSF_Q15, [16]OpusT_opus_int16{2719, 3832, 6061, 12826, 13710, 16302, 18712, 20727, 26430, 30273}; got != want {
		t.Fatalf("prevNLSF_Q15: got %v, want %v", got, want)
	}
	if got, want := ch.Fexc_Q14[:16], []int32{1600, -1600, -1600, -1600, -13504, 1600, -1600, -1600, -1600, -1600, 1600, 1600, -1600, -1600, 1600, -1600}; !equalInt32s(got, want) {
		t.Fatalf("exc_Q14: got %v, want %v", got, want)
	}
	if got, want := ch.FoutBuf[:16], []int16{4, 6, 7, 1, 23, 41, 29, 9, 13, 25, 21, 9, 1, 2, 9, 15}; !equalInt16s(got, want) {
		t.Fatalf("outBuf: got %v, want %v", got, want)
	}
	if got, want := ch.FsLPC_Q14_buf, [16]OpusT_opus_int32{-780, -11752, -10376, -7464, -7868, -12716, -9676, -3736, -3260, -3884, -2332, 1944, 4680, 6808, 3444, 52}; got != want {
		t.Fatalf("sLPC_Q14_buf: got %v, want %v", got, want)
	}
	if got, want := decoder.FnChannelsAPI, int32(2); got != want {
		t.Fatalf("nChannelsAPI: got %d, want %d", got, want)
	}
	if got, want := decoder.FnChannelsInternal, int32(1); got != want {
		t.Fatalf("nChannelsInternal: got %d, want %d", got, want)
	}
	if got, want := decoder.Fprev_decode_only_middle, int32(0); got != want {
		t.Fatalf("prev_decode_only_middle: got %d, want %d", got, want)
	}
	if got, want := decoder.FsStereo, (OpusT_stereo_dec_state{Fpred_prev_Q13: [2]OpusT_opus_int16{0, 0}, FsMid: [2]OpusT_opus_int16{10, 0}, FsSide: [2]OpusT_opus_int16{0, 0}}); got != want {
		t.Fatalf("sStereo: got %+v, want %+v", got, want)
	}
}

// Stereo SILK packets 602/603 of testvector02.bit (stereo NB 60 ms) decoded
// on a fresh decoder, compared against the C reference implementation.
func TestSilkDecodeRealPacketsStereo(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)

	pkt602 := mustHex(t, "1ce27f853110ec825fbabc88e8820cb1a5d6ba25069ec1d8cc14ea25e1a38ccfd4eaafb3"+
		"596c1444c2e02fe8706847d56e57cf605f9732ddaf5efa683bae5ebfd9a8cd3bfa9c7e32"+
		"f0e48418b97fc54503cecaa1cc6b1a46d91123908119b9793b3e6bd408b89c9feab30575"+
		"f971e925bc8d7eb061bb22f4786430e2a148792d3e3097bea8e1186f6b0efdcc107ac6f3"+
		"3d1b0ca884e15c2c3b9f17c11cabb6afd73b717cb2d901532dcb41adeddc6d04352a04a6"+
		"cd4eac275ab18a17b2e1ff1161700e743564dfdbfc58854491b94fa1565204d15f5f4b15"+
		"a39343494601fa5c8c7f9b464063c33266951dc965096200e48d1a119d1946563e3a2804"+
		"09654f1e9941bbf6ec7374d995715c902b914e63262db048048fdc711480660299247c40"+
		"721cfd020a8c9967c7ac6e66ea309821aa31b8fe9de96376c5a58cb0bae1c2f113678600"+
		"412f4c6130d956648101f349d5bbf3a73b6ae43a93b92a5537")
	pkt603 := mustHex(t, "1cc080404d5128c74f619b83306a8e60a6c91ac34020ccd1b90a15517ebda3d901e64e76"+
		"d954e12366b679a968c98e696e105c1703f487adf5cc016da5c644de47e81f16fc0c5ffb"+
		"e3175c169828131108ff493ed1a398fb2d370f09b05f43423db1c055572df9c775aab7b0"+
		"0f553595ddb566d949363df565940ec96cbf2f9a7ab6e67184ab5073809168f903f6f793"+
		"0655cce4c88f0bf69b649bafb600823c6a41c1a486b0727c0894ba3f1b50660762df668e"+
		"f69bf492be64e6cc8fc44a318ab3d87c49f8cbf8690f0feb417fc69b46d47889bc26819b"+
		"1da85ba45a6826650adf772c3956fe0da296ece58954960c16025f3b8fc41dcf625ab839"+
		"92b2823f5a71bdc37af4d69927e74d8b94344abbbbdbecb28230f0a085c0")

	var size int32
	Opus_silk_Get_Decoder_Size(tls, uintptr(unsafe.Pointer(&size)))
	memory := libc.Xmalloc(tls, uint64(size))
	libc.Xmemset(tls, memory, 0, uint64(size))
	if got := Opus_silk_InitDecoder(tls, memory); got != SILK_NO_ERROR {
		t.Fatalf("decoder initialization: got %d", got)
	}
	decoder := (*OpusT_silk_decoder)(unsafe.Pointer(memory))

	var control OpusT_silk_DecControlStruct
	control.FnChannelsAPI = 2
	control.FnChannelsInternal = 2
	control.FAPI_sampleRate = 48000
	control.FinternalSampleRate = 8000
	control.FpayloadSize_ms = 60

	pcm := make([]float32, 5760*2)

	checkSilkCalls(t, "pkt602", runSilkPacket(t, tls, memory, &control, pcm, pkt602), []silkDecodeCall{
		{0, 960, 396, 0x1aa334a4, [8]uint32{}},
		{0, 960, 348, 0xabe77760, [8]uint32{0xbb2e0000, 0xbad40000, 0xbb300000, 0xbad40000, 0xbb300000, 0xbad40000, 0xbb300000, 0xbad40000}},
		{0, 960, 408, 0x632dad28, [8]uint32{0xbc178000, 0xbbe50000, 0xbc088000, 0xbbbe0000, 0xbbec0000, 0xbb920000, 0xbbbf0000, 0xbb420000}},
	})
	checkSilkCalls(t, "pkt603", runSilkPacket(t, tls, memory, &control, pcm, pkt603), []silkDecodeCall{
		{0, 960, 0, 0xf32c7850, [8]uint32{0xbba10000, 0xbb620000, 0xbb930000, 0xbb4c0000, 0xbb830000, 0xbb3a0000, 0xbb660000, 0xbb2a0000}},
		{0, 960, 0, 0x6747419b, [8]uint32{0xba600000, 0xba380000, 0xba500000, 0xba900000, 0xba380000, 0xbac00000, 0xba280000, 0xbae40000}},
		{0, 960, 0, 0xab69c72f, [8]uint32{0xb8c00000, 0xb9800000, 0xb9200000, 0xb9800000, 0xb9400000, 0xb9600000, 0xb9600000, 0xb9600000}},
	})

	ch0 := &decoder.Fchannel_state[0]
	ch1 := &decoder.Fchannel_state[1]
	if got, want := ch0.FnFramesDecoded, int32(3); got != want {
		t.Fatalf("ch0 nFramesDecoded: got %d, want %d", got, want)
	}
	if got, want := ch0.FnFramesPerPacket, int32(3); got != want {
		t.Fatalf("ch0 nFramesPerPacket: got %d, want %d", got, want)
	}
	if got, want := ch0.FVAD_flags, [3]int32{1, 1, 0}; got != want {
		t.Fatalf("ch0 VAD flags: got %v, want %v", got, want)
	}
	if got, want := ch1.FVAD_flags, [3]int32{0, 0, 0}; got != want {
		t.Fatalf("ch1 VAD flags: got %v, want %v", got, want)
	}
	if got, want := ch1.FLBRR_flags, [3]int32{0, 0, 0}; got != want {
		t.Fatalf("ch1 LBRR flags: got %v, want %v", got, want)
	}
	if got, want := ch0.Findices.FsignalType, int8(0); got != want {
		t.Fatalf("ch0 signal type: got %d, want %d", got, want)
	}
	if got, want := ch0.Findices.FquantOffsetType, int8(1); got != want {
		t.Fatalf("ch0 quant offset type: got %d, want %d", got, want)
	}
	if got, want := ch0.Findices.FGainsIndices, [4]OpusT_opus_int8{3, 3, 4, 4}; got != want {
		t.Fatalf("ch0 gain indices: got %v, want %v", got, want)
	}
	if got, want := ch0.Findices.FNLSFInterpCoef_Q2, int8(4); got != want {
		t.Fatalf("ch0 NLSF interp coef: got %d, want %d", got, want)
	}
	if got, want := ch0.Findices.FSeed, int8(1); got != want {
		t.Fatalf("ch0 seed: got %d, want %d", got, want)
	}
	if got, want := ch0.Findices.FlagIndex, OpusT_opus_int16(52); got != want {
		t.Fatalf("ch0 lagIndex: got %d, want %d", got, want)
	}
	if got, want := ch0.Findices.FcontourIndex, int8(8); got != want {
		t.Fatalf("ch0 contour index: got %d, want %d", got, want)
	}
	if got, want := ch0.Findices.FLTPIndex, [4]OpusT_opus_int8{7, 29, 18, 7}; got != want {
		t.Fatalf("ch0 LTP index: got %v, want %v", got, want)
	}
	if got, want := ch0.FLastGainIndex, int8(3); got != want {
		t.Fatalf("ch0 LastGainIndex: got %d, want %d", got, want)
	}
	if got, want := ch0.Fprev_gain_Q16, OpusT_opus_int32(131072); got != want {
		t.Fatalf("ch0 prev_gain_Q16: got %d, want %d", got, want)
	}
	if got, want := ch0.FprevNLSF_Q15, [16]OpusT_opus_int16{1018, 4203, 7265, 10278, 13678, 15280, 18096, 22350, 27673, 30032}; got != want {
		t.Fatalf("ch0 prevNLSF_Q15: got %v, want %v", got, want)
	}
	if got, want := ch1.Findices.FGainsIndices, [4]OpusT_opus_int8{3, 3, 3, 4}; got != want {
		t.Fatalf("ch1 gain indices: got %v, want %v", got, want)
	}
	if got, want := ch1.FprevNLSF_Q15, [16]OpusT_opus_int16{1622, 5810, 7111, 11074, 14047, 14634, 20036, 22390, 26752, 29440}; got != want {
		t.Fatalf("ch1 prevNLSF_Q15: got %v, want %v", got, want)
	}
	if got, want := ch1.FsLPC_Q14_buf, [16]OpusT_opus_int32{2752, -2480, 5296, 12272, 7904, 13280, 10624, 5680, 2592, 11072, 10784, 5024, 400, -2656, 7360, -3008}; got != want {
		t.Fatalf("ch1 sLPC_Q14_buf: got %v, want %v", got, want)
	}
	if got, want := ch0.Fexc_Q14[:16], []int32{-11264, 3840, 35328, -11264, 3840, 3840, 27648, -11264, 18944, -3840, 18944, 3840, -11264, -3840, 3840, 11264}; !equalInt32s(got, want) {
		t.Fatalf("ch0 exc_Q14: got %v, want %v", got, want)
	}
	if got, want := ch1.Fexc_Q14[:16], []int32{-11264, 3840, 18944, 3840, 3840, -3840, -3840, -18944, -11264, -11264, 3840, 3840, -11264, 3840, -3840, 3840}; !equalInt32s(got, want) {
		t.Fatalf("ch1 exc_Q14: got %v, want %v", got, want)
	}
	if got, want := decoder.FnChannelsInternal, int32(2); got != want {
		t.Fatalf("nChannelsInternal: got %d, want %d", got, want)
	}
	if got, want := decoder.Fprev_decode_only_middle, int32(0); got != want {
		t.Fatalf("prev_decode_only_middle: got %d, want %d", got, want)
	}
	if got, want := decoder.FsStereo, (OpusT_stereo_dec_state{Fpred_prev_Q13: [2]OpusT_opus_int16{-1689, 656}, FsMid: [2]OpusT_opus_int16{2, 2}, FsSide: [2]OpusT_opus_int16{1, 0}}); got != want {
		t.Fatalf("sStereo: got %+v, want %+v", got, want)
	}
}

func TestSilkDecodeLostFrameState(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)

	var size int32
	Opus_silk_Get_Decoder_Size(tls, uintptr(unsafe.Pointer(&size)))
	memory := libc.Xmalloc(tls, uint64(size))
	if got := Opus_silk_InitDecoder(tls, memory); got != OPUS_OK {
		t.Fatalf("decoder initialization: got %d", got)
	}
	decoder := (*OpusT_silk_decoder)(unsafe.Pointer(memory))
	decoder.FnChannelsAPI = 1
	decoder.FnChannelsInternal = 1
	decoder.Fchannel_state[0].Ffs_API_hz = 8000
	decoder.Fchannel_state[0].Fnb_subfr = MAX_NB_SUBFR
	if got := Opus_silk_decoder_set_fs(tls, uintptr(unsafe.Pointer(&decoder.Fchannel_state[0])), 8, 8000); got != OPUS_OK {
		t.Fatalf("sample rate setup: got %d", got)
	}
	Opus_silk_PLC_Reset(tls, uintptr(unsafe.Pointer(&decoder.Fchannel_state[0])))
	decoder.Fchannel_state[0].FsPLC.FprevGain_Q16 = [2]OpusT_opus_int32{65536, 65536}
	decoder.Fchannel_state[0].FsPLC.FpitchL_Q8 = 20 << 8
	decoder.Fchannel_state[0].FsPLC.Frand_seed = 12345
	var control OpusT_silk_DecControlStruct
	control.FAPI_sampleRate = 8000
	control.FinternalSampleRate = 8000
	control.FnChannelsAPI = 1
	control.FnChannelsInternal = 1
	output := make([]int16, 320)
	var samples int32
	if got := Opus_silk_Decode(tls, memory, uintptr(unsafe.Pointer(&control)), 1, 1, 0, uintptr(unsafe.Pointer(&output[0])), uintptr(unsafe.Pointer(&samples)), 0); got != OPUS_OK {
		t.Fatalf("decode result: got %d", got)
	}
	if got, want := samples, int32(80); got != want {
		t.Fatalf("sample count: got %d, want %d", got, want)
	}
	if got, want := decoder.Fchannel_state[0].FlossCnt, int32(1); got != want {
		t.Fatalf("loss count: got %d, want %d", got, want)
	}
	if got, want := decoder.Fchannel_state[0].FsPLC.Frand_seed, int32(-1769093111); got != want {
		t.Fatalf("PLC random seed: got %d, want %d", got, want)
	}
}
