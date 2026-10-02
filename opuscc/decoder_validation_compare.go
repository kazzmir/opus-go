//go:build compareopus

package opuscc

import "unsafe"

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
