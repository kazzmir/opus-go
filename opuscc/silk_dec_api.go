// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

package opuscc

import (
	"reflect"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

var _ reflect.Type
var _ unsafe.Pointer

func Opus_silk_decoder_set_fs(tls *libc.TLS, psDec uintptr, fs_kHz int32, fs_API_Hz OpusT_opus_int32) (r int32) {
	var frame_length, ret int32
	_, _ = frame_length, ret
	ret = 0
	if !(fs_kHz == int32(8) || fs_kHz == int32(12) || fs_kHz == int32(16)) {
		Opus_celt_fatal(tls, __ccgo_ts+6261, __ccgo_ts+6323, int32(43))
	}
	if !((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr == int32(MAX_NB_SUBFR) || (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr == int32(MAX_NB_SUBFR)/int32(2)) {
		Opus_celt_fatal(tls, __ccgo_ts+6348, __ccgo_ts+6323, int32(44))
	}
	/* New (sub)frame length */
	(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length = int32(int16(int32(SUB_FRAME_LENGTH_MS))) * int32(int16(fs_kHz))
	frame_length = int32(int16((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr)) * int32(int16((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fsubfr_length))
	/* Initialize resampler when switching internal or external sampling frequency */
	if (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Ffs_kHz != fs_kHz || (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Ffs_API_hz != fs_API_Hz {
		/* Initialize the resampler for dec_API.c preparing resampling from fs_kHz to API_fs_Hz */
		ret = ret + Opus_silk_resampler_init(tls, psDec+2448, int32(int16(fs_kHz))*int32(int16(int32(1000))), fs_API_Hz, 0)
		(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Ffs_API_hz = fs_API_Hz
	}
	if (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Ffs_kHz != fs_kHz || frame_length != (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length {
		if fs_kHz == int32(8) {
			if (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr == int32(MAX_NB_SUBFR) {
				(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fpitch_contour_iCDF = uintptr(unsafe.Pointer(&Opus_silk_pitch_contour_NB_iCDF))
			} else {
				(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fpitch_contour_iCDF = uintptr(unsafe.Pointer(&Opus_silk_pitch_contour_10_ms_NB_iCDF))
			}
		} else {
			if (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fnb_subfr == int32(MAX_NB_SUBFR) {
				(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fpitch_contour_iCDF = uintptr(unsafe.Pointer(&Opus_silk_pitch_contour_iCDF))
			} else {
				(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fpitch_contour_iCDF = uintptr(unsafe.Pointer(&Opus_silk_pitch_contour_10_ms_iCDF))
			}
		}
		if (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Ffs_kHz != fs_kHz {
			(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fltp_mem_length = int32(int16(int32(LTP_MEM_LENGTH_MS))) * int32(int16(fs_kHz))
			if fs_kHz == int32(8) || fs_kHz == int32(12) {
				(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order = int32(MIN_LPC_ORDER)
				(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FpsNLSF_CB = uintptr(unsafe.Pointer(&Opus_silk_NLSF_CB_NB_MB))
			} else {
				(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLPC_order = int32(MAX_LPC_ORDER)
				(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FpsNLSF_CB = uintptr(unsafe.Pointer(&Opus_silk_NLSF_CB_WB))
			}
			if fs_kHz == int32(16) {
				(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fpitch_lag_low_bits_iCDF = uintptr(unsafe.Pointer(&Opus_silk_uniform8_iCDF))
			} else {
				if fs_kHz == int32(12) {
					(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fpitch_lag_low_bits_iCDF = uintptr(unsafe.Pointer(&Opus_silk_uniform6_iCDF))
				} else {
					if fs_kHz == int32(8) {
						(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fpitch_lag_low_bits_iCDF = uintptr(unsafe.Pointer(&Opus_silk_uniform4_iCDF))
					} else {
						/* unsupported sampling rate */
						if !(int32(0) != 0) {
							Opus_celt_fatal(tls, __ccgo_ts+1017, __ccgo_ts+6323, int32(89))
						}
					}
				}
			}
			(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Ffirst_frame_after_reset = int32(1)
			(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FlagPrev = int32(100)
			(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FLastGainIndex = int8(10)
			(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).FprevSignalType = TYPE_NO_VOICE_ACTIVITY
			libc.Xmemset(tls, psDec+1348, 0, uint64(960))
			libc.Xmemset(tls, psDec+1284, 0, uint64(64))
		}
		(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Ffs_kHz = fs_kHz
		(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length = frame_length
	}
	/* Check that settings are valid */
	if !((*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length > 0 && (*OpusT_silk_decoder_state)(unsafe.Pointer(psDec)).Fframe_length <= int32(SUB_FRAME_LENGTH_MS)*int32(MAX_NB_SUBFR)*int32(MAX_FS_KHZ)) {
		Opus_celt_fatal(tls, __ccgo_ts+6435, __ccgo_ts+6323, int32(104))
	}
	return ret
}

const CELT_SIG_SCALE9 = 32768

/* Copyright (C) 2007 Jean-Marc Valin

   File: os_support.h
   This is the (tiny) OS abstraction layer. Aside from math.h, this is the
   only place where system headers are allowed.

   Redistribution and use in source and binary forms, with or without
   modification, are permitted provided that the following conditions are
   met:

   1. Redistributions of source code must retain the above copyright notice,
   this list of conditions and the following disclaimer.

   2. Redistributions in binary form must reproduce the above copyright
   notice, this list of conditions and the following disclaimer in the
   documentation and/or other materials provided with the distribution.

   THIS SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR
   IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
   OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
   DISCLAIMED. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT,
   INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
   (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
   SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
   HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
   STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN
   ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
   POSSIBILITY OF SUCH DAMAGE.
*/

// C documentation
//
//	/************************/
//	/* Decoder Super Struct */
//	/************************/
type OpusT_silk_decoder = struct {
	Fchannel_state           [2]OpusT_silk_decoder_state
	FsStereo                 OpusT_stereo_dec_state
	FnChannelsAPI            int32
	FnChannelsInternal       int32
	Fprev_decode_only_middle int32
}

/*********************/
/* Decoder functions */
/*********************/
func Opus_silk_LoadOSCEModels(tls *libc.TLS, decState uintptr, data uintptr, len1 int32) (r int32) {
	_ = decState
	_ = data
	_ = len1
	return SILK_NO_ERROR
}

func Opus_silk_Get_Decoder_Size(tls *libc.TLS, decSizeBytes uintptr) (r int32) {
	var ret int32
	_ = ret
	ret = SILK_NO_ERROR
	*(*int32)(unsafe.Pointer(decSizeBytes)) = int32(8808)
	return ret
}

// C documentation
//
//	/* Reset decoder state */
func Opus_silk_ResetDecoder(tls *libc.TLS, decState uintptr) (r int32) {
	var channel_state uintptr
	var n, ret int32
	_, _, _ = channel_state, n, ret
	ret = SILK_NO_ERROR
	channel_state = decState
	n = 0
	for {
		if !(n < int32(DECODER_NUM_CHANNELS)) {
			break
		}
		ret = Opus_silk_reset_decoder(tls, channel_state+uintptr(n)*4392)
		goto _1
	_1:
		;
		n = n + 1
	}
	libc.Xmemset(tls, decState+8784, 0, uint64(12))
	/* Not strictly needed, but it's cleaner that way */
	(*OpusT_silk_decoder)(unsafe.Pointer(decState)).Fprev_decode_only_middle = 0
	return ret
}

func Opus_silk_InitDecoder(tls *libc.TLS, decState uintptr) (r int32) {
	var channel_state uintptr
	var n, ret int32
	_, _, _ = channel_state, n, ret
	ret = SILK_NO_ERROR
	channel_state = decState
	/* load osce models */
	Opus_silk_LoadOSCEModels(tls, decState, uintptr(uint32(0)), 0)
	n = 0
	for {
		if !(n < int32(DECODER_NUM_CHANNELS)) {
			break
		}
		ret = Opus_silk_init_decoder(tls, channel_state+uintptr(n)*4392)
		goto _1
	_1:
		;
		n = n + 1
	}
	libc.Xmemset(tls, decState+8784, 0, uint64(12))
	/* Not strictly needed, but it's cleaner that way */
	(*OpusT_silk_decoder)(unsafe.Pointer(decState)).Fprev_decode_only_middle = 0
	return ret
}

// C documentation
//
//	/* Decode a frame */
func Opus_silk_Decode(tls *libc.TLS, decState uintptr, decControl uintptr, lostFlag int32, newPacketFlag int32, psRangeDec uintptr, samplesOut uintptr, nSamplesOut uintptr, arch int32) (r int32) {
	bp := tls.Alloc(656)
	defer tls.Free(656)
	var FrameIndex, condCoding, condCoding1, fs_kHz_dec, has_side, i, n, ret, stereo_to_mono, v51 int32
	var LBRR_symbol OpusT_opus_int32
	var _saved_stack, channel_state, psDec, resample_out_ptr, samplesOut1_tmp_storage1, samplesOut2_tmp, st, v1, v11, v13, v15, v17, v26, v28, v3, v30, v32, v7, v9 uintptr
	var mult_tab [3]int32
	var samplesOut1_tmp [2]uintptr
	var _ /* MS_pred_Q13 at bp+8 */ [2]OpusT_opus_int32
	var _ /* decode_only_middle at bp+0 */ int32
	var _ /* nSamplesOutDec at bp+4 */ OpusT_opus_int32
	var _ /* pulses at bp+16 */ [320]OpusT_opus_int16
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = FrameIndex, LBRR_symbol, _saved_stack, channel_state, condCoding, condCoding1, fs_kHz_dec, has_side, i, mult_tab, n, psDec, resample_out_ptr, ret, samplesOut1_tmp, samplesOut1_tmp_storage1, samplesOut2_tmp, st, stereo_to_mono, v1, v11, v13, v15, v17, v26, v28, v3, v30, v32, v51, v7, v9
	*(*int32)(unsafe.Pointer(bp)) = 0
	ret = SILK_NO_ERROR
	*(*[2]OpusT_opus_int32)(unsafe.Pointer(bp + 8)) = [2]OpusT_opus_int32{}
	psDec = decState
	channel_state = psDec
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
	if !((*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal == int32(1) || (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal == int32(2)) {
		Opus_celt_fatal(tls, __ccgo_ts+6520, __ccgo_ts+6611, int32(165))
	}
	/**********************************/
	/* Test if first frame in payload */
	/**********************************/
	if newPacketFlag != 0 {
		n = 0
		for {
			if !(n < (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal) {
				break
			}
			(*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + uintptr(n)*4392))).FnFramesDecoded = 0 /* Used to count frames in packet */
			goto _5
		_5:
			;
			n = n + 1
		}
	}
	/* If Mono -> Stereo transition in bitstream: init state of second channel */
	if (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal > (*OpusT_silk_decoder)(unsafe.Pointer(psDec)).FnChannelsInternal {
		ret = ret + Opus_silk_init_decoder(tls, channel_state+1*4392)
	}
	stereo_to_mono = libc.BoolInt32((*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal == int32(1) && (*OpusT_silk_decoder)(unsafe.Pointer(psDec)).FnChannelsInternal == int32(2) && (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FinternalSampleRate == int32(1000)*(*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state))).Ffs_kHz)
	if (*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state))).FnFramesDecoded == 0 {
		n = 0
		for {
			if !(n < (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal) {
				break
			}
			if (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FpayloadSize_ms == 0 {
				/* Assuming packet loss, use 10 ms */
				(*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + uintptr(n)*4392))).FnFramesPerPacket = int32(1)
				(*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + uintptr(n)*4392))).Fnb_subfr = int32(2)
			} else {
				if (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FpayloadSize_ms == int32(10) {
					(*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + uintptr(n)*4392))).FnFramesPerPacket = int32(1)
					(*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + uintptr(n)*4392))).Fnb_subfr = int32(2)
				} else {
					if (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FpayloadSize_ms == int32(20) {
						(*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + uintptr(n)*4392))).FnFramesPerPacket = int32(1)
						(*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + uintptr(n)*4392))).Fnb_subfr = int32(4)
					} else {
						if (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FpayloadSize_ms == int32(40) {
							(*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + uintptr(n)*4392))).FnFramesPerPacket = int32(2)
							(*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + uintptr(n)*4392))).Fnb_subfr = int32(4)
						} else {
							if (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FpayloadSize_ms == int32(60) {
								(*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + uintptr(n)*4392))).FnFramesPerPacket = int32(3)
								(*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + uintptr(n)*4392))).Fnb_subfr = int32(4)
							} else {
								if !(int32(0) != 0) {
									Opus_celt_fatal(tls, __ccgo_ts+1017, __ccgo_ts+6611, int32(204))
								}
								st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
								if !(st != 0) {
									v1 = libc.Xmalloc(tls, uint64(16))
									goto _8
								_8:
									st = v1
									if st != 0 {
										libc.Xmemset(tls, st, 0, uint64(16))
									}
									libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
								}
								v3 = st
								goto _10
							_10:
								(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
								return -int32(203)
							}
						}
					}
				}
			}
			fs_kHz_dec = (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FinternalSampleRate>>int32(10) + int32(1)
			if fs_kHz_dec != int32(8) && fs_kHz_dec != int32(12) && fs_kHz_dec != int32(16) {
				if !(int32(0) != 0) {
					Opus_celt_fatal(tls, __ccgo_ts+1017, __ccgo_ts+6611, int32(210))
				}
				st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
				if !(st != 0) {
					v1 = libc.Xmalloc(tls, uint64(16))
					goto _12
				_12:
					st = v1
					if st != 0 {
						libc.Xmemset(tls, st, 0, uint64(16))
					}
					libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
				}
				v3 = st
				goto _14
			_14:
				(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
				return -int32(200)
			}
			ret = ret + Opus_silk_decoder_set_fs(tls, channel_state+uintptr(n)*4392, fs_kHz_dec, (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FAPI_sampleRate)
			goto _6
		_6:
			;
			n = n + 1
		}
	}
	if (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsAPI == int32(2) && (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal == int32(2) && ((*OpusT_silk_decoder)(unsafe.Pointer(psDec)).FnChannelsAPI == int32(1) || (*OpusT_silk_decoder)(unsafe.Pointer(psDec)).FnChannelsInternal == int32(1)) {
		libc.Xmemset(tls, psDec+8784, 0, uint64(4))
		libc.Xmemset(tls, psDec+8784+8, 0, uint64(4))
		libc.Xmemcpy(tls, channel_state+1*4392+2448, channel_state+2448, uint64(400))
	}
	(*OpusT_silk_decoder)(unsafe.Pointer(psDec)).FnChannelsAPI = (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsAPI
	(*OpusT_silk_decoder)(unsafe.Pointer(psDec)).FnChannelsInternal = (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal
	if (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FAPI_sampleRate > int32(MAX_API_FS_KHZ)*int32(1000) || (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FAPI_sampleRate < int32(8000) {
		ret = -int32(200)
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			goto _16
		_16:
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v3 = st
		goto _18
	_18:
		(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
		return ret
	}
	if lostFlag != int32(FLAG_PACKET_LOST) && (*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state))).FnFramesDecoded == 0 {
		/* First decoder call for this payload */
		/* Decode VAD flags and LBRR flag */
		n = 0
		for {
			if !(n < (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal) {
				break
			}
			i = 0
			for {
				if !(i < (*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + uintptr(n)*4392))).FnFramesPerPacket) {
					break
				}
				*(*int32)(unsafe.Pointer(channel_state + uintptr(n)*4392 + 2416 + uintptr(i)*4)) = Opus_ec_dec_bit_logp(tls, psRangeDec, uint32(1))
				goto _20
			_20:
				;
				i = i + 1
			}
			(*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + uintptr(n)*4392))).FLBRR_flag = Opus_ec_dec_bit_logp(tls, psRangeDec, uint32(1))
			goto _19
		_19:
			;
			n = n + 1
		}
		/* Decode LBRR flags */
		n = 0
		for {
			if !(n < (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal) {
				break
			}
			libc.Xmemset(tls, channel_state+uintptr(n)*4392+2432, 0, uint64(12))
			if (*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + uintptr(n)*4392))).FLBRR_flag != 0 {
				if (*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + uintptr(n)*4392))).FnFramesPerPacket == int32(1) {
					*(*int32)(unsafe.Pointer(channel_state + uintptr(n)*4392 + 2432)) = int32(1)
				} else {
					LBRR_symbol = Opus_ec_dec_icdf(tls, psRangeDec, Opus_silk_LBRR_flags_iCDF_ptr[(*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + uintptr(n)*4392))).FnFramesPerPacket-int32(2)], uint32(8)) + int32(1)
					i = 0
					for {
						if !(i < (*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + uintptr(n)*4392))).FnFramesPerPacket) {
							break
						}
						*(*int32)(unsafe.Pointer(channel_state + uintptr(n)*4392 + 2432 + uintptr(i)*4)) = LBRR_symbol >> i & int32(1)
						goto _22
					_22:
						;
						i = i + 1
					}
				}
			}
			goto _21
		_21:
			;
			n = n + 1
		}
		if lostFlag == FLAG_DECODE_NORMAL {
			/* Regular decoding: skip all LBRR data */
			i = 0
			for {
				if !(i < (*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state))).FnFramesPerPacket) {
					break
				}
				n = 0
				for {
					if !(n < (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal) {
						break
					}
					if *(*int32)(unsafe.Pointer(channel_state + uintptr(n)*4392 + 2432 + uintptr(i)*4)) != 0 {
						if (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal == int32(2) && n == 0 {
							Opus_silk_stereo_decode_pred(tls, psRangeDec, bp+8)
							if *(*int32)(unsafe.Pointer(channel_state + 1*4392 + 2432 + uintptr(i)*4)) == 0 {
								Opus_silk_stereo_decode_mid_only(tls, psRangeDec, bp)
							}
						}
						/* Use conditional coding if previous frame available */
						if i > 0 && *(*int32)(unsafe.Pointer(channel_state + uintptr(n)*4392 + 2432 + uintptr(i-int32(1))*4)) != 0 {
							condCoding = int32(CODE_CONDITIONALLY)
						} else {
							condCoding = CODE_INDEPENDENTLY
						}
						Opus_silk_decode_indices(tls, channel_state+uintptr(n)*4392, psRangeDec, i, int32(1), condCoding)
						Opus_silk_decode_pulses(tls, psRangeDec, bp+16, int32((*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + uintptr(n)*4392))).Findices.FsignalType), int32((*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + uintptr(n)*4392))).Findices.FquantOffsetType), (*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + uintptr(n)*4392))).Fframe_length)
					}
					goto _24
				_24:
					;
					n = n + 1
				}
				goto _23
			_23:
				;
				i = i + 1
			}
		}
	}
	/* Get MS predictor index */
	if (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal == int32(2) {
		if lostFlag == FLAG_DECODE_NORMAL || lostFlag == int32(FLAG_DECODE_LBRR) && *(*int32)(unsafe.Pointer(channel_state + 2432 + uintptr((*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state))).FnFramesDecoded)*4)) == int32(1) {
			Opus_silk_stereo_decode_pred(tls, psRangeDec, bp+8)
			/* For LBRR data, decode mid-only flag only if side-channel's LBRR flag is false */
			if lostFlag == FLAG_DECODE_NORMAL && *(*int32)(unsafe.Pointer(channel_state + 1*4392 + 2416 + uintptr((*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state))).FnFramesDecoded)*4)) == 0 || lostFlag == int32(FLAG_DECODE_LBRR) && *(*int32)(unsafe.Pointer(channel_state + 1*4392 + 2432 + uintptr((*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state))).FnFramesDecoded)*4)) == 0 {
				Opus_silk_stereo_decode_mid_only(tls, psRangeDec, bp)
			} else {
				*(*int32)(unsafe.Pointer(bp)) = 0
			}
		} else {
			n = 0
			for {
				if !(n < int32(2)) {
					break
				}
				(*(*[2]OpusT_opus_int32)(unsafe.Pointer(bp + 8)))[n] = int32(*(*OpusT_opus_int16)(unsafe.Pointer(psDec + 8784 + uintptr(n)*2)))
				goto _25
			_25:
				;
				n = n + 1
			}
		}
	}
	/* Reset side channel decoder prediction memory for first frame with side coding */
	if (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal == int32(2) && *(*int32)(unsafe.Pointer(bp)) == 0 && (*OpusT_silk_decoder)(unsafe.Pointer(psDec)).Fprev_decode_only_middle == int32(1) {
		libc.Xmemset(tls, psDec+1*4392+1348, 0, uint64(960))
		libc.Xmemset(tls, psDec+1*4392+1284, 0, uint64(64))
		(*(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec + 1*4392))).FlagPrev = int32(100)
		(*(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec + 1*4392))).FLastGainIndex = int8(10)
		(*(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec + 1*4392))).FprevSignalType = TYPE_NO_VOICE_ACTIVITY
		(*(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec + 1*4392))).Ffirst_frame_after_reset = int32(1)
	}
	/* Check if the temp buffer fits into the output PCM buffer. If it fits,
	   we can delay allocating the temp buffer until after the SILK peak stack
	   usage. We need to use a < and not a <= because of the two extra samples. */
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _27
	_27:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _29
_29:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v7 = libc.Xmalloc(tls, uint64(16))
		goto _31
	_31:
		st = v7
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v9 = st
	goto _33
_33:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(2)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v9)).Fglobal_stack))) & (uint64(uint32(2)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _35
	_35:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _37
_37:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _39
	_39:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _41
_41:
	if !(int64(int32(uint64(uint32((*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal*((*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state))).Fframe_length+int32(2))))*(uint64(2)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+6611, int32(319))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v26 = libc.Xmalloc(tls, uint64(16))
		goto _43
	_43:
		st = v26
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v28 = st
	goto _45
_45:
	*(*uintptr)(unsafe.Pointer(v28 + 8)) += uintptr(uint64(uint32((*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal*((*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state))).Fframe_length+int32(2)))) * (uint64(2) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v30 = libc.Xmalloc(tls, uint64(16))
		goto _47
	_47:
		st = v30
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v32 = st
	goto _49
_49:
	samplesOut1_tmp_storage1 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v32)).Fglobal_stack - uintptr(uint64(uint32((*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal*((*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state))).Fframe_length+int32(2))))*(uint64(2)/uint64(1)))
	samplesOut1_tmp[0] = samplesOut1_tmp_storage1
	samplesOut1_tmp[int32(1)] = samplesOut1_tmp_storage1 + uintptr((*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state))).Fframe_length)*2 + uintptr(2)*2
	if lostFlag == FLAG_DECODE_NORMAL {
		has_side = libc.BoolInt32(!(*(*int32)(unsafe.Pointer(bp)) != 0))
	} else {
		has_side = libc.BoolInt32(!((*OpusT_silk_decoder)(unsafe.Pointer(psDec)).Fprev_decode_only_middle != 0) || (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal == int32(2) && lostFlag == int32(FLAG_DECODE_LBRR) && *(*int32)(unsafe.Pointer(channel_state + 1*4392 + 2432 + uintptr((*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + 1*4392))).FnFramesDecoded)*4)) == int32(1))
	}
	(*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state))).FsPLC.Fenable_deep_plc = (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).Fenable_deep_plc
	/* Call decoder for one frame */
	n = 0
	for {
		if !(n < (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal) {
			break
		}
		if n == 0 || has_side != 0 {
			FrameIndex = (*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state))).FnFramesDecoded - n
			/* Use independent coding if no previous frame available */
			if FrameIndex <= 0 {
				condCoding1 = CODE_INDEPENDENTLY
			} else {
				if lostFlag == int32(FLAG_DECODE_LBRR) {
					if *(*int32)(unsafe.Pointer(channel_state + uintptr(n)*4392 + 2432 + uintptr(FrameIndex-int32(1))*4)) != 0 {
						v51 = int32(CODE_CONDITIONALLY)
					} else {
						v51 = CODE_INDEPENDENTLY
					}
					condCoding1 = v51
				} else {
					if n > 0 && (*OpusT_silk_decoder)(unsafe.Pointer(psDec)).Fprev_decode_only_middle != 0 {
						/* If we skipped a side frame in this packet, we don't
						   need LTP scaling; the LTP state is well-defined. */
						condCoding1 = int32(CODE_INDEPENDENTLY_NO_LTP_SCALING)
					} else {
						condCoding1 = int32(CODE_CONDITIONALLY)
					}
				}
			}
			ret = ret + Opus_silk_decode_frame(tls, channel_state+uintptr(n)*4392, psRangeDec, samplesOut1_tmp[n]+2*2, bp+4, lostFlag, condCoding1, arch)
		} else {
			libc.Xmemset(tls, samplesOut1_tmp[n]+2*2, 0, uint64(uint32(*(*OpusT_opus_int32)(unsafe.Pointer(bp + 4))))*uint64(2))
		}
		(*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + uintptr(n)*4392))).FnFramesDecoded = (*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state + uintptr(n)*4392))).FnFramesDecoded + 1
		goto _50
	_50:
		;
		n = n + 1
	}
	if (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsAPI == int32(2) && (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal == int32(2) {
		/* Convert Mid/Side to Left/Right */
		Opus_silk_stereo_MS_to_LR(tls, psDec+8784, samplesOut1_tmp[0], samplesOut1_tmp[int32(1)], bp+8, (*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state))).Ffs_kHz, *(*OpusT_opus_int32)(unsafe.Pointer(bp + 4)))
	} else {
		/* Buffering */
		libc.Xmemcpy(tls, samplesOut1_tmp[0], psDec+8784+4, uint64(uint32(2))*uint64(2))
		libc.Xmemcpy(tls, psDec+8784+4, samplesOut1_tmp[0]+uintptr(*(*OpusT_opus_int32)(unsafe.Pointer(bp + 4)))*2, uint64(uint32(2))*uint64(2))
	}
	/* Number of output samples */
	*(*OpusT_opus_int32)(unsafe.Pointer(nSamplesOut)) = *(*OpusT_opus_int32)(unsafe.Pointer(bp + 4)) * (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FAPI_sampleRate / (int32(int16((*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state))).Ffs_kHz)) * int32(int16(int32(1000))))
	/* Set up pointers to temp buffers */
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _53
	_53:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _55
_55:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v7 = libc.Xmalloc(tls, uint64(16))
		goto _57
	_57:
		st = v7
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v9 = st
	goto _59
_59:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(2)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v9)).Fglobal_stack))) & (uint64(uint32(2)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _61
	_61:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _63
_63:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _65
	_65:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _67
_67:
	if !(int64(int32(uint64(uint32(*(*OpusT_opus_int32)(unsafe.Pointer(nSamplesOut))))*(uint64(2)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+6611, int32(382))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v26 = libc.Xmalloc(tls, uint64(16))
		goto _69
	_69:
		st = v26
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v28 = st
	goto _71
_71:
	*(*uintptr)(unsafe.Pointer(v28 + 8)) += uintptr(uint64(uint32(*(*OpusT_opus_int32)(unsafe.Pointer(nSamplesOut)))) * (uint64(2) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v30 = libc.Xmalloc(tls, uint64(16))
		goto _73
	_73:
		st = v30
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v32 = st
	goto _75
_75:
	samplesOut2_tmp = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v32)).Fglobal_stack - uintptr(uint64(uint32(*(*OpusT_opus_int32)(unsafe.Pointer(nSamplesOut))))*(uint64(2)/uint64(1)))
	resample_out_ptr = samplesOut2_tmp
	n = 0
	for {
		if (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsAPI < (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal {
			v51 = (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsAPI
		} else {
			v51 = (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal
		}
		if !(n < v51) {
			break
		}
		/* Resample decoded signal to API_sampleRate */
		ret = ret + Opus_silk_resampler(tls, channel_state+uintptr(n)*4392+2448, resample_out_ptr, samplesOut1_tmp[n]+1*2, *(*OpusT_opus_int32)(unsafe.Pointer(bp + 4)))
		/* Interleave if stereo output and stereo stream */
		if (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsAPI == int32(2) {
			i = 0
			for {
				if !(i < *(*OpusT_opus_int32)(unsafe.Pointer(nSamplesOut))) {
					break
				}
				*(*OpusT_opus_res)(unsafe.Pointer(samplesOut + uintptr(n+int32(2)*i)*4)) = float32(float32(*(*OpusT_opus_int16)(unsafe.Pointer(resample_out_ptr + uintptr(i)*2))) * (float32(1) / float32(32768)))
				goto _78
			_78:
				;
				i = i + 1
			}
		} else {
			i = 0
			for {
				if !(i < *(*OpusT_opus_int32)(unsafe.Pointer(nSamplesOut))) {
					break
				}
				*(*OpusT_opus_res)(unsafe.Pointer(samplesOut + uintptr(i)*4)) = float32(float32(*(*OpusT_opus_int16)(unsafe.Pointer(resample_out_ptr + uintptr(i)*2))) * (float32(1) / float32(32768)))
				goto _79
			_79:
				;
				i = i + 1
			}
		}
		goto _76
	_76:
		;
		n = n + 1
	}
	/* Create two channel output from mono stream */
	if (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsAPI == int32(2) && (*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FnChannelsInternal == int32(1) {
		if stereo_to_mono != 0 {
			/* Resample right channel for newly collapsed stereo just in case
			   we weren't doing collapsing when switching to mono */
			ret = ret + Opus_silk_resampler(tls, channel_state+1*4392+2448, resample_out_ptr, samplesOut1_tmp[0]+1*2, *(*OpusT_opus_int32)(unsafe.Pointer(bp + 4)))
			i = 0
			for {
				if !(i < *(*OpusT_opus_int32)(unsafe.Pointer(nSamplesOut))) {
					break
				}
				*(*OpusT_opus_res)(unsafe.Pointer(samplesOut + uintptr(int32(1)+int32(2)*i)*4)) = float32(float32(*(*OpusT_opus_int16)(unsafe.Pointer(resample_out_ptr + uintptr(i)*2))) * (float32(1) / float32(32768)))
				goto _80
			_80:
				;
				i = i + 1
			}
		} else {
			i = 0
			for {
				if !(i < *(*OpusT_opus_int32)(unsafe.Pointer(nSamplesOut))) {
					break
				}
				*(*OpusT_opus_res)(unsafe.Pointer(samplesOut + uintptr(int32(1)+int32(2)*i)*4)) = *(*OpusT_opus_res)(unsafe.Pointer(samplesOut + uintptr(0+int32(2)*i)*4))
				goto _81
			_81:
				;
				i = i + 1
			}
		}
	}
	/* Export pitch lag, measured at 48 kHz sampling rate */
	if (*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state))).FprevSignalType == int32(TYPE_VOICED) {
		mult_tab = [3]int32{
			0: int32(6),
			1: int32(4),
			2: int32(3),
		}
		(*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FprevPitchLag = (*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state))).FlagPrev * mult_tab[((*(*OpusT_silk_decoder_state)(unsafe.Pointer(channel_state))).Ffs_kHz-int32(8))>>int32(2)]
	} else {
		(*OpusT_silk_DecControlStruct)(unsafe.Pointer(decControl)).FprevPitchLag = 0
	}
	if lostFlag == int32(FLAG_PACKET_LOST) {
		/* On packet loss, remove the gain clamping to prevent having the energy "bounce back"
		   if we lose packets when the energy is going down */
		i = 0
		for {
			if !(i < (*OpusT_silk_decoder)(unsafe.Pointer(psDec)).FnChannelsInternal) {
				break
			}
			(*(*OpusT_silk_decoder_state)(unsafe.Pointer(psDec + uintptr(i)*4392))).FLastGainIndex = int8(10)
			goto _82
		_82:
			;
			i = i + 1
		}
	} else {
		(*OpusT_silk_decoder)(unsafe.Pointer(psDec)).Fprev_decode_only_middle = *(*int32)(unsafe.Pointer(bp))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _84
	_84:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _86
_86:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
	return ret
}

const CELT_SIG_SCALE10 = "32768.f"

// C documentation
//
//	/* Gain scalar quantization with hysteresis, uniform on log scale */
