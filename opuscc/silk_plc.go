// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

package opuscc

import (
	"reflect"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

var _ reflect.Type
var _ unsafe.Pointer

func Opus_silk_PLC_Reset(tls *libc.TLS, psDec uintptr) {
	(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FsPLC.FpitchL_Q8 = int32(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length) << (int32(8) - int32(1)))
	*(*OpusT_opus_int32)(unsafe.Pointer(psDec + 4292 + 72)) = int32(65536)
	*(*OpusT_opus_int32)(unsafe.Pointer(psDec + 4292 + 72 + 1*4)) = int32(65536)
	(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FsPLC.Fsubfr_length = int32(20)
	(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FsPLC.Fnb_subfr = int32(2)
}

func Opus_silk_PLC(tls *libc.TLS, psDec uintptr, psDecCtrl uintptr, frame uintptr, lost int32, arch int32) {
	/* PLC control function */
	if (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Ffs_kHz != (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FsPLC.Ffs_kHz {
		Opus_silk_PLC_Reset(tls, psDec)
		(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FsPLC.Ffs_kHz = (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Ffs_kHz
	}
	if lost != 0 {
		/****************************/
		/* Generate Signal          */
		/****************************/
		silk_PLC_conceal(tls, psDec, psDecCtrl, frame, arch)
		(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FlossCnt = (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FlossCnt + 1
	} else {
		/****************************/
		/* Update state             */
		/****************************/
		silk_PLC_update(tls, psDec, psDecCtrl)
	}
}

// C documentation
//
//	/**************************************************/
//	/* Update state of PLC                            */
//	/**************************************************/
func silk_PLC_update(tls *libc.TLS, psDec uintptr, psDecCtrl uintptr) {
	var LTP_Gain_Q14, temp_LTP_Gain_Q14, tmp, tmp1 OpusT_opus_int32
	var i, j, scale_Q10, scale_Q14, v3 int32
	var psPLC uintptr
	_, _, _, _, _, _, _, _, _, _ = LTP_Gain_Q14, i, j, psPLC, scale_Q10, scale_Q14, temp_LTP_Gain_Q14, tmp, tmp1, v3
	psPLC = psDec + 4292
	/* Update parameters used in case of packet loss */
	(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FprevSignalType = int32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FsignalType)
	LTP_Gain_Q14 = 0
	if int32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FsignalType) == int32(TYPE_VOICED) {
		/* Find the parameters for the last subframe which contains a pitch pulse */
		j = 0
		for {
			if !(j*(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length < *(*int32)(unsafe.Pointer(psDecCtrl + uintptr((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr-int32(1))*4))) {
				break
			}
			if j == (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr {
				break
			}
			temp_LTP_Gain_Q14 = 0
			i = 0
			for {
				if !(i < int32(LTP_ORDER)) {
					break
				}
				temp_LTP_Gain_Q14 = temp_LTP_Gain_Q14 + int32(*(*OpusT_opus_int16)(unsafe.Pointer(psDecCtrl + 96 + uintptr(((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr-int32(1)-j)*int32(LTP_ORDER)+i)*2)))
				i = i + 1
			}
			if temp_LTP_Gain_Q14 > LTP_Gain_Q14 {
				LTP_Gain_Q14 = temp_LTP_Gain_Q14
				libc.Xmemcpy(tls, psPLC+4, psDecCtrl+96+uintptr(int32(int16((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr-int32(1)-j))*int32(int16(int32(LTP_ORDER))))*2, uint64(uint32(LTP_ORDER))*uint64(2))
				(*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).FpitchL_Q8 = int32(uint32(*(*int32)(unsafe.Pointer(psDecCtrl + uintptr((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr-int32(1)-j)*4))) << int32(8))
			}
			j = j + 1
		}
		libc.Xmemset(tls, psPLC+4, 0, uint64(uint32(LTP_ORDER))*uint64(2))
		*(*OpusT_opus_int16)(unsafe.Pointer(psPLC + 4 + uintptr(int32(LTP_ORDER)/int32(2))*2)) = int16(LTP_Gain_Q14)
		/* Limit LT coefs */
		if LTP_Gain_Q14 < int32(V_PITCH_GAIN_START_MIN_Q14) {
			tmp = int32(uint32(int32(V_PITCH_GAIN_START_MIN_Q14)) << int32(10))
			if LTP_Gain_Q14 > int32(1) {
				v3 = LTP_Gain_Q14
			} else {
				v3 = int32(1)
			}
			scale_Q10 = tmp / v3
			i = 0
			for {
				if !(i < int32(LTP_ORDER)) {
					break
				}
				*(*OpusT_opus_int16)(unsafe.Pointer(psPLC + 4 + uintptr(i)*2)) = int16(int32(*(*OpusT_opus_int16)(unsafe.Pointer(psPLC + 4 + uintptr(i)*2))) * int32(int16(scale_Q10)) >> int32(10))
				i = i + 1
			}
		} else {
			if LTP_Gain_Q14 > int32(V_PITCH_GAIN_START_MAX_Q14) {
				tmp1 = int32(uint32(int32(V_PITCH_GAIN_START_MAX_Q14)) << int32(14))
				if LTP_Gain_Q14 > int32(1) {
					v3 = LTP_Gain_Q14
				} else {
					v3 = int32(1)
				}
				scale_Q14 = tmp1 / v3
				i = 0
				for {
					if !(i < int32(LTP_ORDER)) {
						break
					}
					*(*OpusT_opus_int16)(unsafe.Pointer(psPLC + 4 + uintptr(i)*2)) = int16(int32(*(*OpusT_opus_int16)(unsafe.Pointer(psPLC + 4 + uintptr(i)*2))) * int32(int16(scale_Q14)) >> int32(14))
					i = i + 1
				}
			}
		}
	} else {
		(*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).FpitchL_Q8 = int32(uint32(int32(int16((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Ffs_kHz))*int32(int16(int32(18)))) << int32(8))
		libc.Xmemset(tls, psPLC+4, 0, uint64(uint32(LTP_ORDER))*uint64(2))
	}
	/* Save LPC coefficients */
	libc.Xmemcpy(tls, psPLC+14, psDecCtrl+32+1*32, uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order))*uint64(2))
	(*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).FprevLTP_scale_Q14 = int16((*OpusT_silk_decoder_control)(unsafe.Pointer(psDecCtrl)).FLTP_scale_Q14)
	/* Save last two gains */
	libc.Xmemcpy(tls, psPLC+72, psDecCtrl+16+uintptr((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr-int32(2))*4, uint64(uint32(2))*uint64(4))
	(*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).Fsubfr_length = (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length
	(*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).Fnb_subfr = (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr
}

func silk_PLC_energy(tls *libc.TLS, energy1 uintptr, shift1 uintptr, energy2 uintptr, shift2 uintptr, exc_Q14 uintptr, prevGain_Q10 uintptr, subfr_length int32, nb_subfr int32) {
	var _saved_stack, exc_buf, exc_buf_ptr, st, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var i, k, v31, v32 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = _saved_stack, exc_buf, exc_buf_ptr, i, k, st, v1, v11, v13, v15, v17, v19, v21, v23, v3, v31, v32, v5, v7, v9
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	_saved_stack = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v7 = st
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(2)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (uint64(uint32(2)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v11 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v15 = st
	if !(int64(int32(uint64(uint32(int32(2)*subfr_length))*(uint64(2)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+6715, int32(199))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v19 = st
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32(int32(2)*subfr_length)) * (uint64(2) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v23 = st
	exc_buf = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32(int32(2)*subfr_length))*(uint64(2)/uint64(1)))
	/* Find random noise component */
	/* Scale previous excitation signal */
	exc_buf_ptr = exc_buf
	k = 0
	for {
		if !(k < int32(2)) {
			break
		}
		i = 0
		for {
			if !(i < subfr_length) {
				break
			}
			if int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(exc_Q14 + uintptr(i+(k+nb_subfr-int32(2))*subfr_length)*4)))*int64(*(*OpusT_opus_int32)(unsafe.Pointer(prevGain_Q10 + uintptr(k)*4)))>>int32(16))>>int32(8) > int32(silk_int16_MAX7) {
				v31 = int32(silk_int16_MAX7)
			} else {
				if int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(exc_Q14 + uintptr(i+(k+nb_subfr-int32(2))*subfr_length)*4)))*int64(*(*OpusT_opus_int32)(unsafe.Pointer(prevGain_Q10 + uintptr(k)*4)))>>int32(16))>>int32(8) < int32(int16(-32768)) {
					v32 = int32(int16(-32768))
				} else {
					v32 = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(exc_Q14 + uintptr(i+(k+nb_subfr-int32(2))*subfr_length)*4)))*int64(*(*OpusT_opus_int32)(unsafe.Pointer(prevGain_Q10 + uintptr(k)*4)))>>int32(16)) >> int32(8)
				}
				v31 = v32
			}
			*(*OpusT_opus_int16)(unsafe.Pointer(exc_buf_ptr + uintptr(i)*2)) = int16(v31)
			i = i + 1
		}
		exc_buf_ptr = exc_buf_ptr + uintptr(subfr_length)*2
		k = k + 1
	}
	/* Find the subframe with lowest energy of the last two and use that as random noise generator */
	Opus_silk_sum_sqr_shift(tls, energy1, shift1, exc_buf, subfr_length)
	Opus_silk_sum_sqr_shift(tls, energy2, shift2, exc_buf+uintptr(subfr_length)*2, subfr_length)
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
}

func silk_PLC_conceal(tls *libc.TLS, psDec uintptr, psDecCtrl uintptr, frame uintptr, arch int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var B_Q14, _saved_stack, pred_lag_ptr, psPLC, rand_ptr, sLPC_Q14_ptr, sLTP, sLTP_Q14, st, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var LPC_pred_Q10, LTP_pred_Q12, b32_inv, b32_nrm, down_scale_Q30, err_Q32, harm_Gain_Q15, invGain_Q30, inv_gain_Q30, rand_Gain_Q15, rand_seed, result, v84, v85, v86, v89 OpusT_opus_int32
	var b_headrm, i, idx, j, k, lag, lshift, sLTP_buf_idx, v53, v54, v55, v57, v58, v59, v60, v62, v63, v64, v65, v67, v68 int32
	var rand_scale_Q14, v79, v80, v81 OpusT_opus_int16
	var _ /* A_Q12 at bp+16 */ [16]OpusT_opus_int16
	var _ /* energy1 at bp+8 */ OpusT_opus_int32
	var _ /* energy2 at bp+12 */ OpusT_opus_int32
	var _ /* prevGain_Q10 at bp+48 */ [2]OpusT_opus_int32
	var _ /* shift1 at bp+0 */ int32
	var _ /* shift2 at bp+4 */ int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = B_Q14, LPC_pred_Q10, LTP_pred_Q12, _saved_stack, b32_inv, b32_nrm, b_headrm, down_scale_Q30, err_Q32, harm_Gain_Q15, i, idx, invGain_Q30, inv_gain_Q30, j, k, lag, lshift, pred_lag_ptr, psPLC, rand_Gain_Q15, rand_ptr, rand_scale_Q14, rand_seed, result, sLPC_Q14_ptr, sLTP, sLTP_Q14, sLTP_buf_idx, st, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v53, v54, v55, v57, v58, v59, v60, v62, v63, v64, v65, v67, v68, v7, v79, v80, v81, v84, v85, v86, v89, v9
	psPLC = psDec + 4292
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	_saved_stack = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v7 = st
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v11 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v15 = st
	if !(int64(int32(uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length+(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+6715, int32(245))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v19 = st
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length+(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v23 = st
	sLTP_Q14 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length+(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length))*(uint64(4)/uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v7 = st
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(2)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (uint64(uint32(2)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v11 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v15 = st
	if !(int64(int32(uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length))*(uint64(2)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+6715, int32(250))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v19 = st
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length)) * (uint64(2) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v23 = st
	sLTP = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length))*(uint64(2)/uint64(1)))
	(*(*[2]OpusT_opus_int32)(unsafe.Pointer(bp + 48)))[0] = *(*OpusT_opus_int32)(unsafe.Pointer(psPLC + 72)) >> int32(6)
	(*(*[2]OpusT_opus_int32)(unsafe.Pointer(bp + 48)))[int32(1)] = *(*OpusT_opus_int32)(unsafe.Pointer(psPLC + 72 + 1*4)) >> int32(6)
	if (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Ffirst_frame_after_reset != 0 {
		libc.Xmemset(tls, psPLC+14, 0, uint64(32))
	}
	silk_PLC_energy(tls, bp+8, bp, bp+12, bp+4, psDec+4, bp+48, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr)
	if *(*OpusT_opus_int32)(unsafe.Pointer(bp + 8))>>*(*int32)(unsafe.Pointer(bp + 4)) < *(*OpusT_opus_int32)(unsafe.Pointer(bp + 12))>>*(*int32)(unsafe.Pointer(bp)) {
		/* First sub-frame has lowest energy */
		v53 = 0
		v54 = ((*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).Fnb_subfr-int32(1))*(*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).Fsubfr_length - int32(RAND_BUF_SIZE)
		if v53 > v54 {
			v57 = v53
		} else {
			v57 = v54
		}
		v55 = v57
		rand_ptr = psDec + 4 + uintptr(v55)*4
	} else {
		/* Second sub-frame has lowest energy */
		v53 = 0
		v54 = (*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).Fnb_subfr*(*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).Fsubfr_length - int32(RAND_BUF_SIZE)
		if v53 > v54 {
			v57 = v53
		} else {
			v57 = v54
		}
		v55 = v57
		rand_ptr = psDec + 4 + uintptr(v55)*4
	}
	/* Set up Gain to random noise component */
	B_Q14 = psPLC + 4
	rand_scale_Q14 = (*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).FrandScale_Q14
	/* Set up attenuation gains */
	v53 = int32(NB_ATT) - int32(1)
	v54 = (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FlossCnt
	if v53 < v54 {
		v57 = v53
	} else {
		v57 = v54
	}
	v55 = v57
	harm_Gain_Q15 = int32(HARM_ATT_Q15[v55])
	if (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FprevSignalType == int32(TYPE_VOICED) {
		v53 = int32(NB_ATT) - int32(1)
		v54 = (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FlossCnt
		if v53 < v54 {
			v57 = v53
		} else {
			v57 = v54
		}
		v55 = v57
		rand_Gain_Q15 = int32(PLC_RAND_ATTENUATE_V_Q15[v55])
	} else {
		v53 = int32(NB_ATT) - int32(1)
		v54 = (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FlossCnt
		if v53 < v54 {
			v57 = v53
		} else {
			v57 = v54
		}
		v55 = v57
		rand_Gain_Q15 = int32(PLC_RAND_ATTENUATE_UV_Q15[v55])
	}
	/* LPC concealment. Apply BWE to previous LPC */
	Opus_silk_bwexpander(tls, psPLC+14, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order, int32(64881))
	/* Preload LPC coefficients to array on stack. Gives small performance gain */
	libc.Xmemcpy(tls, bp+16, psPLC+14, uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order))*uint64(2))
	/* First Lost frame */
	if (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FlossCnt == 0 {
		rand_scale_Q14 = int16(int32(1) << int32(14))
		/* Reduce random noise Gain for voiced frames */
		if (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FprevSignalType == int32(TYPE_VOICED) {
			i = 0
			for {
				if !(i < int32(LTP_ORDER)) {
					break
				}
				rand_scale_Q14 = int16(int32(rand_scale_Q14) - int32(*(*OpusT_opus_int16)(unsafe.Pointer(B_Q14 + uintptr(i)*2))))
				i = i + 1
			}
			v79 = int16(3277)
			v80 = rand_scale_Q14
			if int32(v79) > int32(v80) {
				v53 = int32(v79)
			} else {
				v53 = int32(v80)
			}
			v81 = int16(v53)
			rand_scale_Q14 = v81 /* 0.2 */
			rand_scale_Q14 = int16(int32(rand_scale_Q14) * int32((*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).FprevLTP_scale_Q14) >> int32(14))
		} else {
			_ = arch
			invGain_Q30 = Opus_silk_LPC_inverse_pred_gain_c(tls, psPLC+14, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order)
			v84 = int32(1) << int32(30) >> int32(LOG2_INV_LPC_GAIN_HIGH_THRES)
			v85 = invGain_Q30
			if v84 < v85 {
				v53 = v84
			} else {
				v53 = v85
			}
			v86 = v53
			down_scale_Q30 = v86
			v84 = int32(1) << int32(30) >> int32(LOG2_INV_LPC_GAIN_LOW_THRES)
			v85 = down_scale_Q30
			if v84 > v85 {
				v53 = v84
			} else {
				v53 = v85
			}
			v86 = v53
			down_scale_Q30 = v86
			down_scale_Q30 = int32(uint32(down_scale_Q30) << int32(LOG2_INV_LPC_GAIN_HIGH_THRES))
			rand_Gain_Q15 = int32(int64(down_scale_Q30)*int64(int16(rand_Gain_Q15))>>int32(16)) >> int32(14)
		}
	}
	rand_seed = (*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).Frand_seed
	lag = ((*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).FpitchL_Q8>>(int32(8)-int32(1)) + int32(1)) >> int32(1)
	sLTP_buf_idx = (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length
	/* Rewhiten LTP state */
	idx = (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length - lag - (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order - int32(LTP_ORDER)/int32(2)
	if !(idx > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+6729, __ccgo_ts+6715, int32(319))
	}
	Opus_silk_LPC_analysis_filter(tls, sLTP+uintptr(idx)*2, psDec+1348+uintptr(idx)*2, bp+16, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length-idx, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order, arch)
	/* Scale LTP state */
	v84 = *(*OpusT_opus_int32)(unsafe.Pointer(psPLC + 72 + 1*4))
	v53 = int32(46)
	_ = v84 != int32(0)
	_ = v53 > int32(0)
	if v84 > 0 {
		v54 = v84
	} else {
		v54 = -v84
	}
	v85 = v54
	if v85 != 0 {
		v55 = int32(32) - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, uint32(v85)))
	} else {
		v55 = int32(32)
	}
	v86 = v55
	b_headrm = v86 - int32(1)
	b32_nrm = int32(uint32(v84) << b_headrm)
	b32_inv = int32(silk_int32_MAX) >> int32(2) / (b32_nrm >> int32(16))
	result = int32(uint32(b32_inv) << int32(16))
	err_Q32 = int32(uint32(int32(1)<<int32(29)-int32(int64(b32_nrm)*int64(int16(b32_inv))>>int32(16))) << int32(3))
	result = int32(int64(result) + int64(err_Q32)*int64(b32_inv)>>int32(16))
	lshift = int32(61) - b_headrm - v53
	if lshift <= int32(0) {
		if int32(-2147483648)>>-lshift > int32(silk_int32_MAX)>>-lshift {
			if result > int32(-2147483648)>>-lshift {
				v58 = int32(-2147483648) >> -lshift
			} else {
				if result < int32(silk_int32_MAX)>>-lshift {
					v59 = int32(silk_int32_MAX) >> -lshift
				} else {
					v59 = result
				}
				v58 = v59
			}
			v57 = v58
		} else {
			if result > int32(silk_int32_MAX)>>-lshift {
				v60 = int32(silk_int32_MAX) >> -lshift
			} else {
				if result < int32(-2147483648)>>-lshift {
					v62 = int32(-2147483648) >> -lshift
				} else {
					v62 = result
				}
				v60 = v62
			}
			v57 = v60
		}
		v89 = int32(uint32(v57) << -lshift)
		goto _102
	} else {
		if lshift < int32(32) {
			v89 = result >> lshift
			goto _102
		} else {
			v89 = 0
			goto _102
		}
	}
_102:
	inv_gain_Q30 = v89
	if inv_gain_Q30 < int32(silk_int32_MAX)>>int32(1) {
		v53 = inv_gain_Q30
	} else {
		v53 = int32(silk_int32_MAX) >> int32(1)
	}
	inv_gain_Q30 = v53
	i = idx + (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order
	for {
		if !(i < (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length) {
			break
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(sLTP_Q14 + uintptr(i)*4)) = int32(int64(inv_gain_Q30) * int64(*(*OpusT_opus_int16)(unsafe.Pointer(sLTP + uintptr(i)*2))) >> int32(16))
		i = i + 1
	}
	/***************************/
	/* LTP synthesis filtering */
	/***************************/
	k = 0
	for {
		if !(k < (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr) {
			break
		}
		/* Set up pointer */
		pred_lag_ptr = sLTP_Q14 + uintptr(sLTP_buf_idx-lag+int32(LTP_ORDER)/int32(2))*4
		i = 0
		for {
			if !(i < (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length) {
				break
			}
			/* Unrolled loop */
			/* Avoids introducing a bias because silk_SMLAWB() always rounds to -inf */
			LTP_pred_Q12 = int32(2)
			LTP_pred_Q12 = int32(int64(LTP_pred_Q12) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(pred_lag_ptr)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(B_Q14)))>>int32(16))
			LTP_pred_Q12 = int32(int64(LTP_pred_Q12) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(pred_lag_ptr - uintptr(1)*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(B_Q14 + 1*2)))>>int32(16))
			LTP_pred_Q12 = int32(int64(LTP_pred_Q12) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(pred_lag_ptr - uintptr(2)*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(B_Q14 + 2*2)))>>int32(16))
			LTP_pred_Q12 = int32(int64(LTP_pred_Q12) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(pred_lag_ptr - uintptr(3)*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(B_Q14 + 3*2)))>>int32(16))
			LTP_pred_Q12 = int32(int64(LTP_pred_Q12) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(pred_lag_ptr - uintptr(4)*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(B_Q14 + 4*2)))>>int32(16))
			pred_lag_ptr += 4
			/* Generate LPC excitation */
			rand_seed = int32(uint32(int32(RAND_INCREMENT)) + uint32(rand_seed)*uint32(int32(RAND_MULTIPLIER)))
			idx = rand_seed >> int32(25) & (int32(RAND_BUF_SIZE) - int32(1))
			*(*OpusT_opus_int32)(unsafe.Pointer(sLTP_Q14 + uintptr(sLTP_buf_idx)*4)) = int32(uint32(int32(int64(LTP_pred_Q12)+int64(*(*OpusT_opus_int32)(unsafe.Pointer(rand_ptr + uintptr(idx)*4)))*int64(rand_scale_Q14)>>int32(16))) << int32(2))
			sLTP_buf_idx = sLTP_buf_idx + 1
			i = i + 1
		}
		/* Gradually reduce LTP gain */
		j = 0
		for {
			if !(j < int32(LTP_ORDER)) {
				break
			}
			*(*OpusT_opus_int16)(unsafe.Pointer(B_Q14 + uintptr(j)*2)) = int16(int32(int16(harm_Gain_Q15)) * int32(*(*OpusT_opus_int16)(unsafe.Pointer(B_Q14 + uintptr(j)*2))) >> int32(15))
			j = j + 1
		}
		/* Gradually reduce excitation gain */
		rand_scale_Q14 = int16(int32(rand_scale_Q14) * int32(int16(rand_Gain_Q15)) >> int32(15))
		/* Slowly increase pitch lag */
		(*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).FpitchL_Q8 = int32(int64((*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).FpitchL_Q8) + int64((*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).FpitchL_Q8)*int64(int16(int32(PITCH_DRIFT_FAC_Q16)))>>int32(16))
		v84 = (*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).FpitchL_Q8
		v85 = int32(uint32(int32(int16(int32(MAX_PITCH_LAG_MS)))*int32(int16((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Ffs_kHz))) << int32(8))
		if v84 < v85 {
			v53 = v84
		} else {
			v53 = v85
		}
		v86 = v53
		(*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).FpitchL_Q8 = v86
		lag = ((*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).FpitchL_Q8>>(int32(8)-int32(1)) + int32(1)) >> int32(1)
		k = k + 1
	}
	/***************************/
	/* LPC synthesis filtering */
	/***************************/
	sLPC_Q14_ptr = sLTP_Q14 + uintptr((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length-int32(MAX_LPC_ORDER))*4
	/* Copy LPC state */
	libc.Xmemcpy(tls, sLPC_Q14_ptr, psDec+1284, uint64(uint32(MAX_LPC_ORDER))*uint64(4))
	if !((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order >= int32(10)) {
		Opus_celt_fatal(tls, __ccgo_ts+6755, __ccgo_ts+6715, int32(373))
	} /* check that unrolling works */
	i = 0
	for {
		if !(i < (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length) {
			break
		}
		/* partly unrolled */
		/* Avoids introducing a bias because silk_SMLAWB() always rounds to -inf */
		LPC_pred_Q10 = (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order >> int32(1)
		LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i-int32(1))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 16)))[0])>>int32(16))
		LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i-int32(2))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 16)))[int32(1)])>>int32(16))
		LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i-int32(3))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 16)))[int32(2)])>>int32(16))
		LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i-int32(4))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 16)))[int32(3)])>>int32(16))
		LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i-int32(5))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 16)))[int32(4)])>>int32(16))
		LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i-int32(6))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 16)))[int32(5)])>>int32(16))
		LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i-int32(7))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 16)))[int32(6)])>>int32(16))
		LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i-int32(8))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 16)))[int32(7)])>>int32(16))
		LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i-int32(9))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 16)))[int32(8)])>>int32(16))
		LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i-int32(10))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 16)))[int32(9)])>>int32(16))
		j = int32(10)
		for {
			if !(j < (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order) {
				break
			}
			LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i-j-int32(1))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 16)))[j])>>int32(16))
			j = j + 1
		}
		/* Add prediction to LPC excitation */
		if LPC_pred_Q10 > int32(silk_int32_MAX)>>int32(4) {
			v54 = int32(silk_int32_MAX) >> int32(4)
		} else {
			if LPC_pred_Q10 < int32(-2147483648)>>int32(4) {
				v55 = int32(-2147483648) >> int32(4)
			} else {
				v55 = LPC_pred_Q10
			}
			v54 = v55
		}
		if (uint32(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i)*4)))+uint32(int32(uint32(v54)<<int32(4))))&uint32(0x80000000) == uint32(0) {
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
			if uint32(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i)*4))&int32(uint32(v58)<<int32(4)))&uint32(0x80000000) != uint32(0) {
				v57 = int32(-2147483648)
			} else {
				if LPC_pred_Q10 > int32(silk_int32_MAX)>>int32(4) {
					v60 = int32(silk_int32_MAX) >> int32(4)
				} else {
					if LPC_pred_Q10 < int32(-2147483648)>>int32(4) {
						v62 = int32(-2147483648) >> int32(4)
					} else {
						v62 = LPC_pred_Q10
					}
					v60 = v62
				}
				v57 = *(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i)*4)) + int32(uint32(v60)<<int32(4))
			}
			v53 = v57
		} else {
			if LPC_pred_Q10 > int32(silk_int32_MAX)>>int32(4) {
				v64 = int32(silk_int32_MAX) >> int32(4)
			} else {
				if LPC_pred_Q10 < int32(-2147483648)>>int32(4) {
					v65 = int32(-2147483648) >> int32(4)
				} else {
					v65 = LPC_pred_Q10
				}
				v64 = v65
			}
			if uint32(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i)*4))|int32(uint32(v64)<<int32(4)))&uint32(0x80000000) == uint32(0) {
				v63 = int32(silk_int32_MAX)
			} else {
				if LPC_pred_Q10 > int32(silk_int32_MAX)>>int32(4) {
					v67 = int32(silk_int32_MAX) >> int32(4)
				} else {
					if LPC_pred_Q10 < int32(-2147483648)>>int32(4) {
						v68 = int32(-2147483648) >> int32(4)
					} else {
						v68 = LPC_pred_Q10
					}
					v67 = v68
				}
				v63 = *(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i)*4)) + int32(uint32(v67)<<int32(4))
			}
			v53 = v63
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i)*4)) = v53
		/* Scale with Gain */
		if (int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i)*4)))*int64((*(*[2]OpusT_opus_int32)(unsafe.Pointer(bp + 48)))[int32(1)])>>int32(16))>>(int32(8)-int32(1))+int32(1))>>int32(1) > int32(silk_int16_MAX7) {
			v54 = int32(silk_int16_MAX7)
		} else {
			if (int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i)*4)))*int64((*(*[2]OpusT_opus_int32)(unsafe.Pointer(bp + 48)))[int32(1)])>>int32(16))>>(int32(8)-int32(1))+int32(1))>>int32(1) < int32(int16(-32768)) {
				v55 = int32(int16(-32768))
			} else {
				v55 = (int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i)*4)))*int64((*(*[2]OpusT_opus_int32)(unsafe.Pointer(bp + 48)))[int32(1)])>>int32(16))>>(int32(8)-int32(1)) + int32(1)) >> int32(1)
			}
			v54 = v55
		}
		if v54 > int32(silk_int16_MAX7) {
			v53 = int32(silk_int16_MAX7)
		} else {
			if (int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i)*4)))*int64((*(*[2]OpusT_opus_int32)(unsafe.Pointer(bp + 48)))[int32(1)])>>int32(16))>>(int32(8)-int32(1))+int32(1))>>int32(1) > int32(silk_int16_MAX7) {
				v58 = int32(silk_int16_MAX7)
			} else {
				if (int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i)*4)))*int64((*(*[2]OpusT_opus_int32)(unsafe.Pointer(bp + 48)))[int32(1)])>>int32(16))>>(int32(8)-int32(1))+int32(1))>>int32(1) < int32(int16(-32768)) {
					v59 = int32(int16(-32768))
				} else {
					v59 = (int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i)*4)))*int64((*(*[2]OpusT_opus_int32)(unsafe.Pointer(bp + 48)))[int32(1)])>>int32(16))>>(int32(8)-int32(1)) + int32(1)) >> int32(1)
				}
				v58 = v59
			}
			if v58 < int32(int16(-32768)) {
				v57 = int32(int16(-32768))
			} else {
				if (int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i)*4)))*int64((*(*[2]OpusT_opus_int32)(unsafe.Pointer(bp + 48)))[int32(1)])>>int32(16))>>(int32(8)-int32(1))+int32(1))>>int32(1) > int32(silk_int16_MAX7) {
					v60 = int32(silk_int16_MAX7)
				} else {
					if (int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i)*4)))*int64((*(*[2]OpusT_opus_int32)(unsafe.Pointer(bp + 48)))[int32(1)])>>int32(16))>>(int32(8)-int32(1))+int32(1))>>int32(1) < int32(int16(-32768)) {
						v62 = int32(int16(-32768))
					} else {
						v62 = (int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14_ptr + uintptr(int32(MAX_LPC_ORDER)+i)*4)))*int64((*(*[2]OpusT_opus_int32)(unsafe.Pointer(bp + 48)))[int32(1)])>>int32(16))>>(int32(8)-int32(1)) + int32(1)) >> int32(1)
					}
					v60 = v62
				}
				v57 = v60
			}
			v53 = v57
		}
		*(*OpusT_opus_int16)(unsafe.Pointer(frame + uintptr(i)*2)) = int16(v53)
		i = i + 1
	}
	/* Save LPC state */
	libc.Xmemcpy(tls, psDec+1284, sLPC_Q14_ptr+uintptr((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length)*4, uint64(uint32(MAX_LPC_ORDER))*uint64(4))
	/**************************************/
	/* Update states                      */
	/**************************************/
	(*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).Frand_seed = rand_seed
	(*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).FrandScale_Q14 = rand_scale_Q14
	i = 0
	for {
		if !(i < int32(MAX_NB_SUBFR)) {
			break
		}
		*(*int32)(unsafe.Pointer(psDecCtrl + uintptr(i)*4)) = lag
		i = i + 1
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
}

// C documentation
//
//	/* Glues concealed frames with new good received frames */
func Opus_silk_PLC_glue_frames(tls *libc.TLS, psDec uintptr, frame uintptr, length int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var LZ, frac_Q24, gain_Q16, lzeros, slope_Q16, y, v1, v11, v12, v2, v5, v6, v7 OpusT_opus_int32
	var i, v4, v9 int32
	var m, r, x OpusT_opus_uint32
	var psPLC uintptr
	var _ /* energy at bp+12 */ OpusT_opus_int32
	var _ /* energy_shift at bp+8 */ int32
	var _ /* frac_Q7 at bp+4 */ OpusT_opus_int32
	var _ /* lz at bp+0 */ OpusT_opus_int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = LZ, frac_Q24, gain_Q16, i, lzeros, m, psPLC, r, slope_Q16, x, y, v1, v11, v12, v2, v4, v5, v6, v7, v9
	psPLC = psDec + 4292
	if (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FlossCnt != 0 {
		/* Calculate energy in concealed residual */
		Opus_silk_sum_sqr_shift(tls, psPLC+60, psPLC+64, frame, length)
		(*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).Flast_frame_lost = int32(1)
	} else {
		if (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FsPLC.Flast_frame_lost != 0 {
			/* Calculate residual in decoded signal if last frame was lost */
			Opus_silk_sum_sqr_shift(tls, bp+12, bp+8, frame, length)
			/* Normalize energies */
			if *(*int32)(unsafe.Pointer(bp + 8)) > (*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).Fconc_energy_shift {
				(*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).Fconc_energy = (*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).Fconc_energy >> (*(*int32)(unsafe.Pointer(bp + 8)) - (*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).Fconc_energy_shift)
			} else {
				if *(*int32)(unsafe.Pointer(bp + 8)) < (*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).Fconc_energy_shift {
					*(*OpusT_opus_int32)(unsafe.Pointer(bp + 12)) = *(*OpusT_opus_int32)(unsafe.Pointer(bp + 12)) >> ((*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).Fconc_energy_shift - *(*int32)(unsafe.Pointer(bp + 8)))
				}
			}
			/* Fade in the energy difference */
			if *(*OpusT_opus_int32)(unsafe.Pointer(bp + 12)) > (*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).Fconc_energy {
				v1 = (*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).Fconc_energy
				if v1 != 0 {
					v4 = int32(32) - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, uint32(v1)))
				} else {
					v4 = int32(32)
				}
				v2 = v4
				LZ = v2
				LZ = LZ - int32(1)
				(*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).Fconc_energy = int32(uint32((*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).Fconc_energy) << LZ)
				v1 = int32(24) - LZ
				v2 = 0
				if v1 > v2 {
					v4 = v1
				} else {
					v4 = v2
				}
				v5 = v4
				*(*OpusT_opus_int32)(unsafe.Pointer(bp + 12)) = *(*OpusT_opus_int32)(unsafe.Pointer(bp + 12)) >> v5
				if *(*OpusT_opus_int32)(unsafe.Pointer(bp + 12)) > int32(1) {
					v4 = *(*OpusT_opus_int32)(unsafe.Pointer(bp + 12))
				} else {
					v4 = int32(1)
				}
				frac_Q24 = (*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).Fconc_energy / v4
				v1 = frac_Q24
				if v1 <= int32(0) {
					v2 = 0
					goto _13
				}
				v5 = v1
				v6 = v5
				if v6 != 0 {
					v4 = int32(32) - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, uint32(v6)))
				} else {
					v4 = int32(32)
				}
				v7 = v4
				lzeros = v7
				*(*OpusT_opus_int32)(unsafe.Pointer(bp)) = lzeros
				v11 = v5
				v9 = int32(24) - lzeros
				x = uint32(v11)
				r = uint32(v9)
				m = uint32(-v9)
				if v9 == int32(0) {
					v12 = v11
					goto _22
				} else {
					if v9 < int32(0) {
						v12 = int32(x<<m | x>>(uint32(32)-m))
						goto _22
					} else {
						v12 = int32(x<<(uint32(32)-r) | x>>r)
						goto _22
					}
				}
			_22:
				*(*OpusT_opus_int32)(unsafe.Pointer(bp + 4)) = v12 & int32(0x7f)
				if *(*OpusT_opus_int32)(unsafe.Pointer(bp))&int32(1) != 0 {
					y = int32(32768)
				} else {
					y = int32(46214)
				}
				y = y >> (*(*OpusT_opus_int32)(unsafe.Pointer(bp)) >> int32(1))
				y = int32(int64(y) + int64(y)*int64(int16(int32(int16(int32(213)))*int32(int16(*(*OpusT_opus_int32)(unsafe.Pointer(bp + 4))))))>>int32(16))
				v2 = y
			_13:
				gain_Q16 = int32(uint32(v2) << int32(4))
				slope_Q16 = (int32(1)<<int32(16) - gain_Q16) / length
				/* Make slope 4x steeper to avoid missing onsets after DTX */
				slope_Q16 = int32(uint32(slope_Q16) << int32(2))
				i = 0
				for {
					if !(i < length) {
						break
					}
					*(*OpusT_opus_int16)(unsafe.Pointer(frame + uintptr(i)*2)) = int16(int32(int64(gain_Q16) * int64(*(*OpusT_opus_int16)(unsafe.Pointer(frame + uintptr(i)*2))) >> int32(16)))
					gain_Q16 = gain_Q16 + slope_Q16
					if gain_Q16 > int32(1)<<int32(16) {
						break
					}
					i = i + 1
				}
			}
		}
		(*OpusT_silk_PLC_struct)(unsafe.Pointer(psPLC)).Flast_frame_lost = 0
	}
}

const silk_int16_MAX8 = 0x7FFF

/* shell coder; pulse-subframe length is hardcoded */
