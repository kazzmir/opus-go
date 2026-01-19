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

const CELT_LPC_ORDER = 24
const CELT_MAX_PULSES = 128
const CELT_SIG_SCALE3 = 32768
const COEF_ONE3 = 1
const FILENAME_MAX = 4096
const FINE_OFFSET = 21
const FOPEN_MAX = 1000
const LOG_MAX_PSEUDO = 6
const L_ctermid = 20
const L_cuserid = 20
const L_tmpnam = 20
const MAX_FINE_BITS = 8
const MAX_PSEUDO = 40
const P_tmpdir = "/tmp"
const QTHETA_OFFSET = 4
const QTHETA_OFFSET_TWOPHASE = 16
const SPREAD_AGGRESSIVE = 3
const SPREAD_LIGHT = 1
const SPREAD_NONE = 0
const SPREAD_NORMAL = 2
const TMP_MAX = 10000
const _IOFBF = 0
const _IOLBF = 1
const _IONBF = 2
const celt_inner_prod_norm = "celt_inner_prod"
const celt_inner_prod_norm_shift = "celt_inner_prod"
const celt_pitch_xcorr = "celt_pitch_xcorr_c"

type OpusT_ssize_t = int64

type OpusT_off_t = int64

type _IO_FILE = struct {
	F__x int8
}

type OpusT_FILE = struct {
	F__x int8
}

type OpusT___isoc_va_list = uintptr

type OpusT_fpos_t = struct {
	F__lldata [0]int64
	F__align  [0]float64
	F__opaque [16]int8
}

type _G_fpos64_t = OpusT_fpos_t

type OpusT_cookie_io_functions_t = struct {
	Fread   uintptr
	Fwrite  uintptr
	Fseek   uintptr
	Fclose1 uintptr
}

type _IO_cookie_io_functions_t = OpusT_cookie_io_functions_t

func Opus_celt_fatal(tls *libc.TLS, str uintptr, file uintptr, line int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+3082, libc.VaList(bp+8, file, line, str))
	libc.Xabort(tls)
}

var trim_icdf7 = [11]uint8{
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
var spread_icdf7 = [4]uint8{
	0: uint8(25),
	1: uint8(23),
	2: uint8(2),
}
var tapset_icdf7 = [3]uint8{
	0: uint8(2),
	1: uint8(1),
}

var log2_x_norm_coeff3 = [8]float32{
	0: float32(1),
	1: float32(0.8888888955116272),
	2: float32(0.8),
	3: float32(0.7272727489471436),
	4: float32(0.6666666865348816),
	5: float32(0.6153846383094788),
	6: float32(0.5714285969734192),
	7: float32(0.5333333611488342),
}
var log2_y_norm_coeff3 = [8]float32{
	1: float32(0.1699250042438507),
	2: float32(0.32192808389663696),
	3: float32(0.45943161845207214),
	4: float32(0.5849624872207642),
	5: float32(0.7004396915435791),
	6: float32(0.8073549270629883),
	7: float32(0.9068905711174011),
}

func Opus_resampling_factor(tls *libc.TLS, rate OpusT_opus_int32) (r int32) {
	var ret int32
	_ = ret
	switch rate {
	case int32(48000):
		ret = int32(1)
	case int32(24000):
		ret = int32(2)
	case int32(16000):
		ret = int32(3)
	case int32(12000):
		ret = int32(4)
	case int32(8000):
		ret = int32(6)
	default:
		if !(int32(0) != 0) {
			Opus_celt_fatal(tls, __ccgo_ts+1017, __ccgo_ts+3125, int32(87))
		}
		ret = 0
		break
	}
	return ret
}

// C documentation
//
//	/* This version should be faster on ARM */
func comb_filter_const_c(tls *libc.TLS, y uintptr, x uintptr, T int32, N int32, g10 OpusT_celt_coef, g11 OpusT_celt_coef, g12 OpusT_celt_coef) {
	var i int32
	var x0, x1, x2, x3, x4 OpusT_opus_val32
	_, _, _, _, _, _ = i, x0, x1, x2, x3, x4
	x4 = *(*OpusT_opus_val32)(unsafe.Pointer(x + uintptr(-T-int32(2))*4))
	x3 = *(*OpusT_opus_val32)(unsafe.Pointer(x + uintptr(-T-int32(1))*4))
	x2 = *(*OpusT_opus_val32)(unsafe.Pointer(x + uintptr(-T)*4))
	x1 = *(*OpusT_opus_val32)(unsafe.Pointer(x + uintptr(-T+int32(1))*4))
	i = 0
	for {
		if !(i < N) {
			break
		}
		x0 = *(*OpusT_opus_val32)(unsafe.Pointer(x + uintptr(i-T+int32(2))*4))
		*(*OpusT_opus_val32)(unsafe.Pointer(y + uintptr(i)*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(x + uintptr(i)*4)) + OpusT_celt_coef(g10*x2) + OpusT_celt_coef(g11*(x1+x3)) + OpusT_celt_coef(g12*(x0+x4))
		*(*OpusT_opus_val32)(unsafe.Pointer(y + uintptr(i)*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(y + uintptr(i)*4))
		x4 = x3
		x3 = x2
		x2 = x1
		x1 = x0
		goto _1
	_1:
		;
		i = i + 1
	}
}

func Opus_comb_filter(tls *libc.TLS, y uintptr, x uintptr, T0 int32, T1 int32, N int32, g0 OpusT_opus_val16, g1 OpusT_opus_val16, tapset0 int32, tapset1 int32, window uintptr, overlap int32, arch int32) {
	var f, g00, g01, g02, g10, g11, g12 OpusT_celt_coef
	var i, v1 int32
	var x0, x1, x2, x3, x4 OpusT_opus_val32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = f, g00, g01, g02, g10, g11, g12, i, x0, x1, x2, x3, x4, v1
	if g0 == float32(0) && g1 == float32(0) {
		/* OPT: Happens to work without the OPUS_MOVE(), but only because the current encoder already copies x to y */
		if x != y {
			libc.Xmemmove(tls, y, x, uint64(uint32(N))*uint64(4)+uint64(0*((int64(y)-int64(x))/4)))
		}
		return
	}
	/* When the gain is zero, T0 and/or T1 is set to zero. We need
	   to have then be at least 2 to avoid processing garbage data. */
	if T0 > int32(COMBFILTER_MINPERIOD) {
		v1 = T0
	} else {
		v1 = int32(COMBFILTER_MINPERIOD)
	}
	T0 = v1
	if T1 > int32(COMBFILTER_MINPERIOD) {
		v1 = T1
	} else {
		v1 = int32(COMBFILTER_MINPERIOD)
	}
	T1 = v1
	g00 = OpusT_opus_val16(g0 * *(*OpusT_opus_val16)(unsafe.Pointer(uintptr(unsafe.Pointer(&gains)) + uintptr(tapset0)*12)))
	g01 = OpusT_opus_val16(g0 * *(*OpusT_opus_val16)(unsafe.Pointer(uintptr(unsafe.Pointer(&gains)) + uintptr(tapset0)*12 + 1*4)))
	g02 = OpusT_opus_val16(g0 * *(*OpusT_opus_val16)(unsafe.Pointer(uintptr(unsafe.Pointer(&gains)) + uintptr(tapset0)*12 + 2*4)))
	g10 = OpusT_opus_val16(g1 * *(*OpusT_opus_val16)(unsafe.Pointer(uintptr(unsafe.Pointer(&gains)) + uintptr(tapset1)*12)))
	g11 = OpusT_opus_val16(g1 * *(*OpusT_opus_val16)(unsafe.Pointer(uintptr(unsafe.Pointer(&gains)) + uintptr(tapset1)*12 + 1*4)))
	g12 = OpusT_opus_val16(g1 * *(*OpusT_opus_val16)(unsafe.Pointer(uintptr(unsafe.Pointer(&gains)) + uintptr(tapset1)*12 + 2*4)))
	x1 = *(*OpusT_opus_val32)(unsafe.Pointer(x + uintptr(-T1+int32(1))*4))
	x2 = *(*OpusT_opus_val32)(unsafe.Pointer(x + uintptr(-T1)*4))
	x3 = *(*OpusT_opus_val32)(unsafe.Pointer(x + uintptr(-T1-int32(1))*4))
	x4 = *(*OpusT_opus_val32)(unsafe.Pointer(x + uintptr(-T1-int32(2))*4))
	/* If the filter didn't change, we don't need the overlap */
	if g0 == g1 && T0 == T1 && tapset0 == tapset1 {
		overlap = 0
	}
	i = 0
	for {
		if !(i < overlap) {
			break
		}
		x0 = *(*OpusT_opus_val32)(unsafe.Pointer(x + uintptr(i-T1+int32(2))*4))
		f = OpusT_celt_coef(*(*OpusT_celt_coef)(unsafe.Pointer(window + uintptr(i)*4)) * *(*OpusT_celt_coef)(unsafe.Pointer(window + uintptr(i)*4)))
		*(*OpusT_opus_val32)(unsafe.Pointer(y + uintptr(i)*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(x + uintptr(i)*4)) + float32(float32((float32(1)-f)*g00)**(*OpusT_opus_val32)(unsafe.Pointer(x + uintptr(i-T0)*4))) + float32(float32((float32(1)-f)*g01)*(*(*OpusT_opus_val32)(unsafe.Pointer(x + uintptr(i-T0+int32(1))*4))+*(*OpusT_opus_val32)(unsafe.Pointer(x + uintptr(i-T0-int32(1))*4)))) + float32(float32((float32(1)-f)*g02)*(*(*OpusT_opus_val32)(unsafe.Pointer(x + uintptr(i-T0+int32(2))*4))+*(*OpusT_opus_val32)(unsafe.Pointer(x + uintptr(i-T0-int32(2))*4)))) + OpusT_celt_coef(OpusT_celt_coef(f*g10)*x2) + OpusT_celt_coef(OpusT_celt_coef(f*g11)*(x1+x3)) + OpusT_celt_coef(OpusT_celt_coef(f*g12)*(x0+x4))
		*(*OpusT_opus_val32)(unsafe.Pointer(y + uintptr(i)*4)) = *(*OpusT_opus_val32)(unsafe.Pointer(y + uintptr(i)*4))
		x4 = x3
		x3 = x2
		x2 = x1
		x1 = x0
		goto _3
	_3:
		;
		i = i + 1
	}
	if g1 == float32(0) {
		/* OPT: Happens to work without the OPUS_MOVE(), but only because the current encoder already copies x to y */
		if x != y {
			libc.Xmemmove(tls, y+uintptr(overlap)*4, x+uintptr(overlap)*4, uint64(uint32(N-overlap))*uint64(4)+uint64(0*((int64(y+uintptr(overlap)*4)-int64(x+uintptr(overlap)*4))/4)))
		}
		return
	}
	/* Compute the part with the constant filter. */
	_ = arch
	comb_filter_const_c(tls, y+uintptr(i)*4, x+uintptr(i)*4, T1, N-i, g10, g11, g12)
}

var gains = [3][3]OpusT_opus_val16{
	0: {
		0: float32(0.306640625),
		1: float32(0.2170410156),
		2: float32(0.1296386719),
	},
	1: {
		0: float32(0.4638671875),
		1: float32(0.2680664062),
	},
	2: {
		0: float32(0.7998046875),
		1: float32(0.1000976562),
	},
}

func Opus_init_caps(tls *libc.TLS, m uintptr, cap1 uintptr, LM int32, C int32) {
	var N, i int32
	_, _ = N, i
	i = 0
	for {
		if !(i < (*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands) {
			break
		}
		N = (int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(i+int32(1))*2))) - int32(*(*OpusT_opus_int16)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FeBands + uintptr(i)*2)))) << LM
		*(*int32)(unsafe.Pointer(cap1 + uintptr(i)*4)) = (int32(*(*uint8)(unsafe.Pointer((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).Fcache.Fcaps + uintptr((*OpusT_OpusCustomMode)(unsafe.Pointer(m)).FnbEBands*(int32(2)*LM+C-int32(1))+i)))) + int32(64)) * C * N >> int32(2)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func Opus_opus_strerror(tls *libc.TLS, error1 int32) (r uintptr) {
	if error1 > 0 || error1 < -int32(7) {
		return __ccgo_ts + 3277
	} else {
		return error_strings[-error1]
	}
	return r
}

var error_strings = [8]uintptr{
	0: __ccgo_ts + 3140,
	1: __ccgo_ts + 3148,
	2: __ccgo_ts + 3165,
	3: __ccgo_ts + 3182,
	4: __ccgo_ts + 3197,
	5: __ccgo_ts + 3214,
	6: __ccgo_ts + 3238,
	7: __ccgo_ts + 3252,
}

func Opus_opus_get_version_string(tls *libc.TLS) (r uintptr) {
	return __ccgo_ts + 3291
}

const CELT_SIG_SCALE4 = "32768.f"
const COEF_ONE4 = "1.0f"

var log2_x_norm_coeff4 = [8]float32{
	0: float32(1),
	1: float32(0.8888888955116272),
	2: float32(0.8),
	3: float32(0.7272727489471436),
	4: float32(0.6666666865348816),
	5: float32(0.6153846383094788),
	6: float32(0.5714285969734192),
	7: float32(0.5333333611488342),
}
var log2_y_norm_coeff4 = [8]float32{
	1: float32(0.1699250042438507),
	2: float32(0.32192808389663696),
	3: float32(0.45943161845207214),
	4: float32(0.5849624872207642),
	5: float32(0.7004396915435791),
	6: float32(0.8073549270629883),
	7: float32(0.9068905711174011),
}

var trim_icdf8 = [11]uint8{
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
var spread_icdf8 = [4]uint8{
	0: uint8(25),
	1: uint8(23),
	2: uint8(2),
}
var tapset_icdf8 = [3]uint8{
	0: uint8(2),
	1: uint8(1),
}
