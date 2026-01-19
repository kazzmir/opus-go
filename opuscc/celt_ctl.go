// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

//go:build linux && (amd64 || arm64)

package opuscc

import (
	"reflect"
	"unsafe"

	libc "opusgo/libcshim"
)

var _ reflect.Type
var _ unsafe.Pointer

func Opus_opus_custom_decoder_ctl(tls *libc.TLS, st uintptr, request int32, va uintptr) (r int32) {
	var ap OpusT_va_list
	var decode_buffer_size, i int32
	var oldBandE, oldLogE, oldLogE2, value1, value10, value12, value5, value6, value7, value8 uintptr
	var value, value11, value2, value3, value4, value9 OpusT_opus_int32
	var v2 OpusT_celt_glog
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = ap, decode_buffer_size, i, oldBandE, oldLogE, oldLogE2, value, value1, value10, value11, value12, value2, value3, value4, value5, value6, value7, value8, value9, v2
	ap = va
	switch request {
	case int32(OPUS_SET_COMPLEXITY_REQUEST):
		value = libc.VaInt32(&ap)
		if value < 0 || value > int32(10) {
			goto bad_arg
		}
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fcomplexity = value
	case int32(OPUS_GET_COMPLEXITY_REQUEST):
		value1 = libc.VaUintptr(&ap)
		if !(value1 != 0) {
			goto bad_arg
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(value1)) = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fcomplexity
	case int32(CELT_SET_START_BAND_REQUEST):
		value2 = libc.VaInt32(&ap)
		if value2 < 0 || value2 >= (*OpusT_OpusCustomMode)(unsafe.Pointer((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fmode)).FnbEBands {
			goto bad_arg
		}
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fstart = value2
	case int32(CELT_SET_END_BAND_REQUEST):
		value3 = libc.VaInt32(&ap)
		if value3 < int32(1) || value3 > (*OpusT_OpusCustomMode)(unsafe.Pointer((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fmode)).FnbEBands {
			goto bad_arg
		}
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fend = value3
	case int32(CELT_SET_CHANNELS_REQUEST):
		value4 = libc.VaInt32(&ap)
		if value4 < int32(1) || value4 > int32(2) {
			goto bad_arg
		}
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fstream_channels = value4
	case int32(CELT_GET_AND_CLEAR_ERROR_REQUEST):
		value5 = libc.VaUintptr(&ap)
		if value5 == uintptr(uint32(0)) {
			goto bad_arg
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(value5)) = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Ferror1
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Ferror1 = 0
	case int32(OPUS_GET_LOOKAHEAD_REQUEST):
		value6 = libc.VaUintptr(&ap)
		if value6 == uintptr(uint32(0)) {
			goto bad_arg
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(value6)) = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Foverlap / (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fdownsample
	case int32(OPUS_RESET_STATE):
		decode_buffer_size = int32(DEC_PITCH_BUF_SIZE)
		oldBandE = st + 112 + uintptr((decode_buffer_size+(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Foverlap)*(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fchannels)*4
		oldLogE = oldBandE + uintptr(int32(2)*(*OpusT_OpusCustomMode)(unsafe.Pointer((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fmode)).FnbEBands)*4
		oldLogE2 = oldLogE + uintptr(int32(2)*(*OpusT_OpusCustomMode)(unsafe.Pointer((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fmode)).FnbEBands)*4
		libc.Xmemset(tls, st+48, 0, uint64(int64(opus_custom_decoder_get_size(tls, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fmode, (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fchannels))-(int64(st+48)-int64(st)))*uint64(1))
		i = 0
		for {
			if !(i < int32(2)*(*OpusT_OpusCustomMode)(unsafe.Pointer((*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fmode)).FnbEBands) {
				break
			}
			v2 = -float32(28)
			*(*OpusT_celt_glog)(unsafe.Pointer(oldLogE2 + uintptr(i)*4)) = v2
			*(*OpusT_celt_glog)(unsafe.Pointer(oldLogE + uintptr(i)*4)) = v2
			goto _1
		_1:
			;
			i = i + 1
		}
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fskip_plc = int32(1)
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Flast_frame_type = FRAME_NONE
	case int32(OPUS_GET_PITCH_REQUEST):
		value7 = libc.VaUintptr(&ap)
		if value7 == uintptr(uint32(0)) {
			goto bad_arg
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(value7)) = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fpostfilter_period
	case int32(CELT_GET_MODE_REQUEST):
		value8 = libc.VaUintptr(&ap)
		if value8 == uintptr(0) {
			goto bad_arg
		}
		*(*uintptr)(unsafe.Pointer(value8)) = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fmode
	case int32(CELT_SET_SIGNALLING_REQUEST):
		value9 = libc.VaInt32(&ap)
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fsignalling = value9
	case int32(OPUS_GET_FINAL_RANGE_REQUEST):
		value10 = libc.VaUintptr(&ap)
		if value10 == uintptr(0) {
			goto bad_arg
		}
		*(*OpusT_opus_uint32)(unsafe.Pointer(value10)) = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Frng
	case int32(OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST):
		value11 = libc.VaInt32(&ap)
		if value11 < 0 || value11 > int32(1) {
			goto bad_arg
		}
		(*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fdisable_inv = value11
	case int32(OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST):
		value12 = libc.VaUintptr(&ap)
		if !(value12 != 0) {
			goto bad_arg
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(value12)) = (*OpusT_OpusCustomDecoder)(unsafe.Pointer(st)).Fdisable_inv
	default:
		goto bad_request
	}
	_ = ap
	return OPUS_OK
	goto bad_arg
bad_arg:
	;
	_ = ap
	return -int32(1)
	goto bad_request
bad_request:
	;
	_ = ap
	return -int32(5)
}

const CELT_SIG_SCALE8 = "32768.f"
const Q31ONE2 = "1.0f"
const VERY_SMALL2 = "1e-30f"

var trim_icdf10 = [11]uint8{
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
var spread_icdf10 = [4]uint8{
	0: uint8(25),
	1: uint8(23),
	2: uint8(2),
}
var tapset_icdf10 = [3]uint8{
	0: uint8(2),
	1: uint8(1),
}

var log2_x_norm_coeff9 = [8]float32{
	0: float32(1),
	1: float32(0.8888888955116272),
	2: float32(0.8),
	3: float32(0.7272727489471436),
	4: float32(0.6666666865348816),
	5: float32(0.6153846383094788),
	6: float32(0.5714285969734192),
	7: float32(0.5333333611488342),
}
var log2_y_norm_coeff9 = [8]float32{
	1: float32(0.1699250042438507),
	2: float32(0.32192808389663696),
	3: float32(0.45943161845207214),
	4: float32(0.5849624872207642),
	5: float32(0.7004396915435791),
	6: float32(0.8073549270629883),
	7: float32(0.9068905711174011),
}

func find_best_pitch(tls *libc.TLS, xcorr uintptr, y uintptr, len1 int32, max_pitch int32, best_pitch uintptr) {
	var Syy, xcorr16, v3 OpusT_opus_val32
	var best_den [2]OpusT_opus_val32
	var best_num [2]OpusT_opus_val16
	var i, j int32
	var num OpusT_opus_val16
	_, _, _, _, _, _, _, _ = Syy, best_den, best_num, i, j, num, xcorr16, v3
	Syy = float32(1)
	best_num[0] = float32(-int32(1))
	best_num[int32(1)] = float32(-int32(1))
	best_den[0] = float32(0)
	best_den[int32(1)] = float32(0)
	*(*int32)(unsafe.Pointer(best_pitch)) = 0
	*(*int32)(unsafe.Pointer(best_pitch + 1*4)) = int32(1)
	j = 0
	for {
		if !(j < len1) {
			break
		}
		Syy = Syy + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(y + uintptr(j)*4))**(*OpusT_opus_val16)(unsafe.Pointer(y + uintptr(j)*4)))
		goto _1
	_1:
		;
		j = j + 1
	}
	i = 0
	for {
		if !(i < max_pitch) {
			break
		}
		if *(*OpusT_opus_val32)(unsafe.Pointer(xcorr + uintptr(i)*4)) > float32(0) {
			xcorr16 = *(*OpusT_opus_val32)(unsafe.Pointer(xcorr + uintptr(i)*4))
			/* Considering the range of xcorr16, this should avoid both underflows
			   and overflows (inf) when squaring xcorr16 */
			xcorr16 = xcorr16 * float32(1e-12)
			num = OpusT_opus_val32(xcorr16 * xcorr16)
			if OpusT_opus_val16(num*best_den[int32(1)]) > OpusT_opus_val16(best_num[int32(1)]*Syy) {
				if OpusT_opus_val16(num*best_den[0]) > OpusT_opus_val16(best_num[0]*Syy) {
					best_num[int32(1)] = best_num[0]
					best_den[int32(1)] = best_den[0]
					*(*int32)(unsafe.Pointer(best_pitch + 1*4)) = *(*int32)(unsafe.Pointer(best_pitch))
					best_num[0] = num
					best_den[0] = Syy
					*(*int32)(unsafe.Pointer(best_pitch)) = i
				} else {
					best_num[int32(1)] = num
					best_den[int32(1)] = Syy
					*(*int32)(unsafe.Pointer(best_pitch + 1*4)) = i
				}
			}
		}
		Syy = Syy + (OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(y + uintptr(i+len1)*4))**(*OpusT_opus_val16)(unsafe.Pointer(y + uintptr(i+len1)*4))) - OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(y + uintptr(i)*4))**(*OpusT_opus_val16)(unsafe.Pointer(y + uintptr(i)*4))))
		if float32(int32(1)) > Syy {
			v3 = float32(int32(1))
		} else {
			v3 = Syy
		}
		Syy = v3
		goto _2
	_2:
		;
		i = i + 1
	}
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
	mem0 = float32(0)
	mem1 = float32(0)
	mem2 = float32(0)
	mem3 = float32(0)
	mem4 = float32(0)
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

func Opus_pitch_downsample(tls *libc.TLS, x uintptr, x_lp uintptr, len1 int32, C int32, factor int32, arch int32) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var c1, tmp OpusT_opus_val16
	var i, offset int32
	var _ /* ac at bp+0 */ [5]OpusT_opus_val32
	var _ /* lpc at bp+20 */ [4]OpusT_opus_val16
	var _ /* lpc2 at bp+36 */ [5]OpusT_opus_val16
	_, _, _, _ = c1, i, offset, tmp
	tmp = float32(1)
	c1 = float32(0.8)
	offset = factor / int32(2)
	i = int32(1)
	for {
		if !(i < len1) {
			break
		}
		*(*OpusT_opus_val16)(unsafe.Pointer(x_lp + uintptr(i)*4)) = float32(float32(0.25)**(*OpusT_celt_sig)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(x)) + uintptr(factor*i-offset)*4))) + float32(float32(0.25)**(*OpusT_celt_sig)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(x)) + uintptr(factor*i+offset)*4))) + float32(float32(0.5)**(*OpusT_celt_sig)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(x)) + uintptr(factor*i)*4)))
		goto _1
	_1:
		;
		i = i + 1
	}
	*(*OpusT_opus_val16)(unsafe.Pointer(x_lp)) = float32(float32(0.25)**(*OpusT_celt_sig)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(x)) + uintptr(offset)*4))) + float32(float32(0.5)**(*OpusT_celt_sig)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(x)))))
	if C == int32(2) {
		i = int32(1)
		for {
			if !(i < len1) {
				break
			}
			*(*OpusT_opus_val16)(unsafe.Pointer(x_lp + uintptr(i)*4)) += float32(float32(0.25)**(*OpusT_celt_sig)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(x + 1*8)) + uintptr(factor*i-offset)*4))) + float32(float32(0.25)**(*OpusT_celt_sig)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(x + 1*8)) + uintptr(factor*i+offset)*4))) + float32(float32(0.5)**(*OpusT_celt_sig)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(x + 1*8)) + uintptr(factor*i)*4)))
			goto _2
		_2:
			;
			i = i + 1
		}
		*(*OpusT_opus_val16)(unsafe.Pointer(x_lp)) += float32(float32(0.25)**(*OpusT_celt_sig)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(x + 1*8)) + uintptr(offset)*4))) + float32(float32(0.5)**(*OpusT_celt_sig)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(x + 1*8)))))
	}
	Opus__celt_autocorr(tls, x_lp, bp, uintptr(uint32(0)), 0, int32(4), len1, arch)
	/* Noise floor -40 dB */
	*(*OpusT_opus_val32)(unsafe.Pointer(bp)) *= float32(1.0001)
	/* Lag windowing */
	i = int32(1)
	for {
		if !(i <= int32(4)) {
			break
		}
		/*ac[i] *= exp(-.5*(2*M_PI*.002*i)*(2*M_PI*.002*i));*/
		*(*OpusT_opus_val32)(unsafe.Pointer(bp + uintptr(i)*4)) -= OpusT_opus_val32(OpusT_opus_val32((*(*[5]OpusT_opus_val32)(unsafe.Pointer(bp)))[i]*float32(float32(0.008)*float32(i))) * float32(float32(0.008)*float32(i)))
		goto _3
	_3:
		;
		i = i + 1
	}
	Opus__celt_lpc(tls, bp+20, bp, int32(4))
	i = 0
	for {
		if !(i < int32(4)) {
			break
		}
		tmp = float32(float32(0.9) * tmp)
		(*(*[4]OpusT_opus_val16)(unsafe.Pointer(bp + 20)))[i] = OpusT_opus_val16((*(*[4]OpusT_opus_val16)(unsafe.Pointer(bp + 20)))[i] * tmp)
		goto _4
	_4:
		;
		i = i + 1
	}
	/* Add a zero */
	(*(*[5]OpusT_opus_val16)(unsafe.Pointer(bp + 36)))[0] = (*(*[4]OpusT_opus_val16)(unsafe.Pointer(bp + 20)))[0] + float32(0.8)
	(*(*[5]OpusT_opus_val16)(unsafe.Pointer(bp + 36)))[int32(1)] = (*(*[4]OpusT_opus_val16)(unsafe.Pointer(bp + 20)))[int32(1)] + OpusT_opus_val16(c1*(*(*[4]OpusT_opus_val16)(unsafe.Pointer(bp + 20)))[0])
	(*(*[5]OpusT_opus_val16)(unsafe.Pointer(bp + 36)))[int32(2)] = (*(*[4]OpusT_opus_val16)(unsafe.Pointer(bp + 20)))[int32(2)] + OpusT_opus_val16(c1*(*(*[4]OpusT_opus_val16)(unsafe.Pointer(bp + 20)))[int32(1)])
	(*(*[5]OpusT_opus_val16)(unsafe.Pointer(bp + 36)))[int32(3)] = (*(*[4]OpusT_opus_val16)(unsafe.Pointer(bp + 20)))[int32(3)] + OpusT_opus_val16(c1*(*(*[4]OpusT_opus_val16)(unsafe.Pointer(bp + 20)))[int32(2)])
	(*(*[5]OpusT_opus_val16)(unsafe.Pointer(bp + 36)))[int32(4)] = OpusT_opus_val16(c1 * (*(*[4]OpusT_opus_val16)(unsafe.Pointer(bp + 20)))[int32(3)])
	celt_fir5(tls, x_lp, bp+36, len1)
}

// C documentation
//
//	/* Pure C implementation. */
func Opus_celt_pitch_xcorr_c(tls *libc.TLS, _x uintptr, _y uintptr, xcorr uintptr, len1 int32, max_pitch int32, arch int32) {
	_ = arch
	/*The EDSP version requires that max_pitch is at least 1, and that _x is
	  32-bit aligned.
	 Since it's hard to put asserts in assembly, put them here.*/
	if !(max_pitch > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+4629, __ccgo_ts+4659, int32(265))
	}
	_ = uint64(_x)&uint64(uint32(3)) == uint64(uint32(0))

	xS := unsafe.Slice((*OpusT_opus_val16)(unsafe.Pointer(_x)), int(len1))
	// y must have at least max_pitch+len1 samples, as we access y[i1+i].
	yS := unsafe.Slice((*OpusT_opus_val16)(unsafe.Pointer(_y)), int(max_pitch+len1))
	xcorrS := unsafe.Slice((*OpusT_opus_val32)(unsafe.Pointer(xcorr)), int(max_pitch))

	// The original unrolled path requires len1 >= 3.
	if max_pitch > 3 && len1 < 3 {
		Opus_celt_fatal(tls, __ccgo_ts+3349, __ccgo_ts+3374, int32(69))
	}

	for i1 := int32(0); i1 < max_pitch-3; i1 += 4 {
		sum0 := OpusT_opus_val32(0)
		sum1 := OpusT_opus_val32(0)
		sum2 := OpusT_opus_val32(0)
		sum3 := OpusT_opus_val32(0)

		y0 := yS[i1]
		y1 := yS[i1+1]
		y2 := yS[i1+2]
		var y3 OpusT_opus_val16

		j := int32(0)
		for ; j < len1-3; j += 4 {
			tmp := xS[j]
			y3 = yS[i1+j+3]
			sum0 += OpusT_opus_val32(tmp * y0)
			sum1 += OpusT_opus_val32(tmp * y1)
			sum2 += OpusT_opus_val32(tmp * y2)
			sum3 += OpusT_opus_val32(tmp * y3)

			tmp = xS[j+1]
			y0 = yS[i1+j+4]
			sum0 += OpusT_opus_val32(tmp * y1)
			sum1 += OpusT_opus_val32(tmp * y2)
			sum2 += OpusT_opus_val32(tmp * y3)
			sum3 += OpusT_opus_val32(tmp * y0)

			tmp = xS[j+2]
			y1 = yS[i1+j+5]
			sum0 += OpusT_opus_val32(tmp * y2)
			sum1 += OpusT_opus_val32(tmp * y3)
			sum2 += OpusT_opus_val32(tmp * y0)
			sum3 += OpusT_opus_val32(tmp * y1)

			tmp = xS[j+3]
			y2 = yS[i1+j+6]
			sum0 += OpusT_opus_val32(tmp * y3)
			sum1 += OpusT_opus_val32(tmp * y0)
			sum2 += OpusT_opus_val32(tmp * y1)
			sum3 += OpusT_opus_val32(tmp * y2)
		}

		// Remainder (up to 3 samples).
		idx := j
		if idx < len1 {
			tmp := xS[idx]
			y3 = yS[i1+idx+3]
			sum0 += OpusT_opus_val32(tmp * y0)
			sum1 += OpusT_opus_val32(tmp * y1)
			sum2 += OpusT_opus_val32(tmp * y2)
			sum3 += OpusT_opus_val32(tmp * y3)
			idx++
		}
		if idx < len1 {
			tmp := xS[idx]
			y0 = yS[i1+idx+3]
			sum0 += OpusT_opus_val32(tmp * y1)
			sum1 += OpusT_opus_val32(tmp * y2)
			sum2 += OpusT_opus_val32(tmp * y3)
			sum3 += OpusT_opus_val32(tmp * y0)
			idx++
		}
		if idx < len1 {
			tmp := xS[idx]
			y1 = yS[i1+idx+3]
			sum0 += OpusT_opus_val32(tmp * y2)
			sum1 += OpusT_opus_val32(tmp * y3)
			sum2 += OpusT_opus_val32(tmp * y0)
			sum3 += OpusT_opus_val32(tmp * y1)
		}

		xcorrS[i1] = sum0
		xcorrS[i1+1] = sum1
		xcorrS[i1+2] = sum2
		xcorrS[i1+3] = sum3
	}
	/* In case max_pitch isn't a multiple of 4, do non-unrolled version. */
	for i1 := (max_pitch &^ 3); i1 < max_pitch; i1++ {
		xy := OpusT_opus_val32(0)
		for i := int32(0); i < len1; i++ {
			xy += OpusT_opus_val32(xS[i] * yS[i1+i])
		}
		xcorrS[i1] = xy
	}
}

func Opus_pitch_search(tls *libc.TLS, x_lp uintptr, y1 uintptr, len1 int32, max_pitch int32, pitch uintptr, arch int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var _saved_stack, st, x_lp4, xcorr, y_lp4, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v9 uintptr
	var a, b, c, sum, xy, v81 OpusT_opus_val32
	var i, i1, j, lag, offset int32
	var _ /* best_pitch at bp+0 */ [2]int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = _saved_stack, a, b, c, i, i1, j, lag, offset, st, sum, x_lp4, xcorr, xy, y_lp4, v1, v11, v13, v15, v17, v19, v21, v23, v3, v5, v7, v81, v9
	*(*[2]int32)(unsafe.Pointer(bp)) = [2]int32{}
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
	if !(len1 > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+4675, __ccgo_ts+4659, int32(325))
	}
	if !(max_pitch > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+4629, __ccgo_ts+4659, int32(326))
	}
	lag = len1 + max_pitch
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
	if !(int64(int32(uint64(uint32(len1>>int32(2)))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4659, int32(329))
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
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32(len1>>int32(2))) * (uint64(4) / uint64(1)))
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
	x_lp4 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32(len1>>int32(2)))*(uint64(4)/uint64(1)))
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
	if !(int64(int32(uint64(uint32(lag>>int32(2)))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4659, int32(330))
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
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32(lag>>int32(2))) * (uint64(4) / uint64(1)))
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
	y_lp4 = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32(lag>>int32(2)))*(uint64(4)/uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _54
	_54:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _56
_56:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _58
	_58:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v7 = st
	goto _60
_60:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v7)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
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
	if !(int64(int32(uint64(uint32(max_pitch>>int32(1)))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v11)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v15)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4659, int32(331))
	}
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
	*(*uintptr)(unsafe.Pointer(v19 + 8)) += uintptr(uint64(uint32(max_pitch>>int32(1))) * (uint64(4) / uint64(1)))
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
	xcorr = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v23)).Fglobal_stack - uintptr(uint64(uint32(max_pitch>>int32(1)))*(uint64(4)/uint64(1)))
	/* Downsample by 2 again */
	j = 0
	for {
		if !(j < len1>>int32(2)) {
			break
		}
		*(*OpusT_opus_val16)(unsafe.Pointer(x_lp4 + uintptr(j)*4)) = *(*OpusT_opus_val16)(unsafe.Pointer(x_lp + uintptr(int32(2)*j)*4))
		goto _77
	_77:
		;
		j = j + 1
	}
	j = 0
	for {
		if !(j < lag>>int32(2)) {
			break
		}
		*(*OpusT_opus_val16)(unsafe.Pointer(y_lp4 + uintptr(j)*4)) = *(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(int32(2)*j)*4))
		goto _78
	_78:
		;
		j = j + 1
	}
	/* Coarse search with 4x decimation */
	Opus_celt_pitch_xcorr_c(tls, x_lp4, y_lp4, xcorr, len1>>int32(2), max_pitch>>int32(2), arch)
	find_best_pitch(tls, xcorr, y_lp4, len1>>int32(2), max_pitch>>int32(2), bp)
	/* Finer search with 2x decimation */
	i1 = 0
	for {
		if !(i1 < max_pitch>>int32(1)) {
			break
		}
		*(*OpusT_opus_val32)(unsafe.Pointer(xcorr + uintptr(i1)*4)) = float32(0)
		if libc.Xabs(tls, i1-int32(2)*(*(*[2]int32)(unsafe.Pointer(bp)))[0]) > int32(2) && libc.Xabs(tls, i1-int32(2)*(*(*[2]int32)(unsafe.Pointer(bp)))[int32(1)]) > int32(2) {
			goto _79
		}
		_ = arch
		xy = float32(0)
		i = int32(0)
		for {
			if !(i < len1>>int32(1)) {
				break
			}
			xy = xy + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(x_lp + uintptr(i)*4))**(*OpusT_opus_val16)(unsafe.Pointer(y1 + uintptr(i1)*4 + uintptr(i)*4)))
			goto _80
		_80:
			;
			i = i + 1
		}
		v81 = xy
		goto _82
	_82:
		sum = v81
		if float32(-int32(1)) > sum {
			v81 = float32(-int32(1))
		} else {
			v81 = sum
		}
		*(*OpusT_opus_val32)(unsafe.Pointer(xcorr + uintptr(i1)*4)) = v81
		goto _79
	_79:
		;
		i1 = i1 + 1
	}
	find_best_pitch(tls, xcorr, y1, len1>>int32(1), max_pitch>>int32(1), bp)
	/* Refine by pseudo-interpolation */
	if (*(*[2]int32)(unsafe.Pointer(bp)))[0] > 0 && (*(*[2]int32)(unsafe.Pointer(bp)))[0] < max_pitch>>int32(1)-int32(1) {
		a = *(*OpusT_opus_val32)(unsafe.Pointer(xcorr + uintptr((*(*[2]int32)(unsafe.Pointer(bp)))[0]-int32(1))*4))
		b = *(*OpusT_opus_val32)(unsafe.Pointer(xcorr + uintptr((*(*[2]int32)(unsafe.Pointer(bp)))[0])*4))
		c = *(*OpusT_opus_val32)(unsafe.Pointer(xcorr + uintptr((*(*[2]int32)(unsafe.Pointer(bp)))[0]+int32(1))*4))
		if c-a > float32(float32(0.7)*(b-a)) {
			offset = int32(1)
		} else {
			if a-c > float32(float32(0.7)*(b-c)) {
				offset = -int32(1)
			} else {
				offset = 0
			}
		}
	} else {
		offset = 0
	}
	*(*int32)(unsafe.Pointer(pitch)) = int32(2)*(*(*[2]int32)(unsafe.Pointer(bp)))[0] - offset
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _85
	_85:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _87
_87:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
}

func compute_pitch_gain(tls *libc.TLS, xy OpusT_opus_val32, xx OpusT_opus_val32, yy OpusT_opus_val32) (r OpusT_opus_val16) {
	return xy / float32(libc.Xsqrt(tls, float64(float32(1)+OpusT_opus_val32(xx*yy))))
}

var second_check = [16]int32{
	2:  int32(3),
	3:  int32(2),
	4:  int32(3),
	5:  int32(2),
	6:  int32(5),
	7:  int32(2),
	8:  int32(3),
	9:  int32(2),
	10: int32(3),
	11: int32(2),
	12: int32(5),
	13: int32(2),
	14: int32(3),
	15: int32(2),
}

func Opus_remove_doubling(tls *libc.TLS, x2 uintptr, maxperiod int32, minperiod int32, N2 int32, T0_ uintptr, prev_period int32, prev_gain OpusT_opus_val16, arch int32) (r OpusT_opus_val16) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var T, T0, T1, T1b, i, i1, i2, k, minperiod0, offset, v5 int32
	var _saved_stack, st, yy_lookup, v1, v10, v12, v14, v16, v18, v20, v22, v24, v3, v6, v8 uintptr
	var best_xy, best_yy, xy, xy01, xy02, yy, v33 OpusT_opus_val32
	var cont, g, g0, g1, pg, thresh, v34 OpusT_opus_val16
	var xcorr [3]OpusT_opus_val32
	var v36, v37 OpusT_opus_uint32
	var v44 float32
	var _ /* xx at bp+4 */ OpusT_opus_val32
	var _ /* xy at bp+0 */ OpusT_opus_val32
	var _ /* xy2 at bp+8 */ OpusT_opus_val32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = T, T0, T1, T1b, _saved_stack, best_xy, best_yy, cont, g, g0, g1, i, i1, i2, k, minperiod0, offset, pg, st, thresh, xcorr, xy, xy01, xy02, yy, yy_lookup, v1, v10, v12, v14, v16, v18, v20, v22, v24, v3, v33, v34, v36, v37, v44, v5, v6, v8
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
	minperiod0 = minperiod
	maxperiod = maxperiod / int32(2)
	minperiod = minperiod / int32(2)
	*(*int32)(unsafe.Pointer(T0_)) /= int32(2)
	prev_period = prev_period / int32(2)
	N2 = N2 / int32(2)
	x2 = x2 + uintptr(maxperiod)*4
	if *(*int32)(unsafe.Pointer(T0_)) >= maxperiod {
		*(*int32)(unsafe.Pointer(T0_)) = maxperiod - int32(1)
	}
	v5 = *(*int32)(unsafe.Pointer(T0_))
	T0 = v5
	T = v5
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _7
	_7:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _9
_9:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v6 = libc.Xmalloc(tls, uint64(16))
		goto _11
	_11:
		st = v6
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v8 = st
	goto _13
_13:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v8)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
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
	if !(int64(int32(uint64(uint32(maxperiod+int32(1)))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v12)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v16)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4659, int32(479))
	}
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
	*(*uintptr)(unsafe.Pointer(v20 + 8)) += uintptr(uint64(uint32(maxperiod+int32(1))) * (uint64(4) / uint64(1)))
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
	yy_lookup = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v24)).Fglobal_stack - uintptr(uint64(uint32(maxperiod+int32(1)))*(uint64(4)/uint64(1)))
	_ = arch
	v1 = x2
	xy01 = float32(0)
	xy02 = float32(0)
	i = int32(0)
	for {
		if !(i < N2) {
			break
		}
		xy01 = xy01 + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(v1 + uintptr(i)*4))**(*OpusT_opus_val16)(unsafe.Pointer(x2 + uintptr(i)*4)))
		xy02 = xy02 + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(v1 + uintptr(i)*4))**(*OpusT_opus_val16)(unsafe.Pointer(x2 - uintptr(T0)*4 + uintptr(i)*4)))
		goto _31
	_31:
		;
		i = i + 1
	}
	*(*OpusT_opus_val32)(unsafe.Pointer(bp + 4)) = xy01
	*(*OpusT_opus_val32)(unsafe.Pointer(bp)) = xy02
	*(*OpusT_opus_val32)(unsafe.Pointer(yy_lookup)) = *(*OpusT_opus_val32)(unsafe.Pointer(bp + 4))
	yy = *(*OpusT_opus_val32)(unsafe.Pointer(bp + 4))
	i2 = int32(1)
	for {
		if !(i2 <= maxperiod) {
			break
		}
		yy = yy + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(x2 + uintptr(-i2)*4))**(*OpusT_opus_val16)(unsafe.Pointer(x2 + uintptr(-i2)*4))) - OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(x2 + uintptr(N2-i2)*4))**(*OpusT_opus_val16)(unsafe.Pointer(x2 + uintptr(N2-i2)*4)))
		if float32(int32(0)) > yy {
			v33 = float32(int32(0))
		} else {
			v33 = yy
		}
		*(*OpusT_opus_val32)(unsafe.Pointer(yy_lookup + uintptr(i2)*4)) = v33
		goto _32
	_32:
		;
		i2 = i2 + 1
	}
	yy = *(*OpusT_opus_val32)(unsafe.Pointer(yy_lookup + uintptr(T0)*4))
	best_xy = *(*OpusT_opus_val32)(unsafe.Pointer(bp))
	best_yy = yy
	v34 = compute_pitch_gain(tls, *(*OpusT_opus_val32)(unsafe.Pointer(bp)), *(*OpusT_opus_val32)(unsafe.Pointer(bp + 4)), yy)
	g0 = v34
	g = v34
	/* Look for any pitch at T/k */
	k = int32(2)
	for {
		if !(k <= int32(15)) {
			break
		}
		cont = float32(0)
		v36 = uint32(int32(2) * k)
		_ = v36 > uint32(0)
		v37 = uint32(int32(2)*T0+k) / v36
		goto _38
	_38:
		T1 = int32(v37)
		if T1 < minperiod {
			break
		}
		/* Look for another strong correlation at T1b */
		if k == int32(2) {
			if T1+T0 > maxperiod {
				T1b = T0
			} else {
				T1b = T0 + T1
			}
		} else {
			v36 = uint32(int32(2) * k)
			_ = v36 > uint32(0)
			v37 = uint32(int32(2)*second_check[k]*T0+k) / v36
			goto _41
		_41:
			T1b = int32(v37)
		}
		_ = arch
		v1 = x2
		xy01 = float32(0)
		xy02 = float32(0)
		i = int32(0)
		for {
			if !(i < N2) {
				break
			}
			xy01 = xy01 + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(v1 + uintptr(i)*4))**(*OpusT_opus_val16)(unsafe.Pointer(x2 + uintptr(-T1)*4 + uintptr(i)*4)))
			xy02 = xy02 + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(v1 + uintptr(i)*4))**(*OpusT_opus_val16)(unsafe.Pointer(x2 + uintptr(-T1b)*4 + uintptr(i)*4)))
			goto _43
		_43:
			;
			i = i + 1
		}
		*(*OpusT_opus_val32)(unsafe.Pointer(bp)) = xy01
		*(*OpusT_opus_val32)(unsafe.Pointer(bp + 8)) = xy02
		*(*OpusT_opus_val32)(unsafe.Pointer(bp)) = float32(float32(0.5) * (*(*OpusT_opus_val32)(unsafe.Pointer(bp)) + *(*OpusT_opus_val32)(unsafe.Pointer(bp + 8))))
		yy = float32(float32(0.5) * (*(*OpusT_opus_val32)(unsafe.Pointer(yy_lookup + uintptr(T1)*4)) + *(*OpusT_opus_val32)(unsafe.Pointer(yy_lookup + uintptr(T1b)*4))))
		g1 = compute_pitch_gain(tls, *(*OpusT_opus_val32)(unsafe.Pointer(bp)), *(*OpusT_opus_val32)(unsafe.Pointer(bp + 4)), yy)
		if libc.Xabs(tls, T1-prev_period) <= int32(1) {
			cont = prev_gain
		} else {
			if libc.Xabs(tls, T1-prev_period) <= int32(2) && int32(5)*k*k < T0 {
				cont = float32(float32(0.5) * prev_gain)
			} else {
				cont = float32(0)
			}
		}
		if float32(0.3) > float32(float32(0.7)*g0)-cont {
			v44 = float32(0.3)
		} else {
			v44 = float32(float32(0.7)*g0) - cont
		}
		thresh = v44
		/* Bias against very high pitch (very short period) to avoid false-positives
		   due to short-term correlation */
		if T1 < int32(3)*minperiod {
			if float32(0.4) > float32(float32(0.85)*g0)-cont {
				v44 = float32(0.4)
			} else {
				v44 = float32(float32(0.85)*g0) - cont
			}
			thresh = v44
		} else {
			if T1 < int32(2)*minperiod {
				if float32(0.5) > float32(float32(0.9)*g0)-cont {
					v44 = float32(0.5)
				} else {
					v44 = float32(float32(0.9)*g0) - cont
				}
				thresh = v44
			}
		}
		if g1 > thresh {
			best_xy = *(*OpusT_opus_val32)(unsafe.Pointer(bp))
			best_yy = yy
			T = T1
			g = g1
		}
		goto _35
	_35:
		;
		k = k + 1
	}
	if float32(int32(0)) > best_xy {
		v33 = float32(int32(0))
	} else {
		v33 = best_xy
	}
	best_xy = v33
	if best_yy <= best_xy {
		pg = float32(1)
	} else {
		pg = best_xy / (best_yy + float32(1))
	}
	k = 0
	for {
		if !(k < int32(3)) {
			break
		}
		_ = arch
		xy = float32(0)
		i1 = int32(0)
		for {
			if !(i1 < N2) {
				break
			}
			xy = xy + OpusT_opus_val32(*(*OpusT_opus_val16)(unsafe.Pointer(x2 + uintptr(i1)*4))**(*OpusT_opus_val16)(unsafe.Pointer(x2 - uintptr(T+k-int32(1))*4 + uintptr(i1)*4)))
			goto _49
		_49:
			;
			i1 = i1 + 1
		}
		v33 = xy
		goto _51
	_51:
		xcorr[k] = v33
		goto _48
	_48:
		;
		k = k + 1
	}
	if xcorr[int32(2)]-xcorr[0] > float32(float32(0.7)*(xcorr[int32(1)]-xcorr[0])) {
		offset = int32(1)
	} else {
		if xcorr[0]-xcorr[int32(2)] > float32(float32(0.7)*(xcorr[int32(1)]-xcorr[int32(2)])) {
			offset = -int32(1)
		} else {
			offset = 0
		}
	}
	if pg > g {
		pg = g
	}
	*(*int32)(unsafe.Pointer(T0_)) = int32(2)*T + offset
	if *(*int32)(unsafe.Pointer(T0_)) < minperiod0 {
		*(*int32)(unsafe.Pointer(T0_)) = minperiod0
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _53
	_53:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _55
_55:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
	return pg
}

const Q15ONE2 = "1.0f"

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

/*A range encoder.
  See entdec.c and the references for implementation details \cite{Mar79,MNW98}.

  @INPROCEEDINGS{Mar79,
   author="Martin, G.N.N.",
   title="Range encoding: an algorithm for removing redundancy from a digitised
    message",
   booktitle="Video \& Data Recording Conference",
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
   URL="http://www.stanford.edu/class/ee398/handouts/papers/Moffat98ArithmCoding.pdf"
  }*/

func ec_write_byte(tls *libc.TLS, _this uintptr, _value uint32) (r int32) {
	var v1 OpusT_opus_uint32
	var v2 uintptr
	_, _ = v1, v2
	if (*OpusT_ec_enc)(unsafe.Pointer(_this)).Foffs+(*OpusT_ec_enc)(unsafe.Pointer(_this)).Fend_offs >= (*OpusT_ec_enc)(unsafe.Pointer(_this)).Fstorage {
		return -int32(1)
	}
	v2 = _this + 28
	v1 = *(*OpusT_opus_uint32)(unsafe.Pointer(v2))
	*(*OpusT_opus_uint32)(unsafe.Pointer(v2)) = *(*OpusT_opus_uint32)(unsafe.Pointer(v2)) + 1
	*(*uint8)(unsafe.Pointer((*OpusT_ec_enc)(unsafe.Pointer(_this)).Fbuf + uintptr(v1))) = uint8(_value)
	return 0
}

func ec_write_byte_at_end(tls *libc.TLS, _this uintptr, _value uint32) (r int32) {
	var v1 OpusT_opus_uint32
	var v2 uintptr
	_, _ = v1, v2
	if (*OpusT_ec_enc)(unsafe.Pointer(_this)).Foffs+(*OpusT_ec_enc)(unsafe.Pointer(_this)).Fend_offs >= (*OpusT_ec_enc)(unsafe.Pointer(_this)).Fstorage {
		return -int32(1)
	}
	v2 = _this + 12
	*(*OpusT_opus_uint32)(unsafe.Pointer(v2)) = *(*OpusT_opus_uint32)(unsafe.Pointer(v2)) + 1
	v1 = *(*OpusT_opus_uint32)(unsafe.Pointer(v2))
	*(*uint8)(unsafe.Pointer((*OpusT_ec_enc)(unsafe.Pointer(_this)).Fbuf + uintptr((*OpusT_ec_enc)(unsafe.Pointer(_this)).Fstorage-v1))) = uint8(_value)
	return 0
}

// C documentation
//
//	/*Outputs a symbol, with a carry bit.
//	  If there is a potential to propagate a carry over several symbols, they are
//	   buffered until it can be determined whether or not an actual carry will
//	   occur.
//	  If the counter for the buffered symbols overflows, then the stream becomes
//	   undecodable.
//	  This gives a theoretical limit of a few billion symbols in a single packet on
//	   32-bit systems.
//	  The alternative is to truncate the range in order to force a carry, but
//	   requires similar carry tracking in the decoder, needlessly slowing it down.*/
func ec_enc_carry_out(tls *libc.TLS, _this uintptr, _c int32) {
	var carry int32
	var sym uint32
	var v1 OpusT_opus_uint32
	var v2 uintptr
	_, _, _, _ = carry, sym, v1, v2
	if uint32(_c) != uint32(1)<<int32(EC_SYM_BITS)-uint32(1) {
		carry = _c >> int32(EC_SYM_BITS)
		/*Don't output a byte on the first write.
		  This compare should be taken care of by branch-prediction thereafter.*/
		if (*OpusT_ec_enc)(unsafe.Pointer(_this)).Frem >= 0 {
			*(*int32)(unsafe.Pointer(_this + 48)) |= ec_write_byte(tls, _this, uint32((*OpusT_ec_enc)(unsafe.Pointer(_this)).Frem+carry))
		}
		if (*OpusT_ec_enc)(unsafe.Pointer(_this)).Fext > uint32(0) {
			sym = (uint32(1)<<int32(EC_SYM_BITS) - uint32(1) + uint32(carry)) & (uint32(1)<<int32(EC_SYM_BITS) - uint32(1))
			for {
				*(*int32)(unsafe.Pointer(_this + 48)) |= ec_write_byte(tls, _this, sym)
				goto _3
			_3:
				;
				v2 = _this + 40
				*(*OpusT_opus_uint32)(unsafe.Pointer(v2)) = *(*OpusT_opus_uint32)(unsafe.Pointer(v2)) - 1
				v1 = *(*OpusT_opus_uint32)(unsafe.Pointer(v2))
				if !(v1 > uint32(0)) {
					break
				}
			}
		}
		(*OpusT_ec_enc)(unsafe.Pointer(_this)).Frem = int32(uint32(_c) & (uint32(1)<<int32(EC_SYM_BITS) - uint32(1)))
	} else {
		(*OpusT_ec_enc)(unsafe.Pointer(_this)).Fext = (*OpusT_ec_enc)(unsafe.Pointer(_this)).Fext + 1
	}
}

func ec_enc_normalize(tls *libc.TLS, _this uintptr) {
	/*If the range is too small, output some bits and rescale it.*/
	for (*OpusT_ec_enc)(unsafe.Pointer(_this)).Frng <= uint32(1)<<(int32(EC_CODE_BITS)-int32(1))>>int32(EC_SYM_BITS) {
		ec_enc_carry_out(tls, _this, int32((*OpusT_ec_enc)(unsafe.Pointer(_this)).Fval>>(int32(EC_CODE_BITS)-int32(EC_SYM_BITS)-int32(1))))
		/*Move the next-to-high-order symbol into the high-order position.*/
		(*OpusT_ec_enc)(unsafe.Pointer(_this)).Fval = (*OpusT_ec_enc)(unsafe.Pointer(_this)).Fval << int32(EC_SYM_BITS) & (uint32(1)<<(int32(EC_CODE_BITS)-int32(1)) - uint32(1))
		*(*OpusT_opus_uint32)(unsafe.Pointer(_this + 32)) <<= uint32(int32(EC_SYM_BITS))
		*(*int32)(unsafe.Pointer(_this + 24)) += int32(EC_SYM_BITS)
	}
}

func Opus_ec_enc_init(tls *libc.TLS, _this uintptr, _buf uintptr, _size OpusT_opus_uint32) {
	(*OpusT_ec_enc)(unsafe.Pointer(_this)).Fbuf = _buf
	(*OpusT_ec_enc)(unsafe.Pointer(_this)).Fend_offs = uint32(0)
	(*OpusT_ec_enc)(unsafe.Pointer(_this)).Fend_window = uint32(0)
	(*OpusT_ec_enc)(unsafe.Pointer(_this)).Fnend_bits = 0
	/*This is the offset from which ec_tell() will subtract partial bits.*/
	(*OpusT_ec_enc)(unsafe.Pointer(_this)).Fnbits_total = int32(EC_CODE_BITS) + int32(1)
	(*OpusT_ec_enc)(unsafe.Pointer(_this)).Foffs = uint32(0)
	(*OpusT_ec_enc)(unsafe.Pointer(_this)).Frng = uint32(1) << (int32(EC_CODE_BITS) - int32(1))
	(*OpusT_ec_enc)(unsafe.Pointer(_this)).Frem = -int32(1)
	(*OpusT_ec_enc)(unsafe.Pointer(_this)).Fval = uint32(0)
	(*OpusT_ec_enc)(unsafe.Pointer(_this)).Fext = uint32(0)
	(*OpusT_ec_enc)(unsafe.Pointer(_this)).Fstorage = _size
	(*OpusT_ec_enc)(unsafe.Pointer(_this)).Ferror1 = 0
}

func Opus_ec_encode(tls *libc.TLS, _this uintptr, _fl uint32, _fh uint32, _ft uint32) {
	var r, v1, v2 OpusT_opus_uint32
	_, _, _ = r, v1, v2
	v1 = _ft
	_ = v1 > uint32(0)
	v2 = (*OpusT_ec_enc)(unsafe.Pointer(_this)).Frng / v1
	goto _3
_3:
	r = v2
	if _fl > uint32(0) {
		*(*OpusT_opus_uint32)(unsafe.Pointer(_this + 36)) += (*OpusT_ec_enc)(unsafe.Pointer(_this)).Frng - r*(_ft-_fl)
		(*OpusT_ec_enc)(unsafe.Pointer(_this)).Frng = r * (_fh - _fl)
	} else {
		*(*OpusT_opus_uint32)(unsafe.Pointer(_this + 32)) -= r * (_ft - _fh)
	}
	ec_enc_normalize(tls, _this)
}

func Opus_ec_encode_bin(tls *libc.TLS, _this uintptr, _fl uint32, _fh uint32, _bits uint32) {
	var r OpusT_opus_uint32
	_ = r
	r = (*OpusT_ec_enc)(unsafe.Pointer(_this)).Frng >> _bits
	if _fl > uint32(0) {
		*(*OpusT_opus_uint32)(unsafe.Pointer(_this + 36)) += (*OpusT_ec_enc)(unsafe.Pointer(_this)).Frng - r*(uint32(1)<<_bits-_fl)
		(*OpusT_ec_enc)(unsafe.Pointer(_this)).Frng = r * (_fh - _fl)
	} else {
		*(*OpusT_opus_uint32)(unsafe.Pointer(_this + 32)) -= r * (uint32(1)<<_bits - _fh)
	}
	ec_enc_normalize(tls, _this)
}

// C documentation
//
//	/*The probability of having a "one" is 1/(1<<_logp).*/
func Opus_ec_enc_bit_logp(tls *libc.TLS, _this uintptr, _val int32, _logp uint32) {
	var l, r, s OpusT_opus_uint32
	var v1 uint32
	_, _, _, _ = l, r, s, v1
	r = (*OpusT_ec_enc)(unsafe.Pointer(_this)).Frng
	l = (*OpusT_ec_enc)(unsafe.Pointer(_this)).Fval
	s = r >> _logp
	r = r - s
	if _val != 0 {
		(*OpusT_ec_enc)(unsafe.Pointer(_this)).Fval = l + r
	}
	if _val != 0 {
		v1 = s
	} else {
		v1 = r
	}
	(*OpusT_ec_enc)(unsafe.Pointer(_this)).Frng = v1
	ec_enc_normalize(tls, _this)
}

func Opus_ec_enc_icdf(tls *libc.TLS, _this uintptr, _s int32, _icdf uintptr, _ftb uint32) {
	var r OpusT_opus_uint32
	_ = r
	r = (*OpusT_ec_enc)(unsafe.Pointer(_this)).Frng >> _ftb
	if _s > 0 {
		*(*OpusT_opus_uint32)(unsafe.Pointer(_this + 36)) += (*OpusT_ec_enc)(unsafe.Pointer(_this)).Frng - r*uint32(*(*uint8)(unsafe.Pointer(_icdf + uintptr(_s-int32(1)))))
		(*OpusT_ec_enc)(unsafe.Pointer(_this)).Frng = r * uint32(int32(*(*uint8)(unsafe.Pointer(_icdf + uintptr(_s-int32(1)))))-int32(*(*uint8)(unsafe.Pointer(_icdf + uintptr(_s)))))
	} else {
		*(*OpusT_opus_uint32)(unsafe.Pointer(_this + 32)) -= r * uint32(*(*uint8)(unsafe.Pointer(_icdf + uintptr(_s))))
	}
	ec_enc_normalize(tls, _this)
}

func Opus_ec_enc_icdf16(tls *libc.TLS, _this uintptr, _s int32, _icdf uintptr, _ftb uint32) {
	var r OpusT_opus_uint32
	_ = r
	r = (*OpusT_ec_enc)(unsafe.Pointer(_this)).Frng >> _ftb
	if _s > 0 {
		*(*OpusT_opus_uint32)(unsafe.Pointer(_this + 36)) += (*OpusT_ec_enc)(unsafe.Pointer(_this)).Frng - r*uint32(*(*OpusT_opus_uint16)(unsafe.Pointer(_icdf + uintptr(_s-int32(1))*2)))
		(*OpusT_ec_enc)(unsafe.Pointer(_this)).Frng = r * uint32(int32(*(*OpusT_opus_uint16)(unsafe.Pointer(_icdf + uintptr(_s-int32(1))*2)))-int32(*(*OpusT_opus_uint16)(unsafe.Pointer(_icdf + uintptr(_s)*2))))
	} else {
		*(*OpusT_opus_uint32)(unsafe.Pointer(_this + 32)) -= r * uint32(*(*OpusT_opus_uint16)(unsafe.Pointer(_icdf + uintptr(_s)*2)))
	}
	ec_enc_normalize(tls, _this)
}

func Opus_ec_enc_uint(tls *libc.TLS, _this uintptr, _fl OpusT_opus_uint32, _ft OpusT_opus_uint32) {
	var fl, ft uint32
	var ftb int32
	_, _, _ = fl, ft, ftb
	/*In order to optimize EC_ILOG(), it is undefined for the value 0.*/
	if !(_ft > uint32(1)) {
		Opus_celt_fatal(tls, __ccgo_ts+3569, __ccgo_ts+4699, int32(191))
	}
	_ft = _ft - 1
	ftb = int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, _ft)
	if ftb > int32(EC_UINT_BITS) {
		ftb = ftb - int32(EC_UINT_BITS)
		ft = _ft>>ftb + uint32(1)
		fl = _fl >> ftb
		Opus_ec_encode(tls, _this, fl, fl+uint32(1), ft)
		Opus_ec_enc_bits(tls, _this, _fl&(uint32(1)<<ftb-uint32(1)), uint32(ftb))
	} else {
		Opus_ec_encode(tls, _this, _fl, _fl+uint32(1), _ft+uint32(1))
	}
}

func Opus_ec_enc_bits(tls *libc.TLS, _this uintptr, _fl OpusT_opus_uint32, _bits uint32) {
	var used int32
	var window OpusT_ec_window
	var v1 uintptr
	_, _, _ = used, window, v1
	window = (*OpusT_ec_enc)(unsafe.Pointer(_this)).Fend_window
	used = (*OpusT_ec_enc)(unsafe.Pointer(_this)).Fnend_bits
	if !(_bits > uint32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+4716, __ccgo_ts+4699, int32(209))
	}
	if uint32(used)+_bits > uint32(int32(4)*int32(CHAR_BIT)) {
		for cond := true; cond; cond = used >= int32(EC_SYM_BITS) {
			*(*int32)(unsafe.Pointer(_this + 48)) |= ec_write_byte_at_end(tls, _this, window&(uint32(1)<<int32(EC_SYM_BITS)-uint32(1)))
			window = window >> uint32(int32(EC_SYM_BITS))
			used = used - int32(EC_SYM_BITS)
		}
	}
	window = window | _fl<<used
	used = int32(uint32(used) + _bits)
	(*OpusT_ec_enc)(unsafe.Pointer(_this)).Fend_window = window
	(*OpusT_ec_enc)(unsafe.Pointer(_this)).Fnend_bits = used
	v1 = _this + 24
	*(*int32)(unsafe.Pointer(v1)) = int32(uint32(*(*int32)(unsafe.Pointer(v1))) + _bits)
}

func Opus_ec_enc_patch_initial_bits(tls *libc.TLS, _this uintptr, _val uint32, _nbits uint32) {
	var mask uint32
	var shift int32
	_, _ = mask, shift
	if !(_nbits <= uint32(int32(EC_SYM_BITS))) {
		Opus_celt_fatal(tls, __ccgo_ts+4742, __ccgo_ts+4699, int32(228))
	}
	shift = int32(uint32(int32(EC_SYM_BITS)) - _nbits)
	mask = uint32((int32(1)<<_nbits - int32(1)) << shift)
	if (*OpusT_ec_enc)(unsafe.Pointer(_this)).Foffs > uint32(0) {
		/*The first byte has been finalized.*/
		*(*uint8)(unsafe.Pointer((*OpusT_ec_enc)(unsafe.Pointer(_this)).Fbuf)) = uint8(uint32(*(*uint8)(unsafe.Pointer((*OpusT_ec_enc)(unsafe.Pointer(_this)).Fbuf))) & ^mask | _val<<shift)
	} else {
		if (*OpusT_ec_enc)(unsafe.Pointer(_this)).Frem >= 0 {
			/*The first byte is still awaiting carry propagation.*/
			(*OpusT_ec_enc)(unsafe.Pointer(_this)).Frem = int32(uint32((*OpusT_ec_enc)(unsafe.Pointer(_this)).Frem) & ^mask | _val<<shift)
		} else {
			if (*OpusT_ec_enc)(unsafe.Pointer(_this)).Frng <= uint32(1)<<(int32(EC_CODE_BITS)-int32(1))>>_nbits {
				/*The renormalization loop has never been run.*/
				(*OpusT_ec_enc)(unsafe.Pointer(_this)).Fval = (*OpusT_ec_enc)(unsafe.Pointer(_this)).Fval & ^(mask<<(int32(EC_CODE_BITS)-int32(EC_SYM_BITS)-int32(1))) | _val<<(int32(EC_CODE_BITS)-int32(EC_SYM_BITS)-int32(1)+shift)
			} else {
				(*OpusT_ec_enc)(unsafe.Pointer(_this)).Ferror1 = -int32(1)
			}
		}
	}
}

func Opus_ec_enc_shrink(tls *libc.TLS, _this uintptr, _size OpusT_opus_uint32) {
	if !((*OpusT_ec_enc)(unsafe.Pointer(_this)).Foffs+(*OpusT_ec_enc)(unsafe.Pointer(_this)).Fend_offs <= _size) {
		Opus_celt_fatal(tls, __ccgo_ts+4780, __ccgo_ts+4699, int32(249))
	}
	libc.Xmemmove(tls, (*OpusT_ec_enc)(unsafe.Pointer(_this)).Fbuf+uintptr(_size)-uintptr((*OpusT_ec_enc)(unsafe.Pointer(_this)).Fend_offs), (*OpusT_ec_enc)(unsafe.Pointer(_this)).Fbuf+uintptr((*OpusT_ec_enc)(unsafe.Pointer(_this)).Fstorage)-uintptr((*OpusT_ec_enc)(unsafe.Pointer(_this)).Fend_offs), uint64((*OpusT_ec_enc)(unsafe.Pointer(_this)).Fend_offs)*uint64(1)+uint64(0*(int64((*OpusT_ec_enc)(unsafe.Pointer(_this)).Fbuf+uintptr(_size)-uintptr((*OpusT_ec_enc)(unsafe.Pointer(_this)).Fend_offs))-int64((*OpusT_ec_enc)(unsafe.Pointer(_this)).Fbuf+uintptr((*OpusT_ec_enc)(unsafe.Pointer(_this)).Fstorage)-uintptr((*OpusT_ec_enc)(unsafe.Pointer(_this)).Fend_offs)))))
	(*OpusT_ec_enc)(unsafe.Pointer(_this)).Fstorage = _size
}

func Opus_ec_enc_done(tls *libc.TLS, _this uintptr) {
	var end, msk OpusT_opus_uint32
	var l, used int32
	var window OpusT_ec_window
	var v1 uintptr
	_, _, _, _, _, _ = end, l, msk, used, window, v1
	/*We output the minimum number of bits that ensures that the symbols encoded
	  thus far will be decoded correctly regardless of the bits that follow.*/
	l = int32(EC_CODE_BITS) - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_enc)(unsafe.Pointer(_this)).Frng))
	msk = (uint32(1)<<(int32(EC_CODE_BITS)-int32(1)) - uint32(1)) >> l
	end = ((*OpusT_ec_enc)(unsafe.Pointer(_this)).Fval + msk) & ^msk
	if end|msk >= (*OpusT_ec_enc)(unsafe.Pointer(_this)).Fval+(*OpusT_ec_enc)(unsafe.Pointer(_this)).Frng {
		l = l + 1
		msk = msk >> uint32(1)
		end = ((*OpusT_ec_enc)(unsafe.Pointer(_this)).Fval + msk) & ^msk
	}
	for l > 0 {
		ec_enc_carry_out(tls, _this, int32(end>>(int32(EC_CODE_BITS)-int32(EC_SYM_BITS)-int32(1))))
		end = end << int32(EC_SYM_BITS) & (uint32(1)<<(int32(EC_CODE_BITS)-int32(1)) - uint32(1))
		l = l - int32(EC_SYM_BITS)
	}
	/*If we have a buffered byte flush it into the output buffer.*/
	if (*OpusT_ec_enc)(unsafe.Pointer(_this)).Frem >= 0 || (*OpusT_ec_enc)(unsafe.Pointer(_this)).Fext > uint32(0) {
		ec_enc_carry_out(tls, _this, 0)
	}
	/*If we have buffered extra bits, flush them as well.*/
	window = (*OpusT_ec_enc)(unsafe.Pointer(_this)).Fend_window
	used = (*OpusT_ec_enc)(unsafe.Pointer(_this)).Fnend_bits
	for used >= int32(EC_SYM_BITS) {
		*(*int32)(unsafe.Pointer(_this + 48)) |= ec_write_byte_at_end(tls, _this, window&(uint32(1)<<int32(EC_SYM_BITS)-uint32(1)))
		window = window >> uint32(int32(EC_SYM_BITS))
		used = used - int32(EC_SYM_BITS)
	}
	/*Clear any excess space and add any remaining extra bits to the last byte.*/
	if !((*OpusT_ec_enc)(unsafe.Pointer(_this)).Ferror1 != 0) {
		if (*OpusT_ec_enc)(unsafe.Pointer(_this)).Fbuf != 0 {
			libc.Xmemset(tls, (*OpusT_ec_enc)(unsafe.Pointer(_this)).Fbuf+uintptr((*OpusT_ec_enc)(unsafe.Pointer(_this)).Foffs), 0, uint64((*OpusT_ec_enc)(unsafe.Pointer(_this)).Fstorage-(*OpusT_ec_enc)(unsafe.Pointer(_this)).Foffs-(*OpusT_ec_enc)(unsafe.Pointer(_this)).Fend_offs)*uint64(1))
		}
		if used > 0 {
			/*If there's no range coder data at all, give up.*/
			if (*OpusT_ec_enc)(unsafe.Pointer(_this)).Fend_offs >= (*OpusT_ec_enc)(unsafe.Pointer(_this)).Fstorage {
				(*OpusT_ec_enc)(unsafe.Pointer(_this)).Ferror1 = -int32(1)
			} else {
				l = -l
				/*If we've busted, don't add too many extra bits to the last byte; it
				  would corrupt the range coder data, and that's more important.*/
				if (*OpusT_ec_enc)(unsafe.Pointer(_this)).Foffs+(*OpusT_ec_enc)(unsafe.Pointer(_this)).Fend_offs >= (*OpusT_ec_enc)(unsafe.Pointer(_this)).Fstorage && l < used {
					window = window & uint32(int32(1)<<l-int32(1))
					(*OpusT_ec_enc)(unsafe.Pointer(_this)).Ferror1 = -int32(1)
				}
				v1 = (*OpusT_ec_enc)(unsafe.Pointer(_this)).Fbuf + uintptr((*OpusT_ec_enc)(unsafe.Pointer(_this)).Fstorage-(*OpusT_ec_enc)(unsafe.Pointer(_this)).Fend_offs-uint32(1))
				*(*uint8)(unsafe.Pointer(v1)) = uint8(int32(*(*uint8)(unsafe.Pointer(v1))) | int32(uint8(window)))
			}
		}
	}
}

var trim_icdf11 = [11]uint8{
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
var spread_icdf11 = [4]uint8{
	0: uint8(25),
	1: uint8(23),
	2: uint8(2),
}
var tapset_icdf11 = [3]uint8{
	0: uint8(2),
	1: uint8(1),
}

var log2_x_norm_coeff10 = [8]float32{
	0: float32(1),
	1: float32(0.8888888955116272),
	2: float32(0.8),
	3: float32(0.7272727489471436),
	4: float32(0.6666666865348816),
	5: float32(0.6153846383094788),
	6: float32(0.5714285969734192),
	7: float32(0.5333333611488342),
}
var log2_y_norm_coeff10 = [8]float32{
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
//	/* prediction coefficients: 0.9, 0.8, 0.65, 0.5 */
var pred_coef = [4]OpusT_opus_val16{
	0: float32(float64(29440) / float64(32768)),
	1: float32(float64(26112) / float64(32768)),
	2: float32(float64(21248) / float64(32768)),
	3: float32(float64(16384) / float64(32768)),
}
var beta_coef = [4]OpusT_opus_val16{
	0: float32(float64(30147) / float64(32768)),
	1: float32(float64(22282) / float64(32768)),
	2: float32(float64(12124) / float64(32768)),
	3: float32(float64(6554) / float64(32768)),
}
var beta_intra = float32(float64(4915) / float64(32768))

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

func loss_distortion(tls *libc.TLS, eBands uintptr, oldEBands uintptr, start int32, end int32, len1 int32, C int32) (r OpusT_opus_val32) {
	var c, i, v1 int32
	var d OpusT_celt_glog
	var dist, v4 OpusT_opus_val32
	_, _, _, _, _, _ = c, d, dist, i, v1, v4
	dist = float32(0)
	c = 0
	for {
		i = start
		for {
			if !(i < end) {
				break
			}
			d = *(*OpusT_celt_glog)(unsafe.Pointer(eBands + uintptr(i+c*len1)*4)) - *(*OpusT_celt_glog)(unsafe.Pointer(oldEBands + uintptr(i+c*len1)*4))
			dist = dist + OpusT_opus_val32(d*d)
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
	if float32(int32(200)) < dist {
		v4 = float32(int32(200))
	} else {
		v4 = dist
	}
	return v4
}

func quant_coarse_energy_impl(tls *libc.TLS, m uintptr, start int32, end int32, eBands uintptr, oldEBands uintptr, budget OpusT_opus_int32, tell OpusT_opus_int32, prob_model uintptr, error1 uintptr, enc uintptr, C int32, LM int32, intra int32, max_decay OpusT_celt_glog, lfe int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var badness, bits_left, c, i, pi, qi0, v2, v7, v9 int32
	var beta, coef OpusT_opus_val16
	var decay_bound, oldE, x OpusT_celt_glog
	var f, q, tmp OpusT_opus_val32
	var prev [2]OpusT_opus_val32
	var v4 float32
	var v6 uintptr
	var _ /* qi at bp+0 */ int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = badness, beta, bits_left, c, coef, decay_bound, f, i, oldE, pi, prev, q, qi0, tmp, x, v2, v4, v6, v7, v9
	badness = 0
	prev = [2]OpusT_opus_val32{}
	if tell+int32(3) <= budget {
		Opus_ec_enc_bit_logp(tls, enc, intra, uint32(3))
	}
	if intra != 0 {
		coef = float32(0)
		beta = beta_intra
	} else {
		beta = beta_coef[LM]
		coef = pred_coef[LM]
	}
	/* Encode at a fixed coarse resolution */
	i = start
	for {
		if !(i < end) {
			break
		}
		c = 0
		for {
			x = *(*OpusT_celt_glog)(unsafe.Pointer(eBands + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4))
			if -float32(9) > *(*OpusT_celt_glog)(unsafe.Pointer(oldEBands + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4)) {
				v4 = -float32(9)
			} else {
				v4 = *(*OpusT_celt_glog)(unsafe.Pointer(oldEBands + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4))
			}
			oldE = v4
			f = x - OpusT_celt_glog(coef*oldE) - prev[c]
			/* Rounding to nearest integer here is really important! */
			*(*int32)(unsafe.Pointer(bp)) = int32(libc.Xfloor(tls, float64(float32(0.5)+f)))
			if -float32(28) > *(*OpusT_celt_glog)(unsafe.Pointer(oldEBands + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4)) {
				v4 = -float32(28)
			} else {
				v4 = *(*OpusT_celt_glog)(unsafe.Pointer(oldEBands + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4))
			}
			decay_bound = v4 - max_decay
			/* Prevent the energy from going down too quickly (e.g. for bands
			   that have just one bin) */
			if *(*int32)(unsafe.Pointer(bp)) < 0 && x < decay_bound {
				*(*int32)(unsafe.Pointer(bp)) = *(*int32)(unsafe.Pointer(bp)) + int32(decay_bound-x)
				if *(*int32)(unsafe.Pointer(bp)) > 0 {
					*(*int32)(unsafe.Pointer(bp)) = 0
				}
			}
			qi0 = *(*int32)(unsafe.Pointer(bp))
			/* If we don't have enough bits to encode all the energy, just assume
			   something safe. */
			v6 = enc
			v2 = (*OpusT_ec_ctx)(unsafe.Pointer(v6)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v6)).Frng))
			goto _8
		_8:
			tell = v2
			bits_left = budget - tell - int32(3)*C*(end-i)
			if i != start && bits_left < int32(30) {
				if bits_left < int32(24) {
					if int32(1) < *(*int32)(unsafe.Pointer(bp)) {
						v2 = int32(1)
					} else {
						v2 = *(*int32)(unsafe.Pointer(bp))
					}
					*(*int32)(unsafe.Pointer(bp)) = v2
				}
				if bits_left < int32(16) {
					if -int32(1) > *(*int32)(unsafe.Pointer(bp)) {
						v2 = -int32(1)
					} else {
						v2 = *(*int32)(unsafe.Pointer(bp))
					}
					*(*int32)(unsafe.Pointer(bp)) = v2
				}
			}
			if lfe != 0 && i >= int32(2) {
				if *(*int32)(unsafe.Pointer(bp)) < 0 {
					v2 = *(*int32)(unsafe.Pointer(bp))
				} else {
					v2 = 0
				}
				*(*int32)(unsafe.Pointer(bp)) = v2
			}
			if budget-tell >= int32(15) {
				if i < int32(20) {
					v2 = i
				} else {
					v2 = int32(20)
				}
				pi = int32(2) * v2
				Opus_ec_laplace_encode(tls, enc, bp, uint32(int32(*(*uint8)(unsafe.Pointer(prob_model + uintptr(pi))))<<int32(7)), int32(*(*uint8)(unsafe.Pointer(prob_model + uintptr(pi+int32(1)))))<<int32(6))
			} else {
				if budget-tell >= int32(2) {
					if *(*int32)(unsafe.Pointer(bp)) < int32(1) {
						v7 = *(*int32)(unsafe.Pointer(bp))
					} else {
						v7 = int32(1)
					}
					if -int32(1) > v7 {
						v2 = -int32(1)
					} else {
						if *(*int32)(unsafe.Pointer(bp)) < int32(1) {
							v9 = *(*int32)(unsafe.Pointer(bp))
						} else {
							v9 = int32(1)
						}
						v2 = v9
					}
					*(*int32)(unsafe.Pointer(bp)) = v2
					Opus_ec_enc_icdf(tls, enc, int32(2)**(*int32)(unsafe.Pointer(bp))^-libc.BoolInt32(*(*int32)(unsafe.Pointer(bp)) < 0), uintptr(unsafe.Pointer(&small_energy_icdf)), uint32(2))
				} else {
					if budget-tell >= int32(1) {
						if 0 < *(*int32)(unsafe.Pointer(bp)) {
							v2 = 0
						} else {
							v2 = *(*int32)(unsafe.Pointer(bp))
						}
						*(*int32)(unsafe.Pointer(bp)) = v2
						Opus_ec_enc_bit_logp(tls, enc, -*(*int32)(unsafe.Pointer(bp)), uint32(1))
					} else {
						*(*int32)(unsafe.Pointer(bp)) = -int32(1)
					}
				}
			}
			*(*OpusT_celt_glog)(unsafe.Pointer(error1 + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4)) = f - float32(*(*int32)(unsafe.Pointer(bp)))
			badness = badness + libc.Xabs(tls, qi0-*(*int32)(unsafe.Pointer(bp)))
			q = float32(*(*int32)(unsafe.Pointer(bp)))
			tmp = OpusT_opus_val16(coef*oldE) + prev[c] + q
			*(*OpusT_celt_glog)(unsafe.Pointer(oldEBands + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4)) = tmp
			prev[c] = prev[c] + q - OpusT_opus_val16(beta*q)
			goto _3
		_3:
			;
			c = c + 1
			v2 = c
			if !(v2 < C) {
				break
			}
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	if lfe != 0 {
		v2 = 0
	} else {
		v2 = badness
	}
	return v2
}

func Opus_quant_coarse_energy(tls *libc.TLS, m uintptr, start int32, end int32, effEnd int32, eBands uintptr, oldEBands uintptr, budget OpusT_opus_uint32, error1 uintptr, enc uintptr, C int32, LM int32, nbAvailableBytes int32, force_intra int32, delayedIntra uintptr, two_pass int32, loss_rate int32, lfe int32) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var _saved_stack, error_intra, intra_bits, intra_buf, oldEBands_intra, st, v1, v10, v12, v14, v16, v18, v20, v22, v24, v26, v3, v5 uintptr
	var badness1, badness2, intra, v6 int32
	var intra_bias, tell_intra OpusT_opus_int32
	var max_decay, v9 OpusT_celt_glog
	var new_distortion OpusT_opus_val32
	var nintra_bytes, nstart_bytes, save_bytes, tell, v58 OpusT_opus_uint32
	var _ /* enc_intra_state at bp+56 */ OpusT_ec_enc
	var _ /* enc_start_state at bp+0 */ OpusT_ec_enc
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = _saved_stack, badness1, badness2, error_intra, intra, intra_bias, intra_bits, intra_buf, max_decay, new_distortion, nintra_bytes, nstart_bytes, oldEBands_intra, save_bytes, st, tell, tell_intra, v1, v10, v12, v14, v16, v18, v20, v22, v24, v26, v3, v5, v58, v6, v9
	badness1 = 0
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
	intra = libc.BoolInt32(force_intra != 0 || !(two_pass != 0) && *(*OpusT_opus_val32)(unsafe.Pointer(delayedIntra)) > OpusT_opus_val32(int32(2)*C*(end-start)) && nbAvailableBytes > (end-start)*C)
	intra_bias = int32(OpusT_opus_val32(OpusT_opus_val32(float32(budget)**(*OpusT_opus_val32)(unsafe.Pointer(delayedIntra)))*float32(loss_rate)) / float32(C*int32(512)))
	new_distortion = loss_distortion(tls, eBands, oldEBands, start, effEnd, (*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands, C)
	v1 = enc
	v6 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
	goto _7
_7:
	tell = uint32(v6)
	if tell+uint32(3) > budget {
		v6 = int32(0)
		intra = v6
		two_pass = v6
	}
	max_decay = float32(16)
	if end-start > int32(10) {
		if max_decay < float32(float32(0.125)*float32(nbAvailableBytes)) {
			v9 = max_decay
		} else {
			v9 = float32(float32(0.125) * float32(nbAvailableBytes))
		}
		max_decay = v9
	}
	if lfe != 0 {
		max_decay = float32(3)
	}
	*(*OpusT_ec_enc)(unsafe.Pointer(bp)) = *(*OpusT_ec_enc)(unsafe.Pointer(enc))
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
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _15
	_15:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	goto _17
_17:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v12 = libc.Xmalloc(tls, uint64(16))
		goto _19
	_19:
		st = v12
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v14 = st
	goto _21
_21:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v16 = libc.Xmalloc(tls, uint64(16))
		goto _23
	_23:
		st = v16
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v18 = st
	goto _25
_25:
	if !(int64(int32(uint64(uint32(C*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v14)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v18)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4833, int32(297))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v20 = libc.Xmalloc(tls, uint64(16))
		goto _27
	_27:
		st = v20
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v22 = st
	goto _29
_29:
	*(*uintptr)(unsafe.Pointer(v22 + 8)) += uintptr(uint64(uint32(C*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v24 = libc.Xmalloc(tls, uint64(16))
		goto _31
	_31:
		st = v24
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v26 = st
	goto _33
_33:
	oldEBands_intra = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v26)).Fglobal_stack - uintptr(uint64(uint32(C*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands))*(uint64(4)/uint64(1)))
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
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _39
	_39:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	goto _41
_41:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v12 = libc.Xmalloc(tls, uint64(16))
		goto _43
	_43:
		st = v12
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v14 = st
	goto _45
_45:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v16 = libc.Xmalloc(tls, uint64(16))
		goto _47
	_47:
		st = v16
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v18 = st
	goto _49
_49:
	if !(int64(int32(uint64(uint32(C*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v14)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v18)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4833, int32(298))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v20 = libc.Xmalloc(tls, uint64(16))
		goto _51
	_51:
		st = v20
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v22 = st
	goto _53
_53:
	*(*uintptr)(unsafe.Pointer(v22 + 8)) += uintptr(uint64(uint32(C*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v24 = libc.Xmalloc(tls, uint64(16))
		goto _55
	_55:
		st = v24
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v26 = st
	goto _57
_57:
	error_intra = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v26)).Fglobal_stack - uintptr(uint64(uint32(C*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands))*(uint64(4)/uint64(1)))
	libc.Xmemcpy(tls, oldEBands_intra, oldEBands, uint64(uint32(C*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands))*uint64(4)+uint64(0*((int64(oldEBands_intra)-int64(oldEBands))/4)))
	if two_pass != 0 || intra != 0 {
		badness1 = quant_coarse_energy_impl(tls, m, start, end, eBands, oldEBands_intra, int32(budget), int32(tell), uintptr(unsafe.Pointer(&e_prob_model))+uintptr(LM)*84+1*42, error_intra, enc, C, LM, int32(1), max_decay, lfe)
	}
	if !(intra != 0) {
		tell_intra = int32(Opus_ec_tell_frac(tls, enc))
		*(*OpusT_ec_enc)(unsafe.Pointer(bp + 56)) = *(*OpusT_ec_enc)(unsafe.Pointer(enc))
		v58 = (*OpusT_ec_ctx)(unsafe.Pointer(bp)).Foffs
		goto _59
	_59:
		nstart_bytes = v58
		v58 = (*OpusT_ec_ctx)(unsafe.Pointer(bp + 56)).Foffs
		goto _61
	_61:
		nintra_bytes = v58
		v1 = (*OpusT_ec_ctx)(unsafe.Pointer(bp + 56)).Fbuf
		goto _63
	_63:
		intra_buf = v1 + uintptr(nstart_bytes)
		save_bytes = nintra_bytes - nstart_bytes
		if save_bytes == uint32(0) {
			save_bytes = uint32(ALLOC_NONE)
		}
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
		v10 = st
		goto _71
	_71:
		*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(1)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fglobal_stack))) & (uint64(uint32(1)) - uint64(uint32(1))))
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v12 = libc.Xmalloc(tls, uint64(16))
			goto _73
		_73:
			st = v12
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v14 = st
		goto _75
	_75:
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v16 = libc.Xmalloc(tls, uint64(16))
			goto _77
		_77:
			st = v16
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v18 = st
		goto _79
	_79:
		if !(int64(int32(uint64(save_bytes)*(uint64(1)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v14)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v18)).Fglobal_stack)) {
			Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+4833, int32(328))
		}
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v20 = libc.Xmalloc(tls, uint64(16))
			goto _81
		_81:
			st = v20
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v22 = st
		goto _83
	_83:
		*(*uintptr)(unsafe.Pointer(v22 + 8)) += uintptr(uint64(save_bytes) * (uint64(1) / uint64(1)))
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v24 = libc.Xmalloc(tls, uint64(16))
			goto _85
		_85:
			st = v24
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v26 = st
		goto _87
	_87:
		intra_bits = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v26)).Fglobal_stack - uintptr(uint64(save_bytes)*(uint64(1)/uint64(1)))
		/* Copy bits from intra bit-stream */
		libc.Xmemcpy(tls, intra_bits, intra_buf, uint64(nintra_bytes-nstart_bytes)*uint64(1)+uint64(0*(int64(intra_bits)-int64(intra_buf))))
		*(*OpusT_ec_enc)(unsafe.Pointer(enc)) = *(*OpusT_ec_enc)(unsafe.Pointer(bp))
		badness2 = quant_coarse_energy_impl(tls, m, start, end, eBands, oldEBands, int32(budget), int32(tell), uintptr(unsafe.Pointer(&e_prob_model))+uintptr(LM)*84+uintptr(intra)*42, error1, enc, C, LM, 0, max_decay, lfe)
		if two_pass != 0 && (badness1 < badness2 || badness1 == badness2 && int32(Opus_ec_tell_frac(tls, enc))+intra_bias > tell_intra) {
			*(*OpusT_ec_enc)(unsafe.Pointer(enc)) = *(*OpusT_ec_enc)(unsafe.Pointer(bp + 56))
			/* Copy intra bits to bit-stream */
			libc.Xmemcpy(tls, intra_buf, intra_bits, uint64(nintra_bytes-nstart_bytes)*uint64(1)+uint64(0*(int64(intra_buf)-int64(intra_bits))))
			libc.Xmemcpy(tls, oldEBands, oldEBands_intra, uint64(uint32(C*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands))*uint64(4)+uint64(0*((int64(oldEBands)-int64(oldEBands_intra))/4)))
			libc.Xmemcpy(tls, error1, error_intra, uint64(uint32(C*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands))*uint64(4)+uint64(0*((int64(error1)-int64(error_intra))/4)))
			intra = int32(1)
		}
	} else {
		libc.Xmemcpy(tls, oldEBands, oldEBands_intra, uint64(uint32(C*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands))*uint64(4)+uint64(0*((int64(oldEBands)-int64(oldEBands_intra))/4)))
		libc.Xmemcpy(tls, error1, error_intra, uint64(uint32(C*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands))*uint64(4)+uint64(0*((int64(error1)-int64(error_intra))/4)))
	}
	if intra != 0 {
		*(*OpusT_opus_val32)(unsafe.Pointer(delayedIntra)) = new_distortion
	} else {
		*(*OpusT_opus_val32)(unsafe.Pointer(delayedIntra)) = OpusT_opus_val16(OpusT_opus_val16(pred_coef[LM]*pred_coef[LM])**(*OpusT_opus_val32)(unsafe.Pointer(delayedIntra))) + new_distortion
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _89
	_89:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _91
_91:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
}

func Opus_quant_fine_energy(tls *libc.TLS, m uintptr, start int32, end int32, oldEBands uintptr, error1 uintptr, prev_quant uintptr, extra_quant uintptr, enc uintptr, C int32) {
	var c, i, q2, v3 int32
	var extra, prev OpusT_opus_int16
	var offset OpusT_celt_glog
	var v2 uintptr
	_, _, _, _, _, _, _, _ = c, extra, i, offset, prev, q2, v2, v3
	/* Encode finer resolution */
	i = start
	for {
		if !(i < end) {
			break
		}
		extra = int16(int32(1) << *(*int32)(unsafe.Pointer(extra_quant + uintptr(i)*4)))
		if *(*int32)(unsafe.Pointer(extra_quant + uintptr(i)*4)) <= 0 {
			goto _1
		}
		v2 = enc
		v3 = (*OpusT_ec_ctx)(unsafe.Pointer(v2)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v2)).Frng))
		goto _4
	_4:
		if v3+C**(*int32)(unsafe.Pointer(extra_quant + uintptr(i)*4)) > int32((*OpusT_ec_enc)(unsafe.Pointer(enc)).Fstorage)*int32(8) {
			goto _1
		}
		if prev_quant != uintptr(uint32(0)) {
			v3 = *(*int32)(unsafe.Pointer(prev_quant + uintptr(i)*4))
		} else {
			v3 = 0
		}
		prev = int16(v3)
		c = 0
		for {
			q2 = int32(libc.Xfloor(tls, float64((OpusT_celt_glog(*(*OpusT_celt_glog)(unsafe.Pointer(error1 + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4))*float32(int32(1)<<prev))+float32(0.5))*float32(extra))))
			if q2 > int32(extra)-int32(1) {
				q2 = int32(extra) - int32(1)
			}
			if q2 < 0 {
				q2 = 0
			}
			Opus_ec_enc_bits(tls, enc, uint32(q2), uint32(*(*int32)(unsafe.Pointer(extra_quant + uintptr(i)*4))))
			offset = float32(float32((float32(q2)+float32(0.5))*float32(int32(1)<<(int32(14)-*(*int32)(unsafe.Pointer(extra_quant + uintptr(i)*4)))))*(float32(1)/float32(16384))) - float32(0.5)
			offset = offset * OpusT_celt_glog(float32(int32(1)<<(int32(14)-int32(prev)))*(float32(1)/float32(16384)))
			*(*OpusT_celt_glog)(unsafe.Pointer(oldEBands + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4)) += offset
			*(*OpusT_celt_glog)(unsafe.Pointer(error1 + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4)) -= offset
			/*printf ("%f ", error[i] - offset);*/
			goto _7
		_7:
			;
			c = c + 1
			v3 = c
			if !(v3 < C) {
				break
			}
		}
		goto _1
	_1:
		;
		i = i + 1
	}
}

func Opus_quant_energy_finalise(tls *libc.TLS, m uintptr, start int32, end int32, oldEBands uintptr, error1 uintptr, fine_quant uintptr, fine_priority uintptr, bits_left int32, enc uintptr, C int32) {
	var c, i, prio, q2, v3, v5 int32
	var offset OpusT_celt_glog
	_, _, _, _, _, _, _ = c, i, offset, prio, q2, v3, v5
	/* Use up the remaining bits */
	prio = 0
	for {
		if !(prio < int32(2)) {
			break
		}
		i = start
		for {
			if !(i < end && bits_left >= C) {
				break
			}
			if *(*int32)(unsafe.Pointer(fine_quant + uintptr(i)*4)) >= int32(MAX_FINE_BITS) || *(*int32)(unsafe.Pointer(fine_priority + uintptr(i)*4)) != prio {
				goto _2
			}
			c = 0
			for {
				if *(*OpusT_celt_glog)(unsafe.Pointer(error1 + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4)) < float32(0) {
					v5 = 0
				} else {
					v5 = int32(1)
				}
				q2 = v5
				Opus_ec_enc_bits(tls, enc, uint32(q2), uint32(1))
				offset = OpusT_celt_glog(float32((float32(q2)-float32(0.5))*float32(int32(1)<<(int32(14)-*(*int32)(unsafe.Pointer(fine_quant + uintptr(i)*4))-int32(1)))) * (float32(1) / float32(16384)))
				if oldEBands != uintptr(uint32(0)) {
					*(*OpusT_celt_glog)(unsafe.Pointer(oldEBands + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4)) += offset
				}
				*(*OpusT_celt_glog)(unsafe.Pointer(error1 + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4)) -= offset
				bits_left = bits_left - 1
				goto _4
			_4:
				;
				c = c + 1
				v3 = c
				if !(v3 < C) {
					break
				}
			}
			goto _2
		_2:
			;
			i = i + 1
		}
		goto _1
	_1:
		;
		prio = prio + 1
	}
}

func Opus_unquant_coarse_energy(tls *libc.TLS, m uintptr, start int32, end int32, oldEBands uintptr, intra int32, dec uintptr, C int32, LM int32) {
	var beta, coef OpusT_opus_val16
	var budget, tell OpusT_opus_int32
	var c, i, pi, qi, v2 int32
	var prev [2]OpusT_opus_val64
	var prob_model, v4 uintptr
	var q, tmp OpusT_opus_val32
	var v8 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = beta, budget, c, coef, i, pi, prev, prob_model, q, qi, tell, tmp, v2, v4, v8
	prob_model = uintptr(unsafe.Pointer(&e_prob_model)) + uintptr(LM)*84 + uintptr(intra)*42
	prev = [2]OpusT_opus_val64{}
	if intra != 0 {
		coef = float32(0)
		beta = beta_intra
	} else {
		beta = beta_coef[LM]
		coef = pred_coef[LM]
	}
	budget = int32((*OpusT_ec_dec)(unsafe.Pointer(dec)).Fstorage * uint32(8))
	/* Decode at a fixed coarse resolution */
	i = start
	for {
		if !(i < end) {
			break
		}
		c = 0
		for {
			/* It would be better to express this invariant as a
			   test on C at function entry, but that isn't enough
			   to make the static analyzer happy. */
			_ = c < int32(2)
			v4 = dec
			v2 = (*OpusT_ec_ctx)(unsafe.Pointer(v4)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v4)).Frng))
			goto _6
		_6:
			tell = v2
			if budget-tell >= int32(15) {
				if i < int32(20) {
					v2 = i
				} else {
					v2 = int32(20)
				}
				pi = int32(2) * v2
				qi = Opus_ec_laplace_decode(tls, dec, uint32(int32(*(*uint8)(unsafe.Pointer(prob_model + uintptr(pi))))<<int32(7)), int32(*(*uint8)(unsafe.Pointer(prob_model + uintptr(pi+int32(1)))))<<int32(6))
			} else {
				if budget-tell >= int32(2) {
					qi = Opus_ec_dec_icdf(tls, dec, uintptr(unsafe.Pointer(&small_energy_icdf)), uint32(2))
					qi = qi>>int32(1) ^ -(qi & int32(1))
				} else {
					if budget-tell >= int32(1) {
						qi = -Opus_ec_dec_bit_logp(tls, dec, uint32(1))
					} else {
						qi = -int32(1)
					}
				}
			}
			q = float32(qi)
			if -float32(9) > *(*OpusT_celt_glog)(unsafe.Pointer(oldEBands + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4)) {
				v8 = -float32(9)
			} else {
				v8 = *(*OpusT_celt_glog)(unsafe.Pointer(oldEBands + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4))
			}
			*(*OpusT_celt_glog)(unsafe.Pointer(oldEBands + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4)) = v8
			tmp = OpusT_opus_val16(coef**(*OpusT_celt_glog)(unsafe.Pointer(oldEBands + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4))) + prev[c] + q
			*(*OpusT_celt_glog)(unsafe.Pointer(oldEBands + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4)) = tmp
			prev[c] = prev[c] + q - OpusT_opus_val16(beta*q)
			goto _3
		_3:
			;
			c = c + 1
			v2 = c
			if !(v2 < C) {
				break
			}
		}
		goto _1
	_1:
		;
		i = i + 1
	}
}

func Opus_unquant_fine_energy(tls *libc.TLS, m uintptr, start int32, end int32, oldEBands uintptr, prev_quant uintptr, extra_quant uintptr, dec uintptr, C int32) {
	var c, i, q2, v3 int32
	var extra, prev OpusT_opus_int16
	var offset OpusT_celt_glog
	var v2 uintptr
	_, _, _, _, _, _, _, _ = c, extra, i, offset, prev, q2, v2, v3
	/* Decode finer resolution */
	i = start
	for {
		if !(i < end) {
			break
		}
		extra = int16(*(*int32)(unsafe.Pointer(extra_quant + uintptr(i)*4)))
		if *(*int32)(unsafe.Pointer(extra_quant + uintptr(i)*4)) <= 0 {
			goto _1
		}
		v2 = dec
		v3 = (*OpusT_ec_ctx)(unsafe.Pointer(v2)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v2)).Frng))
		goto _4
	_4:
		if v3+C**(*int32)(unsafe.Pointer(extra_quant + uintptr(i)*4)) > int32((*OpusT_ec_dec)(unsafe.Pointer(dec)).Fstorage)*int32(8) {
			goto _1
		}
		if prev_quant != uintptr(uint32(0)) {
			v3 = *(*int32)(unsafe.Pointer(prev_quant + uintptr(i)*4))
		} else {
			v3 = 0
		}
		prev = int16(v3)
		c = 0
		for {
			q2 = int32(Opus_ec_dec_bits(tls, dec, uint32(uint16(extra))))
			offset = float32(float32((float32(q2)+float32(0.5))*float32(int32(1)<<(int32(14)-int32(extra))))*(float32(1)/float32(16384))) - float32(0.5)
			offset = offset * OpusT_celt_glog(float32(int32(1)<<(int32(14)-int32(prev)))*(float32(1)/float32(16384)))
			*(*OpusT_celt_glog)(unsafe.Pointer(oldEBands + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4)) += offset
			goto _7
		_7:
			;
			c = c + 1
			v3 = c
			if !(v3 < C) {
				break
			}
		}
		goto _1
	_1:
		;
		i = i + 1
	}
}

func Opus_unquant_energy_finalise(tls *libc.TLS, m uintptr, start int32, end int32, oldEBands uintptr, fine_quant uintptr, fine_priority uintptr, bits_left int32, dec uintptr, C int32) {
	var c, i, prio, q2, v3 int32
	var offset OpusT_celt_glog
	_, _, _, _, _, _ = c, i, offset, prio, q2, v3
	/* Use up the remaining bits */
	prio = 0
	for {
		if !(prio < int32(2)) {
			break
		}
		i = start
		for {
			if !(i < end && bits_left >= C) {
				break
			}
			if *(*int32)(unsafe.Pointer(fine_quant + uintptr(i)*4)) >= int32(MAX_FINE_BITS) || *(*int32)(unsafe.Pointer(fine_priority + uintptr(i)*4)) != prio {
				goto _2
			}
			c = 0
			for {
				q2 = int32(Opus_ec_dec_bits(tls, dec, uint32(1)))
				offset = OpusT_celt_glog(float32((float32(q2)-float32(0.5))*float32(int32(1)<<(int32(14)-*(*int32)(unsafe.Pointer(fine_quant + uintptr(i)*4))-int32(1)))) * (float32(1) / float32(16384)))
				if oldEBands != uintptr(uint32(0)) {
					*(*OpusT_celt_glog)(unsafe.Pointer(oldEBands + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4)) += offset
				}
				bits_left = bits_left - 1
				goto _4
			_4:
				;
				c = c + 1
				v3 = c
				if !(v3 < C) {
					break
				}
			}
			goto _2
		_2:
			;
			i = i + 1
		}
		goto _1
	_1:
		;
		prio = prio + 1
	}
}

func Opus_amp2Log2(tls *libc.TLS, m uintptr, effEnd int32, end int32, bandE uintptr, bandLogE uintptr, C int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var c, i, v1 int32
	var integer, range_idx OpusT_opus_int32
	var v4 float32
	var _ /* in at bp+0 */ struct {
		Fi [0]OpusT_opus_uint32
		Ff float32
	}
	_, _, _, _, _, _ = c, i, integer, range_idx, v1, v4
	c = 0
	for {
		i = 0
		for {
			if !(i < effEnd) {
				break
			}
			*(*float32)(unsafe.Pointer(bp)) = *(*OpusT_celt_ener)(unsafe.Pointer(bandE + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4))
			integer = int32(*(*OpusT_opus_uint32)(unsafe.Pointer(bp))>>int32(23)) - int32(127)
			*(*OpusT_opus_uint32)(unsafe.Pointer(bp)) = uint32(int32(*(*OpusT_opus_uint32)(unsafe.Pointer(bp))) - int32(uint32(integer)<<int32(23)))
			range_idx = int32(*(*OpusT_opus_uint32)(unsafe.Pointer(bp)) >> int32(20) & uint32(0x7))
			*(*float32)(unsafe.Pointer(bp)) = float32(*(*float32)(unsafe.Pointer(bp))*log2_x_norm_coeff10[range_idx]) - float32(1.0625)
			*(*float32)(unsafe.Pointer(bp)) = float32(0.08746284246444702) + float32(*(*float32)(unsafe.Pointer(bp))*(float32(1.3578295707702637)+float32(*(*float32)(unsafe.Pointer(bp))*(-float32(0.63897705078125)+float32(*(*float32)(unsafe.Pointer(bp))*(float32(0.4019712507724762)+float32(*(*float32)(unsafe.Pointer(bp))*-float32(0.2841544449329376))))))))
			v4 = float32(integer) + *(*float32)(unsafe.Pointer(bp)) + log2_y_norm_coeff10[range_idx]
			goto _5
		_5:
			*(*OpusT_celt_glog)(unsafe.Pointer(bandLogE + uintptr(i+c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands)*4)) = v4 - Opus_eMeans[i]
			goto _3
		_3:
			;
			i = i + 1
		}
		i = effEnd
		for {
			if !(i < end) {
				break
			}
			*(*OpusT_celt_glog)(unsafe.Pointer(bandLogE + uintptr(c*(*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands+i)*4)) = -float32(14)
			goto _6
		_6:
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

const BITALLOC_SIZE = 11
const TOTAL_MODES = 1

var trim_icdf12 = [11]uint8{
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
var spread_icdf12 = [4]uint8{
	0: uint8(25),
	1: uint8(23),
	2: uint8(2),
}
var tapset_icdf12 = [3]uint8{
	0: uint8(2),
	1: uint8(1),
}

var log2_x_norm_coeff11 = [8]float32{
	0: float32(1),
	1: float32(0.8888888955116272),
	2: float32(0.8),
	3: float32(0.7272727489471436),
	4: float32(0.6666666865348816),
	5: float32(0.6153846383094788),
	6: float32(0.5714285969734192),
	7: float32(0.5333333611488342),
}
var log2_y_norm_coeff11 = [8]float32{
	1: float32(0.1699250042438507),
	2: float32(0.32192808389663696),
	3: float32(0.45943161845207214),
	4: float32(0.5849624872207642),
	5: float32(0.7004396915435791),
	6: float32(0.8073549270629883),
	7: float32(0.9068905711174011),
}

/* Copyright (c) 2010 Xiph.Org Foundation
 * Copyright (c) 2013 Parrot */
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

var eband5ms = [22]OpusT_opus_int16{
	1:  int16(1),
	2:  int16(2),
	3:  int16(3),
	4:  int16(4),
	5:  int16(5),
	6:  int16(6),
	7:  int16(7),
	8:  int16(8),
	9:  int16(10),
	10: int16(12),
	11: int16(14),
	12: int16(16),
	13: int16(20),
	14: int16(24),
	15: int16(28),
	16: int16(34),
	17: int16(40),
	18: int16(48),
	19: int16(60),
	20: int16(78),
	21: int16(100),
}

// C documentation
//
//	/* Alternate tuning (partially derived from Vorbis) */
//	/* Bit allocation table in units of 1/32 bit/sample (0.1875 dB SNR) */
var band_allocation = [231]uint8{
	21:  uint8(90),
	22:  uint8(80),
	23:  uint8(75),
	24:  uint8(69),
	25:  uint8(63),
	26:  uint8(56),
	27:  uint8(49),
	28:  uint8(40),
	29:  uint8(34),
	30:  uint8(29),
	31:  uint8(20),
	32:  uint8(18),
	33:  uint8(10),
	42:  uint8(110),
	43:  uint8(100),
	44:  uint8(90),
	45:  uint8(84),
	46:  uint8(78),
	47:  uint8(71),
	48:  uint8(65),
	49:  uint8(58),
	50:  uint8(51),
	51:  uint8(45),
	52:  uint8(39),
	53:  uint8(32),
	54:  uint8(26),
	55:  uint8(20),
	56:  uint8(12),
	63:  uint8(118),
	64:  uint8(110),
	65:  uint8(103),
	66:  uint8(93),
	67:  uint8(86),
	68:  uint8(80),
	69:  uint8(75),
	70:  uint8(70),
	71:  uint8(65),
	72:  uint8(59),
	73:  uint8(53),
	74:  uint8(47),
	75:  uint8(40),
	76:  uint8(31),
	77:  uint8(23),
	78:  uint8(15),
	79:  uint8(4),
	84:  uint8(126),
	85:  uint8(119),
	86:  uint8(112),
	87:  uint8(104),
	88:  uint8(95),
	89:  uint8(89),
	90:  uint8(83),
	91:  uint8(78),
	92:  uint8(72),
	93:  uint8(66),
	94:  uint8(60),
	95:  uint8(54),
	96:  uint8(47),
	97:  uint8(39),
	98:  uint8(32),
	99:  uint8(25),
	100: uint8(17),
	101: uint8(12),
	102: uint8(1),
	105: uint8(134),
	106: uint8(127),
	107: uint8(120),
	108: uint8(114),
	109: uint8(103),
	110: uint8(97),
	111: uint8(91),
	112: uint8(85),
	113: uint8(78),
	114: uint8(72),
	115: uint8(66),
	116: uint8(60),
	117: uint8(54),
	118: uint8(47),
	119: uint8(41),
	120: uint8(35),
	121: uint8(29),
	122: uint8(23),
	123: uint8(16),
	124: uint8(10),
	125: uint8(1),
	126: uint8(144),
	127: uint8(137),
	128: uint8(130),
	129: uint8(124),
	130: uint8(113),
	131: uint8(107),
	132: uint8(101),
	133: uint8(95),
	134: uint8(88),
	135: uint8(82),
	136: uint8(76),
	137: uint8(70),
	138: uint8(64),
	139: uint8(57),
	140: uint8(51),
	141: uint8(45),
	142: uint8(39),
	143: uint8(33),
	144: uint8(26),
	145: uint8(15),
	146: uint8(1),
	147: uint8(152),
	148: uint8(145),
	149: uint8(138),
	150: uint8(132),
	151: uint8(123),
	152: uint8(117),
	153: uint8(111),
	154: uint8(105),
	155: uint8(98),
	156: uint8(92),
	157: uint8(86),
	158: uint8(80),
	159: uint8(74),
	160: uint8(67),
	161: uint8(61),
	162: uint8(55),
	163: uint8(49),
	164: uint8(43),
	165: uint8(36),
	166: uint8(20),
	167: uint8(1),
	168: uint8(162),
	169: uint8(155),
	170: uint8(148),
	171: uint8(142),
	172: uint8(133),
	173: uint8(127),
	174: uint8(121),
	175: uint8(115),
	176: uint8(108),
	177: uint8(102),
	178: uint8(96),
	179: uint8(90),
	180: uint8(84),
	181: uint8(77),
	182: uint8(71),
	183: uint8(65),
	184: uint8(59),
	185: uint8(53),
	186: uint8(46),
	187: uint8(30),
	188: uint8(1),
	189: uint8(172),
	190: uint8(165),
	191: uint8(158),
	192: uint8(152),
	193: uint8(143),
	194: uint8(137),
	195: uint8(131),
	196: uint8(125),
	197: uint8(118),
	198: uint8(112),
	199: uint8(106),
	200: uint8(100),
	201: uint8(94),
	202: uint8(87),
	203: uint8(81),
	204: uint8(75),
	205: uint8(69),
	206: uint8(63),
	207: uint8(56),
	208: uint8(45),
	209: uint8(20),
	210: uint8(200),
	211: uint8(200),
	212: uint8(200),
	213: uint8(200),
	214: uint8(200),
	215: uint8(200),
	216: uint8(200),
	217: uint8(200),
	218: uint8(198),
	219: uint8(193),
	220: uint8(188),
	221: uint8(183),
	222: uint8(178),
	223: uint8(173),
	224: uint8(168),
	225: uint8(163),
	226: uint8(158),
	227: uint8(153),
	228: uint8(148),
	229: uint8(129),
	230: uint8(104),
}
var window120 = [120]OpusT_celt_coef{
	0:   float32(6.7286966e-05),
	1:   float32(0.00060551348),
	2:   float32(0.001681597),
	3:   float32(0.0032947962),
	4:   float32(0.0054439943),
	5:   float32(0.0081276923),
	6:   float32(0.011344001),
	7:   float32(0.015090633),
	8:   float32(0.019364886),
	9:   float32(0.024163635),
	10:  float32(0.029483315),
	11:  float32(0.035319905),
	12:  float32(0.041668911),
	13:  float32(0.048525347),
	14:  float32(0.055883718),
	15:  float32(0.063737999),
	16:  float32(0.072081616),
	17:  float32(0.080907428),
	18:  float32(0.090207705),
	19:  float32(0.099974111),
	20:  float32(0.11019769),
	21:  float32(0.12086883),
	22:  float32(0.13197729),
	23:  float32(0.14351214),
	24:  float32(0.15546177),
	25:  float32(0.16781389),
	26:  float32(0.1805555),
	27:  float32(0.1936729),
	28:  float32(0.20715171),
	29:  float32(0.22097682),
	30:  float32(0.23513243),
	31:  float32(0.24960208),
	32:  float32(0.2643686),
	33:  float32(0.27941419),
	34:  float32(0.2947204),
	35:  float32(0.31026818),
	36:  float32(0.32603788),
	37:  float32(0.34200931),
	38:  float32(0.35816177),
	39:  float32(0.37447407),
	40:  float32(0.39092462),
	41:  float32(0.40749142),
	42:  float32(0.42415215),
	43:  float32(0.44088423),
	44:  float32(0.45766484),
	45:  float32(0.47447104),
	46:  float32(0.49127978),
	47:  float32(0.50806798),
	48:  float32(0.52481261),
	49:  float32(0.54149077),
	50:  float32(0.55807973),
	51:  float32(0.57455701),
	52:  float32(0.59090049),
	53:  float32(0.60708841),
	54:  float32(0.62309951),
	55:  float32(0.63891306),
	56:  float32(0.65450896),
	57:  float32(0.66986776),
	58:  float32(0.68497077),
	59:  float32(0.6998001),
	60:  float32(0.71433873),
	61:  float32(0.72857055),
	62:  float32(0.74248043),
	63:  float32(0.75605425),
	64:  float32(0.76927895),
	65:  float32(0.78214257),
	66:  float32(0.7946343),
	67:  float32(0.80674445),
	68:  float32(0.81846456),
	69:  float32(0.82978733),
	70:  float32(0.84070669),
	71:  float32(0.85121779),
	72:  float32(0.86131698),
	73:  float32(0.87100183),
	74:  float32(0.88027111),
	75:  float32(0.88912479),
	76:  float32(0.89756398),
	77:  float32(0.90559094),
	78:  float32(0.91320904),
	79:  float32(0.9204227),
	80:  float32(0.92723738),
	81:  float32(0.93365955),
	82:  float32(0.93969656),
	83:  float32(0.94535671),
	84:  float32(0.95064907),
	85:  float32(0.95558353),
	86:  float32(0.96017067),
	87:  float32(0.96442171),
	88:  float32(0.96834849),
	89:  float32(0.97196334),
	90:  float32(0.97527906),
	91:  float32(0.97830883),
	92:  float32(0.98106616),
	93:  float32(0.9835648),
	94:  float32(0.98581869),
	95:  float32(0.98784191),
	96:  float32(0.98964856),
	97:  float32(0.99125274),
	98:  float32(0.99266849),
	99:  float32(0.99390969),
	100: float32(0.99499004),
	101: float32(0.99592297),
	102: float32(0.99672162),
	103: float32(0.99739874),
	104: float32(0.99796667),
	105: float32(0.99843728),
	106: float32(0.99882195),
	107: float32(0.99913147),
	108: float32(0.99937606),
	109: float32(0.99956527),
	110: float32(0.99970802),
	111: float32(0.99981248),
	112: float32(0.99988613),
	113: float32(0.99993565),
	114: float32(0.99996697),
	115: float32(0.99998518),
	116: float32(0.99999457),
	117: float32(0.99999859),
	118: float32(0.99999982),
	119: float32(1),
}
var logN400 = [21]OpusT_opus_int16{
	8:  int16(8),
	9:  int16(8),
	10: int16(8),
	11: int16(8),
	12: int16(16),
	13: int16(16),
	14: int16(16),
	15: int16(21),
	16: int16(21),
	17: int16(24),
	18: int16(29),
	19: int16(34),
	20: int16(36),
}
var cache_index50 = [105]OpusT_opus_int16{
	0:   int16(-int32(1)),
	1:   int16(-int32(1)),
	2:   int16(-int32(1)),
	3:   int16(-int32(1)),
	4:   int16(-int32(1)),
	5:   int16(-int32(1)),
	6:   int16(-int32(1)),
	7:   int16(-int32(1)),
	12:  int16(41),
	13:  int16(41),
	14:  int16(41),
	15:  int16(82),
	16:  int16(82),
	17:  int16(123),
	18:  int16(164),
	19:  int16(200),
	20:  int16(222),
	29:  int16(41),
	30:  int16(41),
	31:  int16(41),
	32:  int16(41),
	33:  int16(123),
	34:  int16(123),
	35:  int16(123),
	36:  int16(164),
	37:  int16(164),
	38:  int16(240),
	39:  int16(266),
	40:  int16(283),
	41:  int16(295),
	42:  int16(41),
	43:  int16(41),
	44:  int16(41),
	45:  int16(41),
	46:  int16(41),
	47:  int16(41),
	48:  int16(41),
	49:  int16(41),
	50:  int16(123),
	51:  int16(123),
	52:  int16(123),
	53:  int16(123),
	54:  int16(240),
	55:  int16(240),
	56:  int16(240),
	57:  int16(266),
	58:  int16(266),
	59:  int16(305),
	60:  int16(318),
	61:  int16(328),
	62:  int16(336),
	63:  int16(123),
	64:  int16(123),
	65:  int16(123),
	66:  int16(123),
	67:  int16(123),
	68:  int16(123),
	69:  int16(123),
	70:  int16(123),
	71:  int16(240),
	72:  int16(240),
	73:  int16(240),
	74:  int16(240),
	75:  int16(305),
	76:  int16(305),
	77:  int16(305),
	78:  int16(318),
	79:  int16(318),
	80:  int16(343),
	81:  int16(351),
	82:  int16(358),
	83:  int16(364),
	84:  int16(240),
	85:  int16(240),
	86:  int16(240),
	87:  int16(240),
	88:  int16(240),
	89:  int16(240),
	90:  int16(240),
	91:  int16(240),
	92:  int16(305),
	93:  int16(305),
	94:  int16(305),
	95:  int16(305),
	96:  int16(343),
	97:  int16(343),
	98:  int16(343),
	99:  int16(351),
	100: int16(351),
	101: int16(370),
	102: int16(376),
	103: int16(382),
	104: int16(387),
}
var cache_bits50 = [392]uint8{
	0:   uint8(40),
	1:   uint8(7),
	2:   uint8(7),
	3:   uint8(7),
	4:   uint8(7),
	5:   uint8(7),
	6:   uint8(7),
	7:   uint8(7),
	8:   uint8(7),
	9:   uint8(7),
	10:  uint8(7),
	11:  uint8(7),
	12:  uint8(7),
	13:  uint8(7),
	14:  uint8(7),
	15:  uint8(7),
	16:  uint8(7),
	17:  uint8(7),
	18:  uint8(7),
	19:  uint8(7),
	20:  uint8(7),
	21:  uint8(7),
	22:  uint8(7),
	23:  uint8(7),
	24:  uint8(7),
	25:  uint8(7),
	26:  uint8(7),
	27:  uint8(7),
	28:  uint8(7),
	29:  uint8(7),
	30:  uint8(7),
	31:  uint8(7),
	32:  uint8(7),
	33:  uint8(7),
	34:  uint8(7),
	35:  uint8(7),
	36:  uint8(7),
	37:  uint8(7),
	38:  uint8(7),
	39:  uint8(7),
	40:  uint8(7),
	41:  uint8(40),
	42:  uint8(15),
	43:  uint8(23),
	44:  uint8(28),
	45:  uint8(31),
	46:  uint8(34),
	47:  uint8(36),
	48:  uint8(38),
	49:  uint8(39),
	50:  uint8(41),
	51:  uint8(42),
	52:  uint8(43),
	53:  uint8(44),
	54:  uint8(45),
	55:  uint8(46),
	56:  uint8(47),
	57:  uint8(47),
	58:  uint8(49),
	59:  uint8(50),
	60:  uint8(51),
	61:  uint8(52),
	62:  uint8(53),
	63:  uint8(54),
	64:  uint8(55),
	65:  uint8(55),
	66:  uint8(57),
	67:  uint8(58),
	68:  uint8(59),
	69:  uint8(60),
	70:  uint8(61),
	71:  uint8(62),
	72:  uint8(63),
	73:  uint8(63),
	74:  uint8(65),
	75:  uint8(66),
	76:  uint8(67),
	77:  uint8(68),
	78:  uint8(69),
	79:  uint8(70),
	80:  uint8(71),
	81:  uint8(71),
	82:  uint8(40),
	83:  uint8(20),
	84:  uint8(33),
	85:  uint8(41),
	86:  uint8(48),
	87:  uint8(53),
	88:  uint8(57),
	89:  uint8(61),
	90:  uint8(64),
	91:  uint8(66),
	92:  uint8(69),
	93:  uint8(71),
	94:  uint8(73),
	95:  uint8(75),
	96:  uint8(76),
	97:  uint8(78),
	98:  uint8(80),
	99:  uint8(82),
	100: uint8(85),
	101: uint8(87),
	102: uint8(89),
	103: uint8(91),
	104: uint8(92),
	105: uint8(94),
	106: uint8(96),
	107: uint8(98),
	108: uint8(101),
	109: uint8(103),
	110: uint8(105),
	111: uint8(107),
	112: uint8(108),
	113: uint8(110),
	114: uint8(112),
	115: uint8(114),
	116: uint8(117),
	117: uint8(119),
	118: uint8(121),
	119: uint8(123),
	120: uint8(124),
	121: uint8(126),
	122: uint8(128),
	123: uint8(40),
	124: uint8(23),
	125: uint8(39),
	126: uint8(51),
	127: uint8(60),
	128: uint8(67),
	129: uint8(73),
	130: uint8(79),
	131: uint8(83),
	132: uint8(87),
	133: uint8(91),
	134: uint8(94),
	135: uint8(97),
	136: uint8(100),
	137: uint8(102),
	138: uint8(105),
	139: uint8(107),
	140: uint8(111),
	141: uint8(115),
	142: uint8(118),
	143: uint8(121),
	144: uint8(124),
	145: uint8(126),
	146: uint8(129),
	147: uint8(131),
	148: uint8(135),
	149: uint8(139),
	150: uint8(142),
	151: uint8(145),
	152: uint8(148),
	153: uint8(150),
	154: uint8(153),
	155: uint8(155),
	156: uint8(159),
	157: uint8(163),
	158: uint8(166),
	159: uint8(169),
	160: uint8(172),
	161: uint8(174),
	162: uint8(177),
	163: uint8(179),
	164: uint8(35),
	165: uint8(28),
	166: uint8(49),
	167: uint8(65),
	168: uint8(78),
	169: uint8(89),
	170: uint8(99),
	171: uint8(107),
	172: uint8(114),
	173: uint8(120),
	174: uint8(126),
	175: uint8(132),
	176: uint8(136),
	177: uint8(141),
	178: uint8(145),
	179: uint8(149),
	180: uint8(153),
	181: uint8(159),
	182: uint8(165),
	183: uint8(171),
	184: uint8(176),
	185: uint8(180),
	186: uint8(185),
	187: uint8(189),
	188: uint8(192),
	189: uint8(199),
	190: uint8(205),
	191: uint8(211),
	192: uint8(216),
	193: uint8(220),
	194: uint8(225),
	195: uint8(229),
	196: uint8(232),
	197: uint8(239),
	198: uint8(245),
	199: uint8(251),
	200: uint8(21),
	201: uint8(33),
	202: uint8(58),
	203: uint8(79),
	204: uint8(97),
	205: uint8(112),
	206: uint8(125),
	207: uint8(137),
	208: uint8(148),
	209: uint8(157),
	210: uint8(166),
	211: uint8(174),
	212: uint8(182),
	213: uint8(189),
	214: uint8(195),
	215: uint8(201),
	216: uint8(207),
	217: uint8(217),
	218: uint8(227),
	219: uint8(235),
	220: uint8(243),
	221: uint8(251),
	222: uint8(17),
	223: uint8(35),
	224: uint8(63),
	225: uint8(86),
	226: uint8(106),
	227: uint8(123),
	228: uint8(139),
	229: uint8(152),
	230: uint8(165),
	231: uint8(177),
	232: uint8(187),
	233: uint8(197),
	234: uint8(206),
	235: uint8(214),
	236: uint8(222),
	237: uint8(230),
	238: uint8(237),
	239: uint8(250),
	240: uint8(25),
	241: uint8(31),
	242: uint8(55),
	243: uint8(75),
	244: uint8(91),
	245: uint8(105),
	246: uint8(117),
	247: uint8(128),
	248: uint8(138),
	249: uint8(146),
	250: uint8(154),
	251: uint8(161),
	252: uint8(168),
	253: uint8(174),
	254: uint8(180),
	255: uint8(185),
	256: uint8(190),
	257: uint8(200),
	258: uint8(208),
	259: uint8(215),
	260: uint8(222),
	261: uint8(229),
	262: uint8(235),
	263: uint8(240),
	264: uint8(245),
	265: uint8(255),
	266: uint8(16),
	267: uint8(36),
	268: uint8(65),
	269: uint8(89),
	270: uint8(110),
	271: uint8(128),
	272: uint8(144),
	273: uint8(159),
	274: uint8(173),
	275: uint8(185),
	276: uint8(196),
	277: uint8(207),
	278: uint8(217),
	279: uint8(226),
	280: uint8(234),
	281: uint8(242),
	282: uint8(250),
	283: uint8(11),
	284: uint8(41),
	285: uint8(74),
	286: uint8(103),
	287: uint8(128),
	288: uint8(151),
	289: uint8(172),
	290: uint8(191),
	291: uint8(209),
	292: uint8(225),
	293: uint8(241),
	294: uint8(255),
	295: uint8(9),
	296: uint8(43),
	297: uint8(79),
	298: uint8(110),
	299: uint8(138),
	300: uint8(163),
	301: uint8(186),
	302: uint8(207),
	303: uint8(227),
	304: uint8(246),
	305: uint8(12),
	306: uint8(39),
	307: uint8(71),
	308: uint8(99),
	309: uint8(123),
	310: uint8(144),
	311: uint8(164),
	312: uint8(182),
	313: uint8(198),
	314: uint8(214),
	315: uint8(228),
	316: uint8(241),
	317: uint8(253),
	318: uint8(9),
	319: uint8(44),
	320: uint8(81),
	321: uint8(113),
	322: uint8(142),
	323: uint8(168),
	324: uint8(192),
	325: uint8(214),
	326: uint8(235),
	327: uint8(255),
	328: uint8(7),
	329: uint8(49),
	330: uint8(90),
	331: uint8(127),
	332: uint8(160),
	333: uint8(191),
	334: uint8(220),
	335: uint8(247),
	336: uint8(6),
	337: uint8(51),
	338: uint8(95),
	339: uint8(134),
	340: uint8(170),
	341: uint8(203),
	342: uint8(234),
	343: uint8(7),
	344: uint8(47),
	345: uint8(87),
	346: uint8(123),
	347: uint8(155),
	348: uint8(184),
	349: uint8(212),
	350: uint8(237),
	351: uint8(6),
	352: uint8(52),
	353: uint8(97),
	354: uint8(137),
	355: uint8(174),
	356: uint8(208),
	357: uint8(240),
	358: uint8(5),
	359: uint8(57),
	360: uint8(106),
	361: uint8(151),
	362: uint8(192),
	363: uint8(231),
	364: uint8(5),
	365: uint8(59),
	366: uint8(111),
	367: uint8(158),
	368: uint8(202),
	369: uint8(243),
	370: uint8(5),
	371: uint8(55),
	372: uint8(103),
	373: uint8(147),
	374: uint8(187),
	375: uint8(224),
	376: uint8(5),
	377: uint8(60),
	378: uint8(113),
	379: uint8(161),
	380: uint8(206),
	381: uint8(248),
	382: uint8(4),
	383: uint8(65),
	384: uint8(122),
	385: uint8(175),
	386: uint8(224),
	387: uint8(4),
	388: uint8(67),
	389: uint8(127),
	390: uint8(182),
	391: uint8(234),
}
var cache_caps50 = [168]uint8{
	0:   uint8(224),
	1:   uint8(224),
	2:   uint8(224),
	3:   uint8(224),
	4:   uint8(224),
	5:   uint8(224),
	6:   uint8(224),
	7:   uint8(224),
	8:   uint8(160),
	9:   uint8(160),
	10:  uint8(160),
	11:  uint8(160),
	12:  uint8(185),
	13:  uint8(185),
	14:  uint8(185),
	15:  uint8(178),
	16:  uint8(178),
	17:  uint8(168),
	18:  uint8(134),
	19:  uint8(61),
	20:  uint8(37),
	21:  uint8(224),
	22:  uint8(224),
	23:  uint8(224),
	24:  uint8(224),
	25:  uint8(224),
	26:  uint8(224),
	27:  uint8(224),
	28:  uint8(224),
	29:  uint8(240),
	30:  uint8(240),
	31:  uint8(240),
	32:  uint8(240),
	33:  uint8(207),
	34:  uint8(207),
	35:  uint8(207),
	36:  uint8(198),
	37:  uint8(198),
	38:  uint8(183),
	39:  uint8(144),
	40:  uint8(66),
	41:  uint8(40),
	42:  uint8(160),
	43:  uint8(160),
	44:  uint8(160),
	45:  uint8(160),
	46:  uint8(160),
	47:  uint8(160),
	48:  uint8(160),
	49:  uint8(160),
	50:  uint8(185),
	51:  uint8(185),
	52:  uint8(185),
	53:  uint8(185),
	54:  uint8(193),
	55:  uint8(193),
	56:  uint8(193),
	57:  uint8(183),
	58:  uint8(183),
	59:  uint8(172),
	60:  uint8(138),
	61:  uint8(64),
	62:  uint8(38),
	63:  uint8(240),
	64:  uint8(240),
	65:  uint8(240),
	66:  uint8(240),
	67:  uint8(240),
	68:  uint8(240),
	69:  uint8(240),
	70:  uint8(240),
	71:  uint8(207),
	72:  uint8(207),
	73:  uint8(207),
	74:  uint8(207),
	75:  uint8(204),
	76:  uint8(204),
	77:  uint8(204),
	78:  uint8(193),
	79:  uint8(193),
	80:  uint8(180),
	81:  uint8(143),
	82:  uint8(66),
	83:  uint8(40),
	84:  uint8(185),
	85:  uint8(185),
	86:  uint8(185),
	87:  uint8(185),
	88:  uint8(185),
	89:  uint8(185),
	90:  uint8(185),
	91:  uint8(185),
	92:  uint8(193),
	93:  uint8(193),
	94:  uint8(193),
	95:  uint8(193),
	96:  uint8(193),
	97:  uint8(193),
	98:  uint8(193),
	99:  uint8(183),
	100: uint8(183),
	101: uint8(172),
	102: uint8(138),
	103: uint8(65),
	104: uint8(39),
	105: uint8(207),
	106: uint8(207),
	107: uint8(207),
	108: uint8(207),
	109: uint8(207),
	110: uint8(207),
	111: uint8(207),
	112: uint8(207),
	113: uint8(204),
	114: uint8(204),
	115: uint8(204),
	116: uint8(204),
	117: uint8(201),
	118: uint8(201),
	119: uint8(201),
	120: uint8(188),
	121: uint8(188),
	122: uint8(176),
	123: uint8(141),
	124: uint8(66),
	125: uint8(40),
	126: uint8(193),
	127: uint8(193),
	128: uint8(193),
	129: uint8(193),
	130: uint8(193),
	131: uint8(193),
	132: uint8(193),
	133: uint8(193),
	134: uint8(193),
	135: uint8(193),
	136: uint8(193),
	137: uint8(193),
	138: uint8(194),
	139: uint8(194),
	140: uint8(194),
	141: uint8(184),
	142: uint8(184),
	143: uint8(173),
	144: uint8(139),
	145: uint8(65),
	146: uint8(39),
	147: uint8(204),
	148: uint8(204),
	149: uint8(204),
	150: uint8(204),
	151: uint8(204),
	152: uint8(204),
	153: uint8(204),
	154: uint8(204),
	155: uint8(201),
	156: uint8(201),
	157: uint8(201),
	158: uint8(201),
	159: uint8(198),
	160: uint8(198),
	161: uint8(198),
	162: uint8(187),
	163: uint8(187),
	164: uint8(175),
	165: uint8(140),
	166: uint8(66),
	167: uint8(40),
}
var fft_twiddles48000_960 = [480]OpusT_kiss_twiddle_cpx{
	0: {
		Fr: float32(1),
		Fi: -float32(0),
	},
	1: {
		Fr: float32(0.99991433),
		Fi: -float32(0.013089596),
	},
	2: {
		Fr: float32(0.99965732),
		Fi: -float32(0.026176948),
	},
	3: {
		Fr: float32(0.99922904),
		Fi: -float32(0.039259816),
	},
	4: {
		Fr: float32(0.99862953),
		Fi: -float32(0.052335956),
	},
	5: {
		Fr: float32(0.99785892),
		Fi: -float32(0.065403129),
	},
	6: {
		Fr: float32(0.99691733),
		Fi: -float32(0.078459096),
	},
	7: {
		Fr: float32(0.99580493),
		Fi: -float32(0.091501619),
	},
	8: {
		Fr: float32(0.9945219),
		Fi: -float32(0.10452846),
	},
	9: {
		Fr: float32(0.99306846),
		Fi: -float32(0.1175374),
	},
	10: {
		Fr: float32(0.99144486),
		Fi: -float32(0.13052619),
	},
	11: {
		Fr: float32(0.98965139),
		Fi: -float32(0.14349262),
	},
	12: {
		Fr: float32(0.98768834),
		Fi: -float32(0.15643447),
	},
	13: {
		Fr: float32(0.98555606),
		Fi: -float32(0.1693495),
	},
	14: {
		Fr: float32(0.98325491),
		Fi: -float32(0.18223553),
	},
	15: {
		Fr: float32(0.98078528),
		Fi: -float32(0.19509032),
	},
	16: {
		Fr: float32(0.9781476),
		Fi: -float32(0.20791169),
	},
	17: {
		Fr: float32(0.97534232),
		Fi: -float32(0.22069744),
	},
	18: {
		Fr: float32(0.97236992),
		Fi: -float32(0.23344536),
	},
	19: {
		Fr: float32(0.96923091),
		Fi: -float32(0.24615329),
	},
	20: {
		Fr: float32(0.96592583),
		Fi: -float32(0.25881905),
	},
	21: {
		Fr: float32(0.96245524),
		Fi: -float32(0.27144045),
	},
	22: {
		Fr: float32(0.95881973),
		Fi: -float32(0.28401534),
	},
	23: {
		Fr: float32(0.95501994),
		Fi: -float32(0.29654157),
	},
	24: {
		Fr: float32(0.95105652),
		Fi: -float32(0.30901699),
	},
	25: {
		Fr: float32(0.94693013),
		Fi: -float32(0.32143947),
	},
	26: {
		Fr: float32(0.94264149),
		Fi: -float32(0.33380686),
	},
	27: {
		Fr: float32(0.93819134),
		Fi: -float32(0.34611706),
	},
	28: {
		Fr: float32(0.93358043),
		Fi: -float32(0.35836795),
	},
	29: {
		Fr: float32(0.92880955),
		Fi: -float32(0.37055744),
	},
	30: {
		Fr: float32(0.92387953),
		Fi: -float32(0.38268343),
	},
	31: {
		Fr: float32(0.91879121),
		Fi: -float32(0.39474386),
	},
	32: {
		Fr: float32(0.91354546),
		Fi: -float32(0.40673664),
	},
	33: {
		Fr: float32(0.90814317),
		Fi: -float32(0.41865974),
	},
	34: {
		Fr: float32(0.90258528),
		Fi: -float32(0.4305111),
	},
	35: {
		Fr: float32(0.89687274),
		Fi: -float32(0.44228869),
	},
	36: {
		Fr: float32(0.89100652),
		Fi: -float32(0.4539905),
	},
	37: {
		Fr: float32(0.88498764),
		Fi: -float32(0.46561452),
	},
	38: {
		Fr: float32(0.87881711),
		Fi: -float32(0.47715876),
	},
	39: {
		Fr: float32(0.87249601),
		Fi: -float32(0.48862124),
	},
	40: {
		Fr: float32(0.8660254),
		Fi: -float32(0.5),
	},
	41: {
		Fr: float32(0.85940641),
		Fi: -float32(0.51129309),
	},
	42: {
		Fr: float32(0.85264016),
		Fi: -float32(0.52249856),
	},
	43: {
		Fr: float32(0.84572782),
		Fi: -float32(0.53361452),
	},
	44: {
		Fr: float32(0.83867057),
		Fi: -float32(0.54463904),
	},
	45: {
		Fr: float32(0.83146961),
		Fi: -float32(0.55557023),
	},
	46: {
		Fr: float32(0.82412619),
		Fi: -float32(0.56640624),
	},
	47: {
		Fr: float32(0.81664156),
		Fi: -float32(0.57714519),
	},
	48: {
		Fr: float32(0.80901699),
		Fi: -float32(0.58778525),
	},
	49: {
		Fr: float32(0.80125381),
		Fi: -float32(0.5983246),
	},
	50: {
		Fr: float32(0.79335334),
		Fi: -float32(0.60876143),
	},
	51: {
		Fr: float32(0.78531693),
		Fi: -float32(0.61909395),
	},
	52: {
		Fr: float32(0.77714596),
		Fi: -float32(0.62932039),
	},
	53: {
		Fr: float32(0.76884183),
		Fi: -float32(0.639439),
	},
	54: {
		Fr: float32(0.76040597),
		Fi: -float32(0.64944805),
	},
	55: {
		Fr: float32(0.75183981),
		Fi: -float32(0.65934582),
	},
	56: {
		Fr: float32(0.74314483),
		Fi: -float32(0.66913061),
	},
	57: {
		Fr: float32(0.73432251),
		Fi: -float32(0.67880075),
	},
	58: {
		Fr: float32(0.72537437),
		Fi: -float32(0.68835458),
	},
	59: {
		Fr: float32(0.71630194),
		Fi: -float32(0.69779046),
	},
	60: {
		Fr: float32(0.70710678),
		Fi: -float32(0.70710678),
	},
	61: {
		Fr: float32(0.69779046),
		Fi: -float32(0.71630194),
	},
	62: {
		Fr: float32(0.68835458),
		Fi: -float32(0.72537437),
	},
	63: {
		Fr: float32(0.67880075),
		Fi: -float32(0.73432251),
	},
	64: {
		Fr: float32(0.66913061),
		Fi: -float32(0.74314483),
	},
	65: {
		Fr: float32(0.65934582),
		Fi: -float32(0.75183981),
	},
	66: {
		Fr: float32(0.64944805),
		Fi: -float32(0.76040597),
	},
	67: {
		Fr: float32(0.639439),
		Fi: -float32(0.76884183),
	},
	68: {
		Fr: float32(0.62932039),
		Fi: -float32(0.77714596),
	},
	69: {
		Fr: float32(0.61909395),
		Fi: -float32(0.78531693),
	},
	70: {
		Fr: float32(0.60876143),
		Fi: -float32(0.79335334),
	},
	71: {
		Fr: float32(0.5983246),
		Fi: -float32(0.80125381),
	},
	72: {
		Fr: float32(0.58778525),
		Fi: -float32(0.80901699),
	},
	73: {
		Fr: float32(0.57714519),
		Fi: -float32(0.81664156),
	},
	74: {
		Fr: float32(0.56640624),
		Fi: -float32(0.82412619),
	},
	75: {
		Fr: float32(0.55557023),
		Fi: -float32(0.83146961),
	},
	76: {
		Fr: float32(0.54463904),
		Fi: -float32(0.83867057),
	},
	77: {
		Fr: float32(0.53361452),
		Fi: -float32(0.84572782),
	},
	78: {
		Fr: float32(0.52249856),
		Fi: -float32(0.85264016),
	},
	79: {
		Fr: float32(0.51129309),
		Fi: -float32(0.85940641),
	},
	80: {
		Fr: float32(0.5),
		Fi: -float32(0.8660254),
	},
	81: {
		Fr: float32(0.48862124),
		Fi: -float32(0.87249601),
	},
	82: {
		Fr: float32(0.47715876),
		Fi: -float32(0.87881711),
	},
	83: {
		Fr: float32(0.46561452),
		Fi: -float32(0.88498764),
	},
	84: {
		Fr: float32(0.4539905),
		Fi: -float32(0.89100652),
	},
	85: {
		Fr: float32(0.44228869),
		Fi: -float32(0.89687274),
	},
	86: {
		Fr: float32(0.4305111),
		Fi: -float32(0.90258528),
	},
	87: {
		Fr: float32(0.41865974),
		Fi: -float32(0.90814317),
	},
	88: {
		Fr: float32(0.40673664),
		Fi: -float32(0.91354546),
	},
	89: {
		Fr: float32(0.39474386),
		Fi: -float32(0.91879121),
	},
	90: {
		Fr: float32(0.38268343),
		Fi: -float32(0.92387953),
	},
	91: {
		Fr: float32(0.37055744),
		Fi: -float32(0.92880955),
	},
	92: {
		Fr: float32(0.35836795),
		Fi: -float32(0.93358043),
	},
	93: {
		Fr: float32(0.34611706),
		Fi: -float32(0.93819134),
	},
	94: {
		Fr: float32(0.33380686),
		Fi: -float32(0.94264149),
	},
	95: {
		Fr: float32(0.32143947),
		Fi: -float32(0.94693013),
	},
	96: {
		Fr: float32(0.30901699),
		Fi: -float32(0.95105652),
	},
	97: {
		Fr: float32(0.29654157),
		Fi: -float32(0.95501994),
	},
	98: {
		Fr: float32(0.28401534),
		Fi: -float32(0.95881973),
	},
	99: {
		Fr: float32(0.27144045),
		Fi: -float32(0.96245524),
	},
	100: {
		Fr: float32(0.25881905),
		Fi: -float32(0.96592583),
	},
	101: {
		Fr: float32(0.24615329),
		Fi: -float32(0.96923091),
	},
	102: {
		Fr: float32(0.23344536),
		Fi: -float32(0.97236992),
	},
	103: {
		Fr: float32(0.22069744),
		Fi: -float32(0.97534232),
	},
	104: {
		Fr: float32(0.20791169),
		Fi: -float32(0.9781476),
	},
	105: {
		Fr: float32(0.19509032),
		Fi: -float32(0.98078528),
	},
	106: {
		Fr: float32(0.18223553),
		Fi: -float32(0.98325491),
	},
	107: {
		Fr: float32(0.1693495),
		Fi: -float32(0.98555606),
	},
	108: {
		Fr: float32(0.15643447),
		Fi: -float32(0.98768834),
	},
	109: {
		Fr: float32(0.14349262),
		Fi: -float32(0.98965139),
	},
	110: {
		Fr: float32(0.13052619),
		Fi: -float32(0.99144486),
	},
	111: {
		Fr: float32(0.1175374),
		Fi: -float32(0.99306846),
	},
	112: {
		Fr: float32(0.10452846),
		Fi: -float32(0.9945219),
	},
	113: {
		Fr: float32(0.091501619),
		Fi: -float32(0.99580493),
	},
	114: {
		Fr: float32(0.078459096),
		Fi: -float32(0.99691733),
	},
	115: {
		Fr: float32(0.065403129),
		Fi: -float32(0.99785892),
	},
	116: {
		Fr: float32(0.052335956),
		Fi: -float32(0.99862953),
	},
	117: {
		Fr: float32(0.039259816),
		Fi: -float32(0.99922904),
	},
	118: {
		Fr: float32(0.026176948),
		Fi: -float32(0.99965732),
	},
	119: {
		Fr: float32(0.013089596),
		Fi: -float32(0.99991433),
	},
	120: {
		Fr: float32(6.123234e-17),
		Fi: -float32(1),
	},
	121: {
		Fr: -float32(0.013089596),
		Fi: -float32(0.99991433),
	},
	122: {
		Fr: -float32(0.026176948),
		Fi: -float32(0.99965732),
	},
	123: {
		Fr: -float32(0.039259816),
		Fi: -float32(0.99922904),
	},
	124: {
		Fr: -float32(0.052335956),
		Fi: -float32(0.99862953),
	},
	125: {
		Fr: -float32(0.065403129),
		Fi: -float32(0.99785892),
	},
	126: {
		Fr: -float32(0.078459096),
		Fi: -float32(0.99691733),
	},
	127: {
		Fr: -float32(0.091501619),
		Fi: -float32(0.99580493),
	},
	128: {
		Fr: -float32(0.10452846),
		Fi: -float32(0.9945219),
	},
	129: {
		Fr: -float32(0.1175374),
		Fi: -float32(0.99306846),
	},
	130: {
		Fr: -float32(0.13052619),
		Fi: -float32(0.99144486),
	},
	131: {
		Fr: -float32(0.14349262),
		Fi: -float32(0.98965139),
	},
	132: {
		Fr: -float32(0.15643447),
		Fi: -float32(0.98768834),
	},
	133: {
		Fr: -float32(0.1693495),
		Fi: -float32(0.98555606),
	},
	134: {
		Fr: -float32(0.18223553),
		Fi: -float32(0.98325491),
	},
	135: {
		Fr: -float32(0.19509032),
		Fi: -float32(0.98078528),
	},
	136: {
		Fr: -float32(0.20791169),
		Fi: -float32(0.9781476),
	},
	137: {
		Fr: -float32(0.22069744),
		Fi: -float32(0.97534232),
	},
	138: {
		Fr: -float32(0.23344536),
		Fi: -float32(0.97236992),
	},
	139: {
		Fr: -float32(0.24615329),
		Fi: -float32(0.96923091),
	},
	140: {
		Fr: -float32(0.25881905),
		Fi: -float32(0.96592583),
	},
	141: {
		Fr: -float32(0.27144045),
		Fi: -float32(0.96245524),
	},
	142: {
		Fr: -float32(0.28401534),
		Fi: -float32(0.95881973),
	},
	143: {
		Fr: -float32(0.29654157),
		Fi: -float32(0.95501994),
	},
	144: {
		Fr: -float32(0.30901699),
		Fi: -float32(0.95105652),
	},
	145: {
		Fr: -float32(0.32143947),
		Fi: -float32(0.94693013),
	},
	146: {
		Fr: -float32(0.33380686),
		Fi: -float32(0.94264149),
	},
	147: {
		Fr: -float32(0.34611706),
		Fi: -float32(0.93819134),
	},
	148: {
		Fr: -float32(0.35836795),
		Fi: -float32(0.93358043),
	},
	149: {
		Fr: -float32(0.37055744),
		Fi: -float32(0.92880955),
	},
	150: {
		Fr: -float32(0.38268343),
		Fi: -float32(0.92387953),
	},
	151: {
		Fr: -float32(0.39474386),
		Fi: -float32(0.91879121),
	},
	152: {
		Fr: -float32(0.40673664),
		Fi: -float32(0.91354546),
	},
	153: {
		Fr: -float32(0.41865974),
		Fi: -float32(0.90814317),
	},
	154: {
		Fr: -float32(0.4305111),
		Fi: -float32(0.90258528),
	},
	155: {
		Fr: -float32(0.44228869),
		Fi: -float32(0.89687274),
	},
	156: {
		Fr: -float32(0.4539905),
		Fi: -float32(0.89100652),
	},
	157: {
		Fr: -float32(0.46561452),
		Fi: -float32(0.88498764),
	},
	158: {
		Fr: -float32(0.47715876),
		Fi: -float32(0.87881711),
	},
	159: {
		Fr: -float32(0.48862124),
		Fi: -float32(0.87249601),
	},
	160: {
		Fr: -float32(0.5),
		Fi: -float32(0.8660254),
	},
	161: {
		Fr: -float32(0.51129309),
		Fi: -float32(0.85940641),
	},
	162: {
		Fr: -float32(0.52249856),
		Fi: -float32(0.85264016),
	},
	163: {
		Fr: -float32(0.53361452),
		Fi: -float32(0.84572782),
	},
	164: {
		Fr: -float32(0.54463904),
		Fi: -float32(0.83867057),
	},
	165: {
		Fr: -float32(0.55557023),
		Fi: -float32(0.83146961),
	},
	166: {
		Fr: -float32(0.56640624),
		Fi: -float32(0.82412619),
	},
	167: {
		Fr: -float32(0.57714519),
		Fi: -float32(0.81664156),
	},
	168: {
		Fr: -float32(0.58778525),
		Fi: -float32(0.80901699),
	},
	169: {
		Fr: -float32(0.5983246),
		Fi: -float32(0.80125381),
	},
	170: {
		Fr: -float32(0.60876143),
		Fi: -float32(0.79335334),
	},
	171: {
		Fr: -float32(0.61909395),
		Fi: -float32(0.78531693),
	},
	172: {
		Fr: -float32(0.62932039),
		Fi: -float32(0.77714596),
	},
	173: {
		Fr: -float32(0.639439),
		Fi: -float32(0.76884183),
	},
	174: {
		Fr: -float32(0.64944805),
		Fi: -float32(0.76040597),
	},
	175: {
		Fr: -float32(0.65934582),
		Fi: -float32(0.75183981),
	},
	176: {
		Fr: -float32(0.66913061),
		Fi: -float32(0.74314483),
	},
	177: {
		Fr: -float32(0.67880075),
		Fi: -float32(0.73432251),
	},
	178: {
		Fr: -float32(0.68835458),
		Fi: -float32(0.72537437),
	},
	179: {
		Fr: -float32(0.69779046),
		Fi: -float32(0.71630194),
	},
	180: {
		Fr: -float32(0.70710678),
		Fi: -float32(0.70710678),
	},
	181: {
		Fr: -float32(0.71630194),
		Fi: -float32(0.69779046),
	},
	182: {
		Fr: -float32(0.72537437),
		Fi: -float32(0.68835458),
	},
	183: {
		Fr: -float32(0.73432251),
		Fi: -float32(0.67880075),
	},
	184: {
		Fr: -float32(0.74314483),
		Fi: -float32(0.66913061),
	},
	185: {
		Fr: -float32(0.75183981),
		Fi: -float32(0.65934582),
	},
	186: {
		Fr: -float32(0.76040597),
		Fi: -float32(0.64944805),
	},
	187: {
		Fr: -float32(0.76884183),
		Fi: -float32(0.639439),
	},
	188: {
		Fr: -float32(0.77714596),
		Fi: -float32(0.62932039),
	},
	189: {
		Fr: -float32(0.78531693),
		Fi: -float32(0.61909395),
	},
	190: {
		Fr: -float32(0.79335334),
		Fi: -float32(0.60876143),
	},
	191: {
		Fr: -float32(0.80125381),
		Fi: -float32(0.5983246),
	},
	192: {
		Fr: -float32(0.80901699),
		Fi: -float32(0.58778525),
	},
	193: {
		Fr: -float32(0.81664156),
		Fi: -float32(0.57714519),
	},
	194: {
		Fr: -float32(0.82412619),
		Fi: -float32(0.56640624),
	},
	195: {
		Fr: -float32(0.83146961),
		Fi: -float32(0.55557023),
	},
	196: {
		Fr: -float32(0.83867057),
		Fi: -float32(0.54463904),
	},
	197: {
		Fr: -float32(0.84572782),
		Fi: -float32(0.53361452),
	},
	198: {
		Fr: -float32(0.85264016),
		Fi: -float32(0.52249856),
	},
	199: {
		Fr: -float32(0.85940641),
		Fi: -float32(0.51129309),
	},
	200: {
		Fr: -float32(0.8660254),
		Fi: -float32(0.5),
	},
	201: {
		Fr: -float32(0.87249601),
		Fi: -float32(0.48862124),
	},
	202: {
		Fr: -float32(0.87881711),
		Fi: -float32(0.47715876),
	},
	203: {
		Fr: -float32(0.88498764),
		Fi: -float32(0.46561452),
	},
	204: {
		Fr: -float32(0.89100652),
		Fi: -float32(0.4539905),
	},
	205: {
		Fr: -float32(0.89687274),
		Fi: -float32(0.44228869),
	},
	206: {
		Fr: -float32(0.90258528),
		Fi: -float32(0.4305111),
	},
	207: {
		Fr: -float32(0.90814317),
		Fi: -float32(0.41865974),
	},
	208: {
		Fr: -float32(0.91354546),
		Fi: -float32(0.40673664),
	},
	209: {
		Fr: -float32(0.91879121),
		Fi: -float32(0.39474386),
	},
	210: {
		Fr: -float32(0.92387953),
		Fi: -float32(0.38268343),
	},
	211: {
		Fr: -float32(0.92880955),
		Fi: -float32(0.37055744),
	},
	212: {
		Fr: -float32(0.93358043),
		Fi: -float32(0.35836795),
	},
	213: {
		Fr: -float32(0.93819134),
		Fi: -float32(0.34611706),
	},
	214: {
		Fr: -float32(0.94264149),
		Fi: -float32(0.33380686),
	},
	215: {
		Fr: -float32(0.94693013),
		Fi: -float32(0.32143947),
	},
	216: {
		Fr: -float32(0.95105652),
		Fi: -float32(0.30901699),
	},
	217: {
		Fr: -float32(0.95501994),
		Fi: -float32(0.29654157),
	},
	218: {
		Fr: -float32(0.95881973),
		Fi: -float32(0.28401534),
	},
	219: {
		Fr: -float32(0.96245524),
		Fi: -float32(0.27144045),
	},
	220: {
		Fr: -float32(0.96592583),
		Fi: -float32(0.25881905),
	},
	221: {
		Fr: -float32(0.96923091),
		Fi: -float32(0.24615329),
	},
	222: {
		Fr: -float32(0.97236992),
		Fi: -float32(0.23344536),
	},
	223: {
		Fr: -float32(0.97534232),
		Fi: -float32(0.22069744),
	},
	224: {
		Fr: -float32(0.9781476),
		Fi: -float32(0.20791169),
	},
	225: {
		Fr: -float32(0.98078528),
		Fi: -float32(0.19509032),
	},
	226: {
		Fr: -float32(0.98325491),
		Fi: -float32(0.18223553),
	},
	227: {
		Fr: -float32(0.98555606),
		Fi: -float32(0.1693495),
	},
	228: {
		Fr: -float32(0.98768834),
		Fi: -float32(0.15643447),
	},
	229: {
		Fr: -float32(0.98965139),
		Fi: -float32(0.14349262),
	},
	230: {
		Fr: -float32(0.99144486),
		Fi: -float32(0.13052619),
	},
	231: {
		Fr: -float32(0.99306846),
		Fi: -float32(0.1175374),
	},
	232: {
		Fr: -float32(0.9945219),
		Fi: -float32(0.10452846),
	},
	233: {
		Fr: -float32(0.99580493),
		Fi: -float32(0.091501619),
	},
	234: {
		Fr: -float32(0.99691733),
		Fi: -float32(0.078459096),
	},
	235: {
		Fr: -float32(0.99785892),
		Fi: -float32(0.065403129),
	},
	236: {
		Fr: -float32(0.99862953),
		Fi: -float32(0.052335956),
	},
	237: {
		Fr: -float32(0.99922904),
		Fi: -float32(0.039259816),
	},
	238: {
		Fr: -float32(0.99965732),
		Fi: -float32(0.026176948),
	},
	239: {
		Fr: -float32(0.99991433),
		Fi: -float32(0.013089596),
	},
	240: {
		Fr: -float32(1),
		Fi: -float32(1.2246468e-16),
	},
	241: {
		Fr: -float32(0.99991433),
		Fi: float32(0.013089596),
	},
	242: {
		Fr: -float32(0.99965732),
		Fi: float32(0.026176948),
	},
	243: {
		Fr: -float32(0.99922904),
		Fi: float32(0.039259816),
	},
	244: {
		Fr: -float32(0.99862953),
		Fi: float32(0.052335956),
	},
	245: {
		Fr: -float32(0.99785892),
		Fi: float32(0.065403129),
	},
	246: {
		Fr: -float32(0.99691733),
		Fi: float32(0.078459096),
	},
	247: {
		Fr: -float32(0.99580493),
		Fi: float32(0.091501619),
	},
	248: {
		Fr: -float32(0.9945219),
		Fi: float32(0.10452846),
	},
	249: {
		Fr: -float32(0.99306846),
		Fi: float32(0.1175374),
	},
	250: {
		Fr: -float32(0.99144486),
		Fi: float32(0.13052619),
	},
	251: {
		Fr: -float32(0.98965139),
		Fi: float32(0.14349262),
	},
	252: {
		Fr: -float32(0.98768834),
		Fi: float32(0.15643447),
	},
	253: {
		Fr: -float32(0.98555606),
		Fi: float32(0.1693495),
	},
	254: {
		Fr: -float32(0.98325491),
		Fi: float32(0.18223553),
	},
	255: {
		Fr: -float32(0.98078528),
		Fi: float32(0.19509032),
	},
	256: {
		Fr: -float32(0.9781476),
		Fi: float32(0.20791169),
	},
	257: {
		Fr: -float32(0.97534232),
		Fi: float32(0.22069744),
	},
	258: {
		Fr: -float32(0.97236992),
		Fi: float32(0.23344536),
	},
	259: {
		Fr: -float32(0.96923091),
		Fi: float32(0.24615329),
	},
	260: {
		Fr: -float32(0.96592583),
		Fi: float32(0.25881905),
	},
	261: {
		Fr: -float32(0.96245524),
		Fi: float32(0.27144045),
	},
	262: {
		Fr: -float32(0.95881973),
		Fi: float32(0.28401534),
	},
	263: {
		Fr: -float32(0.95501994),
		Fi: float32(0.29654157),
	},
	264: {
		Fr: -float32(0.95105652),
		Fi: float32(0.30901699),
	},
	265: {
		Fr: -float32(0.94693013),
		Fi: float32(0.32143947),
	},
	266: {
		Fr: -float32(0.94264149),
		Fi: float32(0.33380686),
	},
	267: {
		Fr: -float32(0.93819134),
		Fi: float32(0.34611706),
	},
	268: {
		Fr: -float32(0.93358043),
		Fi: float32(0.35836795),
	},
	269: {
		Fr: -float32(0.92880955),
		Fi: float32(0.37055744),
	},
	270: {
		Fr: -float32(0.92387953),
		Fi: float32(0.38268343),
	},
	271: {
		Fr: -float32(0.91879121),
		Fi: float32(0.39474386),
	},
	272: {
		Fr: -float32(0.91354546),
		Fi: float32(0.40673664),
	},
	273: {
		Fr: -float32(0.90814317),
		Fi: float32(0.41865974),
	},
	274: {
		Fr: -float32(0.90258528),
		Fi: float32(0.4305111),
	},
	275: {
		Fr: -float32(0.89687274),
		Fi: float32(0.44228869),
	},
	276: {
		Fr: -float32(0.89100652),
		Fi: float32(0.4539905),
	},
	277: {
		Fr: -float32(0.88498764),
		Fi: float32(0.46561452),
	},
	278: {
		Fr: -float32(0.87881711),
		Fi: float32(0.47715876),
	},
	279: {
		Fr: -float32(0.87249601),
		Fi: float32(0.48862124),
	},
	280: {
		Fr: -float32(0.8660254),
		Fi: float32(0.5),
	},
	281: {
		Fr: -float32(0.85940641),
		Fi: float32(0.51129309),
	},
	282: {
		Fr: -float32(0.85264016),
		Fi: float32(0.52249856),
	},
	283: {
		Fr: -float32(0.84572782),
		Fi: float32(0.53361452),
	},
	284: {
		Fr: -float32(0.83867057),
		Fi: float32(0.54463904),
	},
	285: {
		Fr: -float32(0.83146961),
		Fi: float32(0.55557023),
	},
	286: {
		Fr: -float32(0.82412619),
		Fi: float32(0.56640624),
	},
	287: {
		Fr: -float32(0.81664156),
		Fi: float32(0.57714519),
	},
	288: {
		Fr: -float32(0.80901699),
		Fi: float32(0.58778525),
	},
	289: {
		Fr: -float32(0.80125381),
		Fi: float32(0.5983246),
	},
	290: {
		Fr: -float32(0.79335334),
		Fi: float32(0.60876143),
	},
	291: {
		Fr: -float32(0.78531693),
		Fi: float32(0.61909395),
	},
	292: {
		Fr: -float32(0.77714596),
		Fi: float32(0.62932039),
	},
	293: {
		Fr: -float32(0.76884183),
		Fi: float32(0.639439),
	},
	294: {
		Fr: -float32(0.76040597),
		Fi: float32(0.64944805),
	},
	295: {
		Fr: -float32(0.75183981),
		Fi: float32(0.65934582),
	},
	296: {
		Fr: -float32(0.74314483),
		Fi: float32(0.66913061),
	},
	297: {
		Fr: -float32(0.73432251),
		Fi: float32(0.67880075),
	},
	298: {
		Fr: -float32(0.72537437),
		Fi: float32(0.68835458),
	},
	299: {
		Fr: -float32(0.71630194),
		Fi: float32(0.69779046),
	},
	300: {
		Fr: -float32(0.70710678),
		Fi: float32(0.70710678),
	},
	301: {
		Fr: -float32(0.69779046),
		Fi: float32(0.71630194),
	},
	302: {
		Fr: -float32(0.68835458),
		Fi: float32(0.72537437),
	},
	303: {
		Fr: -float32(0.67880075),
		Fi: float32(0.73432251),
	},
	304: {
		Fr: -float32(0.66913061),
		Fi: float32(0.74314483),
	},
	305: {
		Fr: -float32(0.65934582),
		Fi: float32(0.75183981),
	},
	306: {
		Fr: -float32(0.64944805),
		Fi: float32(0.76040597),
	},
	307: {
		Fr: -float32(0.639439),
		Fi: float32(0.76884183),
	},
	308: {
		Fr: -float32(0.62932039),
		Fi: float32(0.77714596),
	},
	309: {
		Fr: -float32(0.61909395),
		Fi: float32(0.78531693),
	},
	310: {
		Fr: -float32(0.60876143),
		Fi: float32(0.79335334),
	},
	311: {
		Fr: -float32(0.5983246),
		Fi: float32(0.80125381),
	},
	312: {
		Fr: -float32(0.58778525),
		Fi: float32(0.80901699),
	},
	313: {
		Fr: -float32(0.57714519),
		Fi: float32(0.81664156),
	},
	314: {
		Fr: -float32(0.56640624),
		Fi: float32(0.82412619),
	},
	315: {
		Fr: -float32(0.55557023),
		Fi: float32(0.83146961),
	},
	316: {
		Fr: -float32(0.54463904),
		Fi: float32(0.83867057),
	},
	317: {
		Fr: -float32(0.53361452),
		Fi: float32(0.84572782),
	},
	318: {
		Fr: -float32(0.52249856),
		Fi: float32(0.85264016),
	},
	319: {
		Fr: -float32(0.51129309),
		Fi: float32(0.85940641),
	},
	320: {
		Fr: -float32(0.5),
		Fi: float32(0.8660254),
	},
	321: {
		Fr: -float32(0.48862124),
		Fi: float32(0.87249601),
	},
	322: {
		Fr: -float32(0.47715876),
		Fi: float32(0.87881711),
	},
	323: {
		Fr: -float32(0.46561452),
		Fi: float32(0.88498764),
	},
	324: {
		Fr: -float32(0.4539905),
		Fi: float32(0.89100652),
	},
	325: {
		Fr: -float32(0.44228869),
		Fi: float32(0.89687274),
	},
	326: {
		Fr: -float32(0.4305111),
		Fi: float32(0.90258528),
	},
	327: {
		Fr: -float32(0.41865974),
		Fi: float32(0.90814317),
	},
	328: {
		Fr: -float32(0.40673664),
		Fi: float32(0.91354546),
	},
	329: {
		Fr: -float32(0.39474386),
		Fi: float32(0.91879121),
	},
	330: {
		Fr: -float32(0.38268343),
		Fi: float32(0.92387953),
	},
	331: {
		Fr: -float32(0.37055744),
		Fi: float32(0.92880955),
	},
	332: {
		Fr: -float32(0.35836795),
		Fi: float32(0.93358043),
	},
	333: {
		Fr: -float32(0.34611706),
		Fi: float32(0.93819134),
	},
	334: {
		Fr: -float32(0.33380686),
		Fi: float32(0.94264149),
	},
	335: {
		Fr: -float32(0.32143947),
		Fi: float32(0.94693013),
	},
	336: {
		Fr: -float32(0.30901699),
		Fi: float32(0.95105652),
	},
	337: {
		Fr: -float32(0.29654157),
		Fi: float32(0.95501994),
	},
	338: {
		Fr: -float32(0.28401534),
		Fi: float32(0.95881973),
	},
	339: {
		Fr: -float32(0.27144045),
		Fi: float32(0.96245524),
	},
	340: {
		Fr: -float32(0.25881905),
		Fi: float32(0.96592583),
	},
	341: {
		Fr: -float32(0.24615329),
		Fi: float32(0.96923091),
	},
	342: {
		Fr: -float32(0.23344536),
		Fi: float32(0.97236992),
	},
	343: {
		Fr: -float32(0.22069744),
		Fi: float32(0.97534232),
	},
	344: {
		Fr: -float32(0.20791169),
		Fi: float32(0.9781476),
	},
	345: {
		Fr: -float32(0.19509032),
		Fi: float32(0.98078528),
	},
	346: {
		Fr: -float32(0.18223553),
		Fi: float32(0.98325491),
	},
	347: {
		Fr: -float32(0.1693495),
		Fi: float32(0.98555606),
	},
	348: {
		Fr: -float32(0.15643447),
		Fi: float32(0.98768834),
	},
	349: {
		Fr: -float32(0.14349262),
		Fi: float32(0.98965139),
	},
	350: {
		Fr: -float32(0.13052619),
		Fi: float32(0.99144486),
	},
	351: {
		Fr: -float32(0.1175374),
		Fi: float32(0.99306846),
	},
	352: {
		Fr: -float32(0.10452846),
		Fi: float32(0.9945219),
	},
	353: {
		Fr: -float32(0.091501619),
		Fi: float32(0.99580493),
	},
	354: {
		Fr: -float32(0.078459096),
		Fi: float32(0.99691733),
	},
	355: {
		Fr: -float32(0.065403129),
		Fi: float32(0.99785892),
	},
	356: {
		Fr: -float32(0.052335956),
		Fi: float32(0.99862953),
	},
	357: {
		Fr: -float32(0.039259816),
		Fi: float32(0.99922904),
	},
	358: {
		Fr: -float32(0.026176948),
		Fi: float32(0.99965732),
	},
	359: {
		Fr: -float32(0.013089596),
		Fi: float32(0.99991433),
	},
	360: {
		Fr: -float32(1.8369702e-16),
		Fi: float32(1),
	},
	361: {
		Fr: float32(0.013089596),
		Fi: float32(0.99991433),
	},
	362: {
		Fr: float32(0.026176948),
		Fi: float32(0.99965732),
	},
	363: {
		Fr: float32(0.039259816),
		Fi: float32(0.99922904),
	},
	364: {
		Fr: float32(0.052335956),
		Fi: float32(0.99862953),
	},
	365: {
		Fr: float32(0.065403129),
		Fi: float32(0.99785892),
	},
	366: {
		Fr: float32(0.078459096),
		Fi: float32(0.99691733),
	},
	367: {
		Fr: float32(0.091501619),
		Fi: float32(0.99580493),
	},
	368: {
		Fr: float32(0.10452846),
		Fi: float32(0.9945219),
	},
	369: {
		Fr: float32(0.1175374),
		Fi: float32(0.99306846),
	},
	370: {
		Fr: float32(0.13052619),
		Fi: float32(0.99144486),
	},
	371: {
		Fr: float32(0.14349262),
		Fi: float32(0.98965139),
	},
	372: {
		Fr: float32(0.15643447),
		Fi: float32(0.98768834),
	},
	373: {
		Fr: float32(0.1693495),
		Fi: float32(0.98555606),
	},
	374: {
		Fr: float32(0.18223553),
		Fi: float32(0.98325491),
	},
	375: {
		Fr: float32(0.19509032),
		Fi: float32(0.98078528),
	},
	376: {
		Fr: float32(0.20791169),
		Fi: float32(0.9781476),
	},
	377: {
		Fr: float32(0.22069744),
		Fi: float32(0.97534232),
	},
	378: {
		Fr: float32(0.23344536),
		Fi: float32(0.97236992),
	},
	379: {
		Fr: float32(0.24615329),
		Fi: float32(0.96923091),
	},
	380: {
		Fr: float32(0.25881905),
		Fi: float32(0.96592583),
	},
	381: {
		Fr: float32(0.27144045),
		Fi: float32(0.96245524),
	},
	382: {
		Fr: float32(0.28401534),
		Fi: float32(0.95881973),
	},
	383: {
		Fr: float32(0.29654157),
		Fi: float32(0.95501994),
	},
	384: {
		Fr: float32(0.30901699),
		Fi: float32(0.95105652),
	},
	385: {
		Fr: float32(0.32143947),
		Fi: float32(0.94693013),
	},
	386: {
		Fr: float32(0.33380686),
		Fi: float32(0.94264149),
	},
	387: {
		Fr: float32(0.34611706),
		Fi: float32(0.93819134),
	},
	388: {
		Fr: float32(0.35836795),
		Fi: float32(0.93358043),
	},
	389: {
		Fr: float32(0.37055744),
		Fi: float32(0.92880955),
	},
	390: {
		Fr: float32(0.38268343),
		Fi: float32(0.92387953),
	},
	391: {
		Fr: float32(0.39474386),
		Fi: float32(0.91879121),
	},
	392: {
		Fr: float32(0.40673664),
		Fi: float32(0.91354546),
	},
	393: {
		Fr: float32(0.41865974),
		Fi: float32(0.90814317),
	},
	394: {
		Fr: float32(0.4305111),
		Fi: float32(0.90258528),
	},
	395: {
		Fr: float32(0.44228869),
		Fi: float32(0.89687274),
	},
	396: {
		Fr: float32(0.4539905),
		Fi: float32(0.89100652),
	},
	397: {
		Fr: float32(0.46561452),
		Fi: float32(0.88498764),
	},
	398: {
		Fr: float32(0.47715876),
		Fi: float32(0.87881711),
	},
	399: {
		Fr: float32(0.48862124),
		Fi: float32(0.87249601),
	},
	400: {
		Fr: float32(0.5),
		Fi: float32(0.8660254),
	},
	401: {
		Fr: float32(0.51129309),
		Fi: float32(0.85940641),
	},
	402: {
		Fr: float32(0.52249856),
		Fi: float32(0.85264016),
	},
	403: {
		Fr: float32(0.53361452),
		Fi: float32(0.84572782),
	},
	404: {
		Fr: float32(0.54463904),
		Fi: float32(0.83867057),
	},
	405: {
		Fr: float32(0.55557023),
		Fi: float32(0.83146961),
	},
	406: {
		Fr: float32(0.56640624),
		Fi: float32(0.82412619),
	},
	407: {
		Fr: float32(0.57714519),
		Fi: float32(0.81664156),
	},
	408: {
		Fr: float32(0.58778525),
		Fi: float32(0.80901699),
	},
	409: {
		Fr: float32(0.5983246),
		Fi: float32(0.80125381),
	},
	410: {
		Fr: float32(0.60876143),
		Fi: float32(0.79335334),
	},
	411: {
		Fr: float32(0.61909395),
		Fi: float32(0.78531693),
	},
	412: {
		Fr: float32(0.62932039),
		Fi: float32(0.77714596),
	},
	413: {
		Fr: float32(0.639439),
		Fi: float32(0.76884183),
	},
	414: {
		Fr: float32(0.64944805),
		Fi: float32(0.76040597),
	},
	415: {
		Fr: float32(0.65934582),
		Fi: float32(0.75183981),
	},
	416: {
		Fr: float32(0.66913061),
		Fi: float32(0.74314483),
	},
	417: {
		Fr: float32(0.67880075),
		Fi: float32(0.73432251),
	},
	418: {
		Fr: float32(0.68835458),
		Fi: float32(0.72537437),
	},
	419: {
		Fr: float32(0.69779046),
		Fi: float32(0.71630194),
	},
	420: {
		Fr: float32(0.70710678),
		Fi: float32(0.70710678),
	},
	421: {
		Fr: float32(0.71630194),
		Fi: float32(0.69779046),
	},
	422: {
		Fr: float32(0.72537437),
		Fi: float32(0.68835458),
	},
	423: {
		Fr: float32(0.73432251),
		Fi: float32(0.67880075),
	},
	424: {
		Fr: float32(0.74314483),
		Fi: float32(0.66913061),
	},
	425: {
		Fr: float32(0.75183981),
		Fi: float32(0.65934582),
	},
	426: {
		Fr: float32(0.76040597),
		Fi: float32(0.64944805),
	},
	427: {
		Fr: float32(0.76884183),
		Fi: float32(0.639439),
	},
	428: {
		Fr: float32(0.77714596),
		Fi: float32(0.62932039),
	},
	429: {
		Fr: float32(0.78531693),
		Fi: float32(0.61909395),
	},
	430: {
		Fr: float32(0.79335334),
		Fi: float32(0.60876143),
	},
	431: {
		Fr: float32(0.80125381),
		Fi: float32(0.5983246),
	},
	432: {
		Fr: float32(0.80901699),
		Fi: float32(0.58778525),
	},
	433: {
		Fr: float32(0.81664156),
		Fi: float32(0.57714519),
	},
	434: {
		Fr: float32(0.82412619),
		Fi: float32(0.56640624),
	},
	435: {
		Fr: float32(0.83146961),
		Fi: float32(0.55557023),
	},
	436: {
		Fr: float32(0.83867057),
		Fi: float32(0.54463904),
	},
	437: {
		Fr: float32(0.84572782),
		Fi: float32(0.53361452),
	},
	438: {
		Fr: float32(0.85264016),
		Fi: float32(0.52249856),
	},
	439: {
		Fr: float32(0.85940641),
		Fi: float32(0.51129309),
	},
	440: {
		Fr: float32(0.8660254),
		Fi: float32(0.5),
	},
	441: {
		Fr: float32(0.87249601),
		Fi: float32(0.48862124),
	},
	442: {
		Fr: float32(0.87881711),
		Fi: float32(0.47715876),
	},
	443: {
		Fr: float32(0.88498764),
		Fi: float32(0.46561452),
	},
	444: {
		Fr: float32(0.89100652),
		Fi: float32(0.4539905),
	},
	445: {
		Fr: float32(0.89687274),
		Fi: float32(0.44228869),
	},
	446: {
		Fr: float32(0.90258528),
		Fi: float32(0.4305111),
	},
	447: {
		Fr: float32(0.90814317),
		Fi: float32(0.41865974),
	},
	448: {
		Fr: float32(0.91354546),
		Fi: float32(0.40673664),
	},
	449: {
		Fr: float32(0.91879121),
		Fi: float32(0.39474386),
	},
	450: {
		Fr: float32(0.92387953),
		Fi: float32(0.38268343),
	},
	451: {
		Fr: float32(0.92880955),
		Fi: float32(0.37055744),
	},
	452: {
		Fr: float32(0.93358043),
		Fi: float32(0.35836795),
	},
	453: {
		Fr: float32(0.93819134),
		Fi: float32(0.34611706),
	},
	454: {
		Fr: float32(0.94264149),
		Fi: float32(0.33380686),
	},
	455: {
		Fr: float32(0.94693013),
		Fi: float32(0.32143947),
	},
	456: {
		Fr: float32(0.95105652),
		Fi: float32(0.30901699),
	},
	457: {
		Fr: float32(0.95501994),
		Fi: float32(0.29654157),
	},
	458: {
		Fr: float32(0.95881973),
		Fi: float32(0.28401534),
	},
	459: {
		Fr: float32(0.96245524),
		Fi: float32(0.27144045),
	},
	460: {
		Fr: float32(0.96592583),
		Fi: float32(0.25881905),
	},
	461: {
		Fr: float32(0.96923091),
		Fi: float32(0.24615329),
	},
	462: {
		Fr: float32(0.97236992),
		Fi: float32(0.23344536),
	},
	463: {
		Fr: float32(0.97534232),
		Fi: float32(0.22069744),
	},
	464: {
		Fr: float32(0.9781476),
		Fi: float32(0.20791169),
	},
	465: {
		Fr: float32(0.98078528),
		Fi: float32(0.19509032),
	},
	466: {
		Fr: float32(0.98325491),
		Fi: float32(0.18223553),
	},
	467: {
		Fr: float32(0.98555606),
		Fi: float32(0.1693495),
	},
	468: {
		Fr: float32(0.98768834),
		Fi: float32(0.15643447),
	},
	469: {
		Fr: float32(0.98965139),
		Fi: float32(0.14349262),
	},
	470: {
		Fr: float32(0.99144486),
		Fi: float32(0.13052619),
	},
	471: {
		Fr: float32(0.99306846),
		Fi: float32(0.1175374),
	},
	472: {
		Fr: float32(0.9945219),
		Fi: float32(0.10452846),
	},
	473: {
		Fr: float32(0.99580493),
		Fi: float32(0.091501619),
	},
	474: {
		Fr: float32(0.99691733),
		Fi: float32(0.078459096),
	},
	475: {
		Fr: float32(0.99785892),
		Fi: float32(0.065403129),
	},
	476: {
		Fr: float32(0.99862953),
		Fi: float32(0.052335956),
	},
	477: {
		Fr: float32(0.99922904),
		Fi: float32(0.039259816),
	},
	478: {
		Fr: float32(0.99965732),
		Fi: float32(0.026176948),
	},
	479: {
		Fr: float32(0.99991433),
		Fi: float32(0.013089596),
	},
}
var fft_bitrev480 = [480]OpusT_opus_int16{
	1:   int16(96),
	2:   int16(192),
	3:   int16(288),
	4:   int16(384),
	5:   int16(32),
	6:   int16(128),
	7:   int16(224),
	8:   int16(320),
	9:   int16(416),
	10:  int16(64),
	11:  int16(160),
	12:  int16(256),
	13:  int16(352),
	14:  int16(448),
	15:  int16(8),
	16:  int16(104),
	17:  int16(200),
	18:  int16(296),
	19:  int16(392),
	20:  int16(40),
	21:  int16(136),
	22:  int16(232),
	23:  int16(328),
	24:  int16(424),
	25:  int16(72),
	26:  int16(168),
	27:  int16(264),
	28:  int16(360),
	29:  int16(456),
	30:  int16(16),
	31:  int16(112),
	32:  int16(208),
	33:  int16(304),
	34:  int16(400),
	35:  int16(48),
	36:  int16(144),
	37:  int16(240),
	38:  int16(336),
	39:  int16(432),
	40:  int16(80),
	41:  int16(176),
	42:  int16(272),
	43:  int16(368),
	44:  int16(464),
	45:  int16(24),
	46:  int16(120),
	47:  int16(216),
	48:  int16(312),
	49:  int16(408),
	50:  int16(56),
	51:  int16(152),
	52:  int16(248),
	53:  int16(344),
	54:  int16(440),
	55:  int16(88),
	56:  int16(184),
	57:  int16(280),
	58:  int16(376),
	59:  int16(472),
	60:  int16(4),
	61:  int16(100),
	62:  int16(196),
	63:  int16(292),
	64:  int16(388),
	65:  int16(36),
	66:  int16(132),
	67:  int16(228),
	68:  int16(324),
	69:  int16(420),
	70:  int16(68),
	71:  int16(164),
	72:  int16(260),
	73:  int16(356),
	74:  int16(452),
	75:  int16(12),
	76:  int16(108),
	77:  int16(204),
	78:  int16(300),
	79:  int16(396),
	80:  int16(44),
	81:  int16(140),
	82:  int16(236),
	83:  int16(332),
	84:  int16(428),
	85:  int16(76),
	86:  int16(172),
	87:  int16(268),
	88:  int16(364),
	89:  int16(460),
	90:  int16(20),
	91:  int16(116),
	92:  int16(212),
	93:  int16(308),
	94:  int16(404),
	95:  int16(52),
	96:  int16(148),
	97:  int16(244),
	98:  int16(340),
	99:  int16(436),
	100: int16(84),
	101: int16(180),
	102: int16(276),
	103: int16(372),
	104: int16(468),
	105: int16(28),
	106: int16(124),
	107: int16(220),
	108: int16(316),
	109: int16(412),
	110: int16(60),
	111: int16(156),
	112: int16(252),
	113: int16(348),
	114: int16(444),
	115: int16(92),
	116: int16(188),
	117: int16(284),
	118: int16(380),
	119: int16(476),
	120: int16(1),
	121: int16(97),
	122: int16(193),
	123: int16(289),
	124: int16(385),
	125: int16(33),
	126: int16(129),
	127: int16(225),
	128: int16(321),
	129: int16(417),
	130: int16(65),
	131: int16(161),
	132: int16(257),
	133: int16(353),
	134: int16(449),
	135: int16(9),
	136: int16(105),
	137: int16(201),
	138: int16(297),
	139: int16(393),
	140: int16(41),
	141: int16(137),
	142: int16(233),
	143: int16(329),
	144: int16(425),
	145: int16(73),
	146: int16(169),
	147: int16(265),
	148: int16(361),
	149: int16(457),
	150: int16(17),
	151: int16(113),
	152: int16(209),
	153: int16(305),
	154: int16(401),
	155: int16(49),
	156: int16(145),
	157: int16(241),
	158: int16(337),
	159: int16(433),
	160: int16(81),
	161: int16(177),
	162: int16(273),
	163: int16(369),
	164: int16(465),
	165: int16(25),
	166: int16(121),
	167: int16(217),
	168: int16(313),
	169: int16(409),
	170: int16(57),
	171: int16(153),
	172: int16(249),
	173: int16(345),
	174: int16(441),
	175: int16(89),
	176: int16(185),
	177: int16(281),
	178: int16(377),
	179: int16(473),
	180: int16(5),
	181: int16(101),
	182: int16(197),
	183: int16(293),
	184: int16(389),
	185: int16(37),
	186: int16(133),
	187: int16(229),
	188: int16(325),
	189: int16(421),
	190: int16(69),
	191: int16(165),
	192: int16(261),
	193: int16(357),
	194: int16(453),
	195: int16(13),
	196: int16(109),
	197: int16(205),
	198: int16(301),
	199: int16(397),
	200: int16(45),
	201: int16(141),
	202: int16(237),
	203: int16(333),
	204: int16(429),
	205: int16(77),
	206: int16(173),
	207: int16(269),
	208: int16(365),
	209: int16(461),
	210: int16(21),
	211: int16(117),
	212: int16(213),
	213: int16(309),
	214: int16(405),
	215: int16(53),
	216: int16(149),
	217: int16(245),
	218: int16(341),
	219: int16(437),
	220: int16(85),
	221: int16(181),
	222: int16(277),
	223: int16(373),
	224: int16(469),
	225: int16(29),
	226: int16(125),
	227: int16(221),
	228: int16(317),
	229: int16(413),
	230: int16(61),
	231: int16(157),
	232: int16(253),
	233: int16(349),
	234: int16(445),
	235: int16(93),
	236: int16(189),
	237: int16(285),
	238: int16(381),
	239: int16(477),
	240: int16(2),
	241: int16(98),
	242: int16(194),
	243: int16(290),
	244: int16(386),
	245: int16(34),
	246: int16(130),
	247: int16(226),
	248: int16(322),
	249: int16(418),
	250: int16(66),
	251: int16(162),
	252: int16(258),
	253: int16(354),
	254: int16(450),
	255: int16(10),
	256: int16(106),
	257: int16(202),
	258: int16(298),
	259: int16(394),
	260: int16(42),
	261: int16(138),
	262: int16(234),
	263: int16(330),
	264: int16(426),
	265: int16(74),
	266: int16(170),
	267: int16(266),
	268: int16(362),
	269: int16(458),
	270: int16(18),
	271: int16(114),
	272: int16(210),
	273: int16(306),
	274: int16(402),
	275: int16(50),
	276: int16(146),
	277: int16(242),
	278: int16(338),
	279: int16(434),
	280: int16(82),
	281: int16(178),
	282: int16(274),
	283: int16(370),
	284: int16(466),
	285: int16(26),
	286: int16(122),
	287: int16(218),
	288: int16(314),
	289: int16(410),
	290: int16(58),
	291: int16(154),
	292: int16(250),
	293: int16(346),
	294: int16(442),
	295: int16(90),
	296: int16(186),
	297: int16(282),
	298: int16(378),
	299: int16(474),
	300: int16(6),
	301: int16(102),
	302: int16(198),
	303: int16(294),
	304: int16(390),
	305: int16(38),
	306: int16(134),
	307: int16(230),
	308: int16(326),
	309: int16(422),
	310: int16(70),
	311: int16(166),
	312: int16(262),
	313: int16(358),
	314: int16(454),
	315: int16(14),
	316: int16(110),
	317: int16(206),
	318: int16(302),
	319: int16(398),
	320: int16(46),
	321: int16(142),
	322: int16(238),
	323: int16(334),
	324: int16(430),
	325: int16(78),
	326: int16(174),
	327: int16(270),
	328: int16(366),
	329: int16(462),
	330: int16(22),
	331: int16(118),
	332: int16(214),
	333: int16(310),
	334: int16(406),
	335: int16(54),
	336: int16(150),
	337: int16(246),
	338: int16(342),
	339: int16(438),
	340: int16(86),
	341: int16(182),
	342: int16(278),
	343: int16(374),
	344: int16(470),
	345: int16(30),
	346: int16(126),
	347: int16(222),
	348: int16(318),
	349: int16(414),
	350: int16(62),
	351: int16(158),
	352: int16(254),
	353: int16(350),
	354: int16(446),
	355: int16(94),
	356: int16(190),
	357: int16(286),
	358: int16(382),
	359: int16(478),
	360: int16(3),
	361: int16(99),
	362: int16(195),
	363: int16(291),
	364: int16(387),
	365: int16(35),
	366: int16(131),
	367: int16(227),
	368: int16(323),
	369: int16(419),
	370: int16(67),
	371: int16(163),
	372: int16(259),
	373: int16(355),
	374: int16(451),
	375: int16(11),
	376: int16(107),
	377: int16(203),
	378: int16(299),
	379: int16(395),
	380: int16(43),
	381: int16(139),
	382: int16(235),
	383: int16(331),
	384: int16(427),
	385: int16(75),
	386: int16(171),
	387: int16(267),
	388: int16(363),
	389: int16(459),
	390: int16(19),
	391: int16(115),
	392: int16(211),
	393: int16(307),
	394: int16(403),
	395: int16(51),
	396: int16(147),
	397: int16(243),
	398: int16(339),
	399: int16(435),
	400: int16(83),
	401: int16(179),
	402: int16(275),
	403: int16(371),
	404: int16(467),
	405: int16(27),
	406: int16(123),
	407: int16(219),
	408: int16(315),
	409: int16(411),
	410: int16(59),
	411: int16(155),
	412: int16(251),
	413: int16(347),
	414: int16(443),
	415: int16(91),
	416: int16(187),
	417: int16(283),
	418: int16(379),
	419: int16(475),
	420: int16(7),
	421: int16(103),
	422: int16(199),
	423: int16(295),
	424: int16(391),
	425: int16(39),
	426: int16(135),
	427: int16(231),
	428: int16(327),
	429: int16(423),
	430: int16(71),
	431: int16(167),
	432: int16(263),
	433: int16(359),
	434: int16(455),
	435: int16(15),
	436: int16(111),
	437: int16(207),
	438: int16(303),
	439: int16(399),
	440: int16(47),
	441: int16(143),
	442: int16(239),
	443: int16(335),
	444: int16(431),
	445: int16(79),
	446: int16(175),
	447: int16(271),
	448: int16(367),
	449: int16(463),
	450: int16(23),
	451: int16(119),
	452: int16(215),
	453: int16(311),
	454: int16(407),
	455: int16(55),
	456: int16(151),
	457: int16(247),
	458: int16(343),
	459: int16(439),
	460: int16(87),
	461: int16(183),
	462: int16(279),
	463: int16(375),
	464: int16(471),
	465: int16(31),
	466: int16(127),
	467: int16(223),
	468: int16(319),
	469: int16(415),
	470: int16(63),
	471: int16(159),
	472: int16(255),
	473: int16(351),
	474: int16(447),
	475: int16(95),
	476: int16(191),
	477: int16(287),
	478: int16(383),
	479: int16(479),
}
var fft_bitrev240 = [240]OpusT_opus_int16{
	1:   int16(48),
	2:   int16(96),
	3:   int16(144),
	4:   int16(192),
	5:   int16(16),
	6:   int16(64),
	7:   int16(112),
	8:   int16(160),
	9:   int16(208),
	10:  int16(32),
	11:  int16(80),
	12:  int16(128),
	13:  int16(176),
	14:  int16(224),
	15:  int16(4),
	16:  int16(52),
	17:  int16(100),
	18:  int16(148),
	19:  int16(196),
	20:  int16(20),
	21:  int16(68),
	22:  int16(116),
	23:  int16(164),
	24:  int16(212),
	25:  int16(36),
	26:  int16(84),
	27:  int16(132),
	28:  int16(180),
	29:  int16(228),
	30:  int16(8),
	31:  int16(56),
	32:  int16(104),
	33:  int16(152),
	34:  int16(200),
	35:  int16(24),
	36:  int16(72),
	37:  int16(120),
	38:  int16(168),
	39:  int16(216),
	40:  int16(40),
	41:  int16(88),
	42:  int16(136),
	43:  int16(184),
	44:  int16(232),
	45:  int16(12),
	46:  int16(60),
	47:  int16(108),
	48:  int16(156),
	49:  int16(204),
	50:  int16(28),
	51:  int16(76),
	52:  int16(124),
	53:  int16(172),
	54:  int16(220),
	55:  int16(44),
	56:  int16(92),
	57:  int16(140),
	58:  int16(188),
	59:  int16(236),
	60:  int16(1),
	61:  int16(49),
	62:  int16(97),
	63:  int16(145),
	64:  int16(193),
	65:  int16(17),
	66:  int16(65),
	67:  int16(113),
	68:  int16(161),
	69:  int16(209),
	70:  int16(33),
	71:  int16(81),
	72:  int16(129),
	73:  int16(177),
	74:  int16(225),
	75:  int16(5),
	76:  int16(53),
	77:  int16(101),
	78:  int16(149),
	79:  int16(197),
	80:  int16(21),
	81:  int16(69),
	82:  int16(117),
	83:  int16(165),
	84:  int16(213),
	85:  int16(37),
	86:  int16(85),
	87:  int16(133),
	88:  int16(181),
	89:  int16(229),
	90:  int16(9),
	91:  int16(57),
	92:  int16(105),
	93:  int16(153),
	94:  int16(201),
	95:  int16(25),
	96:  int16(73),
	97:  int16(121),
	98:  int16(169),
	99:  int16(217),
	100: int16(41),
	101: int16(89),
	102: int16(137),
	103: int16(185),
	104: int16(233),
	105: int16(13),
	106: int16(61),
	107: int16(109),
	108: int16(157),
	109: int16(205),
	110: int16(29),
	111: int16(77),
	112: int16(125),
	113: int16(173),
	114: int16(221),
	115: int16(45),
	116: int16(93),
	117: int16(141),
	118: int16(189),
	119: int16(237),
	120: int16(2),
	121: int16(50),
	122: int16(98),
	123: int16(146),
	124: int16(194),
	125: int16(18),
	126: int16(66),
	127: int16(114),
	128: int16(162),
	129: int16(210),
	130: int16(34),
	131: int16(82),
	132: int16(130),
	133: int16(178),
	134: int16(226),
	135: int16(6),
	136: int16(54),
	137: int16(102),
	138: int16(150),
	139: int16(198),
	140: int16(22),
	141: int16(70),
	142: int16(118),
	143: int16(166),
	144: int16(214),
	145: int16(38),
	146: int16(86),
	147: int16(134),
	148: int16(182),
	149: int16(230),
	150: int16(10),
	151: int16(58),
	152: int16(106),
	153: int16(154),
	154: int16(202),
	155: int16(26),
	156: int16(74),
	157: int16(122),
	158: int16(170),
	159: int16(218),
	160: int16(42),
	161: int16(90),
	162: int16(138),
	163: int16(186),
	164: int16(234),
	165: int16(14),
	166: int16(62),
	167: int16(110),
	168: int16(158),
	169: int16(206),
	170: int16(30),
	171: int16(78),
	172: int16(126),
	173: int16(174),
	174: int16(222),
	175: int16(46),
	176: int16(94),
	177: int16(142),
	178: int16(190),
	179: int16(238),
	180: int16(3),
	181: int16(51),
	182: int16(99),
	183: int16(147),
	184: int16(195),
	185: int16(19),
	186: int16(67),
	187: int16(115),
	188: int16(163),
	189: int16(211),
	190: int16(35),
	191: int16(83),
	192: int16(131),
	193: int16(179),
	194: int16(227),
	195: int16(7),
	196: int16(55),
	197: int16(103),
	198: int16(151),
	199: int16(199),
	200: int16(23),
	201: int16(71),
	202: int16(119),
	203: int16(167),
	204: int16(215),
	205: int16(39),
	206: int16(87),
	207: int16(135),
	208: int16(183),
	209: int16(231),
	210: int16(11),
	211: int16(59),
	212: int16(107),
	213: int16(155),
	214: int16(203),
	215: int16(27),
	216: int16(75),
	217: int16(123),
	218: int16(171),
	219: int16(219),
	220: int16(43),
	221: int16(91),
	222: int16(139),
	223: int16(187),
	224: int16(235),
	225: int16(15),
	226: int16(63),
	227: int16(111),
	228: int16(159),
	229: int16(207),
	230: int16(31),
	231: int16(79),
	232: int16(127),
	233: int16(175),
	234: int16(223),
	235: int16(47),
	236: int16(95),
	237: int16(143),
	238: int16(191),
	239: int16(239),
}
var fft_bitrev120 = [120]OpusT_opus_int16{
	1:   int16(24),
	2:   int16(48),
	3:   int16(72),
	4:   int16(96),
	5:   int16(8),
	6:   int16(32),
	7:   int16(56),
	8:   int16(80),
	9:   int16(104),
	10:  int16(16),
	11:  int16(40),
	12:  int16(64),
	13:  int16(88),
	14:  int16(112),
	15:  int16(4),
	16:  int16(28),
	17:  int16(52),
	18:  int16(76),
	19:  int16(100),
	20:  int16(12),
	21:  int16(36),
	22:  int16(60),
	23:  int16(84),
	24:  int16(108),
	25:  int16(20),
	26:  int16(44),
	27:  int16(68),
	28:  int16(92),
	29:  int16(116),
	30:  int16(1),
	31:  int16(25),
	32:  int16(49),
	33:  int16(73),
	34:  int16(97),
	35:  int16(9),
	36:  int16(33),
	37:  int16(57),
	38:  int16(81),
	39:  int16(105),
	40:  int16(17),
	41:  int16(41),
	42:  int16(65),
	43:  int16(89),
	44:  int16(113),
	45:  int16(5),
	46:  int16(29),
	47:  int16(53),
	48:  int16(77),
	49:  int16(101),
	50:  int16(13),
	51:  int16(37),
	52:  int16(61),
	53:  int16(85),
	54:  int16(109),
	55:  int16(21),
	56:  int16(45),
	57:  int16(69),
	58:  int16(93),
	59:  int16(117),
	60:  int16(2),
	61:  int16(26),
	62:  int16(50),
	63:  int16(74),
	64:  int16(98),
	65:  int16(10),
	66:  int16(34),
	67:  int16(58),
	68:  int16(82),
	69:  int16(106),
	70:  int16(18),
	71:  int16(42),
	72:  int16(66),
	73:  int16(90),
	74:  int16(114),
	75:  int16(6),
	76:  int16(30),
	77:  int16(54),
	78:  int16(78),
	79:  int16(102),
	80:  int16(14),
	81:  int16(38),
	82:  int16(62),
	83:  int16(86),
	84:  int16(110),
	85:  int16(22),
	86:  int16(46),
	87:  int16(70),
	88:  int16(94),
	89:  int16(118),
	90:  int16(3),
	91:  int16(27),
	92:  int16(51),
	93:  int16(75),
	94:  int16(99),
	95:  int16(11),
	96:  int16(35),
	97:  int16(59),
	98:  int16(83),
	99:  int16(107),
	100: int16(19),
	101: int16(43),
	102: int16(67),
	103: int16(91),
	104: int16(115),
	105: int16(7),
	106: int16(31),
	107: int16(55),
	108: int16(79),
	109: int16(103),
	110: int16(15),
	111: int16(39),
	112: int16(63),
	113: int16(87),
	114: int16(111),
	115: int16(23),
	116: int16(47),
	117: int16(71),
	118: int16(95),
	119: int16(119),
}
var fft_bitrev60 = [60]OpusT_opus_int16{
	1:  int16(12),
	2:  int16(24),
	3:  int16(36),
	4:  int16(48),
	5:  int16(4),
	6:  int16(16),
	7:  int16(28),
	8:  int16(40),
	9:  int16(52),
	10: int16(8),
	11: int16(20),
	12: int16(32),
	13: int16(44),
	14: int16(56),
	15: int16(1),
	16: int16(13),
	17: int16(25),
	18: int16(37),
	19: int16(49),
	20: int16(5),
	21: int16(17),
	22: int16(29),
	23: int16(41),
	24: int16(53),
	25: int16(9),
	26: int16(21),
	27: int16(33),
	28: int16(45),
	29: int16(57),
	30: int16(2),
	31: int16(14),
	32: int16(26),
	33: int16(38),
	34: int16(50),
	35: int16(6),
	36: int16(18),
	37: int16(30),
	38: int16(42),
	39: int16(54),
	40: int16(10),
	41: int16(22),
	42: int16(34),
	43: int16(46),
	44: int16(58),
	45: int16(3),
	46: int16(15),
	47: int16(27),
	48: int16(39),
	49: int16(51),
	50: int16(7),
	51: int16(19),
	52: int16(31),
	53: int16(43),
	54: int16(55),
	55: int16(11),
	56: int16(23),
	57: int16(35),
	58: int16(47),
	59: int16(59),
}
var fft_state48000_960_0 = OpusT_kiss_fft_state{
	Fnfft:  int32(480),
	Fscale: float32(0.0020833334),
	Fshift: -int32(1),
	Ffactors: [16]OpusT_opus_int16{
		0: int16(5),
		1: int16(96),
		2: int16(3),
		3: int16(32),
		4: int16(4),
		5: int16(8),
		6: int16(2),
		7: int16(4),
		8: int16(4),
		9: int16(1),
	},
	Fbitrev:   uintptr(unsafe.Pointer(&fft_bitrev480)),
	Ftwiddles: uintptr(unsafe.Pointer(&fft_twiddles48000_960)),
}
var fft_state48000_960_1 = OpusT_kiss_fft_state{
	Fnfft:  int32(240),
	Fscale: float32(0.0041666669),
	Fshift: int32(1),
	Ffactors: [16]OpusT_opus_int16{
		0: int16(5),
		1: int16(48),
		2: int16(3),
		3: int16(16),
		4: int16(4),
		5: int16(4),
		6: int16(4),
		7: int16(1),
	},
	Fbitrev:   uintptr(unsafe.Pointer(&fft_bitrev240)),
	Ftwiddles: uintptr(unsafe.Pointer(&fft_twiddles48000_960)),
}
var fft_state48000_960_2 = OpusT_kiss_fft_state{
	Fnfft:  int32(120),
	Fscale: float32(0.0083333338),
	Fshift: int32(2),
	Ffactors: [16]OpusT_opus_int16{
		0: int16(5),
		1: int16(24),
		2: int16(3),
		3: int16(8),
		4: int16(2),
		5: int16(4),
		6: int16(4),
		7: int16(1),
	},
	Fbitrev:   uintptr(unsafe.Pointer(&fft_bitrev120)),
	Ftwiddles: uintptr(unsafe.Pointer(&fft_twiddles48000_960)),
}
var fft_state48000_960_3 = OpusT_kiss_fft_state{
	Fnfft:  int32(60),
	Fscale: float32(0.016666668),
	Fshift: int32(3),
	Ffactors: [16]OpusT_opus_int16{
		0: int16(5),
		1: int16(12),
		2: int16(3),
		3: int16(4),
		4: int16(4),
		5: int16(1),
	},
	Fbitrev:   uintptr(unsafe.Pointer(&fft_bitrev60)),
	Ftwiddles: uintptr(unsafe.Pointer(&fft_twiddles48000_960)),
}
var mdct_twiddles960 = [1800]OpusT_celt_coef{
	0:    float32(0.99999992),
	1:    float32(0.99999322),
	2:    float32(0.99997582),
	3:    float32(0.99994771),
	4:    float32(0.99990889),
	5:    float32(0.99985936),
	6:    float32(0.99979913),
	7:    float32(0.99972818),
	8:    float32(0.99964653),
	9:    float32(0.99955418),
	10:   float32(0.99945112),
	11:   float32(0.99933736),
	12:   float32(0.99921289),
	13:   float32(0.99907773),
	14:   float32(0.99893186),
	15:   float32(0.9987753),
	16:   float32(0.99860804),
	17:   float32(0.99843009),
	18:   float32(0.99824144),
	19:   float32(0.99804211),
	20:   float32(0.99783209),
	21:   float32(0.99761138),
	22:   float32(0.99737998),
	23:   float32(0.99713791),
	24:   float32(0.99688516),
	25:   float32(0.99662173),
	26:   float32(0.99634763),
	27:   float32(0.99606285),
	28:   float32(0.99576741),
	29:   float32(0.99546131),
	30:   float32(0.99514455),
	31:   float32(0.99481713),
	32:   float32(0.99447905),
	33:   float32(0.99413033),
	34:   float32(0.99377096),
	35:   float32(0.99340095),
	36:   float32(0.99302029),
	37:   float32(0.99262901),
	38:   float32(0.99222709),
	39:   float32(0.99181455),
	40:   float32(0.99139139),
	41:   float32(0.9909576),
	42:   float32(0.99051321),
	43:   float32(0.99005821),
	44:   float32(0.98959261),
	45:   float32(0.98911641),
	46:   float32(0.98862961),
	47:   float32(0.98813223),
	48:   float32(0.98762427),
	49:   float32(0.98710573),
	50:   float32(0.98657662),
	51:   float32(0.98603694),
	52:   float32(0.9854867),
	53:   float32(0.98492591),
	54:   float32(0.98435457),
	55:   float32(0.98377269),
	56:   float32(0.98318028),
	57:   float32(0.98257734),
	58:   float32(0.98196387),
	59:   float32(0.98133989),
	60:   float32(0.98070539),
	61:   float32(0.9800604),
	62:   float32(0.97940491),
	63:   float32(0.97873893),
	64:   float32(0.97806247),
	65:   float32(0.97737554),
	66:   float32(0.97667813),
	67:   float32(0.97597027),
	68:   float32(0.97525196),
	69:   float32(0.9745232),
	70:   float32(0.97378401),
	71:   float32(0.97303439),
	72:   float32(0.97227435),
	73:   float32(0.97150389),
	74:   float32(0.97072303),
	75:   float32(0.96993178),
	76:   float32(0.96913014),
	77:   float32(0.96831812),
	78:   float32(0.96749573),
	79:   float32(0.96666298),
	80:   float32(0.96581987),
	81:   float32(0.96496643),
	82:   float32(0.96410265),
	83:   float32(0.96322854),
	84:   float32(0.96234412),
	85:   float32(0.96144939),
	86:   float32(0.96054437),
	87:   float32(0.95962906),
	88:   float32(0.95870347),
	89:   float32(0.95776762),
	90:   float32(0.95682151),
	91:   float32(0.95586515),
	92:   float32(0.95489856),
	93:   float32(0.95392174),
	94:   float32(0.95293471),
	95:   float32(0.95193746),
	96:   float32(0.95093003),
	97:   float32(0.94991241),
	98:   float32(0.94888462),
	99:   float32(0.94784667),
	100:  float32(0.94679856),
	101:  float32(0.94574032),
	102:  float32(0.94467195),
	103:  float32(0.94359346),
	104:  float32(0.94250486),
	105:  float32(0.94140618),
	106:  float32(0.94029741),
	107:  float32(0.93917857),
	108:  float32(0.93804967),
	109:  float32(0.93691073),
	110:  float32(0.93576176),
	111:  float32(0.93460276),
	112:  float32(0.93343375),
	113:  float32(0.93225475),
	114:  float32(0.93106577),
	115:  float32(0.92986681),
	116:  float32(0.92865789),
	117:  float32(0.92743903),
	118:  float32(0.92621024),
	119:  float32(0.92497153),
	120:  float32(0.92372291),
	121:  float32(0.9224644),
	122:  float32(0.92119602),
	123:  float32(0.91991776),
	124:  float32(0.91862966),
	125:  float32(0.91733172),
	126:  float32(0.91602395),
	127:  float32(0.91470637),
	128:  float32(0.913379),
	129:  float32(0.91204185),
	130:  float32(0.91069493),
	131:  float32(0.90933825),
	132:  float32(0.90797184),
	133:  float32(0.9065957),
	134:  float32(0.90520986),
	135:  float32(0.90381432),
	136:  float32(0.9024091),
	137:  float32(0.90099422),
	138:  float32(0.89956969),
	139:  float32(0.89813553),
	140:  float32(0.89669174),
	141:  float32(0.89523836),
	142:  float32(0.89377538),
	143:  float32(0.89230284),
	144:  float32(0.89082074),
	145:  float32(0.8893291),
	146:  float32(0.88782793),
	147:  float32(0.88631726),
	148:  float32(0.8847971),
	149:  float32(0.88326746),
	150:  float32(0.88172836),
	151:  float32(0.88017982),
	152:  float32(0.87862185),
	153:  float32(0.87705448),
	154:  float32(0.87547771),
	155:  float32(0.87389156),
	156:  float32(0.87229606),
	157:  float32(0.87069121),
	158:  float32(0.86907704),
	159:  float32(0.86745357),
	160:  float32(0.8658208),
	161:  float32(0.86417876),
	162:  float32(0.86252747),
	163:  float32(0.86086694),
	164:  float32(0.85919719),
	165:  float32(0.85751824),
	166:  float32(0.8558301),
	167:  float32(0.85413281),
	168:  float32(0.85242636),
	169:  float32(0.85071078),
	170:  float32(0.8489861),
	171:  float32(0.84725232),
	172:  float32(0.84550947),
	173:  float32(0.84375756),
	174:  float32(0.84199662),
	175:  float32(0.84022666),
	176:  float32(0.83844771),
	177:  float32(0.83665977),
	178:  float32(0.83486287),
	179:  float32(0.83305704),
	180:  float32(0.83124228),
	181:  float32(0.82941862),
	182:  float32(0.82758608),
	183:  float32(0.82574467),
	184:  float32(0.82389442),
	185:  float32(0.82203535),
	186:  float32(0.82016748),
	187:  float32(0.81829082),
	188:  float32(0.8164054),
	189:  float32(0.81451123),
	190:  float32(0.81260835),
	191:  float32(0.81069676),
	192:  float32(0.80877649),
	193:  float32(0.80684755),
	194:  float32(0.80490998),
	195:  float32(0.80296379),
	196:  float32(0.80100899),
	197:  float32(0.79904562),
	198:  float32(0.7970737),
	199:  float32(0.79509323),
	200:  float32(0.79310425),
	201:  float32(0.79110678),
	202:  float32(0.78910084),
	203:  float32(0.78708644),
	204:  float32(0.78506362),
	205:  float32(0.78303239),
	206:  float32(0.78099277),
	207:  float32(0.77894479),
	208:  float32(0.77688847),
	209:  float32(0.77482382),
	210:  float32(0.77275088),
	211:  float32(0.77066967),
	212:  float32(0.7685802),
	213:  float32(0.7664825),
	214:  float32(0.76437659),
	215:  float32(0.7622625),
	216:  float32(0.76014024),
	217:  float32(0.75800984),
	218:  float32(0.75587132),
	219:  float32(0.75372471),
	220:  float32(0.75157003),
	221:  float32(0.7494073),
	222:  float32(0.74723654),
	223:  float32(0.74505779),
	224:  float32(0.74287105),
	225:  float32(0.74067635),
	226:  float32(0.73847373),
	227:  float32(0.7362632),
	228:  float32(0.73404478),
	229:  float32(0.7318185),
	230:  float32(0.72958438),
	231:  float32(0.72734245),
	232:  float32(0.72509273),
	233:  float32(0.72283525),
	234:  float32(0.72057002),
	235:  float32(0.71829708),
	236:  float32(0.71601644),
	237:  float32(0.71372814),
	238:  float32(0.7114322),
	239:  float32(0.70912863),
	240:  float32(0.70681747),
	241:  float32(0.70449874),
	242:  float32(0.70217247),
	243:  float32(0.69983868),
	244:  float32(0.69749739),
	245:  float32(0.69514863),
	246:  float32(0.69279243),
	247:  float32(0.69042881),
	248:  float32(0.6880578),
	249:  float32(0.68567941),
	250:  float32(0.68329369),
	251:  float32(0.68090064),
	252:  float32(0.67850031),
	253:  float32(0.6760927),
	254:  float32(0.67367786),
	255:  float32(0.6712558),
	256:  float32(0.66882656),
	257:  float32(0.66639015),
	258:  float32(0.66394661),
	259:  float32(0.66149595),
	260:  float32(0.65903821),
	261:  float32(0.65657341),
	262:  float32(0.65410159),
	263:  float32(0.65162275),
	264:  float32(0.64913694),
	265:  float32(0.64664418),
	266:  float32(0.64414449),
	267:  float32(0.6416379),
	268:  float32(0.63912444),
	269:  float32(0.63660414),
	270:  float32(0.63407702),
	271:  float32(0.63154311),
	272:  float32(0.62900244),
	273:  float32(0.62645503),
	274:  float32(0.62390091),
	275:  float32(0.62134011),
	276:  float32(0.61877265),
	277:  float32(0.61619857),
	278:  float32(0.61361789),
	279:  float32(0.61103064),
	280:  float32(0.60843685),
	281:  float32(0.60583654),
	282:  float32(0.60322974),
	283:  float32(0.60061648),
	284:  float32(0.59799679),
	285:  float32(0.59537069),
	286:  float32(0.59273822),
	287:  float32(0.5900994),
	288:  float32(0.58745427),
	289:  float32(0.58480284),
	290:  float32(0.58214514),
	291:  float32(0.57948122),
	292:  float32(0.57681109),
	293:  float32(0.57413478),
	294:  float32(0.57145232),
	295:  float32(0.56876374),
	296:  float32(0.56606907),
	297:  float32(0.56336834),
	298:  float32(0.56066158),
	299:  float32(0.55794881),
	300:  float32(0.55523006),
	301:  float32(0.55250537),
	302:  float32(0.54977477),
	303:  float32(0.54703827),
	304:  float32(0.54429592),
	305:  float32(0.54154774),
	306:  float32(0.53879376),
	307:  float32(0.53603401),
	308:  float32(0.53326852),
	309:  float32(0.53049731),
	310:  float32(0.52772043),
	311:  float32(0.5249379),
	312:  float32(0.52214974),
	313:  float32(0.51935599),
	314:  float32(0.51655668),
	315:  float32(0.51375184),
	316:  float32(0.51094149),
	317:  float32(0.50812568),
	318:  float32(0.50530442),
	319:  float32(0.50247775),
	320:  float32(0.4996457),
	321:  float32(0.4968083),
	322:  float32(0.49396558),
	323:  float32(0.49111757),
	324:  float32(0.4882643),
	325:  float32(0.4854058),
	326:  float32(0.4825421),
	327:  float32(0.47967323),
	328:  float32(0.47679923),
	329:  float32(0.47392012),
	330:  float32(0.47103594),
	331:  float32(0.46814671),
	332:  float32(0.46525247),
	333:  float32(0.46235324),
	334:  float32(0.45944907),
	335:  float32(0.45653997),
	336:  float32(0.45362599),
	337:  float32(0.45070714),
	338:  float32(0.44778347),
	339:  float32(0.44485501),
	340:  float32(0.44192178),
	341:  float32(0.43898381),
	342:  float32(0.43604115),
	343:  float32(0.43309382),
	344:  float32(0.43014185),
	345:  float32(0.42718527),
	346:  float32(0.42422412),
	347:  float32(0.42125842),
	348:  float32(0.41828822),
	349:  float32(0.41531353),
	350:  float32(0.4123344),
	351:  float32(0.40935085),
	352:  float32(0.40636291),
	353:  float32(0.40337063),
	354:  float32(0.40037402),
	355:  float32(0.39737313),
	356:  float32(0.39436798),
	357:  float32(0.39135861),
	358:  float32(0.38834505),
	359:  float32(0.38532733),
	360:  float32(0.38230548),
	361:  float32(0.37927953),
	362:  float32(0.37624953),
	363:  float32(0.3732155),
	364:  float32(0.37017747),
	365:  float32(0.36713547),
	366:  float32(0.36408955),
	367:  float32(0.36103972),
	368:  float32(0.35798603),
	369:  float32(0.3549285),
	370:  float32(0.35186718),
	371:  float32(0.34880208),
	372:  float32(0.34573325),
	373:  float32(0.34266072),
	374:  float32(0.33958451),
	375:  float32(0.33650468),
	376:  float32(0.33342123),
	377:  float32(0.33033422),
	378:  float32(0.32724367),
	379:  float32(0.32414961),
	380:  float32(0.32105209),
	381:  float32(0.31795112),
	382:  float32(0.31484675),
	383:  float32(0.31173901),
	384:  float32(0.30862793),
	385:  float32(0.30551354),
	386:  float32(0.30239588),
	387:  float32(0.29927499),
	388:  float32(0.29615089),
	389:  float32(0.29302362),
	390:  float32(0.28989321),
	391:  float32(0.28675969),
	392:  float32(0.2836231),
	393:  float32(0.28048348),
	394:  float32(0.27734085),
	395:  float32(0.27419526),
	396:  float32(0.27104672),
	397:  float32(0.26789529),
	398:  float32(0.26474098),
	399:  float32(0.26158384),
	400:  float32(0.2584239),
	401:  float32(0.25526119),
	402:  float32(0.25209575),
	403:  float32(0.24892761),
	404:  float32(0.2457568),
	405:  float32(0.24258336),
	406:  float32(0.23940732),
	407:  float32(0.23622872),
	408:  float32(0.23304759),
	409:  float32(0.22986396),
	410:  float32(0.22667787),
	411:  float32(0.22348935),
	412:  float32(0.22029844),
	413:  float32(0.21710517),
	414:  float32(0.21390958),
	415:  float32(0.21071169),
	416:  float32(0.20751155),
	417:  float32(0.20430919),
	418:  float32(0.20110463),
	419:  float32(0.19789793),
	420:  float32(0.1946891),
	421:  float32(0.19147819),
	422:  float32(0.18826523),
	423:  float32(0.18505026),
	424:  float32(0.1818333),
	425:  float32(0.17861439),
	426:  float32(0.17539357),
	427:  float32(0.17217088),
	428:  float32(0.16894634),
	429:  float32(0.16571999),
	430:  float32(0.16249186),
	431:  float32(0.159262),
	432:  float32(0.15603043),
	433:  float32(0.15279719),
	434:  float32(0.14956231),
	435:  float32(0.14632583),
	436:  float32(0.14308778),
	437:  float32(0.1398482),
	438:  float32(0.13660713),
	439:  float32(0.13336459),
	440:  float32(0.13012062),
	441:  float32(0.12687526),
	442:  float32(0.12362854),
	443:  float32(0.12038049),
	444:  float32(0.11713116),
	445:  float32(0.11388057),
	446:  float32(0.11062877),
	447:  float32(0.10737578),
	448:  float32(0.10412163),
	449:  float32(0.10086638),
	450:  float32(0.09761004),
	451:  float32(0.094352658),
	452:  float32(0.091094266),
	453:  float32(0.087834897),
	454:  float32(0.084574589),
	455:  float32(0.081313374),
	456:  float32(0.078051289),
	457:  float32(0.074788367),
	458:  float32(0.071524645),
	459:  float32(0.068260157),
	460:  float32(0.064994938),
	461:  float32(0.061729023),
	462:  float32(0.058462447),
	463:  float32(0.055195244),
	464:  float32(0.051927451),
	465:  float32(0.048659101),
	466:  float32(0.045390231),
	467:  float32(0.042120874),
	468:  float32(0.038851066),
	469:  float32(0.035580842),
	470:  float32(0.032310238),
	471:  float32(0.029039287),
	472:  float32(0.025768025),
	473:  float32(0.022496487),
	474:  float32(0.019224708),
	475:  float32(0.015952723),
	476:  float32(0.012680568),
	477:  float32(0.0094082767),
	478:  float32(0.0061358846),
	479:  float32(0.0028634269),
	480:  -float32(0.00040906153),
	481:  -float32(0.0036815456),
	482:  -float32(0.0069539902),
	483:  -float32(0.01022636),
	484:  -float32(0.013498621),
	485:  -float32(0.016770737),
	486:  -float32(0.020042673),
	487:  -float32(0.023314395),
	488:  -float32(0.026585867),
	489:  -float32(0.029857055),
	490:  -float32(0.033127923),
	491:  -float32(0.036398436),
	492:  -float32(0.039668559),
	493:  -float32(0.042938257),
	494:  -float32(0.046207495),
	495:  -float32(0.049476239),
	496:  -float32(0.052744453),
	497:  -float32(0.056012102),
	498:  -float32(0.059279151),
	499:  -float32(0.062545565),
	500:  -float32(0.065811309),
	501:  -float32(0.069076349),
	502:  -float32(0.072340649),
	503:  -float32(0.075604174),
	504:  -float32(0.07886689),
	505:  -float32(0.082128761),
	506:  -float32(0.085389752),
	507:  -float32(0.088649829),
	508:  -float32(0.091908956),
	509:  -float32(0.0951671),
	510:  -float32(0.098424224),
	511:  -float32(0.10168029),
	512:  -float32(0.10493528),
	513:  -float32(0.10818913),
	514:  -float32(0.11144183),
	515:  -float32(0.11469334),
	516:  -float32(0.11794361),
	517:  -float32(0.12119263),
	518:  -float32(0.12444034),
	519:  -float32(0.12768673),
	520:  -float32(0.13093174),
	521:  -float32(0.13417536),
	522:  -float32(0.13741753),
	523:  -float32(0.14065824),
	524:  -float32(0.14389744),
	525:  -float32(0.1471351),
	526:  -float32(0.15037118),
	527:  -float32(0.15360565),
	528:  -float32(0.15683848),
	529:  -float32(0.16006962),
	530:  -float32(0.16329906),
	531:  -float32(0.16652674),
	532:  -float32(0.16975264),
	533:  -float32(0.17297673),
	534:  -float32(0.17619896),
	535:  -float32(0.1794193),
	536:  -float32(0.18263772),
	537:  -float32(0.18585419),
	538:  -float32(0.18906866),
	539:  -float32(0.19228112),
	540:  -float32(0.19549151),
	541:  -float32(0.19869981),
	542:  -float32(0.20190598),
	543:  -float32(0.20510998),
	544:  -float32(0.2083118),
	545:  -float32(0.21151138),
	546:  -float32(0.21470869),
	547:  -float32(0.21790371),
	548:  -float32(0.22109639),
	549:  -float32(0.22428671),
	550:  -float32(0.22747462),
	551:  -float32(0.2306601),
	552:  -float32(0.2338431),
	553:  -float32(0.23702361),
	554:  -float32(0.24020157),
	555:  -float32(0.24337696),
	556:  -float32(0.24654975),
	557:  -float32(0.24971989),
	558:  -float32(0.25288736),
	559:  -float32(0.25605213),
	560:  -float32(0.25921415),
	561:  -float32(0.26237339),
	562:  -float32(0.26552983),
	563:  -float32(0.26868342),
	564:  -float32(0.27183413),
	565:  -float32(0.27498193),
	566:  -float32(0.27812679),
	567:  -float32(0.28126867),
	568:  -float32(0.28440754),
	569:  -float32(0.28754336),
	570:  -float32(0.2906761),
	571:  -float32(0.29380573),
	572:  -float32(0.29693221),
	573:  -float32(0.30005551),
	574:  -float32(0.3031756),
	575:  -float32(0.30629245),
	576:  -float32(0.30940601),
	577:  -float32(0.31251626),
	578:  -float32(0.31562316),
	579:  -float32(0.31872668),
	580:  -float32(0.32182679),
	581:  -float32(0.32492345),
	582:  -float32(0.32801664),
	583:  -float32(0.33110631),
	584:  -float32(0.33419243),
	585:  -float32(0.33727497),
	586:  -float32(0.34035391),
	587:  -float32(0.3434292),
	588:  -float32(0.34650081),
	589:  -float32(0.34956871),
	590:  -float32(0.35263286),
	591:  -float32(0.35569324),
	592:  -float32(0.35874981),
	593:  -float32(0.36180254),
	594:  -float32(0.36485139),
	595:  -float32(0.36789634),
	596:  -float32(0.37093735),
	597:  -float32(0.37397438),
	598:  -float32(0.37700741),
	599:  -float32(0.3800364),
	600:  -float32(0.38306132),
	601:  -float32(0.38608214),
	602:  -float32(0.38909883),
	603:  -float32(0.39211135),
	604:  -float32(0.39511967),
	605:  -float32(0.39812375),
	606:  -float32(0.40112358),
	607:  -float32(0.4041191),
	608:  -float32(0.40711031),
	609:  -float32(0.41009715),
	610:  -float32(0.41307959),
	611:  -float32(0.41605762),
	612:  -float32(0.41903119),
	613:  -float32(0.42200027),
	614:  -float32(0.42496483),
	615:  -float32(0.42792484),
	616:  -float32(0.43088027),
	617:  -float32(0.43383109),
	618:  -float32(0.43677726),
	619:  -float32(0.43971875),
	620:  -float32(0.44265553),
	621:  -float32(0.44558757),
	622:  -float32(0.44851484),
	623:  -float32(0.45143731),
	624:  -float32(0.45435494),
	625:  -float32(0.4572677),
	626:  -float32(0.46017557),
	627:  -float32(0.46307851),
	628:  -float32(0.4659765),
	629:  -float32(0.46886949),
	630:  -float32(0.47175746),
	631:  -float32(0.47464038),
	632:  -float32(0.47751821),
	633:  -float32(0.48039093),
	634:  -float32(0.48325851),
	635:  -float32(0.48612091),
	636:  -float32(0.48897811),
	637:  -float32(0.49183006),
	638:  -float32(0.49467676),
	639:  -float32(0.49751815),
	640:  -float32(0.50035422),
	641:  -float32(0.50318492),
	642:  -float32(0.50601024),
	643:  -float32(0.50883014),
	644:  -float32(0.51164459),
	645:  -float32(0.51445356),
	646:  -float32(0.51725703),
	647:  -float32(0.52005495),
	648:  -float32(0.5228473),
	649:  -float32(0.52563406),
	650:  -float32(0.52841518),
	651:  -float32(0.53119065),
	652:  -float32(0.53396043),
	653:  -float32(0.53672449),
	654:  -float32(0.5394828),
	655:  -float32(0.54223533),
	656:  -float32(0.54498206),
	657:  -float32(0.54772295),
	658:  -float32(0.55045797),
	659:  -float32(0.5531871),
	660:  -float32(0.55591031),
	661:  -float32(0.55862756),
	662:  -float32(0.56133883),
	663:  -float32(0.56404409),
	664:  -float32(0.56674331),
	665:  -float32(0.56943646),
	666:  -float32(0.57212351),
	667:  -float32(0.57480443),
	668:  -float32(0.5774792),
	669:  -float32(0.58014778),
	670:  -float32(0.58281015),
	671:  -float32(0.58546628),
	672:  -float32(0.58811614),
	673:  -float32(0.5907597),
	674:  -float32(0.59339694),
	675:  -float32(0.59602782),
	676:  -float32(0.59865231),
	677:  -float32(0.6012704),
	678:  -float32(0.60388204),
	679:  -float32(0.60648722),
	680:  -float32(0.60908591),
	681:  -float32(0.61167807),
	682:  -float32(0.61426368),
	683:  -float32(0.61684271),
	684:  -float32(0.61941514),
	685:  -float32(0.62198093),
	686:  -float32(0.62454007),
	687:  -float32(0.62709251),
	688:  -float32(0.62963824),
	689:  -float32(0.63217722),
	690:  -float32(0.63470944),
	691:  -float32(0.63723486),
	692:  -float32(0.63975345),
	693:  -float32(0.64226519),
	694:  -float32(0.64477006),
	695:  -float32(0.64726802),
	696:  -float32(0.64975905),
	697:  -float32(0.65224312),
	698:  -float32(0.6547202),
	699:  -float32(0.65719027),
	700:  -float32(0.65965331),
	701:  -float32(0.66210928),
	702:  -float32(0.66455816),
	703:  -float32(0.66699992),
	704:  -float32(0.66943454),
	705:  -float32(0.67186199),
	706:  -float32(0.67428225),
	707:  -float32(0.67669528),
	708:  -float32(0.67910107),
	709:  -float32(0.68149959),
	710:  -float32(0.6838908),
	711:  -float32(0.6862747),
	712:  -float32(0.68865124),
	713:  -float32(0.69102041),
	714:  -float32(0.69338218),
	715:  -float32(0.69573652),
	716:  -float32(0.69808341),
	717:  -float32(0.70042283),
	718:  -float32(0.70275474),
	719:  -float32(0.70507913),
	720:  -float32(0.70739597),
	721:  -float32(0.70970524),
	722:  -float32(0.7120069),
	723:  -float32(0.71430093),
	724:  -float32(0.71658732),
	725:  -float32(0.71886604),
	726:  -float32(0.72113705),
	727:  -float32(0.72340034),
	728:  -float32(0.72565589),
	729:  -float32(0.72790366),
	730:  -float32(0.73014364),
	731:  -float32(0.7323758),
	732:  -float32(0.73460012),
	733:  -float32(0.73681657),
	734:  -float32(0.73902513),
	735:  -float32(0.74122577),
	736:  -float32(0.74341848),
	737:  -float32(0.74560322),
	738:  -float32(0.74777998),
	739:  -float32(0.74994874),
	740:  -float32(0.75210946),
	741:  -float32(0.75426212),
	742:  -float32(0.75640671),
	743:  -float32(0.7585432),
	744:  -float32(0.76067157),
	745:  -float32(0.76279178),
	746:  -float32(0.76490383),
	747:  -float32(0.76700769),
	748:  -float32(0.76910334),
	749:  -float32(0.77119075),
	750:  -float32(0.77326989),
	751:  -float32(0.77534076),
	752:  -float32(0.77740333),
	753:  -float32(0.77945757),
	754:  -float32(0.78150346),
	755:  -float32(0.78354098),
	756:  -float32(0.78557011),
	757:  -float32(0.78759083),
	758:  -float32(0.78960312),
	759:  -float32(0.79160694),
	760:  -float32(0.79360229),
	761:  -float32(0.79558915),
	762:  -float32(0.79756748),
	763:  -float32(0.79953727),
	764:  -float32(0.8014985),
	765:  -float32(0.80345114),
	766:  -float32(0.80539518),
	767:  -float32(0.8073306),
	768:  -float32(0.80925737),
	769:  -float32(0.81117547),
	770:  -float32(0.81308489),
	771:  -float32(0.81498559),
	772:  -float32(0.81687757),
	773:  -float32(0.81876081),
	774:  -float32(0.82063527),
	775:  -float32(0.82250095),
	776:  -float32(0.82435781),
	777:  -float32(0.82620585),
	778:  -float32(0.82804505),
	779:  -float32(0.82987537),
	780:  -float32(0.83169681),
	781:  -float32(0.83350933),
	782:  -float32(0.83531294),
	783:  -float32(0.8371076),
	784:  -float32(0.83889329),
	785:  -float32(0.84067),
	786:  -float32(0.8424377),
	787:  -float32(0.84419639),
	788:  -float32(0.84594603),
	789:  -float32(0.84768662),
	790:  -float32(0.84941812),
	791:  -float32(0.85114053),
	792:  -float32(0.85285383),
	793:  -float32(0.85455799),
	794:  -float32(0.856253),
	795:  -float32(0.85793884),
	796:  -float32(0.85961549),
	797:  -float32(0.86128294),
	798:  -float32(0.86294116),
	799:  -float32(0.86459014),
	800:  -float32(0.86622986),
	801:  -float32(0.86786031),
	802:  -float32(0.86948146),
	803:  -float32(0.8710933),
	804:  -float32(0.87269581),
	805:  -float32(0.87428898),
	806:  -float32(0.87587278),
	807:  -float32(0.8774472),
	808:  -float32(0.87901223),
	809:  -float32(0.88056784),
	810:  -float32(0.88211402),
	811:  -float32(0.88365076),
	812:  -float32(0.88517803),
	813:  -float32(0.88669582),
	814:  -float32(0.88820412),
	815:  -float32(0.8897029),
	816:  -float32(0.89119216),
	817:  -float32(0.89267187),
	818:  -float32(0.89414203),
	819:  -float32(0.8956026),
	820:  -float32(0.89705359),
	821:  -float32(0.89849497),
	822:  -float32(0.89992673),
	823:  -float32(0.90134885),
	824:  -float32(0.90276131),
	825:  -float32(0.90416411),
	826:  -float32(0.90555723),
	827:  -float32(0.90694065),
	828:  -float32(0.90831436),
	829:  -float32(0.90967833),
	830:  -float32(0.91103257),
	831:  -float32(0.91237705),
	832:  -float32(0.91371176),
	833:  -float32(0.91503669),
	834:  -float32(0.91635181),
	835:  -float32(0.91765712),
	836:  -float32(0.91895261),
	837:  -float32(0.92023825),
	838:  -float32(0.92151404),
	839:  -float32(0.92277996),
	840:  -float32(0.924036),
	841:  -float32(0.92528214),
	842:  -float32(0.92651837),
	843:  -float32(0.92774468),
	844:  -float32(0.92896106),
	845:  -float32(0.93016748),
	846:  -float32(0.93136395),
	847:  -float32(0.93255044),
	848:  -float32(0.93372694),
	849:  -float32(0.93489345),
	850:  -float32(0.93604994),
	851:  -float32(0.93719641),
	852:  -float32(0.93833284),
	853:  -float32(0.93945922),
	854:  -float32(0.94057555),
	855:  -float32(0.9416818),
	856:  -float32(0.94277796),
	857:  -float32(0.94386403),
	858:  -float32(0.94493999),
	859:  -float32(0.94600583),
	860:  -float32(0.94706154),
	861:  -float32(0.94810711),
	862:  -float32(0.94914252),
	863:  -float32(0.95016777),
	864:  -float32(0.95118284),
	865:  -float32(0.95218773),
	866:  -float32(0.95318242),
	867:  -float32(0.9541669),
	868:  -float32(0.95514117),
	869:  -float32(0.9561052),
	870:  -float32(0.957059),
	871:  -float32(0.95800255),
	872:  -float32(0.95893583),
	873:  -float32(0.95985885),
	874:  -float32(0.96077159),
	875:  -float32(0.96167404),
	876:  -float32(0.96256619),
	877:  -float32(0.96344803),
	878:  -float32(0.96431956),
	879:  -float32(0.96518076),
	880:  -float32(0.96603162),
	881:  -float32(0.96687213),
	882:  -float32(0.9677023),
	883:  -float32(0.96852209),
	884:  -float32(0.96933152),
	885:  -float32(0.97013057),
	886:  -float32(0.97091922),
	887:  -float32(0.97169748),
	888:  -float32(0.97246533),
	889:  -float32(0.97322277),
	890:  -float32(0.97396979),
	891:  -float32(0.97470637),
	892:  -float32(0.97543252),
	893:  -float32(0.97614822),
	894:  -float32(0.97685347),
	895:  -float32(0.97754825),
	896:  -float32(0.97823257),
	897:  -float32(0.97890641),
	898:  -float32(0.97956977),
	899:  -float32(0.98022263),
	900:  -float32(0.980865),
	901:  -float32(0.98149687),
	902:  -float32(0.98211822),
	903:  -float32(0.98272906),
	904:  -float32(0.98332937),
	905:  -float32(0.98391915),
	906:  -float32(0.9844984),
	907:  -float32(0.9850671),
	908:  -float32(0.98562525),
	909:  -float32(0.98617285),
	910:  -float32(0.98670988),
	911:  -float32(0.98723635),
	912:  -float32(0.98775225),
	913:  -float32(0.98825757),
	914:  -float32(0.9887523),
	915:  -float32(0.98923645),
	916:  -float32(0.98971),
	917:  -float32(0.99017295),
	918:  -float32(0.9906253),
	919:  -float32(0.99106704),
	920:  -float32(0.99149817),
	921:  -float32(0.99191868),
	922:  -float32(0.99232857),
	923:  -float32(0.99272783),
	924:  -float32(0.99311645),
	925:  -float32(0.99349445),
	926:  -float32(0.9938618),
	927:  -float32(0.99421851),
	928:  -float32(0.99456457),
	929:  -float32(0.99489998),
	930:  -float32(0.99522474),
	931:  -float32(0.99553884),
	932:  -float32(0.99584227),
	933:  -float32(0.99613505),
	934:  -float32(0.99641715),
	935:  -float32(0.99668858),
	936:  -float32(0.99694934),
	937:  -float32(0.99719943),
	938:  -float32(0.99743883),
	939:  -float32(0.99766756),
	940:  -float32(0.99788559),
	941:  -float32(0.99809295),
	942:  -float32(0.99828961),
	943:  -float32(0.99847558),
	944:  -float32(0.99865086),
	945:  -float32(0.99881544),
	946:  -float32(0.99896933),
	947:  -float32(0.99911252),
	948:  -float32(0.99924501),
	949:  -float32(0.9993668),
	950:  -float32(0.99947789),
	951:  -float32(0.99957827),
	952:  -float32(0.99966795),
	953:  -float32(0.99974692),
	954:  -float32(0.99981519),
	955:  -float32(0.99987275),
	956:  -float32(0.9999196),
	957:  -float32(0.99995574),
	958:  -float32(0.99998118),
	959:  -float32(0.9999959),
	960:  float32(0.99999967),
	961:  float32(0.99997289),
	962:  float32(0.99990328),
	963:  float32(0.99979084),
	964:  float32(0.99963557),
	965:  float32(0.99943748),
	966:  float32(0.99919658),
	967:  float32(0.99891288),
	968:  float32(0.99858638),
	969:  float32(0.99821711),
	970:  float32(0.99780508),
	971:  float32(0.99735031),
	972:  float32(0.99685281),
	973:  float32(0.99631261),
	974:  float32(0.99572973),
	975:  float32(0.9951042),
	976:  float32(0.99443605),
	977:  float32(0.99372529),
	978:  float32(0.99297196),
	979:  float32(0.9921761),
	980:  float32(0.99133774),
	981:  float32(0.99045692),
	982:  float32(0.98953366),
	983:  float32(0.98856802),
	984:  float32(0.98756003),
	985:  float32(0.98650973),
	986:  float32(0.98541718),
	987:  float32(0.98428242),
	988:  float32(0.98310549),
	989:  float32(0.98188645),
	990:  float32(0.98062534),
	991:  float32(0.97932224),
	992:  float32(0.97797718),
	993:  float32(0.97659022),
	994:  float32(0.97516144),
	995:  float32(0.97369088),
	996:  float32(0.97217861),
	997:  float32(0.97062469),
	998:  float32(0.9690292),
	999:  float32(0.9673922),
	1000: float32(0.96571376),
	1001: float32(0.96399395),
	1002: float32(0.96223284),
	1003: float32(0.96043052),
	1004: float32(0.95858705),
	1005: float32(0.95670253),
	1006: float32(0.95477702),
	1007: float32(0.95281061),
	1008: float32(0.95080338),
	1009: float32(0.94875543),
	1010: float32(0.94666684),
	1011: float32(0.94453769),
	1012: float32(0.94236808),
	1013: float32(0.9401581),
	1014: float32(0.93790786),
	1015: float32(0.93561743),
	1016: float32(0.93328693),
	1017: float32(0.93091644),
	1018: float32(0.92850608),
	1019: float32(0.92605595),
	1020: float32(0.92356614),
	1021: float32(0.92103677),
	1022: float32(0.91846795),
	1023: float32(0.91585979),
	1024: float32(0.91321239),
	1025: float32(0.91052587),
	1026: float32(0.90780035),
	1027: float32(0.90503595),
	1028: float32(0.90223277),
	1029: float32(0.89939095),
	1030: float32(0.89651059),
	1031: float32(0.89359184),
	1032: float32(0.89063481),
	1033: float32(0.88763962),
	1034: float32(0.88460641),
	1035: float32(0.88153531),
	1036: float32(0.87842644),
	1037: float32(0.87527995),
	1038: float32(0.87209596),
	1039: float32(0.86887462),
	1040: float32(0.86561605),
	1041: float32(0.86232041),
	1042: float32(0.85898782),
	1043: float32(0.85561844),
	1044: float32(0.85221241),
	1045: float32(0.84876987),
	1046: float32(0.84529098),
	1047: float32(0.84177587),
	1048: float32(0.83822471),
	1049: float32(0.83463763),
	1050: float32(0.83101481),
	1051: float32(0.82735639),
	1052: float32(0.82366252),
	1053: float32(0.81993338),
	1054: float32(0.81616911),
	1055: float32(0.81236987),
	1056: float32(0.80853584),
	1057: float32(0.80466718),
	1058: float32(0.80076404),
	1059: float32(0.7968266),
	1060: float32(0.79285503),
	1061: float32(0.7888495),
	1062: float32(0.78481017),
	1063: float32(0.78073723),
	1064: float32(0.77663084),
	1065: float32(0.77249118),
	1066: float32(0.76831844),
	1067: float32(0.76411277),
	1068: float32(0.75987438),
	1069: float32(0.75560344),
	1070: float32(0.75130013),
	1071: float32(0.74696464),
	1072: float32(0.74259715),
	1073: float32(0.73819784),
	1074: float32(0.73376692),
	1075: float32(0.72930457),
	1076: float32(0.72481097),
	1077: float32(0.72028632),
	1078: float32(0.71573083),
	1079: float32(0.71114467),
	1080: float32(0.70652804),
	1081: float32(0.70188116),
	1082: float32(0.6972042),
	1083: float32(0.69249738),
	1084: float32(0.6877609),
	1085: float32(0.68299495),
	1086: float32(0.67819975),
	1087: float32(0.6733755),
	1088: float32(0.6685224),
	1089: float32(0.66364066),
	1090: float32(0.6587305),
	1091: float32(0.65379211),
	1092: float32(0.64882573),
	1093: float32(0.64383154),
	1094: float32(0.63880978),
	1095: float32(0.63376065),
	1096: float32(0.62868438),
	1097: float32(0.62358117),
	1098: float32(0.61845126),
	1099: float32(0.61329485),
	1100: float32(0.60811216),
	1101: float32(0.60290343),
	1102: float32(0.59766888),
	1103: float32(0.59240872),
	1104: float32(0.58712318),
	1105: float32(0.58181249),
	1106: float32(0.57647688),
	1107: float32(0.57111658),
	1108: float32(0.56573181),
	1109: float32(0.56032281),
	1110: float32(0.5548898),
	1111: float32(0.54943303),
	1112: float32(0.54395272),
	1113: float32(0.53844911),
	1114: float32(0.53292243),
	1115: float32(0.52737292),
	1116: float32(0.52180083),
	1117: float32(0.51620638),
	1118: float32(0.51058981),
	1119: float32(0.50495138),
	1120: float32(0.49929132),
	1121: float32(0.49360987),
	1122: float32(0.48790727),
	1123: float32(0.48218377),
	1124: float32(0.47643962),
	1125: float32(0.47067506),
	1126: float32(0.46489034),
	1127: float32(0.4590857),
	1128: float32(0.45326139),
	1129: float32(0.44741768),
	1130: float32(0.44155479),
	1131: float32(0.43567299),
	1132: float32(0.42977253),
	1133: float32(0.42385365),
	1134: float32(0.41791662),
	1135: float32(0.41196169),
	1136: float32(0.40598911),
	1137: float32(0.39999914),
	1138: float32(0.39399204),
	1139: float32(0.38796806),
	1140: float32(0.38192746),
	1141: float32(0.3758705),
	1142: float32(0.36979743),
	1143: float32(0.36370853),
	1144: float32(0.35760405),
	1145: float32(0.35148424),
	1146: float32(0.34534939),
	1147: float32(0.33919973),
	1148: float32(0.33303555),
	1149: float32(0.3268571),
	1150: float32(0.32066465),
	1151: float32(0.31445847),
	1152: float32(0.30823881),
	1153: float32(0.30200595),
	1154: float32(0.29576015),
	1155: float32(0.28950169),
	1156: float32(0.28323082),
	1157: float32(0.27694782),
	1158: float32(0.27065295),
	1159: float32(0.26434649),
	1160: float32(0.25802871),
	1161: float32(0.25169988),
	1162: float32(0.24536026),
	1163: float32(0.23901013),
	1164: float32(0.23264977),
	1165: float32(0.22627944),
	1166: float32(0.21989941),
	1167: float32(0.21350997),
	1168: float32(0.20711138),
	1169: float32(0.20070391),
	1170: float32(0.19428785),
	1171: float32(0.18786347),
	1172: float32(0.18143104),
	1173: float32(0.17499084),
	1174: float32(0.16854314),
	1175: float32(0.16208822),
	1176: float32(0.15562636),
	1177: float32(0.14915783),
	1178: float32(0.14268292),
	1179: float32(0.13620189),
	1180: float32(0.12971502),
	1181: float32(0.1232226),
	1182: float32(0.11672491),
	1183: float32(0.11022221),
	1184: float32(0.10371479),
	1185: float32(0.097202924),
	1186: float32(0.090686897),
	1187: float32(0.084166986),
	1188: float32(0.077643468),
	1189: float32(0.071116625),
	1190: float32(0.064586736),
	1191: float32(0.05805408),
	1192: float32(0.051518937),
	1193: float32(0.044981587),
	1194: float32(0.03844231),
	1195: float32(0.031901387),
	1196: float32(0.025359097),
	1197: float32(0.018815721),
	1198: float32(0.012271538),
	1199: float32(0.0057268303),
	1200: -float32(0.000818123),
	1201: -float32(0.0073630412),
	1202: -float32(0.013907644),
	1203: -float32(0.020451651),
	1204: -float32(0.026994782),
	1205: -float32(0.033536757),
	1206: -float32(0.040077295),
	1207: -float32(0.046616116),
	1208: -float32(0.053152941),
	1209: -float32(0.059687488),
	1210: -float32(0.066219479),
	1211: -float32(0.072748633),
	1212: -float32(0.07927467),
	1213: -float32(0.085797312),
	1214: -float32(0.092316279),
	1215: -float32(0.098831291),
	1216: -float32(0.10534207),
	1217: -float32(0.11184834),
	1218: -float32(0.11834981),
	1219: -float32(0.12484622),
	1220: -float32(0.13133727),
	1221: -float32(0.1378227),
	1222: -float32(0.14430223),
	1223: -float32(0.15077558),
	1224: -float32(0.15724246),
	1225: -float32(0.16370261),
	1226: -float32(0.17015575),
	1227: -float32(0.1766016),
	1228: -float32(0.18303989),
	1229: -float32(0.18947033),
	1230: -float32(0.19589266),
	1231: -float32(0.2023066),
	1232: -float32(0.20871187),
	1233: -float32(0.2151082),
	1234: -float32(0.22149531),
	1235: -float32(0.22787294),
	1236: -float32(0.2342408),
	1237: -float32(0.24059864),
	1238: -float32(0.24694616),
	1239: -float32(0.25328311),
	1240: -float32(0.2596092),
	1241: -float32(0.26592418),
	1242: -float32(0.27222777),
	1243: -float32(0.27851969),
	1244: -float32(0.28479968),
	1245: -float32(0.29106748),
	1246: -float32(0.2973228),
	1247: -float32(0.30356539),
	1248: -float32(0.30979497),
	1249: -float32(0.31601129),
	1250: -float32(0.32221406),
	1251: -float32(0.32840304),
	1252: -float32(0.33457794),
	1253: -float32(0.34073852),
	1254: -float32(0.3468845),
	1255: -float32(0.35301562),
	1256: -float32(0.35913161),
	1257: -float32(0.36523223),
	1258: -float32(0.37131719),
	1259: -float32(0.37738626),
	1260: -float32(0.38343915),
	1261: -float32(0.38947562),
	1262: -float32(0.39549541),
	1263: -float32(0.40149825),
	1264: -float32(0.4074839),
	1265: -float32(0.41345209),
	1266: -float32(0.41940257),
	1267: -float32(0.42533508),
	1268: -float32(0.43124938),
	1269: -float32(0.4371452),
	1270: -float32(0.44302229),
	1271: -float32(0.44888041),
	1272: -float32(0.4547193),
	1273: -float32(0.46053871),
	1274: -float32(0.46633839),
	1275: -float32(0.4721181),
	1276: -float32(0.47787758),
	1277: -float32(0.48361659),
	1278: -float32(0.48933489),
	1279: -float32(0.49503222),
	1280: -float32(0.50070835),
	1281: -float32(0.50636303),
	1282: -float32(0.51199602),
	1283: -float32(0.51760707),
	1284: -float32(0.52319595),
	1285: -float32(0.52876243),
	1286: -float32(0.53430625),
	1287: -float32(0.53982718),
	1288: -float32(0.54532499),
	1289: -float32(0.55079944),
	1290: -float32(0.55625029),
	1291: -float32(0.56167732),
	1292: -float32(0.56708028),
	1293: -float32(0.57245896),
	1294: -float32(0.57781311),
	1295: -float32(0.58314251),
	1296: -float32(0.58844693),
	1297: -float32(0.59372614),
	1298: -float32(0.59897992),
	1299: -float32(0.60420805),
	1300: -float32(0.60941029),
	1301: -float32(0.61458642),
	1302: -float32(0.61973623),
	1303: -float32(0.62485949),
	1304: -float32(0.62995598),
	1305: -float32(0.63502549),
	1306: -float32(0.6400678),
	1307: -float32(0.64508268),
	1308: -float32(0.65006994),
	1309: -float32(0.65502934),
	1310: -float32(0.65996069),
	1311: -float32(0.66486377),
	1312: -float32(0.66973837),
	1313: -float32(0.67458427),
	1314: -float32(0.67940128),
	1315: -float32(0.68418919),
	1316: -float32(0.68894779),
	1317: -float32(0.69367688),
	1318: -float32(0.69837625),
	1319: -float32(0.70304571),
	1320: -float32(0.70768504),
	1321: -float32(0.71229407),
	1322: -float32(0.71687258),
	1323: -float32(0.72142039),
	1324: -float32(0.72593729),
	1325: -float32(0.73042309),
	1326: -float32(0.73487761),
	1327: -float32(0.73930064),
	1328: -float32(0.74369201),
	1329: -float32(0.74805152),
	1330: -float32(0.75237898),
	1331: -float32(0.75667422),
	1332: -float32(0.76093704),
	1333: -float32(0.76516727),
	1334: -float32(0.76936471),
	1335: -float32(0.77352921),
	1336: -float32(0.77766056),
	1337: -float32(0.78175861),
	1338: -float32(0.78582316),
	1339: -float32(0.78985406),
	1340: -float32(0.79385112),
	1341: -float32(0.79781417),
	1342: -float32(0.80174305),
	1343: -float32(0.80563758),
	1344: -float32(0.8094976),
	1345: -float32(0.81332295),
	1346: -float32(0.81711346),
	1347: -float32(0.82086896),
	1348: -float32(0.8245893),
	1349: -float32(0.82827432),
	1350: -float32(0.83192386),
	1351: -float32(0.83553776),
	1352: -float32(0.83911587),
	1353: -float32(0.84265803),
	1354: -float32(0.8461641),
	1355: -float32(0.84963392),
	1356: -float32(0.85306735),
	1357: -float32(0.85646423),
	1358: -float32(0.85982442),
	1359: -float32(0.86314779),
	1360: -float32(0.86643418),
	1361: -float32(0.86968345),
	1362: -float32(0.87289547),
	1363: -float32(0.87607009),
	1364: -float32(0.87920719),
	1365: -float32(0.88230663),
	1366: -float32(0.88536827),
	1367: -float32(0.88839199),
	1368: -float32(0.89137765),
	1369: -float32(0.89432512),
	1370: -float32(0.89723429),
	1371: -float32(0.90010502),
	1372: -float32(0.90293719),
	1373: -float32(0.90573069),
	1374: -float32(0.90848539),
	1375: -float32(0.91120117),
	1376: -float32(0.91387791),
	1377: -float32(0.91651551),
	1378: -float32(0.91911385),
	1379: -float32(0.92167282),
	1380: -float32(0.92419231),
	1381: -float32(0.9266722),
	1382: -float32(0.9291124),
	1383: -float32(0.9315128),
	1384: -float32(0.9338733),
	1385: -float32(0.9361938),
	1386: -float32(0.93847419),
	1387: -float32(0.94071438),
	1388: -float32(0.94291427),
	1389: -float32(0.94507377),
	1390: -float32(0.94719279),
	1391: -float32(0.94927123),
	1392: -float32(0.95130901),
	1393: -float32(0.95330604),
	1394: -float32(0.95526223),
	1395: -float32(0.9571775),
	1396: -float32(0.95905177),
	1397: -float32(0.96088496),
	1398: -float32(0.96267699),
	1399: -float32(0.96442777),
	1400: -float32(0.96613725),
	1401: -float32(0.96780534),
	1402: -float32(0.96943197),
	1403: -float32(0.97101707),
	1404: -float32(0.97256058),
	1405: -float32(0.97406243),
	1406: -float32(0.97552255),
	1407: -float32(0.97694089),
	1408: -float32(0.97831737),
	1409: -float32(0.97965195),
	1410: -float32(0.98094456),
	1411: -float32(0.98219515),
	1412: -float32(0.98340367),
	1413: -float32(0.98457006),
	1414: -float32(0.98569428),
	1415: -float32(0.98677627),
	1416: -float32(0.98781599),
	1417: -float32(0.9888134),
	1418: -float32(0.98976845),
	1419: -float32(0.9906811),
	1420: -float32(0.99155132),
	1421: -float32(0.99237906),
	1422: -float32(0.99316428),
	1423: -float32(0.99390697),
	1424: -float32(0.99460708),
	1425: -float32(0.99526458),
	1426: -float32(0.99587945),
	1427: -float32(0.99645166),
	1428: -float32(0.99698119),
	1429: -float32(0.99746801),
	1430: -float32(0.9979121),
	1431: -float32(0.99831344),
	1432: -float32(0.99867202),
	1433: -float32(0.99898782),
	1434: -float32(0.99926082),
	1435: -float32(0.99949102),
	1436: -float32(0.99967841),
	1437: -float32(0.99982297),
	1438: -float32(0.9999247),
	1439: -float32(0.9999836),
	1440: float32(0.99999866),
	1441: float32(0.99989157),
	1442: float32(0.99961315),
	1443: float32(0.99916346),
	1444: float32(0.99854256),
	1445: float32(0.99775057),
	1446: float32(0.99678762),
	1447: float32(0.99565388),
	1448: float32(0.99434953),
	1449: float32(0.99287481),
	1450: float32(0.99122996),
	1451: float32(0.98941527),
	1452: float32(0.98743105),
	1453: float32(0.98527764),
	1454: float32(0.98295541),
	1455: float32(0.98046475),
	1456: float32(0.9778061),
	1457: float32(0.9749799),
	1458: float32(0.97198664),
	1459: float32(0.96882685),
	1460: float32(0.96550104),
	1461: float32(0.9620098),
	1462: float32(0.95835373),
	1463: float32(0.95453345),
	1464: float32(0.95054962),
	1465: float32(0.94640291),
	1466: float32(0.94209404),
	1467: float32(0.93762375),
	1468: float32(0.9329928),
	1469: float32(0.92820199),
	1470: float32(0.92325213),
	1471: float32(0.91814408),
	1472: float32(0.91287871),
	1473: float32(0.90745693),
	1474: float32(0.90187965),
	1475: float32(0.89614785),
	1476: float32(0.89026249),
	1477: float32(0.88422459),
	1478: float32(0.87803519),
	1479: float32(0.87169533),
	1480: float32(0.86520612),
	1481: float32(0.85856866),
	1482: float32(0.85178409),
	1483: float32(0.84485357),
	1484: float32(0.83777828),
	1485: float32(0.83055945),
	1486: float32(0.82319831),
	1487: float32(0.81569611),
	1488: float32(0.80805415),
	1489: float32(0.80027373),
	1490: float32(0.7923562),
	1491: float32(0.78430289),
	1492: float32(0.7761152),
	1493: float32(0.76779452),
	1494: float32(0.75934229),
	1495: float32(0.75075995),
	1496: float32(0.74204897),
	1497: float32(0.73321084),
	1498: float32(0.72424708),
	1499: float32(0.71515923),
	1500: float32(0.70594883),
	1501: float32(0.69661748),
	1502: float32(0.68716676),
	1503: float32(0.6775983),
	1504: float32(0.66791374),
	1505: float32(0.65811474),
	1506: float32(0.64820297),
	1507: float32(0.63818013),
	1508: float32(0.62804795),
	1509: float32(0.61780815),
	1510: float32(0.60746249),
	1511: float32(0.59701275),
	1512: float32(0.58646072),
	1513: float32(0.57580819),
	1514: float32(0.56505701),
	1515: float32(0.554209),
	1516: float32(0.54326604),
	1517: float32(0.53222998),
	1518: float32(0.52110274),
	1519: float32(0.5098862),
	1520: float32(0.4985823),
	1521: float32(0.48719297),
	1522: float32(0.47572016),
	1523: float32(0.46416584),
	1524: float32(0.45253199),
	1525: float32(0.44082059),
	1526: float32(0.42903367),
	1527: float32(0.41717323),
	1528: float32(0.40524131),
	1529: float32(0.39323996),
	1530: float32(0.38117123),
	1531: float32(0.36903718),
	1532: float32(0.3568399),
	1533: float32(0.34458148),
	1534: float32(0.33226402),
	1535: float32(0.31988962),
	1536: float32(0.30746042),
	1537: float32(0.29497853),
	1538: float32(0.2824461),
	1539: float32(0.26986527),
	1540: float32(0.25723821),
	1541: float32(0.24456706),
	1542: float32(0.23185402),
	1543: float32(0.21910124),
	1544: float32(0.20631092),
	1545: float32(0.19348526),
	1546: float32(0.18062644),
	1547: float32(0.16773667),
	1548: float32(0.15481816),
	1549: float32(0.14187312),
	1550: float32(0.12890377),
	1551: float32(0.11591234),
	1552: float32(0.10290104),
	1553: float32(0.089872115),
	1554: float32(0.076827789),
	1555: float32(0.0637703),
	1556: float32(0.050701883),
	1557: float32(0.037624779),
	1558: float32(0.024541229),
	1559: float32(0.011453473),
	1560: -float32(0.0016362454),
	1561: -float32(0.014725683),
	1562: -float32(0.027812598),
	1563: -float32(0.040894747),
	1564: -float32(0.053969889),
	1565: -float32(0.067035784),
	1566: -float32(0.080090192),
	1567: -float32(0.093130877),
	1568: -float32(0.10615561),
	1569: -float32(0.11916214),
	1570: -float32(0.13214826),
	1571: -float32(0.14511174),
	1572: -float32(0.15805036),
	1573: -float32(0.17096189),
	1574: -float32(0.18384413),
	1575: -float32(0.19669487),
	1576: -float32(0.2095119),
	1577: -float32(0.22229304),
	1578: -float32(0.23503609),
	1579: -float32(0.24773886),
	1580: -float32(0.26039919),
	1581: -float32(0.2730149),
	1582: -float32(0.28558383),
	1583: -float32(0.29810383),
	1584: -float32(0.31057274),
	1585: -float32(0.32298845),
	1586: -float32(0.33534881),
	1587: -float32(0.34765171),
	1588: -float32(0.35989504),
	1589: -float32(0.3720767),
	1590: -float32(0.38419461),
	1591: -float32(0.3962467),
	1592: -float32(0.40823088),
	1593: -float32(0.42014512),
	1594: -float32(0.43198737),
	1595: -float32(0.4437556),
	1596: -float32(0.4554478),
	1597: -float32(0.46706195),
	1598: -float32(0.47859608),
	1599: -float32(0.49004821),
	1600: -float32(0.50141636),
	1601: -float32(0.5126986),
	1602: -float32(0.52389299),
	1603: -float32(0.53499762),
	1604: -float32(0.54601058),
	1605: -float32(0.55692998),
	1606: -float32(0.56775395),
	1607: -float32(0.57848064),
	1608: -float32(0.58910822),
	1609: -float32(0.59963485),
	1610: -float32(0.61005873),
	1611: -float32(0.62037809),
	1612: -float32(0.63059115),
	1613: -float32(0.64069616),
	1614: -float32(0.65069139),
	1615: -float32(0.66057513),
	1616: -float32(0.67034568),
	1617: -float32(0.68000137),
	1618: -float32(0.68954054),
	1619: -float32(0.69896157),
	1620: -float32(0.70826283),
	1621: -float32(0.71744274),
	1622: -float32(0.72649972),
	1623: -float32(0.73543221),
	1624: -float32(0.74423869),
	1625: -float32(0.75291765),
	1626: -float32(0.7614676),
	1627: -float32(0.76988708),
	1628: -float32(0.77817464),
	1629: -float32(0.78632887),
	1630: -float32(0.79434836),
	1631: -float32(0.80223175),
	1632: -float32(0.80997767),
	1633: -float32(0.81758481),
	1634: -float32(0.82505187),
	1635: -float32(0.83237755),
	1636: -float32(0.83956061),
	1637: -float32(0.84659981),
	1638: -float32(0.85349396),
	1639: -float32(0.86024186),
	1640: -float32(0.86684237),
	1641: -float32(0.87329434),
	1642: -float32(0.87959669),
	1643: -float32(0.88574831),
	1644: -float32(0.89174817),
	1645: -float32(0.89759523),
	1646: -float32(0.9032885),
	1647: -float32(0.90882699),
	1648: -float32(0.91420976),
	1649: -float32(0.91943588),
	1650: -float32(0.92450446),
	1651: -float32(0.92941463),
	1652: -float32(0.93416555),
	1653: -float32(0.93875641),
	1654: -float32(0.94318642),
	1655: -float32(0.94745482),
	1656: -float32(0.95156087),
	1657: -float32(0.95550388),
	1658: -float32(0.95928317),
	1659: -float32(0.96289809),
	1660: -float32(0.96634802),
	1661: -float32(0.96963238),
	1662: -float32(0.97275059),
	1663: -float32(0.97570213),
	1664: -float32(0.97848649),
	1665: -float32(0.98110318),
	1666: -float32(0.98355177),
	1667: -float32(0.98583184),
	1668: -float32(0.98794298),
	1669: -float32(0.98988485),
	1670: -float32(0.99165711),
	1671: -float32(0.99325945),
	1672: -float32(0.9946916),
	1673: -float32(0.99595331),
	1674: -float32(0.99704438),
	1675: -float32(0.9979646),
	1676: -float32(0.99871383),
	1677: -float32(0.99929194),
	1678: -float32(0.99969882),
	1679: -float32(0.99993441),
	1680: float32(0.99999465),
	1681: float32(0.99956631),
	1682: float32(0.99845292),
	1683: float32(0.99665524),
	1684: float32(0.9941745),
	1685: float32(0.99101241),
	1686: float32(0.98717112),
	1687: float32(0.98265328),
	1688: float32(0.97746197),
	1689: float32(0.97160077),
	1690: float32(0.96507367),
	1691: float32(0.95788516),
	1692: float32(0.95004017),
	1693: float32(0.94154407),
	1694: float32(0.93240267),
	1695: float32(0.92262226),
	1696: float32(0.91220953),
	1697: float32(0.90117161),
	1698: float32(0.88951608),
	1699: float32(0.87725091),
	1700: float32(0.86438452),
	1701: float32(0.85092573),
	1702: float32(0.83688375),
	1703: float32(0.82226822),
	1704: float32(0.80708914),
	1705: float32(0.79135693),
	1706: float32(0.77508236),
	1707: float32(0.75827658),
	1708: float32(0.74095113),
	1709: float32(0.72311786),
	1710: float32(0.704789),
	1711: float32(0.68597711),
	1712: float32(0.66669509),
	1713: float32(0.64695615),
	1714: float32(0.62677382),
	1715: float32(0.60616193),
	1716: float32(0.58513461),
	1717: float32(0.56370626),
	1718: float32(0.54189158),
	1719: float32(0.51970551),
	1720: float32(0.49716327),
	1721: float32(0.47428029),
	1722: float32(0.45107226),
	1723: float32(0.42755509),
	1724: float32(0.4037449),
	1725: float32(0.379658),
	1726: float32(0.3553109),
	1727: float32(0.33072029),
	1728: float32(0.30590302),
	1729: float32(0.2808761),
	1730: float32(0.25565668),
	1731: float32(0.23026205),
	1732: float32(0.2047096),
	1733: float32(0.17901686),
	1734: float32(0.15320143),
	1735: float32(0.127281),
	1736: float32(0.10127334),
	1737: float32(0.075196277),
	1738: float32(0.049067674),
	1739: float32(0.022905443),
	1740: -float32(0.0032724865),
	1741: -float32(0.029448173),
	1742: -float32(0.055603678),
	1743: -float32(0.081721074),
	1744: -float32(0.10778246),
	1745: -float32(0.13376998),
	1746: -float32(0.15966582),
	1747: -float32(0.18545224),
	1748: -float32(0.21111155),
	1749: -float32(0.23662618),
	1750: -float32(0.26197864),
	1751: -float32(0.28715155),
	1752: -float32(0.31212766),
	1753: -float32(0.33688985),
	1754: -float32(0.36142116),
	1755: -float32(0.38570477),
	1756: -float32(0.40972403),
	1757: -float32(0.43346249),
	1758: -float32(0.45690388),
	1759: -float32(0.48003212),
	1760: -float32(0.50283138),
	1761: -float32(0.52528602),
	1762: -float32(0.54738066),
	1763: -float32(0.56910015),
	1764: -float32(0.5904296),
	1765: -float32(0.61135441),
	1766: -float32(0.63186022),
	1767: -float32(0.65193299),
	1768: -float32(0.67155895),
	1769: -float32(0.69072467),
	1770: -float32(0.70941699),
	1771: -float32(0.72762312),
	1772: -float32(0.74533057),
	1773: -float32(0.7625272),
	1774: -float32(0.77920124),
	1775: -float32(0.79534126),
	1776: -float32(0.81093618),
	1777: -float32(0.82597533),
	1778: -float32(0.8404484),
	1779: -float32(0.85434547),
	1780: -float32(0.86765701),
	1781: -float32(0.8803739),
	1782: -float32(0.89248743),
	1783: -float32(0.90398929),
	1784: -float32(0.91487161),
	1785: -float32(0.92512691),
	1786: -float32(0.93474818),
	1787: -float32(0.94372882),
	1788: -float32(0.95206268),
	1789: -float32(0.95974404),
	1790: -float32(0.96676764),
	1791: -float32(0.97312866),
	1792: -float32(0.97882275),
	1793: -float32(0.98384601),
	1794: -float32(0.98819498),
	1795: -float32(0.9918667),
	1796: -float32(0.99485864),
	1797: -float32(0.99716875),
	1798: -float32(0.99879546),
	1799: -float32(0.99973764),
}
var mode48000_960_120 = OpusT_OpusCustomMode{
	FFs:        int32(48000),
	Foverlap:   int32(120),
	FnbEBands:  int32(21),
	FeffEBands: int32(21),
	Fpreemph: [4]OpusT_opus_val16{
		0: float32(0.8500061),
		2: float32(1),
		3: float32(1),
	},
	FeBands:         uintptr(unsafe.Pointer(&eband5ms)),
	FmaxLM:          int32(3),
	FnbShortMdcts:   int32(8),
	FshortMdctSize:  int32(120),
	FnbAllocVectors: int32(11),
	FallocVectors:   uintptr(unsafe.Pointer(&band_allocation)),
	FlogN:           uintptr(unsafe.Pointer(&logN400)),
	Fwindow:         uintptr(unsafe.Pointer(&window120)),
	Fmdct: OpusT_mdct_lookup{
		Fn:        int32(1920),
		Fmaxshift: int32(3),
		Fkfft: [4]uintptr{
			0: uintptr(unsafe.Pointer(&fft_state48000_960_0)),
			1: uintptr(unsafe.Pointer(&fft_state48000_960_1)),
			2: uintptr(unsafe.Pointer(&fft_state48000_960_2)),
			3: uintptr(unsafe.Pointer(&fft_state48000_960_3)),
		},
		Ftrig: uintptr(unsafe.Pointer(&mdct_twiddles960)),
	},
	Fcache: OpusT_PulseCache{
		Fsize:  int32(392),
		Findex: uintptr(unsafe.Pointer(&cache_index50)),
		Fbits:  uintptr(unsafe.Pointer(&cache_bits50)),
		Fcaps:  uintptr(unsafe.Pointer(&cache_caps50)),
	},
}
var static_mode_list = [1]uintptr{
	0: uintptr(unsafe.Pointer(&mode48000_960_120)),
}
