//go:build compareopus

package opuscc

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
