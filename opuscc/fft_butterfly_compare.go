//go:build compareopus

package opuscc

func CompareFFTButterfly4(out *OpusT_kiss_fft_cpx, stride uint64, tw *OpusT_kiss_twiddle_cpx, m, N, mm int32) {
	kf_bfly4(nil, out, stride, tw, m, N, mm)
}

func CompareFFTButterfly2(out *OpusT_kiss_fft_cpx, m, N int32) { kf_bfly2(nil, out, m, N) }
