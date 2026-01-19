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

func Opus__celt_lpc(tls *libc.TLS, _lpc uintptr, ac uintptr, p int32) {
	var error1, r, rr, tmp1, tmp2 OpusT_opus_val32
	var i, j int32
	var lpc uintptr
	_, _, _, _, _, _, _, _ = error1, i, j, lpc, r, rr, tmp1, tmp2
	error1 = *(*OpusT_opus_val32)(unsafe.Pointer(ac))
	lpc = _lpc
	libc.Xmemset(tls, lpc, 0, uint64(uint32(p))*uint64(4))
	if *(*OpusT_opus_val32)(unsafe.Pointer(ac)) > float32(1e-10) {
		i = 0
		for {
			if !(i < p) {
				break
			}
			/* Sum up this iteration's reflection coefficient */
			rr = float32(0)
			j = 0
			for {
				if !(j < i) {
					break
				}
				rr = rr + float32(*(*float32)(unsafe.Pointer(lpc + uintptr(j)*4))**(*OpusT_opus_val32)(unsafe.Pointer(ac + uintptr(i-j)*4)))
				goto _2
			_2:
				;
				j = j + 1
			}
			rr = rr + *(*OpusT_opus_val32)(unsafe.Pointer(ac + uintptr(i+int32(1))*4))
			r = -(rr / error1)
			/*  Update LPC coefficients and total error */
			*(*float32)(unsafe.Pointer(lpc + uintptr(i)*4)) = r
			j = 0
			for {
				if !(j < (i+int32(1))>>int32(1)) {
					break
				}
				tmp1 = *(*float32)(unsafe.Pointer(lpc + uintptr(j)*4))
				tmp2 = *(*float32)(unsafe.Pointer(lpc + uintptr(i-int32(1)-j)*4))
				*(*float32)(unsafe.Pointer(lpc + uintptr(j)*4)) = tmp1 + OpusT_opus_val32(r*tmp2)
				*(*float32)(unsafe.Pointer(lpc + uintptr(i-int32(1)-j)*4)) = tmp2 + OpusT_opus_val32(r*tmp1)
				goto _3
			_3:
				;
				j = j + 1
			}
			error1 = error1 - OpusT_opus_val32(OpusT_opus_val32(r*r)*error1)
			/* Bail out once we get 30 dB gain */
			if error1 <= OpusT_opus_val32(float32(0.001)**(*OpusT_opus_val32)(unsafe.Pointer(ac))) {
				break
			}
			goto _1
		_1:
			;
			i = i + 1
		}
	}
}

func Opus_celt_fir_c(tls *libc.TLS, x1 uintptr, num uintptr, y1 uintptr, N int32, ord int32, arch int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _saved_stack, rnum, st, v1, v11, v13, v15, v17, v19, v21, v23, v25, v27, v3, v31, v32, v33, v35, v36, v37, v5, v7, v9 uintptr
	var i, j, j1, v34, v47, v50 int32
	var sum2 OpusT_opus_val32
	var tmp, tmp1, tmp2, tmp3, y_0, y_1, y_2, y_3 OpusT_opus_val16
	var _ /* sum at bp+0 */ [4]OpusT_opus_val32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = _saved_stack, i, j, j1, rnum, st, sum2, tmp, tmp1, tmp2, tmp3, y_0, y_1, y_2, y_3, v1, v11, v13, v15, v17, v19, v21, v23, v25, v27, v3, v31, v32, v33, v34, v35, v36, v37, v47, v5, v50, v7, v9
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
	if !(x1 != y1) {
		Opus_celt_fatal(tls, __ccgo_ts+3305, __ccgo_ts+3330, int32(157))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _6
	_6:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _8
_8:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _10
	_10:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v7 = st
	goto _12
_12:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		goto _14
	_14:
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v11 = st
	goto _16
_16:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		goto _18
	_18:
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v15 = st
	goto _20
_20:
	if !(int64(int32(uint64(uint32(ord))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3330, int32(158))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		goto _22
	_22:
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v19 = st
	goto _24
_24:
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32(ord)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		goto _26
	_26:
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v23 = st
	goto _28
_28:
	rnum = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32(ord))*(uint64(4)/uint64(1)))
	i = 0
	for {
		if !(i < ord) {
			break
		}
		*(*OpusT_opus_val16)(unsafe.Pointer(rnum + uintptr(i)*4)) = *(*OpusT_opus_val16)(unsafe.Pointer(num + uintptr(ord-i-int32(1))*4))
		goto _29
	_29:
		;
		i = i + 1
	}
	i = 0
	for {
		if !(i < N-int32(3)) {
			break
		}
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[0] = *(*OpusT_opus_val16)(unsafe.Pointer(x1 + uintptr(i)*4))
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(1)] = *(*OpusT_opus_val16)(unsafe.Pointer(x1 + uintptr(i+int32(1))*4))
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(2)] = *(*OpusT_opus_val16)(unsafe.Pointer(x1 + uintptr(i+int32(2))*4))
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)] = *(*OpusT_opus_val16)(unsafe.Pointer(x1 + uintptr(i+int32(3))*4))
		_ = arch
		v1 = rnum
		v3 = x1 + uintptr(i)*4 - uintptr(ord)*4
		v5 = bp
		v34 = ord
		if !(v34 >= int32(3)) {
			Opus_celt_fatal(tls, __ccgo_ts+3349, __ccgo_ts+3374, int32(69))
		}
		y_3 = float32(0)
		v7 = v3
		v3 += 4
		y_0 = *(*OpusT_opus_val16)(unsafe.Pointer(v7))
		v9 = v3
		v3 += 4
		y_1 = *(*OpusT_opus_val16)(unsafe.Pointer(v9))
		v11 = v3
		v3 += 4
		y_2 = *(*OpusT_opus_val16)(unsafe.Pointer(v11))
		j = int32(0)
		for {
			if !(j < v34-int32(3)) {
				break
			}
			v13 = v1
			v1 += 4
			tmp = *(*OpusT_opus_val16)(unsafe.Pointer(v13))
			v15 = v3
			v3 += 4
			y_3 = *(*OpusT_opus_val16)(unsafe.Pointer(v15))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp*y_3)
			v17 = v1
			v1 += 4
			tmp = *(*OpusT_opus_val16)(unsafe.Pointer(v17))
			v19 = v3
			v3 += 4
			y_0 = *(*OpusT_opus_val16)(unsafe.Pointer(v19))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp*y_0)
			v21 = v1
			v1 += 4
			tmp = *(*OpusT_opus_val16)(unsafe.Pointer(v21))
			v23 = v3
			v3 += 4
			y_1 = *(*OpusT_opus_val16)(unsafe.Pointer(v23))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp*y_1)
			v25 = v1
			v1 += 4
			tmp = *(*OpusT_opus_val16)(unsafe.Pointer(v25))
			v27 = v3
			v3 += 4
			y_2 = *(*OpusT_opus_val16)(unsafe.Pointer(v27))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp*y_2)
			goto _38
		_38:
			;
			j = j + int32(4)
		}
		v47 = j
		j = j + 1
		if v47 < v34 {
			v31 = v1
			v1 += 4
			tmp1 = *(*OpusT_opus_val16)(unsafe.Pointer(v31))
			v32 = v3
			v3 += 4
			y_3 = *(*OpusT_opus_val16)(unsafe.Pointer(v32))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp1*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp1*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp1*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp1*y_3)
		}
		v50 = j
		j = j + 1
		if v50 < v34 {
			v33 = v1
			v1 += 4
			tmp2 = *(*OpusT_opus_val16)(unsafe.Pointer(v33))
			v35 = v3
			v3 += 4
			y_0 = *(*OpusT_opus_val16)(unsafe.Pointer(v35))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp2*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp2*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp2*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp2*y_0)
		}
		if j < v34 {
			v36 = v1
			v1 += 4
			tmp3 = *(*OpusT_opus_val16)(unsafe.Pointer(v36))
			v37 = v3
			v3 += 4
			y_1 = *(*OpusT_opus_val16)(unsafe.Pointer(v37))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp3*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp3*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp3*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp3*y_1)
		}
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i)*4)) = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[0]
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+int32(1))*4)) = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(1)]
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+int32(2))*4)) = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(2)]
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+int32(3))*4)) = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)]
		goto _30
	_30:
		;
		i = i + int32(4)
	}
	for {
		if !(i < N) {
			break
		}
		sum2 = *(*OpusT_opus_val16)(unsafe.Pointer(x1 + uintptr(i)*4))
		j1 = 0
		for {
			if !(j1 < ord) {
				break
			}
			sum2 = sum2 + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(rnum + uintptr(j1)*4))**(*OpusT_opus_val16)(unsafe.Pointer(x1 + uintptr(i+j1-ord)*4)))
			goto _56
		_56:
			;
			j1 = j1 + 1
		}
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i)*4)) = sum2
		goto _55
	_55:
		;
		i = i + 1
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _58
	_58:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _60
_60:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
}

func Opus_celt_iir(tls *libc.TLS, _x uintptr, den uintptr, _y uintptr, N int32, ord int32, mem uintptr, arch int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _saved_stack, rden, st, y1, v1, v11, v13, v15, v17, v19, v21, v23, v25, v27, v29, v3, v31, v33, v35, v37, v39, v5, v7, v9 uintptr
	var i, j, j1, v60, v73, v76 int32
	var sum2 OpusT_opus_val32
	var tmp, tmp1, tmp2, tmp3, y_0, y_1, y_2, y_3 OpusT_opus_val16
	var _ /* sum at bp+0 */ [4]OpusT_opus_val32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = _saved_stack, i, j, j1, rden, st, sum2, tmp, tmp1, tmp2, tmp3, y1, y_0, y_1, y_2, y_3, v1, v11, v13, v15, v17, v19, v21, v23, v25, v27, v29, v3, v31, v33, v35, v37, v39, v5, v60, v7, v73, v76, v9
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
	if !(ord&int32(3) == int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+3390, __ccgo_ts+3330, int32(225))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _6
	_6:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _8
_8:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _10
	_10:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v7 = st
	goto _12
_12:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		goto _14
	_14:
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v11 = st
	goto _16
_16:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		goto _18
	_18:
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v15 = st
	goto _20
_20:
	if !(int64(int32(uint64(uint32(ord))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3330, int32(226))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		goto _22
	_22:
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v19 = st
	goto _24
_24:
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32(ord)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		goto _26
	_26:
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v23 = st
	goto _28
_28:
	rden = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32(ord))*(uint64(4)/uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _30
	_30:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _32
_32:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _34
	_34:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v7 = st
	goto _36
_36:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		goto _38
	_38:
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v11 = st
	goto _40
_40:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		goto _42
	_42:
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v15 = st
	goto _44
_44:
	if !(int64(int32(uint64(uint32(N+ord))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3330, int32(227))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		goto _46
	_46:
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v19 = st
	goto _48
_48:
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32(N+ord)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		goto _50
	_50:
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v23 = st
	goto _52
_52:
	y1 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32(N+ord))*(uint64(4)/uint64(1)))
	i = 0
	for {
		if !(i < ord) {
			break
		}
		*(*OpusT_opus_val16)(unsafe.Pointer(rden + uintptr(i)*4)) = *(*OpusT_opus_val16)(unsafe.Pointer(den + uintptr(ord-i-int32(1))*4))
		goto _53
	_53:
		;
		i = i + 1
	}
	i = 0
	for {
		if !(i < ord) {
			break
		}
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i)*4)) = -*(*OpusT_opus_val16)(unsafe.Pointer(mem + uintptr(ord-i-int32(1))*4))
		goto _54
	_54:
		;
		i = i + 1
	}
	for {
		if !(i < N+ord) {
			break
		}
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i)*4)) = float32(0)
		goto _55
	_55:
		;
		i = i + 1
	}
	i = 0
	for {
		if !(i < N-int32(3)) {
			break
		}
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[0] = *(*OpusT_opus_val32)(unsafe.Pointer(_x + uintptr(i)*4))
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(1)] = *(*OpusT_opus_val32)(unsafe.Pointer(_x + uintptr(i+int32(1))*4))
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(2)] = *(*OpusT_opus_val32)(unsafe.Pointer(_x + uintptr(i+int32(2))*4))
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)] = *(*OpusT_opus_val32)(unsafe.Pointer(_x + uintptr(i+int32(3))*4))
		_ = arch
		v1 = rden
		v3 = y1 + uintptr(i)*4
		v5 = bp
		v60 = ord
		if !(v60 >= int32(3)) {
			Opus_celt_fatal(tls, __ccgo_ts+3349, __ccgo_ts+3374, int32(69))
		}
		y_3 = float32(0)
		v7 = v3
		v3 += 4
		y_0 = *(*OpusT_opus_val16)(unsafe.Pointer(v7))
		v9 = v3
		v3 += 4
		y_1 = *(*OpusT_opus_val16)(unsafe.Pointer(v9))
		v11 = v3
		v3 += 4
		y_2 = *(*OpusT_opus_val16)(unsafe.Pointer(v11))
		j = int32(0)
		for {
			if !(j < v60-int32(3)) {
				break
			}
			v13 = v1
			v1 += 4
			tmp = *(*OpusT_opus_val16)(unsafe.Pointer(v13))
			v15 = v3
			v3 += 4
			y_3 = *(*OpusT_opus_val16)(unsafe.Pointer(v15))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp*y_3)
			v17 = v1
			v1 += 4
			tmp = *(*OpusT_opus_val16)(unsafe.Pointer(v17))
			v19 = v3
			v3 += 4
			y_0 = *(*OpusT_opus_val16)(unsafe.Pointer(v19))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp*y_0)
			v21 = v1
			v1 += 4
			tmp = *(*OpusT_opus_val16)(unsafe.Pointer(v21))
			v23 = v3
			v3 += 4
			y_1 = *(*OpusT_opus_val16)(unsafe.Pointer(v23))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp*y_1)
			v25 = v1
			v1 += 4
			tmp = *(*OpusT_opus_val16)(unsafe.Pointer(v25))
			v27 = v3
			v3 += 4
			y_2 = *(*OpusT_opus_val16)(unsafe.Pointer(v27))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp*y_2)
			goto _64
		_64:
			;
			j = j + int32(4)
		}
		v73 = j
		j = j + 1
		if v73 < v60 {
			v29 = v1
			v1 += 4
			tmp1 = *(*OpusT_opus_val16)(unsafe.Pointer(v29))
			v31 = v3
			v3 += 4
			y_3 = *(*OpusT_opus_val16)(unsafe.Pointer(v31))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp1*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp1*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp1*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp1*y_3)
		}
		v76 = j
		j = j + 1
		if v76 < v60 {
			v33 = v1
			v1 += 4
			tmp2 = *(*OpusT_opus_val16)(unsafe.Pointer(v33))
			v35 = v3
			v3 += 4
			y_0 = *(*OpusT_opus_val16)(unsafe.Pointer(v35))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp2*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp2*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp2*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp2*y_0)
		}
		if j < v60 {
			v37 = v1
			v1 += 4
			tmp3 = *(*OpusT_opus_val16)(unsafe.Pointer(v37))
			v39 = v3
			v3 += 4
			y_1 = *(*OpusT_opus_val16)(unsafe.Pointer(v39))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp3*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp3*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp3*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp3*y_1)
		}
		/* Patch up the result to compensate for the fact that this is an IIR */
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+ord)*4)) = -(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[0]
		*(*OpusT_opus_val32)(unsafe.Pointer(_y + uintptr(i)*4)) = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[0]
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(1)] = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(1)] + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+ord)*4))**(*OpusT_opus_val16)(unsafe.Pointer(den)))
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+ord+int32(1))*4)) = -(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(1)]
		*(*OpusT_opus_val32)(unsafe.Pointer(_y + uintptr(i+int32(1))*4)) = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(1)]
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(2)] = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(2)] + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+ord+int32(1))*4))**(*OpusT_opus_val16)(unsafe.Pointer(den)))
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(2)] = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(2)] + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+ord)*4))**(*OpusT_opus_val16)(unsafe.Pointer(den + 1*4)))
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+ord+int32(2))*4)) = -(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(2)]
		*(*OpusT_opus_val32)(unsafe.Pointer(_y + uintptr(i+int32(2))*4)) = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(2)]
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)] = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)] + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+ord+int32(2))*4))**(*OpusT_opus_val16)(unsafe.Pointer(den)))
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)] = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)] + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+ord+int32(1))*4))**(*OpusT_opus_val16)(unsafe.Pointer(den + 1*4)))
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)] = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)] + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+ord)*4))**(*OpusT_opus_val16)(unsafe.Pointer(den + 2*4)))
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+ord+int32(3))*4)) = -(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)]
		*(*OpusT_opus_val32)(unsafe.Pointer(_y + uintptr(i+int32(3))*4)) = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)]
		goto _56
	_56:
		;
		i = i + int32(4)
	}
	for {
		if !(i < N) {
			break
		}
		sum2 = *(*OpusT_opus_val32)(unsafe.Pointer(_x + uintptr(i)*4))
		j1 = 0
		for {
			if !(j1 < ord) {
				break
			}
			sum2 = sum2 - OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(rden + uintptr(j1)*4))**(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+j1)*4)))
			goto _82
		_82:
			;
			j1 = j1 + 1
		}
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+ord)*4)) = sum2
		*(*OpusT_opus_val32)(unsafe.Pointer(_y + uintptr(i)*4)) = sum2
		goto _81
	_81:
		;
		i = i + 1
	}
	i = 0
	for {
		if !(i < ord) {
			break
		}
		*(*OpusT_opus_val16)(unsafe.Pointer(mem + uintptr(i)*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(_y + uintptr(N-i-int32(1))*4))
		goto _83
	_83:
		;
		i = i + 1
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _85
	_85:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _87
_87:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
}

func Opus__celt_autocorr(tls *libc.TLS, x uintptr, ac uintptr, window uintptr, overlap int32, lag int32, n int32, arch int32) (r int32) {
	var _saved_stack, st, xptr, xx, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var d OpusT_opus_val32
	var fastN, i, k, shift int32
	var w OpusT_opus_val16
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = _saved_stack, d, fastN, i, k, shift, st, w, xptr, xx, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9
	fastN = n - lag
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
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _6
	_6:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _8
_8:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _10
	_10:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v7 = st
	goto _12
_12:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		goto _14
	_14:
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v11 = st
	goto _16
_16:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		goto _18
	_18:
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v15 = st
	goto _20
_20:
	if !(int64(int32(uint64(uint32(n))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3330, int32(301))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		goto _22
	_22:
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v19 = st
	goto _24
_24:
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32(n)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		goto _26
	_26:
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v23 = st
	goto _28
_28:
	xx = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32(n))*(uint64(4)/uint64(1)))
	if !(n > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+3419, __ccgo_ts+3330, int32(302))
	}
	if !(overlap >= int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+3441, __ccgo_ts+3330, int32(303))
	}
	if overlap == 0 {
		xptr = x
	} else {
		i = 0
		for {
			if !(i < n) {
				break
			}
			*(*OpusT_opus_val16)(unsafe.Pointer(xx + uintptr(i)*4)) = *(*OpusT_opus_val16)(unsafe.Pointer(x + uintptr(i)*4))
			goto _29
		_29:
			;
			i = i + 1
		}
		i = 0
		for {
			if !(i < overlap) {
				break
			}
			w = *(*OpusT_celt_coef)(unsafe.Pointer(window + uintptr(i)*4))
			*(*OpusT_opus_val16)(unsafe.Pointer(xx + uintptr(i)*4)) = OpusT_opus_val16(*(*OpusT_opus_val16)(unsafe.Pointer(x + uintptr(i)*4)) * w)
			*(*OpusT_opus_val16)(unsafe.Pointer(xx + uintptr(n-i-int32(1))*4)) = OpusT_opus_val16(*(*OpusT_opus_val16)(unsafe.Pointer(x + uintptr(n-i-int32(1))*4)) * w)
			goto _30
		_30:
			;
			i = i + 1
		}
		xptr = xx
	}
	shift = 0
	Opus_celt_pitch_xcorr_c(tls, xptr, xptr, ac, fastN, lag+int32(1), arch)
	k = 0
	for {
		if !(k <= lag) {
			break
		}
		i = k + fastN
		d = float32(0)
		for {
			if !(i < n) {
				break
			}
			d = d + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(xptr + uintptr(i)*4))**(*OpusT_opus_val16)(unsafe.Pointer(xptr + uintptr(i-k)*4)))
			goto _32
		_32:
			;
			i = i + 1
		}
		*(*OpusT_opus_val32)(unsafe.Pointer(ac + uintptr(k)*4)) += d
		goto _31
	_31:
		;
		k = k + 1
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _34
	_34:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _36
_36:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
	return shift
}

var log2_x_norm_coeff5 = [8]float32{
	0: float32(1),
	1: float32(0.8888888955116272),
	2: float32(0.8),
	3: float32(0.7272727489471436),
	4: float32(0.6666666865348816),
	5: float32(0.6153846383094788),
	6: float32(0.5714285969734192),
	7: float32(0.5333333611488342),
}
var log2_y_norm_coeff5 = [8]float32{
	1: float32(0.1699250042438507),
	2: float32(0.32192808389663696),
	3: float32(0.45943161845207214),
	4: float32(0.5849624872207642),
	5: float32(0.7004396915435791),
	6: float32(0.8073549270629883),
	7: float32(0.9068905711174011),
}

/* The guts header contains all the multiplication and addition macros that are defined for
   complex numbers.  It also declares the kf_ internal functions.
*/
