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

func Opus_silk_stereo_MS_to_LR(tls *libc.TLS, state uintptr, x1 uintptr, x2 uintptr, pred_Q13 uintptr, fs_kHz int32, frame_length int32) {
	var delta0_Q13, delta1_Q13, denom_Q16, n, v2, v3 int32
	var diff, pred0_Q13, pred1_Q13, sum OpusT_opus_int32
	_, _, _, _, _, _, _, _, _, _ = delta0_Q13, delta1_Q13, denom_Q16, diff, n, pred0_Q13, pred1_Q13, sum, v2, v3
	/* Buffering */
	libc.Xmemcpy(tls, x1, state+4, uint64(uint32(2))*uint64(2))
	libc.Xmemcpy(tls, x2, state+8, uint64(uint32(2))*uint64(2))
	libc.Xmemcpy(tls, state+4, x1+uintptr(frame_length)*2, uint64(uint32(2))*uint64(2))
	libc.Xmemcpy(tls, state+8, x2+uintptr(frame_length)*2, uint64(uint32(2))*uint64(2))
	/* Interpolate predictors and add prediction to side channel */
	pred0_Q13 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(state)))
	pred1_Q13 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(state + 1*2)))
	denom_Q16 = int32(1) << int32(16) / (int32(STEREO_INTERP_LEN_MS) * fs_kHz)
	delta0_Q13 = (int32(int16(*(*OpusT_opus_int32)(unsafe.Pointer(pred_Q13))-int32(*(*OpusT_opus_int16)(unsafe.Pointer(state)))))*int32(int16(denom_Q16))>>(int32(16)-int32(1)) + int32(1)) >> int32(1)
	delta1_Q13 = (int32(int16(*(*OpusT_opus_int32)(unsafe.Pointer(pred_Q13 + 1*4))-int32(*(*OpusT_opus_int16)(unsafe.Pointer(state + 1*2)))))*int32(int16(denom_Q16))>>(int32(16)-int32(1)) + int32(1)) >> int32(1)
	n = 0
	for {
		if !(n < int32(STEREO_INTERP_LEN_MS)*fs_kHz) {
			break
		}
		pred0_Q13 = pred0_Q13 + delta0_Q13
		pred1_Q13 = pred1_Q13 + delta1_Q13
		sum = int32(uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(x1 + uintptr(n)*2)))+int32(*(*OpusT_opus_int16)(unsafe.Pointer(x1 + uintptr(n+int32(2))*2)))+int32(uint32(uint16(*(*OpusT_opus_int16)(unsafe.Pointer(x1 + uintptr(n+int32(1))*2))))<<int32(1))) << int32(9)) /* Q11 */
		sum = int32(int64(int32(uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(x2 + uintptr(n+int32(1))*2))))<<int32(8))) + int64(sum)*int64(int16(pred0_Q13))>>int32(16))                                                                                                      /* Q8  */
		sum = int32(int64(sum) + int64(int32(uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(x1 + uintptr(n+int32(1))*2))))<<int32(11)))*int64(int16(pred1_Q13))>>int32(16))                                                                                                     /* Q8  */
		if (sum>>(int32(8)-int32(1))+int32(1))>>int32(1) > int32(silk_int16_MAX11) {
			v2 = int32(silk_int16_MAX11)
		} else {
			if (sum>>(int32(8)-int32(1))+int32(1))>>int32(1) < int32(int16(-32768)) {
				v3 = int32(int16(-32768))
			} else {
				v3 = (sum>>(int32(8)-int32(1)) + int32(1)) >> int32(1)
			}
			v2 = v3
		}
		*(*OpusT_opus_int16)(unsafe.Pointer(x2 + uintptr(n+int32(1))*2)) = int16(v2)
		goto _1
	_1:
		;
		n = n + 1
	}
	pred0_Q13 = *(*OpusT_opus_int32)(unsafe.Pointer(pred_Q13))
	pred1_Q13 = *(*OpusT_opus_int32)(unsafe.Pointer(pred_Q13 + 1*4))
	n = int32(STEREO_INTERP_LEN_MS) * fs_kHz
	for {
		if !(n < frame_length) {
			break
		}
		sum = int32(uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(x1 + uintptr(n)*2)))+int32(*(*OpusT_opus_int16)(unsafe.Pointer(x1 + uintptr(n+int32(2))*2)))+int32(uint32(uint16(*(*OpusT_opus_int16)(unsafe.Pointer(x1 + uintptr(n+int32(1))*2))))<<int32(1))) << int32(9)) /* Q11 */
		sum = int32(int64(int32(uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(x2 + uintptr(n+int32(1))*2))))<<int32(8))) + int64(sum)*int64(int16(pred0_Q13))>>int32(16))                                                                                                      /* Q8  */
		sum = int32(int64(sum) + int64(int32(uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(x1 + uintptr(n+int32(1))*2))))<<int32(11)))*int64(int16(pred1_Q13))>>int32(16))                                                                                                     /* Q8  */
		if (sum>>(int32(8)-int32(1))+int32(1))>>int32(1) > int32(silk_int16_MAX11) {
			v2 = int32(silk_int16_MAX11)
		} else {
			if (sum>>(int32(8)-int32(1))+int32(1))>>int32(1) < int32(int16(-32768)) {
				v3 = int32(int16(-32768))
			} else {
				v3 = (sum>>(int32(8)-int32(1)) + int32(1)) >> int32(1)
			}
			v2 = v3
		}
		*(*OpusT_opus_int16)(unsafe.Pointer(x2 + uintptr(n+int32(1))*2)) = int16(v2)
		goto _4
	_4:
		;
		n = n + 1
	}
	*(*OpusT_opus_int16)(unsafe.Pointer(state)) = int16(*(*OpusT_opus_int32)(unsafe.Pointer(pred_Q13)))
	*(*OpusT_opus_int16)(unsafe.Pointer(state + 1*2)) = int16(*(*OpusT_opus_int32)(unsafe.Pointer(pred_Q13 + 1*4)))
	/* Convert to left/right signals */
	n = 0
	for {
		if !(n < frame_length) {
			break
		}
		sum = int32(*(*OpusT_opus_int16)(unsafe.Pointer(x1 + uintptr(n+int32(1))*2))) + int32(*(*OpusT_opus_int16)(unsafe.Pointer(x2 + uintptr(n+int32(1))*2)))
		diff = int32(*(*OpusT_opus_int16)(unsafe.Pointer(x1 + uintptr(n+int32(1))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer(x2 + uintptr(n+int32(1))*2)))
		if sum > int32(silk_int16_MAX11) {
			v2 = int32(silk_int16_MAX11)
		} else {
			if sum < int32(int16(-32768)) {
				v3 = int32(int16(-32768))
			} else {
				v3 = sum
			}
			v2 = v3
		}
		*(*OpusT_opus_int16)(unsafe.Pointer(x1 + uintptr(n+int32(1))*2)) = int16(v2)
		if diff > int32(silk_int16_MAX11) {
			v2 = int32(silk_int16_MAX11)
		} else {
			if diff < int32(int16(-32768)) {
				v3 = int32(int16(-32768))
			} else {
				v3 = diff
			}
			v2 = v3
		}
		*(*OpusT_opus_int16)(unsafe.Pointer(x2 + uintptr(n+int32(1))*2)) = int16(v2)
		goto _7
	_7:
		;
		n = n + 1
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
//	/* Coefficients for 2-band filter bank based on first-order allpass filters */
var A_fb1_20 = int16(int32(5394) << int32(1))
var A_fb1_21 = int16(-int32(24290)) /* (opus_int16)(20623 << 1) */

// C documentation
//
//	/* Split signal into two decimated bands using first-order allpass filters */
func Opus_silk_ana_filt_bank_1(tls *libc.TLS, in uintptr, S uintptr, outL uintptr, outH uintptr, N OpusT_opus_int32) {
	var N2, k, v2, v3 int32
	var X, Y, in32, out_1, out_2 OpusT_opus_int32
	_, _, _, _, _, _, _, _, _ = N2, X, Y, in32, k, out_1, out_2, v2, v3
	N2 = N >> int32(1)
	/* Internal variables and state are in Q10 format */
	k = 0
	for {
		if !(k < N2) {
			break
		}
		/* Convert to Q10 */
		in32 = int32(uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(in + uintptr(int32(2)*k)*2)))) << int32(10))
		/* All-pass section for even input sample */
		Y = in32 - *(*OpusT_opus_int32)(unsafe.Pointer(S))
		X = int32(int64(Y) + int64(Y)*int64(A_fb1_21)>>int32(16))
		out_1 = *(*OpusT_opus_int32)(unsafe.Pointer(S)) + X
		*(*OpusT_opus_int32)(unsafe.Pointer(S)) = in32 + X
		/* Convert to Q10 */
		in32 = int32(uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(in + uintptr(int32(2)*k+int32(1))*2)))) << int32(10))
		/* All-pass section for odd input sample, and add to output of previous section */
		Y = in32 - *(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4))
		X = int32(int64(Y) * int64(A_fb1_20) >> int32(16))
		out_2 = *(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4)) + X
		*(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4)) = in32 + X
		/* Add/subtract, convert back to int16 and store to output */
		if ((out_2+out_1)>>(int32(11)-int32(1))+int32(1))>>int32(1) > int32(silk_int16_MAX11) {
			v2 = int32(silk_int16_MAX11)
		} else {
			if ((out_2+out_1)>>(int32(11)-int32(1))+int32(1))>>int32(1) < int32(int16(-32768)) {
				v3 = int32(int16(-32768))
			} else {
				v3 = ((out_2+out_1)>>(int32(11)-int32(1)) + int32(1)) >> int32(1)
			}
			v2 = v3
		}
		*(*OpusT_opus_int16)(unsafe.Pointer(outL + uintptr(k)*2)) = int16(v2)
		if ((out_2-out_1)>>(int32(11)-int32(1))+int32(1))>>int32(1) > int32(silk_int16_MAX11) {
			v2 = int32(silk_int16_MAX11)
		} else {
			if ((out_2-out_1)>>(int32(11)-int32(1))+int32(1))>>int32(1) < int32(int16(-32768)) {
				v3 = int32(int16(-32768))
			} else {
				v3 = ((out_2-out_1)>>(int32(11)-int32(1)) + int32(1)) >> int32(1)
			}
			v2 = v3
		}
		*(*OpusT_opus_int16)(unsafe.Pointer(outH + uintptr(k)*2)) = int16(v2)
		goto _1
	_1:
		;
		k = k + 1
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
//	/* Second order ARMA filter, alternative implementation */
func Opus_silk_biquad_alt_stride1(tls *libc.TLS, in uintptr, B_Q28 uintptr, A_Q28 uintptr, S uintptr, out uintptr, len1 OpusT_opus_int32) {
	var A0_L_Q28, A0_U_Q28, A1_L_Q28, A1_U_Q28, inval, out32_Q14 OpusT_opus_int32
	var k, v2, v3 int32
	_, _, _, _, _, _, _, _, _ = A0_L_Q28, A0_U_Q28, A1_L_Q28, A1_U_Q28, inval, k, out32_Q14, v2, v3
	/* Negate A_Q28 values and split in two parts */
	A0_L_Q28 = -*(*OpusT_opus_int32)(unsafe.Pointer(A_Q28)) & int32(0x00003FFF)       /* lower part */
	A0_U_Q28 = -*(*OpusT_opus_int32)(unsafe.Pointer(A_Q28)) >> int32(14)              /* upper part */
	A1_L_Q28 = -*(*OpusT_opus_int32)(unsafe.Pointer(A_Q28 + 1*4)) & int32(0x00003FFF) /* lower part */
	A1_U_Q28 = -*(*OpusT_opus_int32)(unsafe.Pointer(A_Q28 + 1*4)) >> int32(14)        /* upper part */
	k = 0
	for {
		if !(k < len1) {
			break
		}
		/* S[ 0 ], S[ 1 ]: Q12 */
		inval = int32(*(*OpusT_opus_int16)(unsafe.Pointer(in + uintptr(k)*2)))
		out32_Q14 = int32(uint32(int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(S)))+int64(*(*OpusT_opus_int32)(unsafe.Pointer(B_Q28)))*int64(int16(inval))>>int32(16))) << int32(2))
		*(*OpusT_opus_int32)(unsafe.Pointer(S)) = *(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4)) + (int32(int64(out32_Q14)*int64(int16(A0_L_Q28))>>int32(16))>>(int32(14)-int32(1))+int32(1))>>int32(1)
		*(*OpusT_opus_int32)(unsafe.Pointer(S)) = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(S))) + int64(out32_Q14)*int64(int16(A0_U_Q28))>>int32(16))
		*(*OpusT_opus_int32)(unsafe.Pointer(S)) = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(S))) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(B_Q28 + 1*4)))*int64(int16(inval))>>int32(16))
		*(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4)) = (int32(int64(out32_Q14)*int64(int16(A1_L_Q28))>>int32(16))>>(int32(14)-int32(1)) + int32(1)) >> int32(1)
		*(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4)) = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4))) + int64(out32_Q14)*int64(int16(A1_U_Q28))>>int32(16))
		*(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4)) = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4))) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(B_Q28 + 2*4)))*int64(int16(inval))>>int32(16))
		/* Scale back to Q0 and saturate */
		if (out32_Q14+int32(1)<<int32(14)-int32(1))>>int32(14) > int32(silk_int16_MAX11) {
			v2 = int32(silk_int16_MAX11)
		} else {
			if (out32_Q14+int32(1)<<int32(14)-int32(1))>>int32(14) < int32(int16(-32768)) {
				v3 = int32(int16(-32768))
			} else {
				v3 = (out32_Q14 + int32(1)<<int32(14) - int32(1)) >> int32(14)
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

func Opus_silk_biquad_alt_stride2_c(tls *libc.TLS, in uintptr, B_Q28 uintptr, A_Q28 uintptr, S uintptr, out uintptr, len1 OpusT_opus_int32) {
	var A0_L_Q28, A0_U_Q28, A1_L_Q28, A1_U_Q28 OpusT_opus_int32
	var k, v2, v3 int32
	var out32_Q14 [2]OpusT_opus_int32
	_, _, _, _, _, _, _, _ = A0_L_Q28, A0_U_Q28, A1_L_Q28, A1_U_Q28, k, out32_Q14, v2, v3
	/* Negate A_Q28 values and split in two parts */
	A0_L_Q28 = -*(*OpusT_opus_int32)(unsafe.Pointer(A_Q28)) & int32(0x00003FFF)       /* lower part */
	A0_U_Q28 = -*(*OpusT_opus_int32)(unsafe.Pointer(A_Q28)) >> int32(14)              /* upper part */
	A1_L_Q28 = -*(*OpusT_opus_int32)(unsafe.Pointer(A_Q28 + 1*4)) & int32(0x00003FFF) /* lower part */
	A1_U_Q28 = -*(*OpusT_opus_int32)(unsafe.Pointer(A_Q28 + 1*4)) >> int32(14)        /* upper part */
	k = 0
	for {
		if !(k < len1) {
			break
		}
		/* S[ 0 ], S[ 1 ], S[ 2 ], S[ 3 ]: Q12 */
		out32_Q14[0] = int32(uint32(int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(S)))+int64(*(*OpusT_opus_int32)(unsafe.Pointer(B_Q28)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(in + uintptr(int32(2)*k+0)*2)))>>int32(16))) << int32(2))
		out32_Q14[int32(1)] = int32(uint32(int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(S + 2*4)))+int64(*(*OpusT_opus_int32)(unsafe.Pointer(B_Q28)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(in + uintptr(int32(2)*k+int32(1))*2)))>>int32(16))) << int32(2))
		*(*OpusT_opus_int32)(unsafe.Pointer(S)) = *(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4)) + (int32(int64(out32_Q14[0])*int64(int16(A0_L_Q28))>>int32(16))>>(int32(14)-int32(1))+int32(1))>>int32(1)
		*(*OpusT_opus_int32)(unsafe.Pointer(S + 2*4)) = *(*OpusT_opus_int32)(unsafe.Pointer(S + 3*4)) + (int32(int64(out32_Q14[int32(1)])*int64(int16(A0_L_Q28))>>int32(16))>>(int32(14)-int32(1))+int32(1))>>int32(1)
		*(*OpusT_opus_int32)(unsafe.Pointer(S)) = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(S))) + int64(out32_Q14[0])*int64(int16(A0_U_Q28))>>int32(16))
		*(*OpusT_opus_int32)(unsafe.Pointer(S + 2*4)) = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(S + 2*4))) + int64(out32_Q14[int32(1)])*int64(int16(A0_U_Q28))>>int32(16))
		*(*OpusT_opus_int32)(unsafe.Pointer(S)) = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(S))) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(B_Q28 + 1*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(in + uintptr(int32(2)*k+0)*2)))>>int32(16))
		*(*OpusT_opus_int32)(unsafe.Pointer(S + 2*4)) = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(S + 2*4))) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(B_Q28 + 1*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(in + uintptr(int32(2)*k+int32(1))*2)))>>int32(16))
		*(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4)) = (int32(int64(out32_Q14[0])*int64(int16(A1_L_Q28))>>int32(16))>>(int32(14)-int32(1)) + int32(1)) >> int32(1)
		*(*OpusT_opus_int32)(unsafe.Pointer(S + 3*4)) = (int32(int64(out32_Q14[int32(1)])*int64(int16(A1_L_Q28))>>int32(16))>>(int32(14)-int32(1)) + int32(1)) >> int32(1)
		*(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4)) = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4))) + int64(out32_Q14[0])*int64(int16(A1_U_Q28))>>int32(16))
		*(*OpusT_opus_int32)(unsafe.Pointer(S + 3*4)) = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(S + 3*4))) + int64(out32_Q14[int32(1)])*int64(int16(A1_U_Q28))>>int32(16))
		*(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4)) = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(S + 1*4))) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(B_Q28 + 2*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(in + uintptr(int32(2)*k+0)*2)))>>int32(16))
		*(*OpusT_opus_int32)(unsafe.Pointer(S + 3*4)) = int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(S + 3*4))) + int64(*(*OpusT_opus_int32)(unsafe.Pointer(B_Q28 + 2*4)))*int64(*(*OpusT_opus_int16)(unsafe.Pointer(in + uintptr(int32(2)*k+int32(1))*2)))>>int32(16))
		/* Scale back to Q0 and saturate */
		if (out32_Q14[0]+int32(1)<<int32(14)-int32(1))>>int32(14) > int32(silk_int16_MAX11) {
			v2 = int32(silk_int16_MAX11)
		} else {
			if (out32_Q14[0]+int32(1)<<int32(14)-int32(1))>>int32(14) < int32(int16(-32768)) {
				v3 = int32(int16(-32768))
			} else {
				v3 = (out32_Q14[0] + int32(1)<<int32(14) - int32(1)) >> int32(14)
			}
			v2 = v3
		}
		*(*OpusT_opus_int16)(unsafe.Pointer(out + uintptr(int32(2)*k+0)*2)) = int16(v2)
		if (out32_Q14[int32(1)]+int32(1)<<int32(14)-int32(1))>>int32(14) > int32(silk_int16_MAX11) {
			v2 = int32(silk_int16_MAX11)
		} else {
			if (out32_Q14[int32(1)]+int32(1)<<int32(14)-int32(1))>>int32(14) < int32(int16(-32768)) {
				v3 = int32(int16(-32768))
			} else {
				v3 = (out32_Q14[int32(1)] + int32(1)<<int32(14) - int32(1)) >> int32(14)
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

const silk_int16_MAX12 = 0x7FFF

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
//	/* Chirp (bandwidth expand) LP AR filter.
//	   This logic is reused in _celt_lpc(). Any bug fixes should also be applied there. */
func Opus_silk_bwexpander_32(tls *libc.TLS, ar uintptr, d int32, chirp_Q16 OpusT_opus_int32) {
	var chirp_minus_one_Q16 OpusT_opus_int32
	var i int32
	_, _ = chirp_minus_one_Q16, i
	chirp_minus_one_Q16 = chirp_Q16 - int32(65536)
	i = 0
	for {
		if !(i < d-int32(1)) {
			break
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(ar + uintptr(i)*4)) = int32(int64(chirp_Q16) * int64(*(*OpusT_opus_int32)(unsafe.Pointer(ar + uintptr(i)*4))) >> int32(16))
		chirp_Q16 = chirp_Q16 + (chirp_Q16*chirp_minus_one_Q16>>(int32(16)-int32(1))+int32(1))>>int32(1)
		goto _1
	_1:
		;
		i = i + 1
	}
	*(*OpusT_opus_int32)(unsafe.Pointer(ar + uintptr(d-int32(1))*4)) = int32(int64(chirp_Q16) * int64(*(*OpusT_opus_int32)(unsafe.Pointer(ar + uintptr(d-int32(1))*4))) >> int32(16))
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
//	/* Chirp (bandwidth expand) LP AR filter */
func Opus_silk_bwexpander(tls *libc.TLS, ar uintptr, d int32, chirp_Q16 OpusT_opus_int32) {
	var chirp_minus_one_Q16 OpusT_opus_int32
	var i int32
	_, _ = chirp_minus_one_Q16, i
	chirp_minus_one_Q16 = chirp_Q16 - int32(65536)
	/* NB: Dont use silk_SMULWB, instead of silk_RSHIFT_ROUND( silk_MUL(), 16 ), below.  */
	/* Bias in silk_SMULWB can lead to unstable filters                                */
	i = 0
	for {
		if !(i < d-int32(1)) {
			break
		}
		*(*OpusT_opus_int16)(unsafe.Pointer(ar + uintptr(i)*2)) = int16((chirp_Q16*int32(*(*OpusT_opus_int16)(unsafe.Pointer(ar + uintptr(i)*2)))>>(int32(16)-int32(1)) + int32(1)) >> int32(1))
		chirp_Q16 = chirp_Q16 + (chirp_Q16*chirp_minus_one_Q16>>(int32(16)-int32(1))+int32(1))>>int32(1)
		goto _1
	_1:
		;
		i = i + 1
	}
	*(*OpusT_opus_int16)(unsafe.Pointer(ar + uintptr(d-int32(1))*2)) = int16((chirp_Q16*int32(*(*OpusT_opus_int16)(unsafe.Pointer(ar + uintptr(d-int32(1))*2)))>>(int32(16)-int32(1)) + int32(1)) >> int32(1))
}

/* config_ccgo.h
 *
 * ccgo build configuration for producing a pure-Go libopus translation.
 *
 * We include the project's generated config.h, then disable all CPU RTCD and
 * x86 intrinsics so only portable C paths are used.
 */

/* config.h.  Generated from config.h.in by configure.  */
/* config.h.in.  Generated from configure.ac by autoheader.  */

/* Get CPU Info by asm method */

/* Get CPU Info by c method */
/* #undef CPU_INFO_BY_C */

/* Custom modes */
/* #undef CUSTOM_MODES */

/* Disable DNN debug float */

/* Disable dot product instructions */
/* #undef DISABLE_DOT_PROD */

/* Do not build the float API */
/* #undef DISABLE_FLOAT_API */

/* Disable bitstream fixes from RFC 8251 */
/* #undef DISABLE_UPDATE_DRAFT */

/* Assertions */
/* #undef ENABLE_ASSERTIONS */

/* Deep PLC */
/* #undef ENABLE_DEEP_PLC */

/* DRED */
/* #undef ENABLE_DRED */

/* Hardening */

/* LOSSGEN */
/* #undef ENABLE_LOSSGEN */

/* Opus custom API */
/* #undef ENABLE_OPUS_CUSTOM_API */

/* Enable Opus Speech Coding Enhancement */
/* #undef ENABLE_OSCE */

/* Enable Opus Speech Coding Enhancement Blind BWE */
/* #undef ENABLE_OSCE_BWE */

/* Enable dumping of OSCE training data */
/* #undef ENABLE_OSCE_TRAINING_DATA */

/* Scalable quality extension */
/* #undef ENABLE_QEXT */

/* 24-bit internal resolution for fixed-point */

/* Debug fixed-point implementation */
/* #undef FIXED_DEBUG */

/* Compile as fixed-point (for machines without a fast enough FPU) */
/* #undef FIXED_POINT */

/* Float approximations */

/* Fuzzing */
/* #undef FUZZING */

/* Define to 1 if you have the <alloca.h> header file. */
/* #undef HAVE_ALLOCA_H */

/* NE10 library is installed on host. Make sure it is on target! */
/* #undef HAVE_ARM_NE10 */

/* Define to 1 if you have the <dlfcn.h> header file. */

/* Define to 1 if you have the `elf_aux_info' function. */
/* #undef HAVE_ELF_AUX_INFO */

/* Define to 1 if you have the <inttypes.h> header file. */

/* Define to 1 if you have the `lrint' function. */

/* Define to 1 if you have the `lrintf' function. */

/* Define to 1 if you have the <stdint.h> header file. */

/* Define to 1 if you have the <stdio.h> header file. */

/* Define to 1 if you have the <stdlib.h> header file. */

/* Define to 1 if you have the <strings.h> header file. */

/* Define to 1 if you have the <string.h> header file. */

/* Define to 1 if you have the <sys/stat.h> header file. */

/* Define to 1 if you have the <sys/types.h> header file. */

/* Define to 1 if you have the <unistd.h> header file. */

/* Define to 1 if you have the `__malloc_hook' function. */
/* #undef HAVE___MALLOC_HOOK */

/* Define to the sub-directory where libtool stores uninstalled libraries. */

/* Make use of ARM asm optimization */
/* #undef OPUS_ARM_ASM */

/* Use generic ARMv4 inline asm optimizations */
/* #undef OPUS_ARM_INLINE_ASM */

/* Use ARMv5E inline asm optimizations */
/* #undef OPUS_ARM_INLINE_EDSP */

/* Use ARMv6 inline asm optimizations */
/* #undef OPUS_ARM_INLINE_MEDIA */

/* Use ARM NEON inline asm optimizations */
/* #undef OPUS_ARM_INLINE_NEON */

/* Compiler supports Aarch64 DOTPROD Intrinsics */
/* #undef OPUS_ARM_MAY_HAVE_DOTPROD */

/* Define if assembler supports EDSP instructions */
/* #undef OPUS_ARM_MAY_HAVE_EDSP */

/* Define if assembler supports ARMv6 media instructions */
/* #undef OPUS_ARM_MAY_HAVE_MEDIA */

/* Define if compiler supports NEON instructions */
/* #undef OPUS_ARM_MAY_HAVE_NEON */

/* Compiler supports ARMv7/Aarch64 Neon Intrinsics */
/* #undef OPUS_ARM_MAY_HAVE_NEON_INTR */

/* Define if binary requires Aarch64 Neon Intrinsics */
/* #undef OPUS_ARM_PRESUME_AARCH64_NEON_INTR */

/* Define if binary requires Aarch64 dotprod Intrinsics */
/* #undef OPUS_ARM_PRESUME_DOTPROD */

/* Define if binary requires EDSP instruction support */
/* #undef OPUS_ARM_PRESUME_EDSP */

/* Define if binary requires ARMv6 media instruction support */
/* #undef OPUS_ARM_PRESUME_MEDIA */

/* Define if binary requires NEON instruction support */
/* #undef OPUS_ARM_PRESUME_NEON */

/* Define if binary requires NEON intrinsics support */
/* #undef OPUS_ARM_PRESUME_NEON_INTR */

/* This is a build of OPUS */

/* Run bit-exactness checks between optimized and C implementations */
/* #undef OPUS_CHECK_ASM */

/* Use run-time CPU capabilities detection */

/* Compiler supports X86 AVX2 Intrinsics */

/* Compiler supports X86 SSE Intrinsics */

/* Compiler supports X86 SSE2 Intrinsics */

/* Compiler supports X86 SSE4.1 Intrinsics */

/* Define if binary requires AVX2 intrinsics support */
/* #undef OPUS_X86_PRESUME_AVX2 */

/* Define if binary requires SSE intrinsics support */

/* Define if binary requires SSE2 intrinsics support */

/* Define if binary requires SSE4.1 intrinsics support */
/* #undef OPUS_X86_PRESUME_SSE4_1 */

/* Define to the address where bug reports for this package should be sent. */

/* Define to the full name of this package. */

/* Define to the full name and version of this package. */

/* Define to the one symbol short name of this package. */

/* Define to the home page for this package. */

/* Define to the version of this package. */

/* Define to 1 if all of the C90 standard headers exist (not just the ones
   required in a freestanding environment). This macro is provided for
   backward compatibility; new code need not use it. */

/* Make use of alloca */
/* #undef USE_ALLOCA */

/* Use C99 variable-size arrays */

/* Define to empty if `const' does not conform to ANSI C. */
/* #undef const */

/* Define to `__inline__' or `__inline' if that's what the C compiler
   calls it, or to nothing if 'inline' is not supported under any name.  */
/* #undef inline */

/* Define to the equivalent of the C99 'restrict' keyword, or to
   nothing if this is not supported.  Do not define if restrict is
   supported directly.  */
/* Work around a bug in Sun C++: it does not support _Restrict or
   __restrict__, even though the corresponding Sun C compiler ends up with
   "#define restrict _Restrict" or "#define restrict __restrict__" in the
   previous line.  Perhaps some future version of Sun C++ will work with
   restrict; if so, hopefully it defines __RESTRICT like Sun C does.  */

/* Enable ccgo-specific concurrency-safe pseudostack plumbing.
 * See celt/stack_alloc.h for details.
 */

/* Temp allocation mode: NONTHREADSAFE_PSEUDOSTACK is fastest, but the upstream
 * implementation uses global variables. We'll patch the generated Go to store
 * the pseudostack state in per-decoder TLS (no shared globals).
 */

/* Disable runtime CPU detection tables and any x86 intrinsic paths.
 *
 * Important: many libopus headers use `#if defined(OPUS_...)` checks, so we
 * must UNDEFINE these macros (not merely set them to 0).
 */

/* Also request that intrinsics are not used even if headers try. */
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

type OpusT_prevent_empty_translation_unit_warning = int32

const PE_D_SRCH_LENGTH = 24
const PE_FLATCONTOUR_BIAS = "0.05f"
const PE_MAX_FS_KHZ = 16
const PE_MAX_LAG_MS = 18
const PE_MAX_NB_SUBFR = 4
const PE_MIN_LAG_MS = 2
const PE_NB_CBKS_STAGE2 = 3
const PE_NB_CBKS_STAGE2_10MS = 3
const PE_NB_CBKS_STAGE2_EXT = 11
const PE_NB_CBKS_STAGE3_10MS = 12
const PE_NB_CBKS_STAGE3_MAX = 34
const PE_NB_CBKS_STAGE3_MID = 24
const PE_NB_CBKS_STAGE3_MIN = 16
const PE_NB_STAGE3_LAGS = 5
const PE_PREVLAG_BIAS = "0.2f"
const PE_SHORTLAG_BIAS = "0.2f"
const PE_SUBFR_LENGTH_MS = 5
const SILK_PE_MAX_COMPLEX = 2
const SILK_PE_MID_COMPLEX = 1
const SILK_PE_MIN_COMPLEX = 0

func Opus_silk_decode_pitch(tls *libc.TLS, lagIndex OpusT_opus_int16, contourIndex OpusT_opus_int8, pitch_lags uintptr, Fs_kHz int32, nb_subfr int32) {
	var Lag_CB_ptr uintptr
	var cbk_size, k, lag, max_lag, min_lag, v2, v3, v4, v5, v6 int32
	_, _, _, _, _, _, _, _, _, _, _ = Lag_CB_ptr, cbk_size, k, lag, max_lag, min_lag, v2, v3, v4, v5, v6
	if Fs_kHz == int32(8) {
		if nb_subfr == int32(PE_MAX_NB_SUBFR) {
			Lag_CB_ptr = uintptr(unsafe.Pointer(&Opus_silk_CB_lags_stage2))
			cbk_size = int32(PE_NB_CBKS_STAGE2_EXT)
		} else {
			if !(nb_subfr == int32(PE_MAX_NB_SUBFR)>>int32(1)) {
				Opus_celt_fatal(tls, __ccgo_ts+7059, __ccgo_ts+7110, int32(54))
			}
			Lag_CB_ptr = uintptr(unsafe.Pointer(&Opus_silk_CB_lags_stage2_10_ms))
			cbk_size = int32(PE_NB_CBKS_STAGE2_10MS)
		}
	} else {
		if nb_subfr == int32(PE_MAX_NB_SUBFR) {
			Lag_CB_ptr = uintptr(unsafe.Pointer(&Opus_silk_CB_lags_stage3))
			cbk_size = int32(PE_NB_CBKS_STAGE3_MAX)
		} else {
			if !(nb_subfr == int32(PE_MAX_NB_SUBFR)>>int32(1)) {
				Opus_celt_fatal(tls, __ccgo_ts+7059, __ccgo_ts+7110, int32(63))
			}
			Lag_CB_ptr = uintptr(unsafe.Pointer(&Opus_silk_CB_lags_stage3_10_ms))
			cbk_size = int32(PE_NB_CBKS_STAGE3_10MS)
		}
	}
	min_lag = int32(int16(int32(PE_MIN_LAG_MS))) * int32(int16(Fs_kHz))
	max_lag = int32(int16(int32(PE_MAX_LAG_MS))) * int32(int16(Fs_kHz))
	lag = min_lag + int32(lagIndex)
	k = 0
	for {
		if !(k < nb_subfr) {
			break
		}
		*(*int32)(unsafe.Pointer(pitch_lags + uintptr(k)*4)) = lag + int32(*(*OpusT_opus_int8)(unsafe.Pointer(Lag_CB_ptr + uintptr(k*cbk_size+int32(contourIndex)))))
		if min_lag > max_lag {
			if *(*int32)(unsafe.Pointer(pitch_lags + uintptr(k)*4)) > min_lag {
				v3 = min_lag
			} else {
				if *(*int32)(unsafe.Pointer(pitch_lags + uintptr(k)*4)) < max_lag {
					v4 = max_lag
				} else {
					v4 = *(*int32)(unsafe.Pointer(pitch_lags + uintptr(k)*4))
				}
				v3 = v4
			}
			v2 = v3
		} else {
			if *(*int32)(unsafe.Pointer(pitch_lags + uintptr(k)*4)) > max_lag {
				v5 = max_lag
			} else {
				if *(*int32)(unsafe.Pointer(pitch_lags + uintptr(k)*4)) < min_lag {
					v6 = min_lag
				} else {
					v6 = *(*int32)(unsafe.Pointer(pitch_lags + uintptr(k)*4))
				}
				v5 = v6
			}
			v2 = v5
		}
		*(*int32)(unsafe.Pointer(pitch_lags + uintptr(k)*4)) = v2
		goto _1
	_1:
		;
		k = k + 1
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

func Opus_silk_inner_prod_aligned_scale(tls *libc.TLS, inVec1 uintptr, inVec2 uintptr, scale int32, len1 int32) (r OpusT_opus_int32) {
	var i int32
	var sum OpusT_opus_int32
	_, _ = i, sum
	sum = 0
	i = 0
	for {
		if !(i < len1) {
			break
		}
		sum = sum + int32(*(*OpusT_opus_int16)(unsafe.Pointer(inVec1 + uintptr(i)*2)))*int32(*(*OpusT_opus_int16)(unsafe.Pointer(inVec2 + uintptr(i)*2)))>>scale
		goto _1
	_1:
		;
		i = i + 1
	}
	return sum
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
//	/* Approximation of 128 * log2() (very close inverse of silk_log2lin()) */
//	/* Convert input to a log scale    */
func Opus_silk_lin2log(tls *libc.TLS, inLin OpusT_opus_int32) (r1 OpusT_opus_int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var lzeros, v1, v2, v3, v6, v8 OpusT_opus_int32
	var m, r, x OpusT_opus_uint32
	var v5, v7 int32
	var _ /* frac_Q7 at bp+4 */ OpusT_opus_int32
	var _ /* lz at bp+0 */ OpusT_opus_int32
	_, _, _, _, _, _, _, _, _, _, _ = lzeros, m, r, x, v1, v2, v3, v5, v6, v7, v8
	v1 = inLin
	v2 = v1
	if v2 != 0 {
		v5 = int32(32) - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, uint32(v2)))
	} else {
		v5 = int32(32)
	}
	v3 = v5
	goto _4
_4:
	lzeros = v3
	*(*OpusT_opus_int32)(unsafe.Pointer(bp)) = lzeros
	v6 = v1
	v7 = int32(24) - lzeros
	x = uint32(v6)
	r = uint32(v7)
	m = uint32(-v7)
	if v7 == int32(0) {
		v8 = v6
		goto _9
	} else {
		if v7 < int32(0) {
			v8 = int32(x<<m | x>>(uint32(32)-m))
			goto _9
		} else {
			v8 = int32(x<<(uint32(32)-r) | x>>r)
			goto _9
		}
	}
_9:
	*(*OpusT_opus_int32)(unsafe.Pointer(bp + 4)) = v8 & int32(0x7f)
	/* Piece-wise parabolic approximation */
	return int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(bp + 4)))+int64(*(*OpusT_opus_int32)(unsafe.Pointer(bp + 4))*(int32(128)-*(*OpusT_opus_int32)(unsafe.Pointer(bp + 4))))*int64(int16(int32(179)))>>int32(16)) + int32(uint32(int32(31)-*(*OpusT_opus_int32)(unsafe.Pointer(bp)))<<int32(7))
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
//	/* Approximation of 2^() (very close inverse of silk_lin2log()) */
//	/* Convert input to a linear scale    */
func Opus_silk_log2lin(tls *libc.TLS, inLog_Q7 OpusT_opus_int32) (r OpusT_opus_int32) {
	var frac_Q7, out OpusT_opus_int32
	_, _ = frac_Q7, out
	if inLog_Q7 < 0 {
		return 0
	} else {
		if inLog_Q7 >= int32(3967) {
			return int32(silk_int32_MAX)
		}
	}
	out = int32(uint32(int32(1)) << (inLog_Q7 >> int32(7)))
	frac_Q7 = inLog_Q7 & int32(0x7F)
	if inLog_Q7 < int32(2048) {
		/* Piece-wise parabolic approximation */
		out = out + out*int32(int64(frac_Q7)+int64(int32(int16(frac_Q7))*int32(int16(int32(128)-frac_Q7)))*int64(int16(-int32(174)))>>int32(16))>>int32(7)
	} else {
		/* Piece-wise parabolic approximation */
		out = out + out>>int32(7)*int32(int64(frac_Q7)+int64(int32(int16(frac_Q7))*int32(int16(int32(128)-frac_Q7)))*int64(int16(-int32(174)))>>int32(16))
	}
	return out
}

const USE_CELT_FIR = 0
const silk_int16_MAX13 = 32767

/*******************************************/
/* LPC analysis filter                     */
/* NB! State is kept internally and the    */
/* filter always starts with zero state    */
/* first d output samples are set to zero  */
/*******************************************/

/* OPT: Using celt_fir() for this function should be faster, but it may cause
   integer overflows in intermediate values (not final results), which the
   current implementation silences by casting to unsigned. Enabling
   this should be safe in pretty much all cases, even though it is not technically
   C89-compliant. */

func Opus_silk_LPC_analysis_filter(tls *libc.TLS, out uintptr, in uintptr, B uintptr, len1 OpusT_opus_int32, d OpusT_opus_int32, arch int32) {
	var in_ptr uintptr
	var ix, j, v3, v4 int32
	var out32, out32_Q12 OpusT_opus_int32
	_, _, _, _, _, _, _ = in_ptr, ix, j, out32, out32_Q12, v3, v4
	if !(d >= int32(6)) {
		Opus_celt_fatal(tls, __ccgo_ts+7133, __ccgo_ts+7158, int32(67))
	}
	if !(d&int32(1) == int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+7188, __ccgo_ts+7158, int32(68))
	}
	if !(d <= len1) {
		Opus_celt_fatal(tls, __ccgo_ts+7219, __ccgo_ts+7158, int32(69))
	}
	_ = arch
	ix = d
	for {
		if !(ix < len1) {
			break
		}
		in_ptr = in + uintptr(ix-int32(1))*2
		out32_Q12 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(in_ptr))) * int32(*(*OpusT_opus_int16)(unsafe.Pointer(B)))
		/* Allowing wrap around so that two wraps can cancel each other. The rare
		   cases where the result wraps around can only be triggered by invalid streams*/
		out32_Q12 = int32(uint32(out32_Q12) + uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(in_ptr-uintptr(1)*2)))*int32(*(*OpusT_opus_int16)(unsafe.Pointer(B + 1*2)))))
		out32_Q12 = int32(uint32(out32_Q12) + uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(in_ptr-uintptr(2)*2)))*int32(*(*OpusT_opus_int16)(unsafe.Pointer(B + 2*2)))))
		out32_Q12 = int32(uint32(out32_Q12) + uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(in_ptr-uintptr(3)*2)))*int32(*(*OpusT_opus_int16)(unsafe.Pointer(B + 3*2)))))
		out32_Q12 = int32(uint32(out32_Q12) + uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(in_ptr-uintptr(4)*2)))*int32(*(*OpusT_opus_int16)(unsafe.Pointer(B + 4*2)))))
		out32_Q12 = int32(uint32(out32_Q12) + uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(in_ptr-uintptr(5)*2)))*int32(*(*OpusT_opus_int16)(unsafe.Pointer(B + 5*2)))))
		j = int32(6)
		for {
			if !(j < d) {
				break
			}
			out32_Q12 = int32(uint32(out32_Q12) + uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(in_ptr + uintptr(-j)*2)))*int32(*(*OpusT_opus_int16)(unsafe.Pointer(B + uintptr(j)*2)))))
			out32_Q12 = int32(uint32(out32_Q12) + uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(in_ptr + uintptr(-j-int32(1))*2)))*int32(*(*OpusT_opus_int16)(unsafe.Pointer(B + uintptr(j+int32(1))*2)))))
			goto _2
		_2:
			;
			j = j + int32(2)
		}
		/* Subtract prediction */
		out32_Q12 = int32(uint32(int32(uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(in_ptr + 1*2))))<<int32(12))) - uint32(out32_Q12))
		/* Scale to Q0 */
		out32 = (out32_Q12>>(int32(12)-int32(1)) + int32(1)) >> int32(1)
		/* Saturate output */
		if out32 > int32(silk_int16_MAX13) {
			v3 = int32(silk_int16_MAX13)
		} else {
			if out32 < int32(int16(-32768)) {
				v4 = int32(int16(-32768))
			} else {
				v4 = out32
			}
			v3 = v4
		}
		*(*OpusT_opus_int16)(unsafe.Pointer(out + uintptr(ix)*2)) = int16(v3)
		goto _1
	_1:
		;
		ix = ix + 1
	}
	/* Set first d output samples to zero */
	libc.Xmemset(tls, out, 0, uint64(uint32(d))*uint64(2))
}

const MAX_PREDICTION_POWER_GAIN1 = 10000
const QA = 24
const silk_int16_MAX14 = 0x7FFF

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

/******************/
/* Error messages */
/******************/

/**************************/
/* Encoder error messages */
/**************************/

/* Input length is not a multiple of 10 ms, or length is longer than the packet length */

/* Sampling frequency not 8000, 12000 or 16000 Hertz */

/* Packet size not 10, 20, 40, or 60 ms */

/* Allocated payload buffer too short */

/* Loss rate not between 0 and 100 percent */

/* Complexity setting not valid, use 0...10 */

/* Inband FEC setting not valid, use 0 or 1 */

/* DTX setting not valid, use 0 or 1 */

/* CBR setting not valid, use 0 or 1 */

/* Internal encoder error */

/* Internal encoder error */

/**************************/
/* Decoder error messages */
/**************************/

/* Output sampling frequency lower than internal decoded sampling frequency */

/* Payload size exceeded the maximum allowed 1024 bytes */

/* Payload has bit errors */

/* Payload has bit errors */

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

/* Max number of encoder channels (1/2) */
/* Number of decoder channels (1/2) */

/* Limits on bitrate */

/* LBRR thresholds */

/* DTX settings */

/* VAD decision */

/* Maximum sampling frequency */

/* Signal types */

/* Conditional coding types */

/* Settings for stereo processing */

/* Range of pitch lag estimates */

/* Maximum number of subframes */

/* Number of samples per frame */

/* Milliseconds of lookahead for pitch analysis */

/* Order of LPC used in find pitch */

/* Length of LPC window used in find pitch */

/* Milliseconds of lookahead for noise shape analysis */

/* Maximum length of LPC window used in noise shape analysis */

/* dB level of lowest gain quantization level */
/* dB level of highest gain quantization level */
/* Number of gain quantization levels */
/* Max increase in gain quantization index */
/* Max decrease in gain quantization index */

/* Quantization offsets (multiples of 4) */

/* Maximum numbers of iterations used to stabilize an LPC vector */

/* Find Pred Coef defines */

/* LTP quantization settings */

/* Flag to use harmonic noise shaping */

/* Max LPC order of noise shaping filters */

/* Maximum number of delayed decision states */

/* Number of subframes for excitation entropy coding */

/* Number of rate levels, for entropy coding of excitation */

/* Maximum sum of pulses per shell coding frame */

/***************************/
/* Voice activity detector */
/***************************/

/* Sigmoid settings */

/* smoothing for SNR measurement */

/* Size of the piecewise linear cosine approximation table for the LSFs */

/******************/
/* NLSF quantizer */
/******************/

/* Transition filtering for mode switching */

/* BWE factors to apply after packet loss */

/* Defines for CN generation */

// C documentation
//
//	/* Compute inverse of LPC prediction gain, and                          */
//	/* test if LPC coefficients are stable (all poles within unit circle)   */
func LPC_inverse_pred_gain_QA_c(tls *libc.TLS, A_QA uintptr, order int32) (r OpusT_opus_int32) {
	var b32_inv, b32_nrm, err_Q32, invGain_Q30, rc_Q31, rc_mult1_Q30, rc_mult2, result, tmp1, tmp2, v10, v3, v4, v7 OpusT_opus_int32
	var b_headrm, k, lshift, mult2Q, n, v13, v16, v17, v18, v19, v2, v6, v8, v9 int32
	var tmp64 OpusT_opus_int64
	var v22 int64
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = b32_inv, b32_nrm, b_headrm, err_Q32, invGain_Q30, k, lshift, mult2Q, n, rc_Q31, rc_mult1_Q30, rc_mult2, result, tmp1, tmp2, tmp64, v10, v13, v16, v17, v18, v19, v2, v22, v3, v4, v6, v7, v8, v9
	invGain_Q30 = int32(1073741824)
	k = order - int32(1)
	for {
		if !(k > 0) {
			break
		}
		/* Check for stability */
		if *(*OpusT_opus_int32)(unsafe.Pointer(A_QA + uintptr(k)*4)) > int32(16773022) || *(*OpusT_opus_int32)(unsafe.Pointer(A_QA + uintptr(k)*4)) < -int32(16773022) {
			return 0
		}
		/* Set RC equal to negated AR coef */
		rc_Q31 = -int32(uint32(*(*OpusT_opus_int32)(unsafe.Pointer(A_QA + uintptr(k)*4))) << (int32(31) - int32(QA)))
		/* rc_mult1_Q30 range: [ 1 : 2^30 ] */
		rc_mult1_Q30 = int32(1073741824) - int32(int64(rc_Q31)*int64(rc_Q31)>>int32(32))
		_ = rc_mult1_Q30 > int32(1)<<int32(15) /* reduce A_LIMIT if fails */
		_ = rc_mult1_Q30 <= int32(1)<<int32(30)
		/* Update inverse gain */
		/* invGain_Q30 range: [ 0 : 2^30 ] */
		invGain_Q30 = int32(uint32(int32(int64(invGain_Q30)*int64(rc_mult1_Q30)>>int32(32))) << int32(2))
		_ = invGain_Q30 >= int32(0)
		_ = invGain_Q30 <= int32(1)<<int32(30)
		if invGain_Q30 < int32(107374) {
			return 0
		}
		/* rc_mult2 range: [ 2^30 : silk_int32_MAX ] */
		if rc_mult1_Q30 > 0 {
			v2 = rc_mult1_Q30
		} else {
			v2 = -rc_mult1_Q30
		}
		v3 = v2
		if v3 != 0 {
			v6 = int32(32) - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, uint32(v3)))
		} else {
			v6 = int32(32)
		}
		v4 = v6
		goto _5
	_5:
		mult2Q = int32(32) - v4
		v3 = rc_mult1_Q30
		v2 = mult2Q + int32(30)
		_ = v3 != int32(0)
		_ = v2 > int32(0)
		if v3 > 0 {
			v6 = v3
		} else {
			v6 = -v3
		}
		v4 = v6
		if v4 != 0 {
			v8 = int32(32) - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, uint32(v4)))
		} else {
			v8 = int32(32)
		}
		v7 = v8
		goto _12
	_12:
		b_headrm = v7 - int32(1)
		b32_nrm = int32(uint32(v3) << b_headrm)
		b32_inv = int32(silk_int32_MAX) >> int32(2) / (b32_nrm >> int32(16))
		result = int32(uint32(b32_inv) << int32(16))
		err_Q32 = int32(uint32(int32(1)<<int32(29)-int32(int64(b32_nrm)*int64(int16(b32_inv))>>int32(16))) << int32(3))
		result = int32(int64(result) + int64(err_Q32)*int64(b32_inv)>>int32(16))
		lshift = int32(61) - b_headrm - v2
		if lshift <= int32(0) {
			if int32(-2147483648)>>-lshift > int32(silk_int32_MAX)>>-lshift {
				if result > int32(-2147483648)>>-lshift {
					v13 = int32(-2147483648) >> -lshift
				} else {
					if result < int32(silk_int32_MAX)>>-lshift {
						v16 = int32(silk_int32_MAX) >> -lshift
					} else {
						v16 = result
					}
					v13 = v16
				}
				v9 = v13
			} else {
				if result > int32(silk_int32_MAX)>>-lshift {
					v17 = int32(silk_int32_MAX) >> -lshift
				} else {
					if result < int32(-2147483648)>>-lshift {
						v18 = int32(-2147483648) >> -lshift
					} else {
						v18 = result
					}
					v17 = v18
				}
				v9 = v17
			}
			v10 = int32(uint32(v9) << -lshift)
			goto _15
		} else {
			if lshift < int32(32) {
				v10 = result >> lshift
				goto _15
			} else {
				v10 = 0
				goto _15
			}
		}
	_15:
		rc_mult2 = v10
		/* Update AR coefficient */
		n = 0
		for {
			if !(n < (k+int32(1))>>int32(1)) {
				break
			}
			tmp1 = *(*OpusT_opus_int32)(unsafe.Pointer(A_QA + uintptr(n)*4))
			tmp2 = *(*OpusT_opus_int32)(unsafe.Pointer(A_QA + uintptr(k-n-int32(1))*4))
			if mult2Q == int32(1) {
				if (uint32(tmp1)-uint32(int32((int64(tmp2)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1))))&uint32(0x80000000) == uint32(0) {
					if uint32(tmp1)&(uint32(int32((int64(tmp2)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1)))^uint32(0x80000000))&uint32(0x80000000) != 0 {
						v6 = int32(-2147483648)
					} else {
						v6 = tmp1 - int32((int64(tmp2)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1))
					}
					v2 = v6
				} else {
					if (uint32(tmp1)^uint32(0x80000000))&uint32(int32((int64(tmp2)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1)))&uint32(0x80000000) != 0 {
						v8 = int32(silk_int32_MAX)
					} else {
						v8 = tmp1 - int32((int64(tmp2)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1))
					}
					v2 = v8
				}
				if (uint32(tmp1)-uint32(int32((int64(tmp2)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1))))&uint32(0x80000000) == uint32(0) {
					if uint32(tmp1)&(uint32(int32((int64(tmp2)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1)))^uint32(0x80000000))&uint32(0x80000000) != 0 {
						v13 = int32(-2147483648)
					} else {
						v13 = tmp1 - int32((int64(tmp2)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1))
					}
					v9 = v13
				} else {
					if (uint32(tmp1)^uint32(0x80000000))&uint32(int32((int64(tmp2)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1)))&uint32(0x80000000) != 0 {
						v16 = int32(silk_int32_MAX)
					} else {
						v16 = tmp1 - int32((int64(tmp2)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1))
					}
					v9 = v16
				}
				v22 = int64(v2)*int64(rc_mult2)>>int32(1) + int64(v9)*int64(rc_mult2)&int64(1)
			} else {
				if (uint32(tmp1)-uint32(int32((int64(tmp2)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1))))&uint32(0x80000000) == uint32(0) {
					if uint32(tmp1)&(uint32(int32((int64(tmp2)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1)))^uint32(0x80000000))&uint32(0x80000000) != 0 {
						v18 = int32(-2147483648)
					} else {
						v18 = tmp1 - int32((int64(tmp2)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1))
					}
					v17 = v18
				} else {
					if (uint32(tmp1)^uint32(0x80000000))&uint32(int32((int64(tmp2)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1)))&uint32(0x80000000) != 0 {
						v19 = int32(silk_int32_MAX)
					} else {
						v19 = tmp1 - int32((int64(tmp2)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1))
					}
					v17 = v19
				}
				v22 = (int64(v17)*int64(rc_mult2)>>(mult2Q-int32(1)) + int64(1)) >> int32(1)
			}
			tmp64 = v22
			if tmp64 > int64(silk_int32_MAX) || tmp64 < int64(int32(-2147483648)) {
				return 0
			}
			*(*OpusT_opus_int32)(unsafe.Pointer(A_QA + uintptr(n)*4)) = int32(tmp64)
			if mult2Q == int32(1) {
				if (uint32(tmp2)-uint32(int32((int64(tmp1)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1))))&uint32(0x80000000) == uint32(0) {
					if uint32(tmp2)&(uint32(int32((int64(tmp1)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1)))^uint32(0x80000000))&uint32(0x80000000) != 0 {
						v6 = int32(-2147483648)
					} else {
						v6 = tmp2 - int32((int64(tmp1)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1))
					}
					v2 = v6
				} else {
					if (uint32(tmp2)^uint32(0x80000000))&uint32(int32((int64(tmp1)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1)))&uint32(0x80000000) != 0 {
						v8 = int32(silk_int32_MAX)
					} else {
						v8 = tmp2 - int32((int64(tmp1)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1))
					}
					v2 = v8
				}
				if (uint32(tmp2)-uint32(int32((int64(tmp1)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1))))&uint32(0x80000000) == uint32(0) {
					if uint32(tmp2)&(uint32(int32((int64(tmp1)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1)))^uint32(0x80000000))&uint32(0x80000000) != 0 {
						v13 = int32(-2147483648)
					} else {
						v13 = tmp2 - int32((int64(tmp1)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1))
					}
					v9 = v13
				} else {
					if (uint32(tmp2)^uint32(0x80000000))&uint32(int32((int64(tmp1)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1)))&uint32(0x80000000) != 0 {
						v16 = int32(silk_int32_MAX)
					} else {
						v16 = tmp2 - int32((int64(tmp1)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1))
					}
					v9 = v16
				}
				v22 = int64(v2)*int64(rc_mult2)>>int32(1) + int64(v9)*int64(rc_mult2)&int64(1)
			} else {
				if (uint32(tmp2)-uint32(int32((int64(tmp1)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1))))&uint32(0x80000000) == uint32(0) {
					if uint32(tmp2)&(uint32(int32((int64(tmp1)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1)))^uint32(0x80000000))&uint32(0x80000000) != 0 {
						v18 = int32(-2147483648)
					} else {
						v18 = tmp2 - int32((int64(tmp1)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1))
					}
					v17 = v18
				} else {
					if (uint32(tmp2)^uint32(0x80000000))&uint32(int32((int64(tmp1)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1)))&uint32(0x80000000) != 0 {
						v19 = int32(silk_int32_MAX)
					} else {
						v19 = tmp2 - int32((int64(tmp1)*int64(rc_Q31)>>(int32(31)-int32(1))+int64(1))>>int32(1))
					}
					v17 = v19
				}
				v22 = (int64(v17)*int64(rc_mult2)>>(mult2Q-int32(1)) + int64(1)) >> int32(1)
			}
			tmp64 = v22
			if tmp64 > int64(silk_int32_MAX) || tmp64 < int64(int32(-2147483648)) {
				return 0
			}
			*(*OpusT_opus_int32)(unsafe.Pointer(A_QA + uintptr(k-n-int32(1))*4)) = int32(tmp64)
			goto _21
		_21:
			;
			n = n + 1
		}
		goto _1
	_1:
		;
		k = k - 1
	}
	/* Check for stability */
	if *(*OpusT_opus_int32)(unsafe.Pointer(A_QA + uintptr(k)*4)) > int32(16773022) || *(*OpusT_opus_int32)(unsafe.Pointer(A_QA + uintptr(k)*4)) < -int32(16773022) {
		return 0
	}
	/* Set RC equal to negated AR coef */
	rc_Q31 = -int32(uint32(*(*OpusT_opus_int32)(unsafe.Pointer(A_QA))) << (int32(31) - int32(QA)))
	/* Range: [ 1 : 2^30 ] */
	rc_mult1_Q30 = int32(1073741824) - int32(int64(rc_Q31)*int64(rc_Q31)>>int32(32))
	/* Update inverse gain */
	/* Range: [ 0 : 2^30 ] */
	invGain_Q30 = int32(uint32(int32(int64(invGain_Q30)*int64(rc_mult1_Q30)>>int32(32))) << int32(2))
	_ = invGain_Q30 >= int32(0)
	_ = invGain_Q30 <= int32(1)<<int32(30)
	if invGain_Q30 < int32(107374) {
		return 0
	}
	return invGain_Q30
}

// C documentation
//
//	/* For input in Q12 domain */
func Opus_silk_LPC_inverse_pred_gain_c(tls *libc.TLS, A_Q12 uintptr, order int32) (r OpusT_opus_int32) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var DC_resp OpusT_opus_int32
	var k int32
	var _ /* Atmp_QA at bp+0 */ [24]OpusT_opus_int32
	_, _ = DC_resp, k
	DC_resp = 0
	/* Increase Q domain of the AR coefficients */
	k = 0
	for {
		if !(k < order) {
			break
		}
		DC_resp = DC_resp + int32(*(*OpusT_opus_int16)(unsafe.Pointer(A_Q12 + uintptr(k)*2)))
		(*(*[24]OpusT_opus_int32)(unsafe.Pointer(bp)))[k] = int32(uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(A_Q12 + uintptr(k)*2)))) << (int32(QA) - int32(12)))
		goto _1
	_1:
		;
		k = k + 1
	}
	/* If the DC is unstable, we don't even need to do the full calculations */
	if DC_resp >= int32(4096) {
		return 0
	}
	return LPC_inverse_pred_gain_QA_c(tls, bp, order)
}

const silk_int16_MAX15 = 32767

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
//	/* Convert int32 coefficients to int16 coefs and make sure there's no wrap-around.
//	   This logic is reused in _celt_lpc(). Any bug fixes should also be applied there. */
func Opus_silk_LPC_fit(tls *libc.TLS, a_QOUT uintptr, a_QIN uintptr, QOUT int32, QIN int32, d int32) {
	var absval, chirp_Q16, maxabs OpusT_opus_int32
	var i, idx, k, v3, v4, v5, v7, v8 int32
	_, _, _, _, _, _, _, _, _, _, _ = absval, chirp_Q16, i, idx, k, maxabs, v3, v4, v5, v7, v8
	idx = 0
	/* Limit the maximum absolute value of the prediction coefficients, so that they'll fit in int16 */
	i = 0
	for {
		if !(i < int32(10)) {
			break
		}
		/* Find maximum absolute value and its index */
		maxabs = 0
		k = 0
		for {
			if !(k < d) {
				break
			}
			if *(*OpusT_opus_int32)(unsafe.Pointer(a_QIN + uintptr(k)*4)) > 0 {
				v3 = *(*OpusT_opus_int32)(unsafe.Pointer(a_QIN + uintptr(k)*4))
			} else {
				v3 = -*(*OpusT_opus_int32)(unsafe.Pointer(a_QIN + uintptr(k)*4))
			}
			absval = v3
			if absval > maxabs {
				maxabs = absval
				idx = k
			}
			goto _2
		_2:
			;
			k = k + 1
		}
		if QIN-QOUT == int32(1) {
			v3 = maxabs>>int32(1) + maxabs&int32(1)
		} else {
			v3 = (maxabs>>(QIN-QOUT-int32(1)) + int32(1)) >> int32(1)
		}
		maxabs = v3
		if maxabs > int32(silk_int16_MAX15) {
			/* Reduce magnitude of prediction coefficients */
			if maxabs < int32(163838) {
				v3 = maxabs
			} else {
				v3 = int32(163838)
			}
			maxabs = v3 /* ( silk_int32_MAX >> 14 ) + silk_int16_MAX = 163838 */
			chirp_Q16 = int32(65470) - int32(uint32(maxabs-int32(silk_int16_MAX15))<<int32(14))/(maxabs*(idx+int32(1))>>int32(2))
			Opus_silk_bwexpander_32(tls, a_QIN, d, chirp_Q16)
		} else {
			break
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	if i == int32(10) {
		/* Reached the last iteration, clip the coefficients */
		k = 0
		for {
			if !(k < d) {
				break
			}
			if QIN-QOUT == int32(1) {
				v4 = *(*OpusT_opus_int32)(unsafe.Pointer(a_QIN + uintptr(k)*4))>>int32(1) + *(*OpusT_opus_int32)(unsafe.Pointer(a_QIN + uintptr(k)*4))&int32(1)
			} else {
				v4 = (*(*OpusT_opus_int32)(unsafe.Pointer(a_QIN + uintptr(k)*4))>>(QIN-QOUT-int32(1)) + int32(1)) >> int32(1)
			}
			if v4 > int32(silk_int16_MAX15) {
				v3 = int32(silk_int16_MAX15)
			} else {
				if QIN-QOUT == int32(1) {
					v7 = *(*OpusT_opus_int32)(unsafe.Pointer(a_QIN + uintptr(k)*4))>>int32(1) + *(*OpusT_opus_int32)(unsafe.Pointer(a_QIN + uintptr(k)*4))&int32(1)
				} else {
					v7 = (*(*OpusT_opus_int32)(unsafe.Pointer(a_QIN + uintptr(k)*4))>>(QIN-QOUT-int32(1)) + int32(1)) >> int32(1)
				}
				if v7 < int32(int16(-32768)) {
					v5 = int32(int16(-32768))
				} else {
					if QIN-QOUT == int32(1) {
						v8 = *(*OpusT_opus_int32)(unsafe.Pointer(a_QIN + uintptr(k)*4))>>int32(1) + *(*OpusT_opus_int32)(unsafe.Pointer(a_QIN + uintptr(k)*4))&int32(1)
					} else {
						v8 = (*(*OpusT_opus_int32)(unsafe.Pointer(a_QIN + uintptr(k)*4))>>(QIN-QOUT-int32(1)) + int32(1)) >> int32(1)
					}
					v5 = v8
				}
				v3 = v5
			}
			*(*OpusT_opus_int16)(unsafe.Pointer(a_QOUT + uintptr(k)*2)) = int16(v3)
			*(*OpusT_opus_int32)(unsafe.Pointer(a_QIN + uintptr(k)*4)) = int32(uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(a_QOUT + uintptr(k)*2)))) << (QIN - QOUT))
			goto _6
		_6:
			;
			k = k + 1
		}
	} else {
		k = 0
		for {
			if !(k < d) {
				break
			}
			if QIN-QOUT == int32(1) {
				v3 = *(*OpusT_opus_int32)(unsafe.Pointer(a_QIN + uintptr(k)*4))>>int32(1) + *(*OpusT_opus_int32)(unsafe.Pointer(a_QIN + uintptr(k)*4))&int32(1)
			} else {
				v3 = (*(*OpusT_opus_int32)(unsafe.Pointer(a_QIN + uintptr(k)*4))>>(QIN-QOUT-int32(1)) + int32(1)) >> int32(1)
			}
			*(*OpusT_opus_int16)(unsafe.Pointer(a_QOUT + uintptr(k)*2)) = int16(v3)
			goto _12
		_12:
			;
			k = k + 1
		}
	}
}

const MAX_PREDICTION_POWER_GAIN2 = "1e4f"
const silk_int16_MAX16 = 0x7FFF
const QA1 = 16

// C documentation
//
//	/* helper function for NLSF2A(..) */
func silk_NLSF2A_find_poly(tls *libc.TLS, out uintptr, cLSF uintptr, dd int32) {
	var ftmp OpusT_opus_int32
	var k, n int32
	_, _, _ = ftmp, k, n
	*(*OpusT_opus_int32)(unsafe.Pointer(out)) = int32(uint32(int32(1)) << int32(QA1))
	*(*OpusT_opus_int32)(unsafe.Pointer(out + 1*4)) = -*(*OpusT_opus_int32)(unsafe.Pointer(cLSF))
	k = int32(1)
	for {
		if !(k < dd) {
			break
		}
		ftmp = *(*OpusT_opus_int32)(unsafe.Pointer(cLSF + uintptr(int32(2)*k)*4)) /* QA*/
		*(*OpusT_opus_int32)(unsafe.Pointer(out + uintptr(k+int32(1))*4)) = int32(uint32(*(*OpusT_opus_int32)(unsafe.Pointer(out + uintptr(k-int32(1))*4)))<<int32(1)) - int32((int64(ftmp)*int64(*(*OpusT_opus_int32)(unsafe.Pointer(out + uintptr(k)*4)))>>(int32(QA1)-int32(1))+int64(1))>>int32(1))
		n = k
		for {
			if !(n > int32(1)) {
				break
			}
			*(*OpusT_opus_int32)(unsafe.Pointer(out + uintptr(n)*4)) += *(*OpusT_opus_int32)(unsafe.Pointer(out + uintptr(n-int32(2))*4)) - int32((int64(ftmp)*int64(*(*OpusT_opus_int32)(unsafe.Pointer(out + uintptr(n-int32(1))*4)))>>(int32(QA1)-int32(1))+int64(1))>>int32(1))
			goto _2
		_2:
			;
			n = n - 1
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(out + 1*4)) -= ftmp
		goto _1
	_1:
		;
		k = k + 1
	}
}

// C documentation
//
//	/* compute whitening filter coefficients from normalized line spectral frequencies */
