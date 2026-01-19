// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

package opuscc

import (
	"reflect"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

var _ reflect.Type
var _ unsafe.Pointer

func Opus_silk_VAD_Init(tls *libc.TLS, psSilk_VAD uintptr) (r int32) {
	var b1, ret, v6 int32
	var v2, v3, v4 OpusT_opus_int32
	_, _, _, _, _, _ = b1, ret, v2, v3, v4, v6
	ret = 0
	/* reset state memory */
	libc.Xmemset(tls, psSilk_VAD, 0, uint64(112))
	/* init noise levels */
	/* Initialize array with approx pink noise levels (psd proportional to inverse of frequency) */
	b1 = 0
	for {
		if !(b1 < int32(VAD_N_BANDS)) {
			break
		}
		v2 = int32(VAD_NOISE_LEVELS_BIAS) / (b1 + int32(1))
		v3 = int32(1)
		if v2 > v3 {
			v6 = v2
		} else {
			v6 = v3
		}
		v4 = v6
		*(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 92 + uintptr(b1)*4)) = v4
		b1 = b1 + 1
	}
	/* Initialize state */
	b1 = 0
	for {
		if !(b1 < int32(VAD_N_BANDS)) {
			break
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 60 + uintptr(b1)*4)) = int32(100) * *(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 92 + uintptr(b1)*4))
		*(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 76 + uintptr(b1)*4)) = int32(silk_int32_MAX) / *(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 60 + uintptr(b1)*4))
		b1 = b1 + 1
	}
	(*OpusT_silk_VAD_state)(unsafe.Pointer(psSilk_VAD)).Fcounter = int32(15)
	/* init smoothed energy-to-noise ratio*/
	b1 = 0
	for {
		if !(b1 < int32(VAD_N_BANDS)) {
			break
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 40 + uintptr(b1)*4)) = int32(100) * int32(256) /* 100 * 256 --> 20 dB SNR */
		b1 = b1 + 1
	}
	return ret
}

// C documentation
//
//	/* Weighting factors for tilt measure */
var tiltWeights = [4]OpusT_opus_int32{
	0: int32(30000),
	1: int32(6000),
	2: -int32(12000),
	3: -int32(12000),
}

// C documentation
//
//	/***************************************/
//	/* Get the speech activity level in Q8 */
//	/***************************************/
func Opus_silk_VAD_GetSA_Q8_c(tls *libc.TLS, psEncC uintptr, pIn uintptr) (r1 int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var HPstateTmp OpusT_opus_int16
	var NrgToNoiseRatio_Q8 [4]OpusT_opus_int32
	var SA_Q15, SNR_Q7, b1, dec_subframe_length, dec_subframe_offset, decimated_framelength, decimated_framelength1, decimated_framelength2, i, input_tilt, pSNR_dB_Q7, ret, s, v33, v34, v35, v37 int32
	var X, _saved_stack, psSilk_VAD, st, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var X_offset [4]int32
	var lzeros, smooth_coef_Q16, speech_nrg, sumSquared, x_tmp, y, v43, v44, v46, v47, v48, v51, v53 OpusT_opus_int32
	var m, r, x OpusT_opus_uint32
	var _ /* Xnrg at bp+8 */ [4]OpusT_opus_int32
	var _ /* frac_Q7 at bp+4 */ OpusT_opus_int32
	var _ /* lz at bp+0 */ OpusT_opus_int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = HPstateTmp, NrgToNoiseRatio_Q8, SA_Q15, SNR_Q7, X, X_offset, _saved_stack, b1, dec_subframe_length, dec_subframe_offset, decimated_framelength, decimated_framelength1, decimated_framelength2, i, input_tilt, lzeros, m, pSNR_dB_Q7, psSilk_VAD, r, ret, s, smooth_coef_Q16, speech_nrg, st, sumSquared, x, x_tmp, y, v1, v11, v13, v15, v17, v19, v21, v23, v3, v33, v34, v35, v37, v43, v44, v46, v47, v48, v5, v51, v53, v7, v9
	ret = 0
	psSilk_VAD = psEncC + 36
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
	/* Safety checks */
	_ = true
	if !(int32(SUB_FRAME_LENGTH_MS)*int32(MAX_NB_SUBFR)*int32(MAX_FS_KHZ) >= (*OpusT_silk_encoder_state)(unsafe.Pointer(psEncC)).Fframe_length) {
		Opus_celt_fatal(tls, __ccgo_ts+6796, __ccgo_ts+6855, int32(104))
	}
	if !((*OpusT_silk_encoder_state)(unsafe.Pointer(psEncC)).Fframe_length <= int32(512)) {
		Opus_celt_fatal(tls, __ccgo_ts+6869, __ccgo_ts+6855, int32(105))
	}
	if !((*OpusT_silk_encoder_state)(unsafe.Pointer(psEncC)).Fframe_length == int32(8)*((*OpusT_silk_encoder_state)(unsafe.Pointer(psEncC)).Fframe_length>>int32(3))) {
		Opus_celt_fatal(tls, __ccgo_ts+6915, __ccgo_ts+6855, int32(106))
	}
	/***********************/
	/* Filter and Decimate */
	/***********************/
	decimated_framelength1 = (*OpusT_silk_encoder_state)(unsafe.Pointer(psEncC)).Fframe_length >> int32(1)
	decimated_framelength2 = (*OpusT_silk_encoder_state)(unsafe.Pointer(psEncC)).Fframe_length >> int32(2)
	decimated_framelength = (*OpusT_silk_encoder_state)(unsafe.Pointer(psEncC)).Fframe_length >> int32(3)
	/* Decimate into 4 bands:
	   0       L      3L       L              3L                             5L
	           -      --       -              --                             --
	           8       8       2               4                              4
	   [0-1 kHz| temp. |1-2 kHz|    2-4 kHz    |            4-8 kHz           |
	   They're arranged to allow the minimal ( frame_length / 4 ) extra
	   scratch space during the downsampling process */
	X_offset[0] = 0
	X_offset[int32(1)] = decimated_framelength + decimated_framelength2
	X_offset[int32(2)] = X_offset[int32(1)] + decimated_framelength
	X_offset[int32(3)] = X_offset[int32(2)] + decimated_framelength2
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
	if !(int64(int32(uint64(uint32(X_offset[int32(3)]+decimated_framelength1))*(uint64(2)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+6855, int32(127))
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
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32(X_offset[int32(3)]+decimated_framelength1)) * (uint64(2) / uint64(1)))
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
	X = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32(X_offset[int32(3)]+decimated_framelength1))*(uint64(2)/uint64(1)))
	/* 0-8 kHz to 0-4 kHz and 4-8 kHz */
	Opus_silk_ana_filt_bank_1(tls, pIn, psSilk_VAD, X, X+uintptr(X_offset[int32(3)])*2, (*OpusT_silk_encoder_state)(unsafe.Pointer(psEncC)).Fframe_length)
	/* 0-4 kHz to 0-2 kHz and 2-4 kHz */
	Opus_silk_ana_filt_bank_1(tls, X, psSilk_VAD+8, X, X+uintptr(X_offset[int32(2)])*2, decimated_framelength1)
	/* 0-2 kHz to 0-1 kHz and 1-2 kHz */
	Opus_silk_ana_filt_bank_1(tls, X, psSilk_VAD+16, X, X+uintptr(X_offset[int32(1)])*2, decimated_framelength2)
	/*********************************************/
	/* HP filter on lowest band (differentiator) */
	/*********************************************/
	*(*OpusT_opus_int16)(unsafe.Pointer(X + uintptr(decimated_framelength-int32(1))*2)) = int16(int32(*(*OpusT_opus_int16)(unsafe.Pointer(X + uintptr(decimated_framelength-int32(1))*2))) >> int32(1))
	HPstateTmp = *(*OpusT_opus_int16)(unsafe.Pointer(X + uintptr(decimated_framelength-int32(1))*2))
	i = decimated_framelength - int32(1)
	for {
		if !(i > 0) {
			break
		}
		*(*OpusT_opus_int16)(unsafe.Pointer(X + uintptr(i-int32(1))*2)) = int16(int32(*(*OpusT_opus_int16)(unsafe.Pointer(X + uintptr(i-int32(1))*2))) >> int32(1))
		v1 = X + uintptr(i)*2
		*(*OpusT_opus_int16)(unsafe.Pointer(v1)) = OpusT_opus_int16(int32(*(*OpusT_opus_int16)(unsafe.Pointer(v1))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer(X + uintptr(i-int32(1))*2))))
		i = i - 1
	}
	v1 = X
	*(*OpusT_opus_int16)(unsafe.Pointer(v1)) = OpusT_opus_int16(int32(*(*OpusT_opus_int16)(unsafe.Pointer(v1))) - int32((*OpusT_silk_VAD_state)(unsafe.Pointer(psSilk_VAD)).FHPstate))
	(*OpusT_silk_VAD_state)(unsafe.Pointer(psSilk_VAD)).FHPstate = HPstateTmp
	/*************************************/
	/* Calculate the energy in each band */
	/*************************************/
	b1 = 0
	for {
		if !(b1 < int32(VAD_N_BANDS)) {
			break
		}
		/* Find the decimated framelength in the non-uniformly divided bands */
		v33 = int32(VAD_N_BANDS) - b1
		v34 = int32(VAD_N_BANDS) - int32(1)
		if v33 < v34 {
			v37 = v33
		} else {
			v37 = v34
		}
		v35 = v37
		decimated_framelength = (*OpusT_silk_encoder_state)(unsafe.Pointer(psEncC)).Fframe_length >> v35
		/* Split length into subframe lengths */
		dec_subframe_length = decimated_framelength >> int32(VAD_INTERNAL_SUBFRAMES_LOG2)
		dec_subframe_offset = 0
		/* Compute energy per sub-frame */
		/* initialize with summed energy of last subframe */
		(*(*[4]OpusT_opus_int32)(unsafe.Pointer(bp + 8)))[b1] = *(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 24 + uintptr(b1)*4))
		s = 0
		for {
			if !(s < int32(1)<<int32(VAD_INTERNAL_SUBFRAMES_LOG2)) {
				break
			}
			sumSquared = 0
			i = 0
			for {
				if !(i < dec_subframe_length) {
					break
				}
				/* The energy will be less than dec_subframe_length * ( silk_int16_MIN / 8 ) ^ 2.            */
				/* Therefore we can accumulate with no risk of overflow (unless dec_subframe_length > 128)  */
				x_tmp = int32(*(*OpusT_opus_int16)(unsafe.Pointer(X + uintptr(X_offset[b1]+i+dec_subframe_offset)*2))) >> int32(3)
				sumSquared = sumSquared + int32(int16(x_tmp))*int32(int16(x_tmp))
				/* Safety check */
				_ = sumSquared >= int32(0)
				i = i + 1
			}
			/* Add/saturate summed energy of current subframe */
			if s < int32(1)<<int32(VAD_INTERNAL_SUBFRAMES_LOG2)-int32(1) {
				if (uint32((*(*[4]OpusT_opus_int32)(unsafe.Pointer(bp + 8)))[b1])+uint32(sumSquared))&uint32(0x80000000) != 0 {
					v33 = int32(silk_int32_MAX)
				} else {
					v33 = (*(*[4]OpusT_opus_int32)(unsafe.Pointer(bp + 8)))[b1] + sumSquared
				}
				(*(*[4]OpusT_opus_int32)(unsafe.Pointer(bp + 8)))[b1] = v33
			} else {
				/* Look-ahead subframe */
				if (uint32((*(*[4]OpusT_opus_int32)(unsafe.Pointer(bp + 8)))[b1])+uint32(sumSquared>>int32(1)))&uint32(0x80000000) != 0 {
					v33 = int32(silk_int32_MAX)
				} else {
					v33 = (*(*[4]OpusT_opus_int32)(unsafe.Pointer(bp + 8)))[b1] + sumSquared>>int32(1)
				}
				(*(*[4]OpusT_opus_int32)(unsafe.Pointer(bp + 8)))[b1] = v33
			}
			dec_subframe_offset = dec_subframe_offset + dec_subframe_length
			s = s + 1
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 24 + uintptr(b1)*4)) = sumSquared
		b1 = b1 + 1
	}
	/********************/
	/* Noise estimation */
	/********************/
	silk_VAD_GetNoiseLevels(tls, bp+8, psSilk_VAD)
	/***********************************************/
	/* Signal-plus-noise to noise ratio estimation */
	/***********************************************/
	sumSquared = 0
	input_tilt = 0
	b1 = 0
	for {
		if !(b1 < int32(VAD_N_BANDS)) {
			break
		}
		speech_nrg = (*(*[4]OpusT_opus_int32)(unsafe.Pointer(bp + 8)))[b1] - *(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 60 + uintptr(b1)*4))
		if speech_nrg > 0 {
			/* Divide, with sufficient resolution */
			if uint32((*(*[4]OpusT_opus_int32)(unsafe.Pointer(bp + 8)))[b1])&uint32(0xFF800000) == uint32(0) {
				NrgToNoiseRatio_Q8[b1] = int32(uint32((*(*[4]OpusT_opus_int32)(unsafe.Pointer(bp + 8)))[b1])<<int32(8)) / (*(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 60 + uintptr(b1)*4)) + int32(1))
			} else {
				NrgToNoiseRatio_Q8[b1] = (*(*[4]OpusT_opus_int32)(unsafe.Pointer(bp + 8)))[b1] / (*(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 60 + uintptr(b1)*4))>>int32(8) + int32(1))
			}
			/* Convert to log domain */
			SNR_Q7 = Opus_silk_lin2log(tls, NrgToNoiseRatio_Q8[b1]) - int32(8)*int32(128)
			/* Sum-of-squares */
			sumSquared = sumSquared + int32(int16(SNR_Q7))*int32(int16(SNR_Q7)) /* Q14 */
			/* Tilt measure */
			if speech_nrg < int32(1)<<int32(20) {
				/* Scale down SNR value for small subband speech energies */
				v43 = speech_nrg
				if v43 <= int32(0) {
					v44 = 0
					goto _45
				}
				v46 = v43
				v47 = v46
				if v47 != 0 {
					v33 = int32(32) - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, uint32(v47)))
				} else {
					v33 = int32(32)
				}
				v48 = v33
				lzeros = v48
				*(*OpusT_opus_int32)(unsafe.Pointer(bp)) = lzeros
				v51 = v46
				v34 = int32(24) - lzeros
				x = uint32(v51)
				r = uint32(v34)
				m = uint32(-v34)
				if v34 == int32(0) {
					v53 = v51
					goto _54
				} else {
					if v34 < int32(0) {
						v53 = int32(x<<m | x>>(uint32(32)-m))
						goto _54
					} else {
						v53 = int32(x<<(uint32(32)-r) | x>>r)
						goto _54
					}
				}
			_54:
				*(*OpusT_opus_int32)(unsafe.Pointer(bp + 4)) = v53 & int32(0x7f)
				if *(*OpusT_opus_int32)(unsafe.Pointer(bp))&int32(1) != 0 {
					y = int32(32768)
				} else {
					y = int32(46214)
				}
				y = y >> (*(*OpusT_opus_int32)(unsafe.Pointer(bp)) >> int32(1))
				y = int32(int64(y) + int64(y)*int64(int16(int32(int16(int32(213)))*int32(int16(*(*OpusT_opus_int32)(unsafe.Pointer(bp + 4))))))>>int32(16))
				v44 = y
			_45:
				SNR_Q7 = int32(int64(int32(uint32(v44)<<int32(6))) * int64(int16(SNR_Q7)) >> int32(16))
			}
			input_tilt = int32(int64(input_tilt) + int64(tiltWeights[b1])*int64(int16(SNR_Q7))>>int32(16))
		} else {
			NrgToNoiseRatio_Q8[b1] = int32(256)
		}
		b1 = b1 + 1
	}
	/* Mean-of-squares */
	sumSquared = sumSquared / int32(VAD_N_BANDS) /* Q14 */
	/* Root-mean-square approximation, scale to dBs, and write to output pointer */
	v43 = sumSquared
	if v43 <= int32(0) {
		v44 = 0
		goto _57
	}
	v46 = v43
	v47 = v46
	if v47 != 0 {
		v33 = int32(32) - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, uint32(v47)))
	} else {
		v33 = int32(32)
	}
	v48 = v33
	lzeros = v48
	*(*OpusT_opus_int32)(unsafe.Pointer(bp)) = lzeros
	v51 = v46
	v34 = int32(24) - lzeros
	x = uint32(v51)
	r = uint32(v34)
	m = uint32(-v34)
	if v34 == int32(0) {
		v53 = v51
		goto _66
	} else {
		if v34 < int32(0) {
			v53 = int32(x<<m | x>>(uint32(32)-m))
			goto _66
		} else {
			v53 = int32(x<<(uint32(32)-r) | x>>r)
			goto _66
		}
	}
_66:
	*(*OpusT_opus_int32)(unsafe.Pointer(bp + 4)) = v53 & int32(0x7f)
	if *(*OpusT_opus_int32)(unsafe.Pointer(bp))&int32(1) != 0 {
		y = int32(32768)
	} else {
		y = int32(46214)
	}
	y = y >> (*(*OpusT_opus_int32)(unsafe.Pointer(bp)) >> int32(1))
	y = int32(int64(y) + int64(y)*int64(int16(int32(int16(int32(213)))*int32(int16(*(*OpusT_opus_int32)(unsafe.Pointer(bp + 4))))))>>int32(16))
	v44 = y
_57:
	pSNR_dB_Q7 = int32(int16(int32(3) * v44)) /* Q7 */
	/*********************************/
	/* Speech Probability Estimation */
	/*********************************/
	SA_Q15 = Opus_silk_sigm_Q15(tls, int32(int64(int32(VAD_SNR_FACTOR_Q16))*int64(int16(pSNR_dB_Q7))>>int32(16))-int32(VAD_NEGATIVE_OFFSET_Q5))
	/**************************/
	/* Frequency Tilt Measure */
	/**************************/
	(*OpusT_silk_encoder_state)(unsafe.Pointer(psEncC)).Finput_tilt_Q15 = int32(uint32(Opus_silk_sigm_Q15(tls, input_tilt)-int32(16384)) << int32(1))
	/**************************************************/
	/* Scale the sigmoid output based on power levels */
	/**************************************************/
	speech_nrg = 0
	b1 = 0
	for {
		if !(b1 < int32(VAD_N_BANDS)) {
			break
		}
		/* Accumulate signal-without-noise energies, higher frequency bands have more weight */
		speech_nrg = speech_nrg + (b1+int32(1))*(((*(*[4]OpusT_opus_int32)(unsafe.Pointer(bp + 8)))[b1]-*(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 60 + uintptr(b1)*4)))>>int32(4))
		b1 = b1 + 1
	}
	if (*OpusT_silk_encoder_state)(unsafe.Pointer(psEncC)).Fframe_length == int32(20)*(*OpusT_silk_encoder_state)(unsafe.Pointer(psEncC)).Ffs_kHz {
		speech_nrg = speech_nrg >> int32(1)
	}
	/* Power scaling */
	if speech_nrg <= 0 {
		SA_Q15 = SA_Q15 >> int32(1)
	} else {
		if speech_nrg < int32(16384) {
			speech_nrg = int32(uint32(speech_nrg) << int32(16))
			/* square-root */
			v43 = speech_nrg
			if v43 <= int32(0) {
				v44 = 0
				goto _70
			}
			v46 = v43
			v47 = v46
			if v47 != 0 {
				v33 = int32(32) - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, uint32(v47)))
			} else {
				v33 = int32(32)
			}
			v48 = v33
			lzeros = v48
			*(*OpusT_opus_int32)(unsafe.Pointer(bp)) = lzeros
			v51 = v46
			v34 = int32(24) - lzeros
			x = uint32(v51)
			r = uint32(v34)
			m = uint32(-v34)
			if v34 == int32(0) {
				v53 = v51
				goto _79
			} else {
				if v34 < int32(0) {
					v53 = int32(x<<m | x>>(uint32(32)-m))
					goto _79
				} else {
					v53 = int32(x<<(uint32(32)-r) | x>>r)
					goto _79
				}
			}
		_79:
			*(*OpusT_opus_int32)(unsafe.Pointer(bp + 4)) = v53 & int32(0x7f)
			if *(*OpusT_opus_int32)(unsafe.Pointer(bp))&int32(1) != 0 {
				y = int32(32768)
			} else {
				y = int32(46214)
			}
			y = y >> (*(*OpusT_opus_int32)(unsafe.Pointer(bp)) >> int32(1))
			y = int32(int64(y) + int64(y)*int64(int16(int32(int16(int32(213)))*int32(int16(*(*OpusT_opus_int32)(unsafe.Pointer(bp + 4))))))>>int32(16))
			v44 = y
		_70:
			speech_nrg = v44
			SA_Q15 = int32(int64(int32(32768)+speech_nrg) * int64(int16(SA_Q15)) >> int32(16))
		}
	}
	/* Copy the resulting speech activity in Q8 */
	v33 = SA_Q15 >> int32(7)
	v34 = int32(silk_uint8_MAX1)
	if v33 < v34 {
		v37 = v33
	} else {
		v37 = v34
	}
	v35 = v37
	(*OpusT_silk_encoder_state)(unsafe.Pointer(psEncC)).Fspeech_activity_Q8 = v35
	/***********************************/
	/* Energy Level and SNR estimation */
	/***********************************/
	/* Smoothing coefficient */
	smooth_coef_Q16 = int32(int64(int32(VAD_SNR_SMOOTH_COEF_Q18)) * int64(int16(int32(int64(SA_Q15)*int64(int16(SA_Q15))>>int32(16)))) >> int32(16))
	if (*OpusT_silk_encoder_state)(unsafe.Pointer(psEncC)).Fframe_length == int32(10)*(*OpusT_silk_encoder_state)(unsafe.Pointer(psEncC)).Ffs_kHz {
		smooth_coef_Q16 = smooth_coef_Q16 >> int32(1)
	}
	b1 = 0
	for {
		if !(b1 < int32(VAD_N_BANDS)) {
			break
		}
		/* compute smoothed energy-to-noise ratio per band */
		*(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 40 + uintptr(b1)*4)) = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 40 + uintptr(b1)*4))) + int64(NrgToNoiseRatio_Q8[b1]-*(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 40 + uintptr(b1)*4)))*int64(int16(smooth_coef_Q16))>>int32(16))
		/* signal to noise ratio in dB per band */
		SNR_Q7 = int32(3) * (Opus_silk_lin2log(tls, *(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 40 + uintptr(b1)*4))) - int32(8)*int32(128))
		/* quality = sigmoid( 0.25 * ( SNR_dB - 16 ) ); */
		*(*int32)(unsafe.Pointer(psEncC + 4712 + uintptr(b1)*4)) = Opus_silk_sigm_Q15(tls, (SNR_Q7-int32(16)*int32(128))>>int32(4))
		b1 = b1 + 1
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
	return ret
}

// C documentation
//
//	/**************************/
//	/* Noise level estimation */
//	/**************************/
func silk_VAD_GetNoiseLevels(tls *libc.TLS, pX uintptr, psSilk_VAD uintptr) {
	var coef, k, min_coef, v2, v3, v4, v5 int32
	var inv_nrg, nl, nrg OpusT_opus_int32
	_, _, _, _, _, _, _, _, _, _ = coef, inv_nrg, k, min_coef, nl, nrg, v2, v3, v4, v5
	/* Initially faster smoothing */
	if (*OpusT_silk_VAD_state)(unsafe.Pointer(psSilk_VAD)).Fcounter < int32(1000) { /* 1000 = 20 sec */
		min_coef = int32(silk_int16_MAX9) / ((*OpusT_silk_VAD_state)(unsafe.Pointer(psSilk_VAD)).Fcounter>>int32(4) + int32(1))
		/* Increment frame counter */
		(*OpusT_silk_VAD_state)(unsafe.Pointer(psSilk_VAD)).Fcounter = (*OpusT_silk_VAD_state)(unsafe.Pointer(psSilk_VAD)).Fcounter + 1
	} else {
		min_coef = 0
	}
	k = 0
	for {
		if !(k < int32(VAD_N_BANDS)) {
			break
		}
		/* Get old noise level estimate for current band */
		nl = *(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 60 + uintptr(k)*4))
		_ = nl >= int32(0)
		/* Add bias */
		if (uint32(*(*OpusT_opus_int32)(unsafe.Pointer(pX + uintptr(k)*4)))+uint32(*(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 92 + uintptr(k)*4))))&uint32(0x80000000) != 0 {
			v2 = int32(silk_int32_MAX)
		} else {
			v2 = *(*OpusT_opus_int32)(unsafe.Pointer(pX + uintptr(k)*4)) + *(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 92 + uintptr(k)*4))
		}
		nrg = v2
		_ = nrg > int32(0)
		/* Invert energies */
		inv_nrg = int32(silk_int32_MAX) / nrg
		_ = inv_nrg >= int32(0)
		/* Less update when subband energy is high */
		if nrg > int32(uint32(nl)<<int32(3)) {
			coef = int32(VAD_NOISE_LEVEL_SMOOTH_COEF_Q16) >> int32(3)
		} else {
			if nrg < nl {
				coef = int32(VAD_NOISE_LEVEL_SMOOTH_COEF_Q16)
			} else {
				coef = int32(int64(int32(int64(inv_nrg)*int64(nl)>>int32(16))) * int64(int16(int32(VAD_NOISE_LEVEL_SMOOTH_COEF_Q16)<<int32(1))) >> int32(16))
			}
		}
		/* Initially faster smoothing */
		v2 = coef
		v3 = min_coef
		if v2 > v3 {
			v5 = v2
		} else {
			v5 = v3
		}
		v4 = v5
		coef = v4
		/* Smooth inverse energies */
		*(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 76 + uintptr(k)*4)) = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 76 + uintptr(k)*4))) + int64(inv_nrg-*(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 76 + uintptr(k)*4)))*int64(int16(coef))>>int32(16))
		_ = *(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 76 + uintptr(k)*4)) >= int32(0)
		/* Compute noise level by inverting again */
		nl = int32(silk_int32_MAX) / *(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 76 + uintptr(k)*4))
		_ = nl >= int32(0)
		/* Limit noise levels (guarantee 7 bits of head room) */
		if nl < int32(0x00FFFFFF) {
			v2 = nl
		} else {
			v2 = int32(0x00FFFFFF)
		}
		nl = v2
		/* Store as part of state */
		*(*OpusT_opus_int32)(unsafe.Pointer(psSilk_VAD + 60 + uintptr(k)*4)) = nl
		k = k + 1
	}
}

const silk_int16_MAX10 = 0x7FFF
const silk_uint8_MAX2 = 0xFF

// C documentation
//
//	/* Compute quantization errors for an LPC_order element input vector for a VQ codebook */
func Opus_silk_NLSF_VQ(tls *libc.TLS, err_Q24 uintptr, in_Q15 uintptr, pCB_Q8 uintptr, pWght_Q9 uintptr, K int32, LPC_order int32) {
	var cb_Q8_ptr, w_Q9_ptr uintptr
	var diff_Q15, diffw_Q24, pred_Q24, sum_error_Q24 OpusT_opus_int32
	var i, m, v3 int32
	_, _, _, _, _, _, _, _, _ = cb_Q8_ptr, diff_Q15, diffw_Q24, i, m, pred_Q24, sum_error_Q24, w_Q9_ptr, v3
	if !(LPC_order&int32(1) == int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+7000, __ccgo_ts+7041, int32(49))
	}
	/* Loop over codebook */
	cb_Q8_ptr = pCB_Q8
	w_Q9_ptr = pWght_Q9
	i = 0
	for {
		if !(i < K) {
			break
		}
		sum_error_Q24 = 0
		pred_Q24 = 0
		m = LPC_order - int32(2)
		for {
			if !(m >= 0) {
				break
			}
			/* Compute weighted absolute predictive quantization error for index m + 1 */
			diff_Q15 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(in_Q15 + uintptr(m+int32(1))*2))) - int32(uint32(int32(*(*OpusT_opus_uint8)(unsafe.Pointer(cb_Q8_ptr + uintptr(m+int32(1))))))<<int32(7)) /* range: [ -32767 : 32767 ]*/
			diffw_Q24 = int32(int16(diff_Q15)) * int32(*(*OpusT_opus_int16)(unsafe.Pointer(w_Q9_ptr + uintptr(m+int32(1))*2)))
			if diffw_Q24-pred_Q24>>int32(1) > 0 {
				v3 = diffw_Q24 - pred_Q24>>int32(1)
			} else {
				v3 = -(diffw_Q24 - pred_Q24>>int32(1))
			}
			sum_error_Q24 = sum_error_Q24 + v3
			pred_Q24 = diffw_Q24
			/* Compute weighted absolute predictive quantization error for index m */
			diff_Q15 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(in_Q15 + uintptr(m)*2))) - int32(uint32(int32(*(*OpusT_opus_uint8)(unsafe.Pointer(cb_Q8_ptr + uintptr(m)))))<<int32(7)) /* range: [ -32767 : 32767 ]*/
			diffw_Q24 = int32(int16(diff_Q15)) * int32(*(*OpusT_opus_int16)(unsafe.Pointer(w_Q9_ptr + uintptr(m)*2)))
			if diffw_Q24-pred_Q24>>int32(1) > 0 {
				v3 = diffw_Q24 - pred_Q24>>int32(1)
			} else {
				v3 = -(diffw_Q24 - pred_Q24>>int32(1))
			}
			sum_error_Q24 = sum_error_Q24 + v3
			pred_Q24 = diffw_Q24
			_ = sum_error_Q24 >= int32(0)
			m = m - int32(2)
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(err_Q24 + uintptr(i)*4)) = sum_error_Q24
		cb_Q8_ptr = cb_Q8_ptr + uintptr(LPC_order)
		w_Q9_ptr = w_Q9_ptr + uintptr(LPC_order)*2
		i = i + 1
	}
}

// C documentation
//
//	/* Unpack predictor values and indices for entropy coding tables */
func Opus_silk_NLSF_unpack(tls *libc.TLS, ec_ix uintptr, pred_Q8 uintptr, psNLSF_CB uintptr, CB1_index int32) {
	var ec_sel_ptr, v2 uintptr
	var entry OpusT_opus_uint8
	var i int32
	_, _, _, _ = ec_sel_ptr, entry, i, v2
	ec_sel_ptr = (*OpusT_silk_NLSF_CB_struct)(unsafe.Pointer(psNLSF_CB)).Fec_sel + uintptr(CB1_index*int32((*OpusT_silk_NLSF_CB_struct)(unsafe.Pointer(psNLSF_CB)).Forder)/int32(2))
	i = 0
	for {
		if !(i < int32((*OpusT_silk_NLSF_CB_struct)(unsafe.Pointer(psNLSF_CB)).Forder)) {
			break
		}
		v2 = ec_sel_ptr
		ec_sel_ptr = ec_sel_ptr + 1
		entry = *(*OpusT_opus_uint8)(unsafe.Pointer(v2))
		*(*OpusT_opus_int16)(unsafe.Pointer(ec_ix + uintptr(i)*2)) = int16(int32(int16(int32(entry)>>int32(1)&int32(7))) * int32(int16(int32(2)*int32(NLSF_QUANT_MAX_AMPLITUDE)+int32(1))))
		*(*OpusT_opus_uint8)(unsafe.Pointer(pred_Q8 + uintptr(i))) = *(*OpusT_opus_uint8)(unsafe.Pointer((*OpusT_silk_NLSF_CB_struct)(unsafe.Pointer(psNLSF_CB)).Fpred_Q8 + uintptr(i+int32(entry)&int32(1)*(int32((*OpusT_silk_NLSF_CB_struct)(unsafe.Pointer(psNLSF_CB)).Forder)-int32(1)))))
		*(*OpusT_opus_int16)(unsafe.Pointer(ec_ix + uintptr(i+int32(1))*2)) = int16(int32(int16(int32(entry)>>int32(5)&int32(7))) * int32(int16(int32(2)*int32(NLSF_QUANT_MAX_AMPLITUDE)+int32(1))))
		*(*OpusT_opus_uint8)(unsafe.Pointer(pred_Q8 + uintptr(i+int32(1)))) = *(*OpusT_opus_uint8)(unsafe.Pointer((*OpusT_silk_NLSF_CB_struct)(unsafe.Pointer(psNLSF_CB)).Fpred_Q8 + uintptr(i+int32(entry)>>int32(4)&int32(1)*(int32((*OpusT_silk_NLSF_CB_struct)(unsafe.Pointer(psNLSF_CB)).Forder)-int32(1))+int32(1))))
		i = i + int32(2)
	}
}

// C documentation
//
//	/* Delayed-decision quantizer for NLSF residuals */
func Opus_silk_NLSF_del_dec_quant(tls *libc.TLS, indices uintptr, x_Q10 uintptr, w_Q5 uintptr, pred_coef_Q8 uintptr, ec_ix uintptr, ec_rates_Q5 uintptr, quant_step_size_Q16 int32, inv_quant_step_size_Q6 OpusT_opus_int16, mu_Q20 OpusT_opus_int32, order OpusT_opus_int16) (r OpusT_opus_int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var RD_Q25 [8]OpusT_opus_int32
	var RD_max_Q25, RD_min_Q25 [4]OpusT_opus_int32
	var RD_tmp_Q25, max_min_Q25, min_Q25, min_max_Q25 OpusT_opus_int32
	var diff_Q10, i, in_Q10, ind_max_min, ind_min_max, ind_tmp, j, nStates, pred_Q10, rate0_Q5, rate1_Q5, res_Q10, v4, v5 int32
	var ind_sort [4]int32
	var out0_Q10, out1_Q10 OpusT_opus_int16
	var out0_Q10_table, out1_Q10_table [20]int32
	var prev_out_Q10 [8]OpusT_opus_int16
	var rates_Q5, v11 uintptr
	var _ /* ind at bp+0 */ [4][16]OpusT_opus_int8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = RD_Q25, RD_max_Q25, RD_min_Q25, RD_tmp_Q25, diff_Q10, i, in_Q10, ind_max_min, ind_min_max, ind_sort, ind_tmp, j, max_min_Q25, min_Q25, min_max_Q25, nStates, out0_Q10, out0_Q10_table, out1_Q10, out1_Q10_table, pred_Q10, prev_out_Q10, rate0_Q5, rate1_Q5, rates_Q5, res_Q10, v11, v4, v5
	i = -int32(NLSF_QUANT_MAX_AMPLITUDE_EXT)
	for {
		if !(i <= int32(NLSF_QUANT_MAX_AMPLITUDE_EXT)-int32(1)) {
			break
		}
		out0_Q10 = int16(int32(uint32(i) << int32(10)))
		out1_Q10 = int16(int32(out0_Q10) + int32(1024))
		if i > 0 {
			out0_Q10 = int16(int32(out0_Q10) - int32(102))
			out1_Q10 = int16(int32(out1_Q10) - int32(102))
		} else {
			if i == 0 {
				out1_Q10 = int16(int32(out1_Q10) - int32(102))
			} else {
				if i == -int32(1) {
					out0_Q10 = int16(int32(out0_Q10) + int32(102))
				} else {
					out0_Q10 = int16(int32(out0_Q10) + int32(102))
					out1_Q10 = int16(int32(out1_Q10) + int32(102))
				}
			}
		}
		out0_Q10_table[i+int32(NLSF_QUANT_MAX_AMPLITUDE_EXT)] = int32(out0_Q10) * int32(int16(quant_step_size_Q16)) >> int32(16)
		out1_Q10_table[i+int32(NLSF_QUANT_MAX_AMPLITUDE_EXT)] = int32(out1_Q10) * int32(int16(quant_step_size_Q16)) >> int32(16)
		i = i + 1
	}
	_ = int32(1)<<int32(NLSF_QUANT_DEL_DEC_STATES_LOG2)&(int32(1)<<int32(NLSF_QUANT_DEL_DEC_STATES_LOG2)-int32(1)) == int32(0) /* must be power of two */
	nStates = int32(1)
	RD_Q25[0] = 0
	prev_out_Q10[0] = 0
	i = int32(order) - int32(1)
	for {
		if !(i >= 0) {
			break
		}
		rates_Q5 = ec_rates_Q5 + uintptr(*(*OpusT_opus_int16)(unsafe.Pointer(ec_ix + uintptr(i)*2)))
		in_Q10 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(x_Q10 + uintptr(i)*2)))
		j = 0
		for {
			if !(j < nStates) {
				break
			}
			pred_Q10 = int32(int16(*(*OpusT_opus_uint8)(unsafe.Pointer(pred_coef_Q8 + uintptr(i))))) * int32(prev_out_Q10[j]) >> int32(8)
			res_Q10 = in_Q10 - pred_Q10
			ind_tmp = int32(inv_quant_step_size_Q6) * int32(int16(res_Q10)) >> int32(16)
			if ind_tmp > int32(NLSF_QUANT_MAX_AMPLITUDE_EXT)-int32(1) {
				v4 = int32(NLSF_QUANT_MAX_AMPLITUDE_EXT) - int32(1)
			} else {
				if ind_tmp < -int32(NLSF_QUANT_MAX_AMPLITUDE_EXT) {
					v5 = -int32(NLSF_QUANT_MAX_AMPLITUDE_EXT)
				} else {
					v5 = ind_tmp
				}
				v4 = v5
			}
			ind_tmp = v4
			*(*OpusT_opus_int8)(unsafe.Pointer(bp + uintptr(j)*16 + uintptr(i))) = int8(ind_tmp)
			/* compute outputs for ind_tmp and ind_tmp + 1 */
			out0_Q10 = int16(out0_Q10_table[ind_tmp+int32(NLSF_QUANT_MAX_AMPLITUDE_EXT)])
			out1_Q10 = int16(out1_Q10_table[ind_tmp+int32(NLSF_QUANT_MAX_AMPLITUDE_EXT)])
			out0_Q10 = int16(int32(out0_Q10) + pred_Q10)
			out1_Q10 = int16(int32(out1_Q10) + pred_Q10)
			prev_out_Q10[j] = out0_Q10
			prev_out_Q10[j+nStates] = out1_Q10
			/* compute RD for ind_tmp and ind_tmp + 1 */
			if ind_tmp+int32(1) >= int32(NLSF_QUANT_MAX_AMPLITUDE) {
				if ind_tmp+int32(1) == int32(NLSF_QUANT_MAX_AMPLITUDE) {
					rate0_Q5 = int32(*(*OpusT_opus_uint8)(unsafe.Pointer(rates_Q5 + uintptr(ind_tmp+int32(NLSF_QUANT_MAX_AMPLITUDE)))))
					rate1_Q5 = int32(280)
				} else {
					rate0_Q5 = int32(280) - int32(43)*int32(NLSF_QUANT_MAX_AMPLITUDE) + int32(int16(int32(43)))*int32(int16(ind_tmp))
					rate1_Q5 = rate0_Q5 + int32(43)
				}
			} else {
				if ind_tmp <= -int32(NLSF_QUANT_MAX_AMPLITUDE) {
					if ind_tmp == -int32(NLSF_QUANT_MAX_AMPLITUDE) {
						rate0_Q5 = int32(280)
						rate1_Q5 = int32(*(*OpusT_opus_uint8)(unsafe.Pointer(rates_Q5 + uintptr(ind_tmp+int32(1)+int32(NLSF_QUANT_MAX_AMPLITUDE)))))
					} else {
						rate0_Q5 = int32(280) - int32(43)*int32(NLSF_QUANT_MAX_AMPLITUDE) + int32(int16(-int32(43)))*int32(int16(ind_tmp))
						rate1_Q5 = rate0_Q5 - int32(43)
					}
				} else {
					rate0_Q5 = int32(*(*OpusT_opus_uint8)(unsafe.Pointer(rates_Q5 + uintptr(ind_tmp+int32(NLSF_QUANT_MAX_AMPLITUDE)))))
					rate1_Q5 = int32(*(*OpusT_opus_uint8)(unsafe.Pointer(rates_Q5 + uintptr(ind_tmp+int32(1)+int32(NLSF_QUANT_MAX_AMPLITUDE)))))
				}
			}
			RD_tmp_Q25 = RD_Q25[j]
			diff_Q10 = in_Q10 - int32(out0_Q10)
			RD_Q25[j] = RD_tmp_Q25 + int32(int16(diff_Q10))*int32(int16(diff_Q10))*int32(*(*OpusT_opus_int16)(unsafe.Pointer(w_Q5 + uintptr(i)*2))) + int32(int16(mu_Q20))*int32(int16(rate0_Q5))
			diff_Q10 = in_Q10 - int32(out1_Q10)
			RD_Q25[j+nStates] = RD_tmp_Q25 + int32(int16(diff_Q10))*int32(int16(diff_Q10))*int32(*(*OpusT_opus_int16)(unsafe.Pointer(w_Q5 + uintptr(i)*2))) + int32(int16(mu_Q20))*int32(int16(rate1_Q5))
			j = j + 1
		}
		if nStates <= int32(1)<<int32(NLSF_QUANT_DEL_DEC_STATES_LOG2)/int32(2) {
			/* double number of states and copy */
			j = 0
			for {
				if !(j < nStates) {
					break
				}
				*(*OpusT_opus_int8)(unsafe.Pointer(bp + uintptr(j+nStates)*16 + uintptr(i))) = int8(int32(*(*OpusT_opus_int8)(unsafe.Pointer(bp + uintptr(j)*16 + uintptr(i)))) + int32(1))
				j = j + 1
			}
			nStates = int32(uint32(nStates) << int32(1))
			j = nStates
			for {
				if !(j < int32(1)<<int32(NLSF_QUANT_DEL_DEC_STATES_LOG2)) {
					break
				}
				*(*OpusT_opus_int8)(unsafe.Pointer(bp + uintptr(j)*16 + uintptr(i))) = *(*OpusT_opus_int8)(unsafe.Pointer(bp + uintptr(j-nStates)*16 + uintptr(i)))
				j = j + 1
			}
		} else {
			/* sort lower and upper half of RD_Q25, pairwise */
			j = 0
			for {
				if !(j < int32(1)<<int32(NLSF_QUANT_DEL_DEC_STATES_LOG2)) {
					break
				}
				if RD_Q25[j] > RD_Q25[j+int32(1)<<int32(NLSF_QUANT_DEL_DEC_STATES_LOG2)] {
					RD_max_Q25[j] = RD_Q25[j]
					RD_min_Q25[j] = RD_Q25[j+int32(1)<<int32(NLSF_QUANT_DEL_DEC_STATES_LOG2)]
					RD_Q25[j] = RD_min_Q25[j]
					RD_Q25[j+int32(1)<<int32(NLSF_QUANT_DEL_DEC_STATES_LOG2)] = RD_max_Q25[j]
					/* swap prev_out values */
					out0_Q10 = prev_out_Q10[j]
					prev_out_Q10[j] = prev_out_Q10[j+int32(1)<<int32(NLSF_QUANT_DEL_DEC_STATES_LOG2)]
					prev_out_Q10[j+int32(1)<<int32(NLSF_QUANT_DEL_DEC_STATES_LOG2)] = out0_Q10
					ind_sort[j] = j + int32(1)<<int32(NLSF_QUANT_DEL_DEC_STATES_LOG2)
				} else {
					RD_min_Q25[j] = RD_Q25[j]
					RD_max_Q25[j] = RD_Q25[j+int32(1)<<int32(NLSF_QUANT_DEL_DEC_STATES_LOG2)]
					ind_sort[j] = j
				}
				j = j + 1
			}
			/* compare the highest RD values of the winning half with the lowest one in the losing half, and copy if necessary */
			/* afterwards ind_sort[] will contain the indices of the NLSF_QUANT_DEL_DEC_STATES winning RD values */
			for int32(1) != 0 {
				min_max_Q25 = int32(silk_int32_MAX)
				max_min_Q25 = 0
				ind_min_max = 0
				ind_max_min = 0
				j = 0
				for {
					if !(j < int32(1)<<int32(NLSF_QUANT_DEL_DEC_STATES_LOG2)) {
						break
					}
					if min_max_Q25 > RD_max_Q25[j] {
						min_max_Q25 = RD_max_Q25[j]
						ind_min_max = j
					}
					if max_min_Q25 < RD_min_Q25[j] {
						max_min_Q25 = RD_min_Q25[j]
						ind_max_min = j
					}
					j = j + 1
				}
				if min_max_Q25 >= max_min_Q25 {
					break
				}
				/* copy ind_min_max to ind_max_min */
				ind_sort[ind_max_min] = ind_sort[ind_min_max] ^ int32(1)<<int32(NLSF_QUANT_DEL_DEC_STATES_LOG2)
				RD_Q25[ind_max_min] = RD_Q25[ind_min_max+int32(1)<<int32(NLSF_QUANT_DEL_DEC_STATES_LOG2)]
				prev_out_Q10[ind_max_min] = prev_out_Q10[ind_min_max+int32(1)<<int32(NLSF_QUANT_DEL_DEC_STATES_LOG2)]
				RD_min_Q25[ind_max_min] = 0
				RD_max_Q25[ind_min_max] = int32(silk_int32_MAX)
				libc.Xmemcpy(tls, bp+uintptr(ind_max_min)*16, bp+uintptr(ind_min_max)*16, uint64(uint32(MAX_LPC_ORDER))*uint64(1))
			}
			/* increment index if it comes from the upper half */
			j = 0
			for {
				if !(j < int32(1)<<int32(NLSF_QUANT_DEL_DEC_STATES_LOG2)) {
					break
				}
				v11 = bp + uintptr(j)*16 + uintptr(i)
				*(*OpusT_opus_int8)(unsafe.Pointer(v11)) = OpusT_opus_int8(int32(*(*OpusT_opus_int8)(unsafe.Pointer(v11))) + ind_sort[j]>>int32(NLSF_QUANT_DEL_DEC_STATES_LOG2))
				j = j + 1
			}
		}
		i = i - 1
	}
	/* last sample: find winner, copy indices and return RD value */
	ind_tmp = 0
	min_Q25 = int32(silk_int32_MAX)
	j = 0
	for {
		if !(j < int32(2)*(int32(1)<<int32(NLSF_QUANT_DEL_DEC_STATES_LOG2))) {
			break
		}
		if min_Q25 > RD_Q25[j] {
			min_Q25 = RD_Q25[j]
			ind_tmp = j
		}
		j = j + 1
	}
	j = 0
	for {
		if !(j < int32(order)) {
			break
		}
		*(*OpusT_opus_int8)(unsafe.Pointer(indices + uintptr(j))) = *(*OpusT_opus_int8)(unsafe.Pointer(bp + uintptr(ind_tmp&(int32(1)<<int32(NLSF_QUANT_DEL_DEC_STATES_LOG2)-int32(1)))*16 + uintptr(j)))
		_ = int32(*(*OpusT_opus_int8)(unsafe.Pointer(indices + uintptr(j)))) >= -int32(NLSF_QUANT_MAX_AMPLITUDE_EXT)
		_ = int32(*(*OpusT_opus_int8)(unsafe.Pointer(indices + uintptr(j)))) <= int32(NLSF_QUANT_MAX_AMPLITUDE_EXT)
		j = j + 1
	}
	v11 = indices
	*(*OpusT_opus_int8)(unsafe.Pointer(v11)) = OpusT_opus_int8(int32(*(*OpusT_opus_int8)(unsafe.Pointer(v11))) + ind_tmp>>int32(NLSF_QUANT_DEL_DEC_STATES_LOG2))
	_ = int32(*(*OpusT_opus_int8)(unsafe.Pointer(indices))) <= int32(NLSF_QUANT_MAX_AMPLITUDE_EXT)
	_ = min_Q25 >= int32(0)
	return min_Q25
}

const silk_int16_MAX11 = 32767

// C documentation
//
//	/* Convert adaptive Mid/Side representation to Left/Right stereo signal */
