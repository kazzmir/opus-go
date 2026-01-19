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

func Opus_silk_resampler_init(tls *libc.TLS, S uintptr, Fs_Hz_in OpusT_opus_int32, Fs_Hz_out OpusT_opus_int32, forEnc int32) (r int32) {
	var up2x, v1, v2 int32
	_, _, _ = up2x, v1, v2
	/* Clear state */
	libc.Xmemset(tls, S, 0, uint64(400))
	/* Input checking */
	if forEnc != 0 {
		if Fs_Hz_in != int32(8000) && Fs_Hz_in != int32(12000) && Fs_Hz_in != int32(16000) && Fs_Hz_in != int32(24000) && Fs_Hz_in != int32(48000) || Fs_Hz_out != int32(8000) && Fs_Hz_out != int32(12000) && Fs_Hz_out != int32(16000) {
			if !(int32(0) != 0) {
				Opus_celt_fatal(tls, __ccgo_ts+1017, __ccgo_ts+7386, int32(99))
			}
			return -int32(1)
		}
		if int32(5) < (Fs_Hz_in>>int32(12)-libc.BoolInt32(Fs_Hz_in > int32(16000)))>>libc.BoolInt32(Fs_Hz_in > int32(24000))-int32(1) {
			v1 = int32(5)
		} else {
			v1 = (Fs_Hz_in>>int32(12)-libc.BoolInt32(Fs_Hz_in > int32(16000)))>>libc.BoolInt32(Fs_Hz_in > int32(24000)) - int32(1)
		}
		if int32(5) < (Fs_Hz_out>>int32(12)-libc.BoolInt32(Fs_Hz_out > int32(16000)))>>libc.BoolInt32(Fs_Hz_out > int32(24000))-int32(1) {
			v2 = int32(5)
		} else {
			v2 = (Fs_Hz_out>>int32(12)-libc.BoolInt32(Fs_Hz_out > int32(16000)))>>libc.BoolInt32(Fs_Hz_out > int32(24000)) - int32(1)
		}
		(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FinputDelay = int32(*(*OpusT_opus_int8)(unsafe.Pointer(uintptr(unsafe.Pointer(&delay_matrix_enc)) + uintptr(v1)*3 + uintptr(v2))))
	} else {
		if Fs_Hz_in != int32(8000) && Fs_Hz_in != int32(12000) && Fs_Hz_in != int32(16000) || Fs_Hz_out != int32(8000) && Fs_Hz_out != int32(12000) && Fs_Hz_out != int32(16000) && Fs_Hz_out != int32(24000) && Fs_Hz_out != int32(48000) {
			if !(int32(0) != 0) {
				Opus_celt_fatal(tls, __ccgo_ts+1017, __ccgo_ts+7386, int32(110))
			}
			return -int32(1)
		}
		if int32(5) < (Fs_Hz_in>>int32(12)-libc.BoolInt32(Fs_Hz_in > int32(16000)))>>libc.BoolInt32(Fs_Hz_in > int32(24000))-int32(1) {
			v1 = int32(5)
		} else {
			v1 = (Fs_Hz_in>>int32(12)-libc.BoolInt32(Fs_Hz_in > int32(16000)))>>libc.BoolInt32(Fs_Hz_in > int32(24000)) - int32(1)
		}
		if int32(5) < (Fs_Hz_out>>int32(12)-libc.BoolInt32(Fs_Hz_out > int32(16000)))>>libc.BoolInt32(Fs_Hz_out > int32(24000))-int32(1) {
			v2 = int32(5)
		} else {
			v2 = (Fs_Hz_out>>int32(12)-libc.BoolInt32(Fs_Hz_out > int32(16000)))>>libc.BoolInt32(Fs_Hz_out > int32(24000)) - int32(1)
		}
		(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FinputDelay = int32(*(*OpusT_opus_int8)(unsafe.Pointer(uintptr(unsafe.Pointer(&delay_matrix_dec)) + uintptr(v1)*6 + uintptr(v2))))
	}
	(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFs_in_kHz = Fs_Hz_in / int32(1000)
	(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFs_out_kHz = Fs_Hz_out / int32(1000)
	/* Number of samples processed per batch */
	(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FbatchSize = (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFs_in_kHz * int32(RESAMPLER_MAX_BATCH_SIZE_MS)
	/* Find resampler with the right sampling ratio */
	up2x = 0
	if Fs_Hz_out > Fs_Hz_in {
		/* Upsample */
		if Fs_Hz_out == Fs_Hz_in*int32(2) { /* Fs_out : Fs_in = 2 : 1 */
			/* Special case: directly use 2x upsampler */
			(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).Fresampler_function = int32(USE_silk_resampler_private_up2_HQ_wrapper)
		} else {
			/* Default resampler */
			(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).Fresampler_function = int32(USE_silk_resampler_private_IIR_FIR)
			up2x = int32(1)
		}
	} else {
		if Fs_Hz_out < Fs_Hz_in {
			/* Downsample */
			(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).Fresampler_function = int32(USE_silk_resampler_private_down_FIR)
			if Fs_Hz_out*int32(4) == Fs_Hz_in*int32(3) { /* Fs_out : Fs_in = 3 : 4 */
				(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFIR_Fracs = int32(3)
				(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFIR_Order = int32(RESAMPLER_DOWN_ORDER_FIR0)
				(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FCoefs = uintptr(unsafe.Pointer(&Opus_silk_Resampler_3_4_COEFS))
			} else {
				if Fs_Hz_out*int32(3) == Fs_Hz_in*int32(2) { /* Fs_out : Fs_in = 2 : 3 */
					(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFIR_Fracs = int32(2)
					(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFIR_Order = int32(RESAMPLER_DOWN_ORDER_FIR0)
					(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FCoefs = uintptr(unsafe.Pointer(&Opus_silk_Resampler_2_3_COEFS))
				} else {
					if Fs_Hz_out*int32(2) == Fs_Hz_in { /* Fs_out : Fs_in = 1 : 2 */
						(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFIR_Fracs = int32(1)
						(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFIR_Order = int32(RESAMPLER_DOWN_ORDER_FIR1)
						(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FCoefs = uintptr(unsafe.Pointer(&Opus_silk_Resampler_1_2_COEFS))
					} else {
						if Fs_Hz_out*int32(3) == Fs_Hz_in { /* Fs_out : Fs_in = 1 : 3 */
							(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFIR_Fracs = int32(1)
							(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFIR_Order = int32(RESAMPLER_DOWN_ORDER_FIR2)
							(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FCoefs = uintptr(unsafe.Pointer(&Opus_silk_Resampler_1_3_COEFS))
						} else {
							if Fs_Hz_out*int32(4) == Fs_Hz_in { /* Fs_out : Fs_in = 1 : 4 */
								(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFIR_Fracs = int32(1)
								(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFIR_Order = int32(RESAMPLER_DOWN_ORDER_FIR2)
								(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FCoefs = uintptr(unsafe.Pointer(&Opus_silk_Resampler_1_4_COEFS))
							} else {
								if Fs_Hz_out*int32(6) == Fs_Hz_in { /* Fs_out : Fs_in = 1 : 6 */
									(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFIR_Fracs = int32(1)
									(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFIR_Order = int32(RESAMPLER_DOWN_ORDER_FIR2)
									(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FCoefs = uintptr(unsafe.Pointer(&Opus_silk_Resampler_1_6_COEFS))
								} else {
									/* None available */
									if !(int32(0) != 0) {
										Opus_celt_fatal(tls, __ccgo_ts+1017, __ccgo_ts+7386, int32(163))
									}
									return -int32(1)
								}
							}
						}
					}
				}
			}
		} else {
			/* Input and output sampling rates are equal: copy */
			(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).Fresampler_function = USE_silk_resampler_copy
		}
	}
	/* Ratio of input/output samples */
	(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FinvRatio_Q16 = int32(uint32(int32(uint32(Fs_Hz_in)<<(int32(14)+up2x))/Fs_Hz_out) << int32(2))
	/* Make sure the ratio is rounded up */
	for int32(int64((*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FinvRatio_Q16)*int64(Fs_Hz_out)>>int32(16)) < int32(uint32(Fs_Hz_in)<<up2x) {
		(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FinvRatio_Q16 = (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FinvRatio_Q16 + 1
	}
	return 0
}

// C documentation
//
//	/* Resampler: convert from one sampling rate to another */
//	/* Input and output sampling rate are at most 48000 Hz  */
func Opus_silk_resampler(tls *libc.TLS, S uintptr, out uintptr, in uintptr, inLen OpusT_opus_int32) (r int32) {
	var nSamples int32
	_ = nSamples
	/* Need at least 1 ms of input data */
	if !(inLen >= (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFs_in_kHz) {
		Opus_celt_fatal(tls, __ccgo_ts+7406, __ccgo_ts+7386, int32(193))
	}
	/* Delay can't exceed the 1 ms of buffering */
	if !((*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FinputDelay <= (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFs_in_kHz) {
		Opus_celt_fatal(tls, __ccgo_ts+7446, __ccgo_ts+7386, int32(195))
	}
	nSamples = (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFs_in_kHz - (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FinputDelay
	/* Copy to delay buffer */
	libc.Xmemcpy(tls, S+168+uintptr((*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FinputDelay)*2, in, uint64(uint32(nSamples))*uint64(2))
	switch (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).Fresampler_function {
	case int32(USE_silk_resampler_private_up2_HQ_wrapper):
		Opus_silk_resampler_private_up2_HQ_wrapper(tls, S, out, S+168, (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFs_in_kHz)
		Opus_silk_resampler_private_up2_HQ_wrapper(tls, S, out+uintptr((*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFs_out_kHz)*2, in+uintptr(nSamples)*2, inLen-(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFs_in_kHz)
	case int32(USE_silk_resampler_private_IIR_FIR):
		Opus_silk_resampler_private_IIR_FIR(tls, S, out, S+168, (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFs_in_kHz)
		Opus_silk_resampler_private_IIR_FIR(tls, S, out+uintptr((*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFs_out_kHz)*2, in+uintptr(nSamples)*2, inLen-(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFs_in_kHz)
	case int32(USE_silk_resampler_private_down_FIR):
		Opus_silk_resampler_private_down_FIR(tls, S, out, S+168, (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFs_in_kHz)
		Opus_silk_resampler_private_down_FIR(tls, S, out+uintptr((*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFs_out_kHz)*2, in+uintptr(nSamples)*2, inLen-(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFs_in_kHz)
	default:
		libc.Xmemcpy(tls, out, S+168, uint64(uint32((*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFs_in_kHz))*uint64(2))
		libc.Xmemcpy(tls, out+uintptr((*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFs_out_kHz)*2, in+uintptr(nSamples)*2, uint64(uint32(inLen-(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFs_in_kHz))*uint64(2))
	}
	/* Copy to delay buffer */
	libc.Xmemcpy(tls, S+168, in+uintptr(inLen-(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FinputDelay)*2, uint64(uint32((*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FinputDelay))*uint64(2))
	return 0
}

const ORDER_FIR = 4
const silk_int16_MAX19 = 32767

var silk_resampler_down2_01 = int16(9872)
var silk_resampler_down2_11 = int16(int32(39809) - int32(65536))
var silk_resampler_up2_hq_01 = [3]OpusT_opus_int16{
	0: int16(1746),
	1: int16(14986),
	2: int16(int32(39083) - int32(65536)),
}
var silk_resampler_up2_hq_11 = [3]OpusT_opus_int16{
	0: int16(6854),
	1: int16(25769),
	2: int16(int32(55542) - int32(65536)),
}

// C documentation
//
//	/* Downsample by a factor 2/3, low quality */
func Opus_silk_resampler_down2_3(tls *libc.TLS, S uintptr, out uintptr, in uintptr, inLen OpusT_opus_int32) {
	var _saved_stack, buf, buf_ptr, st, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var counter, nSamplesIn, res_Q6 OpusT_opus_int32
	var v29, v31 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = _saved_stack, buf, buf_ptr, counter, nSamplesIn, res_Q6, st, v1, v11, v13, v15, v17, v19, v21, v23, v29, v3, v31, v5, v7, v9
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
	if !(int64(int32(uint64(uint32(int32(RESAMPLER_MAX_BATCH_SIZE_MS)*int32(RESAMPLER_MAX_FS_KHZ)+int32(ORDER_FIR)))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+7494, int32(51))
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
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32(int32(RESAMPLER_MAX_BATCH_SIZE_MS)*int32(RESAMPLER_MAX_FS_KHZ)+int32(ORDER_FIR))) * (uint64(4) / uint64(1)))
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
	buf = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32(int32(RESAMPLER_MAX_BATCH_SIZE_MS)*int32(RESAMPLER_MAX_FS_KHZ)+int32(ORDER_FIR)))*(uint64(4)/uint64(1)))
	/* Copy buffered samples to start of buffer */
	libc.Xmemcpy(tls, buf, S, uint64(uint32(ORDER_FIR))*uint64(4))
	/* Iterate over blocks of frameSizeIn input samples */
	for int32(1) != 0 {
		if inLen < int32(RESAMPLER_MAX_BATCH_SIZE_MS)*int32(RESAMPLER_MAX_FS_KHZ) {
			v29 = inLen
		} else {
			v29 = int32(RESAMPLER_MAX_BATCH_SIZE_MS) * int32(RESAMPLER_MAX_FS_KHZ)
		}
		nSamplesIn = v29
		/* Second-order AR filter (output in Q8) */
		Opus_silk_resampler_private_AR2(tls, S+4*4, buf+4*4, in, uintptr(unsafe.Pointer(&Opus_silk_Resampler_2_3_COEFS_LQ)), nSamplesIn)
		/* Interpolate filtered signal */
		buf_ptr = buf
		counter = nSamplesIn
		for counter > int32(2) {
			/* Inner product */
			res_Q6 = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr))) * int64(Opus_silk_Resampler_2_3_COEFS_LQ[int32(2)]) >> int32(16))
			res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 1*4)))*int64(Opus_silk_Resampler_2_3_COEFS_LQ[int32(3)])>>int32(16))
			res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 2*4)))*int64(Opus_silk_Resampler_2_3_COEFS_LQ[int32(5)])>>int32(16))
			res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 3*4)))*int64(Opus_silk_Resampler_2_3_COEFS_LQ[int32(4)])>>int32(16))
			/* Scale down, saturate and store in output array */
			v1 = out
			out += 2
			if (res_Q6>>(int32(6)-int32(1))+int32(1))>>int32(1) > int32(silk_int16_MAX19) {
				v29 = int32(silk_int16_MAX19)
			} else {
				if (res_Q6>>(int32(6)-int32(1))+int32(1))>>int32(1) < int32(int16(-32768)) {
					v31 = int32(int16(-32768))
				} else {
					v31 = (res_Q6>>(int32(6)-int32(1)) + int32(1)) >> int32(1)
				}
				v29 = v31
			}
			*(*OpusT_opus_int16)(unsafe.Pointer(v1)) = int16(v29)
			res_Q6 = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 1*4))) * int64(Opus_silk_Resampler_2_3_COEFS_LQ[int32(4)]) >> int32(16))
			res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 2*4)))*int64(Opus_silk_Resampler_2_3_COEFS_LQ[int32(5)])>>int32(16))
			res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 3*4)))*int64(Opus_silk_Resampler_2_3_COEFS_LQ[int32(3)])>>int32(16))
			res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 4*4)))*int64(Opus_silk_Resampler_2_3_COEFS_LQ[int32(2)])>>int32(16))
			/* Scale down, saturate and store in output array */
			v1 = out
			out += 2
			if (res_Q6>>(int32(6)-int32(1))+int32(1))>>int32(1) > int32(silk_int16_MAX19) {
				v29 = int32(silk_int16_MAX19)
			} else {
				if (res_Q6>>(int32(6)-int32(1))+int32(1))>>int32(1) < int32(int16(-32768)) {
					v31 = int32(int16(-32768))
				} else {
					v31 = (res_Q6>>(int32(6)-int32(1)) + int32(1)) >> int32(1)
				}
				v29 = v31
			}
			*(*OpusT_opus_int16)(unsafe.Pointer(v1)) = int16(v29)
			buf_ptr = buf_ptr + uintptr(3)*4
			counter = counter - int32(3)
		}
		in = in + uintptr(nSamplesIn)*2
		inLen = inLen - nSamplesIn
		if inLen > 0 {
			/* More iterations to do; copy last part of filtered signal to beginning of buffer */
			libc.Xmemcpy(tls, buf, buf+uintptr(nSamplesIn)*4, uint64(uint32(ORDER_FIR))*uint64(4))
		} else {
			break
		}
	}
	/* Copy last part of filtered signal to the state for the next call */
	libc.Xmemcpy(tls, S, buf+uintptr(nSamplesIn)*4, uint64(uint32(ORDER_FIR))*uint64(4))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _37
	_37:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _39
_39:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
}

var silk_resampler_down2_02 = int16(9872)
var silk_resampler_down2_12 = int16(int32(39809) - int32(65536))
var silk_resampler_up2_hq_02 = [3]OpusT_opus_int16{
	0: int16(1746),
	1: int16(14986),
	2: int16(int32(39083) - int32(65536)),
}
var silk_resampler_up2_hq_12 = [3]OpusT_opus_int16{
	0: int16(6854),
	1: int16(25769),
	2: int16(int32(55542) - int32(65536)),
}

// C documentation
//
//	/* Downsample by a factor 2 */
func Opus_silk_resampler_down2(tls *libc.TLS, S uintptr, out uintptr, in uintptr, inLen OpusT_opus_int32) {
	var X, Y, in32, k, len2, out32 OpusT_opus_int32
	var v2, v3 int32
	_, _, _, _, _, _, _, _ = X, Y, in32, k, len2, out32, v2, v3
	len2 = inLen >> int32(1)
	if !(int32(silk_resampler_down2_02) > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+7522, __ccgo_ts+7567, int32(46))
	}
	if !(int32(silk_resampler_down2_12) < int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+7593, __ccgo_ts+7567, int32(47))
	}
	/* Internal variables and state are in Q10 format */
	k = 0
	for {
		if !(k < len2) {
			break
		}
		/* Convert to Q10 */
		in32 = int32(uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(in + uintptr(int32(2)*k)*2)))) << int32(10))
		/* All-pass section for even input sample */
		Y = in32 - *(*OpusT_opus_int32)(unsafe.Pointer(S))
		X = int32(int64(Y) + int64(Y)*int64(silk_resampler_down2_12)>>int32(16))
		out32 = *(*OpusT_opus_int32)(unsafe.Pointer(S)) + X
		*(*OpusT_opus_int32)(unsafe.Pointer(S)) = in32 + X
		/* Convert to Q10 */
		in32 = int32(uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(in + uintptr(int32(2)*k+int32(1))*2)))) << int32(10))
		/* All-pass section for odd input sample, and add to output of previous section */
		Y = in32 - *(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4))
		X = int32(int64(Y) * int64(silk_resampler_down2_02) >> int32(16))
		out32 = out32 + *(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4))
		out32 = out32 + X
		*(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4)) = in32 + X
		/* Add, convert back to int16 and store to output */
		if (out32>>(int32(11)-int32(1))+int32(1))>>int32(1) > int32(silk_int16_MAX19) {
			v2 = int32(silk_int16_MAX19)
		} else {
			if (out32>>(int32(11)-int32(1))+int32(1))>>int32(1) < int32(int16(-32768)) {
				v3 = int32(int16(-32768))
			} else {
				v3 = (out32>>(int32(11)-int32(1)) + int32(1)) >> int32(1)
			}
			v2 = v3
		}
		*(*OpusT_opus_int16)(unsafe.Pointer(out + uintptr(k)*2)) = int16(v2)
		goto _1
	_1:
		;
		k = k + 1
	}
}

const silk_int16_MAX20 = 0x7FFF

var silk_resampler_down2_03 = int16(9872)
var silk_resampler_down2_13 = int16(int32(39809) - int32(65536))
var silk_resampler_up2_hq_03 = [3]OpusT_opus_int16{
	0: int16(1746),
	1: int16(14986),
	2: int16(int32(39083) - int32(65536)),
}
var silk_resampler_up2_hq_13 = [3]OpusT_opus_int16{
	0: int16(6854),
	1: int16(25769),
	2: int16(int32(55542) - int32(65536)),
}

// C documentation
//
//	/* Second order AR filter with single delay elements */
func Opus_silk_resampler_private_AR2(tls *libc.TLS, S uintptr, out_Q8 uintptr, in uintptr, A_Q14 uintptr, len1 OpusT_opus_int32) {
	var k, out32 OpusT_opus_int32
	_, _ = k, out32
	k = 0
	for {
		if !(k < len1) {
			break
		}
		out32 = *(*OpusT_opus_int32)(unsafe.Pointer(S)) + int32(uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(in + uintptr(k)*2))))<<int32(8))
		*(*OpusT_opus_int32)(unsafe.Pointer(out_Q8 + uintptr(k)*4)) = out32
		out32 = int32(uint32(out32) << int32(2))
		*(*OpusT_opus_int32)(unsafe.Pointer(S)) = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4))) + int64(out32)*int64(*(*OpusT_opus_int16)(unsafe.Pointer(A_Q14)))>>int32(16))
		*(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4)) = int32(int64(out32) * int64(*(*OpusT_opus_int16)(unsafe.Pointer(A_Q14 + 1*2))) >> int32(16))
		goto _1
	_1:
		;
		k = k + 1
	}
}

const silk_int16_MAX21 = 32767

var silk_resampler_down2_04 = int16(9872)
var silk_resampler_down2_14 = int16(int32(39809) - int32(65536))
var silk_resampler_up2_hq_04 = [3]OpusT_opus_int16{
	0: int16(1746),
	1: int16(14986),
	2: int16(int32(39083) - int32(65536)),
}
var silk_resampler_up2_hq_14 = [3]OpusT_opus_int16{
	0: int16(6854),
	1: int16(25769),
	2: int16(int32(55542) - int32(65536)),
}

func silk_resampler_private_down_FIR_INTERPOL(tls *libc.TLS, out uintptr, buf uintptr, FIR_Coefs uintptr, FIR_Order int32, FIR_Fracs int32, max_index_Q16 OpusT_opus_int32, index_increment_Q16 OpusT_opus_int32) (r uintptr) {
	var buf_ptr, interpol_ptr, v9 uintptr
	var index_Q16, interpol_ind, res_Q6 OpusT_opus_int32
	var v10, v11 int32
	_, _, _, _, _, _, _, _ = buf_ptr, index_Q16, interpol_ind, interpol_ptr, res_Q6, v10, v11, v9
	switch FIR_Order {
	case int32(RESAMPLER_DOWN_ORDER_FIR0):
		goto _1
	case int32(RESAMPLER_DOWN_ORDER_FIR1):
		goto _2
	case int32(RESAMPLER_DOWN_ORDER_FIR2):
		goto _3
	default:
		goto _4
	}
	goto _5
_1:
	;
	index_Q16 = 0
_8:
	;
	if !(index_Q16 < max_index_Q16) {
		goto _6
	}
	/* Integer part gives pointer to buffered input */
	buf_ptr = buf + uintptr(index_Q16>>int32(16))*4
	/* Fractional part gives interpolation coefficients */
	interpol_ind = int32(int64(index_Q16&int32(0xFFFF)) * int64(int16(FIR_Fracs)) >> int32(16))
	/* Inner product */
	interpol_ptr = FIR_Coefs + uintptr(int32(RESAMPLER_DOWN_ORDER_FIR0)/int32(2)*interpol_ind)*2
	res_Q6 = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr))) * int64(*(*OpusT_opus_int16)(unsafe.Pointer(interpol_ptr))) >> int32(16))
	res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 1*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(interpol_ptr + 1*2)))>>int32(16))
	res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 2*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(interpol_ptr + 2*2)))>>int32(16))
	res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 3*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(interpol_ptr + 3*2)))>>int32(16))
	res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 4*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(interpol_ptr + 4*2)))>>int32(16))
	res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 5*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(interpol_ptr + 5*2)))>>int32(16))
	res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 6*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(interpol_ptr + 6*2)))>>int32(16))
	res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 7*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(interpol_ptr + 7*2)))>>int32(16))
	res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 8*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(interpol_ptr + 8*2)))>>int32(16))
	interpol_ptr = FIR_Coefs + uintptr(int32(RESAMPLER_DOWN_ORDER_FIR0)/int32(2)*(FIR_Fracs-int32(1)-interpol_ind))*2
	res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 17*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(interpol_ptr)))>>int32(16))
	res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 16*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(interpol_ptr + 1*2)))>>int32(16))
	res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 15*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(interpol_ptr + 2*2)))>>int32(16))
	res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 14*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(interpol_ptr + 3*2)))>>int32(16))
	res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 13*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(interpol_ptr + 4*2)))>>int32(16))
	res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 12*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(interpol_ptr + 5*2)))>>int32(16))
	res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 11*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(interpol_ptr + 6*2)))>>int32(16))
	res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 10*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(interpol_ptr + 7*2)))>>int32(16))
	res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 9*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(interpol_ptr + 8*2)))>>int32(16))
	/* Scale down, saturate and store in output array */
	v9 = out
	out += 2
	if (res_Q6>>(int32(6)-int32(1))+int32(1))>>int32(1) > int32(silk_int16_MAX21) {
		v10 = int32(silk_int16_MAX21)
	} else {
		if (res_Q6>>(int32(6)-int32(1))+int32(1))>>int32(1) < int32(int16(-32768)) {
			v11 = int32(int16(-32768))
		} else {
			v11 = (res_Q6>>(int32(6)-int32(1)) + int32(1)) >> int32(1)
		}
		v10 = v11
	}
	*(*OpusT_opus_int16)(unsafe.Pointer(v9)) = int16(v10)
	goto _7
_7:
	;
	index_Q16 = index_Q16 + index_increment_Q16
	goto _8
	goto _6
_6:
	;
	goto _5
_2:
	;
	index_Q16 = 0
	for {
		if !(index_Q16 < max_index_Q16) {
			break
		}
		/* Integer part gives pointer to buffered input */
		buf_ptr = buf + uintptr(index_Q16>>int32(16))*4
		/* Inner product */
		res_Q6 = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 23*4))) * int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs))) >> int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 1*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 22*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 1*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 2*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 21*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 2*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 3*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 20*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 3*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 4*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 19*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 4*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 5*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 18*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 5*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 6*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 17*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 6*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 7*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 16*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 7*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 8*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 15*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 8*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 9*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 14*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 9*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 10*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 13*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 10*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 11*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 12*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 11*2)))>>int32(16))
		/* Scale down, saturate and store in output array */
		v9 = out
		out += 2
		if (res_Q6>>(int32(6)-int32(1))+int32(1))>>int32(1) > int32(silk_int16_MAX21) {
			v10 = int32(silk_int16_MAX21)
		} else {
			if (res_Q6>>(int32(6)-int32(1))+int32(1))>>int32(1) < int32(int16(-32768)) {
				v11 = int32(int16(-32768))
			} else {
				v11 = (res_Q6>>(int32(6)-int32(1)) + int32(1)) >> int32(1)
			}
			v10 = v11
		}
		*(*OpusT_opus_int16)(unsafe.Pointer(v9)) = int16(v10)
		goto _12
	_12:
		;
		index_Q16 = index_Q16 + index_increment_Q16
	}
	goto _5
_3:
	;
	index_Q16 = 0
	for {
		if !(index_Q16 < max_index_Q16) {
			break
		}
		/* Integer part gives pointer to buffered input */
		buf_ptr = buf + uintptr(index_Q16>>int32(16))*4
		/* Inner product */
		res_Q6 = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 35*4))) * int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs))) >> int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 1*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 34*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 1*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 2*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 33*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 2*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 3*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 32*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 3*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 4*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 31*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 4*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 5*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 30*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 5*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 6*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 29*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 6*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 7*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 28*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 7*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 8*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 27*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 8*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 9*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 26*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 9*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 10*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 25*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 10*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 11*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 24*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 11*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 12*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 23*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 12*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 13*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 22*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 13*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 14*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 21*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 14*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 15*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 20*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 15*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 16*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 19*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 16*2)))>>int32(16))
		res_Q6 = int32(int64(res_Q6) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 17*4))+*(*OpusT_opus_int32)(unsafe.Pointer(buf_ptr + 18*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(FIR_Coefs + 17*2)))>>int32(16))
		/* Scale down, saturate and store in output array */
		v9 = out
		out += 2
		if (res_Q6>>(int32(6)-int32(1))+int32(1))>>int32(1) > int32(silk_int16_MAX21) {
			v10 = int32(silk_int16_MAX21)
		} else {
			if (res_Q6>>(int32(6)-int32(1))+int32(1))>>int32(1) < int32(int16(-32768)) {
				v11 = int32(int16(-32768))
			} else {
				v11 = (res_Q6>>(int32(6)-int32(1)) + int32(1)) >> int32(1)
			}
			v10 = v11
		}
		*(*OpusT_opus_int16)(unsafe.Pointer(v9)) = int16(v10)
		goto _16
	_16:
		;
		index_Q16 = index_Q16 + index_increment_Q16
	}
	goto _5
_4:
	;
	if !(int32(0) != 0) {
		Opus_celt_fatal(tls, __ccgo_ts+1017, __ccgo_ts+7638, int32(139))
	}
_5:
	;
	return out
}

// C documentation
//
//	/* Resample with a 2nd order AR filter followed by FIR interpolation */
func Opus_silk_resampler_private_down_FIR(tls *libc.TLS, SS uintptr, out uintptr, in uintptr, inLen OpusT_opus_int32) {
	var FIR_Coefs, S, _saved_stack, buf, st, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var index_increment_Q16, max_index_Q16, nSamplesIn OpusT_opus_int32
	var v29 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = FIR_Coefs, S, _saved_stack, buf, index_increment_Q16, max_index_Q16, nSamplesIn, st, v1, v11, v13, v15, v17, v19, v21, v23, v29, v3, v5, v7, v9
	S = SS
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
	if !(int64(int32(uint64(uint32((*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FbatchSize+(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFIR_Order))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+7638, int32(159))
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
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32((*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FbatchSize+(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFIR_Order)) * (uint64(4) / uint64(1)))
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
	buf = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32((*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FbatchSize+(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFIR_Order))*(uint64(4)/uint64(1)))
	/* Copy buffered samples to start of buffer */
	libc.Xmemcpy(tls, buf, S+24, uint64(uint32((*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFIR_Order))*uint64(4))
	FIR_Coefs = (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FCoefs + 2*2
	/* Iterate over blocks of frameSizeIn input samples */
	index_increment_Q16 = (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FinvRatio_Q16
	for int32(1) != 0 {
		if inLen < (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FbatchSize {
			v29 = inLen
		} else {
			v29 = (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FbatchSize
		}
		nSamplesIn = v29
		/* Second-order AR filter (output in Q8) */
		Opus_silk_resampler_private_AR2(tls, S, buf+uintptr((*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFIR_Order)*4, in, (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FCoefs, nSamplesIn)
		max_index_Q16 = int32(uint32(nSamplesIn) << int32(16))
		/* Interpolate filtered signal */
		out = silk_resampler_private_down_FIR_INTERPOL(tls, out, buf, FIR_Coefs, (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFIR_Order, (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFIR_Fracs, max_index_Q16, index_increment_Q16)
		in = in + uintptr(nSamplesIn)*2
		inLen = inLen - nSamplesIn
		if inLen > int32(1) {
			/* More iterations to do; copy last part of filtered signal to beginning of buffer */
			libc.Xmemcpy(tls, buf, buf+uintptr(nSamplesIn)*4, uint64(uint32((*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFIR_Order))*uint64(4))
		} else {
			break
		}
	}
	/* Copy last part of filtered signal to the state for the next call */
	libc.Xmemcpy(tls, S+24, buf+uintptr(nSamplesIn)*4, uint64(uint32((*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FFIR_Order))*uint64(4))
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
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
}

var silk_resampler_down2_05 = int16(9872)
var silk_resampler_down2_15 = int16(int32(39809) - int32(65536))
var silk_resampler_up2_hq_05 = [3]OpusT_opus_int16{
	0: int16(1746),
	1: int16(14986),
	2: int16(int32(39083) - int32(65536)),
}
var silk_resampler_up2_hq_15 = [3]OpusT_opus_int16{
	0: int16(6854),
	1: int16(25769),
	2: int16(int32(55542) - int32(65536)),
}

func silk_resampler_private_IIR_FIR_INTERPOL(tls *libc.TLS, out uintptr, buf uintptr, max_index_Q16 OpusT_opus_int32, index_increment_Q16 OpusT_opus_int32) (r uintptr) {
	var buf_ptr, v2 uintptr
	var index_Q16, res_Q15, table_index OpusT_opus_int32
	var v3, v4 int32
	_, _, _, _, _, _, _ = buf_ptr, index_Q16, res_Q15, table_index, v2, v3, v4
	/* Interpolate upsampled signal and store in output array */
	index_Q16 = 0
	for {
		if !(index_Q16 < max_index_Q16) {
			break
		}
		table_index = int32(int64(index_Q16&int32(0xFFFF)) * int64(int16(int32(12))) >> int32(16))
		buf_ptr = buf + uintptr(index_Q16>>int32(16))*2
		res_Q15 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(buf_ptr))) * int32(*(*OpusT_opus_int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_silk_resampler_frac_FIR_12)) + uintptr(table_index)*8)))
		res_Q15 = res_Q15 + int32(*(*OpusT_opus_int16)(unsafe.Pointer(buf_ptr + 1*2)))*int32(*(*OpusT_opus_int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_silk_resampler_frac_FIR_12)) + uintptr(table_index)*8 + 1*2)))
		res_Q15 = res_Q15 + int32(*(*OpusT_opus_int16)(unsafe.Pointer(buf_ptr + 2*2)))*int32(*(*OpusT_opus_int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_silk_resampler_frac_FIR_12)) + uintptr(table_index)*8 + 2*2)))
		res_Q15 = res_Q15 + int32(*(*OpusT_opus_int16)(unsafe.Pointer(buf_ptr + 3*2)))*int32(*(*OpusT_opus_int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_silk_resampler_frac_FIR_12)) + uintptr(table_index)*8 + 3*2)))
		res_Q15 = res_Q15 + int32(*(*OpusT_opus_int16)(unsafe.Pointer(buf_ptr + 4*2)))*int32(*(*OpusT_opus_int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_silk_resampler_frac_FIR_12)) + uintptr(int32(11)-table_index)*8 + 3*2)))
		res_Q15 = res_Q15 + int32(*(*OpusT_opus_int16)(unsafe.Pointer(buf_ptr + 5*2)))*int32(*(*OpusT_opus_int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_silk_resampler_frac_FIR_12)) + uintptr(int32(11)-table_index)*8 + 2*2)))
		res_Q15 = res_Q15 + int32(*(*OpusT_opus_int16)(unsafe.Pointer(buf_ptr + 6*2)))*int32(*(*OpusT_opus_int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_silk_resampler_frac_FIR_12)) + uintptr(int32(11)-table_index)*8 + 1*2)))
		res_Q15 = res_Q15 + int32(*(*OpusT_opus_int16)(unsafe.Pointer(buf_ptr + 7*2)))*int32(*(*OpusT_opus_int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_silk_resampler_frac_FIR_12)) + uintptr(int32(11)-table_index)*8)))
		v2 = out
		out += 2
		if (res_Q15>>(int32(15)-int32(1))+int32(1))>>int32(1) > int32(silk_int16_MAX21) {
			v3 = int32(silk_int16_MAX21)
		} else {
			if (res_Q15>>(int32(15)-int32(1))+int32(1))>>int32(1) < int32(int16(-32768)) {
				v4 = int32(int16(-32768))
			} else {
				v4 = (res_Q15>>(int32(15)-int32(1)) + int32(1)) >> int32(1)
			}
			v3 = v4
		}
		*(*OpusT_opus_int16)(unsafe.Pointer(v2)) = int16(v3)
		goto _1
	_1:
		;
		index_Q16 = index_Q16 + index_increment_Q16
	}
	return out
}

// C documentation
//
//	/* Upsample using a combination of allpass-based 2x upsampling and FIR interpolation */
func Opus_silk_resampler_private_IIR_FIR(tls *libc.TLS, SS uintptr, out uintptr, in uintptr, inLen OpusT_opus_int32) {
	var S, _saved_stack, buf, st, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var index_increment_Q16, max_index_Q16, nSamplesIn OpusT_opus_int32
	var v29 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = S, _saved_stack, buf, index_increment_Q16, max_index_Q16, nSamplesIn, st, v1, v11, v13, v15, v17, v19, v21, v23, v29, v3, v5, v7, v9
	S = SS
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
	if !(int64(int32(uint64(uint32(int32(2)*(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FbatchSize+int32(RESAMPLER_ORDER_FIR_12)))*(uint64(2)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+7675, int32(78))
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
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32(int32(2)*(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FbatchSize+int32(RESAMPLER_ORDER_FIR_12))) * (uint64(2) / uint64(1)))
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
	buf = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32(int32(2)*(*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FbatchSize+int32(RESAMPLER_ORDER_FIR_12)))*(uint64(2)/uint64(1)))
	/* Copy buffered samples to start of buffer */
	libc.Xmemcpy(tls, buf, S+24, uint64(uint32(RESAMPLER_ORDER_FIR_12))*uint64(2))
	/* Iterate over blocks of frameSizeIn input samples */
	index_increment_Q16 = (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FinvRatio_Q16
	for int32(1) != 0 {
		if inLen < (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FbatchSize {
			v29 = inLen
		} else {
			v29 = (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S)).FbatchSize
		}
		nSamplesIn = v29
		/* Upsample 2x */
		Opus_silk_resampler_private_up2_HQ(tls, S, buf+8*2, in, nSamplesIn)
		max_index_Q16 = int32(uint32(nSamplesIn) << (int32(16) + int32(1))) /* + 1 because 2x upsampling */
		out = silk_resampler_private_IIR_FIR_INTERPOL(tls, out, buf, max_index_Q16, index_increment_Q16)
		in = in + uintptr(nSamplesIn)*2
		inLen = inLen - nSamplesIn
		if inLen > 0 {
			/* More iterations to do; copy last part of filtered signal to beginning of buffer */
			libc.Xmemcpy(tls, buf, buf+uintptr(nSamplesIn<<int32(1))*2, uint64(uint32(RESAMPLER_ORDER_FIR_12))*uint64(2))
		} else {
			break
		}
	}
	/* Copy last part of filtered signal to the state for the next call */
	libc.Xmemcpy(tls, S+24, buf+uintptr(nSamplesIn<<int32(1))*2, uint64(uint32(RESAMPLER_ORDER_FIR_12))*uint64(2))
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
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
}

var silk_resampler_down2_06 = int16(9872)
var silk_resampler_down2_16 = int16(int32(39809) - int32(65536))
var silk_resampler_up2_hq_06 = [3]OpusT_opus_int16{
	0: int16(1746),
	1: int16(14986),
	2: int16(int32(39083) - int32(65536)),
}
var silk_resampler_up2_hq_16 = [3]OpusT_opus_int16{
	0: int16(6854),
	1: int16(25769),
	2: int16(int32(55542) - int32(65536)),
}

// C documentation
//
//	/* Upsample by a factor 2, high quality */
//	/* Uses 2nd order allpass filters for the 2x upsampling, followed by a      */
//	/* notch filter just above Nyquist.                                         */
func Opus_silk_resampler_private_up2_HQ(tls *libc.TLS, S uintptr, out uintptr, in uintptr, len1 OpusT_opus_int32) {
	var X, Y, in32, k, out32_1, out32_2 OpusT_opus_int32
	var v2, v3 int32
	_, _, _, _, _, _, _, _ = X, Y, in32, k, out32_1, out32_2, v2, v3
	_ = int32(silk_resampler_up2_hq_06[0]) > int32(0)
	_ = int32(silk_resampler_up2_hq_06[int32(1)]) > int32(0)
	_ = int32(silk_resampler_up2_hq_06[int32(2)]) < int32(0)
	_ = int32(silk_resampler_up2_hq_16[0]) > int32(0)
	_ = int32(silk_resampler_up2_hq_16[int32(1)]) > int32(0)
	_ = int32(silk_resampler_up2_hq_16[int32(2)]) < int32(0)
	/* Internal variables and state are in Q10 format */
	k = 0
	for {
		if !(k < len1) {
			break
		}
		/* Convert to Q10 */
		in32 = int32(uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(in + uintptr(k)*2)))) << int32(10))
		/* First all-pass section for even output sample */
		Y = in32 - *(*OpusT_opus_int32)(unsafe.Pointer(S))
		X = int32(int64(Y) * int64(silk_resampler_up2_hq_06[0]) >> int32(16))
		out32_1 = *(*OpusT_opus_int32)(unsafe.Pointer(S)) + X
		*(*OpusT_opus_int32)(unsafe.Pointer(S)) = in32 + X
		/* Second all-pass section for even output sample */
		Y = out32_1 - *(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4))
		X = int32(int64(Y) * int64(silk_resampler_up2_hq_06[int32(1)]) >> int32(16))
		out32_2 = *(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4)) + X
		*(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4)) = out32_1 + X
		/* Third all-pass section for even output sample */
		Y = out32_2 - *(*OpusT_opus_int32)(unsafe.Pointer(S + 2*4))
		X = int32(int64(Y) + int64(Y)*int64(silk_resampler_up2_hq_06[int32(2)])>>int32(16))
		out32_1 = *(*OpusT_opus_int32)(unsafe.Pointer(S + 2*4)) + X
		*(*OpusT_opus_int32)(unsafe.Pointer(S + 2*4)) = out32_2 + X
		/* Apply gain in Q15, convert back to int16 and store to output */
		if (out32_1>>(int32(10)-int32(1))+int32(1))>>int32(1) > int32(silk_int16_MAX21) {
			v2 = int32(silk_int16_MAX21)
		} else {
			if (out32_1>>(int32(10)-int32(1))+int32(1))>>int32(1) < int32(int16(-32768)) {
				v3 = int32(int16(-32768))
			} else {
				v3 = (out32_1>>(int32(10)-int32(1)) + int32(1)) >> int32(1)
			}
			v2 = v3
		}
		*(*OpusT_opus_int16)(unsafe.Pointer(out + uintptr(int32(2)*k)*2)) = int16(v2)
		/* First all-pass section for odd output sample */
		Y = in32 - *(*OpusT_opus_int32)(unsafe.Pointer(S + 3*4))
		X = int32(int64(Y) * int64(silk_resampler_up2_hq_16[0]) >> int32(16))
		out32_1 = *(*OpusT_opus_int32)(unsafe.Pointer(S + 3*4)) + X
		*(*OpusT_opus_int32)(unsafe.Pointer(S + 3*4)) = in32 + X
		/* Second all-pass section for odd output sample */
		Y = out32_1 - *(*OpusT_opus_int32)(unsafe.Pointer(S + 4*4))
		X = int32(int64(Y) * int64(silk_resampler_up2_hq_16[int32(1)]) >> int32(16))
		out32_2 = *(*OpusT_opus_int32)(unsafe.Pointer(S + 4*4)) + X
		*(*OpusT_opus_int32)(unsafe.Pointer(S + 4*4)) = out32_1 + X
		/* Third all-pass section for odd output sample */
		Y = out32_2 - *(*OpusT_opus_int32)(unsafe.Pointer(S + 5*4))
		X = int32(int64(Y) + int64(Y)*int64(silk_resampler_up2_hq_16[int32(2)])>>int32(16))
		out32_1 = *(*OpusT_opus_int32)(unsafe.Pointer(S + 5*4)) + X
		*(*OpusT_opus_int32)(unsafe.Pointer(S + 5*4)) = out32_2 + X
		/* Apply gain in Q15, convert back to int16 and store to output */
		if (out32_1>>(int32(10)-int32(1))+int32(1))>>int32(1) > int32(silk_int16_MAX21) {
			v2 = int32(silk_int16_MAX21)
		} else {
			if (out32_1>>(int32(10)-int32(1))+int32(1))>>int32(1) < int32(int16(-32768)) {
				v3 = int32(int16(-32768))
			} else {
				v3 = (out32_1>>(int32(10)-int32(1)) + int32(1)) >> int32(1)
			}
			v2 = v3
		}
		*(*OpusT_opus_int16)(unsafe.Pointer(out + uintptr(int32(2)*k+int32(1))*2)) = int16(v2)
		goto _1
	_1:
		;
		k = k + 1
	}
}

func Opus_silk_resampler_private_up2_HQ_wrapper(tls *libc.TLS, SS uintptr, out uintptr, in uintptr, len1 OpusT_opus_int32) {
	var S uintptr
	_ = S
	S = SS
	Opus_silk_resampler_private_up2_HQ(tls, S, out, in, len1)
}

const silk_int16_MAX22 = 0x7FFF

var silk_resampler_down2_07 = int16(9872)
var silk_resampler_down2_17 = int16(int32(39809) - int32(65536))
var silk_resampler_up2_hq_07 = [3]OpusT_opus_int16{
	0: int16(1746),
	1: int16(14986),
	2: int16(int32(39083) - int32(65536)),
}
var silk_resampler_up2_hq_17 = [3]OpusT_opus_int16{
	0: int16(6854),
	1: int16(25769),
	2: int16(int32(55542) - int32(65536)),
}

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

/***********************************************************************
Copyright (c) 2006-2011, Skype Limited. All rights reserved.
Copyright (C) 2012 Xiph.Org Foundation
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

/* Redefine macro functions with extensive assertion in DEBUG mode.
   As functions can't be undefined, this file can't work with SigProcFIX_MacroCount.h */

// C documentation
//
//	/* fprintf(1, '%d, ', round(1024 * ([1 ./ (1 + exp(-(1:5))), 1] - 1 ./ (1 + exp(-(0:5)))))); */
var sigm_LUT_slope_Q10 = [6]OpusT_opus_int32{
	0: int32(237),
	1: int32(153),
	2: int32(73),
	3: int32(30),
	4: int32(12),
	5: int32(7),
}

// C documentation
//
//	/* fprintf(1, '%d, ', round(32767 * 1 ./ (1 + exp(-(0:5))))); */
var sigm_LUT_pos_Q15 = [6]OpusT_opus_int32{
	0: int32(16384),
	1: int32(23955),
	2: int32(28861),
	3: int32(31213),
	4: int32(32178),
	5: int32(32548),
}

// C documentation
//
//	/* fprintf(1, '%d, ', round(32767 * 1 ./ (1 + exp((0:5))))); */
var sigm_LUT_neg_Q15 = [6]OpusT_opus_int32{
	0: int32(16384),
	1: int32(8812),
	2: int32(3906),
	3: int32(1554),
	4: int32(589),
	5: int32(219),
}

func Opus_silk_sigm_Q15(tls *libc.TLS, in_Q5 int32) (r int32) {
	var ind int32
	_ = ind
	if in_Q5 < 0 {
		/* Negative input */
		in_Q5 = -in_Q5
		if in_Q5 >= int32(6)*int32(32) {
			return 0 /* Clip */
		} else {
			/* Linear interpolation of look up table */
			ind = in_Q5 >> int32(5)
			return sigm_LUT_neg_Q15[ind] - int32(int16(sigm_LUT_slope_Q10[ind]))*int32(int16(in_Q5&int32(0x1F)))
		}
	} else {
		/* Positive input */
		if in_Q5 >= int32(6)*int32(32) {
			return int32(32767) /* clip */
		} else {
			/* Linear interpolation of look up table */
			ind = in_Q5 >> int32(5)
			return sigm_LUT_pos_Q15[ind] + int32(int16(sigm_LUT_slope_Q10[ind]))*int32(int16(in_Q5&int32(0x1F)))
		}
	}
	return r
}

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

/***********************************************************************
Copyright (c) 2006-2011, Skype Limited. All rights reserved.
Copyright (C) 2012 Xiph.Org Foundation
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

/* Redefine macro functions with extensive assertion in DEBUG mode.
   As functions can't be undefined, this file can't work with SigProcFIX_MacroCount.h */

func Opus_silk_insertion_sort_increasing(tls *libc.TLS, a uintptr, idx uintptr, L int32, K int32) {
	var i, j int32
	var value OpusT_opus_int32
	_, _, _ = i, j, value
	/* Safety checks */
	if !(K > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+7711, __ccgo_ts+7735, int32(51))
	}
	if !(L > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+7750, __ccgo_ts+7735, int32(52))
	}
	if !(L >= K) {
		Opus_celt_fatal(tls, __ccgo_ts+7774, __ccgo_ts+7735, int32(53))
	}
	/* Write start indices in index vector */
	i = 0
	for {
		if !(i < K) {
			break
		}
		*(*int32)(unsafe.Pointer(idx + uintptr(i)*4)) = i
		goto _1
	_1:
		;
		i = i + 1
	}
	/* Sort vector elements by value, increasing order */
	i = int32(1)
	for {
		if !(i < K) {
			break
		}
		value = *(*OpusT_opus_int32)(unsafe.Pointer(a + uintptr(i)*4))
		j = i - int32(1)
		for {
			if !(j >= 0 && value < *(*OpusT_opus_int32)(unsafe.Pointer(a + uintptr(j)*4))) {
				break
			}
			*(*OpusT_opus_int32)(unsafe.Pointer(a + uintptr(j+int32(1))*4)) = *(*OpusT_opus_int32)(unsafe.Pointer(a + uintptr(j)*4)) /* Shift value */
			*(*int32)(unsafe.Pointer(idx + uintptr(j+int32(1))*4)) = *(*int32)(unsafe.Pointer(idx + uintptr(j)*4))                   /* Shift index */
			goto _3
		_3:
			;
			j = j - 1
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(a + uintptr(j+int32(1))*4)) = value /* Write value */
		*(*int32)(unsafe.Pointer(idx + uintptr(j+int32(1))*4)) = i              /* Write index */
		goto _2
	_2:
		;
		i = i + 1
	}
	/* If less than L values are asked for, check the remaining values, */
	/* but only spend CPU to ensure that the K first values are correct */
	i = K
	for {
		if !(i < L) {
			break
		}
		value = *(*OpusT_opus_int32)(unsafe.Pointer(a + uintptr(i)*4))
		if value < *(*OpusT_opus_int32)(unsafe.Pointer(a + uintptr(K-int32(1))*4)) {
			j = K - int32(2)
			for {
				if !(j >= 0 && value < *(*OpusT_opus_int32)(unsafe.Pointer(a + uintptr(j)*4))) {
					break
				}
				*(*OpusT_opus_int32)(unsafe.Pointer(a + uintptr(j+int32(1))*4)) = *(*OpusT_opus_int32)(unsafe.Pointer(a + uintptr(j)*4)) /* Shift value */
				*(*int32)(unsafe.Pointer(idx + uintptr(j+int32(1))*4)) = *(*int32)(unsafe.Pointer(idx + uintptr(j)*4))                   /* Shift index */
				goto _5
			_5:
				;
				j = j - 1
			}
			*(*OpusT_opus_int32)(unsafe.Pointer(a + uintptr(j+int32(1))*4)) = value /* Write value */
			*(*int32)(unsafe.Pointer(idx + uintptr(j+int32(1))*4)) = i              /* Write index */
		}
		goto _4
	_4:
		;
		i = i + 1
	}
}

func Opus_silk_insertion_sort_increasing_all_values_int16(tls *libc.TLS, a uintptr, L int32) {
	var i, j, value int32
	_, _, _ = i, j, value
	/* Safety checks */
	if !(L > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+7750, __ccgo_ts+7735, int32(144))
	}
	/* Sort vector elements by value, increasing order */
	i = int32(1)
	for {
		if !(i < L) {
			break
		}
		value = int32(*(*OpusT_opus_int16)(unsafe.Pointer(a + uintptr(i)*2)))
		j = i - int32(1)
		for {
			if !(j >= 0 && value < int32(*(*OpusT_opus_int16)(unsafe.Pointer(a + uintptr(j)*2)))) {
				break
			}
			*(*OpusT_opus_int16)(unsafe.Pointer(a + uintptr(j+int32(1))*2)) = *(*OpusT_opus_int16)(unsafe.Pointer(a + uintptr(j)*2)) /* Shift value */
			goto _2
		_2:
			;
			j = j - 1
		}
		*(*OpusT_opus_int16)(unsafe.Pointer(a + uintptr(j+int32(1))*2)) = int16(value) /* Write value */
		goto _1
	_1:
		;
		i = i + 1
	}
}

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

/***********************************************************************
Copyright (c) 2006-2011, Skype Limited. All rights reserved.
Copyright (C) 2012 Xiph.Org Foundation
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

/* Redefine macro functions with extensive assertion in DEBUG mode.
   As functions can't be undefined, this file can't work with SigProcFIX_MacroCount.h */

// C documentation
//
//	/* Compute number of bits to right shift the sum of squares of a vector */
//	/* of int16s to make it fit in an int32                                 */
func Opus_silk_sum_sqr_shift(tls *libc.TLS, energy uintptr, shift uintptr, x uintptr, len1 int32) {
	var i, shft, v4, v9 int32
	var nrg, v1, v10, v2, v6, v7 OpusT_opus_int32
	var nrg_tmp OpusT_opus_uint32
	_, _, _, _, _, _, _, _, _, _, _ = i, nrg, nrg_tmp, shft, v1, v10, v2, v4, v6, v7, v9
	/* Do a first run with the maximum shift we could have. */
	v1 = len1
	if v1 != 0 {
		v4 = int32(32) - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, uint32(v1)))
	} else {
		v4 = int32(32)
	}
	v2 = v4
	goto _3
_3:
	shft = int32(31) - v2
	/* Let's be conservative with rounding and start with nrg=len. */
	nrg = len1
	i = 0
	for {
		if !(i < len1-int32(1)) {
			break
		}
		nrg_tmp = uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(x + uintptr(i)*2))) * int32(*(*OpusT_opus_int16)(unsafe.Pointer(x + uintptr(i)*2))))
		nrg_tmp = uint32(int32(nrg_tmp + uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(x + uintptr(i+int32(1))*2)))*int32(*(*OpusT_opus_int16)(unsafe.Pointer(x + uintptr(i+int32(1))*2))))))
		nrg = int32(uint32(nrg) + nrg_tmp>>shft)
		goto _5
	_5:
		;
		i = i + int32(2)
	}
	if i < len1 {
		/* One sample left to process */
		nrg_tmp = uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(x + uintptr(i)*2))) * int32(*(*OpusT_opus_int16)(unsafe.Pointer(x + uintptr(i)*2))))
		nrg = int32(uint32(nrg) + nrg_tmp>>shft)
	}
	_ = nrg >= int32(0)
	/* Make sure the result will fit in a 32-bit signed integer with two bits
	   of headroom. */
	v1 = nrg
	if v1 != 0 {
		v4 = int32(32) - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, uint32(v1)))
	} else {
		v4 = int32(32)
	}
	v2 = v4
	goto _8
_8:
	v6 = 0
	v7 = shft + int32(3) - v2
	if v6 > v7 {
		v9 = v6
	} else {
		v9 = v7
	}
	v10 = v9
	goto _13
_13:
	shft = v10
	nrg = 0
	i = 0
	for {
		if !(i < len1-int32(1)) {
			break
		}
		nrg_tmp = uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(x + uintptr(i)*2))) * int32(*(*OpusT_opus_int16)(unsafe.Pointer(x + uintptr(i)*2))))
		nrg_tmp = uint32(int32(nrg_tmp + uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(x + uintptr(i+int32(1))*2)))*int32(*(*OpusT_opus_int16)(unsafe.Pointer(x + uintptr(i+int32(1))*2))))))
		nrg = int32(uint32(nrg) + nrg_tmp>>shft)
		goto _15
	_15:
		;
		i = i + int32(2)
	}
	if i < len1 {
		/* One sample left to process */
		nrg_tmp = uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(x + uintptr(i)*2))) * int32(*(*OpusT_opus_int16)(unsafe.Pointer(x + uintptr(i)*2))))
		nrg = int32(uint32(nrg) + nrg_tmp>>shft)
	}
	_ = nrg >= int32(0)
	/* Output arguments */
	*(*int32)(unsafe.Pointer(shift)) = shft
	*(*OpusT_opus_int32)(unsafe.Pointer(energy)) = nrg
}

// C documentation
//
//	/* Decode mid/side predictors */
