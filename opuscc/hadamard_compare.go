//go:build compareopus

package opuscc

func CompareQuantN1(encode, resynth int32, remaining *int32, ec *OpusT_ec_ctx, x, y, low *float32) uint32 {
	ctx := band_ctx{Fencode: encode, Fresynth: resynth, Fremaining_bits: *remaining}
	mask := quant_band_n1(nil, &ctx, ec, x, y, low)
	*remaining = ctx.Fremaining_bits
	return mask
}

func CompareInterleaveHadamard(x *float32, n0, stride, hadamard int32) {
	interleave_hadamard(nil, x, n0, stride, hadamard)
}

func CompareDeinterleaveHadamard(x *float32, n0, stride, hadamard int32) {
	deinterleave_hadamard(nil, x, n0, stride, hadamard)
}
