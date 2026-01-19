// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

//go:build linux && (amd64 || arm64)

package opuscc

import (
	"reflect"
	"unsafe"

	libc "opusgo/libcshim"
)

var _ reflect.Type
var _ unsafe.Pointer

func Opus_silk_reset_decoder(tls *libc.TLS, psDec uintptr) (r int32) {
	var v1 int32
	_ = v1
	/* Clear the entire encoder state, except anything copied */
	libc.Xmemset(tls, psDec, 0, uint64(4392)-uint64(int64(psDec)-int64(psDec)))
	/* Used to deactivate LSF interpolation */
	(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Ffirst_frame_after_reset = int32(1)
	(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fprev_gain_Q16 = int32(65536)
	v1 = 0
	goto _2
_2:
	(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Farch = v1
	/* Reset CNG state */
	Opus_silk_CNG_Reset(tls, psDec)
	/* Reset PLC state */
	Opus_silk_PLC_Reset(tls, psDec)
	return 0
}

// C documentation
//
//	/************************/
//	/* Init Decoder State   */
//	/************************/
func Opus_silk_init_decoder(tls *libc.TLS, psDec uintptr) (r int32) {
	/* Clear the entire encoder state, except anything copied */
	libc.Xmemset(tls, psDec, 0, uint64(4392))
	Opus_silk_reset_decoder(tls, psDec)
	return 0
}

const silk_int16_MAX3 = 32767

// C documentation
//
//	/**********************************************************/
//	/* Core decoder. Performs inverse NSQ operation LTP + LPC */
//	/**********************************************************/
func Opus_silk_decode_core(tls *libc.TLS, psDec uintptr, psDecCtrl uintptr, xq uintptr, pulses uintptr, arch int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var A_Q12, B_Q14, _saved_stack, pexc_Q14, pred_lag_ptr, pres_Q14, pxq, res_Q14, sLPC_Q14, sLTP, sLTP_Q15, st, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var Gain_Q10, LPC_pred_Q10, LTP_pred_Q13, a32_nrm, b32_inv, b32_inv1, b32_nrm, b32_nrm1, err_Q32, gain_adj_Q16, inv_gain_Q31, offset_Q10, rand_seed, result, result1, v103, v106, v107, v110, v117, v118, v121 OpusT_opus_int32
	var NLSF_interpolation_flag, a_headrm, b_headrm, b_headrm1, i, k, lag, lshift, lshift1, sLTP_buf_idx, signalType, start_idx, v104, v105, v109, v112, v113, v114, v115, v116, v119, v120, v124, v125, v129 int32
	var _ /* A_Q12_tmp at bp+0 */ [16]OpusT_opus_int16
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = A_Q12, B_Q14, Gain_Q10, LPC_pred_Q10, LTP_pred_Q13, NLSF_interpolation_flag, _saved_stack, a32_nrm, a_headrm, b32_inv, b32_inv1, b32_nrm, b32_nrm1, b_headrm, b_headrm1, err_Q32, gain_adj_Q16, i, inv_gain_Q31, k, lag, lshift, lshift1, offset_Q10, pexc_Q14, pred_lag_ptr, pres_Q14, pxq, rand_seed, res_Q14, result, result1, sLPC_Q14, sLTP, sLTP_Q15, sLTP_buf_idx, signalType, st, start_idx, v1, v103, v104, v105, v106, v107, v109, v11, v110, v112, v113, v114, v115, v116, v117, v118, v119, v120, v121, v124, v125, v129, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9
	lag = 0
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
	_ = (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fprev_gain_Q16 != int32(0)
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
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(2)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (uint64(uint32(2)) - uint64(uint32(1))))
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
	if !(int64(int32(uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length))*(uint64(2)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5844, int32(58))
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
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length)) * (uint64(2) / uint64(1)))
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
	sLTP = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length))*(uint64(2)/uint64(1)))
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
	if !(int64(int32(uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length+(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5844, int32(59))
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
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length+(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length)) * (uint64(4) / uint64(1)))
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
	sLTP_Q15 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length+(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length))*(uint64(4)/uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _54
	_54:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _56
_56:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _58
	_58:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v7 = st
	goto _60
_60:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		goto _62
	_62:
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v11 = st
	goto _64
_64:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		goto _66
	_66:
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v15 = st
	goto _68
_68:
	if !(int64(int32(uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5844, int32(60))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		goto _70
	_70:
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v19 = st
	goto _72
_72:
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		goto _74
	_74:
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v23 = st
	goto _76
_76:
	res_Q14 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length))*(uint64(4)/uint64(1)))
	/* Work around a clang bug (verified with clang 6.0 through clang 20.1.0) that causes the last
	   memset to be flagged as an invalid read by valgrind (not caught by asan). */
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _78
	_78:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _80
_80:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _82
	_82:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v7 = st
	goto _84
_84:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		goto _86
	_86:
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v11 = st
	goto _88
_88:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		goto _90
	_90:
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v15 = st
	goto _92
_92:
	if !(int64(int32(uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length+int32(MAX_LPC_ORDER)))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5844, int32(66))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		goto _94
	_94:
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v19 = st
	goto _96
_96:
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length+int32(MAX_LPC_ORDER))) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		goto _98
	_98:
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v23 = st
	goto _100
_100:
	sLPC_Q14 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length+int32(MAX_LPC_ORDER)))*(uint64(4)/uint64(1)))
	offset_Q10 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_silk_Quantization_Offsets_Q10)) + uintptr(int32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FsignalType)>>int32(1))*4 + uintptr((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FquantOffsetType)*2)))
	if int32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FNLSFInterpCoef_Q2) < int32(1)<<int32(2) {
		NLSF_interpolation_flag = int32(1)
	} else {
		NLSF_interpolation_flag = 0
	}
	/* Decode excitation */
	rand_seed = int32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FSeed)
	i = 0
	for {
		if !(i < (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length) {
			break
		}
		rand_seed = int32(uint32(int32(RAND_INCREMENT)) + uint32(rand_seed)*uint32(int32(RAND_MULTIPLIER)))
		*(*OpusT_opus_int32)(unsafe.Pointer(psDec + 4 + uintptr(i)*4)) = int32(uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(pulses + uintptr(i)*2)))) << int32(14))
		if *(*OpusT_opus_int32)(unsafe.Pointer(psDec + 4 + uintptr(i)*4)) > 0 {
			*(*OpusT_opus_int32)(unsafe.Pointer(psDec + 4 + uintptr(i)*4)) -= int32(QUANT_LEVEL_ADJUST_Q10) << int32(4)
		} else {
			if *(*OpusT_opus_int32)(unsafe.Pointer(psDec + 4 + uintptr(i)*4)) < 0 {
				*(*OpusT_opus_int32)(unsafe.Pointer(psDec + 4 + uintptr(i)*4)) += int32(QUANT_LEVEL_ADJUST_Q10) << int32(4)
			}
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(psDec + 4 + uintptr(i)*4)) += offset_Q10 << int32(4)
		if rand_seed < 0 {
			*(*OpusT_opus_int32)(unsafe.Pointer(psDec + 4 + uintptr(i)*4)) = -*(*OpusT_opus_int32)(unsafe.Pointer(psDec + 4 + uintptr(i)*4))
		}
		rand_seed = int32(uint32(rand_seed) + uint32(uint16(*(*OpusT_opus_int16)(unsafe.Pointer(pulses + uintptr(i)*2)))))
		goto _101
	_101:
		;
		i = i + 1
	}
	/* Copy LPC state */
	libc.Xmemcpy(tls, sLPC_Q14, psDec+1284, uint64(uint32(MAX_LPC_ORDER))*uint64(4))
	pexc_Q14 = psDec + 4
	pxq = xq
	sLTP_buf_idx = (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length
	/* Loop over subframes */
	k = 0
	for {
		if !(k < (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr) {
			break
		}
		pres_Q14 = res_Q14
		A_Q12 = psDecCtrl + 32 + uintptr(k>>int32(1))*32
		/* Preload LPC coefficients to array on stack. Gives small performance gain */
		libc.Xmemcpy(tls, bp, A_Q12, uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order))*uint64(2))
		B_Q14 = psDecCtrl + 96 + uintptr(k*int32(LTP_ORDER))*2
		signalType = int32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FsignalType)
		Gain_Q10 = *(*OpusT_opus_int32)(unsafe.Pointer(psDecCtrl + 16 + uintptr(k)*4)) >> int32(6)
		v103 = *(*OpusT_opus_int32)(unsafe.Pointer(psDecCtrl + 16 + uintptr(k)*4))
		v104 = int32(47)
		_ = v103 != int32(0)
		_ = v104 > int32(0)
		if v103 > 0 {
			v105 = v103
		} else {
			v105 = -v103
		}
		v106 = v105
		if v106 != 0 {
			v109 = int32(32) - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, uint32(v106)))
		} else {
			v109 = int32(32)
		}
		v107 = v109
		goto _108
	_108:
		b_headrm1 = v107 - int32(1)
		b32_nrm1 = int32(uint32(v103) << b_headrm1)
		b32_inv1 = int32(silk_int32_MAX) >> int32(2) / (b32_nrm1 >> int32(16))
		result1 = int32(uint32(b32_inv1) << int32(16))
		err_Q32 = int32(uint32(int32(1)<<int32(29)-int32(int64(b32_nrm1)*int64(int16(b32_inv1))>>int32(16))) << int32(3))
		result1 = int32(int64(result1) + int64(err_Q32)*int64(b32_inv1)>>int32(16))
		lshift1 = int32(61) - b_headrm1 - v104
		if lshift1 <= int32(0) {
			if int32(-2147483648)>>-lshift1 > int32(silk_int32_MAX)>>-lshift1 {
				if result1 > int32(-2147483648)>>-lshift1 {
					v113 = int32(-2147483648) >> -lshift1
				} else {
					if result1 < int32(silk_int32_MAX)>>-lshift1 {
						v114 = int32(silk_int32_MAX) >> -lshift1
					} else {
						v114 = result1
					}
					v113 = v114
				}
				v112 = v113
			} else {
				if result1 > int32(silk_int32_MAX)>>-lshift1 {
					v115 = int32(silk_int32_MAX) >> -lshift1
				} else {
					if result1 < int32(-2147483648)>>-lshift1 {
						v116 = int32(-2147483648) >> -lshift1
					} else {
						v116 = result1
					}
					v115 = v116
				}
				v112 = v115
			}
			v110 = int32(uint32(v112) << -lshift1)
			goto _111
		} else {
			if lshift1 < int32(32) {
				v110 = result1 >> lshift1
				goto _111
			} else {
				v110 = 0
				goto _111
			}
		}
	_111:
		inv_gain_Q31 = v110
		/* Calculate gain adjustment factor */
		if *(*OpusT_opus_int32)(unsafe.Pointer(psDecCtrl + 16 + uintptr(k)*4)) != (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fprev_gain_Q16 {
			v103 = (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fprev_gain_Q16
			v106 = *(*OpusT_opus_int32)(unsafe.Pointer(psDecCtrl + 16 + uintptr(k)*4))
			v104 = int32(16)
			_ = v106 != int32(0)
			_ = v104 >= int32(0)
			if v103 > 0 {
				v105 = v103
			} else {
				v105 = -v103
			}
			v107 = v105
			if v107 != 0 {
				v109 = int32(32) - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, uint32(v107)))
			} else {
				v109 = int32(32)
			}
			v110 = v109
			goto _123
		_123:
			a_headrm = v110 - int32(1)
			a32_nrm = int32(uint32(v103) << a_headrm)
			if v106 > 0 {
				v112 = v106
			} else {
				v112 = -v106
			}
			v117 = v112
			if v117 != 0 {
				v113 = int32(32) - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, uint32(v117)))
			} else {
				v113 = int32(32)
			}
			v118 = v113
			goto _128
		_128:
			b_headrm = v118 - int32(1)
			b32_nrm = int32(uint32(v106) << b_headrm)
			b32_inv = int32(silk_int32_MAX) >> int32(2) / (b32_nrm >> int32(16))
			result = int32(int64(a32_nrm) * int64(int16(b32_inv)) >> int32(16))
			a32_nrm = int32(uint32(a32_nrm) - uint32(int32(uint32(int32(int64(b32_nrm)*int64(result)>>int32(32)))<<int32(3))))
			result = int32(int64(result) + int64(a32_nrm)*int64(int16(b32_inv))>>int32(16))
			lshift = int32(29) + a_headrm - b_headrm - v104
			if lshift < int32(0) {
				if int32(-2147483648)>>-lshift > int32(silk_int32_MAX)>>-lshift {
					if result > int32(-2147483648)>>-lshift {
						v115 = int32(-2147483648) >> -lshift
					} else {
						if result < int32(silk_int32_MAX)>>-lshift {
							v116 = int32(silk_int32_MAX) >> -lshift
						} else {
							v116 = result
						}
						v115 = v116
					}
					v114 = v115
				} else {
					if result > int32(silk_int32_MAX)>>-lshift {
						v119 = int32(silk_int32_MAX) >> -lshift
					} else {
						if result < int32(-2147483648)>>-lshift {
							v120 = int32(-2147483648) >> -lshift
						} else {
							v120 = result
						}
						v119 = v120
					}
					v114 = v119
				}
				v121 = int32(uint32(v114) << -lshift)
				goto _131
			} else {
				if lshift < int32(32) {
					v121 = result >> lshift
					goto _131
				} else {
					v121 = 0
					goto _131
				}
			}
		_131:
			gain_adj_Q16 = v121
			/* Scale short term state */
			i = 0
			for {
				if !(i < int32(MAX_LPC_ORDER)) {
					break
				}
				*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14 + uintptr(i)*4)) = int32(int64(gain_adj_Q16) * int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14 + uintptr(i)*4))) >> int32(16))
				goto _137
			_137:
				;
				i = i + 1
			}
		} else {
			gain_adj_Q16 = int32(1) << int32(16)
		}
		/* Save inv_gain */
		_ = inv_gain_Q31 != int32(0)
		(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fprev_gain_Q16 = *(*OpusT_opus_int32)(unsafe.Pointer(psDecCtrl + 16 + uintptr(k)*4))
		/* Avoid abrupt transition from voiced PLC to unvoiced normal decoding */
		if (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FlossCnt != 0 && (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FprevSignalType == int32(TYPE_VOICED) && int32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FsignalType) != int32(TYPE_VOICED) && k < int32(MAX_NB_SUBFR)/int32(2) {
			libc.Xmemset(tls, B_Q14, 0, uint64(uint32(LTP_ORDER))*uint64(2))
			*(*OpusT_opus_int16)(unsafe.Pointer(B_Q14 + uintptr(int32(LTP_ORDER)/int32(2))*2)) = int16(4096)
			signalType = int32(TYPE_VOICED)
			*(*int32)(unsafe.Pointer(psDecCtrl + uintptr(k)*4)) = (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FlagPrev
		}
		if signalType == int32(TYPE_VOICED) {
			/* Voiced */
			lag = *(*int32)(unsafe.Pointer(psDecCtrl + uintptr(k)*4))
			/* Re-whitening */
			if k == 0 || k == int32(2) && NLSF_interpolation_flag != 0 {
				/* Rewhiten with new A coefs */
				start_idx = (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length - lag - (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order - int32(LTP_ORDER)/int32(2)
				if !(start_idx > int32(0)) {
					Opus_celt_fatal(tls, __ccgo_ts+5866, __ccgo_ts+5844, int32(150))
				}
				if k == int32(2) {
					libc.Xmemcpy(tls, psDec+1348+uintptr((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length)*2, xq, uint64(uint32(int32(2)*(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length))*uint64(2))
				}
				Opus_silk_LPC_analysis_filter(tls, sLTP+uintptr(start_idx)*2, psDec+1348+uintptr(start_idx+k*(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length)*2, A_Q12, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length-start_idx, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order, arch)
				/* After rewhitening the LTP state is unscaled */
				if k == 0 {
					/* Do LTP downscaling to reduce inter-packet dependency */
					inv_gain_Q31 = int32(uint32(int32(int64(inv_gain_Q31)*int64(int16((*OpusT_silk_decoder_control)(unsafe.Pointer(psDecCtrl)).FLTP_scale_Q14))>>int32(16))) << int32(2))
				}
				i = 0
				for {
					if !(i < lag+int32(LTP_ORDER)/int32(2)) {
						break
					}
					*(*OpusT_opus_int32)(unsafe.Pointer(sLTP_Q15 + uintptr(sLTP_buf_idx-i-int32(1))*4)) = int32(int64(inv_gain_Q31) * int64(*(*OpusT_opus_int16)(unsafe.Pointer(sLTP + uintptr((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length-i-int32(1))*2))) >> int32(16))
					goto _138
				_138:
					;
					i = i + 1
				}
			} else {
				/* Update LTP state when Gain changes */
				if gain_adj_Q16 != int32(1)<<int32(16) {
					i = 0
					for {
						if !(i < lag+int32(LTP_ORDER)/int32(2)) {
							break
						}
						*(*OpusT_opus_int32)(unsafe.Pointer(sLTP_Q15 + uintptr(sLTP_buf_idx-i-int32(1))*4)) = int32(int64(gain_adj_Q16) * int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLTP_Q15 + uintptr(sLTP_buf_idx-i-int32(1))*4))) >> int32(16))
						goto _139
					_139:
						;
						i = i + 1
					}
				}
			}
		}
		/* Long-term prediction */
		if signalType == int32(TYPE_VOICED) {
			/* Set up pointer */
			pred_lag_ptr = sLTP_Q15 + uintptr(sLTP_buf_idx-lag+int32(LTP_ORDER)/int32(2))*4
			i = 0
			for {
				if !(i < (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length) {
					break
				}
				/* Unrolled loop */
				/* Avoids introducing a bias because silk_SMLAWB() always rounds to -inf */
				LTP_pred_Q13 = int32(2)
				LTP_pred_Q13 = int32(int64(LTP_pred_Q13) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(pred_lag_ptr)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(B_Q14)))>>int32(16))
				LTP_pred_Q13 = int32(int64(LTP_pred_Q13) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(pred_lag_ptr - uintptr(1)*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(B_Q14 + 1*2)))>>int32(16))
				LTP_pred_Q13 = int32(int64(LTP_pred_Q13) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(pred_lag_ptr - uintptr(2)*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(B_Q14 + 2*2)))>>int32(16))
				LTP_pred_Q13 = int32(int64(LTP_pred_Q13) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(pred_lag_ptr - uintptr(3)*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(B_Q14 + 3*2)))>>int32(16))
				LTP_pred_Q13 = int32(int64(LTP_pred_Q13) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(pred_lag_ptr - uintptr(4)*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(B_Q14 + 4*2)))>>int32(16))
				pred_lag_ptr += 4
				/* Generate LPC excitation */
				*(*OpusT_opus_int32)(unsafe.Pointer(pres_Q14 + uintptr(i)*4)) = *(*OpusT_opus_int32)(unsafe.Pointer(pexc_Q14 + uintptr(i)*4)) + int32(uint32(LTP_pred_Q13)<<int32(1))
				/* Update states */
				*(*OpusT_opus_int32)(unsafe.Pointer(sLTP_Q15 + uintptr(sLTP_buf_idx)*4)) = int32(uint32(*(*OpusT_opus_int32)(unsafe.Pointer(pres_Q14 + uintptr(i)*4))) << int32(1))
				sLTP_buf_idx = sLTP_buf_idx + 1
				goto _140
			_140:
				;
				i = i + 1
			}
		} else {
			pres_Q14 = pexc_Q14
		}
		i = 0
		for {
			if !(i < (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length) {
				break
			}
			/* Short-term prediction */
			if !((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order == int32(10) || (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order == int32(16)) {
				Opus_celt_fatal(tls, __ccgo_ts+5777, __ccgo_ts+5844, int32(205))
			}
			/* Avoids introducing a bias because silk_SMLAWB() always rounds to -inf */
			LPC_pred_Q10 = (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order >> int32(1)
			LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(1))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp)))[0])>>int32(16))
			LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(2))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp)))[int32(1)])>>int32(16))
			LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(3))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp)))[int32(2)])>>int32(16))
			LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(4))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp)))[int32(3)])>>int32(16))
			LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(5))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp)))[int32(4)])>>int32(16))
			LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(6))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp)))[int32(5)])>>int32(16))
			LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(7))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp)))[int32(6)])>>int32(16))
			LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(8))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp)))[int32(7)])>>int32(16))
			LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(9))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp)))[int32(8)])>>int32(16))
			LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(10))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp)))[int32(9)])>>int32(16))
			if (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order == int32(16) {
				LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(11))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp)))[int32(10)])>>int32(16))
				LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(12))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp)))[int32(11)])>>int32(16))
				LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(13))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp)))[int32(12)])>>int32(16))
				LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(14))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp)))[int32(13)])>>int32(16))
				LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(15))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp)))[int32(14)])>>int32(16))
				LPC_pred_Q10 = int32(int64(LPC_pred_Q10) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14 + uintptr(int32(MAX_LPC_ORDER)+i-int32(16))*4)))*int64((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp)))[int32(15)])>>int32(16))
			}
			/* Add prediction to LPC excitation */
			if LPC_pred_Q10 > int32(silk_int32_MAX)>>int32(4) {
				v105 = int32(silk_int32_MAX) >> int32(4)
			} else {
				if LPC_pred_Q10 < int32(-2147483648)>>int32(4) {
					v109 = int32(-2147483648) >> int32(4)
				} else {
					v109 = LPC_pred_Q10
				}
				v105 = v109
			}
			if (uint32(*(*OpusT_opus_int32)(unsafe.Pointer(pres_Q14 + uintptr(i)*4)))+uint32(int32(uint32(v105)<<int32(4))))&uint32(0x80000000) == uint32(0) {
				if LPC_pred_Q10 > int32(silk_int32_MAX)>>int32(4) {
					v113 = int32(silk_int32_MAX) >> int32(4)
				} else {
					if LPC_pred_Q10 < int32(-2147483648)>>int32(4) {
						v114 = int32(-2147483648) >> int32(4)
					} else {
						v114 = LPC_pred_Q10
					}
					v113 = v114
				}
				if uint32(*(*OpusT_opus_int32)(unsafe.Pointer(pres_Q14 + uintptr(i)*4))&int32(uint32(v113)<<int32(4)))&uint32(0x80000000) != uint32(0) {
					v112 = int32(-2147483648)
				} else {
					if LPC_pred_Q10 > int32(silk_int32_MAX)>>int32(4) {
						v115 = int32(silk_int32_MAX) >> int32(4)
					} else {
						if LPC_pred_Q10 < int32(-2147483648)>>int32(4) {
							v116 = int32(-2147483648) >> int32(4)
						} else {
							v116 = LPC_pred_Q10
						}
						v115 = v116
					}
					v112 = *(*OpusT_opus_int32)(unsafe.Pointer(pres_Q14 + uintptr(i)*4)) + int32(uint32(v115)<<int32(4))
				}
				v104 = v112
			} else {
				if LPC_pred_Q10 > int32(silk_int32_MAX)>>int32(4) {
					v120 = int32(silk_int32_MAX) >> int32(4)
				} else {
					if LPC_pred_Q10 < int32(-2147483648)>>int32(4) {
						v124 = int32(-2147483648) >> int32(4)
					} else {
						v124 = LPC_pred_Q10
					}
					v120 = v124
				}
				if uint32(*(*OpusT_opus_int32)(unsafe.Pointer(pres_Q14 + uintptr(i)*4))|int32(uint32(v120)<<int32(4)))&uint32(0x80000000) == uint32(0) {
					v119 = int32(silk_int32_MAX)
				} else {
					if LPC_pred_Q10 > int32(silk_int32_MAX)>>int32(4) {
						v125 = int32(silk_int32_MAX) >> int32(4)
					} else {
						if LPC_pred_Q10 < int32(-2147483648)>>int32(4) {
							v129 = int32(-2147483648) >> int32(4)
						} else {
							v129 = LPC_pred_Q10
						}
						v125 = v129
					}
					v119 = *(*OpusT_opus_int32)(unsafe.Pointer(pres_Q14 + uintptr(i)*4)) + int32(uint32(v125)<<int32(4))
				}
				v104 = v119
			}
			*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14 + uintptr(int32(MAX_LPC_ORDER)+i)*4)) = v104
			/* Scale with gain */
			if (int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14 + uintptr(int32(MAX_LPC_ORDER)+i)*4)))*int64(Gain_Q10)>>int32(16))>>(int32(8)-int32(1))+int32(1))>>int32(1) > int32(silk_int16_MAX3) {
				v104 = int32(silk_int16_MAX3)
			} else {
				if (int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14 + uintptr(int32(MAX_LPC_ORDER)+i)*4)))*int64(Gain_Q10)>>int32(16))>>(int32(8)-int32(1))+int32(1))>>int32(1) < int32(int16(-32768)) {
					v105 = int32(int16(-32768))
				} else {
					v105 = (int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(sLPC_Q14 + uintptr(int32(MAX_LPC_ORDER)+i)*4)))*int64(Gain_Q10)>>int32(16))>>(int32(8)-int32(1)) + int32(1)) >> int32(1)
				}
				v104 = v105
			}
			*(*OpusT_opus_int16)(unsafe.Pointer(pxq + uintptr(i)*2)) = int16(v104)
			goto _141
		_141:
			;
			i = i + 1
		}
		/* Update LPC filter state */
		libc.Xmemcpy(tls, sLPC_Q14, sLPC_Q14+uintptr((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length)*4, uint64(uint32(MAX_LPC_ORDER))*uint64(4))
		pexc_Q14 = pexc_Q14 + uintptr((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length)*4
		pxq = pxq + uintptr((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length)*2
		goto _102
	_102:
		;
		k = k + 1
	}
	/* Save LPC state */
	libc.Xmemcpy(tls, psDec+1284, sLPC_Q14, uint64(uint32(MAX_LPC_ORDER))*uint64(4))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _158
	_158:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _160
_160:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
}

const silk_int16_MAX4 = 0x7FFF

/***********************************************************************
Copyright (c) 2006-2011, Skype Limited. All rights reserved.
Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions
are met:
- Redistributions of source code must retain the above copyright notice,
this list of conditions and the following disclaimer.
- Redistributions in binary form must reproduce the above copyright
notice, this list of conditions and the following disclaimer in the
documentation and/or other materials provided with the distribution.
- Neither the name of Internet Society, IETF or IETF Trust, nor the
names of specific contributors, may be used to endorse or promote
products derived from this software without specific prior written
permission.
THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE
LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
POSSIBILITY OF SUCH DAMAGE.
***********************************************************************/

// C documentation
//
//	/****************/
//	/* Decode frame */
//	/****************/
