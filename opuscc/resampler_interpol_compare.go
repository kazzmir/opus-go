//go:build compareopus

package opuscc

func CompareIIRFIRInterpol(out, in *int16, limit, step int32) int32 {
	return silk_resampler_private_IIR_FIR_INTERPOL(nil, out, in, limit, step)
}
