//go:build compareopus

package opuscc

import "unsafe"
import libc "github.com/kazzmir/opus-go/libcshim"

func CompareCleanupRecordLayout() [4]uint64 {
	var state __ptcb
	return [4]uint64{uint64(unsafe.Sizeof(state)), uint64(unsafe.Offsetof(state.F__f)), uint64(unsafe.Offsetof(state.F__x)), uint64(unsafe.Offsetof(state.F__next))}
}

func CompareTimezoneLayout() [3]uint64 {
	var state tm
	return [3]uint64{uint64(unsafe.Sizeof(state)), uint64(unsafe.Offsetof(state.Ftm_gmtoff)), uint64(unsafe.Offsetof(state.Ftm_zone))}
}

func CompareAllocationDriver(tls *libc.TLS, mode *OpusT_OpusCustomMode, a *[7][23]int32, s *[3]int32, cfg *[12]int32, ec *OpusT_ec_ctx) int32 {
	return CompareAllocationDriverAlias(tls, mode, a, s, cfg, ec, 0)
}
func CompareAllocationDriverAlias(tls *libc.TLS, mode *OpusT_OpusCustomMode, a *[7][23]int32, s *[3]int32, cfg *[12]int32, ec *OpusT_ec_ctx, alias int32) (result int32) {
	defer func() {
		if recover() != nil {
			result = -99
		}
	}()
	balance, intensity, dual := &s[0], &s[1], &s[2]
	bits, fine, priority := &a[4][1], &a[5][1], &a[6][1]
	switch alias {
	case 1:
		bits = &a[0][1]
	case 2:
		dual = intensity
	case 3:
		balance = &a[5][cfg[1]]
	case 4:
		priority = fine
	case 5:
		intensity = bits
	case 6:
		bits = &a[3][1]
	case 7:
		fine = &a[3][1]
	}
	return clt_compute_allocation(tls, mode, cfg[0], cfg[1], &a[0][1], &a[3][1], cfg[2], intensity, dual, cfg[3], balance, bits, fine, priority, cfg[7], cfg[8], ec, cfg[9], cfg[10], cfg[11])
}

func CompareAllocationInterp(mode *OpusT_OpusCustomMode, a *[7][23]int32, s *[3]int32, cfg *[12]int32, ec *OpusT_ec_ctx) int32 {
	return CompareAllocationInterpAlias(mode, a, s, cfg, ec, 0)
}
func CompareAllocationInterpAlias(mode *OpusT_OpusCustomMode, a *[7][23]int32, s *[3]int32, cfg *[12]int32, ec *OpusT_ec_ctx, alias int32) (result int32) {
	defer func() {
		if recover() != nil {
			result = -99
		}
	}()
	balance, intensity, dual := &s[0], &s[1], &s[2]
	bits, fine, priority := &a[4][1], &a[5][1], &a[6][1]
	switch alias {
	case 1:
		bits = &a[0][1]
	case 2:
		dual = intensity
	case 3:
		balance = &a[5][cfg[1]]
	case 4:
		priority = fine
	case 5:
		intensity = &a[4][1]
	}
	return interp_bits2pulses(nil, mode, cfg[0], cfg[1], cfg[2], &a[0][1], &a[1][1], &a[2][1], &a[3][1], cfg[3], balance, cfg[4], intensity, cfg[5], dual, cfg[6], bits, fine, priority, cfg[7], cfg[8], ec, cfg[9], cfg[10], cfg[11])
}

func ComparePrefilterFold(tls *libc.TLS, data []byte, N int32) {
	st := (*OpusT_OpusCustomDecoder)(unsafe.Pointer(unsafe.SliceData(data)))
	st.Fmode = &mode48000_960_120
	prefilter_and_fold(tls, st, N)
	st.Fmode = nil
}

func CompareCeltSynthesis(tls *libc.TLS, X, energy, left, right *float32, start, end, C, CC, transient, LM, downsample, silence int32) {
	out := [2]*float32{left, right}
	celt_synthesis(tls, &mode48000_960_120, X, &out[0], energy, start, end, C, CC, transient, LM, downsample, silence, 0)
}

func CompareDeemphasis(tls *libc.TLS, left, right, pcm *float32, N, C, downsample int32, coef *float32, mem *float32, accum int32) {
	input := [2]*float32{left, right}
	deemphasis(tls, &input[0], pcm, N, C, downsample, coef, mem, accum)
}

func CompareProjectionCtl(data []byte, request, value, alias int32) (int32, uint32) {
	st := (*OpusT_OpusProjectionDecoder)(unsafe.Pointer(unsafe.SliceData(data)))
	ms := get_multistream_decoder(nil, st)
	offset := int(uintptr(unsafe.Pointer(ms)) - uintptr(unsafe.Pointer(st)))
	if alias >= 0 {
		alias -= int32(offset)
	}
	r, out := compareMSCtl(data[offset:], request, value, alias, func(_ *OpusT_OpusMSDecoder, request int32, a OpusDecoderCtlArgs) int32 {
		return Opus_opus_projection_decoder_ctl_typed(nil, st, request, a)
	})
	if r == 0 && request == OPUS_MULTISTREAM_GET_DECODER_STATE_REQUEST {
		out += uint32(offset)
	}
	return r, out
}

func CompareMSCtl(data []byte, request, value, alias int32) (int32, uint32) {
	return compareMSCtl(data, request, value, alias, func(st *OpusT_OpusMSDecoder, request int32, a OpusDecoderCtlArgs) int32 {
		return Opus_opus_multistream_decoder_ctl_typed(nil, st, request, a)
	})
}
func compareMSCtl(data []byte, request, value, alias int32, ctl func(*OpusT_OpusMSDecoder, int32, OpusDecoderCtlArgs) int32) (int32, uint32) {
	st := (*OpusT_OpusMSDecoder)(unsafe.Pointer(unsafe.SliceData(data)))
	streams, coupled := st.Flayout.Fnb_streams, st.Flayout.Fnb_coupled_streams
	visit := func(mode *OpusT_OpusCustomMode) {
		offset := int((unsafe.Sizeof(*st) + 7) &^ uintptr(7))
		for i := int32(0); i < streams; i++ {
			dec := (*OpusT_OpusDecoder)(unsafe.Add(unsafe.Pointer(st), offset))
			celt := (*OpusT_OpusCustomDecoder)(unsafe.Add(unsafe.Pointer(dec), dec.Fcelt_dec_offset))
			celt.Fmode = mode
			ch := int32(1)
			if i < coupled {
				ch = 2
			}
			offset += int((uint32(Opus_opus_decoder_get_size(nil, ch)) + 7) &^ uint32(7))
		}
	}
	visit(&mode48000_960_120)
	out := uint32(77)
	var decoder *OpusT_OpusDecoder
	a := OpusDecoderCtlArgs{Value: value}
	if alias != -2 {
		a.Decoder = &decoder
		if alias >= 0 {
			p := unsafe.Add(unsafe.Pointer(st), alias)
			a.I32 = (*int32)(p)
			a.U32 = (*uint32)(p)
		} else {
			a.I32 = (*int32)(unsafe.Pointer(&out))
			a.U32 = &out
		}
	}
	r := ctl(st, request, a)
	if request == OPUS_MULTISTREAM_GET_DECODER_STATE_REQUEST && decoder != nil {
		out = uint32(uintptr(unsafe.Pointer(decoder)) - uintptr(unsafe.Pointer(st)))
	}
	visit(nil)
	return r, out
}

func CompareOpusCtl(data []byte, request, value, alias int32) (int32, uint32) {
	st := (*OpusT_OpusDecoder)(unsafe.Pointer(unsafe.SliceData(data)))
	celt := (*OpusT_OpusCustomDecoder)(unsafe.Add(unsafe.Pointer(st), st.Fcelt_dec_offset))
	celt.Fmode = &mode48000_960_120
	out := uint32(77)
	a := OpusDecoderCtlArgs{Value: value}
	if alias != -2 {
		if alias >= 0 {
			p := unsafe.Add(unsafe.Pointer(st), alias)
			a.I32 = (*int32)(p)
			a.U32 = (*uint32)(p)
		} else {
			a.I32 = (*int32)(unsafe.Pointer(&out))
			a.U32 = &out
		}
	}
	r := Opus_opus_decoder_ctl_typed(nil, st, request, a)
	celt.Fmode = nil
	return r, out
}

func CompareCustomCtl(data []byte, request, value, alias int32) (int32, uint32) {
	st := (*OpusT_OpusCustomDecoder)(unsafe.Pointer(unsafe.SliceData(data)))
	st.Fmode = &mode48000_960_120
	out := uint32(77)
	var mode *OpusT_OpusCustomMode
	a := OpusDecoderCtlArgs{Value: value}
	if alias != -2 {
		if alias >= 0 {
			p := unsafe.Add(unsafe.Pointer(st), alias)
			a.I32 = (*int32)(p)
			a.U32 = (*uint32)(p)
			a.Mode = (**OpusT_OpusCustomMode)(p)
		} else {
			a.I32 = (*int32)(unsafe.Pointer(&out))
			a.U32 = &out
			a.Mode = &mode
		}
	}
	r := Opus_opus_custom_decoder_ctl_typed(nil, st, request, a)
	if request == CELT_GET_MODE_REQUEST {
		out = 0
		if mode == st.Fmode {
			out = 1
		}
	}
	st.Fmode = nil
	return r, out
}

func CompareModePulseRate(m *OpusT_OpusCustomMode, band, LM, bits, pulses int32) (int32, int32) {
	cache := modePulseCache(m, (LM+1)*m.FnbEBands+band)
	return modeBits2Pulses(cache, bits), modePulses2Bits(cache, pulses)
}

func CompareCustomDecoderInit(st *OpusT_OpusCustomDecoder, mode *OpusT_OpusCustomMode, channels int32) int32 {
	return opus_custom_decoder_init(nil, st, mode, channels)
}

func CompareCeltReset(st *OpusT_OpusCustomDecoder) { celt_decoder_reset(nil, st) }

func ComparePLCPitchSearch(left, right *float32, C int32) int32 {
	return celt_plc_pitch_search(nil, nil, left, right, C, 0)
}

func CompareTFDecode(start, end, transient int32, out *int32, LM int32, dec *OpusT_ec_dec) {
	tf_decode(nil, start, end, transient, out, LM, dec)
}

func CompareCustomDecoderSize(mode *OpusT_OpusCustomMode, channels int32) int32 {
	return opus_custom_decoder_get_size(nil, mode, channels)
}

func CompareMSPacketValidation(data *byte, length, streams, Fs int32) int32 {
	return opus_multistream_packet_validate(nil, data, length, streams, Fs)
}

func CompareMSValidation(st *OpusT_OpusMSDecoder) int32 {
	validate_ms_decoder(nil, st)
	return Opus_validate_layout(nil, &st.Flayout)
}

func CompareCeltValidation(st *OpusT_OpusCustomDecoder) (panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	Opus_validate_celt_decoder(nil, st)
	return false
}

func CompareOpusValidation(st *OpusT_OpusDecoder) (panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	validate_opus_decoder(nil, st)
	return false
}
