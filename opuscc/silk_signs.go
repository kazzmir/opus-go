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

func Opus_silk_encode_signs(tls *libc.TLS, psRangeEnc uintptr, pulses uintptr, length int32, signalType int32, quantOffsetType int32, sum_pulses uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i, j, p, v2 int32
	var icdf_ptr, q_ptr uintptr
	var _ /* icdf at bp+0 */ [2]OpusT_opus_uint8
	_, _, _, _, _, _ = i, icdf_ptr, j, p, q_ptr, v2
	(*(*[2]OpusT_opus_uint8)(unsafe.Pointer(bp)))[int32(1)] = uint8(0)
	q_ptr = pulses
	i = int32(int16(int32(7))) * int32(int16(quantOffsetType+int32(uint32(signalType)<<int32(1))))
	icdf_ptr = uintptr(unsafe.Pointer(&Opus_silk_sign_iCDF)) + uintptr(i)
	length = (length + int32(SHELL_CODEC_FRAME_LENGTH)/int32(2)) >> int32(LOG2_SHELL_CODEC_FRAME_LENGTH)
	i = 0
	for {
		if !(i < length) {
			break
		}
		p = *(*int32)(unsafe.Pointer(sum_pulses + uintptr(i)*4))
		if p > 0 {
			if p&int32(0x1F) < int32(6) {
				v2 = p & int32(0x1F)
			} else {
				v2 = int32(6)
			}
			(*(*[2]OpusT_opus_uint8)(unsafe.Pointer(bp)))[0] = *(*OpusT_opus_uint8)(unsafe.Pointer(icdf_ptr + uintptr(v2)))
			j = 0
			for {
				if !(j < int32(SHELL_CODEC_FRAME_LENGTH)) {
					break
				}
				if int32(*(*OpusT_opus_int8)(unsafe.Pointer(q_ptr + uintptr(j)))) != 0 {
					Opus_ec_enc_icdf(tls, psRangeEnc, int32(*(*OpusT_opus_int8)(unsafe.Pointer(q_ptr + uintptr(j))))>>int32(15)+int32(1), bp, uint32(8))
				}
				goto _3
			_3:
				;
				j = j + 1
			}
		}
		q_ptr = q_ptr + uintptr(SHELL_CODEC_FRAME_LENGTH)
		goto _1
	_1:
		;
		i = i + 1
	}
}

// C documentation
//
//	/* Decodes signs of excitation */
func Opus_silk_decode_signs(tls *libc.TLS, psRangeDec uintptr, pulses uintptr, length int32, signalType int32, quantOffsetType int32, sum_pulses uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i, j, p, v2 int32
	var icdf_ptr, q_ptr, v4 uintptr
	var _ /* icdf at bp+0 */ [2]OpusT_opus_uint8
	_, _, _, _, _, _, _ = i, icdf_ptr, j, p, q_ptr, v2, v4
	(*(*[2]OpusT_opus_uint8)(unsafe.Pointer(bp)))[int32(1)] = uint8(0)
	q_ptr = pulses
	i = int32(int16(int32(7))) * int32(int16(quantOffsetType+int32(uint32(signalType)<<int32(1))))
	icdf_ptr = uintptr(unsafe.Pointer(&Opus_silk_sign_iCDF)) + uintptr(i)
	length = (length + int32(SHELL_CODEC_FRAME_LENGTH)/int32(2)) >> int32(LOG2_SHELL_CODEC_FRAME_LENGTH)
	i = 0
	for {
		if !(i < length) {
			break
		}
		p = *(*int32)(unsafe.Pointer(sum_pulses + uintptr(i)*4))
		if p > 0 {
			if p&int32(0x1F) < int32(6) {
				v2 = p & int32(0x1F)
			} else {
				v2 = int32(6)
			}
			(*(*[2]OpusT_opus_uint8)(unsafe.Pointer(bp)))[0] = *(*OpusT_opus_uint8)(unsafe.Pointer(icdf_ptr + uintptr(v2)))
			j = 0
			for {
				if !(j < int32(SHELL_CODEC_FRAME_LENGTH)) {
					break
				}
				if int32(*(*OpusT_opus_int16)(unsafe.Pointer(q_ptr + uintptr(j)*2))) > 0 {
					/* attach sign */
					/* implementation with shift, subtraction, multiplication */
					v4 = q_ptr + uintptr(j)*2
					*(*OpusT_opus_int16)(unsafe.Pointer(v4)) = OpusT_opus_int16(int32(*(*OpusT_opus_int16)(unsafe.Pointer(v4))) * (int32(uint32(Opus_ec_dec_icdf(tls, psRangeDec, bp, uint32(8)))<<int32(1)) - int32(1)))
				}
				goto _3
			_3:
				;
				j = j + 1
			}
		}
		q_ptr = q_ptr + uintptr(SHELL_CODEC_FRAME_LENGTH)*2
		goto _1
	_1:
		;
		i = i + 1
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

// C documentation
//
//	/************************/
//	/* Reset Decoder State  */
//	/************************/
