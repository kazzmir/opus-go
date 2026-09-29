// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

package opuscc

import (
	"math/bits"
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
		state := (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S))
		Opus_silk_resampler_private_up2_HQ_wrapper(tls, state, (*OpusT_opus_int16)(unsafe.Pointer(out)), &state.FdelayBuf[0], state.FFs_in_kHz)
		Opus_silk_resampler_private_up2_HQ_wrapper(tls, state, (*OpusT_opus_int16)(unsafe.Pointer(out+uintptr(state.FFs_out_kHz)*2)), (*OpusT_opus_int16)(unsafe.Pointer(in+uintptr(nSamples)*2)), inLen-state.FFs_in_kHz)
	case int32(USE_silk_resampler_private_IIR_FIR):
		state := (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S))
		Opus_silk_resampler_private_IIR_FIR(tls, state, (*int16)(unsafe.Pointer(out)), &state.FdelayBuf[0], state.FFs_in_kHz)
		Opus_silk_resampler_private_IIR_FIR(tls, state, (*int16)(unsafe.Pointer(out+uintptr(state.FFs_out_kHz)*2)), (*int16)(unsafe.Pointer(in+uintptr(nSamples)*2)), inLen-state.FFs_in_kHz)
	case int32(USE_silk_resampler_private_down_FIR):
		state := (*OpusT_silk_resampler_state_struct)(unsafe.Pointer(S))
		coefs := (*int16)(unsafe.Pointer(state.FCoefs))
		Opus_silk_resampler_private_down_FIR(tls, state, coefs, (*int16)(unsafe.Pointer(out)), &state.FdelayBuf[0], state.FFs_in_kHz)
		Opus_silk_resampler_private_down_FIR(tls, state, coefs, (*int16)(unsafe.Pointer(out+uintptr(state.FFs_out_kHz)*2)), (*int16)(unsafe.Pointer(in+uintptr(nSamples)*2)), inLen-state.FFs_in_kHz)
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
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
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
	if !(int64(int32(uint64(uint32(int32(RESAMPLER_MAX_BATCH_SIZE_MS)*int32(RESAMPLER_MAX_FS_KHZ)+int32(ORDER_FIR)))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+7494, int32(51))
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
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v19)).Fglobal_stack += uintptr(uint64(uint32(int32(RESAMPLER_MAX_BATCH_SIZE_MS)*int32(RESAMPLER_MAX_FS_KHZ)+int32(ORDER_FIR))) * (uint64(4) / uint64(1)))
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
		Opus_silk_resampler_private_AR2(tls, (*OpusT_opus_int32)(unsafe.Pointer(S+4*4)), (*OpusT_opus_int32)(unsafe.Pointer(buf+4*4)), (*OpusT_opus_int16)(unsafe.Pointer(in)), &Opus_silk_Resampler_2_3_COEFS_LQ[0], nSamplesIn)
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
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
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
func Opus_silk_resampler_down2(tls *libc.TLS, S *[2]OpusT_opus_int32, out *OpusT_opus_int16, in *OpusT_opus_int16, inLen OpusT_opus_int32) {
	len2 := int(inLen >> 1)
	if !(int32(silk_resampler_down2_02) > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+7522, __ccgo_ts+7567, int32(46))
	}
	if !(int32(silk_resampler_down2_12) < int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+7593, __ccgo_ts+7567, int32(47))
	}
	if len2 <= 0 {
		return
	}
	// Only complete input pairs are consumed, matching C's floor(inLen/2).
	input := unsafe.Slice(in, 2*len2)
	output := unsafe.Slice(out, len2)
	for k := range output {
		// Internal variables and state are Q10, with 32-bit wraparound.
		in32 := int32(uint32(int32(input[2*k])) << 10)
		Y := in32 - S[0]
		X := int32(int64(Y) + ((int64(Y) * int64(silk_resampler_down2_12)) >> 16))
		out32 := S[0] + X
		S[0] = in32 + X
		in32 = int32(uint32(int32(input[2*k+1])) << 10)
		Y = in32 - S[1]
		X = int32((int64(Y) * int64(silk_resampler_down2_02)) >> 16)
		out32 += S[1]
		out32 += X
		S[1] = in32 + X
		value := ((out32 >> 10) + 1) >> 1
		output[k] = int16(min(max(value, -32768), 32767))
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
func Opus_silk_resampler_private_AR2(tls *libc.TLS, S *OpusT_opus_int32, out_Q8 *OpusT_opus_int32, in *OpusT_opus_int16, A_Q14 *OpusT_opus_int16, len1 OpusT_opus_int32) {
	if len1 <= 0 {
		return
	}
	state := unsafe.Slice(S, 2)
	out := unsafe.Slice(out_Q8, int(len1))
	input := unsafe.Slice(in, int(len1))
	coefs := unsafe.Slice(A_Q14, 2)
	for k, sample := range input {
		out32 := state[0] + int32(uint32(int32(sample))<<8)
		out[k] = out32
		out32 = int32(uint32(out32) << 2)
		state[0] = int32(int64(state[1]) + (int64(out32)*int64(coefs[0]))>>16)
		state[1] = int32((int64(out32) * int64(coefs[1])) >> 16)
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

func silk_resampler_private_down_FIR_INTERPOL(tls *libc.TLS, out *int16, buf *int32, coefs *int16, order, fracs, maxIndex, increment int32) int32 {
	if order != 18 && order != 24 && order != 36 {
		Opus_celt_fatal(tls, __ccgo_ts+1017, __ccgo_ts+7638, 139)
	}
	if maxIndex <= 0 {
		return 0
	}
	count := 1 + (maxIndex-1)/increment
	input := unsafe.Slice(buf, ((count-1)*increment>>16)+order)
	output := unsafe.Slice(out, count)
	coefCount := order / 2
	if order == 18 {
		coefCount *= fracs
	}
	coeff := unsafe.Slice(coefs, coefCount)
	for i, index := int32(0), int32(0); i < count; i, index = i+1, index+increment {
		x := input[index>>16:]
		var sum int32
		if order == 18 {
			phase := int32(int64(index&65535) * int64(int16(fracs)) >> 16)
			for j := int32(0); j < 9; j++ {
				sum += int32(int64(x[j]) * int64(coeff[9*phase+j]) >> 16)
			}
			for j := int32(0); j < 9; j++ {
				sum += int32(int64(x[17-j]) * int64(coeff[9*(fracs-1-phase)+j]) >> 16)
			}
		} else {
			for j := int32(0); j < order/2; j++ {
				// ADD32 narrows before SMULWB, not after the multiply.
				pair := x[j] + x[order-1-j]
				sum += int32(int64(pair) * int64(coeff[j]) >> 16)
			}
		}
		output[i] = int16(max(-32768, min(32767, ((sum>>5)+1)>>1)))
	}
	return count
}

// C documentation
//
//	/* Resample with a 2nd order AR filter followed by FIR interpolation */
//
// Coefficients are explicit so the driver need not recover a Go allocation
// from the legacy state's FCoefs uintptr field.
func Opus_silk_resampler_private_down_FIR(tls *libc.TLS, state *OpusT_silk_resampler_state_struct, coefs, out, in *int16, inLen int32) {
	order := state.FFIR_Order
	buf := make([]int32, state.FbatchSize+order)
	copy(buf, state.FsFIR.Fi32[:order])
	firCoefs := (*int16)(unsafe.Add(unsafe.Pointer(coefs), 4))
	input := unsafe.Slice(in, inLen)
	increment := state.FinvRatio_Q16
	for {
		n := min(int32(len(input)), state.FbatchSize)
		Opus_silk_resampler_private_AR2(tls, &state.FsIIR[0], &buf[order], unsafe.SliceData(input), coefs, n)
		written := silk_resampler_private_down_FIR_INTERPOL(tls, out, &buf[0], firCoefs, order, state.FFIR_Fracs, n<<16, increment)
		input = input[n:]
		// Match C: a lone remainder after a batch is not processed.
		if len(input) <= 1 {
			copy(state.FsFIR.Fi32[:order], buf[n:n+order])
			break
		}
		out = (*int16)(unsafe.Add(unsafe.Pointer(out), int(written)*2))
		copy(buf[:order], buf[n:n+order])
	}
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

// Return the number of output samples, rather than a one-past-end pointer.
func silk_resampler_private_IIR_FIR_INTERPOL(tls *libc.TLS, out, buf *int16, maxIndex, increment int32) int32 {
	if maxIndex <= 0 {
		return 0
	}
	count := 1 + (maxIndex-1)/increment
	input := unsafe.Slice(buf, ((count-1)*increment>>16)+8)
	output := unsafe.Slice(out, count)
	for i, index := int32(0), int32(0); i < count; i, index = i+1, index+increment {
		phase := (index & 65535) * 12 >> 16
		x := input[index>>16:]
		var sum int32
		for j := 0; j < 4; j++ {
			sum += int32(x[j]) * int32(Opus_silk_resampler_frac_FIR_12[phase][j])
		}
		for j := 0; j < 4; j++ {
			sum += int32(x[j+4]) * int32(Opus_silk_resampler_frac_FIR_12[11-phase][3-j])
		}
		output[i] = int16(max(-32768, min(32767, ((sum>>14)+1)>>1)))
	}
	return count
}

// C documentation
//
//	/* Upsample using a combination of allpass-based 2x upsampling and FIR interpolation */
func Opus_silk_resampler_private_IIR_FIR(tls *libc.TLS, state *OpusT_silk_resampler_state_struct, out, in *int16, inLen int32) {
	buf := make([]int16, 2*state.FbatchSize+8)
	// The C union uses its first eight int16 elements in this mode.
	history := (*[8]int16)(unsafe.Pointer(&state.FsFIR.Fi32[0]))
	copy(buf, history[:])
	input := unsafe.Slice(in, inLen)
	increment := state.FinvRatio_Q16
	for {
		n := min(int32(len(input)), state.FbatchSize)
		Opus_silk_resampler_private_up2_HQ(tls, &state.FsIIR, &buf[8], unsafe.SliceData(input), n)
		written := silk_resampler_private_IIR_FIR_INTERPOL(tls, out, &buf[0], n<<17, increment)
		input = input[n:]
		if len(input) == 0 {
			copy(history[:], buf[2*n:2*n+8])
			break
		}
		out = (*int16)(unsafe.Add(unsafe.Pointer(out), int(written)*2))
		copy(buf[:8], buf[2*n:2*n+8])
	}
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
func Opus_silk_resampler_private_up2_HQ(tls *libc.TLS, S *[6]OpusT_opus_int32, out *OpusT_opus_int16, in *OpusT_opus_int16, len1 OpusT_opus_int32) {
	if len1 <= 0 {
		return
	}
	input := unsafe.Slice(in, int(len1))
	output := unsafe.Slice(out, 2*int(len1))
	for k, sample := range input {
		// Internal variables and state are Q10, with 32-bit wraparound.
		in32 := int32(uint32(int32(sample)) << 10)
		Y := in32 - S[0]
		X := int32((int64(Y) * int64(silk_resampler_up2_hq_06[0])) >> 16)
		out1 := S[0] + X
		S[0] = in32 + X
		Y = out1 - S[1]
		X = int32((int64(Y) * int64(silk_resampler_up2_hq_06[1])) >> 16)
		out2 := S[1] + X
		S[1] = out1 + X
		Y = out2 - S[2]
		X = int32(int64(Y) + ((int64(Y) * int64(silk_resampler_up2_hq_06[2])) >> 16))
		out1 = S[2] + X
		S[2] = out2 + X
		value := ((out1 >> 9) + 1) >> 1
		output[2*k] = int16(min(max(value, -32768), 32767))

		Y = in32 - S[3]
		X = int32((int64(Y) * int64(silk_resampler_up2_hq_16[0])) >> 16)
		out1 = S[3] + X
		S[3] = in32 + X
		Y = out1 - S[4]
		X = int32((int64(Y) * int64(silk_resampler_up2_hq_16[1])) >> 16)
		out2 = S[4] + X
		S[4] = out1 + X
		Y = out2 - S[5]
		X = int32(int64(Y) + ((int64(Y) * int64(silk_resampler_up2_hq_16[2])) >> 16))
		out1 = S[5] + X
		S[5] = out2 + X
		value = ((out1 >> 9) + 1) >> 1
		output[2*k+1] = int16(min(max(value, -32768), 32767))
	}
}

func Opus_silk_resampler_private_up2_HQ_wrapper(tls *libc.TLS, state *OpusT_silk_resampler_state_struct, out, in *OpusT_opus_int16, len1 OpusT_opus_int32) {
	Opus_silk_resampler_private_up2_HQ(tls, &state.FsIIR, out, in, len1)
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

func Opus_silk_insertion_sort_increasing(tls *libc.TLS, a *OpusT_opus_int32, idx *int32, L int32, K int32) {
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
	values := unsafe.Slice(a, int(L))
	indices := unsafe.Slice(idx, int(K))
	for i := range indices {
		indices[i] = int32(i)
	}
	for i := 1; i < int(K); i++ {
		value := values[i]
		j := i - 1
		for ; j >= 0 && value < values[j]; j-- {
			values[j+1], indices[j+1] = values[j], indices[j]
		}
		values[j+1], indices[j+1] = value, int32(i)
	}
	// Only the first K entries are sorted; the tail is read but not changed.
	for i := int(K); i < int(L); i++ {
		value := values[i]
		if value < values[K-1] {
			j := int(K) - 2
			for ; j >= 0 && value < values[j]; j-- {
				values[j+1], indices[j+1] = values[j], indices[j]
			}
			values[j+1], indices[j+1] = value, int32(i)
		}
	}
}

func Opus_silk_insertion_sort_increasing_all_values_int16(tls *libc.TLS, a *OpusT_opus_int16, L int32) {
	if !(L > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+7750, __ccgo_ts+7735, int32(144))
	}
	values := unsafe.Slice(a, int(L))
	for i := 1; i < len(values); i++ {
		value := values[i]
		j := i - 1
		for ; j >= 0 && value < values[j]; j-- {
			values[j+1] = values[j]
		}
		values[j+1] = value
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
func Opus_silk_sum_sqr_shift(tls *libc.TLS, energy *OpusT_opus_int32, shift *int32, x *OpusT_opus_int16, len1 int32) {
	input := unsafe.Slice(x, int(len1))
	shft := int32(31 - bits.LeadingZeros32(uint32(len1)))
	// Start conservatively with nrg=len, then recompute with two headroom bits.
	nrg := len1
	for pass := 0; pass < 2; pass++ {
		if pass == 1 {
			shft = max(0, shft+3-int32(bits.LeadingZeros32(uint32(nrg))))
			nrg = 0
		}
		i := 0
		for ; i+1 < len(input); i += 2 {
			a, b := int32(input[i]), int32(input[i+1])
			// The sum of two squares can set bit 31: shift it unsigned.
			pair := uint32(a*a) + uint32(b*b)
			nrg = int32(uint32(nrg) + (pair >> shft))
		}
		if i < len(input) {
			a := int32(input[i])
			nrg = int32(uint32(nrg) + (uint32(a*a) >> shft))
		}
	}
	*shift = shft
	*energy = nrg
}

// C documentation
//
//	/* Decode mid/side predictors */
