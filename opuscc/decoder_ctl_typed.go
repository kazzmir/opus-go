package opuscc

import (
	libc "github.com/kazzmir/opus-go/libcshim"
	"unsafe"
)

// OpusDecoderCtlArgs carries the single request's value or typed output slot.
// Mode and Decoder outputs use pointer stores (and therefore Go write barriers).
// Unused fields are ignored, just as unused C varargs are not read.
type OpusDecoderCtlArgs struct {
	Value   int32
	I32     *int32
	U32     *uint32
	Mode    **OpusT_OpusCustomMode
	Decoder **OpusT_OpusDecoder
}

func Opus_opus_projection_decoder_ctl_typed(tls *libc.TLS, st *OpusT_OpusProjectionDecoder, request int32, a OpusDecoderCtlArgs) int32 {
	return Opus_opus_multistream_decoder_ctl_typed(tls, get_multistream_decoder(tls, st), request, a)
}

func msCtlLegacyArgs(st *OpusT_OpusMSDecoder, request int32, ap uintptr) (a OpusDecoderCtlArgs) {
	switch request {
	case OPUS_GET_BANDWIDTH_REQUEST, OPUS_GET_SAMPLE_RATE_REQUEST, OPUS_GET_GAIN_REQUEST, OPUS_GET_LAST_PACKET_DURATION_REQUEST, OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST, OPUS_GET_COMPLEXITY_REQUEST, OPUS_GET_FINAL_RANGE_REQUEST, OPUS_SET_GAIN_REQUEST, OPUS_SET_COMPLEXITY_REQUEST, OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST:
		return opusCtlLegacyArgs(request, ap)
	case OPUS_MULTISTREAM_GET_DECODER_STATE_REQUEST:
		a.Value = libc.VaInt32(&ap)
		// Do not read the second vararg for an invalid stream ID.
		if a.Value >= 0 && a.Value < st.Flayout.Fnb_streams {
			a.Decoder = (**OpusT_OpusDecoder)(unsafe.Pointer(libc.VaUintptr(&ap)))
		}
	}
	return
}

func Opus_opus_multistream_decoder_ctl_typed(tls *libc.TLS, st *OpusT_OpusMSDecoder, request int32, a OpusDecoderCtlArgs) int32 {
	coupled := uint((uint32(Opus_opus_decoder_get_size(tls, 2)) + 7) &^ uint32(7))
	mono := uint((uint32(Opus_opus_decoder_get_size(tls, 1)) + 7) &^ uint32(7))
	offset := (uint(unsafe.Sizeof(*st)) + 7) &^ uint(7)
	next := func(s int32) {
		if s < st.Flayout.Fnb_coupled_streams {
			offset += coupled
		} else {
			offset += mono
		}
	}
	switch request {
	case OPUS_GET_BANDWIDTH_REQUEST, OPUS_GET_SAMPLE_RATE_REQUEST, OPUS_GET_GAIN_REQUEST, OPUS_GET_LAST_PACKET_DURATION_REQUEST, OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST, OPUS_GET_COMPLEXITY_REQUEST:
		return Opus_opus_decoder_ctl_typed(tls, opusMSDecoderAt(st, offset), request, a)
	case OPUS_GET_FINAL_RANGE_REQUEST:
		if a.U32 == nil {
			return -1
		}
		*a.U32 = 0
		for s := int32(0); s < st.Flayout.Fnb_streams; s++ {
			dec := opusMSDecoderAt(st, offset)
			next(s)
			var tmp uint32
			if ret := Opus_opus_decoder_ctl_typed(tls, dec, request, OpusDecoderCtlArgs{U32: &tmp}); ret != OPUS_OK {
				return ret
			}
			*a.U32 ^= tmp
		}
	case OPUS_RESET_STATE, OPUS_SET_GAIN_REQUEST, OPUS_SET_COMPLEXITY_REQUEST, OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST:
		for s := int32(0); s < st.Flayout.Fnb_streams; s++ {
			dec := opusMSDecoderAt(st, offset)
			next(s)
			if ret := Opus_opus_decoder_ctl_typed(tls, dec, request, a); ret != OPUS_OK {
				return ret
			}
		}
	case OPUS_MULTISTREAM_GET_DECODER_STATE_REQUEST:
		if a.Value < 0 || a.Value >= st.Flayout.Fnb_streams {
			return -1
		}
		if a.Decoder == nil {
			return -1
		}
		for s := int32(0); s < a.Value; s++ {
			next(s)
		}
		*a.Decoder = opusMSDecoderAt(st, offset)
	default:
		return -5
	}
	// Offsets advance before child calls, but no unused one-past pointer is made.
	return OPUS_OK
}

func opusCtlLegacyArgs(request int32, ap uintptr) (a OpusDecoderCtlArgs) {
	switch request {
	case OPUS_SET_COMPLEXITY_REQUEST, OPUS_SET_GAIN_REQUEST, OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST, OPUS_SET_IGNORE_EXTENSIONS_REQUEST:
		a.Value = libc.VaInt32(&ap)
	case OPUS_GET_BANDWIDTH_REQUEST, OPUS_GET_COMPLEXITY_REQUEST, OPUS_GET_SAMPLE_RATE_REQUEST, OPUS_GET_PITCH_REQUEST, OPUS_GET_GAIN_REQUEST, OPUS_GET_LAST_PACKET_DURATION_REQUEST, OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST, OPUS_GET_IGNORE_EXTENSIONS_REQUEST:
		a.I32 = (*int32)(unsafe.Pointer(libc.VaUintptr(&ap)))
	case OPUS_GET_FINAL_RANGE_REQUEST:
		a.U32 = (*uint32)(unsafe.Pointer(libc.VaUintptr(&ap)))
	}
	return
}

func Opus_opus_decoder_ctl_typed(tls *libc.TLS, st *OpusT_OpusDecoder, request int32, a OpusDecoderCtlArgs) int32 {
	// C derives these interiors before dispatch, even for an unknown request.
	silk := (*OpusT_silk_decoder)(unsafe.Add(unsafe.Pointer(st), st.Fsilk_dec_offset))
	celt := (*OpusT_OpusCustomDecoder)(unsafe.Add(unsafe.Pointer(st), st.Fcelt_dec_offset))
	switch request {
	case OPUS_GET_BANDWIDTH_REQUEST:
		if a.I32 == nil {
			return -1
		}
		*a.I32 = st.Fbandwidth
	case OPUS_SET_COMPLEXITY_REQUEST:
		if a.Value < 0 || a.Value > 10 {
			return -1
		}
		st.Fcomplexity = a.Value
		// The C caller intentionally ignores the child return value here.
		Opus_opus_custom_decoder_ctl_typed(tls, celt, request, a)
	case OPUS_GET_COMPLEXITY_REQUEST:
		if a.I32 == nil {
			return -1
		}
		*a.I32 = st.Fcomplexity
	case OPUS_GET_FINAL_RANGE_REQUEST:
		if a.U32 == nil {
			return -1
		}
		*a.U32 = st.FrangeFinal
	case OPUS_RESET_STATE:
		clear(unsafe.Slice((*byte)(unsafe.Pointer(&st.Fstream_channels)), int(unsafe.Sizeof(*st)-unsafe.Offsetof(st.Fstream_channels))))
		Opus_opus_custom_decoder_ctl_typed(tls, celt, request, a)
		Opus_silk_ResetDecoder(tls, silk)
		st.Fstream_channels = st.Fchannels
		st.Fframe_size = st.FFs / 400
	case OPUS_GET_SAMPLE_RATE_REQUEST:
		if a.I32 == nil {
			return -1
		}
		*a.I32 = st.FFs
	case OPUS_GET_PITCH_REQUEST:
		if a.I32 == nil {
			return -1
		}
		if st.Fprev_mode == MODE_CELT_ONLY {
			return Opus_opus_custom_decoder_ctl_typed(tls, celt, request, a)
		}
		*a.I32 = st.FDecControl.FprevPitchLag
	case OPUS_GET_GAIN_REQUEST:
		if a.I32 == nil {
			return -1
		}
		*a.I32 = st.Fdecode_gain
	case OPUS_SET_GAIN_REQUEST:
		if a.Value < -32768 || a.Value > 32767 {
			return -1
		}
		st.Fdecode_gain = a.Value
	case OPUS_GET_LAST_PACKET_DURATION_REQUEST:
		if a.I32 == nil {
			return -1
		}
		*a.I32 = st.Flast_packet_duration
	case OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST:
		if a.Value < 0 || a.Value > 1 {
			return -1
		}
		return Opus_opus_custom_decoder_ctl_typed(tls, celt, request, a)
	case OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST:
		if a.I32 == nil {
			return -1
		}
		return Opus_opus_custom_decoder_ctl_typed(tls, celt, request, a)
	case OPUS_SET_IGNORE_EXTENSIONS_REQUEST:
		if a.Value < 0 || a.Value > 1 {
			return -1
		}
		st.Fignore_extensions = a.Value
	case OPUS_GET_IGNORE_EXTENSIONS_REQUEST:
		if a.I32 == nil {
			return -1
		}
		*a.I32 = st.Fignore_extensions
	default:
		return -5
	}
	return OPUS_OK
}

func customCtlLegacyArgs(request int32, ap uintptr) (a OpusDecoderCtlArgs) {
	switch request {
	case OPUS_SET_COMPLEXITY_REQUEST, CELT_SET_START_BAND_REQUEST, CELT_SET_END_BAND_REQUEST, CELT_SET_CHANNELS_REQUEST, CELT_SET_SIGNALLING_REQUEST, OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST:
		a.Value = libc.VaInt32(&ap)
	case OPUS_GET_COMPLEXITY_REQUEST, CELT_GET_AND_CLEAR_ERROR_REQUEST, OPUS_GET_LOOKAHEAD_REQUEST, OPUS_GET_PITCH_REQUEST, OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST:
		a.I32 = (*int32)(unsafe.Pointer(libc.VaUintptr(&ap)))
	case OPUS_GET_FINAL_RANGE_REQUEST:
		a.U32 = (*uint32)(unsafe.Pointer(libc.VaUintptr(&ap)))
	case CELT_GET_MODE_REQUEST:
		a.Mode = (**OpusT_OpusCustomMode)(unsafe.Pointer(libc.VaUintptr(&ap)))
	}
	return
}

func Opus_opus_custom_decoder_ctl_typed(tls *libc.TLS, st *OpusT_OpusCustomDecoder, request int32, a OpusDecoderCtlArgs) int32 {
	switch request {
	case OPUS_SET_COMPLEXITY_REQUEST:
		if a.Value < 0 || a.Value > 10 {
			return -1
		}
		st.Fcomplexity = a.Value
	case OPUS_GET_COMPLEXITY_REQUEST:
		if a.I32 == nil {
			return -1
		}
		*a.I32 = st.Fcomplexity
	case CELT_SET_START_BAND_REQUEST:
		if a.Value < 0 || a.Value >= st.Fmode.FnbEBands {
			return -1
		}
		st.Fstart = a.Value
	case CELT_SET_END_BAND_REQUEST:
		if a.Value < 1 || a.Value > st.Fmode.FnbEBands {
			return -1
		}
		st.Fend = a.Value
	case CELT_SET_CHANNELS_REQUEST:
		if a.Value < 1 || a.Value > 2 {
			return -1
		}
		st.Fstream_channels = a.Value
	case CELT_GET_AND_CLEAR_ERROR_REQUEST:
		if a.I32 == nil {
			return -1
		}
		*a.I32 = st.Ferror1
		st.Ferror1 = 0
	case OPUS_GET_LOOKAHEAD_REQUEST:
		if a.I32 == nil {
			return -1
		}
		*a.I32 = st.Foverlap / st.Fdownsample
	case OPUS_RESET_STATE:
		celt_decoder_reset(tls, st)
	case OPUS_GET_PITCH_REQUEST:
		if a.I32 == nil {
			return -1
		}
		*a.I32 = st.Fpostfilter_period
	case CELT_GET_MODE_REQUEST:
		if a.Mode == nil {
			return -1
		}
		*a.Mode = st.Fmode
	case CELT_SET_SIGNALLING_REQUEST:
		st.Fsignalling = a.Value
	case OPUS_GET_FINAL_RANGE_REQUEST:
		if a.U32 == nil {
			return -1
		}
		*a.U32 = st.Frng
	case OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST:
		if a.Value < 0 || a.Value > 1 {
			return -1
		}
		st.Fdisable_inv = a.Value
	case OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST:
		if a.I32 == nil {
			return -1
		}
		*a.I32 = st.Fdisable_inv
	default:
		return -5
	}
	return OPUS_OK
}
