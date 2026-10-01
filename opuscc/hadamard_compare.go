//go:build compareopus

package opuscc

func CompareQuantN1(encode, resynth int32, remaining *int32, ec *OpusT_ec_ctx, x, y, low *float32) uint32 {
	ctx := band_ctx{Fencode: encode, Fresynth: resynth, Fremaining_bits: *remaining}
	mask := quant_band_n1(nil, &ctx, ec, x, y, low)
	*remaining = ctx.Fremaining_bits
	return mask
}

func CompareCoarseEnergyImpl(nb, start, end int32, energy, old *float32, budget, tell int32, errors *float32, enc *OpusT_ec_enc, C, LM, intra int32, maxDecay float32, lfe int32) int32 {
	return quant_coarse_energy_impl(nil, nb, start, end, energy, old, budget, tell, &e_prob_model[LM][intra][0], errors, enc, C, LM, intra, maxDecay, lfe)
}

func CompareInterleaveHadamard(x *float32, n0, stride, hadamard int32) {
	interleave_hadamard(nil, x, n0, stride, hadamard)
}

func CompareDeinterleaveHadamard(x *float32, n0, stride, hadamard int32) {
	deinterleave_hadamard(nil, x, n0, stride, hadamard)
}
