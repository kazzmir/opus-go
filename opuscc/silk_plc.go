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

func Opus_silk_PLC_Reset(tls *libc.TLS, decoder *OpusT_silk_decoder_state) {
	plc := &decoder.FsPLC
	plc.FpitchL_Q8 = int32(uint32(decoder.Fframe_length) << (int32(8) - int32(1)))
	plc.FprevGain_Q16[0] = 65536
	plc.FprevGain_Q16[1] = 65536
	plc.Fsubfr_length = 20
	plc.Fnb_subfr = 2
}

//go:uintptrescapes
func Opus_silk_PLC(tls *libc.TLS, psDec, psDecCtrl, frame uintptr, lost, arch int32) {
	silk_PLC(tls, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)), (*OpusT_silk_decoder_control)(unsafe.Pointer(psDecCtrl)), (*int16)(unsafe.Pointer(frame)), lost, arch)
}
func silkPLCRate(tls *libc.TLS, decoder *OpusT_silk_decoder_state) {
	if decoder.Ffs_kHz != decoder.FsPLC.Ffs_kHz {
		Opus_silk_PLC_Reset(tls, decoder)
		decoder.FsPLC.Ffs_kHz = decoder.Ffs_kHz
	}
}
func silk_PLC(tls *libc.TLS, decoder *OpusT_silk_decoder_state, control *OpusT_silk_decoder_control, frame *int16, lost, arch int32) {
	silkPLCRate(tls, decoder)
	if lost != 0 {
		/****************************/
		/* Generate Signal          */
		/****************************/
		silk_PLC_conceal(tls, decoder, control, frame, arch)
		decoder.FlossCnt = decoder.FlossCnt + 1
	} else {
		/****************************/
		/* Update state             */
		/****************************/
		silk_PLC_update(tls, decoder, control)
	}
}

// C documentation
//
//	/**************************************************/
//	/* Update state of PLC                            */
//	/**************************************************/
func silk_PLC_update(tls *libc.TLS, decoder *OpusT_silk_decoder_state, control *OpusT_silk_decoder_control) {
	var LTP_Gain_Q14, temp_LTP_Gain_Q14, tmp, tmp1 OpusT_opus_int32
	var i, j, scale_Q10, scale_Q14, v3 int32
	_, _, _, _, _, _, _, _, _ = LTP_Gain_Q14, i, j, scale_Q10, scale_Q14, temp_LTP_Gain_Q14, tmp, tmp1, v3
	plc := &decoder.FsPLC
	/* Update parameters used in case of packet loss */
	decoder.FprevSignalType = int32(decoder.Findices.FsignalType)
	LTP_Gain_Q14 = 0
	if int32(decoder.Findices.FsignalType) == int32(TYPE_VOICED) {
		/* Find the parameters for the last subframe which contains a pitch pulse */
		j = 0
		for {
			if !(j*decoder.Fsubfr_length < control.FpitchL[decoder.Fnb_subfr-int32(1)]) {
				break
			}
			if j == decoder.Fnb_subfr {
				break
			}
			temp_LTP_Gain_Q14 = 0
			i = 0
			for {
				if !(i < int32(LTP_ORDER)) {
					break
				}
				temp_LTP_Gain_Q14 = temp_LTP_Gain_Q14 + int32(control.FLTPCoef_Q14[(decoder.Fnb_subfr-int32(1)-j)*int32(LTP_ORDER)+i])
				i = i + 1
			}
			if temp_LTP_Gain_Q14 > LTP_Gain_Q14 {
				LTP_Gain_Q14 = temp_LTP_Gain_Q14
				copy(plc.FLTPCoef_Q14[:], control.FLTPCoef_Q14[(decoder.Fnb_subfr-1-j)*LTP_ORDER:(decoder.Fnb_subfr-j)*LTP_ORDER])
				plc.FpitchL_Q8 = int32(uint32(control.FpitchL[decoder.Fnb_subfr-int32(1)-j]) << int32(8))
			}
			j = j + 1
		}
		clear(plc.FLTPCoef_Q14[:])
		plc.FLTPCoef_Q14[int32(LTP_ORDER)/int32(2)] = int16(LTP_Gain_Q14)
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
				plc.FLTPCoef_Q14[i] = int16(int32(plc.FLTPCoef_Q14[i]) * int32(int16(scale_Q10)) >> int32(10))
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
					plc.FLTPCoef_Q14[i] = int16(int32(plc.FLTPCoef_Q14[i]) * int32(int16(scale_Q14)) >> int32(14))
					i = i + 1
				}
			}
		}
	} else {
		plc.FpitchL_Q8 = int32(uint32(int32(int16(decoder.Ffs_kHz))*int32(int16(int32(18)))) << int32(8))
		clear(plc.FLTPCoef_Q14[:])
	}
	/* Save LPC coefficients */
	copy(plc.FprevLPC_Q12[:decoder.FLPC_order], control.FPredCoef_Q12[1][:decoder.FLPC_order])
	plc.FprevLTP_scale_Q14 = int16(control.FLTP_scale_Q14)
	/* Save last two gains */
	copy(plc.FprevGain_Q16[:], control.FGains_Q16[decoder.Fnb_subfr-2:decoder.Fnb_subfr])
	plc.Fsubfr_length = decoder.Fsubfr_length
	plc.Fnb_subfr = decoder.Fnb_subfr
}

func silk_PLC_energy(tls *libc.TLS, energy1, shift1, energy2, shift2 *int32, exc_Q14 *OpusT_opus_int32, prevGain_Q10 *[2]OpusT_opus_int32, subfr_length, nb_subfr int32) {
	excitation := unsafe.Slice(exc_Q14, subfr_length*nb_subfr)
	buffer := make([]OpusT_opus_int16, 2*subfr_length)
	for k := int32(0); k < 2; k++ {
		for i := int32(0); i < subfr_length; i++ {
			// SMULWW narrows to int32 before the Q8 shift and int16 saturation.
			value := int32((int64(excitation[i+(k+nb_subfr-2)*subfr_length])*int64(prevGain_Q10[k]))>>16) >> 8
			buffer[k*subfr_length+i] = int16(min(max(value, -32768), 32767))
		}
	}
	// Finish scaling both subframes before writing any output, as in C.
	Opus_silk_sum_sqr_shift(tls, energy1, shift1, unsafe.SliceData(buffer), subfr_length)
	Opus_silk_sum_sqr_shift(tls, energy2, shift2, unsafe.SliceData(buffer[subfr_length:]), subfr_length)
}

func silkPLCAddPrediction(excitation, prediction int32) int32 {
	prediction = min(max(prediction, -2147483648>>4), 2147483647>>4)
	shifted := int32(uint32(prediction) << 4)
	return int32(min(max(int64(excitation)+int64(shifted), int64(-2147483648)), int64(2147483647)))
}
func silkPLCLPC(tls *libc.TLS, decoder *OpusT_silk_decoder_state, history []int32, A *[MAX_LPC_ORDER]int16, frame *int16, gain int32) {
	copy(history[:MAX_LPC_ORDER], decoder.FsLPC_Q14_buf[:])
	if decoder.FLPC_order < 10 {
		Opus_celt_fatal(tls, __ccgo_ts+6755, __ccgo_ts+6715, 373)
	}
	pcm := unsafe.Slice(frame, decoder.Fframe_length)
	for i := int32(0); i < decoder.Fframe_length; i++ {
		prediction := decoder.FLPC_order >> 1
		for j := int32(0); j < 10; j++ {
			prediction = int32(int64(prediction) + (int64(history[MAX_LPC_ORDER+i-j-1]) * int64(A[j]) >> 16))
		}
		for j := int32(10); j < decoder.FLPC_order; j++ {
			prediction = int32(int64(prediction) + (int64(history[MAX_LPC_ORDER+i-j-1]) * int64(A[j]) >> 16))
		}
		history[MAX_LPC_ORDER+i] = silkPLCAddPrediction(history[MAX_LPC_ORDER+i], prediction)
		silkPLCPCMStore(pcm, i, history[MAX_LPC_ORDER+i], gain)
	}
	copy(decoder.FsLPC_Q14_buf[:], history[decoder.Fframe_length:decoder.Fframe_length+MAX_LPC_ORDER])
}
func silkPLCPCMStore(pcm []int16, index, sample, gain int32) { pcm[index] = silkPLCPCM(sample, gain) }
func silkPLCPCM(sample, gain int32) int16 {
	scaled := int32(int64(sample) * int64(gain) >> 16)
	rounded := ((scaled >> 7) + 1) >> 1
	return int16(min(max(rounded, -32768), 32767))
}
func silkPLCRandom(decoder *OpusT_silk_decoder_state, offset int32) []int32 {
	return decoder.Fexc_Q14[offset : offset+RAND_BUF_SIZE]
}
func silkPLCNoise(prediction int32, random []int32, index int32, scale int16) int32 {
	return int32(uint32(int32(int64(prediction)+(int64(random[index])*int64(scale)>>16))) << 2)
}
func silkPLCDecayLTP(coefficients *[LTP_ORDER]int16, gain int32) {
	for i := range coefficients {
		coefficients[i] = int16(int32(int16(gain)) * int32(coefficients[i]) >> 15)
	}
}
func silkPLCSynthesisBuffer(decoder *OpusT_silk_decoder_state) []int32 {
	return make([]int32, decoder.Fltp_mem_length+decoder.Fframe_length)
}
func silkPLCLTPPrediction(history []int32, index int32, coefficients *[LTP_ORDER]int16) int32 {
	prediction := int32(2)
	for j := int32(0); j < LTP_ORDER; j++ {
		prediction = int32(int64(prediction) + (int64(history[index-j]) * int64(coefficients[j]) >> 16))
	}
	return prediction
}
func silkPLCWhiten(tls *libc.TLS, decoder *OpusT_silk_decoder_state, samples []int16, A *[MAX_LPC_ORDER]int16, index, arch int32) {
	Opus_silk_LPC_analysis_filter(tls, unsafe.SliceData(samples[index:]), &decoder.FoutBuf[index], &A[0], decoder.Fltp_mem_length-index, decoder.FLPC_order, arch)
}
func silkPLCConcealState(decoder *OpusT_silk_decoder_state) *OpusT_silk_PLC_struct {
	return &decoder.FsPLC
}

// Both scratch arrays are Go-owned; no TLS cursor is consumed or restored.
func silk_PLC_conceal(tls *libc.TLS, psDec *OpusT_silk_decoder_state, psDecCtrl *OpusT_silk_decoder_control, frame *int16, arch int32) {
	var psPLC *OpusT_silk_PLC_struct
	var B_Q14 *[LTP_ORDER]int16
	var rand_ptr []int32
	var sLTP []int16
	var sLTP_Q14 []int32
	var pred_index int32
	var LTP_pred_Q12, b32_inv, b32_nrm, down_scale_Q30, err_Q32, harm_Gain_Q15, invGain_Q30, inv_gain_Q30, rand_Gain_Q15, rand_seed, result, v84, v85, v86, v89 OpusT_opus_int32
	var b_headrm, i, idx, k, lag, lshift, sLTP_buf_idx, v53, v54, v55, v57, v58, v59, v60, v62 int32
	var rand_scale_Q14, v79, v80, v81 OpusT_opus_int16
	var energy1, energy2 OpusT_opus_int32
	var shift1, shift2 int32
	var A_Q12 [MAX_LPC_ORDER]OpusT_opus_int16
	var prevGain_Q10 [2]OpusT_opus_int32
	decoder, control := psDec, psDecCtrl
	plc := silkPLCConcealState(decoder)
	psPLC = plc
	sLTP_Q14 = silkPLCSynthesisBuffer(decoder)
	sLTP = make([]int16, decoder.Fltp_mem_length)
	prevGain_Q10[0] = plc.FprevGain_Q16[0] >> int32(6)
	prevGain_Q10[1] = plc.FprevGain_Q16[1] >> int32(6)
	if decoder.Ffirst_frame_after_reset != 0 {
		clear(plc.FprevLPC_Q12[:])
	}
	silk_PLC_energy(tls, &energy1, &shift1, &energy2, &shift2, &decoder.Fexc_Q14[0], &prevGain_Q10, decoder.Fsubfr_length, decoder.Fnb_subfr)
	if energy1>>shift2 < energy2>>shift1 {
		/* First sub-frame has lowest energy */
		v53 = 0
		v54 = (psPLC.Fnb_subfr-int32(1))*psPLC.Fsubfr_length - int32(RAND_BUF_SIZE)
		if v53 > v54 {
			v57 = v53
		} else {
			v57 = v54
		}
		v55 = v57
		rand_ptr = silkPLCRandom(decoder, v55)
	} else {
		/* Second sub-frame has lowest energy */
		v53 = 0
		v54 = psPLC.Fnb_subfr*psPLC.Fsubfr_length - int32(RAND_BUF_SIZE)
		if v53 > v54 {
			v57 = v53
		} else {
			v57 = v54
		}
		v55 = v57
		rand_ptr = silkPLCRandom(decoder, v55)
	}
	/* Set up Gain to random noise component */
	B_Q14 = &plc.FLTPCoef_Q14
	rand_scale_Q14 = plc.FrandScale_Q14
	/* Set up attenuation gains */
	v53 = int32(NB_ATT) - int32(1)
	v54 = psDec.FlossCnt
	if v53 < v54 {
		v57 = v53
	} else {
		v57 = v54
	}
	v55 = v57
	harm_Gain_Q15 = int32(HARM_ATT_Q15[v55])
	if psDec.FprevSignalType == int32(TYPE_VOICED) {
		v53 = int32(NB_ATT) - int32(1)
		v54 = psDec.FlossCnt
		if v53 < v54 {
			v57 = v53
		} else {
			v57 = v54
		}
		v55 = v57
		rand_Gain_Q15 = int32(PLC_RAND_ATTENUATE_V_Q15[v55])
	} else {
		v53 = int32(NB_ATT) - int32(1)
		v54 = psDec.FlossCnt
		if v53 < v54 {
			v57 = v53
		} else {
			v57 = v54
		}
		v55 = v57
		rand_Gain_Q15 = int32(PLC_RAND_ATTENUATE_UV_Q15[v55])
	}
	/* LPC concealment. Apply BWE to previous LPC */
	Opus_silk_bwexpander(tls, &plc.FprevLPC_Q12[0], decoder.FLPC_order, int32(64881))
	/* Preload LPC coefficients to array on stack. Gives small performance gain */
	copy(A_Q12[:decoder.FLPC_order], plc.FprevLPC_Q12[:decoder.FLPC_order])
	/* First Lost frame */
	if psDec.FlossCnt == 0 {
		rand_scale_Q14 = int16(int32(1) << int32(14))
		/* Reduce random noise Gain for voiced frames */
		if psDec.FprevSignalType == int32(TYPE_VOICED) {
			i = 0
			for {
				if !(i < int32(LTP_ORDER)) {
					break
				}
				rand_scale_Q14 = int16(int32(rand_scale_Q14) - int32(B_Q14[i]))
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
			rand_scale_Q14 = int16(int32(rand_scale_Q14) * int32(psPLC.FprevLTP_scale_Q14) >> int32(14))
		} else {
			_ = arch
			invGain_Q30 = Opus_silk_LPC_inverse_pred_gain_c(tls, &plc.FprevLPC_Q12[0], decoder.FLPC_order)
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
	rand_seed = psPLC.Frand_seed
	lag = (psPLC.FpitchL_Q8>>(int32(8)-int32(1)) + int32(1)) >> int32(1)
	sLTP_buf_idx = psDec.Fltp_mem_length
	/* Rewhiten LTP state */
	idx = psDec.Fltp_mem_length - lag - psDec.FLPC_order - int32(LTP_ORDER)/int32(2)
	if !(idx > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+6729, __ccgo_ts+6715, int32(319))
	}
	silkPLCWhiten(tls, decoder, sLTP, &A_Q12, idx, arch)
	/* Scale LTP state */
	v84 = plc.FprevGain_Q16[1]
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
	i = idx + psDec.FLPC_order
	for {
		if !(i < psDec.Fltp_mem_length) {
			break
		}
		sLTP_Q14[i] = int32(int64(inv_gain_Q30) * int64(sLTP[i]) >> 16)
		i = i + 1
	}
	/***************************/
	/* LTP synthesis filtering */
	/***************************/
	k = 0
	for {
		if !(k < psDec.Fnb_subfr) {
			break
		}
		/* Set up pointer */
		pred_index = sLTP_buf_idx - lag + LTP_ORDER/2
		i = 0
		for {
			if !(i < psDec.Fsubfr_length) {
				break
			}
			/* Unrolled loop */
			/* Avoids introducing a bias because silk_SMLAWB() always rounds to -inf */
			LTP_pred_Q12 = silkPLCLTPPrediction(sLTP_Q14, pred_index, B_Q14)
			pred_index++
			/* Generate LPC excitation */
			rand_seed = int32(uint32(int32(RAND_INCREMENT)) + uint32(rand_seed)*uint32(int32(RAND_MULTIPLIER)))
			idx = rand_seed >> int32(25) & (int32(RAND_BUF_SIZE) - int32(1))
			sLTP_Q14[sLTP_buf_idx] = silkPLCNoise(LTP_pred_Q12, rand_ptr, idx, rand_scale_Q14)
			sLTP_buf_idx = sLTP_buf_idx + 1
			i = i + 1
		}
		/* Gradually reduce LTP gain */
		silkPLCDecayLTP(B_Q14, harm_Gain_Q15)
		/* Gradually reduce excitation gain */
		rand_scale_Q14 = int16(int32(rand_scale_Q14) * int32(int16(rand_Gain_Q15)) >> int32(15))
		/* Slowly increase pitch lag */
		psPLC.FpitchL_Q8 = int32(int64(psPLC.FpitchL_Q8) + int64(psPLC.FpitchL_Q8)*int64(int16(int32(PITCH_DRIFT_FAC_Q16)))>>int32(16))
		v84 = psPLC.FpitchL_Q8
		v85 = int32(uint32(int32(int16(int32(MAX_PITCH_LAG_MS)))*int32(int16(psDec.Ffs_kHz))) << int32(8))
		if v84 < v85 {
			v53 = v84
		} else {
			v53 = v85
		}
		v86 = v53
		psPLC.FpitchL_Q8 = v86
		lag = (psPLC.FpitchL_Q8>>(int32(8)-int32(1)) + int32(1)) >> int32(1)
		k = k + 1
	}
	// LPC uses the live tail of the same typed synthesis buffer.
	history := sLTP_Q14[decoder.Fltp_mem_length-MAX_LPC_ORDER : decoder.Fltp_mem_length+decoder.Fframe_length]
	silkPLCLPC(tls, decoder, history, &A_Q12, frame, prevGain_Q10[1])
	/**************************************/
	/* Update states                      */
	/**************************************/
	plc.Frand_seed = rand_seed
	plc.FrandScale_Q14 = rand_scale_Q14
	i = 0
	for {
		if !(i < int32(MAX_NB_SUBFR)) {
			break
		}
		control.FpitchL[i] = lag
		i = i + 1
	}
}

// C documentation
//
//	/* Glues concealed frames with new good received frames */
func Opus_silk_PLC_glue_frames(tls *libc.TLS, dec *OpusT_silk_decoder_state, frame *int16, length int32) {
	plc := &dec.FsPLC
	if dec.FlossCnt != 0 {
		Opus_silk_sum_sqr_shift(tls, &plc.Fconc_energy, &plc.Fconc_energy_shift, frame, length)
		plc.Flast_frame_lost = 1
		return
	}
	if plc.Flast_frame_lost != 0 {
		var energy, shift int32
		Opus_silk_sum_sqr_shift(tls, &energy, &shift, frame, length)
		if shift > plc.Fconc_energy_shift {
			plc.Fconc_energy >>= shift - plc.Fconc_energy_shift
		} else if shift < plc.Fconc_energy_shift {
			energy >>= plc.Fconc_energy_shift - shift
		}
		if energy > plc.Fconc_energy {
			lz := int32(bits.LeadingZeros32(uint32(plc.Fconc_energy))) - 1
			plc.Fconc_energy = int32(uint32(plc.Fconc_energy) << lz)
			energy >>= max(24-lz, int32(0))
			fraction := plc.Fconc_energy / max(energy, int32(1))
			root := int32(0)
			if fraction > 0 {
				zeros := bits.LeadingZeros32(uint32(fraction))
				fracQ7 := int32(bits.RotateLeft32(uint32(fraction), zeros-24)) & 127
				root = 46214
				if zeros&1 != 0 {
					root = 32768
				}
				root >>= zeros >> 1
				root += int32(int64(root) * int64(int16(213*fracQ7)) >> 16)
			}
			gain := int32(uint32(root) << 4)
			slope := int32(uint32(((1<<16)-gain)/length) << 2)
			samples := unsafe.Slice(frame, length)
			for i := range samples {
				samples[i] = int16(int64(gain) * int64(samples[i]) >> 16)
				gain += slope
				if gain > 1<<16 {
					break
				}
			}
		}
	}
	plc.Flast_frame_lost = 0
}

const silk_int16_MAX8 = 0x7FFF

/* shell coder; pulse-subframe length is hardcoded */
