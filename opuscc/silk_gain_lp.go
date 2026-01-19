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

func Opus_silk_gains_quant(tls *libc.TLS, ind uintptr, gain_Q16 uintptr, prev_ind uintptr, conditional int32, nb_subfr int32) {
	var double_step_size_threshold, k, v2, v3, v4, v5, v6 int32
	var v11 uintptr
	var v19, v20, v21 OpusT_opus_int32
	_, _, _, _, _, _, _, _, _, _, _ = double_step_size_threshold, k, v11, v19, v2, v20, v21, v3, v4, v5, v6
	k = 0
	for {
		if !(k < nb_subfr) {
			break
		}
		/* Convert to log scale, scale, floor() */
		*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k))) = int8(int32(int64(int32(65536)*(int32(N_LEVELS_QGAIN)-int32(1))/((int32(MAX_QGAIN_DB)-int32(MIN_QGAIN_DB))*int32(128)/int32(6))) * int64(int16(Opus_silk_lin2log(tls, *(*OpusT_opus_int32)(unsafe.Pointer(gain_Q16 + uintptr(k)*4)))-(int32(MIN_QGAIN_DB)*int32(128)/int32(6)+int32(16)*int32(128)))) >> int32(16)))
		/* Round towards previous quantized gain (hysteresis) */
		if int32(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k)))) < int32(*(*OpusT_opus_int8)(unsafe.Pointer(prev_ind))) {
			*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k))) = *(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k))) + 1
		}
		if int32(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k)))) > int32(N_LEVELS_QGAIN)-int32(1) {
			v2 = int32(N_LEVELS_QGAIN) - int32(1)
		} else {
			if int32(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k)))) < 0 {
				v3 = 0
			} else {
				v3 = int32(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k))))
			}
			v2 = v3
		}
		*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k))) = int8(v2)
		/* Compute delta indices and limit */
		if k == 0 && conditional == 0 {
			/* Full index */
			if int32(*(*OpusT_opus_int8)(unsafe.Pointer(prev_ind)))+-int32(4) > int32(N_LEVELS_QGAIN)-int32(1) {
				if int32(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k)))) > int32(*(*OpusT_opus_int8)(unsafe.Pointer(prev_ind)))+-int32(4) {
					v3 = int32(*(*OpusT_opus_int8)(unsafe.Pointer(prev_ind))) + -int32(4)
				} else {
					if int32(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k)))) < int32(N_LEVELS_QGAIN)-int32(1) {
						v4 = int32(N_LEVELS_QGAIN) - int32(1)
					} else {
						v4 = int32(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k))))
					}
					v3 = v4
				}
				v2 = v3
			} else {
				if int32(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k)))) > int32(N_LEVELS_QGAIN)-int32(1) {
					v5 = int32(N_LEVELS_QGAIN) - int32(1)
				} else {
					if int32(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k)))) < int32(*(*OpusT_opus_int8)(unsafe.Pointer(prev_ind)))+-int32(4) {
						v6 = int32(*(*OpusT_opus_int8)(unsafe.Pointer(prev_ind))) + -int32(4)
					} else {
						v6 = int32(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k))))
					}
					v5 = v6
				}
				v2 = v5
			}
			*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k))) = int8(v2)
			*(*OpusT_opus_int8)(unsafe.Pointer(prev_ind)) = *(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k)))
		} else {
			/* Delta index */
			*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k))) = int8(int32(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k)))) - int32(*(*OpusT_opus_int8)(unsafe.Pointer(prev_ind))))
			/* Double the quantization step size for large gain increases, so that the max gain level can be reached */
			double_step_size_threshold = int32(2)*int32(MAX_DELTA_GAIN_QUANT) - int32(N_LEVELS_QGAIN) + int32(*(*OpusT_opus_int8)(unsafe.Pointer(prev_ind)))
			if int32(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k)))) > double_step_size_threshold {
				*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k))) = int8(double_step_size_threshold + (int32(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k))))-double_step_size_threshold+int32(1))>>int32(1))
			}
			if int32(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k)))) > int32(MAX_DELTA_GAIN_QUANT) {
				v2 = int32(MAX_DELTA_GAIN_QUANT)
			} else {
				if int32(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k)))) < -int32(4) {
					v3 = -int32(4)
				} else {
					v3 = int32(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k))))
				}
				v2 = v3
			}
			*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k))) = int8(v2)
			/* Accumulate deltas */
			if int32(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k)))) > double_step_size_threshold {
				v11 = prev_ind
				*(*OpusT_opus_int8)(unsafe.Pointer(v11)) = OpusT_opus_int8(int32(*(*OpusT_opus_int8)(unsafe.Pointer(v11))) + (int32(uint32(uint8(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k)))))<<int32(1)) - double_step_size_threshold))
				v2 = int32(*(*OpusT_opus_int8)(unsafe.Pointer(prev_ind)))
				v3 = int32(N_LEVELS_QGAIN) - int32(1)
				if v2 < v3 {
					v5 = v2
				} else {
					v5 = v3
				}
				v4 = v5
				goto _15
			_15:
				*(*OpusT_opus_int8)(unsafe.Pointer(prev_ind)) = int8(v4)
			} else {
				v11 = prev_ind
				*(*OpusT_opus_int8)(unsafe.Pointer(v11)) = OpusT_opus_int8(int32(*(*OpusT_opus_int8)(unsafe.Pointer(v11))) + int32(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k)))))
			}
			/* Shift to make non-negative */
			v11 = ind + uintptr(k)
			*(*OpusT_opus_int8)(unsafe.Pointer(v11)) = OpusT_opus_int8(int32(*(*OpusT_opus_int8)(unsafe.Pointer(v11))) - -int32(4))
		}
		/* Scale and convert to linear scale */
		v19 = int32(int64(int32(65536)*((int32(MAX_QGAIN_DB)-int32(MIN_QGAIN_DB))*int32(128)/int32(6))/(int32(N_LEVELS_QGAIN)-int32(1)))*int64(int16(*(*OpusT_opus_int8)(unsafe.Pointer(prev_ind))))>>int32(16)) + (int32(MIN_QGAIN_DB)*int32(128)/int32(6) + int32(16)*int32(128))
		v20 = int32(3967)
		if v19 < v20 {
			v2 = v19
		} else {
			v2 = v20
		}
		v21 = v2
		goto _22
	_22:
		*(*OpusT_opus_int32)(unsafe.Pointer(gain_Q16 + uintptr(k)*4)) = Opus_silk_log2lin(tls, v21) /* 3967 = 31 in Q7 */
		goto _1
	_1:
		;
		k = k + 1
	}
}

// C documentation
//
//	/* Gains scalar dequantization, uniform on log scale */
func Opus_silk_gains_dequant(tls *libc.TLS, gain_Q16 uintptr, ind uintptr, prev_ind uintptr, conditional int32, nb_subfr int32) {
	var double_step_size_threshold, ind_tmp, k, v2, v3, v4, v6 int32
	var v11, v12, v13 OpusT_opus_int32
	var v7 uintptr
	_, _, _, _, _, _, _, _, _, _, _ = double_step_size_threshold, ind_tmp, k, v11, v12, v13, v2, v3, v4, v6, v7
	k = 0
	for {
		if !(k < nb_subfr) {
			break
		}
		if k == 0 && conditional == 0 {
			/* Gain index is not allowed to go down more than 16 steps (~21.8 dB) */
			v2 = int32(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k))))
			v3 = int32(*(*OpusT_opus_int8)(unsafe.Pointer(prev_ind))) - int32(16)
			if v2 > v3 {
				v6 = v2
			} else {
				v6 = v3
			}
			v4 = v6
			goto _5
		_5:
			*(*OpusT_opus_int8)(unsafe.Pointer(prev_ind)) = int8(v4)
		} else {
			/* Delta index */
			ind_tmp = int32(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k)))) + -int32(4)
			/* Accumulate deltas */
			double_step_size_threshold = int32(2)*int32(MAX_DELTA_GAIN_QUANT) - int32(N_LEVELS_QGAIN) + int32(*(*OpusT_opus_int8)(unsafe.Pointer(prev_ind)))
			if ind_tmp > double_step_size_threshold {
				v7 = prev_ind
				*(*OpusT_opus_int8)(unsafe.Pointer(v7)) = OpusT_opus_int8(int32(*(*OpusT_opus_int8)(unsafe.Pointer(v7))) + (int32(uint32(ind_tmp)<<int32(1)) - double_step_size_threshold))
			} else {
				v7 = prev_ind
				*(*OpusT_opus_int8)(unsafe.Pointer(v7)) = OpusT_opus_int8(int32(*(*OpusT_opus_int8)(unsafe.Pointer(v7))) + ind_tmp)
			}
		}
		if int32(*(*OpusT_opus_int8)(unsafe.Pointer(prev_ind))) > int32(N_LEVELS_QGAIN)-int32(1) {
			v2 = int32(N_LEVELS_QGAIN) - int32(1)
		} else {
			if int32(*(*OpusT_opus_int8)(unsafe.Pointer(prev_ind))) < 0 {
				v3 = 0
			} else {
				v3 = int32(*(*OpusT_opus_int8)(unsafe.Pointer(prev_ind)))
			}
			v2 = v3
		}
		*(*OpusT_opus_int8)(unsafe.Pointer(prev_ind)) = int8(v2)
		/* Scale and convert to linear scale */
		v11 = int32(int64(int32(65536)*((int32(MAX_QGAIN_DB)-int32(MIN_QGAIN_DB))*int32(128)/int32(6))/(int32(N_LEVELS_QGAIN)-int32(1)))*int64(int16(*(*OpusT_opus_int8)(unsafe.Pointer(prev_ind))))>>int32(16)) + (int32(MIN_QGAIN_DB)*int32(128)/int32(6) + int32(16)*int32(128))
		v12 = int32(3967)
		if v11 < v12 {
			v2 = v11
		} else {
			v2 = v12
		}
		v13 = v2
		goto _14
	_14:
		*(*OpusT_opus_int32)(unsafe.Pointer(gain_Q16 + uintptr(k)*4)) = Opus_silk_log2lin(tls, v13) /* 3967 = 31 in Q7 */
		goto _1
	_1:
		;
		k = k + 1
	}
}

// C documentation
//
//	/* Compute unique identifier of gain indices vector */
func Opus_silk_gains_ID(tls *libc.TLS, ind uintptr, nb_subfr int32) (r OpusT_opus_int32) {
	var gainsID OpusT_opus_int32
	var k int32
	_, _ = gainsID, k
	gainsID = 0
	k = 0
	for {
		if !(k < nb_subfr) {
			break
		}
		gainsID = int32(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k)))) + int32(uint32(gainsID)<<int32(8))
		goto _1
	_1:
		;
		k = k + 1
	}
	return gainsID
}

// C documentation
//
//	/* Interpolate two vectors */
func Opus_silk_interpolate(tls *libc.TLS, xi uintptr, x0 uintptr, x1 uintptr, ifact_Q2 int32, d int32) {
	var i int32
	_ = i
	if !(ifact_Q2 >= int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+6629, __ccgo_ts+6661, int32(45))
	}
	if !(ifact_Q2 <= int32(4)) {
		Opus_celt_fatal(tls, __ccgo_ts+6683, __ccgo_ts+6661, int32(46))
	}
	i = 0
	for {
		if !(i < d) {
			break
		}
		*(*OpusT_opus_int16)(unsafe.Pointer(xi + uintptr(i)*2)) = int16(int32(*(*OpusT_opus_int16)(unsafe.Pointer(x0 + uintptr(i)*2))) + int32(int16(int32(*(*OpusT_opus_int16)(unsafe.Pointer(x1 + uintptr(i)*2)))-int32(*(*OpusT_opus_int16)(unsafe.Pointer(x0 + uintptr(i)*2)))))*int32(int16(ifact_Q2))>>int32(2))
		goto _1
	_1:
		;
		i = i + 1
	}
}

const silk_int16_MAX5 = 32767

// C documentation
//
//	/* Helper function, interpolates the filter taps */
func silk_LP_interpolate_filter_taps(tls *libc.TLS, B_Q28 uintptr, A_Q28 uintptr, ind int32, fac_Q16 OpusT_opus_int32) {
	var na, nb, v3, v4 int32
	_, _, _, _ = na, nb, v3, v4
	if ind < int32(TRANSITION_INT_NUM)-int32(1) {
		if fac_Q16 > 0 {
			if fac_Q16 < int32(32768) { /* fac_Q16 is in range of a 16-bit int */
				/* Piece-wise linear interpolation of B and A */
				nb = 0
				for {
					if !(nb < int32(TRANSITION_NB)) {
						break
					}
					*(*OpusT_opus_int32)(unsafe.Pointer(B_Q28 + uintptr(nb)*4)) = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_silk_Transition_LP_B_Q28)) + uintptr(ind)*12 + uintptr(nb)*4))) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_silk_Transition_LP_B_Q28)) + uintptr(ind+int32(1))*12 + uintptr(nb)*4))-*(*OpusT_opus_int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_silk_Transition_LP_B_Q28)) + uintptr(ind)*12 + uintptr(nb)*4)))*int64(int16(fac_Q16))>>int32(16))
					goto _1
				_1:
					;
					nb = nb + 1
				}
				na = 0
				for {
					if !(na < int32(TRANSITION_NA)) {
						break
					}
					*(*OpusT_opus_int32)(unsafe.Pointer(A_Q28 + uintptr(na)*4)) = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_silk_Transition_LP_A_Q28)) + uintptr(ind)*8 + uintptr(na)*4))) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_silk_Transition_LP_A_Q28)) + uintptr(ind+int32(1))*8 + uintptr(na)*4))-*(*OpusT_opus_int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_silk_Transition_LP_A_Q28)) + uintptr(ind)*8 + uintptr(na)*4)))*int64(int16(fac_Q16))>>int32(16))
					goto _2
				_2:
					;
					na = na + 1
				}
			} else { /* ( fac_Q16 - ( 1 << 16 ) ) is in range of a 16-bit int */
				if fac_Q16-int32(1)<<int32(16) > int32(silk_int16_MAX5) {
					v3 = int32(silk_int16_MAX5)
				} else {
					if fac_Q16-int32(1)<<int32(16) < int32(int16(-32768)) {
						v4 = int32(int16(-32768))
					} else {
						v4 = fac_Q16 - int32(1)<<int32(16)
					}
					v3 = v4
				}
				_ = fac_Q16-int32(1)<<int32(16) == v3
				/* Piece-wise linear interpolation of B and A */
				nb = 0
				for {
					if !(nb < int32(TRANSITION_NB)) {
						break
					}
					*(*OpusT_opus_int32)(unsafe.Pointer(B_Q28 + uintptr(nb)*4)) = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_silk_Transition_LP_B_Q28)) + uintptr(ind+int32(1))*12 + uintptr(nb)*4))) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_silk_Transition_LP_B_Q28)) + uintptr(ind+int32(1))*12 + uintptr(nb)*4))-*(*OpusT_opus_int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_silk_Transition_LP_B_Q28)) + uintptr(ind)*12 + uintptr(nb)*4)))*int64(int16(fac_Q16-int32(1)<<int32(16)))>>int32(16))
					goto _5
				_5:
					;
					nb = nb + 1
				}
				na = 0
				for {
					if !(na < int32(TRANSITION_NA)) {
						break
					}
					*(*OpusT_opus_int32)(unsafe.Pointer(A_Q28 + uintptr(na)*4)) = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_silk_Transition_LP_A_Q28)) + uintptr(ind+int32(1))*8 + uintptr(na)*4))) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_silk_Transition_LP_A_Q28)) + uintptr(ind+int32(1))*8 + uintptr(na)*4))-*(*OpusT_opus_int32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_silk_Transition_LP_A_Q28)) + uintptr(ind)*8 + uintptr(na)*4)))*int64(int16(fac_Q16-int32(1)<<int32(16)))>>int32(16))
					goto _6
				_6:
					;
					na = na + 1
				}
			}
		} else {
			libc.Xmemcpy(tls, B_Q28, uintptr(unsafe.Pointer(&Opus_silk_Transition_LP_B_Q28))+uintptr(ind)*12, uint64(uint32(TRANSITION_NB))*uint64(4))
			libc.Xmemcpy(tls, A_Q28, uintptr(unsafe.Pointer(&Opus_silk_Transition_LP_A_Q28))+uintptr(ind)*8, uint64(uint32(TRANSITION_NA))*uint64(4))
		}
	} else {
		libc.Xmemcpy(tls, B_Q28, uintptr(unsafe.Pointer(&Opus_silk_Transition_LP_B_Q28))+uintptr(int32(TRANSITION_INT_NUM)-int32(1))*12, uint64(uint32(TRANSITION_NB))*uint64(4))
		libc.Xmemcpy(tls, A_Q28, uintptr(unsafe.Pointer(&Opus_silk_Transition_LP_A_Q28))+uintptr(int32(TRANSITION_INT_NUM)-int32(1))*8, uint64(uint32(TRANSITION_NA))*uint64(4))
	}
}

// C documentation
//
//	/* Low-pass filter with variable cutoff frequency based on  */
//	/* piece-wise linear interpolation between elliptic filters */
//	/* Start by setting psEncC->mode <> 0;                      */
//	/* Deactivate by setting psEncC->mode = 0;                  */
func Opus_silk_LP_variable_cutoff(tls *libc.TLS, psLP uintptr, frame uintptr, frame_length int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var fac_Q16 OpusT_opus_int32
	var ind, v1, v2 int32
	var _ /* A_Q28 at bp+16 */ [2]OpusT_opus_int32
	var _ /* B_Q28 at bp+0 */ [3]OpusT_opus_int32
	_, _, _, _ = fac_Q16, ind, v1, v2
	fac_Q16 = 0
	ind = 0
	_ = (*OpusT_silk_LP_state)(unsafe.Pointer(psLP)).Ftransition_frame_no >= 0 && (*OpusT_silk_LP_state)(unsafe.Pointer(psLP)).Ftransition_frame_no <= int32(TRANSITION_TIME_MS)/(int32(SUB_FRAME_LENGTH_MS)*int32(MAX_NB_SUBFR))
	/* Run filter if needed */
	if (*OpusT_silk_LP_state)(unsafe.Pointer(psLP)).Fmode != 0 {
		/* Calculate index and interpolation factor for interpolation */
		fac_Q16 = int32(uint32(int32(TRANSITION_TIME_MS)/(int32(SUB_FRAME_LENGTH_MS)*int32(MAX_NB_SUBFR))-(*OpusT_silk_LP_state)(unsafe.Pointer(psLP)).Ftransition_frame_no) << (int32(16) - int32(6)))
		ind = fac_Q16 >> int32(16)
		fac_Q16 = fac_Q16 - int32(uint32(ind)<<int32(16))
		_ = ind >= int32(0)
		_ = ind < int32(TRANSITION_INT_NUM)
		/* Interpolate filter coefficients */
		silk_LP_interpolate_filter_taps(tls, bp, bp+16, ind, fac_Q16)
		/* Update transition frame number for next frame */
		if (*OpusT_silk_LP_state)(unsafe.Pointer(psLP)).Ftransition_frame_no+(*OpusT_silk_LP_state)(unsafe.Pointer(psLP)).Fmode > int32(TRANSITION_TIME_MS)/(int32(SUB_FRAME_LENGTH_MS)*int32(MAX_NB_SUBFR)) {
			v1 = int32(TRANSITION_TIME_MS) / (int32(SUB_FRAME_LENGTH_MS) * int32(MAX_NB_SUBFR))
		} else {
			if (*OpusT_silk_LP_state)(unsafe.Pointer(psLP)).Ftransition_frame_no+(*OpusT_silk_LP_state)(unsafe.Pointer(psLP)).Fmode < 0 {
				v2 = 0
			} else {
				v2 = (*OpusT_silk_LP_state)(unsafe.Pointer(psLP)).Ftransition_frame_no + (*OpusT_silk_LP_state)(unsafe.Pointer(psLP)).Fmode
			}
			v1 = v2
		}
		(*OpusT_silk_LP_state)(unsafe.Pointer(psLP)).Ftransition_frame_no = v1
		/* ARMA low-pass filtering */
		_ = libc.Bool(true) && libc.Bool(true)
		Opus_silk_biquad_alt_stride1(tls, frame, bp, bp+16, psLP, frame, frame_length)
	}
}

const silk_int16_MAX6 = 0x7FFF

// C documentation
//
//	/* Predictive dequantizer for NLSF residuals */
func silk_NLSF_residual_dequant(tls *libc.TLS, x_Q10 uintptr, indices uintptr, pred_coef_Q8 uintptr, quant_step_size_Q16 int32, order OpusT_opus_int16) {
	var i, out_Q10, pred_Q10 int32
	_, _, _ = i, out_Q10, pred_Q10
	out_Q10 = 0
	i = int32(order) - int32(1)
	for {
		if !(i >= 0) {
			break
		}
		pred_Q10 = int32(int16(out_Q10)) * int32(int16(*(*OpusT_opus_uint8)(unsafe.Pointer(pred_coef_Q8 + uintptr(i))))) >> int32(8)
		out_Q10 = int32(uint32(uint8(*(*OpusT_opus_int8)(unsafe.Pointer(indices + uintptr(i))))) << int32(10))
		if out_Q10 > 0 {
			out_Q10 = out_Q10 - int32(102)
		} else {
			if out_Q10 < 0 {
				out_Q10 = out_Q10 + int32(102)
			}
		}
		out_Q10 = int32(int64(pred_Q10) + int64(out_Q10)*int64(int16(quant_step_size_Q16))>>int32(16))
		*(*OpusT_opus_int16)(unsafe.Pointer(x_Q10 + uintptr(i)*2)) = int16(out_Q10)
		goto _1
	_1:
		;
		i = i - 1
	}
}

// C documentation
//
//	/***********************/
//	/* NLSF vector decoder */
//	/***********************/
