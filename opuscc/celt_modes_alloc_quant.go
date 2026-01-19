// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

package opuscc

import (
	"reflect"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

var _ reflect.Type
var _ unsafe.Pointer

func Opus_opus_custom_mode_create(tls *libc.TLS, Fs OpusT_opus_int32, frame_size int32) (uintptr, error) {
	var i, j int32
	_, _ = i, j
	i = 0
	for {
		if !(i < int32(TOTAL_MODES)) {
			break
		}
		j = 0
		for {
			if !(j < int32(4)) {
				break
			}
			if Fs == (*OpusT_OpusCustomMode)(unsafe.Pointer(static_mode_list[i])).FFs && frame_size<<j == (*OpusT_OpusCustomMode)(unsafe.Pointer(static_mode_list[i])).FshortMdctSize*(*OpusT_OpusCustomMode)(unsafe.Pointer(static_mode_list[i])).FnbShortMdcts {
				return static_mode_list[i], nil
			}
			goto _2
		_2:
			;
			j = j + 1
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return uintptr(uint32(0)), opusErrorFromCode(-int32(1))
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

func exp_rotation1(tls *libc.TLS, X uintptr, len1 int32, stride int32, c OpusT_opus_val16, s OpusT_opus_val16) {
	var Xptr, v2 uintptr
	var i int32
	var ms OpusT_opus_val16
	var x1, x11, x2, x21 OpusT_celt_norm
	_, _, _, _, _, _, _, _ = Xptr, i, ms, x1, x11, x2, x21, v2
	Xptr = X
	ms = -s
	i = 0
	for {
		if !(i < len1-stride) {
			break
		}
		x1 = *(*OpusT_celt_norm)(unsafe.Pointer(Xptr))
		x2 = *(*OpusT_celt_norm)(unsafe.Pointer(Xptr + uintptr(stride)*4))
		*(*OpusT_celt_norm)(unsafe.Pointer(Xptr + uintptr(stride)*4)) = OpusT_opus_val32(c*x2) + OpusT_opus_val32(s*x1)
		v2 = Xptr
		Xptr += 4
		*(*OpusT_celt_norm)(unsafe.Pointer(v2)) = OpusT_opus_val32(c*x1) + OpusT_opus_val32(ms*x2)
		goto _1
	_1:
		;
		i = i + 1
	}
	Xptr = X + uintptr(len1-int32(2)*stride-int32(1))*4
	i = len1 - int32(2)*stride - int32(1)
	for {
		if !(i >= 0) {
			break
		}
		x11 = *(*OpusT_celt_norm)(unsafe.Pointer(Xptr))
		x21 = *(*OpusT_celt_norm)(unsafe.Pointer(Xptr + uintptr(stride)*4))
		*(*OpusT_celt_norm)(unsafe.Pointer(Xptr + uintptr(stride)*4)) = OpusT_opus_val32(c*x21) + OpusT_opus_val32(s*x11)
		v2 = Xptr
		Xptr -= 4
		*(*OpusT_celt_norm)(unsafe.Pointer(v2)) = OpusT_opus_val32(c*x11) + OpusT_opus_val32(ms*x21)
		goto _3
	_3:
		;
		i = i - 1
	}
}

func Opus_exp_rotation(tls *libc.TLS, X uintptr, len1 int32, dir int32, stride int32, K int32, spread int32) {
	var c, gain, s, theta OpusT_opus_val16
	var factor, i, stride2 int32
	var v1, v2 OpusT_opus_uint32
	_, _, _, _, _, _, _, _, _ = c, factor, gain, i, s, stride2, theta, v1, v2
	stride2 = 0
	if int32(2)*K >= len1 || spread == SPREAD_NONE {
		return
	}
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
	goto _3
_3:
	len1 = int32(v2)
	i = 0
	for {
		if !(i < stride) {
			break
		}
		if dir < 0 {
			if stride2 != 0 {
				exp_rotation1(tls, X+uintptr(i*len1)*4, len1, stride2, s, c)
			}
			exp_rotation1(tls, X+uintptr(i*len1)*4, len1, int32(1), c, s)
		} else {
			exp_rotation1(tls, X+uintptr(i*len1)*4, len1, int32(1), c, -s)
			if stride2 != 0 {
				exp_rotation1(tls, X+uintptr(i*len1)*4, len1, stride2, s, -c)
			}
		}
		goto _4
	_4:
		;
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
func normalise_residual(tls *libc.TLS, iy uintptr, X uintptr, N int32, Ryy OpusT_opus_val32, gain OpusT_opus_val32, shift int32) {
	var g, t OpusT_opus_val32
	var i, v1 int32
	_, _, _, _ = g, i, t, v1
	t = Ryy
	g = float32(float32(1) / float32(libc.Xsqrt(tls, float64(t))) * gain)
	i = 0
	_ = shift
	for {
		*(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(i)*4)) = OpusT_opus_val32(float32(*(*int32)(unsafe.Pointer(iy + uintptr(i)*4))) * g)
		goto _2
	_2:
		;
		i = i + 1
		v1 = i
		if !(v1 < N) {
			break
		}
	}
}

func extract_collapse_mask(tls *libc.TLS, iy uintptr, N int32, B int32) (r uint32) {
	var N0, i, j, v4 int32
	var collapse_mask, tmp uint32
	var v1, v2 OpusT_opus_uint32
	_, _, _, _, _, _, _, _ = N0, collapse_mask, i, j, tmp, v1, v2, v4
	if B <= int32(1) {
		return uint32(1)
	}
	/*NOTE: As a minor optimization, we could be passing around log2(B), not B, for both this and for
	  exp_rotation().*/
	v1 = uint32(B)
	_ = v1 > uint32(0)
	v2 = uint32(N) / v1
	goto _3
_3:
	N0 = int32(v2)
	collapse_mask = uint32(0)
	i = 0
	for {
		tmp = uint32(0)
		j = 0
		for {
			tmp = tmp | uint32(*(*int32)(unsafe.Pointer(iy + uintptr(i*N0+j)*4)))
			goto _7
		_7:
			;
			j = j + 1
			v4 = j
			if !(v4 < N0) {
				break
			}
		}
		collapse_mask = collapse_mask | uint32(libc.BoolInt32(tmp != uint32(0))<<i)
		goto _5
	_5:
		;
		i = i + 1
		v4 = i
		if !(v4 < B) {
			break
		}
	}
	return collapse_mask
}

func Opus_op_pvq_search_c(tls *libc.TLS, X uintptr, iy uintptr, K int32, N int32, arch int32) (r OpusT_opus_val16) {
	var Rxy, Ryy, best_den, rcp, tmp, yy, v55 OpusT_opus_val16
	var _saved_stack, signx, st, y, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var best_id, i, j, pulsesLeft, v53 int32
	var best_num, sum, xy OpusT_opus_val32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = Rxy, Ryy, _saved_stack, best_den, best_id, best_num, i, j, pulsesLeft, rcp, signx, st, sum, tmp, xy, y, yy, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v53, v55, v7, v9
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
	_ = arch
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _6
	_6:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _8
_8:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _10
	_10:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v7 = st
	goto _12
_12:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		goto _14
	_14:
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v11 = st
	goto _16
_16:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		goto _18
	_18:
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v15 = st
	goto _20
_20:
	if !(int64(int32(uint64(uint32(N))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4855, int32(217))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		goto _22
	_22:
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v19 = st
	goto _24
_24:
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32(N)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		goto _26
	_26:
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v23 = st
	goto _28
_28:
	y = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32(N))*(uint64(4)/uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _30
	_30:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _32
_32:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _34
	_34:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v7 = st
	goto _36
_36:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		goto _38
	_38:
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v11 = st
	goto _40
_40:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		goto _42
	_42:
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v15 = st
	goto _44
_44:
	if !(int64(int32(uint64(uint32(N))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4855, int32(218))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		goto _46
	_46:
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v19 = st
	goto _48
_48:
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32(N)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		goto _50
	_50:
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v23 = st
	goto _52
_52:
	signx = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32(N))*(uint64(4)/uint64(1)))
	/* Get rid of the sign */
	sum = float32(0)
	j = 0
	for {
		*(*int32)(unsafe.Pointer(signx + uintptr(j)*4)) = libc.BoolInt32(*(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(j)*4)) < float32(0))
		/* OPT: Make sure the compiler doesn't use a branch on ABS16(). */
		*(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(j)*4)) = float32(libc.Xfabs(tls, float64(*(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(j)*4)))))
		*(*int32)(unsafe.Pointer(iy + uintptr(j)*4)) = 0
		*(*OpusT_celt_norm)(unsafe.Pointer(y + uintptr(j)*4)) = float32(0)
		goto _54
	_54:
		;
		j = j + 1
		v53 = j
		if !(v53 < N) {
			break
		}
	}
	v55 = float32(0)
	yy = v55
	xy = v55
	pulsesLeft = K
	/* Do a pre-search by projecting on the pyramid */
	if K > N>>int32(1) {
		j = 0
		for {
			sum = sum + *(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(j)*4))
			goto _57
		_57:
			;
			j = j + 1
			v53 = j
			if !(v53 < N) {
				break
			}
		}
		/* If X is too small, just replace it with a pulse at 0 */
		/* Prevents infinities and NaNs from causing too many pulses
		   to be allocated. 64 is an approximation of infinity here. */
		if !(sum > float32(1e-15) && sum < float32(64)) {
			*(*OpusT_celt_norm)(unsafe.Pointer(X)) = float32(1)
			j = int32(1)
			for {
				*(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(j)*4)) = float32(0)
				goto _59
			_59:
				;
				j = j + 1
				v53 = j
				if !(v53 < N) {
					break
				}
			}
			sum = float32(1)
		}
		/* Using K+e with e < 1 guarantees we cannot get more than K pulses. */
		rcp = float32((float32(K) + float32(0.8)) * (float32(1) / sum))
		j = 0
		for {
			*(*int32)(unsafe.Pointer(iy + uintptr(j)*4)) = int32(libc.Xfloor(tls, float64(rcp**(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(j)*4)))))
			*(*OpusT_celt_norm)(unsafe.Pointer(y + uintptr(j)*4)) = float32(*(*int32)(unsafe.Pointer(iy + uintptr(j)*4)))
			yy = yy + OpusT_opus_val16(*(*OpusT_celt_norm)(unsafe.Pointer(y + uintptr(j)*4))**(*OpusT_celt_norm)(unsafe.Pointer(y + uintptr(j)*4)))
			xy = xy + OpusT_opus_val32(*(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(j)*4))**(*OpusT_celt_norm)(unsafe.Pointer(y + uintptr(j)*4)))
			*(*OpusT_celt_norm)(unsafe.Pointer(y + uintptr(j)*4)) *= float32(2)
			pulsesLeft = pulsesLeft - *(*int32)(unsafe.Pointer(iy + uintptr(j)*4))
			goto _61
		_61:
			;
			j = j + 1
			v53 = j
			if !(v53 < N) {
				break
			}
		}
	}
	_ = pulsesLeft >= int32(0)
	/* This should never happen, but just in case it does (e.g. on silence)
	   we fill the first bin with pulses. */
	if pulsesLeft > N+int32(3) {
		tmp = float32(pulsesLeft)
		yy = yy + OpusT_opus_val16(tmp*tmp)
		yy = yy + OpusT_opus_val16(tmp**(*OpusT_celt_norm)(unsafe.Pointer(y)))
		*(*int32)(unsafe.Pointer(iy)) += pulsesLeft
		pulsesLeft = 0
	}
	i = 0
	for {
		if !(i < pulsesLeft) {
			break
		}
		best_id = 0
		/* The squared magnitude term gets added anyway, so we might as well
		   add it outside the loop */
		yy = yy + float32(int32(1))
		/* Calculations for position 0 are out of the loop, in part to reduce
		   mispredicted branches (since the if condition is usually false)
		   in the loop. */
		/* Temporary sums of the new pulse(s) */
		Rxy = xy + *(*OpusT_celt_norm)(unsafe.Pointer(X))
		/* We're multiplying y[j] by two so we don't have to do it here */
		Ryy = yy + *(*OpusT_celt_norm)(unsafe.Pointer(y))
		/* Approximate score: we maximise Rxy/sqrt(Ryy) (we're guaranteed that
		   Rxy is positive because the sign is pre-computed) */
		Rxy = OpusT_opus_val16(Rxy * Rxy)
		best_den = Ryy
		best_num = Rxy
		j = int32(1)
		for {
			/* Temporary sums of the new pulse(s) */
			Rxy = xy + *(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(j)*4))
			/* We're multiplying y[j] by two so we don't have to do it here */
			Ryy = yy + *(*OpusT_celt_norm)(unsafe.Pointer(y + uintptr(j)*4))
			/* Approximate score: we maximise Rxy/sqrt(Ryy) (we're guaranteed that
			   Rxy is positive because the sign is pre-computed) */
			Rxy = OpusT_opus_val16(Rxy * Rxy)
			/* The idea is to check for num/den >= best_num/best_den, but that way
			   we can do it without any division */
			/* OPT: It's not clear whether a cmov is faster than a branch here
			   since the condition is more often false than true and using
			   a cmov introduces data dependencies across iterations. The optimal
			   choice may be architecture-dependent. */
			if libc.BoolInt64(!!(OpusT_opus_val32(best_den*Rxy) > OpusT_opus_val32(Ryy*best_num))) != 0 {
				best_den = Ryy
				best_num = Rxy
				best_id = j
			}
			goto _64
		_64:
			;
			j = j + 1
			v53 = j
			if !(v53 < N) {
				break
			}
		}
		/* Updating the sums of the new pulse(s) */
		xy = xy + *(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(best_id)*4))
		/* We're multiplying y[j] by two so we don't have to do it here */
		yy = yy + *(*OpusT_celt_norm)(unsafe.Pointer(y + uintptr(best_id)*4))
		/* Only now that we've made the final choice, update y/iy */
		/* Multiplying y[j] by 2 so we don't have to do it everywhere else */
		*(*OpusT_celt_norm)(unsafe.Pointer(y + uintptr(best_id)*4)) += float32(2)
		*(*int32)(unsafe.Pointer(iy + uintptr(best_id)*4)) = *(*int32)(unsafe.Pointer(iy + uintptr(best_id)*4)) + 1
		goto _62
	_62:
		;
		i = i + 1
	}
	/* Put the original sign back */
	j = 0
	for {
		/*iy[j] = signx[j] ? -iy[j] : iy[j];*/
		/* OPT: The is more likely to be compiled without a branch than the code above
		   but has the same performance otherwise. */
		*(*int32)(unsafe.Pointer(iy + uintptr(j)*4)) = *(*int32)(unsafe.Pointer(iy + uintptr(j)*4)) ^ -*(*int32)(unsafe.Pointer(signx + uintptr(j)*4)) + *(*int32)(unsafe.Pointer(signx + uintptr(j)*4))
		goto _66
	_66:
		;
		j = j + 1
		v53 = j
		if !(v53 < N) {
			break
		}
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _68
	_68:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _70
_70:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
	return yy
}

func Opus_alg_quant(tls *libc.TLS, X uintptr, N int32, K int32, spread int32, B int32, enc uintptr, gain OpusT_opus_val32, resynth int32, arch int32) (r uint32) {
	var _saved_stack, iy, st, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var collapse_mask uint32
	var yy OpusT_opus_val32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = _saved_stack, collapse_mask, iy, st, yy, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9
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
	if !(K > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+4868, __ccgo_ts+4855, int32(562))
	}
	if !(N > int32(1)) {
		Opus_celt_fatal(tls, __ccgo_ts+4927, __ccgo_ts+4855, int32(563))
	}
	/* Covers vectorization by up to 4. */
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _6
	_6:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _8
_8:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _10
	_10:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v7 = st
	goto _12
_12:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		goto _14
	_14:
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v11 = st
	goto _16
_16:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		goto _18
	_18:
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v15 = st
	goto _20
_20:
	if !(int64(int32(uint64(uint32(N+int32(3)))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4855, int32(566))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		goto _22
	_22:
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v19 = st
	goto _24
_24:
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32(N+int32(3))) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		goto _26
	_26:
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v23 = st
	goto _28
_28:
	iy = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32(N+int32(3)))*(uint64(4)/uint64(1)))
	Opus_exp_rotation(tls, X, N, int32(1), B, K, spread)
	yy = Opus_op_pvq_search_c(tls, X, iy, K, N, arch)
	collapse_mask = extract_collapse_mask(tls, iy, N, B)
	Opus_encode_pulses(tls, iy, N, K, enc)
	if resynth != 0 {
		normalise_residual(tls, iy, X, N, yy, gain, 0)
	}
	if resynth != 0 {
		Opus_exp_rotation(tls, X, N, -int32(1), B, K, spread)
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _30
	_30:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _32
_32:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
	return collapse_mask
}

// C documentation
//
//	/** Decode pulse vector and combine the result with the pitch vector to produce
//	    the final normalised signal in the current band. */
func Opus_alg_unquant(tls *libc.TLS, X uintptr, N int32, K int32, spread int32, B int32, dec uintptr, gain OpusT_opus_val32) (r uint32) {
	var Ryy OpusT_opus_val32
	var _saved_stack, iy, st, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var collapse_mask uint32
	var yy_shift int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = Ryy, _saved_stack, collapse_mask, iy, st, yy_shift, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9
	yy_shift = 0
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
	if !(K > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+4991, __ccgo_ts+4855, int32(629))
	}
	if !(N > int32(1)) {
		Opus_celt_fatal(tls, __ccgo_ts+5052, __ccgo_ts+4855, int32(630))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _6
	_6:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _8
_8:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _10
	_10:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v7 = st
	goto _12
_12:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		goto _14
	_14:
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v11 = st
	goto _16
_16:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		goto _18
	_18:
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v15 = st
	goto _20
_20:
	if !(int64(int32(uint64(uint32(N))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4855, int32(631))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		goto _22
	_22:
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v19 = st
	goto _24
_24:
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32(N)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		goto _26
	_26:
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v23 = st
	goto _28
_28:
	iy = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32(N))*(uint64(4)/uint64(1)))
	Ryy = Opus_decode_pulses(tls, iy, N, K, dec)
	normalise_residual(tls, iy, X, N, Ryy, gain, yy_shift)
	Opus_exp_rotation(tls, X, N, -int32(1), B, K, spread)
	collapse_mask = extract_collapse_mask(tls, iy, N, B)
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _30
	_30:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _32
_32:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
	return collapse_mask
}

func Opus_renormalise_vector(tls *libc.TLS, X uintptr, N1 int32, gain OpusT_opus_val32, arch int32) {
	var E, t, xy, v2 OpusT_opus_val32
	var g OpusT_opus_val16
	var i, i1 int32
	var xptr uintptr
	_, _, _, _, _, _, _, _ = E, g, i, i1, t, xptr, xy, v2
	_ = arch
	xy = float32(0)
	i = int32(0)
	for {
		if !(i < N1) {
			break
		}
		xy = xy + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(X + uintptr(i)*4))**(*OpusT_opus_val16)(unsafe.Pointer(X + uintptr(i)*4)))
		goto _1
	_1:
		;
		i = i + 1
	}
	v2 = xy
	goto _3
_3:
	E = float32(1e-15) + v2
	t = E
	g = float32(float32(1) / float32(libc.Xsqrt(tls, float64(t))) * gain)
	xptr = X
	i1 = 0
	for {
		if !(i1 < N1) {
			break
		}
		*(*OpusT_celt_norm)(unsafe.Pointer(xptr)) = OpusT_opus_val32(g * *(*OpusT_celt_norm)(unsafe.Pointer(xptr)))
		xptr += 4
		goto _4
	_4:
		;
		i1 = i1 + 1
	}
	/*return celt_sqrt(E);*/
}

func Opus_stereo_itheta(tls *libc.TLS, X uintptr, Y uintptr, stereo int32, N1 int32, arch int32) (r OpusT_opus_int32) {
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
			m = *(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(i1)*4)) + *(*OpusT_celt_norm)(unsafe.Pointer(Y + uintptr(i1)*4))
			s = *(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(i1)*4)) - *(*OpusT_celt_norm)(unsafe.Pointer(Y + uintptr(i1)*4))
			Emid = Emid + OpusT_opus_val32(m*m)
			Eside = Eside + OpusT_opus_val32(s*s)
			goto _2
		_2:
			;
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
			xy = xy + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(X + uintptr(i)*4))**(*OpusT_opus_val16)(unsafe.Pointer(X + uintptr(i)*4)))
			goto _3
		_3:
			;
			i = i + 1
		}
		v1 = xy
		goto _5
	_5:
		Emid = Emid + v1
		_ = arch
		xy = float32(0)
		i = int32(0)
		for {
			if !(i < N1) {
				break
			}
			xy = xy + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(Y + uintptr(i)*4))**(*OpusT_opus_val16)(unsafe.Pointer(Y + uintptr(i)*4)))
			goto _6
		_6:
			;
			i = i + 1
		}
		v1 = xy
		goto _8
	_8:
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
		goto _15
	_15:
		v11 = v14
		goto _12
	} else {
		v16 = v10 / v9
		x_sq = float32(v16 * v16)
		v17 = float32(float32(0.636619772367581) * (v16 + float32(float32(v16*x_sq)*(-float32(0.3333165943622589)+float32(x_sq*(float32(0.19962704181671143)+float32(x_sq*(-float32(0.13976582884788513)+float32(x_sq*(float32(0.09794234484434128)+float32(x_sq*(-float32(0.057773590087890625)+float32(x_sq*(float32(0.023040136322379112)+float32(x_sq*-float32(0.0043554059229791164))))))))))))))))
		goto _18
	_18:
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

func interp_bits2pulses(tls *libc.TLS, m uintptr, start int32, end int32, skip_start int32, bits1 uintptr, bits2 uintptr, thresh uintptr, cap1 uintptr, total OpusT_opus_int32, _balance uintptr, skip_rsv int32, intensity uintptr, intensity_rsv int32, dual_stereo uintptr, dual_stereo_rsv int32, bits uintptr, ebits uintptr, fine_priority uintptr, C int32, LM int32, ec uintptr, encode int32, prev int32, signalBandwidth int32) (r int32) {
	var N, N0, NClogN, alloc_floor, band_bits, band_width, codedBands, den, depth_threshold, done, extra_bits, extra_fine, hi, i, j, lo, logM, mid, offset, rem, stereo, tmp, tmp1, tmp2, v7, v8 int32
	var _saved_stack, st, v1, v3 uintptr
	var balance, bit, excess, left, percoeff, psum OpusT_opus_int32
	var v13, v14 OpusT_opus_uint32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = N, N0, NClogN, _saved_stack, alloc_floor, balance, band_bits, band_width, bit, codedBands, den, depth_threshold, done, excess, extra_bits, extra_fine, hi, i, j, left, lo, logM, mid, offset, percoeff, psum, rem, st, stereo, tmp, tmp1, tmp2, v1, v13, v14, v3, v7, v8
	codedBands = -int32(1)
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
			tmp = *(*int32)(unsafe.Pointer(bits1 + uintptr(j)*4)) + mid**(*int32)(unsafe.Pointer(bits2 + uintptr(j)*4))>>int32(ALLOC_STEPS)
			if tmp >= *(*int32)(unsafe.Pointer(thresh + uintptr(j)*4)) || done != 0 {
				done = int32(1)
				/* Don't allocate more than we can actually use */
				if tmp < *(*int32)(unsafe.Pointer(cap1 + uintptr(j)*4)) {
					v7 = tmp
				} else {
					v7 = *(*int32)(unsafe.Pointer(cap1 + uintptr(j)*4))
				}
				psum = psum + v7
			} else {
				if tmp >= alloc_floor {
					psum = psum + alloc_floor
				}
			}
			goto _6
		_6:
		}
		if psum > total {
			hi = mid
		} else {
			lo = mid
		}
		goto _5
	_5:
		;
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
		tmp1 = *(*int32)(unsafe.Pointer(bits1 + uintptr(j)*4)) + lo**(*int32)(unsafe.Pointer(bits2 + uintptr(j)*4))>>int32(ALLOC_STEPS)
		if tmp1 < *(*int32)(unsafe.Pointer(thresh + uintptr(j)*4)) && !(done != 0) {
			if tmp1 >= alloc_floor {
				tmp1 = alloc_floor
			} else {
				tmp1 = 0
			}
		} else {
			done = int32(1)
		}
		/* Don't allocate more than we can actually use */
		if tmp1 < *(*int32)(unsafe.Pointer(cap1 + uintptr(j)*4)) {
			v7 = tmp1
		} else {
			v7 = *(*int32)(unsafe.Pointer(cap1 + uintptr(j)*4))
		}
		tmp1 = v7
		*(*int32)(unsafe.Pointer(bits + uintptr(j)*4)) = tmp1
		psum = psum + tmp1
		goto _9
	_9:
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
		v13 = uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(codedBands)*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(start)*2))))
		_ = v13 > uint32(0)
		v14 = uint32(left) / v13
		goto _15
	_15:
		percoeff = int32(v14)
		left = left - (int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(codedBands)*2)))-int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(start)*2))))*percoeff
		if left-(int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j)*2)))-int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(start)*2)))) > 0 {
			v7 = left - (int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j)*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(start)*2))))
		} else {
			v7 = 0
		}
		rem = v7
		band_width = int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(codedBands)*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j)*2)))
		band_bits = *(*int32)(unsafe.Pointer(bits + uintptr(j)*4)) + percoeff*band_width + rem
		/*Only code a skip decision if we're above the threshold for this band.
		  Otherwise it is force-skipped.
		  This ensures that we have enough bits to code the skip flag.*/
		if *(*int32)(unsafe.Pointer(thresh + uintptr(j)*4)) > alloc_floor+int32(1)<<int32(BITRES) {
			v7 = *(*int32)(unsafe.Pointer(thresh + uintptr(j)*4))
		} else {
			v7 = alloc_floor + int32(1)<<int32(BITRES)
		}
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
					Opus_ec_enc_bit_logp(tls, ec, int32(1), uint32(1))
					break
				}
				Opus_ec_enc_bit_logp(tls, ec, 0, uint32(1))
			} else {
				if Opus_ec_dec_bit_logp(tls, ec, uint32(1)) != 0 {
					break
				}
			}
			/*We used a bit to skip this band.*/
			psum = psum + int32(1)<<int32(BITRES)
			band_bits = band_bits - int32(1)<<int32(BITRES)
		}
		/*Reclaim the bits originally allocated to this band.*/
		psum = psum - (*(*int32)(unsafe.Pointer(bits + uintptr(j)*4)) + intensity_rsv)
		if intensity_rsv > 0 {
			intensity_rsv = int32(LOG2_FRAC_TABLE[j-start])
		}
		psum = psum + intensity_rsv
		if band_bits >= alloc_floor {
			/*If we have enough for a fine energy bit per channel, use it.*/
			psum = psum + alloc_floor
			*(*int32)(unsafe.Pointer(bits + uintptr(j)*4)) = alloc_floor
		} else {
			/*Otherwise this band gets nothing at all.*/
			*(*int32)(unsafe.Pointer(bits + uintptr(j)*4)) = 0
		}
		goto _12
	_12:
		;
		codedBands = codedBands - 1
	}
	if !(codedBands > start) {
		Opus_celt_fatal(tls, __ccgo_ts+5118, __ccgo_ts+5155, int32(394))
	}
	/* Code the intensity and dual stereo parameters. */
	if intensity_rsv > 0 {
		if encode != 0 {
			if *(*int32)(unsafe.Pointer(intensity)) < codedBands {
				v7 = *(*int32)(unsafe.Pointer(intensity))
			} else {
				v7 = codedBands
			}
			*(*int32)(unsafe.Pointer(intensity)) = v7
			Opus_ec_enc_uint(tls, ec, uint32(*(*int32)(unsafe.Pointer(intensity))-start), uint32(codedBands+int32(1)-start))
		} else {
			*(*int32)(unsafe.Pointer(intensity)) = int32(uint32(start) + Opus_ec_dec_uint(tls, ec, uint32(codedBands+int32(1)-start)))
		}
	} else {
		*(*int32)(unsafe.Pointer(intensity)) = 0
	}
	if *(*int32)(unsafe.Pointer(intensity)) <= start {
		total = total + dual_stereo_rsv
		dual_stereo_rsv = 0
	}
	if dual_stereo_rsv > 0 {
		if encode != 0 {
			Opus_ec_enc_bit_logp(tls, ec, *(*int32)(unsafe.Pointer(dual_stereo)), uint32(1))
		} else {
			*(*int32)(unsafe.Pointer(dual_stereo)) = Opus_ec_dec_bit_logp(tls, ec, uint32(1))
		}
	} else {
		*(*int32)(unsafe.Pointer(dual_stereo)) = 0
	}
	/* Allocate the remaining bits */
	left = total - psum
	v13 = uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(codedBands)*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(start)*2))))
	_ = v13 > uint32(0)
	v14 = uint32(left) / v13
	goto _22
_22:
	percoeff = int32(v14)
	left = left - (int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(codedBands)*2)))-int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(start)*2))))*percoeff
	j = start
	for {
		if !(j < codedBands) {
			break
		}
		*(*int32)(unsafe.Pointer(bits + uintptr(j)*4)) += percoeff * (int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j+int32(1))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j)*2))))
		goto _23
	_23:
		;
		j = j + 1
	}
	j = start
	for {
		if !(j < codedBands) {
			break
		}
		if left < int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j+int32(1))*2)))-int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j)*2))) {
			v7 = left
		} else {
			v7 = int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j+int32(1))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j)*2)))
		}
		tmp2 = v7
		*(*int32)(unsafe.Pointer(bits + uintptr(j)*4)) += tmp2
		left = left - tmp2
		goto _24
	_24:
		;
		j = j + 1
	}
	/*for (j=0;j<end;j++)printf("%d ", bits[j]);printf("\n");*/
	balance = 0
	j = start
	for {
		if !(j < codedBands) {
			break
		}
		if !(*(*int32)(unsafe.Pointer(bits + uintptr(j)*4)) >= int32(0)) {
			Opus_celt_fatal(tls, __ccgo_ts+5170, __ccgo_ts+5155, int32(445))
		}
		N0 = int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j+int32(1))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j)*2)))
		N = N0 << LM
		bit = *(*int32)(unsafe.Pointer(bits + uintptr(j)*4)) + balance
		if N > int32(1) {
			if bit-*(*int32)(unsafe.Pointer(cap1 + uintptr(j)*4)) > 0 {
				v7 = bit - *(*int32)(unsafe.Pointer(cap1 + uintptr(j)*4))
			} else {
				v7 = 0
			}
			excess = v7
			*(*int32)(unsafe.Pointer(bits + uintptr(j)*4)) = bit - excess
			/* Compensate for the extra DoF in stereo */
			if C == int32(2) && N > int32(2) && !(*(*int32)(unsafe.Pointer(dual_stereo)) != 0) && j < *(*int32)(unsafe.Pointer(intensity)) {
				v7 = int32(1)
			} else {
				v7 = 0
			}
			den = C*N + v7
			NClogN = den * (int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FlogN + uintptr(j)*2))) + logM)
			/* Offset for the number of fine bits by log2(N)/2 + FINE_OFFSET
			   compared to their "fair share" of total/N */
			offset = NClogN>>int32(1) - den*int32(FINE_OFFSET)
			/* N=2 is the only point that doesn't match the curve */
			if N == int32(2) {
				offset = offset + den<<int32(BITRES)>>int32(2)
			}
			/* Changing the offset for allocating the second and third
			   fine energy bit */
			if *(*int32)(unsafe.Pointer(bits + uintptr(j)*4))+offset < den*int32(2)<<int32(BITRES) {
				offset = offset + NClogN>>int32(2)
			} else {
				if *(*int32)(unsafe.Pointer(bits + uintptr(j)*4))+offset < den*int32(3)<<int32(BITRES) {
					offset = offset + NClogN>>int32(3)
				}
			}
			/* Divide with rounding */
			if 0 > *(*int32)(unsafe.Pointer(bits + uintptr(j)*4))+offset+den<<(int32(BITRES)-int32(1)) {
				v7 = 0
			} else {
				v7 = *(*int32)(unsafe.Pointer(bits + uintptr(j)*4)) + offset + den<<(int32(BITRES)-int32(1))
			}
			*(*int32)(unsafe.Pointer(ebits + uintptr(j)*4)) = v7
			v13 = uint32(den)
			_ = v13 > uint32(0)
			v14 = uint32(*(*int32)(unsafe.Pointer(ebits + uintptr(j)*4))) / v13
			goto _32
		_32:
			*(*int32)(unsafe.Pointer(ebits + uintptr(j)*4)) = int32(v14 >> int32(BITRES))
			/* Make sure not to bust */
			if C**(*int32)(unsafe.Pointer(ebits + uintptr(j)*4)) > *(*int32)(unsafe.Pointer(bits + uintptr(j)*4))>>int32(BITRES) {
				*(*int32)(unsafe.Pointer(ebits + uintptr(j)*4)) = *(*int32)(unsafe.Pointer(bits + uintptr(j)*4)) >> stereo >> int32(BITRES)
			}
			/* More than that is useless because that's about as far as PVQ can go */
			if *(*int32)(unsafe.Pointer(ebits + uintptr(j)*4)) < int32(MAX_FINE_BITS) {
				v7 = *(*int32)(unsafe.Pointer(ebits + uintptr(j)*4))
			} else {
				v7 = int32(MAX_FINE_BITS)
			}
			*(*int32)(unsafe.Pointer(ebits + uintptr(j)*4)) = v7
			/* If we rounded down or capped this band, make it a candidate for the
			   final fine energy pass */
			*(*int32)(unsafe.Pointer(fine_priority + uintptr(j)*4)) = libc.BoolInt32(*(*int32)(unsafe.Pointer(ebits + uintptr(j)*4))*(den<<int32(BITRES)) >= *(*int32)(unsafe.Pointer(bits + uintptr(j)*4))+offset)
			/* Remove the allocated fine bits; the rest are assigned to PVQ */
			*(*int32)(unsafe.Pointer(bits + uintptr(j)*4)) -= C * *(*int32)(unsafe.Pointer(ebits + uintptr(j)*4)) << int32(BITRES)
		} else {
			/* For N=1, all bits go to fine energy except for a single sign bit */
			if 0 > bit-C<<int32(BITRES) {
				v7 = 0
			} else {
				v7 = bit - C<<int32(BITRES)
			}
			excess = v7
			*(*int32)(unsafe.Pointer(bits + uintptr(j)*4)) = bit - excess
			*(*int32)(unsafe.Pointer(ebits + uintptr(j)*4)) = 0
			*(*int32)(unsafe.Pointer(fine_priority + uintptr(j)*4)) = int32(1)
		}
		/* Fine energy can't take advantage of the re-balancing in
		   quant_all_bands().
		  Instead, do the re-balancing here.*/
		if excess > 0 {
			if excess>>(stereo+int32(BITRES)) < int32(MAX_FINE_BITS)-*(*int32)(unsafe.Pointer(ebits + uintptr(j)*4)) {
				v7 = excess >> (stereo + int32(BITRES))
			} else {
				v7 = int32(MAX_FINE_BITS) - *(*int32)(unsafe.Pointer(ebits + uintptr(j)*4))
			}
			extra_fine = v7
			*(*int32)(unsafe.Pointer(ebits + uintptr(j)*4)) += extra_fine
			extra_bits = extra_fine * C << int32(BITRES)
			*(*int32)(unsafe.Pointer(fine_priority + uintptr(j)*4)) = libc.BoolInt32(extra_bits >= excess-balance)
			excess = excess - extra_bits
		}
		balance = excess
		if !(*(*int32)(unsafe.Pointer(bits + uintptr(j)*4)) >= int32(0)) {
			Opus_celt_fatal(tls, __ccgo_ts+5170, __ccgo_ts+5155, int32(516))
		}
		if !(*(*int32)(unsafe.Pointer(ebits + uintptr(j)*4)) >= int32(0)) {
			Opus_celt_fatal(tls, __ccgo_ts+5201, __ccgo_ts+5155, int32(517))
		}
		goto _26
	_26:
		;
		j = j + 1
	}
	/* Save any remaining bits over the cap for the rebalancing in
	   quant_all_bands(). */
	*(*OpusT_opus_int32)(unsafe.Pointer(_balance)) = balance
	/* The skipped bands use all their bits for fine energy. */
	for {
		if !(j < end) {
			break
		}
		*(*int32)(unsafe.Pointer(ebits + uintptr(j)*4)) = *(*int32)(unsafe.Pointer(bits + uintptr(j)*4)) >> stereo >> int32(BITRES)
		if !(C**(*int32)(unsafe.Pointer(ebits + uintptr(j)*4))<<int32(BITRES) == *(*int32)(unsafe.Pointer(bits + uintptr(j)*4))) {
			Opus_celt_fatal(tls, __ccgo_ts+5233, __ccgo_ts+5155, int32(527))
		}
		*(*int32)(unsafe.Pointer(bits + uintptr(j)*4)) = 0
		*(*int32)(unsafe.Pointer(fine_priority + uintptr(j)*4)) = libc.BoolInt32(*(*int32)(unsafe.Pointer(ebits + uintptr(j)*4)) < int32(1))
		goto _36
	_36:
		;
		j = j + 1
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _38
	_38:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _40
_40:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
	return codedBands
}

func Opus_clt_compute_allocation(tls *libc.TLS, m uintptr, start int32, end int32, offsets uintptr, cap1 uintptr, alloc_trim int32, intensity uintptr, dual_stereo uintptr, total OpusT_opus_int32, balance uintptr, pulses uintptr, ebits uintptr, fine_priority uintptr, C int32, LM int32, ec uintptr, encode int32, prev int32, signalBandwidth int32) (r int32) {
	var N, N1, bits1j, bits2j, bitsj, codedBands, done, dual_stereo_rsv, hi, intensity_rsv, j, len1, lo, mid, psum, skip_rsv, skip_start, v5 int32
	var _saved_stack, bits1, bits2, st, thresh, trim_offset, v1, v11, v13, v15, v17, v19, v21, v23, v25, v27, v3, v9 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = N, N1, _saved_stack, bits1, bits1j, bits2, bits2j, bitsj, codedBands, done, dual_stereo_rsv, hi, intensity_rsv, j, len1, lo, mid, psum, skip_rsv, skip_start, st, thresh, trim_offset, v1, v11, v13, v15, v17, v19, v21, v23, v25, v27, v3, v5, v9
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
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _10
	_10:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _12
_12:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		goto _14
	_14:
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v11 = st
	goto _16
_16:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		goto _18
	_18:
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v15 = st
	goto _20
_20:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		goto _22
	_22:
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v19 = st
	goto _24
_24:
	if !(int64(int32(uint64(uint32(len1))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v19)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5155, int32(570))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		goto _26
	_26:
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v23 = st
	goto _28
_28:
	*(*uintptr)(unsafe.Pointer(v23 + 8)) += uintptr(uint64(uint32(len1)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v25 = libc.Xmalloc(tls, uint64(16))
		goto _30
	_30:
		st = v25
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v27 = st
	goto _32
_32:
	bits1 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v27)).Fglobal_stack - uintptr(uint64(uint32(len1))*(uint64(4)/uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _34
	_34:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _36
_36:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		goto _38
	_38:
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v11 = st
	goto _40
_40:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		goto _42
	_42:
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v15 = st
	goto _44
_44:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		goto _46
	_46:
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v19 = st
	goto _48
_48:
	if !(int64(int32(uint64(uint32(len1))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v19)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5155, int32(571))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		goto _50
	_50:
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v23 = st
	goto _52
_52:
	*(*uintptr)(unsafe.Pointer(v23 + 8)) += uintptr(uint64(uint32(len1)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v25 = libc.Xmalloc(tls, uint64(16))
		goto _54
	_54:
		st = v25
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v27 = st
	goto _56
_56:
	bits2 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v27)).Fglobal_stack - uintptr(uint64(uint32(len1))*(uint64(4)/uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _58
	_58:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _60
_60:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		goto _62
	_62:
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v11 = st
	goto _64
_64:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		goto _66
	_66:
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v15 = st
	goto _68
_68:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		goto _70
	_70:
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v19 = st
	goto _72
_72:
	if !(int64(int32(uint64(uint32(len1))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v19)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5155, int32(572))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		goto _74
	_74:
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v23 = st
	goto _76
_76:
	*(*uintptr)(unsafe.Pointer(v23 + 8)) += uintptr(uint64(uint32(len1)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v25 = libc.Xmalloc(tls, uint64(16))
		goto _78
	_78:
		st = v25
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v27 = st
	goto _80
_80:
	thresh = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v27)).Fglobal_stack - uintptr(uint64(uint32(len1))*(uint64(4)/uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _82
	_82:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _84
_84:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		goto _86
	_86:
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v11 = st
	goto _88
_88:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		goto _90
	_90:
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v15 = st
	goto _92
_92:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		goto _94
	_94:
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v19 = st
	goto _96
_96:
	if !(int64(int32(uint64(uint32(len1))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v19)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5155, int32(573))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		goto _98
	_98:
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v23 = st
	goto _100
_100:
	*(*uintptr)(unsafe.Pointer(v23 + 8)) += uintptr(uint64(uint32(len1)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v25 = libc.Xmalloc(tls, uint64(16))
		goto _102
	_102:
		st = v25
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v27 = st
	goto _104
_104:
	trim_offset = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v27)).Fglobal_stack - uintptr(uint64(uint32(len1))*(uint64(4)/uint64(1)))
	j = start
	for {
		if !(j < end) {
			break
		}
		/* Below this threshold, we're sure not to allocate any PVQ bits */
		if C<<int32(BITRES) > int32(3)*(int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j+int32(1))*2)))-int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j)*2))))<<LM<<int32(BITRES)>>int32(4) {
			v5 = C << int32(BITRES)
		} else {
			v5 = int32(3) * (int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j+int32(1))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j)*2)))) << LM << int32(BITRES) >> int32(4)
		}
		*(*int32)(unsafe.Pointer(thresh + uintptr(j)*4)) = v5
		/* Tilt of the allocation curve */
		*(*int32)(unsafe.Pointer(trim_offset + uintptr(j)*4)) = C * (int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j+int32(1))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j)*2)))) * (alloc_trim - int32(5) - LM) * (end - j - int32(1)) * (int32(1) << (LM + int32(BITRES))) >> int32(6)
		/* Giving less resolution to single-coefficient bands because they get
		   more benefit from having one coarse value per coefficient*/
		if (int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j+int32(1))*2)))-int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j)*2))))<<LM == int32(1) {
			*(*int32)(unsafe.Pointer(trim_offset + uintptr(j)*4)) -= C << int32(BITRES)
		}
		goto _105
	_105:
		;
		j = j + 1
	}
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
			N = int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j+int32(1))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j)*2)))
			bitsj = C * N * int32(*(*uint8)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FallocVectors + uintptr(mid*len1+j)))) << LM >> int32(2)
			if bitsj > 0 {
				if 0 > bitsj+*(*int32)(unsafe.Pointer(trim_offset + uintptr(j)*4)) {
					v5 = 0
				} else {
					v5 = bitsj + *(*int32)(unsafe.Pointer(trim_offset + uintptr(j)*4))
				}
				bitsj = v5
			}
			bitsj = bitsj + *(*int32)(unsafe.Pointer(offsets + uintptr(j)*4))
			if bitsj >= *(*int32)(unsafe.Pointer(thresh + uintptr(j)*4)) || done != 0 {
				done = int32(1)
				/* Don't allocate more than we can actually use */
				if bitsj < *(*int32)(unsafe.Pointer(cap1 + uintptr(j)*4)) {
					v5 = bitsj
				} else {
					v5 = *(*int32)(unsafe.Pointer(cap1 + uintptr(j)*4))
				}
				psum = psum + v5
			} else {
				if bitsj >= C<<int32(BITRES) {
					psum = psum + C<<int32(BITRES)
				}
			}
			goto _107
		_107:
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
		N1 = int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j+int32(1))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(j)*2)))
		bits1j = C * N1 * int32(*(*uint8)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FallocVectors + uintptr(lo*len1+j)))) << LM >> int32(2)
		if hi >= (*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbAllocVectors {
			v5 = *(*int32)(unsafe.Pointer(cap1 + uintptr(j)*4))
		} else {
			v5 = C * N1 * int32(*(*uint8)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FallocVectors + uintptr(hi*len1+j)))) << LM >> int32(2)
		}
		bits2j = v5
		if bits1j > 0 {
			if 0 > bits1j+*(*int32)(unsafe.Pointer(trim_offset + uintptr(j)*4)) {
				v5 = 0
			} else {
				v5 = bits1j + *(*int32)(unsafe.Pointer(trim_offset + uintptr(j)*4))
			}
			bits1j = v5
		}
		if bits2j > 0 {
			if 0 > bits2j+*(*int32)(unsafe.Pointer(trim_offset + uintptr(j)*4)) {
				v5 = 0
			} else {
				v5 = bits2j + *(*int32)(unsafe.Pointer(trim_offset + uintptr(j)*4))
			}
			bits2j = v5
		}
		if lo > 0 {
			bits1j = bits1j + *(*int32)(unsafe.Pointer(offsets + uintptr(j)*4))
		}
		bits2j = bits2j + *(*int32)(unsafe.Pointer(offsets + uintptr(j)*4))
		if *(*int32)(unsafe.Pointer(offsets + uintptr(j)*4)) > 0 {
			skip_start = j
		}
		if 0 > bits2j-bits1j {
			v5 = 0
		} else {
			v5 = bits2j - bits1j
		}
		bits2j = v5
		*(*int32)(unsafe.Pointer(bits1 + uintptr(j)*4)) = bits1j
		*(*int32)(unsafe.Pointer(bits2 + uintptr(j)*4)) = bits2j
		goto _112
	_112:
		;
		j = j + 1
	}
	codedBands = interp_bits2pulses(tls, m, start, end, skip_start, bits1, bits2, thresh, cap1, total, balance, skip_rsv, intensity, intensity_rsv, dual_stereo, dual_stereo_rsv, pulses, ebits, fine_priority, C, LM, ec, encode, prev, signalBandwidth)
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _118
	_118:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _120
_120:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
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
func Opus_ec_tell_frac(tls *libc.TLS, _this uintptr) (r1 OpusT_opus_uint32) {
	var b uint32
	var l int32
	var nbits, r OpusT_opus_uint32
	_, _, _, _ = b, l, nbits, r
	nbits = uint32((*OpusT_ec_ctx)(unsafe.Pointer(_this)).Fnbits_total << int32(BITRES))
	l = int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(_this)).Frng)
	r = (*OpusT_ec_ctx)(unsafe.Pointer(_this)).Frng >> (l - int32(16))
	b = r>>int32(12) - uint32(8)
	b = b + libc.BoolUint32(r > correction[b])
	l = int32(uint32(l<<int32(3)) + b)
	return nbits - uint32(l)
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

func Opus_hysteresis_decision(tls *libc.TLS, val OpusT_opus_val16, thresholds uintptr, hysteresis uintptr, N int32, prev int32) (r int32) {
	var i int32
	_ = i
	i = 0
	for {
		if !(i < N) {
			break
		}
		if val < *(*OpusT_opus_val16)(unsafe.Pointer(thresholds + uintptr(i)*4)) {
			break
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	if i > prev && val < *(*OpusT_opus_val16)(unsafe.Pointer(thresholds + uintptr(prev)*4))+*(*OpusT_opus_val16)(unsafe.Pointer(hysteresis + uintptr(prev)*4)) {
		i = prev
	}
	if i < prev && val > *(*OpusT_opus_val16)(unsafe.Pointer(thresholds + uintptr(prev-int32(1))*4))-*(*OpusT_opus_val16)(unsafe.Pointer(hysteresis + uintptr(prev-int32(1))*4)) {
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
func Opus_compute_band_energies(tls *libc.TLS, m uintptr, X uintptr, bandE uintptr, end int32, C int32, LM int32, arch int32) {
	var N1, c, i, i1, v1 int32
	var eBands uintptr
	var sum, xy, v5 OpusT_opus_val32
	_, _, _, _, _, _, _, _, _ = N1, c, eBands, i, i1, sum, xy, v1, v5
	eBands = (*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands
	N1 = (*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FshortMdctSize << LM
	c = 0
	for {
		i1 = 0
		for {
			if !(i1 < end) {
				break
			}
			_ = arch
			xy = float32(0)
			i = int32(0)
			for {
				if !(i < (int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1+int32(1))*2)))-int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1)*2))))<<LM) {
					break
				}
				xy = xy + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(X + uintptr(c*N1+int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1)*2)))<<LM)*4 + uintptr(i)*4))**(*OpusT_opus_val16)(unsafe.Pointer(X + uintptr(c*N1+int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1)*2)))<<LM)*4 + uintptr(i)*4)))
				goto _4
			_4:
				;
				i = i + 1
			}
			v5 = xy
			goto _6
		_6:
			sum = float32(1e-27) + v5
			*(*OpusT_celt_ener)(unsafe.Pointer(bandE + uintptr(i1+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4)) = float32(libc.Xsqrt(tls, float64(sum)))
			/*printf ("%f ", bandE[i+c*m->nbEBands]);*/
			goto _3
		_3:
			;
			i1 = i1 + 1
		}
		goto _2
	_2:
		;
		c = c + 1
		v1 = c
		if !(v1 < C) {
			break
		}
	}
	/*printf ("\n");*/
}

// C documentation
//
//	/* Normalise each band such that the energy is one. */
func Opus_normalise_bands(tls *libc.TLS, m uintptr, freq uintptr, X uintptr, bandE uintptr, end int32, C int32, M int32) {
	var N, c, i, j, v1 int32
	var eBands uintptr
	var g OpusT_opus_val16
	_, _, _, _, _, _, _ = N, c, eBands, g, i, j, v1
	eBands = (*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands
	N = M * (*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FshortMdctSize
	c = 0
	for {
		i = 0
		for {
			if !(i < end) {
				break
			}
			g = float32(1) / (float32(1e-27) + *(*OpusT_celt_ener)(unsafe.Pointer(bandE + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4)))
			j = M * int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i)*2)))
			for {
				if !(j < M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i+int32(1))*2)))) {
					break
				}
				*(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(j+c*N)*4)) = OpusT_celt_norm(*(*OpusT_celt_sig)(unsafe.Pointer(freq + uintptr(j+c*N)*4)) * g)
				goto _4
			_4:
				;
				j = j + 1
			}
			goto _3
		_3:
			;
			i = i + 1
		}
		goto _2
	_2:
		;
		c = c + 1
		v1 = c
		if !(v1 < C) {
			break
		}
	}
}

// C documentation
//
//	/* De-normalise the energy to produce the synthesis from the unit-energy bands */
func Opus_denormalise_bands(tls *libc.TLS, m uintptr, X uintptr, freq uintptr, bandLogE uintptr, start int32, end int32, M int32, downsample int32, silence int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var N, band_end, bound, i, j, v1 int32
	var eBands, f, x1, v4 uintptr
	var frac, v6, v7, v8 float32
	var g OpusT_opus_val32
	var integer OpusT_opus_int32
	var lg OpusT_celt_glog
	var _ /* res at bp+0 */ struct {
		Fi [0]OpusT_opus_uint32
		Ff float32
	}
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = N, band_end, bound, eBands, f, frac, g, i, integer, j, lg, x1, v1, v4, v6, v7, v8
	eBands = (*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands
	N = M * (*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FshortMdctSize
	bound = M * int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(end)*2)))
	if downsample != int32(1) {
		if bound < N/downsample {
			v1 = bound
		} else {
			v1 = N / downsample
		}
		bound = v1
	}
	if silence != 0 {
		bound = 0
		v1 = int32(0)
		end = v1
		start = v1
	}
	f = freq
	x1 = X + uintptr(M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(start)*2))))*4
	if start != 0 {
		i = 0
		for {
			if !(i < M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(start)*2)))) {
				break
			}
			v4 = f
			f += 4
			*(*OpusT_celt_sig)(unsafe.Pointer(v4)) = float32(0)
			goto _3
		_3:
			;
			i = i + 1
		}
	} else {
		f = f + uintptr(M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(start)*2))))*4
	}
	i = start
	for {
		if !(i < end) {
			break
		}
		j = M * int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i)*2)))
		band_end = M * int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i+int32(1))*2)))
		lg = *(*OpusT_celt_glog)(unsafe.Pointer(bandLogE + uintptr(i)*4)) + Opus_eMeans[i]
		if float32(32) < lg {
			v6 = float32(32)
		} else {
			v6 = lg
		}
		v7 = v6
		integer = int32(libc.Xfloor(tls, float64(v7)))
		if integer < -int32(50) {
			v8 = float32(0)
			goto _9
		}
		frac = v7 - float32(integer)
		*(*float32)(unsafe.Pointer(bp)) = float32(0.9999999403953552) + float32(frac*(float32(0.6931530833244324)+float32(frac*(float32(0.24015361070632935)+float32(frac*(float32(0.05582631751894951)+float32(frac*(float32(0.00898933969438076)+float32(frac*float32(0.0018775766948238015))))))))))
		*(*OpusT_opus_uint32)(unsafe.Pointer(bp)) = uint32(int32(*(*OpusT_opus_uint32)(unsafe.Pointer(bp)))+int32(uint32(integer)<<int32(23))) & uint32(0x7fffffff)
		v8 = *(*float32)(unsafe.Pointer(bp))
		goto _9
	_9:
		g = v8
		for {
			v4 = f
			f += 4
			*(*OpusT_celt_sig)(unsafe.Pointer(v4)) = OpusT_celt_norm(*(*OpusT_celt_norm)(unsafe.Pointer(x1)) * g)
			x1 += 4
			goto _11
		_11:
			;
			j = j + 1
			v1 = j
			if !(v1 < band_end) {
				break
			}
		}
		goto _5
	_5:
		;
		i = i + 1
	}
	if !(start <= end) {
		Opus_celt_fatal(tls, __ccgo_ts+5281, __ccgo_ts+5312, int32(254))
	}
	libc.Xmemset(tls, freq+uintptr(bound)*4, 0, uint64(uint32(N-bound))*uint64(4))
}

// C documentation
//
//	/* This prevents energy collapse for transients with multiple short MDCTs */
func Opus_anti_collapse(tls *libc.TLS, m uintptr, X_ uintptr, collapse_masks uintptr, LM int32, C int32, size int32, start int32, end int32, logE uintptr, prev1logE uintptr, prev2logE uintptr, pulses uintptr, seed OpusT_opus_uint32, encode int32, arch int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var Ediff, v13 OpusT_opus_val32
	var N0, c, depth, i, j, k, renormalize, v8 int32
	var X uintptr
	var frac, v5, v6 float32
	var integer OpusT_opus_int32
	var prev1, prev2, v10 OpusT_celt_glog
	var r, v20 OpusT_celt_norm
	var sqrt_1, thresh, v17 OpusT_opus_val16
	var v2, v3 OpusT_opus_uint32
	var _ /* res at bp+0 */ struct {
		Fi [0]OpusT_opus_uint32
		Ff float32
	}
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = Ediff, N0, X, c, depth, frac, i, integer, j, k, prev1, prev2, r, renormalize, sqrt_1, thresh, v10, v13, v17, v2, v20, v3, v5, v6, v8
	i = start
	for {
		if !(i < end) {
			break
		}
		N0 = int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(i+int32(1))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(i)*2)))
		/* depth in 1/8 bits */
		_ = *(*int32)(unsafe.Pointer(pulses + uintptr(i)*4)) >= int32(0)
		v2 = uint32(int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(i+int32(1))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(i)*2))))
		_ = v2 > uint32(0)
		v3 = uint32(int32(1)+*(*int32)(unsafe.Pointer(pulses + uintptr(i)*4))) / v2
		goto _4
	_4:
		depth = int32(v3 >> LM)
		v5 = float32(-float32(0.125) * float32(depth))
		integer = int32(libc.Xfloor(tls, float64(v5)))
		if integer < -int32(50) {
			v6 = float32(0)
			goto _7
		}
		frac = v5 - float32(integer)
		*(*float32)(unsafe.Pointer(bp)) = float32(0.9999999403953552) + float32(frac*(float32(0.6931530833244324)+float32(frac*(float32(0.24015361070632935)+float32(frac*(float32(0.05582631751894951)+float32(frac*(float32(0.00898933969438076)+float32(frac*float32(0.0018775766948238015))))))))))
		*(*OpusT_opus_uint32)(unsafe.Pointer(bp)) = uint32(int32(*(*OpusT_opus_uint32)(unsafe.Pointer(bp)))+int32(uint32(integer)<<int32(23))) & uint32(0x7fffffff)
		v6 = *(*float32)(unsafe.Pointer(bp))
		goto _7
	_7:
		thresh = OpusT_opus_val16(float32(0.5) * v6)
		sqrt_1 = float32(1) / float32(libc.Xsqrt(tls, float64(N0<<LM)))
		c = 0
		for {
			renormalize = 0
			prev1 = *(*OpusT_celt_glog)(unsafe.Pointer(prev1logE + uintptr(c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands+i)*4))
			prev2 = *(*OpusT_celt_glog)(unsafe.Pointer(prev2logE + uintptr(c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands+i)*4))
			if !(encode != 0) && C == int32(1) {
				if prev1 > *(*OpusT_celt_glog)(unsafe.Pointer(prev1logE + uintptr((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands+i)*4)) {
					v10 = prev1
				} else {
					v10 = *(*OpusT_celt_glog)(unsafe.Pointer(prev1logE + uintptr((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands+i)*4))
				}
				prev1 = v10
				if prev2 > *(*OpusT_celt_glog)(unsafe.Pointer(prev2logE + uintptr((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands+i)*4)) {
					v10 = prev2
				} else {
					v10 = *(*OpusT_celt_glog)(unsafe.Pointer(prev2logE + uintptr((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands+i)*4))
				}
				prev2 = v10
			}
			if prev1 < prev2 {
				v10 = prev1
			} else {
				v10 = prev2
			}
			Ediff = *(*OpusT_celt_glog)(unsafe.Pointer(logE + uintptr(c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands+i)*4)) - v10
			if float32(int32(0)) > Ediff {
				v13 = float32(int32(0))
			} else {
				v13 = Ediff
			}
			Ediff = v13
			/* r needs to be multiplied by 2 or 2*sqrt(2) depending on LM because
			   short blocks don't have the same energy as long */
			v5 = -Ediff
			integer = int32(libc.Xfloor(tls, float64(v5)))
			if integer < -int32(50) {
				v6 = float32(0)
				goto _16
			}
			frac = v5 - float32(integer)
			*(*float32)(unsafe.Pointer(bp)) = float32(0.9999999403953552) + float32(frac*(float32(0.6931530833244324)+float32(frac*(float32(0.24015361070632935)+float32(frac*(float32(0.05582631751894951)+float32(frac*(float32(0.00898933969438076)+float32(frac*float32(0.0018775766948238015))))))))))
			*(*OpusT_opus_uint32)(unsafe.Pointer(bp)) = uint32(int32(*(*OpusT_opus_uint32)(unsafe.Pointer(bp)))+int32(uint32(integer)<<int32(23))) & uint32(0x7fffffff)
			v6 = *(*float32)(unsafe.Pointer(bp))
			goto _16
		_16:
			r = OpusT_celt_norm(float32(2) * v6)
			if LM == int32(3) {
				r = r * float32(1.41421356)
			}
			if thresh < r {
				v17 = thresh
			} else {
				v17 = r
			}
			r = v17
			r = OpusT_celt_norm(r * sqrt_1)
			X = X_ + uintptr(c*size)*4 + uintptr(int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(i)*2)))<<LM)*4
			k = 0
			for {
				if !(k < int32(1)<<LM) {
					break
				}
				/* Detect collapse */
				if !(int32(*(*uint8)(unsafe.Pointer(collapse_masks + uintptr(i*C+c))))&(int32(1)<<k) != 0) {
					/* Fill with noise */
					j = 0
					for {
						if !(j < N0) {
							break
						}
						seed = Opus_celt_lcg_rand(tls, seed)
						if seed&uint32(0x8000) != 0 {
							v20 = r
						} else {
							v20 = -r
						}
						*(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(j<<LM+k)*4)) = v20
						goto _19
					_19:
						;
						j = j + 1
					}
					renormalize = int32(1)
				}
				goto _18
			_18:
				;
				k = k + 1
			}
			/* We just added some energy, so we need to renormalise */
			if renormalize != 0 {
				Opus_renormalise_vector(tls, X, N0<<LM, float32(1), arch)
			}
			goto _9
		_9:
			;
			c = c + 1
			v8 = c
			if !(v8 < C) {
				break
			}
		}
		goto _1
	_1:
		;
		i = i + 1
	}
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
func compute_channel_weights(tls *libc.TLS, Ex OpusT_celt_ener, Ey OpusT_celt_ener, w uintptr) {
	var minE, v1 OpusT_celt_ener
	_, _ = minE, v1
	if Ex < Ey {
		v1 = Ex
	} else {
		v1 = Ey
	}
	minE = v1
	/* Adjustment to make the weights a bit more conservative. */
	Ex = Ex + minE/float32(3)
	Ey = Ey + minE/float32(3)
	*(*OpusT_opus_val16)(unsafe.Pointer(w)) = Ex
	*(*OpusT_opus_val16)(unsafe.Pointer(w + 1*4)) = Ey
}

func intensity_stereo(tls *libc.TLS, m uintptr, X uintptr, Y uintptr, bandE uintptr, bandID int32, N int32) {
	var a1, a2, left, norm, right OpusT_opus_val16
	var i, j int32
	_, _, _, _, _, _, _ = a1, a2, i, j, left, norm, right
	i = bandID
	left = *(*OpusT_celt_ener)(unsafe.Pointer(bandE + uintptr(i)*4))
	right = *(*OpusT_celt_ener)(unsafe.Pointer(bandE + uintptr(i+(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4))
	norm = float32(1e-15) + float32(libc.Xsqrt(tls, float64(float32(1e-15)+OpusT_opus_val32(left*left)+OpusT_opus_val32(right*right))))
	a1 = left / norm
	a2 = right / norm
	j = 0
	for {
		if !(j < N) {
			break
		}
		*(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(j)*4)) = OpusT_opus_val16(a1**(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(j)*4))) + OpusT_opus_val16(a2**(*OpusT_celt_norm)(unsafe.Pointer(Y + uintptr(j)*4)))
		/* Side is not encoded, no need to calculate */
		goto _1
	_1:
		;
		j = j + 1
	}
}

func stereo_split(tls *libc.TLS, X uintptr, Y uintptr, N int32) {
	var j int32
	var l, r OpusT_opus_val32
	_, _, _ = j, l, r
	j = 0
	for {
		if !(j < N) {
			break
		}
		l = float32(float32(0.70710678) * *(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(j)*4)))
		r = float32(float32(0.70710678) * *(*OpusT_celt_norm)(unsafe.Pointer(Y + uintptr(j)*4)))
		*(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(j)*4)) = l + r
		*(*OpusT_celt_norm)(unsafe.Pointer(Y + uintptr(j)*4)) = r - l
		goto _1
	_1:
		;
		j = j + 1
	}
}

func stereo_merge(tls *libc.TLS, X uintptr, Y uintptr, mid OpusT_opus_val32, N1 int32, arch int32) {
	var El, Er, lgain, rgain, side, t, xp, xy, v2 OpusT_opus_val32
	var i, j int32
	var l, r OpusT_celt_norm
	_, _, _, _, _, _, _, _, _, _, _, _, _ = El, Er, i, j, l, lgain, r, rgain, side, t, xp, xy, v2
	xp = float32(0)
	side = float32(0)
	/* Compute the norm of X+Y and X-Y as |X|^2 + |Y|^2 +/- sum(xy) */
	_ = arch
	xy = float32(0)
	i = int32(0)
	for {
		if !(i < N1) {
			break
		}
		xy = xy + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(Y + uintptr(i)*4))**(*OpusT_opus_val16)(unsafe.Pointer(X + uintptr(i)*4)))
		goto _1
	_1:
		;
		i = i + 1
	}
	v2 = xy
	goto _3
_3:
	xp = v2
	_ = arch
	xy = float32(0)
	i = int32(0)
	for {
		if !(i < N1) {
			break
		}
		xy = xy + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(Y + uintptr(i)*4))**(*OpusT_opus_val16)(unsafe.Pointer(Y + uintptr(i)*4)))
		goto _4
	_4:
		;
		i = i + 1
	}
	v2 = xy
	goto _6
_6:
	side = v2
	/* Compensating for the mid normalization */
	xp = OpusT_opus_val32(mid * xp)
	/* mid and side are in Q15, not Q14 like X and Y */
	El = OpusT_opus_val32(mid*mid) + side - OpusT_opus_val32(float32(2)*xp)
	Er = OpusT_opus_val32(mid*mid) + side + OpusT_opus_val32(float32(2)*xp)
	if Er < float32(0.0006) || El < float32(0.0006) {
		libc.Xmemcpy(tls, Y, X, uint64(uint32(N1))*uint64(4)+uint64(0*((int64(Y)-int64(X))/4)))
		return
	}
	t = El
	lgain = float32(1) / float32(libc.Xsqrt(tls, float64(t)))
	t = Er
	rgain = float32(1) / float32(libc.Xsqrt(tls, float64(t)))
	j = 0
	for {
		if !(j < N1) {
			break
		}
		/* Apply mid scaling (side is already scaled) */
		l = OpusT_opus_val32(mid * *(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(j)*4)))
		r = *(*OpusT_celt_norm)(unsafe.Pointer(Y + uintptr(j)*4))
		*(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(j)*4)) = OpusT_opus_val32(lgain * (l - r))
		*(*OpusT_celt_norm)(unsafe.Pointer(Y + uintptr(j)*4)) = OpusT_opus_val32(rgain * (l + r))
		goto _7
	_7:
		;
		j = j + 1
	}
}

// C documentation
//
//	/* Decide whether we should spread the pulses in the current frame */
func Opus_spreading_decision(tls *libc.TLS, m uintptr, X uintptr, average uintptr, last_decision int32, hf_average uintptr, tapset_decision uintptr, update_hf int32, end int32, C int32, M int32, spread_weight uintptr) (r int32) {
	var N, N0, c, decision, hf_sum, i, j, nbBands, sum, tmp, v1 int32
	var eBands, x uintptr
	var tcount [3]int32
	var x2N OpusT_opus_val32
	var v5, v6 OpusT_opus_uint32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = N, N0, c, decision, eBands, hf_sum, i, j, nbBands, sum, tcount, tmp, x, x2N, v1, v5, v6
	sum = 0
	nbBands = 0
	eBands = (*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands
	hf_sum = 0
	if !(end > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+5328, __ccgo_ts+5312, int32(480))
	}
	N0 = M * (*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FshortMdctSize
	if M*(int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(end)*2)))-int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(end-int32(1))*2)))) <= int32(8) {
		return SPREAD_NONE
	}
	c = 0
	for {
		i = 0
		for {
			if !(i < end) {
				break
			}
			tmp = 0
			tcount = [3]int32{}
			x = X + uintptr(M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i)*2))))*4 + uintptr(c*N0)*4
			N = M * (int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i+int32(1))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i)*2))))
			if N <= int32(8) {
				goto _3
			}
			/* Compute rough CDF of |x[j]| */
			j = 0
			for {
				if !(j < N) {
					break
				}
				/* Q13 */
				x2N = OpusT_opus_val32(OpusT_celt_norm(*(*OpusT_celt_norm)(unsafe.Pointer(x + uintptr(j)*4))**(*OpusT_celt_norm)(unsafe.Pointer(x + uintptr(j)*4))) * float32(N))
				if x2N < float32(0.25) {
					tcount[0] = tcount[0] + 1
				}
				if x2N < float32(0.0625) {
					tcount[int32(1)] = tcount[int32(1)] + 1
				}
				if x2N < float32(0.015625) {
					tcount[int32(2)] = tcount[int32(2)] + 1
				}
				goto _4
			_4:
				;
				j = j + 1
			}
			/* Only include four last bands (8 kHz and up) */
			if i > (*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands-int32(4) {
				v5 = uint32(N)
				_ = v5 > uint32(0)
				v6 = uint32(int32(32)*(tcount[int32(1)]+tcount[0])) / v5
				goto _7
			_7:
				hf_sum = int32(uint32(hf_sum) + v6)
			}
			tmp = libc.BoolInt32(int32(2)*tcount[int32(2)] >= N) + libc.BoolInt32(int32(2)*tcount[int32(1)] >= N) + libc.BoolInt32(int32(2)*tcount[0] >= N)
			sum = sum + tmp**(*int32)(unsafe.Pointer(spread_weight + uintptr(i)*4))
			nbBands = nbBands + *(*int32)(unsafe.Pointer(spread_weight + uintptr(i)*4))
			goto _3
		_3:
			;
			i = i + 1
		}
		goto _2
	_2:
		;
		c = c + 1
		v1 = c
		if !(v1 < C) {
			break
		}
	}
	if update_hf != 0 {
		if hf_sum != 0 {
			v5 = uint32(C * (int32(4) - (*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands + end))
			_ = v5 > uint32(0)
			v6 = uint32(hf_sum) / v5
			goto _10
		_10:
			hf_sum = int32(v6)
		}
		*(*int32)(unsafe.Pointer(hf_average)) = (*(*int32)(unsafe.Pointer(hf_average)) + hf_sum) >> int32(1)
		hf_sum = *(*int32)(unsafe.Pointer(hf_average))
		if *(*int32)(unsafe.Pointer(tapset_decision)) == int32(2) {
			hf_sum = hf_sum + int32(4)
		} else {
			if *(*int32)(unsafe.Pointer(tapset_decision)) == 0 {
				hf_sum = hf_sum - int32(4)
			}
		}
		if hf_sum > int32(22) {
			*(*int32)(unsafe.Pointer(tapset_decision)) = int32(2)
		} else {
			if hf_sum > int32(18) {
				*(*int32)(unsafe.Pointer(tapset_decision)) = int32(1)
			} else {
				*(*int32)(unsafe.Pointer(tapset_decision)) = 0
			}
		}
	}
	/*printf("%d %d %d\n", hf_sum, *hf_average, *tapset_decision);*/
	if !(nbBands > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+5352, __ccgo_ts+5312, int32(536))
	} /* end has to be non-zero */
	if !(sum >= int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+5380, __ccgo_ts+5312, int32(537))
	}
	v5 = uint32(nbBands)
	_ = v5 > uint32(0)
	v6 = uint32(sum<<int32(8)) / v5
	goto _13
_13:
	sum = int32(v6)
	/* Recursive averaging */
	sum = (sum + *(*int32)(unsafe.Pointer(average))) >> int32(1)
	*(*int32)(unsafe.Pointer(average)) = sum
	/* Hysteresis */
	sum = (int32(3)*sum + ((int32(3)-last_decision)<<int32(7) + int32(64)) + int32(2)) >> int32(2)
	if sum < int32(80) {
		decision = int32(SPREAD_AGGRESSIVE)
	} else {
		if sum < int32(256) {
			decision = int32(SPREAD_NORMAL)
		} else {
			if sum < int32(384) {
				decision = int32(SPREAD_LIGHT)
			} else {
				decision = SPREAD_NONE
			}
		}
	}
	return decision
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

func deinterleave_hadamard(tls *libc.TLS, X uintptr, N0 int32, stride int32, hadamard int32) {
	var N, i, j int32
	var _saved_stack, ordery, st, tmp, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = N, _saved_stack, i, j, ordery, st, tmp, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9
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
	N = N0 * stride
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _6
	_6:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _8
_8:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _10
	_10:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v7 = st
	goto _12
_12:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		goto _14
	_14:
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v11 = st
	goto _16
_16:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		goto _18
	_18:
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v15 = st
	goto _20
_20:
	if !(int64(int32(uint64(uint32(N))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5312, int32(581))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		goto _22
	_22:
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v19 = st
	goto _24
_24:
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32(N)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		goto _26
	_26:
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v23 = st
	goto _28
_28:
	tmp = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32(N))*(uint64(4)/uint64(1)))
	if !(stride > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+5405, __ccgo_ts+5312, int32(582))
	}
	if hadamard != 0 {
		ordery = uintptr(unsafe.Pointer(&ordery_table)) + uintptr(stride)*4 - uintptr(2)*4
		i = 0
		for {
			if !(i < stride) {
				break
			}
			j = 0
			for {
				if !(j < N0) {
					break
				}
				*(*OpusT_celt_norm)(unsafe.Pointer(tmp + uintptr(*(*int32)(unsafe.Pointer(ordery + uintptr(i)*4))*N0+j)*4)) = *(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(j*stride+i)*4))
				goto _30
			_30:
				;
				j = j + 1
			}
			goto _29
		_29:
			;
			i = i + 1
		}
	} else {
		i = 0
		for {
			if !(i < stride) {
				break
			}
			j = 0
			for {
				if !(j < N0) {
					break
				}
				*(*OpusT_celt_norm)(unsafe.Pointer(tmp + uintptr(i*N0+j)*4)) = *(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(j*stride+i)*4))
				goto _32
			_32:
				;
				j = j + 1
			}
			goto _31
		_31:
			;
			i = i + 1
		}
	}
	libc.Xmemcpy(tls, X, tmp, uint64(uint32(N))*uint64(4)+uint64(0*((int64(X)-int64(tmp))/4)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _34
	_34:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _36
_36:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
}

func interleave_hadamard(tls *libc.TLS, X uintptr, N0 int32, stride int32, hadamard int32) {
	var N, i, j int32
	var _saved_stack, ordery, st, tmp, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = N, _saved_stack, i, j, ordery, st, tmp, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9
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
	N = N0 * stride
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _6
	_6:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _8
_8:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _10
	_10:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v7 = st
	goto _12
_12:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		goto _14
	_14:
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v11 = st
	goto _16
_16:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		goto _18
	_18:
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v15 = st
	goto _20
_20:
	if !(int64(int32(uint64(uint32(N))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5312, int32(607))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		goto _22
	_22:
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v19 = st
	goto _24
_24:
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32(N)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		goto _26
	_26:
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v23 = st
	goto _28
_28:
	tmp = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32(N))*(uint64(4)/uint64(1)))
	if hadamard != 0 {
		ordery = uintptr(unsafe.Pointer(&ordery_table)) + uintptr(stride)*4 - uintptr(2)*4
		i = 0
		for {
			if !(i < stride) {
				break
			}
			j = 0
			for {
				if !(j < N0) {
					break
				}
				*(*OpusT_celt_norm)(unsafe.Pointer(tmp + uintptr(j*stride+i)*4)) = *(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(*(*int32)(unsafe.Pointer(ordery + uintptr(i)*4))*N0+j)*4))
				goto _30
			_30:
				;
				j = j + 1
			}
			goto _29
		_29:
			;
			i = i + 1
		}
	} else {
		i = 0
		for {
			if !(i < stride) {
				break
			}
			j = 0
			for {
				if !(j < N0) {
					break
				}
				*(*OpusT_celt_norm)(unsafe.Pointer(tmp + uintptr(j*stride+i)*4)) = *(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(i*N0+j)*4))
				goto _32
			_32:
				;
				j = j + 1
			}
			goto _31
		_31:
			;
			i = i + 1
		}
	}
	libc.Xmemcpy(tls, X, tmp, uint64(uint32(N))*uint64(4)+uint64(0*((int64(X)-int64(tmp))/4)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _34
	_34:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _36
_36:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
}

func Opus_haar1(tls *libc.TLS, X uintptr, N0 int32, stride int32) {
	var i, j int32
	var tmp1, tmp2 OpusT_opus_val32
	_, _, _, _ = i, j, tmp1, tmp2
	N0 = N0 >> int32(1)
	i = 0
	for {
		if !(i < stride) {
			break
		}
		j = 0
		for {
			if !(j < N0) {
				break
			}
			tmp1 = float32(float32(0.70710678) * *(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(stride*int32(2)*j+i)*4)))
			tmp2 = float32(float32(0.70710678) * *(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(stride*(int32(2)*j+int32(1))+i)*4)))
			*(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(stride*int32(2)*j+i)*4)) = tmp1 + tmp2
			*(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(stride*(int32(2)*j+int32(1))+i)*4)) = tmp1 - tmp2
			goto _2
		_2:
			;
			j = j + 1
		}
		goto _1
	_1:
		;
		i = i + 1
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
	goto _3
_3:
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
	Fm                 uintptr
	Fi                 int32
	Fintensity         int32
	Fspread            int32
	Ftf_change         int32
	Fec                uintptr
	Fremaining_bits    OpusT_opus_int32
	FbandE             uintptr
	Fseed              OpusT_opus_uint32
	Farch              int32
	Ftheta_round       int32
	Fdisable_inv       int32
	Favoid_split_noise int32
}

type split_ctx = struct {
	Finv    int32
	Fimid   int32
	Fiside  int32
	Fdelta  int32
	Fitheta int32
	Fqalloc int32
}

func compute_theta(tls *libc.TLS, ctx uintptr, sctx uintptr, X uintptr, Y uintptr, N int32, b uintptr, B int32, B0 int32, LM int32, stereo int32, fill uintptr) {
	var bandE, ec, m uintptr
	var bias, delta, down, encode, fl, fl1, fm, fs, fs1, ft, ft1, i, imid, intensity, inv, iside, itheta, itheta_q30, j, offset, p0, pulse_cap, qalloc, qn, unquantized, x, x0, v1, v5, v6, v7 int32
	var tell OpusT_opus_int32
	var v2, v3 OpusT_opus_uint32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bandE, bias, delta, down, ec, encode, fl, fl1, fm, fs, fs1, ft, ft1, i, imid, intensity, inv, iside, itheta, itheta_q30, j, m, offset, p0, pulse_cap, qalloc, qn, tell, unquantized, x, x0, v1, v2, v3, v5, v6, v7
	itheta = 0
	itheta_q30 = 0
	inv = 0
	encode = (*band_ctx)(unsafe.Pointer(ctx)).Fencode
	m = (*band_ctx)(unsafe.Pointer(ctx)).Fm
	i = (*band_ctx)(unsafe.Pointer(ctx)).Fi
	intensity = (*band_ctx)(unsafe.Pointer(ctx)).Fintensity
	ec = (*band_ctx)(unsafe.Pointer(ctx)).Fec
	bandE = (*band_ctx)(unsafe.Pointer(ctx)).FbandE
	/* Decide on the resolution to give to the split parameter theta */
	pulse_cap = int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FlogN + uintptr(i)*2))) + LM*(int32(1)<<int32(BITRES))
	if stereo != 0 && N == int32(2) {
		v1 = int32(QTHETA_OFFSET_TWOPHASE)
	} else {
		v1 = int32(QTHETA_OFFSET)
	}
	offset = pulse_cap>>int32(1) - v1
	qn = compute_qn(tls, N, *(*int32)(unsafe.Pointer(b)), offset, pulse_cap, stereo)
	if stereo != 0 && i >= intensity {
		qn = int32(1)
	}
	if encode != 0 {
		/* theta is the atan() of the ratio between the (normalized)
		   side and mid. With just that parameter, we can re-scale both
		   mid and side because we know that 1) they have unit norm and
		   2) they are orthogonal. */
		itheta_q30 = Opus_stereo_itheta(tls, X, Y, stereo, N, (*band_ctx)(unsafe.Pointer(ctx)).Farch)
		itheta = itheta_q30 >> int32(16)
	}
	tell = int32(Opus_ec_tell_frac(tls, ec))
	if qn != int32(1) {
		if encode != 0 {
			if !(stereo != 0) || (*band_ctx)(unsafe.Pointer(ctx)).Ftheta_round == 0 {
				itheta = (itheta*qn + int32(8192)) >> int32(14)
				if !(stereo != 0) && (*band_ctx)(unsafe.Pointer(ctx)).Favoid_split_noise != 0 && itheta > 0 && itheta < qn {
					v2 = uint32(qn)
					_ = v2 > uint32(0)
					v3 = uint32(itheta*int32(16384)) / v2
					goto _4
				_4:
					/* Check if the selected value of theta will cause the bit allocation
					   to inject noise on one side. If so, make sure the energy of that side
					   is zero. */
					unquantized = int32(v3)
					imid = int32(Opus_bitexact_cos(tls, int16(unquantized)))
					iside = int32(Opus_bitexact_cos(tls, int16(int32(16384)-unquantized)))
					delta = (int32(16384) + int32(int16((N-int32(1))<<int32(7)))*int32(int16(Opus_bitexact_log2tan(tls, iside, imid)))) >> int32(15)
					if delta > *(*int32)(unsafe.Pointer(b)) {
						itheta = qn
					} else {
						if delta < -*(*int32)(unsafe.Pointer(b)) {
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
				if (*band_ctx)(unsafe.Pointer(ctx)).Ftheta_round < 0 {
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
				Opus_ec_encode(tls, ec, uint32(v1), uint32(v5), uint32(ft))
			} else {
				fs = int32(Opus_ec_decode(tls, ec, uint32(ft)))
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
				Opus_ec_dec_update(tls, ec, uint32(v1), uint32(v5), uint32(ft))
				itheta = x
			}
		} else {
			if B0 > int32(1) || stereo != 0 {
				/* Uniform pdf */
				if encode != 0 {
					Opus_ec_enc_uint(tls, ec, uint32(itheta), uint32(qn+int32(1)))
				} else {
					itheta = int32(Opus_ec_dec_uint(tls, ec, uint32(qn+int32(1))))
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
					Opus_ec_encode(tls, ec, uint32(fl), uint32(fl+fs1), uint32(ft1))
				} else {
					/* Triangular pdf */
					fl1 = 0
					fm = int32(Opus_ec_decode(tls, ec, uint32(ft1)))
					if fm < qn>>int32(1)*(qn>>int32(1)+int32(1))>>int32(1) {
						itheta = int32((Opus_isqrt32(tls, uint32(8)*uint32(fm)+uint32(1)) - uint32(1)) >> int32(1))
						fs1 = itheta + int32(1)
						fl1 = itheta * (itheta + int32(1)) >> int32(1)
					} else {
						itheta = int32((uint32(int32(2)*(qn+int32(1))) - Opus_isqrt32(tls, uint32(8)*uint32(ft1-fm-int32(1))+uint32(1))) >> int32(1))
						fs1 = qn + int32(1) - itheta
						fl1 = ft1 - (qn+int32(1)-itheta)*(qn+int32(2)-itheta)>>int32(1)
					}
					Opus_ec_dec_update(tls, ec, uint32(fl1), uint32(fl1+fs1), uint32(ft1))
				}
			}
		}
		if !(itheta >= int32(0)) {
			Opus_celt_fatal(tls, __ccgo_ts+5460, __ccgo_ts+5312, int32(840))
		}
		v2 = uint32(qn)
		_ = v2 > uint32(0)
		v3 = uint32(itheta*int32(16384)) / v2
		goto _17
	_17:
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
				inv = libc.BoolInt32(itheta > int32(8192) && !((*band_ctx)(unsafe.Pointer(ctx)).Fdisable_inv != 0))
				if inv != 0 {
					j = 0
					for {
						if !(j < N) {
							break
						}
						*(*OpusT_celt_norm)(unsafe.Pointer(Y + uintptr(j)*4)) = -*(*OpusT_celt_norm)(unsafe.Pointer(Y + uintptr(j)*4))
						goto _18
					_18:
						;
						j = j + 1
					}
				}
				intensity_stereo(tls, m, X, Y, bandE, i, N)
			}
			if *(*int32)(unsafe.Pointer(b)) > int32(2)<<int32(BITRES) && (*band_ctx)(unsafe.Pointer(ctx)).Fremaining_bits > int32(2)<<int32(BITRES) {
				if encode != 0 {
					Opus_ec_enc_bit_logp(tls, ec, inv, uint32(2))
				} else {
					inv = Opus_ec_dec_bit_logp(tls, ec, uint32(2))
				}
			} else {
				inv = 0
			}
			/* inv flag override to avoid problems with downmixing. */
			if (*band_ctx)(unsafe.Pointer(ctx)).Fdisable_inv != 0 {
				inv = 0
			}
			itheta = 0
			itheta_q30 = 0
		}
	}
	qalloc = int32(Opus_ec_tell_frac(tls, ec) - uint32(tell))
	*(*int32)(unsafe.Pointer(b)) -= qalloc
	if itheta == 0 {
		imid = int32(32767)
		iside = 0
		*(*int32)(unsafe.Pointer(fill)) &= int32(1)<<B - int32(1)
		delta = -int32(16384)
	} else {
		if itheta == int32(16384) {
			imid = 0
			iside = int32(32767)
			*(*int32)(unsafe.Pointer(fill)) &= (int32(1)<<B - int32(1)) << B
			delta = int32(16384)
		} else {
			imid = int32(Opus_bitexact_cos(tls, int16(itheta)))
			iside = int32(Opus_bitexact_cos(tls, int16(int32(16384)-itheta)))
			/* This is the mid vs side allocation that minimizes squared error
			   in that band. */
			delta = (int32(16384) + int32(int16((N-int32(1))<<int32(7)))*int32(int16(Opus_bitexact_log2tan(tls, iside, imid)))) >> int32(15)
		}
	}
	(*split_ctx)(unsafe.Pointer(sctx)).Finv = inv
	(*split_ctx)(unsafe.Pointer(sctx)).Fimid = imid
	(*split_ctx)(unsafe.Pointer(sctx)).Fiside = iside
	(*split_ctx)(unsafe.Pointer(sctx)).Fdelta = delta
	(*split_ctx)(unsafe.Pointer(sctx)).Fitheta = itheta
	(*split_ctx)(unsafe.Pointer(sctx)).Fqalloc = qalloc
}

func quant_band_n1(tls *libc.TLS, ctx uintptr, X uintptr, Y uintptr, lowband_out uintptr) (r uint32) {
	var c, encode, sign, stereo, v1 int32
	var ec, x uintptr
	var v3 float32
	_, _, _, _, _, _, _, _ = c, ec, encode, sign, stereo, x, v1, v3
	x = X
	encode = (*band_ctx)(unsafe.Pointer(ctx)).Fencode
	ec = (*band_ctx)(unsafe.Pointer(ctx)).Fec
	stereo = libc.BoolInt32(Y != uintptr(uint32(0)))
	c = 0
	for {
		sign = 0
		if (*band_ctx)(unsafe.Pointer(ctx)).Fremaining_bits >= int32(1)<<int32(BITRES) {
			if encode != 0 {
				sign = libc.BoolInt32(*(*OpusT_celt_norm)(unsafe.Pointer(x)) < float32(0))
				Opus_ec_enc_bits(tls, ec, uint32(sign), uint32(1))
			} else {
				sign = int32(Opus_ec_dec_bits(tls, ec, uint32(1)))
			}
			*(*OpusT_opus_int32)(unsafe.Pointer(ctx + 40)) -= int32(1) << int32(BITRES)
		}
		if (*band_ctx)(unsafe.Pointer(ctx)).Fresynth != 0 {
			if sign != 0 {
				v3 = -float32(1)
			} else {
				v3 = float32(1)
			}
			*(*OpusT_celt_norm)(unsafe.Pointer(x)) = v3
		}
		x = Y
		goto _2
	_2:
		;
		c = c + 1
		v1 = c
		if !(v1 < int32(1)+stereo) {
			break
		}
	}
	if lowband_out != 0 {
		*(*OpusT_celt_norm)(unsafe.Pointer(lowband_out)) = *(*OpusT_celt_norm)(unsafe.Pointer(X))
	}
	return uint32(1)
}

// C documentation
//
//	/* This function is responsible for encoding and decoding a mono partition.
//	   It can split the band in two and transmit the energy difference with
//	   the two half-bands. It can be called recursively so bands can end up being
//	   split in 8 parts. */
func quant_partition(tls *libc.TLS, ctx uintptr, X uintptr, N int32, _b int32, B int32, lowband uintptr, LM2 int32, gain OpusT_opus_val32, _fill int32) (r uint32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*int32)(unsafe.Pointer(bp)) = _b
	*(*int32)(unsafe.Pointer(bp + 4)) = _fill
	var B0, K, curr_bits, delta, encode, hi, i1, i2, imid, iside, itheta, j, lo, mbits, mid, q, qalloc, sbits, spread, v1, v2, v3, v4 int32
	var Y, cache, cache1, cache2, ec, m2, next_lowband2, v5 uintptr
	var cm, cm_mask uint32
	var mid1, side OpusT_opus_val32
	var rebalance OpusT_opus_int32
	var tmp, v30 OpusT_opus_val16
	var _ /* sctx at bp+8 */ split_ctx
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = B0, K, Y, cache, cache1, cache2, cm, cm_mask, curr_bits, delta, ec, encode, hi, i1, i2, imid, iside, itheta, j, lo, m2, mbits, mid, mid1, next_lowband2, q, qalloc, rebalance, sbits, side, spread, tmp, v1, v2, v3, v30, v4, v5
	imid = 0
	iside = 0
	B0 = B
	mid1 = float32(0)
	side = float32(0)
	cm = uint32(0)
	Y = uintptr(uint32(0))
	encode = (*band_ctx)(unsafe.Pointer(ctx)).Fencode
	m2 = (*band_ctx)(unsafe.Pointer(ctx)).Fm
	i2 = (*band_ctx)(unsafe.Pointer(ctx)).Fi
	spread = (*band_ctx)(unsafe.Pointer(ctx)).Fspread
	ec = (*band_ctx)(unsafe.Pointer(ctx)).Fec
	/* If we need 1.5 more bit than we can produce, split the band in two. */
	cache2 = (*OpusT_OpusCustomMode)(unsafe.Pointer(m2)).Fcache.Fbits + uintptr(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m2)).Fcache.Findex + uintptr((LM2+int32(1))*(*OpusT_OpusCustomMode)(unsafe.Pointer(m2)).FnbEBands+i2)*2)))
	if LM2 != -int32(1) && *(*int32)(unsafe.Pointer(bp)) > int32(*(*uint8)(unsafe.Pointer(cache2 + uintptr(*(*uint8)(unsafe.Pointer(cache2))))))+int32(12) && N > int32(2) {
		next_lowband2 = uintptr(uint32(0))
		N = N >> int32(1)
		Y = X + uintptr(N)*4
		LM2 = LM2 - int32(1)
		if B == int32(1) {
			*(*int32)(unsafe.Pointer(bp + 4)) = *(*int32)(unsafe.Pointer(bp + 4))&int32(1) | *(*int32)(unsafe.Pointer(bp + 4))<<int32(1)
		}
		B = (B + int32(1)) >> int32(1)
		compute_theta(tls, ctx, bp+8, X, Y, N, bp, B, B0, LM2, 0, bp+4)
		imid = (*(*split_ctx)(unsafe.Pointer(bp + 8))).Fimid
		iside = (*(*split_ctx)(unsafe.Pointer(bp + 8))).Fiside
		delta = (*(*split_ctx)(unsafe.Pointer(bp + 8))).Fdelta
		itheta = (*(*split_ctx)(unsafe.Pointer(bp + 8))).Fitheta
		qalloc = (*(*split_ctx)(unsafe.Pointer(bp + 8))).Fqalloc
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
		if *(*int32)(unsafe.Pointer(bp)) < (*(*int32)(unsafe.Pointer(bp))-delta)/int32(2) {
			v2 = *(*int32)(unsafe.Pointer(bp))
		} else {
			v2 = (*(*int32)(unsafe.Pointer(bp)) - delta) / int32(2)
		}
		if 0 > v2 {
			v1 = 0
		} else {
			if *(*int32)(unsafe.Pointer(bp)) < (*(*int32)(unsafe.Pointer(bp))-delta)/int32(2) {
				v3 = *(*int32)(unsafe.Pointer(bp))
			} else {
				v3 = (*(*int32)(unsafe.Pointer(bp)) - delta) / int32(2)
			}
			v1 = v3
		}
		mbits = v1
		sbits = *(*int32)(unsafe.Pointer(bp)) - mbits
		*(*OpusT_opus_int32)(unsafe.Pointer(ctx + 40)) -= qalloc
		if lowband != 0 {
			next_lowband2 = lowband + uintptr(N)*4
		} /* >32-bit split case */
		rebalance = (*band_ctx)(unsafe.Pointer(ctx)).Fremaining_bits
		if mbits >= sbits {
			cm = quant_partition(tls, ctx, X, N, mbits, B, lowband, LM2, OpusT_opus_val32(gain*mid1), *(*int32)(unsafe.Pointer(bp + 4)))
			rebalance = mbits - (rebalance - (*band_ctx)(unsafe.Pointer(ctx)).Fremaining_bits)
			if rebalance > int32(3)<<int32(BITRES) && itheta != 0 {
				sbits = sbits + (rebalance - int32(3)<<int32(BITRES))
			}
			cm = cm | quant_partition(tls, ctx, Y, N, sbits, B, next_lowband2, LM2, OpusT_opus_val32(gain*side), *(*int32)(unsafe.Pointer(bp + 4))>>B)<<(B0>>int32(1))
		} else {
			cm = quant_partition(tls, ctx, Y, N, sbits, B, next_lowband2, LM2, OpusT_opus_val32(gain*side), *(*int32)(unsafe.Pointer(bp + 4))>>B) << (B0 >> int32(1))
			rebalance = sbits - (rebalance - (*band_ctx)(unsafe.Pointer(ctx)).Fremaining_bits)
			if rebalance > int32(3)<<int32(BITRES) && itheta != int32(16384) {
				mbits = mbits + (rebalance - int32(3)<<int32(BITRES))
			}
			cm = cm | quant_partition(tls, ctx, X, N, mbits, B, lowband, LM2, OpusT_opus_val32(gain*mid1), *(*int32)(unsafe.Pointer(bp + 4)))
		}
	} else {
		/* This is the basic no-split case */
		v5 = m2
		v1 = LM2
		v2 = *(*int32)(unsafe.Pointer(bp))
		v1 = v1 + 1
		cache = (*OpusT_OpusCustomMode)(unsafe.Pointer(v5)).Fcache.Fbits + uintptr(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(v5)).Fcache.Findex + uintptr(v1*(*OpusT_OpusCustomMode)(unsafe.Pointer(v5)).FnbEBands+i2)*2)))
		lo = 0
		hi = int32(*(*uint8)(unsafe.Pointer(cache)))
		v2 = v2 - 1
		i1 = int32(0)
		for {
			if !(i1 < int32(LOG_MAX_PSEUDO)) {
				break
			}
			mid = (lo + hi + int32(1)) >> int32(1)
			if int32(*(*uint8)(unsafe.Pointer(cache + uintptr(mid)))) >= v2 {
				hi = mid
			} else {
				lo = mid
			}
			goto _8
		_8:
			;
			i1 = i1 + 1
		}
		if lo == 0 {
			v3 = -int32(1)
		} else {
			v3 = int32(*(*uint8)(unsafe.Pointer(cache + uintptr(lo))))
		}
		if v2-v3 <= int32(*(*uint8)(unsafe.Pointer(cache + uintptr(hi))))-v2 {
			v4 = lo
			goto _11
		} else {
			v4 = hi
			goto _11
		}
	_11:
		q = v4
		v5 = m2
		v1 = LM2
		v2 = q
		v1 = v1 + 1
		cache1 = (*OpusT_OpusCustomMode)(unsafe.Pointer(v5)).Fcache.Fbits + uintptr(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(v5)).Fcache.Findex + uintptr(v1*(*OpusT_OpusCustomMode)(unsafe.Pointer(v5)).FnbEBands+i2)*2)))
		if v2 == 0 {
			v4 = 0
		} else {
			v4 = int32(*(*uint8)(unsafe.Pointer(cache1 + uintptr(v2)))) + int32(1)
		}
		v3 = v4
		goto _16
	_16:
		curr_bits = v3
		*(*OpusT_opus_int32)(unsafe.Pointer(ctx + 40)) -= curr_bits
		/* Ensures we can never bust the budget */
		for (*band_ctx)(unsafe.Pointer(ctx)).Fremaining_bits < 0 && q > 0 {
			*(*OpusT_opus_int32)(unsafe.Pointer(ctx + 40)) += curr_bits
			q = q - 1
			v5 = m2
			v1 = LM2
			v2 = q
			v1 = v1 + 1
			cache1 = (*OpusT_OpusCustomMode)(unsafe.Pointer(v5)).Fcache.Fbits + uintptr(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(v5)).Fcache.Findex + uintptr(v1*(*OpusT_OpusCustomMode)(unsafe.Pointer(v5)).FnbEBands+i2)*2)))
			if v2 == 0 {
				v4 = 0
			} else {
				v4 = int32(*(*uint8)(unsafe.Pointer(cache1 + uintptr(v2)))) + int32(1)
			}
			v3 = v4
			goto _22
		_22:
			curr_bits = v3
			*(*OpusT_opus_int32)(unsafe.Pointer(ctx + 40)) -= curr_bits
		}
		if q != 0 {
			v1 = q
			if v1 < int32(8) {
				v3 = v1
			} else {
				v3 = (int32(8) + v1&int32(7)) << (v1>>int32(3) - int32(1))
			}
			v2 = v3
			goto _26
		_26:
			K = v2
			/* Finally do the actual quantization */
			if encode != 0 {
				cm = Opus_alg_quant(tls, X, N, K, spread, B, ec, gain, (*band_ctx)(unsafe.Pointer(ctx)).Fresynth, (*band_ctx)(unsafe.Pointer(ctx)).Farch)
			} else {
				cm = Opus_alg_unquant(tls, X, N, K, spread, B, ec, gain)
			}
		} else {
			if (*band_ctx)(unsafe.Pointer(ctx)).Fresynth != 0 {
				/* B can be as large as 16, so this shift might overflow an int on a
				   16-bit platform; use a long to get defined behavior.*/
				cm_mask = uint32(uint64(1)<<B) - uint32(1)
				*(*int32)(unsafe.Pointer(bp + 4)) = int32(uint32(*(*int32)(unsafe.Pointer(bp + 4))) & cm_mask)
				if !(*(*int32)(unsafe.Pointer(bp + 4)) != 0) {
					libc.Xmemset(tls, X, 0, uint64(uint32(N))*uint64(4))
				} else {
					if lowband == uintptr(uint32(0)) {
						/* Noise */
						j = 0
						for {
							if !(j < N) {
								break
							}
							(*band_ctx)(unsafe.Pointer(ctx)).Fseed = Opus_celt_lcg_rand(tls, (*band_ctx)(unsafe.Pointer(ctx)).Fseed)
							*(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(j)*4)) = float32(int32((*band_ctx)(unsafe.Pointer(ctx)).Fseed) >> int32(20))
							goto _28
						_28:
							;
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
							(*band_ctx)(unsafe.Pointer(ctx)).Fseed = Opus_celt_lcg_rand(tls, (*band_ctx)(unsafe.Pointer(ctx)).Fseed)
							/* About 48 dB below the "normal" folding level */
							tmp = float32(1) / float32(256)
							if (*band_ctx)(unsafe.Pointer(ctx)).Fseed&uint32(0x8000) != 0 {
								v30 = tmp
							} else {
								v30 = -tmp
							}
							tmp = v30
							*(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(j)*4)) = *(*OpusT_celt_norm)(unsafe.Pointer(lowband + uintptr(j)*4)) + tmp
							goto _29
						_29:
							;
							j = j + 1
						}
						cm = uint32(*(*int32)(unsafe.Pointer(bp + 4)))
					}
					Opus_renormalise_vector(tls, X, N, gain, (*band_ctx)(unsafe.Pointer(ctx)).Farch)
				}
			}
		}
	}
	return cm
}

// C documentation
//
//	/* This function is responsible for encoding and decoding a band for the mono case. */
func quant_band(tls *libc.TLS, ctx uintptr, X uintptr, N int32, b int32, B int32, lowband uintptr, LM int32, lowband_out uintptr, gain OpusT_opus_val32, lowband_scratch uintptr, fill int32) (r uint32) {
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
	encode = (*band_ctx)(unsafe.Pointer(ctx)).Fencode
	tf_change = (*band_ctx)(unsafe.Pointer(ctx)).Ftf_change
	longBlocks = libc.BoolInt32(B0 == int32(1))
	v1 = uint32(B)
	_ = v1 > uint32(0)
	v2 = uint32(N_B) / v1
	goto _3
_3:
	N_B = int32(v2)
	/* Special case for one sample */
	if N == int32(1) {
		return quant_band_n1(tls, ctx, X, uintptr(uint32(0)), lowband_out)
	}
	if tf_change > 0 {
		recombine = tf_change
	}
	/* Band recombining to increase frequency resolution */
	if lowband_scratch != 0 && lowband != 0 && (recombine != 0 || N_B&int32(1) == 0 && tf_change < 0 || B0 > int32(1)) {
		libc.Xmemcpy(tls, lowband_scratch, lowband, uint64(uint32(N))*uint64(4)+uint64(0*((int64(lowband_scratch)-int64(lowband))/4)))
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
		if lowband != 0 {
			Opus_haar1(tls, lowband, N>>k, int32(1)<<k)
		}
		fill = int32(bit_interleave_table[fill&int32(0xF)]) | int32(bit_interleave_table[fill>>int32(4)])<<int32(2)
		goto _4
	_4:
		;
		k = k + 1
	}
	B = B >> recombine
	N_B = N_B << recombine
	/* Increasing the time resolution */
	for N_B&int32(1) == 0 && tf_change < 0 {
		if encode != 0 {
			Opus_haar1(tls, X, N_B, B)
		}
		if lowband != 0 {
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
		if lowband != 0 {
			deinterleave_hadamard(tls, lowband, N_B>>recombine, B0<<recombine, longBlocks)
		}
	}
	cm = quant_partition(tls, ctx, X, N, b, B, lowband, LM, gain, fill)
	/* This code is used by the decoder and by the resynthesis-enabled encoder */
	if (*band_ctx)(unsafe.Pointer(ctx)).Fresynth != 0 {
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
			goto _5
		_5:
			;
			k = k + 1
		}
		k = 0
		for {
			if !(k < recombine) {
				break
			}
			cm = uint32(bit_deinterleave_table[cm])
			Opus_haar1(tls, X, N0>>k, int32(1)<<k)
			goto _6
		_6:
			;
			k = k + 1
		}
		B = B << recombine
		/* Scale output for later folding */
		if lowband_out != 0 {
			n1 = float32(libc.Xsqrt(tls, float64(N0)))
			j = 0
			for {
				if !(j < N0) {
					break
				}
				*(*OpusT_celt_norm)(unsafe.Pointer(lowband_out + uintptr(j)*4)) = OpusT_opus_val16(n1 * *(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(j)*4)))
				goto _7
			_7:
				;
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
func quant_band_stereo(tls *libc.TLS, ctx uintptr, X uintptr, Y uintptr, N int32, _b int32, B int32, lowband uintptr, LM int32, lowband_out uintptr, lowband_scratch uintptr, _fill int32) (r uint32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	*(*int32)(unsafe.Pointer(bp)) = _b
	*(*int32)(unsafe.Pointer(bp + 4)) = _fill
	var c, delta, encode, imid, inv, iside, itheta, j, mbits, orig_fill, qalloc, sbits, sign, v3, v4, v5 int32
	var cm uint32
	var ec, x2, y2, v1 uintptr
	var mid, side OpusT_opus_val32
	var rebalance OpusT_opus_int32
	var tmp OpusT_celt_norm
	var _ /* sctx at bp+8 */ split_ctx
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, cm, delta, ec, encode, imid, inv, iside, itheta, j, mbits, mid, orig_fill, qalloc, rebalance, sbits, side, sign, tmp, x2, y2, v1, v3, v4, v5
	imid = 0
	iside = 0
	inv = 0
	mid = float32(0)
	side = float32(0)
	cm = uint32(0)
	encode = (*band_ctx)(unsafe.Pointer(ctx)).Fencode
	ec = (*band_ctx)(unsafe.Pointer(ctx)).Fec
	/* Special case for one sample */
	if N == int32(1) {
		return quant_band_n1(tls, ctx, X, Y, lowband_out)
	}
	orig_fill = *(*int32)(unsafe.Pointer(bp + 4))
	if encode != 0 {
		if *(*OpusT_celt_ener)(unsafe.Pointer((*band_ctx)(unsafe.Pointer(ctx)).FbandE + uintptr((*band_ctx)(unsafe.Pointer(ctx)).Fi)*4)) < float32(1e-10) || *(*OpusT_celt_ener)(unsafe.Pointer((*band_ctx)(unsafe.Pointer(ctx)).FbandE + uintptr((*OpusT_OpusCustomMode)(unsafe.Pointer((*band_ctx)(unsafe.Pointer(ctx)).Fm)).FnbEBands+(*band_ctx)(unsafe.Pointer(ctx)).Fi)*4)) < float32(1e-10) {
			if *(*OpusT_celt_ener)(unsafe.Pointer((*band_ctx)(unsafe.Pointer(ctx)).FbandE + uintptr((*band_ctx)(unsafe.Pointer(ctx)).Fi)*4)) > *(*OpusT_celt_ener)(unsafe.Pointer((*band_ctx)(unsafe.Pointer(ctx)).FbandE + uintptr((*OpusT_OpusCustomMode)(unsafe.Pointer((*band_ctx)(unsafe.Pointer(ctx)).Fm)).FnbEBands+(*band_ctx)(unsafe.Pointer(ctx)).Fi)*4)) {
				libc.Xmemcpy(tls, Y, X, uint64(uint32(N))*uint64(4)+uint64(0*((int64(Y)-int64(X))/4)))
			} else {
				libc.Xmemcpy(tls, X, Y, uint64(uint32(N))*uint64(4)+uint64(0*((int64(X)-int64(Y))/4)))
			}
		}
	}
	compute_theta(tls, ctx, bp+8, X, Y, N, bp, B, B, LM, int32(1), bp+4)
	inv = (*(*split_ctx)(unsafe.Pointer(bp + 8))).Finv
	imid = (*(*split_ctx)(unsafe.Pointer(bp + 8))).Fimid
	iside = (*(*split_ctx)(unsafe.Pointer(bp + 8))).Fiside
	delta = (*(*split_ctx)(unsafe.Pointer(bp + 8))).Fdelta
	itheta = (*(*split_ctx)(unsafe.Pointer(bp + 8))).Fitheta
	qalloc = (*(*split_ctx)(unsafe.Pointer(bp + 8))).Fqalloc
	mid = OpusT_opus_val32(float32(1) / float32(32768) * float32(imid))
	side = OpusT_opus_val32(float32(1) / float32(32768) * float32(iside))
	/* This is a special case for N=2 that only works for stereo and takes
	   advantage of the fact that mid and side are orthogonal to encode
	   the side with just one bit. */
	if N == int32(2) {
		sign = 0
		mbits = *(*int32)(unsafe.Pointer(bp))
		sbits = 0
		/* Only need one bit for the side. */
		if itheta != 0 && itheta != int32(16384) {
			sbits = int32(1) << int32(BITRES)
		}
		mbits = mbits - sbits
		c = libc.BoolInt32(itheta > int32(8192))
		*(*OpusT_opus_int32)(unsafe.Pointer(ctx + 40)) -= qalloc + sbits
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
				sign = libc.BoolInt32(OpusT_celt_norm(*(*OpusT_celt_norm)(unsafe.Pointer(x2))**(*OpusT_celt_norm)(unsafe.Pointer(y2 + 1*4)))-OpusT_celt_norm(*(*OpusT_celt_norm)(unsafe.Pointer(x2 + 1*4))**(*OpusT_celt_norm)(unsafe.Pointer(y2))) < float32(0))
				Opus_ec_enc_bits(tls, ec, uint32(sign), uint32(1))
			} else {
				sign = int32(Opus_ec_dec_bits(tls, ec, uint32(1)))
			}
		}
		sign = int32(1) - int32(2)*sign
		/* We use orig_fill here because we want to fold the side, but if
		   itheta==16384, we'll have cleared the low bits of fill. */
		cm = quant_band(tls, ctx, x2, N, mbits, B, lowband, LM, lowband_out, float32(1), lowband_scratch, orig_fill)
		/* We don't split N=2 bands, so cm is either 1 or 0 (for a fold-collapse),
		   and there's no need to worry about mixing with the other channel. */
		*(*OpusT_celt_norm)(unsafe.Pointer(y2)) = OpusT_celt_norm(float32(-sign) * *(*OpusT_celt_norm)(unsafe.Pointer(x2 + 1*4)))
		*(*OpusT_celt_norm)(unsafe.Pointer(y2 + 1*4)) = OpusT_celt_norm(float32(sign) * *(*OpusT_celt_norm)(unsafe.Pointer(x2)))
		if (*band_ctx)(unsafe.Pointer(ctx)).Fresynth != 0 {
			*(*OpusT_celt_norm)(unsafe.Pointer(X)) = OpusT_opus_val32(mid * *(*OpusT_celt_norm)(unsafe.Pointer(X)))
			*(*OpusT_celt_norm)(unsafe.Pointer(X + 1*4)) = OpusT_opus_val32(mid * *(*OpusT_celt_norm)(unsafe.Pointer(X + 1*4)))
			*(*OpusT_celt_norm)(unsafe.Pointer(Y)) = OpusT_opus_val32(side * *(*OpusT_celt_norm)(unsafe.Pointer(Y)))
			*(*OpusT_celt_norm)(unsafe.Pointer(Y + 1*4)) = OpusT_opus_val32(side * *(*OpusT_celt_norm)(unsafe.Pointer(Y + 1*4)))
			tmp = *(*OpusT_celt_norm)(unsafe.Pointer(X))
			*(*OpusT_celt_norm)(unsafe.Pointer(X)) = tmp - *(*OpusT_celt_norm)(unsafe.Pointer(Y))
			*(*OpusT_celt_norm)(unsafe.Pointer(Y)) = tmp + *(*OpusT_celt_norm)(unsafe.Pointer(Y))
			tmp = *(*OpusT_celt_norm)(unsafe.Pointer(X + 1*4))
			*(*OpusT_celt_norm)(unsafe.Pointer(X + 1*4)) = tmp - *(*OpusT_celt_norm)(unsafe.Pointer(Y + 1*4))
			*(*OpusT_celt_norm)(unsafe.Pointer(Y + 1*4)) = tmp + *(*OpusT_celt_norm)(unsafe.Pointer(Y + 1*4))
		}
	} else {
		if *(*int32)(unsafe.Pointer(bp)) < (*(*int32)(unsafe.Pointer(bp))-delta)/int32(2) {
			v4 = *(*int32)(unsafe.Pointer(bp))
		} else {
			v4 = (*(*int32)(unsafe.Pointer(bp)) - delta) / int32(2)
		}
		if 0 > v4 {
			v3 = 0
		} else {
			if *(*int32)(unsafe.Pointer(bp)) < (*(*int32)(unsafe.Pointer(bp))-delta)/int32(2) {
				v5 = *(*int32)(unsafe.Pointer(bp))
			} else {
				v5 = (*(*int32)(unsafe.Pointer(bp)) - delta) / int32(2)
			}
			v3 = v5
		}
		mbits = v3
		sbits = *(*int32)(unsafe.Pointer(bp)) - mbits
		*(*OpusT_opus_int32)(unsafe.Pointer(ctx + 40)) -= qalloc
		rebalance = (*band_ctx)(unsafe.Pointer(ctx)).Fremaining_bits
		if mbits >= sbits {
			/* In stereo mode, we do not apply a scaling to the mid because we need the normalized
			   mid for folding later. */
			cm = quant_band(tls, ctx, X, N, mbits, B, lowband, LM, lowband_out, float32(1), lowband_scratch, *(*int32)(unsafe.Pointer(bp + 4)))
			rebalance = mbits - (rebalance - (*band_ctx)(unsafe.Pointer(ctx)).Fremaining_bits)
			if rebalance > int32(3)<<int32(BITRES) && itheta != 0 {
				sbits = sbits + (rebalance - int32(3)<<int32(BITRES))
			}
			/* For a stereo split, the high bits of fill are always zero, so no
			   folding will be done to the side. */
			cm = cm | quant_band(tls, ctx, Y, N, sbits, B, uintptr(uint32(0)), LM, uintptr(uint32(0)), side, uintptr(uint32(0)), *(*int32)(unsafe.Pointer(bp + 4))>>B)
		} else {
			/* For a stereo split, the high bits of fill are always zero, so no
			   folding will be done to the side. */
			cm = quant_band(tls, ctx, Y, N, sbits, B, uintptr(uint32(0)), LM, uintptr(uint32(0)), side, uintptr(uint32(0)), *(*int32)(unsafe.Pointer(bp + 4))>>B)
			rebalance = sbits - (rebalance - (*band_ctx)(unsafe.Pointer(ctx)).Fremaining_bits)
			if rebalance > int32(3)<<int32(BITRES) && itheta != int32(16384) {
				mbits = mbits + (rebalance - int32(3)<<int32(BITRES))
			}
			/* In stereo mode, we do not apply a scaling to the mid because we need the normalized
			   mid for folding later. */
			cm = cm | quant_band(tls, ctx, X, N, mbits, B, lowband, LM, lowband_out, float32(1), lowband_scratch, *(*int32)(unsafe.Pointer(bp + 4)))
		}
	}
	/* This code is used by the decoder and by the resynthesis-enabled encoder */
	if (*band_ctx)(unsafe.Pointer(ctx)).Fresynth != 0 {
		if N != int32(2) {
			stereo_merge(tls, X, Y, mid, N, (*band_ctx)(unsafe.Pointer(ctx)).Farch)
		}
		if inv != 0 {
			j = 0
			for {
				if !(j < N) {
					break
				}
				*(*OpusT_celt_norm)(unsafe.Pointer(Y + uintptr(j)*4)) = -*(*OpusT_celt_norm)(unsafe.Pointer(Y + uintptr(j)*4))
				goto _6
			_6:
				;
				j = j + 1
			}
		}
	}
	return cm
}

func special_hybrid_folding(tls *libc.TLS, m uintptr, norm uintptr, norm2 uintptr, start int32, M int32, dual_stereo int32) {
	var eBands uintptr
	var n1, n2 int32
	_, _, _ = eBands, n1, n2
	eBands = (*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands
	n1 = M * (int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(start+int32(1))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(start)*2))))
	n2 = M * (int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(start+int32(2))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(start+int32(1))*2))))
	/* Duplicate enough of the first band folding data to be able to fold the second band.
	   Copies no data for CELT-only mode. */
	libc.Xmemcpy(tls, norm+uintptr(n1)*4, norm+uintptr(int32(2)*n1-n2)*4, uint64(uint32(n2-n1))*uint64(4)+uint64(0*((OpusT___predefined_ptrdiff_t(norm+uintptr(n1)*4)-OpusT___predefined_ptrdiff_t(norm+uintptr(int32(2)*n1-n2)*4))/4)))
	if dual_stereo != 0 {
		libc.Xmemcpy(tls, norm2+uintptr(n1)*4, norm2+uintptr(int32(2)*n1-n2)*4, uint64(uint32(n2-n1))*uint64(4)+uint64(0*((OpusT___predefined_ptrdiff_t(norm2+uintptr(n1)*4)-OpusT___predefined_ptrdiff_t(norm2+uintptr(int32(2)*n1-n2)*4))/4)))
	}
}

func Opus_quant_all_bands(tls *libc.TLS, encode int32, m uintptr, start int32, end int32, X_ uintptr, Y_ uintptr, collapse_masks uintptr, bandE uintptr, pulses uintptr, shortBlocks int32, spread int32, dual_stereo int32, intensity int32, tf_res uintptr, total_bits OpusT_opus_int32, balance OpusT_opus_int32, ec uintptr, LM int32, codedBands int32, seed uintptr, complexity int32, arch int32, disable_inv int32) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var B, C, M, N1, b, effective_lowband, fold_end, fold_i, fold_start, i, i1, j, last, lowband_offset, nend_bytes, norm_offset, nstart_bytes, resynth, resynth_alloc, save_bytes, tf_change, theta_rdo, update_lowband, v1, v183, v196, v201, v203, v207, v6 int32
	var X, X_save, X_save2, Y, Y_save, Y_save2, _lowband_scratch, _norm, _saved_stack, bytes_buf, bytes_save, eBands, lowband_scratch, norm, norm2, norm_save2, st, v11, v13, v15, v17, v19, v2, v21, v23, v25, v4, v7, v9 uintptr
	var cm, cm2, x_cm, y_cm, v217 uint32
	var ctx_save, ctx_save2 band_ctx
	var curr_balance, remaining_bits, tell, v204, v205 OpusT_opus_int32
	var dist0, dist1, xy, v229, v232 OpusT_opus_val32
	var ec_save, ec_save2 OpusT_ec_ctx
	var _ /* ctx at bp+0 */ band_ctx
	var _ /* w at bp+80 */ [2]OpusT_opus_val16
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = B, C, M, N1, X, X_save, X_save2, Y, Y_save, Y_save2, _lowband_scratch, _norm, _saved_stack, b, bytes_buf, bytes_save, cm, cm2, ctx_save, ctx_save2, curr_balance, dist0, dist1, eBands, ec_save, ec_save2, effective_lowband, fold_end, fold_i, fold_start, i, i1, j, last, lowband_offset, lowband_scratch, nend_bytes, norm, norm2, norm_offset, norm_save2, nstart_bytes, remaining_bits, resynth, resynth_alloc, save_bytes, st, tell, tf_change, theta_rdo, update_lowband, x_cm, xy, y_cm, v1, v11, v13, v15, v17, v183, v19, v196, v2, v201, v203, v204, v205, v207, v21, v217, v229, v23, v232, v25, v4, v6, v7, v9
	eBands = (*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands
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
		goto _3
	_3:
		st = v2
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v4 = st
	goto _5
_5:
	_saved_stack = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v4)).Fglobal_stack
	M = int32(1) << LM
	if shortBlocks != 0 {
		v1 = M
	} else {
		v1 = int32(1)
	}
	B = v1
	norm_offset = M * int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(start)*2)))
	/* No need to allocate norm for the last band because we don't need an
	   output in that band. */
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v2 = libc.Xmalloc(tls, uint64(16))
		goto _8
	_8:
		st = v2
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v4 = st
	goto _10
_10:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v7 = libc.Xmalloc(tls, uint64(16))
		goto _12
	_12:
		st = v7
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v9 = st
	goto _14
_14:
	*(*uintptr)(unsafe.Pointer(v4 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v9)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _16
	_16:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _18
_18:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _20
	_20:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _22
_22:
	if !(int64(int32(uint64(uint32(C*(M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands-int32(1))*2)))-norm_offset)))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5312, int32(1638))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _24
	_24:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _26
_26:
	*(*uintptr)(unsafe.Pointer(v21 + 8)) += uintptr(uint64(uint32(C*(M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands-int32(1))*2)))-norm_offset))) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v23 = libc.Xmalloc(tls, uint64(16))
		goto _28
	_28:
		st = v23
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v25 = st
	goto _30
_30:
	_norm = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v25)).Fglobal_stack - uintptr(uint64(uint32(C*(M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands-int32(1))*2)))-norm_offset)))*(uint64(4)/uint64(1)))
	norm = _norm
	norm2 = norm + uintptr(M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands-int32(1))*2))))*4 - uintptr(norm_offset)*4
	/* For decoding, we can use the last band as scratch space because we don't need that
	   scratch space for the last band and we don't care about the data there until we're
	   decoding the last band. */
	if encode != 0 && resynth != 0 {
		resynth_alloc = M * (int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands-int32(1))*2))))
	} else {
		resynth_alloc = ALLOC_NONE
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v2 = libc.Xmalloc(tls, uint64(16))
		goto _32
	_32:
		st = v2
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v4 = st
	goto _34
_34:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v7 = libc.Xmalloc(tls, uint64(16))
		goto _36
	_36:
		st = v7
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v9 = st
	goto _38
_38:
	*(*uintptr)(unsafe.Pointer(v4 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v9)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _40
	_40:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _42
_42:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _44
	_44:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _46
_46:
	if !(int64(int32(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5312, int32(1649))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _48
	_48:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _50
_50:
	*(*uintptr)(unsafe.Pointer(v21 + 8)) += uintptr(uint64(uint32(resynth_alloc)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v23 = libc.Xmalloc(tls, uint64(16))
		goto _52
	_52:
		st = v23
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v25 = st
	goto _54
_54:
	_lowband_scratch = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v25)).Fglobal_stack - uintptr(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))
	if encode != 0 && resynth != 0 {
		lowband_scratch = _lowband_scratch
	} else {
		lowband_scratch = X_ + uintptr(M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeffEBands-int32(1))*2))))*4
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v2 = libc.Xmalloc(tls, uint64(16))
		goto _56
	_56:
		st = v2
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v4 = st
	goto _58
_58:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v7 = libc.Xmalloc(tls, uint64(16))
		goto _60
	_60:
		st = v7
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v9 = st
	goto _62
_62:
	*(*uintptr)(unsafe.Pointer(v4 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v9)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _64
	_64:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _66
_66:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _68
	_68:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _70
_70:
	if !(int64(int32(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5312, int32(1654))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _72
	_72:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _74
_74:
	*(*uintptr)(unsafe.Pointer(v21 + 8)) += uintptr(uint64(uint32(resynth_alloc)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v23 = libc.Xmalloc(tls, uint64(16))
		goto _76
	_76:
		st = v23
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v25 = st
	goto _78
_78:
	X_save = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v25)).Fglobal_stack - uintptr(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v2 = libc.Xmalloc(tls, uint64(16))
		goto _80
	_80:
		st = v2
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v4 = st
	goto _82
_82:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v7 = libc.Xmalloc(tls, uint64(16))
		goto _84
	_84:
		st = v7
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v9 = st
	goto _86
_86:
	*(*uintptr)(unsafe.Pointer(v4 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v9)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _88
	_88:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _90
_90:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _92
	_92:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _94
_94:
	if !(int64(int32(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5312, int32(1655))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _96
	_96:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _98
_98:
	*(*uintptr)(unsafe.Pointer(v21 + 8)) += uintptr(uint64(uint32(resynth_alloc)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v23 = libc.Xmalloc(tls, uint64(16))
		goto _100
	_100:
		st = v23
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v25 = st
	goto _102
_102:
	Y_save = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v25)).Fglobal_stack - uintptr(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v2 = libc.Xmalloc(tls, uint64(16))
		goto _104
	_104:
		st = v2
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v4 = st
	goto _106
_106:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v7 = libc.Xmalloc(tls, uint64(16))
		goto _108
	_108:
		st = v7
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v9 = st
	goto _110
_110:
	*(*uintptr)(unsafe.Pointer(v4 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v9)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _112
	_112:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _114
_114:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _116
	_116:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _118
_118:
	if !(int64(int32(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5312, int32(1656))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _120
	_120:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _122
_122:
	*(*uintptr)(unsafe.Pointer(v21 + 8)) += uintptr(uint64(uint32(resynth_alloc)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v23 = libc.Xmalloc(tls, uint64(16))
		goto _124
	_124:
		st = v23
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v25 = st
	goto _126
_126:
	X_save2 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v25)).Fglobal_stack - uintptr(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v2 = libc.Xmalloc(tls, uint64(16))
		goto _128
	_128:
		st = v2
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v4 = st
	goto _130
_130:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v7 = libc.Xmalloc(tls, uint64(16))
		goto _132
	_132:
		st = v7
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v9 = st
	goto _134
_134:
	*(*uintptr)(unsafe.Pointer(v4 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v9)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _136
	_136:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _138
_138:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _140
	_140:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _142
_142:
	if !(int64(int32(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5312, int32(1657))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _144
	_144:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _146
_146:
	*(*uintptr)(unsafe.Pointer(v21 + 8)) += uintptr(uint64(uint32(resynth_alloc)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v23 = libc.Xmalloc(tls, uint64(16))
		goto _148
	_148:
		st = v23
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v25 = st
	goto _150
_150:
	Y_save2 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v25)).Fglobal_stack - uintptr(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v2 = libc.Xmalloc(tls, uint64(16))
		goto _152
	_152:
		st = v2
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v4 = st
	goto _154
_154:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v7 = libc.Xmalloc(tls, uint64(16))
		goto _156
	_156:
		st = v7
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v9 = st
	goto _158
_158:
	*(*uintptr)(unsafe.Pointer(v4 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v9)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _160
	_160:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _162
_162:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _164
	_164:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _166
_166:
	if !(int64(int32(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5312, int32(1658))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _168
	_168:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _170
_170:
	*(*uintptr)(unsafe.Pointer(v21 + 8)) += uintptr(uint64(uint32(resynth_alloc)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v23 = libc.Xmalloc(tls, uint64(16))
		goto _172
	_172:
		st = v23
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v25 = st
	goto _174
_174:
	norm_save2 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v25)).Fglobal_stack - uintptr(uint64(uint32(resynth_alloc))*(uint64(4)/uint64(1)))
	lowband_offset = 0
	(*(*band_ctx)(unsafe.Pointer(bp))).FbandE = bandE
	(*(*band_ctx)(unsafe.Pointer(bp))).Fec = ec
	(*(*band_ctx)(unsafe.Pointer(bp))).Fencode = encode
	(*(*band_ctx)(unsafe.Pointer(bp))).Fintensity = intensity
	(*(*band_ctx)(unsafe.Pointer(bp))).Fm = m
	(*(*band_ctx)(unsafe.Pointer(bp))).Fseed = *(*OpusT_opus_uint32)(unsafe.Pointer(seed))
	(*(*band_ctx)(unsafe.Pointer(bp))).Fspread = spread
	(*(*band_ctx)(unsafe.Pointer(bp))).Farch = arch
	(*(*band_ctx)(unsafe.Pointer(bp))).Fdisable_inv = disable_inv
	(*(*band_ctx)(unsafe.Pointer(bp))).Fresynth = resynth
	(*(*band_ctx)(unsafe.Pointer(bp))).Ftheta_round = 0
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v2 = libc.Xmalloc(tls, uint64(16))
		goto _176
	_176:
		st = v2
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v4 = st
	goto _178
_178:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v7 = libc.Xmalloc(tls, uint64(16))
		goto _180
	_180:
		st = v7
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v9 = st
	goto _182
_182:
	*(*uintptr)(unsafe.Pointer(v4 + 8)) += uintptr((uint64(uint32(1)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v9)).Fglobal_stack))) & (uint64(uint32(1)) - uint64(uint32(1))))
	if theta_rdo != 0 {
		v1 = int32(1275)
	} else {
		v1 = ALLOC_NONE
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _185
	_185:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _187
_187:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _189
	_189:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _191
_191:
	if !(int64(int32(uint64(uint32(v1))*(uint64(1)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5312, int32(1679))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _193
	_193:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _195
_195:
	if theta_rdo != 0 {
		v6 = int32(1275)
	} else {
		v6 = ALLOC_NONE
	}
	*(*uintptr)(unsafe.Pointer(v21 + 8)) += uintptr(uint64(uint32(v6)) * (uint64(1) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v23 = libc.Xmalloc(tls, uint64(16))
		goto _198
	_198:
		st = v23
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v25 = st
	goto _200
_200:
	if theta_rdo != 0 {
		v183 = int32(1275)
	} else {
		v183 = ALLOC_NONE
	}
	bytes_save = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v25)).Fglobal_stack - uintptr(uint64(uint32(v183))*(uint64(1)/uint64(1)))
	/* Avoid injecting noise in the first band on transients. */
	(*(*band_ctx)(unsafe.Pointer(bp))).Favoid_split_noise = libc.BoolInt32(B > int32(1))
	i1 = start
	for {
		if !(i1 < end) {
			break
		}
		effective_lowband = -int32(1)
		tf_change = 0
		(*(*band_ctx)(unsafe.Pointer(bp))).Fi = i1
		last = libc.BoolInt32(i1 == end-int32(1))
		X = X_ + uintptr(M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1)*2))))*4
		if Y_ != uintptr(uint32(0)) {
			Y = Y_ + uintptr(M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1)*2))))*4
		} else {
			Y = uintptr(uint32(0))
		}
		N1 = M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1+int32(1))*2))) - M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1)*2)))
		if !(N1 > int32(0)) {
			Opus_celt_fatal(tls, __ccgo_ts+5488, __ccgo_ts+5312, int32(1705))
		}
		tell = int32(Opus_ec_tell_frac(tls, ec))
		/* Compute how many bits we want to allocate to this band */
		if i1 != start {
			balance = balance - tell
		}
		remaining_bits = total_bits - tell - int32(1)
		(*(*band_ctx)(unsafe.Pointer(bp))).Fremaining_bits = remaining_bits
		if i1 <= codedBands-int32(1) {
			if int32(3) < codedBands-i1 {
				v1 = int32(3)
			} else {
				v1 = codedBands - i1
			}
			v204 = v1
			_ = v204 > int32(0)
			v205 = balance / v204
			goto _206
		_206:
			curr_balance = v205
			if remaining_bits+int32(1) < *(*int32)(unsafe.Pointer(pulses + uintptr(i1)*4))+curr_balance {
				v183 = remaining_bits + int32(1)
			} else {
				v183 = *(*int32)(unsafe.Pointer(pulses + uintptr(i1)*4)) + curr_balance
			}
			if int32(16383) < v183 {
				v6 = int32(16383)
			} else {
				if remaining_bits+int32(1) < *(*int32)(unsafe.Pointer(pulses + uintptr(i1)*4))+curr_balance {
					v196 = remaining_bits + int32(1)
				} else {
					v196 = *(*int32)(unsafe.Pointer(pulses + uintptr(i1)*4)) + curr_balance
				}
				v6 = v196
			}
			if 0 > v6 {
				v1 = 0
			} else {
				if remaining_bits+int32(1) < *(*int32)(unsafe.Pointer(pulses + uintptr(i1)*4))+curr_balance {
					v203 = remaining_bits + int32(1)
				} else {
					v203 = *(*int32)(unsafe.Pointer(pulses + uintptr(i1)*4)) + curr_balance
				}
				if int32(16383) < v203 {
					v201 = int32(16383)
				} else {
					if remaining_bits+int32(1) < *(*int32)(unsafe.Pointer(pulses + uintptr(i1)*4))+curr_balance {
						v207 = remaining_bits + int32(1)
					} else {
						v207 = *(*int32)(unsafe.Pointer(pulses + uintptr(i1)*4)) + curr_balance
					}
					v201 = v207
				}
				v1 = v201
			}
			b = v1
		} else {
			b = 0
		}
		if resynth != 0 && (M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1)*2)))-N1 >= M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(start)*2))) || i1 == start+int32(1)) && (update_lowband != 0 || lowband_offset == 0) {
			lowband_offset = i1
		}
		if i1 == start+int32(1) {
			special_hybrid_folding(tls, m, norm, norm2, start, M, dual_stereo)
		}
		tf_change = *(*int32)(unsafe.Pointer(tf_res + uintptr(i1)*4))
		(*(*band_ctx)(unsafe.Pointer(bp))).Ftf_change = tf_change
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
			if 0 > M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(lowband_offset)*2)))-norm_offset-N1 {
				v1 = 0
			} else {
				v1 = M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(lowband_offset)*2))) - norm_offset - N1
			}
			effective_lowband = v1
			fold_start = lowband_offset
			for {
				fold_start = fold_start - 1
				v1 = fold_start
				if !(M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(v1)*2))) > effective_lowband+norm_offset) {
					break
				}
			}
			fold_end = lowband_offset - int32(1)
			for {
				fold_end = fold_end + 1
				v1 = fold_end
				if !(v1 < i1 && M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(fold_end)*2))) < effective_lowband+norm_offset+N1) {
					break
				}
			}
			v217 = uint32(0)
			y_cm = v217
			x_cm = v217
			fold_i = fold_start
			for {
				x_cm = x_cm | uint32(*(*uint8)(unsafe.Pointer(collapse_masks + uintptr(fold_i*C+0))))
				y_cm = y_cm | uint32(*(*uint8)(unsafe.Pointer(collapse_masks + uintptr(fold_i*C+C-int32(1)))))
				goto _219
			_219:
				;
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
					if !(j < M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1)*2)))-norm_offset) {
						break
					}
					*(*OpusT_celt_norm)(unsafe.Pointer(norm + uintptr(j)*4)) = float32(float32(0.5) * (*(*OpusT_celt_norm)(unsafe.Pointer(norm + uintptr(j)*4)) + *(*OpusT_celt_norm)(unsafe.Pointer(norm2 + uintptr(j)*4))))
					goto _221
				_221:
					;
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
				v4 = norm + uintptr(M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1)*2))))*4 - uintptr(norm_offset)*4
			}
			x_cm = quant_band(tls, bp, X, N1, b/int32(2), B, v2, LM, v4, float32(1), lowband_scratch, int32(x_cm))
			if effective_lowband != -int32(1) {
				v2 = norm2 + uintptr(effective_lowband)*4
			} else {
				v2 = uintptr(uint32(0))
			}
			if last != 0 {
				v4 = uintptr(uint32(0))
			} else {
				v4 = norm2 + uintptr(M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1)*2))))*4 - uintptr(norm_offset)*4
			}
			y_cm = quant_band(tls, bp, Y, N1, b/int32(2), B, v2, LM, v4, float32(1), lowband_scratch, int32(y_cm))
		} else {
			if Y != uintptr(uint32(0)) {
				if theta_rdo != 0 && i1 < intensity {
					compute_channel_weights(tls, *(*OpusT_celt_ener)(unsafe.Pointer(bandE + uintptr(i1)*4)), *(*OpusT_celt_ener)(unsafe.Pointer(bandE + uintptr(i1+(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4)), bp+80)
					/* Make a copy. */
					cm = x_cm | y_cm
					ec_save = *(*OpusT_ec_ctx)(unsafe.Pointer(ec))
					ctx_save = *(*band_ctx)(unsafe.Pointer(bp))
					libc.Xmemcpy(tls, X_save, X, uint64(uint32(N1))*uint64(4)+uint64(0*((int64(X_save)-int64(X))/4)))
					libc.Xmemcpy(tls, Y_save, Y, uint64(uint32(N1))*uint64(4)+uint64(0*((int64(Y_save)-int64(Y))/4)))
					/* Encode and round down. */
					(*(*band_ctx)(unsafe.Pointer(bp))).Ftheta_round = -int32(1)
					if effective_lowband != -int32(1) {
						v2 = norm + uintptr(effective_lowband)*4
					} else {
						v2 = uintptr(uint32(0))
					}
					if last != 0 {
						v4 = uintptr(uint32(0))
					} else {
						v4 = norm + uintptr(M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1)*2))))*4 - uintptr(norm_offset)*4
					}
					x_cm = quant_band_stereo(tls, bp, X, Y, N1, b, B, v2, LM, v4, lowband_scratch, int32(cm))
					_ = arch
					xy = float32(0)
					i = int32(0)
					for {
						if !(i < N1) {
							break
						}
						xy = xy + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(X_save + uintptr(i)*4))**(*OpusT_opus_val16)(unsafe.Pointer(X + uintptr(i)*4)))
						goto _228
					_228:
						;
						i = i + 1
					}
					v229 = xy
					goto _230
				_230:
					_ = arch
					xy = float32(0)
					i = int32(0)
					for {
						if !(i < N1) {
							break
						}
						xy = xy + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(Y_save + uintptr(i)*4))**(*OpusT_opus_val16)(unsafe.Pointer(Y + uintptr(i)*4)))
						goto _231
					_231:
						;
						i = i + 1
					}
					v232 = xy
					goto _233
				_233:
					dist0 = OpusT_opus_val16((*(*[2]OpusT_opus_val16)(unsafe.Pointer(bp + 80)))[0]*v229) + OpusT_opus_val16((*(*[2]OpusT_opus_val16)(unsafe.Pointer(bp + 80)))[int32(1)]*v232)
					/* Save first result. */
					cm2 = x_cm
					ec_save2 = *(*OpusT_ec_ctx)(unsafe.Pointer(ec))
					ctx_save2 = *(*band_ctx)(unsafe.Pointer(bp))
					libc.Xmemcpy(tls, X_save2, X, uint64(uint32(N1))*uint64(4)+uint64(0*((int64(X_save2)-int64(X))/4)))
					libc.Xmemcpy(tls, Y_save2, Y, uint64(uint32(N1))*uint64(4)+uint64(0*((int64(Y_save2)-int64(Y))/4)))
					if !(last != 0) {
						libc.Xmemcpy(tls, norm_save2, norm+uintptr(M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1)*2))))*4-uintptr(norm_offset)*4, uint64(uint32(N1))*uint64(4)+uint64(0*((int64(norm_save2)-int64(norm+uintptr(M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1)*2))))*4-uintptr(norm_offset)*4))/4)))
					}
					nstart_bytes = int32(ec_save.Foffs)
					nend_bytes = int32(ec_save.Fstorage)
					bytes_buf = ec_save.Fbuf + uintptr(nstart_bytes)
					save_bytes = nend_bytes - nstart_bytes
					libc.Xmemcpy(tls, bytes_save, bytes_buf, uint64(uint32(save_bytes))*uint64(1)+uint64(0*(int64(bytes_save)-int64(bytes_buf))))
					/* Restore */
					*(*OpusT_ec_ctx)(unsafe.Pointer(ec)) = ec_save
					*(*band_ctx)(unsafe.Pointer(bp)) = ctx_save
					libc.Xmemcpy(tls, X, X_save, uint64(uint32(N1))*uint64(4)+uint64(0*((int64(X)-int64(X_save))/4)))
					libc.Xmemcpy(tls, Y, Y_save, uint64(uint32(N1))*uint64(4)+uint64(0*((int64(Y)-int64(Y_save))/4)))
					if i1 == start+int32(1) {
						special_hybrid_folding(tls, m, norm, norm2, start, M, dual_stereo)
					}
					/* Encode and round up. */
					(*(*band_ctx)(unsafe.Pointer(bp))).Ftheta_round = int32(1)
					if effective_lowband != -int32(1) {
						v2 = norm + uintptr(effective_lowband)*4
					} else {
						v2 = uintptr(uint32(0))
					}
					if last != 0 {
						v4 = uintptr(uint32(0))
					} else {
						v4 = norm + uintptr(M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1)*2))))*4 - uintptr(norm_offset)*4
					}
					x_cm = quant_band_stereo(tls, bp, X, Y, N1, b, B, v2, LM, v4, lowband_scratch, int32(cm))
					_ = arch
					xy = float32(0)
					i = int32(0)
					for {
						if !(i < N1) {
							break
						}
						xy = xy + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(X_save + uintptr(i)*4))**(*OpusT_opus_val16)(unsafe.Pointer(X + uintptr(i)*4)))
						goto _236
					_236:
						;
						i = i + 1
					}
					v229 = xy
					goto _238
				_238:
					_ = arch
					xy = float32(0)
					i = int32(0)
					for {
						if !(i < N1) {
							break
						}
						xy = xy + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(Y_save + uintptr(i)*4))**(*OpusT_opus_val16)(unsafe.Pointer(Y + uintptr(i)*4)))
						goto _239
					_239:
						;
						i = i + 1
					}
					v232 = xy
					goto _241
				_241:
					dist1 = OpusT_opus_val16((*(*[2]OpusT_opus_val16)(unsafe.Pointer(bp + 80)))[0]*v229) + OpusT_opus_val16((*(*[2]OpusT_opus_val16)(unsafe.Pointer(bp + 80)))[int32(1)]*v232)
					if dist0 >= dist1 {
						x_cm = cm2
						*(*OpusT_ec_ctx)(unsafe.Pointer(ec)) = ec_save2
						*(*band_ctx)(unsafe.Pointer(bp)) = ctx_save2
						libc.Xmemcpy(tls, X, X_save2, uint64(uint32(N1))*uint64(4)+uint64(0*((int64(X)-int64(X_save2))/4)))
						libc.Xmemcpy(tls, Y, Y_save2, uint64(uint32(N1))*uint64(4)+uint64(0*((int64(Y)-int64(Y_save2))/4)))
						if !(last != 0) {
							libc.Xmemcpy(tls, norm+uintptr(M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1)*2))))*4-uintptr(norm_offset)*4, norm_save2, uint64(uint32(N1))*uint64(4)+uint64(0*((int64(norm+uintptr(M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1)*2))))*4-uintptr(norm_offset)*4)-int64(norm_save2))/4)))
						}
						libc.Xmemcpy(tls, bytes_buf, bytes_save, uint64(uint32(save_bytes))*uint64(1)+uint64(0*(int64(bytes_buf)-int64(bytes_save))))
					}
				} else {
					(*(*band_ctx)(unsafe.Pointer(bp))).Ftheta_round = 0
					if effective_lowband != -int32(1) {
						v2 = norm + uintptr(effective_lowband)*4
					} else {
						v2 = uintptr(uint32(0))
					}
					if last != 0 {
						v4 = uintptr(uint32(0))
					} else {
						v4 = norm + uintptr(M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1)*2))))*4 - uintptr(norm_offset)*4
					}
					x_cm = quant_band_stereo(tls, bp, X, Y, N1, b, B, v2, LM, v4, lowband_scratch, int32(x_cm|y_cm))
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
					v4 = norm + uintptr(M*int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1)*2))))*4 - uintptr(norm_offset)*4
				}
				x_cm = quant_band(tls, bp, X, N1, b, B, v2, LM, v4, float32(1), lowband_scratch, int32(x_cm|y_cm))
			}
			y_cm = x_cm
		}
		*(*uint8)(unsafe.Pointer(collapse_masks + uintptr(i1*C+0))) = uint8(x_cm)
		*(*uint8)(unsafe.Pointer(collapse_masks + uintptr(i1*C+C-int32(1)))) = uint8(y_cm)
		balance = balance + (*(*int32)(unsafe.Pointer(pulses + uintptr(i1)*4)) + tell)
		/* Update the folding position only as long as we have 1 bit/sample depth. */
		update_lowband = libc.BoolInt32(b > N1<<int32(BITRES))
		/* We only need to avoid noise on a split for the first band. After that, we
		   have folding. */
		(*(*band_ctx)(unsafe.Pointer(bp))).Favoid_split_noise = 0
		goto _202
	_202:
		;
		i1 = i1 + 1
	}
	*(*OpusT_opus_uint32)(unsafe.Pointer(seed)) = (*(*band_ctx)(unsafe.Pointer(bp))).Fseed
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v2 = libc.Xmalloc(tls, uint64(16))
		goto _247
	_247:
		st = v2
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v4 = st
	goto _249
_249:
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
