// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

//go:build linux && (amd64 || arm64)

package opuscc

import (
	"reflect"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

var _ reflect.Type
var _ unsafe.Pointer

func Opus_clt_mdct_forward_c(tls *libc.TLS, l uintptr, in uintptr, out uintptr, window uintptr, overlap int32, shift int32, stride int32, arch int32) {
	var N, N2, N4, i int32
	var _saved_stack, f, f2, fp, st, st1, t, t1, trig, wp1, wp2, xp1, xp2, yp, yp1, yp11, yp2, v1, v10, v12, v14, v16, v18, v20, v22, v24, v3, v6, v8 uintptr
	var im, re, t0, t01, t11, t12, yi, yi1, yr, yr1 float32
	var scale OpusT_celt_coef
	var yc OpusT_kiss_fft_cpx
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = N, N2, N4, _saved_stack, f, f2, fp, i, im, re, scale, st, st1, t, t0, t01, t1, t11, t12, trig, wp1, wp2, xp1, xp2, yc, yi, yi1, yp, yp1, yp11, yp2, yr, yr1, v1, v10, v12, v14, v16, v18, v20, v22, v24, v3, v6, v8
	st1 = *(*uintptr)(unsafe.Pointer(l + 8 + uintptr(shift)*8))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _2
	_2:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _4
_4:
	_saved_stack = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack
	_ = arch
	scale = (*OpusT_kiss_fft_state)(unsafe.Pointer(st1)).Fscale
	N = (*OpusT_mdct_lookup)(unsafe.Pointer(l)).Fn
	trig = (*OpusT_mdct_lookup)(unsafe.Pointer(l)).Ftrig
	i = 0
	for {
		if !(i < shift) {
			break
		}
		N = N >> int32(1)
		trig = trig + uintptr(N)*4
		goto _5
	_5:
		;
		i = i + 1
	}
	N2 = N >> int32(1)
	N4 = N >> int32(2)
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _7
	_7:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _9
_9:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v6 = libc.Xmalloc(tls, uint64(16))
		goto _11
	_11:
		st = v6
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v8 = st
	goto _13
_13:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v8)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v10 = libc.Xmalloc(tls, uint64(16))
		goto _15
	_15:
		st = v10
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v12 = st
	goto _17
_17:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v14 = libc.Xmalloc(tls, uint64(16))
		goto _19
	_19:
		st = v14
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v16 = st
	goto _21
_21:
	if !(int64(int32(uint64(uint32(N2))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v12)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v16)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5512, int32(152))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v18 = libc.Xmalloc(tls, uint64(16))
		goto _23
	_23:
		st = v18
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v20 = st
	goto _25
_25:
	*(*uintptr)(unsafe.Pointer(v20 + 8)) += uintptr(uint64(uint32(N2)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v22 = libc.Xmalloc(tls, uint64(16))
		goto _27
	_27:
		st = v22
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v24 = st
	goto _29
_29:
	f = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v24)).Fglobal_stack - uintptr(uint64(uint32(N2))*(uint64(4)/uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _31
	_31:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _33
_33:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v6 = libc.Xmalloc(tls, uint64(16))
		goto _35
	_35:
		st = v6
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v8 = st
	goto _37
_37:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v8)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v10 = libc.Xmalloc(tls, uint64(16))
		goto _39
	_39:
		st = v10
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v12 = st
	goto _41
_41:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v14 = libc.Xmalloc(tls, uint64(16))
		goto _43
	_43:
		st = v14
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v16 = st
	goto _45
_45:
	if !(int64(int32(uint64(uint32(N4))*(uint64(8)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v12)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v16)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5512, int32(153))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v18 = libc.Xmalloc(tls, uint64(16))
		goto _47
	_47:
		st = v18
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v20 = st
	goto _49
_49:
	*(*uintptr)(unsafe.Pointer(v20 + 8)) += uintptr(uint64(uint32(N4)) * (uint64(8) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v22 = libc.Xmalloc(tls, uint64(16))
		goto _51
	_51:
		st = v22
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v24 = st
	goto _53
_53:
	f2 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v24)).Fglobal_stack - uintptr(uint64(uint32(N4))*(uint64(8)/uint64(1)))
	/* Consider the input to be composed of four blocks: [a, b, c, d] */
	/* Window, shuffle, fold */
	/* Temp pointers to make it really clear to the compiler what we're doing */
	xp1 = in + uintptr(overlap>>int32(1))*4
	xp2 = in + uintptr(N2)*4 - uintptr(1)*4 + uintptr(overlap>>int32(1))*4
	yp = f
	wp1 = window + uintptr(overlap>>int32(1))*4
	wp2 = window + uintptr(overlap>>int32(1))*4 - uintptr(1)*4
	i = 0
	for {
		if !(i < (overlap+int32(3))>>int32(2)) {
			break
		}
		/* Real part arranged as -d-cR, Imag part arranged as -b+aR*/
		v1 = yp
		yp += 4
		*(*float32)(unsafe.Pointer(v1)) = float32(*(*float32)(unsafe.Pointer(xp1 + uintptr(N2)*4))**(*OpusT_celt_coef)(unsafe.Pointer(wp2))) + float32(*(*float32)(unsafe.Pointer(xp2))**(*OpusT_celt_coef)(unsafe.Pointer(wp1)))
		v1 = yp
		yp += 4
		*(*float32)(unsafe.Pointer(v1)) = float32(*(*float32)(unsafe.Pointer(xp1))**(*OpusT_celt_coef)(unsafe.Pointer(wp1))) - float32(*(*float32)(unsafe.Pointer(xp2 + uintptr(-N2)*4))**(*OpusT_celt_coef)(unsafe.Pointer(wp2)))
		xp1 = xp1 + uintptr(2)*4
		xp2 = xp2 - uintptr(2)*4
		wp1 = wp1 + uintptr(2)*4
		wp2 = wp2 - uintptr(2)*4
		goto _54
	_54:
		;
		i = i + 1
	}
	wp1 = window
	wp2 = window + uintptr(overlap)*4 - uintptr(1)*4
	for {
		if !(i < N4-(overlap+int32(3))>>int32(2)) {
			break
		}
		/* Real part arranged as a-bR, Imag part arranged as -c-dR */
		v1 = yp
		yp += 4
		*(*float32)(unsafe.Pointer(v1)) = *(*float32)(unsafe.Pointer(xp2))
		v1 = yp
		yp += 4
		*(*float32)(unsafe.Pointer(v1)) = *(*float32)(unsafe.Pointer(xp1))
		xp1 = xp1 + uintptr(2)*4
		xp2 = xp2 - uintptr(2)*4
		goto _57
	_57:
		;
		i = i + 1
	}
	for {
		if !(i < N4) {
			break
		}
		/* Real part arranged as a-bR, Imag part arranged as -c-dR */
		v1 = yp
		yp += 4
		*(*float32)(unsafe.Pointer(v1)) = -float32(*(*float32)(unsafe.Pointer(xp1 + uintptr(-N2)*4))**(*OpusT_celt_coef)(unsafe.Pointer(wp1))) + float32(*(*float32)(unsafe.Pointer(xp2))**(*OpusT_celt_coef)(unsafe.Pointer(wp2)))
		v1 = yp
		yp += 4
		*(*float32)(unsafe.Pointer(v1)) = float32(*(*float32)(unsafe.Pointer(xp1))**(*OpusT_celt_coef)(unsafe.Pointer(wp2))) + float32(*(*float32)(unsafe.Pointer(xp2 + uintptr(N2)*4))**(*OpusT_celt_coef)(unsafe.Pointer(wp1)))
		xp1 = xp1 + uintptr(2)*4
		xp2 = xp2 - uintptr(2)*4
		wp1 = wp1 + uintptr(2)*4
		wp2 = wp2 - uintptr(2)*4
		goto _60
	_60:
		;
		i = i + 1
	}
	/* Pre-rotation */
	yp1 = f
	t = trig
	i = 0
	for {
		if !(i < N4) {
			break
		}
		t0 = *(*float32)(unsafe.Pointer(t + uintptr(i)*4))
		t11 = *(*float32)(unsafe.Pointer(t + uintptr(N4+i)*4))
		v1 = yp1
		yp1 += 4
		re = *(*float32)(unsafe.Pointer(v1))
		v1 = yp1
		yp1 += 4
		im = *(*float32)(unsafe.Pointer(v1))
		yr = float32(re*t0) - float32(im*t11)
		yi = float32(im*t0) + float32(re*t11)
		/* For QEXT, it's best to scale before the FFT, but otherwise it's best to scale after.
		   For floating-point it doesn't matter. */
		yc.Fr = float32(yr * scale)
		yc.Fi = float32(yi * scale)
		*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(f2 + uintptr(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_kiss_fft_state)(unsafe.Pointer(st1)).Fbitrev + uintptr(i)*2)))*8)) = yc
		goto _63
	_63:
		;
		i = i + 1
	}
	/* N/4 complex FFT, does not downscale anymore */
	Opus_opus_fft_impl(tls, st1, f2)
	/* Post-rotate */
	/* Temp pointers to make it really clear to the compiler what we're doing */
	fp = f2
	yp11 = out
	yp2 = out + uintptr(stride*(N2-int32(1)))*4
	t1 = trig
	/* Temp pointers to make it really clear to the compiler what we're doing */
	i = 0
	for {
		if !(i < N4) {
			break
		}
		t01 = *(*float32)(unsafe.Pointer(t1 + uintptr(i)*4))
		t12 = *(*float32)(unsafe.Pointer(t1 + uintptr(N4+i)*4))
		yr1 = float32((*OpusT_kiss_fft_cpx)(unsafe.Pointer(fp)).Fi*t12) - float32((*OpusT_kiss_fft_cpx)(unsafe.Pointer(fp)).Fr*t01)
		yi1 = float32((*OpusT_kiss_fft_cpx)(unsafe.Pointer(fp)).Fr*t12) + float32((*OpusT_kiss_fft_cpx)(unsafe.Pointer(fp)).Fi*t01)
		*(*float32)(unsafe.Pointer(yp11)) = yr1
		*(*float32)(unsafe.Pointer(yp2)) = yi1
		fp += 8
		yp11 = yp11 + uintptr(int32(2)*stride)*4
		yp2 = yp2 - uintptr(int32(2)*stride)*4
		goto _66
	_66:
		;
		i = i + 1
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _68
	_68:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _70
_70:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
}

func Opus_clt_mdct_backward_c(tls *libc.TLS, l uintptr, in uintptr, out uintptr, window uintptr, overlap int32, shift int32, stride int32, arch int32) {
	var N, N2, N4, i, rev int32
	var bitrev, t, t1, trig, wp1, wp2, xp1, xp11, xp2, yp, yp0, yp1, yp11, v3 uintptr
	var im, re, t0, t11, x11, x21, yi, yi1, yr, yr1 float32
	var x1, x2 OpusT_opus_val32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = N, N2, N4, bitrev, i, im, re, rev, t, t0, t1, t11, trig, wp1, wp2, x1, x11, x2, x21, xp1, xp11, xp2, yi, yi1, yp, yp0, yp1, yp11, yr, yr1, v3
	_ = arch
	N = (*OpusT_mdct_lookup)(unsafe.Pointer(l)).Fn
	trig = (*OpusT_mdct_lookup)(unsafe.Pointer(l)).Ftrig
	i = 0
	for {
		if !(i < shift) {
			break
		}
		N = N >> int32(1)
		trig = trig + uintptr(N)*4
		goto _1
	_1:
		;
		i = i + 1
	}
	N2 = N >> int32(1)
	N4 = N >> int32(2)
	/* Pre-rotate */
	/* Temp pointers to make it really clear to the compiler what we're doing */
	xp1 = in
	xp2 = in + uintptr(stride*(N2-int32(1)))*4
	yp = out + uintptr(overlap>>int32(1))*4
	t = trig
	bitrev = (*OpusT_kiss_fft_state)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(l + 8 + uintptr(shift)*8)))).Fbitrev
	i = 0
	for {
		if !(i < N4) {
			break
		}
		v3 = bitrev
		bitrev += 2
		rev = int32(*(*OpusT_opus_int16)(unsafe.Pointer(v3)))
		x1 = *(*float32)(unsafe.Pointer(xp1))
		x2 = *(*float32)(unsafe.Pointer(xp2))
		yr = OpusT_opus_val32(x2**(*float32)(unsafe.Pointer(t + uintptr(i)*4))) + OpusT_opus_val32(x1**(*float32)(unsafe.Pointer(t + uintptr(N4+i)*4)))
		yi = OpusT_opus_val32(x1**(*float32)(unsafe.Pointer(t + uintptr(i)*4))) - OpusT_opus_val32(x2**(*float32)(unsafe.Pointer(t + uintptr(N4+i)*4)))
		/* We swap real and imag because we use an FFT instead of an IFFT. */
		*(*float32)(unsafe.Pointer(yp + uintptr(int32(2)*rev+int32(1))*4)) = yr
		*(*float32)(unsafe.Pointer(yp + uintptr(int32(2)*rev)*4)) = yi
		/* Storing the pre-rotation directly in the bitrev order. */
		xp1 = xp1 + uintptr(int32(2)*stride)*4
		xp2 = xp2 - uintptr(int32(2)*stride)*4
		goto _2
	_2:
		;
		i = i + 1
	}
	Opus_opus_fft_impl(tls, *(*uintptr)(unsafe.Pointer(l + 8 + uintptr(shift)*8)), out+uintptr(overlap>>int32(1))*4)
	/* Post-rotate and de-shuffle from both ends of the buffer at once to make
	   it in-place. */
	yp0 = out + uintptr(overlap>>int32(1))*4
	yp1 = out + uintptr(overlap>>int32(1))*4 + uintptr(N2)*4 - uintptr(2)*4
	t1 = trig
	/* Loop to (N4+1)>>1 to handle odd N4. When N4 is odd, the
	   middle pair will be computed twice. */
	i = 0
	for {
		if !(i < (N4+int32(1))>>int32(1)) {
			break
		}
		/* We swap real and imag because we're using an FFT instead of an IFFT. */
		re = *(*float32)(unsafe.Pointer(yp0 + 1*4))
		im = *(*float32)(unsafe.Pointer(yp0))
		t0 = *(*float32)(unsafe.Pointer(t1 + uintptr(i)*4))
		t11 = *(*float32)(unsafe.Pointer(t1 + uintptr(N4+i)*4))
		/* We'd scale up by 2 here, but instead it's done when mixing the windows */
		yr1 = float32(re*t0) + float32(im*t11)
		yi1 = float32(re*t11) - float32(im*t0)
		/* We swap real and imag because we're using an FFT instead of an IFFT. */
		re = *(*float32)(unsafe.Pointer(yp1 + 1*4))
		im = *(*float32)(unsafe.Pointer(yp1))
		*(*float32)(unsafe.Pointer(yp0)) = yr1
		*(*float32)(unsafe.Pointer(yp1 + 1*4)) = yi1
		t0 = *(*float32)(unsafe.Pointer(t1 + uintptr(N4-i-int32(1))*4))
		t11 = *(*float32)(unsafe.Pointer(t1 + uintptr(N2-i-int32(1))*4))
		/* We'd scale up by 2 here, but instead it's done when mixing the windows */
		yr1 = float32(re*t0) + float32(im*t11)
		yi1 = float32(re*t11) - float32(im*t0)
		*(*float32)(unsafe.Pointer(yp1)) = yr1
		*(*float32)(unsafe.Pointer(yp0 + 1*4)) = yi1
		yp0 = yp0 + uintptr(2)*4
		yp1 = yp1 - uintptr(2)*4
		goto _4
	_4:
		;
		i = i + 1
	}
	/* Mirror on both sides for TDAC */
	xp11 = out + uintptr(overlap)*4 - uintptr(1)*4
	yp11 = out
	wp1 = window
	wp2 = window + uintptr(overlap)*4 - uintptr(1)*4
	i = 0
	for {
		if !(i < overlap/int32(2)) {
			break
		}
		x11 = *(*float32)(unsafe.Pointer(xp11))
		x21 = *(*float32)(unsafe.Pointer(yp11))
		v3 = yp11
		yp11 += 4
		*(*float32)(unsafe.Pointer(v3)) = float32(x21**(*OpusT_celt_coef)(unsafe.Pointer(wp2))) - float32(x11**(*OpusT_celt_coef)(unsafe.Pointer(wp1)))
		v3 = xp11
		xp11 -= 4
		*(*float32)(unsafe.Pointer(v3)) = float32(x21**(*OpusT_celt_coef)(unsafe.Pointer(wp1))) + float32(x11**(*OpusT_celt_coef)(unsafe.Pointer(wp2)))
		wp1 += 4
		wp2 -= 4
		goto _5
	_5:
		;
		i = i + 1
	}
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

func kf_bfly21(tls *libc.TLS, Fout uintptr, fstride OpusT_size_t, st OpusT_mini_kiss_fft_cfg, m int32) {
	var Fout2, tw1 uintptr
	var t OpusT_mini_kiss_fft_cpx
	var v1 int32
	_, _, _, _ = Fout2, t, tw1, v1
	tw1 = st + 264
	Fout2 = Fout + uintptr(m)*8
	for {
		t.Fr = float32((*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout2)).Fr*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw1)).Fr) - float32((*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout2)).Fi*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw1)).Fi)
		t.Fi = float32((*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout2)).Fr*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw1)).Fi) + float32((*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout2)).Fi*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw1)).Fr)
		tw1 = tw1 + uintptr(fstride)*8
		(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout2)).Fr = (*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fr - t.Fr
		(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout2)).Fi = (*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fi - t.Fi
		(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fr += t.Fr
		(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fi += t.Fi
		Fout2 += 8
		Fout += 8
		goto _2
	_2:
		;
		m = m - 1
		v1 = m
		if !(v1 != 0) {
			break
		}
	}
}

func kf_bfly41(tls *libc.TLS, Fout uintptr, fstride OpusT_size_t, st OpusT_mini_kiss_fft_cfg, m OpusT_size_t) {
	var k, m2, m3, v3 OpusT_size_t
	var scratch [6]OpusT_mini_kiss_fft_cpx
	var tw1, tw2, tw3, v1, v2 uintptr
	_, _, _, _, _, _, _, _, _, _ = k, m2, m3, scratch, tw1, tw2, tw3, v1, v2, v3
	k = m
	m2 = uint64(2) * m
	m3 = uint64(3) * m
	v2 = st + 264
	tw1 = v2
	v1 = v2
	tw2 = v1
	tw3 = v1
	for {
		scratch[0].Fr = float32((*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fr*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw1)).Fr) - float32((*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fi*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw1)).Fi)
		scratch[0].Fi = float32((*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fr*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw1)).Fi) + float32((*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fi*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw1)).Fr)
		scratch[int32(1)].Fr = float32((*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fr*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw2)).Fr) - float32((*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fi*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw2)).Fi)
		scratch[int32(1)].Fi = float32((*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fr*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw2)).Fi) + float32((*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fi*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw2)).Fr)
		scratch[int32(2)].Fr = float32((*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m3)*8))).Fr*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw3)).Fr) - float32((*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m3)*8))).Fi*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw3)).Fi)
		scratch[int32(2)].Fi = float32((*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m3)*8))).Fr*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw3)).Fi) + float32((*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m3)*8))).Fi*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw3)).Fr)
		scratch[int32(5)].Fr = (*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fr - scratch[int32(1)].Fr
		scratch[int32(5)].Fi = (*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fi - scratch[int32(1)].Fi
		(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fr += scratch[int32(1)].Fr
		(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fi += scratch[int32(1)].Fi
		scratch[int32(3)].Fr = scratch[0].Fr + scratch[int32(2)].Fr
		scratch[int32(3)].Fi = scratch[0].Fi + scratch[int32(2)].Fi
		scratch[int32(4)].Fr = scratch[0].Fr - scratch[int32(2)].Fr
		scratch[int32(4)].Fi = scratch[0].Fi - scratch[int32(2)].Fi
		(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fr = (*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fr - scratch[int32(3)].Fr
		(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fi = (*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fi - scratch[int32(3)].Fi
		tw1 = tw1 + uintptr(fstride)*8
		tw2 = tw2 + uintptr(fstride*uint64(2))*8
		tw3 = tw3 + uintptr(fstride*uint64(3))*8
		(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fr += scratch[int32(3)].Fr
		(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fi += scratch[int32(3)].Fi
		if (*mini_kiss_fft_state)(unsafe.Pointer(st)).Finverse != 0 {
			(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fr = scratch[int32(5)].Fr - scratch[int32(4)].Fi
			(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fi = scratch[int32(5)].Fi + scratch[int32(4)].Fr
			(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m3)*8))).Fr = scratch[int32(5)].Fr + scratch[int32(4)].Fi
			(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m3)*8))).Fi = scratch[int32(5)].Fi - scratch[int32(4)].Fr
		} else {
			(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fr = scratch[int32(5)].Fr + scratch[int32(4)].Fi
			(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fi = scratch[int32(5)].Fi - scratch[int32(4)].Fr
			(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m3)*8))).Fr = scratch[int32(5)].Fr - scratch[int32(4)].Fi
			(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m3)*8))).Fi = scratch[int32(5)].Fi + scratch[int32(4)].Fr
		}
		Fout += 8
		goto _4
	_4:
		;
		k = k - 1
		v3 = k
		if !(v3 != 0) {
			break
		}
	}
}

func kf_bfly31(tls *libc.TLS, Fout uintptr, fstride OpusT_size_t, st OpusT_mini_kiss_fft_cfg, m OpusT_size_t) {
	var epi3 OpusT_mini_kiss_fft_cpx
	var k, m2, v2 OpusT_size_t
	var scratch [5]OpusT_mini_kiss_fft_cpx
	var tw1, tw2, v1 uintptr
	_, _, _, _, _, _, _, _ = epi3, k, m2, scratch, tw1, tw2, v1, v2
	k = m
	m2 = uint64(2) * m
	epi3 = *(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(st + 264 + uintptr(fstride*m)*8))
	v1 = st + 264
	tw2 = v1
	tw1 = v1
	for {
		scratch[int32(1)].Fr = float32((*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fr*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw1)).Fr) - float32((*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fi*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw1)).Fi)
		scratch[int32(1)].Fi = float32((*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fr*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw1)).Fi) + float32((*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fi*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw1)).Fr)
		scratch[int32(2)].Fr = float32((*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fr*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw2)).Fr) - float32((*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fi*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw2)).Fi)
		scratch[int32(2)].Fi = float32((*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fr*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw2)).Fi) + float32((*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fi*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw2)).Fr)
		scratch[int32(3)].Fr = scratch[int32(1)].Fr + scratch[int32(2)].Fr
		scratch[int32(3)].Fi = scratch[int32(1)].Fi + scratch[int32(2)].Fi
		scratch[0].Fr = scratch[int32(1)].Fr - scratch[int32(2)].Fr
		scratch[0].Fi = scratch[int32(1)].Fi - scratch[int32(2)].Fi
		tw1 = tw1 + uintptr(fstride)*8
		tw2 = tw2 + uintptr(fstride*uint64(2))*8
		(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fr = (*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fr - float32(scratch[int32(3)].Fr*float32(0.5))
		(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fi = (*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fi - float32(scratch[int32(3)].Fi*float32(0.5))
		scratch[0].Fr *= epi3.Fi
		scratch[0].Fi *= epi3.Fi
		(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fr += scratch[int32(3)].Fr
		(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fi += scratch[int32(3)].Fi
		(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fr = (*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fr + scratch[0].Fi
		(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fi = (*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fi - scratch[0].Fr
		(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fr -= scratch[0].Fi
		(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fi += scratch[0].Fr
		Fout += 8
		goto _3
	_3:
		;
		k = k - 1
		v2 = k
		if !(v2 != 0) {
			break
		}
	}
}

func kf_bfly51(tls *libc.TLS, Fout uintptr, fstride OpusT_size_t, st OpusT_mini_kiss_fft_cfg, m int32) {
	var Fout0, Fout1, Fout2, Fout3, Fout4, tw, twiddles uintptr
	var scratch [13]OpusT_mini_kiss_fft_cpx
	var u int32
	var ya, yb OpusT_mini_kiss_fft_cpx
	_, _, _, _, _, _, _, _, _, _, _ = Fout0, Fout1, Fout2, Fout3, Fout4, scratch, tw, twiddles, u, ya, yb
	twiddles = st + 264
	ya = *(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(twiddles + uintptr(fstride*uint64(uint32(m)))*8))
	yb = *(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(twiddles + uintptr(fstride*uint64(2)*uint64(uint32(m)))*8))
	Fout0 = Fout
	Fout1 = Fout0 + uintptr(m)*8
	Fout2 = Fout0 + uintptr(int32(2)*m)*8
	Fout3 = Fout0 + uintptr(int32(3)*m)*8
	Fout4 = Fout0 + uintptr(int32(4)*m)*8
	tw = st + 264
	u = 0
	for {
		if !(u < m) {
			break
		}
		scratch[0] = *(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout0))
		scratch[int32(1)].Fr = float32((*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout1)).Fr*(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(u))*fstride)*8))).Fr) - float32((*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout1)).Fi*(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(u))*fstride)*8))).Fi)
		scratch[int32(1)].Fi = float32((*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout1)).Fr*(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(u))*fstride)*8))).Fi) + float32((*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout1)).Fi*(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(u))*fstride)*8))).Fr)
		scratch[int32(2)].Fr = float32((*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout2)).Fr*(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(2)*u))*fstride)*8))).Fr) - float32((*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout2)).Fi*(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(2)*u))*fstride)*8))).Fi)
		scratch[int32(2)].Fi = float32((*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout2)).Fr*(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(2)*u))*fstride)*8))).Fi) + float32((*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout2)).Fi*(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(2)*u))*fstride)*8))).Fr)
		scratch[int32(3)].Fr = float32((*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout3)).Fr*(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(3)*u))*fstride)*8))).Fr) - float32((*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout3)).Fi*(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(3)*u))*fstride)*8))).Fi)
		scratch[int32(3)].Fi = float32((*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout3)).Fr*(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(3)*u))*fstride)*8))).Fi) + float32((*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout3)).Fi*(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(3)*u))*fstride)*8))).Fr)
		scratch[int32(4)].Fr = float32((*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout4)).Fr*(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(4)*u))*fstride)*8))).Fr) - float32((*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout4)).Fi*(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(4)*u))*fstride)*8))).Fi)
		scratch[int32(4)].Fi = float32((*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout4)).Fr*(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(4)*u))*fstride)*8))).Fi) + float32((*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout4)).Fi*(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(4)*u))*fstride)*8))).Fr)
		scratch[int32(7)].Fr = scratch[int32(1)].Fr + scratch[int32(4)].Fr
		scratch[int32(7)].Fi = scratch[int32(1)].Fi + scratch[int32(4)].Fi
		scratch[int32(10)].Fr = scratch[int32(1)].Fr - scratch[int32(4)].Fr
		scratch[int32(10)].Fi = scratch[int32(1)].Fi - scratch[int32(4)].Fi
		scratch[int32(8)].Fr = scratch[int32(2)].Fr + scratch[int32(3)].Fr
		scratch[int32(8)].Fi = scratch[int32(2)].Fi + scratch[int32(3)].Fi
		scratch[int32(9)].Fr = scratch[int32(2)].Fr - scratch[int32(3)].Fr
		scratch[int32(9)].Fi = scratch[int32(2)].Fi - scratch[int32(3)].Fi
		*(*float32)(unsafe.Pointer(Fout0)) += scratch[int32(7)].Fr + scratch[int32(8)].Fr
		*(*float32)(unsafe.Pointer(Fout0 + 4)) += scratch[int32(7)].Fi + scratch[int32(8)].Fi
		scratch[int32(5)].Fr = scratch[0].Fr + float32(scratch[int32(7)].Fr*ya.Fr) + float32(scratch[int32(8)].Fr*yb.Fr)
		scratch[int32(5)].Fi = scratch[0].Fi + float32(scratch[int32(7)].Fi*ya.Fr) + float32(scratch[int32(8)].Fi*yb.Fr)
		scratch[int32(6)].Fr = float32(scratch[int32(10)].Fi*ya.Fi) + float32(scratch[int32(9)].Fi*yb.Fi)
		scratch[int32(6)].Fi = -float32(scratch[int32(10)].Fr*ya.Fi) - float32(scratch[int32(9)].Fr*yb.Fi)
		(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout1)).Fr = scratch[int32(5)].Fr - scratch[int32(6)].Fr
		(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout1)).Fi = scratch[int32(5)].Fi - scratch[int32(6)].Fi
		(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout4)).Fr = scratch[int32(5)].Fr + scratch[int32(6)].Fr
		(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout4)).Fi = scratch[int32(5)].Fi + scratch[int32(6)].Fi
		scratch[int32(11)].Fr = scratch[0].Fr + float32(scratch[int32(7)].Fr*yb.Fr) + float32(scratch[int32(8)].Fr*ya.Fr)
		scratch[int32(11)].Fi = scratch[0].Fi + float32(scratch[int32(7)].Fi*yb.Fr) + float32(scratch[int32(8)].Fi*ya.Fr)
		scratch[int32(12)].Fr = -float32(scratch[int32(10)].Fi*yb.Fi) + float32(scratch[int32(9)].Fi*ya.Fi)
		scratch[int32(12)].Fi = float32(scratch[int32(10)].Fr*yb.Fi) - float32(scratch[int32(9)].Fr*ya.Fi)
		(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout2)).Fr = scratch[int32(11)].Fr + scratch[int32(12)].Fr
		(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout2)).Fi = scratch[int32(11)].Fi + scratch[int32(12)].Fi
		(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout3)).Fr = scratch[int32(11)].Fr - scratch[int32(12)].Fr
		(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout3)).Fi = scratch[int32(11)].Fi - scratch[int32(12)].Fi
		Fout0 += 8
		Fout1 += 8
		Fout2 += 8
		Fout3 += 8
		Fout4 += 8
		goto _1
	_1:
		;
		u = u + 1
	}
}

func kf_work(tls *libc.TLS, Fout uintptr, f uintptr, fstride OpusT_size_t, in_stride int32, factors uintptr, st OpusT_mini_kiss_fft_cfg) {
	var Fout_beg, Fout_end, v1, v2, v3 uintptr
	var m, p int32
	var v6 bool
	_, _, _, _, _, _, _, _ = Fout_beg, Fout_end, m, p, v1, v2, v3, v6
	Fout_beg = Fout
	v1 = factors
	factors += 4
	p = *(*int32)(unsafe.Pointer(v1))
	v2 = factors
	factors += 4                      /* the radix  */
	m = *(*int32)(unsafe.Pointer(v2)) /* stage's fft length/p */
	Fout_end = Fout + uintptr(p*m)*8
	if m == int32(1) {
		for {
			*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(Fout)) = *(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(f))
			f = f + uintptr(fstride*uint64(uint32(in_stride)))*8
			goto _4
		_4:
			;
			Fout += 8
			v3 = Fout
			if !(v3 != Fout_end) {
				break
			}
		}
	} else {
		for {
			/* recursive call:
			   DFT of size m*p performed by doing
			   p instances of smaller DFTs of size m,
			   each one takes a decimated version of the input */
			kf_work(tls, Fout, f, fstride*uint64(uint32(p)), in_stride, factors, st)
			f = f + uintptr(fstride*uint64(uint32(in_stride)))*8
			goto _5
		_5:
			;
			Fout = Fout + uintptr(m)*8
			if !(Fout != Fout_end) {
				break
			}
		}
	}
	Fout = Fout_beg
	/* recombine the p smaller DFTs*/
	switch p {
	case int32(2):
		kf_bfly21(tls, Fout, fstride, st, m)
	case int32(3):
		kf_bfly31(tls, Fout, fstride, st, uint64(uint32(m)))
	case int32(4):
		kf_bfly41(tls, Fout, fstride, st, uint64(uint32(m)))
	case int32(5):
		kf_bfly51(tls, Fout, fstride, st, m)
	default:
		if v6 = libc.Bool(0 != 0); !v6 {
			libc.X__assert_fail(tls, __ccgo_ts+5527, __ccgo_ts+5529, int32(317), uintptr(unsafe.Pointer(&__func__)))
		}
		_ = v6 || libc.Bool(int32(0) != 0)
	}
}

var __func__ = [8]int8{'k', 'f', '_', 'w', 'o', 'r', 'k'}

// C documentation
//
//	/*  facbuf is populated by p1,m1,p2,m2, ...
//	    where
//	    p[i] * m[i] = m[i-1]
//	    m0 = n                  */
func kf_factor(tls *libc.TLS, n int32, facbuf uintptr) {
	var floor_sqrt float64
	var p int32
	var v1 uintptr
	_, _, _ = floor_sqrt, p, v1
	p = int32(4)
	floor_sqrt = libc.Xfloor(tls, libc.Xsqrt(tls, float64(n)))
	/*factor out powers of 4, powers of 2, then any remaining primes */
	for cond := true; cond; cond = n > int32(1) {
		for n%p != 0 {
			switch p {
			case int32(4):
				p = int32(2)
			case int32(2):
				p = int32(3)
			default:
				p = p + int32(2)
				break
			}
			if float64(p) > floor_sqrt {
				p = n
			} /* no more factors, skip to end */
		}
		n = n / p
		v1 = facbuf
		facbuf += 4
		*(*int32)(unsafe.Pointer(v1)) = p
		v1 = facbuf
		facbuf += 4
		*(*int32)(unsafe.Pointer(v1)) = n
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
func Opus_mini_kiss_fft_alloc(tls *libc.TLS, nfft int32, inverse_fft int32, mem uintptr, lenmem uintptr) (r OpusT_mini_kiss_fft_cfg) {
	var i int32
	var memneeded OpusT_size_t
	var phase, pi float64
	var st OpusT_mini_kiss_fft_cfg
	_, _, _, _, _ = i, memneeded, phase, pi, st
	st = uintptr(uint32(0))
	memneeded = uint64(272) + uint64(8)*uint64(uint32(nfft-int32(1))) /* twiddle factors*/
	if lenmem == uintptr(uint32(0)) {
		st = libc.Xmalloc(tls, memneeded)
	} else {
		if mem != uintptr(uint32(0)) && *(*OpusT_size_t)(unsafe.Pointer(lenmem)) >= memneeded {
			st = mem
		}
		*(*OpusT_size_t)(unsafe.Pointer(lenmem)) = memneeded
	}
	if st != 0 {
		(*mini_kiss_fft_state)(unsafe.Pointer(st)).Fnfft = nfft
		(*mini_kiss_fft_state)(unsafe.Pointer(st)).Finverse = inverse_fft
		i = 0
		for {
			if !(i < nfft) {
				break
			}
			pi = float64(3.141592653589793)
			phase = float64(float64(float64(-int32(2))*pi)*float64(i)) / float64(nfft)
			if (*mini_kiss_fft_state)(unsafe.Pointer(st)).Finverse != 0 {
				phase = phase * float64(-int32(1))
			}
			(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(st + 264 + uintptr(i)*8)).Fr = float32(libc.Xcos(tls, phase))
			(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(st + 264 + uintptr(i)*8)).Fi = float32(libc.Xsin(tls, phase))
			goto _1
		_1:
			;
			i = i + 1
		}
		kf_factor(tls, nfft, st+8)
	}
	return st
}

func Opus_mini_kiss_fft_stride(tls *libc.TLS, st OpusT_mini_kiss_fft_cfg, fin uintptr, fout uintptr, in_stride int32) {
	var v1 bool
	_ = v1
	if v1 = fin != fout; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+5549, __ccgo_ts+5529, int32(391), uintptr(unsafe.Pointer(&__func__1)))
	}
	_ = v1 || libc.Bool(int32(0) != 0)
	kf_work(tls, fout, fin, uint64(1), in_stride, st+8, st)
}

var __func__1 = [21]int8{'m', 'i', 'n', 'i', '_', 'k', 'i', 's', 's', '_', 'f', 'f', 't', '_', 's', 't', 'r', 'i', 'd', 'e'}

func Opus_mini_kiss_fft(tls *libc.TLS, cfg OpusT_mini_kiss_fft_cfg, fin uintptr, fout uintptr) {
	Opus_mini_kiss_fft_stride(tls, cfg, fin, fout, int32(1))
}

type OpusT_mini_kiss_fftr_cfg = uintptr

type mini_kiss_fftr_state = struct {
	Fsubstate       OpusT_mini_kiss_fft_cfg
	Ftmpbuf         uintptr
	Fsuper_twiddles uintptr
}

type OpusT_mini_kiss_fftr_state = struct {
	Fsubstate       OpusT_mini_kiss_fft_cfg
	Ftmpbuf         uintptr
	Fsuper_twiddles uintptr
}

func Opus_mini_kiss_fftr_alloc(tls *libc.TLS, nfft int32, inverse_fft int32, mem uintptr, lenmem uintptr) (r OpusT_mini_kiss_fftr_cfg) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i int32
	var memneeded OpusT_size_t
	var phase float64
	var st OpusT_mini_kiss_fftr_cfg
	var v1 bool
	var _ /* subsize at bp+0 */ OpusT_size_t
	_, _, _, _, _ = i, memneeded, phase, st, v1
	st = uintptr(uint32(0))
	*(*OpusT_size_t)(unsafe.Pointer(bp)) = uint64(0)
	if v1 = nfft&int32(1) == 0; !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+5561, __ccgo_ts+5529, int32(416), uintptr(unsafe.Pointer(&__func__2)))
	}
	_ = v1 || libc.Bool(int32(0) != 0)
	nfft = nfft >> int32(1)
	Opus_mini_kiss_fft_alloc(tls, nfft, inverse_fft, uintptr(uint32(0)), bp)
	memneeded = uint64(24) + *(*OpusT_size_t)(unsafe.Pointer(bp)) + uint64(8)*uint64(uint32(nfft*int32(3)/int32(2)))
	if lenmem == uintptr(uint32(0)) {
		st = libc.Xmalloc(tls, memneeded)
	} else {
		if *(*OpusT_size_t)(unsafe.Pointer(lenmem)) >= memneeded {
			st = mem
		}
		*(*OpusT_size_t)(unsafe.Pointer(lenmem)) = memneeded
	}
	if !(st != 0) {
		return uintptr(uint32(0))
	}
	(*mini_kiss_fftr_state)(unsafe.Pointer(st)).Fsubstate = st + uintptr(uint32(1))*24 /*just beyond kiss_fftr_state struct */
	(*mini_kiss_fftr_state)(unsafe.Pointer(st)).Ftmpbuf = (*mini_kiss_fftr_state)(unsafe.Pointer(st)).Fsubstate + uintptr(*(*OpusT_size_t)(unsafe.Pointer(bp)))
	(*mini_kiss_fftr_state)(unsafe.Pointer(st)).Fsuper_twiddles = (*mini_kiss_fftr_state)(unsafe.Pointer(st)).Ftmpbuf + uintptr(nfft)*8
	Opus_mini_kiss_fft_alloc(tls, nfft, inverse_fft, (*mini_kiss_fftr_state)(unsafe.Pointer(st)).Fsubstate, bp)
	i = 0
	for {
		if !(i < nfft/int32(2)) {
			break
		}
		phase = float64(-float64(3.141592653589793) * (float64(i+int32(1))/float64(nfft) + float64(0.5)))
		if inverse_fft != 0 {
			phase = phase * float64(-int32(1))
		}
		(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer((*mini_kiss_fftr_state)(unsafe.Pointer(st)).Fsuper_twiddles + uintptr(i)*8)).Fr = float32(libc.Xcos(tls, phase))
		(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer((*mini_kiss_fftr_state)(unsafe.Pointer(st)).Fsuper_twiddles + uintptr(i)*8)).Fi = float32(libc.Xsin(tls, phase))
		goto _2
	_2:
		;
		i = i + 1
	}
	return st
}

var __func__2 = [21]int8{'m', 'i', 'n', 'i', '_', 'k', 'i', 's', 's', '_', 'f', 'f', 't', 'r', '_', 'a', 'l', 'l', 'o', 'c'}

func Opus_mini_kiss_fftr(tls *libc.TLS, st OpusT_mini_kiss_fftr_cfg, timedata uintptr, freqdata uintptr) {
	var f1k, f2k, fpk, fpnk, tdc, tw OpusT_mini_kiss_fft_cpx
	var k, ncfft int32
	var v1 bool
	var v2 float32
	_, _, _, _, _, _, _, _, _, _ = f1k, f2k, fpk, fpnk, k, ncfft, tdc, tw, v1, v2
	if v1 = !((*mini_kiss_fft_state)(unsafe.Pointer((*mini_kiss_fftr_state)(unsafe.Pointer(st)).Fsubstate)).Finverse != 0); !v1 {
		libc.X__assert_fail(tls, __ccgo_ts+5577, __ccgo_ts+5529, int32(453), uintptr(unsafe.Pointer(&__func__3)))
	}
	_ = v1 || libc.Bool(int32(0) != 0)
	ncfft = (*mini_kiss_fft_state)(unsafe.Pointer((*mini_kiss_fftr_state)(unsafe.Pointer(st)).Fsubstate)).Fnfft
	/*perform the parallel fft of two real signals packed in real,imag*/
	Opus_mini_kiss_fft(tls, (*mini_kiss_fftr_state)(unsafe.Pointer(st)).Fsubstate, timedata, (*mini_kiss_fftr_state)(unsafe.Pointer(st)).Ftmpbuf)
	/* The real part of the DC element of the frequency spectrum in st->tmpbuf
	 * contains the sum of the even-numbered elements of the input time sequence
	 * The imag part is the sum of the odd-numbered elements
	 *
	 * The sum of tdc.r and tdc.i is the sum of the input time sequence.
	 *      yielding DC of input time sequence
	 * The difference of tdc.r - tdc.i is the sum of the input (dot product) [1,-1,1,-1...
	 *      yielding Nyquist bin of input time sequence
	 */
	tdc.Fr = (*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer((*mini_kiss_fftr_state)(unsafe.Pointer(st)).Ftmpbuf))).Fr
	tdc.Fi = (*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer((*mini_kiss_fftr_state)(unsafe.Pointer(st)).Ftmpbuf))).Fi
	(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(freqdata))).Fr = tdc.Fr + tdc.Fi
	(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(freqdata + uintptr(ncfft)*8))).Fr = tdc.Fr - tdc.Fi
	v2 = float32(0)
	(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(freqdata))).Fi = v2
	(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(freqdata + uintptr(ncfft)*8))).Fi = v2
	k = int32(1)
	for {
		if !(k <= ncfft/int32(2)) {
			break
		}
		fpk = *(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer((*mini_kiss_fftr_state)(unsafe.Pointer(st)).Ftmpbuf + uintptr(k)*8))
		fpnk.Fr = (*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer((*mini_kiss_fftr_state)(unsafe.Pointer(st)).Ftmpbuf + uintptr(ncfft-k)*8))).Fr
		fpnk.Fi = -(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer((*mini_kiss_fftr_state)(unsafe.Pointer(st)).Ftmpbuf + uintptr(ncfft-k)*8))).Fi
		f1k.Fr = fpk.Fr + fpnk.Fr
		f1k.Fi = fpk.Fi + fpnk.Fi
		f2k.Fr = fpk.Fr - fpnk.Fr
		f2k.Fi = fpk.Fi - fpnk.Fi
		tw.Fr = float32(f2k.Fr*(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer((*mini_kiss_fftr_state)(unsafe.Pointer(st)).Fsuper_twiddles + uintptr(k-int32(1))*8))).Fr) - float32(f2k.Fi*(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer((*mini_kiss_fftr_state)(unsafe.Pointer(st)).Fsuper_twiddles + uintptr(k-int32(1))*8))).Fi)
		tw.Fi = float32(f2k.Fr*(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer((*mini_kiss_fftr_state)(unsafe.Pointer(st)).Fsuper_twiddles + uintptr(k-int32(1))*8))).Fi) + float32(f2k.Fi*(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer((*mini_kiss_fftr_state)(unsafe.Pointer(st)).Fsuper_twiddles + uintptr(k-int32(1))*8))).Fr)
		(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(freqdata + uintptr(k)*8))).Fr = float32((f1k.Fr + tw.Fr) * float32(0.5))
		(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(freqdata + uintptr(k)*8))).Fi = float32((f1k.Fi + tw.Fi) * float32(0.5))
		(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(freqdata + uintptr(ncfft-k)*8))).Fr = float32((f1k.Fr - tw.Fr) * float32(0.5))
		(*(*OpusT_mini_kiss_fft_cpx)(unsafe.Pointer(freqdata + uintptr(ncfft-k)*8))).Fi = float32((tw.Fi - f1k.Fi) * float32(0.5))
		goto _3
	_3:
		;
		k = k + 1
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

func Opus_ec_laplace_encode(tls *libc.TLS, enc uintptr, value uintptr, fs uint32, decay int32) {
	var di, i, ndi_max, s, val, v2 int32
	var fl, v3 uint32
	_, _, _, _, _, _, _, _ = di, fl, i, ndi_max, s, val, v2, v3
	val = *(*int32)(unsafe.Pointer(value))
	fl = uint32(0)
	if val != 0 {
		s = -libc.BoolInt32(val < 0)
		val = val + s ^ s
		fl = fs
		fs = ec_laplace_get_freq1(tls, fs, decay)
		/* Search the decaying part of the PDF.*/
		i = int32(1)
		for {
			if !(fs > uint32(0) && i < val) {
				break
			}
			fs = fs * uint32(2)
			fl = fl + (fs + uint32(int32(2)*(int32(1)<<int32(LAPLACE_LOG_MINP))))
			fs = fs * uint32(decay) >> int32(15)
			goto _1
		_1:
			;
			i = i + 1
		}
		/* Everything beyond that has probability LAPLACE_MINP. */
		if !(fs != 0) {
			ndi_max = int32((uint32(32768) - fl + uint32(int32(1)<<int32(LAPLACE_LOG_MINP)) - uint32(1)) >> LAPLACE_LOG_MINP)
			ndi_max = (ndi_max - s) >> int32(1)
			if val-i < ndi_max-int32(1) {
				v2 = val - i
			} else {
				v2 = ndi_max - int32(1)
			}
			di = v2
			fl = fl + uint32((int32(2)*di+int32(1)+s)*(int32(1)<<int32(LAPLACE_LOG_MINP)))
			if uint32(int32(1)<<int32(LAPLACE_LOG_MINP)) < uint32(32768)-fl {
				v3 = uint32(int32(1) << int32(LAPLACE_LOG_MINP))
			} else {
				v3 = uint32(32768) - fl
			}
			fs = v3
			*(*int32)(unsafe.Pointer(value)) = i + di + s ^ s
		} else {
			fs = fs + uint32(int32(1)<<int32(LAPLACE_LOG_MINP))
			fl = fl + fs&uint32(^s)
		}
		if !(fl+fs <= uint32(32768)) {
			Opus_celt_fatal(tls, __ccgo_ts+5600, __ccgo_ts+5631, int32(88))
		}
		if !(fs > uint32(0)) {
			Opus_celt_fatal(tls, __ccgo_ts+5649, __ccgo_ts+5631, int32(89))
		}
	}
	Opus_ec_encode_bin(tls, enc, fl, fl+fs, uint32(15))
}

func Opus_ec_laplace_decode(tls *libc.TLS, dec uintptr, fs uint32, decay int32) (r int32) {
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

func Opus_ec_laplace_encode_p0(tls *libc.TLS, enc uintptr, value int32, p0 OpusT_opus_uint16, decay OpusT_opus_uint16) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i, s, v1, v2 int32
	var _ /* icdf at bp+6 */ [8]OpusT_opus_uint16
	var _ /* sign_icdf at bp+0 */ [3]OpusT_opus_uint16
	_, _, _, _ = i, s, v1, v2
	(*(*[3]OpusT_opus_uint16)(unsafe.Pointer(bp)))[0] = uint16(int32(32768) - int32(p0))
	(*(*[3]OpusT_opus_uint16)(unsafe.Pointer(bp)))[int32(1)] = uint16(int32((*(*[3]OpusT_opus_uint16)(unsafe.Pointer(bp)))[0]) / int32(2))
	(*(*[3]OpusT_opus_uint16)(unsafe.Pointer(bp)))[int32(2)] = uint16(0)
	if value == 0 {
		v1 = 0
	} else {
		if value > 0 {
			v2 = int32(1)
		} else {
			v2 = int32(2)
		}
		v1 = v2
	}
	s = v1
	Opus_ec_enc_icdf16(tls, enc, s, bp, uint32(15))
	value = libc.Xabs(tls, value)
	if value != 0 {
		if int32(7) > int32(decay) {
			v1 = int32(7)
		} else {
			v1 = int32(decay)
		}
		(*(*[8]OpusT_opus_uint16)(unsafe.Pointer(bp + 6)))[0] = uint16(v1)
		i = int32(1)
		for {
			if !(i < int32(7)) {
				break
			}
			if int32(7)-i > int32((*(*[8]OpusT_opus_uint16)(unsafe.Pointer(bp + 6)))[i-int32(1)])*int32(decay)>>int32(15) {
				v1 = int32(7) - i
			} else {
				v1 = int32((*(*[8]OpusT_opus_uint16)(unsafe.Pointer(bp + 6)))[i-int32(1)]) * int32(decay) >> int32(15)
			}
			(*(*[8]OpusT_opus_uint16)(unsafe.Pointer(bp + 6)))[i] = uint16(v1)
			goto _4
		_4:
			;
			i = i + 1
		}
		(*(*[8]OpusT_opus_uint16)(unsafe.Pointer(bp + 6)))[int32(7)] = uint16(0)
		value = value - 1
		for cond := true; cond; cond = value >= 0 {
			if value < int32(7) {
				v1 = value
			} else {
				v1 = int32(7)
			}
			Opus_ec_enc_icdf16(tls, enc, v1, bp+6, uint32(15))
			value = value - int32(7)
		}
	}
}

func Opus_ec_laplace_decode_p0(tls *libc.TLS, dec uintptr, p0 OpusT_opus_uint16, decay OpusT_opus_uint16) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var i, s, v, value, v1 int32
	var _ /* icdf at bp+6 */ [8]OpusT_opus_uint16
	var _ /* sign_icdf at bp+0 */ [3]OpusT_opus_uint16
	_, _, _, _, _ = i, s, v, value, v1
	(*(*[3]OpusT_opus_uint16)(unsafe.Pointer(bp)))[0] = uint16(int32(32768) - int32(p0))
	(*(*[3]OpusT_opus_uint16)(unsafe.Pointer(bp)))[int32(1)] = uint16(int32((*(*[3]OpusT_opus_uint16)(unsafe.Pointer(bp)))[0]) / int32(2))
	(*(*[3]OpusT_opus_uint16)(unsafe.Pointer(bp)))[int32(2)] = uint16(0)
	s = Opus_ec_dec_icdf16(tls, dec, bp, uint32(15))
	if s == int32(2) {
		s = -int32(1)
	}
	if s != 0 {
		if int32(7) > int32(decay) {
			v1 = int32(7)
		} else {
			v1 = int32(decay)
		}
		(*(*[8]OpusT_opus_uint16)(unsafe.Pointer(bp + 6)))[0] = uint16(v1)
		i = int32(1)
		for {
			if !(i < int32(7)) {
				break
			}
			if int32(7)-i > int32((*(*[8]OpusT_opus_uint16)(unsafe.Pointer(bp + 6)))[i-int32(1)])*int32(decay)>>int32(15) {
				v1 = int32(7) - i
			} else {
				v1 = int32((*(*[8]OpusT_opus_uint16)(unsafe.Pointer(bp + 6)))[i-int32(1)]) * int32(decay) >> int32(15)
			}
			(*(*[8]OpusT_opus_uint16)(unsafe.Pointer(bp + 6)))[i] = uint16(v1)
			goto _2
		_2:
			;
			i = i + 1
		}
		(*(*[8]OpusT_opus_uint16)(unsafe.Pointer(bp + 6)))[int32(7)] = uint16(0)
		value = int32(1)
		for cond := true; cond; cond = v == int32(7) {
			v = Opus_ec_dec_icdf16(tls, dec, bp+6, uint32(15))
			value = value + v
		}
		return s * value
	} else {
		return 0
	}
	return r
}
