// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

package opuscc

import (
	"reflect"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

var _ reflect.Type
var _ unsafe.Pointer

func Opus_silk_NLSF2A(tls *libc.TLS, a_Q12 *OpusT_opus_int16, NLSF *OpusT_opus_int16, d int32, arch int32) {
	a := unsafe.Slice(a_Q12, d)
	nlsf := unsafe.Slice(NLSF, d)
	var Ptmp, Qtmp, cos_val, delta, f_frac, f_int OpusT_opus_int32
	var dd, i, k int32
	var ordering []uint8
	var cos_LSF_QA [SILK_MAX_ORDER_LPC]OpusT_opus_int32
	var P [SILK_MAX_ORDER_LPC/2 + 1]OpusT_opus_int32
	var Q [SILK_MAX_ORDER_LPC/2 + 1]OpusT_opus_int32
	var a32_QA1 [SILK_MAX_ORDER_LPC]OpusT_opus_int32
	if !(d == int32(10) || d == int32(16)) {
		Opus_celt_fatal(tls, __ccgo_ts+7246, __ccgo_ts+7279, int32(89))
	}
	/* convert LSFs to 2*cos(LSF), using piecewise linear curve from table */
	if d == int32(16) {
		ordering = ordering16[:]
	} else {
		ordering = ordering10[:]
	}
	k = 0
	for {
		if !(k < d) {
			break
		}
		_ = int32(nlsf[k]) >= int32(0)
		/* f_int on a scale 0-127 (rounded down) */
		f_int = int32(nlsf[k]) >> (int32(15) - int32(7))
		/* f_frac, range: 0..255 */
		f_frac = int32(nlsf[k]) - int32(uint32(f_int)<<(int32(15)-int32(7)))
		_ = f_int >= int32(0)
		_ = f_int < int32(LSF_COS_TAB_SZ_FIX)
		/* Read start and end value from table */
		cos_val = int32(Opus_silk_LSFCosTab_FIX_Q12[f_int])                  /* Q12 */
		delta = int32(Opus_silk_LSFCosTab_FIX_Q12[f_int+int32(1)]) - cos_val /* Q12, with a range of 0..200 */
		/* Linear interpolation */
		cos_LSF_QA[ordering[k]] = ((int32(uint32(cos_val)<<int32(8))+delta*f_frac)>>(int32(20)-int32(QA1)-int32(1)) + int32(1)) >> int32(1) /* QA */
		k = k + 1
	}
	dd = d >> int32(1)
	/* generate even and odd polynomials using convolution */
	silk_NLSF2A_find_poly(tls, &P[0], &cos_LSF_QA[0], dd)
	silk_NLSF2A_find_poly(tls, &Q[0], &cos_LSF_QA[1], dd)
	/* convert even and odd polynomials to opus_int32 Q12 filter coefs */
	k = 0
	for {
		if !(k < dd) {
			break
		}
		Ptmp = P[k+int32(1)] + P[k]
		Qtmp = Q[k+int32(1)] - Q[k]
		/* the Ptmp and Qtmp values at this stage need to fit in int32 */
		a32_QA1[k] = -Qtmp - Ptmp           /* QA+1 */
		a32_QA1[d-k-int32(1)] = Qtmp - Ptmp /* QA+1 */
		k = k + 1
	}
	/* Convert int32 coefficients to Q12 int16 coefs */
	Opus_silk_LPC_fit(tls, a_Q12, &a32_QA1[0], int32(12), int32(QA1)+int32(1), d)
	i = 0
	for {
		if !(Opus_silk_LPC_inverse_pred_gain_c(tls, a_Q12, d) == 0 && i < int32(MAX_LPC_STABILIZE_ITERATIONS)) {
			break
		}
		/* Prediction coefficients are (too close to) unstable; apply bandwidth expansion   */
		/* on the unscaled coefficients, convert to Q12 and measure again                   */
		Opus_silk_bwexpander_32(tls, &a32_QA1[0], d, int32(65536)-int32(uint32(int32(2))<<i))
		k = 0
		for {
			if !(k < d) {
				break
			}
			a[k] = int16((a32_QA1[k]>>(int32(QA1)+int32(1)-int32(12)-int32(1)) + int32(1)) >> int32(1)) /* QA+1 -> Q12 */
			k = k + 1
		}
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
func Opus_silk_NLSF_stabilize(tls *libc.TLS, NLSF_Q15 *OpusT_opus_int16, NDeltaMin_Q15 *OpusT_opus_int16, L int32) {
	x, delta := unsafe.Slice(NLSF_Q15, int(L)), unsafe.Slice(NDeltaMin_Q15, int(L)+1)
	for loops := 0; loops < MAX_LOOPS; loops++ {
		gap, at := int32(x[0])-int32(delta[0]), 0
		for i := 1; i < len(x); i++ {
			d := int32(x[i]) - (int32(x[i-1]) + int32(delta[i]))
			if d < gap {
				gap, at = d, i
			}
		}
		d := int32(1<<15) - (int32(x[L-1]) + int32(delta[L]))
		if d < gap {
			gap, at = d, int(L)
		}
		if gap >= 0 {
			return
		}
		if at == 0 {
			x[0] = delta[0]
		} else if at == int(L) {
			x[L-1] = int16((1 << 15) - int32(delta[L]))
		} else {
			var lower int32
			for k := 0; k < at; k++ {
				lower += int32(delta[k])
			}
			lower += int32(delta[at]) >> 1
			upper := int32(1 << 15)
			for k := int(L); k > at; k-- {
				upper -= int32(delta[k])
			}
			upper -= int32(delta[at]) >> 1
			sum := int32(x[at-1]) + int32(x[at])
			center := int16(min(max((sum>>1)+(sum&1), min(lower, upper)), max(lower, upper)))
			x[at-1] = int16(int32(center) - (int32(delta[at]) >> 1))
			x[at] = int16(int32(x[at-1]) + int32(delta[at]))
		}
	}
	// C's bounded-iteration fallback: sort, forward saturated spacing, then backward spacing.
	Opus_silk_insertion_sort_increasing_all_values_int16(tls, NLSF_Q15, L)
	x[0] = max(x[0], delta[0])
	for i := 1; i < len(x); i++ {
		bound := int16(min(max(int32(x[i-1])+int32(delta[i]), -32768), 32767))
		x[i] = max(x[i], bound)
	}
	x[L-1] = int16(min(int32(x[L-1]), (1<<15)-int32(delta[L])))
	for i := int(L) - 2; i >= 0; i-- {
		x[i] = int16(min(int32(x[i]), int32(x[i+1])-int32(delta[i+1])))
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
func Opus_silk_NLSF_VQ_weights_laroia(tls *libc.TLS, pNLSFW_Q_OUT *OpusT_opus_int16, pNLSF_Q15 *OpusT_opus_int16, D int32) {
	if !(D > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+7296, __ccgo_ts+7320, int32(51))
	}
	if !(D&int32(1) == int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+7353, __ccgo_ts+7320, int32(52))
	}
	input := unsafe.Slice(pNLSF_Q15, int(D))
	output := unsafe.Slice(pNLSFW_Q_OUT, int(D))
	const numerator = int32(1) << (15 + NLSF_W_Q)
	previous := numerator / max(int32(input[0]), 1)
	for i := range output {
		gap := int32(1<<15) - int32(input[i])
		if i+1 < len(input) {
			gap = int32(input[i+1]) - int32(input[i])
		}
		next := numerator / max(gap, 1)
		// Read the next gap before writing, preserving C's in-place behavior.
		output[i] = int16(min(previous+next, 32767))
		previous = next
	}
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
