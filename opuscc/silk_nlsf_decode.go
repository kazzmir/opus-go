// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

package opuscc

import (
	"reflect"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

var _ reflect.Type
var _ unsafe.Pointer

func Opus_silk_NLSF_decode(tls *libc.TLS, pNLSF_Q15 uintptr, NLSFIndices uintptr, psNLSF_CB uintptr) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var NLSF_Q15_tmp OpusT_opus_int32
	var i, v2, v3 int32
	var pCB_Wght_Q9, pCB_element uintptr
	var _ /* ec_ix at bp+16 */ [16]OpusT_opus_int16
	var _ /* pred_Q8 at bp+0 */ [16]OpusT_opus_uint8
	var _ /* res_Q10 at bp+48 */ [16]OpusT_opus_int16
	_, _, _, _, _, _ = NLSF_Q15_tmp, i, pCB_Wght_Q9, pCB_element, v2, v3
	/* Unpack entropy table indices and predictor for current CB1 index */
	Opus_silk_NLSF_unpack(tls, bp+16, bp, psNLSF_CB, int32(*(*OpusT_opus_int8)(unsafe.Pointer(NLSFIndices))))
	/* Predictive residual dequantizer */
	silk_NLSF_residual_dequant(tls, bp+48, NLSFIndices+1, bp, int32((*OpusT_silk_NLSF_CB_struct)(unsafe.Pointer(psNLSF_CB)).FquantStepSize_Q16), (*OpusT_silk_NLSF_CB_struct)(unsafe.Pointer(psNLSF_CB)).Forder)
	/* Apply inverse square-rooted weights to first stage and add to output */
	pCB_element = (*OpusT_silk_NLSF_CB_struct)(unsafe.Pointer(psNLSF_CB)).FCB1_NLSF_Q8 + uintptr(int32(*(*OpusT_opus_int8)(unsafe.Pointer(NLSFIndices)))*int32((*OpusT_silk_NLSF_CB_struct)(unsafe.Pointer(psNLSF_CB)).Forder))
	pCB_Wght_Q9 = (*OpusT_silk_NLSF_CB_struct)(unsafe.Pointer(psNLSF_CB)).FCB1_Wght_Q9 + uintptr(int32(*(*OpusT_opus_int8)(unsafe.Pointer(NLSFIndices)))*int32((*OpusT_silk_NLSF_CB_struct)(unsafe.Pointer(psNLSF_CB)).Forder))*2
	i = 0
	for {
		if !(i < int32((*OpusT_silk_NLSF_CB_struct)(unsafe.Pointer(psNLSF_CB)).Forder)) {
			break
		}
		NLSF_Q15_tmp = int32(uint32(int32((*(*[16]OpusT_opus_int16)(unsafe.Pointer(bp + 48)))[i]))<<int32(14))/int32(*(*OpusT_opus_int16)(unsafe.Pointer(pCB_Wght_Q9 + uintptr(i)*2))) + int32(uint32(uint16(int16(*(*OpusT_opus_uint8)(unsafe.Pointer(pCB_element + uintptr(i))))))<<int32(7))
		if NLSF_Q15_tmp > int32(32767) {
			v2 = int32(32767)
		} else {
			if NLSF_Q15_tmp < 0 {
				v3 = 0
			} else {
				v3 = NLSF_Q15_tmp
			}
			v2 = v3
		}
		*(*OpusT_opus_int16)(unsafe.Pointer(pNLSF_Q15 + uintptr(i)*2)) = int16(v2)
		i = i + 1
	}
	/* NLSF stabilization */
	Opus_silk_NLSF_stabilize(tls, pNLSF_Q15, (*OpusT_silk_NLSF_CB_struct)(unsafe.Pointer(psNLSF_CB)).FdeltaMin_Q15, int32((*OpusT_silk_NLSF_CB_struct)(unsafe.Pointer(psNLSF_CB)).Forder))
}

const NB_ATT = 2
const silk_int16_MAX7 = 32767

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

var HARM_ATT_Q15 = [2]OpusT_opus_int16{
	0: int16(32440),
	1: int16(31130),
} /* 0.99, 0.95 */
var PLC_RAND_ATTENUATE_V_Q15 = [2]OpusT_opus_int16{
	0: int16(31130),
	1: int16(26214),
} /* 0.95, 0.8 */
var PLC_RAND_ATTENUATE_UV_Q15 = [2]OpusT_opus_int16{
	0: int16(32440),
	1: int16(29491),
}
