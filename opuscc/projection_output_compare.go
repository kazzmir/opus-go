//go:build compareopus

package opuscc

func CompareProjectionInt24(dst *int32, ds, dc int32, src *float32, ss, n int32, matrix *OpusT_MappingMatrix) {
	opus_projection_copy_channel_out_int24(nil, dst, ds, dc, src, ss, n, matrix)
}

func CompareProjectionShort(dst *int16, ds, dc int32, src *float32, ss, n int32, matrix *OpusT_MappingMatrix) {
	opus_projection_copy_channel_out_short(nil, dst, ds, dc, src, ss, n, matrix)
}

func CompareProjectionFloat(dst *float32, ds, dc int32, src *float32, ss, n int32, matrix *OpusT_MappingMatrix) {
	opus_projection_copy_channel_out_float(nil, dst, ds, dc, src, ss, n, matrix)
}
