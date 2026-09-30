//go:build compareopus

package opuscc

func CompareMiniFFTButterfly3(out, tw *OpusT_mini_kiss_fft_cpx, stride, m uint64, inverse int32) {
	kf_bfly31(nil, out, stride, tw, m)
}

func CompareMiniFFTButterfly4(out, tw *OpusT_mini_kiss_fft_cpx, stride, m uint64, inverse int32) {
	kf_bfly41(nil, out, stride, tw, m, inverse)
}

func CompareMiniFFTButterfly2(out, tw *OpusT_mini_kiss_fft_cpx, stride, m uint64, inverse int32) {
	kf_bfly21(nil, out, stride, tw, int32(m))
}

func CompareFFTButterfly5(out *OpusT_kiss_fft_cpx, stride uint64, tw *OpusT_kiss_twiddle_cpx, m, N, mm int32) {
	kf_bfly5(nil, out, stride, tw, m, N, mm)
}

func CompareFFTButterfly3(out *OpusT_kiss_fft_cpx, stride uint64, tw *OpusT_kiss_twiddle_cpx, m, N, mm int32) {
	kf_bfly3(nil, out, stride, tw, m, N, mm)
}

func CompareFFTButterfly4(out *OpusT_kiss_fft_cpx, stride uint64, tw *OpusT_kiss_twiddle_cpx, m, N, mm int32) {
	kf_bfly4(nil, out, stride, tw, m, N, mm)
}

func CompareFFTButterfly2(out *OpusT_kiss_fft_cpx, m, N int32) { kf_bfly2(nil, out, m, N) }
