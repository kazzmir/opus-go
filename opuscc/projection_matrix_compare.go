//go:build compareopus

package opuscc

func CompareProjectionMultistream(st *OpusT_OpusProjectionDecoder) *OpusT_OpusMSDecoder {
	return get_multistream_decoder(nil, st)
}

func CompareProjectionMatrix(st *OpusT_OpusProjectionDecoder) *OpusT_MappingMatrix {
	return get_dec_demixing_matrix(nil, st)
}
