// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

package opuscc

import (
	"reflect"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

var _ reflect.Type
var _ unsafe.Pointer

func ec_read_byte(tls *libc.TLS, _this uintptr) (r int32) {
	var v1 int32
	var v2 OpusT_opus_uint32
	var v3 uintptr
	_, _, _ = v1, v2, v3
	if (*OpusT_ec_dec)(unsafe.Pointer(_this)).Foffs < (*OpusT_ec_dec)(unsafe.Pointer(_this)).Fstorage {
		v3 = _this + 28
		v2 = *(*OpusT_opus_uint32)(unsafe.Pointer(v3))
		*(*OpusT_opus_uint32)(unsafe.Pointer(v3)) = *(*OpusT_opus_uint32)(unsafe.Pointer(v3)) + 1
		v1 = int32(*(*uint8)(unsafe.Pointer((*OpusT_ec_dec)(unsafe.Pointer(_this)).Fbuf + uintptr(v2))))
	} else {
		v1 = 0
	}
	return v1
}

func ec_read_byte_from_end(tls *libc.TLS, _this uintptr) (r int32) {
	var v1 int32
	var v2 OpusT_opus_uint32
	var v3 uintptr
	_, _, _ = v1, v2, v3
	if (*OpusT_ec_dec)(unsafe.Pointer(_this)).Fend_offs < (*OpusT_ec_dec)(unsafe.Pointer(_this)).Fstorage {
		v3 = _this + 12
		*(*OpusT_opus_uint32)(unsafe.Pointer(v3)) = *(*OpusT_opus_uint32)(unsafe.Pointer(v3)) + 1
		v2 = *(*OpusT_opus_uint32)(unsafe.Pointer(v3))
		v1 = int32(*(*uint8)(unsafe.Pointer((*OpusT_ec_dec)(unsafe.Pointer(_this)).Fbuf + uintptr((*OpusT_ec_dec)(unsafe.Pointer(_this)).Fstorage-v2))))
	} else {
		v1 = 0
	}
	return v1
}

// C documentation
//
//	/*Normalizes the contents of val and rng so that rng lies entirely in the
//	   high-order symbol.*/
func ec_dec_normalize(tls *libc.TLS, _this uintptr) {
	var sym int32
	_ = sym
	/*If the range is too small, rescale it and input some bits.*/
	for (*OpusT_ec_dec)(unsafe.Pointer(_this)).Frng <= uint32(1)<<(int32(EC_CODE_BITS)-int32(1))>>int32(EC_SYM_BITS) {
		*(*int32)(unsafe.Pointer(_this + 24)) += int32(EC_SYM_BITS)
		*(*OpusT_opus_uint32)(unsafe.Pointer(_this + 32)) <<= uint32(int32(EC_SYM_BITS))
		/*Use up the remaining bits from our last symbol.*/
		sym = (*OpusT_ec_dec)(unsafe.Pointer(_this)).Frem
		/*Read the next value from the input.*/
		(*OpusT_ec_dec)(unsafe.Pointer(_this)).Frem = ec_read_byte(tls, _this)
		/*Take the rest of the bits we need from this new symbol.*/
		sym = (sym<<int32(EC_SYM_BITS) | (*OpusT_ec_dec)(unsafe.Pointer(_this)).Frem) >> (int32(EC_SYM_BITS) - ((int32(EC_CODE_BITS)-int32(2))%int32(EC_SYM_BITS) + int32(1)))
		/*And subtract them from val, capped to be less than EC_CODE_TOP.*/
		(*OpusT_ec_dec)(unsafe.Pointer(_this)).Fval = ((*OpusT_ec_dec)(unsafe.Pointer(_this)).Fval<<int32(EC_SYM_BITS) + (uint32(1)<<int32(EC_SYM_BITS)-uint32(1))&uint32(^sym)) & (uint32(1)<<(int32(EC_CODE_BITS)-int32(1)) - uint32(1))
	}
}

func Opus_ec_dec_init(tls *libc.TLS, _this uintptr, _buf uintptr, _storage OpusT_opus_uint32) {
	(*OpusT_ec_dec)(unsafe.Pointer(_this)).Fbuf = _buf
	(*OpusT_ec_dec)(unsafe.Pointer(_this)).Fstorage = _storage
	(*OpusT_ec_dec)(unsafe.Pointer(_this)).Fend_offs = uint32(0)
	(*OpusT_ec_dec)(unsafe.Pointer(_this)).Fend_window = uint32(0)
	(*OpusT_ec_dec)(unsafe.Pointer(_this)).Fnend_bits = 0
	/*This is the offset from which ec_tell() will subtract partial bits.
	  The final value after the ec_dec_normalize() call will be the same as in
	   the encoder, but we have to compensate for the bits that are added there.*/
	(*OpusT_ec_dec)(unsafe.Pointer(_this)).Fnbits_total = int32(EC_CODE_BITS) + int32(1) - (int32(EC_CODE_BITS)-((int32(EC_CODE_BITS)-int32(2))%int32(EC_SYM_BITS)+int32(1)))/int32(EC_SYM_BITS)*int32(EC_SYM_BITS)
	(*OpusT_ec_dec)(unsafe.Pointer(_this)).Foffs = uint32(0)
	(*OpusT_ec_dec)(unsafe.Pointer(_this)).Frng = uint32(1) << ((int32(EC_CODE_BITS)-int32(2))%int32(EC_SYM_BITS) + int32(1))
	(*OpusT_ec_dec)(unsafe.Pointer(_this)).Frem = ec_read_byte(tls, _this)
	(*OpusT_ec_dec)(unsafe.Pointer(_this)).Fval = (*OpusT_ec_dec)(unsafe.Pointer(_this)).Frng - uint32(1) - uint32((*OpusT_ec_dec)(unsafe.Pointer(_this)).Frem>>(int32(EC_SYM_BITS)-((int32(EC_CODE_BITS)-int32(2))%int32(EC_SYM_BITS)+int32(1))))
	(*OpusT_ec_dec)(unsafe.Pointer(_this)).Ferror1 = 0
	/*Normalize the interval.*/
	ec_dec_normalize(tls, _this)
}

func Opus_ec_decode(tls *libc.TLS, _this uintptr, _ft uint32) (r uint32) {
	var s uint32
	var v1, v2 OpusT_opus_uint32
	_, _, _ = s, v1, v2
	v1 = _ft
	_ = v1 > uint32(0)
	v2 = (*OpusT_ec_dec)(unsafe.Pointer(_this)).Frng / v1
	goto _3
_3:
	(*OpusT_ec_dec)(unsafe.Pointer(_this)).Fext = v2
	s = (*OpusT_ec_dec)(unsafe.Pointer(_this)).Fval / (*OpusT_ec_dec)(unsafe.Pointer(_this)).Fext
	return _ft - (s + uint32(1) + (_ft-(s+uint32(1)))&uint32(-libc.BoolInt32(_ft < s+uint32(1))))
}

func Opus_ec_decode_bin(tls *libc.TLS, _this uintptr, _bits uint32) (r uint32) {
	var s uint32
	_ = s
	(*OpusT_ec_dec)(unsafe.Pointer(_this)).Fext = (*OpusT_ec_dec)(unsafe.Pointer(_this)).Frng >> _bits
	s = (*OpusT_ec_dec)(unsafe.Pointer(_this)).Fval / (*OpusT_ec_dec)(unsafe.Pointer(_this)).Fext
	return uint32(1)<<_bits - (s + uint32(1) + (uint32(1)<<_bits-(s+uint32(1)))&uint32(-libc.BoolInt32(uint32(1)<<_bits < s+uint32(1))))
}

func Opus_ec_dec_update(tls *libc.TLS, _this uintptr, _fl uint32, _fh uint32, _ft uint32) {
	var s OpusT_opus_uint32
	var v1 uint32
	_, _ = s, v1
	s = (*OpusT_ec_dec)(unsafe.Pointer(_this)).Fext * (_ft - _fh)
	*(*OpusT_opus_uint32)(unsafe.Pointer(_this + 36)) -= s
	if _fl > uint32(0) {
		v1 = (*OpusT_ec_dec)(unsafe.Pointer(_this)).Fext * (_fh - _fl)
	} else {
		v1 = (*OpusT_ec_dec)(unsafe.Pointer(_this)).Frng - s
	}
	(*OpusT_ec_dec)(unsafe.Pointer(_this)).Frng = v1
	ec_dec_normalize(tls, _this)
}

// C documentation
//
//	/*The probability of having a "one" is 1/(1<<_logp).*/
func Opus_ec_dec_bit_logp(tls *libc.TLS, _this uintptr, _logp uint32) (r1 int32) {
	var d, r, s OpusT_opus_uint32
	var ret int32
	var v1 uint32
	_, _, _, _, _ = d, r, ret, s, v1
	r = (*OpusT_ec_dec)(unsafe.Pointer(_this)).Frng
	d = (*OpusT_ec_dec)(unsafe.Pointer(_this)).Fval
	s = r >> _logp
	ret = libc.BoolInt32(d < s)
	if !(ret != 0) {
		(*OpusT_ec_dec)(unsafe.Pointer(_this)).Fval = d - s
	}
	if ret != 0 {
		v1 = s
	} else {
		v1 = r - s
	}
	(*OpusT_ec_dec)(unsafe.Pointer(_this)).Frng = v1
	ec_dec_normalize(tls, _this)
	return ret
}

func Opus_ec_dec_icdf(tls *libc.TLS, _this uintptr, _icdf uintptr, _ftb uint32) (r1 int32) {
	var d, r, s, t OpusT_opus_uint32
	var ret, v1 int32
	_, _, _, _, _, _ = d, r, ret, s, t, v1
	s = (*OpusT_ec_dec)(unsafe.Pointer(_this)).Frng
	d = (*OpusT_ec_dec)(unsafe.Pointer(_this)).Fval
	r = s >> _ftb
	ret = -int32(1)
	for cond := true; cond; cond = d < s {
		t = s
		ret = ret + 1
		v1 = ret
		s = r * uint32(*(*uint8)(unsafe.Pointer(_icdf + uintptr(v1))))
	}
	(*OpusT_ec_dec)(unsafe.Pointer(_this)).Fval = d - s
	(*OpusT_ec_dec)(unsafe.Pointer(_this)).Frng = t - s
	ec_dec_normalize(tls, _this)
	return ret
}

func Opus_ec_dec_icdf16(tls *libc.TLS, _this uintptr, _icdf uintptr, _ftb uint32) (r1 int32) {
	var d, r, s, t OpusT_opus_uint32
	var ret, v1 int32
	_, _, _, _, _, _ = d, r, ret, s, t, v1
	s = (*OpusT_ec_dec)(unsafe.Pointer(_this)).Frng
	d = (*OpusT_ec_dec)(unsafe.Pointer(_this)).Fval
	r = s >> _ftb
	ret = -int32(1)
	for cond := true; cond; cond = d < s {
		t = s
		ret = ret + 1
		v1 = ret
		s = r * uint32(*(*OpusT_opus_uint16)(unsafe.Pointer(_icdf + uintptr(v1)*2)))
	}
	(*OpusT_ec_dec)(unsafe.Pointer(_this)).Fval = d - s
	(*OpusT_ec_dec)(unsafe.Pointer(_this)).Frng = t - s
	ec_dec_normalize(tls, _this)
	return ret
}

func Opus_ec_dec_uint(tls *libc.TLS, _this uintptr, _ft OpusT_opus_uint32) (r OpusT_opus_uint32) {
	var ft, s uint32
	var ftb int32
	var t OpusT_opus_uint32
	_, _, _, _ = ft, ftb, s, t
	/*In order to optimize EC_ILOG(), it is undefined for the value 0.*/
	if !(_ft > uint32(1)) {
		Opus_celt_fatal(tls, __ccgo_ts+3569, __ccgo_ts+3593, int32(224))
	}
	_ft = _ft - 1
	ftb = int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, _ft)
	if ftb > int32(EC_UINT_BITS) {
		ftb = ftb - int32(EC_UINT_BITS)
		ft = _ft>>ftb + uint32(1)
		s = Opus_ec_decode(tls, _this, ft)
		Opus_ec_dec_update(tls, _this, s, s+uint32(1), ft)
		t = s<<ftb | Opus_ec_dec_bits(tls, _this, uint32(ftb))
		if t <= _ft {
			return t
		}
		(*OpusT_ec_dec)(unsafe.Pointer(_this)).Ferror1 = int32(1)
		return _ft
	} else {
		_ft = _ft + 1
		s = Opus_ec_decode(tls, _this, _ft)
		Opus_ec_dec_update(tls, _this, s, s+uint32(1), _ft)
		return s
	}
	return r
}

func Opus_ec_dec_bits(tls *libc.TLS, _this uintptr, _bits uint32) (r OpusT_opus_uint32) {
	var available int32
	var ret OpusT_opus_uint32
	var window OpusT_ec_window
	var v1 uintptr
	_, _, _, _ = available, ret, window, v1
	window = (*OpusT_ec_dec)(unsafe.Pointer(_this)).Fend_window
	available = (*OpusT_ec_dec)(unsafe.Pointer(_this)).Fnend_bits
	if uint32(available) < _bits {
		for cond := true; cond; cond = available <= int32(4)*int32(CHAR_BIT)-int32(EC_SYM_BITS) {
			window = window | uint32(ec_read_byte_from_end(tls, _this))<<available
			available = available + int32(EC_SYM_BITS)
		}
	}
	ret = window & (uint32(1)<<_bits - uint32(1))
	window = window >> _bits
	available = int32(uint32(available) - _bits)
	(*OpusT_ec_dec)(unsafe.Pointer(_this)).Fend_window = window
	(*OpusT_ec_dec)(unsafe.Pointer(_this)).Fnend_bits = available
	v1 = _this + 24
	*(*int32)(unsafe.Pointer(v1)) = int32(uint32(*(*int32)(unsafe.Pointer(v1))) + _bits)
	return ret
}

var log2_x_norm_coeff7 = [8]float32{
	0: float32(1),
	1: float32(0.8888888955116272),
	2: float32(0.8),
	3: float32(0.7272727489471436),
	4: float32(0.6666666865348816),
	5: float32(0.6153846383094788),
	6: float32(0.5714285969734192),
	7: float32(0.5333333611488342),
}
var log2_y_norm_coeff7 = [8]float32{
	1: float32(0.1699250042438507),
	2: float32(0.32192808389663696),
	3: float32(0.45943161845207214),
	4: float32(0.5849624872207642),
	5: float32(0.7004396915435791),
	6: float32(0.8073549270629883),
	7: float32(0.9068905711174011),
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

/*Although derived separately, the pulse vector coding scheme is equivalent to
   a Pyramid Vector Quantizer \cite{Fis86}.
  Some additional notes about an early version appear at
   https://people.xiph.org/~tterribe/notes/cwrs.html, but the codebook ordering
   and the definitions of some terms have evolved since that was written.

  The conversion from a pulse vector to an integer index (encoding) and back
   (decoding) is governed by two related functions, V(N,K) and U(N,K).

  V(N,K) = the number of combinations, with replacement, of N items, taken K
   at a time, when a sign bit is added to each item taken at least once (i.e.,
   the number of N-dimensional unit pulse vectors with K pulses).
  One way to compute this is via
    V(N,K) = K>0 ? sum(k=1...K,2**k*choose(N,k)*choose(K-1,k-1)) : 1,
   where choose() is the binomial function.
  A table of values for N<10 and K<10 looks like:
  V[10][10] = {
    {1,  0,   0,    0,    0,     0,     0,      0,      0,       0},
    {1,  2,   2,    2,    2,     2,     2,      2,      2,       2},
    {1,  4,   8,   12,   16,    20,    24,     28,     32,      36},
    {1,  6,  18,   38,   66,   102,   146,    198,    258,     326},
    {1,  8,  32,   88,  192,   360,   608,    952,   1408,    1992},
    {1, 10,  50,  170,  450,  1002,  1970,   3530,   5890,    9290},
    {1, 12,  72,  292,  912,  2364,  5336,  10836,  20256,   35436},
    {1, 14,  98,  462, 1666,  4942, 12642,  28814,  59906,  115598},
    {1, 16, 128,  688, 2816,  9424, 27008,  68464, 157184,  332688},
    {1, 18, 162,  978, 4482, 16722, 53154, 148626, 374274,  864146}
  };

  U(N,K) = the number of such combinations wherein N-1 objects are taken at
   most K-1 at a time.
  This is given by
    U(N,K) = sum(k=0...K-1,V(N-1,k))
           = K>0 ? (V(N-1,K-1) + V(N,K-1))/2 : 0.
  The latter expression also makes clear that U(N,K) is half the number of such
   combinations wherein the first object is taken at least once.
  Although it may not be clear from either of these definitions, U(N,K) is the
   natural function to work with when enumerating the pulse vector codebooks,
   not V(N,K).
  U(N,K) is not well-defined for N=0, but with the extension
    U(0,K) = K>0 ? 0 : 1,
   the function becomes symmetric: U(N,K) = U(K,N), with a similar table:
  U[10][10] = {
    {1, 0,  0,   0,    0,    0,     0,     0,      0,      0},
    {0, 1,  1,   1,    1,    1,     1,     1,      1,      1},
    {0, 1,  3,   5,    7,    9,    11,    13,     15,     17},
    {0, 1,  5,  13,   25,   41,    61,    85,    113,    145},
    {0, 1,  7,  25,   63,  129,   231,   377,    575,    833},
    {0, 1,  9,  41,  129,  321,   681,  1289,   2241,   3649},
    {0, 1, 11,  61,  231,  681,  1683,  3653,   7183,  13073},
    {0, 1, 13,  85,  377, 1289,  3653,  8989,  19825,  40081},
    {0, 1, 15, 113,  575, 2241,  7183, 19825,  48639, 108545},
    {0, 1, 17, 145,  833, 3649, 13073, 40081, 108545, 265729}
  };

  With this extension, V(N,K) may be written in terms of U(N,K):
    V(N,K) = U(N,K) + U(N,K+1)
   for all N>=0, K>=0.
  Thus U(N,K+1) represents the number of combinations where the first element
   is positive or zero, and U(N,K) represents the number of combinations where
   it is negative.
  With a large enough table of U(N,K) values, we could write O(N) encoding
   and O(min(N*log(K),N+K)) decoding routines, but such a table would be
   prohibitively large for small embedded devices (K may be as large as 32767
   for small N, and N may be as large as 200).

  Both functions obey the same recurrence relation:
    V(N,K) = V(N-1,K) + V(N,K-1) + V(N-1,K-1),
    U(N,K) = U(N-1,K) + U(N,K-1) + U(N-1,K-1),
   for all N>0, K>0, with different initial conditions at N=0 or K=0.
  This allows us to construct a row of one of the tables above given the
   previous row or the next row.
  Thus we can derive O(NK) encoding and decoding routines with O(K) memory
   using only addition and subtraction.

  When encoding, we build up from the U(2,K) row and work our way forwards.
  When decoding, we need to start at the U(N,K) row and work our way backwards,
   which requires a means of computing U(N,K).
  U(N,K) may be computed from two previous values with the same N:
    U(N,K) = ((2*N-1)*U(N,K-1) - U(N,K-2))/(K-1) + U(N,K-2)
   for all N>1, and since U(N,K) is symmetric, a similar relation holds for two
   previous values with the same K:
    U(N,K>1) = ((2*K-1)*U(N-1,K) - U(N-2,K))/(N-1) + U(N-2,K)
   for all K>1.
  This allows us to construct an arbitrary row of the U(N,K) table by starting
   with the first two values, which are constants.
  This saves roughly 2/3 the work in our O(NK) decoding routine, but costs O(K)
   multiplications.
  Similar relations can be derived for V(N,K), but are not used here.

  For N>0 and K>0, U(N,K) and V(N,K) take on the form of an (N-1)-degree
   polynomial for fixed N.
  The first few are
    U(1,K) = 1,
    U(2,K) = 2*K-1,
    U(3,K) = (2*K-2)*K+1,
    U(4,K) = (((4*K-6)*K+8)*K-3)/3,
    U(5,K) = ((((2*K-4)*K+10)*K-8)*K+3)/3,
   and
    V(1,K) = 2,
    V(2,K) = 4*K,
    V(3,K) = 4*K*K+2,
    V(4,K) = 8*(K*K+2)*K/3,
    V(5,K) = ((4*K*K+20)*K*K+6)/3,
   for all K>0.
  This allows us to derive O(N) encoding and O(N*log(K)) decoding routines for
   small N (and indeed decoding is also O(N) for N<3).

  @ARTICLE{Fis86,
    author="Thomas R. Fischer",
    title="A Pyramid Vector Quantizer",
    journal="IEEE Transactions on Information Theory",
    volume="IT-32",
    number=4,
    pages="568--583",
    month=Jul,
    year=1986
  }*/

/*U(N,K) = U(K,N) := N>0?K>0?U(N-1,K)+U(N,K-1)+U(N-1,K-1):0:K>0?1:0*/
/*V(N,K) := U(N,K)+U(N,K+1) = the number of PVQ codewords for a band of size N
  with K pulses allocated to it.*/

// C documentation
//
//	/*For each V(N,K) supported, we will access element U(min(N,K+1),max(N,K+1)).
//	  Thus, the number of entries in row I is the larger of the maximum number of
//	   pulses we will ever allocate for a given N=I (K=128, or however many fit in
//	   32 bits, whichever is smaller), plus one, and the maximum N for which
//	   K=I-1 pulses fit in 32 bits.
//	  The largest band size in an Opus Custom mode is 208.
//	  Otherwise, we can limit things to the set of N which can be achieved by
//	   splitting a band from a standard Opus mode: 176, 144, 96, 88, 72, 64, 48,
//	   44, 36, 32, 24, 22, 18, 16, 8, 4, 2).*/
var CELT_PVQ_U_DATA = [1272]OpusT_opus_uint32{
	0:    uint32(1),
	177:  uint32(1),
	178:  uint32(1),
	179:  uint32(1),
	180:  uint32(1),
	181:  uint32(1),
	182:  uint32(1),
	183:  uint32(1),
	184:  uint32(1),
	185:  uint32(1),
	186:  uint32(1),
	187:  uint32(1),
	188:  uint32(1),
	189:  uint32(1),
	190:  uint32(1),
	191:  uint32(1),
	192:  uint32(1),
	193:  uint32(1),
	194:  uint32(1),
	195:  uint32(1),
	196:  uint32(1),
	197:  uint32(1),
	198:  uint32(1),
	199:  uint32(1),
	200:  uint32(1),
	201:  uint32(1),
	202:  uint32(1),
	203:  uint32(1),
	204:  uint32(1),
	205:  uint32(1),
	206:  uint32(1),
	207:  uint32(1),
	208:  uint32(1),
	209:  uint32(1),
	210:  uint32(1),
	211:  uint32(1),
	212:  uint32(1),
	213:  uint32(1),
	214:  uint32(1),
	215:  uint32(1),
	216:  uint32(1),
	217:  uint32(1),
	218:  uint32(1),
	219:  uint32(1),
	220:  uint32(1),
	221:  uint32(1),
	222:  uint32(1),
	223:  uint32(1),
	224:  uint32(1),
	225:  uint32(1),
	226:  uint32(1),
	227:  uint32(1),
	228:  uint32(1),
	229:  uint32(1),
	230:  uint32(1),
	231:  uint32(1),
	232:  uint32(1),
	233:  uint32(1),
	234:  uint32(1),
	235:  uint32(1),
	236:  uint32(1),
	237:  uint32(1),
	238:  uint32(1),
	239:  uint32(1),
	240:  uint32(1),
	241:  uint32(1),
	242:  uint32(1),
	243:  uint32(1),
	244:  uint32(1),
	245:  uint32(1),
	246:  uint32(1),
	247:  uint32(1),
	248:  uint32(1),
	249:  uint32(1),
	250:  uint32(1),
	251:  uint32(1),
	252:  uint32(1),
	253:  uint32(1),
	254:  uint32(1),
	255:  uint32(1),
	256:  uint32(1),
	257:  uint32(1),
	258:  uint32(1),
	259:  uint32(1),
	260:  uint32(1),
	261:  uint32(1),
	262:  uint32(1),
	263:  uint32(1),
	264:  uint32(1),
	265:  uint32(1),
	266:  uint32(1),
	267:  uint32(1),
	268:  uint32(1),
	269:  uint32(1),
	270:  uint32(1),
	271:  uint32(1),
	272:  uint32(1),
	273:  uint32(1),
	274:  uint32(1),
	275:  uint32(1),
	276:  uint32(1),
	277:  uint32(1),
	278:  uint32(1),
	279:  uint32(1),
	280:  uint32(1),
	281:  uint32(1),
	282:  uint32(1),
	283:  uint32(1),
	284:  uint32(1),
	285:  uint32(1),
	286:  uint32(1),
	287:  uint32(1),
	288:  uint32(1),
	289:  uint32(1),
	290:  uint32(1),
	291:  uint32(1),
	292:  uint32(1),
	293:  uint32(1),
	294:  uint32(1),
	295:  uint32(1),
	296:  uint32(1),
	297:  uint32(1),
	298:  uint32(1),
	299:  uint32(1),
	300:  uint32(1),
	301:  uint32(1),
	302:  uint32(1),
	303:  uint32(1),
	304:  uint32(1),
	305:  uint32(1),
	306:  uint32(1),
	307:  uint32(1),
	308:  uint32(1),
	309:  uint32(1),
	310:  uint32(1),
	311:  uint32(1),
	312:  uint32(1),
	313:  uint32(1),
	314:  uint32(1),
	315:  uint32(1),
	316:  uint32(1),
	317:  uint32(1),
	318:  uint32(1),
	319:  uint32(1),
	320:  uint32(1),
	321:  uint32(1),
	322:  uint32(1),
	323:  uint32(1),
	324:  uint32(1),
	325:  uint32(1),
	326:  uint32(1),
	327:  uint32(1),
	328:  uint32(1),
	329:  uint32(1),
	330:  uint32(1),
	331:  uint32(1),
	332:  uint32(1),
	333:  uint32(1),
	334:  uint32(1),
	335:  uint32(1),
	336:  uint32(1),
	337:  uint32(1),
	338:  uint32(1),
	339:  uint32(1),
	340:  uint32(1),
	341:  uint32(1),
	342:  uint32(1),
	343:  uint32(1),
	344:  uint32(1),
	345:  uint32(1),
	346:  uint32(1),
	347:  uint32(1),
	348:  uint32(1),
	349:  uint32(1),
	350:  uint32(1),
	351:  uint32(1),
	352:  uint32(1),
	353:  uint32(3),
	354:  uint32(5),
	355:  uint32(7),
	356:  uint32(9),
	357:  uint32(11),
	358:  uint32(13),
	359:  uint32(15),
	360:  uint32(17),
	361:  uint32(19),
	362:  uint32(21),
	363:  uint32(23),
	364:  uint32(25),
	365:  uint32(27),
	366:  uint32(29),
	367:  uint32(31),
	368:  uint32(33),
	369:  uint32(35),
	370:  uint32(37),
	371:  uint32(39),
	372:  uint32(41),
	373:  uint32(43),
	374:  uint32(45),
	375:  uint32(47),
	376:  uint32(49),
	377:  uint32(51),
	378:  uint32(53),
	379:  uint32(55),
	380:  uint32(57),
	381:  uint32(59),
	382:  uint32(61),
	383:  uint32(63),
	384:  uint32(65),
	385:  uint32(67),
	386:  uint32(69),
	387:  uint32(71),
	388:  uint32(73),
	389:  uint32(75),
	390:  uint32(77),
	391:  uint32(79),
	392:  uint32(81),
	393:  uint32(83),
	394:  uint32(85),
	395:  uint32(87),
	396:  uint32(89),
	397:  uint32(91),
	398:  uint32(93),
	399:  uint32(95),
	400:  uint32(97),
	401:  uint32(99),
	402:  uint32(101),
	403:  uint32(103),
	404:  uint32(105),
	405:  uint32(107),
	406:  uint32(109),
	407:  uint32(111),
	408:  uint32(113),
	409:  uint32(115),
	410:  uint32(117),
	411:  uint32(119),
	412:  uint32(121),
	413:  uint32(123),
	414:  uint32(125),
	415:  uint32(127),
	416:  uint32(129),
	417:  uint32(131),
	418:  uint32(133),
	419:  uint32(135),
	420:  uint32(137),
	421:  uint32(139),
	422:  uint32(141),
	423:  uint32(143),
	424:  uint32(145),
	425:  uint32(147),
	426:  uint32(149),
	427:  uint32(151),
	428:  uint32(153),
	429:  uint32(155),
	430:  uint32(157),
	431:  uint32(159),
	432:  uint32(161),
	433:  uint32(163),
	434:  uint32(165),
	435:  uint32(167),
	436:  uint32(169),
	437:  uint32(171),
	438:  uint32(173),
	439:  uint32(175),
	440:  uint32(177),
	441:  uint32(179),
	442:  uint32(181),
	443:  uint32(183),
	444:  uint32(185),
	445:  uint32(187),
	446:  uint32(189),
	447:  uint32(191),
	448:  uint32(193),
	449:  uint32(195),
	450:  uint32(197),
	451:  uint32(199),
	452:  uint32(201),
	453:  uint32(203),
	454:  uint32(205),
	455:  uint32(207),
	456:  uint32(209),
	457:  uint32(211),
	458:  uint32(213),
	459:  uint32(215),
	460:  uint32(217),
	461:  uint32(219),
	462:  uint32(221),
	463:  uint32(223),
	464:  uint32(225),
	465:  uint32(227),
	466:  uint32(229),
	467:  uint32(231),
	468:  uint32(233),
	469:  uint32(235),
	470:  uint32(237),
	471:  uint32(239),
	472:  uint32(241),
	473:  uint32(243),
	474:  uint32(245),
	475:  uint32(247),
	476:  uint32(249),
	477:  uint32(251),
	478:  uint32(253),
	479:  uint32(255),
	480:  uint32(257),
	481:  uint32(259),
	482:  uint32(261),
	483:  uint32(263),
	484:  uint32(265),
	485:  uint32(267),
	486:  uint32(269),
	487:  uint32(271),
	488:  uint32(273),
	489:  uint32(275),
	490:  uint32(277),
	491:  uint32(279),
	492:  uint32(281),
	493:  uint32(283),
	494:  uint32(285),
	495:  uint32(287),
	496:  uint32(289),
	497:  uint32(291),
	498:  uint32(293),
	499:  uint32(295),
	500:  uint32(297),
	501:  uint32(299),
	502:  uint32(301),
	503:  uint32(303),
	504:  uint32(305),
	505:  uint32(307),
	506:  uint32(309),
	507:  uint32(311),
	508:  uint32(313),
	509:  uint32(315),
	510:  uint32(317),
	511:  uint32(319),
	512:  uint32(321),
	513:  uint32(323),
	514:  uint32(325),
	515:  uint32(327),
	516:  uint32(329),
	517:  uint32(331),
	518:  uint32(333),
	519:  uint32(335),
	520:  uint32(337),
	521:  uint32(339),
	522:  uint32(341),
	523:  uint32(343),
	524:  uint32(345),
	525:  uint32(347),
	526:  uint32(349),
	527:  uint32(351),
	528:  uint32(13),
	529:  uint32(25),
	530:  uint32(41),
	531:  uint32(61),
	532:  uint32(85),
	533:  uint32(113),
	534:  uint32(145),
	535:  uint32(181),
	536:  uint32(221),
	537:  uint32(265),
	538:  uint32(313),
	539:  uint32(365),
	540:  uint32(421),
	541:  uint32(481),
	542:  uint32(545),
	543:  uint32(613),
	544:  uint32(685),
	545:  uint32(761),
	546:  uint32(841),
	547:  uint32(925),
	548:  uint32(1013),
	549:  uint32(1105),
	550:  uint32(1201),
	551:  uint32(1301),
	552:  uint32(1405),
	553:  uint32(1513),
	554:  uint32(1625),
	555:  uint32(1741),
	556:  uint32(1861),
	557:  uint32(1985),
	558:  uint32(2113),
	559:  uint32(2245),
	560:  uint32(2381),
	561:  uint32(2521),
	562:  uint32(2665),
	563:  uint32(2813),
	564:  uint32(2965),
	565:  uint32(3121),
	566:  uint32(3281),
	567:  uint32(3445),
	568:  uint32(3613),
	569:  uint32(3785),
	570:  uint32(3961),
	571:  uint32(4141),
	572:  uint32(4325),
	573:  uint32(4513),
	574:  uint32(4705),
	575:  uint32(4901),
	576:  uint32(5101),
	577:  uint32(5305),
	578:  uint32(5513),
	579:  uint32(5725),
	580:  uint32(5941),
	581:  uint32(6161),
	582:  uint32(6385),
	583:  uint32(6613),
	584:  uint32(6845),
	585:  uint32(7081),
	586:  uint32(7321),
	587:  uint32(7565),
	588:  uint32(7813),
	589:  uint32(8065),
	590:  uint32(8321),
	591:  uint32(8581),
	592:  uint32(8845),
	593:  uint32(9113),
	594:  uint32(9385),
	595:  uint32(9661),
	596:  uint32(9941),
	597:  uint32(10225),
	598:  uint32(10513),
	599:  uint32(10805),
	600:  uint32(11101),
	601:  uint32(11401),
	602:  uint32(11705),
	603:  uint32(12013),
	604:  uint32(12325),
	605:  uint32(12641),
	606:  uint32(12961),
	607:  uint32(13285),
	608:  uint32(13613),
	609:  uint32(13945),
	610:  uint32(14281),
	611:  uint32(14621),
	612:  uint32(14965),
	613:  uint32(15313),
	614:  uint32(15665),
	615:  uint32(16021),
	616:  uint32(16381),
	617:  uint32(16745),
	618:  uint32(17113),
	619:  uint32(17485),
	620:  uint32(17861),
	621:  uint32(18241),
	622:  uint32(18625),
	623:  uint32(19013),
	624:  uint32(19405),
	625:  uint32(19801),
	626:  uint32(20201),
	627:  uint32(20605),
	628:  uint32(21013),
	629:  uint32(21425),
	630:  uint32(21841),
	631:  uint32(22261),
	632:  uint32(22685),
	633:  uint32(23113),
	634:  uint32(23545),
	635:  uint32(23981),
	636:  uint32(24421),
	637:  uint32(24865),
	638:  uint32(25313),
	639:  uint32(25765),
	640:  uint32(26221),
	641:  uint32(26681),
	642:  uint32(27145),
	643:  uint32(27613),
	644:  uint32(28085),
	645:  uint32(28561),
	646:  uint32(29041),
	647:  uint32(29525),
	648:  uint32(30013),
	649:  uint32(30505),
	650:  uint32(31001),
	651:  uint32(31501),
	652:  uint32(32005),
	653:  uint32(32513),
	654:  uint32(33025),
	655:  uint32(33541),
	656:  uint32(34061),
	657:  uint32(34585),
	658:  uint32(35113),
	659:  uint32(35645),
	660:  uint32(36181),
	661:  uint32(36721),
	662:  uint32(37265),
	663:  uint32(37813),
	664:  uint32(38365),
	665:  uint32(38921),
	666:  uint32(39481),
	667:  uint32(40045),
	668:  uint32(40613),
	669:  uint32(41185),
	670:  uint32(41761),
	671:  uint32(42341),
	672:  uint32(42925),
	673:  uint32(43513),
	674:  uint32(44105),
	675:  uint32(44701),
	676:  uint32(45301),
	677:  uint32(45905),
	678:  uint32(46513),
	679:  uint32(47125),
	680:  uint32(47741),
	681:  uint32(48361),
	682:  uint32(48985),
	683:  uint32(49613),
	684:  uint32(50245),
	685:  uint32(50881),
	686:  uint32(51521),
	687:  uint32(52165),
	688:  uint32(52813),
	689:  uint32(53465),
	690:  uint32(54121),
	691:  uint32(54781),
	692:  uint32(55445),
	693:  uint32(56113),
	694:  uint32(56785),
	695:  uint32(57461),
	696:  uint32(58141),
	697:  uint32(58825),
	698:  uint32(59513),
	699:  uint32(60205),
	700:  uint32(60901),
	701:  uint32(61601),
	702:  uint32(63),
	703:  uint32(129),
	704:  uint32(231),
	705:  uint32(377),
	706:  uint32(575),
	707:  uint32(833),
	708:  uint32(1159),
	709:  uint32(1561),
	710:  uint32(2047),
	711:  uint32(2625),
	712:  uint32(3303),
	713:  uint32(4089),
	714:  uint32(4991),
	715:  uint32(6017),
	716:  uint32(7175),
	717:  uint32(8473),
	718:  uint32(9919),
	719:  uint32(11521),
	720:  uint32(13287),
	721:  uint32(15225),
	722:  uint32(17343),
	723:  uint32(19649),
	724:  uint32(22151),
	725:  uint32(24857),
	726:  uint32(27775),
	727:  uint32(30913),
	728:  uint32(34279),
	729:  uint32(37881),
	730:  uint32(41727),
	731:  uint32(45825),
	732:  uint32(50183),
	733:  uint32(54809),
	734:  uint32(59711),
	735:  uint32(64897),
	736:  uint32(70375),
	737:  uint32(76153),
	738:  uint32(82239),
	739:  uint32(88641),
	740:  uint32(95367),
	741:  uint32(102425),
	742:  uint32(109823),
	743:  uint32(117569),
	744:  uint32(125671),
	745:  uint32(134137),
	746:  uint32(142975),
	747:  uint32(152193),
	748:  uint32(161799),
	749:  uint32(171801),
	750:  uint32(182207),
	751:  uint32(193025),
	752:  uint32(204263),
	753:  uint32(215929),
	754:  uint32(228031),
	755:  uint32(240577),
	756:  uint32(253575),
	757:  uint32(267033),
	758:  uint32(280959),
	759:  uint32(295361),
	760:  uint32(310247),
	761:  uint32(325625),
	762:  uint32(341503),
	763:  uint32(357889),
	764:  uint32(374791),
	765:  uint32(392217),
	766:  uint32(410175),
	767:  uint32(428673),
	768:  uint32(447719),
	769:  uint32(467321),
	770:  uint32(487487),
	771:  uint32(508225),
	772:  uint32(529543),
	773:  uint32(551449),
	774:  uint32(573951),
	775:  uint32(597057),
	776:  uint32(620775),
	777:  uint32(645113),
	778:  uint32(670079),
	779:  uint32(695681),
	780:  uint32(721927),
	781:  uint32(748825),
	782:  uint32(776383),
	783:  uint32(804609),
	784:  uint32(833511),
	785:  uint32(863097),
	786:  uint32(893375),
	787:  uint32(924353),
	788:  uint32(956039),
	789:  uint32(988441),
	790:  uint32(1021567),
	791:  uint32(1055425),
	792:  uint32(1090023),
	793:  uint32(1125369),
	794:  uint32(1161471),
	795:  uint32(1198337),
	796:  uint32(1235975),
	797:  uint32(1274393),
	798:  uint32(1313599),
	799:  uint32(1353601),
	800:  uint32(1394407),
	801:  uint32(1436025),
	802:  uint32(1478463),
	803:  uint32(1521729),
	804:  uint32(1565831),
	805:  uint32(1610777),
	806:  uint32(1656575),
	807:  uint32(1703233),
	808:  uint32(1750759),
	809:  uint32(1799161),
	810:  uint32(1848447),
	811:  uint32(1898625),
	812:  uint32(1949703),
	813:  uint32(2001689),
	814:  uint32(2054591),
	815:  uint32(2108417),
	816:  uint32(2163175),
	817:  uint32(2218873),
	818:  uint32(2275519),
	819:  uint32(2333121),
	820:  uint32(2391687),
	821:  uint32(2451225),
	822:  uint32(2511743),
	823:  uint32(2573249),
	824:  uint32(2635751),
	825:  uint32(2699257),
	826:  uint32(2763775),
	827:  uint32(2829313),
	828:  uint32(2895879),
	829:  uint32(2963481),
	830:  uint32(3032127),
	831:  uint32(3101825),
	832:  uint32(3172583),
	833:  uint32(3244409),
	834:  uint32(3317311),
	835:  uint32(3391297),
	836:  uint32(3466375),
	837:  uint32(3542553),
	838:  uint32(3619839),
	839:  uint32(3698241),
	840:  uint32(3777767),
	841:  uint32(3858425),
	842:  uint32(3940223),
	843:  uint32(4023169),
	844:  uint32(4107271),
	845:  uint32(4192537),
	846:  uint32(4278975),
	847:  uint32(4366593),
	848:  uint32(4455399),
	849:  uint32(4545401),
	850:  uint32(4636607),
	851:  uint32(4729025),
	852:  uint32(4822663),
	853:  uint32(4917529),
	854:  uint32(5013631),
	855:  uint32(5110977),
	856:  uint32(5209575),
	857:  uint32(5309433),
	858:  uint32(5410559),
	859:  uint32(5512961),
	860:  uint32(5616647),
	861:  uint32(5721625),
	862:  uint32(5827903),
	863:  uint32(5935489),
	864:  uint32(6044391),
	865:  uint32(6154617),
	866:  uint32(6266175),
	867:  uint32(6379073),
	868:  uint32(6493319),
	869:  uint32(6608921),
	870:  uint32(6725887),
	871:  uint32(6844225),
	872:  uint32(6963943),
	873:  uint32(7085049),
	874:  uint32(7207551),
	875:  uint32(321),
	876:  uint32(681),
	877:  uint32(1289),
	878:  uint32(2241),
	879:  uint32(3649),
	880:  uint32(5641),
	881:  uint32(8361),
	882:  uint32(11969),
	883:  uint32(16641),
	884:  uint32(22569),
	885:  uint32(29961),
	886:  uint32(39041),
	887:  uint32(50049),
	888:  uint32(63241),
	889:  uint32(78889),
	890:  uint32(97281),
	891:  uint32(118721),
	892:  uint32(143529),
	893:  uint32(172041),
	894:  uint32(204609),
	895:  uint32(241601),
	896:  uint32(283401),
	897:  uint32(330409),
	898:  uint32(383041),
	899:  uint32(441729),
	900:  uint32(506921),
	901:  uint32(579081),
	902:  uint32(658689),
	903:  uint32(746241),
	904:  uint32(842249),
	905:  uint32(947241),
	906:  uint32(1061761),
	907:  uint32(1186369),
	908:  uint32(1321641),
	909:  uint32(1468169),
	910:  uint32(1626561),
	911:  uint32(1797441),
	912:  uint32(1981449),
	913:  uint32(2179241),
	914:  uint32(2391489),
	915:  uint32(2618881),
	916:  uint32(2862121),
	917:  uint32(3121929),
	918:  uint32(3399041),
	919:  uint32(3694209),
	920:  uint32(4008201),
	921:  uint32(4341801),
	922:  uint32(4695809),
	923:  uint32(5071041),
	924:  uint32(5468329),
	925:  uint32(5888521),
	926:  uint32(6332481),
	927:  uint32(6801089),
	928:  uint32(7295241),
	929:  uint32(7815849),
	930:  uint32(8363841),
	931:  uint32(8940161),
	932:  uint32(9545769),
	933:  uint32(10181641),
	934:  uint32(10848769),
	935:  uint32(11548161),
	936:  uint32(12280841),
	937:  uint32(13047849),
	938:  uint32(13850241),
	939:  uint32(14689089),
	940:  uint32(15565481),
	941:  uint32(16480521),
	942:  uint32(17435329),
	943:  uint32(18431041),
	944:  uint32(19468809),
	945:  uint32(20549801),
	946:  uint32(21675201),
	947:  uint32(22846209),
	948:  uint32(24064041),
	949:  uint32(25329929),
	950:  uint32(26645121),
	951:  uint32(28010881),
	952:  uint32(29428489),
	953:  uint32(30899241),
	954:  uint32(32424449),
	955:  uint32(34005441),
	956:  uint32(35643561),
	957:  uint32(37340169),
	958:  uint32(39096641),
	959:  uint32(40914369),
	960:  uint32(42794761),
	961:  uint32(44739241),
	962:  uint32(46749249),
	963:  uint32(48826241),
	964:  uint32(50971689),
	965:  uint32(53187081),
	966:  uint32(55473921),
	967:  uint32(57833729),
	968:  uint32(60268041),
	969:  uint32(62778409),
	970:  uint32(65366401),
	971:  uint32(68033601),
	972:  uint32(70781609),
	973:  uint32(73612041),
	974:  uint32(76526529),
	975:  uint32(79526721),
	976:  uint32(82614281),
	977:  uint32(85790889),
	978:  uint32(89058241),
	979:  uint32(92418049),
	980:  uint32(95872041),
	981:  uint32(99421961),
	982:  uint32(103069569),
	983:  uint32(106816641),
	984:  uint32(110664969),
	985:  uint32(114616361),
	986:  uint32(118672641),
	987:  uint32(122835649),
	988:  uint32(127107241),
	989:  uint32(131489289),
	990:  uint32(135983681),
	991:  uint32(140592321),
	992:  uint32(145317129),
	993:  uint32(150160041),
	994:  uint32(155123009),
	995:  uint32(160208001),
	996:  uint32(165417001),
	997:  uint32(170752009),
	998:  uint32(176215041),
	999:  uint32(181808129),
	1000: uint32(187533321),
	1001: uint32(193392681),
	1002: uint32(199388289),
	1003: uint32(205522241),
	1004: uint32(211796649),
	1005: uint32(218213641),
	1006: uint32(224775361),
	1007: uint32(231483969),
	1008: uint32(238341641),
	1009: uint32(245350569),
	1010: uint32(252512961),
	1011: uint32(259831041),
	1012: uint32(267307049),
	1013: uint32(274943241),
	1014: uint32(282741889),
	1015: uint32(290705281),
	1016: uint32(298835721),
	1017: uint32(307135529),
	1018: uint32(315607041),
	1019: uint32(324252609),
	1020: uint32(333074601),
	1021: uint32(342075401),
	1022: uint32(351257409),
	1023: uint32(360623041),
	1024: uint32(370174729),
	1025: uint32(379914921),
	1026: uint32(389846081),
	1027: uint32(399970689),
	1028: uint32(410291241),
	1029: uint32(420810249),
	1030: uint32(431530241),
	1031: uint32(442453761),
	1032: uint32(453583369),
	1033: uint32(464921641),
	1034: uint32(476471169),
	1035: uint32(488234561),
	1036: uint32(500214441),
	1037: uint32(512413449),
	1038: uint32(524834241),
	1039: uint32(537479489),
	1040: uint32(550351881),
	1041: uint32(563454121),
	1042: uint32(576788929),
	1043: uint32(590359041),
	1044: uint32(604167209),
	1045: uint32(618216201),
	1046: uint32(632508801),
	1047: uint32(1683),
	1048: uint32(3653),
	1049: uint32(7183),
	1050: uint32(13073),
	1051: uint32(22363),
	1052: uint32(36365),
	1053: uint32(56695),
	1054: uint32(85305),
	1055: uint32(124515),
	1056: uint32(177045),
	1057: uint32(246047),
	1058: uint32(335137),
	1059: uint32(448427),
	1060: uint32(590557),
	1061: uint32(766727),
	1062: uint32(982729),
	1063: uint32(1244979),
	1064: uint32(1560549),
	1065: uint32(1937199),
	1066: uint32(2383409),
	1067: uint32(2908411),
	1068: uint32(3522221),
	1069: uint32(4235671),
	1070: uint32(5060441),
	1071: uint32(6009091),
	1072: uint32(7095093),
	1073: uint32(8332863),
	1074: uint32(9737793),
	1075: uint32(11326283),
	1076: uint32(13115773),
	1077: uint32(15124775),
	1078: uint32(17372905),
	1079: uint32(19880915),
	1080: uint32(22670725),
	1081: uint32(25765455),
	1082: uint32(29189457),
	1083: uint32(32968347),
	1084: uint32(37129037),
	1085: uint32(41699767),
	1086: uint32(46710137),
	1087: uint32(52191139),
	1088: uint32(58175189),
	1089: uint32(64696159),
	1090: uint32(71789409),
	1091: uint32(79491819),
	1092: uint32(87841821),
	1093: uint32(96879431),
	1094: uint32(106646281),
	1095: uint32(117185651),
	1096: uint32(128542501),
	1097: uint32(140763503),
	1098: uint32(153897073),
	1099: uint32(167993403),
	1100: uint32(183104493),
	1101: uint32(199284183),
	1102: uint32(216588185),
	1103: uint32(235074115),
	1104: uint32(254801525),
	1105: uint32(275831935),
	1106: uint32(298228865),
	1107: uint32(322057867),
	1108: uint32(347386557),
	1109: uint32(374284647),
	1110: uint32(402823977),
	1111: uint32(433078547),
	1112: uint32(465124549),
	1113: uint32(499040399),
	1114: uint32(534906769),
	1115: uint32(572806619),
	1116: uint32(612825229),
	1117: uint32(655050231),
	1118: uint32(699571641),
	1119: uint32(746481891),
	1120: uint32(795875861),
	1121: uint32(847850911),
	1122: uint32(902506913),
	1123: uint32(959946283),
	1124: uint32(1020274013),
	1125: uint32(1083597703),
	1126: uint32(1150027593),
	1127: uint32(1219676595),
	1128: uint32(1292660325),
	1129: uint32(1369097135),
	1130: uint32(1449108145),
	1131: uint32(1532817275),
	1132: uint32(1620351277),
	1133: uint32(1711839767),
	1134: uint32(1807415257),
	1135: uint32(1907213187),
	1136: uint32(2011371957),
	1137: uint32(2120032959),
	1138: uint32(8989),
	1139: uint32(19825),
	1140: uint32(40081),
	1141: uint32(75517),
	1142: uint32(134245),
	1143: uint32(227305),
	1144: uint32(369305),
	1145: uint32(579125),
	1146: uint32(880685),
	1147: uint32(1303777),
	1148: uint32(1884961),
	1149: uint32(2668525),
	1150: uint32(3707509),
	1151: uint32(5064793),
	1152: uint32(6814249),
	1153: uint32(9041957),
	1154: uint32(11847485),
	1155: uint32(15345233),
	1156: uint32(19665841),
	1157: uint32(24957661),
	1158: uint32(31388293),
	1159: uint32(39146185),
	1160: uint32(48442297),
	1161: uint32(59511829),
	1162: uint32(72616013),
	1163: uint32(88043969),
	1164: uint32(106114625),
	1165: uint32(127178701),
	1166: uint32(151620757),
	1167: uint32(179861305),
	1168: uint32(212358985),
	1169: uint32(249612805),
	1170: uint32(292164445),
	1171: uint32(340600625),
	1172: uint32(395555537),
	1173: uint32(457713341),
	1174: uint32(527810725),
	1175: uint32(606639529),
	1176: uint32(695049433),
	1177: uint32(793950709),
	1178: uint32(904317037),
	1179: uint32(1027188385),
	1180: uint32(1163673953),
	1181: uint32(1314955181),
	1182: uint32(1482288821),
	1183: uint32(1667010073),
	1184: uint32(1870535785),
	1185: uint32(2094367717),
	1186: uint32(48639),
	1187: uint32(108545),
	1188: uint32(224143),
	1189: uint32(433905),
	1190: uint32(795455),
	1191: uint32(1392065),
	1192: uint32(2340495),
	1193: uint32(3800305),
	1194: uint32(5984767),
	1195: uint32(9173505),
	1196: uint32(13726991),
	1197: uint32(20103025),
	1198: uint32(28875327),
	1199: uint32(40754369),
	1200: uint32(56610575),
	1201: uint32(77500017),
	1202: uint32(104692735),
	1203: uint32(139703809),
	1204: uint32(184327311),
	1205: uint32(240673265),
	1206: uint32(311207743),
	1207: uint32(398796225),
	1208: uint32(506750351),
	1209: uint32(638878193),
	1210: uint32(799538175),
	1211: uint32(993696769),
	1212: uint32(1226990095),
	1213: uint32(1505789553),
	1214: uint32(1837271615),
	1215: uint32(2229491905),
	1216: uint32(265729),
	1217: uint32(598417),
	1218: uint32(1256465),
	1219: uint32(2485825),
	1220: uint32(4673345),
	1221: uint32(8405905),
	1222: uint32(14546705),
	1223: uint32(24331777),
	1224: uint32(39490049),
	1225: uint32(62390545),
	1226: uint32(96220561),
	1227: uint32(145198913),
	1228: uint32(214828609),
	1229: uint32(312193553),
	1230: uint32(446304145),
	1231: uint32(628496897),
	1232: uint32(872893441),
	1233: uint32(1196924561),
	1234: uint32(1621925137),
	1235: uint32(2173806145),
	1236: uint32(1462563),
	1237: uint32(3317445),
	1238: uint32(7059735),
	1239: uint32(14218905),
	1240: uint32(27298155),
	1241: uint32(50250765),
	1242: uint32(89129247),
	1243: uint32(152951073),
	1244: uint32(254831667),
	1245: uint32(413442773),
	1246: uint32(654862247),
	1247: uint32(1014889769),
	1248: uint32(1541911931),
	1249: uint32(2300409629),
	1250: uint32(3375210671),
	1251: uint32(8097453),
	1252: uint32(18474633),
	1253: uint32(39753273),
	1254: uint32(81270333),
	1255: uint32(158819253),
	1256: uint32(298199265),
	1257: uint32(540279585),
	1258: uint32(948062325),
	1259: uint32(1616336765),
	1260: uint32(45046719),
	1261: uint32(103274625),
	1262: uint32(224298231),
	1263: uint32(464387817),
	1264: uint32(921406335),
	1265: uint32(1759885185),
	1266: uint32(3248227095),
	1267: uint32(251595969),
	1268: uint32(579168825),
	1269: uint32(1267854873),
	1270: uint32(2653649025),
	1271: uint32(1409933619),
}

var CELT_PVQ_U_ROW = [15]uintptr{
	0:  uintptr(unsafe.Pointer(&CELT_PVQ_U_DATA)) + uintptr(0)*4,
	1:  uintptr(unsafe.Pointer(&CELT_PVQ_U_DATA)) + uintptr(176)*4,
	2:  uintptr(unsafe.Pointer(&CELT_PVQ_U_DATA)) + uintptr(351)*4,
	3:  uintptr(unsafe.Pointer(&CELT_PVQ_U_DATA)) + uintptr(525)*4,
	4:  uintptr(unsafe.Pointer(&CELT_PVQ_U_DATA)) + uintptr(698)*4,
	5:  uintptr(unsafe.Pointer(&CELT_PVQ_U_DATA)) + uintptr(870)*4,
	6:  uintptr(unsafe.Pointer(&CELT_PVQ_U_DATA)) + uintptr(1041)*4,
	7:  uintptr(unsafe.Pointer(&CELT_PVQ_U_DATA)) + uintptr(1131)*4,
	8:  uintptr(unsafe.Pointer(&CELT_PVQ_U_DATA)) + uintptr(1178)*4,
	9:  uintptr(unsafe.Pointer(&CELT_PVQ_U_DATA)) + uintptr(1207)*4,
	10: uintptr(unsafe.Pointer(&CELT_PVQ_U_DATA)) + uintptr(1226)*4,
	11: uintptr(unsafe.Pointer(&CELT_PVQ_U_DATA)) + uintptr(1240)*4,
	12: uintptr(unsafe.Pointer(&CELT_PVQ_U_DATA)) + uintptr(1248)*4,
	13: uintptr(unsafe.Pointer(&CELT_PVQ_U_DATA)) + uintptr(1254)*4,
	14: uintptr(unsafe.Pointer(&CELT_PVQ_U_DATA)) + uintptr(1257)*4,
}

func icwrs(tls *libc.TLS, _n int32, _y uintptr) (r OpusT_opus_uint32) {
	var i OpusT_opus_uint32
	var j, k, v1, v2 int32
	_, _, _, _, _ = i, j, k, v1, v2
	if !(_n >= int32(2)) {
		Opus_celt_fatal(tls, __ccgo_ts+3610, __ccgo_ts+3634, int32(448))
	}
	j = _n - int32(1)
	i = libc.BoolUint32(*(*int32)(unsafe.Pointer(_y + uintptr(j)*4)) < 0)
	k = libc.Xabs(tls, *(*int32)(unsafe.Pointer(_y + uintptr(j)*4)))
	for cond := true; cond; cond = j > 0 {
		j = j - 1
		if _n-j < k {
			v1 = _n - j
		} else {
			v1 = k
		}
		if _n-j > k {
			v2 = _n - j
		} else {
			v2 = k
		}
		i = i + *(*OpusT_opus_uint32)(unsafe.Pointer(CELT_PVQ_U_ROW[v1] + uintptr(v2)*4))
		k = k + libc.Xabs(tls, *(*int32)(unsafe.Pointer(_y + uintptr(j)*4)))
		if *(*int32)(unsafe.Pointer(_y + uintptr(j)*4)) < 0 {
			if _n-j < k+int32(1) {
				v1 = _n - j
			} else {
				v1 = k + int32(1)
			}
			if _n-j > k+int32(1) {
				v2 = _n - j
			} else {
				v2 = k + int32(1)
			}
			i = i + *(*OpusT_opus_uint32)(unsafe.Pointer(CELT_PVQ_U_ROW[v1] + uintptr(v2)*4))
		}
	}
	return i
}

func Opus_encode_pulses(tls *libc.TLS, _y uintptr, _n int32, _k int32, _enc uintptr) {
	var v1, v2, v3, v4 int32
	_, _, _, _ = v1, v2, v3, v4
	if !(_k > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+3649, __ccgo_ts+3634, int32(463))
	}
	if _n < _k {
		v1 = _n
	} else {
		v1 = _k
	}
	if _n > _k {
		v2 = _n
	} else {
		v2 = _k
	}
	if _n < _k+int32(1) {
		v3 = _n
	} else {
		v3 = _k + int32(1)
	}
	if _n > _k+int32(1) {
		v4 = _n
	} else {
		v4 = _k + int32(1)
	}
	Opus_ec_enc_uint(tls, _enc, icwrs(tls, _n, _y), *(*OpusT_opus_uint32)(unsafe.Pointer(CELT_PVQ_U_ROW[v1] + uintptr(v2)*4))+*(*OpusT_opus_uint32)(unsafe.Pointer(CELT_PVQ_U_ROW[v3] + uintptr(v4)*4)))
}

func cwrsi(tls *libc.TLS, _n int32, _k int32, _i OpusT_opus_uint32, _y uintptr) (r OpusT_opus_val32) {
	var k0, s, v1 int32
	var p, q OpusT_opus_uint32
	var row, v3 uintptr
	var val OpusT_opus_int16
	var yy OpusT_opus_val32
	_, _, _, _, _, _, _, _, _ = k0, p, q, row, s, val, yy, v1, v3
	yy = float32(0)
	if !(_k > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+3649, __ccgo_ts+3634, int32(473))
	}
	if !(_n > int32(1)) {
		Opus_celt_fatal(tls, __ccgo_ts+3672, __ccgo_ts+3634, int32(474))
	}
	for _n > int32(2) {
		/*Lots of pulses case:*/
		if _k >= _n {
			row = CELT_PVQ_U_ROW[_n]
			/*Are the pulses in this dimension negative?*/
			p = *(*OpusT_opus_uint32)(unsafe.Pointer(row + uintptr(_k+int32(1))*4))
			s = -libc.BoolInt32(_i >= p)
			_i = _i - p&uint32(s)
			/*Count how many pulses were placed in this dimension.*/
			k0 = _k
			q = *(*OpusT_opus_uint32)(unsafe.Pointer(row + uintptr(_n)*4))
			if q > _i {
				_ = p > q
				_k = _n
				for cond := true; cond; cond = p > _i {
					_k = _k - 1
					v1 = _k
					p = *(*OpusT_opus_uint32)(unsafe.Pointer(CELT_PVQ_U_ROW[v1] + uintptr(_n)*4))
				}
			} else {
				p = *(*OpusT_opus_uint32)(unsafe.Pointer(row + uintptr(_k)*4))
				for {
					if !(p > _i) {
						break
					}
					_k = _k - 1
					goto _2
				_2:
					;
					p = *(*OpusT_opus_uint32)(unsafe.Pointer(row + uintptr(_k)*4))
				}
			}
			_i = _i - p
			val = int16(k0 - _k + s ^ s)
			v3 = _y
			_y += 4
			*(*int32)(unsafe.Pointer(v3)) = int32(val)
			yy = yy + OpusT_opus_val32(float32(val)*float32(val))
		} else {
			/*Are there any pulses in this dimension at all?*/
			p = *(*OpusT_opus_uint32)(unsafe.Pointer(CELT_PVQ_U_ROW[_k] + uintptr(_n)*4))
			q = *(*OpusT_opus_uint32)(unsafe.Pointer(CELT_PVQ_U_ROW[_k+int32(1)] + uintptr(_n)*4))
			if p <= _i && _i < q {
				_i = _i - p
				v3 = _y
				_y += 4
				*(*int32)(unsafe.Pointer(v3)) = 0
			} else {
				/*Are the pulses in this dimension negative?*/
				s = -libc.BoolInt32(_i >= q)
				_i = _i - q&uint32(s)
				/*Count how many pulses were placed in this dimension.*/
				k0 = _k
				for cond := true; cond; cond = p > _i {
					_k = _k - 1
					v1 = _k
					p = *(*OpusT_opus_uint32)(unsafe.Pointer(CELT_PVQ_U_ROW[v1] + uintptr(_n)*4))
				}
				_i = _i - p
				val = int16(k0 - _k + s ^ s)
				v3 = _y
				_y += 4
				*(*int32)(unsafe.Pointer(v3)) = int32(val)
				yy = yy + OpusT_opus_val32(float32(val)*float32(val))
			}
		}
		_n = _n - 1
	}
	/*_n==2*/
	p = uint32(int32(2)*_k + int32(1))
	s = -libc.BoolInt32(_i >= p)
	_i = _i - p&uint32(s)
	k0 = _k
	_k = int32((_i + uint32(1)) >> int32(1))
	if _k != 0 {
		_i = _i - uint32(int32(2)*_k-int32(1))
	}
	val = int16(k0 - _k + s ^ s)
	v3 = _y
	_y += 4
	*(*int32)(unsafe.Pointer(v3)) = int32(val)
	yy = yy + OpusT_opus_val32(float32(val)*float32(val))
	/*_n==1*/
	s = -int32(_i)
	val = int16(_k + s ^ s)
	*(*int32)(unsafe.Pointer(_y)) = int32(val)
	yy = yy + OpusT_opus_val32(float32(val)*float32(val))
	return yy
}

func Opus_decode_pulses(tls *libc.TLS, _y uintptr, _n int32, _k int32, _dec uintptr) (r OpusT_opus_val32) {
	var v1, v2, v3, v4 int32
	_, _, _, _ = v1, v2, v3, v4
	if _n < _k {
		v1 = _n
	} else {
		v1 = _k
	}
	if _n > _k {
		v2 = _n
	} else {
		v2 = _k
	}
	if _n < _k+int32(1) {
		v3 = _n
	} else {
		v3 = _k + int32(1)
	}
	if _n > _k+int32(1) {
		v4 = _n
	} else {
		v4 = _k + int32(1)
	}
	return cwrsi(tls, _n, _k, Opus_ec_dec_uint(tls, _dec, *(*OpusT_opus_uint32)(unsafe.Pointer(CELT_PVQ_U_ROW[v1] + uintptr(v2)*4))+*(*OpusT_opus_uint32)(unsafe.Pointer(CELT_PVQ_U_ROW[v3] + uintptr(v4)*4))), _y)
}

const CELT_SIG_SCALE7 = 32768
const DECODER_RESET_START = "rng"
const DECODE_BUFFER_SIZE = "DEC_PITCH_BUF_SIZE"
const FRAME_DRED = 5
const FRAME_NONE = 0
const FRAME_NORMAL = 1
const FRAME_PLC_NEURAL = 4
const FRAME_PLC_NOISE = 2
const FRAME_PLC_PERIODIC = 3
const PLC_PITCH_LAG_MAX = 720
const PLC_PITCH_LAG_MIN = 100
const PLC_UPDATE_FRAMES = 4
const Q15ONE1 = 1
const Q31ONE1 = 1
const VERY_SMALL1 = 1e-30
const qext_bytes = 0

type OpusT_OpusCustomDecoder = struct {
	Fmode                  uintptr
	Foverlap               int32
	Fchannels              int32
	Fstream_channels       int32
	Fdownsample            int32
	Fstart                 int32
	Fend                   int32
	Fsignalling            int32
	Fdisable_inv           int32
	Fcomplexity            int32
	Farch                  int32
	Frng                   OpusT_opus_uint32
	Ferror1                int32
	Flast_pitch_index      int32
	Floss_duration         int32
	Fplc_duration          int32
	Flast_frame_type       int32
	Fskip_plc              int32
	Fpostfilter_period     int32
	Fpostfilter_period_old int32
	Fpostfilter_gain       OpusT_opus_val16
	Fpostfilter_gain_old   OpusT_opus_val16
	Fpostfilter_tapset     int32
	Fpostfilter_tapset_old int32
	Fprefilter_and_fold    int32
	Fpreemph_memD          [2]OpusT_celt_sig
	F_decode_mem           [1]OpusT_celt_sig
}

var trim_icdf9 = [11]uint8{
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
var spread_icdf9 = [4]uint8{
	0: uint8(25),
	1: uint8(23),
	2: uint8(2),
}
var tapset_icdf9 = [3]uint8{
	0: uint8(2),
	1: uint8(1),
}

var log2_x_norm_coeff8 = [8]float32{
	0: float32(1),
	1: float32(0.8888888955116272),
	2: float32(0.8),
	3: float32(0.7272727489471436),
	4: float32(0.6666666865348816),
	5: float32(0.6153846383094788),
	6: float32(0.5714285969734192),
	7: float32(0.5333333611488342),
}
var log2_y_norm_coeff8 = [8]float32{
	1: float32(0.1699250042438507),
	2: float32(0.32192808389663696),
	3: float32(0.45943161845207214),
	4: float32(0.5849624872207642),
	5: float32(0.7004396915435791),
	6: float32(0.8073549270629883),
	7: float32(0.9068905711174011),
}

/* The maximum pitch lag to allow in the pitch-based PLC. It's possible to save
   CPU time in the PLC pitch search by making this smaller than MAX_PERIOD. The
   current value corresponds to a pitch of 66.67 Hz. */
/* The minimum pitch lag to allow in the pitch-based PLC. This corresponds to a
   pitch of 480 Hz. */

/**********************************************************************/
/*                                                                    */
/*                             DECODER                                */
/*                                                                    */
/**********************************************************************/

/*
  - Decoder state
    @brief Decoder state
*/
type OpusCustomDecoder = struct {
	Fmode                  uintptr
	Foverlap               int32
	Fchannels              int32
	Fstream_channels       int32
	Fdownsample            int32
	Fstart                 int32
	Fend                   int32
	Fsignalling            int32
	Fdisable_inv           int32
	Fcomplexity            int32
	Farch                  int32
	Frng                   OpusT_opus_uint32
	Ferror1                int32
	Flast_pitch_index      int32
	Floss_duration         int32
	Fplc_duration          int32
	Flast_frame_type       int32
	Fskip_plc              int32
	Fpostfilter_period     int32
	Fpostfilter_period_old int32
	Fpostfilter_gain       OpusT_opus_val16
	Fpostfilter_gain_old   OpusT_opus_val16
	Fpostfilter_tapset     int32
	Fpostfilter_tapset_old int32
	Fprefilter_and_fold    int32
	Fpreemph_memD          [2]OpusT_celt_sig
	F_decode_mem           [1]OpusT_celt_sig
}

// C documentation
//
//	/* Make basic checks on the CELT state to ensure we don't end
//	   up writing all over memory. */
