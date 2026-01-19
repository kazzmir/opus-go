// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

//go:build amd64 || arm64

package opuscc

import (
	"reflect"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

var _ reflect.Type
var _ unsafe.Pointer

func Opus_silk_decode_frame(tls *libc.TLS, psDec uintptr, psRangeDec uintptr, pOut uintptr, pN uintptr, lostFlag int32, condCoding int32, arch int32) (r int32) {
	var L, mv_len, ret int32
	var _saved_stack, psDecCtrl, pulses, st, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = L, _saved_stack, mv_len, psDecCtrl, pulses, ret, st, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9
	ret = 0
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
	L = (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length
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
	if !(int64(int32(uint64(uint32(int32(1)))*(uint64(140)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5898, int32(64))
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
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32(int32(1))) * (uint64(140) / uint64(1)))
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
	psDecCtrl = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32(int32(1)))*(uint64(140)/uint64(1)))
	(*OpusT_silk_decoder_control)(unsafe.Pointer(psDecCtrl)).FLTP_scale_Q14 = 0
	/* Safety checks */
	if !(L > 0 && L <= int32(SUB_FRAME_LENGTH_MS)*int32(MAX_NB_SUBFR)*int32(MAX_FS_KHZ)) {
		Opus_celt_fatal(tls, __ccgo_ts+5921, __ccgo_ts+5898, int32(68))
	}
	if lostFlag == FLAG_DECODE_NORMAL || lostFlag == int32(FLAG_DECODE_LBRR) && *(*int32)(unsafe.Pointer(psDec + 2432 + uintptr((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FnFramesDecoded)*4)) == int32(1) {
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
		*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(2)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (uint64(uint32(2)) - uint64(uint32(1))))
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
		if !(int64(int32(uint64(uint32((L+int32(SHELL_CODEC_FRAME_LENGTH)-int32(1)) & ^(int32(SHELL_CODEC_FRAME_LENGTH)-int32(1))))*(uint64(2)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
			Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5898, int32(78))
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
		*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32((L+int32(SHELL_CODEC_FRAME_LENGTH)-int32(1)) & ^(int32(SHELL_CODEC_FRAME_LENGTH)-int32(1)))) * (uint64(2) / uint64(1)))
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
		pulses = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32((L+int32(SHELL_CODEC_FRAME_LENGTH)-int32(1)) & ^(int32(SHELL_CODEC_FRAME_LENGTH)-int32(1))))*(uint64(2)/uint64(1)))
		/*********************************************/
		/* Decode quantization indices of side info  */
		/*********************************************/
		Opus_silk_decode_indices(tls, psDec, psRangeDec, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FnFramesDecoded, lostFlag, condCoding)
		/*********************************************/
		/* Decode quantization indices of excitation */
		/*********************************************/
		Opus_silk_decode_pulses(tls, psRangeDec, pulses, int32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FsignalType), int32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FquantOffsetType), (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length)
		/********************************************/
		/* Decode parameters and pulse signal       */
		/********************************************/
		Opus_silk_decode_parameters(tls, psDec, psDecCtrl, condCoding)
		/********************************************************/
		/* Run inverse NSQ                                      */
		/********************************************************/
		Opus_silk_decode_core(tls, psDec, psDecCtrl, pOut, pulses, arch)
		/*************************/
		/* Update output buffer. */
		/*************************/
		if !((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length >= (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length) {
			Opus_celt_fatal(tls, __ccgo_ts+5970, __ccgo_ts+5898, int32(104))
		}
		mv_len = (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length - (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length
		libc.Xmemmove(tls, psDec+1348, psDec+1348+uintptr((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length)*2, uint64(uint32(mv_len))*uint64(2))
		libc.Xmemcpy(tls, psDec+1348+uintptr(mv_len)*2, pOut, uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length))*uint64(2))
		/********************************************************/
		/* Update PLC state                                     */
		/********************************************************/
		Opus_silk_PLC(tls, psDec, psDecCtrl, pOut, 0, arch)
		(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FlossCnt = 0
		(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FprevSignalType = int32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FsignalType)
		if !((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FprevSignalType >= 0 && (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FprevSignalType <= int32(2)) {
			Opus_celt_fatal(tls, __ccgo_ts+6033, __ccgo_ts+5898, int32(127))
		}
		/* A frame has been decoded without errors */
		(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Ffirst_frame_after_reset = 0
	} else {
		/* Handle packet loss by extrapolation */
		Opus_silk_PLC(tls, psDec, psDecCtrl, pOut, int32(1), arch)
		/*************************/
		/* Update output buffer. */
		/*************************/
		if !((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length >= (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length) {
			Opus_celt_fatal(tls, __ccgo_ts+5970, __ccgo_ts+5898, int32(145))
		}
		mv_len = (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length - (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length
		libc.Xmemmove(tls, psDec+1348, psDec+1348+uintptr((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length)*2, uint64(uint32(mv_len))*uint64(2))
		libc.Xmemcpy(tls, psDec+1348+uintptr(mv_len)*2, pOut, uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length))*uint64(2))
	}
	/************************************************/
	/* Comfort noise generation / estimation        */
	/************************************************/
	Opus_silk_CNG(tls, psDec, psDecCtrl, pOut, L)
	/****************************************************************/
	/* Ensure smooth connection of extrapolated and good frames     */
	/****************************************************************/
	Opus_silk_PLC_glue_frames(tls, psDec, pOut, L)
	/* Update some decoder state variables */
	(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FlagPrev = *(*int32)(unsafe.Pointer(psDecCtrl + uintptr((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr-int32(1))*4))
	/* Set output frame length */
	*(*OpusT_opus_int32)(unsafe.Pointer(pN)) = L
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
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
	return ret
}

// C documentation
//
//	/* Decode parameters from payload */
func Opus_silk_decode_parameters(tls *libc.TLS, psDec uintptr, psDecCtrl uintptr, condCoding int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var Ix, i, k int32
	var cbk_ptr_Q7 uintptr
	var _ /* pNLSF0_Q15 at bp+32 */ [16]OpusT_opus_int16
	var _ /* pNLSF_Q15 at bp+0 */ [16]OpusT_opus_int16
	_, _, _, _ = Ix, cbk_ptr_Q7, i, k
	/* Dequant Gains */
	Opus_silk_gains_dequant(tls, psDecCtrl+16, psDec+2856, psDec+2312, libc.BoolInt32(condCoding == int32(CODE_CONDITIONALLY)), (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr)
	/****************/
	/* Decode NLSFs */
	/****************/
	Opus_silk_NLSF_decode(tls, bp, psDec+2856+8, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FpsNLSF_CB)
	/* Convert NLSF parameters to AR prediction filter coefficients */
	Opus_silk_NLSF2A(tls, psDecCtrl+32+1*32, bp, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Farch)
	/* If just reset, e.g., because internal Fs changed, do not allow interpolation */
	/* improves the case of packet loss in the first frame after a switch           */
	if (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Ffirst_frame_after_reset == int32(1) {
		(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FNLSFInterpCoef_Q2 = int8(4)
	}
	if int32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FNLSFInterpCoef_Q2) < int32(4) {
		/* Calculation of the interpolated NLSF0 vector from the interpolation factor, */
		/* the previous NLSF1, and the current NLSF1                                   */
		i = 0
		for {
			if !(i < (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order) {
				break
			}
			(*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 32)))[i] = int16(int32(*(*OpusT_opus_int16)(unsafe.Pointer(psDec + 2344 + uintptr(i)*2))) + int32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FNLSFInterpCoef_Q2)*(int32((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp)))[i])-int32(*(*OpusT_opus_int16)(unsafe.Pointer(psDec + 2344 + uintptr(i)*2))))>>int32(2))
			goto _1
		_1:
			;
			i = i + 1
		}
		/* Convert NLSF parameters to AR prediction filter coefficients */
		Opus_silk_NLSF2A(tls, psDecCtrl+32, bp+32, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Farch)
	} else {
		/* Copy LPC coefficients for first half from second half */
		libc.Xmemcpy(tls, psDecCtrl+32, psDecCtrl+32+1*32, uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order))*uint64(2))
	}
	libc.Xmemcpy(tls, psDec+2344, bp, uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order))*uint64(2))
	/* After a packet loss do BWE of LPC coefs */
	if (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FlossCnt != 0 {
		Opus_silk_bwexpander(tls, psDecCtrl+32, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order, int32(BWE_AFTER_LOSS_Q16))
		Opus_silk_bwexpander(tls, psDecCtrl+32+1*32, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order, int32(BWE_AFTER_LOSS_Q16))
	}
	if int32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FsignalType) == int32(TYPE_VOICED) {
		/*********************/
		/* Decode pitch lags */
		/*********************/
		/* Decode pitch values */
		Opus_silk_decode_pitch(tls, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FlagIndex, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FcontourIndex, psDecCtrl, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Ffs_kHz, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr)
		/* Decode Codebook Index */
		cbk_ptr_Q7 = Opus_silk_LTP_vq_ptrs_Q7[(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FPERIndex] /* set pointer to start of codebook */
		k = 0
		for {
			if !(k < (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr) {
				break
			}
			Ix = int32(*(*OpusT_opus_int8)(unsafe.Pointer(psDec + 2856 + 4 + uintptr(k))))
			i = 0
			for {
				if !(i < int32(LTP_ORDER)) {
					break
				}
				*(*OpusT_opus_int16)(unsafe.Pointer(psDecCtrl + 96 + uintptr(k*int32(LTP_ORDER)+i)*2)) = int16(int32(uint32(uint8(*(*OpusT_opus_int8)(unsafe.Pointer(cbk_ptr_Q7 + uintptr(Ix*int32(LTP_ORDER)+i))))) << int32(7)))
				goto _3
			_3:
				;
				i = i + 1
			}
			goto _2
		_2:
			;
			k = k + 1
		}
		/**********************/
		/* Decode LTP scaling */
		/**********************/
		Ix = int32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FLTP_scaleIndex)
		(*OpusT_silk_decoder_control)(unsafe.Pointer(psDecCtrl)).FLTP_scale_Q14 = int32(Opus_silk_LTPScales_table_Q14[Ix])
	} else {
		libc.Xmemset(tls, psDecCtrl, 0, uint64(uint32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr))*uint64(4))
		libc.Xmemset(tls, psDecCtrl+96, 0, uint64(uint32(int32(LTP_ORDER)*(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr))*uint64(2))
		(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FPERIndex = 0
		(*OpusT_silk_decoder_control)(unsafe.Pointer(psDecCtrl)).FLTP_scale_Q14 = 0
	}
}

// C documentation
//
//	/* Decode side-information parameters from payload */
func Opus_silk_decode_indices(tls *libc.TLS, psDec uintptr, psRangeDec uintptr, FrameIndex int32, decode_LBRR int32, condCoding int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var Ix, decode_absolute_lagIndex, delta_lagIndex, i, k int32
	var v1 uintptr
	var _ /* ec_ix at bp+0 */ [16]OpusT_opus_int16
	var _ /* pred_Q8 at bp+32 */ [16]OpusT_opus_uint8
	_, _, _, _, _, _ = Ix, decode_absolute_lagIndex, delta_lagIndex, i, k, v1
	/*******************************************/
	/* Decode signal type and quantizer offset */
	/*******************************************/
	if decode_LBRR != 0 || *(*int32)(unsafe.Pointer(psDec + 2416 + uintptr(FrameIndex)*4)) != 0 {
		Ix = Opus_ec_dec_icdf(tls, psRangeDec, uintptr(unsafe.Pointer(&Opus_silk_type_offset_VAD_iCDF)), uint32(8)) + int32(2)
	} else {
		Ix = Opus_ec_dec_icdf(tls, psRangeDec, uintptr(unsafe.Pointer(&Opus_silk_type_offset_no_VAD_iCDF)), uint32(8))
	}
	(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FsignalType = int8(Ix >> int32(1))
	(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FquantOffsetType = int8(Ix & int32(1))
	/****************/
	/* Decode gains */
	/****************/
	/* First subframe */
	if condCoding == int32(CODE_CONDITIONALLY) {
		/* Conditional coding */
		*(*OpusT_opus_int8)(unsafe.Pointer(psDec + 2856)) = int8(Opus_ec_dec_icdf(tls, psRangeDec, uintptr(unsafe.Pointer(&Opus_silk_delta_gain_iCDF)), uint32(8)))
	} else {
		/* Independent coding, in two stages: MSB bits followed by 3 LSBs */
		*(*OpusT_opus_int8)(unsafe.Pointer(psDec + 2856)) = int8(int32(uint32(Opus_ec_dec_icdf(tls, psRangeDec, uintptr(unsafe.Pointer(&Opus_silk_gain_iCDF))+uintptr((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FsignalType)*8, uint32(8))) << int32(3)))
		v1 = psDec + 2856
		*(*OpusT_opus_int8)(unsafe.Pointer(v1)) = OpusT_opus_int8(int32(*(*OpusT_opus_int8)(unsafe.Pointer(v1))) + int32(int8(Opus_ec_dec_icdf(tls, psRangeDec, uintptr(unsafe.Pointer(&Opus_silk_uniform8_iCDF)), uint32(8)))))
	}
	/* Remaining subframes */
	i = int32(1)
	for {
		if !(i < (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr) {
			break
		}
		*(*OpusT_opus_int8)(unsafe.Pointer(psDec + 2856 + uintptr(i))) = int8(Opus_ec_dec_icdf(tls, psRangeDec, uintptr(unsafe.Pointer(&Opus_silk_delta_gain_iCDF)), uint32(8)))
		goto _2
	_2:
		;
		i = i + 1
	}
	/**********************/
	/* Decode LSF Indices */
	/**********************/
	*(*OpusT_opus_int8)(unsafe.Pointer(psDec + 2856 + 8)) = int8(Opus_ec_dec_icdf(tls, psRangeDec, (*OpusT_silk_NLSF_CB_struct)(unsafe.Pointer((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FpsNLSF_CB)).FCB1_iCDF+uintptr(int32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FsignalType)>>int32(1)*int32((*OpusT_silk_NLSF_CB_struct)(unsafe.Pointer((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FpsNLSF_CB)).FnVectors)), uint32(8)))
	Opus_silk_NLSF_unpack(tls, bp, bp+32, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FpsNLSF_CB, int32(*(*OpusT_opus_int8)(unsafe.Pointer(psDec + 2856 + 8))))
	if !(int32((*OpusT_silk_NLSF_CB_struct)(unsafe.Pointer((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FpsNLSF_CB)).Forder) == (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order) {
		Opus_celt_fatal(tls, __ccgo_ts+6108, __ccgo_ts+6170, int32(82))
	}
	i = 0
	for {
		if !(i < int32((*OpusT_silk_NLSF_CB_struct)(unsafe.Pointer((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FpsNLSF_CB)).Forder)) {
			break
		}
		Ix = Opus_ec_dec_icdf(tls, psRangeDec, (*OpusT_silk_NLSF_CB_struct)(unsafe.Pointer((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FpsNLSF_CB)).Fec_iCDF+uintptr((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp)))[i]), uint32(8))
		if Ix == 0 {
			Ix = Ix - Opus_ec_dec_icdf(tls, psRangeDec, uintptr(unsafe.Pointer(&Opus_silk_NLSF_EXT_iCDF)), uint32(8))
		} else {
			if Ix == int32(2)*int32(NLSF_QUANT_MAX_AMPLITUDE) {
				Ix = Ix + Opus_ec_dec_icdf(tls, psRangeDec, uintptr(unsafe.Pointer(&Opus_silk_NLSF_EXT_iCDF)), uint32(8))
			}
		}
		*(*OpusT_opus_int8)(unsafe.Pointer(psDec + 2856 + 8 + uintptr(i+int32(1)))) = int8(Ix - int32(NLSF_QUANT_MAX_AMPLITUDE))
		goto _3
	_3:
		;
		i = i + 1
	}
	/* Decode LSF interpolation factor */
	if (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr == int32(MAX_NB_SUBFR) {
		(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FNLSFInterpCoef_Q2 = int8(Opus_ec_dec_icdf(tls, psRangeDec, uintptr(unsafe.Pointer(&Opus_silk_NLSF_interpolation_factor_iCDF)), uint32(8)))
	} else {
		(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FNLSFInterpCoef_Q2 = int8(4)
	}
	if int32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FsignalType) == int32(TYPE_VOICED) {
		/*********************/
		/* Decode pitch lags */
		/*********************/
		/* Get lag index */
		decode_absolute_lagIndex = int32(1)
		if condCoding == int32(CODE_CONDITIONALLY) && (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fec_prevSignalType == int32(TYPE_VOICED) {
			/* Decode Delta index */
			delta_lagIndex = int32(int16(Opus_ec_dec_icdf(tls, psRangeDec, uintptr(unsafe.Pointer(&Opus_silk_pitch_delta_iCDF)), uint32(8))))
			if delta_lagIndex > 0 {
				delta_lagIndex = delta_lagIndex - int32(9)
				(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FlagIndex = int16(int32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fec_prevLagIndex) + delta_lagIndex)
				decode_absolute_lagIndex = 0
			}
		}
		if decode_absolute_lagIndex != 0 {
			/* Absolute decoding */
			(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FlagIndex = int16(int32(int16(Opus_ec_dec_icdf(tls, psRangeDec, uintptr(unsafe.Pointer(&Opus_silk_pitch_lag_iCDF)), uint32(8)))) * ((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Ffs_kHz >> int32(1)))
			v1 = psDec + 2856 + 26
			*(*OpusT_opus_int16)(unsafe.Pointer(v1)) = OpusT_opus_int16(int32(*(*OpusT_opus_int16)(unsafe.Pointer(v1))) + int32(int16(Opus_ec_dec_icdf(tls, psRangeDec, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fpitch_lag_low_bits_iCDF, uint32(8)))))
		}
		(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fec_prevLagIndex = (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FlagIndex
		/* Get contour index */
		(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FcontourIndex = int8(Opus_ec_dec_icdf(tls, psRangeDec, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fpitch_contour_iCDF, uint32(8)))
		/********************/
		/* Decode LTP gains */
		/********************/
		/* Decode PERIndex value */
		(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FPERIndex = int8(Opus_ec_dec_icdf(tls, psRangeDec, uintptr(unsafe.Pointer(&Opus_silk_LTP_per_index_iCDF)), uint32(8)))
		k = 0
		for {
			if !(k < (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr) {
				break
			}
			*(*OpusT_opus_int8)(unsafe.Pointer(psDec + 2856 + 4 + uintptr(k))) = int8(Opus_ec_dec_icdf(tls, psRangeDec, Opus_silk_LTP_gain_iCDF_ptrs[(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FPERIndex], uint32(8)))
			goto _5
		_5:
			;
			k = k + 1
		}
		/**********************/
		/* Decode LTP scaling */
		/**********************/
		if condCoding == CODE_INDEPENDENTLY {
			(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FLTP_scaleIndex = int8(Opus_ec_dec_icdf(tls, psRangeDec, uintptr(unsafe.Pointer(&Opus_silk_LTPscale_iCDF)), uint32(8)))
		} else {
			(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FLTP_scaleIndex = 0
		}
	}
	(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fec_prevSignalType = int32((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FsignalType)
	/***************/
	/* Decode seed */
	/***************/
	(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Findices.FSeed = int8(Opus_ec_dec_icdf(tls, psRangeDec, uintptr(unsafe.Pointer(&Opus_silk_uniform4_iCDF)), uint32(8)))
}

// C documentation
//
//	/*********************************************/
//	/* Decode quantization indices of excitation */
//	/*********************************************/
func Opus_silk_decode_pulses(tls *libc.TLS, psRangeDec uintptr, pulses uintptr, signalType int32, quantOffsetType int32, frame_length int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var RateLevelIndex, abs_q, i, iter, j, k, nLS int32
	var cdf_ptr, pulses_ptr uintptr
	var nLshifts [20]int32
	var _ /* sum_pulses at bp+0 */ [20]int32
	_, _, _, _, _, _, _, _, _, _ = RateLevelIndex, abs_q, cdf_ptr, i, iter, j, k, nLS, nLshifts, pulses_ptr
	/*********************/
	/* Decode rate level */
	/*********************/
	RateLevelIndex = Opus_ec_dec_icdf(tls, psRangeDec, uintptr(unsafe.Pointer(&Opus_silk_rate_levels_iCDF))+uintptr(signalType>>int32(1))*9, uint32(8))
	/* Calculate number of shell blocks */
	_ = int32(1)<<int32(LOG2_SHELL_CODEC_FRAME_LENGTH) == int32(SHELL_CODEC_FRAME_LENGTH)
	iter = frame_length >> int32(LOG2_SHELL_CODEC_FRAME_LENGTH)
	if iter*int32(SHELL_CODEC_FRAME_LENGTH) < frame_length {
		if !(frame_length == int32(12)*int32(10)) {
			Opus_celt_fatal(tls, __ccgo_ts+6195, __ccgo_ts+6237, int32(59))
		} /* Make sure only happens for 10 ms @ 12 kHz */
		iter = iter + 1
	}
	/***************************************************/
	/* Sum-Weighted-Pulses Decoding                    */
	/***************************************************/
	cdf_ptr = uintptr(unsafe.Pointer(&Opus_silk_pulses_per_block_iCDF)) + uintptr(RateLevelIndex)*18
	i = 0
	for {
		if !(i < iter) {
			break
		}
		nLshifts[i] = 0
		(*(*[20]int32)(unsafe.Pointer(bp)))[i] = Opus_ec_dec_icdf(tls, psRangeDec, cdf_ptr, uint32(8))
		/* LSB indication */
		for (*(*[20]int32)(unsafe.Pointer(bp)))[i] == int32(SILK_MAX_PULSES)+int32(1) {
			nLshifts[i] = nLshifts[i] + 1
			/* When we've already got 10 LSBs, we shift the table to not allow (SILK_MAX_PULSES + 1) */
			(*(*[20]int32)(unsafe.Pointer(bp)))[i] = Opus_ec_dec_icdf(tls, psRangeDec, uintptr(unsafe.Pointer(&Opus_silk_pulses_per_block_iCDF))+uintptr(int32(N_RATE_LEVELS)-int32(1))*18+libc.BoolUintptr(nLshifts[i] == int32(10)), uint32(8))
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	/***************************************************/
	/* Shell decoding                                  */
	/***************************************************/
	i = 0
	for {
		if !(i < iter) {
			break
		}
		if (*(*[20]int32)(unsafe.Pointer(bp)))[i] > 0 {
			Opus_silk_shell_decoder(tls, pulses+uintptr(int32(int16(i))*int32(int16(int32(SHELL_CODEC_FRAME_LENGTH))))*2, psRangeDec, (*(*[20]int32)(unsafe.Pointer(bp)))[i])
		} else {
			libc.Xmemset(tls, pulses+uintptr(int32(int16(i))*int32(int16(int32(SHELL_CODEC_FRAME_LENGTH))))*2, 0, uint64(uint32(SHELL_CODEC_FRAME_LENGTH))*uint64(2))
		}
		goto _2
	_2:
		;
		i = i + 1
	}
	/***************************************************/
	/* LSB Decoding                                    */
	/***************************************************/
	i = 0
	for {
		if !(i < iter) {
			break
		}
		if nLshifts[i] > 0 {
			nLS = nLshifts[i]
			pulses_ptr = pulses + uintptr(int32(int16(i))*int32(int16(int32(SHELL_CODEC_FRAME_LENGTH))))*2
			k = 0
			for {
				if !(k < int32(SHELL_CODEC_FRAME_LENGTH)) {
					break
				}
				abs_q = int32(*(*OpusT_opus_int16)(unsafe.Pointer(pulses_ptr + uintptr(k)*2)))
				j = 0
				for {
					if !(j < nLS) {
						break
					}
					abs_q = int32(uint32(abs_q) << int32(1))
					abs_q = abs_q + Opus_ec_dec_icdf(tls, psRangeDec, uintptr(unsafe.Pointer(&Opus_silk_lsb_iCDF)), uint32(8))
					goto _5
				_5:
					;
					j = j + 1
				}
				*(*OpusT_opus_int16)(unsafe.Pointer(pulses_ptr + uintptr(k)*2)) = int16(abs_q)
				goto _4
			_4:
				;
				k = k + 1
			}
			/* Mark the number of pulses non-zero for sign decoding. */
			*(*int32)(unsafe.Pointer(bp + uintptr(i)*4)) |= nLS << int32(5)
		}
		goto _3
	_3:
		;
		i = i + 1
	}
	/****************************************/
	/* Decode and add signs to pulse signal */
	/****************************************/
	Opus_silk_decode_signs(tls, psRangeDec, pulses, frame_length, signalType, quantOffsetType, bp)
}

// C documentation
//
//	/* Set decoder sampling rate */
