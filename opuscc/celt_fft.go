// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

package opuscc

import (
	"reflect"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

var _ reflect.Type

func kf_bfly2(tls *libc.TLS, out *OpusT_kiss_fft_cpx, m, N int32) {
	// The no-custom-modes radix-2 stage always follows a radix-4 stage.
	if m != 4 {
		opusCeltFatal(tls, opusDiagnosticString(3470), opusDiagnosticString(3493), 80)
	}
	if N <= 0 {
		return
	}
	data := unsafe.Slice(out, 8*N)
	const tw = float32(0.7071067812)
	for i := int32(0); i < N; i++ {
		f := data[8*i : 8*i+8]
		for j := 0; j < 4; j++ {
			t := f[j+4]
			switch j {
			case 1:
				t = OpusT_kiss_fft_cpx{Fr: float32((t.Fr + t.Fi) * tw), Fi: float32((t.Fi - t.Fr) * tw)}
			case 2:
				t = OpusT_kiss_fft_cpx{Fr: t.Fi, Fi: -t.Fr}
			case 3:
				t = OpusT_kiss_fft_cpx{Fr: float32((t.Fi - t.Fr) * tw), Fi: float32(-(t.Fi + t.Fr) * tw)}
			}
			f[j+4].Fr = f[j].Fr - t.Fr
			f[j+4].Fi = f[j].Fi - t.Fi
			f[j].Fr += t.Fr
			f[j].Fi += t.Fi
		}
	}
}

// Twiddles are explicit so a local table is not hidden in the legacy FFT state.
func kf_bfly4(tls *libc.TLS, out *OpusT_kiss_fft_cpx, stride OpusT_size_t, twiddles *OpusT_kiss_twiddle_cpx, m, N, mm int32) {
	if N <= 0 {
		return
	}
	if m == 1 {
		data := unsafe.Slice(out, 4*N)
		for i := int32(0); i < N; i++ {
			f := data[4*i : 4*i+4]
			s0 := fftSub(f[0], f[2])
			f[0] = fftAdd(f[0], f[2])
			s1 := fftAdd(f[1], f[3])
			f[2] = fftSub(f[0], s1)
			f[0] = fftAdd(f[0], s1)
			s1 = fftSub(f[1], f[3])
			f[1] = OpusT_kiss_fft_cpx{Fr: s0.Fr + s1.Fi, Fi: s0.Fi - s1.Fr}
			f[3] = OpusT_kiss_fft_cpx{Fr: s0.Fr - s1.Fi, Fi: s0.Fi + s1.Fr}
		}
		return
	}
	if m <= 0 {
		return
	}
	data := unsafe.Slice(out, (N-1)*mm+4*m)
	tw := unsafe.Slice(twiddles, 3*uint64(m-1)*stride+1)
	for i := int32(0); i < N; i++ {
		f := data[i*mm : i*mm+4*m]
		for j := int32(0); j < m; j++ {
			s0 := fftMul(f[j+m], tw[uint64(j)*stride])
			s1 := fftMul(f[j+2*m], tw[2*uint64(j)*stride])
			s2 := fftMul(f[j+3*m], tw[3*uint64(j)*stride])
			s5 := fftSub(f[j], s1)
			f[j] = fftAdd(f[j], s1)
			s3 := fftAdd(s0, s2)
			s4 := fftSub(s0, s2)
			f[j+2*m] = fftSub(f[j], s3)
			f[j] = fftAdd(f[j], s3)
			f[j+m] = OpusT_kiss_fft_cpx{Fr: s5.Fr + s4.Fi, Fi: s5.Fi - s4.Fr}
			f[j+3*m] = OpusT_kiss_fft_cpx{Fr: s5.Fr - s4.Fi, Fi: s5.Fi + s4.Fr}
		}
	}
}

func fftAdd(a, b OpusT_kiss_fft_cpx) OpusT_kiss_fft_cpx {
	return OpusT_kiss_fft_cpx{Fr: a.Fr + b.Fr, Fi: a.Fi + b.Fi}
}
func fftSub(a, b OpusT_kiss_fft_cpx) OpusT_kiss_fft_cpx {
	return OpusT_kiss_fft_cpx{Fr: a.Fr - b.Fr, Fi: a.Fi - b.Fi}
}
func fftMul(a OpusT_kiss_fft_cpx, b OpusT_kiss_twiddle_cpx) OpusT_kiss_fft_cpx {
	// Explicit product rounding prevents fusion across C's float32 boundaries.
	return OpusT_kiss_fft_cpx{Fr: float32(a.Fr*b.Fr) - float32(a.Fi*b.Fi), Fi: float32(a.Fr*b.Fi) + float32(a.Fi*b.Fr)}
}

func kf_bfly3(tls *libc.TLS, out *OpusT_kiss_fft_cpx, stride OpusT_size_t, twiddles *OpusT_kiss_twiddle_cpx, m, N, mm int32) {
	// C requires positive m; no output or twiddle access for an empty Go call.
	if N <= 0 || m <= 0 {
		return
	}
	data := unsafe.Slice(out, (N-1)*mm+3*m)
	maxTw := uint64(m) * stride
	if v := 2 * uint64(m-1) * stride; v > maxTw {
		maxTw = v
	}
	tw := unsafe.Slice(twiddles, maxTw+1)
	epi3 := tw[uint64(m)*stride]
	for i := int32(0); i < N; i++ {
		f := data[i*mm : i*mm+3*m]
		for j := int32(0); j < m; j++ {
			s1 := fftMul(f[j+m], tw[uint64(j)*stride])
			s2 := fftMul(f[j+2*m], tw[2*uint64(j)*stride])
			s3 := fftAdd(s1, s2)
			s0 := fftSub(s1, s2)
			f[j+m] = OpusT_kiss_fft_cpx{Fr: f[j].Fr - float32(s3.Fr*.5), Fi: f[j].Fi - float32(s3.Fi*.5)}
			// Explicit rounding prevents ARM64 from fusing these products with
			// the following additions/subtractions (C stores scratch first).
			s0.Fr = float32(s0.Fr * epi3.Fi)
			s0.Fi = float32(s0.Fi * epi3.Fi)
			f[j] = fftAdd(f[j], s3)
			f[j+2*m] = OpusT_kiss_fft_cpx{Fr: f[j+m].Fr + s0.Fi, Fi: f[j+m].Fi - s0.Fr}
			f[j+m].Fr -= s0.Fi
			f[j+m].Fi += s0.Fr
		}
	}
}

func kf_bfly5(tls *libc.TLS, out *OpusT_kiss_fft_cpx, stride OpusT_size_t, twiddles *OpusT_kiss_twiddle_cpx, m, N, mm int32) {
	if N <= 0 || m <= 0 {
		return
	}
	data := unsafe.Slice(out, (N-1)*mm+5*m)
	maxTw := 2 * uint64(m) * stride
	if v := 4 * uint64(m-1) * stride; v > maxTw {
		maxTw = v
	}
	tw := unsafe.Slice(twiddles, maxTw+1)
	ya := tw[uint64(m)*stride]
	yb := tw[2*uint64(m)*stride]
	for i := int32(0); i < N; i++ {
		f := data[i*mm : i*mm+5*m]
		for u := int32(0); u < m; u++ {
			var s [13]OpusT_kiss_fft_cpx
			s[0] = f[u]
			for k := int32(1); k <= 4; k++ {
				s[k] = fftMul(f[u+k*m], tw[uint64(k*u)*stride])
			}
			s[7] = fftAdd(s[1], s[4])
			s[10] = fftSub(s[1], s[4])
			s[8] = fftAdd(s[2], s[3])
			s[9] = fftSub(s[2], s[3])
			f[u] = fftAdd(f[u], fftAdd(s[7], s[8]))
			s[5].Fr = s[0].Fr + (float32(s[7].Fr*ya.Fr) + float32(s[8].Fr*yb.Fr))
			s[5].Fi = s[0].Fi + (float32(s[7].Fi*ya.Fr) + float32(s[8].Fi*yb.Fr))
			s[6].Fr = float32(s[10].Fi*ya.Fi) + float32(s[9].Fi*yb.Fi)
			s[6].Fi = -(float32(s[10].Fr*ya.Fi) + float32(s[9].Fr*yb.Fi))
			f[u+m] = fftSub(s[5], s[6])
			f[u+4*m] = fftAdd(s[5], s[6])
			s[11].Fr = s[0].Fr + (float32(s[7].Fr*yb.Fr) + float32(s[8].Fr*ya.Fr))
			s[11].Fi = s[0].Fi + (float32(s[7].Fi*yb.Fr) + float32(s[8].Fi*ya.Fr))
			s[12].Fr = float32(s[9].Fi*ya.Fi) - float32(s[10].Fi*yb.Fi)
			s[12].Fi = float32(s[10].Fr*yb.Fi) - float32(s[9].Fr*ya.Fi)
			f[u+2*m] = fftAdd(s[11], s[12])
			f[u+3*m] = fftSub(s[11], s[12])
		}
	}
}

func Opus_opus_fft_impl(tls *libc.TLS, state *OpusT_kiss_fft_state, twiddles *OpusT_kiss_twiddle_cpx, fout *OpusT_kiss_fft_cpx) {
	var L, i, m, m2, p, shift, v1 int32
	var fstride [8]int32
	_, _, _, _, _, _, _, _ = L, fstride, i, m, m2, p, shift, v1
	/* st->shift can be -1 */
	if state.Fshift > 0 {
		v1 = state.Fshift
	} else {
		v1 = 0
	}
	shift = v1
	fstride[0] = int32(1)
	L = 0
	for cond := true; cond; cond = m != int32(1) {
		p = int32(state.Ffactors[2*L])
		m = int32(state.Ffactors[2*L+1])
		fstride[L+int32(1)] = fstride[L] * p
		L = L + 1
	}
	m = int32(state.Ffactors[2*L-1])
	i = L - int32(1)
	for {
		if !(i >= 0) {
			break
		}
		if i != 0 {
			m2 = int32(state.Ffactors[2*i-1])
		} else {
			m2 = int32(1)
		}
		switch int32(state.Ffactors[2*i]) {
		case int32(2):
			kf_bfly2(tls, fout, m, fstride[i])
		case int32(4):
			kf_bfly4(tls, fout, uint64(uint32(fstride[i]<<shift)), twiddles, m, fstride[i], m2)
		case int32(3):
			kf_bfly3(tls, fout, uint64(uint32(fstride[i]<<shift)), twiddles, m, fstride[i], m2)
		case int32(5):
			kf_bfly5(tls, fout, uint64(uint32(fstride[i]<<shift)), twiddles, m, fstride[i], m2)
			break
		}
		m = m2
		i = i - 1
	}
}

func Opus_opus_fft_c(tls *libc.TLS, st *OpusT_kiss_fft_state, bitrev *int16, twiddles *OpusT_kiss_twiddle_cpx, fin, fout *OpusT_kiss_fft_cpx) {
	if fin == fout {
		opusCeltFatal(tls, opusDiagnosticString(3512), opusDiagnosticString(3493), 626)
	}
	scale := st.Fscale
	in := unsafe.Slice(fin, st.Fnfft)
	out := unsafe.Slice(fout, st.Fnfft)
	rev := unsafe.Slice(bitrev, st.Fnfft)
	// Forward iteration intentionally preserves C's behavior on partial overlap.
	for i := range in {
		x := in[i]
		out[rev[i]].Fr = float32(x.Fr * scale)
		out[rev[i]].Fi = float32(x.Fi * scale)
	}
	Opus_opus_fft_impl(tls, st, twiddles, fout)
}

func Opus_opus_ifft_c(tls *libc.TLS, st *OpusT_kiss_fft_state, bitrev *int16, twiddles *OpusT_kiss_twiddle_cpx, fin, fout *OpusT_kiss_fft_cpx) {
	if fin == fout {
		opusCeltFatal(tls, opusDiagnosticString(3512), opusDiagnosticString(3493), 641)
	}
	in := unsafe.Slice(fin, st.Fnfft)
	out := unsafe.Slice(fout, st.Fnfft)
	rev := unsafe.Slice(bitrev, st.Fnfft)
	for i := range in {
		out[rev[i]] = in[i]
	}
	// Preserve both conjugation passes, including the sign bit of zero.
	for i := range out {
		out[i].Fi = -out[i].Fi
	}
	Opus_opus_fft_impl(tls, st, twiddles, fout)
	for i := range out {
		out[i].Fi = -out[i].Fi
	}
}

const CELT_SIG_SCALE5 = 32768

var log2_x_norm_coeff6 = [8]float32{
	0: float32(1),
	1: float32(0.8888888955116272),
	2: float32(0.8),
	3: float32(0.7272727489471436),
	4: float32(0.6666666865348816),
	5: float32(0.6153846383094788),
	6: float32(0.5714285969734192),
	7: float32(0.5333333611488342),
}
var log2_y_norm_coeff6 = [8]float32{
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
//	/*Compute floor(sqrt(_val)) with exact arithmetic.
//	  _val must be greater than 0.
//	  This has been tested on all possible 32-bit inputs greater than 0.*/
func Opus_isqrt32(tls *libc.TLS, _val OpusT_opus_uint32) (r uint32) {
	var b, g uint32
	var bshift int32
	var t OpusT_opus_uint32
	_, _, _, _ = b, bshift, g, t
	/*Uses the second method from
	   http://www.azillionmonkeys.com/qed/sqroot.html
	  The main idea is to search for the largest binary digit b such that
	   (g+b)*(g+b) <= _val, and add it to the solution g.*/
	g = uint32(0)
	bshift = (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, _val) - int32(1)) >> int32(1)
	b = uint32(1) << bshift
	for cond := true; cond; cond = bshift >= 0 {
		t = (g<<int32(1) + b) << bshift
		if t <= _val {
			g = g + b
			_val = _val - t
		}
		b = b >> uint32(1)
		bshift = bshift - 1
	}
	return g
}

func Opus_celt_float2int16_c(tls *libc.TLS, in *float32, out *int16, cnt int32) {
	if cnt <= 0 {
		return
	}
	inS := unsafe.Slice(in, int(cnt))
	outS := unsafe.Slice(out, int(cnt))
	for i := 0; i < int(cnt); i++ {
		v := float32(inS[i] * float32(32768))
		if v < float32(-int32(32768)) {
			v = float32(-int32(32768))
		}
		if v > float32(int32(32767)) {
			v = float32(int32(32767))
		}
		outS[i] = int16(libc.Xlrintf(tls, v))
	}
}

func Opus_opus_limit2_checkwithin1_c(tls *libc.TLS, samples *float32, cnt int32) (r int32) {
	if cnt <= 0 {
		return 1
	}
	pcm := unsafe.Slice(samples, int(cnt))
	for i, value := range pcm {
		// Match C's FMAX/FMIN comparisons, preserving NaNs and signed zero.
		if value < -2 {
			value = -2
		}
		if value > 2 {
			value = 2
		}
		pcm[i] = value
	}
	/* C implementation can't provide quick hint. Assume it might exceed -1/+1. */
	return 0
}

const CELT_SIG_SCALE6 = "32768.f"
const EC_CODE_BITS = 32
const EC_SYM_BITS = 8
const _mfrngcode_H = 1

/* Copyright (c) 2001-2008 Timothy B. Terriberry
   Copyright (c) 2008-2009 Xiph.Org Foundation */
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

/* Copyright (c) 2001-2011 Timothy B. Terriberry
   Copyright (c) 2008-2009 Xiph.Org Foundation */
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

/* (C) COPYRIGHT 1994-2002 Xiph.Org Foundation */
/* Modified by Jean-Marc Valin */
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
/* opus_types.h based on ogg_types.h from libogg */

/**
  @file opus_types.h
  @brief Opus reference implementation types
*/
/* Copyright (c) 2010-2011 Xiph.Org Foundation, Skype Limited
   Written by Jean-Marc Valin and Koen Vos */
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

/**
 * @file opus_defines.h
 * @brief Opus reference implementation constants
 */

/*Constants used by the entropy encoder/decoder.*/

/*The number of bits to output at a time.*/
/*The total number of bits in each of the state registers.*/
/*The maximum symbol value.*/
/*Bits to shift by to move a symbol into the high-order position.*/
/*Carry bit of the high-order range symbol.*/
/*Low-order bit of the high-order range symbol.*/
/*The number of bits available for the last, partial symbol in the code field.*/

/*A range decoder.
  This is an entropy decoder based upon \cite{Mar79}, which is itself a
   rediscovery of the FIFO arithmetic code introduced by \cite{Pas76}.
  It is very similar to arithmetic encoding, except that encoding is done with
   digits in any base, instead of with bits, and so it is faster when using
   larger bases (i.e.: a byte).
  The author claims an average waste of $\frac{1}{2}\log_b(2b)$ bits, where $b$
   is the base, longer than the theoretical optimum, but to my knowledge there
   is no published justification for this claim.
  This only seems true when using near-infinite precision arithmetic so that
   the process is carried out with no rounding errors.

  An excellent description of implementation details is available at
   http://www.arturocampos.com/ac_range.html
  A recent work \cite{MNW98} which proposes several changes to arithmetic
   encoding for efficiency actually re-discovers many of the principles
   behind range encoding, and presents a good theoretical analysis of them.

  End of stream is handled by writing out the smallest number of bits that
   ensures that the stream will be correctly decoded regardless of the value of
   any subsequent bits.
  ec_tell() can be used to determine how many bits were needed to decode
   all the symbols thus far; other data can be packed in the remaining bits of
   the input buffer.
  @PHDTHESIS{Pas76,
    author="Richard Clark Pasco",
    title="Source coding algorithms for fast data compression",
    school="Dept. of Electrical Engineering, Stanford University",
    address="Stanford, CA",
    month=May,
    year=1976
  }
  @INPROCEEDINGS{Mar79,
   author="Martin, G.N.N.",
   title="Range encoding: an algorithm for removing redundancy from a digitised
    message",
   booktitle="Video & Data Recording Conference",
   year=1979,
   address="Southampton",
   month=Jul
  }
  @ARTICLE{MNW98,
   author="Alistair Moffat and Radford Neal and Ian H. Witten",
   title="Arithmetic Coding Revisited",
   journal="{ACM} Transactions on Information Systems",
   year=1998,
   volume=16,
   number=3,
   pages="256--294",
   month=Jul,
   URL="http://www.stanford.edu/class/ee398a/handouts/papers/Moffat98ArithmCoding.pdf"
  }*/
