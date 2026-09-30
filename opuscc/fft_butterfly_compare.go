//go:build compareopus

package opuscc

func CompareFFTButterfly2(out *OpusT_kiss_fft_cpx, m, N int32) { kf_bfly2(nil, out, m, N) }
