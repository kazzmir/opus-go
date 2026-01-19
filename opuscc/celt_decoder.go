// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

package opuscc

import (
	"reflect"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

var _ reflect.Type
var _ unsafe.Pointer

func Opus_validate_celt_decoder(tls *libc.TLS, st uintptr) {
	mode, _ := Opus_opus_custom_mode_create(tls, int32(48000), int32(960))
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fmode == mode) {
		Opus_celt_fatal(tls, __ccgo_ts+3695, __ccgo_ts+3767, int32(147))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Foverlap == int32(120)) {
		Opus_celt_fatal(tls, __ccgo_ts+3790, __ccgo_ts+3767, int32(148))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fend <= int32(21)) {
		Opus_celt_fatal(tls, __ccgo_ts+3827, __ccgo_ts+3767, int32(149))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fchannels == int32(1) || (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fchannels == int32(2)) {
		Opus_celt_fatal(tls, __ccgo_ts, __ccgo_ts+3767, int32(157))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fstream_channels == int32(1) || (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fstream_channels == int32(2)) {
		Opus_celt_fatal(tls, __ccgo_ts+925, __ccgo_ts+3767, int32(158))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fdownsample > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+3859, __ccgo_ts+3767, int32(159))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fstart == 0 || (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fstart == int32(17)) {
		Opus_celt_fatal(tls, __ccgo_ts+3896, __ccgo_ts+3767, int32(160))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fstart < (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fend) {
		Opus_celt_fatal(tls, __ccgo_ts+3948, __ccgo_ts+3767, int32(161))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Farch >= int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+849, __ccgo_ts+3767, int32(163))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Farch <= int32(OPUS_ARCHMASK)) {
		Opus_celt_fatal(tls, __ccgo_ts+881, __ccgo_ts+3767, int32(164))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Flast_pitch_index <= int32(PLC_PITCH_LAG_MAX)) {
		Opus_celt_fatal(tls, __ccgo_ts+3986, __ccgo_ts+3767, int32(167))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Flast_pitch_index >= int32(PLC_PITCH_LAG_MIN) || (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Flast_pitch_index == 0) {
		Opus_celt_fatal(tls, __ccgo_ts+4046, __ccgo_ts+3767, int32(168))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fpostfilter_period < int32(MAX_PERIOD)) {
		Opus_celt_fatal(tls, __ccgo_ts+4135, __ccgo_ts+3767, int32(170))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fpostfilter_period >= int32(COMBFILTER_MINPERIOD) || (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fpostfilter_period == 0) {
		Opus_celt_fatal(tls, __ccgo_ts+4188, __ccgo_ts+3767, int32(171))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fpostfilter_period_old < int32(MAX_PERIOD)) {
		Opus_celt_fatal(tls, __ccgo_ts+4282, __ccgo_ts+3767, int32(172))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fpostfilter_period_old >= int32(COMBFILTER_MINPERIOD) || (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fpostfilter_period_old == 0) {
		Opus_celt_fatal(tls, __ccgo_ts+4339, __ccgo_ts+3767, int32(173))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fpostfilter_tapset <= int32(2)) {
		Opus_celt_fatal(tls, __ccgo_ts+4441, __ccgo_ts+3767, int32(174))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fpostfilter_tapset >= int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+4486, __ccgo_ts+3767, int32(175))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fpostfilter_tapset_old <= int32(2)) {
		Opus_celt_fatal(tls, __ccgo_ts+4531, __ccgo_ts+3767, int32(176))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fpostfilter_tapset_old >= int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+4580, __ccgo_ts+3767, int32(177))
	}
}

func Opus_celt_decoder_get_size(tls *libc.TLS, channels int32) (r int32) {
	var mode uintptr
	_ = mode
	mode, _ = Opus_opus_custom_mode_create(tls, int32(48000), int32(960))
	return opus_custom_decoder_get_size(tls, mode, channels)
}

func opus_custom_decoder_get_size(tls *libc.TLS, mode uintptr, channels int32) (r int32) {
	var size int32
	_ = size
	size = int32(uint64(120) + uint64(uint32(channels*(int32(DEC_PITCH_BUF_SIZE)+(*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Foverlap)-int32(1)))*uint64(4) + uint64(uint32(int32(4)*int32(2)*(*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FnbEBands))*uint64(4) + uint64(uint32(channels*int32(CELT_LPC_ORDER)))*uint64(4))
	return size
}

func Opus_celt_decoder_init(tls *libc.TLS, st uintptr, sampling_rate OpusT_opus_int32, channels int32) (r int32) {
	var ret int32
	_ = ret
	mode, _ := Opus_opus_custom_mode_create(tls, int32(48000), int32(960))
	ret = opus_custom_decoder_init(tls, st, mode, channels)
	if ret != OPUS_OK {
		return ret
	}
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fdownsample = Opus_resampling_factor(tls, sampling_rate)
	if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fdownsample == 0 {
		return -int32(1)
	} else {
		return OPUS_OK
	}
	return r
}

func opus_custom_decoder_init(tls *libc.TLS, st uintptr, mode uintptr, channels int32) (r int32) {
	var v1 int32
	_ = v1
	if channels < 0 || channels > int32(2) {
		return -int32(1)
	}
	if st == uintptr(uint32(0)) {
		return -int32(7)
	}
	libc.Xmemset(tls, st, 0, uint64(uint32(opus_custom_decoder_get_size(tls, mode, channels)))*uint64(1))
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fmode = mode
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Foverlap = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Foverlap
	v1 = channels
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fchannels = v1
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fstream_channels = v1
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fdownsample = int32(1)
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fstart = 0
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fend = (*OpusT_OpusCustomMode)(unsafe.Pointer((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fmode)).FeffEBands
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fsignalling = int32(1)
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fdisable_inv = libc.BoolInt32(channels == int32(1))
	v1 = 0
	goto _3
_3:
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Farch = v1
	Opus_opus_custom_decoder_ctl(tls, st, int32(OPUS_RESET_STATE), 0)
	return OPUS_OK
}

// C documentation
//
//	/* Special case for stereo with no downsampling and no accumulation. This is
//	   quite common and we can make it faster by processing both channels in the
//	   same loop, reducing overhead due to the dependency loop in the IIR filter. */
func deemphasis_stereo_simple(tls *libc.TLS, in uintptr, pcm uintptr, N int32, coef0 OpusT_opus_val16, mem uintptr) {
	var j int32
	var m0, m1, tmp0, tmp1 OpusT_celt_sig
	var x0, x1 uintptr
	_, _, _, _, _, _, _ = j, m0, m1, tmp0, tmp1, x0, x1
	x0 = *(*uintptr)(unsafe.Pointer(in))
	x1 = *(*uintptr)(unsafe.Pointer(in + 1*8))
	m0 = *(*OpusT_celt_sig)(unsafe.Pointer(mem))
	m1 = *(*OpusT_celt_sig)(unsafe.Pointer(mem + 1*4))
	j = 0
	for {
		if !(j < N) {
			break
		}
		/* Add VERY_SMALL to x[] first to reduce dependency chain. */
		tmp0 = *(*OpusT_celt_sig)(unsafe.Pointer(x0 + uintptr(j)*4)) + float32(1e-30) + m0
		tmp1 = *(*OpusT_celt_sig)(unsafe.Pointer(x1 + uintptr(j)*4)) + float32(1e-30) + m1
		m0 = OpusT_opus_val16(coef0 * tmp0)
		m1 = OpusT_opus_val16(coef0 * tmp1)
		*(*OpusT_opus_res)(unsafe.Pointer(pcm + uintptr(int32(2)*j)*4)) = float32(float32(1) / float32(32768) * tmp0)
		*(*OpusT_opus_res)(unsafe.Pointer(pcm + uintptr(int32(2)*j+int32(1))*4)) = float32(float32(1) / float32(32768) * tmp1)
		goto _1
	_1:
		;
		j = j + 1
	}
	*(*OpusT_celt_sig)(unsafe.Pointer(mem)) = m0
	*(*OpusT_celt_sig)(unsafe.Pointer(mem + 1*4)) = m1
}

func deemphasis(tls *libc.TLS, in uintptr, pcm uintptr, N int32, C int32, downsample int32, coef uintptr, mem uintptr, accum int32) {
	var Nd, apply_downsampling, c, j, v29 int32
	var _saved_stack, scratch, st, x, y, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var coef0 OpusT_opus_val16
	var m, tmp, tmp1, tmp2 OpusT_celt_sig
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = Nd, _saved_stack, apply_downsampling, c, coef0, j, m, scratch, st, tmp, tmp1, tmp2, x, y, v1, v11, v13, v15, v17, v19, v21, v23, v29, v3, v5, v7, v9
	apply_downsampling = 0
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
	/* Short version for common case. */
	if downsample == int32(1) && C == int32(2) && !(accum != 0) {
		deemphasis_stereo_simple(tls, in, pcm, N, *(*OpusT_opus_val16)(unsafe.Pointer(coef)), mem)
		return
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
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(335))
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
	scratch = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32(N))*(uint64(4)/uint64(1)))
	coef0 = *(*OpusT_opus_val16)(unsafe.Pointer(coef))
	Nd = N / downsample
	c = 0
	for {
		m = *(*OpusT_celt_sig)(unsafe.Pointer(mem + uintptr(c)*4))
		x = *(*uintptr)(unsafe.Pointer(in + uintptr(c)*8))
		y = pcm + uintptr(c)*4
		if downsample > int32(1) {
			/* Shortcut for the standard (non-custom modes) case */
			j = 0
			for {
				if !(j < N) {
					break
				}
				tmp = *(*OpusT_celt_sig)(unsafe.Pointer(x + uintptr(j)*4)) + float32(1e-30) + m
				m = OpusT_opus_val16(coef0 * tmp)
				*(*OpusT_celt_sig)(unsafe.Pointer(scratch + uintptr(j)*4)) = tmp
				goto _31
			_31:
				;
				j = j + 1
			}
			apply_downsampling = int32(1)
		} else {
			/* Shortcut for the standard (non-custom modes) case */
			if accum != 0 {
				j = 0
				for {
					if !(j < N) {
						break
					}
					tmp1 = *(*OpusT_celt_sig)(unsafe.Pointer(x + uintptr(j)*4)) + m + float32(1e-30)
					m = OpusT_opus_val16(coef0 * tmp1)
					*(*OpusT_opus_res)(unsafe.Pointer(y + uintptr(j*C)*4)) = *(*OpusT_opus_res)(unsafe.Pointer(y + uintptr(j*C)*4)) + float32(float32(1)/float32(32768)*tmp1)
					goto _32
				_32:
					;
					j = j + 1
				}
			} else {
				j = 0
				for {
					if !(j < N) {
						break
					}
					tmp2 = *(*OpusT_celt_sig)(unsafe.Pointer(x + uintptr(j)*4)) + float32(1e-30) + m
					m = OpusT_opus_val16(coef0 * tmp2)
					*(*OpusT_opus_res)(unsafe.Pointer(y + uintptr(j*C)*4)) = float32(float32(1) / float32(32768) * tmp2)
					goto _33
				_33:
					;
					j = j + 1
				}
			}
		}
		*(*OpusT_celt_sig)(unsafe.Pointer(mem + uintptr(c)*4)) = m
		if apply_downsampling != 0 {
			/* Perform down-sampling */
			if accum != 0 {
				j = 0
				for {
					if !(j < Nd) {
						break
					}
					*(*OpusT_opus_res)(unsafe.Pointer(y + uintptr(j*C)*4)) = *(*OpusT_opus_res)(unsafe.Pointer(y + uintptr(j*C)*4)) + float32(float32(1)/float32(32768)**(*OpusT_celt_sig)(unsafe.Pointer(scratch + uintptr(j*downsample)*4)))
					goto _34
				_34:
					;
					j = j + 1
				}
			} else {
				j = 0
				for {
					if !(j < Nd) {
						break
					}
					*(*OpusT_opus_res)(unsafe.Pointer(y + uintptr(j*C)*4)) = float32(float32(1) / float32(32768) * *(*OpusT_celt_sig)(unsafe.Pointer(scratch + uintptr(j*downsample)*4)))
					goto _35
				_35:
					;
					j = j + 1
				}
			}
		}
		goto _30
	_30:
		;
		c = c + 1
		v29 = c
		if !(v29 < C) {
			break
		}
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _37
	_37:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _39
_39:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
}

func celt_synthesis(tls *libc.TLS, mode uintptr, X uintptr, out_syn uintptr, oldBandE uintptr, start int32, effEnd int32, C int32, CC int32, isTransient int32, LM int32, downsample int32, silence int32, arch int32) {
	var B, M, N, NB, b, c, i, nbEBands, overlap, shift, v33 int32
	var _saved_stack, freq, freq2, freq21, st, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = B, M, N, NB, _saved_stack, b, c, freq, freq2, freq21, i, nbEBands, overlap, shift, st, v1, v11, v13, v15, v17, v19, v21, v23, v3, v33, v5, v7, v9
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
	overlap = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Foverlap
	nbEBands = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FnbEBands
	N = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize << LM
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
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(432))
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
	freq = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32(N))*(uint64(4)/uint64(1))) /**< Interleaved signal MDCTs */
	M = int32(1) << LM
	if isTransient != 0 {
		B = M
		NB = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize
		shift = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FmaxLM
	} else {
		B = int32(1)
		NB = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize << LM
		shift = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FmaxLM - LM
	}
	if CC == int32(2) && C == int32(1) {
		Opus_denormalise_bands(tls, mode, X, freq, oldBandE, start, effEnd, M, downsample, silence)
		/* Store a temporary copy in the output buffer because the IMDCT destroys its input. */
		freq2 = *(*uintptr)(unsafe.Pointer(out_syn + 1*8)) + uintptr(overlap/int32(2))*4
		libc.Xmemcpy(tls, freq2, freq, uint64(uint32(N))*uint64(4)+uint64(0*((int64(freq2)-int64(freq))/4)))
		b = 0
		for {
			if !(b < B) {
				break
			}
			Opus_clt_mdct_backward_c(tls, mode+80, freq2+uintptr(b)*4, *(*uintptr)(unsafe.Pointer(out_syn))+uintptr(NB*b)*4, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Fwindow, overlap, shift, B, arch)
			goto _29
		_29:
			;
			b = b + 1
		}
		b = 0
		for {
			if !(b < B) {
				break
			}
			Opus_clt_mdct_backward_c(tls, mode+80, freq+uintptr(b)*4, *(*uintptr)(unsafe.Pointer(out_syn + 1*8))+uintptr(NB*b)*4, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Fwindow, overlap, shift, B, arch)
			goto _30
		_30:
			;
			b = b + 1
		}
	} else {
		if CC == int32(1) && C == int32(2) {
			freq21 = *(*uintptr)(unsafe.Pointer(out_syn)) + uintptr(overlap/int32(2))*4
			Opus_denormalise_bands(tls, mode, X, freq, oldBandE, start, effEnd, M, downsample, silence)
			/* Use the output buffer as temp array before downmixing. */
			Opus_denormalise_bands(tls, mode, X+uintptr(N)*4, freq21, oldBandE+uintptr(nbEBands)*4, start, effEnd, M, downsample, silence)
			i = 0
			for {
				if !(i < N) {
					break
				}
				*(*OpusT_celt_sig)(unsafe.Pointer(freq + uintptr(i)*4)) = float32(float32(0.5)**(*OpusT_celt_sig)(unsafe.Pointer(freq + uintptr(i)*4))) + float32(float32(0.5)**(*OpusT_celt_sig)(unsafe.Pointer(freq21 + uintptr(i)*4)))
				goto _31
			_31:
				;
				i = i + 1
			}
			b = 0
			for {
				if !(b < B) {
					break
				}
				Opus_clt_mdct_backward_c(tls, mode+80, freq+uintptr(b)*4, *(*uintptr)(unsafe.Pointer(out_syn))+uintptr(NB*b)*4, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Fwindow, overlap, shift, B, arch)
				goto _32
			_32:
				;
				b = b + 1
			}
		} else {
			/* Normal case (mono or stereo) */
			c = 0
			for {
				Opus_denormalise_bands(tls, mode, X+uintptr(c*N)*4, freq, oldBandE+uintptr(c*nbEBands)*4, start, effEnd, M, downsample, silence)
				b = 0
				for {
					if !(b < B) {
						break
					}
					Opus_clt_mdct_backward_c(tls, mode+80, freq+uintptr(b)*4, *(*uintptr)(unsafe.Pointer(out_syn + uintptr(c)*8))+uintptr(NB*b)*4, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Fwindow, overlap, shift, B, arch)
					goto _35
				_35:
					;
					b = b + 1
				}
				goto _34
			_34:
				;
				c = c + 1
				v33 = c
				if !(v33 < CC) {
					break
				}
			}
		}
	}
	/* Saturate IMDCT output so that we can't overflow in the pitch postfilter
	   or in the */
	c = 0
	for {
		i = 0
		for {
			if !(i < N) {
				break
			}
			*(*OpusT_celt_sig)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(out_syn + uintptr(c)*8)) + uintptr(i)*4)) = *(*OpusT_celt_sig)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(out_syn + uintptr(c)*8)) + uintptr(i)*4))
			goto _38
		_38:
			;
			i = i + 1
		}
		goto _37
	_37:
		;
		c = c + 1
		v33 = c
		if !(v33 < CC) {
			break
		}
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _40
	_40:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _42
_42:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
}

func tf_decode(tls *libc.TLS, start int32, end int32, isTransient int32, tf_res uintptr, LM int32, dec uintptr) {
	var budget, tell OpusT_opus_uint32
	var curr, i, logp, tf_changed, tf_select, tf_select_rsv, v2 int32
	var v1 uintptr
	_, _, _, _, _, _, _, _, _, _ = budget, curr, i, logp, tell, tf_changed, tf_select, tf_select_rsv, v1, v2
	budget = (*OpusT_ec_dec)(unsafe.Pointer(dec)).Fstorage * uint32(8)
	v1 = dec
	v2 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
	goto _3
_3:
	tell = uint32(v2)
	if isTransient != 0 {
		v2 = int32(2)
	} else {
		v2 = int32(4)
	}
	logp = v2
	tf_select_rsv = libc.BoolInt32(LM > 0 && tell+uint32(logp)+uint32(1) <= budget)
	budget = budget - uint32(tf_select_rsv)
	v2 = int32(0)
	curr = v2
	tf_changed = v2
	i = start
	for {
		if !(i < end) {
			break
		}
		if tell+uint32(logp) <= budget {
			curr = curr ^ Opus_ec_dec_bit_logp(tls, dec, uint32(logp))
			v1 = dec
			v2 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
			goto _9
		_9:
			tell = uint32(v2)
			tf_changed = tf_changed | curr
		}
		*(*int32)(unsafe.Pointer(tf_res + uintptr(i)*4)) = curr
		if isTransient != 0 {
			v2 = int32(4)
		} else {
			v2 = int32(5)
		}
		logp = v2
		goto _6
	_6:
		;
		i = i + 1
	}
	tf_select = 0
	if tf_select_rsv != 0 && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_tf_select_table)) + uintptr(LM)*8 + uintptr(int32(4)*isTransient+0+tf_changed)))) != int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_tf_select_table)) + uintptr(LM)*8 + uintptr(int32(4)*isTransient+int32(2)+tf_changed)))) {
		tf_select = Opus_ec_dec_bit_logp(tls, dec, uint32(1))
	}
	i = start
	for {
		if !(i < end) {
			break
		}
		*(*int32)(unsafe.Pointer(tf_res + uintptr(i)*4)) = int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(&Opus_tf_select_table)) + uintptr(LM)*8 + uintptr(int32(4)*isTransient+int32(2)*tf_select+*(*int32)(unsafe.Pointer(tf_res + uintptr(i)*4))))))
		goto _11
	_11:
		;
		i = i + 1
	}
}

func celt_plc_pitch_search(tls *libc.TLS, st1 uintptr, decode_mem uintptr, C int32, arch int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _saved_stack, lp_pitch_buf, st, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var _ /* pitch_index at bp+0 */ int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = _saved_stack, lp_pitch_buf, st, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9
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
	_ = st1
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
	if !(int64(int32(uint64(uint32(int32(DEC_PITCH_BUF_SIZE)>>int32(1)))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(565))
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
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32(int32(DEC_PITCH_BUF_SIZE)>>int32(1))) * (uint64(4) / uint64(1)))
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
	lp_pitch_buf = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32(int32(DEC_PITCH_BUF_SIZE)>>int32(1)))*(uint64(4)/uint64(1)))
	Opus_pitch_downsample(tls, decode_mem, lp_pitch_buf, int32(DEC_PITCH_BUF_SIZE)>>int32(1), C, int32(2), arch)
	Opus_pitch_search(tls, lp_pitch_buf+uintptr(int32(PLC_PITCH_LAG_MAX)>>int32(1))*4, lp_pitch_buf, int32(DEC_PITCH_BUF_SIZE)-int32(PLC_PITCH_LAG_MAX), int32(PLC_PITCH_LAG_MAX)-int32(PLC_PITCH_LAG_MIN), bp, arch)
	*(*int32)(unsafe.Pointer(bp)) = int32(PLC_PITCH_LAG_MAX) - *(*int32)(unsafe.Pointer(bp))
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
	return *(*int32)(unsafe.Pointer(bp))
}

func prefilter_and_fold(tls *libc.TLS, st1 uintptr, N int32) {
	var CC, c, decode_buffer_size, i, overlap, v29 int32
	var _saved_stack, etmp, mode, st, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var decode_mem [2]uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = CC, _saved_stack, c, decode_buffer_size, decode_mem, etmp, i, mode, overlap, st, v1, v11, v13, v15, v17, v19, v21, v23, v29, v3, v5, v7, v9
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
	decode_buffer_size = int32(DEC_PITCH_BUF_SIZE)
	mode = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fmode
	overlap = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Foverlap
	CC = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fchannels
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
	if !(int64(int32(uint64(uint32(overlap))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(597))
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
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32(overlap)) * (uint64(4) / uint64(1)))
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
	etmp = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32(overlap))*(uint64(4)/uint64(1)))
	c = 0
	for {
		decode_mem[c] = st1 + 112 + uintptr(c*(decode_buffer_size+overlap))*4
		goto _30
	_30:
		;
		c = c + 1
		v29 = c
		if !(v29 < CC) {
			break
		}
	}
	c = 0
	for {
		/* Apply the pre-filter to the MDCT overlap for the next frame because
		   the post-filter will be re-applied in the decoder after the MDCT
		   overlap. */
		Opus_comb_filter(tls, etmp, decode_mem[c]+uintptr(decode_buffer_size)*4-uintptr(N)*4, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period, overlap, -(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain_old, -(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset_old, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset, uintptr(uint32(0)), 0, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
		/* Simulate TDAC on the concealed audio so that it blends with the
		   MDCT of the next frame. */
		i = 0
		for {
			if !(i < overlap/int32(2)) {
				break
			}
			*(*OpusT_celt_sig)(unsafe.Pointer(decode_mem[c] + uintptr(decode_buffer_size-N+i)*4)) = OpusT_celt_coef(*(*OpusT_celt_coef)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Fwindow + uintptr(i)*4))**(*OpusT_opus_val32)(unsafe.Pointer(etmp + uintptr(overlap-int32(1)-i)*4))) + OpusT_celt_coef(*(*OpusT_celt_coef)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Fwindow + uintptr(overlap-i-int32(1))*4))**(*OpusT_opus_val32)(unsafe.Pointer(etmp + uintptr(i)*4)))
			goto _33
		_33:
			;
			i = i + 1
		}
		goto _32
	_32:
		;
		c = c + 1
		v29 = c
		if !(v29 < CC) {
			break
		}
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _35
	_35:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _37
_37:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
}

func celt_decode_lost(tls *libc.TLS, st1 uintptr, N int32, LM int32) {
	bp := tls.Alloc(240)
	defer tls.Free(240)
	var C, blen, boffs, c, curr_frame_type, curr_neural, decay_length, decode_buffer_size, effEnd, end, exc_length, extrapolation_len, extrapolation_offset, i, j, j1, last_neural, loss_duration, max_period, nbEBands, overlap, pitch_index, start, v5, v7, v8 int32
	var E1, E2, S1, S2, v103 OpusT_opus_val32
	var X, _exc, _saved_stack, backgroundLogE, buf, eBands, exc, fir_tmp, lpc, mode, oldBandE, oldLogE, oldLogE2, st, window, v1, v10, v12, v14, v16, v18, v20, v22, v24, v26, v28, v3 uintptr
	var attenuation, decay1, e, fade, ratio, tmp, tmp1, tmp_g OpusT_opus_val16
	var decay, v40 OpusT_celt_glog
	var seed OpusT_opus_uint32
	var v36 float32
	var _ /* ac at bp+32 */ [25]OpusT_opus_val32
	var _ /* decode_mem at bp+0 */ [2]uintptr
	var _ /* lpc_mem at bp+132 */ [24]OpusT_opus_val16
	var _ /* out_syn at bp+16 */ [2]uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, E1, E2, S1, S2, X, _exc, _saved_stack, attenuation, backgroundLogE, blen, boffs, buf, c, curr_frame_type, curr_neural, decay, decay1, decay_length, decode_buffer_size, e, eBands, effEnd, end, exc, exc_length, extrapolation_len, extrapolation_offset, fade, fir_tmp, i, j, j1, last_neural, loss_duration, lpc, max_period, mode, nbEBands, oldBandE, oldLogE, oldLogE2, overlap, pitch_index, ratio, seed, st, start, tmp, tmp1, tmp_g, window, v1, v10, v103, v12, v14, v16, v18, v20, v22, v24, v26, v28, v3, v36, v40, v5, v7, v8
	C = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fchannels
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
	decode_buffer_size = int32(DEC_PITCH_BUF_SIZE)
	max_period = int32(MAX_PERIOD)
	mode = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fmode
	nbEBands = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FnbEBands
	overlap = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Foverlap
	eBands = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FeBands
	c = 0
	for {
		(*(*[2]uintptr)(unsafe.Pointer(bp)))[c] = st1 + 112 + uintptr(c*(decode_buffer_size+overlap))*4
		(*(*[2]uintptr)(unsafe.Pointer(bp + 16)))[c] = (*(*[2]uintptr)(unsafe.Pointer(bp)))[c] + uintptr(decode_buffer_size)*4 - uintptr(N)*4
		goto _6
	_6:
		;
		c = c + 1
		v5 = c
		if !(v5 < C) {
			break
		}
	}
	oldBandE = st1 + 112 + uintptr((decode_buffer_size+overlap)*C)*4
	oldLogE = oldBandE + uintptr(int32(2)*nbEBands)*4
	oldLogE2 = oldLogE + uintptr(int32(2)*nbEBands)*4
	backgroundLogE = oldLogE2 + uintptr(int32(2)*nbEBands)*4
	lpc = backgroundLogE + uintptr(int32(2)*nbEBands)*4
	loss_duration = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Floss_duration
	start = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fstart
	curr_frame_type = int32(FRAME_PLC_PERIODIC)
	if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fplc_duration >= int32(40) || start != 0 || (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fskip_plc != 0 {
		curr_frame_type = int32(FRAME_PLC_NOISE)
	}
	if curr_frame_type == int32(FRAME_PLC_NOISE) {
		end = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fend
		if end < (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FeffEBands {
			v7 = end
		} else {
			v7 = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FeffEBands
		}
		if start > v7 {
			v5 = start
		} else {
			if end < (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FeffEBands {
				v8 = end
			} else {
				v8 = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FeffEBands
			}
			v5 = v8
		}
		effEnd = v5
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			goto _11
		_11:
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v3 = st
		goto _13
	_13:
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v10 = libc.Xmalloc(tls, uint64(16))
			goto _15
		_15:
			st = v10
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v12 = st
		goto _17
	_17:
		*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v12)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v14 = libc.Xmalloc(tls, uint64(16))
			goto _19
		_19:
			st = v14
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v16 = st
		goto _21
	_21:
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v18 = libc.Xmalloc(tls, uint64(16))
			goto _23
		_23:
			st = v18
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v20 = st
		goto _25
	_25:
		if !(int64(int32(uint64(uint32(C*N))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v16)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v20)).Fglobal_stack)) {
			Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(749))
		}
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v22 = libc.Xmalloc(tls, uint64(16))
			goto _27
		_27:
			st = v22
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v24 = st
		goto _29
	_29:
		*(*uintptr)(unsafe.Pointer(v24 + 8)) += uintptr(uint64(uint32(C*N)) * (uint64(4) / uint64(1)))
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v26 = libc.Xmalloc(tls, uint64(16))
			goto _31
		_31:
			st = v26
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v28 = st
		goto _33
	_33:
		X = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v28)).Fglobal_stack - uintptr(uint64(uint32(C*N))*(uint64(4)/uint64(1))) /**< Interleaved normalised MDCTs */
		c = 0
		for {
			libc.Xmemmove(tls, (*(*[2]uintptr)(unsafe.Pointer(bp)))[c], (*(*[2]uintptr)(unsafe.Pointer(bp)))[c]+uintptr(N)*4, uint64(uint32(decode_buffer_size-N+overlap))*uint64(4)+uint64(0*((int64((*(*[2]uintptr)(unsafe.Pointer(bp)))[c])-int64((*(*[2]uintptr)(unsafe.Pointer(bp)))[c]+uintptr(N)*4))/4)))
			goto _35
		_35:
			;
			c = c + 1
			v5 = c
			if !(v5 < C) {
				break
			}
		}
		if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fprefilter_and_fold != 0 {
			prefilter_and_fold(tls, st1, N)
		}
		/* Energy decay */
		if loss_duration == 0 {
			v36 = float32(1.5)
		} else {
			v36 = float32(0.5)
		}
		decay = v36
		c = 0
		for {
			i = start
			for {
				if !(i < end) {
					break
				}
				if *(*OpusT_celt_glog)(unsafe.Pointer(backgroundLogE + uintptr(c*nbEBands+i)*4)) > *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4))-decay {
					v40 = *(*OpusT_celt_glog)(unsafe.Pointer(backgroundLogE + uintptr(c*nbEBands+i)*4))
				} else {
					v40 = *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4)) - decay
				}
				*(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4)) = v40
				goto _39
			_39:
				;
				i = i + 1
			}
			goto _38
		_38:
			;
			c = c + 1
			v5 = c
			if !(v5 < C) {
				break
			}
		}
		seed = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Frng
		c = 0
		for {
			if !(c < C) {
				break
			}
			i = start
			for {
				if !(i < effEnd) {
					break
				}
				boffs = N*c + int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i)*2)))<<LM
				blen = (int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i+int32(1))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i)*2)))) << LM
				j = 0
				for {
					if !(j < blen) {
						break
					}
					seed = Opus_celt_lcg_rand(tls, seed)
					*(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(boffs+j)*4)) = float32(int32(seed) >> int32(20))
					goto _43
				_43:
					;
					j = j + 1
				}
				Opus_renormalise_vector(tls, X+uintptr(boffs)*4, blen, float32(1), (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
				goto _42
			_42:
				;
				i = i + 1
			}
			goto _41
		_41:
			;
			c = c + 1
		}
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Frng = seed
		celt_synthesis(tls, mode, X, bp+16, oldBandE, start, effEnd, C, C, 0, LM, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdownsample, 0, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
		/* Run the postfilter with the last parameters. */
		c = 0
		for {
			if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period > int32(COMBFILTER_MINPERIOD) {
				v7 = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period
			} else {
				v7 = int32(COMBFILTER_MINPERIOD)
			}
			(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period = v7
			if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old > int32(COMBFILTER_MINPERIOD) {
				v5 = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old
			} else {
				v5 = int32(COMBFILTER_MINPERIOD)
			}
			(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old = v5
			Opus_comb_filter(tls, (*(*[2]uintptr)(unsafe.Pointer(bp + 16)))[c], (*(*[2]uintptr)(unsafe.Pointer(bp + 16)))[c], (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain_old, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset_old, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Fwindow, overlap, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
			if LM != 0 {
				Opus_comb_filter(tls, (*(*[2]uintptr)(unsafe.Pointer(bp + 16)))[c]+uintptr((*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize)*4, (*(*[2]uintptr)(unsafe.Pointer(bp + 16)))[c]+uintptr((*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize)*4, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period, N-(*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Fwindow, overlap, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
			}
			goto _45
		_45:
			;
			c = c + 1
			v5 = c
			if !(v5 < C) {
				break
			}
		}
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain_old = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset_old = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fprefilter_and_fold = 0
		/* Skip regular PLC until we get two consecutive packets. */
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fskip_plc = int32(1)
	} else {
		fade = float32(1)
		curr_neural = libc.BoolInt32(curr_frame_type == int32(FRAME_PLC_NEURAL) || curr_frame_type == int32(FRAME_DRED))
		last_neural = libc.BoolInt32((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Flast_frame_type == int32(FRAME_PLC_NEURAL) || (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Flast_frame_type == int32(FRAME_DRED))
		if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Flast_frame_type != int32(FRAME_PLC_PERIODIC) && !(last_neural != 0 && curr_neural != 0) {
			v5 = celt_plc_pitch_search(tls, st1, bp, C, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
			pitch_index = v5
			(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Flast_pitch_index = v5
		} else {
			pitch_index = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Flast_pitch_index
			fade = float32(0.8)
		}
		/* We want the excitation for 2 pitch periods in order to look for a
		   decaying signal, but we can't get more than MAX_PERIOD. */
		if int32(2)*pitch_index < max_period {
			v5 = int32(2) * pitch_index
		} else {
			v5 = max_period
		}
		exc_length = v5
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			goto _51
		_51:
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v3 = st
		goto _53
	_53:
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v10 = libc.Xmalloc(tls, uint64(16))
			goto _55
		_55:
			st = v10
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v12 = st
		goto _57
	_57:
		*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v12)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v14 = libc.Xmalloc(tls, uint64(16))
			goto _59
		_59:
			st = v14
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v16 = st
		goto _61
	_61:
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v18 = libc.Xmalloc(tls, uint64(16))
			goto _63
		_63:
			st = v18
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v20 = st
		goto _65
	_65:
		if !(int64(int32(uint64(uint32(max_period+int32(CELT_LPC_ORDER)))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v16)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v20)).Fglobal_stack)) {
			Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(837))
		}
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v22 = libc.Xmalloc(tls, uint64(16))
			goto _67
		_67:
			st = v22
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v24 = st
		goto _69
	_69:
		*(*uintptr)(unsafe.Pointer(v24 + 8)) += uintptr(uint64(uint32(max_period+int32(CELT_LPC_ORDER))) * (uint64(4) / uint64(1)))
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v26 = libc.Xmalloc(tls, uint64(16))
			goto _71
		_71:
			st = v26
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v28 = st
		goto _73
	_73:
		_exc = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v28)).Fglobal_stack - uintptr(uint64(uint32(max_period+int32(CELT_LPC_ORDER)))*(uint64(4)/uint64(1)))
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			goto _75
		_75:
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v3 = st
		goto _77
	_77:
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v10 = libc.Xmalloc(tls, uint64(16))
			goto _79
		_79:
			st = v10
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v12 = st
		goto _81
	_81:
		*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v12)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v14 = libc.Xmalloc(tls, uint64(16))
			goto _83
		_83:
			st = v14
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v16 = st
		goto _85
	_85:
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v18 = libc.Xmalloc(tls, uint64(16))
			goto _87
		_87:
			st = v18
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v20 = st
		goto _89
	_89:
		if !(int64(int32(uint64(uint32(exc_length))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v16)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v20)).Fglobal_stack)) {
			Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(838))
		}
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v22 = libc.Xmalloc(tls, uint64(16))
			goto _91
		_91:
			st = v22
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v24 = st
		goto _93
	_93:
		*(*uintptr)(unsafe.Pointer(v24 + 8)) += uintptr(uint64(uint32(exc_length)) * (uint64(4) / uint64(1)))
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v26 = libc.Xmalloc(tls, uint64(16))
			goto _95
		_95:
			st = v26
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v28 = st
		goto _97
	_97:
		fir_tmp = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v28)).Fglobal_stack - uintptr(uint64(uint32(exc_length))*(uint64(4)/uint64(1)))
		exc = _exc + uintptr(CELT_LPC_ORDER)*4
		window = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Fwindow
		c = 0
		for {
			S1 = float32(0)
			buf = (*(*[2]uintptr)(unsafe.Pointer(bp)))[c]
			i = 0
			for {
				if !(i < max_period+int32(CELT_LPC_ORDER)) {
					break
				}
				*(*OpusT_opus_val16)(unsafe.Pointer(exc + uintptr(i-int32(CELT_LPC_ORDER))*4)) = *(*OpusT_celt_sig)(unsafe.Pointer(buf + uintptr(decode_buffer_size-max_period-int32(CELT_LPC_ORDER)+i)*4))
				goto _100
			_100:
				;
				i = i + 1
			}
			if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Flast_frame_type != int32(FRAME_PLC_PERIODIC) && !(last_neural != 0 && curr_neural != 0) {
				/* Compute LPC coefficients for the last MAX_PERIOD samples before
				   the first loss so we can work in the excitation-filter domain. */
				Opus__celt_autocorr(tls, exc, bp+32, window, overlap, int32(CELT_LPC_ORDER), max_period, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
				/* Add a noise floor of -40 dB. */
				*(*OpusT_opus_val32)(unsafe.Pointer(bp + 32)) *= float32(1.0001)
				/* Use lag windowing to stabilize the Levinson-Durbin recursion. */
				i = int32(1)
				for {
					if !(i <= int32(CELT_LPC_ORDER)) {
						break
					}
					/*ac[i] *= exp(-.5*(2*M_PI*.002*i)*(2*M_PI*.002*i));*/
					*(*OpusT_opus_val32)(unsafe.Pointer(bp + 32 + uintptr(i)*4)) -= OpusT_opus_val32(OpusT_opus_val32(OpusT_opus_val32((*(*[25]OpusT_opus_val32)(unsafe.Pointer(bp + 32)))[i]*float32(float32(0.008)*float32(0.008)))*float32(i)) * float32(i))
					goto _101
				_101:
					;
					i = i + 1
				}
				Opus__celt_lpc(tls, lpc+uintptr(c*int32(CELT_LPC_ORDER))*4, bp+32, int32(CELT_LPC_ORDER))
			}
			/* Initialize the LPC history with the samples just before the start
			   of the region for which we're computing the excitation. */
			/* Compute the excitation for exc_length samples before the loss. We need the copy
			   because celt_fir() cannot filter in-place. */
			Opus_celt_fir_c(tls, exc+uintptr(max_period)*4-uintptr(exc_length)*4, lpc+uintptr(c*int32(CELT_LPC_ORDER))*4, fir_tmp, exc_length, int32(CELT_LPC_ORDER), (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
			libc.Xmemcpy(tls, exc+uintptr(max_period)*4-uintptr(exc_length)*4, fir_tmp, uint64(uint32(exc_length))*uint64(4)+uint64(0*((int64(exc+uintptr(max_period)*4-uintptr(exc_length)*4)-int64(fir_tmp))/4)))
			/* Check if the waveform is decaying, and if so how fast.
			   We do this to avoid adding energy when concealing in a segment
			   with decaying energy. */
			E1 = float32(1)
			E2 = float32(1)
			decay_length = exc_length >> int32(1)
			i = 0
			for {
				if !(i < decay_length) {
					break
				}
				e = *(*OpusT_opus_val16)(unsafe.Pointer(exc + uintptr(max_period-decay_length+i)*4))
				E1 = E1 + OpusT_opus_val32(e*e)
				e = *(*OpusT_opus_val16)(unsafe.Pointer(exc + uintptr(max_period-int32(2)*decay_length+i)*4))
				E2 = E2 + OpusT_opus_val32(e*e)
				goto _102
			_102:
				;
				i = i + 1
			}
			if E1 < E2 {
				v103 = E1
			} else {
				v103 = E2
			}
			E1 = v103
			decay1 = float32(libc.Xsqrt(tls, float64(E1/E2)))
			/* Move the decoder memory one frame to the left to give us room to
			   add the data for the new frame. We ignore the overlap that extends
			   past the end of the buffer, because we aren't going to use it. */
			libc.Xmemmove(tls, buf, buf+uintptr(N)*4, uint64(uint32(decode_buffer_size-N))*uint64(4)+uint64(0*((int64(buf)-int64(buf+uintptr(N)*4))/4)))
			/* Extrapolate from the end of the excitation with a period of
			   "pitch_index", scaling down each period by an additional factor of
			   "decay". */
			extrapolation_offset = max_period - pitch_index
			/* We need to extrapolate enough samples to cover a complete MDCT
			   window (including overlap/2 samples on both sides). */
			extrapolation_len = N + overlap
			/* We also apply fading if this is not the first loss. */
			attenuation = OpusT_opus_val16(fade * decay1)
			v5 = int32(0)
			j1 = v5
			i = v5
			for {
				if !(i < extrapolation_len) {
					break
				}
				if j1 >= pitch_index {
					j1 = j1 - pitch_index
					attenuation = OpusT_opus_val16(attenuation * decay1)
				}
				*(*OpusT_celt_sig)(unsafe.Pointer(buf + uintptr(decode_buffer_size-N+i)*4)) = OpusT_opus_val16(attenuation * *(*OpusT_opus_val16)(unsafe.Pointer(exc + uintptr(extrapolation_offset+j1)*4)))
				/* Compute the energy of the previously decoded signal whose
				   excitation we're copying. */
				tmp = *(*OpusT_celt_sig)(unsafe.Pointer(buf + uintptr(decode_buffer_size-max_period-N+extrapolation_offset+j1)*4))
				S1 = S1 + OpusT_opus_val32(tmp*tmp)
				goto _104
			_104:
				;
				i = i + 1
				j1 = j1 + 1
			}
			/* Copy the last decoded samples (prior to the overlap region) to
			   synthesis filter memory so we can have a continuous signal. */
			i = 0
			for {
				if !(i < int32(CELT_LPC_ORDER)) {
					break
				}
				(*(*[24]OpusT_opus_val16)(unsafe.Pointer(bp + 132)))[i] = *(*OpusT_celt_sig)(unsafe.Pointer(buf + uintptr(decode_buffer_size-N-int32(1)-i)*4))
				goto _106
			_106:
				;
				i = i + 1
			}
			/* Apply the synthesis filter to convert the excitation back into
			   the signal domain. */
			Opus_celt_iir(tls, buf+uintptr(decode_buffer_size)*4-uintptr(N)*4, lpc+uintptr(c*int32(CELT_LPC_ORDER))*4, buf+uintptr(decode_buffer_size)*4-uintptr(N)*4, extrapolation_len, int32(CELT_LPC_ORDER), bp+132, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
			/* Check if the synthesis energy is higher than expected, which can
			   happen with the signal changes during our window. If so,
			   attenuate. */
			S2 = float32(0)
			i = 0
			for {
				if !(i < extrapolation_len) {
					break
				}
				tmp1 = *(*OpusT_celt_sig)(unsafe.Pointer(buf + uintptr(decode_buffer_size-N+i)*4))
				S2 = S2 + OpusT_opus_val32(tmp1*tmp1)
				goto _107
			_107:
				;
				i = i + 1
			}
			/* This checks for an "explosion" in the synthesis. */
			/* The float test is written this way to catch NaNs in the output
			   of the IIR filter at the same time. */
			if !(S1 > OpusT_opus_val32(float32(0.2)*S2)) {
				i = 0
				for {
					if !(i < extrapolation_len) {
						break
					}
					*(*OpusT_celt_sig)(unsafe.Pointer(buf + uintptr(decode_buffer_size-N+i)*4)) = float32(0)
					goto _108
				_108:
					;
					i = i + 1
				}
			} else {
				if S1 < S2 {
					ratio = float32(libc.Xsqrt(tls, float64((S1+float32(1))/(S2+float32(1)))))
					i = 0
					for {
						if !(i < overlap) {
							break
						}
						tmp_g = float32(1) - OpusT_celt_coef(*(*OpusT_celt_coef)(unsafe.Pointer(window + uintptr(i)*4))*(float32(1)-ratio))
						*(*OpusT_celt_sig)(unsafe.Pointer(buf + uintptr(decode_buffer_size-N+i)*4)) = OpusT_opus_val16(tmp_g * *(*OpusT_celt_sig)(unsafe.Pointer(buf + uintptr(decode_buffer_size-N+i)*4)))
						goto _109
					_109:
						;
						i = i + 1
					}
					i = overlap
					for {
						if !(i < extrapolation_len) {
							break
						}
						*(*OpusT_celt_sig)(unsafe.Pointer(buf + uintptr(decode_buffer_size-N+i)*4)) = OpusT_opus_val16(ratio * *(*OpusT_celt_sig)(unsafe.Pointer(buf + uintptr(decode_buffer_size-N+i)*4)))
						goto _110
					_110:
						;
						i = i + 1
					}
				}
			}
			goto _99
		_99:
			;
			c = c + 1
			v5 = c
			if !(v5 < C) {
				break
			}
		}
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fprefilter_and_fold = int32(1)
	}
	/* Saturate to something large to avoid wrap-around. */
	if int32(10000) < loss_duration+int32(1)<<LM {
		v5 = int32(10000)
	} else {
		v5 = loss_duration + int32(1)<<LM
	}
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Floss_duration = v5
	if int32(10000) < (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fplc_duration+int32(1)<<LM {
		v5 = int32(10000)
	} else {
		v5 = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fplc_duration + int32(1)<<LM
	}
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fplc_duration = v5
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Flast_frame_type = curr_frame_type
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _114
	_114:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _116
_116:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
}

func Opus_celt_decode_with_ec_dred(tls *libc.TLS, st1 uintptr, data uintptr, len1 int32, pcm uintptr, frame_size int32, dec uintptr, accum int32) (r int32) {
	bp := tls.Alloc(96)
	defer tls.Free(96)
	var C, CC, LM, M, N, alloc_trim, anti_collapse_on, anti_collapse_rsv, boost, c, codedBands, decode_buffer_size, dynalloc_logp, dynalloc_loop_logp, effEnd, end, flag, i, intra_ener, isTransient, missing, nbEBands, octave, overlap, postfilter_pitch, postfilter_tapset, qg, quanta, shortBlocks, silence, spread_decision, start, width, v28, v37, v40 int32
	var E0, E1, E2, slope, v57 OpusT_opus_val32
	var X, _saved_stack, backgroundLogE, cap1, collapse_masks, eBands, fine_priority, fine_quant, mode, offsets, oldBandE, oldLogE, oldLogE2, pulses, st, tf_res, v1, v10, v11, v13, v15, v17, v19, v21, v3, v5, v6, v8 uintptr
	var bits, tell, total_bits OpusT_opus_int32
	var decode_mem [2]uintptr
	var max_background_increase, safety, v35, v56, v61 OpusT_celt_glog
	var postfilter_gain OpusT_opus_val16
	var v60 float32
	var _ /* _dec at bp+0 */ OpusT_ec_dec
	var _ /* balance at bp+80 */ OpusT_opus_int32
	var _ /* dual_stereo at bp+76 */ int32
	var _ /* intensity at bp+72 */ int32
	var _ /* out_syn at bp+56 */ [2]uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, CC, E0, E1, E2, LM, M, N, X, _saved_stack, alloc_trim, anti_collapse_on, anti_collapse_rsv, backgroundLogE, bits, boost, c, cap1, codedBands, collapse_masks, decode_buffer_size, decode_mem, dynalloc_logp, dynalloc_loop_logp, eBands, effEnd, end, fine_priority, fine_quant, flag, i, intra_ener, isTransient, max_background_increase, missing, mode, nbEBands, octave, offsets, oldBandE, oldLogE, oldLogE2, overlap, postfilter_gain, postfilter_pitch, postfilter_tapset, pulses, qg, quanta, safety, shortBlocks, silence, slope, spread_decision, st, start, tell, tf_res, total_bits, width, v1, v10, v11, v13, v15, v17, v19, v21, v28, v3, v35, v37, v40, v5, v56, v57, v6, v60, v61, v8
	CC = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fchannels
	*(*int32)(unsafe.Pointer(bp + 72)) = 0
	*(*int32)(unsafe.Pointer(bp + 76)) = 0
	anti_collapse_on = 0
	C = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fstream_channels
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
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v6 = libc.Xmalloc(tls, uint64(16))
		goto _7
	_7:
		st = v6
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v8 = st
	goto _9
_9:
	if (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v8)).Fglobal_stack == uintptr(0) {
		v13 = libc.Xmalloc(tls, uint64(GLOBAL_STACK_SIZE))
		goto _14
	_14:
		v11 = v13
		goto _12
	_12:
		v10 = v11
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v15 = libc.Xmalloc(tls, uint64(16))
			goto _16
		_16:
			st = v15
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v17 = st
		goto _18
	_18:
		(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fscratch_ptr = v10
		v5 = v10
	} else {
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v19 = libc.Xmalloc(tls, uint64(16))
			goto _20
		_20:
			st = v19
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v21 = st
		goto _22
	_22:
		v5 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack
	}
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = v5
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _24
	_24:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _26
_26:
	_saved_stack = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack
	decode_buffer_size = int32(DEC_PITCH_BUF_SIZE)
	Opus_validate_celt_decoder(tls, st1)
	mode = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fmode
	nbEBands = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FnbEBands
	overlap = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Foverlap
	eBands = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FeBands
	start = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fstart
	end = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fend
	frame_size = frame_size * (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdownsample
	oldBandE = st1 + 112 + uintptr((decode_buffer_size+overlap)*CC)*4
	oldLogE = oldBandE + uintptr(int32(2)*nbEBands)*4
	oldLogE2 = oldLogE + uintptr(int32(2)*nbEBands)*4
	backgroundLogE = oldLogE2 + uintptr(int32(2)*nbEBands)*4
	LM = 0
	for {
		if !(LM <= (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FmaxLM) {
			break
		}
		if (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize<<LM == frame_size {
			break
		}
		goto _27
	_27:
		;
		LM = LM + 1
	}
	if LM > (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FmaxLM {
		return -int32(1)
	}
	M = int32(1) << LM
	if len1 < 0 || len1 > int32(1275) || pcm == uintptr(uint32(0)) {
		return -int32(1)
	}
	N = M * (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize
	c = 0
	for {
		decode_mem[c] = st1 + 112 + uintptr(c*(decode_buffer_size+overlap))*4
		(*(*[2]uintptr)(unsafe.Pointer(bp + 56)))[c] = decode_mem[c] + uintptr(decode_buffer_size)*4 - uintptr(N)*4
		goto _29
	_29:
		;
		c = c + 1
		v28 = c
		if !(v28 < CC) {
			break
		}
	}
	effEnd = end
	if effEnd > (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FeffEBands {
		effEnd = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FeffEBands
	}
	if data == uintptr(uint32(0)) || len1 <= int32(1) {
		celt_decode_lost(tls, st1, N, LM)
		deemphasis(tls, bp+56, pcm, N, CC, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdownsample, mode+16, st1+104, accum)
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			goto _31
		_31:
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v3 = st
		goto _33
	_33:
		(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
		return frame_size / (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdownsample
	}
	/* Check if there are at least two packets received consecutively before
	 * turning on the pitch-based PLC */
	if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Floss_duration == 0 {
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fskip_plc = 0
	}
	if dec == uintptr(uint32(0)) {
		Opus_ec_dec_init(tls, bp, data, uint32(len1))
		dec = bp
	}
	if C == int32(1) {
		i = 0
		for {
			if !(i < nbEBands) {
				break
			}
			if *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(i)*4)) > *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(nbEBands+i)*4)) {
				v35 = *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(i)*4))
			} else {
				v35 = *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(nbEBands+i)*4))
			}
			*(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(i)*4)) = v35
			goto _34
		_34:
			;
			i = i + 1
		}
	}
	total_bits = len1 * int32(8)
	v1 = dec
	v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
	goto _38
_38:
	tell = v28
	if tell >= total_bits {
		silence = int32(1)
	} else {
		if tell == int32(1) {
			silence = Opus_ec_dec_bit_logp(tls, dec, uint32(15))
		} else {
			silence = 0
		}
	}
	if silence != 0 {
		/* Pretend we've read all the remaining bits */
		tell = len1 * int32(8)
		v1 = dec
		v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
		goto _41
	_41:
		*(*int32)(unsafe.Pointer(dec + 24)) += tell - v28
	}
	postfilter_gain = float32(0)
	postfilter_pitch = 0
	postfilter_tapset = 0
	if start == 0 && tell+int32(16) <= total_bits {
		if Opus_ec_dec_bit_logp(tls, dec, uint32(1)) != 0 {
			octave = int32(Opus_ec_dec_uint(tls, dec, uint32(6)))
			postfilter_pitch = int32(uint32(int32(16)<<octave) + Opus_ec_dec_bits(tls, dec, uint32(int32(4)+octave)) - uint32(1))
			qg = int32(Opus_ec_dec_bits(tls, dec, uint32(3)))
			v1 = dec
			v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
			goto _44
		_44:
			if v28+int32(2) <= total_bits {
				postfilter_tapset = Opus_ec_dec_icdf(tls, dec, uintptr(unsafe.Pointer(&tapset_icdf9)), uint32(2))
			}
			postfilter_gain = OpusT_opus_val16(float32(0.09375) * float32(qg+int32(1)))
		}
		v1 = dec
		v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
		goto _47
	_47:
		tell = v28
	}
	if LM > 0 && tell+int32(3) <= total_bits {
		isTransient = Opus_ec_dec_bit_logp(tls, dec, uint32(3))
		v1 = dec
		v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
		goto _50
	_50:
		tell = v28
	} else {
		isTransient = 0
	}
	if isTransient != 0 {
		shortBlocks = M
	} else {
		shortBlocks = 0
	}
	/* Decode the global flags (first symbols in the stream) */
	if tell+int32(3) <= total_bits {
		v28 = Opus_ec_dec_bit_logp(tls, dec, uint32(3))
	} else {
		v28 = 0
	}
	intra_ener = v28
	/* If recovering from packet loss, make sure we make the energy prediction safe to reduce the
	   risk of getting loud artifacts. */
	if !(intra_ener != 0) && (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Floss_duration != 0 {
		c = 0
		for {
			safety = float32(0)
			if int32(10) < (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Floss_duration>>LM {
				v37 = int32(10)
			} else {
				v37 = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Floss_duration >> LM
			}
			missing = v37
			if LM == 0 {
				safety = float32(1.5)
			} else {
				if LM == int32(1) {
					safety = float32(0.5)
				}
			}
			i = start
			for {
				if !(i < end) {
					break
				}
				if *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(c*nbEBands+i)*4)) > *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE2 + uintptr(c*nbEBands+i)*4)) {
					v35 = *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(c*nbEBands+i)*4))
				} else {
					v35 = *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE2 + uintptr(c*nbEBands+i)*4))
				}
				if *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4)) < v35 {
					E0 = *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4))
					E1 = *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(c*nbEBands+i)*4))
					E2 = *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE2 + uintptr(c*nbEBands+i)*4))
					if E1-E0 > float32(float32(0.5)*(E2-E0)) {
						v57 = E1 - E0
					} else {
						v57 = float32(float32(0.5) * (E2 - E0))
					}
					slope = v57
					if slope < float32(2) {
						v57 = slope
					} else {
						v57 = float32(2)
					}
					slope = v57
					if float32(int32(0)) > OpusT_opus_val32(float32(int32(1)+missing)*slope) {
						v57 = float32(int32(0))
					} else {
						v57 = OpusT_opus_val32(float32(int32(1)+missing) * slope)
					}
					E0 = E0 - v57
					if -float32(20) > E0 {
						v60 = -float32(20)
					} else {
						v60 = E0
					}
					*(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4)) = v60
				} else {
					/* Otherwise take the min of the last frames. */
					if *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4)) < *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(c*nbEBands+i)*4)) {
						v56 = *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4))
					} else {
						v56 = *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(c*nbEBands+i)*4))
					}
					if v56 < *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE2 + uintptr(c*nbEBands+i)*4)) {
						if *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4)) < *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(c*nbEBands+i)*4)) {
							v61 = *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4))
						} else {
							v61 = *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(c*nbEBands+i)*4))
						}
						v35 = v61
					} else {
						v35 = *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE2 + uintptr(c*nbEBands+i)*4))
					}
					*(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4)) = v35
				}
				/* Shorter frames have more natural fluctuations -- play it safe. */
				*(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4)) -= safety
				goto _55
			_55:
				;
				i = i + 1
			}
			goto _53
		_53:
			;
			c = c + 1
			v28 = c
			if !(v28 < int32(2)) {
				break
			}
		}
	}
	/* Get band energies */
	Opus_unquant_coarse_energy(tls, mode, start, end, oldBandE, intra_ener, dec, C, LM)
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _65
	_65:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _67
_67:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _69
	_69:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	goto _71
_71:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _73
	_73:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	goto _75
_75:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _77
	_77:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _79
_79:
	if !(int64(int32(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(1395))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _81
	_81:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _83
_83:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(uint64(uint32(nbEBands)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _85
	_85:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _87
_87:
	tf_res = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))
	tf_decode(tls, start, end, isTransient, tf_res, LM, dec)
	v1 = dec
	v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
	goto _90
_90:
	tell = v28
	spread_decision = int32(SPREAD_NORMAL)
	if tell+int32(4) <= total_bits {
		spread_decision = Opus_ec_dec_icdf(tls, dec, uintptr(unsafe.Pointer(&spread_icdf9)), uint32(5))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _92
	_92:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _94
_94:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _96
	_96:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	goto _98
_98:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _100
	_100:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	goto _102
_102:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _104
	_104:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _106
_106:
	if !(int64(int32(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(1403))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _108
	_108:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _110
_110:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(uint64(uint32(nbEBands)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _112
	_112:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _114
_114:
	cap1 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))
	Opus_init_caps(tls, mode, cap1, LM, C)
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _116
	_116:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _118
_118:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _120
	_120:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	goto _122
_122:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _124
	_124:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	goto _126
_126:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _128
	_128:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _130
_130:
	if !(int64(int32(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(1407))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _132
	_132:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _134
_134:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(uint64(uint32(nbEBands)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _136
	_136:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _138
_138:
	offsets = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))
	dynalloc_logp = int32(6)
	total_bits = total_bits << int32(BITRES)
	tell = int32(Opus_ec_tell_frac(tls, dec))
	i = start
	for {
		if !(i < end) {
			break
		}
		width = C * (int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i+int32(1))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i)*2)))) << LM
		/* quanta is 6 bits, but no more than 1 bit/sample
		   and no less than 1/8 bit/sample */
		if int32(6)<<int32(BITRES) > width {
			v37 = int32(6) << int32(BITRES)
		} else {
			v37 = width
		}
		if width<<int32(BITRES) < v37 {
			v28 = width << int32(BITRES)
		} else {
			if int32(6)<<int32(BITRES) > width {
				v40 = int32(6) << int32(BITRES)
			} else {
				v40 = width
			}
			v28 = v40
		}
		quanta = v28
		dynalloc_loop_logp = dynalloc_logp
		boost = 0
		for tell+dynalloc_loop_logp<<int32(BITRES) < total_bits && boost < *(*int32)(unsafe.Pointer(cap1 + uintptr(i)*4)) {
			flag = Opus_ec_dec_bit_logp(tls, dec, uint32(dynalloc_loop_logp))
			tell = int32(Opus_ec_tell_frac(tls, dec))
			if !(flag != 0) {
				break
			}
			boost = boost + quanta
			total_bits = total_bits - quanta
			dynalloc_loop_logp = int32(1)
		}
		*(*int32)(unsafe.Pointer(offsets + uintptr(i)*4)) = boost
		/* Making dynalloc more likely */
		if boost > 0 {
			if int32(2) > dynalloc_logp-int32(1) {
				v28 = int32(2)
			} else {
				v28 = dynalloc_logp - int32(1)
			}
			dynalloc_logp = v28
		}
		goto _139
	_139:
		;
		i = i + 1
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _145
	_145:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _147
_147:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _149
	_149:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	goto _151
_151:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _153
	_153:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	goto _155
_155:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _157
	_157:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _159
_159:
	if !(int64(int32(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(1440))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _161
	_161:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _163
_163:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(uint64(uint32(nbEBands)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _165
	_165:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _167
_167:
	fine_quant = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))
	if tell+int32(6)<<int32(BITRES) <= total_bits {
		v28 = Opus_ec_dec_icdf(tls, dec, uintptr(unsafe.Pointer(&trim_icdf9)), uint32(7))
	} else {
		v28 = int32(5)
	}
	alloc_trim = v28
	bits = len1*int32(8)<<int32(BITRES) - int32(Opus_ec_tell_frac(tls, dec)) - int32(1)
	if isTransient != 0 && LM >= int32(2) && bits >= (LM+int32(2))<<int32(BITRES) {
		v28 = int32(1) << int32(BITRES)
	} else {
		v28 = 0
	}
	anti_collapse_rsv = v28
	bits = bits - anti_collapse_rsv
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _171
	_171:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _173
_173:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _175
	_175:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	goto _177
_177:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _179
	_179:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	goto _181
_181:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _183
	_183:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _185
_185:
	if !(int64(int32(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(1448))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _187
	_187:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _189
_189:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(uint64(uint32(nbEBands)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _191
	_191:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _193
_193:
	pulses = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _195
	_195:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _197
_197:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _199
	_199:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	goto _201
_201:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _203
	_203:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	goto _205
_205:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _207
	_207:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _209
_209:
	if !(int64(int32(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(1449))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _211
	_211:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _213
_213:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(uint64(uint32(nbEBands)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _215
	_215:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _217
_217:
	fine_priority = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(nbEBands))*(uint64(4)/uint64(1)))
	codedBands = Opus_clt_compute_allocation(tls, mode, start, end, offsets, cap1, alloc_trim, bp+72, bp+76, bits, bp+80, pulses, fine_quant, fine_priority, C, LM, dec, 0, 0, 0)
	Opus_unquant_fine_energy(tls, mode, start, end, oldBandE, uintptr(uint32(0)), fine_quant, dec, C)
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _219
	_219:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _221
_221:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _223
	_223:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	goto _225
_225:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _227
	_227:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	goto _229
_229:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _231
	_231:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _233
_233:
	if !(int64(int32(uint64(uint32(C*N))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(1457))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _235
	_235:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _237
_237:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(uint64(uint32(C*N)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _239
	_239:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _241
_241:
	X = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(C*N))*(uint64(4)/uint64(1))) /**< Interleaved normalised MDCTs */
	c = 0
	for {
		libc.Xmemmove(tls, decode_mem[c], decode_mem[c]+uintptr(N)*4, uint64(uint32(decode_buffer_size-N+overlap))*uint64(4)+uint64(0*((int64(decode_mem[c])-int64(decode_mem[c]+uintptr(N)*4))/4)))
		goto _243
	_243:
		;
		c = c + 1
		v28 = c
		if !(v28 < CC) {
			break
		}
	}
	/* Decode fixed codebook */
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _245
	_245:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _247
_247:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _249
	_249:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	goto _251
_251:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(1)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(1)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _253
	_253:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	goto _255
_255:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _257
	_257:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _259
_259:
	if !(int64(int32(uint64(uint32(C*nbEBands))*(uint64(1)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+3767, int32(1487))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _261
	_261:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _263
_263:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(uint64(uint32(C*nbEBands)) * (uint64(1) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _265
	_265:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _267
_267:
	collapse_masks = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(C*nbEBands))*(uint64(1)/uint64(1)))
	if C == int32(2) {
		v1 = X + uintptr(N)*4
	} else {
		v1 = uintptr(uint32(0))
	}
	Opus_quant_all_bands(tls, 0, mode, start, end, X, v1, collapse_masks, uintptr(uint32(0)), pulses, shortBlocks, spread_decision, *(*int32)(unsafe.Pointer(bp + 76)), *(*int32)(unsafe.Pointer(bp + 72)), tf_res, len1*(int32(8)<<int32(BITRES))-anti_collapse_rsv, *(*OpusT_opus_int32)(unsafe.Pointer(bp + 80)), dec, LM, codedBands, st1+48, 0, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdisable_inv)
	if anti_collapse_rsv > 0 {
		anti_collapse_on = int32(Opus_ec_dec_bits(tls, dec, uint32(1)))
	}
	v1 = dec
	v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
	goto _271
_271:
	Opus_unquant_energy_finalise(tls, mode, start, end, oldBandE, fine_quant, fine_priority, len1*int32(8)-v28, dec, C)
	if anti_collapse_on != 0 {
		Opus_anti_collapse(tls, mode, X, collapse_masks, LM, C, N, start, end, oldBandE, oldLogE, oldLogE2, pulses, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Frng, 0, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
	}
	if silence != 0 {
		i = 0
		for {
			if !(i < C*nbEBands) {
				break
			}
			*(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(i)*4)) = -float32(28)
			goto _272
		_272:
			;
			i = i + 1
		}
	}
	if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fprefilter_and_fold != 0 {
		prefilter_and_fold(tls, st1, N)
	}
	celt_synthesis(tls, mode, X, bp+56, oldBandE, start, effEnd, C, CC, isTransient, LM, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdownsample, silence, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
	c = 0
	for {
		if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period > int32(COMBFILTER_MINPERIOD) {
			v37 = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period
		} else {
			v37 = int32(COMBFILTER_MINPERIOD)
		}
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period = v37
		if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old > int32(COMBFILTER_MINPERIOD) {
			v28 = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old
		} else {
			v28 = int32(COMBFILTER_MINPERIOD)
		}
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old = v28
		Opus_comb_filter(tls, (*(*[2]uintptr)(unsafe.Pointer(bp + 56)))[c], (*(*[2]uintptr)(unsafe.Pointer(bp + 56)))[c], (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain_old, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset_old, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Fwindow, overlap, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
		if LM != 0 {
			Opus_comb_filter(tls, (*(*[2]uintptr)(unsafe.Pointer(bp + 56)))[c]+uintptr((*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize)*4, (*(*[2]uintptr)(unsafe.Pointer(bp + 56)))[c]+uintptr((*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize)*4, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period, postfilter_pitch, N-(*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain, postfilter_gain, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset, postfilter_tapset, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Fwindow, overlap, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
		}
		goto _274
	_274:
		;
		c = c + 1
		v28 = c
		if !(v28 < CC) {
			break
		}
	}
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain_old = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset_old = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period = postfilter_pitch
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain = postfilter_gain
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset = postfilter_tapset
	if LM != 0 {
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period_old = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_period
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain_old = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_gain
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset_old = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fpostfilter_tapset
	}
	if C == int32(1) {
		libc.Xmemcpy(tls, oldBandE+uintptr(nbEBands)*4, oldBandE, uint64(uint32(nbEBands))*uint64(4)+uint64(0*((OpusT___predefined_ptrdiff_t(oldBandE+uintptr(nbEBands)*4)-int64(oldBandE))/4)))
	}
	if !(isTransient != 0) {
		libc.Xmemcpy(tls, oldLogE2, oldLogE, uint64(uint32(int32(2)*nbEBands))*uint64(4)+uint64(0*((int64(oldLogE2)-int64(oldLogE))/4)))
		libc.Xmemcpy(tls, oldLogE, oldBandE, uint64(uint32(int32(2)*nbEBands))*uint64(4)+uint64(0*((int64(oldLogE)-int64(oldBandE))/4)))
	} else {
		i = 0
		for {
			if !(i < int32(2)*nbEBands) {
				break
			}
			if *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(i)*4)) < *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(i)*4)) {
				v35 = *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(i)*4))
			} else {
				v35 = *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(i)*4))
			}
			*(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(i)*4)) = v35
			goto _277
		_277:
			;
			i = i + 1
		}
	}
	/* In normal circumstances, we only allow the noise floor to increase by
	   up to 2.4 dB/second, but when we're in DTX we give the weight of
	   all missing packets to the update packet. */
	if int32(160) < (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Floss_duration+M {
		v28 = int32(160)
	} else {
		v28 = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Floss_duration + M
	}
	max_background_increase = OpusT_celt_glog(float32(v28) * float32(0.001))
	i = 0
	for {
		if !(i < int32(2)*nbEBands) {
			break
		}
		if *(*OpusT_celt_glog)(unsafe.Pointer(backgroundLogE + uintptr(i)*4))+max_background_increase < *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(i)*4)) {
			v35 = *(*OpusT_celt_glog)(unsafe.Pointer(backgroundLogE + uintptr(i)*4)) + max_background_increase
		} else {
			v35 = *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(i)*4))
		}
		*(*OpusT_celt_glog)(unsafe.Pointer(backgroundLogE + uintptr(i)*4)) = v35
		goto _280
	_280:
		;
		i = i + 1
	}
	/* In case start or end were to change */
	c = 0
	for {
		i = 0
		for {
			if !(i < start) {
				break
			}
			*(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4)) = float32(0)
			v35 = -float32(28)
			*(*OpusT_celt_glog)(unsafe.Pointer(oldLogE2 + uintptr(c*nbEBands+i)*4)) = v35
			*(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(c*nbEBands+i)*4)) = v35
			goto _284
		_284:
			;
			i = i + 1
		}
		i = end
		for {
			if !(i < nbEBands) {
				break
			}
			*(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4)) = float32(0)
			v35 = -float32(28)
			*(*OpusT_celt_glog)(unsafe.Pointer(oldLogE2 + uintptr(c*nbEBands+i)*4)) = v35
			*(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(c*nbEBands+i)*4)) = v35
			goto _286
		_286:
			;
			i = i + 1
		}
		goto _283
	_283:
		;
		c = c + 1
		v28 = c
		if !(v28 < int32(2)) {
			break
		}
	}
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Frng = (*OpusT_ec_dec)(unsafe.Pointer(dec)).Frng
	deemphasis(tls, bp+56, pcm, N, CC, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdownsample, mode+16, st1+104, accum)
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Floss_duration = 0
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fplc_duration = 0
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Flast_frame_type = int32(FRAME_NORMAL)
	(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fprefilter_and_fold = 0
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _289
	_289:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _291
_291:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
	v1 = dec
	v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
	goto _294
_294:
	if v28 > int32(8)*len1 {
		return -int32(3)
	}
	v28 = (*OpusT_ec_ctx)(unsafe.Pointer(dec)).Ferror1
	goto _296
_296:
	if v28 != 0 {
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Ferror1 = int32(1)
	}
	return frame_size / (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdownsample
}

func Opus_celt_decode_with_ec(tls *libc.TLS, st uintptr, data uintptr, len1 int32, pcm uintptr, frame_size int32, dec uintptr, accum int32) (r int32) {
	return Opus_celt_decode_with_ec_dred(tls, st, data, len1, pcm, frame_size, dec, accum)
}
