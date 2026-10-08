// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

package opuscc

import (
	"reflect"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

var _ reflect.Type

//go:uintptrescapes
func Opus_silk_decode_frame(tls *libc.TLS, psDec, psRangeDec, pOut, pN uintptr, lostFlag, condCoding, arch int32) int32 {
	return silk_decode_frame(tls, (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)), (*OpusT_ec_dec)(unsafe.Pointer(psRangeDec)), (*int16)(unsafe.Pointer(pOut)), (*int32)(unsafe.Pointer(pN)), lostFlag, condCoding, arch)
}

func silkDecodeFramePulses(length int32) []int16 {
	return make([]int16, (length+SHELL_CODEC_FRAME_LENGTH-1)&^(SHELL_CODEC_FRAME_LENGTH-1))
}

func silkDecodeFrameFinish(decoder *OpusT_silk_decoder_state, control *OpusT_silk_decoder_control, count *int32, length int32) {
	decoder.FlagPrev = control.FpitchL[decoder.Fnb_subfr-1]
	*count = length
}

func silkDecodeFrameHistory(decoder *OpusT_silk_decoder_state, frame *int16) {
	move := decoder.Fltp_mem_length - decoder.Fframe_length
	copy(decoder.FoutBuf[:move], decoder.FoutBuf[decoder.Fframe_length:decoder.Fframe_length+move])
	copy(decoder.FoutBuf[move:move+decoder.Fframe_length], unsafe.Slice(frame, decoder.Fframe_length))
}

func silk_decode_frame(tls *libc.TLS, psDec *OpusT_silk_decoder_state, psRangeDec *OpusT_ec_dec, pOut *int16, pN *int32, lostFlag int32, condCoding int32, arch int32) (r int32) {
	var L, ret int32
	var psDecCtrl *OpusT_silk_decoder_control
	var pulses []int16
	decoder := psDec
	ret = 0
	L = decoder.Fframe_length
	psDecCtrl = new(OpusT_silk_decoder_control)
	control := psDecCtrl
	psDecCtrl.FLTP_scale_Q14 = 0
	/* Safety checks */
	if !(L > 0 && L <= int32(SUB_FRAME_LENGTH_MS)*int32(MAX_NB_SUBFR)*int32(MAX_FS_KHZ)) {
		Opus_celt_fatal(tls, __ccgo_ts+5921, __ccgo_ts+5898, int32(68))
	}
	decode := lostFlag == FLAG_DECODE_NORMAL || lostFlag == FLAG_DECODE_LBRR && decoder.FLBRR_flags[decoder.FnFramesDecoded] == 1
	if decode {
		pulses = silkDecodeFramePulses(L)
		/*********************************************/
		/* Decode quantization indices of side info  */
		/*********************************************/
		Opus_silk_decode_indices(tls, psDec, psRangeDec, psDec.FnFramesDecoded, lostFlag, condCoding)
		/*********************************************/
		/* Decode quantization indices of excitation */
		/*********************************************/
		Opus_silk_decode_pulses(tls, psRangeDec, unsafe.SliceData(pulses), int32(psDec.Findices.FsignalType), int32(psDec.Findices.FquantOffsetType), psDec.Fframe_length)
		/********************************************/
		/* Decode parameters and pulse signal       */
		/********************************************/
		Opus_silk_decode_parameters(tls, psDec, psDecCtrl, condCoding)
		/********************************************************/
		/* Run inverse NSQ                                      */
		/********************************************************/
		silk_decode_core(tls, decoder, psDecCtrl, pOut, unsafe.SliceData(pulses), arch)
		/*************************/
		/* Update output buffer. */
		/*************************/
		if !(psDec.Fltp_mem_length >= psDec.Fframe_length) {
			Opus_celt_fatal(tls, __ccgo_ts+5970, __ccgo_ts+5898, int32(104))
		}
		silkDecodeFrameHistory(decoder, pOut)
		/********************************************************/
		/* Update PLC state                                     */
		/********************************************************/
		silk_PLC(tls, decoder, psDecCtrl, pOut, 0, arch)
		psDec.FlossCnt = 0
		psDec.FprevSignalType = int32(psDec.Findices.FsignalType)
		if !(psDec.FprevSignalType >= 0 && psDec.FprevSignalType <= int32(2)) {
			Opus_celt_fatal(tls, __ccgo_ts+6033, __ccgo_ts+5898, int32(127))
		}
		/* A frame has been decoded without errors */
		psDec.Ffirst_frame_after_reset = 0
	} else {
		/* Handle packet loss by extrapolation */
		silk_PLC(tls, decoder, psDecCtrl, pOut, 1, arch)
		/*************************/
		/* Update output buffer. */
		/*************************/
		if !(psDec.Fltp_mem_length >= psDec.Fframe_length) {
			Opus_celt_fatal(tls, __ccgo_ts+5970, __ccgo_ts+5898, int32(145))
		}
		silkDecodeFrameHistory(decoder, pOut)
	}
	/************************************************/
	/* Comfort noise generation / estimation        */
	/************************************************/
	Opus_silk_CNG(tls, decoder, psDecCtrl, pOut, L)
	/****************************************************************/
	/* Ensure smooth connection of extrapolated and good frames     */
	/****************************************************************/
	Opus_silk_PLC_glue_frames(tls, decoder, pOut, L)
	// Lag update precedes the final live count store.
	silkDecodeFrameFinish(decoder, control, pN, L)
	return ret
}

// C documentation
//
//	/* Decode parameters from payload */
func Opus_silk_decode_parameters(tls *libc.TLS, decoder *OpusT_silk_decoder_state, control *OpusT_silk_decoder_control, condCoding int32) {
	var nlsf, nlsf0 [MAX_LPC_ORDER]int16
	Opus_silk_gains_dequant(tls, &control.FGains_Q16[0], &decoder.Findices.FGainsIndices[0], &decoder.FLastGainIndex, libc.BoolInt32(condCoding == CODE_CONDITIONALLY), decoder.Fnb_subfr)
	Opus_silk_NLSF_decode(tls, &nlsf[0], &decoder.Findices.FNLSFIndices[0], decoder.FpsNLSF_CB)
	Opus_silk_NLSF2A(tls, &control.FPredCoef_Q12[1][0], &nlsf[0], decoder.FLPC_order, decoder.Farch)
	if decoder.Ffirst_frame_after_reset == 1 {
		decoder.Findices.FNLSFInterpCoef_Q2 = 4
	}
	if decoder.Findices.FNLSFInterpCoef_Q2 < 4 {
		for i := int32(0); i < decoder.FLPC_order; i++ {
			nlsf0[i] = int16(int32(decoder.FprevNLSF_Q15[i]) + (int32(decoder.Findices.FNLSFInterpCoef_Q2) * (int32(nlsf[i]) - int32(decoder.FprevNLSF_Q15[i])) >> 2))
		}
		Opus_silk_NLSF2A(tls, &control.FPredCoef_Q12[0][0], &nlsf0[0], decoder.FLPC_order, decoder.Farch)
	} else {
		copy(control.FPredCoef_Q12[0][:decoder.FLPC_order], control.FPredCoef_Q12[1][:decoder.FLPC_order])
	}
	copy(decoder.FprevNLSF_Q15[:decoder.FLPC_order], nlsf[:decoder.FLPC_order])
	if decoder.FlossCnt != 0 {
		Opus_silk_bwexpander(tls, &control.FPredCoef_Q12[0][0], decoder.FLPC_order, BWE_AFTER_LOSS_Q16)
		Opus_silk_bwexpander(tls, &control.FPredCoef_Q12[1][0], decoder.FLPC_order, BWE_AFTER_LOSS_Q16)
	}
	if decoder.Findices.FsignalType == TYPE_VOICED {
		Opus_silk_decode_pitch(tls, decoder.Findices.FlagIndex, decoder.Findices.FcontourIndex, &control.FpitchL[0], decoder.Ffs_kHz, decoder.Fnb_subfr)
		periodicity := decoder.Findices.FPERIndex
		table := unsafe.Slice(Opus_silk_LTP_vq_ptrs_Q7[periodicity], 8<<uint(periodicity))
		for k := int32(0); k < decoder.Fnb_subfr; k++ {
			row := table[decoder.Findices.FLTPIndex[k]]
			for i := int32(0); i < LTP_ORDER; i++ {
				control.FLTPCoef_Q14[k*LTP_ORDER+i] = int16(int32(row[i]) << 7)
			}
		}
		control.FLTP_scale_Q14 = int32(Opus_silk_LTPScales_table_Q14[decoder.Findices.FLTP_scaleIndex])
	} else {
		clear(control.FpitchL[:decoder.Fnb_subfr])
		clear(control.FLTPCoef_Q14[:LTP_ORDER*decoder.Fnb_subfr])
		decoder.Findices.FPERIndex = 0
		control.FLTP_scale_Q14 = 0
	}
}

// ICDF callers consume the selected entry; do not expose a guessed full row.
func silkDecodeICDFAt(table *byte, offset int32) *byte {
	if offset == 0 {
		return table
	}
	if offset < 0 {
		return (*byte)(unsafe.Add(unsafe.Pointer(table), offset))
	}
	return &unsafe.Slice(table, int64(offset)+1)[offset]
}

// C documentation
//
//	/* Decode side-information parameters from payload */
func Opus_silk_decode_indices(tls *libc.TLS, decoder *OpusT_silk_decoder_state, dec *OpusT_ec_dec, frame, lbrr, cond int32) {
	indices := &decoder.Findices
	var ix int32
	if lbrr != 0 || decoder.FVAD_flags[frame] != 0 {
		ix = Opus_ec_dec_icdf(tls, dec, &Opus_silk_type_offset_VAD_iCDF[0], 8) + 2
	} else {
		ix = Opus_ec_dec_icdf(tls, dec, &Opus_silk_type_offset_no_VAD_iCDF[0], 8)
	}
	indices.FsignalType = int8(ix >> 1)
	indices.FquantOffsetType = int8(ix & 1)
	if cond == CODE_CONDITIONALLY {
		indices.FGainsIndices[0] = int8(Opus_ec_dec_icdf(tls, dec, &Opus_silk_delta_gain_iCDF[0], 8))
	} else {
		indices.FGainsIndices[0] = int8(Opus_ec_dec_icdf(tls, dec, &Opus_silk_gain_iCDF[indices.FsignalType][0], 8) << 3)
		indices.FGainsIndices[0] += int8(Opus_ec_dec_icdf(tls, dec, &Opus_silk_uniform8_iCDF[0], 8))
	}
	for i := int32(1); i < decoder.Fnb_subfr; i++ {
		indices.FGainsIndices[i] = int8(Opus_ec_dec_icdf(tls, dec, &Opus_silk_delta_gain_iCDF[0], 8))
	}
	// Retain concrete codebook/table bases while selecting consumed rows.
	cb := decoder.FpsNLSF_CB
	first := cb.FCB1_iCDF
	residual := cb.Fec_iCDF
	indices.FNLSFIndices[0] = int8(Opus_ec_dec_icdf(tls, dec, silkDecodeICDFAt(first, int32(indices.FsignalType>>1)*int32(cb.FnVectors)), 8))
	var ecIX [MAX_LPC_ORDER]int16
	var pred [MAX_LPC_ORDER]byte
	Opus_silk_NLSF_unpack(tls, &ecIX[0], &pred[0], cb, int32(indices.FNLSFIndices[0]))
	if int32(cb.Forder) != decoder.FLPC_order {
		Opus_celt_fatal(tls, __ccgo_ts+6108, __ccgo_ts+6170, 82)
	}
	for i := int32(0); i < int32(cb.Forder); i++ {
		ix = Opus_ec_dec_icdf(tls, dec, silkDecodeICDFAt(residual, int32(ecIX[i])), 8)
		if ix == 0 {
			ix -= Opus_ec_dec_icdf(tls, dec, &Opus_silk_NLSF_EXT_iCDF[0], 8)
		} else if ix == 2*NLSF_QUANT_MAX_AMPLITUDE {
			ix += Opus_ec_dec_icdf(tls, dec, &Opus_silk_NLSF_EXT_iCDF[0], 8)
		}
		indices.FNLSFIndices[i+1] = int8(ix - NLSF_QUANT_MAX_AMPLITUDE)
	}
	if decoder.Fnb_subfr == MAX_NB_SUBFR {
		indices.FNLSFInterpCoef_Q2 = int8(Opus_ec_dec_icdf(tls, dec, &Opus_silk_NLSF_interpolation_factor_iCDF[0], 8))
	} else {
		indices.FNLSFInterpCoef_Q2 = 4
	}
	if indices.FsignalType == TYPE_VOICED {
		absolute := true
		if cond == CODE_CONDITIONALLY && decoder.Fec_prevSignalType == TYPE_VOICED {
			delta := int32(int16(Opus_ec_dec_icdf(tls, dec, &Opus_silk_pitch_delta_iCDF[0], 8)))
			if delta > 0 {
				indices.FlagIndex = int16(int32(decoder.Fec_prevLagIndex) + delta - 9)
				absolute = false
			}
		}
		if absolute {
			indices.FlagIndex = int16(int32(int16(Opus_ec_dec_icdf(tls, dec, &Opus_silk_pitch_lag_iCDF[0], 8))) * (decoder.Ffs_kHz >> 1))
			indices.FlagIndex += int16(Opus_ec_dec_icdf(tls, dec, decoder.Fpitch_lag_low_bits_iCDF, 8))
		}
		decoder.Fec_prevLagIndex = indices.FlagIndex
		indices.FcontourIndex = int8(Opus_ec_dec_icdf(tls, dec, decoder.Fpitch_contour_iCDF, 8))
		indices.FPERIndex = int8(Opus_ec_dec_icdf(tls, dec, &Opus_silk_LTP_per_index_iCDF[0], 8))
		for k := int32(0); k < decoder.Fnb_subfr; k++ {
			indices.FLTPIndex[k] = int8(Opus_ec_dec_icdf(tls, dec, Opus_silk_LTP_gain_iCDF_ptrs[indices.FPERIndex], 8))
		}
		if cond == CODE_INDEPENDENTLY {
			indices.FLTP_scaleIndex = int8(Opus_ec_dec_icdf(tls, dec, &Opus_silk_LTPscale_iCDF[0], 8))
		} else {
			indices.FLTP_scaleIndex = 0
		}
	}
	decoder.Fec_prevSignalType = int32(indices.FsignalType)
	indices.FSeed = int8(Opus_ec_dec_icdf(tls, dec, &Opus_silk_uniform4_iCDF[0], 8))
}

// C documentation
//
//	/*********************************************/
//	/* Decode quantization indices of excitation */
//	/*********************************************/
func Opus_silk_decode_pulses(tls *libc.TLS, dec *OpusT_ec_dec, pulses *OpusT_opus_int16, signalType, quantOffsetType, frame_length int32) {
	var shifts, sums [20]int32 // MAX_NB_SHELL_BLOCKS: 320 samples / 16.
	rate := ec_dec_icdf(tls, dec, &Opus_silk_rate_levels_iCDF[signalType>>1][0], 8)
	blocks := frame_length >> LOG2_SHELL_CODEC_FRAME_LENGTH
	if blocks*SHELL_CODEC_FRAME_LENGTH < frame_length {
		if frame_length != 120 {
			Opus_celt_fatal(tls, __ccgo_ts+6195, __ccgo_ts+6237, 59)
		}
		blocks++ // 10 ms at 12 kHz has a padded final shell block.
	}
	q := unsafe.Slice(pulses, blocks*SHELL_CODEC_FRAME_LENGTH)
	for i := int32(0); i < blocks; i++ {
		sums[i] = ec_dec_icdf(tls, dec, &Opus_silk_pulses_per_block_iCDF[rate][0], 8)
		for sums[i] == SILK_MAX_PULSES+1 {
			shifts[i]++
			start := 0
			if shifts[i] == 10 {
				start = 1
			} // Disallow another escape after ten LSBs.
			sums[i] = ec_dec_icdf(tls, dec, &Opus_silk_pulses_per_block_iCDF[N_RATE_LEVELS-1][start], 8)
		}
	}
	for i := int32(0); i < blocks; i++ {
		block := q[i*SHELL_CODEC_FRAME_LENGTH : (i+1)*SHELL_CODEC_FRAME_LENGTH]
		if sums[i] > 0 {
			Opus_silk_shell_decoder(tls, (*[16]OpusT_opus_int16)(block), dec, sums[i])
		} else {
			clear(block)
		}
	}
	for i := int32(0); i < blocks; i++ {
		if shifts[i] > 0 {
			block := q[i*SHELL_CODEC_FRAME_LENGTH : (i+1)*SHELL_CODEC_FRAME_LENGTH]
			for k, pulse := range block {
				value := int32(pulse)
				for j := int32(0); j < shifts[i]; j++ {
					value = int32(uint32(value)<<1) + ec_dec_icdf(tls, dec, &Opus_silk_lsb_iCDF[0], 8)
				}
				block[k] = int16(value)
			}
			sums[i] |= shifts[i] << 5 // Keep sign decoding active after LSB reconstruction.
		}
	}
	Opus_silk_decode_signs(tls, dec, pulses, frame_length, signalType, quantOffsetType, &sums[0])
}

// C documentation
//
//	/* Set decoder sampling rate */
