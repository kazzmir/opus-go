//go:build compareopus

package opuscc

func CompareDownFIRInterpol(out *int16, in *int32, coefs *int16, order, fracs, limit, step int32) int32 {
	return silk_resampler_private_down_FIR_INTERPOL(nil, out, in, coefs, order, fracs, limit, step)
}

func CompareIIRFIRInterpol(out, in *int16, limit, step int32) int32 {
	return silk_resampler_private_IIR_FIR_INTERPOL(nil, out, in, limit, step)
}
