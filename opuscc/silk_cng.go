// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

package opuscc

import (
	"reflect"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

var _ reflect.Type
var _ unsafe.Pointer

// C documentation
//
//	/* Generates excitation for CNG LPC synthesis */
func silk_CNG_exc(tls *libc.TLS, exc_Q14 uintptr, exc_buf_Q14 uintptr, length int32, rand_seed uintptr) {
	var exc_mask, i, idx int32
	var seed OpusT_opus_int32
	_, _, _, _ = exc_mask, i, idx, seed
	exc_mask = int32(CNG_BUF_MASK_MAX)
	for exc_mask > length {
		exc_mask = exc_mask >> int32(1)
	}
	seed = *(*OpusT_opus_int32)(unsafe.Pointer(rand_seed))
	i = 0
	for {
		if !(i < length) {
			break
		}
		seed = int32(uint32(int32(RAND_INCREMENT)) + uint32(seed)*uint32(int32(RAND_MULTIPLIER)))
		idx = seed >> int32(24) & exc_mask
		_ = idx >= int32(0)
		_ = idx <= int32(CNG_BUF_MASK_MAX)
		*(*OpusT_opus_int32)(unsafe.Pointer(exc_Q14 + uintptr(i)*4)) = *(*OpusT_opus_int32)(unsafe.Pointer(exc_buf_Q14 + uintptr(idx)*4))
		goto _1
	_1:
		;
		i = i + 1
	}
	*(*OpusT_opus_int32)(unsafe.Pointer(rand_seed)) = seed
}

func Opus_silk_CNG_Reset(tls *libc.TLS, psDec uintptr) {
	var NLSF_acc_Q15, NLSF_step_Q15, i int32
	_, _, _ = NLSF_acc_Q15, NLSF_step_Q15, i
	NLSF_step_Q15 = int32(silk_int16_MAX1) / ((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order + int32(1))
	NLSF_acc_Q15 = 0
	i = 0
	for {
		if !(i < (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order) {
			break
		}
		NLSF_acc_Q15 = NLSF_acc_Q15 + NLSF_step_Q15
		*(*OpusT_opus_int16)(unsafe.Pointer(psDec + 2892 + 1280 + uintptr(i)*2)) = int16(NLSF_acc_Q15)
		goto _1
	_1:
		;
		i = i + 1
	}
	(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FsCNG.FCNG_smth_Gain_Q16 = 0
	(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FsCNG.Frand_seed = int32(3176576)
}

// C documentation
//
//	/* Updates CNG estimate, and applies the CNG when packet was lost   */
func Opus_silk_CNG(tls *libc.TLS, psDec uintptr, psDecCtrl uintptr, frame uintptr, length int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var CNG_sig_Q14, _saved_stack, psCNG, st, v1, v11, v13, v15, v17, v19, v21, v23, v25, v3, v6, v9 uintptr
	var LPC_pred_Q10, gain_Q10, gain_Q16, lzeros, max_Gain_Q16, y, v33, v34, v36, v37, v38, v41, v43 OpusT_opus_int32
	var i, subfr, v40, v42, v52, v54, v58, v59, v60, v61, v62, v63, v64, v65, v66 int32
	var m, r, x OpusT_opus_uint32
	var _ /* A_Q12 at bp+8 */ [16]OpusT_opus_int16
	var _ /* frac_Q7 at bp+4 */ OpusT_opus_int32
	var _ /* lz at bp+0 */ OpusT_opus_int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = CNG_sig_Q14, LPC_pred_Q10, _saved_stack, gain_Q10, gain_Q16, i, lzeros, m, max_Gain_Q16, psCNG, r, st, subfr, x, y, v1, v11, v13, v15, v17, v19, v21, v23, v25, v3, v33, v34, v36, v37, v38, v40, v41, v42, v43, v52, v54, v58, v59, v6, v60, v61, v62, v63, v64, v65, v66, v9
	psCNG = psDec + 2892
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
	if (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Ffs_kHz != (*OpusT_silk_CNG_struct)(unsafe.Pointer(psCNG)).Ffs_kHz {
		/* Reset state */
		Opus_silk_CNG_Reset(tls, psDec)
		(*OpusT_silk_CNG_struct)(unsafe.Pointer(psCNG)).Ffs_kHz = (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Ffs_kHz
	}
	if (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FlossCnt == 0 && (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FprevSignalType == TYPE_NO_VOICE_ACTIVITY {
		/* Update CNG parameters */
		/* Smoothing of LSF's  */
		i = 0
		for {
			if !(i < (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order) {
				break
			}
			v1 = psCNG + 1280 + uintptr(i)*2
			*(*OpusT_opus_int16)(unsafe.Pointer(v1)) = OpusT_opus_int16(int32(*(*OpusT_opus_int16)(unsafe.Pointer(v1))) + int32(int64(int32(*(*OpusT_opus_int16)(unsafe.Pointer(psDec + 2344 + uintptr(i)*2)))-int32(*(*OpusT_opus_int16)(unsafe.Pointer(psCNG + 1280 + uintptr(i)*2))))*int64(int16(int32(CNG_NLSF_SMTH_Q16)))>>int32(16)))
			goto _5
		_5:
			;
			i = i + 1
		}
		/* Find the subframe with the highest gain */
		max_Gain_Q16 = 0
		subfr = 0
		i = 0
		for {
			if !(i < (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr) {
				break
			}
			if *(*OpusT_opus_int32)(unsafe.Pointer(psDecCtrl + 16 + uintptr(i)*4)) > max_Gain_Q16 {
				max_Gain_Q16 = *(*OpusT_opus_int32)(unsafe.Pointer(psDecCtrl + 16 + uintptr(i)*4))
				subfr = i
			}
			goto _7
		_7:
			;
			i = i + 1
		}
		/* Update CNG excitation buffer with excitation from this subframe */
		libc.Xmemmove(tls, psCNG+uintptr((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length)*4, psCNG, uint64(uint32(((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr-int32(1))*(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length))*uint64(4))
		libc.Xmemcpy(tls, psCNG, psDec+4+uintptr(subfr*(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length)*4, uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length))*uint64(4))
		/* Smooth gains */
		i = 0
		for {
			if !(i < (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr) {
				break
			}
			*(*OpusT_opus_int32)(unsafe.Pointer(psCNG + 1376)) += int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(psDecCtrl + 16 + uintptr(i)*4))-(*OpusT_silk_CNG_struct)(unsafe.Pointer(psCNG)).FCNG_smth_Gain_Q16) * int64(int16(int32(CNG_GAIN_SMTH_Q16))) >> int32(16))
			/* If the smoothed gain is 3 dB greater than this subframe's gain, use this subframe's gain to adapt faster. */
			if int32(int64((*OpusT_silk_CNG_struct)(unsafe.Pointer(psCNG)).FCNG_smth_Gain_Q16)*int64(int32(CNG_GAIN_SMTH_THRESHOLD_Q16))>>int32(16)) > *(*OpusT_opus_int32)(unsafe.Pointer(psDecCtrl + 16 + uintptr(i)*4)) {
				(*OpusT_silk_CNG_struct)(unsafe.Pointer(psCNG)).FCNG_smth_Gain_Q16 = *(*OpusT_opus_int32)(unsafe.Pointer(psDecCtrl + 16 + uintptr(i)*4))
			}
			goto _8
		_8:
			;
			i = i + 1
		}
	}
	/* Add CNG when packet is lost or during DTX */
	if (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FlossCnt != 0 {
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			goto _10
		_10:
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v3 = st
		goto _12
	_12:
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v6 = libc.Xmalloc(tls, uint64(16))
			goto _14
		_14:
			st = v6
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v9 = st
		goto _16
	_16:
		*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v9)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v11 = libc.Xmalloc(tls, uint64(16))
			goto _18
		_18:
			st = v11
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v13 = st
		goto _20
	_20:
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v15 = libc.Xmalloc(tls, uint64(16))
			goto _22
		_22:
			st = v15
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v17 = st
		goto _24
	_24:
		if !(int64(int32(uint64(uint32(length+int32(MAX_LPC_ORDER)))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack)) {
			Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5763, int32(131))
		}
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v19 = libc.Xmalloc(tls, uint64(16))
			goto _26
		_26:
			st = v19
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v21 = st
		goto _28
	_28:
		*(*uintptr)(unsafe.Pointer(v21 + 8)) += uintptr(uint64(uint32(length+int32(MAX_LPC_ORDER))) * (uint64(4) / uint64(1)))
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v23 = libc.Xmalloc(tls, uint64(16))
			goto _30
		_30:
			st = v23
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v25 = st
		goto _32
	_32:
		CNG_sig_Q14 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v25)).Fglobal_stack - uintptr(uint64(uint32(length+int32(MAX_LPC_ORDER)))*(uint64(4)/uint64(1)))
		/* Generate CNG excitation */
		gain_Q16 = int32(int64((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FsPLC.FrandScale_Q14) * int64(*(*OpusT_opus_int32)(unsafe.Pointer(psDec + 4292 + 72 + 1*4))) >> int32(16))
		if gain_Q16 >= int32(1)<<int32(21) || (*OpusT_silk_CNG_struct)(unsafe.Pointer(psCNG)).FCNG_smth_Gain_Q16 > int32(1)<<int32(23) {
			gain_Q16 = gain_Q16 >> int32(16) * (gain_Q16 >> int32(16))
			gain_Q16 = (*OpusT_silk_CNG_struct)(unsafe.Pointer(psCNG)).FCNG_smth_Gain_Q16>>int32(16)*((*OpusT_silk_CNG_struct)(unsafe.Pointer(psCNG)).FCNG_smth_Gain_Q16>>int32(16)) - int32(uint32(gain_Q16)<<int32(5))
			v33 = gain_Q16
			if v33 <= int32(0) {
				v34 = 0
				goto _35
			}
			v36 = v33
			v37 = v36
			if v37 != 0 {
				v40 = int32(32) - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, uint32(v37)))
			} else {
				v40 = int32(32)
			}
			v38 = v40
			goto _39
		_39:
			lzeros = v38
			*(*OpusT_opus_int32)(unsafe.Pointer(bp)) = lzeros
			v41 = v36
			v42 = int32(24) - lzeros
			x = uint32(v41)
			r = uint32(v42)
			m = uint32(-v42)
			if v42 == int32(0) {
				v43 = v41
				goto _44
			} else {
				if v42 < int32(0) {
					v43 = int32(x<<m | x>>(uint32(32)-m))
					goto _44
				} else {
					v43 = int32(x<<(uint32(32)-r) | x>>r)
					goto _44
				}
			}
		_44:
			*(*OpusT_opus_int32)(unsafe.Pointer(bp + 4)) = v43 & int32(0x7f)
			if *(*OpusT_opus_int32)(unsafe.Pointer(bp))&int32(1) != 0 {
				y = int32(32768)
			} else {
				y = int32(46214)
			}
			y = y >> (*(*OpusT_opus_int32)(unsafe.Pointer(bp)) >> int32(1))
			y = int32(int64(y) + int64(y)*int64(int16(int32(int16(int32(213)))*int32(int16(*(*OpusT_opus_int32)(unsafe.Pointer(bp + 4))))))>>int32(16))
			v34 = y
			goto _35
		_35:
			gain_Q16 = int32(uint32(v34) << int32(16))
		} else {
			gain_Q16 = int32(int64(gain_Q16) * int64(gain_Q16) >> int32(16))
			gain_Q16 = int32(int64((*OpusT_silk_CNG_struct)(unsafe.Pointer(psCNG)).FCNG_smth_Gain_Q16)*int64((*OpusT_silk_CNG_struct)(unsafe.Pointer(psCNG)).FCNG_smth_Gain_Q16)>>int32(16)) - int32(uint32(gain_Q16)<<int32(5))
			v33 = gain_Q16
			if v33 <= int32(0) {
				v34 = 0
				goto _47
			}
			v36 = v33
			v37 = v36
			if v37 != 0 {
				v40 = int32(32) - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, uint32(v37)))
			} else {
				v40 = int32(32)
			}
			v38 = v40
			goto _51
		_51:
			lzeros = v38
			*(*OpusT_opus_int32)(unsafe.Pointer(bp)) = lzeros
			v41 = v36
			v42 = int32(24) - lzeros
			x = uint32(v41)
			r = uint32(v42)
			m = uint32(-v42)
			if v42 == int32(0) {
				v43 = v41
				goto _56
			} else {
				if v42 < int32(0) {
					v43 = int32(x<<m | x>>(uint32(32)-m))
					goto _56
				} else {
					v43 = int32(x<<(uint32(32)-r) | x>>r)
					goto _56
				}
			}
		_56:
			*(*OpusT_opus_int32)(unsafe.Pointer(bp + 4)) = v43 & int32(0x7f)
			if *(*OpusT_opus_int32)(unsafe.Pointer(bp))&int32(1) != 0 {
				y = int32(32768)
			} else {
				y = int32(46214)
			}
			y = y >> (*(*OpusT_opus_int32)(unsafe.Pointer(bp)) >> int32(1))
			y = int32(int64(y) + int64(y)*int64(int16(int32(int16(int32(213)))*int32(int16(*(*OpusT_opus_int32)(unsafe.Pointer(bp + 4))))))>>int32(16))
			v34 = y
			goto _47
		_47:
			gain_Q16 = int32(uint32(v34) << int32(8))
		}
		gain_Q10 = gain_Q16 >> int32(6)
		silk_CNG_exc(tls, CNG_sig_Q14+uintptr(MAX_LPC_ORDER)*4, psCNG, length, psCNG+1380)
		/* Convert CNG NLSF to filter representation */
		Opus_silk_NLSF2A(tls, bp+8, psCNG+1280, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Farch)
		/* Generate CNG signal, by synthesis filtering */
		libc.Xmemcpy(tls, CNG_sig_Q14, psCNG+1312, uint64(uint32(MAX_LPC_ORDER))*uint64(4))
		if !((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order == int32(10) || (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order == int32(16)) {
			Opus_celt_fatal(tls, __ccgo_ts+5777, __ccgo_ts+5763, int32(153))
		}
		i = 0
		for {
			if !(i < length) {
				break
			}
			/* Avoids introducing a bias because silk_SMLAWB() always rounds to -inf */
			LPC_pred_Q10 = (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order >> int32(1)
			LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(1))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 8)))[0])>>int32(16))
			LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(2))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 8)))[int32(1)])>>int32(16))
			LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(3))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 8)))[int32(2)])>>int32(16))
			LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(4))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 8)))[int32(3)])>>int32(16))
			LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(5))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 8)))[int32(4)])>>int32(16))
			LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(6))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 8)))[int32(5)])>>int32(16))
			LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(7))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 8)))[int32(6)])>>int32(16))
			LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(8))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 8)))[int32(7)])>>int32(16))
			LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(9))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 8)))[int32(8)])>>int32(16))
			LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(10))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 8)))[int32(9)])>>int32(16))
			if (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order == int32(16) {
				LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(11))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 8)))[int32(10)])>>int32(16))
				LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(12))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 8)))[int32(11)])>>int32(16))
				LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(13))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 8)))[int32(12)])>>int32(16))
				LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(14))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 8)))[int32(13)])>>int32(16))
				LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(15))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 8)))[int32(14)])>>int32(16))
				LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(16))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 8)))[int32(15)])>>int32(16))
			}
			/* Update states */
			if LPC_pred_Q10 > int32(silk_int32_MAX)>>int32(4) {
				v42 = int32(silk_int32_MAX) >> int32(4)
			} else {
				if LPC_pred_Q10 < int32(-2147483648)>>int32(4) {
					v52 = int32(-2147483648) >> int32(4)
				} else {
					v52 = LPC_pred_Q10
				}
				v42 = v52
			}
			if (uint32(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i)*4)))+uint32(int32(uint32(v42)<<int32(4))))&uint32(0x80000000) == uint32(0) {
				if LPC_pred_Q10 > int32(silk_int32_MAX)>>int32(4) {
					v58 = int32(silk_int32_MAX) >> int32(4)
				} else {
					if LPC_pred_Q10 < int32(-2147483648)>>int32(4) {
						v59 = int32(-2147483648) >> int32(4)
					} else {
						v59 = LPC_pred_Q10
					}
					v58 = v59
				}
				if uint32(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i)*4))&int32(uint32(v58)<<int32(4)))&uint32(0x80000000) != uint32(0) {
					v54 = int32(-2147483648)
				} else {
					if LPC_pred_Q10 > int32(silk_int32_MAX)>>int32(4) {
						v60 = int32(silk_int32_MAX) >> int32(4)
					} else {
						if LPC_pred_Q10 < int32(-2147483648)>>int32(4) {
							v61 = int32(-2147483648) >> int32(4)
						} else {
							v61 = LPC_pred_Q10
						}
						v60 = v61
					}
					v54 = *(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i)*4)) + int32(uint32(v60)<<int32(4))
				}
				v40 = v54
			} else {
				if LPC_pred_Q10 > int32(silk_int32_MAX)>>int32(4) {
					v63 = int32(silk_int32_MAX) >> int32(4)
				} else {
					if LPC_pred_Q10 < int32(-2147483648)>>int32(4) {
						v64 = int32(-2147483648) >> int32(4)
					} else {
						v64 = LPC_pred_Q10
					}
					v63 = v64
				}
				if uint32(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i)*4))|int32(uint32(v63)<<int32(4)))&uint32(0x80000000) == uint32(0) {
					v62 = int32(silk_int32_MAX)
				} else {
					if LPC_pred_Q10 > int32(silk_int32_MAX)>>int32(4) {
						v65 = int32(silk_int32_MAX) >> int32(4)
					} else {
						if LPC_pred_Q10 < int32(-2147483648)>>int32(4) {
							v66 = int32(-2147483648) >> int32(4)
						} else {
							v66 = LPC_pred_Q10
						}
						v65 = v66
					}
					v62 = *(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i)*4)) + int32(uint32(v65)<<int32(4))
				}
				v40 = v62
			}
			*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i)*4)) = v40
			/* Scale with Gain and add to input signal */
			if (int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i)*4)))*int64(gain_Q10)>>int32(16))>>(int32(8)-int32(1))+int32(1))>>int32(1) > int32(silk_int16_MAX1) {
				v42 = int32(silk_int16_MAX1)
			} else {
				if (int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i)*4)))*int64(gain_Q10)>>int32(16))>>(int32(8)-int32(1))+int32(1))>>int32(1) < int32(int16(-32768)) {
					v52 = int32(int16(-32768))
				} else {
					v52 = (int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i)*4)))*int64(gain_Q10)>>int32(16))>>(int32(8)-int32(1)) + int32(1)) >> int32(1)
				}
				v42 = v52
			}
			if int32(*(*OpusT_opus_int16)(unsafe.Pointer(frame + uintptr(i)*2)))+v42 > int32(silk_int16_MAX1) {
				v40 = int32(silk_int16_MAX1)
			} else {
				if (int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i)*4)))*int64(gain_Q10)>>int32(16))>>(int32(8)-int32(1))+int32(1))>>int32(1) > int32(silk_int16_MAX1) {
					v58 = int32(silk_int16_MAX1)
				} else {
					if (int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i)*4)))*int64(gain_Q10)>>int32(16))>>(int32(8)-int32(1))+int32(1))>>int32(1) < int32(int16(-32768)) {
						v59 = int32(int16(-32768))
					} else {
						v59 = (int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i)*4)))*int64(gain_Q10)>>int32(16))>>(int32(8)-int32(1)) + int32(1)) >> int32(1)
					}
					v58 = v59
				}
				if int32(*(*OpusT_opus_int16)(unsafe.Pointer(frame + uintptr(i)*2)))+v58 < int32(int16(-32768)) {
					v54 = int32(int16(-32768))
				} else {
					if (int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i)*4)))*int64(gain_Q10)>>int32(16))>>(int32(8)-int32(1))+int32(1))>>int32(1) > int32(silk_int16_MAX1) {
						v60 = int32(silk_int16_MAX1)
					} else {
						if (int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i)*4)))*int64(gain_Q10)>>int32(16))>>(int32(8)-int32(1))+int32(1))>>int32(1) < int32(int16(-32768)) {
							v61 = int32(int16(-32768))
						} else {
							v61 = (int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(CNG_sig_Q14 + uintptr(int32(MAX_LPC_ORDER)+i)*4)))*int64(gain_Q10)>>int32(16))>>(int32(8)-int32(1)) + int32(1)) >> int32(1)
						}
						v60 = v61
					}
					v54 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(frame + uintptr(i)*2))) + v60
				}
				v40 = v54
			}
			*(*OpusT_opus_int16)(unsafe.Pointer(frame + uintptr(i)*2)) = int16(v40)
			goto _57
		_57:
			;
			i = i + 1
		}
		libc.Xmemcpy(tls, psCNG+1312, CNG_sig_Q14+uintptr(length)*4, uint64(uint32(MAX_LPC_ORDER))*uint64(4))
	} else {
		libc.Xmemset(tls, psCNG+1312, 0, uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order))*uint64(4))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _80
	_80:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _82
_82:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
}

const silk_int16_MAX2 = 0x7FFF

/*#define silk_enc_map(a)                ((a) > 0 ? 1 : 0)*/
/*#define silk_dec_map(a)                ((a) > 0 ? 1 : -1)*/
/* shifting avoids if-statement */

// C documentation
//
//	/* Encodes signs of excitation */
