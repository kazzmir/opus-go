// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

package opuscc

import (
	"reflect"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

var _ reflect.Type
var _ unsafe.Pointer

func Opus_silk_NLSF2A(tls *libc.TLS, a_Q12 uintptr, NLSF uintptr, d int32, arch int32) {
	bp := tls.Alloc(304)
	defer tls.Free(304)
	var Ptmp, Qtmp, cos_val, delta, f_frac, f_int OpusT_opus_int32
	var dd, i, k int32
	var ordering, v1 uintptr
	var _ /* P at bp+96 */ [13]OpusT_opus_int32
	var _ /* Q at bp+148 */ [13]OpusT_opus_int32
	var _ /* a32_QA1 at bp+200 */ [24]OpusT_opus_int32
	var _ /* cos_LSF_QA at bp+0 */ [24]OpusT_opus_int32
	_, _, _, _, _, _, _, _, _, _, _ = Ptmp, Qtmp, cos_val, dd, delta, f_frac, f_int, i, k, ordering, v1
	_ = true
	if !(d == int32(10) || d == int32(16)) {
		Opus_celt_fatal(tls, __ccgo_ts+7246, __ccgo_ts+7279, int32(89))
	}
	/* convert LSFs to 2*cos(LSF), using piecewise linear curve from table */
	if d == int32(16) {
		v1 = uintptr(unsafe.Pointer(&ordering16))
	} else {
		v1 = uintptr(unsafe.Pointer(&ordering10))
	}
	ordering = v1
	k = 0
	for {
		if !(k < d) {
			break
		}
		_ = int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF + uintptr(k)*2))) >= int32(0)
		/* f_int on a scale 0-127 (rounded down) */
		f_int = int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF + uintptr(k)*2))) >> (int32(15) - int32(7))
		/* f_frac, range: 0..255 */
		f_frac = int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF + uintptr(k)*2))) - int32(uint32(f_int)<<(int32(15)-int32(7)))
		_ = f_int >= int32(0)
		_ = f_int < int32(LSF_COS_TAB_SZ_FIX)
		/* Read start and end value from table */
		cos_val = int32(Opus_silk_LSFCosTab_FIX_Q12[f_int])                  /* Q12 */
		delta = int32(Opus_silk_LSFCosTab_FIX_Q12[f_int+int32(1)]) - cos_val /* Q12, with a range of 0..200 */
		/* Linear interpolation */
		(*(*[24]OpusT_opus_int32)(unsafe.Pointer(bp)))[*(*uint8)(unsafe.Pointer(ordering + uintptr(k)))] = ((int32(uint32(cos_val)<<int32(8))+delta*f_frac)>>(int32(20)-int32(QA1)-int32(1)) + int32(1)) >> int32(1) /* QA */
		goto _2
	_2:
		;
		k = k + 1
	}
	dd = d >> int32(1)
	/* generate even and odd polynomials using convolution */
	silk_NLSF2A_find_poly(tls, bp+96, bp, dd)
	silk_NLSF2A_find_poly(tls, bp+148, bp+1*4, dd)
	/* convert even and odd polynomials to opus_int32 Q12 filter coefs */
	k = 0
	for {
		if !(k < dd) {
			break
		}
		Ptmp = (*(*[13]OpusT_opus_int32)(unsafe.Pointer(bp + 96)))[k+int32(1)] + (*(*[13]OpusT_opus_int32)(unsafe.Pointer(bp + 96)))[k]
		Qtmp = (*(*[13]OpusT_opus_int32)(unsafe.Pointer(bp + 148)))[k+int32(1)] - (*(*[13]OpusT_opus_int32)(unsafe.Pointer(bp + 148)))[k]
		/* the Ptmp and Qtmp values at this stage need to fit in int32 */
		(*(*[24]OpusT_opus_int32)(unsafe.Pointer(bp + 200)))[k] = -Qtmp - Ptmp           /* QA+1 */
		(*(*[24]OpusT_opus_int32)(unsafe.Pointer(bp + 200)))[d-k-int32(1)] = Qtmp - Ptmp /* QA+1 */
		goto _3
	_3:
		;
		k = k + 1
	}
	/* Convert int32 coefficients to Q12 int16 coefs */
	Opus_silk_LPC_fit(tls, a_Q12, bp+200, int32(12), int32(QA1)+int32(1), d)
	i = 0
	for {
		_ = arch
		if !(Opus_silk_LPC_inverse_pred_gain_c(tls, a_Q12, d) == 0 && i < int32(MAX_LPC_STABILIZE_ITERATIONS)) {
			break
		}
		/* Prediction coefficients are (too close to) unstable; apply bandwidth expansion   */
		/* on the unscaled coefficients, convert to Q12 and measure again                   */
		Opus_silk_bwexpander_32(tls, bp+200, d, int32(65536)-int32(uint32(int32(2))<<i))
		k = 0
		for {
			if !(k < d) {
				break
			}
			*(*OpusT_opus_int16)(unsafe.Pointer(a_Q12 + uintptr(k)*2)) = int16(((*(*[24]OpusT_opus_int32)(unsafe.Pointer(bp + 200)))[k]>>(int32(QA1)+int32(1)-int32(12)-int32(1)) + int32(1)) >> int32(1)) /* QA+1 -> Q12 */
			goto _5
		_5:
			;
			k = k + 1
		}
		goto _4
	_4:
		;
		i = i + 1
	}
}

/*
This ordering was found to maximize quality. It improves numerical accuracy of

	silk_NLSF2A_find_poly() compared to "standard" ordering.
*/
var ordering16 = [16]uint8{
	1:  uint8(15),
	2:  uint8(8),
	3:  uint8(7),
	4:  uint8(4),
	5:  uint8(11),
	6:  uint8(12),
	7:  uint8(3),
	8:  uint8(2),
	9:  uint8(13),
	10: uint8(10),
	11: uint8(5),
	12: uint8(6),
	13: uint8(9),
	14: uint8(14),
	15: uint8(1),
}

var ordering10 = [10]uint8{
	1: uint8(9),
	2: uint8(6),
	3: uint8(3),
	4: uint8(4),
	5: uint8(5),
	6: uint8(8),
	7: uint8(1),
	8: uint8(2),
	9: uint8(7),
}

const MAX_LOOPS = 20
const silk_int16_MAX17 = 32767

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

/* Constant Definitions */

// C documentation
//
//	/* NLSF stabilizer, for a single input data vector */
func Opus_silk_NLSF_stabilize(tls *libc.TLS, NLSF_Q15 uintptr, NDeltaMin_Q15 uintptr, L int32) {
	var I, i, k, loops, v10, v5, v6, v7, v8, v9 int32
	var center_freq_Q15 OpusT_opus_int16
	var diff_Q15, max_center_Q15, min_center_Q15, min_diff_Q15 OpusT_opus_int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = I, center_freq_Q15, diff_Q15, i, k, loops, max_center_Q15, min_center_Q15, min_diff_Q15, v10, v5, v6, v7, v8, v9
	I = 0
	/* This is necessary to ensure an output within range of a opus_int16 */
	_ = int32(*(*OpusT_opus_int16)(unsafe.Pointer(NDeltaMin_Q15 + uintptr(L)*2))) >= int32(1)
	loops = 0
	for {
		if !(loops < int32(MAX_LOOPS)) {
			break
		}
		/**************************/
		/* Find smallest distance */
		/**************************/
		/* First element */
		min_diff_Q15 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer(NDeltaMin_Q15)))
		I = 0
		/* Middle elements */
		i = int32(1)
		for {
			if !(i <= L-int32(1)) {
				break
			}
			diff_Q15 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(i)*2))) - (int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(i-int32(1))*2))) + int32(*(*OpusT_opus_int16)(unsafe.Pointer(NDeltaMin_Q15 + uintptr(i)*2))))
			if diff_Q15 < min_diff_Q15 {
				min_diff_Q15 = diff_Q15
				I = i
			}
			goto _2
		_2:
			;
			i = i + 1
		}
		/* Last element */
		diff_Q15 = int32(1)<<int32(15) - (int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(L-int32(1))*2))) + int32(*(*OpusT_opus_int16)(unsafe.Pointer(NDeltaMin_Q15 + uintptr(L)*2))))
		if diff_Q15 < min_diff_Q15 {
			min_diff_Q15 = diff_Q15
			I = L
		}
		/***************************************************/
		/* Now check if the smallest distance non-negative */
		/***************************************************/
		if min_diff_Q15 >= 0 {
			return
		}
		if I == 0 {
			/* Move away from lower limit */
			*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15)) = *(*OpusT_opus_int16)(unsafe.Pointer(NDeltaMin_Q15))
		} else {
			if I == L {
				/* Move away from higher limit */
				*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(L-int32(1))*2)) = int16(int32(1)<<int32(15) - int32(*(*OpusT_opus_int16)(unsafe.Pointer(NDeltaMin_Q15 + uintptr(L)*2))))
			} else {
				/* Find the lower extreme for the location of the current center frequency */
				min_center_Q15 = 0
				k = 0
				for {
					if !(k < I) {
						break
					}
					min_center_Q15 = min_center_Q15 + int32(*(*OpusT_opus_int16)(unsafe.Pointer(NDeltaMin_Q15 + uintptr(k)*2)))
					goto _3
				_3:
					;
					k = k + 1
				}
				min_center_Q15 = min_center_Q15 + int32(*(*OpusT_opus_int16)(unsafe.Pointer(NDeltaMin_Q15 + uintptr(I)*2)))>>int32(1)
				/* Find the upper extreme for the location of the current center frequency */
				max_center_Q15 = int32(1) << int32(15)
				k = L
				for {
					if !(k > I) {
						break
					}
					max_center_Q15 = max_center_Q15 - int32(*(*OpusT_opus_int16)(unsafe.Pointer(NDeltaMin_Q15 + uintptr(k)*2)))
					goto _4
				_4:
					;
					k = k - 1
				}
				max_center_Q15 = max_center_Q15 - int32(*(*OpusT_opus_int16)(unsafe.Pointer(NDeltaMin_Q15 + uintptr(I)*2)))>>int32(1)
				/* Move apart, sorted by value, keeping the same center frequency */
				if min_center_Q15 > max_center_Q15 {
					if (int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I-int32(1))*2)))+int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I)*2))))>>int32(1)+(int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I-int32(1))*2)))+int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I)*2))))&int32(1) > min_center_Q15 {
						v6 = min_center_Q15
					} else {
						if (int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I-int32(1))*2)))+int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I)*2))))>>int32(1)+(int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I-int32(1))*2)))+int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I)*2))))&int32(1) < max_center_Q15 {
							v7 = max_center_Q15
						} else {
							v7 = (int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I-int32(1))*2)))+int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I)*2))))>>int32(1) + (int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I-int32(1))*2)))+int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I)*2))))&int32(1)
						}
						v6 = v7
					}
					v5 = v6
				} else {
					if (int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I-int32(1))*2)))+int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I)*2))))>>int32(1)+(int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I-int32(1))*2)))+int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I)*2))))&int32(1) > max_center_Q15 {
						v8 = max_center_Q15
					} else {
						if (int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I-int32(1))*2)))+int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I)*2))))>>int32(1)+(int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I-int32(1))*2)))+int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I)*2))))&int32(1) < min_center_Q15 {
							v9 = min_center_Q15
						} else {
							v9 = (int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I-int32(1))*2)))+int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I)*2))))>>int32(1) + (int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I-int32(1))*2)))+int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I)*2))))&int32(1)
						}
						v8 = v9
					}
					v5 = v8
				}
				center_freq_Q15 = int16(v5)
				*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I-int32(1))*2)) = int16(int32(center_freq_Q15) - int32(*(*OpusT_opus_int16)(unsafe.Pointer(NDeltaMin_Q15 + uintptr(I)*2)))>>int32(1))
				*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I)*2)) = int16(int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(I-int32(1))*2))) + int32(*(*OpusT_opus_int16)(unsafe.Pointer(NDeltaMin_Q15 + uintptr(I)*2))))
			}
		}
		goto _1
	_1:
		;
		loops = loops + 1
	}
	/* Safe and simple fall back method, which is less ideal than the above */
	if loops == int32(MAX_LOOPS) {
		/* Insertion sort (fast for already almost sorted arrays):   */
		/* Best case:  O(n)   for an already sorted array            */
		/* Worst case: O(n^2) for an inversely sorted array          */
		Opus_silk_insertion_sort_increasing_all_values_int16(tls, NLSF_Q15, L)
		/* First NLSF should be no less than NDeltaMin[0] */
		v5 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15)))
		v6 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(NDeltaMin_Q15)))
		if v5 > v6 {
			v8 = v5
		} else {
			v8 = v6
		}
		v7 = v8
		goto _13
	_13:
		*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15)) = int16(v7)
		/* Keep delta_min distance between the NLSFs */
		i = int32(1)
		for {
			if !(i < L) {
				break
			}
			if int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(i-int32(1))*2)))+int32(*(*OpusT_opus_int16)(unsafe.Pointer(NDeltaMin_Q15 + uintptr(i)*2))) > int32(silk_int16_MAX17) {
				v5 = int32(silk_int16_MAX17)
			} else {
				if int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(i-int32(1))*2)))+int32(*(*OpusT_opus_int16)(unsafe.Pointer(NDeltaMin_Q15 + uintptr(i)*2))) < int32(int16(-32768)) {
					v6 = int32(int16(-32768))
				} else {
					v6 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(i-int32(1))*2))) + int32(*(*OpusT_opus_int16)(unsafe.Pointer(NDeltaMin_Q15 + uintptr(i)*2)))
				}
				v5 = v6
			}
			v7 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(i)*2)))
			v8 = int32(int16(v5))
			if v7 > v8 {
				v10 = v7
			} else {
				v10 = v8
			}
			v9 = v10
			goto _21
		_21:
			*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(i)*2)) = int16(v9)
			goto _15
		_15:
			;
			i = i + 1
		}
		/* Last NLSF should be no higher than 1 - NDeltaMin[L] */
		v5 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(L-int32(1))*2)))
		v6 = int32(1)<<int32(15) - int32(*(*OpusT_opus_int16)(unsafe.Pointer(NDeltaMin_Q15 + uintptr(L)*2)))
		if v5 < v6 {
			v8 = v5
		} else {
			v8 = v6
		}
		v7 = v8
		goto _26
	_26:
		*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(L-int32(1))*2)) = int16(v7)
		/* Keep NDeltaMin distance between the NLSFs */
		i = L - int32(2)
		for {
			if !(i >= 0) {
				break
			}
			v5 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(i)*2)))
			v6 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(i+int32(1))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer(NDeltaMin_Q15 + uintptr(i+int32(1))*2)))
			if v5 < v6 {
				v8 = v5
			} else {
				v8 = v6
			}
			v7 = v8
			goto _32
		_32:
			*(*OpusT_opus_int16)(unsafe.Pointer(NLSF_Q15 + uintptr(i)*2)) = int16(v7)
			goto _28
		_28:
			;
			i = i - 1
		}
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

/*
R. Laroia, N. Phamdo and N. Farvardin, "Robust and Efficient Quantization of Speech LSP
Parameters Using Structured Vector Quantization", Proc. IEEE Int. Conf. Acoust., Speech,
Signal Processing, pp. 641-644, 1991.
*/

// C documentation
//
//	/* Laroia low complexity NLSF weights */
func Opus_silk_NLSF_VQ_weights_laroia(tls *libc.TLS, pNLSFW_Q_OUT uintptr, pNLSF_Q15 uintptr, D int32) {
	var k, v1, v2, v3, v5 int32
	var tmp1_int, tmp2_int OpusT_opus_int32
	_, _, _, _, _, _, _ = k, tmp1_int, tmp2_int, v1, v2, v3, v5
	if !(D > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+7296, __ccgo_ts+7320, int32(51))
	}
	if !(D&int32(1) == int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+7353, __ccgo_ts+7320, int32(52))
	}
	/* First value */
	v1 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(pNLSF_Q15)))
	v2 = int32(1)
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	tmp1_int = v3
	tmp1_int = int32(1) << (int32(15) + int32(NLSF_W_Q)) / tmp1_int
	v1 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(pNLSF_Q15 + 1*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer(pNLSF_Q15)))
	v2 = int32(1)
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _9
_9:
	tmp2_int = v3
	tmp2_int = int32(1) << (int32(15) + int32(NLSF_W_Q)) / tmp2_int
	v1 = tmp1_int + tmp2_int
	v2 = int32(silk_int16_MAX17)
	if v1 < v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _14
_14:
	*(*OpusT_opus_int16)(unsafe.Pointer(pNLSFW_Q_OUT)) = int16(v3)
	_ = int32(*(*OpusT_opus_int16)(unsafe.Pointer(pNLSFW_Q_OUT))) > int32(0)
	/* Main loop */
	k = int32(1)
	for {
		if !(k < D-int32(1)) {
			break
		}
		v1 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(pNLSF_Q15 + uintptr(k+int32(1))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer(pNLSF_Q15 + uintptr(k)*2)))
		v2 = int32(1)
		if v1 > v2 {
			v5 = v1
		} else {
			v5 = v2
		}
		v3 = v5
		goto _20
	_20:
		tmp1_int = v3
		tmp1_int = int32(1) << (int32(15) + int32(NLSF_W_Q)) / tmp1_int
		v1 = tmp1_int + tmp2_int
		v2 = int32(silk_int16_MAX17)
		if v1 < v2 {
			v5 = v1
		} else {
			v5 = v2
		}
		v3 = v5
		goto _25
	_25:
		*(*OpusT_opus_int16)(unsafe.Pointer(pNLSFW_Q_OUT + uintptr(k)*2)) = int16(v3)
		_ = int32(*(*OpusT_opus_int16)(unsafe.Pointer(pNLSFW_Q_OUT + uintptr(k)*2))) > int32(0)
		v1 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(pNLSF_Q15 + uintptr(k+int32(2))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer(pNLSF_Q15 + uintptr(k+int32(1))*2)))
		v2 = int32(1)
		if v1 > v2 {
			v5 = v1
		} else {
			v5 = v2
		}
		v3 = v5
		goto _30
	_30:
		tmp2_int = v3
		tmp2_int = int32(1) << (int32(15) + int32(NLSF_W_Q)) / tmp2_int
		v1 = tmp1_int + tmp2_int
		v2 = int32(silk_int16_MAX17)
		if v1 < v2 {
			v5 = v1
		} else {
			v5 = v2
		}
		v3 = v5
		goto _35
	_35:
		*(*OpusT_opus_int16)(unsafe.Pointer(pNLSFW_Q_OUT + uintptr(k+int32(1))*2)) = int16(v3)
		_ = int32(*(*OpusT_opus_int16)(unsafe.Pointer(pNLSFW_Q_OUT + uintptr(k+int32(1))*2))) > int32(0)
		goto _16
	_16:
		;
		k = k + int32(2)
	}
	/* Last value */
	v1 = int32(1)<<int32(15) - int32(*(*OpusT_opus_int16)(unsafe.Pointer(pNLSF_Q15 + uintptr(D-int32(1))*2)))
	v2 = int32(1)
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _40
_40:
	tmp1_int = v3
	tmp1_int = int32(1) << (int32(15) + int32(NLSF_W_Q)) / tmp1_int
	v1 = tmp1_int + tmp2_int
	v2 = int32(silk_int16_MAX17)
	if v1 < v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _45
_45:
	*(*OpusT_opus_int16)(unsafe.Pointer(pNLSFW_Q_OUT + uintptr(D-int32(1))*2)) = int16(v3)
	_ = int32(*(*OpusT_opus_int16)(unsafe.Pointer(pNLSFW_Q_OUT + uintptr(D-int32(1))*2))) > int32(0)
}

const silk_int16_MAX18 = 0x7FFF
const RESAMPLER_DOWN_ORDER_FIR0 = 18
const RESAMPLER_DOWN_ORDER_FIR1 = 24
const RESAMPLER_DOWN_ORDER_FIR2 = 36
const RESAMPLER_MAX_BATCH_SIZE_MS = 10
const RESAMPLER_MAX_FS_KHZ = 48
const RESAMPLER_ORDER_FIR_12 = 8
const USE_silk_resampler_copy = 0
const USE_silk_resampler_private_IIR_FIR = 2
const USE_silk_resampler_private_down_FIR = 3
const USE_silk_resampler_private_up2_HQ_wrapper = 1

var silk_resampler_down2_0 = int16(9872)
var silk_resampler_down2_1 = int16(int32(39809) - int32(65536))
var silk_resampler_up2_hq_0 = [3]OpusT_opus_int16{
	0: int16(1746),
	1: int16(14986),
	2: int16(int32(39083) - int32(65536)),
}
var silk_resampler_up2_hq_1 = [3]OpusT_opus_int16{
	0: int16(6854),
	1: int16(25769),
	2: int16(int32(55542) - int32(65536)),
}

// C documentation
//
//	/* Tables with delay compensation values to equalize total delay for different modes */
var delay_matrix_enc = [6][3]OpusT_opus_int8{
	0: {
		0: int8(6),
		2: int8(3),
	},
	1: {
		1: int8(7),
		2: int8(3),
	},
	2: {
		1: int8(1),
		2: int8(10),
	},
	3: {
		1: int8(2),
		2: int8(6),
	},
	4: {
		0: int8(18),
		1: int8(10),
		2: int8(12),
	},
	5: {
		2: int8(44),
	},
}

var delay_matrix_dec = [3][6]OpusT_opus_int8{
	0: {
		0: int8(4),
		2: int8(2),
	},
	1: {
		1: int8(9),
		2: int8(4),
		3: int8(7),
		4: int8(4),
		5: int8(4),
	},
	2: {
		1: int8(3),
		2: int8(12),
		3: int8(7),
		4: int8(7),
		5: int8(7),
	},
}

/* Simple way to make [8000, 12000, 16000, 24000, 48000] to [0, 1, 2, 3, 4] */

// C documentation
//
//	/* Initialize/reset the resampler state for a given pair of input/output sampling rates */
