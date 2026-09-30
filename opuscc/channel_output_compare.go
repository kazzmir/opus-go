//go:build compareopus

package opuscc

func CompareChannelInt24(dst *int32, ds, dc int32, src *float32, ss, n int32) {
	opus_copy_channel_out_int24(nil, dst, ds, dc, src, ss, n)
}

func CompareChannelShort(dst *int16, ds, dc int32, src *float32, ss, n int32) {
	opus_copy_channel_out_short(nil, dst, ds, dc, src, ss, n)
}

func CompareChannelFloat(dst *float32, ds, dc int32, src *float32, ss, n int32) {
	opus_copy_channel_out_float(nil, dst, ds, dc, src, ss, n)
}
