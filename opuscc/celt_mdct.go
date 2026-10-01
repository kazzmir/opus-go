// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

package opuscc

import (
	"math"
	"reflect"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

var _ reflect.Type
var _ unsafe.Pointer

// Keep signed strides and zero-stride store order without integer addresses.
func mdctStridedSlice(p *float32, n, stride int32) ([]float32, int32) {
	last := (n - 1) * stride
	first := min(int32(0), last)
	end := max(int32(0), last)
	base := (*float32)(unsafe.Add(unsafe.Pointer(p), int64(first)*4))
	return unsafe.Slice(base, end-first+1), -first
}

func Opus_clt_mdct_forward_c(tls *libc.TLS, l *OpusT_mdct_lookup, bitrev *int16, twiddles *OpusT_kiss_twiddle_cpx, in, out, window *float32, overlap, shift, stride, arch int32) {
	_ = arch
	st := l.Fkfft[shift]
	scale := st.Fscale
	N := l.Fn
	offset := int32(0)
	for i := int32(0); i < shift; i++ {
		N >>= 1
		offset += N
	}
	N2, N4 := N>>1, N>>2
	trig := unsafe.Slice(l.Ftrig, offset+N2)[offset:]
	input := unsafe.Slice(in, N2+overlap)
	output, outputBase := mdctStridedSlice(out, N2, stride)
	win := unsafe.Slice(window, overlap)
	rev := unsafe.Slice(bitrev, N4)
	f := make([]float32, N2)
	f2 := make([]OpusT_kiss_fft_cpx, N4)
	var i int32
	xp1, xp2 := overlap>>1, N2-1+(overlap>>1)
	wp1, wp2 := overlap>>1, (overlap>>1)-1
	for ; i < (overlap+3)>>2; i++ {
		f[2*i] = float32(input[xp1+N2]*win[wp2]) + float32(input[xp2]*win[wp1])
		f[2*i+1] = float32(input[xp1]*win[wp1]) - float32(input[xp2-N2]*win[wp2])
		xp1 += 2
		xp2 -= 2
		wp1 += 2
		wp2 -= 2
	}
	wp1, wp2 = 0, overlap-1
	for ; i < N4-((overlap+3)>>2); i++ {
		f[2*i] = input[xp2]
		f[2*i+1] = input[xp1]
		xp1 += 2
		xp2 -= 2
	}
	for ; i < N4; i++ {
		f[2*i] = -float32(input[xp1-N2]*win[wp1]) + float32(input[xp2]*win[wp2])
		f[2*i+1] = float32(input[xp1]*win[wp2]) + float32(input[xp2+N2]*win[wp1])
		xp1 += 2
		xp2 -= 2
		wp1 += 2
		wp2 -= 2
	}
	for i = 0; i < N4; i++ {
		re, im := f[2*i], f[2*i+1]
		t0, t1 := trig[i], trig[N4+i]
		yr := float32(re*t0) - float32(im*t1)
		yi := float32(im*t0) + float32(re*t1)
		f2[rev[i]] = OpusT_kiss_fft_cpx{Fr: float32(yr * scale), Fi: float32(yi * scale)}
	}
	Opus_opus_fft_impl(tls, st, twiddles, &f2[0])
	for i = 0; i < N4; i++ {
		t0, t1 := trig[i], trig[N4+i]
		v := f2[i]
		yr := float32(v.Fi*t1) - float32(v.Fr*t0)
		yi := float32(v.Fr*t1) + float32(v.Fi*t0)
		output[outputBase+2*i*stride] = yr
		output[outputBase+(N2-1-2*i)*stride] = yi
	}
}

func Opus_clt_mdct_backward_c(tls *libc.TLS, l *OpusT_mdct_lookup, bitrev *int16, twiddles *OpusT_kiss_twiddle_cpx, in, out, window *float32, overlap, shift, stride, arch int32) {
	_ = arch
	N := l.Fn
	offset := int32(0)
	for i := int32(0); i < shift; i++ {
		N >>= 1
		offset += N
	}
	N2, N4 := N>>1, N>>2
	trig := unsafe.Slice(l.Ftrig, offset+N2)[offset:]
	input, inputBase := mdctStridedSlice(in, N2, stride)
	output := unsafe.Slice(out, N2+(overlap>>1))
	win := unsafe.Slice(window, overlap)
	rev := unsafe.Slice(bitrev, N4)
	base := overlap >> 1
	for i := int32(0); i < N4; i++ {
		x1, x2 := input[inputBase+2*i*stride], input[inputBase+(N2-1-2*i)*stride]
		yr := float32(x2*trig[i]) + float32(x1*trig[N4+i])
		yi := float32(x1*trig[i]) - float32(x2*trig[N4+i])
		r := int32(rev[i])
		output[base+2*r+1] = yr
		output[base+2*r] = yi
	}
	Opus_opus_fft_impl(tls, l.Fkfft[shift], twiddles, (*OpusT_kiss_fft_cpx)(unsafe.Pointer(&output[base])))
	// Capture both ends before storing: odd N4 computes the middle pair twice.
	p0, p1 := base, base+N2-2
	for i := int32(0); i < (N4+1)>>1; i++ {
		re, im := output[p0+1], output[p0]
		t0, t1 := trig[i], trig[N4+i]
		yr := float32(re*t0) + float32(im*t1)
		yi := float32(re*t1) - float32(im*t0)
		re, im = output[p1+1], output[p1]
		output[p0] = yr
		output[p1+1] = yi
		t0, t1 = trig[N4-i-1], trig[N2-i-1]
		yr = float32(re*t0) + float32(im*t1)
		yi = float32(re*t1) - float32(im*t0)
		output[p1] = yr
		output[p0+1] = yi
		p0 += 2
		p1 -= 2
	}
	for i := int32(0); i < overlap/2; i++ {
		j := overlap - 1 - i
		x1, x2 := output[j], output[i]
		output[i] = float32(x2*win[j]) - float32(x1*win[i])
		output[j] = float32(x2*win[i]) + float32(x1*win[j])
	}
}

func mdct_backward_legacy(tls *libc.TLS, l, in, out, window uintptr, overlap, shift, stride, arch int32) {
	lookup := (*OpusT_mdct_lookup)(unsafe.Pointer(l))
	st := lookup.Fkfft[shift]
	Opus_clt_mdct_backward_c(tls, lookup, st.Fbitrev, st.Ftwiddles, (*float32)(unsafe.Pointer(in)), (*float32)(unsafe.Pointer(out)), (*float32)(unsafe.Pointer(window)), overlap, shift, stride, arch)
}

const MINI_MAXFACTORS = 32
const mini_kiss_fft_scalar = "float"

type OpusT_mini_kiss_fft_cpx = struct {
	Fr float32
	Fi float32
}

type OpusT_mini_kiss_fft_cfg = uintptr

type mini_kiss_fft_state = struct {
	Fnfft     int32
	Finverse  int32
	Ffactors  [64]int32
	Ftwiddles [1]OpusT_mini_kiss_fft_cpx
}

/* e.g. an fft of length 128 has 4 factors
as far as kissfft is concerned
4*4*4*2
*/

type OpusT_mini_kiss_fft_state = struct {
	Fnfft     int32
	Finverse  int32
	Ffactors  [64]int32
	Ftwiddles [1]OpusT_mini_kiss_fft_cpx
}

/*
  Explanation of macros dealing with complex math:

   C_MUL(m,a,b)         : m = a*b
   C_FIXDIV( c , div )  : if a fixed point impl., c /= div. noop otherwise
   C_SUB( res, a,b)     : res = a - b
   C_SUBFROM( res , a)  : res -= a
   C_ADDTO( res , a)    : res += a
 * */

func kf_bfly21(tls *libc.TLS, out *OpusT_mini_kiss_fft_cpx, stride OpusT_size_t, twiddles *OpusT_mini_kiss_fft_cpx, m int32) {
	// The native do/while requires positive m.
	if m <= 0 {
		return
	}
	data := unsafe.Slice(out, 2*m)
	tw := unsafe.Slice(twiddles, uint64(m-1)*stride+1)
	for j := int32(0); j < m; j++ {
		t := fftMul(data[j+m], tw[uint64(j)*stride])
		data[j+m] = fftSub(data[j], t)
		data[j] = fftAdd(data[j], t)
	}
}

func kf_bfly41(tls *libc.TLS, out *OpusT_mini_kiss_fft_cpx, stride OpusT_size_t, twiddles *OpusT_mini_kiss_fft_cpx, m OpusT_size_t, inverse int32) {
	if m == 0 {
		return
	}
	data := unsafe.Slice(out, 4*m)
	tw := unsafe.Slice(twiddles, 3*(m-1)*stride+1)
	for j := uint64(0); j < m; j++ {
		s0 := fftMul(data[j+m], tw[j*stride])
		s1 := fftMul(data[j+2*m], tw[2*j*stride])
		s2 := fftMul(data[j+3*m], tw[3*j*stride])
		s5 := fftSub(data[j], s1)
		data[j] = fftAdd(data[j], s1)
		s3 := fftAdd(s0, s2)
		s4 := fftSub(s0, s2)
		data[j+2*m] = fftSub(data[j], s3)
		data[j] = fftAdd(data[j], s3)
		if inverse != 0 {
			data[j+m] = OpusT_mini_kiss_fft_cpx{Fr: s5.Fr - s4.Fi, Fi: s5.Fi + s4.Fr}
			data[j+3*m] = OpusT_mini_kiss_fft_cpx{Fr: s5.Fr + s4.Fi, Fi: s5.Fi - s4.Fr}
		} else {
			data[j+m] = OpusT_mini_kiss_fft_cpx{Fr: s5.Fr + s4.Fi, Fi: s5.Fi - s4.Fr}
			data[j+3*m] = OpusT_mini_kiss_fft_cpx{Fr: s5.Fr - s4.Fi, Fi: s5.Fi + s4.Fr}
		}
	}
}

func kf_bfly31(tls *libc.TLS, out *OpusT_mini_kiss_fft_cpx, stride OpusT_size_t, twiddles *OpusT_mini_kiss_fft_cpx, m OpusT_size_t) {
	if m == 0 {
		return
	}
	data := unsafe.Slice(out, 3*m)
	maxTw := m * stride
	if v := 2 * (m - 1) * stride; v > maxTw {
		maxTw = v
	}
	tw := unsafe.Slice(twiddles, maxTw+1)
	epi3 := tw[m*stride]
	for j := uint64(0); j < m; j++ {
		s1 := fftMul(data[j+m], tw[j*stride])
		s2 := fftMul(data[j+2*m], tw[2*j*stride])
		s3 := fftAdd(s1, s2)
		s0 := fftSub(s1, s2)
		data[j+m] = OpusT_mini_kiss_fft_cpx{Fr: data[j].Fr - float32(s3.Fr*.5), Fi: data[j].Fi - float32(s3.Fi*.5)}
		// Preserve the native scratch-product rounding before later sums.
		s0.Fr = float32(s0.Fr * epi3.Fi)
		s0.Fi = float32(s0.Fi * epi3.Fi)
		data[j] = fftAdd(data[j], s3)
		data[j+2*m] = OpusT_mini_kiss_fft_cpx{Fr: data[j+m].Fr + s0.Fi, Fi: data[j+m].Fi - s0.Fr}
		data[j+m].Fr -= s0.Fi
		data[j+m].Fi += s0.Fr
	}
}

func kf_bfly51(tls *libc.TLS, out *OpusT_mini_kiss_fft_cpx, stride OpusT_size_t, twiddles *OpusT_mini_kiss_fft_cpx, m int32) {
	if m <= 0 {
		return
	}
	data := unsafe.Slice(out, 5*m)
	maxTw := 2 * uint64(m) * stride
	if v := 4 * uint64(m-1) * stride; v > maxTw {
		maxTw = v
	}
	tw := unsafe.Slice(twiddles, maxTw+1)
	ya := tw[uint64(m)*stride]
	yb := tw[2*uint64(m)*stride]
	for u := int32(0); u < m; u++ {
		var s [13]OpusT_mini_kiss_fft_cpx
		s[0] = data[u]
		for k := int32(1); k <= 4; k++ {
			s[k] = fftMul(data[u+k*m], tw[uint64(k*u)*stride])
		}
		s[7] = fftAdd(s[1], s[4])
		s[10] = fftSub(s[1], s[4])
		s[8] = fftAdd(s[2], s[3])
		s[9] = fftSub(s[2], s[3])
		data[u] = fftAdd(data[u], fftAdd(s[7], s[8]))
		// mini_kfft.c uses left-associated sums, unlike kiss_fft.c.
		s[5].Fr = float32(s[0].Fr+float32(s[7].Fr*ya.Fr)) + float32(s[8].Fr*yb.Fr)
		s[5].Fi = float32(s[0].Fi+float32(s[7].Fi*ya.Fr)) + float32(s[8].Fi*yb.Fr)
		s[6].Fr = float32(s[10].Fi*ya.Fi) + float32(s[9].Fi*yb.Fi)
		s[6].Fi = -float32(s[10].Fr*ya.Fi) - float32(s[9].Fr*yb.Fi)
		data[u+m] = fftSub(s[5], s[6])
		data[u+4*m] = fftAdd(s[5], s[6])
		s[11].Fr = float32(s[0].Fr+float32(s[7].Fr*yb.Fr)) + float32(s[8].Fr*ya.Fr)
		s[11].Fi = float32(s[0].Fi+float32(s[7].Fi*yb.Fr)) + float32(s[8].Fi*ya.Fr)
		s[12].Fr = -float32(s[10].Fi*yb.Fi) + float32(s[9].Fi*ya.Fi)
		s[12].Fi = float32(s[10].Fr*yb.Fi) - float32(s[9].Fr*ya.Fi)
		data[u+2*m] = fftAdd(s[11], s[12])
		data[u+3*m] = fftSub(s[11], s[12])
	}
}

// st must belong to its complete variable-sized allocation, including all twiddles.
func kf_work(tls *libc.TLS, out, in *OpusT_mini_kiss_fft_cpx, fstride OpusT_size_t, inStride int32, factors []int32, st *OpusT_mini_kiss_fft_state) {
	p, m := factors[0], factors[1]
	data := unsafe.Slice(out, p*m)
	step := fstride * uint64(uint32(inStride))
	input := unsafe.Slice(in, uint64(p*m-1)*step+1)
	if m == 1 {
		for i := int32(0); i < p; i++ {
			data[i] = input[uint64(i)*step]
		}
	} else {
		for i := int32(0); i < p; i++ {
			kf_work(tls, &data[i*m], &input[uint64(i)*step], fstride*uint64(uint32(p)), inStride, factors[2:], st)
		}
	}
	tw := &st.Ftwiddles[0]
	switch p {
	case 2:
		kf_bfly21(tls, out, fstride, tw, m)
	case 3:
		kf_bfly31(tls, out, fstride, tw, uint64(uint32(m)))
	case 4:
		kf_bfly41(tls, out, fstride, tw, uint64(uint32(m)), st.Finverse)
	case 5:
		kf_bfly51(tls, out, fstride, tw, m)
	default:
		libc.X__assert_fail(tls, __ccgo_ts+5527, __ccgo_ts+5529, 317, uintptr(unsafe.Pointer(&__func__)))
	}
}

var __func__ = [8]int8{'k', 'f', '_', 'w', 'o', 'r', 'k'}

// C documentation
//
//	/*  facbuf is populated by p1,m1,p2,m2, ...
//	    where
//	    p[i] * m[i] = m[i-1]
//	    m0 = n                  */
//
// kf_factor requires positive n, like the C helper, and leaves unused factors untouched.
func kf_factor(tls *libc.TLS, n int32, factors *[2 * MINI_MAXFACTORS]int32) int {
	p := int32(4)
	floorSqrt := math.Floor(math.Sqrt(float64(n)))
	used := 0
	for {
		for n%p != 0 {
			switch p {
			case 4:
				p = 2
			case 2:
				p = 3
			default:
				p += 2
			}
			if float64(p) > floorSqrt {
				p = n
			}
		}
		n /= p
		factors[used] = p
		factors[used+1] = n
		used += 2
		if n <= 1 {
			return used
		}
	}
}

// C documentation
//
//	/*
//	 *
//	 * User-callable function to allocate all necessary storage space for the fft.
//	 *
//	 * The return value is a contiguous block of memory, allocated with malloc.  As such,
//	 * It can be freed with free(), rather than a kiss_fft-specific function.
//	 * */
func Opus_mini_kiss_fft_alloc(tls *libc.TLS, nfft, inverse int32, mem *byte, lenmem *OpusT_size_t) *OpusT_mini_kiss_fft_state {
	needed := uint64(unsafe.Sizeof(OpusT_mini_kiss_fft_state{})) + 8*uint64(uint32(nfft-1))
	var state *OpusT_mini_kiss_fft_state
	if lenmem == nil {
		// The returned interior pointer owns the complete flexible-array allocation.
		backing := make([]uint64, (needed+7)/8)
		state = (*OpusT_mini_kiss_fft_state)(unsafe.Pointer(unsafe.SliceData(backing)))
	} else {
		if mem != nil && *lenmem >= needed {
			state = (*OpusT_mini_kiss_fft_state)(unsafe.Pointer(mem))
		}
		*lenmem = needed
	}
	if state == nil {
		return nil
	}
	state.Fnfft = nfft
	state.Finverse = inverse
	twiddles := unsafe.Slice(&state.Ftwiddles[0], nfft)
	for i := int32(0); i < nfft; i++ {
		phase := float64(float64(-2)*float64(3.141592653589793)) * float64(i) / float64(nfft)
		if inverse != 0 {
			phase *= -1
		}
		twiddles[i].Fr = float32(libc.Xcos(tls, phase))
		twiddles[i].Fi = float32(libc.Xsin(tls, phase))
	}
	kf_factor(tls, nfft, &state.Ffactors)
	return state
}

func Opus_mini_kiss_fft_stride(tls *libc.TLS, st *OpusT_mini_kiss_fft_state, fin, fout *OpusT_mini_kiss_fft_cpx, inStride int32) {
	if fin == fout {
		libc.X__assert_fail(tls, __ccgo_ts+5549, __ccgo_ts+5529, 391, uintptr(unsafe.Pointer(&__func__1)))
	}
	kf_work(tls, fout, fin, 1, inStride, st.Ffactors[:], st)
}

var __func__1 = [21]int8{'m', 'i', 'n', 'i', '_', 'k', 'i', 's', 's', '_', 'f', 'f', 't', '_', 's', 't', 'r', 'i', 'd', 'e'}

func Opus_mini_kiss_fft(tls *libc.TLS, cfg *OpusT_mini_kiss_fft_state, fin, fout *OpusT_mini_kiss_fft_cpx) {
	Opus_mini_kiss_fft_stride(tls, cfg, fin, fout, 1)
}

type OpusT_mini_kiss_fftr_cfg = uintptr

type mini_kiss_fftr_state = OpusT_mini_kiss_fftr_state

type OpusT_mini_kiss_fftr_state = struct {
	Fsubstate       *OpusT_mini_kiss_fft_state
	Ftmpbuf         *OpusT_mini_kiss_fft_cpx
	Fsuper_twiddles *OpusT_mini_kiss_fft_cpx
}

func Opus_mini_kiss_fftr_alloc(tls *libc.TLS, nfft, inverse int32, mem *byte, lenmem *OpusT_size_t) *OpusT_mini_kiss_fftr_state {
	if nfft&1 != 0 {
		libc.X__assert_fail(tls, __ccgo_ts+5561, __ccgo_ts+5529, 416, uintptr(unsafe.Pointer(&__func__2)))
	}
	nfft >>= 1
	var subsize OpusT_size_t
	Opus_mini_kiss_fft_alloc(tls, nfft, inverse, nil, &subsize)
	header := unsafe.Sizeof(OpusT_mini_kiss_fftr_state{})
	needed := uint64(header) + subsize + 8*uint64(uint32(nfft*3/2))
	var st *OpusT_mini_kiss_fftr_state
	if lenmem == nil {
		// Every stored pointer is an interior of this complete owning allocation.
		backing := make([]uint64, (needed+7)/8)
		st = (*OpusT_mini_kiss_fftr_state)(unsafe.Pointer(unsafe.SliceData(backing)))
	} else {
		if *lenmem >= needed {
			st = (*OpusT_mini_kiss_fftr_state)(unsafe.Pointer(mem))
		}
		*lenmem = needed
	}
	if st == nil {
		return nil
	}
	submem := (*byte)(unsafe.Add(unsafe.Pointer(st), header))
	st.Fsubstate = Opus_mini_kiss_fft_alloc(tls, nfft, inverse, submem, &subsize)
	st.Ftmpbuf = (*OpusT_mini_kiss_fft_cpx)(unsafe.Add(unsafe.Pointer(submem), subsize))
	st.Fsuper_twiddles = (*OpusT_mini_kiss_fft_cpx)(unsafe.Add(unsafe.Pointer(st.Ftmpbuf), int(nfft)*8))
	tw := unsafe.Slice(st.Fsuper_twiddles, nfft/2)
	for i := int32(0); i < nfft/2; i++ {
		phase := -float64(3.141592653589793) * (float64(i+1)/float64(nfft) + 0.5)
		if inverse != 0 {
			phase *= -1
		}
		tw[i].Fr = float32(libc.Xcos(tls, phase))
		tw[i].Fi = float32(libc.Xsin(tls, phase))
	}
	return st
}

var __func__2 = [21]int8{'m', 'i', 'n', 'i', '_', 'k', 'i', 's', 's', '_', 'f', 'f', 't', 'r', '_', 'a', 'l', 'l', 'o', 'c'}

func Opus_mini_kiss_fftr(tls *libc.TLS, st *OpusT_mini_kiss_fftr_state, timedata *float32, freqdata *OpusT_mini_kiss_fft_cpx) {
	if st.Fsubstate.Finverse != 0 {
		libc.X__assert_fail(tls, __ccgo_ts+5577, __ccgo_ts+5529, 453, uintptr(unsafe.Pointer(&__func__3)))
	}
	ncfft := st.Fsubstate.Fnfft
	// Pack even/odd real samples into complex input; the complete input stays typed.
	Opus_mini_kiss_fft(tls, st.Fsubstate, (*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(timedata)), st.Ftmpbuf)
	tmp := unsafe.Slice(st.Ftmpbuf, ncfft)
	twiddles := unsafe.Slice(st.Fsuper_twiddles, ncfft/2)
	out := unsafe.Slice(freqdata, ncfft+1)
	tdc := tmp[0]
	out[0].Fr = tdc.Fr + tdc.Fi
	out[ncfft].Fr = tdc.Fr - tdc.Fi
	out[0].Fi = 0
	out[ncfft].Fi = 0
	for k := int32(1); k <= ncfft/2; k++ {
		fpk := tmp[k]
		fpnk := tmp[ncfft-k]
		fpnk.Fi = -fpnk.Fi
		f1 := fftAdd(fpk, fpnk)
		f2 := fftSub(fpk, fpnk)
		tw := fftMul(f2, twiddles[k-1])
		out[k].Fr = float32((f1.Fr + tw.Fr) * 0.5)
		out[k].Fi = float32((f1.Fi + tw.Fi) * 0.5)
		out[ncfft-k].Fr = float32((f1.Fr - tw.Fr) * 0.5)
		out[ncfft-k].Fi = float32((tw.Fi - f1.Fi) * 0.5)
	}
}

var __func__3 = [15]int8{'m', 'i', 'n', 'i', '_', 'k', 'i', 's', 's', '_', 'f', 'f', 't', 'r'}

const LAPLACE_LOG_MINP = 0
const LAPLACE_NMIN = 16

var log2_x_norm_coeff16 = [8]float32{
	0: float32(1),
	1: float32(0.8888888955116272),
	2: float32(0.8),
	3: float32(0.7272727489471436),
	4: float32(0.6666666865348816),
	5: float32(0.6153846383094788),
	6: float32(0.5714285969734192),
	7: float32(0.5333333611488342),
}
var log2_y_norm_coeff16 = [8]float32{
	1: float32(0.1699250042438507),
	2: float32(0.32192808389663696),
	3: float32(0.45943161845207214),
	4: float32(0.5849624872207642),
	5: float32(0.7004396915435791),
	6: float32(0.8073549270629883),
	7: float32(0.9068905711174011),
}

/* The minimum probability of an energy delta (out of 32768). */
/* The minimum number of guaranteed representable energy deltas (in one
   direction). */

// C documentation
//
//	/* When called, decay is positive and at most 11456. */
func ec_laplace_get_freq1(tls *libc.TLS, fs0 uint32, decay int32) (r uint32) {
	var ft uint32
	_ = ft
	ft = uint32(int32(32768)-int32(1)<<int32(LAPLACE_LOG_MINP)*(int32(2)*int32(LAPLACE_NMIN))) - fs0
	return ft * uint32(int32(16384)-decay) >> int32(15)
}

func Opus_ec_laplace_encode(tls *libc.TLS, enc *OpusT_ec_enc, value *int32, fs uint32, decay int32) {
	val := *value
	fl := uint32(0)
	if val != 0 {
		s := int32(0)
		if val < 0 {
			s = -1
		}
		val = (val + s) ^ s
		fl = fs
		fs = ec_laplace_get_freq1(tls, fs, decay)
		i := int32(1)
		for fs > 0 && i < val {
			fs *= 2
			fl += fs + 2*(1<<LAPLACE_LOG_MINP)
			fs = fs * uint32(decay) >> 15
			i++
		}
		if fs == 0 {
			maximum := int32((32768 - fl + (1 << LAPLACE_LOG_MINP) - 1) >> LAPLACE_LOG_MINP)
			maximum = (maximum - s) >> 1
			di := min(val-i, maximum-1)
			fl += uint32((2*di + 1 + s) * (1 << LAPLACE_LOG_MINP))
			fs = min(uint32(1<<LAPLACE_LOG_MINP), 32768-fl)
			// Clip the caller's symbol before updating the entropy context (C order).
			*value = (i + di + s) ^ s
		} else {
			fs += 1 << LAPLACE_LOG_MINP
			fl += fs & uint32(^s)
		}
		if fl+fs > 32768 {
			Opus_celt_fatal(tls, __ccgo_ts+5600, __ccgo_ts+5631, 88)
		}
		if fs == 0 {
			Opus_celt_fatal(tls, __ccgo_ts+5649, __ccgo_ts+5631, 89)
		}
	}
	Opus_ec_encode_bin(tls, enc, fl, fl+fs, 15)
}

func Opus_ec_laplace_decode(tls *libc.TLS, dec *OpusT_ec_dec, fs uint32, decay int32) (r int32) {
	var di, val int32
	var fl, fm, v1 uint32
	_, _, _, _, _ = di, fl, fm, val, v1
	val = 0
	fm = Opus_ec_decode_bin(tls, dec, uint32(15))
	fl = uint32(0)
	if fm >= fs {
		val = val + 1
		fl = fs
		fs = ec_laplace_get_freq1(tls, fs, decay) + uint32(int32(1)<<int32(LAPLACE_LOG_MINP))
		/* Search the decaying part of the PDF.*/
		for fs > uint32(int32(1)<<int32(LAPLACE_LOG_MINP)) && fm >= fl+uint32(2)*fs {
			fs = fs * uint32(2)
			fl = fl + fs
			fs = (fs - uint32(int32(2)*(int32(1)<<int32(LAPLACE_LOG_MINP)))) * uint32(decay) >> int32(15)
			fs = fs + uint32(int32(1)<<int32(LAPLACE_LOG_MINP))
			val = val + 1
		}
		/* Everything beyond that has probability LAPLACE_MINP. */
		if fs <= uint32(int32(1)<<int32(LAPLACE_LOG_MINP)) {
			di = int32((fm - fl) >> (int32(LAPLACE_LOG_MINP) + int32(1)))
			val = val + di
			fl = fl + uint32(int32(2)*di*(int32(1)<<int32(LAPLACE_LOG_MINP)))
		}
		if fm < fl+fs {
			val = -val
		} else {
			fl = fl + fs
		}
	}
	if !(fl < uint32(32768)) {
		Opus_celt_fatal(tls, __ccgo_ts+5672, __ccgo_ts+5631, int32(128))
	}
	if !(fs > uint32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+5649, __ccgo_ts+5631, int32(129))
	}
	if !(fl <= fm) {
		Opus_celt_fatal(tls, __ccgo_ts+5699, __ccgo_ts+5631, int32(130))
	}
	if fl+fs < uint32(int32(32768)) {
		v1 = fl + fs
	} else {
		v1 = uint32(int32(32768))
	}
	if !(fm < v1) {
		Opus_celt_fatal(tls, __ccgo_ts+5724, __ccgo_ts+5631, int32(131))
	}
	if fl+fs < uint32(int32(32768)) {
		v1 = fl + fs
	} else {
		v1 = uint32(int32(32768))
	}
	Opus_ec_dec_update(tls, dec, fl, v1, uint32(32768))
	return val
}

func Opus_ec_laplace_encode_p0(tls *libc.TLS, enc *OpusT_ec_enc, value int32, p0 OpusT_opus_uint16, decay OpusT_opus_uint16) {
	var signICDF [3]uint16
	signICDF[0] = uint16(32768 - int32(p0))
	signICDF[1] = signICDF[0] / 2
	s := int32(0)
	if value > 0 {
		s = 1
	} else if value < 0 {
		s = 2
	}
	Opus_ec_enc_icdf16(tls, enc, s, &signICDF[0], 15)
	if value < 0 {
		value = -value
	}
	if value != 0 {
		var icdf [8]uint16
		icdf[0] = max(uint16(7), decay)
		for i := int32(1); i < 7; i++ {
			icdf[i] = uint16(max(7-i, int32(icdf[i-1])*int32(decay)>>15))
		}
		value--
		for {
			Opus_ec_enc_icdf16(tls, enc, min(value, int32(7)), &icdf[0], 15)
			value -= 7
			if value < 0 {
				break
			}
		}
	}
}

func Opus_ec_laplace_decode_p0(tls *libc.TLS, dec *OpusT_ec_dec, p0 OpusT_opus_uint16, decay OpusT_opus_uint16) (r int32) {
	var i, s, v, value, v1 int32
	var icdf [8]OpusT_opus_uint16
	var sign_icdf [3]OpusT_opus_uint16
	_, _, _, _, _ = i, s, v, value, v1
	sign_icdf[0] = uint16(32768 - int32(p0))
	sign_icdf[1] = uint16(int32(sign_icdf[0]) / 2)
	sign_icdf[2] = 0
	s = Opus_ec_dec_icdf16(tls, dec, &sign_icdf[0], uint32(15))
	if s == int32(2) {
		s = -int32(1)
	}
	if s != 0 {
		if int32(7) > int32(decay) {
			v1 = int32(7)
		} else {
			v1 = int32(decay)
		}
		icdf[0] = uint16(v1)
		i = int32(1)
		for {
			if !(i < int32(7)) {
				break
			}
			if 7-i > int32(icdf[i-1])*int32(decay)>>int32(15) {
				v1 = int32(7) - i
			} else {
				v1 = int32(icdf[i-1]) * int32(decay) >> int32(15)
			}
			icdf[i] = uint16(v1)
			i = i + 1
		}
		icdf[7] = 0
		value = int32(1)
		for cond := true; cond; cond = v == int32(7) {
			v = Opus_ec_dec_icdf16(tls, dec, &icdf[0], uint32(15))
			value = value + v
		}
		return s * value
	} else {
		return 0
	}
	return r
}
