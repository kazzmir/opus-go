// Code generated for linux/amd64 by 'ccgo --package-name opusccenc --prefix-external Opus_ --prefix-typename OpusT_ -o opusccenc/libopus_enc.go -I .. -I ../include -I ../src -I ../celt -I ../silk -I ../silk/float -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ccgo_support/math_lrint.c ../src/opus.c ../src/opus_decoder.c ../src/opus_encoder.c ../src/extensions.c ../src/opus_multistream.c ../src/opus_multistream_encoder.c ../src/opus_multistream_decoder.c ../src/repacketizer.c ../src/opus_projection_encoder.c ../src/opus_projection_decoder.c ../src/mapping_matrix.c ../src/analysis.c ../src/mlp.c ../src/mlp_data.c ../celt/bands.c ../celt/celt.c ../celt/celt_encoder.c ../celt/celt_decoder.c ../celt/cwrs.c ../celt/entcode.c ../celt/entdec.c ../celt/entenc.c ../celt/kiss_fft.c ../celt/laplace.c ../celt/mathops.c ../celt/mdct.c ../celt/modes.c ../celt/pitch.c ../celt/celt_lpc.c ../celt/quant_bands.c ../celt/rate.c ../celt/vq.c ../celt/mini_kfft.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/enc_API.c ../silk/encode_indices.c ../silk/encode_pulses.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/NSQ.c ../silk/NSQ_del_dec.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/control_audio_bandwidth.c ../silk/quant_LTP_gains.c ../silk/VQ_WMat_EC.c ../silk/HP_variable_cutoff.c ../silk/NLSF_encode.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/process_NLSFs.c ../silk/stereo_LR_to_MS.c ../silk/stereo_MS_to_LR.c ../silk/check_control_input.c ../silk/control_SNR.c ../silk/init_encoder.c ../silk/control_codec.c ../silk/A2NLSF.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c ../silk/stereo_encode_pred.c ../silk/stereo_find_predictor.c ../silk/stereo_quant_pred.c ../silk/LPC_fit.c ../silk/float/apply_sine_window_FLP.c ../silk/float/corrMatrix_FLP.c ../silk/float/encode_frame_FLP.c ../silk/float/find_LPC_FLP.c ../silk/float/find_LTP_FLP.c ../silk/float/find_pitch_lags_FLP.c ../silk/float/find_pred_coefs_FLP.c ../silk/float/LPC_analysis_filter_FLP.c ../silk/float/LTP_analysis_filter_FLP.c ../silk/float/LTP_scale_ctrl_FLP.c ../silk/float/noise_shape_analysis_FLP.c ../silk/float/process_gains_FLP.c ../silk/float/regularize_correlations_FLP.c ../silk/float/residual_energy_FLP.c ../silk/float/warped_autocorrelation_FLP.c ../silk/float/wrappers_FLP.c ../silk/float/autocorrelation_FLP.c ../silk/float/burg_modified_FLP.c ../silk/float/bwexpander_FLP.c ../silk/float/energy_FLP.c ../silk/float/inner_product_FLP.c ../silk/float/k2a_FLP.c ../silk/float/LPC_inv_pred_gain_FLP.c ../silk/float/pitch_analysis_core_FLP.c ../silk/float/scale_copy_vector_FLP.c ../silk/float/scale_vector_FLP.c ../silk/float/schur_FLP.c ../silk/float/sort_FLP.c', DO NOT EDIT.

package opusccenc

import (
	"reflect"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

var _ reflect.Type
var _ unsafe.Pointer

func Opus__celt_autocorr(tls *libc.TLS, x uintptr, ac uintptr, window uintptr, overlap int32, lag int32, n int32, arch int32) (r int32) {
	var _saved_stack, st, xptr, xx, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var d OpusT_opus_val32
	var fastN, i, k, shift int32
	var w OpusT_opus_val16
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = _saved_stack, d, fastN, i, k, shift, st, w, xptr, xx, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9
	fastN = n - lag
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _2
	_2:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _4
_4:
	_saved_stack = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _6
	_6:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _8
_8:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _10
	_10:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v7 = st
	goto _12
_12:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		goto _14
	_14:
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v11 = st
	goto _16
_16:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		goto _18
	_18:
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v15 = st
	goto _20
_20:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(n)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5955, int32(301))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		goto _22
	_22:
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v19 = st
	goto _24
_24:
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(libc.Uint64FromInt32(n) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		goto _26
	_26:
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v23 = st
	goto _28
_28:
	xx = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(libc.Uint64FromInt32(n)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	if !(n > libc.Int32FromInt32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+6003, __ccgo_ts+5955, int32(302))
	}
	if !(overlap >= libc.Int32FromInt32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+6025, __ccgo_ts+5955, int32(303))
	}
	if overlap == 0 {
		xptr = x
	} else {
		i = 0
		for {
			if !(i < n) {
				break
			}
			*(*OpusT_opus_val16)(unsafe.Pointer(xx + uintptr(i)*4)) = *(*OpusT_opus_val16)(unsafe.Pointer(x + uintptr(i)*4))
			goto _29
		_29:
			;
			i = i + 1
		}
		i = 0
		for {
			if !(i < overlap) {
				break
			}
			w = *(*OpusT_celt_coef)(unsafe.Pointer(window + uintptr(i)*4))
			*(*OpusT_opus_val16)(unsafe.Pointer(xx + uintptr(i)*4)) = OpusT_opus_val16(*(*OpusT_opus_val16)(unsafe.Pointer(x + uintptr(i)*4)) * w)
			*(*OpusT_opus_val16)(unsafe.Pointer(xx + uintptr(n-i-int32(1))*4)) = OpusT_opus_val16(*(*OpusT_opus_val16)(unsafe.Pointer(x + uintptr(n-i-int32(1))*4)) * w)
			goto _30
		_30:
			;
			i = i + 1
		}
		xptr = xx
	}
	shift = 0
	Opus_celt_pitch_xcorr_c(tls, xptr, xptr, ac, fastN, lag+int32(1), arch)
	k = 0
	for {
		if !(k <= lag) {
			break
		}
		i = k + fastN
		d = libc.Float32FromInt32(0)
		for {
			if !(i < n) {
				break
			}
			d = d + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(xptr + uintptr(i)*4))**(*OpusT_opus_val16)(unsafe.Pointer(xptr + uintptr(i-k)*4)))
			goto _32
		_32:
			;
			i = i + 1
		}
		*(*OpusT_opus_val32)(unsafe.Pointer(ac + uintptr(k)*4)) += d
		goto _31
	_31:
		;
		k = k + 1
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _34
	_34:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _36
_36:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
	return shift
}

var trim_icdf19 = [11]uint8{
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
var spread_icdf19 = [4]uint8{
	0: uint8(25),
	1: uint8(23),
	2: uint8(2),
}
var tapset_icdf19 = [3]uint8{
	0: uint8(2),
	1: uint8(1),
}

var log2_x_norm_coeff19 = [8]float32{
	0: libc.Float32FromFloat32(1),
	1: libc.Float32FromFloat32(0.8888888955116272),
	2: libc.Float32FromFloat32(0.8),
	3: libc.Float32FromFloat32(0.7272727489471436),
	4: libc.Float32FromFloat32(0.6666666865348816),
	5: libc.Float32FromFloat32(0.6153846383094788),
	6: libc.Float32FromFloat32(0.5714285969734192),
	7: libc.Float32FromFloat32(0.5333333611488342),
}
var log2_y_norm_coeff19 = [8]float32{
	1: libc.Float32FromFloat32(0.1699250042438507),
	2: libc.Float32FromFloat32(0.32192808389663696),
	3: libc.Float32FromFloat32(0.45943161845207214),
	4: libc.Float32FromFloat32(0.5849624872207642),
	5: libc.Float32FromFloat32(0.7004396915435791),
	6: libc.Float32FromFloat32(0.8073549270629883),
	7: libc.Float32FromFloat32(0.9068905711174011),
}

// C documentation
//
//	/* prediction coefficients: 0.9, 0.8, 0.65, 0.5 */
var pred_coef = [4]OpusT_opus_val16{
	0: float32(libc.Float64FromInt32(29440) / libc.Float64FromFloat64(32768)),
	1: float32(libc.Float64FromInt32(26112) / libc.Float64FromFloat64(32768)),
	2: float32(libc.Float64FromInt32(21248) / libc.Float64FromFloat64(32768)),
	3: float32(libc.Float64FromInt32(16384) / libc.Float64FromFloat64(32768)),
}
var beta_coef = [4]OpusT_opus_val16{
	0: float32(libc.Float64FromInt32(30147) / libc.Float64FromFloat64(32768)),
	1: float32(libc.Float64FromInt32(22282) / libc.Float64FromFloat64(32768)),
	2: float32(libc.Float64FromInt32(12124) / libc.Float64FromFloat64(32768)),
	3: float32(libc.Float64FromInt32(6554) / libc.Float64FromFloat64(32768)),
}
var beta_intra = float32(libc.Float64FromInt32(4915) / libc.Float64FromFloat64(32768))

// C documentation
//
//	/*Parameters of the Laplace-like probability models used for the coarse energy.
//	  There is one pair of parameters for each frame size, prediction type
//	   (inter/intra), and band number.
//	  The first number of each pair is the probability of 0, and the second is the
//	   decay rate, both in Q8 precision.*/
var e_prob_model = [4][2][42]uint8{
	0: {
		0: {
			0:  uint8(72),
			1:  uint8(127),
			2:  uint8(65),
			3:  uint8(129),
			4:  uint8(66),
			5:  uint8(128),
			6:  uint8(65),
			7:  uint8(128),
			8:  uint8(64),
			9:  uint8(128),
			10: uint8(62),
			11: uint8(128),
			12: uint8(64),
			13: uint8(128),
			14: uint8(64),
			15: uint8(128),
			16: uint8(92),
			17: uint8(78),
			18: uint8(92),
			19: uint8(79),
			20: uint8(92),
			21: uint8(78),
			22: uint8(90),
			23: uint8(79),
			24: uint8(116),
			25: uint8(41),
			26: uint8(115),
			27: uint8(40),
			28: uint8(114),
			29: uint8(40),
			30: uint8(132),
			31: uint8(26),
			32: uint8(132),
			33: uint8(26),
			34: uint8(145),
			35: uint8(17),
			36: uint8(161),
			37: uint8(12),
			38: uint8(176),
			39: uint8(10),
			40: uint8(177),
			41: uint8(11),
		},
		1: {
			0:  uint8(24),
			1:  uint8(179),
			2:  uint8(48),
			3:  uint8(138),
			4:  uint8(54),
			5:  uint8(135),
			6:  uint8(54),
			7:  uint8(132),
			8:  uint8(53),
			9:  uint8(134),
			10: uint8(56),
			11: uint8(133),
			12: uint8(55),
			13: uint8(132),
			14: uint8(55),
			15: uint8(132),
			16: uint8(61),
			17: uint8(114),
			18: uint8(70),
			19: uint8(96),
			20: uint8(74),
			21: uint8(88),
			22: uint8(75),
			23: uint8(88),
			24: uint8(87),
			25: uint8(74),
			26: uint8(89),
			27: uint8(66),
			28: uint8(91),
			29: uint8(67),
			30: uint8(100),
			31: uint8(59),
			32: uint8(108),
			33: uint8(50),
			34: uint8(120),
			35: uint8(40),
			36: uint8(122),
			37: uint8(37),
			38: uint8(97),
			39: uint8(43),
			40: uint8(78),
			41: uint8(50),
		},
	},
	1: {
		0: {
			0:  uint8(83),
			1:  uint8(78),
			2:  uint8(84),
			3:  uint8(81),
			4:  uint8(88),
			5:  uint8(75),
			6:  uint8(86),
			7:  uint8(74),
			8:  uint8(87),
			9:  uint8(71),
			10: uint8(90),
			11: uint8(73),
			12: uint8(93),
			13: uint8(74),
			14: uint8(93),
			15: uint8(74),
			16: uint8(109),
			17: uint8(40),
			18: uint8(114),
			19: uint8(36),
			20: uint8(117),
			21: uint8(34),
			22: uint8(117),
			23: uint8(34),
			24: uint8(143),
			25: uint8(17),
			26: uint8(145),
			27: uint8(18),
			28: uint8(146),
			29: uint8(19),
			30: uint8(162),
			31: uint8(12),
			32: uint8(165),
			33: uint8(10),
			34: uint8(178),
			35: uint8(7),
			36: uint8(189),
			37: uint8(6),
			38: uint8(190),
			39: uint8(8),
			40: uint8(177),
			41: uint8(9),
		},
		1: {
			0:  uint8(23),
			1:  uint8(178),
			2:  uint8(54),
			3:  uint8(115),
			4:  uint8(63),
			5:  uint8(102),
			6:  uint8(66),
			7:  uint8(98),
			8:  uint8(69),
			9:  uint8(99),
			10: uint8(74),
			11: uint8(89),
			12: uint8(71),
			13: uint8(91),
			14: uint8(73),
			15: uint8(91),
			16: uint8(78),
			17: uint8(89),
			18: uint8(86),
			19: uint8(80),
			20: uint8(92),
			21: uint8(66),
			22: uint8(93),
			23: uint8(64),
			24: uint8(102),
			25: uint8(59),
			26: uint8(103),
			27: uint8(60),
			28: uint8(104),
			29: uint8(60),
			30: uint8(117),
			31: uint8(52),
			32: uint8(123),
			33: uint8(44),
			34: uint8(138),
			35: uint8(35),
			36: uint8(133),
			37: uint8(31),
			38: uint8(97),
			39: uint8(38),
			40: uint8(77),
			41: uint8(45),
		},
	},
	2: {
		0: {
			0:  uint8(61),
			1:  uint8(90),
			2:  uint8(93),
			3:  uint8(60),
			4:  uint8(105),
			5:  uint8(42),
			6:  uint8(107),
			7:  uint8(41),
			8:  uint8(110),
			9:  uint8(45),
			10: uint8(116),
			11: uint8(38),
			12: uint8(113),
			13: uint8(38),
			14: uint8(112),
			15: uint8(38),
			16: uint8(124),
			17: uint8(26),
			18: uint8(132),
			19: uint8(27),
			20: uint8(136),
			21: uint8(19),
			22: uint8(140),
			23: uint8(20),
			24: uint8(155),
			25: uint8(14),
			26: uint8(159),
			27: uint8(16),
			28: uint8(158),
			29: uint8(18),
			30: uint8(170),
			31: uint8(13),
			32: uint8(177),
			33: uint8(10),
			34: uint8(187),
			35: uint8(8),
			36: uint8(192),
			37: uint8(6),
			38: uint8(175),
			39: uint8(9),
			40: uint8(159),
			41: uint8(10),
		},
		1: {
			0:  uint8(21),
			1:  uint8(178),
			2:  uint8(59),
			3:  uint8(110),
			4:  uint8(71),
			5:  uint8(86),
			6:  uint8(75),
			7:  uint8(85),
			8:  uint8(84),
			9:  uint8(83),
			10: uint8(91),
			11: uint8(66),
			12: uint8(88),
			13: uint8(73),
			14: uint8(87),
			15: uint8(72),
			16: uint8(92),
			17: uint8(75),
			18: uint8(98),
			19: uint8(72),
			20: uint8(105),
			21: uint8(58),
			22: uint8(107),
			23: uint8(54),
			24: uint8(115),
			25: uint8(52),
			26: uint8(114),
			27: uint8(55),
			28: uint8(112),
			29: uint8(56),
			30: uint8(129),
			31: uint8(51),
			32: uint8(132),
			33: uint8(40),
			34: uint8(150),
			35: uint8(33),
			36: uint8(140),
			37: uint8(29),
			38: uint8(98),
			39: uint8(35),
			40: uint8(77),
			41: uint8(42),
		},
	},
	3: {
		0: {
			0:  uint8(42),
			1:  uint8(121),
			2:  uint8(96),
			3:  uint8(66),
			4:  uint8(108),
			5:  uint8(43),
			6:  uint8(111),
			7:  uint8(40),
			8:  uint8(117),
			9:  uint8(44),
			10: uint8(123),
			11: uint8(32),
			12: uint8(120),
			13: uint8(36),
			14: uint8(119),
			15: uint8(33),
			16: uint8(127),
			17: uint8(33),
			18: uint8(134),
			19: uint8(34),
			20: uint8(139),
			21: uint8(21),
			22: uint8(147),
			23: uint8(23),
			24: uint8(152),
			25: uint8(20),
			26: uint8(158),
			27: uint8(25),
			28: uint8(154),
			29: uint8(26),
			30: uint8(166),
			31: uint8(21),
			32: uint8(173),
			33: uint8(16),
			34: uint8(184),
			35: uint8(13),
			36: uint8(184),
			37: uint8(10),
			38: uint8(150),
			39: uint8(13),
			40: uint8(139),
			41: uint8(15),
		},
		1: {
			0:  uint8(22),
			1:  uint8(178),
			2:  uint8(63),
			3:  uint8(114),
			4:  uint8(74),
			5:  uint8(82),
			6:  uint8(84),
			7:  uint8(83),
			8:  uint8(92),
			9:  uint8(82),
			10: uint8(103),
			11: uint8(62),
			12: uint8(96),
			13: uint8(72),
			14: uint8(96),
			15: uint8(67),
			16: uint8(101),
			17: uint8(73),
			18: uint8(107),
			19: uint8(72),
			20: uint8(113),
			21: uint8(55),
			22: uint8(118),
			23: uint8(52),
			24: uint8(125),
			25: uint8(52),
			26: uint8(118),
			27: uint8(52),
			28: uint8(117),
			29: uint8(55),
			30: uint8(135),
			31: uint8(49),
			32: uint8(137),
			33: uint8(39),
			34: uint8(157),
			35: uint8(32),
			36: uint8(145),
			37: uint8(29),
			38: uint8(97),
			39: uint8(33),
			40: uint8(77),
			41: uint8(40),
		},
	},
}

var small_energy_icdf = [3]uint8{
	0: uint8(2),
	1: uint8(1),
}

func Opus__celt_lpc(tls *libc.TLS, _lpc uintptr, ac uintptr, p int32) {
	var error1, r, rr, tmp1, tmp2 OpusT_opus_val32
	var i, j int32
	var lpc uintptr
	_, _, _, _, _, _, _, _ = error1, i, j, lpc, r, rr, tmp1, tmp2
	error1 = *(*OpusT_opus_val32)(unsafe.Pointer(ac))
	lpc = _lpc
	libc.Xmemset(tls, lpc, 0, libc.Uint64FromInt32(p)*uint64(4))
	if *(*OpusT_opus_val32)(unsafe.Pointer(ac)) > libc.Float32FromFloat32(1e-10) {
		i = 0
		for {
			if !(i < p) {
				break
			}
			/* Sum up this iteration's reflection coefficient */
			rr = libc.Float32FromInt32(0)
			j = 0
			for {
				if !(j < i) {
					break
				}
				rr = rr + float32(*(*float32)(unsafe.Pointer(lpc + uintptr(j)*4))**(*OpusT_opus_val32)(unsafe.Pointer(ac + uintptr(i-j)*4)))
				goto _2
			_2:
				;
				j = j + 1
			}
			rr = rr + *(*OpusT_opus_val32)(unsafe.Pointer(ac + uintptr(i+int32(1))*4))
			r = -(rr / error1)
			/*  Update LPC coefficients and total error */
			*(*float32)(unsafe.Pointer(lpc + uintptr(i)*4)) = r
			j = 0
			for {
				if !(j < (i+int32(1))>>int32(1)) {
					break
				}
				tmp1 = *(*float32)(unsafe.Pointer(lpc + uintptr(j)*4))
				tmp2 = *(*float32)(unsafe.Pointer(lpc + uintptr(i-int32(1)-j)*4))
				*(*float32)(unsafe.Pointer(lpc + uintptr(j)*4)) = tmp1 + OpusT_opus_val32(r*tmp2)
				*(*float32)(unsafe.Pointer(lpc + uintptr(i-int32(1)-j)*4)) = tmp2 + OpusT_opus_val32(r*tmp1)
				goto _3
			_3:
				;
				j = j + 1
			}
			error1 = error1 - OpusT_opus_val32(OpusT_opus_val32(r*r)*error1)
			/* Bail out once we get 30 dB gain */
			if error1 <= OpusT_opus_val32(libc.Float32FromFloat32(0.001)**(*OpusT_opus_val32)(unsafe.Pointer(ac))) {
				break
			}
			goto _1
		_1:
			;
			i = i + 1
		}
	}
}

func Opus_celt_decode_with_ec(tls *libc.TLS, st uintptr, data uintptr, len1 int32, pcm uintptr, frame_size int32, dec uintptr, accum int32) (r int32) {
	return Opus_celt_decode_with_ec_dred(tls, st, data, len1, pcm, frame_size, dec, accum)
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
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _2
	_2:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _4
_4:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v6 = libc.Xmalloc(tls, uint64(16))
		goto _7
	_7:
		st = v6
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
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
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v15 = libc.Xmalloc(tls, uint64(16))
			goto _16
		_16:
			st = v15
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v17 = st
		goto _18
	_18:
		(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fscratch_ptr = v10
		v5 = v10
	} else {
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v19 = libc.Xmalloc(tls, uint64(16))
			goto _20
		_20:
			st = v19
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v21 = st
		goto _22
	_22:
		v5 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack
	}
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = v5
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _24
	_24:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
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
	if len1 < 0 || len1 > int32(1275) || pcm == libc.UintptrFromInt32(0) {
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
	if data == libc.UintptrFromInt32(0) || len1 <= int32(1) {
		celt_decode_lost(tls, st1, N, LM)
		deemphasis(tls, bp+56, pcm, N, CC, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdownsample, mode+16, st1+104, accum)
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			goto _31
		_31:
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
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
	if dec == libc.UintptrFromInt32(0) {
		Opus_ec_dec_init(tls, bp, data, libc.Uint32FromInt32(len1))
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
	v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (libc.Int32FromInt64(4)*libc.Int32FromInt32(__CHAR_BIT__) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
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
		v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (libc.Int32FromInt64(4)*libc.Int32FromInt32(__CHAR_BIT__) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
		goto _41
	_41:
		*(*int32)(unsafe.Pointer(dec + 24)) += tell - v28
	}
	postfilter_gain = libc.Float32FromInt32(0)
	postfilter_pitch = 0
	postfilter_tapset = 0
	if start == 0 && tell+int32(16) <= total_bits {
		if Opus_ec_dec_bit_logp(tls, dec, uint32(1)) != 0 {
			octave = libc.Int32FromUint32(Opus_ec_dec_uint(tls, dec, uint32(6)))
			postfilter_pitch = libc.Int32FromUint32(libc.Uint32FromInt32(libc.Int32FromInt32(16)<<octave) + Opus_ec_dec_bits(tls, dec, libc.Uint32FromInt32(int32(4)+octave)) - uint32(1))
			qg = libc.Int32FromUint32(Opus_ec_dec_bits(tls, dec, uint32(3)))
			v1 = dec
			v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (libc.Int32FromInt64(4)*libc.Int32FromInt32(__CHAR_BIT__) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
			goto _44
		_44:
			if v28+int32(2) <= total_bits {
				postfilter_tapset = Opus_ec_dec_icdf(tls, dec, uintptr(unsafe.Pointer(&tapset_icdf15)), uint32(2))
			}
			postfilter_gain = OpusT_opus_val16(libc.Float32FromFloat32(0.09375) * float32(qg+libc.Int32FromInt32(1)))
		}
		v1 = dec
		v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (libc.Int32FromInt64(4)*libc.Int32FromInt32(__CHAR_BIT__) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
		goto _47
	_47:
		tell = v28
	}
	if LM > 0 && tell+int32(3) <= total_bits {
		isTransient = Opus_ec_dec_bit_logp(tls, dec, uint32(3))
		v1 = dec
		v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (libc.Int32FromInt64(4)*libc.Int32FromInt32(__CHAR_BIT__) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
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
			safety = libc.Float32FromInt32(0)
			if int32(10) < (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Floss_duration>>LM {
				v37 = int32(10)
			} else {
				v37 = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Floss_duration >> LM
			}
			missing = v37
			if LM == 0 {
				safety = libc.Float32FromFloat32(1.5)
			} else {
				if LM == int32(1) {
					safety = libc.Float32FromFloat32(0.5)
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
					if E1-E0 > float32(libc.Float32FromFloat32(0.5)*(E2-E0)) {
						v57 = E1 - E0
					} else {
						v57 = float32(libc.Float32FromFloat32(0.5) * (E2 - E0))
					}
					slope = v57
					if slope < libc.Float32FromFloat32(2) {
						v57 = slope
					} else {
						v57 = libc.Float32FromFloat32(2)
					}
					slope = v57
					if float32(libc.Int32FromInt32(0)) > OpusT_opus_val32(float32(libc.Int32FromInt32(1)+missing)*slope) {
						v57 = float32(libc.Int32FromInt32(0))
					} else {
						v57 = OpusT_opus_val32(float32(libc.Int32FromInt32(1)+missing) * slope)
					}
					E0 = E0 - v57
					if -libc.Float32FromFloat32(20) > E0 {
						v60 = -libc.Float32FromFloat32(20)
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
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _65
	_65:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _67
_67:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _69
	_69:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _71
_71:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _73
	_73:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _75
_75:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _77
	_77:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _79
_79:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4420, int32(1395))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _81
	_81:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _83
_83:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(nbEBands) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _85
	_85:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _87
_87:
	tf_res = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	tf_decode(tls, start, end, isTransient, tf_res, LM, dec)
	v1 = dec
	v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (libc.Int32FromInt64(4)*libc.Int32FromInt32(__CHAR_BIT__) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
	goto _90
_90:
	tell = v28
	spread_decision = int32(SPREAD_NORMAL)
	if tell+int32(4) <= total_bits {
		spread_decision = Opus_ec_dec_icdf(tls, dec, uintptr(unsafe.Pointer(&spread_icdf15)), uint32(5))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _92
	_92:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _94
_94:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _96
	_96:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _98
_98:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _100
	_100:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _102
_102:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _104
	_104:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _106
_106:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4420, int32(1403))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _108
	_108:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _110
_110:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(nbEBands) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _112
	_112:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _114
_114:
	cap1 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	Opus_init_caps(tls, mode, cap1, LM, C)
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _116
	_116:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _118
_118:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _120
	_120:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _122
_122:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _124
	_124:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _126
_126:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _128
	_128:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _130
_130:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4420, int32(1407))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _132
	_132:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _134
_134:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(nbEBands) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _136
	_136:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _138
_138:
	offsets = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	dynalloc_logp = int32(6)
	total_bits = total_bits << int32(BITRES)
	tell = libc.Int32FromUint32(Opus_ec_tell_frac(tls, dec))
	i = start
	for {
		if !(i < end) {
			break
		}
		width = C * (int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i+int32(1))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i)*2)))) << LM
		/* quanta is 6 bits, but no more than 1 bit/sample
		   and no less than 1/8 bit/sample */
		if libc.Int32FromInt32(6)<<libc.Int32FromInt32(BITRES) > width {
			v37 = libc.Int32FromInt32(6) << libc.Int32FromInt32(BITRES)
		} else {
			v37 = width
		}
		if width<<int32(BITRES) < v37 {
			v28 = width << int32(BITRES)
		} else {
			if libc.Int32FromInt32(6)<<libc.Int32FromInt32(BITRES) > width {
				v40 = libc.Int32FromInt32(6) << libc.Int32FromInt32(BITRES)
			} else {
				v40 = width
			}
			v28 = v40
		}
		quanta = v28
		dynalloc_loop_logp = dynalloc_logp
		boost = 0
		for tell+dynalloc_loop_logp<<int32(BITRES) < total_bits && boost < *(*int32)(unsafe.Pointer(cap1 + uintptr(i)*4)) {
			flag = Opus_ec_dec_bit_logp(tls, dec, libc.Uint32FromInt32(dynalloc_loop_logp))
			tell = libc.Int32FromUint32(Opus_ec_tell_frac(tls, dec))
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
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _145
	_145:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _147
_147:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _149
	_149:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _151
_151:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _153
	_153:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _155
_155:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _157
	_157:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _159
_159:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4420, int32(1440))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _161
	_161:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _163
_163:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(nbEBands) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _165
	_165:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _167
_167:
	fine_quant = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	if tell+libc.Int32FromInt32(6)<<libc.Int32FromInt32(BITRES) <= total_bits {
		v28 = Opus_ec_dec_icdf(tls, dec, uintptr(unsafe.Pointer(&trim_icdf15)), uint32(7))
	} else {
		v28 = int32(5)
	}
	alloc_trim = v28
	bits = len1*int32(8)<<int32(BITRES) - libc.Int32FromUint32(Opus_ec_tell_frac(tls, dec)) - int32(1)
	if isTransient != 0 && LM >= int32(2) && bits >= (LM+int32(2))<<int32(BITRES) {
		v28 = libc.Int32FromInt32(1) << libc.Int32FromInt32(BITRES)
	} else {
		v28 = 0
	}
	anti_collapse_rsv = v28
	bits = bits - anti_collapse_rsv
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _171
	_171:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _173
_173:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _175
	_175:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _177
_177:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _179
	_179:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _181
_181:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _183
	_183:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _185
_185:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4420, int32(1448))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _187
	_187:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _189
_189:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(nbEBands) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _191
	_191:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _193
_193:
	pulses = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _195
	_195:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _197
_197:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _199
	_199:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _201
_201:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _203
	_203:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _205
_205:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _207
	_207:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _209
_209:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4420, int32(1449))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _211
	_211:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _213
_213:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(nbEBands) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _215
	_215:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _217
_217:
	fine_priority = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	codedBands = Opus_clt_compute_allocation(tls, mode, start, end, offsets, cap1, alloc_trim, bp+72, bp+76, bits, bp+80, pulses, fine_quant, fine_priority, C, LM, dec, 0, 0, 0)
	Opus_unquant_fine_energy(tls, mode, start, end, oldBandE, libc.UintptrFromInt32(0), fine_quant, dec, C)
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _219
	_219:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _221
_221:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _223
	_223:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _225
_225:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _227
	_227:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _229
_229:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _231
	_231:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _233
_233:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(C*N)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4420, int32(1457))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _235
	_235:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _237
_237:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(C*N) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _239
	_239:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _241
_241:
	X = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(C*N)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1))) /**< Interleaved normalised MDCTs */
	c = 0
	for {
		libc.Xmemmove(tls, decode_mem[c], decode_mem[c]+uintptr(N)*4, libc.Uint64FromInt32(decode_buffer_size-N+overlap)*uint64(4)+libc.Uint64FromInt64(0*((int64(decode_mem[c])-int64(decode_mem[c]+uintptr(N)*4))/4)))
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
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _245
	_245:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _247
_247:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _249
	_249:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _251
_251:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(1) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(1) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _253
	_253:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _255
_255:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _257
	_257:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _259
_259:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(C*nbEBands)*(libc.Uint64FromInt64(1)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4420, int32(1487))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _261
	_261:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _263
_263:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(C*nbEBands) * (libc.Uint64FromInt64(1) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _265
	_265:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _267
_267:
	collapse_masks = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(C*nbEBands)*(libc.Uint64FromInt64(1)/libc.Uint64FromInt64(1)))
	if C == int32(2) {
		v1 = X + uintptr(N)*4
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	Opus_quant_all_bands(tls, 0, mode, start, end, X, v1, collapse_masks, libc.UintptrFromInt32(0), pulses, shortBlocks, spread_decision, *(*int32)(unsafe.Pointer(bp + 76)), *(*int32)(unsafe.Pointer(bp + 72)), tf_res, len1*(libc.Int32FromInt32(8)<<libc.Int32FromInt32(BITRES))-anti_collapse_rsv, *(*OpusT_opus_int32)(unsafe.Pointer(bp + 80)), dec, LM, codedBands, st1+48, 0, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Fdisable_inv)
	if anti_collapse_rsv > 0 {
		anti_collapse_on = libc.Int32FromUint32(Opus_ec_dec_bits(tls, dec, uint32(1)))
	}
	v1 = dec
	v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (libc.Int32FromInt64(4)*libc.Int32FromInt32(__CHAR_BIT__) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
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
			*(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(i)*4)) = -libc.Float32FromFloat32(28)
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
		libc.Xmemcpy(tls, oldBandE+uintptr(nbEBands)*4, oldBandE, libc.Uint64FromInt32(nbEBands)*uint64(4)+libc.Uint64FromInt64(0*((OpusT___predefined_ptrdiff_t(oldBandE+uintptr(nbEBands)*4)-int64(oldBandE))/4)))
	}
	if !(isTransient != 0) {
		libc.Xmemcpy(tls, oldLogE2, oldLogE, libc.Uint64FromInt32(libc.Int32FromInt32(2)*nbEBands)*uint64(4)+libc.Uint64FromInt64(0*((int64(oldLogE2)-int64(oldLogE))/4)))
		libc.Xmemcpy(tls, oldLogE, oldBandE, libc.Uint64FromInt32(libc.Int32FromInt32(2)*nbEBands)*uint64(4)+libc.Uint64FromInt64(0*((int64(oldLogE)-int64(oldBandE))/4)))
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
	max_background_increase = OpusT_celt_glog(float32(v28) * libc.Float32FromFloat32(0.001))
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
			*(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4)) = libc.Float32FromInt32(0)
			v35 = -libc.Float32FromFloat32(28)
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
			*(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i)*4)) = libc.Float32FromInt32(0)
			v35 = -libc.Float32FromFloat32(28)
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
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _289
	_289:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _291
_291:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
	v1 = dec
	v28 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (libc.Int32FromInt64(4)*libc.Int32FromInt32(__CHAR_BIT__) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
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

func Opus_celt_decoder_get_size(tls *libc.TLS, channels int32) (r int32) {
	var mode uintptr
	_ = mode
	mode = Opus_opus_custom_mode_create(tls, int32(48000), int32(960), libc.UintptrFromInt32(0))
	return opus_custom_decoder_get_size(tls, mode, channels)
}

func Opus_celt_decoder_init(tls *libc.TLS, st uintptr, sampling_rate OpusT_opus_int32, channels int32) (r int32) {
	var ret int32
	_ = ret
	ret = opus_custom_decoder_init(tls, st, Opus_opus_custom_mode_create(tls, int32(48000), int32(960), libc.UintptrFromInt32(0)), channels)
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

func Opus_celt_encode_with_ec(tls *libc.TLS, st1 uintptr, pcm uintptr, frame_size1 int32, compressed uintptr, nbCompressedBytes int32, enc uintptr) (r int32) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var C, CC, LM, M, N, adjust, alloc_trim, allow_weak_transients, anti_collapse_on, anti_collapse_rsv, boost, c, codedBands, count, count_dynalloc, dynalloc_logp, dynalloc_loop_logp, effEnd, effectiveBytes, enable_tf_analysis, enabled, end, flag, hybrid, i, i1, isTransient, j, lambda, lm_diff, mask_end, midband, min_bandwidth, nbAvailableBytes, nbEBands, nbFilledBytes, need_clip, octave, overlap, packet_size_cap, pf_on, pitch_change, prefilter_tapset, qext_bytes, quanta, secondMdct, shortBlocks, signalBandwidth, silence, start, tf_select, transient_got_disabled, width, v38, v40, v43, v44, v45, v46, v47 int32
	var X, _saved_stack, bandE, bandLogE, bandLogE2, cap1, collapse_masks, eBands, energyError, error1, fine_priority, fine_quant, freq, importance, in1, mode, offsets, oldBandE, oldLogE, oldLogE2, prefilter_mem, pulses, spread_weight, st, surround_dynalloc, tf_res, v1, v10, v11, v13, v15, v17, v19, v21, v3, v5, v6, v8 uintptr
	var alpha, mask16, maxval, minval, tone_freq, v81, v82, v85, v88, v89, v92 OpusT_opus_val16
	var base_target, bits, delta, equiv_rate, max_allowed, min_allowed, target, tell, tell0_frac, tmp, total_bits, total_boost, vbr_bound, vbr_rate, v36 OpusT_opus_int32
	var diff, frame_avg, lin, mask_avg, sample_max, v78, v83, v90 OpusT_opus_val32
	var follow, mask, maxDepth, offset, surround_masking, surround_trim, temporal_vbr, unmask, v247, v248, v249 OpusT_celt_glog
	var v112, v215 bool
	var v218 OpusT_celt_ener
	var v259, v263, v264 float32
	var _ /* _enc at bp+8 */ OpusT_ec_enc
	var _ /* balance at bp+76 */ OpusT_opus_int32
	var _ /* dual_stereo at bp+72 */ int32
	var _ /* gain1 at bp+68 */ OpusT_opus_val16
	var _ /* in at bp+0 */ struct {
		Fi [0]OpusT_opus_uint32
		Ff float32
	}
	var _ /* pitch_index at bp+64 */ int32
	var _ /* qg at bp+100 */ int32
	var _ /* tf_chan at bp+80 */ int32
	var _ /* tf_estimate at bp+84 */ OpusT_opus_val16
	var _ /* toneishness at bp+96 */ OpusT_opus_val32
	var _ /* tot_boost at bp+88 */ OpusT_opus_int32
	var _ /* weak_transient at bp+92 */ int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, CC, LM, M, N, X, _saved_stack, adjust, alloc_trim, allow_weak_transients, alpha, anti_collapse_on, anti_collapse_rsv, bandE, bandLogE, bandLogE2, base_target, bits, boost, c, cap1, codedBands, collapse_masks, count, count_dynalloc, delta, diff, dynalloc_logp, dynalloc_loop_logp, eBands, effEnd, effectiveBytes, enable_tf_analysis, enabled, end, energyError, equiv_rate, error1, fine_priority, fine_quant, flag, follow, frame_avg, freq, hybrid, i, i1, importance, in1, isTransient, j, lambda, lin, lm_diff, mask, mask16, mask_avg, mask_end, maxDepth, max_allowed, maxval, midband, min_allowed, min_bandwidth, minval, mode, nbAvailableBytes, nbEBands, nbFilledBytes, need_clip, octave, offset, offsets, oldBandE, oldLogE, oldLogE2, overlap, packet_size_cap, pf_on, pitch_change, prefilter_mem, prefilter_tapset, pulses, qext_bytes, quanta, sample_max, secondMdct, shortBlocks, signalBandwidth, silence, spread_weight, st, start, surround_dynalloc, surround_masking, surround_trim, target, tell, tell0_frac, temporal_vbr, tf_res, tf_select, tmp, tone_freq, total_bits, total_boost, transient_got_disabled, unmask, vbr_bound, vbr_rate, width, v1, v10, v11, v112, v13, v15, v17, v19, v21, v215, v218, v247, v248, v249, v259, v263, v264, v3, v36, v38, v40, v43, v44, v45, v46, v47, v5, v6, v78, v8, v81, v82, v83, v85, v88, v89, v90, v92
	shortBlocks = 0
	isTransient = 0
	CC = (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fchannels
	C = (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fstream_channels
	*(*int32)(unsafe.Pointer(bp + 64)) = int32(COMBFILTER_MINPERIOD)
	*(*OpusT_opus_val16)(unsafe.Pointer(bp + 68)) = libc.Float32FromInt32(0)
	*(*int32)(unsafe.Pointer(bp + 72)) = 0
	prefilter_tapset = 0
	anti_collapse_on = 0
	silence = 0
	*(*int32)(unsafe.Pointer(bp + 80)) = 0
	pitch_change = 0
	transient_got_disabled = 0
	surround_masking = libc.Float32FromInt32(0)
	temporal_vbr = libc.Float32FromInt32(0)
	surround_trim = libc.Float32FromInt32(0)
	*(*int32)(unsafe.Pointer(bp + 92)) = 0
	tone_freq = float32(-libc.Int32FromInt32(1))
	*(*OpusT_opus_val32)(unsafe.Pointer(bp + 96)) = libc.Float32FromInt32(0)
	qext_bytes = 0
	packet_size_cap = int32(1275)
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _2
	_2:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _4
_4:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v6 = libc.Xmalloc(tls, uint64(16))
		goto _7
	_7:
		st = v6
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
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
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v15 = libc.Xmalloc(tls, uint64(16))
			goto _16
		_16:
			st = v15
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v17 = st
		goto _18
	_18:
		(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v17)).Fscratch_ptr = v10
		v5 = v10
	} else {
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v19 = libc.Xmalloc(tls, uint64(16))
			goto _20
		_20:
			st = v19
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v21 = st
		goto _22
	_22:
		v5 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack
	}
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = v5
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _24
	_24:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _26
_26:
	_saved_stack = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack
	mode = (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fmode
	nbEBands = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FnbEBands
	overlap = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Foverlap
	eBands = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FeBands
	start = (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fstart
	end = (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fend
	hybrid = libc.BoolInt32(start != 0)
	*(*OpusT_opus_val16)(unsafe.Pointer(bp + 84)) = libc.Float32FromInt32(0)
	if nbCompressedBytes < int32(2) || pcm == libc.UintptrFromInt32(0) {
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			goto _28
		_28:
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v3 = st
		goto _30
	_30:
		(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
		return -int32(1)
	}
	frame_size1 = frame_size1 * (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fupsample
	LM = 0
	for {
		if !(LM <= (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FmaxLM) {
			break
		}
		if (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize<<LM == frame_size1 {
			break
		}
		goto _31
	_31:
		;
		LM = LM + 1
	}
	if LM > (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FmaxLM {
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			goto _33
		_33:
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v3 = st
		goto _35
	_35:
		(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
		return -int32(1)
	}
	M = int32(1) << LM
	N = M * (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize
	prefilter_mem = st1 + 252 + uintptr(CC*overlap)*4
	oldBandE = st1 + 252 + uintptr(CC*(overlap+libc.Int32FromInt32(COMBFILTER_MAXPERIOD)))*4
	oldLogE = oldBandE + uintptr(CC*nbEBands)*4
	oldLogE2 = oldLogE + uintptr(CC*nbEBands)*4
	energyError = oldLogE2 + uintptr(CC*nbEBands)*4
	if enc == libc.UintptrFromInt32(0) {
		v36 = libc.Int32FromInt32(1)
		tell = v36
		tell0_frac = v36
		nbFilledBytes = 0
	} else {
		tell0_frac = libc.Int32FromUint32(Opus_ec_tell_frac(tls, enc))
		v1 = enc
		v38 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (libc.Int32FromInt64(4)*libc.Int32FromInt32(__CHAR_BIT__) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
		goto _39
	_39:
		tell = v38
		nbFilledBytes = (tell + int32(4)) >> int32(3)
	}
	if !((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fsignalling == libc.Int32FromInt32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+4213, __ccgo_ts+4084, int32(1895))
	}
	/* Can't produce more than 1275 output bytes for the main payload, plus any QEXT extra data. */
	if nbCompressedBytes < packet_size_cap {
		v38 = nbCompressedBytes
	} else {
		v38 = packet_size_cap
	}
	nbCompressedBytes = v38
	if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fvbr != 0 && (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fbitrate != -int32(1) {
		v36 = (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fbitrate * int32(6) / (int32(6) * (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FFs / frame_size1)
		goto _42
	_42:
		vbr_rate = v36 << int32(BITRES)
		effectiveBytes = vbr_rate >> (libc.Int32FromInt32(3) + libc.Int32FromInt32(BITRES))
	} else {
		vbr_rate = 0
		tmp = (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fbitrate * frame_size1
		if tell > int32(1) {
			tmp = tmp + tell*(*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FFs
		}
		if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fbitrate != -int32(1) {
			if nbCompressedBytes < (tmp+int32(4)*(*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FFs)/(int32(8)*(*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FFs)-libc.BoolInt32(!!((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fsignalling != 0)) {
				v40 = nbCompressedBytes
			} else {
				v40 = (tmp+int32(4)*(*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FFs)/(int32(8)*(*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FFs) - libc.BoolInt32(!!((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fsignalling != 0))
			}
			if int32(2) > v40 {
				v38 = int32(2)
			} else {
				if nbCompressedBytes < (tmp+int32(4)*(*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FFs)/(int32(8)*(*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FFs)-libc.BoolInt32(!!((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fsignalling != 0)) {
					v43 = nbCompressedBytes
				} else {
					v43 = (tmp+int32(4)*(*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FFs)/(int32(8)*(*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FFs) - libc.BoolInt32(!!((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fsignalling != 0))
				}
				v38 = v43
			}
			nbCompressedBytes = v38
			if enc != libc.UintptrFromInt32(0) {
				Opus_ec_enc_shrink(tls, enc, libc.Uint32FromInt32(nbCompressedBytes))
			}
		}
		effectiveBytes = nbCompressedBytes - nbFilledBytes
	}
	nbAvailableBytes = nbCompressedBytes - nbFilledBytes
	equiv_rate = nbCompressedBytes*int32(8)*int32(50)<<(int32(3)-LM) - (int32(40)*C+int32(20))*(int32(400)>>LM-int32(50))
	if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fbitrate != -int32(1) {
		if equiv_rate < (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fbitrate-(int32(40)*C+int32(20))*(int32(400)>>LM-int32(50)) {
			v38 = equiv_rate
		} else {
			v38 = (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fbitrate - (int32(40)*C+int32(20))*(int32(400)>>LM-int32(50))
		}
		equiv_rate = v38
	}
	if enc == libc.UintptrFromInt32(0) {
		Opus_ec_enc_init(tls, bp+8, compressed, libc.Uint32FromInt32(nbCompressedBytes))
		enc = bp + 8
	}
	if vbr_rate > 0 {
		/* Computes the max bit-rate allowed in VBR mode to avoid violating the
		    target rate and buffering.
		   We must do this up front so that bust-prevention logic triggers
		    correctly if we don't have enough bits. */
		if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fconstrained_vbr != 0 {
			/* We could use any multiple of vbr_rate as bound (depending on the
			    delay).
			   This is clamped to ensure we use at least two bytes if the encoder
			    was entirely empty, but to allow 0 in hybrid mode. */
			vbr_bound = vbr_rate
			if tell == int32(1) {
				v43 = int32(2)
			} else {
				v43 = 0
			}
			if v43 > (vbr_rate+vbr_bound-(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fvbr_reservoir)>>(libc.Int32FromInt32(BITRES)+libc.Int32FromInt32(3)) {
				if tell == int32(1) {
					v44 = int32(2)
				} else {
					v44 = 0
				}
				v40 = v44
			} else {
				v40 = (vbr_rate + vbr_bound - (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fvbr_reservoir) >> (libc.Int32FromInt32(BITRES) + libc.Int32FromInt32(3))
			}
			if v40 < nbAvailableBytes {
				if tell == int32(1) {
					v46 = int32(2)
				} else {
					v46 = 0
				}
				if v46 > (vbr_rate+vbr_bound-(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fvbr_reservoir)>>(libc.Int32FromInt32(BITRES)+libc.Int32FromInt32(3)) {
					if tell == int32(1) {
						v47 = int32(2)
					} else {
						v47 = 0
					}
					v45 = v47
				} else {
					v45 = (vbr_rate + vbr_bound - (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fvbr_reservoir) >> (libc.Int32FromInt32(BITRES) + libc.Int32FromInt32(3))
				}
				v38 = v45
			} else {
				v38 = nbAvailableBytes
			}
			max_allowed = v38
			if max_allowed < nbAvailableBytes {
				nbCompressedBytes = nbFilledBytes + max_allowed
				nbAvailableBytes = max_allowed
				Opus_ec_enc_shrink(tls, enc, libc.Uint32FromInt32(nbCompressedBytes))
			}
		}
	}
	total_bits = nbCompressedBytes * int32(8)
	effEnd = end
	if effEnd > (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FeffEBands {
		effEnd = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FeffEBands
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _55
	_55:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _57
_57:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _59
	_59:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _61
_61:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _63
	_63:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _65
_65:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _67
	_67:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _69
_69:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(CC*(N+overlap))*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4084, int32(1967))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _71
	_71:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _73
_73:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(CC*(N+overlap)) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _75
	_75:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _77
_77:
	in1 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(CC*(N+overlap))*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	v1 = pcm
	maxval = libc.Float32FromInt32(0)
	minval = libc.Float32FromInt32(0)
	i = libc.Int32FromInt32(0)
	for {
		if !(i < C*(N-overlap)/(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fupsample) {
			break
		}
		if maxval > *(*OpusT_opus_val16)(unsafe.Pointer(v1 + uintptr(i)*4)) {
			v81 = maxval
		} else {
			v81 = *(*OpusT_opus_val16)(unsafe.Pointer(v1 + uintptr(i)*4))
		}
		maxval = v81
		if minval < *(*OpusT_opus_val16)(unsafe.Pointer(v1 + uintptr(i)*4)) {
			v82 = minval
		} else {
			v82 = *(*OpusT_opus_val16)(unsafe.Pointer(v1 + uintptr(i)*4))
		}
		minval = v82
		goto _80
	_80:
		;
		i = i + 1
	}
	if maxval > -minval {
		v85 = maxval
	} else {
		v85 = -minval
	}
	v83 = v85
	goto _84
_84:
	if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Foverlap_max > v83 {
		v78 = (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Foverlap_max
	} else {
		v3 = pcm
		maxval = libc.Float32FromInt32(0)
		minval = libc.Float32FromInt32(0)
		i = libc.Int32FromInt32(0)
		for {
			if !(i < C*(N-overlap)/(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fupsample) {
				break
			}
			if maxval > *(*OpusT_opus_val16)(unsafe.Pointer(v3 + uintptr(i)*4)) {
				v88 = maxval
			} else {
				v88 = *(*OpusT_opus_val16)(unsafe.Pointer(v3 + uintptr(i)*4))
			}
			maxval = v88
			if minval < *(*OpusT_opus_val16)(unsafe.Pointer(v3 + uintptr(i)*4)) {
				v89 = minval
			} else {
				v89 = *(*OpusT_opus_val16)(unsafe.Pointer(v3 + uintptr(i)*4))
			}
			minval = v89
			goto _87
		_87:
			;
			i = i + 1
		}
		if maxval > -minval {
			v92 = maxval
		} else {
			v92 = -minval
		}
		v90 = v92
		goto _91
	_91:
		v78 = v90
	}
	sample_max = v78
	v1 = pcm + uintptr(C*(N-overlap)/(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fupsample)*4
	maxval = libc.Float32FromInt32(0)
	minval = libc.Float32FromInt32(0)
	i = libc.Int32FromInt32(0)
	for {
		if !(i < C*overlap/(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fupsample) {
			break
		}
		if maxval > *(*OpusT_opus_val16)(unsafe.Pointer(v1 + uintptr(i)*4)) {
			v81 = maxval
		} else {
			v81 = *(*OpusT_opus_val16)(unsafe.Pointer(v1 + uintptr(i)*4))
		}
		maxval = v81
		if minval < *(*OpusT_opus_val16)(unsafe.Pointer(v1 + uintptr(i)*4)) {
			v82 = minval
		} else {
			v82 = *(*OpusT_opus_val16)(unsafe.Pointer(v1 + uintptr(i)*4))
		}
		minval = v82
		goto _94
	_94:
		;
		i = i + 1
	}
	if maxval > -minval {
		v85 = maxval
	} else {
		v85 = -minval
	}
	v78 = v85
	goto _98
_98:
	(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Foverlap_max = v78
	if sample_max > (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Foverlap_max {
		v78 = sample_max
	} else {
		v78 = (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Foverlap_max
	}
	sample_max = v78
	silence = libc.BoolInt32(sample_max <= libc.Float32FromInt32(1)/float32(libc.Int32FromInt32(1)<<(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Flsb_depth))
	if tell == int32(1) {
		Opus_ec_enc_bit_logp(tls, enc, silence, uint32(15))
	} else {
		silence = 0
	}
	if silence != 0 {
		/*In VBR mode there is no need to send more than the minimum. */
		if vbr_rate > 0 {
			if nbCompressedBytes < nbFilledBytes+int32(2) {
				v40 = nbCompressedBytes
			} else {
				v40 = nbFilledBytes + int32(2)
			}
			v38 = v40
			nbCompressedBytes = v38
			effectiveBytes = v38
			total_bits = nbCompressedBytes * int32(8)
			nbAvailableBytes = int32(2)
			Opus_ec_enc_shrink(tls, enc, libc.Uint32FromInt32(nbCompressedBytes))
		}
		/* Pretend we've filled all the remaining bits with zeros
		   (that's what the initialiser did anyway) */
		tell = nbCompressedBytes * int32(8)
		v1 = enc
		v38 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (libc.Int32FromInt64(4)*libc.Int32FromInt32(__CHAR_BIT__) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
		goto _105
	_105:
		*(*int32)(unsafe.Pointer(enc + 24)) += tell - v38
	}
	c = 0
	for {
		need_clip = 0
		need_clip = libc.BoolInt32((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fclip != 0 && sample_max > libc.Float32FromFloat32(65536))
		Opus_celt_preemphasis(tls, pcm+uintptr(c)*4, in1+uintptr(c*(N+overlap))*4+uintptr(overlap)*4, N, CC, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fupsample, mode+16, st1+196+uintptr(c)*4, need_clip)
		libc.Xmemcpy(tls, in1+uintptr(c*(N+overlap))*4, prefilter_mem+uintptr((int32(1)+c)*int32(COMBFILTER_MAXPERIOD)-overlap)*4, libc.Uint64FromInt32(overlap)*uint64(4)+libc.Uint64FromInt64(0*((int64(in1+uintptr(c*(N+overlap))*4)-OpusT___predefined_ptrdiff_t(prefilter_mem+uintptr((int32(1)+c)*int32(COMBFILTER_MAXPERIOD)-overlap)*4))/4)))
		goto _107
	_107:
		;
		c = c + 1
		v38 = c
		if !(v38 < CC) {
			break
		}
	}
	tone_freq = tone_detect(tls, in1, CC, N+overlap, bp+96, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FFs)
	isTransient = 0
	shortBlocks = 0
	if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fcomplexity >= int32(1) && !((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Flfe != 0) {
		/* Reduces the likelihood of energy instability on fricatives at low bitrate
		   in hybrid mode. It seems like we still want to have real transients on vowels
		   though (small SILK quantization offset value). */
		allow_weak_transients = libc.BoolInt32(hybrid != 0 && effectiveBytes < int32(15) && (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fsilk_info.FsignalType != int32(2))
		isTransient = transient_analysis(tls, in1, N+overlap, CC, bp+84, bp+80, allow_weak_transients, bp+92, tone_freq, *(*OpusT_opus_val32)(unsafe.Pointer(bp + 96)))
	}
	if *(*OpusT_opus_val32)(unsafe.Pointer(bp + 96)) < libc.Float32FromFloat32(1)-*(*OpusT_opus_val16)(unsafe.Pointer(bp + 84)) {
		v78 = *(*OpusT_opus_val32)(unsafe.Pointer(bp + 96))
	} else {
		v78 = libc.Float32FromFloat32(1) - *(*OpusT_opus_val16)(unsafe.Pointer(bp + 84))
	}
	*(*OpusT_opus_val32)(unsafe.Pointer(bp + 96)) = v78
	/* Find pitch period and gain */
	enabled = libc.BoolInt32(((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Flfe != 0 && nbAvailableBytes > int32(3) || nbAvailableBytes > int32(12)*C) && !(hybrid != 0) && !(silence != 0) && tell+int32(16) <= total_bits && !((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fdisable_pf != 0))
	prefilter_tapset = (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Ftapset_decision
	pf_on = run_prefilter(tls, st1, in1, prefilter_mem, CC, N, prefilter_tapset, bp+64, bp+68, bp+100, enabled, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fcomplexity, *(*OpusT_opus_val16)(unsafe.Pointer(bp + 84)), nbAvailableBytes, st1+124, tone_freq, *(*OpusT_opus_val32)(unsafe.Pointer(bp + 96)))
	if (*(*OpusT_opus_val16)(unsafe.Pointer(bp + 68)) > libc.Float32FromFloat32(0.4) || (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fprefilter_gain > libc.Float32FromFloat32(0.4)) && (!((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fanalysis.Fvalid != 0) || float64((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fanalysis.Ftonality) > float64(0.3)) && (float64(*(*int32)(unsafe.Pointer(bp + 64))) > float64(float64(1.26)*float64((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fprefilter_period)) || float64(*(*int32)(unsafe.Pointer(bp + 64))) < float64(float64(0.79)*float64((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fprefilter_period))) {
		pitch_change = int32(1)
	}
	if pf_on == 0 {
		if !(hybrid != 0) && tell+int32(16) <= total_bits {
			Opus_ec_enc_bit_logp(tls, enc, 0, uint32(1))
		}
	} else {
		Opus_ec_enc_bit_logp(tls, enc, int32(1), uint32(1))
		*(*int32)(unsafe.Pointer(bp + 64)) = *(*int32)(unsafe.Pointer(bp + 64)) + int32(1)
		octave = libc.Int32FromInt64(4)*libc.Int32FromInt32(__CHAR_BIT__) - libc.X__builtin_clz(tls, libc.Uint32FromInt32(*(*int32)(unsafe.Pointer(bp + 64)))) - int32(5)
		Opus_ec_enc_uint(tls, enc, libc.Uint32FromInt32(octave), uint32(6))
		Opus_ec_enc_bits(tls, enc, libc.Uint32FromInt32(*(*int32)(unsafe.Pointer(bp + 64))-int32(16)<<octave), libc.Uint32FromInt32(int32(4)+octave))
		*(*int32)(unsafe.Pointer(bp + 64)) = *(*int32)(unsafe.Pointer(bp + 64)) - int32(1)
		Opus_ec_enc_bits(tls, enc, libc.Uint32FromInt32(*(*int32)(unsafe.Pointer(bp + 100))), uint32(3))
		Opus_ec_enc_icdf(tls, enc, prefilter_tapset, uintptr(unsafe.Pointer(&tapset_icdf14)), uint32(2))
	}
	if v112 = LM > 0; v112 {
		v1 = enc
		v38 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (libc.Int32FromInt64(4)*libc.Int32FromInt32(__CHAR_BIT__) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
		goto _111
	_111:
	}
	if v112 && v38+int32(3) <= total_bits {
		if isTransient != 0 {
			shortBlocks = M
		}
	} else {
		isTransient = 0
		transient_got_disabled = int32(1)
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _114
	_114:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _116
_116:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _118
	_118:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _120
_120:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _122
	_122:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _124
_124:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _126
	_126:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _128
_128:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(CC*N)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4084, int32(2072))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _130
	_130:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _132
_132:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(CC*N) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _134
	_134:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _136
_136:
	freq = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(CC*N)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1))) /**< Interleaved signal MDCTs */
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _138
	_138:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _140
_140:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _142
	_142:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _144
_144:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _146
	_146:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _148
_148:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _150
	_150:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _152
_152:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(nbEBands*CC)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4084, int32(2073))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _154
	_154:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _156
_156:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(nbEBands*CC) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _158
	_158:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _160
_160:
	bandE = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(nbEBands*CC)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _162
	_162:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _164
_164:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _166
	_166:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _168
_168:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _170
	_170:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _172
_172:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _174
	_174:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _176
_176:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(nbEBands*CC)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4084, int32(2074))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _178
	_178:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _180
_180:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(nbEBands*CC) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _182
	_182:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _184
_184:
	bandLogE = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(nbEBands*CC)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	secondMdct = libc.BoolInt32(shortBlocks != 0 && (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fcomplexity >= int32(8))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _186
	_186:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _188
_188:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _190
	_190:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _192
_192:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _194
	_194:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _196
_196:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _198
	_198:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _200
_200:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(C*nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4084, int32(2077))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _202
	_202:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _204
_204:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(C*nbEBands) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _206
	_206:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _208
_208:
	bandLogE2 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(C*nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	if secondMdct != 0 {
		compute_mdcts(tls, mode, 0, in1, freq, C, CC, LM, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fupsample, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Farch)
		Opus_compute_band_energies(tls, mode, freq, bandE, effEnd, C, LM, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Farch)
		Opus_amp2Log2(tls, mode, effEnd, end, bandE, bandLogE2, C)
		c = 0
		for {
			if !(c < C) {
				break
			}
			i1 = 0
			for {
				if !(i1 < end) {
					break
				}
				*(*OpusT_celt_glog)(unsafe.Pointer(bandLogE2 + uintptr(nbEBands*c+i1)*4)) += float32(libc.Float32FromFloat32(0.5) * float32(LM))
				goto _210
			_210:
				;
				i1 = i1 + 1
			}
			goto _209
		_209:
			;
			c = c + 1
		}
	}
	compute_mdcts(tls, mode, shortBlocks, in1, freq, C, CC, LM, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fupsample, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Farch)
	/* This should catch any NaN in the CELT input. Since we're not supposed to see any (they're filtered
	   at the Opus layer), just abort. */
	*(*float32)(unsafe.Pointer(bp)) = *(*OpusT_celt_sig)(unsafe.Pointer(freq))
	v38 = libc.BoolInt32(*(*OpusT_opus_uint32)(unsafe.Pointer(bp))>>libc.Int32FromInt32(23)&uint32(0xFF) == uint32(0xFF) && *(*OpusT_opus_uint32)(unsafe.Pointer(bp))&uint32(0x007FFFFF) != uint32(0))
	goto _212
_212:
	;
	if v215 = !(v38 != 0); v215 {
		if v112 = C == int32(1); !v112 {
			*(*float32)(unsafe.Pointer(bp)) = *(*OpusT_celt_sig)(unsafe.Pointer(freq + uintptr(N)*4))
			v40 = libc.BoolInt32(*(*OpusT_opus_uint32)(unsafe.Pointer(bp))>>libc.Int32FromInt32(23)&uint32(0xFF) == uint32(0xFF) && *(*OpusT_opus_uint32)(unsafe.Pointer(bp))&uint32(0x007FFFFF) != uint32(0))
			goto _214
		_214:
		}
	}
	if !(v215 && (v112 || !(v40 != 0))) {
		Opus_celt_fatal(tls, __ccgo_ts+4249, __ccgo_ts+4084, int32(2093))
	}
	if CC == int32(2) && C == int32(1) {
		*(*int32)(unsafe.Pointer(bp + 80)) = 0
	}
	Opus_compute_band_energies(tls, mode, freq, bandE, effEnd, C, LM, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Farch)
	if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Flfe != 0 {
		i1 = int32(2)
		for {
			if !(i1 < end) {
				break
			}
			if *(*OpusT_celt_ener)(unsafe.Pointer(bandE + uintptr(i1)*4)) < float32(libc.Float32FromFloat32(0.0001)**(*OpusT_celt_ener)(unsafe.Pointer(bandE))) {
				v218 = *(*OpusT_celt_ener)(unsafe.Pointer(bandE + uintptr(i1)*4))
			} else {
				v218 = float32(libc.Float32FromFloat32(0.0001) * *(*OpusT_celt_ener)(unsafe.Pointer(bandE)))
			}
			*(*OpusT_celt_ener)(unsafe.Pointer(bandE + uintptr(i1)*4)) = v218
			if *(*OpusT_celt_ener)(unsafe.Pointer(bandE + uintptr(i1)*4)) > libc.Float32FromFloat32(1e-15) {
				v218 = *(*OpusT_celt_ener)(unsafe.Pointer(bandE + uintptr(i1)*4))
			} else {
				v218 = libc.Float32FromFloat32(1e-15)
			}
			*(*OpusT_celt_ener)(unsafe.Pointer(bandE + uintptr(i1)*4)) = v218
			goto _217
		_217:
			;
			i1 = i1 + 1
		}
	}
	Opus_amp2Log2(tls, mode, effEnd, end, bandE, bandLogE, C)
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _221
	_221:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _223
_223:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _225
	_225:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _227
_227:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _229
	_229:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _231
_231:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _233
	_233:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _235
_235:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(C*nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4084, int32(2108))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _237
	_237:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _239
_239:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(C*nbEBands) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _241
	_241:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _243
_243:
	surround_dynalloc = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(C*nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	libc.Xmemset(tls, surround_dynalloc, 0, libc.Uint64FromInt32(end)*uint64(4))
	/* This computes how much masking takes place between surround channels */
	if !(hybrid != 0) && (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fenergy_mask != 0 && !((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Flfe != 0) {
		mask_avg = libc.Float32FromInt32(0)
		diff = libc.Float32FromInt32(0)
		count = 0
		if int32(2) > (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).FlastCodedBands {
			v38 = int32(2)
		} else {
			v38 = (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).FlastCodedBands
		}
		mask_end = v38
		c = 0
		for {
			if !(c < C) {
				break
			}
			i1 = 0
			for {
				if !(i1 < mask_end) {
					break
				}
				if *(*OpusT_celt_glog)(unsafe.Pointer((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fenergy_mask + uintptr(nbEBands*c+i1)*4)) < libc.Float32FromFloat32(0.25) {
					v248 = *(*OpusT_celt_glog)(unsafe.Pointer((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fenergy_mask + uintptr(nbEBands*c+i1)*4))
				} else {
					v248 = libc.Float32FromFloat32(0.25)
				}
				if v248 > -libc.Float32FromFloat32(2) {
					if *(*OpusT_celt_glog)(unsafe.Pointer((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fenergy_mask + uintptr(nbEBands*c+i1)*4)) < libc.Float32FromFloat32(0.25) {
						v249 = *(*OpusT_celt_glog)(unsafe.Pointer((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fenergy_mask + uintptr(nbEBands*c+i1)*4))
					} else {
						v249 = libc.Float32FromFloat32(0.25)
					}
					v247 = v249
				} else {
					v247 = -libc.Float32FromFloat32(2)
				}
				mask = v247
				if mask > libc.Float32FromInt32(0) {
					mask = float32(libc.Float32FromFloat32(0.5) * mask)
				}
				mask16 = mask
				mask_avg = mask_avg + OpusT_opus_val32(mask16*float32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1+int32(1))*2)))-int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1)*2)))))
				count = count + (int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1+int32(1))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1)*2))))
				diff = diff + OpusT_opus_val32(mask16*float32(libc.Int32FromInt32(1)+libc.Int32FromInt32(2)*i1-mask_end))
				goto _246
			_246:
				;
				i1 = i1 + 1
			}
			goto _245
		_245:
			;
			c = c + 1
		}
		if !(count > libc.Int32FromInt32(0)) {
			Opus_celt_fatal(tls, __ccgo_ts+4322, __ccgo_ts+4084, int32(2136))
		}
		mask_avg = mask_avg / float32(count)
		mask_avg = mask_avg + libc.Float32FromFloat32(0.2)
		diff = OpusT_opus_val32(diff*libc.Float32FromInt32(6)) / float32(C*(mask_end-libc.Int32FromInt32(1))*(mask_end+libc.Int32FromInt32(1))*mask_end)
		/* Again, being conservative */
		diff = float32(libc.Float32FromFloat32(0.5) * diff)
		if diff < libc.Float32FromFloat32(0.031) {
			v83 = diff
		} else {
			v83 = libc.Float32FromFloat32(0.031)
		}
		if v83 > -libc.Float32FromFloat32(0.031) {
			if diff < libc.Float32FromFloat32(0.031) {
				v90 = diff
			} else {
				v90 = libc.Float32FromFloat32(0.031)
			}
			v78 = v90
		} else {
			v78 = -libc.Float32FromFloat32(0.031)
		}
		diff = v78
		/* Find the band that's in the middle of the coded spectrum */
		midband = 0
		for {
			if !(int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(midband+int32(1))*2))) < int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(mask_end)*2)))/int32(2)) {
				break
			}
			goto _253
		_253:
			;
			midband = midband + 1
		}
		count_dynalloc = 0
		i1 = 0
		for {
			if !(i1 < mask_end) {
				break
			}
			lin = mask_avg + OpusT_opus_val32(diff*float32(i1-midband))
			if C == int32(2) {
				if *(*OpusT_celt_glog)(unsafe.Pointer((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fenergy_mask + uintptr(i1)*4)) > *(*OpusT_celt_glog)(unsafe.Pointer((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fenergy_mask + uintptr(nbEBands+i1)*4)) {
					v247 = *(*OpusT_celt_glog)(unsafe.Pointer((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fenergy_mask + uintptr(i1)*4))
				} else {
					v247 = *(*OpusT_celt_glog)(unsafe.Pointer((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fenergy_mask + uintptr(nbEBands+i1)*4))
				}
				unmask = v247
			} else {
				unmask = *(*OpusT_celt_glog)(unsafe.Pointer((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fenergy_mask + uintptr(i1)*4))
			}
			if unmask < libc.Float32FromFloat32(0) {
				v247 = unmask
			} else {
				v247 = libc.Float32FromFloat32(0)
			}
			unmask = v247
			unmask = unmask - lin
			if unmask > libc.Float32FromFloat32(0.25) {
				*(*OpusT_celt_glog)(unsafe.Pointer(surround_dynalloc + uintptr(i1)*4)) = unmask - libc.Float32FromFloat32(0.25)
				count_dynalloc = count_dynalloc + 1
			}
			goto _254
		_254:
			;
			i1 = i1 + 1
		}
		if count_dynalloc >= int32(3) {
			/* If we need dynalloc in many bands, it's probably because our
			   initial masking rate was too low. */
			mask_avg = mask_avg + libc.Float32FromFloat32(0.25)
			if mask_avg > libc.Float32FromInt32(0) {
				/* Something went really wrong in the original calculations,
				   disabling masking. */
				mask_avg = libc.Float32FromInt32(0)
				diff = libc.Float32FromInt32(0)
				libc.Xmemset(tls, surround_dynalloc, 0, libc.Uint64FromInt32(mask_end)*uint64(4))
			} else {
				i1 = 0
				for {
					if !(i1 < mask_end) {
						break
					}
					if float32(libc.Int32FromInt32(0)) > *(*OpusT_celt_glog)(unsafe.Pointer(surround_dynalloc + uintptr(i1)*4))-libc.Float32FromFloat32(0.25) {
						v247 = float32(libc.Int32FromInt32(0))
					} else {
						v247 = *(*OpusT_celt_glog)(unsafe.Pointer(surround_dynalloc + uintptr(i1)*4)) - libc.Float32FromFloat32(0.25)
					}
					*(*OpusT_celt_glog)(unsafe.Pointer(surround_dynalloc + uintptr(i1)*4)) = v247
					goto _257
				_257:
					;
					i1 = i1 + 1
				}
			}
		}
		mask_avg = mask_avg + libc.Float32FromFloat32(0.2)
		/* Convert to 1/64th units used for the trim */
		surround_trim = OpusT_celt_glog(libc.Float32FromInt32(64) * diff)
		/*printf("%d %d ", mask_avg, surround_trim);*/
		surround_masking = mask_avg
	}
	/* Temporal VBR (but not for LFE) */
	if !((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Flfe != 0) {
		follow = -libc.Float32FromFloat32(10)
		frame_avg = libc.Float32FromInt32(0)
		if shortBlocks != 0 {
			v259 = float32(libc.Float32FromFloat32(0.5) * float32(LM))
		} else {
			v259 = libc.Float32FromInt32(0)
		}
		offset = v259
		i1 = start
		for {
			if !(i1 < end) {
				break
			}
			if follow-libc.Float32FromFloat32(1) > *(*OpusT_celt_glog)(unsafe.Pointer(bandLogE + uintptr(i1)*4))-offset {
				v247 = follow - libc.Float32FromFloat32(1)
			} else {
				v247 = *(*OpusT_celt_glog)(unsafe.Pointer(bandLogE + uintptr(i1)*4)) - offset
			}
			follow = v247
			if C == int32(2) {
				if follow > *(*OpusT_celt_glog)(unsafe.Pointer(bandLogE + uintptr(i1+nbEBands)*4))-offset {
					v247 = follow
				} else {
					v247 = *(*OpusT_celt_glog)(unsafe.Pointer(bandLogE + uintptr(i1+nbEBands)*4)) - offset
				}
				follow = v247
			}
			frame_avg = frame_avg + follow
			goto _260
		_260:
			;
			i1 = i1 + 1
		}
		frame_avg = frame_avg / float32(end-start)
		temporal_vbr = frame_avg - (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fspec_avg
		if -libc.Float32FromFloat32(1.5) > temporal_vbr {
			v263 = -libc.Float32FromFloat32(1.5)
		} else {
			v263 = temporal_vbr
		}
		if libc.Float32FromFloat32(3) < v263 {
			v259 = libc.Float32FromFloat32(3)
		} else {
			if -libc.Float32FromFloat32(1.5) > temporal_vbr {
				v264 = -libc.Float32FromFloat32(1.5)
			} else {
				v264 = temporal_vbr
			}
			v259 = v264
		}
		temporal_vbr = v259
		*(*OpusT_celt_glog)(unsafe.Pointer(st1 + 248)) += float32(libc.Float32FromFloat32(0.02) * temporal_vbr)
	}
	/*for (i=0;i<21;i++)
	   printf("%f ", bandLogE[i]);
	printf("\n");*/
	if !(secondMdct != 0) {
		libc.Xmemcpy(tls, bandLogE2, bandLogE, libc.Uint64FromInt32(C*nbEBands)*uint64(4)+libc.Uint64FromInt64(0*((int64(bandLogE2)-int64(bandLogE))/4)))
	}
	/* Last chance to catch any transient we might have missed in the
	   time-domain analysis */
	if v112 = LM > 0; v112 {
		v1 = enc
		v38 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (libc.Int32FromInt64(4)*libc.Int32FromInt32(__CHAR_BIT__) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
		goto _268
	_268:
	}
	if v112 && v38+int32(3) <= total_bits && !(isTransient != 0) && (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fcomplexity >= int32(5) && !((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Flfe != 0) && !(hybrid != 0) {
		if patch_transient_decision(tls, bandLogE, oldBandE, nbEBands, start, end, C) != 0 {
			isTransient = int32(1)
			shortBlocks = M
			compute_mdcts(tls, mode, shortBlocks, in1, freq, C, CC, LM, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fupsample, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Farch)
			Opus_compute_band_energies(tls, mode, freq, bandE, effEnd, C, LM, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Farch)
			Opus_amp2Log2(tls, mode, effEnd, end, bandE, bandLogE, C)
			/* Compensate for the scaling of short vs long mdcts */
			c = 0
			for {
				if !(c < C) {
					break
				}
				i1 = 0
				for {
					if !(i1 < end) {
						break
					}
					*(*OpusT_celt_glog)(unsafe.Pointer(bandLogE2 + uintptr(nbEBands*c+i1)*4)) += float32(libc.Float32FromFloat32(0.5) * float32(LM))
					goto _271
				_271:
					;
					i1 = i1 + 1
				}
				goto _270
			_270:
				;
				c = c + 1
			}
			*(*OpusT_opus_val16)(unsafe.Pointer(bp + 84)) = libc.Float32FromFloat32(0.2)
		}
	}
	if v112 = LM > 0; v112 {
		v1 = enc
		v38 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (libc.Int32FromInt64(4)*libc.Int32FromInt32(__CHAR_BIT__) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
		goto _274
	_274:
	}
	if v112 && v38+int32(3) <= total_bits {
		Opus_ec_enc_bit_logp(tls, enc, isTransient, uint32(3))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _277
	_277:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _279
_279:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _281
	_281:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _283
_283:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _285
	_285:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _287
_287:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _289
	_289:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _291
_291:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(C*N)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4084, int32(2237))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _293
	_293:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _295
_295:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(C*N) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _297
	_297:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _299
_299:
	X = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(C*N)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1))) /**< Interleaved normalised MDCTs */
	/* Band normalisation */
	Opus_normalise_bands(tls, mode, freq, X, bandE, effEnd, C, M)
	enable_tf_analysis = libc.BoolInt32(effectiveBytes >= int32(15)*C && !(hybrid != 0) && (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fcomplexity >= int32(2) && !((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Flfe != 0) && *(*OpusT_opus_val32)(unsafe.Pointer(bp + 96)) < libc.Float32FromFloat32(0.98))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _301
	_301:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _303
_303:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _305
	_305:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _307
_307:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _309
	_309:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _311
_311:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _313
	_313:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _315
_315:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4084, int32(2244))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _317
	_317:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _319
_319:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(nbEBands) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _321
	_321:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _323
_323:
	offsets = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _325
	_325:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _327
_327:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _329
	_329:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _331
_331:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _333
	_333:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _335
_335:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _337
	_337:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _339
_339:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4084, int32(2245))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _341
	_341:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _343
_343:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(nbEBands) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _345
	_345:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _347
_347:
	importance = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _349
	_349:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _351
_351:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _353
	_353:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _355
_355:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _357
	_357:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _359
_359:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _361
	_361:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _363
_363:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4084, int32(2246))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _365
	_365:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _367
_367:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(nbEBands) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _369
	_369:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _371
_371:
	spread_weight = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	maxDepth = dynalloc_analysis(tls, bandLogE, bandLogE2, oldBandE, nbEBands, start, end, C, offsets, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Flsb_depth, (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FlogN, isTransient, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fvbr, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fconstrained_vbr, eBands, LM, effectiveBytes, bp+88, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Flfe, surround_dynalloc, st1+124, importance, spread_weight, tone_freq, *(*OpusT_opus_val32)(unsafe.Pointer(bp + 96)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _373
	_373:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _375
_375:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _377
	_377:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _379
_379:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _381
	_381:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _383
_383:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _385
	_385:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _387
_387:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4084, int32(2252))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _389
	_389:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _391
_391:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(nbEBands) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _393
	_393:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _395
_395:
	tf_res = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	/* Disable variable tf resolution for hybrid and at very low bitrate */
	if enable_tf_analysis != 0 {
		if int32(80) > int32(20480)/effectiveBytes+int32(2) {
			v38 = int32(80)
		} else {
			v38 = int32(20480)/effectiveBytes + int32(2)
		}
		lambda = v38
		tf_select = tf_analysis(tls, mode, effEnd, isTransient, tf_res, lambda, X, N, LM, *(*OpusT_opus_val16)(unsafe.Pointer(bp + 84)), *(*int32)(unsafe.Pointer(bp + 80)), importance)
		i1 = effEnd
		for {
			if !(i1 < end) {
				break
			}
			*(*int32)(unsafe.Pointer(tf_res + uintptr(i1)*4)) = *(*int32)(unsafe.Pointer(tf_res + uintptr(effEnd-int32(1))*4))
			goto _397
		_397:
			;
			i1 = i1 + 1
		}
	} else {
		if hybrid != 0 && *(*int32)(unsafe.Pointer(bp + 92)) != 0 {
			/* For weak transients, we rely on the fact that improving time resolution using
			   TF on a long window is imperfect and will not result in an energy collapse at
			   low bitrate. */
			i1 = 0
			for {
				if !(i1 < end) {
					break
				}
				*(*int32)(unsafe.Pointer(tf_res + uintptr(i1)*4)) = int32(1)
				goto _398
			_398:
				;
				i1 = i1 + 1
			}
			tf_select = 0
		} else {
			if hybrid != 0 && effectiveBytes < int32(15) && (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fsilk_info.FsignalType != int32(2) {
				/* For low bitrate hybrid, we force temporal resolution to 5 ms rather than 2.5 ms. */
				i1 = 0
				for {
					if !(i1 < end) {
						break
					}
					*(*int32)(unsafe.Pointer(tf_res + uintptr(i1)*4)) = 0
					goto _399
				_399:
					;
					i1 = i1 + 1
				}
				tf_select = isTransient
			} else {
				i1 = 0
				for {
					if !(i1 < end) {
						break
					}
					*(*int32)(unsafe.Pointer(tf_res + uintptr(i1)*4)) = isTransient
					goto _400
				_400:
					;
					i1 = i1 + 1
				}
				tf_select = 0
			}
		}
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _402
	_402:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _404
_404:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _406
	_406:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _408
_408:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _410
	_410:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _412
_412:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _414
	_414:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _416
_416:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(C*nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4084, int32(2281))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _418
	_418:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _420
_420:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(C*nbEBands) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _422
	_422:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _424
_424:
	error1 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(C*nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	c = 0
	for {
		i1 = start
		for {
			if !(i1 < end) {
				break
			}
			/* When the energy is stable, slightly bias energy quantization towards
			   the previous error to make the gain more stable (a constant offset is
			   better than fluctuations). */
			if float32(libc.Xfabs(tls, float64(*(*OpusT_celt_glog)(unsafe.Pointer(bandLogE + uintptr(i1+c*nbEBands)*4))-*(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(i1+c*nbEBands)*4))))) < libc.Float32FromFloat32(2) {
				*(*OpusT_celt_glog)(unsafe.Pointer(bandLogE + uintptr(i1+c*nbEBands)*4)) -= float32(libc.Float32FromFloat32(0.25) * *(*OpusT_celt_glog)(unsafe.Pointer(energyError + uintptr(i1+c*nbEBands)*4)))
			}
			goto _427
		_427:
			;
			i1 = i1 + 1
		}
		goto _426
	_426:
		;
		c = c + 1
		v38 = c
		if !(v38 < C) {
			break
		}
	}
	Opus_quant_coarse_energy(tls, mode, start, end, effEnd, bandLogE, oldBandE, libc.Uint32FromInt32(total_bits), error1, enc, C, LM, nbAvailableBytes, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fforce_intra, st1+88, libc.BoolInt32((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fcomplexity >= int32(4)), (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Floss_rate, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Flfe)
	tf_encode(tls, start, end, isTransient, tf_res, LM, tf_select, enc)
	v1 = enc
	v38 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (libc.Int32FromInt64(4)*libc.Int32FromInt32(__CHAR_BIT__) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
	goto _430
_430:
	if v38+int32(4) <= total_bits {
		if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Flfe != 0 {
			(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Ftapset_decision = 0
			(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fspread_decision = int32(SPREAD_NORMAL)
		} else {
			if hybrid != 0 {
				if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fcomplexity == 0 {
					(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fspread_decision = SPREAD_NONE
				} else {
					if isTransient != 0 {
						(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fspread_decision = int32(SPREAD_NORMAL)
					} else {
						(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fspread_decision = int32(SPREAD_AGGRESSIVE)
					}
				}
			} else {
				if shortBlocks != 0 || (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fcomplexity < int32(3) || nbAvailableBytes < int32(10)*C {
					if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fcomplexity == 0 {
						(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fspread_decision = SPREAD_NONE
					} else {
						(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fspread_decision = int32(SPREAD_NORMAL)
					}
				} else {
					/* Disable new spreading+tapset estimator until we can show it works
					   better than the old one. So far it seems like spreading_decision()
					   works best. */
					(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fspread_decision = Opus_spreading_decision(tls, mode, X, st1+92, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fspread_decision, st1+100, st1+104, libc.BoolInt32(pf_on != 0 && !(shortBlocks != 0)), effEnd, C, M, spread_weight)
					/*printf("%d %d\n", st->tapset_decision, st->spread_decision);*/
					/*printf("%f %d %f %d\n\n", st->analysis.tonality, st->spread_decision, st->analysis.tonality_slope, st->tapset_decision);*/
				}
			}
		}
		Opus_ec_enc_icdf(tls, enc, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fspread_decision, uintptr(unsafe.Pointer(&spread_icdf14)), uint32(5))
	} else {
		(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fspread_decision = int32(SPREAD_NORMAL)
	}
	/* For LFE, everything interesting is in the first band */
	if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Flfe != 0 {
		if int32(8) < effectiveBytes/int32(3) {
			v38 = int32(8)
		} else {
			v38 = effectiveBytes / int32(3)
		}
		*(*int32)(unsafe.Pointer(offsets)) = v38
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _433
	_433:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _435
_435:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _437
	_437:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _439
_439:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _441
	_441:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _443
_443:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _445
	_445:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _447
_447:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4084, int32(2353))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _449
	_449:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _451
_451:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(nbEBands) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _453
	_453:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _455
_455:
	cap1 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	Opus_init_caps(tls, mode, cap1, LM, C)
	dynalloc_logp = int32(6)
	total_bits = total_bits << int32(BITRES)
	total_boost = 0
	tell = libc.Int32FromUint32(Opus_ec_tell_frac(tls, enc))
	i1 = start
	for {
		if !(i1 < end) {
			break
		}
		width = C * (int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1+int32(1))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer(eBands + uintptr(i1)*2)))) << LM
		/* quanta is 6 bits, but no more than 1 bit/sample
		   and no less than 1/8 bit/sample */
		if libc.Int32FromInt32(6)<<libc.Int32FromInt32(BITRES) > width {
			v40 = libc.Int32FromInt32(6) << libc.Int32FromInt32(BITRES)
		} else {
			v40 = width
		}
		if width<<int32(BITRES) < v40 {
			v38 = width << int32(BITRES)
		} else {
			if libc.Int32FromInt32(6)<<libc.Int32FromInt32(BITRES) > width {
				v43 = libc.Int32FromInt32(6) << libc.Int32FromInt32(BITRES)
			} else {
				v43 = width
			}
			v38 = v43
		}
		quanta = v38
		dynalloc_loop_logp = dynalloc_logp
		boost = 0
		j = 0
		for {
			if !(tell+dynalloc_loop_logp<<int32(BITRES) < total_bits-total_boost && boost < *(*int32)(unsafe.Pointer(cap1 + uintptr(i1)*4))) {
				break
			}
			flag = libc.BoolInt32(j < *(*int32)(unsafe.Pointer(offsets + uintptr(i1)*4)))
			Opus_ec_enc_bit_logp(tls, enc, flag, libc.Uint32FromInt32(dynalloc_loop_logp))
			tell = libc.Int32FromUint32(Opus_ec_tell_frac(tls, enc))
			if !(flag != 0) {
				break
			}
			boost = boost + quanta
			total_boost = total_boost + quanta
			dynalloc_loop_logp = int32(1)
			goto _460
		_460:
			;
			j = j + 1
		}
		/* Making dynalloc more likely */
		if j != 0 {
			if int32(2) > dynalloc_logp-int32(1) {
				v38 = int32(2)
			} else {
				v38 = dynalloc_logp - int32(1)
			}
			dynalloc_logp = v38
		}
		*(*int32)(unsafe.Pointer(offsets + uintptr(i1)*4)) = boost
		goto _456
	_456:
		;
		i1 = i1 + 1
	}
	if C == int32(2) {
		/* Always use MS for 2.5 ms frames until we can do a better analysis */
		if LM != 0 {
			*(*int32)(unsafe.Pointer(bp + 72)) = stereo_analysis(tls, mode, X, LM, N)
		}
		(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fintensity = Opus_hysteresis_decision(tls, float32(equiv_rate/libc.Int32FromInt32(1000)), uintptr(unsafe.Pointer(&intensity_thresholds)), uintptr(unsafe.Pointer(&intensity_histeresis)), int32(21), (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fintensity)
		if start > (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fintensity {
			v40 = start
		} else {
			v40 = (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fintensity
		}
		if end < v40 {
			v38 = end
		} else {
			if start > (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fintensity {
				v43 = start
			} else {
				v43 = (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fintensity
			}
			v38 = v43
		}
		(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fintensity = v38
	}
	alloc_trim = int32(5)
	if tell+libc.Int32FromInt32(6)<<libc.Int32FromInt32(BITRES) <= total_bits-total_boost {
		if start > 0 || (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Flfe != 0 {
			(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fstereo_saving = libc.Float32FromInt32(0)
			alloc_trim = int32(5)
		} else {
			alloc_trim = alloc_trim_analysis(tls, mode, X, bandLogE, end, LM, C, N, st1+124, st1+232, *(*OpusT_opus_val16)(unsafe.Pointer(bp + 84)), (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fintensity, surround_trim, equiv_rate, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Farch)
		}
		Opus_ec_enc_icdf(tls, enc, alloc_trim, uintptr(unsafe.Pointer(&trim_icdf14)), uint32(7))
		tell = libc.Int32FromUint32(Opus_ec_tell_frac(tls, enc))
	}
	/* In VBR mode the frame size must not be reduced so much that it would
	    result in the encoder running out of bits.
	   The margin of 2 bytes ensures that none of the bust-prevention logic
	    in the decoder will have triggered so far. */
	min_allowed = (tell+total_boost+libc.Int32FromInt32(1)<<(libc.Int32FromInt32(BITRES)+libc.Int32FromInt32(3))-int32(1))>>(libc.Int32FromInt32(BITRES)+libc.Int32FromInt32(3)) + int32(2)
	/* Take into account the 37 bits we need to have left in the packet to
	   signal a redundant frame in hybrid mode. Creating a shorter packet would
	   create an entropy coder desync. */
	if hybrid != 0 {
		if min_allowed > (tell0_frac+libc.Int32FromInt32(37)<<libc.Int32FromInt32(BITRES)+total_boost+libc.Int32FromInt32(1)<<(libc.Int32FromInt32(BITRES)+libc.Int32FromInt32(3))-int32(1))>>(libc.Int32FromInt32(BITRES)+libc.Int32FromInt32(3)) {
			v38 = min_allowed
		} else {
			v38 = (tell0_frac + libc.Int32FromInt32(37)<<libc.Int32FromInt32(BITRES) + total_boost + libc.Int32FromInt32(1)<<(libc.Int32FromInt32(BITRES)+libc.Int32FromInt32(3)) - int32(1)) >> (libc.Int32FromInt32(BITRES) + libc.Int32FromInt32(3))
		}
		min_allowed = v38
	}
	/* Variable bitrate */
	if vbr_rate > 0 {
		lm_diff = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FmaxLM - LM
		/* Don't attempt to use more than 510 kb/s, even for frames smaller than 20 ms.
		   The CELT allocator will just not be able to use more than that anyway. */
		if nbCompressedBytes < packet_size_cap>>(int32(3)-LM) {
			v38 = nbCompressedBytes
		} else {
			v38 = packet_size_cap >> (int32(3) - LM)
		}
		nbCompressedBytes = v38
		if !(hybrid != 0) {
			base_target = vbr_rate - (int32(40)*C+int32(20))<<int32(BITRES)
		} else {
			if 0 > vbr_rate-(int32(9)*C+int32(4))<<int32(BITRES) {
				v38 = 0
			} else {
				v38 = vbr_rate - (int32(9)*C+int32(4))<<int32(BITRES)
			}
			base_target = v38
		}
		if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fconstrained_vbr != 0 {
			base_target = base_target + (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fvbr_offset>>lm_diff
		}
		if !(hybrid != 0) {
			target = compute_vbr(tls, mode, st1+124, base_target, LM, equiv_rate, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).FlastCodedBands, C, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fintensity, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fconstrained_vbr, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fstereo_saving, *(*OpusT_opus_int32)(unsafe.Pointer(bp + 88)), *(*OpusT_opus_val16)(unsafe.Pointer(bp + 84)), pitch_change, maxDepth, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Flfe, libc.BoolInt32((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fenergy_mask != libc.UintptrFromInt32(0)), surround_masking, temporal_vbr)
		} else {
			target = base_target
			/* Tonal frames (offset<100) need more bits than noisy (offset>100) ones. */
			if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fsilk_info.Foffset < int32(100) {
				target = target + libc.Int32FromInt32(12)<<libc.Int32FromInt32(BITRES)>>(int32(3)-LM)
			}
			if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fsilk_info.Foffset > int32(100) {
				target = target - libc.Int32FromInt32(18)<<libc.Int32FromInt32(BITRES)>>(int32(3)-LM)
			}
			/* Boosting bitrate on transients and vowels with significant temporal
			   spikes. */
			target = target + int32(OpusT_opus_val16((*(*OpusT_opus_val16)(unsafe.Pointer(bp + 84))-libc.Float32FromFloat32(0.25))*float32(libc.Int32FromInt32(50)<<libc.Int32FromInt32(BITRES))))
			/* If we have a strong transient, let's make sure it has enough bits to code
			   the first two bands, so that it can use folding rather than noise. */
			if *(*OpusT_opus_val16)(unsafe.Pointer(bp + 84)) > libc.Float32FromFloat32(0.7) {
				if target > libc.Int32FromInt32(50)<<libc.Int32FromInt32(BITRES) {
					v38 = target
				} else {
					v38 = libc.Int32FromInt32(50) << libc.Int32FromInt32(BITRES)
				}
				target = v38
			}
		}
		/* The current offset is removed from the target and the space used
		   so far is added*/
		target = target + tell
		nbAvailableBytes = (target + libc.Int32FromInt32(1)<<(libc.Int32FromInt32(BITRES)+libc.Int32FromInt32(2))) >> (libc.Int32FromInt32(BITRES) + libc.Int32FromInt32(3))
		if min_allowed > nbAvailableBytes {
			v38 = min_allowed
		} else {
			v38 = nbAvailableBytes
		}
		nbAvailableBytes = v38
		if nbCompressedBytes < nbAvailableBytes {
			v38 = nbCompressedBytes
		} else {
			v38 = nbAvailableBytes
		}
		nbAvailableBytes = v38
		/* By how much did we "miss" the target on that frame */
		delta = target - vbr_rate
		target = nbAvailableBytes << (libc.Int32FromInt32(BITRES) + libc.Int32FromInt32(3))
		/*If the frame is silent we don't adjust our drift, otherwise
		  the encoder will shoot to very high rates after hitting a
		  span of silence, but we do allow the bitres to refill.
		  This means that we'll undershoot our target in CVBR/VBR modes
		  on files with lots of silence. */
		if silence != 0 {
			nbAvailableBytes = int32(2)
			target = libc.Int32FromInt32(2) * libc.Int32FromInt32(8) << libc.Int32FromInt32(BITRES)
			delta = 0
		}
		if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fvbr_count < int32(970) {
			(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fvbr_count = (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fvbr_count + 1
			alpha = libc.Float32FromFloat32(1) / float32((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fvbr_count+libc.Int32FromInt32(20))
		} else {
			alpha = libc.Float32FromFloat32(0.001)
		}
		/* How many bits have we used in excess of what we're allowed */
		if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fconstrained_vbr != 0 {
			*(*OpusT_opus_int32)(unsafe.Pointer(st1 + 212)) += target - vbr_rate
		}
		/*printf ("%d\n", st->vbr_reservoir);*/
		/* Compute the offset we need to apply in order to reach the target */
		if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fconstrained_vbr != 0 {
			*(*OpusT_opus_int32)(unsafe.Pointer(st1 + 216)) += int32(OpusT_opus_val16(alpha * float32(delta*(libc.Int32FromInt32(1)<<lm_diff)-(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fvbr_offset-(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fvbr_drift)))
			(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fvbr_offset = -(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fvbr_drift
		}
		/*printf ("%d\n", st->vbr_drift);*/
		if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fconstrained_vbr != 0 && (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fvbr_reservoir < 0 {
			/* We're under the min value -- increase rate */
			adjust = -(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fvbr_reservoir / (libc.Int32FromInt32(8) << libc.Int32FromInt32(BITRES))
			/* Unless we're just coding silence */
			if silence != 0 {
				v38 = 0
			} else {
				v38 = adjust
			}
			nbAvailableBytes = nbAvailableBytes + v38
			(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fvbr_reservoir = 0
			/*printf ("+%d\n", adjust);*/
		}
		if nbCompressedBytes < nbAvailableBytes {
			v38 = nbCompressedBytes
		} else {
			v38 = nbAvailableBytes
		}
		nbCompressedBytes = v38
		/*printf("%d\n", nbCompressedBytes*50*8);*/
		/* This moves the raw bits to take into account the new compressed size */
		Opus_ec_enc_shrink(tls, enc, libc.Uint32FromInt32(nbCompressedBytes))
	}
	/* Bit allocation */
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _474
	_474:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _476
_476:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _478
	_478:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _480
_480:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _482
	_482:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _484
_484:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _486
	_486:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _488
_488:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4084, int32(2598))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _490
	_490:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _492
_492:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(nbEBands) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _494
	_494:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _496
_496:
	fine_quant = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _498
	_498:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _500
_500:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _502
	_502:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _504
_504:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _506
	_506:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _508
_508:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _510
	_510:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _512
_512:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4084, int32(2599))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _514
	_514:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _516
_516:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(nbEBands) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _518
	_518:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _520
_520:
	pulses = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _522
	_522:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _524
_524:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _526
	_526:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _528
_528:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _530
	_530:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _532
_532:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _534
	_534:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _536
_536:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4084, int32(2600))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _538
	_538:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _540
_540:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(nbEBands) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _542
	_542:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _544
_544:
	fine_priority = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(nbEBands)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	/* bits =           packet size                    - where we are - safety*/
	bits = nbCompressedBytes*int32(8)<<int32(BITRES) - libc.Int32FromUint32(Opus_ec_tell_frac(tls, enc)) - int32(1)
	if isTransient != 0 && LM >= int32(2) && bits >= (LM+int32(2))<<int32(BITRES) {
		v38 = libc.Int32FromInt32(1) << libc.Int32FromInt32(BITRES)
	} else {
		v38 = 0
	}
	anti_collapse_rsv = v38
	bits = bits - anti_collapse_rsv
	signalBandwidth = end - int32(1)
	if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fanalysis.Fvalid != 0 {
		if equiv_rate < libc.Int32FromInt32(32000)*C {
			min_bandwidth = int32(13)
		} else {
			if equiv_rate < libc.Int32FromInt32(48000)*C {
				min_bandwidth = int32(16)
			} else {
				if equiv_rate < libc.Int32FromInt32(60000)*C {
					min_bandwidth = int32(18)
				} else {
					if equiv_rate < libc.Int32FromInt32(80000)*C {
						min_bandwidth = int32(19)
					} else {
						min_bandwidth = int32(20)
					}
				}
			}
		}
		if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fanalysis.Fbandwidth > min_bandwidth {
			v38 = (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fanalysis.Fbandwidth
		} else {
			v38 = min_bandwidth
		}
		signalBandwidth = v38
	}
	if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Flfe != 0 {
		signalBandwidth = int32(1)
	}
	codedBands = Opus_clt_compute_allocation(tls, mode, start, end, offsets, cap1, alloc_trim, st1+236, bp+72, bits, bp+76, pulses, fine_quant, fine_priority, C, LM, enc, int32(1), (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).FlastCodedBands, signalBandwidth)
	if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).FlastCodedBands != 0 {
		if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).FlastCodedBands-int32(1) > codedBands {
			v40 = (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).FlastCodedBands - int32(1)
		} else {
			v40 = codedBands
		}
		if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).FlastCodedBands+int32(1) < v40 {
			v38 = (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).FlastCodedBands + int32(1)
		} else {
			if (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).FlastCodedBands-int32(1) > codedBands {
				v43 = (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).FlastCodedBands - int32(1)
			} else {
				v43 = codedBands
			}
			v38 = v43
		}
		(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).FlastCodedBands = v38
	} else {
		(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).FlastCodedBands = codedBands
	}
	Opus_quant_fine_energy(tls, mode, start, end, oldBandE, error1, libc.UintptrFromInt32(0), fine_quant, enc, C)
	libc.Xmemset(tls, energyError, 0, libc.Uint64FromInt32(nbEBands*CC)*uint64(4))
	/* Residual quantisation */
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _551
	_551:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _553
_553:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _555
	_555:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v6 = st
	goto _557
_557:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(1) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (libc.Uint64FromInt32(1) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _559
	_559:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v10 = st
	goto _561
_561:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _563
	_563:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v13 = st
	goto _565
_565:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(C*nbEBands)*(libc.Uint64FromInt64(1)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4084, int32(2669))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _567
	_567:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v17 = st
	goto _569
_569:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(libc.Uint64FromInt32(C*nbEBands) * (libc.Uint64FromInt64(1) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _571
	_571:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v21 = st
	goto _573
_573:
	collapse_masks = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(libc.Uint64FromInt32(C*nbEBands)*(libc.Uint64FromInt64(1)/libc.Uint64FromInt64(1)))
	if C == int32(2) {
		v1 = X + uintptr(N)*4
	} else {
		v1 = libc.UintptrFromInt32(0)
	}
	Opus_quant_all_bands(tls, int32(1), mode, start, end, X, v1, collapse_masks, bandE, pulses, shortBlocks, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fspread_decision, *(*int32)(unsafe.Pointer(bp + 72)), (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fintensity, tf_res, nbCompressedBytes*(libc.Int32FromInt32(8)<<libc.Int32FromInt32(BITRES))-anti_collapse_rsv, *(*OpusT_opus_int32)(unsafe.Pointer(bp + 76)), enc, LM, codedBands, st1+80, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fcomplexity, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Farch, (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fdisable_inv)
	if anti_collapse_rsv > 0 {
		anti_collapse_on = libc.BoolInt32((*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fconsec_transient < int32(2))
		Opus_ec_enc_bits(tls, enc, libc.Uint32FromInt32(anti_collapse_on), uint32(1))
	}
	if qext_bytes == 0 {
		v1 = enc
		v38 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (libc.Int32FromInt64(4)*libc.Int32FromInt32(__CHAR_BIT__) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
		goto _577
	_577:
		Opus_quant_energy_finalise(tls, mode, start, end, oldBandE, error1, fine_quant, fine_priority, nbCompressedBytes*int32(8)-v38, enc, C)
	}
	c = 0
	for {
		i1 = start
		for {
			if !(i1 < end) {
				break
			}
			if libc.Float32FromFloat32(0.5) < *(*OpusT_celt_glog)(unsafe.Pointer(error1 + uintptr(i1+c*nbEBands)*4)) {
				v263 = libc.Float32FromFloat32(0.5)
			} else {
				v263 = *(*OpusT_celt_glog)(unsafe.Pointer(error1 + uintptr(i1+c*nbEBands)*4))
			}
			if -libc.Float32FromFloat32(0.5) > v263 {
				v259 = -libc.Float32FromFloat32(0.5)
			} else {
				if libc.Float32FromFloat32(0.5) < *(*OpusT_celt_glog)(unsafe.Pointer(error1 + uintptr(i1+c*nbEBands)*4)) {
					v264 = libc.Float32FromFloat32(0.5)
				} else {
					v264 = *(*OpusT_celt_glog)(unsafe.Pointer(error1 + uintptr(i1+c*nbEBands)*4))
				}
				v259 = v264
			}
			*(*OpusT_celt_glog)(unsafe.Pointer(energyError + uintptr(i1+c*nbEBands)*4)) = v259
			goto _580
		_580:
			;
			i1 = i1 + 1
		}
		goto _579
	_579:
		;
		c = c + 1
		v38 = c
		if !(v38 < C) {
			break
		}
	}
	if silence != 0 {
		i1 = 0
		for {
			if !(i1 < C*nbEBands) {
				break
			}
			*(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(i1)*4)) = -libc.Float32FromFloat32(28)
			goto _584
		_584:
			;
			i1 = i1 + 1
		}
	}
	(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fprefilter_period = *(*int32)(unsafe.Pointer(bp + 64))
	(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fprefilter_gain = *(*OpusT_opus_val16)(unsafe.Pointer(bp + 68))
	(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fprefilter_tapset = prefilter_tapset
	if CC == int32(2) && C == int32(1) {
		libc.Xmemcpy(tls, oldBandE+uintptr(nbEBands)*4, oldBandE, libc.Uint64FromInt32(nbEBands)*uint64(4)+libc.Uint64FromInt64(0*((OpusT___predefined_ptrdiff_t(oldBandE+uintptr(nbEBands)*4)-int64(oldBandE))/4)))
	}
	if !(isTransient != 0) {
		libc.Xmemcpy(tls, oldLogE2, oldLogE, libc.Uint64FromInt32(CC*nbEBands)*uint64(4)+libc.Uint64FromInt64(0*((int64(oldLogE2)-int64(oldLogE))/4)))
		libc.Xmemcpy(tls, oldLogE, oldBandE, libc.Uint64FromInt32(CC*nbEBands)*uint64(4)+libc.Uint64FromInt64(0*((int64(oldLogE)-int64(oldBandE))/4)))
	} else {
		i1 = 0
		for {
			if !(i1 < CC*nbEBands) {
				break
			}
			if *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(i1)*4)) < *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(i1)*4)) {
				v247 = *(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(i1)*4))
			} else {
				v247 = *(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(i1)*4))
			}
			*(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(i1)*4)) = v247
			goto _585
		_585:
			;
			i1 = i1 + 1
		}
	}
	/* In case start or end were to change */
	c = 0
	for {
		i1 = 0
		for {
			if !(i1 < start) {
				break
			}
			*(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i1)*4)) = libc.Float32FromInt32(0)
			v247 = -libc.Float32FromFloat32(28)
			*(*OpusT_celt_glog)(unsafe.Pointer(oldLogE2 + uintptr(c*nbEBands+i1)*4)) = v247
			*(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(c*nbEBands+i1)*4)) = v247
			goto _589
		_589:
			;
			i1 = i1 + 1
		}
		i1 = end
		for {
			if !(i1 < nbEBands) {
				break
			}
			*(*OpusT_celt_glog)(unsafe.Pointer(oldBandE + uintptr(c*nbEBands+i1)*4)) = libc.Float32FromInt32(0)
			v247 = -libc.Float32FromFloat32(28)
			*(*OpusT_celt_glog)(unsafe.Pointer(oldLogE2 + uintptr(c*nbEBands+i1)*4)) = v247
			*(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(c*nbEBands+i1)*4)) = v247
			goto _591
		_591:
			;
			i1 = i1 + 1
		}
		goto _588
	_588:
		;
		c = c + 1
		v38 = c
		if !(v38 < CC) {
			break
		}
	}
	if isTransient != 0 || transient_got_disabled != 0 {
		(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fconsec_transient = (*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fconsec_transient + 1
	} else {
		(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Fconsec_transient = 0
	}
	(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st1)).Frng = (*OpusT_ec_enc)(unsafe.Pointer(enc)).Frng
	/* If there's any room left (can only happen for very high rates),
	   it's already filled with zeros */
	Opus_ec_enc_done(tls, enc)
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _594
	_594:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _596
_596:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
	v38 = (*OpusT_ec_ctx)(unsafe.Pointer(enc)).Ferror1
	goto _598
_598:
	if v38 != 0 {
		return -int32(3)
	} else {
		return nbCompressedBytes
	}
	return r
}

var intensity_thresholds = [21]OpusT_opus_val16{
	0:  libc.Float32FromInt32(1),
	1:  libc.Float32FromInt32(2),
	2:  libc.Float32FromInt32(3),
	3:  libc.Float32FromInt32(4),
	4:  libc.Float32FromInt32(5),
	5:  libc.Float32FromInt32(6),
	6:  libc.Float32FromInt32(7),
	7:  libc.Float32FromInt32(8),
	8:  libc.Float32FromInt32(16),
	9:  libc.Float32FromInt32(24),
	10: libc.Float32FromInt32(36),
	11: libc.Float32FromInt32(44),
	12: libc.Float32FromInt32(50),
	13: libc.Float32FromInt32(56),
	14: libc.Float32FromInt32(62),
	15: libc.Float32FromInt32(67),
	16: libc.Float32FromInt32(72),
	17: libc.Float32FromInt32(79),
	18: libc.Float32FromInt32(88),
	19: libc.Float32FromInt32(106),
	20: libc.Float32FromInt32(134),
}

var intensity_histeresis = [21]OpusT_opus_val16{
	0:  libc.Float32FromInt32(1),
	1:  libc.Float32FromInt32(1),
	2:  libc.Float32FromInt32(1),
	3:  libc.Float32FromInt32(1),
	4:  libc.Float32FromInt32(1),
	5:  libc.Float32FromInt32(1),
	6:  libc.Float32FromInt32(1),
	7:  libc.Float32FromInt32(2),
	8:  libc.Float32FromInt32(2),
	9:  libc.Float32FromInt32(2),
	10: libc.Float32FromInt32(2),
	11: libc.Float32FromInt32(2),
	12: libc.Float32FromInt32(2),
	13: libc.Float32FromInt32(2),
	14: libc.Float32FromInt32(3),
	15: libc.Float32FromInt32(3),
	16: libc.Float32FromInt32(4),
	17: libc.Float32FromInt32(5),
	18: libc.Float32FromInt32(6),
	19: libc.Float32FromInt32(8),
	20: libc.Float32FromInt32(8),
}

func Opus_celt_encoder_get_size(tls *libc.TLS, channels int32) (r int32) {
	var mode uintptr
	_ = mode
	mode = Opus_opus_custom_mode_create(tls, int32(48000), int32(960), libc.UintptrFromInt32(0))
	return opus_custom_encoder_get_size(tls, mode, channels)
}

func Opus_celt_encoder_init(tls *libc.TLS, st uintptr, sampling_rate OpusT_opus_int32, channels int32, arch int32) (r int32) {
	var ret int32
	_ = ret
	ret = opus_custom_encoder_init_arch(tls, st, Opus_opus_custom_mode_create(tls, int32(48000), int32(960), libc.UintptrFromInt32(0)), channels, arch)
	if ret != OPUS_OK {
		return ret
	}
	(*OpusT_OpusCustomEncoder)(unsafe.Pointer(st)).Fupsample = Opus_resampling_factor(tls, sampling_rate)
	return OPUS_OK
}

func Opus_celt_fatal(tls *libc.TLS, str uintptr, file uintptr, line int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+3861, libc.VaList(bp+8, file, line, str))
	libc.Xabort(tls)
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

var log2_x_norm_coeff8 = [8]float32{
	0: libc.Float32FromFloat32(1),
	1: libc.Float32FromFloat32(0.8888888955116272),
	2: libc.Float32FromFloat32(0.8),
	3: libc.Float32FromFloat32(0.7272727489471436),
	4: libc.Float32FromFloat32(0.6666666865348816),
	5: libc.Float32FromFloat32(0.6153846383094788),
	6: libc.Float32FromFloat32(0.5714285969734192),
	7: libc.Float32FromFloat32(0.5333333611488342),
}
var log2_y_norm_coeff8 = [8]float32{
	1: libc.Float32FromFloat32(0.1699250042438507),
	2: libc.Float32FromFloat32(0.32192808389663696),
	3: libc.Float32FromFloat32(0.45943161845207214),
	4: libc.Float32FromFloat32(0.5849624872207642),
	5: libc.Float32FromFloat32(0.7004396915435791),
	6: libc.Float32FromFloat32(0.8073549270629883),
	7: libc.Float32FromFloat32(0.9068905711174011),
}

func Opus_celt_fir_c(tls *libc.TLS, x1 uintptr, num uintptr, y1 uintptr, N int32, ord int32, arch int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _saved_stack, rnum, st, v1, v11, v13, v15, v17, v19, v21, v23, v25, v27, v3, v31, v32, v33, v35, v36, v37, v5, v7, v9 uintptr
	var i, j, j1, v34, v47, v50 int32
	var sum2 OpusT_opus_val32
	var tmp, tmp1, tmp2, tmp3, y_0, y_1, y_2, y_3 OpusT_opus_val16
	var _ /* sum at bp+0 */ [4]OpusT_opus_val32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = _saved_stack, i, j, j1, rnum, st, sum2, tmp, tmp1, tmp2, tmp3, y_0, y_1, y_2, y_3, v1, v11, v13, v15, v17, v19, v21, v23, v25, v27, v3, v31, v32, v33, v34, v35, v36, v37, v47, v5, v50, v7, v9
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _2
	_2:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _4
_4:
	_saved_stack = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack
	if !(x1 != y1) {
		Opus_celt_fatal(tls, __ccgo_ts+5930, __ccgo_ts+5955, int32(157))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _6
	_6:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _8
_8:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _10
	_10:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v7 = st
	goto _12
_12:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		goto _14
	_14:
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v11 = st
	goto _16
_16:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		goto _18
	_18:
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v15 = st
	goto _20
_20:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(ord)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5955, int32(158))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		goto _22
	_22:
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v19 = st
	goto _24
_24:
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(libc.Uint64FromInt32(ord) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		goto _26
	_26:
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v23 = st
	goto _28
_28:
	rnum = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(libc.Uint64FromInt32(ord)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	i = 0
	for {
		if !(i < ord) {
			break
		}
		*(*OpusT_opus_val16)(unsafe.Pointer(rnum + uintptr(i)*4)) = *(*OpusT_opus_val16)(unsafe.Pointer(num + uintptr(ord-i-int32(1))*4))
		goto _29
	_29:
		;
		i = i + 1
	}
	i = 0
	for {
		if !(i < N-int32(3)) {
			break
		}
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[0] = *(*OpusT_opus_val16)(unsafe.Pointer(x1 + uintptr(i)*4))
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(1)] = *(*OpusT_opus_val16)(unsafe.Pointer(x1 + uintptr(i+int32(1))*4))
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(2)] = *(*OpusT_opus_val16)(unsafe.Pointer(x1 + uintptr(i+int32(2))*4))
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)] = *(*OpusT_opus_val16)(unsafe.Pointer(x1 + uintptr(i+int32(3))*4))
		_ = arch
		v1 = rnum
		v3 = x1 + uintptr(i)*4 - uintptr(ord)*4
		v5 = bp
		v34 = ord
		if !(v34 >= libc.Int32FromInt32(3)) {
			Opus_celt_fatal(tls, __ccgo_ts+5865, __ccgo_ts+5890, int32(69))
		}
		y_3 = libc.Float32FromInt32(0)
		v7 = v3
		v3 += 4
		y_0 = *(*OpusT_opus_val16)(unsafe.Pointer(v7))
		v9 = v3
		v3 += 4
		y_1 = *(*OpusT_opus_val16)(unsafe.Pointer(v9))
		v11 = v3
		v3 += 4
		y_2 = *(*OpusT_opus_val16)(unsafe.Pointer(v11))
		j = libc.Int32FromInt32(0)
		for {
			if !(j < v34-libc.Int32FromInt32(3)) {
				break
			}
			v13 = v1
			v1 += 4
			tmp = *(*OpusT_opus_val16)(unsafe.Pointer(v13))
			v15 = v3
			v3 += 4
			y_3 = *(*OpusT_opus_val16)(unsafe.Pointer(v15))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp*y_3)
			v17 = v1
			v1 += 4
			tmp = *(*OpusT_opus_val16)(unsafe.Pointer(v17))
			v19 = v3
			v3 += 4
			y_0 = *(*OpusT_opus_val16)(unsafe.Pointer(v19))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp*y_0)
			v21 = v1
			v1 += 4
			tmp = *(*OpusT_opus_val16)(unsafe.Pointer(v21))
			v23 = v3
			v3 += 4
			y_1 = *(*OpusT_opus_val16)(unsafe.Pointer(v23))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp*y_1)
			v25 = v1
			v1 += 4
			tmp = *(*OpusT_opus_val16)(unsafe.Pointer(v25))
			v27 = v3
			v3 += 4
			y_2 = *(*OpusT_opus_val16)(unsafe.Pointer(v27))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp*y_2)
			goto _38
		_38:
			;
			j = j + int32(4)
		}
		v47 = j
		j = j + 1
		if v47 < v34 {
			v31 = v1
			v1 += 4
			tmp1 = *(*OpusT_opus_val16)(unsafe.Pointer(v31))
			v32 = v3
			v3 += 4
			y_3 = *(*OpusT_opus_val16)(unsafe.Pointer(v32))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp1*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp1*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp1*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp1*y_3)
		}
		v50 = j
		j = j + 1
		if v50 < v34 {
			v33 = v1
			v1 += 4
			tmp2 = *(*OpusT_opus_val16)(unsafe.Pointer(v33))
			v35 = v3
			v3 += 4
			y_0 = *(*OpusT_opus_val16)(unsafe.Pointer(v35))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp2*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp2*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp2*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp2*y_0)
		}
		if j < v34 {
			v36 = v1
			v1 += 4
			tmp3 = *(*OpusT_opus_val16)(unsafe.Pointer(v36))
			v37 = v3
			v3 += 4
			y_1 = *(*OpusT_opus_val16)(unsafe.Pointer(v37))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp3*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp3*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp3*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp3*y_1)
		}
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i)*4)) = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[0]
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+int32(1))*4)) = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(1)]
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+int32(2))*4)) = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(2)]
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+int32(3))*4)) = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)]
		goto _30
	_30:
		;
		i = i + int32(4)
	}
	for {
		if !(i < N) {
			break
		}
		sum2 = *(*OpusT_opus_val16)(unsafe.Pointer(x1 + uintptr(i)*4))
		j1 = 0
		for {
			if !(j1 < ord) {
				break
			}
			sum2 = sum2 + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(rnum + uintptr(j1)*4))**(*OpusT_opus_val16)(unsafe.Pointer(x1 + uintptr(i+j1-ord)*4)))
			goto _56
		_56:
			;
			j1 = j1 + 1
		}
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i)*4)) = sum2
		goto _55
	_55:
		;
		i = i + 1
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _58
	_58:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _60
_60:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
}

func Opus_celt_float2int16_c(tls *libc.TLS, in uintptr, out uintptr, cnt int32) {
	var i int32
	var v2, v3, v4 float32
	var v5 OpusT_opus_int16
	_, _, _, _, _ = i, v2, v3, v4, v5
	i = 0
	for {
		if !(i < cnt) {
			break
		}
		v2 = *(*float32)(unsafe.Pointer(in + uintptr(i)*4))
		v2 = float32(v2 * libc.Float32FromFloat32(32768))
		if v2 > float32(-libc.Int32FromInt32(32768)) {
			v3 = v2
		} else {
			v3 = float32(-libc.Int32FromInt32(32768))
		}
		v2 = v3
		if v2 < float32(libc.Int32FromInt32(32767)) {
			v4 = v2
		} else {
			v4 = float32(libc.Int32FromInt32(32767))
		}
		v2 = v4
		v5 = int16(Opus_lrintf(tls, v2))
		goto _6
	_6:
		*(*int16)(unsafe.Pointer(out + uintptr(i)*2)) = v5
		goto _1
	_1:
		;
		i = i + 1
	}
}

func Opus_celt_iir(tls *libc.TLS, _x uintptr, den uintptr, _y uintptr, N int32, ord int32, mem uintptr, arch int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _saved_stack, rden, st, y1, v1, v11, v13, v15, v17, v19, v21, v23, v25, v27, v29, v3, v31, v33, v35, v37, v39, v5, v7, v9 uintptr
	var i, j, j1, v60, v73, v76 int32
	var sum2 OpusT_opus_val32
	var tmp, tmp1, tmp2, tmp3, y_0, y_1, y_2, y_3 OpusT_opus_val16
	var _ /* sum at bp+0 */ [4]OpusT_opus_val32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = _saved_stack, i, j, j1, rden, st, sum2, tmp, tmp1, tmp2, tmp3, y1, y_0, y_1, y_2, y_3, v1, v11, v13, v15, v17, v19, v21, v23, v25, v27, v29, v3, v31, v33, v35, v37, v39, v5, v60, v7, v73, v76, v9
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _2
	_2:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _4
_4:
	_saved_stack = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack
	if !(ord&libc.Int32FromInt32(3) == libc.Int32FromInt32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+5974, __ccgo_ts+5955, int32(225))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _6
	_6:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _8
_8:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _10
	_10:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v7 = st
	goto _12
_12:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		goto _14
	_14:
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v11 = st
	goto _16
_16:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		goto _18
	_18:
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v15 = st
	goto _20
_20:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(ord)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5955, int32(226))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		goto _22
	_22:
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v19 = st
	goto _24
_24:
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(libc.Uint64FromInt32(ord) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		goto _26
	_26:
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v23 = st
	goto _28
_28:
	rden = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(libc.Uint64FromInt32(ord)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _30
	_30:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _32
_32:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _34
	_34:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v7 = st
	goto _36
_36:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		goto _38
	_38:
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v11 = st
	goto _40
_40:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		goto _42
	_42:
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v15 = st
	goto _44
_44:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(N+ord)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+5955, int32(227))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		goto _46
	_46:
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v19 = st
	goto _48
_48:
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(libc.Uint64FromInt32(N+ord) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		goto _50
	_50:
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v23 = st
	goto _52
_52:
	y1 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(libc.Uint64FromInt32(N+ord)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	i = 0
	for {
		if !(i < ord) {
			break
		}
		*(*OpusT_opus_val16)(unsafe.Pointer(rden + uintptr(i)*4)) = *(*OpusT_opus_val16)(unsafe.Pointer(den + uintptr(ord-i-int32(1))*4))
		goto _53
	_53:
		;
		i = i + 1
	}
	i = 0
	for {
		if !(i < ord) {
			break
		}
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i)*4)) = -*(*OpusT_opus_val16)(unsafe.Pointer(mem + uintptr(ord-i-int32(1))*4))
		goto _54
	_54:
		;
		i = i + 1
	}
	for {
		if !(i < N+ord) {
			break
		}
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i)*4)) = libc.Float32FromInt32(0)
		goto _55
	_55:
		;
		i = i + 1
	}
	i = 0
	for {
		if !(i < N-int32(3)) {
			break
		}
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[0] = *(*OpusT_opus_val32)(unsafe.Pointer(_x + uintptr(i)*4))
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(1)] = *(*OpusT_opus_val32)(unsafe.Pointer(_x + uintptr(i+int32(1))*4))
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(2)] = *(*OpusT_opus_val32)(unsafe.Pointer(_x + uintptr(i+int32(2))*4))
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)] = *(*OpusT_opus_val32)(unsafe.Pointer(_x + uintptr(i+int32(3))*4))
		_ = arch
		v1 = rden
		v3 = y1 + uintptr(i)*4
		v5 = bp
		v60 = ord
		if !(v60 >= libc.Int32FromInt32(3)) {
			Opus_celt_fatal(tls, __ccgo_ts+5865, __ccgo_ts+5890, int32(69))
		}
		y_3 = libc.Float32FromInt32(0)
		v7 = v3
		v3 += 4
		y_0 = *(*OpusT_opus_val16)(unsafe.Pointer(v7))
		v9 = v3
		v3 += 4
		y_1 = *(*OpusT_opus_val16)(unsafe.Pointer(v9))
		v11 = v3
		v3 += 4
		y_2 = *(*OpusT_opus_val16)(unsafe.Pointer(v11))
		j = libc.Int32FromInt32(0)
		for {
			if !(j < v60-libc.Int32FromInt32(3)) {
				break
			}
			v13 = v1
			v1 += 4
			tmp = *(*OpusT_opus_val16)(unsafe.Pointer(v13))
			v15 = v3
			v3 += 4
			y_3 = *(*OpusT_opus_val16)(unsafe.Pointer(v15))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp*y_3)
			v17 = v1
			v1 += 4
			tmp = *(*OpusT_opus_val16)(unsafe.Pointer(v17))
			v19 = v3
			v3 += 4
			y_0 = *(*OpusT_opus_val16)(unsafe.Pointer(v19))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp*y_0)
			v21 = v1
			v1 += 4
			tmp = *(*OpusT_opus_val16)(unsafe.Pointer(v21))
			v23 = v3
			v3 += 4
			y_1 = *(*OpusT_opus_val16)(unsafe.Pointer(v23))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp*y_1)
			v25 = v1
			v1 += 4
			tmp = *(*OpusT_opus_val16)(unsafe.Pointer(v25))
			v27 = v3
			v3 += 4
			y_2 = *(*OpusT_opus_val16)(unsafe.Pointer(v27))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp*y_2)
			goto _64
		_64:
			;
			j = j + int32(4)
		}
		v73 = j
		j = j + 1
		if v73 < v60 {
			v29 = v1
			v1 += 4
			tmp1 = *(*OpusT_opus_val16)(unsafe.Pointer(v29))
			v31 = v3
			v3 += 4
			y_3 = *(*OpusT_opus_val16)(unsafe.Pointer(v31))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp1*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp1*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp1*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp1*y_3)
		}
		v76 = j
		j = j + 1
		if v76 < v60 {
			v33 = v1
			v1 += 4
			tmp2 = *(*OpusT_opus_val16)(unsafe.Pointer(v33))
			v35 = v3
			v3 += 4
			y_0 = *(*OpusT_opus_val16)(unsafe.Pointer(v35))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp2*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp2*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp2*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp2*y_0)
		}
		if j < v60 {
			v37 = v1
			v1 += 4
			tmp3 = *(*OpusT_opus_val16)(unsafe.Pointer(v37))
			v39 = v3
			v3 += 4
			y_1 = *(*OpusT_opus_val16)(unsafe.Pointer(v39))
			*(*OpusT_opus_val32)(unsafe.Pointer(v5)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5)) + OpusT_opus_val32(tmp3*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 1*4)) + OpusT_opus_val32(tmp3*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 2*4)) + OpusT_opus_val32(tmp3*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v5 + 3*4)) + OpusT_opus_val32(tmp3*y_1)
		}
		/* Patch up the result to compensate for the fact that this is an IIR */
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+ord)*4)) = -(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[0]
		*(*OpusT_opus_val32)(unsafe.Pointer(_y + uintptr(i)*4)) = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[0]
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(1)] = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(1)] + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+ord)*4))**(*OpusT_opus_val16)(unsafe.Pointer(den)))
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+ord+int32(1))*4)) = -(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(1)]
		*(*OpusT_opus_val32)(unsafe.Pointer(_y + uintptr(i+int32(1))*4)) = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(1)]
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(2)] = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(2)] + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+ord+int32(1))*4))**(*OpusT_opus_val16)(unsafe.Pointer(den)))
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(2)] = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(2)] + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+ord)*4))**(*OpusT_opus_val16)(unsafe.Pointer(den + 1*4)))
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+ord+int32(2))*4)) = -(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(2)]
		*(*OpusT_opus_val32)(unsafe.Pointer(_y + uintptr(i+int32(2))*4)) = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(2)]
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)] = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)] + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+ord+int32(2))*4))**(*OpusT_opus_val16)(unsafe.Pointer(den)))
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)] = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)] + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+ord+int32(1))*4))**(*OpusT_opus_val16)(unsafe.Pointer(den + 1*4)))
		(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)] = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)] + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+ord)*4))**(*OpusT_opus_val16)(unsafe.Pointer(den + 2*4)))
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+ord+int32(3))*4)) = -(*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)]
		*(*OpusT_opus_val32)(unsafe.Pointer(_y + uintptr(i+int32(3))*4)) = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)]
		goto _56
	_56:
		;
		i = i + int32(4)
	}
	for {
		if !(i < N) {
			break
		}
		sum2 = *(*OpusT_opus_val32)(unsafe.Pointer(_x + uintptr(i)*4))
		j1 = 0
		for {
			if !(j1 < ord) {
				break
			}
			sum2 = sum2 - OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(rden + uintptr(j1)*4))**(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+j1)*4)))
			goto _82
		_82:
			;
			j1 = j1 + 1
		}
		*(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i+ord)*4)) = sum2
		*(*OpusT_opus_val32)(unsafe.Pointer(_y + uintptr(i)*4)) = sum2
		goto _81
	_81:
		;
		i = i + 1
	}
	i = 0
	for {
		if !(i < ord) {
			break
		}
		*(*OpusT_opus_val16)(unsafe.Pointer(mem + uintptr(i)*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(_y + uintptr(N-i-int32(1))*4))
		goto _83
	_83:
		;
		i = i + 1
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _85
	_85:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _87
_87:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
}

func Opus_celt_lcg_rand(tls *libc.TLS, seed OpusT_opus_uint32) (r OpusT_opus_uint32) {
	return uint32(1664525)*seed + uint32(1013904223)
}

// C documentation
//
//	/* This is a cos() approximation designed to be bit-exact on any platform. Bit exactness
//	   with this approximation is important because it has an impact on the bit allocation */

func Opus_celt_pitch_xcorr_c(tls *libc.TLS, _x uintptr, _y uintptr, xcorr uintptr, len1 int32, max_pitch int32, arch int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var i, i1, j, v18, v21, v5 int32
	var sum2, xy, v28 OpusT_opus_val32
	var tmp, tmp1, tmp2, tmp3, y_0, y_1, y_2, y_3 OpusT_opus_val16
	var v10, v11, v12, v13, v14, v15, v16, v17, v19, v2, v20, v22, v23, v24, v25, v3, v4, v6, v7, v8 uintptr
	var _ /* sum at bp+0 */ [4]OpusT_opus_val32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = i, i1, j, sum2, tmp, tmp1, tmp2, tmp3, xy, y_0, y_1, y_2, y_3, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v2, v20, v21, v22, v23, v24, v25, v28, v3, v4, v5, v6, v7, v8
	/*The EDSP version requires that max_pitch is at least 1, and that _x is
	  32-bit aligned.
	 Since it's hard to put asserts in assembly, put them here.*/
	if !(max_pitch > libc.Int32FromInt32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+5819, __ccgo_ts+5849, int32(265))
	}
	_ = uint64(_x)&libc.Uint64FromInt32(3) == libc.Uint64FromInt32(0)
	i1 = 0
	for {
		if !(i1 < max_pitch-int32(3)) {
			break
		}
		*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)) = [4]OpusT_opus_val32{}
		_ = arch
		v2 = _x
		v3 = _y + uintptr(i1)*4
		v4 = bp
		v5 = len1
		if !(v5 >= libc.Int32FromInt32(3)) {
			Opus_celt_fatal(tls, __ccgo_ts+5865, __ccgo_ts+5890, int32(69))
		}
		y_3 = libc.Float32FromInt32(0)
		v6 = v3
		v3 += 4
		y_0 = *(*OpusT_opus_val16)(unsafe.Pointer(v6))
		v7 = v3
		v3 += 4
		y_1 = *(*OpusT_opus_val16)(unsafe.Pointer(v7))
		v8 = v3
		v3 += 4
		y_2 = *(*OpusT_opus_val16)(unsafe.Pointer(v8))
		j = libc.Int32FromInt32(0)
		for {
			if !(j < v5-libc.Int32FromInt32(3)) {
				break
			}
			v10 = v2
			v2 += 4
			tmp = *(*OpusT_opus_val16)(unsafe.Pointer(v10))
			v11 = v3
			v3 += 4
			y_3 = *(*OpusT_opus_val16)(unsafe.Pointer(v11))
			*(*OpusT_opus_val32)(unsafe.Pointer(v4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4)) + OpusT_opus_val32(tmp*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v4 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4 + 1*4)) + OpusT_opus_val32(tmp*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v4 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4 + 2*4)) + OpusT_opus_val32(tmp*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v4 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4 + 3*4)) + OpusT_opus_val32(tmp*y_3)
			v12 = v2
			v2 += 4
			tmp = *(*OpusT_opus_val16)(unsafe.Pointer(v12))
			v13 = v3
			v3 += 4
			y_0 = *(*OpusT_opus_val16)(unsafe.Pointer(v13))
			*(*OpusT_opus_val32)(unsafe.Pointer(v4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4)) + OpusT_opus_val32(tmp*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v4 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4 + 1*4)) + OpusT_opus_val32(tmp*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v4 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4 + 2*4)) + OpusT_opus_val32(tmp*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v4 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4 + 3*4)) + OpusT_opus_val32(tmp*y_0)
			v14 = v2
			v2 += 4
			tmp = *(*OpusT_opus_val16)(unsafe.Pointer(v14))
			v15 = v3
			v3 += 4
			y_1 = *(*OpusT_opus_val16)(unsafe.Pointer(v15))
			*(*OpusT_opus_val32)(unsafe.Pointer(v4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4)) + OpusT_opus_val32(tmp*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v4 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4 + 1*4)) + OpusT_opus_val32(tmp*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v4 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4 + 2*4)) + OpusT_opus_val32(tmp*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v4 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4 + 3*4)) + OpusT_opus_val32(tmp*y_1)
			v16 = v2
			v2 += 4
			tmp = *(*OpusT_opus_val16)(unsafe.Pointer(v16))
			v17 = v3
			v3 += 4
			y_2 = *(*OpusT_opus_val16)(unsafe.Pointer(v17))
			*(*OpusT_opus_val32)(unsafe.Pointer(v4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4)) + OpusT_opus_val32(tmp*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v4 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4 + 1*4)) + OpusT_opus_val32(tmp*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v4 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4 + 2*4)) + OpusT_opus_val32(tmp*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v4 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4 + 3*4)) + OpusT_opus_val32(tmp*y_2)
			goto _9
		_9:
			;
			j = j + int32(4)
		}
		v18 = j
		j = j + 1
		if v18 < v5 {
			v19 = v2
			v2 += 4
			tmp1 = *(*OpusT_opus_val16)(unsafe.Pointer(v19))
			v20 = v3
			v3 += 4
			y_3 = *(*OpusT_opus_val16)(unsafe.Pointer(v20))
			*(*OpusT_opus_val32)(unsafe.Pointer(v4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4)) + OpusT_opus_val32(tmp1*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v4 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4 + 1*4)) + OpusT_opus_val32(tmp1*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v4 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4 + 2*4)) + OpusT_opus_val32(tmp1*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v4 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4 + 3*4)) + OpusT_opus_val32(tmp1*y_3)
		}
		v21 = j
		j = j + 1
		if v21 < v5 {
			v22 = v2
			v2 += 4
			tmp2 = *(*OpusT_opus_val16)(unsafe.Pointer(v22))
			v23 = v3
			v3 += 4
			y_0 = *(*OpusT_opus_val16)(unsafe.Pointer(v23))
			*(*OpusT_opus_val32)(unsafe.Pointer(v4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4)) + OpusT_opus_val32(tmp2*y_1)
			*(*OpusT_opus_val32)(unsafe.Pointer(v4 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4 + 1*4)) + OpusT_opus_val32(tmp2*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v4 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4 + 2*4)) + OpusT_opus_val32(tmp2*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v4 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4 + 3*4)) + OpusT_opus_val32(tmp2*y_0)
		}
		if j < v5 {
			v24 = v2
			v2 += 4
			tmp3 = *(*OpusT_opus_val16)(unsafe.Pointer(v24))
			v25 = v3
			v3 += 4
			y_1 = *(*OpusT_opus_val16)(unsafe.Pointer(v25))
			*(*OpusT_opus_val32)(unsafe.Pointer(v4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4)) + OpusT_opus_val32(tmp3*y_2)
			*(*OpusT_opus_val32)(unsafe.Pointer(v4 + 1*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4 + 1*4)) + OpusT_opus_val32(tmp3*y_3)
			*(*OpusT_opus_val32)(unsafe.Pointer(v4 + 2*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4 + 2*4)) + OpusT_opus_val32(tmp3*y_0)
			*(*OpusT_opus_val32)(unsafe.Pointer(v4 + 3*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(v4 + 3*4)) + OpusT_opus_val32(tmp3*y_1)
		}
		*(*OpusT_opus_val32)(unsafe.Pointer(xcorr + uintptr(i1)*4)) = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[0]
		*(*OpusT_opus_val32)(unsafe.Pointer(xcorr + uintptr(i1+int32(1))*4)) = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(1)]
		*(*OpusT_opus_val32)(unsafe.Pointer(xcorr + uintptr(i1+int32(2))*4)) = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(2)]
		*(*OpusT_opus_val32)(unsafe.Pointer(xcorr + uintptr(i1+int32(3))*4)) = (*(*[4]OpusT_opus_val32)(unsafe.Pointer(bp)))[int32(3)]
		goto _1
	_1:
		;
		i1 = i1 + int32(4)
	}
	/* In case max_pitch isn't a multiple of 4, do non-unrolled version. */
	for {
		if !(i1 < max_pitch) {
			break
		}
		_ = arch
		xy = libc.Float32FromInt32(0)
		i = libc.Int32FromInt32(0)
		for {
			if !(i < len1) {
				break
			}
			xy = xy + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(_x + uintptr(i)*4))**(*OpusT_opus_val16)(unsafe.Pointer(_y + uintptr(i1)*4 + uintptr(i)*4)))
			goto _27
		_27:
			;
			i = i + 1
		}
		v28 = xy
		goto _29
	_29:
		sum2 = v28
		*(*OpusT_opus_val32)(unsafe.Pointer(xcorr + uintptr(i1)*4)) = sum2
		goto _26
	_26:
		;
		i1 = i1 + 1
	}
}

func Opus_celt_preemphasis(tls *libc.TLS, pcmp uintptr, inp uintptr, N int32, CC int32, upsample int32, coef uintptr, mem uintptr, clip int32) {
	var Nu, i int32
	var coef0 OpusT_opus_val16
	var m, x, x1 OpusT_celt_sig
	var v4, v5, v6 float32
	_, _, _, _, _, _, _, _, _ = Nu, coef0, i, m, x, x1, v4, v5, v6
	coef0 = *(*OpusT_opus_val16)(unsafe.Pointer(coef))
	m = *(*OpusT_celt_sig)(unsafe.Pointer(mem))
	/* Fast path for the normal 48kHz case and no clipping */
	if *(*OpusT_opus_val16)(unsafe.Pointer(coef + 1*4)) == libc.Float32FromInt32(0) && upsample == int32(1) && !(clip != 0) {
		i = 0
		for {
			if !(i < N) {
				break
			}
			x = float32(libc.Float32FromFloat32(32768) * *(*OpusT_opus_res)(unsafe.Pointer(pcmp + uintptr(CC*i)*4)))
			/* Apply pre-emphasis */
			*(*OpusT_celt_sig)(unsafe.Pointer(inp + uintptr(i)*4)) = x - m
			m = OpusT_opus_val16(coef0 * x)
			goto _1
		_1:
			;
			i = i + 1
		}
		*(*OpusT_celt_sig)(unsafe.Pointer(mem)) = m
		return
	}
	Nu = N / upsample
	if upsample != int32(1) {
		libc.Xmemset(tls, inp, 0, libc.Uint64FromInt32(N)*uint64(4))
	}
	i = 0
	for {
		if !(i < Nu) {
			break
		}
		*(*OpusT_celt_sig)(unsafe.Pointer(inp + uintptr(i*upsample)*4)) = float32(libc.Float32FromFloat32(32768) * *(*OpusT_opus_res)(unsafe.Pointer(pcmp + uintptr(CC*i)*4)))
		goto _2
	_2:
		;
		i = i + 1
	}
	if clip != 0 {
		/* Clip input to avoid encoding non-portable files */
		i = 0
		for {
			if !(i < Nu) {
				break
			}
			if libc.Float32FromFloat32(65536) < *(*OpusT_celt_sig)(unsafe.Pointer(inp + uintptr(i*upsample)*4)) {
				v5 = libc.Float32FromFloat32(65536)
			} else {
				v5 = *(*OpusT_celt_sig)(unsafe.Pointer(inp + uintptr(i*upsample)*4))
			}
			if -libc.Float32FromFloat32(65536) > v5 {
				v4 = -libc.Float32FromFloat32(65536)
			} else {
				if libc.Float32FromFloat32(65536) < *(*OpusT_celt_sig)(unsafe.Pointer(inp + uintptr(i*upsample)*4)) {
					v6 = libc.Float32FromFloat32(65536)
				} else {
					v6 = *(*OpusT_celt_sig)(unsafe.Pointer(inp + uintptr(i*upsample)*4))
				}
				v4 = v6
			}
			*(*OpusT_celt_sig)(unsafe.Pointer(inp + uintptr(i*upsample)*4)) = v4
			goto _3
		_3:
			;
			i = i + 1
		}
	}
	i = 0
	for {
		if !(i < N) {
			break
		}
		x1 = *(*OpusT_celt_sig)(unsafe.Pointer(inp + uintptr(i)*4))
		/* Apply pre-emphasis */
		*(*OpusT_celt_sig)(unsafe.Pointer(inp + uintptr(i)*4)) = x1 - m
		m = OpusT_opus_val16(coef0 * x1)
		goto _7
	_7:
		;
		i = i + 1
	}
	*(*OpusT_celt_sig)(unsafe.Pointer(mem)) = m
}

func Opus_validate_celt_decoder(tls *libc.TLS, st uintptr) {
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fmode == Opus_opus_custom_mode_create(tls, int32(48000), int32(960), libc.UintptrFromInt32(0))) {
		Opus_celt_fatal(tls, __ccgo_ts+4348, __ccgo_ts+4420, int32(147))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Foverlap == libc.Int32FromInt32(120)) {
		Opus_celt_fatal(tls, __ccgo_ts+4443, __ccgo_ts+4420, int32(148))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fend <= libc.Int32FromInt32(21)) {
		Opus_celt_fatal(tls, __ccgo_ts+4480, __ccgo_ts+4420, int32(149))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fchannels == int32(1) || (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fchannels == int32(2)) {
		Opus_celt_fatal(tls, __ccgo_ts, __ccgo_ts+4420, int32(157))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fstream_channels == int32(1) || (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fstream_channels == int32(2)) {
		Opus_celt_fatal(tls, __ccgo_ts+925, __ccgo_ts+4420, int32(158))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fdownsample > libc.Int32FromInt32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+4512, __ccgo_ts+4420, int32(159))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fstart == 0 || (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fstart == int32(17)) {
		Opus_celt_fatal(tls, __ccgo_ts+4549, __ccgo_ts+4420, int32(160))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fstart < (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fend) {
		Opus_celt_fatal(tls, __ccgo_ts+4601, __ccgo_ts+4420, int32(161))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Farch >= libc.Int32FromInt32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+849, __ccgo_ts+4420, int32(163))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Farch <= libc.Int32FromInt32(OPUS_ARCHMASK)) {
		Opus_celt_fatal(tls, __ccgo_ts+881, __ccgo_ts+4420, int32(164))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Flast_pitch_index <= libc.Int32FromInt32(PLC_PITCH_LAG_MAX)) {
		Opus_celt_fatal(tls, __ccgo_ts+4639, __ccgo_ts+4420, int32(167))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Flast_pitch_index >= int32(PLC_PITCH_LAG_MIN) || (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Flast_pitch_index == 0) {
		Opus_celt_fatal(tls, __ccgo_ts+4699, __ccgo_ts+4420, int32(168))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fpostfilter_period < libc.Int32FromInt32(MAX_PERIOD)) {
		Opus_celt_fatal(tls, __ccgo_ts+4788, __ccgo_ts+4420, int32(170))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fpostfilter_period >= int32(COMBFILTER_MINPERIOD) || (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fpostfilter_period == 0) {
		Opus_celt_fatal(tls, __ccgo_ts+4841, __ccgo_ts+4420, int32(171))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fpostfilter_period_old < libc.Int32FromInt32(MAX_PERIOD)) {
		Opus_celt_fatal(tls, __ccgo_ts+4935, __ccgo_ts+4420, int32(172))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fpostfilter_period_old >= int32(COMBFILTER_MINPERIOD) || (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fpostfilter_period_old == 0) {
		Opus_celt_fatal(tls, __ccgo_ts+4992, __ccgo_ts+4420, int32(173))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fpostfilter_tapset <= libc.Int32FromInt32(2)) {
		Opus_celt_fatal(tls, __ccgo_ts+5094, __ccgo_ts+4420, int32(174))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fpostfilter_tapset >= libc.Int32FromInt32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+5139, __ccgo_ts+4420, int32(175))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fpostfilter_tapset_old <= libc.Int32FromInt32(2)) {
		Opus_celt_fatal(tls, __ccgo_ts+5184, __ccgo_ts+4420, int32(176))
	}
	if !((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fpostfilter_tapset_old >= libc.Int32FromInt32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+5233, __ccgo_ts+4420, int32(177))
	}
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
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _2
	_2:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
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
	lpc = backgroundLogE + uintptr(libc.Int32FromInt32(2)*nbEBands)*4
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
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			goto _11
		_11:
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v3 = st
		goto _13
	_13:
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v10 = libc.Xmalloc(tls, uint64(16))
			goto _15
		_15:
			st = v10
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v12 = st
		goto _17
	_17:
		*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v12)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v14 = libc.Xmalloc(tls, uint64(16))
			goto _19
		_19:
			st = v14
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v16 = st
		goto _21
	_21:
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v18 = libc.Xmalloc(tls, uint64(16))
			goto _23
		_23:
			st = v18
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v20 = st
		goto _25
	_25:
		if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(C*N)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v16)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v20)).Fglobal_stack)) {
			Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4420, int32(749))
		}
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v22 = libc.Xmalloc(tls, uint64(16))
			goto _27
		_27:
			st = v22
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v24 = st
		goto _29
	_29:
		*(*uintptr)(unsafe.Pointer(v24 + 8)) += uintptr(libc.Uint64FromInt32(C*N) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v26 = libc.Xmalloc(tls, uint64(16))
			goto _31
		_31:
			st = v26
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v28 = st
		goto _33
	_33:
		X = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v28)).Fglobal_stack - uintptr(libc.Uint64FromInt32(C*N)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1))) /**< Interleaved normalised MDCTs */
		c = 0
		for {
			libc.Xmemmove(tls, (*(*[2]uintptr)(unsafe.Pointer(bp)))[c], (*(*[2]uintptr)(unsafe.Pointer(bp)))[c]+uintptr(N)*4, libc.Uint64FromInt32(decode_buffer_size-N+overlap)*uint64(4)+libc.Uint64FromInt64(0*((int64((*(*[2]uintptr)(unsafe.Pointer(bp)))[c])-int64((*(*[2]uintptr)(unsafe.Pointer(bp)))[c]+uintptr(N)*4))/4)))
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
			v36 = libc.Float32FromFloat32(1.5)
		} else {
			v36 = libc.Float32FromFloat32(0.5)
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
					*(*OpusT_celt_norm)(unsafe.Pointer(X + uintptr(boffs+j)*4)) = float32(libc.Int32FromUint32(seed) >> libc.Int32FromInt32(20))
					goto _43
				_43:
					;
					j = j + 1
				}
				Opus_renormalise_vector(tls, X+uintptr(boffs)*4, blen, libc.Float32FromFloat32(1), (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
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
		fade = libc.Float32FromFloat32(1)
		curr_neural = libc.BoolInt32(curr_frame_type == int32(FRAME_PLC_NEURAL) || curr_frame_type == int32(FRAME_DRED))
		last_neural = libc.BoolInt32((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Flast_frame_type == int32(FRAME_PLC_NEURAL) || (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Flast_frame_type == int32(FRAME_DRED))
		if (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Flast_frame_type != int32(FRAME_PLC_PERIODIC) && !(last_neural != 0 && curr_neural != 0) {
			v5 = celt_plc_pitch_search(tls, st1, bp, C, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Farch)
			pitch_index = v5
			(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Flast_pitch_index = v5
		} else {
			pitch_index = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st1)).Flast_pitch_index
			fade = libc.Float32FromFloat32(0.8)
		}
		/* We want the excitation for 2 pitch periods in order to look for a
		   decaying signal, but we can't get more than MAX_PERIOD. */
		if int32(2)*pitch_index < max_period {
			v5 = int32(2) * pitch_index
		} else {
			v5 = max_period
		}
		exc_length = v5
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			goto _51
		_51:
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v3 = st
		goto _53
	_53:
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v10 = libc.Xmalloc(tls, uint64(16))
			goto _55
		_55:
			st = v10
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v12 = st
		goto _57
	_57:
		*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v12)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v14 = libc.Xmalloc(tls, uint64(16))
			goto _59
		_59:
			st = v14
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v16 = st
		goto _61
	_61:
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v18 = libc.Xmalloc(tls, uint64(16))
			goto _63
		_63:
			st = v18
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v20 = st
		goto _65
	_65:
		if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(max_period+libc.Int32FromInt32(CELT_LPC_ORDER))*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v16)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v20)).Fglobal_stack)) {
			Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4420, int32(837))
		}
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v22 = libc.Xmalloc(tls, uint64(16))
			goto _67
		_67:
			st = v22
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v24 = st
		goto _69
	_69:
		*(*uintptr)(unsafe.Pointer(v24 + 8)) += uintptr(libc.Uint64FromInt32(max_period+libc.Int32FromInt32(CELT_LPC_ORDER)) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v26 = libc.Xmalloc(tls, uint64(16))
			goto _71
		_71:
			st = v26
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v28 = st
		goto _73
	_73:
		_exc = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v28)).Fglobal_stack - uintptr(libc.Uint64FromInt32(max_period+libc.Int32FromInt32(CELT_LPC_ORDER))*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			goto _75
		_75:
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v3 = st
		goto _77
	_77:
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v10 = libc.Xmalloc(tls, uint64(16))
			goto _79
		_79:
			st = v10
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v12 = st
		goto _81
	_81:
		*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v12)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v14 = libc.Xmalloc(tls, uint64(16))
			goto _83
		_83:
			st = v14
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v16 = st
		goto _85
	_85:
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v18 = libc.Xmalloc(tls, uint64(16))
			goto _87
		_87:
			st = v18
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v20 = st
		goto _89
	_89:
		if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(exc_length)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v16)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v20)).Fglobal_stack)) {
			Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4420, int32(838))
		}
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v22 = libc.Xmalloc(tls, uint64(16))
			goto _91
		_91:
			st = v22
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v24 = st
		goto _93
	_93:
		*(*uintptr)(unsafe.Pointer(v24 + 8)) += uintptr(libc.Uint64FromInt32(exc_length) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
		st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
		if !(st != 0) {
			v26 = libc.Xmalloc(tls, uint64(16))
			goto _95
		_95:
			st = v26
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
		}
		v28 = st
		goto _97
	_97:
		fir_tmp = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v28)).Fglobal_stack - uintptr(libc.Uint64FromInt32(exc_length)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
		exc = _exc + uintptr(CELT_LPC_ORDER)*4
		window = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Fwindow
		c = 0
		for {
			S1 = libc.Float32FromInt32(0)
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
				*(*OpusT_opus_val32)(unsafe.Pointer(bp + 32)) *= libc.Float32FromFloat32(1.0001)
				/* Use lag windowing to stabilize the Levinson-Durbin recursion. */
				i = int32(1)
				for {
					if !(i <= int32(CELT_LPC_ORDER)) {
						break
					}
					/*ac[i] *= exp(-.5*(2*M_PI*.002*i)*(2*M_PI*.002*i));*/
					*(*OpusT_opus_val32)(unsafe.Pointer(bp + 32 + uintptr(i)*4)) -= OpusT_opus_val32(OpusT_opus_val32(OpusT_opus_val32((*(*[25]OpusT_opus_val32)(unsafe.Pointer(bp + 32)))[i]*float32(libc.Float32FromFloat32(0.008)*libc.Float32FromFloat32(0.008)))*float32(i)) * float32(i))
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
			libc.Xmemcpy(tls, exc+uintptr(max_period)*4-uintptr(exc_length)*4, fir_tmp, libc.Uint64FromInt32(exc_length)*uint64(4)+libc.Uint64FromInt64(0*((int64(exc+uintptr(max_period)*4-uintptr(exc_length)*4)-int64(fir_tmp))/4)))
			/* Check if the waveform is decaying, and if so how fast.
			   We do this to avoid adding energy when concealing in a segment
			   with decaying energy. */
			E1 = libc.Float32FromInt32(1)
			E2 = libc.Float32FromInt32(1)
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
			libc.Xmemmove(tls, buf, buf+uintptr(N)*4, libc.Uint64FromInt32(decode_buffer_size-N)*uint64(4)+libc.Uint64FromInt64(0*((int64(buf)-int64(buf+uintptr(N)*4))/4)))
			/* Extrapolate from the end of the excitation with a period of
			   "pitch_index", scaling down each period by an additional factor of
			   "decay". */
			extrapolation_offset = max_period - pitch_index
			/* We need to extrapolate enough samples to cover a complete MDCT
			   window (including overlap/2 samples on both sides). */
			extrapolation_len = N + overlap
			/* We also apply fading if this is not the first loss. */
			attenuation = OpusT_opus_val16(fade * decay1)
			v5 = libc.Int32FromInt32(0)
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
			S2 = libc.Float32FromInt32(0)
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
			if !(S1 > OpusT_opus_val32(libc.Float32FromFloat32(0.2)*S2)) {
				i = 0
				for {
					if !(i < extrapolation_len) {
						break
					}
					*(*OpusT_celt_sig)(unsafe.Pointer(buf + uintptr(decode_buffer_size-N+i)*4)) = libc.Float32FromInt32(0)
					goto _108
				_108:
					;
					i = i + 1
				}
			} else {
				if S1 < S2 {
					ratio = float32(libc.Xsqrt(tls, float64((S1+libc.Float32FromInt32(1))/(S2+libc.Float32FromInt32(1)))))
					i = 0
					for {
						if !(i < overlap) {
							break
						}
						tmp_g = libc.Float32FromFloat32(1) - OpusT_celt_coef(*(*OpusT_celt_coef)(unsafe.Pointer(window + uintptr(i)*4))*(libc.Float32FromFloat32(1)-ratio))
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
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _114
	_114:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _116
_116:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
}

func celt_fir5(tls *libc.TLS, x uintptr, num uintptr, N int32) {
	var i int32
	var mem0, mem1, mem2, mem3, mem4, sum OpusT_opus_val32
	var num0, num1, num2, num3, num4 OpusT_opus_val16
	_, _, _, _, _, _, _, _, _, _, _, _ = i, mem0, mem1, mem2, mem3, mem4, num0, num1, num2, num3, num4, sum
	num0 = *(*OpusT_opus_val16)(unsafe.Pointer(num))
	num1 = *(*OpusT_opus_val16)(unsafe.Pointer(num + 1*4))
	num2 = *(*OpusT_opus_val16)(unsafe.Pointer(num + 2*4))
	num3 = *(*OpusT_opus_val16)(unsafe.Pointer(num + 3*4))
	num4 = *(*OpusT_opus_val16)(unsafe.Pointer(num + 4*4))
	mem0 = libc.Float32FromInt32(0)
	mem1 = libc.Float32FromInt32(0)
	mem2 = libc.Float32FromInt32(0)
	mem3 = libc.Float32FromInt32(0)
	mem4 = libc.Float32FromInt32(0)
	i = 0
	for {
		if !(i < N) {
			break
		}
		sum = *(*OpusT_opus_val16)(unsafe.Pointer(x + uintptr(i)*4))
		sum = sum + OpusT_opus_val32(num0*mem0)
		sum = sum + OpusT_opus_val32(num1*mem1)
		sum = sum + OpusT_opus_val32(num2*mem2)
		sum = sum + OpusT_opus_val32(num3*mem3)
		sum = sum + OpusT_opus_val32(num4*mem4)
		mem4 = mem3
		mem3 = mem2
		mem2 = mem1
		mem1 = mem0
		mem0 = *(*OpusT_opus_val16)(unsafe.Pointer(x + uintptr(i)*4))
		*(*OpusT_opus_val16)(unsafe.Pointer(x + uintptr(i)*4)) = sum
		goto _1
	_1:
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
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _2
	_2:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _4
_4:
	_saved_stack = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack
	_ = st1
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _6
	_6:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _8
_8:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _10
	_10:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v7 = st
	goto _12
_12:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		goto _14
	_14:
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v11 = st
	goto _16
_16:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		goto _18
	_18:
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v15 = st
	goto _20
_20:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(libc.Int32FromInt32(DEC_PITCH_BUF_SIZE)>>libc.Int32FromInt32(1))*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4420, int32(565))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		goto _22
	_22:
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v19 = st
	goto _24
_24:
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(libc.Uint64FromInt32(libc.Int32FromInt32(DEC_PITCH_BUF_SIZE)>>libc.Int32FromInt32(1)) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		goto _26
	_26:
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v23 = st
	goto _28
_28:
	lp_pitch_buf = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(libc.Uint64FromInt32(libc.Int32FromInt32(DEC_PITCH_BUF_SIZE)>>libc.Int32FromInt32(1))*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))
	Opus_pitch_downsample(tls, decode_mem, lp_pitch_buf, libc.Int32FromInt32(DEC_PITCH_BUF_SIZE)>>libc.Int32FromInt32(1), C, int32(2), arch)
	Opus_pitch_search(tls, lp_pitch_buf+uintptr(libc.Int32FromInt32(PLC_PITCH_LAG_MAX)>>libc.Int32FromInt32(1))*4, lp_pitch_buf, libc.Int32FromInt32(DEC_PITCH_BUF_SIZE)-libc.Int32FromInt32(PLC_PITCH_LAG_MAX), libc.Int32FromInt32(PLC_PITCH_LAG_MAX)-libc.Int32FromInt32(PLC_PITCH_LAG_MIN), bp, arch)
	*(*int32)(unsafe.Pointer(bp)) = int32(PLC_PITCH_LAG_MAX) - *(*int32)(unsafe.Pointer(bp))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _30
	_30:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _32
_32:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
	return *(*int32)(unsafe.Pointer(bp))
}

func celt_synthesis(tls *libc.TLS, mode uintptr, X uintptr, out_syn uintptr, oldBandE uintptr, start int32, effEnd int32, C int32, CC int32, isTransient int32, LM int32, downsample int32, silence int32, arch int32) {
	var B, M, N, NB, b, c, i, nbEBands, overlap, shift, v33 int32
	var _saved_stack, freq, freq2, freq21, st, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = B, M, N, NB, _saved_stack, b, c, freq, freq2, freq21, i, nbEBands, overlap, shift, st, v1, v11, v13, v15, v17, v19, v21, v23, v3, v33, v5, v7, v9
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _2
	_2:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _4
_4:
	_saved_stack = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack
	overlap = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).Foverlap
	nbEBands = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FnbEBands
	N = (*OpusT_OpusCustomMode)(unsafe.Pointer(mode)).FshortMdctSize << LM
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _6
	_6:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _8
_8:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _10
	_10:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v7 = st
	goto _12
_12:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((libc.Uint64FromInt32(4) - libc.Uint64FromInt64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (libc.Uint64FromInt32(4) - libc.Uint64FromInt32(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v9 = libc.Xmalloc(tls, uint64(16))
		goto _14
	_14:
		st = v9
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v11 = st
	goto _16
_16:
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v13 = libc.Xmalloc(tls, uint64(16))
		goto _18
	_18:
		st = v13
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v15 = st
	goto _20
_20:
	if !(int64(libc.Int32FromUint64(libc.Uint64FromInt32(N)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4420, int32(432))
	}
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v17 = libc.Xmalloc(tls, uint64(16))
		goto _22
	_22:
		st = v17
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v19 = st
	goto _24
_24:
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(libc.Uint64FromInt32(N) * (libc.Uint64FromInt64(4) / libc.Uint64FromInt64(1)))
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v21 = libc.Xmalloc(tls, uint64(16))
		goto _26
	_26:
		st = v21
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v23 = st
	goto _28
_28:
	freq = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(libc.Uint64FromInt32(N)*(libc.Uint64FromInt64(4)/libc.Uint64FromInt64(1))) /**< Interleaved signal MDCTs */
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
		libc.Xmemcpy(tls, freq2, freq, libc.Uint64FromInt32(N)*uint64(4)+libc.Uint64FromInt64(0*((int64(freq2)-int64(freq))/4)))
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
				*(*OpusT_celt_sig)(unsafe.Pointer(freq + uintptr(i)*4)) = float32(libc.Float32FromFloat32(0.5)**(*OpusT_celt_sig)(unsafe.Pointer(freq + uintptr(i)*4))) + float32(libc.Float32FromFloat32(0.5)**(*OpusT_celt_sig)(unsafe.Pointer(freq21 + uintptr(i)*4)))
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
	st = libc.Xpthread_getspecific(tls, libc.Uint32FromUint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _40
	_40:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, libc.Uint32FromUint32(0x6f707573), st)
	}
	v3 = st
	goto _42
_42:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
}
