// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

package opuscc

import (
	"math"
	"math/bits"
	"reflect"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

var _ reflect.Type
var _ unsafe.Pointer

// modeLogN keeps the table address typed, including in remaining legacy callers.
func modeLogN(m *OpusT_OpusCustomMode, band int32) int16 {
	return *(*int16)(unsafe.Add(unsafe.Pointer(m.FlogN), uintptr(band)*2))
}

func modePulseIndex(m *OpusT_OpusCustomMode, index int32) int16 {
	return *(*int16)(unsafe.Add(unsafe.Pointer(m.Fcache.Findex), uintptr(index)*2))
}

func modePulseCache(m *OpusT_OpusCustomMode, index int32) *byte {
	bits := m.Fcache.Fbits
	offset := modePulseIndex(m, index)
	return (*byte)(unsafe.Add(unsafe.Pointer(bits), int(offset)))
}

func modePulseByte(cache *byte, index int32) byte {
	return *(*byte)(unsafe.Add(unsafe.Pointer(cache), int(index)))
}

// These are the scalar rate.h searches, with C int32 wrapping and tie order.
func modeBits2Pulses(cache *byte, bits int32) int32 {
	lo, hi := int32(0), int32(*cache)
	bits--
	for i := 0; i < LOG_MAX_PSEUDO; i++ {
		mid := (lo + hi + 1) >> 1
		if int32(modePulseByte(cache, mid)) >= bits {
			hi = mid
		} else {
			lo = mid
		}
	}
	lowBits := int32(-1)
	if lo != 0 {
		lowBits = int32(modePulseByte(cache, lo))
	}
	if bits-lowBits <= int32(modePulseByte(cache, hi))-bits {
		return lo
	}
	return hi
}

func modePulses2Bits(cache *byte, pulses int32) int32 {
	if pulses == 0 {
		return 0
	}
	return int32(modePulseByte(cache, pulses)) + 1
}

func modeBand(m *OpusT_OpusCustomMode, band int32) int16 {
	return *(*int16)(unsafe.Add(unsafe.Pointer(m.FeBands), uintptr(band)*2))
}

func Opus_opus_custom_mode_create(tls *libc.TLS, Fs OpusT_opus_int32, frameSize int32) (*OpusT_OpusCustomMode, error) {
	for _, mode := range static_mode_list {
		for j := uint(0); j < 4; j++ {
			if Fs == mode.FFs && frameSize<<j == mode.FshortMdctSize*mode.FnbShortMdcts {
				return mode, nil
			}
		}
	}
	return nil, opusErrorFromCode(OPUS_BAD_ARG)
}

const EPSILON1 = 1e-15
const Q15ONE3 = 1

var log2_x_norm_coeff12 = [8]float32{
	0: float32(1),
	1: float32(0.8888888955116272),
	2: float32(0.8),
	3: float32(0.7272727489471436),
	4: float32(0.6666666865348816),
	5: float32(0.6153846383094788),
	6: float32(0.5714285969734192),
	7: float32(0.5333333611488342),
}
var log2_y_norm_coeff12 = [8]float32{
	1: float32(0.1699250042438507),
	2: float32(0.32192808389663696),
	3: float32(0.45943161845207214),
	4: float32(0.5849624872207642),
	5: float32(0.7004396915435791),
	6: float32(0.8073549270629883),
	7: float32(0.9068905711174011),
}

var trim_icdf13 = [11]uint8{
	0: uint8(126),
	1: uint8(124),
	2: uint8(119),
	3: uint8(109),
	4: uint8(87),
	5: uint8(41),
	6: uint8(19),
	7: uint8(9),
	8: uint8(4),
	9: uint8(2),
}
var spread_icdf13 = [4]uint8{
	0: uint8(25),
	1: uint8(23),
	2: uint8(2),
}
var tapset_icdf13 = [3]uint8{
	0: uint8(2),
	1: uint8(1),
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

func exp_rotation1(tls *libc.TLS, X *OpusT_celt_norm, len1 int32, stride int32, c OpusT_opus_val16, s OpusT_opus_val16) {
	if len1 <= 0 {
		return
	}
	x := unsafe.Slice(X, int(len1))
	ms := -s
	for i := int32(0); i < len1-stride; i++ {
		x1, x2 := x[i], x[i+stride]
		x[i+stride] = OpusT_opus_val32(c*x2) + OpusT_opus_val32(s*x1)
		x[i] = OpusT_opus_val32(c*x1) + OpusT_opus_val32(ms*x2)
	}
	for i := len1 - 2*stride - 1; i >= 0; i-- {
		x1, x2 := x[i], x[i+stride]
		x[i+stride] = OpusT_opus_val32(c*x2) + OpusT_opus_val32(s*x1)
		x[i] = OpusT_opus_val32(c*x1) + OpusT_opus_val32(ms*x2)
	}
}

func Opus_exp_rotation(tls *libc.TLS, X *OpusT_celt_norm, len1 int32, dir int32, stride int32, K int32, spread int32) {
	var c, gain, s, theta OpusT_opus_val16
	var factor, i, stride2 int32
	var v1, v2 OpusT_opus_uint32
	_, _, _, _, _, _, _, _, _ = c, factor, gain, i, s, stride2, theta, v1, v2
	stride2 = 0
	if int32(2)*K >= len1 || spread == SPREAD_NONE {
		return
	}
	samples := unsafe.Slice(X, int(len1))
	factor = SPREAD_FACTOR[spread-int32(1)]
	gain = OpusT_opus_val32(float32(1)*float32(len1)) / float32(len1+factor*K)
	theta = float32(float32(0.5) * OpusT_opus_val16(gain*gain))
	c = float32(libc.Xcos(tls, float64(float64(float64(0.5)*float64(3.141592653589793))*float64(theta))))
	s = float32(libc.Xcos(tls, float64(float64(float64(0.5)*float64(3.141592653589793))*float64(float32(1)-theta)))) /*  sin(theta) */
	if len1 >= int32(8)*stride {
		stride2 = int32(1)
		/* This is just a simple (equivalent) way of computing sqrt(len/stride) with rounding.
		   It's basically incrementing long as (stride2+0.5)^2 < len/stride. */
		for (stride2*stride2+stride2)*stride+stride>>int32(2) < len1 {
			stride2 = stride2 + 1
		}
	}
	/*NOTE: As a minor optimization, we could be passing around log2(B), not B, for both this and for
	  extract_collapse_mask().*/
	v1 = uint32(stride)
	_ = v1 > uint32(0)
	v2 = uint32(len1) / v1
	len1 = int32(v2)
	i = 0
	for {
		if !(i < stride) {
			break
		}
		if dir < 0 {
			if stride2 != 0 {
				exp_rotation1(tls, &samples[i*len1], len1, stride2, s, c)
			}
			exp_rotation1(tls, &samples[i*len1], len1, int32(1), c, s)
		} else {
			exp_rotation1(tls, &samples[i*len1], len1, int32(1), c, -s)
			if stride2 != 0 {
				exp_rotation1(tls, &samples[i*len1], len1, stride2, s, -c)
			}
		}
		i = i + 1
	}
}

var SPREAD_FACTOR = [3]int32{
	0: int32(15),
	1: int32(10),
	2: int32(5),
}

// C documentation
//
//	/** Normalizes the decoded integer pvq codeword to unit norm. */
func normalise_residual(tls *libc.TLS, iy *int32, X *OpusT_celt_norm, N int32, Ryy OpusT_opus_val32, gain OpusT_opus_val32, shift int32) {
	// The C caller supplies N > 0 and positive residual energy.
	input := unsafe.Slice(iy, int(N))
	output := unsafe.Slice(X, int(N))
	g := float32(float32(1) / float32(libc.Xsqrt(tls, float64(Ryy))) * gain)
	_ = shift // Used only by fixed-point builds, as in C.
	for i, value := range input {
		output[i] = OpusT_opus_val32(float32(value) * g)
	}
}

func extract_collapse_mask(tls *libc.TLS, iy *int32, N int32, B int32) (r uint32) {
	if B <= 1 {
		return 1
	}
	// Valid bands contain at least one coefficient per block (N >= B).
	blockSize := int(uint32(N) / uint32(B))
	coefficients := unsafe.Slice(iy, int(N))
	var mask uint32
	for i := 0; i < int(B); i++ {
		var combined uint32
		for _, value := range coefficients[i*blockSize : (i+1)*blockSize] {
			combined |= uint32(value)
		}
		if combined != 0 {
			mask |= uint32(1) << i
		}
	}
	return mask
}

func Opus_op_pvq_search_c(tls *libc.TLS, X *float32, iy *int32, K, N, arch int32) float32 {
	x := unsafe.Slice(X, N)
	pulses := unsafe.Slice(iy, N)
	y := make([]float32, N)
	signs := make([]int32, N)
	for j := range x {
		signs[j] = libc.BoolInt32(x[j] < 0)
		x[j] = float32(math.Abs(float64(x[j])))
		pulses[j] = 0
		y[j] = 0
	}
	var xy, yy float32
	left := K
	if K > N>>1 {
		var sum float32
		for _, v := range x {
			sum += v
		}
		if !(sum > float32(1e-15) && sum < 64) {
			x[0] = 1
			clear(x[1:])
			sum = 1
		}
		rcp := float32((float32(K) + float32(.8)) * float32(1/sum))
		for j := range x {
			pulses[j] = int32(math.Floor(float64(float32(rcp * x[j]))))
			y[j] = float32(pulses[j])
			yy += float32(y[j] * y[j])
			xy += float32(x[j] * y[j])
			y[j] *= 2
			left -= pulses[j]
		}
	}
	// Retain the non-debug C fallback and its two separate energy additions.
	if left > N+3 {
		tmp := float32(left)
		yy += float32(tmp * tmp)
		yy += float32(tmp * y[0])
		pulses[0] += left
		left = 0
	}
	for i := int32(0); i < left; i++ {
		yy += 1
		rxy := xy + x[0]
		bestNum := float32(rxy * rxy)
		bestDen := yy + y[0]
		best := 0
		for j := 1; j < len(x); j++ {
			rxy = xy + x[j]
			rxy = float32(rxy * rxy)
			ryy := yy + y[j]
			if float32(bestDen*rxy) > float32(ryy*bestNum) {
				bestDen = ryy
				bestNum = rxy
				best = j
			}
		}
		xy += x[best]
		yy += y[best]
		y[best] += 2
		pulses[best]++
	}
	for j := range x {
		pulses[j] = (pulses[j] ^ (-signs[j])) + signs[j]
	}
	return yy
}

func Opus_alg_quant(tls *libc.TLS, X *OpusT_celt_norm, N, K, spread, B int32, enc *OpusT_ec_enc, gain OpusT_opus_val32, resynth, arch int32) uint32 {
	if K <= 0 {
		Opus_celt_fatal(tls, __ccgo_ts+4868, __ccgo_ts+4855, 562)
	}
	if N <= 1 {
		Opus_celt_fatal(tls, __ccgo_ts+4927, __ccgo_ts+4855, 563)
	}
	// Preserve C's three padding words for up-to-four-lane pulse search.
	pulses := make([]int32, N+3)
	Opus_exp_rotation(tls, X, N, 1, B, K, spread)
	yy := Opus_op_pvq_search_c(tls, X, &pulses[0], K, N, arch)
	mask := extract_collapse_mask(tls, &pulses[0], N, B)
	Opus_encode_pulses(tls, &pulses[0], N, K, enc)
	if resynth != 0 {
		normalise_residual(tls, &pulses[0], X, N, yy, gain, 0)
		Opus_exp_rotation(tls, X, N, -1, B, K, spread)
	}
	return mask
}

// The partition driver still owns integer-addressed band/context buffers.
// C documentation
//
//	/** Decode pulse vector and combine the result with the pitch vector to produce
//	    the final normalised signal in the current band. */
func Opus_alg_unquant(tls *libc.TLS, X *OpusT_celt_norm, N, K, spread, B int32, dec *OpusT_ec_dec, gain OpusT_opus_val32) uint32 {
	if K <= 0 {
		Opus_celt_fatal(tls, __ccgo_ts+4991, __ccgo_ts+4855, 629)
	}
	if N <= 1 {
		Opus_celt_fatal(tls, __ccgo_ts+5052, __ccgo_ts+4855, 630)
	}
	// Keep the decoded pulse vector owned by Go throughout reconstruction.
	pulses := make([]int32, N)
	energy := Opus_decode_pulses(tls, &pulses[0], N, K, dec)
	normalise_residual(tls, &pulses[0], X, N, energy, gain, 0)
	Opus_exp_rotation(tls, X, N, -1, B, K, spread)
	return extract_collapse_mask(tls, &pulses[0], N, B)
}

func Opus_renormalise_vector(tls *libc.TLS, X *OpusT_celt_norm, N1 int32, gain OpusT_opus_val32, arch int32) {
	if N1 <= 0 {
		return
	}
	_ = arch
	x := unsafe.Slice(X, int(N1))
	var xy OpusT_opus_val32
	for _, value := range x {
		// Keep scalar float32 products and accumulation in C's order.
		xy = xy + OpusT_opus_val32(value*value)
	}
	E := float32(1e-15) + xy
	g := float32(float32(1) / float32(libc.Xsqrt(tls, float64(E))) * gain)
	for i := range x {
		x[i] = OpusT_opus_val32(g * x[i])
	}
}

func Opus_stereo_itheta(tls *libc.TLS, X, Y *OpusT_celt_norm, stereo int32, N1 int32, arch int32) (r OpusT_opus_int32) {
	x, y := unsafe.Slice(X, int(N1)), unsafe.Slice(Y, int(N1))
	var Emid, Eside, mid, side, xy, v1 OpusT_opus_val32
	var i, i1, itheta int32
	var m, s OpusT_celt_norm
	var x_sq, v10, v11, v13, v14, v16, v17, v9 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = Emid, Eside, i, i1, itheta, m, mid, s, side, x_sq, xy, v1, v10, v11, v13, v14, v16, v17, v9
	v1 = float32(0)
	Eside = v1
	Emid = v1
	if stereo != 0 {
		i1 = 0
		for {
			if !(i1 < N1) {
				break
			}
			m = x[i1] + y[i1]
			s = x[i1] - y[i1]
			Emid = Emid + OpusT_opus_val32(m*m)
			Eside = Eside + OpusT_opus_val32(s*s)
			i1 = i1 + 1
		}
	} else {
		_ = arch
		xy = float32(0)
		i = int32(0)
		for {
			if !(i < N1) {
				break
			}
			xy = xy + OpusT_opus_val32(x[i]*x[i])
			i = i + 1
		}
		v1 = xy
		Emid = Emid + v1
		_ = arch
		xy = float32(0)
		i = int32(0)
		for {
			if !(i < N1) {
				break
			}
			xy = xy + OpusT_opus_val32(y[i]*y[i])
			i = i + 1
		}
		v1 = xy
		Eside = Eside + v1
	}
	mid = float32(libc.Xsqrt(tls, float64(Emid)))
	side = float32(libc.Xsqrt(tls, float64(Eside)))
	v9 = side
	v10 = mid
	_ = v10 >= float32(0) && v9 >= float32(0)
	if float32(v10*v10)+float32(v9*v9) < float32(1e-18) {
		v11 = float32(0)
		goto _12
	}
	if v9 < v10 {
		v13 = v9 / v10
		x_sq = float32(v13 * v13)
		v14 = float32(float32(0.636619772367581) * (v13 + float32(float32(v13*x_sq)*(-float32(0.3333165943622589)+float32(x_sq*(float32(0.19962704181671143)+float32(x_sq*(-float32(0.13976582884788513)+float32(x_sq*(float32(0.09794234484434128)+float32(x_sq*(-float32(0.057773590087890625)+float32(x_sq*(float32(0.023040136322379112)+float32(x_sq*-float32(0.0043554059229791164))))))))))))))))
		v11 = v14
		goto _12
	} else {
		v16 = v10 / v9
		x_sq = float32(v16 * v16)
		v17 = float32(float32(0.636619772367581) * (v16 + float32(float32(v16*x_sq)*(-float32(0.3333165943622589)+float32(x_sq*(float32(0.19962704181671143)+float32(x_sq*(-float32(0.13976582884788513)+float32(x_sq*(float32(0.09794234484434128)+float32(x_sq*(-float32(0.057773590087890625)+float32(x_sq*(float32(0.023040136322379112)+float32(x_sq*-float32(0.0043554059229791164))))))))))))))))
		v11 = float32(1) - v17
		goto _12
	}
_12:
	itheta = int32(libc.Xfloor(tls, float64(float32(0.5)+float32(float32(float32(65536)*float32(16384))*v11))))
	return itheta
}

const ALLOC_STEPS = 6
const EPSILON2 = "1e-15f"
const Q15ONE4 = "1.0f"

var trim_icdf14 = [11]uint8{
	0: uint8(126),
	1: uint8(124),
	2: uint8(119),
	3: uint8(109),
	4: uint8(87),
	5: uint8(41),
	6: uint8(19),
	7: uint8(9),
	8: uint8(4),
	9: uint8(2),
}
var spread_icdf14 = [4]uint8{
	0: uint8(25),
	1: uint8(23),
	2: uint8(2),
}
var tapset_icdf14 = [3]uint8{
	0: uint8(2),
	1: uint8(1),
}

var log2_x_norm_coeff13 = [8]float32{
	0: float32(1),
	1: float32(0.8888888955116272),
	2: float32(0.8),
	3: float32(0.7272727489471436),
	4: float32(0.6666666865348816),
	5: float32(0.6153846383094788),
	6: float32(0.5714285969734192),
	7: float32(0.5333333611488342),
}
var log2_y_norm_coeff13 = [8]float32{
	1: float32(0.1699250042438507),
	2: float32(0.32192808389663696),
	3: float32(0.45943161845207214),
	4: float32(0.5849624872207642),
	5: float32(0.7004396915435791),
	6: float32(0.8073549270629883),
	7: float32(0.9068905711174011),
}

var LOG2_FRAC_TABLE = [24]uint8{
	1:  uint8(8),
	2:  uint8(13),
	3:  uint8(16),
	4:  uint8(19),
	5:  uint8(21),
	6:  uint8(23),
	7:  uint8(24),
	8:  uint8(26),
	9:  uint8(27),
	10: uint8(28),
	11: uint8(29),
	12: uint8(30),
	13: uint8(31),
	14: uint8(32),
	15: uint8(32),
	16: uint8(33),
	17: uint8(34),
	18: uint8(34),
	19: uint8(35),
	20: uint8(36),
	21: uint8(36),
	22: uint8(37),
	23: uint8(37),
}

func allocationInterpBit(a1, a2 []int32, j, mid int32) int32 { return a1[j] + mid*a2[j]>>ALLOC_STEPS }
func interp_bits2pulses(tls *libc.TLS, m *OpusT_OpusCustomMode, start int32, end int32, skip_start int32, bits1 *int32, bits2 *int32, thresh *int32, cap1 *int32, total OpusT_opus_int32, _balance *int32, skip_rsv int32, intensity *int32, intensity_rsv int32, dual_stereo *int32, dual_stereo_rsv int32, bits *int32, ebits *int32, fine_priority *int32, C int32, LM int32, ec *OpusT_ec_ctx, encode int32, prev int32, signalBandwidth int32) (r int32) {
	var alloc_floor, band_bits, band_width, codedBands, depth_threshold, done, hi, i, j, lo, logM, mid, rem, stereo, tmp, tmp1, v7, v8 int32
	var left, percoeff, psum int32
	pulse, fine, priority := unsafe.Slice(bits, end), unsafe.Slice(ebits, end), unsafe.Slice(fine_priority, end)
	var v13, v14 OpusT_opus_uint32
	a1, a2, thresholds, caps := unsafe.Slice(bits1, end), unsafe.Slice(bits2, end), unsafe.Slice(thresh, end), unsafe.Slice(cap1, end)
	codedBands = -1
	alloc_floor = C << int32(BITRES)
	stereo = libc.BoolInt32(C > int32(1))
	logM = LM << int32(BITRES)
	lo = 0
	hi = int32(1) << int32(ALLOC_STEPS)
	i = 0
	for {
		if !(i < int32(ALLOC_STEPS)) {
			break
		}
		mid = (lo + hi) >> int32(1)
		psum = 0
		done = 0
		j = end
		for {
			v7 = j
			j = j - 1
			if !(v7 > start) {
				break
			}
			tmp = allocationInterpBit(a1, a2, j, mid)
			if tmp >= thresholds[j] || done != 0 {
				done = 1
				v7 = min(tmp, caps[j])
				psum = psum + v7
			} else {
				if tmp >= alloc_floor {
					psum = psum + alloc_floor
				}
			}
		}
		if psum > total {
			hi = mid
		} else {
			lo = mid
		}
		i = i + 1
	}
	psum = 0
	/*printf ("interp bisection gave %d\n", lo);*/
	done = 0
	j = end
	for {
		v7 = j
		j = j - 1
		if !(v7 > start) {
			break
		}
		tmp1 = allocationInterpBit(a1, a2, j, lo)
		if tmp1 < thresholds[j] && !(done != 0) {
			if tmp1 >= alloc_floor {
				tmp1 = alloc_floor
			} else {
				tmp1 = 0
			}
		} else {
			done = int32(1)
		}
		/* Don't allocate more than we can actually use */
		v7 = min(tmp1, caps[j])
		tmp1 = v7
		pulse[j] = tmp1
		psum = psum + tmp1
	}
	/* Decide which bands to skip, working backwards from the end. */
	codedBands = end
	for {
		j = codedBands - int32(1)
		/* Never skip the first band, nor a band that has been boosted by
		    dynalloc.
		   In the first case, we'd be coding a bit to signal we're going to waste
		    all the other bits.
		   In the second case, we'd be coding a bit to redistribute all the bits
		    we just signaled should be concentrated in this band. */
		if j <= skip_start {
			/* Give the bit we reserved to end skipping back. */
			total = total + skip_rsv
			break
		}
		/*Figure out how many left-over bits we would be adding to this band.
		  This can include bits we've stolen back from higher, skipped bands.*/
		left = total - psum
		v13 = uint32(int32(modeBand((*OpusT_OpusCustomMode)(unsafe.Pointer(m)), codedBands)) - int32(modeBand((*OpusT_OpusCustomMode)(unsafe.Pointer(m)), start)))
		_ = v13 > uint32(0)
		v14 = uint32(left) / v13
		percoeff = int32(v14)
		left = left - (int32(modeBand((*OpusT_OpusCustomMode)(unsafe.Pointer(m)), codedBands))-int32(modeBand((*OpusT_OpusCustomMode)(unsafe.Pointer(m)), start)))*percoeff
		if left-(int32(modeBand((*OpusT_OpusCustomMode)(unsafe.Pointer(m)), j))-int32(modeBand((*OpusT_OpusCustomMode)(unsafe.Pointer(m)), start))) > 0 {
			v7 = left - (int32(modeBand((*OpusT_OpusCustomMode)(unsafe.Pointer(m)), j)) - int32(modeBand((*OpusT_OpusCustomMode)(unsafe.Pointer(m)), start)))
		} else {
			v7 = 0
		}
		rem = v7
		band_width = int32(modeBand((*OpusT_OpusCustomMode)(unsafe.Pointer(m)), codedBands)) - int32(modeBand((*OpusT_OpusCustomMode)(unsafe.Pointer(m)), j))
		band_bits = pulse[j] + percoeff*band_width + rem
		/*Only code a skip decision if we're above the threshold for this band.
		  Otherwise it is force-skipped.
		  This ensures that we have enough bits to code the skip flag.*/
		v7 = max(thresholds[j], alloc_floor+int32(1)<<int32(BITRES))
		if band_bits >= v7 {
			if encode != 0 {
				/*We choose a threshold with some hysteresis to keep bands from
				  fluctuating in and out, but we try not to fold below a certain point. */
				if codedBands > int32(17) {
					if j < prev {
						v8 = int32(7)
					} else {
						v8 = int32(9)
					}
					depth_threshold = v8
				} else {
					depth_threshold = 0
				}
				if codedBands <= start+int32(2) || band_bits > depth_threshold*band_width<<LM<<int32(BITRES)>>int32(4) && j <= signalBandwidth {
					Opus_ec_enc_bit_logp(tls, (*OpusT_ec_enc)(unsafe.Pointer(ec)), int32(1), uint32(1))
					break
				}
				Opus_ec_enc_bit_logp(tls, (*OpusT_ec_enc)(unsafe.Pointer(ec)), 0, uint32(1))
			} else {
				if Opus_ec_dec_bit_logp(tls, (*OpusT_ec_dec)(unsafe.Pointer(ec)), uint32(1)) != 0 {
					break
				}
			}
			/*We used a bit to skip this band.*/
			psum = psum + int32(1)<<int32(BITRES)
			band_bits = band_bits - int32(1)<<int32(BITRES)
		}
		/*Reclaim the bits originally allocated to this band.*/
		psum = psum - (pulse[j] + intensity_rsv)
		if intensity_rsv > 0 {
			intensity_rsv = int32(LOG2_FRAC_TABLE[j-start])
		}
		psum = psum + intensity_rsv
		if band_bits >= alloc_floor {
			/*If we have enough for a fine energy bit per channel, use it.*/
			psum = psum + alloc_floor
			pulse[j] = alloc_floor
		} else {
			/*Otherwise this band gets nothing at all.*/
			pulse[j] = 0
		}
		codedBands = codedBands - 1
	}
	if !(codedBands > start) {
		Opus_celt_fatal(tls, __ccgo_ts+5118, __ccgo_ts+5155, int32(394))
	}
	// Scalar stores and array accesses remain live: callers may alias them.
	if intensity_rsv > 0 {
		if encode != 0 {
			*intensity = min(*intensity, codedBands)
			Opus_ec_enc_uint(tls, ec, uint32(*intensity-start), uint32(codedBands+1-start))
		} else {
			*intensity = int32(uint32(start) + Opus_ec_dec_uint(tls, ec, uint32(codedBands+1-start)))
		}
	} else {
		*intensity = 0
	}
	if *intensity <= start {
		total += dual_stereo_rsv
		dual_stereo_rsv = 0
	}
	if dual_stereo_rsv > 0 {
		if encode != 0 {
			Opus_ec_enc_bit_logp(tls, ec, *dual_stereo, 1)
		} else {
			*dual_stereo = Opus_ec_dec_bit_logp(tls, ec, 1)
		}
	} else {
		*dual_stereo = 0
	}
	left = total - psum
	percoeff = int32(uint32(left) / uint32(int32(modeBand(m, codedBands))-int32(modeBand(m, start))))
	left -= (int32(modeBand(m, codedBands)) - int32(modeBand(m, start))) * percoeff
	for j = start; j < codedBands; j++ {
		pulse[j] += percoeff * (int32(modeBand(m, j+1)) - int32(modeBand(m, j)))
	}
	for j = start; j < codedBands; j++ {
		tmp := min(left, int32(modeBand(m, j+1))-int32(modeBand(m, j)))
		pulse[j] += tmp
		left -= tmp
	}
	balance := int32(0)
	for j = start; j < codedBands; j++ {
		if pulse[j] < 0 {
			Opus_celt_fatal(tls, __ccgo_ts+5170, __ccgo_ts+5155, 445)
		}
		N := (int32(modeBand(m, j+1)) - int32(modeBand(m, j))) << LM
		bit := pulse[j] + balance
		var excess int32
		if N > 1 {
			excess = max(bit-caps[j], 0)
			pulse[j] = bit - excess
			den := C*N + libc.BoolInt32(C == 2 && N > 2 && *dual_stereo == 0 && j < *intensity)
			NClogN := den * (int32(modeLogN(m, j)) + logM)
			offset := (NClogN >> 1) - den*FINE_OFFSET
			if N == 2 {
				offset += den << BITRES >> 2
			}
			if pulse[j]+offset < den*2<<BITRES {
				offset += NClogN >> 2
			} else if pulse[j]+offset < den*3<<BITRES {
				offset += NClogN >> 3
			}
			fine[j] = max(int32(0), pulse[j]+offset+(den<<(BITRES-1)))
			fine[j] = int32(uint32(fine[j]) / uint32(den) >> BITRES)
			if C*fine[j] > pulse[j]>>BITRES {
				fine[j] = pulse[j] >> stereo >> BITRES
			}
			fine[j] = min(fine[j], int32(MAX_FINE_BITS))
			priority[j] = libc.BoolInt32(fine[j]*(den<<BITRES) >= pulse[j]+offset)
			pulse[j] -= C * fine[j] << BITRES
		} else {
			excess = max(int32(0), bit-(C<<BITRES))
			pulse[j] = bit - excess
			fine[j] = 0
			priority[j] = 1
		}
		if excess > 0 {
			extraFine := min(excess>>(stereo+BITRES), int32(MAX_FINE_BITS)-fine[j])
			fine[j] += extraFine
			extraBits := extraFine * C << BITRES
			priority[j] = libc.BoolInt32(extraBits >= excess-balance)
			excess -= extraBits
		}
		balance = excess
		if pulse[j] < 0 {
			Opus_celt_fatal(tls, __ccgo_ts+5170, __ccgo_ts+5155, 516)
		}
		if fine[j] < 0 {
			Opus_celt_fatal(tls, __ccgo_ts+5201, __ccgo_ts+5155, 517)
		}
	}
	*_balance = balance
	for ; j < end; j++ {
		fine[j] = pulse[j] >> stereo >> BITRES
		if C*fine[j]<<BITRES != pulse[j] {
			Opus_celt_fatal(tls, __ccgo_ts+5233, __ccgo_ts+5155, 527)
		}
		pulse[j] = 0
		priority[j] = libc.BoolInt32(fine[j] < 1)
	}
	return codedBands
}

// Callers pass addresses of Go locals (intensity, dual_stereo, balance) as
// uintptr. Keep those locals alive at stable addresses across nested calls:
// otherwise stack growth can leave these output pointers referring to the old
// stack, making decoded PCM depend on the caller's stack depth.
//
//go:uintptrescapes
func Opus_clt_compute_allocation(tls *libc.TLS, m uintptr, start, end int32, offsets, cap1 uintptr, trim int32, intensity, dual uintptr, total int32, balance, pulses, ebits, priority uintptr, C, LM int32, ec uintptr, encode, prev, bandwidth int32) int32 {
	return clt_compute_allocation(tls, (*OpusT_OpusCustomMode)(unsafe.Pointer(m)), start, end, (*int32)(unsafe.Pointer(offsets)), (*int32)(unsafe.Pointer(cap1)), trim, (*int32)(unsafe.Pointer(intensity)), (*int32)(unsafe.Pointer(dual)), total, (*int32)(unsafe.Pointer(balance)), (*int32)(unsafe.Pointer(pulses)), (*int32)(unsafe.Pointer(ebits)), (*int32)(unsafe.Pointer(priority)), C, LM, (*OpusT_ec_ctx)(unsafe.Pointer(ec)), encode, prev, bandwidth)
}
func allocationVector(mode *OpusT_OpusCustomMode, stride, vector, band int32) byte {
	return unsafe.Slice(mode.FallocVectors, mode.FnbAllocVectors*stride)[vector*stride+band]
}

// Only [start,end) is initialized and consumed, like the C VAR_ARRAYS path.
func allocationCurve(m *OpusT_OpusCustomMode, start, end, length, trim, C, LM int32) (threshold, tilt []int32) {
	threshold, tilt = make([]int32, length), make([]int32, length)
	for j := start; j < end; j++ {
		width := int32(modeBand(m, j+1)) - int32(modeBand(m, j))
		threshold[j] = max(C<<BITRES, 3*width<<LM<<BITRES>>4)
		tilt[j] = C * width * (trim - 5 - LM) * (end - j - 1) * (int32(1) << (LM + BITRES)) >> 6
		if width<<LM == 1 {
			tilt[j] -= C << BITRES
		}
	}
	return
}
func allocationInputs(offsets, cap *int32, end int32) ([]int32, []int32) {
	return unsafe.Slice(offsets, end), unsafe.Slice(cap, end)
}
func clt_compute_allocation(tls *libc.TLS, m *OpusT_OpusCustomMode, start int32, end int32, offsets *int32, cap1 *int32, alloc_trim int32, intensity *int32, dual_stereo *int32, total OpusT_opus_int32, balance *int32, pulses *int32, ebits *int32, fine_priority *int32, C int32, LM int32, ec *OpusT_ec_ctx, encode int32, prev int32, signalBandwidth int32) (r int32) {
	off, caps := allocationInputs(offsets, cap1, end)
	var N, N1, bits1j, bits2j, bitsj, codedBands, done, dual_stereo_rsv, hi, intensity_rsv, j, len1, lo, mid, psum, skip_rsv, skip_start, v5 int32
	if total > 0 {
		v5 = total
	} else {
		v5 = 0
	}
	total = v5
	len1 = (*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands
	skip_start = start
	/* Reserve a bit to signal the end of manually skipped bands. */
	if total >= int32(1)<<int32(BITRES) {
		v5 = int32(1) << int32(BITRES)
	} else {
		v5 = 0
	}
	skip_rsv = v5
	total = total - skip_rsv
	/* Reserve bits for the intensity and dual stereo parameters. */
	v5 = int32(0)
	dual_stereo_rsv = v5
	intensity_rsv = v5
	if C == int32(2) {
		intensity_rsv = int32(LOG2_FRAC_TABLE[end-start])
		if intensity_rsv > total {
			intensity_rsv = 0
		} else {
			total = total - intensity_rsv
			if total >= int32(1)<<int32(BITRES) {
				v5 = int32(1) << int32(BITRES)
			} else {
				v5 = 0
			}
			dual_stereo_rsv = v5
			total = total - dual_stereo_rsv
		}
	}
	bits1, bits2 := make([]int32, len1), make([]int32, len1)
	thresh, trim_offset := allocationCurve(m, start, end, len1, alloc_trim, C, LM)
	lo = int32(1)
	hi = (*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbAllocVectors - int32(1)
	for cond := true; cond; cond = lo <= hi {
		done = 0
		psum = 0
		mid = (lo + hi) >> int32(1)
		j = end
		for {
			v5 = j
			j = j - 1
			if !(v5 > start) {
				break
			}
			N = int32(modeBand((*OpusT_OpusCustomMode)(unsafe.Pointer(m)), j+int32(1))) - int32(modeBand((*OpusT_OpusCustomMode)(unsafe.Pointer(m)), j))
			bitsj = C * N * int32(allocationVector(m, len1, mid, j)) << LM >> int32(2)
			if bitsj > 0 {
				v5 = max(int32(0), bitsj+trim_offset[j])
				bitsj = v5
			}
			bitsj += off[j]
			if bitsj >= thresh[j] || done != 0 {
				done = int32(1)
				/* Don't allocate more than we can actually use */
				v5 = min(bitsj, caps[j])
				psum = psum + v5
			} else {
				if bitsj >= C<<int32(BITRES) {
					psum = psum + C<<int32(BITRES)
				}
			}
		}
		if psum > total {
			hi = mid - int32(1)
		} else {
			lo = mid + int32(1)
		}
		/*printf ("lo = %d, hi = %d\n", lo, hi);*/
	}
	v5 = lo
	lo = lo - 1
	hi = v5
	/*printf ("interp between %d and %d\n", lo, hi);*/
	j = start
	for {
		if !(j < end) {
			break
		}
		N1 = int32(modeBand((*OpusT_OpusCustomMode)(unsafe.Pointer(m)), j+int32(1))) - int32(modeBand((*OpusT_OpusCustomMode)(unsafe.Pointer(m)), j))
		bits1j = C * N1 * int32(allocationVector(m, len1, lo, j)) << LM >> int32(2)
		if hi >= m.FnbAllocVectors {
			v5 = caps[j]
		} else {
			v5 = C * N1 * int32(allocationVector(m, len1, hi, j)) << LM >> int32(2)
		}
		bits2j = v5
		if bits1j > 0 {
			v5 = max(int32(0), bits1j+trim_offset[j])
			bits1j = v5
		}
		if bits2j > 0 {
			v5 = max(int32(0), bits2j+trim_offset[j])
			bits2j = v5
		}
		if lo > 0 {
			bits1j += off[j]
		}
		bits2j += off[j]
		if off[j] > 0 {
			skip_start = j
		}
		if 0 > bits2j-bits1j {
			v5 = 0
		} else {
			v5 = bits2j - bits1j
		}
		bits2j = v5
		bits1[j] = bits1j
		bits2[j] = bits2j
		j = j + 1
	}
	codedBands = interp_bits2pulses(tls, m, start, end, skip_start, unsafe.SliceData(bits1), unsafe.SliceData(bits2), unsafe.SliceData(thresh), cap1, total, balance, skip_rsv, intensity, intensity_rsv, dual_stereo, dual_stereo_rsv, pulses, ebits, fine_priority, C, LM, ec, encode, prev, signalBandwidth)
	return codedBands
}

/* Copyright (c) 2003-2008 Jean-Marc Valin
   Copyright (c) 2007-2008 CSIRO
   Copyright (c) 2007-2009 Xiph.Org Foundation
   Written by Jean-Marc Valin */
/**
  @file arch.h
  @brief Various architecture definitions for CELT
*/
/*
   Redistribution and use in source and binary forms, with or without
   modification, are permitted provided that the following conditions
   are met:

   - Redistributions of source code must retain the above copyright
   notice, this list of conditions and the following disclaimer.

   - Redistributions in binary form must reproduce the above copyright
   notice, this list of conditions and the following disclaimer in the
   documentation and/or other materials provided with the distribution.

   THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
   ``AS IS'' AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
   LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
   A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER
   OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
   EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
   PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
   PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
   LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
   NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
   SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

// C documentation
//
//	/* This is a faster version of ec_tell_frac() that takes advantage
//	   of the low (1/8 bit) resolution to use just a linear function
//	   followed by a lookup to determine the exact transition thresholds. */
func Opus_ec_tell_frac(tls *libc.TLS, ctx *OpusT_ec_ctx) OpusT_opus_uint32 {
	nbits := uint32(ctx.Fnbits_total) << BITRES
	l := bits.Len32(ctx.Frng)
	// Entropy contexts supply a normalized range with at least 16 bits.
	r := ctx.Frng >> (l - 16)
	b := (r >> 12) - 8
	if r > correction[b] {
		b++
	}
	return nbits - (uint32(l<<3) + b)
}

var correction = [8]uint32{
	0: uint32(35733),
	1: uint32(38967),
	2: uint32(42495),
	3: uint32(46340),
	4: uint32(50535),
	5: uint32(55109),
	6: uint32(60097),
	7: uint32(65535),
}

const EPSILON3 = 1e-15
const MIN_STEREO_ENERGY = 1e-10
const NORM_SCALING1 = 1
const Q31ONE3 = 1

var trim_icdf15 = [11]uint8{
	0: uint8(126),
	1: uint8(124),
	2: uint8(119),
	3: uint8(109),
	4: uint8(87),
	5: uint8(41),
	6: uint8(19),
	7: uint8(9),
	8: uint8(4),
	9: uint8(2),
}
var spread_icdf15 = [4]uint8{
	0: uint8(25),
	1: uint8(23),
	2: uint8(2),
}
var tapset_icdf15 = [3]uint8{
	0: uint8(2),
	1: uint8(1),
}

var log2_x_norm_coeff14 = [8]float32{
	0: float32(1),
	1: float32(0.8888888955116272),
	2: float32(0.8),
	3: float32(0.7272727489471436),
	4: float32(0.6666666865348816),
	5: float32(0.6153846383094788),
	6: float32(0.5714285969734192),
	7: float32(0.5333333611488342),
}
var log2_y_norm_coeff14 = [8]float32{
	1: float32(0.1699250042438507),
	2: float32(0.32192808389663696),
	3: float32(0.45943161845207214),
	4: float32(0.5849624872207642),
	5: float32(0.7004396915435791),
	6: float32(0.8073549270629883),
	7: float32(0.9068905711174011),
}

func Opus_hysteresis_decision(tls *libc.TLS, val OpusT_opus_val16, thresholds *OpusT_opus_val16, hysteresis *OpusT_opus_val16, N int32, prev int32) (r int32) {
	// There are N thresholds and N+1 states; prev is in [0, N].
	levels := unsafe.Slice(thresholds, int(N))
	margins := unsafe.Slice(hysteresis, int(N))
	var i int32
	for i < N {
		if val < levels[i] {
			break
		}
		i++
	}
	if i > prev && val < levels[prev]+margins[prev] {
		i = prev
	}
	if i < prev && val > levels[prev-1]-margins[prev-1] {
		i = prev
	}
	return i
}

func Opus_celt_lcg_rand(tls *libc.TLS, seed OpusT_opus_uint32) (r OpusT_opus_uint32) {
	return uint32(1664525)*seed + uint32(1013904223)
}

// C documentation
//
//	/* This is a cos() approximation designed to be bit-exact on any platform. Bit exactness
//	   with this approximation is important because it has an impact on the bit allocation */
func Opus_bitexact_cos(tls *libc.TLS, x OpusT_opus_int16) (r OpusT_opus_int16) {
	var tmp OpusT_opus_int32
	var x2 OpusT_opus_int16
	_, _ = tmp, x2
	tmp = (int32(4096) + int32(x)*int32(x)) >> int32(13)
	_ = tmp <= int32(32767)
	x2 = int16(tmp)
	x2 = int16(int32(32767) - int32(x2) + (int32(16384)+int32(x2)*int32(int16(-int32(7651)+(int32(16384)+int32(x2)*int32(int16(int32(8277)+(int32(16384)+int32(int16(-int32(626)))*int32(x2))>>int32(15))))>>int32(15))))>>int32(15))
	_ = int32(x2) <= int32(32766)
	return int16(int32(1) + int32(x2))
}

func Opus_bitexact_log2tan(tls *libc.TLS, isin int32, icos int32) (r int32) {
	var lc, ls int32
	_, _ = lc, ls
	lc = int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, uint32(icos))
	ls = int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, uint32(isin))
	icos = icos << (int32(15) - lc)
	isin = isin << (int32(15) - ls)
	return (ls-lc)*(int32(1)<<int32(11)) + (int32(16384)+int32(int16(isin))*int32(int16((int32(16384)+int32(int16(isin))*int32(int16(-int32(2597))))>>int32(15)+int32(7932))))>>int32(15) - (int32(16384)+int32(int16(icos))*int32(int16((int32(16384)+int32(int16(icos))*int32(int16(-int32(2597))))>>int32(15)+int32(7932))))>>int32(15)
}

// C documentation
//
//	/* Compute the amplitude (sqrt energy) in each of the bands */
func Opus_compute_band_energies(tls *libc.TLS, bands *int16, shortMdctSize, nbEBands int32, X, bandE *float32, end, C, LM, arch int32) {
	if end <= 0 {
		return
	}
	// As in the C do/while, channel zero is visited even for C=0.
	channels := max(C, 1)
	N := shortMdctSize << LM
	eBands := unsafe.Slice(bands, end+1)
	in := unsafe.Slice(X, (channels-1)*N+(int32(eBands[end])<<LM))
	out := unsafe.Slice(bandE, (channels-1)*nbEBands+end)
	for c := int32(0); c < channels; c++ {
		for band := int32(0); band < end; band++ {
			xy := float32(0)
			for j := int32(eBands[band]) << LM; j < int32(eBands[band+1])<<LM; j++ {
				sample := in[c*N+j]
				// Keep the product rounded before accumulation, including on ARM64.
				xy += float32(sample * sample)
			}
			sum := float32(1e-27) + xy
			out[c*nbEBands+band] = float32(math.Sqrt(float64(sum)))
		}
	}
}

// C documentation
//
//	/* Normalise each band such that the energy is one. */
func Opus_normalise_bands(tls *libc.TLS, bands *int16, shortMdctSize, nbEBands int32, freq, X, bandE *float32, end, C, M int32) {
	if end <= 0 {
		return
	}
	// C's do/while visits channel zero even if C is zero.
	channels := max(C, 1)
	N := M * shortMdctSize
	eBands := unsafe.Slice(bands, end+1)
	extent := (channels-1)*N + M*int32(eBands[end])
	in := unsafe.Slice(freq, extent)
	out := unsafe.Slice(X, extent)
	energy := unsafe.Slice(bandE, (channels-1)*nbEBands+end)
	for c := int32(0); c < channels; c++ {
		for i := int32(0); i < end; i++ {
			g := float32(1) / (float32(1e-27) + energy[i+c*nbEBands])
			for j := M * int32(eBands[i]); j < M*int32(eBands[i+1]); j++ {
				out[j+c*N] = in[j+c*N] * g
			}
		}
	}
}

// C documentation
//
//	/* De-normalise the energy to produce the synthesis from the unit-energy bands */
func Opus_denormalise_bands(tls *libc.TLS, bands *OpusT_opus_int16, shortMdctSize int32, X *OpusT_celt_norm, freq *OpusT_celt_sig, bandLogE *OpusT_celt_glog, start, end, M, downsample, silence int32) {
	// Pass the two mode inputs explicitly so the table stays GC-visible here.
	eBands := unsafe.Slice(bands, end+1)
	N := M * shortMdctSize
	bound := M * int32(eBands[end])
	if downsample != 1 {
		bound = min(bound, N/downsample)
	}
	output := unsafe.Slice(freq, N)
	if silence != 0 {
		clear(output)
		return
	}
	input := unsafe.Slice(X, M*int32(eBands[end]))
	logE := unsafe.Slice(bandLogE, end)
	if start != 0 {
		clear(output[:M*int32(eBands[start])])
	}
	for i := start; i < end; i++ {
		lg := logE[i] + Opus_eMeans[i]
		if float32(32) < lg {
			lg = 32
		}
		integer := int32(math.Floor(float64(lg)))
		var gain float32
		if integer >= -50 {
			frac := lg - float32(integer)
			polynomial := float32(0.9999999403953552) + float32(frac*(float32(0.6931530833244324)+float32(frac*(float32(0.24015361070632935)+float32(frac*(float32(0.05582631751894951)+float32(frac*(float32(0.00898933969438076)+float32(frac*float32(0.0018775766948238015))))))))))
			encoded := (math.Float32bits(polynomial) + (uint32(integer) << 23)) & 0x7fffffff
			gain = math.Float32frombits(encoded)
		}
		for j := M * int32(eBands[i]); j < M*int32(eBands[i+1]); j++ {
			output[j] = input[j] * gain
		}
	}
	if start > end {
		Opus_celt_fatal(tls, __ccgo_ts+5281, __ccgo_ts+5312, 254)
	}
	clear(output[bound:])
}

// C documentation
//
//	/* This prevents energy collapse for transients with multiple short MDCTs */
func collapseExp2(x float32) float32 {
	integer := int32(math.Floor(float64(x)))
	if integer < -50 {
		return 0
	}
	frac := x - float32(integer)
	p := float32(0.00898933969438076) + float32(frac*float32(0.0018775766948238015))
	p = float32(0.05582631751894951) + float32(frac*p)
	p = float32(0.24015361070632935) + float32(frac*p)
	p = float32(0.6931530833244324) + float32(frac*p)
	p = float32(0.9999999403953552) + float32(frac*p)
	return math.Float32frombits((math.Float32bits(p) + (uint32(integer) << 23)) & 0x7fffffff)
}

func Opus_anti_collapse(tls *libc.TLS, bands *int16, nbBands int32, X *float32, masks *byte, LM, C, size, start, end int32, logE, prev1logE, prev2logE *float32, pulses *int32, seed uint32, encode, arch int32) {
	if start >= end {
		return
	}
	b := unsafe.Slice(bands, end+1)
	x := unsafe.Slice(X, max(C, 1)*size)
	mask := unsafe.Slice(masks, max(end*C, 1))
	pulse := unsafe.Slice(pulses, end)
	energy := unsafe.Slice(logE, max(C, 1)*nbBands)
	historyChannels := max(C, 1)
	if encode == 0 && C == 1 {
		historyChannels = 2
	}
	p1 := unsafe.Slice(prev1logE, historyChannels*nbBands)
	p2 := unsafe.Slice(prev2logE, historyChannels*nbBands)
	for i := start; i < end; i++ {
		n0 := int32(b[i+1]) - int32(b[i])
		depth := int32((uint32(1+pulse[i]) / uint32(n0)) >> LM)
		threshold := float32(.5) * collapseExp2(float32(-.125)*float32(depth))
		sqrt1 := float32(1) / float32(math.Sqrt(float64(n0<<LM)))
		for c := int32(0); c < max(C, 1); c++ {
			index := c*nbBands + i
			prev1, prev2 := p1[index], p2[index]
			if encode == 0 && C == 1 {
				if !(prev1 > p1[nbBands+i]) {
					prev1 = p1[nbBands+i]
				}
				if !(prev2 > p2[nbBands+i]) {
					prev2 = p2[nbBands+i]
				}
			}
			previous := prev2
			if prev1 < prev2 {
				previous = prev1
			}
			diff := energy[index] - previous
			if 0 > diff {
				diff = 0
			}
			r := float32(2) * collapseExp2(-diff)
			if LM == 3 {
				r = float32(r * float32(1.41421356))
			}
			if threshold < r {
				r = threshold
			}
			r = float32(r * sqrt1)
			offset := c*size + (int32(b[i]) << LM)
			renormalize := false
			for k := int32(0); k < 1<<LM; k++ {
				if int32(mask[i*C+c])&(1<<k) == 0 {
					for j := int32(0); j < n0; j++ {
						seed = Opus_celt_lcg_rand(tls, seed)
						value := -r
						if seed&0x8000 != 0 {
							value = r
						}
						x[offset+(j<<LM)+k] = value
					}
					renormalize = true
				}
			}
			if renormalize {
				Opus_renormalise_vector(tls, &x[offset], n0<<LM, 1, arch)
			}
		}
	}
}

func anti_collapse_legacy(tls *libc.TLS, m, X, masks uintptr, LM, C, size, start, end int32, logE, p1, p2, pulses uintptr, seed uint32, encode, arch int32) {
	mode := (*OpusT_OpusCustomMode)(unsafe.Pointer(m))
	Opus_anti_collapse(tls, mode.FeBands, mode.FnbEBands, (*float32)(unsafe.Pointer(X)), (*byte)(unsafe.Pointer(masks)), LM, C, size, start, end, (*float32)(unsafe.Pointer(logE)), (*float32)(unsafe.Pointer(p1)), (*float32)(unsafe.Pointer(p2)), (*int32)(unsafe.Pointer(pulses)), seed, encode, arch)
}

// C documentation
//
//	/* Compute the weights to use for optimizing normalized distortion across
//	   channels. We use the amplitude to weight square distortion, which means
//	   that we use the square root of the value we would have been using if we
//	   wanted to minimize the MSE in the non-normalized domain. This roughly
//	   corresponds to some quick-and-dirty perceptual experiments I ran to
//	   measure inter-aural masking (there doesn't seem to be any published data
//	   on the topic). */
func compute_channel_weights(tls *libc.TLS, Ex OpusT_celt_ener, Ey OpusT_celt_ener, w *[2]OpusT_opus_val16) {
	minE := Ey
	if Ex < Ey {
		minE = Ex
	}
	/* Adjustment to make the weights a bit more conservative. */
	w[0] = Ex + minE/float32(3)
	w[1] = Ey + minE/float32(3)
}

func intensity_stereo(tls *libc.TLS, m *OpusT_OpusCustomMode, X *OpusT_celt_norm, Y *OpusT_celt_norm, bandE *OpusT_celt_ener, bandID int32, N int32) {
	energies := unsafe.Slice(bandE, int(m.FnbEBands)+int(bandID)+1)
	left, right := energies[bandID], energies[int(bandID)+int(m.FnbEBands)]
	norm := float32(1e-15) + float32(libc.Xsqrt(tls, float64(float32(1e-15)+OpusT_opus_val32(left*left)+OpusT_opus_val32(right*right))))
	a1, a2 := left/norm, right/norm
	x, y := unsafe.Slice(X, int(N)), unsafe.Slice(Y, int(N))
	for j := range x {
		x[j] = OpusT_opus_val16(a1*x[j]) + OpusT_opus_val16(a2*y[j])
		// Side is not encoded, no need to calculate.
	}
}

func stereo_split(tls *libc.TLS, X *OpusT_celt_norm, Y *OpusT_celt_norm, N int32) {
	if N <= 0 {
		return
	}
	x := unsafe.Slice(X, int(N))
	y := unsafe.Slice(Y, int(N))
	for j := range x {
		// Preserve float32 rounding before the sum/difference.
		l := float32(float32(0.70710678) * x[j])
		r := float32(float32(0.70710678) * y[j])
		x[j] = l + r
		y[j] = r - l
	}
}

func stereo_merge(tls *libc.TLS, X *OpusT_celt_norm, Y *OpusT_celt_norm, mid OpusT_opus_val32, N1 int32, arch int32) {
	if N1 <= 0 {
		return
	}
	x, y := unsafe.Slice(X, int(N1)), unsafe.Slice(Y, int(N1))
	_ = arch
	var xp, side OpusT_opus_val32
	for i := range x {
		xp = xp + OpusT_opus_val32(y[i]*x[i])
	}
	for _, value := range y {
		side = side + OpusT_opus_val32(value*value)
	}
	xp = OpusT_opus_val32(mid * xp)
	El := OpusT_opus_val32(mid*mid) + side - OpusT_opus_val32(float32(2)*xp)
	Er := OpusT_opus_val32(mid*mid) + side + OpusT_opus_val32(float32(2)*xp)
	if Er < float32(0.0006) || El < float32(0.0006) {
		copy(y, x)
		return
	}
	lgain := float32(1) / float32(libc.Xsqrt(tls, float64(El)))
	rgain := float32(1) / float32(libc.Xsqrt(tls, float64(Er)))
	for i := range x {
		l, r := OpusT_opus_val32(mid*x[i]), y[i]
		x[i] = OpusT_opus_val32(lgain * (l - r))
		y[i] = OpusT_opus_val32(rgain * (l + r))
	}
}

// C documentation
//
//	/* Decide whether we should spread the pulses in the current frame */
func Opus_spreading_decision(tls *libc.TLS, bands *int16, nbBands, shortMdctSize int32, X *float32, average *int32, lastDecision int32, hfAverage, tapset *int32, updateHF, end, C, M int32, spreadWeight *int32) int32 {
	if end <= 0 {
		Opus_celt_fatal(tls, __ccgo_ts+5328, __ccgo_ts+5312, 480)
	}
	b := unsafe.Slice(bands, end+1)
	n0 := M * shortMdctSize
	if M*(int32(b[end])-int32(b[end-1])) <= 8 {
		return SPREAD_NONE
	}
	x := unsafe.Slice(X, max(C, 1)*n0)
	weights := unsafe.Slice(spreadWeight, end)
	var sum, weightedBands, hfSum int32
	for c := int32(0); c < max(C, 1); c++ {
		for i := int32(0); i < end; i++ {
			n := M * (int32(b[i+1]) - int32(b[i]))
			if n <= 8 {
				continue
			}
			var count [3]int32
			offset := M*int32(b[i]) + c*n0
			for j := int32(0); j < n; j++ {
				value := x[offset+j]
				square := float32(value * value)
				v := float32(square * float32(n))
				if v < .25 {
					count[0]++
				}
				if v < .0625 {
					count[1]++
				}
				if v < .015625 {
					count[2]++
				}
			}
			if i > nbBands-4 {
				hfSum = int32(uint32(hfSum) + uint32(32*(count[1]+count[0]))/uint32(n))
			}
			tmp := libc.BoolInt32(2*count[2] >= n) + libc.BoolInt32(2*count[1] >= n) + libc.BoolInt32(2*count[0] >= n)
			sum += tmp * weights[i]
			weightedBands += weights[i]
		}
	}
	if updateHF != 0 {
		if hfSum != 0 {
			hfSum = int32(uint32(hfSum) / uint32(C*(4-nbBands+end)))
		}
		*hfAverage = (*hfAverage + hfSum) >> 1
		hfSum = *hfAverage
		if *tapset == 2 {
			hfSum += 4
		} else if *tapset == 0 {
			hfSum -= 4
		}
		if hfSum > 22 {
			*tapset = 2
		} else if hfSum > 18 {
			*tapset = 1
		} else {
			*tapset = 0
		}
	}
	if weightedBands <= 0 {
		Opus_celt_fatal(tls, __ccgo_ts+5352, __ccgo_ts+5312, 536)
	}
	if sum < 0 {
		Opus_celt_fatal(tls, __ccgo_ts+5380, __ccgo_ts+5312, 537)
	}
	sum = int32(uint32(sum<<8) / uint32(weightedBands))
	sum = (sum + *average) >> 1
	*average = sum
	sum = (3*sum + ((3 - lastDecision) << 7) + 64 + 2) >> 2
	if sum < 80 {
		return SPREAD_AGGRESSIVE
	}
	if sum < 256 {
		return SPREAD_NORMAL
	}
	if sum < 384 {
		return SPREAD_LIGHT
	}
	return SPREAD_NONE
}

// C documentation
//
//	/* Indexing table for converting from natural Hadamard to ordery Hadamard
//	   This is essentially a bit-reversed Gray, on top of which we've added
//	   an inversion of the order because we want the DC at the end rather than
//	   the beginning. The lines are for N=2, 4, 8, 16 */
var ordery_table = [30]int32{
	0:  int32(1),
	2:  int32(3),
	4:  int32(2),
	5:  int32(1),
	6:  int32(7),
	8:  int32(4),
	9:  int32(3),
	10: int32(6),
	11: int32(1),
	12: int32(5),
	13: int32(2),
	14: int32(15),
	16: int32(8),
	17: int32(7),
	18: int32(12),
	19: int32(3),
	20: int32(11),
	21: int32(4),
	22: int32(14),
	23: int32(1),
	24: int32(9),
	25: int32(6),
	26: int32(13),
	27: int32(2),
	28: int32(10),
	29: int32(5),
}

func deinterleave_hadamard(tls *libc.TLS, X *OpusT_celt_norm, N0 int32, stride int32, hadamard int32) {
	if stride <= 0 {
		Opus_celt_fatal(tls, __ccgo_ts+5405, __ccgo_ts+5312, 582)
	}
	if N0 == 0 {
		return
	}
	values := unsafe.Slice(X, N0*stride)
	tmp := make([]OpusT_celt_norm, len(values))
	for i := int32(0); i < stride; i++ {
		row := i
		if hadamard != 0 {
			row = ordery_table[stride-2+i]
		}
		for j := int32(0); j < N0; j++ {
			tmp[row*N0+j] = values[j*stride+i]
		}
	}
	copy(values, tmp)
}

func interleave_hadamard(tls *libc.TLS, X *OpusT_celt_norm, N0 int32, stride int32, hadamard int32) {
	N := N0 * stride
	if N == 0 {
		return
	}
	values := unsafe.Slice(X, N)
	tmp := make([]OpusT_celt_norm, len(values))
	for i := int32(0); i < stride; i++ {
		row := i
		if hadamard != 0 {
			row = ordery_table[stride-2+i]
		}
		for j := int32(0); j < N0; j++ {
			tmp[j*stride+i] = values[row*N0+j]
		}
	}
	copy(values, tmp)
}

func Opus_haar1(tls *libc.TLS, X *OpusT_celt_norm, N0 int32, stride int32) {
	pairs := int(N0 >> 1)
	if pairs <= 0 || stride <= 0 {
		return
	}
	step := int(stride)
	// C processes complete pairs only; an odd trailing row is untouched.
	x := unsafe.Slice(X, 2*pairs*step)
	for i := 0; i < step; i++ {
		for j := 0; j < pairs; j++ {
			a, b := step*2*j+i, step*(2*j+1)+i
			tmp1 := float32(float32(0.70710678) * x[a])
			tmp2 := float32(float32(0.70710678) * x[b])
			x[a] = tmp1 + tmp2
			x[b] = tmp1 - tmp2
		}
	}
}

func compute_qn(tls *libc.TLS, N int32, b int32, offset int32, pulse_cap int32, stereo int32) (r int32) {
	var N2, qb, qn, v4 int32
	var v1, v2 OpusT_opus_int32
	_, _, _, _, _, _ = N2, qb, qn, v1, v2, v4
	N2 = int32(2)*N - int32(1)
	if stereo != 0 && N == int32(2) {
		N2 = N2 - 1
	}
	/* The upper limit ensures that in a stereo split with itheta==16384, we'll
	   always have enough bits left over to code at least one pulse in the
	   side; otherwise it would collapse, since it doesn't get folded. */
	v1 = N2
	_ = v1 > int32(0)
	v2 = (b + N2*offset) / v1
	qb = v2
	if b-pulse_cap-int32(4)<<int32(BITRES) < qb {
		v4 = b - pulse_cap - int32(4)<<int32(BITRES)
	} else {
		v4 = qb
	}
	qb = v4
	if int32(8)<<int32(BITRES) < qb {
		v4 = int32(8) << int32(BITRES)
	} else {
		v4 = qb
	}
	qb = v4
	if qb < int32(1)<<int32(BITRES)>>int32(1) {
		qn = int32(1)
	} else {
		qn = int32(exp2_table8[qb&int32(0x7)]) >> (int32(14) - qb>>int32(BITRES))
		qn = (qn + int32(1)) >> int32(1) << int32(1)
	}
	if !(qn <= int32(256)) {
		Opus_celt_fatal(tls, __ccgo_ts+5432, __ccgo_ts+5312, int32(660))
	}
	return qn
}

var exp2_table8 = [8]OpusT_opus_int16{
	0: int16(16384),
	1: int16(17866),
	2: int16(19483),
	3: int16(21247),
	4: int16(23170),
	5: int16(25267),
	6: int16(27554),
	7: int16(30048),
}

type band_ctx = struct {
	Fencode            int32
	Fresynth           int32
	Fm                 *OpusT_OpusCustomMode
	Fi                 int32
	Fintensity         int32
	Fspread            int32
	Ftf_change         int32
	Fec                *OpusT_ec_ctx
	Fremaining_bits    OpusT_opus_int32
	FbandE             *OpusT_celt_ener
	Fseed              OpusT_opus_uint32
	Farch              int32
	Ftheta_round       int32
	Fdisable_inv       int32
	Favoid_split_noise int32
}

func celtNormAdd(p *float32, offset int32) *float32 {
	return (*float32)(unsafe.Add(unsafe.Pointer(p), int(offset)*4))
}

func bandContextEnergy(ctx *band_ctx, index int32) float32 {
	return *(*float32)(unsafe.Add(unsafe.Pointer(ctx.FbandE), int(index)*4))
}

type split_ctx = struct {
	Finv    int32
	Fimid   int32
	Fiside  int32
	Fdelta  int32
	Fitheta int32
	Fqalloc int32
}

// Context, spectra and local outputs remain visible across entropy calls.
func compute_theta(tls *libc.TLS, ctx *band_ctx, sctx *split_ctx, X *float32, Y *float32, N int32, b *int32, B int32, B0 int32, LM int32, stereo int32, fill *int32) {
	var bandE *OpusT_celt_ener
	var ec *OpusT_ec_ctx
	var m *OpusT_OpusCustomMode
	var bias, delta, down, encode, fl, fl1, fm, fs, fs1, ft, ft1, i, imid, intensity, inv, iside, itheta, itheta_q30, j, offset, p0, pulse_cap, qalloc, qn, unquantized, x, x0, v1, v5, v6, v7 int32
	var tell OpusT_opus_int32
	var v2, v3 OpusT_opus_uint32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bandE, bias, delta, down, ec, encode, fl, fl1, fm, fs, fs1, ft, ft1, i, imid, intensity, inv, iside, itheta, itheta_q30, j, m, offset, p0, pulse_cap, qalloc, qn, tell, unquantized, x, x0, v1, v2, v3, v5, v6, v7
	itheta = 0
	itheta_q30 = 0
	inv = 0
	encode = ctx.Fencode
	m = ctx.Fm
	i = ctx.Fi
	intensity = ctx.Fintensity
	ec = ctx.Fec
	bandE = ctx.FbandE
	/* Decide on the resolution to give to the split parameter theta */
	pulse_cap = int32(modeLogN(m, i)) + LM*(int32(1)<<int32(BITRES))
	if stereo != 0 && N == int32(2) {
		v1 = int32(QTHETA_OFFSET_TWOPHASE)
	} else {
		v1 = int32(QTHETA_OFFSET)
	}
	offset = pulse_cap>>int32(1) - v1
	qn = compute_qn(tls, N, *b, offset, pulse_cap, stereo)
	if stereo != 0 && i >= intensity {
		qn = int32(1)
	}
	if encode != 0 {
		/* theta is the atan() of the ratio between the (normalized)
		   side and mid. With just that parameter, we can re-scale both
		   mid and side because we know that 1) they have unit norm and
		   2) they are orthogonal. */
		itheta_q30 = Opus_stereo_itheta(tls, X, Y, stereo, N, ctx.Farch)
		itheta = itheta_q30 >> int32(16)
	}
	tell = int32(Opus_ec_tell_frac(tls, (*OpusT_ec_ctx)(unsafe.Pointer(ec))))
	if qn != int32(1) {
		if encode != 0 {
			if !(stereo != 0) || ctx.Ftheta_round == 0 {
				itheta = (itheta*qn + int32(8192)) >> int32(14)
				if !(stereo != 0) && ctx.Favoid_split_noise != 0 && itheta > 0 && itheta < qn {
					v2 = uint32(qn)
					_ = v2 > uint32(0)
					v3 = uint32(itheta*int32(16384)) / v2
					/* Check if the selected value of theta will cause the bit allocation
					   to inject noise on one side. If so, make sure the energy of that side
					   is zero. */
					unquantized = int32(v3)
					imid = int32(Opus_bitexact_cos(tls, int16(unquantized)))
					iside = int32(Opus_bitexact_cos(tls, int16(int32(16384)-unquantized)))
					delta = (int32(16384) + int32(int16((N-int32(1))<<int32(7)))*int32(int16(Opus_bitexact_log2tan(tls, iside, imid)))) >> int32(15)
					if delta > *b {
						itheta = qn
					} else {
						if delta < -*b {
							itheta = 0
						}
					}
				}
			} else {
				if itheta > int32(8192) {
					v1 = int32(32767) / qn
				} else {
					v1 = -int32(32767) / qn
				}
				/* Bias quantization towards itheta=0 and itheta=16384. */
				bias = v1
				if 0 > (itheta*qn+bias)>>int32(14) {
					v6 = 0
				} else {
					v6 = (itheta*qn + bias) >> int32(14)
				}
				if qn-int32(1) < v6 {
					v5 = qn - int32(1)
				} else {
					if 0 > (itheta*qn+bias)>>int32(14) {
						v7 = 0
					} else {
						v7 = (itheta*qn + bias) >> int32(14)
					}
					v5 = v7
				}
				down = v5
				if ctx.Ftheta_round < 0 {
					itheta = down
				} else {
					itheta = down + int32(1)
				}
			}
		}
		/* Entropy coding of the angle. We use a uniform pdf for the
		   time split, a step for stereo, and a triangular one for the rest. */
		if stereo != 0 && N > int32(2) {
			p0 = int32(3)
			x = itheta
			x0 = qn / int32(2)
			ft = p0*(x0+int32(1)) + x0
			/* Use a probability of p0 up to itheta=8192 and then use 1 after */
			if encode != 0 {
				if x <= x0 {
					v1 = p0 * x
				} else {
					v1 = x - int32(1) - x0 + (x0+int32(1))*p0
				}
				if x <= x0 {
					v5 = p0 * (x + int32(1))
				} else {
					v5 = x - x0 + (x0+int32(1))*p0
				}
				Opus_ec_encode(tls, (*OpusT_ec_enc)(unsafe.Pointer(ec)), uint32(v1), uint32(v5), uint32(ft))
			} else {
				fs = int32(Opus_ec_decode(tls, (*OpusT_ec_dec)(unsafe.Pointer(ec)), uint32(ft)))
				if fs < (x0+int32(1))*p0 {
					x = fs / p0
				} else {
					x = x0 + int32(1) + (fs - (x0+int32(1))*p0)
				}
				if x <= x0 {
					v1 = p0 * x
				} else {
					v1 = x - int32(1) - x0 + (x0+int32(1))*p0
				}
				if x <= x0 {
					v5 = p0 * (x + int32(1))
				} else {
					v5 = x - x0 + (x0+int32(1))*p0
				}
				Opus_ec_dec_update(tls, (*OpusT_ec_dec)(unsafe.Pointer(ec)), uint32(v1), uint32(v5), uint32(ft))
				itheta = x
			}
		} else {
			if B0 > int32(1) || stereo != 0 {
				/* Uniform pdf */
				if encode != 0 {
					Opus_ec_enc_uint(tls, (*OpusT_ec_enc)(unsafe.Pointer(ec)), uint32(itheta), uint32(qn+int32(1)))
				} else {
					itheta = int32(Opus_ec_dec_uint(tls, (*OpusT_ec_dec)(unsafe.Pointer(ec)), uint32(qn+int32(1))))
				}
			} else {
				fs1 = int32(1)
				ft1 = (qn>>int32(1) + int32(1)) * (qn>>int32(1) + int32(1))
				if encode != 0 {
					if itheta <= qn>>int32(1) {
						v1 = itheta + int32(1)
					} else {
						v1 = qn + int32(1) - itheta
					}
					fs1 = v1
					if itheta <= qn>>int32(1) {
						v1 = itheta * (itheta + int32(1)) >> int32(1)
					} else {
						v1 = ft1 - (qn+int32(1)-itheta)*(qn+int32(2)-itheta)>>int32(1)
					}
					fl = v1
					Opus_ec_encode(tls, (*OpusT_ec_enc)(unsafe.Pointer(ec)), uint32(fl), uint32(fl+fs1), uint32(ft1))
				} else {
					/* Triangular pdf */
					fl1 = 0
					fm = int32(Opus_ec_decode(tls, (*OpusT_ec_dec)(unsafe.Pointer(ec)), uint32(ft1)))
					if fm < qn>>int32(1)*(qn>>int32(1)+int32(1))>>int32(1) {
						itheta = int32((Opus_isqrt32(tls, uint32(8)*uint32(fm)+uint32(1)) - uint32(1)) >> int32(1))
						fs1 = itheta + int32(1)
						fl1 = itheta * (itheta + int32(1)) >> int32(1)
					} else {
						itheta = int32((uint32(int32(2)*(qn+int32(1))) - Opus_isqrt32(tls, uint32(8)*uint32(ft1-fm-int32(1))+uint32(1))) >> int32(1))
						fs1 = qn + int32(1) - itheta
						fl1 = ft1 - (qn+int32(1)-itheta)*(qn+int32(2)-itheta)>>int32(1)
					}
					Opus_ec_dec_update(tls, (*OpusT_ec_dec)(unsafe.Pointer(ec)), uint32(fl1), uint32(fl1+fs1), uint32(ft1))
				}
			}
		}
		if !(itheta >= int32(0)) {
			Opus_celt_fatal(tls, __ccgo_ts+5460, __ccgo_ts+5312, int32(840))
		}
		v2 = uint32(qn)
		_ = v2 > uint32(0)
		v3 = uint32(itheta*int32(16384)) / v2
		itheta = int32(v3)
		if encode != 0 && stereo != 0 {
			if itheta == 0 {
				intensity_stereo(tls, m, X, Y, bandE, i, N)
			} else {
				stereo_split(tls, X, Y, N)
			}
		}
		/* NOTE: Renormalising X and Y *may* help fixed-point a bit at very high rate.
		   Let's do that at higher complexity */
	} else {
		if stereo != 0 {
			if encode != 0 {
				inv = libc.BoolInt32(itheta > int32(8192) && !(ctx.Fdisable_inv != 0))
				if inv != 0 {
					j = 0
					for {
						if !(j < N) {
							break
						}
						unsafe.Slice(Y, N)[j] = -unsafe.Slice(Y, N)[j]
						j = j + 1
					}
				}
				intensity_stereo(tls, m, X, Y, bandE, i, N)
			}
			if *b > int32(2)<<int32(BITRES) && ctx.Fremaining_bits > int32(2)<<int32(BITRES) {
				if encode != 0 {
					Opus_ec_enc_bit_logp(tls, (*OpusT_ec_enc)(unsafe.Pointer(ec)), inv, uint32(2))
				} else {
					inv = Opus_ec_dec_bit_logp(tls, (*OpusT_ec_dec)(unsafe.Pointer(ec)), uint32(2))
				}
			} else {
				inv = 0
			}
			/* inv flag override to avoid problems with downmixing. */
			if ctx.Fdisable_inv != 0 {
				inv = 0
			}
			itheta = 0
			itheta_q30 = 0
		}
	}
	qalloc = int32(Opus_ec_tell_frac(tls, (*OpusT_ec_ctx)(unsafe.Pointer(ec))) - uint32(tell))
	*b -= qalloc
	if itheta == 0 {
		imid = int32(32767)
		iside = 0
		*fill &= int32(1)<<B - int32(1)
		delta = -int32(16384)
	} else {
		if itheta == int32(16384) {
			imid = 0
			iside = int32(32767)
			*fill &= (int32(1)<<B - int32(1)) << B
			delta = int32(16384)
		} else {
			imid = int32(Opus_bitexact_cos(tls, int16(itheta)))
			iside = int32(Opus_bitexact_cos(tls, int16(int32(16384)-itheta)))
			/* This is the mid vs side allocation that minimizes squared error
			   in that band. */
			delta = (int32(16384) + int32(int16((N-int32(1))<<int32(7)))*int32(int16(Opus_bitexact_log2tan(tls, iside, imid)))) >> int32(15)
		}
	}
	sctx.Finv = inv
	sctx.Fimid = imid
	sctx.Fiside = iside
	sctx.Fdelta = delta
	sctx.Fitheta = itheta
	sctx.Fqalloc = qalloc
}

func quant_band_n1(tls *libc.TLS, ctx *band_ctx, ec *OpusT_ec_ctx, X, Y, lowband *OpusT_celt_norm) uint32 {
	encode := ctx.Fencode // C caches this before either channel can emit bytes.
	x := X
	channels := 1
	if Y != nil {
		channels = 2
	}
	for c := 0; c < channels; c++ {
		sign := int32(0)
		if ctx.Fremaining_bits >= 1<<BITRES {
			if encode != 0 {
				sign = libc.BoolInt32(*x < 0)
				Opus_ec_enc_bits(tls, ec, uint32(sign), 1)
			} else {
				sign = int32(Opus_ec_dec_bits(tls, ec, 1))
			}
			ctx.Fremaining_bits -= 1 << BITRES
		}
		if ctx.Fresynth != 0 {
			if sign != 0 {
				*x = -1
			} else {
				*x = 1
			}
		}
		x = Y
	}
	if lowband != nil {
		*lowband = *X
	}
	return 1
}

// C documentation
//
//	/* This function is responsible for encoding and decoding a mono partition.
//	   It can split the band in two and transmit the energy difference with
//	   the two half-bands. It can be called recursively so bands can end up being
//	   split in 8 parts. */
func quant_partition(tls *libc.TLS, ctx *band_ctx, X *float32, N int32, _b int32, B int32, lowband *float32, LM2 int32, gain OpusT_opus_val32, _fill int32) (r uint32) {
	b := _b
	fill := _fill
	var B0, K, curr_bits, delta, encode, hi, i1, i2, imid, iside, itheta, j, lo, mbits, mid, q, qalloc, sbits, spread, v1, v2, v3, v4 int32
	var Y, next_lowband2 *float32
	var v5 uintptr
	var ec *OpusT_ec_ctx
	var m2 *OpusT_OpusCustomMode
	var cache, cache1, cache2 *byte
	var cm, cm_mask uint32
	var mid1, side OpusT_opus_val32
	var rebalance OpusT_opus_int32
	var tmp, v30 OpusT_opus_val16
	var sctx split_ctx
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = B0, K, Y, cache, cache1, cache2, cm, cm_mask, curr_bits, delta, ec, encode, hi, i1, i2, imid, iside, itheta, j, lo, m2, mbits, mid, mid1, next_lowband2, q, qalloc, rebalance, sbits, side, spread, tmp, v1, v2, v3, v30, v4, v5
	imid = 0
	iside = 0
	B0 = B
	mid1 = float32(0)
	side = float32(0)
	cm = uint32(0)
	Y = nil
	encode = ctx.Fencode
	m2 = ctx.Fm
	i2 = ctx.Fi
	spread = ctx.Fspread
	ec = ctx.Fec
	/* If we need 1.5 more bit than we can produce, split the band in two. */
	cache2 = modePulseCache(m2, (LM2+int32(1))*m2.FnbEBands+i2)
	if LM2 != -int32(1) && b > int32(modePulseByte(cache2, int32(*cache2)))+int32(12) && N > int32(2) {
		next_lowband2 = nil
		N = N >> int32(1)
		Y = celtNormAdd(X, N)
		LM2 = LM2 - int32(1)
		if B == int32(1) {
			fill = fill&int32(1) | fill<<int32(1)
		}
		B = (B + int32(1)) >> int32(1)
		compute_theta(tls, ctx, &sctx, X, Y, N, &b, B, B0, LM2, 0, &fill)
		imid = sctx.Fimid
		iside = sctx.Fiside
		delta = sctx.Fdelta
		itheta = sctx.Fitheta
		qalloc = sctx.Fqalloc
		mid1 = OpusT_opus_val32(float32(1) / float32(32768) * float32(imid))
		side = OpusT_opus_val32(float32(1) / float32(32768) * float32(iside))
		/* Give more bits to low-energy MDCTs than they would otherwise deserve */
		if B0 > int32(1) && itheta&int32(0x3fff) != 0 {
			if itheta > int32(8192) {
				/* Rough approximation for pre-echo masking */
				delta = delta - delta>>(int32(4)-LM2)
			} else {
				/* Corresponds to a forward-masking slope of 1.5 dB per 10 ms */
				if 0 < delta+N<<int32(BITRES)>>(int32(5)-LM2) {
					v1 = 0
				} else {
					v1 = delta + N<<int32(BITRES)>>(int32(5)-LM2)
				}
				delta = v1
			}
		}
		if b < (b-delta)/int32(2) {
			v2 = b
		} else {
			v2 = (b - delta) / int32(2)
		}
		if 0 > v2 {
			v1 = 0
		} else {
			if b < (b-delta)/int32(2) {
				v3 = b
			} else {
				v3 = (b - delta) / int32(2)
			}
			v1 = v3
		}
		mbits = v1
		sbits = b - mbits
		ctx.Fremaining_bits -= qalloc
		if lowband != nil {
			next_lowband2 = celtNormAdd(lowband, N)
		} /* >32-bit split case */
		rebalance = ctx.Fremaining_bits
		if mbits >= sbits {
			cm = quant_partition(tls, ctx, X, N, mbits, B, lowband, LM2, OpusT_opus_val32(gain*mid1), fill)
			rebalance = mbits - (rebalance - ctx.Fremaining_bits)
			if rebalance > int32(3)<<int32(BITRES) && itheta != 0 {
				sbits = sbits + (rebalance - int32(3)<<int32(BITRES))
			}
			cm = cm | quant_partition(tls, ctx, Y, N, sbits, B, next_lowband2, LM2, OpusT_opus_val32(gain*side), fill>>B)<<(B0>>int32(1))
		} else {
			cm = quant_partition(tls, ctx, Y, N, sbits, B, next_lowband2, LM2, OpusT_opus_val32(gain*side), fill>>B) << (B0 >> int32(1))
			rebalance = sbits - (rebalance - ctx.Fremaining_bits)
			if rebalance > int32(3)<<int32(BITRES) && itheta != int32(16384) {
				mbits = mbits + (rebalance - int32(3)<<int32(BITRES))
			}
			cm = cm | quant_partition(tls, ctx, X, N, mbits, B, lowband, LM2, OpusT_opus_val32(gain*mid1), fill)
		}
	} else {
		/* This is the basic no-split case. Reload the cache for each C helper. */
		cache = modePulseCache(m2, (LM2+1)*m2.FnbEBands+i2)
		q = modeBits2Pulses(cache, b)
		cache1 = modePulseCache(m2, (LM2+1)*m2.FnbEBands+i2)
		curr_bits = modePulses2Bits(cache1, q)
		ctx.Fremaining_bits -= curr_bits
		/* Ensures we can never bust the budget */
		for ctx.Fremaining_bits < 0 && q > 0 {
			ctx.Fremaining_bits += curr_bits
			q = q - 1
			cache1 = modePulseCache(m2, (LM2+1)*m2.FnbEBands+i2)
			curr_bits = modePulses2Bits(cache1, q)
			ctx.Fremaining_bits -= curr_bits
		}
		if q != 0 {
			v1 = q
			if v1 < int32(8) {
				v3 = v1
			} else {
				v3 = (int32(8) + v1&int32(7)) << (v1>>int32(3) - int32(1))
			}
			v2 = v3
			K = v2
			/* Finally do the actual quantization */
			if encode != 0 {
				cm = Opus_alg_quant(tls, X, N, K, spread, B, ec, gain, ctx.Fresynth, ctx.Farch)
			} else {
				cm = Opus_alg_unquant(tls, X, N, K, spread, B, ec, gain)
			}
		} else {
			if ctx.Fresynth != 0 {
				/* B can be as large as 16, so this shift might overflow an int on a
				   16-bit platform; use a long to get defined behavior.*/
				cm_mask = uint32(uint64(1)<<B) - uint32(1)
				fill = int32(uint32(fill) & cm_mask)
				if !(fill != 0) {
					clear(unsafe.Slice(X, N))
				} else {
					if lowband == nil {
						/* Noise */
						j = 0
						for {
							if !(j < N) {
								break
							}
							ctx.Fseed = Opus_celt_lcg_rand(tls, ctx.Fseed)
							unsafe.Slice(X, N)[j] = float32(int32(ctx.Fseed) >> int32(20))
							j = j + 1
						}
						cm = cm_mask
					} else {
						/* Folded spectrum */
						j = 0
						for {
							if !(j < N) {
								break
							}
							ctx.Fseed = Opus_celt_lcg_rand(tls, ctx.Fseed)
							/* About 48 dB below the "normal" folding level */
							tmp = float32(1) / float32(256)
							if ctx.Fseed&uint32(0x8000) != 0 {
								v30 = tmp
							} else {
								v30 = -tmp
							}
							tmp = v30
							unsafe.Slice(X, N)[j] = unsafe.Slice(lowband, N)[j] + tmp
							j = j + 1
						}
						cm = uint32(fill)
					}
					Opus_renormalise_vector(tls, X, N, gain, ctx.Farch)
				}
			}
		}
	}
	return cm
}

// C documentation
//
//	/* This function is responsible for encoding and decoding a band for the mono case. */
func quant_band_legacy(tls *libc.TLS, ctx *band_ctx, X uintptr, N, b, B int32, lowband uintptr, LM int32, lowbandOut uintptr, gain float32, scratch uintptr, fill int32) uint32 {
	return quant_band(tls, ctx, (*float32)(unsafe.Pointer(X)), N, b, B, (*float32)(unsafe.Pointer(lowband)), LM, (*float32)(unsafe.Pointer(lowbandOut)), gain, (*float32)(unsafe.Pointer(scratch)), fill)
}

func quant_band(tls *libc.TLS, ctx *band_ctx, X *float32, N int32, b int32, B int32, lowband *float32, LM int32, lowband_out *float32, gain OpusT_opus_val32, lowband_scratch *float32, fill int32) (r uint32) {
	var B0, N0, N_B, N_B0, encode, j, k, longBlocks, recombine, tf_change, time_divide int32
	var cm uint32
	var n1 OpusT_opus_val16
	var v1, v2 OpusT_opus_uint32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = B0, N0, N_B, N_B0, cm, encode, j, k, longBlocks, n1, recombine, tf_change, time_divide, v1, v2
	N0 = N
	N_B = N
	B0 = B
	time_divide = 0
	recombine = 0
	cm = uint32(0)
	encode = ctx.Fencode
	tf_change = ctx.Ftf_change
	longBlocks = libc.BoolInt32(B0 == int32(1))
	v1 = uint32(B)
	_ = v1 > uint32(0)
	v2 = uint32(N_B) / v1
	N_B = int32(v2)
	/* Special case for one sample */
	if N == int32(1) {
		return quant_band_n1(tls, ctx, ctx.Fec, X, nil, lowband_out)
	}
	if tf_change > 0 {
		recombine = tf_change
	}
	/* Band recombining to increase frequency resolution */
	if lowband_scratch != nil && lowband != nil && (recombine != 0 || N_B&int32(1) == 0 && tf_change < 0 || B0 > int32(1)) {
		copy(unsafe.Slice(lowband_scratch, N), unsafe.Slice(lowband, N))
		lowband = lowband_scratch
	}
	k = 0
	for {
		if !(k < recombine) {
			break
		}
		if encode != 0 {
			Opus_haar1(tls, X, N>>k, int32(1)<<k)
		}
		if lowband != nil {
			Opus_haar1(tls, lowband, N>>k, int32(1)<<k)
		}
		fill = int32(bit_interleave_table[fill&int32(0xF)]) | int32(bit_interleave_table[fill>>int32(4)])<<int32(2)
		k = k + 1
	}
	B = B >> recombine
	N_B = N_B << recombine
	/* Increasing the time resolution */
	for N_B&int32(1) == 0 && tf_change < 0 {
		if encode != 0 {
			Opus_haar1(tls, X, N_B, B)
		}
		if lowband != nil {
			Opus_haar1(tls, lowband, N_B, B)
		}
		fill = fill | fill<<B
		B = B << int32(1)
		N_B = N_B >> int32(1)
		time_divide = time_divide + 1
		tf_change = tf_change + 1
	}
	B0 = B
	N_B0 = N_B
	/* Reorganize the samples in time order instead of frequency order */
	if B0 > int32(1) {
		if encode != 0 {
			deinterleave_hadamard(tls, X, N_B>>recombine, B0<<recombine, longBlocks)
		}
		if lowband != nil {
			deinterleave_hadamard(tls, lowband, N_B>>recombine, B0<<recombine, longBlocks)
		}
	}
	cm = quant_partition(tls, ctx, X, N, b, B, lowband, LM, gain, fill)
	/* This code is used by the decoder and by the resynthesis-enabled encoder */
	if ctx.Fresynth != 0 {
		/* Undo the sample reorganization going from time order to frequency order */
		if B0 > int32(1) {
			interleave_hadamard(tls, X, N_B>>recombine, B0<<recombine, longBlocks)
		}
		/* Undo time-freq changes that we did earlier */
		N_B = N_B0
		B = B0
		k = 0
		for {
			if !(k < time_divide) {
				break
			}
			B = B >> int32(1)
			N_B = N_B << int32(1)
			cm = cm | cm>>B
			Opus_haar1(tls, X, N_B, B)
			k = k + 1
		}
		k = 0
		for {
			if !(k < recombine) {
				break
			}
			cm = uint32(bit_deinterleave_table[cm])
			Opus_haar1(tls, X, N0>>k, int32(1)<<k)
			k = k + 1
		}
		B = B << recombine
		/* Scale output for later folding */
		if lowband_out != nil {
			n1 = float32(libc.Xsqrt(tls, float64(N0)))
			j = 0
			for {
				if !(j < N0) {
					break
				}
				unsafe.Slice(lowband_out, N0)[j] = OpusT_opus_val16(n1 * unsafe.Slice(X, N0)[j])
				j = j + 1
			}
		}
		cm = cm & uint32(int32(1)<<B-int32(1))
	}
	return cm
}

var bit_interleave_table = [16]uint8{
	1:  uint8(1),
	2:  uint8(1),
	3:  uint8(1),
	4:  uint8(2),
	5:  uint8(3),
	6:  uint8(3),
	7:  uint8(3),
	8:  uint8(2),
	9:  uint8(3),
	10: uint8(3),
	11: uint8(3),
	12: uint8(2),
	13: uint8(3),
	14: uint8(3),
	15: uint8(3),
}

var bit_deinterleave_table = [16]uint8{
	1:  uint8(0x03),
	2:  uint8(0x0C),
	3:  uint8(0x0F),
	4:  uint8(0x30),
	5:  uint8(0x33),
	6:  uint8(0x3C),
	7:  uint8(0x3F),
	8:  uint8(0xC0),
	9:  uint8(0xC3),
	10: uint8(0xCC),
	11: uint8(0xCF),
	12: uint8(0xF0),
	13: uint8(0xF3),
	14: uint8(0xFC),
	15: uint8(0xFF),
}

// C documentation
//
//	/* This function is responsible for encoding and decoding a band for the stereo case. */
func quant_band_stereo_legacy(tls *libc.TLS, ctx *band_ctx, X, Y uintptr, N, b, B int32, lowband uintptr, LM int32, out, scratch uintptr, fill int32) uint32 {
	return quant_band_stereo(tls, ctx, (*float32)(unsafe.Pointer(X)), (*float32)(unsafe.Pointer(Y)), N, b, B, (*float32)(unsafe.Pointer(lowband)), LM, (*float32)(unsafe.Pointer(out)), (*float32)(unsafe.Pointer(scratch)), fill)
}

func quant_band_stereo(tls *libc.TLS, ctx *band_ctx, X *float32, Y *float32, N int32, _b int32, B int32, lowband *float32, LM int32, lowband_out *float32, lowband_scratch *float32, _fill int32) (r uint32) {
	b := _b
	fill := _fill
	var c, delta, encode, imid, inv, iside, itheta, j, mbits, orig_fill, qalloc, sbits, sign, v3, v4, v5 int32
	var cm uint32
	var x2, y2, v1 *float32
	var ec *OpusT_ec_ctx
	var mid, side OpusT_opus_val32
	var rebalance OpusT_opus_int32
	var tmp OpusT_celt_norm
	var sctx split_ctx
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, cm, delta, ec, encode, imid, inv, iside, itheta, j, mbits, mid, orig_fill, qalloc, rebalance, sbits, side, sign, tmp, x2, y2, v1, v3, v4, v5
	imid = 0
	iside = 0
	inv = 0
	mid = float32(0)
	side = float32(0)
	cm = uint32(0)
	bandContext := ctx
	encode = bandContext.Fencode
	ec = bandContext.Fec
	/* Special case for one sample */
	if N == int32(1) {
		return quant_band_n1(tls, bandContext, bandContext.Fec, X, Y, lowband_out)
	}
	orig_fill = fill
	if encode != 0 {
		if bandContextEnergy(bandContext, bandContext.Fi) < float32(1e-10) || bandContextEnergy(bandContext, bandContext.Fm.FnbEBands+bandContext.Fi) < float32(1e-10) {
			if bandContextEnergy(bandContext, bandContext.Fi) > bandContextEnergy(bandContext, bandContext.Fm.FnbEBands+bandContext.Fi) {
				copy(unsafe.Slice(Y, N), unsafe.Slice(X, N))
			} else {
				copy(unsafe.Slice(X, N), unsafe.Slice(Y, N))
			}
		}
	}
	compute_theta(tls, bandContext, &sctx, X, Y, N, &b, B, B, LM, int32(1), &fill)
	inv = sctx.Finv
	imid = sctx.Fimid
	iside = sctx.Fiside
	delta = sctx.Fdelta
	itheta = sctx.Fitheta
	qalloc = sctx.Fqalloc
	mid = OpusT_opus_val32(float32(1) / float32(32768) * float32(imid))
	side = OpusT_opus_val32(float32(1) / float32(32768) * float32(iside))
	/* This is a special case for N=2 that only works for stereo and takes
	   advantage of the fact that mid and side are orthogonal to encode
	   the side with just one bit. */
	if N == int32(2) {
		sign = 0
		mbits = b
		sbits = 0
		/* Only need one bit for the side. */
		if itheta != 0 && itheta != int32(16384) {
			sbits = int32(1) << int32(BITRES)
		}
		mbits = mbits - sbits
		c = libc.BoolInt32(itheta > int32(8192))
		bandContext.Fremaining_bits -= qalloc + sbits
		if c != 0 {
			v1 = Y
		} else {
			v1 = X
		}
		x2 = v1
		if c != 0 {
			v1 = X
		} else {
			v1 = Y
		}
		y2 = v1
		if sbits != 0 {
			if encode != 0 {
				/* Here we only need to encode a sign for the side. */
				/* FIXME: Need to increase fixed-point precision? */
				sign = libc.BoolInt32(OpusT_celt_norm(*x2*unsafe.Slice(y2, N)[1])-OpusT_celt_norm(unsafe.Slice(x2, N)[1]**y2) < float32(0))
				Opus_ec_enc_bits(tls, (*OpusT_ec_enc)(unsafe.Pointer(ec)), uint32(sign), uint32(1))
			} else {
				sign = int32(Opus_ec_dec_bits(tls, (*OpusT_ec_dec)(unsafe.Pointer(ec)), uint32(1)))
			}
		}
		sign = int32(1) - int32(2)*sign
		/* We use orig_fill here because we want to fold the side, but if
		   itheta==16384, we'll have cleared the low bits of fill. */
		cm = quant_band(tls, bandContext, x2, N, mbits, B, lowband, LM, lowband_out, float32(1), lowband_scratch, orig_fill)
		/* We don't split N=2 bands, so cm is either 1 or 0 (for a fold-collapse),
		   and there's no need to worry about mixing with the other channel. */
		*y2 = OpusT_celt_norm(float32(-sign) * unsafe.Slice(x2, N)[1])
		unsafe.Slice(y2, N)[1] = OpusT_celt_norm(float32(sign) * *x2)
		if bandContext.Fresynth != 0 {
			*X = OpusT_opus_val32(mid * *X)
			unsafe.Slice(X, N)[1] = OpusT_opus_val32(mid * unsafe.Slice(X, N)[1])
			*Y = OpusT_opus_val32(side * *Y)
			unsafe.Slice(Y, N)[1] = OpusT_opus_val32(side * unsafe.Slice(Y, N)[1])
			tmp = *X
			*X = tmp - *Y
			*Y = tmp + *Y
			tmp = unsafe.Slice(X, N)[1]
			unsafe.Slice(X, N)[1] = tmp - unsafe.Slice(Y, N)[1]
			unsafe.Slice(Y, N)[1] = tmp + unsafe.Slice(Y, N)[1]
		}
	} else {
		if b < (b-delta)/int32(2) {
			v4 = b
		} else {
			v4 = (b - delta) / int32(2)
		}
		if 0 > v4 {
			v3 = 0
		} else {
			if b < (b-delta)/int32(2) {
				v5 = b
			} else {
				v5 = (b - delta) / int32(2)
			}
			v3 = v5
		}
		mbits = v3
		sbits = b - mbits
		bandContext.Fremaining_bits -= qalloc
		rebalance = bandContext.Fremaining_bits
		if mbits >= sbits {
			/* In stereo mode, we do not apply a scaling to the mid because we need the normalized
			   mid for folding later. */
			cm = quant_band(tls, bandContext, X, N, mbits, B, lowband, LM, lowband_out, float32(1), lowband_scratch, fill)
			rebalance = mbits - (rebalance - bandContext.Fremaining_bits)
			if rebalance > int32(3)<<int32(BITRES) && itheta != 0 {
				sbits = sbits + (rebalance - int32(3)<<int32(BITRES))
			}
			/* For a stereo split, the high bits of fill are always zero, so no
			   folding will be done to the side. */
			cm = cm | quant_band(tls, bandContext, Y, N, sbits, B, nil, LM, nil, side, nil, fill>>B)
		} else {
			/* For a stereo split, the high bits of fill are always zero, so no
			   folding will be done to the side. */
			cm = quant_band(tls, bandContext, Y, N, sbits, B, nil, LM, nil, side, nil, fill>>B)
			rebalance = sbits - (rebalance - bandContext.Fremaining_bits)
			if rebalance > int32(3)<<int32(BITRES) && itheta != int32(16384) {
				mbits = mbits + (rebalance - int32(3)<<int32(BITRES))
			}
			/* In stereo mode, we do not apply a scaling to the mid because we need the normalized
			   mid for folding later. */
			cm = cm | quant_band(tls, bandContext, X, N, mbits, B, lowband, LM, lowband_out, float32(1), lowband_scratch, fill)
		}
	}
	/* This code is used by the decoder and by the resynthesis-enabled encoder */
	if bandContext.Fresynth != 0 {
		if N != int32(2) {
			stereo_merge(tls, X, Y, mid, N, bandContext.Farch)
		}
		if inv != 0 {
			j = 0
			for {
				if !(j < N) {
					break
				}
				unsafe.Slice(Y, N)[j] = -unsafe.Slice(Y, N)[j]
				j = j + 1
			}
		}
	}
	return cm
}

func special_hybrid_folding(tls *libc.TLS, bands *OpusT_opus_int16, norm, norm2 *OpusT_celt_norm, start, M, dual_stereo int32) {
	eBands := unsafe.Slice(bands, start+3)
	n1 := M * (int32(eBands[start+1]) - int32(eBands[start]))
	n2 := M * (int32(eBands[start+2]) - int32(eBands[start+1]))
	// Valid folding widths satisfy n1 <= n2 <= 2*n1; source ends at n1.
	// CELT-only equal-width bands copy nothing.
	if n1 == n2 {
		return
	}
	first := unsafe.Slice(norm, n2)
	copy(first[n1:n2], first[2*n1-n2:n1])
	if dual_stereo != 0 {
		second := unsafe.Slice(norm2, n2)
		copy(second[n1:n2], second[2*n1-n2:n1])
	}
}

func quantAllBandsByteStorage(enabled int32) []byte {
	if enabled == 0 {
		return nil
	}
	return make([]byte, 1275)
}
func quantAllBandsByteSave(saved []byte, ec *OpusT_ec_ctx, start, count int32) []byte {
	if count == 0 {
		return nil
	}
	window := unsafe.Slice(ec.Fbuf, ec.Fstorage)[start : start+count]
	copy(saved[:count], window)
	return window
}
func quantAllBandsByteRestore(window, saved []byte, count int32) { copy(window[:count], saved[:count]) }

func quantAllBandsNormCopy(saved, norm *float32, offset, N int32, restore bool) {
	if N == 0 {
		return
	}
	values := unsafe.Slice(norm, offset+N)[offset : offset+N]
	if restore {
		copy(values, unsafe.Slice(saved, N))
	} else {
		copy(unsafe.Slice(saved, N), values)
	}
}

func quantAllBandsDot(left, right *float32, N int32) float32 {
	if N == 0 {
		return 0
	}
	x, y := unsafe.Slice(left, N), unsafe.Slice(right, N)
	sum := float32(0)
	for i := int32(0); i < N; i++ {
		sum = sum + float32(x[i]*y[i])
	}
	return sum
}

func quantAllBandsCopy(dst, src *float32, N int32) {
	if N == 0 {
		return
	}
	copy(unsafe.Slice(dst, N), unsafe.Slice(src, N))
}

func quantAllBandsNormLength(bands *int16, index, M, channels, offset int32) int32 {
	return channels * (M*quantAllBandsBoundary(bands, index) - offset)
}

func quantAllBandsBoundary(bands *int16, index int32) int32 {
	return int32(unsafe.Slice(bands, index+1)[index])
}

func quantAllBandsMask(masks *byte, index int32) uint32 {
	return uint32(unsafe.Slice(masks, index+1)[index])
}
func quantAllBandsMaskStore(masks *byte, band, channels int32, left, right uint32) {
	first := band * channels
	unsafe.Slice(masks, first+1)[first] = uint8(left)
	last := first + channels - 1
	unsafe.Slice(masks, last+1)[last] = uint8(right)
}

func quantAllBandsPulse(pulses *int32, band int32) int32 { return unsafe.Slice(pulses, band+1)[band] }

func quantAllBandsTF(flags *int32, band int32) int32 { return unsafe.Slice(flags, band+1)[band] }

func quantAllBandsSetEnergy(ctx *band_ctx, energy *float32) { ctx.FbandE = energy }
func quantAllBandsChannelWeights(tls *libc.TLS, mode *OpusT_OpusCustomMode, energy *float32, band int32, w *[2]float32) {
	left := unsafe.Slice(energy, band+1)[band]
	rightIndex := band + mode.FnbEBands
	right := unsafe.Slice(energy, rightIndex+1)[rightIndex]
	compute_channel_weights(tls, left, right, w)
}

func quantAllBandsReadSeed(ctx *band_ctx, seed *uint32)  { ctx.Fseed = *seed }
func quantAllBandsWriteSeed(seed *uint32, ctx *band_ctx) { *seed = ctx.Fseed }

func quantAllBandsSetEntropy(ctx *band_ctx, ec *OpusT_ec_ctx) { ctx.Fec = ec }

func quantAllBandsSetMode(ctx *band_ctx, mode *OpusT_OpusCustomMode) { ctx.Fm = mode }

// Legacy pointer ABI: retain Go-owned caller scratch across recursive calls.
//
//go:uintptrescapes
func Opus_quant_all_bands(tls *libc.TLS, encode int32, m uintptr, start int32, end int32, X_ uintptr, Y_ uintptr, collapse_masks uintptr, bandE uintptr, pulses uintptr, shortBlocks int32, spread int32, dual_stereo int32, intensity int32, tf_res uintptr, total_bits OpusT_opus_int32, balance OpusT_opus_int32, ec uintptr, LM int32, codedBands int32, seed uintptr, complexity int32, arch int32, disable_inv int32) {
	quant_all_bands(tls, encode, (*OpusT_OpusCustomMode)(unsafe.Pointer(m)), start, end, X_, Y_, (*byte)(unsafe.Pointer(collapse_masks)), (*float32)(unsafe.Pointer(bandE)), (*int32)(unsafe.Pointer(pulses)), shortBlocks, spread, dual_stereo, intensity, (*int32)(unsafe.Pointer(tf_res)), total_bits, balance, (*OpusT_ec_ctx)(unsafe.Pointer(ec)), LM, codedBands, (*uint32)(unsafe.Pointer(seed)), complexity, arch, disable_inv)
}

// Internal owners are migrated independently of the remaining legacy views.
// The integer spectrum/mask/scratch arguments still require escape retention.
//
//go:uintptrescapes
func quant_all_bands(tls *libc.TLS, encode int32, m *OpusT_OpusCustomMode, start, end int32, X_, Y_ uintptr, collapse_masks *byte, bandE *float32, pulses *int32, shortBlocks, spread, dual_stereo, intensity int32, tf_res *int32, total_bits, balance int32, ec *OpusT_ec_ctx, LM, codedBands int32, seed *uint32, complexity, arch, disable_inv int32) {
	/* ctx keeps the transpiled uintptr calling convention into
	   quant_band/quant_band_stereo, so it is allocated on the C heap:
	   a Go stack local whose address is laundered
	   through uintptr would be left behind by a goroutine stack growth in
	   the PVQ recursion. */
	// Go storage scans the context's mode, codec and energy references.
	ctx := new(band_ctx)
	// Channel weights now stay Go-visible across calls and stack growth.
	w := new([2]OpusT_opus_val16)
	var B, C, M, N1, b, effective_lowband, fold_end, fold_i, fold_start, i, i1, j, last, lowband_offset, nend_bytes, norm_offset, nstart_bytes, resynth, resynth_alloc, save_bytes, tf_change, theta_rdo, update_lowband, v1, v183, v196, v201, v203, v207, v6 int32
	var eBands *int16
	var bytes_buf, bytes_save []byte
	var X, X_save, X_save2, Y, Y_save, Y_save2, _lowband_scratch, _norm, _saved_stack, lowband_scratch, norm, norm2, norm_save2, st, v11, v13, v15, v17, v19, v2, v21, v23, v25, v4, v7, v9 uintptr
	var cm, cm2, x_cm, y_cm, v217 uint32
	var ctx_save, ctx_save2 band_ctx
	var curr_balance, remaining_bits, tell, v204, v205 OpusT_opus_int32
	var dist0, dist1, xy, v229, v232 OpusT_opus_val32
	var ec_save, ec_save2 OpusT_ec_ctx
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = B, C, M, N1, X, X_save, X_save2, Y, Y_save, Y_save2, _lowband_scratch, _norm, _saved_stack, b, bytes_buf, bytes_save, cm, cm2, ctx_save, ctx_save2, curr_balance, dist0, dist1, eBands, ec_save, ec_save2, effective_lowband, fold_end, fold_i, fold_start, i, i1, j, last, lowband_offset, lowband_scratch, nend_bytes, norm, norm2, norm_offset, norm_save2, nstart_bytes, remaining_bits, resynth, resynth_alloc, save_bytes, st, tell, tf_change, theta_rdo, update_lowband, x_cm, xy, y_cm, v1, v11, v13, v15, v17, v183, v19, v196, v2, v201, v203, v204, v205, v207, v21, v217, v229, v23, v232, v25, v4, v6, v7, v9
	eBands = m.FeBands
	update_lowband = int32(1)
	if Y_ != uintptr(uint32(0)) {
		v1 = int32(2)
	} else {
		v1 = int32(1)
	}
	C = v1
	theta_rdo = libc.BoolInt32(encode != 0 && Y_ != uintptr(uint32(0)) && !(dual_stereo != 0) && complexity >= int32(8))
	resynth = libc.BoolInt32(!(encode != 0) || theta_rdo != 0)
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v2 = libc.Xmalloc(tls, uint64(16))
		st = v2
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v4 = st
	_saved_stack = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v4)).Fglobal_stack
	M = int32(1) << LM
	if shortBlocks != 0 {
		v1 = M
	} else {
		v1 = int32(1)
	}
	B = v1
	norm_offset = M * quantAllBandsBoundary(eBands, start)
	/* No need to allocate norm for the last band because we don't need an
	   output in that band. */
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v2 = libc.Xmalloc(tls, uint64(16))
		st = v2
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v4 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v7 = libc.Xmalloc(tls, uint64(16))
		st = v7
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v9 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v4)).Fglobal_stack += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v9)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	if !(int64(int32(uint64(uint32(quantAllBandsNormLength(eBands, m.FnbEBands-1, M, C, norm_offset)))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5312, int32(1638))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack += uintptr(uint64(uint32(quantAllBandsNormLength(eBands, m.FnbEBands-1, M, C, norm_offset))) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v23 = libc.Xmalloc(tls, uint64(16))
		st = v23
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v25 = st
	_norm = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v25)).Fglobal_stack - uintptr(uint64(uint32(quantAllBandsNormLength(eBands, m.FnbEBands-1, M, C, norm_offset)))*(uint64(4)/uint64(1)))
	norm = _norm
	norm2 = norm + uintptr(M*quantAllBandsBoundary(eBands, m.FnbEBands-1))*4 - uintptr(norm_offset)*4
	/* For decoding, we can use the last band as scratch space because we don't need that
	   scratch space for the last band and we don't care about the data there until we're
	   decoding the last band. */
	if encode != 0 && resynth != 0 {
		resynth_alloc = M * (quantAllBandsBoundary(eBands, m.FnbEBands) - quantAllBandsBoundary(eBands, m.FnbEBands-1))
	} else {
		resynth_alloc = ALLOC_NONE
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v2 = libc.Xmalloc(tls, uint64(16))
		st = v2
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v4 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v7 = libc.Xmalloc(tls, uint64(16))
		st = v7
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v9 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v4)).Fglobal_stack += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v9)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	if !(int64(int32(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5312, int32(1649))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack += uintptr(uint64(uint32(resynth_alloc)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v23 = libc.Xmalloc(tls, uint64(16))
		st = v23
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v25 = st
	_lowband_scratch = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v25)).Fglobal_stack - uintptr(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))
	if encode != 0 && resynth != 0 {
		lowband_scratch = _lowband_scratch
	} else {
		lowband_scratch = X_ + uintptr(M*quantAllBandsBoundary(eBands, m.FeffEBands-1))*4
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v2 = libc.Xmalloc(tls, uint64(16))
		st = v2
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v4 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v7 = libc.Xmalloc(tls, uint64(16))
		st = v7
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v9 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v4)).Fglobal_stack += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v9)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	if !(int64(int32(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5312, int32(1654))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack += uintptr(uint64(uint32(resynth_alloc)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v23 = libc.Xmalloc(tls, uint64(16))
		st = v23
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v25 = st
	X_save = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v25)).Fglobal_stack - uintptr(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v2 = libc.Xmalloc(tls, uint64(16))
		st = v2
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v4 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v7 = libc.Xmalloc(tls, uint64(16))
		st = v7
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v9 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v4)).Fglobal_stack += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v9)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	if !(int64(int32(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5312, int32(1655))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack += uintptr(uint64(uint32(resynth_alloc)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v23 = libc.Xmalloc(tls, uint64(16))
		st = v23
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v25 = st
	Y_save = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v25)).Fglobal_stack - uintptr(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v2 = libc.Xmalloc(tls, uint64(16))
		st = v2
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v4 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v7 = libc.Xmalloc(tls, uint64(16))
		st = v7
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v9 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v4)).Fglobal_stack += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v9)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	if !(int64(int32(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5312, int32(1656))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack += uintptr(uint64(uint32(resynth_alloc)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v23 = libc.Xmalloc(tls, uint64(16))
		st = v23
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v25 = st
	X_save2 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v25)).Fglobal_stack - uintptr(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v2 = libc.Xmalloc(tls, uint64(16))
		st = v2
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v4 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v7 = libc.Xmalloc(tls, uint64(16))
		st = v7
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v9 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v4)).Fglobal_stack += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v9)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	if !(int64(int32(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5312, int32(1657))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack += uintptr(uint64(uint32(resynth_alloc)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v23 = libc.Xmalloc(tls, uint64(16))
		st = v23
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v25 = st
	Y_save2 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v25)).Fglobal_stack - uintptr(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v2 = libc.Xmalloc(tls, uint64(16))
		st = v2
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v4 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v7 = libc.Xmalloc(tls, uint64(16))
		st = v7
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v9 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v4)).Fglobal_stack += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v9)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	if !(int64(int32(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5312, int32(1658))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack += uintptr(uint64(uint32(resynth_alloc)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v23 = libc.Xmalloc(tls, uint64(16))
		st = v23
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v25 = st
	norm_save2 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v25)).Fglobal_stack - uintptr(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))
	lowband_offset = 0
	quantAllBandsSetEnergy(ctx, bandE)
	quantAllBandsSetEntropy(ctx, ec)
	ctx.Fencode = encode
	ctx.Fintensity = intensity
	quantAllBandsSetMode(ctx, m)
	quantAllBandsReadSeed(ctx, seed)
	ctx.Fspread = spread
	ctx.Farch = arch
	ctx.Fdisable_inv = disable_inv
	ctx.Fresynth = resynth
	ctx.Ftheta_round = 0
	bytes_save = quantAllBandsByteStorage(theta_rdo)
	/* Avoid injecting noise in the first band on transients. */
	ctx.Favoid_split_noise = libc.BoolInt32(B > int32(1))
	i1 = start
	for {
		if !(i1 < end) {
			break
		}
		effective_lowband = -int32(1)
		tf_change = 0
		ctx.Fi = i1
		last = libc.BoolInt32(i1 == end-int32(1))
		X = X_ + uintptr(M*quantAllBandsBoundary(eBands, i1))*4
		if Y_ != uintptr(uint32(0)) {
			Y = Y_ + uintptr(M*quantAllBandsBoundary(eBands, i1))*4
		} else {
			Y = uintptr(uint32(0))
		}
		N1 = M*quantAllBandsBoundary(eBands, i1+1) - M*quantAllBandsBoundary(eBands, i1)
		if !(N1 > int32(0)) {
			Opus_celt_fatal(tls, __ccgo_ts+5488, __ccgo_ts+5312, int32(1705))
		}
		tell = int32(Opus_ec_tell_frac(tls, (*OpusT_ec_ctx)(unsafe.Pointer(ec))))
		/* Compute how many bits we want to allocate to this band */
		if i1 != start {
			balance = balance - tell
		}
		remaining_bits = total_bits - tell - int32(1)
		ctx.Fremaining_bits = remaining_bits
		if i1 <= codedBands-int32(1) {
			if int32(3) < codedBands-i1 {
				v1 = int32(3)
			} else {
				v1 = codedBands - i1
			}
			v204 = v1
			_ = v204 > int32(0)
			v205 = balance / v204
			curr_balance = v205
			if remaining_bits+1 < quantAllBandsPulse(pulses, i1)+curr_balance {
				v183 = remaining_bits + 1
			} else {
				v183 = quantAllBandsPulse(pulses, i1) + curr_balance
			}
			if 16383 < v183 {
				v6 = 16383
			} else {
				if remaining_bits+1 < quantAllBandsPulse(pulses, i1)+curr_balance {
					v196 = remaining_bits + 1
				} else {
					v196 = quantAllBandsPulse(pulses, i1) + curr_balance
				}
				v6 = v196
			}
			if 0 > v6 {
				v1 = 0
			} else {
				if remaining_bits+1 < quantAllBandsPulse(pulses, i1)+curr_balance {
					v203 = remaining_bits + 1
				} else {
					v203 = quantAllBandsPulse(pulses, i1) + curr_balance
				}
				if 16383 < v203 {
					v201 = 16383
				} else {
					if remaining_bits+1 < quantAllBandsPulse(pulses, i1)+curr_balance {
						v207 = remaining_bits + 1
					} else {
						v207 = quantAllBandsPulse(pulses, i1) + curr_balance
					}
					v201 = v207
				}
				v1 = v201
			}
			b = v1
		} else {
			b = 0
		}
		if resynth != 0 && (M*quantAllBandsBoundary(eBands, i1)-N1 >= M*quantAllBandsBoundary(eBands, start) || i1 == start+1) && (update_lowband != 0 || lowband_offset == 0) {
			lowband_offset = i1
		}
		if i1 == start+int32(1) {
			special_hybrid_folding(tls, (*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands, (*OpusT_celt_norm)(unsafe.Pointer(norm)), (*OpusT_celt_norm)(unsafe.Pointer(norm2)), start, M, dual_stereo)
		}
		tf_change = quantAllBandsTF(tf_res, i1)
		ctx.Ftf_change = tf_change
		if i1 >= (*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeffEBands {
			X = norm
			if Y_ != uintptr(uint32(0)) {
				Y = norm
			}
			lowband_scratch = uintptr(uint32(0))
		}
		if last != 0 && !(theta_rdo != 0) {
			lowband_scratch = uintptr(uint32(0))
		}
		/* Get a conservative estimate of the collapse_mask's for the bands we're
		   going to be folding from. */
		if lowband_offset != 0 && (spread != int32(SPREAD_AGGRESSIVE) || B > int32(1) || tf_change < 0) {
			/* This ensures we never repeat spectral content within one band */
			if 0 > M*quantAllBandsBoundary(eBands, lowband_offset)-norm_offset-N1 {
				v1 = 0
			} else {
				v1 = M*quantAllBandsBoundary(eBands, lowband_offset) - norm_offset - N1
			}
			effective_lowband = v1
			fold_start = lowband_offset
			for {
				fold_start = fold_start - 1
				v1 = fold_start
				if !(M*quantAllBandsBoundary(eBands, v1) > effective_lowband+norm_offset) {
					break
				}
			}
			fold_end = lowband_offset - int32(1)
			for {
				fold_end = fold_end + 1
				v1 = fold_end
				if !(v1 < i1 && M*quantAllBandsBoundary(eBands, fold_end) < effective_lowband+norm_offset+N1) {
					break
				}
			}
			v217 = uint32(0)
			y_cm = v217
			x_cm = v217
			fold_i = fold_start
			for {
				x_cm = x_cm | quantAllBandsMask(collapse_masks, fold_i*C)
				y_cm = y_cm | quantAllBandsMask(collapse_masks, fold_i*C+C-1)
				fold_i = fold_i + 1
				v1 = fold_i
				if !(v1 < fold_end) {
					break
				}
			}
		} else {
			v217 = uint32(int32(1)<<B - int32(1))
			y_cm = v217
			x_cm = v217
		}
		if dual_stereo != 0 && i1 == intensity {
			/* Switch off dual stereo to do intensity. */
			dual_stereo = 0
			if resynth != 0 {
				j = 0
				for {
					if !(j < M*quantAllBandsBoundary(eBands, i1)-norm_offset) {
						break
					}
					*(*OpusT_celt_norm)(unsafe.Pointer(norm + uintptr(j)*4)) = float32(float32(0.5) * (*(*OpusT_celt_norm)(unsafe.Pointer(norm + uintptr(j)*4)) + *(*OpusT_celt_norm)(unsafe.Pointer(norm2 + uintptr(j)*4))))
					j = j + 1
				}
			}
		}
		if dual_stereo != 0 {
			if effective_lowband != -int32(1) {
				v2 = norm + uintptr(effective_lowband)*4
			} else {
				v2 = uintptr(uint32(0))
			}
			if last != 0 {
				v4 = uintptr(uint32(0))
			} else {
				v4 = norm + uintptr(M*quantAllBandsBoundary(eBands, i1))*4 - uintptr(norm_offset)*4
			}
			x_cm = quant_band_legacy(tls, ctx, X, N1, b/int32(2), B, v2, LM, v4, float32(1), lowband_scratch, int32(x_cm))
			if effective_lowband != -int32(1) {
				v2 = norm2 + uintptr(effective_lowband)*4
			} else {
				v2 = uintptr(uint32(0))
			}
			if last != 0 {
				v4 = uintptr(uint32(0))
			} else {
				v4 = norm2 + uintptr(M*quantAllBandsBoundary(eBands, i1))*4 - uintptr(norm_offset)*4
			}
			y_cm = quant_band_legacy(tls, ctx, Y, N1, b/int32(2), B, v2, LM, v4, float32(1), lowband_scratch, int32(y_cm))
		} else {
			if Y != uintptr(uint32(0)) {
				if theta_rdo != 0 && i1 < intensity {
					quantAllBandsChannelWeights(tls, m, bandE, i1, w)
					/* Make a copy. */
					cm = x_cm | y_cm
					ec_save = *(*OpusT_ec_ctx)(unsafe.Pointer(ec))
					ctx_save = *ctx
					quantAllBandsCopy((*float32)(unsafe.Pointer(X_save)), (*float32)(unsafe.Pointer(X)), N1)
					quantAllBandsCopy((*float32)(unsafe.Pointer(Y_save)), (*float32)(unsafe.Pointer(Y)), N1)
					/* Encode and round down. */
					ctx.Ftheta_round = -1
					if effective_lowband != -1 {
						v2 = norm + uintptr(effective_lowband)*4
					} else {
						v2 = uintptr(uint32(0))
					}
					if last != 0 {
						v4 = uintptr(uint32(0))
					} else {
						v4 = norm + uintptr(M*quantAllBandsBoundary(eBands, i1))*4 - uintptr(norm_offset)*4
					}
					x_cm = quant_band_stereo_legacy(tls, ctx, X, Y, N1, b, B, v2, LM, v4, lowband_scratch, int32(cm))
					_ = arch
					v229 = quantAllBandsDot((*float32)(unsafe.Pointer(X_save)), (*float32)(unsafe.Pointer(X)), N1)
					v232 = quantAllBandsDot((*float32)(unsafe.Pointer(Y_save)), (*float32)(unsafe.Pointer(Y)), N1)
					dist0 = OpusT_opus_val16(w[0]*v229) + OpusT_opus_val16(w[1]*v232)
					/* Save first result. */
					cm2 = x_cm
					ec_save2 = *(*OpusT_ec_ctx)(unsafe.Pointer(ec))
					ctx_save2 = *ctx
					quantAllBandsCopy((*float32)(unsafe.Pointer(X_save2)), (*float32)(unsafe.Pointer(X)), N1)
					quantAllBandsCopy((*float32)(unsafe.Pointer(Y_save2)), (*float32)(unsafe.Pointer(Y)), N1)
					if !(last != 0) {
						quantAllBandsNormCopy((*float32)(unsafe.Pointer(norm_save2)), (*float32)(unsafe.Pointer(norm)), M*quantAllBandsBoundary(eBands, i1)-norm_offset, N1, false)
					}
					nstart_bytes = int32(ec_save.Foffs)
					nend_bytes = int32(ec_save.Fstorage)
					save_bytes = nend_bytes - nstart_bytes
					bytes_buf = quantAllBandsByteSave(bytes_save, &ec_save, nstart_bytes, save_bytes)
					/* Restore */
					*(*OpusT_ec_ctx)(unsafe.Pointer(ec)) = ec_save
					*ctx = ctx_save
					quantAllBandsCopy((*float32)(unsafe.Pointer(X)), (*float32)(unsafe.Pointer(X_save)), N1)
					quantAllBandsCopy((*float32)(unsafe.Pointer(Y)), (*float32)(unsafe.Pointer(Y_save)), N1)
					if i1 == start+int32(1) {
						special_hybrid_folding(tls, (*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands, (*OpusT_celt_norm)(unsafe.Pointer(norm)), (*OpusT_celt_norm)(unsafe.Pointer(norm2)), start, M, dual_stereo)
					}
					/* Encode and round up. */
					ctx.Ftheta_round = 1
					if effective_lowband != -1 {
						v2 = norm + uintptr(effective_lowband)*4
					} else {
						v2 = uintptr(uint32(0))
					}
					if last != 0 {
						v4 = uintptr(uint32(0))
					} else {
						v4 = norm + uintptr(M*quantAllBandsBoundary(eBands, i1))*4 - uintptr(norm_offset)*4
					}
					x_cm = quant_band_stereo_legacy(tls, ctx, X, Y, N1, b, B, v2, LM, v4, lowband_scratch, int32(cm))
					_ = arch
					v229 = quantAllBandsDot((*float32)(unsafe.Pointer(X_save)), (*float32)(unsafe.Pointer(X)), N1)
					v232 = quantAllBandsDot((*float32)(unsafe.Pointer(Y_save)), (*float32)(unsafe.Pointer(Y)), N1)
					dist1 = OpusT_opus_val16(w[0]*v229) + OpusT_opus_val16(w[1]*v232)
					if dist0 >= dist1 {
						x_cm = cm2
						*(*OpusT_ec_ctx)(unsafe.Pointer(ec)) = ec_save2
						*ctx = ctx_save2
						quantAllBandsCopy((*float32)(unsafe.Pointer(X)), (*float32)(unsafe.Pointer(X_save2)), N1)
						quantAllBandsCopy((*float32)(unsafe.Pointer(Y)), (*float32)(unsafe.Pointer(Y_save2)), N1)
						if !(last != 0) {
							quantAllBandsNormCopy((*float32)(unsafe.Pointer(norm_save2)), (*float32)(unsafe.Pointer(norm)), M*quantAllBandsBoundary(eBands, i1)-norm_offset, N1, true)
						}
						quantAllBandsByteRestore(bytes_buf, bytes_save, save_bytes)
					}
				} else {
					ctx.Ftheta_round = 0
					if effective_lowband != -int32(1) {
						v2 = norm + uintptr(effective_lowband)*4
					} else {
						v2 = uintptr(uint32(0))
					}
					if last != 0 {
						v4 = uintptr(uint32(0))
					} else {
						v4 = norm + uintptr(M*quantAllBandsBoundary(eBands, i1))*4 - uintptr(norm_offset)*4
					}
					x_cm = quant_band_stereo_legacy(tls, ctx, X, Y, N1, b, B, v2, LM, v4, lowband_scratch, int32(x_cm|y_cm))
				}
			} else {
				if effective_lowband != -int32(1) {
					v2 = norm + uintptr(effective_lowband)*4
				} else {
					v2 = uintptr(uint32(0))
				}
				if last != 0 {
					v4 = uintptr(uint32(0))
				} else {
					v4 = norm + uintptr(M*quantAllBandsBoundary(eBands, i1))*4 - uintptr(norm_offset)*4
				}
				x_cm = quant_band_legacy(tls, ctx, X, N1, b, B, v2, LM, v4, float32(1), lowband_scratch, int32(x_cm|y_cm))
			}
			y_cm = x_cm
		}
		quantAllBandsMaskStore(collapse_masks, i1, C, x_cm, y_cm)
		balance = balance + (quantAllBandsPulse(pulses, i1) + tell)
		/* Update the folding position only as long as we have 1 bit/sample depth. */
		update_lowband = libc.BoolInt32(b > N1<<int32(BITRES))
		/* We only need to avoid noise on a split for the first band. After that, we
		   have folding. */
		ctx.Favoid_split_noise = 0
		i1 = i1 + 1
	}
	quantAllBandsWriteSeed(seed, ctx)
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v2 = libc.Xmalloc(tls, uint64(16))
		st = v2
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v4 = st
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v4)).Fglobal_stack = _saved_stack
}

const EPSILON4 = "1e-15f"
const NORM_SCALING2 = "1.f"
const Q31ONE4 = "1.0f"

var log2_x_norm_coeff15 = [8]float32{
	0: float32(1),
	1: float32(0.8888888955116272),
	2: float32(0.8),
	3: float32(0.7272727489471436),
	4: float32(0.6666666865348816),
	5: float32(0.6153846383094788),
	6: float32(0.5714285969734192),
	7: float32(0.5333333611488342),
}
var log2_y_norm_coeff15 = [8]float32{
	1: float32(0.1699250042438507),
	2: float32(0.32192808389663696),
	3: float32(0.45943161845207214),
	4: float32(0.5849624872207642),
	5: float32(0.7004396915435791),
	6: float32(0.8073549270629883),
	7: float32(0.9068905711174011),
}

// C documentation
//
//	/* Forward MDCT trashes the input array */
