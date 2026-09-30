//go:build compareopus

package opuscc

func CompareOpusValidation(st *OpusT_OpusDecoder) (panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	validate_opus_decoder(nil, st)
	return false
}
