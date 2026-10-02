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
