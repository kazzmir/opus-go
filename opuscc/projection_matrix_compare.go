//go:build compareopus

package opuscc

func CompareProjectionMatrix(st *OpusT_OpusProjectionDecoder) *OpusT_MappingMatrix {
	return get_dec_demixing_matrix(nil, st)
}
