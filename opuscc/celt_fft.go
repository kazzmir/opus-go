// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

package opuscc

import (
	"reflect"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

var _ reflect.Type
var _ unsafe.Pointer

func kf_bfly2(tls *libc.TLS, Fout uintptr, m int32, N int32) {
	var Fout2 uintptr
	var i int32
	var t OpusT_kiss_fft_cpx
	var tw OpusT_celt_coef
	_, _, _, _ = Fout2, i, t, tw
	_ = m
	tw = float32(0.7071067812)
	/* We know that m==4 here because the radix-2 is just after a radix-4 */
	if !(m == int32(4)) {
		Opus_celt_fatal(tls, __ccgo_ts+3470, __ccgo_ts+3493, int32(80))
	}
	i = 0
	for {
		if !(i < N) {
			break
		}
		Fout2 = Fout + uintptr(4)*8
		t = *(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2))
		(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2))).Fr = (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout))).Fr - t.Fr
		(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2))).Fi = (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout))).Fi - t.Fi
		(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout))).Fr += t.Fr
		(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout))).Fi += t.Fi
		t.Fr = float32(((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2 + 1*8))).Fr + (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2 + 1*8))).Fi) * tw)
		t.Fi = float32(((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2 + 1*8))).Fi - (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2 + 1*8))).Fr) * tw)
		(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2 + 1*8))).Fr = (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 1*8))).Fr - t.Fr
		(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2 + 1*8))).Fi = (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 1*8))).Fi - t.Fi
		(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 1*8))).Fr += t.Fr
		(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 1*8))).Fi += t.Fi
		t.Fr = (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2 + 2*8))).Fi
		t.Fi = -(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2 + 2*8))).Fr
		(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2 + 2*8))).Fr = (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 2*8))).Fr - t.Fr
		(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2 + 2*8))).Fi = (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 2*8))).Fi - t.Fi
		(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 2*8))).Fr += t.Fr
		(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 2*8))).Fi += t.Fi
		t.Fr = float32(((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2 + 3*8))).Fi - (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2 + 3*8))).Fr) * tw)
		t.Fi = float32(-((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2 + 3*8))).Fi + (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2 + 3*8))).Fr) * tw)
		(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2 + 3*8))).Fr = (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 3*8))).Fr - t.Fr
		(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2 + 3*8))).Fi = (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 3*8))).Fi - t.Fi
		(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 3*8))).Fr += t.Fr
		(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 3*8))).Fi += t.Fi
		Fout = Fout + uintptr(8)*8
		i = i + 1
	}
}

func kf_bfly4(tls *libc.TLS, Fout uintptr, fstride OpusT_size_t, st uintptr, m int32, N int32, mm int32) {
	var Fout_beg, tw1, tw2, tw3, v3, v4 uintptr
	var i, j, m2, m3 int32
	var scratch [6]OpusT_kiss_fft_cpx
	var scratch0, scratch1 OpusT_kiss_fft_cpx
	_, _, _, _, _, _, _, _, _, _, _, _, _ = Fout_beg, i, j, m2, m3, scratch, scratch0, scratch1, tw1, tw2, tw3, v3, v4
	if m == int32(1) {
		/* Degenerate case where all the twiddles are 1. */
		i = 0
		for {
			if !(i < N) {
				break
			}
			scratch0.Fr = (*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fr - (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 2*8))).Fr
			scratch0.Fi = (*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fi - (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 2*8))).Fi
			(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fr += (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 2*8))).Fr
			(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fi += (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 2*8))).Fi
			scratch1.Fr = (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 1*8))).Fr + (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 3*8))).Fr
			scratch1.Fi = (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 1*8))).Fi + (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 3*8))).Fi
			(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 2*8))).Fr = (*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fr - scratch1.Fr
			(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 2*8))).Fi = (*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fi - scratch1.Fi
			(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fr += scratch1.Fr
			(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fi += scratch1.Fi
			scratch1.Fr = (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 1*8))).Fr - (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 3*8))).Fr
			scratch1.Fi = (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 1*8))).Fi - (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 3*8))).Fi
			(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 1*8))).Fr = scratch0.Fr + scratch1.Fi
			(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 1*8))).Fi = scratch0.Fi - scratch1.Fr
			(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 3*8))).Fr = scratch0.Fr - scratch1.Fi
			(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + 3*8))).Fi = scratch0.Fi + scratch1.Fr
			Fout = Fout + uintptr(4)*8
			i = i + 1
		}
	} else {
		m2 = int32(2) * m
		m3 = int32(3) * m
		Fout_beg = Fout
		i = 0
		for {
			if !(i < N) {
				break
			}
			Fout = Fout_beg + uintptr(i*mm)*8
			v4 = (*OpusT_kiss_fft_state)(unsafe.Pointer(st)).Ftwiddles
			tw1 = v4
			v3 = v4
			tw2 = v3
			tw3 = v3
			/* m is guaranteed to be a multiple of 4. */
			j = 0
			for {
				if !(j < m) {
					break
				}
				scratch[0].Fr = float32((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fr*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw1)).Fr) - float32((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fi*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw1)).Fi)
				scratch[0].Fi = float32((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fr*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw1)).Fi) + float32((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fi*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw1)).Fr)
				scratch[int32(1)].Fr = float32((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fr*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw2)).Fr) - float32((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fi*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw2)).Fi)
				scratch[int32(1)].Fi = float32((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fr*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw2)).Fi) + float32((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fi*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw2)).Fr)
				scratch[int32(2)].Fr = float32((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m3)*8))).Fr*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw3)).Fr) - float32((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m3)*8))).Fi*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw3)).Fi)
				scratch[int32(2)].Fi = float32((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m3)*8))).Fr*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw3)).Fi) + float32((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m3)*8))).Fi*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw3)).Fr)
				scratch[int32(5)].Fr = (*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fr - scratch[int32(1)].Fr
				scratch[int32(5)].Fi = (*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fi - scratch[int32(1)].Fi
				(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fr += scratch[int32(1)].Fr
				(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fi += scratch[int32(1)].Fi
				scratch[int32(3)].Fr = scratch[0].Fr + scratch[int32(2)].Fr
				scratch[int32(3)].Fi = scratch[0].Fi + scratch[int32(2)].Fi
				scratch[int32(4)].Fr = scratch[0].Fr - scratch[int32(2)].Fr
				scratch[int32(4)].Fi = scratch[0].Fi - scratch[int32(2)].Fi
				(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fr = (*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fr - scratch[int32(3)].Fr
				(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fi = (*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fi - scratch[int32(3)].Fi
				tw1 = tw1 + uintptr(fstride)*8
				tw2 = tw2 + uintptr(fstride*uint64(2))*8
				tw3 = tw3 + uintptr(fstride*uint64(3))*8
				(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fr += scratch[int32(3)].Fr
				(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fi += scratch[int32(3)].Fi
				(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fr = scratch[int32(5)].Fr + scratch[int32(4)].Fi
				(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fi = scratch[int32(5)].Fi - scratch[int32(4)].Fr
				(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m3)*8))).Fr = scratch[int32(5)].Fr - scratch[int32(4)].Fi
				(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m3)*8))).Fi = scratch[int32(5)].Fi + scratch[int32(4)].Fr
				Fout += 8
				j = j + 1
			}
			i = i + 1
		}
	}
}

func kf_bfly3(tls *libc.TLS, Fout uintptr, fstride OpusT_size_t, st uintptr, m int32, N int32, mm int32) {
	var Fout_beg, tw1, tw2, v2 uintptr
	var epi3 OpusT_kiss_twiddle_cpx
	var i int32
	var k, m2, v3 OpusT_size_t
	var scratch [5]OpusT_kiss_fft_cpx
	_, _, _, _, _, _, _, _, _, _ = Fout_beg, epi3, i, k, m2, scratch, tw1, tw2, v2, v3
	m2 = uint64(uint32(int32(2) * m))
	Fout_beg = Fout
	epi3 = *(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer((*OpusT_kiss_fft_state)(unsafe.Pointer(st)).Ftwiddles + uintptr(fstride*uint64(uint32(m)))*8))
	i = 0
	for {
		if !(i < N) {
			break
		}
		Fout = Fout_beg + uintptr(i*mm)*8
		v2 = (*OpusT_kiss_fft_state)(unsafe.Pointer(st)).Ftwiddles
		tw2 = v2
		tw1 = v2
		/* For non-custom modes, m is guaranteed to be a multiple of 4. */
		k = uint64(uint32(m))
		for {
			scratch[int32(1)].Fr = float32((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fr*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw1)).Fr) - float32((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fi*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw1)).Fi)
			scratch[int32(1)].Fi = float32((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fr*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw1)).Fi) + float32((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fi*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw1)).Fr)
			scratch[int32(2)].Fr = float32((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fr*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw2)).Fr) - float32((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fi*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw2)).Fi)
			scratch[int32(2)].Fi = float32((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fr*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw2)).Fi) + float32((*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fi*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw2)).Fr)
			scratch[int32(3)].Fr = scratch[int32(1)].Fr + scratch[int32(2)].Fr
			scratch[int32(3)].Fi = scratch[int32(1)].Fi + scratch[int32(2)].Fi
			scratch[0].Fr = scratch[int32(1)].Fr - scratch[int32(2)].Fr
			scratch[0].Fi = scratch[int32(1)].Fi - scratch[int32(2)].Fi
			tw1 = tw1 + uintptr(fstride)*8
			tw2 = tw2 + uintptr(fstride*uint64(2))*8
			(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fr = (*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fr - float32(scratch[int32(3)].Fr*float32(0.5))
			(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fi = (*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fi - float32(scratch[int32(3)].Fi*float32(0.5))
			scratch[0].Fr *= epi3.Fi
			scratch[0].Fi *= epi3.Fi
			(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fr += scratch[int32(3)].Fr
			(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout)).Fi += scratch[int32(3)].Fi
			(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fr = (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fr + scratch[0].Fi
			(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m2)*8))).Fi = (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fi - scratch[0].Fr
			(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fr = (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fr - scratch[0].Fi
			(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fi = (*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout + uintptr(m)*8))).Fi + scratch[0].Fr
			Fout += 8
			k = k - 1
			v3 = k
			if !(v3 != 0) {
				break
			}
		}
		i = i + 1
	}
}

func kf_bfly5(tls *libc.TLS, Fout uintptr, fstride OpusT_size_t, st uintptr, m int32, N int32, mm int32) {
	var Fout0, Fout1, Fout2, Fout3, Fout4, Fout_beg, tw uintptr
	var i, u int32
	var scratch [13]OpusT_kiss_fft_cpx
	var ya, yb OpusT_kiss_twiddle_cpx
	_, _, _, _, _, _, _, _, _, _, _, _ = Fout0, Fout1, Fout2, Fout3, Fout4, Fout_beg, i, scratch, tw, u, ya, yb
	Fout_beg = Fout
	ya = *(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer((*OpusT_kiss_fft_state)(unsafe.Pointer(st)).Ftwiddles + uintptr(fstride*uint64(uint32(m)))*8))
	yb = *(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer((*OpusT_kiss_fft_state)(unsafe.Pointer(st)).Ftwiddles + uintptr(fstride*uint64(2)*uint64(uint32(m)))*8))
	tw = (*OpusT_kiss_fft_state)(unsafe.Pointer(st)).Ftwiddles
	i = 0
	for {
		if !(i < N) {
			break
		}
		Fout = Fout_beg + uintptr(i*mm)*8
		Fout0 = Fout
		Fout1 = Fout0 + uintptr(m)*8
		Fout2 = Fout0 + uintptr(int32(2)*m)*8
		Fout3 = Fout0 + uintptr(int32(3)*m)*8
		Fout4 = Fout0 + uintptr(int32(4)*m)*8
		/* For non-custom modes, m is guaranteed to be a multiple of 4. */
		u = 0
		for {
			if !(u < m) {
				break
			}
			scratch[0] = *(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout0))
			scratch[int32(1)].Fr = float32((*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout1)).Fr*(*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(u))*fstride)*8))).Fr) - float32((*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout1)).Fi*(*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(u))*fstride)*8))).Fi)
			scratch[int32(1)].Fi = float32((*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout1)).Fr*(*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(u))*fstride)*8))).Fi) + float32((*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout1)).Fi*(*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(u))*fstride)*8))).Fr)
			scratch[int32(2)].Fr = float32((*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2)).Fr*(*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(2)*u))*fstride)*8))).Fr) - float32((*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2)).Fi*(*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(2)*u))*fstride)*8))).Fi)
			scratch[int32(2)].Fi = float32((*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2)).Fr*(*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(2)*u))*fstride)*8))).Fi) + float32((*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2)).Fi*(*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(2)*u))*fstride)*8))).Fr)
			scratch[int32(3)].Fr = float32((*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout3)).Fr*(*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(3)*u))*fstride)*8))).Fr) - float32((*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout3)).Fi*(*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(3)*u))*fstride)*8))).Fi)
			scratch[int32(3)].Fi = float32((*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout3)).Fr*(*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(3)*u))*fstride)*8))).Fi) + float32((*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout3)).Fi*(*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(3)*u))*fstride)*8))).Fr)
			scratch[int32(4)].Fr = float32((*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout4)).Fr*(*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(4)*u))*fstride)*8))).Fr) - float32((*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout4)).Fi*(*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(4)*u))*fstride)*8))).Fi)
			scratch[int32(4)].Fi = float32((*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout4)).Fr*(*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(4)*u))*fstride)*8))).Fi) + float32((*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout4)).Fi*(*(*OpusT_kiss_twiddle_cpx)(unsafe.Pointer(tw + uintptr(uint64(uint32(int32(4)*u))*fstride)*8))).Fr)
			scratch[int32(7)].Fr = scratch[int32(1)].Fr + scratch[int32(4)].Fr
			scratch[int32(7)].Fi = scratch[int32(1)].Fi + scratch[int32(4)].Fi
			scratch[int32(10)].Fr = scratch[int32(1)].Fr - scratch[int32(4)].Fr
			scratch[int32(10)].Fi = scratch[int32(1)].Fi - scratch[int32(4)].Fi
			scratch[int32(8)].Fr = scratch[int32(2)].Fr + scratch[int32(3)].Fr
			scratch[int32(8)].Fi = scratch[int32(2)].Fi + scratch[int32(3)].Fi
			scratch[int32(9)].Fr = scratch[int32(2)].Fr - scratch[int32(3)].Fr
			scratch[int32(9)].Fi = scratch[int32(2)].Fi - scratch[int32(3)].Fi
			(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout0)).Fr = (*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout0)).Fr + (scratch[int32(7)].Fr + scratch[int32(8)].Fr)
			(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout0)).Fi = (*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout0)).Fi + (scratch[int32(7)].Fi + scratch[int32(8)].Fi)
			scratch[int32(5)].Fr = scratch[0].Fr + (float32(scratch[int32(7)].Fr*ya.Fr) + float32(scratch[int32(8)].Fr*yb.Fr))
			scratch[int32(5)].Fi = scratch[0].Fi + (float32(scratch[int32(7)].Fi*ya.Fr) + float32(scratch[int32(8)].Fi*yb.Fr))
			scratch[int32(6)].Fr = float32(scratch[int32(10)].Fi*ya.Fi) + float32(scratch[int32(9)].Fi*yb.Fi)
			scratch[int32(6)].Fi = -(float32(scratch[int32(10)].Fr*ya.Fi) + float32(scratch[int32(9)].Fr*yb.Fi))
			(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout1)).Fr = scratch[int32(5)].Fr - scratch[int32(6)].Fr
			(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout1)).Fi = scratch[int32(5)].Fi - scratch[int32(6)].Fi
			(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout4)).Fr = scratch[int32(5)].Fr + scratch[int32(6)].Fr
			(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout4)).Fi = scratch[int32(5)].Fi + scratch[int32(6)].Fi
			scratch[int32(11)].Fr = scratch[0].Fr + (float32(scratch[int32(7)].Fr*yb.Fr) + float32(scratch[int32(8)].Fr*ya.Fr))
			scratch[int32(11)].Fi = scratch[0].Fi + (float32(scratch[int32(7)].Fi*yb.Fr) + float32(scratch[int32(8)].Fi*ya.Fr))
			scratch[int32(12)].Fr = float32(scratch[int32(9)].Fi*ya.Fi) - float32(scratch[int32(10)].Fi*yb.Fi)
			scratch[int32(12)].Fi = float32(scratch[int32(10)].Fr*yb.Fi) - float32(scratch[int32(9)].Fr*ya.Fi)
			(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2)).Fr = scratch[int32(11)].Fr + scratch[int32(12)].Fr
			(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout2)).Fi = scratch[int32(11)].Fi + scratch[int32(12)].Fi
			(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout3)).Fr = scratch[int32(11)].Fr - scratch[int32(12)].Fr
			(*OpusT_kiss_fft_cpx)(unsafe.Pointer(Fout3)).Fi = scratch[int32(11)].Fi - scratch[int32(12)].Fi
			Fout0 += 8
			Fout1 += 8
			Fout2 += 8
			Fout3 += 8
			Fout4 += 8
			u = u + 1
		}
		i = i + 1
	}
}

func Opus_opus_fft_impl(tls *libc.TLS, st uintptr, fout uintptr) {
	var L, i, m, m2, p, shift, v1 int32
	var fstride [8]int32
	_, _, _, _, _, _, _, _ = L, fstride, i, m, m2, p, shift, v1
	/* st->shift can be -1 */
	if (*OpusT_kiss_fft_state)(unsafe.Pointer(st)).Fshift > 0 {
		v1 = (*OpusT_kiss_fft_state)(unsafe.Pointer(st)).Fshift
	} else {
		v1 = 0
	}
	shift = v1
	fstride[0] = int32(1)
	L = 0
	for cond := true; cond; cond = m != int32(1) {
		p = int32(*(*OpusT_opus_int16)(unsafe.Pointer(st + 12 + uintptr(int32(2)*L)*2)))
		m = int32(*(*OpusT_opus_int16)(unsafe.Pointer(st + 12 + uintptr(int32(2)*L+int32(1))*2)))
		fstride[L+int32(1)] = fstride[L] * p
		L = L + 1
	}
	m = int32(*(*OpusT_opus_int16)(unsafe.Pointer(st + 12 + uintptr(int32(2)*L-int32(1))*2)))
	i = L - int32(1)
	for {
		if !(i >= 0) {
			break
		}
		if i != 0 {
			m2 = int32(*(*OpusT_opus_int16)(unsafe.Pointer(st + 12 + uintptr(int32(2)*i-int32(1))*2)))
		} else {
			m2 = int32(1)
		}
		switch int32(*(*OpusT_opus_int16)(unsafe.Pointer(st + 12 + uintptr(int32(2)*i)*2))) {
		case int32(2):
			kf_bfly2(tls, fout, m, fstride[i])
		case int32(4):
			kf_bfly4(tls, fout, uint64(uint32(fstride[i]<<shift)), st, m, fstride[i], m2)
		case int32(3):
			kf_bfly3(tls, fout, uint64(uint32(fstride[i]<<shift)), st, m, fstride[i], m2)
		case int32(5):
			kf_bfly5(tls, fout, uint64(uint32(fstride[i]<<shift)), st, m, fstride[i], m2)
			break
		}
		m = m2
		i = i - 1
	}
}

func Opus_opus_fft_c(tls *libc.TLS, st uintptr, fin uintptr, fout uintptr) {
	var i int32
	var scale OpusT_celt_coef
	var x OpusT_kiss_fft_cpx
	_, _, _ = i, scale, x
	scale = (*OpusT_kiss_fft_state)(unsafe.Pointer(st)).Fscale
	if !(fin != fout) {
		Opus_celt_fatal(tls, __ccgo_ts+3512, __ccgo_ts+3493, int32(626))
	}
	/* Bit-reverse the input */
	i = 0
	for {
		if !(i < (*OpusT_kiss_fft_state)(unsafe.Pointer(st)).Fnfft) {
			break
		}
		x = *(*OpusT_kiss_fft_cpx)(unsafe.Pointer(fin + uintptr(i)*8))
		(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(fout + uintptr(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_kiss_fft_state)(unsafe.Pointer(st)).Fbitrev + uintptr(i)*2)))*8))).Fr = float32(x.Fr * scale)
		(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(fout + uintptr(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_kiss_fft_state)(unsafe.Pointer(st)).Fbitrev + uintptr(i)*2)))*8))).Fi = float32(x.Fi * scale)
		i = i + 1
	}
	Opus_opus_fft_impl(tls, st, fout)
}

func Opus_opus_ifft_c(tls *libc.TLS, st uintptr, fin uintptr, fout uintptr) {
	var i int32
	_ = i
	if !(fin != fout) {
		Opus_celt_fatal(tls, __ccgo_ts+3512, __ccgo_ts+3493, int32(641))
	}
	/* Bit-reverse the input */
	i = 0
	for {
		if !(i < (*OpusT_kiss_fft_state)(unsafe.Pointer(st)).Fnfft) {
			break
		}
		*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(fout + uintptr(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_kiss_fft_state)(unsafe.Pointer(st)).Fbitrev + uintptr(i)*2)))*8)) = *(*OpusT_kiss_fft_cpx)(unsafe.Pointer(fin + uintptr(i)*8))
		i = i + 1
	}
	i = 0
	for {
		if !(i < (*OpusT_kiss_fft_state)(unsafe.Pointer(st)).Fnfft) {
			break
		}
		(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(fout + uintptr(i)*8))).Fi = -(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(fout + uintptr(i)*8))).Fi
		i = i + 1
	}
	Opus_opus_fft_impl(tls, st, fout)
	i = 0
	for {
		if !(i < (*OpusT_kiss_fft_state)(unsafe.Pointer(st)).Fnfft) {
			break
		}
		(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(fout + uintptr(i)*8))).Fi = -(*(*OpusT_kiss_fft_cpx)(unsafe.Pointer(fout + uintptr(i)*8))).Fi
		i = i + 1
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

func Opus_celt_float2int16_c(tls *libc.TLS, in uintptr, out uintptr, cnt int32) {
	if cnt <= 0 {
		return
	}
	inS := unsafe.Slice((*float32)(unsafe.Pointer(in)), int(cnt))
	outS := unsafe.Slice((*int16)(unsafe.Pointer(out)), int(cnt))
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

func Opus_opus_limit2_checkwithin1_c(tls *libc.TLS, samples uintptr, cnt int32) (r int32) {
	var clippedVal, v2 float32
	var i int32
	_, _, _ = clippedVal, i, v2
	if cnt <= 0 {
		return int32(1)
	}
	i = 0
	for {
		if !(i < cnt) {
			break
		}
		clippedVal = *(*float32)(unsafe.Pointer(samples + uintptr(i)*4))
		if -float32(2) > clippedVal {
			v2 = -float32(2)
		} else {
			v2 = clippedVal
		}
		clippedVal = v2
		if float32(2) < clippedVal {
			v2 = float32(2)
		} else {
			v2 = clippedVal
		}
		clippedVal = v2
		*(*float32)(unsafe.Pointer(samples + uintptr(i)*4)) = clippedVal
		i = i + 1
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
