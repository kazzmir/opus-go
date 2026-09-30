//go:build compareopus

package opuscc

func CompareMiniFFT(st *OpusT_mini_kiss_fft_state, in, out *OpusT_mini_kiss_fft_cpx, stride int32) {
	Opus_mini_kiss_fft(nil, st, in, out)
}

func CompareMiniFFTStride(st *OpusT_mini_kiss_fft_state, in, out *OpusT_mini_kiss_fft_cpx, stride int32) {
	Opus_mini_kiss_fft_stride(nil, st, in, out, stride)
}

func CompareMiniFFTWork(st *OpusT_mini_kiss_fft_state, in, out *OpusT_mini_kiss_fft_cpx, stride int32) {
	kf_work(nil, out, in, 1, stride, st.Ffactors[:], st)
}

func CompareMiniFFTButterfly5(out, tw *OpusT_mini_kiss_fft_cpx, stride, m uint64, inverse int32) {
	kf_bfly51(nil, out, stride, tw, int32(m))
}

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
