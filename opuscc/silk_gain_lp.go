// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

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
				*(*OpusT_opus_int8)(unsafe.Pointer(v11)) = OpusT_opus_int8(int32(*(*OpusT_opus_int8)(unsafe.Pointer(v11))) + (int32(*(*OpusT_opus_int8)(unsafe.Pointer(ind + uintptr(k))))<<int32(1) - double_step_size_threshold))
				v2 = int32(*(*OpusT_opus_int8)(unsafe.Pointer(prev_ind)))
				v3 = int32(N_LEVELS_QGAIN) - int32(1)
				if v2 < v3 {
					v5 = v2
				} else {
					v5 = v3
				}
				v4 = v5
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
		*(*OpusT_opus_int32)(unsafe.Pointer(gain_Q16 + uintptr(k)*4)) = Opus_silk_log2lin(tls, v21) /* 3967 = 31 in Q7 */
		k = k + 1
	}
}

// C documentation
//
//	/* Gains scalar dequantization, uniform on log scale */
func Opus_silk_gains_dequant(tls *libc.TLS, gain_Q16 *OpusT_opus_int32, ind *OpusT_opus_int8, prev_ind *OpusT_opus_int8, conditional int32, nb_subfr int32) {
	if nb_subfr <= 0 {
		return
	}
	gains := unsafe.Slice(gain_Q16, int(nb_subfr))
	indices := unsafe.Slice(ind, int(nb_subfr))
	const inverseScale = int32(65536) * ((MAX_QGAIN_DB - MIN_QGAIN_DB) * 128 / 6) / (N_LEVELS_QGAIN - 1)
	const offset = int32(MIN_QGAIN_DB)*128/6 + 16*128
	for k := range gains {
		if k == 0 && conditional == 0 {
			*prev_ind = int8(max(int32(indices[k]), int32(*prev_ind)-16))
		} else {
			delta := int32(indices[k]) - 4
			threshold := 2*int32(MAX_DELTA_GAIN_QUANT) - N_LEVELS_QGAIN + int32(*prev_ind)
			if delta > threshold {
				*prev_ind = int8(int32(*prev_ind) + ((delta << 1) - threshold))
			} else {
				*prev_ind = int8(int32(*prev_ind) + delta)
			}
		}
		// C narrows the state to int8 before limiting it to the gain range.
		*prev_ind = int8(min(max(int32(*prev_ind), 0), N_LEVELS_QGAIN-1))
		logGain := int32((int64(inverseScale)*int64(*prev_ind))>>16) + offset
		gains[k] = Opus_silk_log2lin(tls, min(logGain, 3967))
	}
}

// C documentation
//
//	/* Compute unique identifier of gain indices vector */
func Opus_silk_gains_ID(tls *libc.TLS, ind *OpusT_opus_int8, nb_subfr int32) (r OpusT_opus_int32) {
	if nb_subfr <= 0 {
		return 0
	}
	var gainsID OpusT_opus_int32
	for _, index := range unsafe.Slice(ind, int(nb_subfr)) {
		gainsID = int32(index) + int32(uint32(gainsID)<<8)
	}
	return gainsID
}

// C documentation
//
//	/* Interpolate two vectors */
func Opus_silk_interpolate(tls *libc.TLS, xi *OpusT_opus_int16, x0 *OpusT_opus_int16, x1 *OpusT_opus_int16, ifact_Q2 int32, d int32) {
	if !(ifact_Q2 >= int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+6629, __ccgo_ts+6661, int32(45))
	}
	if !(ifact_Q2 <= int32(4)) {
		Opus_celt_fatal(tls, __ccgo_ts+6683, __ccgo_ts+6661, int32(46))
	}
	if d <= 0 {
		return
	}
	out := unsafe.Slice(xi, int(d))
	left, right := unsafe.Slice(x0, int(d)), unsafe.Slice(x1, int(d))
	for i := range out {
		// silk_SMULBB narrows the difference to signed 16 bits first.
		delta := int32(int16(int32(right[i]) - int32(left[i])))
		out[i] = int16(int32(left[i]) + ((delta * int32(int16(ifact_Q2))) >> 2))
	}
}

const silk_int16_MAX5 = 32767

// C documentation
//
//	/* Helper function, interpolates the filter taps */
func silk_LP_interpolate_filter_taps(tls *libc.TLS, B_Q28 *[3]OpusT_opus_int32, A_Q28 *[2]OpusT_opus_int32, ind int32, fac_Q16 OpusT_opus_int32) {
	if ind >= TRANSITION_INT_NUM-1 {
		*B_Q28 = Opus_silk_Transition_LP_B_Q28[TRANSITION_INT_NUM-1]
		*A_Q28 = Opus_silk_Transition_LP_A_Q28[TRANSITION_INT_NUM-1]
		return
	}
	if fac_Q16 <= 0 {
		*B_Q28 = Opus_silk_Transition_LP_B_Q28[ind]
		*A_Q28 = Opus_silk_Transition_LP_A_Q28[ind]
		return
	}
	base := ind
	factor := fac_Q16
	if factor >= 32768 {
		// Use the upper endpoint so the multiplier fits signed 16 bits.
		base++
		factor -= 1 << 16
	}
	for i := range B_Q28 {
		delta := Opus_silk_Transition_LP_B_Q28[ind+1][i] - Opus_silk_Transition_LP_B_Q28[ind][i]
		B_Q28[i] = int32(int64(Opus_silk_Transition_LP_B_Q28[base][i]) + ((int64(delta) * int64(int16(factor))) >> 16))
	}
	for i := range A_Q28 {
		delta := Opus_silk_Transition_LP_A_Q28[ind+1][i] - Opus_silk_Transition_LP_A_Q28[ind][i]
		A_Q28[i] = int32(int64(Opus_silk_Transition_LP_A_Q28[base][i]) + ((int64(delta) * int64(int16(factor))) >> 16))
	}
}

// C documentation
//
//	/* Low-pass filter with variable cutoff frequency based on  */
//	/* piece-wise linear interpolation between elliptic filters */
//	/* Start by setting psEncC->mode <> 0;                      */
//	/* Deactivate by setting psEncC->mode = 0;                  */
func Opus_silk_LP_variable_cutoff(tls *libc.TLS, psLP *OpusT_silk_LP_state, frame *OpusT_opus_int16, frame_length int32) {
	if psLP.Fmode == 0 {
		return
	}
	const transitionFrames = int32(TRANSITION_TIME_MS) / (int32(SUB_FRAME_LENGTH_MS) * int32(MAX_NB_SUBFR))
	// C requires a transition position in [0, transitionFrames].
	fac_Q16 := (transitionFrames - psLP.Ftransition_frame_no) << 10
	ind := fac_Q16 >> 16
	fac_Q16 -= ind << 16
	var B_Q28 [3]OpusT_opus_int32
	var A_Q28 [2]OpusT_opus_int32
	silk_LP_interpolate_filter_taps(tls, &B_Q28, &A_Q28, ind, fac_Q16)
	psLP.Ftransition_frame_no = min(max(psLP.Ftransition_frame_no+psLP.Fmode, 0), transitionFrames)
	Opus_silk_biquad_alt_stride1(tls, frame, &B_Q28, &A_Q28, &psLP.FIn_LP_State, frame, frame_length)
}

const silk_int16_MAX6 = 0x7FFF

// C documentation
//
//	/* Predictive dequantizer for NLSF residuals */
func silk_NLSF_residual_dequant(tls *libc.TLS, x_Q10 *OpusT_opus_int16, indices *OpusT_opus_int8, pred_coef_Q8 *OpusT_opus_uint8, quant_step_size_Q16 int32, order OpusT_opus_int16) {
	if order <= 0 {
		return
	}
	x := unsafe.Slice(x_Q10, int(order))
	idx := unsafe.Slice(indices, int(order))
	pred := unsafe.Slice(pred_coef_Q8, int(order))
	var out_Q10 int32
	for i := int(order) - 1; i >= 0; i-- {
		// Match silk_SMULBB's signed 16-bit narrowing before prediction.
		pred_Q10 := (int32(int16(out_Q10)) * int32(pred[i])) >> 8
		out_Q10 = int32(idx[i]) << 10
		if out_Q10 > 0 {
			out_Q10 -= 102
		} else if out_Q10 < 0 {
			out_Q10 += 102
		}
		// silk_SMLAWB uses the signed low 16 bits of the quantization step.
		out_Q10 = int32(int64(pred_Q10) + ((int64(out_Q10) * int64(int16(quant_step_size_Q16))) >> 16))
		x[i] = int16(out_Q10)
	}
}

// C documentation
//
//	/***********************/
//	/* NLSF vector decoder */
//	/***********************/
