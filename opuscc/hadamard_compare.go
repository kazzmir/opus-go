//go:build compareopus

package opuscc

import "unsafe"

func CompareBandContextLayout() [4]uint64 {
	var ctx band_ctx
	return [4]uint64{uint64(unsafe.Sizeof(ctx)), uint64(unsafe.Offsetof(ctx.Fm)), uint64(unsafe.Offsetof(ctx.Fec)), uint64(unsafe.Offsetof(ctx.FbandE))}
}

func CompareStereoBand(ec *OpusT_ec_ctx, x, y, low *float32, cfg [9]int32) (uint32, [2]uint32) {
	ctx := band_ctx{Fm: &mode48000_960_120, Fec: ec, Fresynth: cfg[4], Ftf_change: cfg[5], Fremaining_bits: cfg[6], Fseed: 123456, Fintensity: cfg[8]}
	mask := quant_band_stereo(nil, &ctx, uintptr(unsafe.Pointer(x)), uintptr(unsafe.Pointer(y)), cfg[0], cfg[3], cfg[1], 0, cfg[2], uintptr(unsafe.Pointer(low)), 0, cfg[7])
	return mask, [2]uint32{uint32(ctx.Fremaining_bits), ctx.Fseed}
}

func CompareMonoBand(ec *OpusT_ec_ctx, x, low *float32, cfg [8]int32) (uint32, [2]uint32) {
	ctx := band_ctx{Fm: &mode48000_960_120, Fec: ec, Fresynth: cfg[4], Ftf_change: cfg[5], Fremaining_bits: cfg[6], Fseed: 123456}
	mask := quant_band(nil, &ctx, uintptr(unsafe.Pointer(x)), cfg[0], cfg[3], cfg[1], 0, cfg[2], uintptr(unsafe.Pointer(low)), 1, 0, cfg[7])
	return mask, [2]uint32{uint32(ctx.Fremaining_bits), ctx.Fseed}
}

func CompareTheta(ec *OpusT_ec_ctx, cfg [12]int32) [9]int32 {
	log := [1]int16{int16(cfg[8])}
	m := OpusT_OpusCustomMode{FnbEBands: 1, FlogN: &log[0]}
	ctx := band_ctx{Fm: &m, Fec: ec, Fintensity: cfg[7], Fremaining_bits: cfg[9], Fdisable_inv: cfg[10]}
	var split split_ctx
	b, fill := cfg[5], cfg[6]
	bp, fp := &b, &fill
	switch cfg[11] {
	case 1:
		fp = bp
	case 2:
		ctx.Fremaining_bits = b
		bp = &ctx.Fremaining_bits
	case 3:
		split.Fitheta = b
		split.Fqalloc = fill
		bp = &split.Fitheta
		fp = &split.Fqalloc
	}
	compute_theta(nil, &ctx, &split, 0, 0, cfg[0], bp, cfg[1], cfg[2], cfg[3], cfg[4], fp)
	return [9]int32{split.Finv, split.Fimid, split.Fiside, split.Fdelta, split.Fitheta, split.Fqalloc, *bp, *fp, ctx.Fremaining_bits}
}

func CompareQuantN1(encode, resynth int32, remaining *int32, ec *OpusT_ec_ctx, x, y, low *float32) uint32 {
	ctx := band_ctx{Fencode: encode, Fresynth: resynth, Fremaining_bits: *remaining, Fec: ec}
	mask := quant_band_n1(nil, &ctx, ctx.Fec, x, y, low)
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
