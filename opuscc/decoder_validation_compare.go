//go:build compareopus

package opuscc

func CompareCustomDecoderSize(mode *OpusT_OpusCustomMode, channels int32) int32 {
	return opus_custom_decoder_get_size(nil, mode, channels)
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
