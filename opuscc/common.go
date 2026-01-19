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

const ARG_MAX = 131072
const ATAN2_2_OVER_PI = 0.636619772367581
const ATAN2_COEFF_A05 = 0.19962704181671143
const ATAN2_COEFF_A09 = 0.09794234484434128
const ATAN2_COEFF_A13 = 0.023040136322379112
const BC_BASE_MAX = 99
const BC_DIM_MAX = 2048
const BC_SCALE_MAX = 99
const BC_STRING_MAX = 1000
const BITRES = 3
const CCGO_TLS_PSEUDOSTACK = 1
const CELTDecoder = "OpusCustomDecoder"
const CELTEncoder = "OpusCustomEncoder"
const CELTMode = "OpusCustomMode"
const CELT_GET_AND_CLEAR_ERROR_REQUEST = 10007
const CELT_GET_MODE_REQUEST = 10015
const CELT_SET_ANALYSIS_REQUEST = 10022
const CELT_SET_CHANNELS_REQUEST = 10008
const CELT_SET_END_BAND_REQUEST = 10012
const CELT_SET_INPUT_CLIPPING_REQUEST = 10004
const CELT_SET_PREDICTION_REQUEST = 10002
const CELT_SET_SIGNALLING_REQUEST = 10016
const CELT_SET_SILK_INFO_REQUEST = 10028
const CELT_SET_START_BAND_REQUEST = 10010
const CELT_SET_TONALITY_REQUEST = 10018
const CELT_SET_TONALITY_SLOPE_REQUEST = 10020
const CELT_SIG_SCALE = "32768.f"
const CHARCLASS_NAME_MAX = 14
const CHAR_BIT = 8
const CHAR_MAX = 255
const CHAR_MIN = 0
const COEF_ONE = "1.0f"
const COLL_WEIGHTS_MAX = 2
const COMBFILTER_MAXPERIOD = 1024
const COMBFILTER_MINPERIOD = 15
const COS_COEFF_A0 = 0.9999999403953552
const COS_COEFF_A4 = 0.2536507546901703
const COS_COEFF_A8 = 0.0008581906440667808
const CPU_INFO_BY_ASM = 1
const DELAYTIMER_MAX = 0x7fffffff
const DISABLE_DEBUG_FLOAT = 1
const EC_UINT_BITS = 8
const ENABLE_HARDENING = 1
const ENABLE_RES24 = 1
const EPSILON = "1e-15f"
const EXIT_FAILURE = 1
const EXIT_SUCCESS = 0
const EXP2_COEFF_A0 = 0.9999999403953552
const EXP2_COEFF_A1 = 0.6931530833244324
const EXP2_COEFF_A2 = 0.24015361070632935
const EXP2_COEFF_A3 = 0.05582631751894951
const EXP2_COEFF_A4 = 0.00898933969438076
const EXP2_COEFF_A5 = 0.0018775766948238015
const EXPR_NEST_MAX = 32
const FILESIZEBITS = 64
const FLOAT_APPROX = 1
const FP_ILOGB0 = "FP_ILOGBNAN"
const FP_INFINITE = 1
const FP_NAN = 0
const FP_NORMAL = 4
const FP_SUBNORMAL = 3
const FP_ZERO = 2
const GLOBAL_STACK_SIZE = 120000
const HAVE_DLFCN_H = 1
const HAVE_INTTYPES_H = 1
const HAVE_LRINT = 1
const HAVE_LRINTF = 1
const HAVE_STDINT_H = 1
const HAVE_STDIO_H = 1
const HAVE_STDLIB_H = 1
const HAVE_STRINGS_H = 1
const HAVE_STRING_H = 1
const HAVE_SYS_STAT_H = 1
const HAVE_SYS_TYPES_H = 1
const HAVE_UNISTD_H = 1
const HOST_NAME_MAX = 255
const HUGE = 3.40282346638528859812e+38
const HUGE_VALF = "INFINITY"
const INT16_MAX = 0x7fff
const INT32_MAX = 0x7fffffff
const INT64_MAX = 0x7fffffffffffffff
const INT8_MAX = 0x7f
const INTMAX_MAX = "INT64_MAX"
const INTMAX_MIN = "INT64_MIN"
const INTPTR_MAX = "INT64_MAX"
const INTPTR_MIN = "INT64_MIN"
const INT_FAST16_MAX = "INT32_MAX"
const INT_FAST16_MIN = "INT32_MIN"
const INT_FAST32_MAX = "INT32_MAX"
const INT_FAST32_MIN = "INT32_MIN"
const INT_FAST64_MAX = "INT64_MAX"
const INT_FAST64_MIN = "INT64_MIN"
const INT_FAST8_MAX = "INT8_MAX"
const INT_FAST8_MIN = "INT8_MIN"
const INT_LEAST16_MAX = "INT16_MAX"
const INT_LEAST16_MIN = "INT16_MIN"
const INT_LEAST32_MAX = "INT32_MAX"
const INT_LEAST32_MIN = "INT32_MIN"
const INT_LEAST64_MAX = "INT64_MAX"
const INT_LEAST64_MIN = "INT64_MIN"
const INT_LEAST8_MAX = "INT8_MAX"
const INT_LEAST8_MIN = "INT8_MIN"
const INT_MAX = 0x7fffffff
const IOV_MAX = 1024
const KF_SUFFIX = "_celt_single"
const KISS_FFT_MALLOC = "opus_alloc"
const LEAK_BANDS = 19
const LINE_MAX = 4096
const LLONG_MAX = 0x7fffffffffffffff
const LOG2_COEFF_A0 = 0.08746284246444702
const LOG2_COEFF_A1 = 1.3578295707702637
const LOG2_COEFF_A3 = 0.4019712507724762
const LOGIN_NAME_MAX = 256
const LONG_BIT = 64
const LONG_MAX = "__LONG_MAX"
const LT_OBJDIR = ".libs/"
const MATH_ERREXCEPT = 2
const MATH_ERRNO = 1
const MAXFACTORS = 8
const MAX_ENCODING_DEPTH = 24
const MB_LEN_MAX = 4
const MODE_CELT_ONLY = 1002
const MODE_HYBRID = 1001
const MODE_SILK_ONLY = 1000
const MQ_PRIO_MAX = 32768
const M_1_PI = 0.31830988618379067154
const M_2_PI = 0.63661977236758134308
const M_2_SQRTPI = 1.12837916709551257390
const M_E = 2.7182818284590452354
const M_LN10 = 2.30258509299404568402
const M_LN2 = 0.69314718055994530942
const M_LOG10E = 0.43429448190325182765
const M_LOG2E = 1.4426950408889634074
const M_PI = 3.14159265358979323846
const M_PI_2 = 1.57079632679489661923
const M_PI_4 = 0.78539816339744830962
const M_SQRT1_2 = 0.70710678118654752440
const M_SQRT2 = 1.41421356237309504880
const NAME_MAX = 255
const NGROUPS_MAX = 32
const NL_ARGMAX = 9
const NL_LANGMAX = 32
const NL_MSGMAX = 32767
const NL_NMAX = 16
const NL_SETMAX = 255
const NL_TEXTMAX = 2048
const NONTHREADSAFE_PSEUDOSTACK = 1
const NORM_SCALING = "1.f"
const NZERO = 20
const OPUS_APPLICATION_AUDIO = 2049
const OPUS_APPLICATION_RESTRICTED_CELT = 2053
const OPUS_APPLICATION_RESTRICTED_LOWDELAY = 2051
const OPUS_APPLICATION_RESTRICTED_SILK = 2052
const OPUS_APPLICATION_VOIP = 2048
const OPUS_ARCHMASK = 0
const OPUS_BANDWIDTH_FULLBAND = 1105
const OPUS_BANDWIDTH_MEDIUMBAND = 1102
const OPUS_BANDWIDTH_NARROWBAND = 1101
const OPUS_BANDWIDTH_SUPERWIDEBAND = 1104
const OPUS_BANDWIDTH_WIDEBAND = 1103
const OPUS_DISABLE_INTRINSICS = 1
const OPUS_FAST_INT64 = 1
const OPUS_FRAMESIZE_100_MS = 5008
const OPUS_FRAMESIZE_10_MS = 5003
const OPUS_FRAMESIZE_120_MS = 5009
const OPUS_FRAMESIZE_20_MS = 5004
const OPUS_FRAMESIZE_2_5_MS = 5001
const OPUS_FRAMESIZE_40_MS = 5005
const OPUS_FRAMESIZE_5_MS = 5002
const OPUS_FRAMESIZE_60_MS = 5006
const OPUS_FRAMESIZE_80_MS = 5007
const OPUS_FRAMESIZE_ARG = 5000
const OPUS_GET_APPLICATION_REQUEST = 4001
const OPUS_GET_BANDWIDTH_REQUEST = 4009
const OPUS_GET_BITRATE_REQUEST = 4003
const OPUS_GET_COMPLEXITY_REQUEST = 4011
const OPUS_GET_DRED_DURATION_REQUEST = 4051
const OPUS_GET_DTX_REQUEST = 4017
const OPUS_GET_EXPERT_FRAME_DURATION_REQUEST = 4041
const OPUS_GET_FINAL_RANGE_REQUEST = 4031
const OPUS_GET_FORCE_CHANNELS_REQUEST = 4023
const OPUS_GET_GAIN_REQUEST = 4045
const OPUS_GET_IGNORE_EXTENSIONS_REQUEST = 4059
const OPUS_GET_INBAND_FEC_REQUEST = 4013
const OPUS_GET_IN_DTX_REQUEST = 4049
const OPUS_GET_LAST_PACKET_DURATION_REQUEST = 4039
const OPUS_GET_LOOKAHEAD_REQUEST = 4027
const OPUS_GET_LSB_DEPTH_REQUEST = 4037
const OPUS_GET_MAX_BANDWIDTH_REQUEST = 4005
const OPUS_GET_OSCE_BWE_REQUEST = 4055
const OPUS_GET_PACKET_LOSS_PERC_REQUEST = 4015
const OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST = 4047
const OPUS_GET_PITCH_REQUEST = 4033
const OPUS_GET_PREDICTION_DISABLED_REQUEST = 4043
const OPUS_GET_QEXT_REQUEST = 4057
const OPUS_GET_SAMPLE_RATE_REQUEST = 4029
const OPUS_GET_SIGNAL_REQUEST = 4025
const OPUS_GET_VBR_CONSTRAINT_REQUEST = 4021
const OPUS_GET_VBR_REQUEST = 4007
const OPUS_GET_VOICE_RATIO_REQUEST = 11019
const OPUS_INLINE = "inline"
const OPUS_OK = 0
const OPUS_RESET_STATE = 4028
const OPUS_RESTRICT = "restrict"
const OPUS_SET_APPLICATION_REQUEST = 4000
const OPUS_SET_BANDWIDTH_REQUEST = 4008
const OPUS_SET_BITRATE_REQUEST = 4002
const OPUS_SET_COMPLEXITY_REQUEST = 4010
const OPUS_SET_DNN_BLOB_REQUEST = 4052
const OPUS_SET_DRED_DURATION_REQUEST = 4050
const OPUS_SET_DTX_REQUEST = 4016
const OPUS_SET_ENERGY_MASK_REQUEST = 10026
const OPUS_SET_EXPERT_FRAME_DURATION_REQUEST = 4040
const OPUS_SET_FORCE_CHANNELS_REQUEST = 4022
const OPUS_SET_FORCE_MODE_REQUEST = 11002
const OPUS_SET_GAIN_REQUEST = 4034
const OPUS_SET_IGNORE_EXTENSIONS_REQUEST = 4058
const OPUS_SET_INBAND_FEC_REQUEST = 4012
const OPUS_SET_LFE_REQUEST = 10024
const OPUS_SET_LSB_DEPTH_REQUEST = 4036
const OPUS_SET_MAX_BANDWIDTH_REQUEST = 4004
const OPUS_SET_OSCE_BWE_REQUEST = 4054
const OPUS_SET_PACKET_LOSS_PERC_REQUEST = 4014
const OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST = 4046
const OPUS_SET_PREDICTION_DISABLED_REQUEST = 4042
const OPUS_SET_QEXT_REQUEST = 4056
const OPUS_SET_SIGNAL_REQUEST = 4024
const OPUS_SET_VBR_CONSTRAINT_REQUEST = 4020
const OPUS_SET_VBR_REQUEST = 4006
const OPUS_SET_VOICE_RATIO_REQUEST = 11018
const OPUS_SIGNAL_MUSIC = 3002
const OPUS_SIGNAL_VOICE = 3001
const OPUS_X86_PRESUME_SSE = 1
const OPUS_X86_PRESUME_SSE2 = 1
const PACKAGE_BUGREPORT = "opus@xiph.org"
const PACKAGE_NAME = "opus"
const PACKAGE_STRING = "opus 1.6.1"
const PACKAGE_TARNAME = "opus"
const PACKAGE_URL = ""
const PACKAGE_VERSION = "1.6.1"
const PAGESIZE = 4096
const PAGE_SIZE = "PAGESIZE"
const PATH_MAX = 4096
const PI = 3.141592653589793
const PIPE_BUF = 4096
const PTHREAD_DESTRUCTOR_ITERATIONS = 4
const PTHREAD_KEYS_MAX = 128
const PTHREAD_STACK_MIN = 2048
const PTRDIFF_MAX = "INT64_MAX"
const PTRDIFF_MIN = "INT64_MIN"
const Q15ONE = "1.0f"
const Q31ONE = "1.0f"
const QEXT_EXTENSION_ID = 124
const RAND_MAX = 0x7fffffff
const RE_DUP_MAX = 255
const SCHAR_MAX = 127
const SEM_NSEMS_MAX = 256
const SEM_VALUE_MAX = 0x7fffffff
const SHRT_MAX = 0x7fff
const SIG_ATOMIC_MAX = "INT32_MAX"
const SIG_ATOMIC_MIN = "INT32_MIN"
const SIZE_MAX = "UINT64_MAX"
const SSIZE_MAX = "LONG_MAX"
const STDC_HEADERS = 1
const SYMLOOP_MAX = 40
const TTY_NAME_MAX = 32
const TZNAME_MAX = 6
const UCHAR_MAX = 255
const UINT16_MAX = 0xffff
const UINT32_MAX = "0xffffffffu"
const UINT64_MAX = "0xffffffffffffffffu"
const UINT8_MAX = 0xff
const UINTMAX_MAX = "UINT64_MAX"
const UINTPTR_MAX = "UINT64_MAX"
const UINT_FAST16_MAX = "UINT32_MAX"
const UINT_FAST32_MAX = "UINT32_MAX"
const UINT_FAST64_MAX = "UINT64_MAX"
const UINT_FAST8_MAX = "UINT8_MAX"
const UINT_LEAST16_MAX = "UINT16_MAX"
const UINT_LEAST32_MAX = "UINT32_MAX"
const UINT_LEAST64_MAX = "UINT64_MAX"
const UINT_LEAST8_MAX = "UINT8_MAX"
const UINT_MAX = 0xffffffff
const USHRT_MAX = 0xffff
const VERY_LARGE16 = "1e15f"
const VERY_SMALL = "1e-30f"
const WINT_MAX = "UINT32_MAX"
const WINT_MIN = 0
const WNOHANG = 1
const WORD_BIT = 32
const WUNTRACED = 2
const _FORTIFY_SOURCE = 3
const _GNU_SOURCE = 1
const _LP64 = 1
const _POSIX2_BC_BASE_MAX = 99
const _POSIX2_BC_DIM_MAX = 2048
const _POSIX2_BC_SCALE_MAX = 99
const _POSIX2_BC_STRING_MAX = 1000
const _POSIX2_CHARCLASS_NAME_MAX = 14
const _POSIX2_COLL_WEIGHTS_MAX = 2
const _POSIX2_EXPR_NEST_MAX = 32
const _POSIX2_LINE_MAX = 2048
const _POSIX2_RE_DUP_MAX = 255
const _POSIX_AIO_LISTIO_MAX = 2
const _POSIX_AIO_MAX = 1
const _POSIX_ARG_MAX = 4096
const _POSIX_CHILD_MAX = 25
const _POSIX_CLOCKRES_MIN = 20000000
const _POSIX_DELAYTIMER_MAX = 32
const _POSIX_HOST_NAME_MAX = 255
const _POSIX_LINK_MAX = 8
const _POSIX_LOGIN_NAME_MAX = 9
const _POSIX_MAX_CANON = 255
const _POSIX_MAX_INPUT = 255
const _POSIX_MQ_OPEN_MAX = 8
const _POSIX_MQ_PRIO_MAX = 32
const _POSIX_NAME_MAX = 14
const _POSIX_NGROUPS_MAX = 8
const _POSIX_OPEN_MAX = 20
const _POSIX_PATH_MAX = 256
const _POSIX_PIPE_BUF = 512
const _POSIX_RE_DUP_MAX = 255
const _POSIX_RTSIG_MAX = 8
const _POSIX_SEM_NSEMS_MAX = 256
const _POSIX_SEM_VALUE_MAX = 32767
const _POSIX_SIGQUEUE_MAX = 32
const _POSIX_SSIZE_MAX = 32767
const _POSIX_SS_REPL_MAX = 4
const _POSIX_STREAM_MAX = 8
const _POSIX_SYMLINK_MAX = 255
const _POSIX_SYMLOOP_MAX = 8
const _POSIX_THREAD_DESTRUCTOR_ITERATIONS = 4
const _POSIX_THREAD_KEYS_MAX = 128
const _POSIX_THREAD_THREADS_MAX = 64
const _POSIX_TIMER_MAX = 32
const _POSIX_TRACE_EVENT_NAME_MAX = 30
const _POSIX_TRACE_NAME_MAX = 8
const _POSIX_TRACE_SYS_MAX = 8
const _POSIX_TRACE_USER_EVENT_MAX = 32
const _POSIX_TTY_NAME_MAX = 9
const _POSIX_TZNAME_MAX = 6
const _STDC_PREDEF_H = 1
const _XOPEN_IOV_MAX = 16
const _XOPEN_NAME_MAX = 255
const _XOPEN_PATH_MAX = 1024
const __ATOMIC_ACQUIRE = 2
const __ATOMIC_ACQ_REL = 4
const __ATOMIC_CONSUME = 1
const __ATOMIC_HLE_ACQUIRE = 65536
const __ATOMIC_HLE_RELEASE = 131072
const __ATOMIC_RELAXED = 0
const __ATOMIC_RELEASE = 3
const __ATOMIC_SEQ_CST = 5
const __BFLT16_DECIMAL_DIG__ = 4
const __BFLT16_DENORM_MIN__ = "9.18354961579912115600575419704879436e-41B"
const __BFLT16_DIG__ = 2
const __BFLT16_EPSILON__ = "7.81250000000000000000000000000000000e-3B"
const __BFLT16_HAS_DENORM__ = 1
const __BFLT16_HAS_INFINITY__ = 1
const __BFLT16_HAS_QUIET_NAN__ = 1
const __BFLT16_IS_IEC_60559__ = 0
const __BFLT16_MANT_DIG__ = 8
const __BFLT16_MAX_10_EXP__ = 38
const __BFLT16_MAX_EXP__ = 128
const __BFLT16_MAX__ = "3.38953138925153547590470800371487867e+38B"
const __BFLT16_MIN__ = "1.17549435082228750796873653722224568e-38B"
const __BFLT16_NORM_MAX__ = "3.38953138925153547590470800371487867e+38B"
const __BIGGEST_ALIGNMENT__ = 16
const __BIG_ENDIAN = 4321
const __BYTE_ORDER = 1234
const __BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __CCGO__ = 1
const __CET__ = 3
const __CHAR_BIT__ = 8
const __DBL_DECIMAL_DIG__ = 17
const __DBL_DIG__ = 15
const __DBL_HAS_DENORM__ = 1
const __DBL_HAS_INFINITY__ = 1
const __DBL_HAS_QUIET_NAN__ = 1
const __DBL_IS_IEC_60559__ = 1
const __DBL_MANT_DIG__ = 53
const __DBL_MAX_10_EXP__ = 308
const __DBL_MAX_EXP__ = 1024
const __DEC128_EPSILON__ = 1e-33
const __DEC128_MANT_DIG__ = 34
const __DEC128_MAX_EXP__ = 6145
const __DEC128_MAX__ = "9.999999999999999999999999999999999E6144"
const __DEC128_MIN__ = 1e-6143
const __DEC128_SUBNORMAL_MIN__ = 0.000000000000000000000000000000001e-6143
const __DEC32_EPSILON__ = 1e-6
const __DEC32_MANT_DIG__ = 7
const __DEC32_MAX_EXP__ = 97
const __DEC32_MAX__ = 9.999999e96
const __DEC32_MIN__ = 1e-95
const __DEC32_SUBNORMAL_MIN__ = 0.000001e-95
const __DEC64_EPSILON__ = 1e-15
const __DEC64_MANT_DIG__ = 16
const __DEC64_MAX_EXP__ = 385
const __DEC64_MAX__ = "9.999999999999999E384"
const __DEC64_MIN__ = 1e-383
const __DEC64_SUBNORMAL_MIN__ = 0.000000000000001e-383
const __DECIMAL_BID_FORMAT__ = 1
const __DECIMAL_DIG__ = 17
const __DEC_EVAL_METHOD__ = 2
const __ELF__ = 1
const __FINITE_MATH_ONLY__ = 0
const __FLOAT_WORD_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __FLT128_DECIMAL_DIG__ = 36
const __FLT128_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __FLT128_DIG__ = 33
const __FLT128_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __FLT128_HAS_DENORM__ = 1
const __FLT128_HAS_INFINITY__ = 1
const __FLT128_HAS_QUIET_NAN__ = 1
const __FLT128_IS_IEC_60559__ = 1
const __FLT128_MANT_DIG__ = 113
const __FLT128_MAX_10_EXP__ = 4932
const __FLT128_MAX_EXP__ = 16384
const __FLT128_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT128_MIN__ = 3.36210314311209350626267781732175260e-4932
const __FLT128_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT16_DECIMAL_DIG__ = 5
const __FLT16_DENORM_MIN__ = 5.96046447753906250000000000000000000e-8
const __FLT16_DIG__ = 3
const __FLT16_EPSILON__ = 9.76562500000000000000000000000000000e-4
const __FLT16_HAS_DENORM__ = 1
const __FLT16_HAS_INFINITY__ = 1
const __FLT16_HAS_QUIET_NAN__ = 1
const __FLT16_IS_IEC_60559__ = 1
const __FLT16_MANT_DIG__ = 11
const __FLT16_MAX_10_EXP__ = 4
const __FLT16_MAX_EXP__ = 16
const __FLT16_MAX__ = 6.55040000000000000000000000000000000e+4
const __FLT16_MIN__ = 6.10351562500000000000000000000000000e-5
const __FLT16_NORM_MAX__ = 6.55040000000000000000000000000000000e+4
const __FLT32X_DECIMAL_DIG__ = 17
const __FLT32X_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __FLT32X_DIG__ = 15
const __FLT32X_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __FLT32X_HAS_DENORM__ = 1
const __FLT32X_HAS_INFINITY__ = 1
const __FLT32X_HAS_QUIET_NAN__ = 1
const __FLT32X_IS_IEC_60559__ = 1
const __FLT32X_MANT_DIG__ = 53
const __FLT32X_MAX_10_EXP__ = 308
const __FLT32X_MAX_EXP__ = 1024
const __FLT32X_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT32X_MIN__ = 2.22507385850720138309023271733240406e-308
const __FLT32X_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT32_DECIMAL_DIG__ = 9
const __FLT32_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const __FLT32_DIG__ = 6
const __FLT32_EPSILON__ = 1.19209289550781250000000000000000000e-7
const __FLT32_HAS_DENORM__ = 1
const __FLT32_HAS_INFINITY__ = 1
const __FLT32_HAS_QUIET_NAN__ = 1
const __FLT32_IS_IEC_60559__ = 1
const __FLT32_MANT_DIG__ = 24
const __FLT32_MAX_10_EXP__ = 38
const __FLT32_MAX_EXP__ = 128
const __FLT32_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT32_MIN__ = 1.17549435082228750796873653722224568e-38
const __FLT32_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT64X_DECIMAL_DIG__ = 36
const __FLT64X_DENORM_MIN__ = 6.47517511943802511092443895822764655e-4966
const __FLT64X_DIG__ = 33
const __FLT64X_EPSILON__ = 1.92592994438723585305597794258492732e-34
const __FLT64X_HAS_DENORM__ = 1
const __FLT64X_HAS_INFINITY__ = 1
const __FLT64X_HAS_QUIET_NAN__ = 1
const __FLT64X_IS_IEC_60559__ = 1
const __FLT64X_MANT_DIG__ = 113
const __FLT64X_MAX_10_EXP__ = 4932
const __FLT64X_MAX_EXP__ = 16384
const __FLT64X_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT64X_MIN__ = 3.36210314311209350626267781732175260e-4932
const __FLT64X_NORM_MAX__ = "1.18973149535723176508575932662800702e+4932"
const __FLT64_DECIMAL_DIG__ = 17
const __FLT64_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __FLT64_DIG__ = 15
const __FLT64_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __FLT64_HAS_DENORM__ = 1
const __FLT64_HAS_INFINITY__ = 1
const __FLT64_HAS_QUIET_NAN__ = 1
const __FLT64_IS_IEC_60559__ = 1
const __FLT64_MANT_DIG__ = 53
const __FLT64_MAX_10_EXP__ = 308
const __FLT64_MAX_EXP__ = 1024
const __FLT64_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT64_MIN__ = 2.22507385850720138309023271733240406e-308
const __FLT64_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __FLT_DECIMAL_DIG__ = 9
const __FLT_DENORM_MIN__ = 1.40129846432481707092372958328991613e-45
const __FLT_DIG__ = 6
const __FLT_EPSILON__ = 1.19209289550781250000000000000000000e-7
const __FLT_EVAL_METHOD_TS_18661_3__ = 0
const __FLT_EVAL_METHOD__ = 0
const __FLT_HAS_DENORM__ = 1
const __FLT_HAS_INFINITY__ = 1
const __FLT_HAS_QUIET_NAN__ = 1
const __FLT_IS_IEC_60559__ = 1
const __FLT_MANT_DIG__ = 24
const __FLT_MAX_10_EXP__ = 38
const __FLT_MAX_EXP__ = 128
const __FLT_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT_MIN__ = 1.17549435082228750796873653722224568e-38
const __FLT_NORM_MAX__ = 3.40282346638528859811704183484516925e+38
const __FLT_RADIX__ = 2
const __FUNCTION__ = "__func__"
const __FXSR__ = 1
const __GCC_ASM_FLAG_OUTPUTS__ = 1
const __GCC_ATOMIC_BOOL_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR16_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR32_T_LOCK_FREE = 2
const __GCC_ATOMIC_CHAR_LOCK_FREE = 2
const __GCC_ATOMIC_INT_LOCK_FREE = 2
const __GCC_ATOMIC_LLONG_LOCK_FREE = 2
const __GCC_ATOMIC_LONG_LOCK_FREE = 2
const __GCC_ATOMIC_POINTER_LOCK_FREE = 2
const __GCC_ATOMIC_SHORT_LOCK_FREE = 2
const __GCC_ATOMIC_TEST_AND_SET_TRUEVAL = 1
const __GCC_ATOMIC_WCHAR_T_LOCK_FREE = 2
const __GCC_CONSTRUCTIVE_SIZE = 64
const __GCC_DESTRUCTIVE_SIZE = 64
const __GCC_HAVE_DWARF2_CFI_ASM = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_1 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_2 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_4 = 1
const __GCC_HAVE_SYNC_COMPARE_AND_SWAP_8 = 1
const __GCC_IEC_559 = 2
const __GCC_IEC_559_COMPLEX = 2
const __GNUC_EXECUTION_CHARSET_NAME = "UTF-8"
const __GNUC_MINOR__ = 3
const __GNUC_PATCHLEVEL__ = 0
const __GNUC_STDC_INLINE__ = 1
const __GNUC_WIDE_EXECUTION_CHARSET_NAME = "UTF-32LE"
const __GNUC__ = 13
const __GXX_ABI_VERSION = 1018
const __HAVE_SPECULATION_SAFE_VALUE = 1
const __INT16_MAX__ = 0x7fff
const __INT32_MAX__ = 0x7fffffff
const __INT32_TYPE__ = "int"
const __INT64_MAX__ = 0x7fffffffffffffff
const __INT8_MAX__ = 0x7f
const __INTMAX_MAX__ = 0x7fffffffffffffff
const __INTMAX_WIDTH__ = 64
const __INTPTR_MAX__ = 0x7fffffffffffffff
const __INTPTR_WIDTH__ = 64
const __INT_FAST16_MAX__ = 0x7fffffffffffffff
const __INT_FAST16_WIDTH__ = 64
const __INT_FAST32_MAX__ = 0x7fffffffffffffff
const __INT_FAST32_WIDTH__ = 64
const __INT_FAST64_MAX__ = 0x7fffffffffffffff
const __INT_FAST64_WIDTH__ = 64
const __INT_FAST8_MAX__ = 0x7f
const __INT_FAST8_WIDTH__ = 8
const __INT_LEAST16_MAX__ = 0x7fff
const __INT_LEAST16_WIDTH__ = 16
const __INT_LEAST32_MAX__ = 0x7fffffff
const __INT_LEAST32_TYPE__ = "int"
const __INT_LEAST32_WIDTH__ = 32
const __INT_LEAST64_MAX__ = 0x7fffffffffffffff
const __INT_LEAST64_WIDTH__ = 64
const __INT_LEAST8_MAX__ = 0x7f
const __INT_LEAST8_WIDTH__ = 8
const __INT_MAX__ = 0x7fffffff
const __INT_WIDTH__ = 32
const __LDBL_DECIMAL_DIG__ = 17
const __LDBL_DENORM_MIN__ = 4.94065645841246544176568792868221372e-324
const __LDBL_DIG__ = 15
const __LDBL_EPSILON__ = 2.22044604925031308084726333618164062e-16
const __LDBL_HAS_DENORM__ = 1
const __LDBL_HAS_INFINITY__ = 1
const __LDBL_HAS_QUIET_NAN__ = 1
const __LDBL_IS_IEC_60559__ = 1
const __LDBL_MANT_DIG__ = 53
const __LDBL_MAX_10_EXP__ = 308
const __LDBL_MAX_EXP__ = 1024
const __LDBL_MAX__ = 1.79769313486231570814527423731704357e+308
const __LDBL_MIN__ = 2.22507385850720138309023271733240406e-308
const __LDBL_NORM_MAX__ = 1.79769313486231570814527423731704357e+308
const __LITTLE_ENDIAN = 1234
const __LONG_DOUBLE_64__ = 1
const __LONG_LONG_MAX__ = 0x7fffffffffffffff
const __LONG_LONG_WIDTH__ = 64
const __LONG_MAX = 0x7fffffffffffffff
const __LONG_MAX__ = 0x7fffffffffffffff
const __LONG_WIDTH__ = 64
const __LP64__ = 1
const __MMX_WITH_SSE__ = 1
const __MMX__ = 1
const __OPTIMIZE__ = 1
const __ORDER_BIG_ENDIAN__ = 4321
const __ORDER_LITTLE_ENDIAN__ = 1234
const __ORDER_PDP_ENDIAN__ = 3412
const __PIC__ = 2
const __PIE__ = 2
const __PRAGMA_REDEFINE_EXTNAME = 1
const __PRETTY_FUNCTION__ = "__func__"
const __PTRDIFF_MAX__ = 0x7fffffffffffffff
const __PTRDIFF_WIDTH__ = 64
const __SCHAR_MAX__ = 0x7f
const __SCHAR_WIDTH__ = 8
const __SEG_FS = 1
const __SEG_GS = 1
const __SHRT_MAX__ = 0x7fff
const __SHRT_WIDTH__ = 16
const __SIG_ATOMIC_MAX__ = 0x7fffffff
const __SIG_ATOMIC_TYPE__ = "int"
const __SIG_ATOMIC_WIDTH__ = 32
const __SIZEOF_DOUBLE__ = 8
const __SIZEOF_FLOAT128__ = 16
const __SIZEOF_FLOAT80__ = 16
const __SIZEOF_FLOAT__ = 4
const __SIZEOF_INT128__ = 16
const __SIZEOF_INT__ = 4
const __SIZEOF_LONG_DOUBLE__ = 8
const __SIZEOF_LONG_LONG__ = 8
const __SIZEOF_LONG__ = 8
const __SIZEOF_POINTER__ = 8
const __SIZEOF_PTRDIFF_T__ = 8
const __SIZEOF_SHORT__ = 2
const __SIZEOF_SIZE_T__ = 8
const __SIZEOF_WCHAR_T__ = 4
const __SIZEOF_WINT_T__ = 4
const __SIZE_MAX__ = 0xffffffffffffffff
const __SIZE_WIDTH__ = 64
const __SSE2_MATH__ = 1
const __SSE_MATH__ = 1
const __SSP_STRONG__ = 3
const __STDC_HOSTED__ = 1
const __STDC_IEC_559_COMPLEX__ = 1
const __STDC_IEC_559__ = 1
const __STDC_IEC_60559_BFP__ = 201404
const __STDC_IEC_60559_COMPLEX__ = 201404
const __STDC_ISO_10646__ = 201706
const __STDC_VERSION__ = 199901
const __STDC__ = 1
const __STRICT_ANSI__ = 1
const __UINT16_MAX__ = 0xffff
const __UINT32_MAX__ = 0xffffffff
const __UINT64_MAX__ = 0xffffffffffffffff
const __UINT8_MAX__ = 0xff
const __UINTMAX_MAX__ = 0xffffffffffffffff
const __UINTPTR_MAX__ = 0xffffffffffffffff
const __UINT_FAST16_MAX__ = 0xffffffffffffffff
const __UINT_FAST32_MAX__ = 0xffffffffffffffff
const __UINT_FAST64_MAX__ = 0xffffffffffffffff
const __UINT_FAST8_MAX__ = 0xff
const __UINT_LEAST16_MAX__ = 0xffff
const __UINT_LEAST32_MAX__ = 0xffffffff
const __UINT_LEAST64_MAX__ = 0xffffffffffffffff
const __UINT_LEAST8_MAX__ = 0xff
const __USE_TIME_BITS64 = 1
const __VERSION__ = "13.3.0"
const __WCHAR_MAX__ = 0x7fffffff
const __WCHAR_TYPE__ = "int"
const __WCHAR_WIDTH__ = 32
const __WINT_MAX__ = 0xffffffff
const __WINT_MIN__ = 0
const __WINT_WIDTH__ = 32
const __amd64 = 1
const __amd64__ = 1
const __code_model_small__ = 1
const __gnu_linux__ = 1
const __inline = "inline"
const __k8 = 1
const __k8__ = 1
const __linux = 1
const __linux__ = 1
const __pic__ = 2
const __pie__ = 2
const __restrict = "restrict"
const __restrict_arr = "restrict"
const __unix = 1
const __unix__ = 1
const __x86_64 = 1
const __x86_64__ = 1
const _ecintrin_H = 1
const _entcode_H = 1
const _entdec_H = 1
const _entenc_H = 1
const alloca = "__builtin_alloca"
const celt_decoder_ctl = "opus_custom_decoder_ctl"
const celt_encoder_ctl = "opus_custom_encoder_ctl"
const celt_exp2_db = "celt_exp2"
const celt_log2_db = "celt_log2"
const celt_maxabs_res = "celt_maxabs16"
const kiss_fft_scalar = "float"
const kiss_twiddle_scalar = "float"
const math_errhandling = 2
const opus_int = "int"
const restrict = "__restrict"

type OpusT___builtin_va_list = uintptr

type OpusT___predefined_size_t = uint64

type OpusT___predefined_wchar_t = int32

type OpusT___predefined_ptrdiff_t = int64

type OpusT_uintptr_t = uint64

type OpusT_intptr_t = int64

type OpusT_int8_t = int8

type OpusT_int16_t = int16

type OpusT_int32_t = int32

type OpusT_int64_t = int64

type OpusT_intmax_t = int64

type OpusT_uint8_t = uint8

type OpusT_uint16_t = uint16

type OpusT_uint32_t = uint32

type OpusT_uint64_t = uint64

type OpusT_uintmax_t = uint64

type OpusT_int_fast8_t = int8

type OpusT_int_fast64_t = int64

type OpusT_int_least8_t = int8

type OpusT_int_least16_t = int16

type OpusT_int_least32_t = int32

type OpusT_int_least64_t = int64

type OpusT_uint_fast8_t = uint8

type OpusT_uint_fast64_t = uint64

type OpusT_uint_least8_t = uint8

type OpusT_uint_least16_t = uint16

type OpusT_uint_least32_t = uint32

type OpusT_uint_least64_t = uint64

type OpusT_int_fast16_t = int32

type OpusT_int_fast32_t = int32

type OpusT_uint_fast16_t = uint32

type OpusT_uint_fast32_t = uint32

type OpusT_opus_int8 = int8

type OpusT_opus_uint8 = uint8

type OpusT_opus_int16 = int16

type OpusT_opus_uint16 = uint16

type OpusT_opus_int32 = int32

type OpusT_opus_uint32 = uint32

type OpusT_opus_int64 = int64

type OpusT_opus_uint64 = uint64

type OpusT_OpusRepacketizer = struct {
	Ftoc               uint8
	Fnb_frames         int32
	Fframes            [48]uintptr
	Flen1              [48]OpusT_opus_int16
	Fframesize         int32
	Fpaddings          [48]uintptr
	Fpadding_len       [48]OpusT_opus_int32
	Fpadding_nb_frames [48]uint8
}

type OpusT_opus_val16 = float32

type OpusT_opus_val32 = float32

type OpusT_opus_val64 = float32

type OpusT_celt_sig = float32

type OpusT_celt_norm = float32

type OpusT_celt_ener = float32

type OpusT_celt_glog = float32

type OpusT_opus_res = float32

type OpusT_celt_coef = float32

type OpusT_wchar_t = int32

type OpusT_size_t = uint64

type OpusT_ptrdiff_t = int64

type OpusT_float_t = float32

type OpusT_double_t = float64

type OpusT_ec_window = uint32

type OpusT_ec_ctx = struct {
	Fbuf         uintptr
	Fstorage     OpusT_opus_uint32
	Fend_offs    OpusT_opus_uint32
	Fend_window  OpusT_ec_window
	Fnend_bits   int32
	Fnbits_total int32
	Foffs        OpusT_opus_uint32
	Frng         OpusT_opus_uint32
	Fval         OpusT_opus_uint32
	Fext         OpusT_opus_uint32
	Frem         int32
	Ferror1      int32
}

type OpusT_ec_enc = struct {
	Fbuf         uintptr
	Fstorage     OpusT_opus_uint32
	Fend_offs    OpusT_opus_uint32
	Fend_window  OpusT_ec_window
	Fnend_bits   int32
	Fnbits_total int32
	Foffs        OpusT_opus_uint32
	Frng         OpusT_opus_uint32
	Fval         OpusT_opus_uint32
	Fext         OpusT_opus_uint32
	Frem         int32
	Ferror1      int32
}

type ec_ctx = OpusT_ec_enc

type OpusT_ec_dec = struct {
	Fbuf         uintptr
	Fstorage     OpusT_opus_uint32
	Fend_offs    OpusT_opus_uint32
	Fend_window  OpusT_ec_window
	Fnend_bits   int32
	Fnbits_total int32
	Foffs        OpusT_opus_uint32
	Frng         OpusT_opus_uint32
	Fval         OpusT_opus_uint32
	Fext         OpusT_opus_uint32
	Frem         int32
	Ferror1      int32
}

type OpusT_locale_t = uintptr

type OpusT_div_t = struct {
	Fquot int32
	Frem  int32
}

type OpusT_ldiv_t = struct {
	Fquot int64
	Frem  int64
}

type OpusT_lldiv_t = struct {
	Fquot int64
	Frem  int64
}

var log2_x_norm_coeff = [8]float32{
	0: float32(1),
	1: float32(0.8888888955116272),
	2: float32(0.8),
	3: float32(0.7272727489471436),
	4: float32(0.6666666865348816),
	5: float32(0.6153846383094788),
	6: float32(0.5714285969734192),
	7: float32(0.5333333611488342),
}
var log2_y_norm_coeff = [8]float32{
	1: float32(0.1699250042438507),
	2: float32(0.32192808389663696),
	3: float32(0.45943161845207214),
	4: float32(0.5849624872207642),
	5: float32(0.7004396915435791),
	6: float32(0.8073549270629883),
	7: float32(0.9068905711174011),
}

type OpusT_kiss_fft_cpx = struct {
	Fr float32
	Fi float32
}

type OpusT_kiss_twiddle_cpx = struct {
	Fr float32
	Fi float32
}

type OpusT_arch_fft_state = struct {
	Fis_supported int32
	Fpriv         uintptr
}

type OpusT_kiss_fft_state = struct {
	Fnfft     int32
	Fscale    OpusT_celt_coef
	Fshift    int32
	Ffactors  [16]OpusT_opus_int16
	Fbitrev   uintptr
	Ftwiddles uintptr
	Farch_fft uintptr
}

type OpusT_AnalysisInfo = struct {
	Fvalid                int32
	Ftonality             float32
	Ftonality_slope       float32
	Fnoisiness            float32
	Factivity             float32
	Fmusic_prob           float32
	Fmusic_prob_min       float32
	Fmusic_prob_max       float32
	Fbandwidth            int32
	Factivity_probability float32
	Fmax_pitch_ratio      float32
	Fleak_boost           [19]uint8
}

type OpusT_SILKInfo = struct {
	FsignalType int32
	Foffset     int32
}

var trim_icdf = [11]uint8{
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
var spread_icdf = [4]uint8{
	0: uint8(25),
	1: uint8(23),
	2: uint8(2),
}
var tapset_icdf = [3]uint8{
	0: uint8(2),
	1: uint8(1),
}

type OpusT_va_list = uintptr

type OpusRepacketizer = struct {
	Ftoc               uint8
	Fnb_frames         int32
	Fframes            [48]uintptr
	Flen1              [48]OpusT_opus_int16
	Fframesize         int32
	Fpaddings          [48]uintptr
	Fpadding_len       [48]OpusT_opus_int32
	Fpadding_nb_frames [48]uint8
}

type OpusT_OpusExtensionIterator = struct {
	Fdata               uintptr
	Fcurr_data          uintptr
	Frepeat_data        uintptr
	Flast_long          uintptr
	Fsrc_data           uintptr
	Flen1               OpusT_opus_int32
	Fcurr_len           OpusT_opus_int32
	Frepeat_len         OpusT_opus_int32
	Fsrc_len            OpusT_opus_int32
	Ftrailing_short_len OpusT_opus_int32
	Fnb_frames          int32
	Fframe_max          int32
	Fcurr_frame         int32
	Frepeat_frame       int32
	Frepeat_l           uint8
}

type OpusT_opus_extension_data = struct {
	Fid    int32
	Fframe int32
	Fdata  uintptr
	Flen1  OpusT_opus_int32
}

type OpusT_ChannelLayout = struct {
	Fnb_channels        int32
	Fnb_streams         int32
	Fnb_coupled_streams int32
	Fmapping            [256]uint8
}

type OpusT_MappingType = int32

const MAPPING_TYPE_NONE = 0
const MAPPING_TYPE_SURROUND = 1
const MAPPING_TYPE_AMBISONICS = 2

type OpusMSEncoder = struct {
	Flayout            OpusT_ChannelLayout
	Farch              int32
	Flfe_stream        int32
	Fapplication       int32
	FFs                OpusT_opus_int32
	Fvariable_duration int32
	Fmapping_type      OpusT_MappingType
	Fbitrate_bps       OpusT_opus_int32
}

type OpusMSDecoder = struct {
	Flayout OpusT_ChannelLayout
}

type OpusT_opus_copy_channel_in_func = uintptr

type OpusT_opus_copy_channel_out_func = uintptr

type OpusT_downmix_func = uintptr

func Opus_opus_pcm_soft_clip_impl(tls *libc.TLS, _x uintptr, N int32, C int32, declip_mem uintptr, arch int32) {
	var a, delta, maxval, offset, x0, v7, v8, v9 float32
	var all_within_neg1pos1, c, curr, end, i, peak_pos, special, start, v4 int32
	var x uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a, all_within_neg1pos1, c, curr, delta, end, i, maxval, offset, peak_pos, special, start, x, x0, v4, v7, v8, v9
	if C < int32(1) || N < int32(1) || !(_x != 0) || !(declip_mem != 0) {
		return
	}
	/* Clamp everything within the range [-2, +2] which is the domain of the soft
	      clipping non-linearity. Outside the defined range the derivative will be zero,
	      therefore there is no discontinuity introduced here. The implementation
	      might provide a hint if all input samples are within the [-1, +1] range.
	   `opus_limit2_checkwithin1()`:
	      - Clamps all samples within the valid range [-2, +2].
	      - Generic C implementation:
	         * Does not attempt early detection whether samples are within hinted range.
	         * Always returns 0.
	      - Architecture specific implementation:
	         * Uses SIMD instructions to efficiently detect if all samples are
	           within the hinted range [-1, +1].
	         * Returns 1 if no samples exceed the hinted range, 0 otherwise.
	   `all_within_neg1pos1`:
	      - Optimization hint to skip per-sample out-of-bound checks.
	        If true, the check can be skipped. */
	_ = arch
	all_within_neg1pos1 = Opus_opus_limit2_checkwithin1_c(tls, _x, N*C)
	c = 0
	for {
		if !(c < C) {
			break
		}
		x = _x + uintptr(c)*4
		a = *(*float32)(unsafe.Pointer(declip_mem + uintptr(c)*4))
		/* Continue applying the non-linearity from the previous frame to avoid
		   any discontinuity. */
		i = 0
		for {
			if !(i < N) {
				break
			}
			if float32(*(*float32)(unsafe.Pointer(x + uintptr(i*C)*4))*a) >= float32(0) {
				break
			}
			*(*float32)(unsafe.Pointer(x + uintptr(i*C)*4)) = *(*float32)(unsafe.Pointer(x + uintptr(i*C)*4)) + float32(float32(a**(*float32)(unsafe.Pointer(x + uintptr(i*C)*4)))**(*float32)(unsafe.Pointer(x + uintptr(i*C)*4)))
			goto _2
		_2:
			;
			i = i + 1
		}
		curr = 0
		x0 = *(*float32)(unsafe.Pointer(x))
		for int32(1) != 0 {
			special = 0
			/* Detection for early exit can be skipped if hinted by `all_within_neg1pos1` */
			if all_within_neg1pos1 != 0 {
				i = N
			} else {
				i = curr
				for {
					if !(i < N) {
						break
					}
					if *(*float32)(unsafe.Pointer(x + uintptr(i*C)*4)) > float32(1) || *(*float32)(unsafe.Pointer(x + uintptr(i*C)*4)) < float32(-int32(1)) {
						break
					}
					goto _3
				_3:
					;
					i = i + 1
				}
			}
			if i == N {
				a = float32(0)
				break
			}
			peak_pos = i
			v4 = i
			end = v4
			start = v4
			maxval = float32(libc.Xfabs(tls, float64(*(*float32)(unsafe.Pointer(x + uintptr(i*C)*4)))))
			/* Look for first zero crossing before clipping */
			for start > 0 && float32(*(*float32)(unsafe.Pointer(x + uintptr(i*C)*4))**(*float32)(unsafe.Pointer(x + uintptr((start-int32(1))*C)*4))) >= float32(0) {
				start = start - 1
			}
			/* Look for first zero crossing after clipping */
			for end < N && float32(*(*float32)(unsafe.Pointer(x + uintptr(i*C)*4))**(*float32)(unsafe.Pointer(x + uintptr(end*C)*4))) >= float32(0) {
				/* Look for other peaks until the next zero-crossing. */
				if float32(libc.Xfabs(tls, float64(*(*float32)(unsafe.Pointer(x + uintptr(end*C)*4))))) > maxval {
					maxval = float32(libc.Xfabs(tls, float64(*(*float32)(unsafe.Pointer(x + uintptr(end*C)*4)))))
					peak_pos = end
				}
				end = end + 1
			}
			/* Detect the special case where we clip before the first zero crossing */
			special = libc.BoolInt32(start == 0 && float32(*(*float32)(unsafe.Pointer(x + uintptr(i*C)*4))**(*float32)(unsafe.Pointer(x))) >= float32(0))
			/* Compute a such that maxval + a*maxval^2 = 1 */
			a = (maxval - float32(1)) / float32(maxval*maxval)
			/* Slightly boost "a" by 2^-22. This is just enough to ensure -ffast-math
			   does not cause output values larger than +/-1, but small enough not
			   to matter even for 24-bit output.  */
			a = a + float32(a*float32(2.4e-07))
			if *(*float32)(unsafe.Pointer(x + uintptr(i*C)*4)) > float32(0) {
				a = -a
			}
			/* Apply soft clipping */
			i = start
			for {
				if !(i < end) {
					break
				}
				*(*float32)(unsafe.Pointer(x + uintptr(i*C)*4)) = *(*float32)(unsafe.Pointer(x + uintptr(i*C)*4)) + float32(float32(a**(*float32)(unsafe.Pointer(x + uintptr(i*C)*4)))**(*float32)(unsafe.Pointer(x + uintptr(i*C)*4)))
				goto _5
			_5:
				;
				i = i + 1
			}
			if special != 0 && peak_pos >= int32(2) {
				offset = x0 - *(*float32)(unsafe.Pointer(x))
				delta = offset / float32(peak_pos)
				i = curr
				for {
					if !(i < peak_pos) {
						break
					}
					offset = offset - delta
					*(*float32)(unsafe.Pointer(x + uintptr(i*C)*4)) += offset
					if float32(1) < *(*float32)(unsafe.Pointer(x + uintptr(i*C)*4)) {
						v8 = float32(1)
					} else {
						v8 = *(*float32)(unsafe.Pointer(x + uintptr(i*C)*4))
					}
					if -float32(1) > v8 {
						v7 = -float32(1)
					} else {
						if float32(1) < *(*float32)(unsafe.Pointer(x + uintptr(i*C)*4)) {
							v9 = float32(1)
						} else {
							v9 = *(*float32)(unsafe.Pointer(x + uintptr(i*C)*4))
						}
						v7 = v9
					}
					*(*float32)(unsafe.Pointer(x + uintptr(i*C)*4)) = v7
					goto _6
				_6:
					;
					i = i + 1
				}
			}
			curr = end
			if curr == N {
				break
			}
		}
		*(*float32)(unsafe.Pointer(declip_mem + uintptr(c)*4)) = a
		goto _1
	_1:
		;
		c = c + 1
	}
}

func Opus_opus_pcm_soft_clip(tls *libc.TLS, _x uintptr, N int32, C int32, declip_mem uintptr) {
	Opus_opus_pcm_soft_clip_impl(tls, _x, N, C, declip_mem, 0)
}

func Opus_encode_size(tls *libc.TLS, size int32, data uintptr) (r int32) {
	if size < int32(252) {
		*(*uint8)(unsafe.Pointer(data)) = uint8(size)
		return int32(1)
	} else {
		*(*uint8)(unsafe.Pointer(data)) = uint8(int32(252) + size&int32(0x3))
		*(*uint8)(unsafe.Pointer(data + 1)) = uint8((size - int32(*(*uint8)(unsafe.Pointer(data)))) >> int32(2))
		return int32(2)
	}
	return r
}

func parse_size(tls *libc.TLS, data uintptr, len1 OpusT_opus_int32, size uintptr) (r int32) {
	if len1 < int32(1) {
		*(*OpusT_opus_int16)(unsafe.Pointer(size)) = int16(-int32(1))
		return -int32(1)
	} else {
		if int32(*(*uint8)(unsafe.Pointer(data))) < int32(252) {
			*(*OpusT_opus_int16)(unsafe.Pointer(size)) = int16(*(*uint8)(unsafe.Pointer(data)))
			return int32(1)
		} else {
			if len1 < int32(2) {
				*(*OpusT_opus_int16)(unsafe.Pointer(size)) = int16(-int32(1))
				return -int32(1)
			} else {
				*(*OpusT_opus_int16)(unsafe.Pointer(size)) = int16(int32(4)*int32(*(*uint8)(unsafe.Pointer(data + 1))) + int32(*(*uint8)(unsafe.Pointer(data))))
				return int32(2)
			}
		}
	}
	return r
}

func Opus_opus_packet_get_samples_per_frame(tls *libc.TLS, data uintptr, Fs OpusT_opus_int32) (r int32) {
	var audiosize, v1 int32
	_, _ = audiosize, v1
	if int32(*(*uint8)(unsafe.Pointer(data)))&int32(0x80) != 0 {
		audiosize = int32(*(*uint8)(unsafe.Pointer(data))) >> int32(3) & int32(0x3)
		audiosize = Fs << audiosize / int32(400)
	} else {
		if int32(*(*uint8)(unsafe.Pointer(data)))&int32(0x60) == int32(0x60) {
			if int32(*(*uint8)(unsafe.Pointer(data)))&int32(0x08) != 0 {
				v1 = Fs / int32(50)
			} else {
				v1 = Fs / int32(100)
			}
			audiosize = v1
		} else {
			audiosize = int32(*(*uint8)(unsafe.Pointer(data))) >> int32(3) & int32(0x3)
			if audiosize == int32(3) {
				audiosize = Fs * int32(60) / int32(1000)
			} else {
				audiosize = Fs << audiosize / int32(100)
			}
		}
	}
	return audiosize
}

func Opus_opus_packet_parse_impl(tls *libc.TLS, data uintptr, len1 OpusT_opus_int32, self_delimited int32, out_toc uintptr, frames uintptr, size uintptr, payload_offset uintptr, packet_offset uintptr, padding uintptr, padding_len uintptr) (r int32) {
	var bytes, cbr, count, framesize, i, p, tmp, v4 int32
	var ch, toc uint8
	var data0, v1 uintptr
	var last_size, pad OpusT_opus_int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = bytes, cbr, ch, count, data0, framesize, i, last_size, p, pad, tmp, toc, v1, v4
	pad = 0
	data0 = data
	/* Make sure we return NULL/0 on error. */
	if padding != uintptr(uint32(0)) {
		*(*uintptr)(unsafe.Pointer(padding)) = uintptr(uint32(0))
		*(*OpusT_opus_int32)(unsafe.Pointer(padding_len)) = 0
	}
	if size == uintptr(uint32(0)) || len1 < 0 {
		return -int32(1)
	}
	if len1 == 0 {
		return -int32(4)
	}
	framesize = Opus_opus_packet_get_samples_per_frame(tls, data, int32(48000))
	cbr = 0
	v1 = data
	data = data + 1
	toc = *(*uint8)(unsafe.Pointer(v1))
	len1 = len1 - 1
	last_size = len1
	switch int32(toc) & int32(0x3) {
	/* One frame */
	case 0:
		count = int32(1)
		break
		/* Two CBR frames */
		fallthrough
	case int32(1):
		count = int32(2)
		cbr = int32(1)
		if !(self_delimited != 0) {
			if len1&int32(0x1) != 0 {
				return -int32(4)
			}
			last_size = len1 / int32(2)
			/* If last_size doesn't fit in size[0], we'll catch it later */
			*(*OpusT_opus_int16)(unsafe.Pointer(size)) = int16(last_size)
		}
		break
		/* Two VBR frames */
		fallthrough
	case int32(2):
		count = int32(2)
		bytes = parse_size(tls, data, len1, size)
		len1 = len1 - bytes
		if int32(*(*OpusT_opus_int16)(unsafe.Pointer(size))) < 0 || int32(*(*OpusT_opus_int16)(unsafe.Pointer(size))) > len1 {
			return -int32(4)
		}
		data = data + uintptr(bytes)
		last_size = len1 - int32(*(*OpusT_opus_int16)(unsafe.Pointer(size)))
		break
		/* Multiple CBR/VBR frames (from 0 to 120 ms) */
		fallthrough
	default: /*case 3:*/
		if len1 < int32(1) {
			return -int32(4)
		}
		/* Number of frames encoded in bits 0 to 5 */
		v1 = data
		data = data + 1
		ch = *(*uint8)(unsafe.Pointer(v1))
		count = int32(ch) & int32(0x3F)
		if count <= 0 || framesize*count > int32(5760) {
			return -int32(4)
		}
		len1 = len1 - 1
		/* Padding flag is bit 6 */
		if int32(ch)&int32(0x40) != 0 {
			for cond := true; cond; cond = p == int32(255) {
				if len1 <= 0 {
					return -int32(4)
				}
				v1 = data
				data = data + 1
				p = int32(*(*uint8)(unsafe.Pointer(v1)))
				len1 = len1 - 1
				if p == int32(255) {
					v4 = int32(254)
				} else {
					v4 = p
				}
				tmp = v4
				len1 = len1 - tmp
				pad = pad + tmp
			}
		}
		if len1 < 0 {
			return -int32(4)
		}
		/* VBR flag is bit 7 */
		cbr = libc.BoolInt32(!(int32(ch)&int32(0x80) != 0))
		if !(cbr != 0) {
			/* VBR case */
			last_size = len1
			i = 0
			for {
				if !(i < count-int32(1)) {
					break
				}
				bytes = parse_size(tls, data, len1, size+uintptr(i)*2)
				len1 = len1 - bytes
				if int32(*(*OpusT_opus_int16)(unsafe.Pointer(size + uintptr(i)*2))) < 0 || int32(*(*OpusT_opus_int16)(unsafe.Pointer(size + uintptr(i)*2))) > len1 {
					return -int32(4)
				}
				data = data + uintptr(bytes)
				last_size = last_size - (bytes + int32(*(*OpusT_opus_int16)(unsafe.Pointer(size + uintptr(i)*2))))
				goto _5
			_5:
				;
				i = i + 1
			}
			if last_size < 0 {
				return -int32(4)
			}
		} else {
			if !(self_delimited != 0) {
				/* CBR case */
				last_size = len1 / count
				if last_size*count != len1 {
					return -int32(4)
				}
				i = 0
				for {
					if !(i < count-int32(1)) {
						break
					}
					*(*OpusT_opus_int16)(unsafe.Pointer(size + uintptr(i)*2)) = int16(last_size)
					goto _6
				_6:
					;
					i = i + 1
				}
			}
		}
		break
	}
	/* Self-delimited framing has an extra size for the last frame. */
	if self_delimited != 0 {
		bytes = parse_size(tls, data, len1, size+uintptr(count)*2-uintptr(1)*2)
		len1 = len1 - bytes
		if int32(*(*OpusT_opus_int16)(unsafe.Pointer(size + uintptr(count-int32(1))*2))) < 0 || int32(*(*OpusT_opus_int16)(unsafe.Pointer(size + uintptr(count-int32(1))*2))) > len1 {
			return -int32(4)
		}
		data = data + uintptr(bytes)
		/* For CBR packets, apply the size to all the frames. */
		if cbr != 0 {
			if int32(*(*OpusT_opus_int16)(unsafe.Pointer(size + uintptr(count-int32(1))*2)))*count > len1 {
				return -int32(4)
			}
			i = 0
			for {
				if !(i < count-int32(1)) {
					break
				}
				*(*OpusT_opus_int16)(unsafe.Pointer(size + uintptr(i)*2)) = *(*OpusT_opus_int16)(unsafe.Pointer(size + uintptr(count-int32(1))*2))
				goto _7
			_7:
				;
				i = i + 1
			}
		} else {
			if bytes+int32(*(*OpusT_opus_int16)(unsafe.Pointer(size + uintptr(count-int32(1))*2))) > last_size {
				return -int32(4)
			}
		}
	} else {
		/* Because it's not encoded explicitly, it's possible the size of the
		   last packet (or all the packets, for the CBR case) is larger than
		   1275. Reject them here.*/
		if last_size > int32(1275) {
			return -int32(4)
		}
		*(*OpusT_opus_int16)(unsafe.Pointer(size + uintptr(count-int32(1))*2)) = int16(last_size)
	}
	if payload_offset != 0 {
		*(*int32)(unsafe.Pointer(payload_offset)) = int32(int64(data) - int64(data0))
	}
	i = 0
	for {
		if !(i < count) {
			break
		}
		if frames != 0 {
			*(*uintptr)(unsafe.Pointer(frames + uintptr(i)*8)) = data
		}
		data = data + uintptr(*(*OpusT_opus_int16)(unsafe.Pointer(size + uintptr(i)*2)))
		goto _8
	_8:
		;
		i = i + 1
	}
	if padding != uintptr(uint32(0)) {
		*(*uintptr)(unsafe.Pointer(padding)) = data
		*(*OpusT_opus_int32)(unsafe.Pointer(padding_len)) = pad
	}
	if packet_offset != 0 {
		*(*OpusT_opus_int32)(unsafe.Pointer(packet_offset)) = pad + int32(int64(data)-int64(data0))
	}
	if out_toc != 0 {
		*(*uint8)(unsafe.Pointer(out_toc)) = toc
	}
	return count
}

func Opus_opus_packet_parse(tls *libc.TLS, data uintptr, len1 OpusT_opus_int32, out_toc uintptr, frames uintptr, size uintptr, payload_offset uintptr) (r int32) {
	return Opus_opus_packet_parse_impl(tls, data, len1, 0, out_toc, frames, size, payload_offset, uintptr(uint32(0)), uintptr(uint32(0)), uintptr(uint32(0)))
}

const ALLOC_NONE = 0
const BWE_AFTER_LOSS_Q16 = 63570
const CELT_SIG_SCALE1 = 32768
const CLOCKS_PER_SEC = 1000000
const CLOCK_BOOTTIME = 7
const CLOCK_BOOTTIME_ALARM = 9
const CLOCK_MONOTONIC = 1
const CLOCK_MONOTONIC_COARSE = 6
const CLOCK_MONOTONIC_RAW = 4
const CLOCK_PROCESS_CPUTIME_ID = 2
const CLOCK_REALTIME = 0
const CLOCK_REALTIME_ALARM = 8
const CLOCK_REALTIME_COARSE = 5
const CLOCK_SGI_CYCLE = 10
const CLOCK_TAI = 11
const CLOCK_THREAD_CPUTIME_ID = 3
const CLONE_CHILD_CLEARTID = 0x00200000
const CLONE_CHILD_SETTID = 0x01000000
const CLONE_DETACHED = 0x00400000
const CLONE_FILES = 0x00000400
const CLONE_FS = 0x00000200
const CLONE_IO = 0x80000000
const CLONE_NEWCGROUP = 0x02000000
const CLONE_NEWIPC = 0x08000000
const CLONE_NEWNET = 0x40000000
const CLONE_NEWNS = 0x00020000
const CLONE_NEWPID = 0x20000000
const CLONE_NEWTIME = 0x00000080
const CLONE_NEWUSER = 0x10000000
const CLONE_NEWUTS = 0x04000000
const CLONE_PARENT = 0x00008000
const CLONE_PARENT_SETTID = 0x00100000
const CLONE_PIDFD = 0x00001000
const CLONE_PTRACE = 0x00002000
const CLONE_SETTLS = 0x00080000
const CLONE_SIGHAND = 0x00000800
const CLONE_SYSVSEM = 0x00040000
const CLONE_THREAD = 0x00010000
const CLONE_UNTRACED = 0x00800000
const CLONE_VFORK = 0x00004000
const CLONE_VM = 0x00000100
const CNG_BUF_MASK_MAX = 255
const CNG_GAIN_SMTH_Q16 = 4634
const CNG_GAIN_SMTH_THRESHOLD_Q16 = 46396
const CNG_NLSF_SMTH_Q16 = 16348
const CODE_CONDITIONALLY = 2
const CODE_INDEPENDENTLY = 0
const CODE_INDEPENDENTLY_NO_LTP_SCALING = 1
const COEF_ONE1 = 1
const CPU_SETSIZE = 1024
const CSIGNAL = 0x000000ff
const DBL_DECIMAL_DIG = 17
const DBL_DIG = 15
const DBL_EPSILON = 2.22044604925031308085e-16
const DBL_HAS_SUBNORM = 1
const DBL_MANT_DIG = 53
const DBL_MAX = 1.79769313486231570815e+308
const DBL_MAX_10_EXP = 308
const DBL_MAX_EXP = 1024
const DBL_MIN = 2.22507385850720138309e-308
const DBL_TRUE_MIN = 4.94065645841246544177e-324
const DECIMAL_DIG = 17
const DECISION_DELAY = 40
const DECODER_NUM_CHANNELS = 2
const DEC_PITCH_BUF_SIZE = 2048
const DTX_ACTIVITY_THRESHOLD = "0.1f"
const ENCODER_NUM_CHANNELS = 2
const FLAG_DECODE_LBRR = 2
const FLAG_DECODE_NORMAL = 0
const FLAG_PACKET_LOST = 1
const FLT_DECIMAL_DIG = 9
const FLT_DIG = 6
const FLT_EPSILON = 1.1920928955078125e-07
const FLT_EVAL_METHOD = 0
const FLT_HAS_SUBNORM = 1
const FLT_MANT_DIG = 24
const FLT_MAX = 3.40282346638528859812e+38
const FLT_MAX_10_EXP = 38
const FLT_MAX_EXP = 128
const FLT_MIN = 1.17549435082228750797e-38
const FLT_RADIX = 2
const FLT_TRUE_MIN = 1.40129846432481707092e-45
const HARM_SHAPE_FIR_TAPS = 3
const LA_PITCH_MS = 2
const LA_SHAPE_MS = 5
const LBRR_MB_MIN_RATE_BPS = 14000
const LBRR_NB_MIN_RATE_BPS = 12000
const LBRR_WB_MIN_RATE_BPS = 16000
const LDBL_DECIMAL_DIG = "DECIMAL_DIG"
const LDBL_DIG = 15
const LDBL_EPSILON = 2.22044604925031308085e-16
const LDBL_HAS_SUBNORM = 1
const LDBL_MANT_DIG = 53
const LDBL_MAX = 1.79769313486231570815e+308
const LDBL_MAX_10_EXP = 308
const LDBL_MAX_EXP = 1024
const LDBL_MIN = 2.22507385850720138309e-308
const LDBL_TRUE_MIN = 4.94065645841246544177e-324
const LOG2_SHELL_CODEC_FRAME_LENGTH = 4
const LSF_COS_TAB_SZ_FIX = 128
const LTP_BUF_LENGTH = 512
const LTP_MEM_LENGTH_MS = 20
const LTP_ORDER = 5
const MAX_API_FS_KHZ = 48
const MAX_CONSECUTIVE_DTX = 20
const MAX_DELTA_GAIN_QUANT = 36
const MAX_DEL_DEC_STATES = 4
const MAX_FIND_PITCH_LPC_ORDER = 16
const MAX_FRAMES_PER_PACKET = 3
const MAX_FS_KHZ = 16
const MAX_LPC_ORDER = 16
const MAX_LPC_STABILIZE_ITERATIONS = 16
const MAX_MATRIX_SIZE = "MAX_LPC_ORDER"
const MAX_NB_SUBFR = 4
const MAX_PERIOD = 1024
const MAX_PREDICTION_POWER_GAIN = "1e4f"
const MAX_PREDICTION_POWER_GAIN_AFTER_RESET = "1e2f"
const MAX_QGAIN_DB = 88
const MAX_SHAPE_LPC_ORDER = 24
const MAX_TARGET_RATE_BPS = 80000
const MIN_LPC_ORDER = 10
const MIN_QGAIN_DB = 2
const MIN_TARGET_RATE_BPS = 5000
const NB_LTP_CBKS = 3
const NB_SPEECH_FRAMES_BEFORE_DTX = 10
const NLSF_QUANT_DEL_DEC_STATES_LOG2 = 2
const NLSF_QUANT_LEVEL_ADJ = 0.1
const NLSF_QUANT_MAX_AMPLITUDE = 4
const NLSF_QUANT_MAX_AMPLITUDE_EXT = 10
const NLSF_VQ_MAX_VECTORS = 32
const NLSF_W_Q = 2
const NSQ_LPC_BUF_LENGTH = "MAX_LPC_ORDER"
const N_LEVELS_QGAIN = 64
const N_RATE_LEVELS = 10
const OFFSET_UVH_Q10 = 240
const OFFSET_UVL_Q10 = 100
const OFFSET_VH_Q10 = 100
const OFFSET_VL_Q10 = 32
const OPTIONAL_CLIP = 1
const OPUS_CCGO_PSEUDOSTACK_KEY = 1869641075
const OPUS_DECODER_RESET_START = "stream_channels"
const OPUS_FPRINTF = "void"
const PITCH_EST_MAX_LAG_MS = 18
const PITCH_EST_MIN_LAG_MS = 2
const PTHREAD_CANCEL_ASYNCHRONOUS = 1
const PTHREAD_CANCEL_DEFERRED = 0
const PTHREAD_CANCEL_DISABLE = 1
const PTHREAD_CANCEL_ENABLE = 0
const PTHREAD_CANCEL_MASKED = 2
const PTHREAD_CREATE_DETACHED = 1
const PTHREAD_CREATE_JOINABLE = 0
const PTHREAD_EXPLICIT_SCHED = 1
const PTHREAD_INHERIT_SCHED = 0
const PTHREAD_MUTEX_DEFAULT = 0
const PTHREAD_MUTEX_ERRORCHECK = 2
const PTHREAD_MUTEX_NORMAL = 0
const PTHREAD_MUTEX_RECURSIVE = 1
const PTHREAD_MUTEX_ROBUST = 1
const PTHREAD_MUTEX_STALLED = 0
const PTHREAD_ONCE_INIT = 0
const PTHREAD_PRIO_INHERIT = 1
const PTHREAD_PRIO_NONE = 0
const PTHREAD_PRIO_PROTECT = 2
const PTHREAD_PROCESS_PRIVATE = 0
const PTHREAD_PROCESS_SHARED = 1
const PTHREAD_SCOPE_PROCESS = 1
const PTHREAD_SCOPE_SYSTEM = 0
const QUANT_LEVEL_ADJUST_Q10 = 80
const RAND_INCREMENT = 907633515
const RAND_MULTIPLIER = 196314165
const SCHED_BATCH = 3
const SCHED_DEADLINE = 6
const SCHED_FIFO = 1
const SCHED_IDLE = 5
const SCHED_OTHER = 0
const SCHED_RESET_ON_FORK = 0x40000000
const SCHED_RR = 2
const SHELL_CODEC_FRAME_LENGTH = 16
const SILK_DECODER_STATE_RESET_START = "prev_gain_Q16"
const SILK_MAX_FRAMES_PER_PACKET = 3
const SILK_MAX_ORDER_LPC = 24
const SILK_MAX_PULSES = 16
const SILK_NO_ERROR = 0
const SILK_RESAMPLER_MAX_FIR_ORDER = 36
const SILK_RESAMPLER_MAX_IIR_ORDER = 6
const STEREO_INTERP_LEN_MS = 8
const STEREO_QUANT_SUB_STEPS = 5
const STEREO_QUANT_TAB_SIZE = 16
const STEREO_RATIO_SMOOTH_COEF = 0.01
const SUB_FRAME_LENGTH_MS = 5
const TIMER_ABSTIME = 1
const TIME_UTC = 1
const TRANSITION_INT_NUM = 5
const TRANSITION_NA = 2
const TRANSITION_NB = 3
const TRANSITION_TIME_MS = 5120
const TYPE_NO_VOICE_ACTIVITY = 0
const TYPE_UNVOICED = 1
const TYPE_VOICED = 2
const USE_HARM_SHAPING = 1
const VAD_ACTIVITY = 1
const VAD_INTERNAL_SUBFRAMES_LOG2 = 2
const VAD_NEGATIVE_OFFSET_Q5 = 128
const VAD_NOISE_LEVELS_BIAS = 50
const VAD_NOISE_LEVEL_SMOOTH_COEF_Q16 = 1024
const VAD_NO_ACTIVITY = 0
const VAD_N_BANDS = 4
const VAD_SNR_FACTOR_Q16 = 45000
const VAD_SNR_SMOOTH_COEF_Q18 = 4096
const _ISOC99_SOURCE = 1
const _ISOC9X_SOURCE = 1
const __USE_ISOC99 = 1
const __USE_ISOC9X = 1
const __tm_gmtoff = "tm_gmtoff"
const __tm_zone = "tm_zone"
const silk_FALSE = 0
const silk_LIMIT_16 = "silk_LIMIT"
const silk_LIMIT_32 = "silk_LIMIT"
const silk_LIMIT_int = "silk_LIMIT"
const silk_TRUE = 1
const silk_float = "float"
const silk_float_MAX = "FLT_MAX"
const silk_int16_MAX = 0x7FFF
const silk_int32_MAX = 2147483647
const silk_int8_MAX = 0x7F
const silk_uint8_MAX = 0xFF

type OpusT_OpusCustomMode = struct {
	FFs             OpusT_opus_int32
	Foverlap        int32
	FnbEBands       int32
	FeffEBands      int32
	Fpreemph        [4]OpusT_opus_val16
	FeBands         uintptr
	FmaxLM          int32
	FnbShortMdcts   int32
	FshortMdctSize  int32
	FnbAllocVectors int32
	FallocVectors   uintptr
	FlogN           uintptr
	Fwindow         uintptr
	Fmdct           OpusT_mdct_lookup
	Fcache          OpusT_PulseCache
}

var trim_icdf1 = [11]uint8{
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
var spread_icdf1 = [4]uint8{
	0: uint8(25),
	1: uint8(23),
	2: uint8(2),
}
var tapset_icdf1 = [3]uint8{
	0: uint8(2),
	1: uint8(1),
}

type OpusT_OpusDecoder = struct {
	Fcelt_dec_offset      int32
	Fsilk_dec_offset      int32
	Fchannels             int32
	FFs                   OpusT_opus_int32
	FDecControl           OpusT_silk_DecControlStruct
	Fdecode_gain          int32
	Fcomplexity           int32
	Fignore_extensions    int32
	Farch                 int32
	Fstream_channels      int32
	Fbandwidth            int32
	Fmode                 int32
	Fprev_mode            int32
	Fframe_size           int32
	Fprev_redundancy      int32
	Flast_packet_duration int32
	Fsoftclip_mem         [2]OpusT_opus_val16
	FrangeFinal           OpusT_opus_uint32
}

type OpusT_OpusDREDDecoder = struct {
	Floaded int32
	Farch   int32
	Fmagic  OpusT_opus_uint32
}

type OpusT_mdct_lookup = struct {
	Fn        int32
	Fmaxshift int32
	Fkfft     [4]uintptr
	Ftrig     uintptr
}

type OpusT_PulseCache = struct {
	Fsize  int32
	Findex uintptr
	Fbits  uintptr
	Fcaps  uintptr
}

type OpusCustomMode = struct {
	FFs             OpusT_opus_int32
	Foverlap        int32
	FnbEBands       int32
	FeffEBands      int32
	Fpreemph        [4]OpusT_opus_val16
	FeBands         uintptr
	FmaxLM          int32
	FnbShortMdcts   int32
	FshortMdctSize  int32
	FnbAllocVectors int32
	FallocVectors   uintptr
	FlogN           uintptr
	Fwindow         uintptr
	Fmdct           OpusT_mdct_lookup
	Fcache          OpusT_PulseCache
}

type OpusT_silk_EncControlStruct = struct {
	FnChannelsAPI              OpusT_opus_int32
	FnChannelsInternal         OpusT_opus_int32
	FAPI_sampleRate            OpusT_opus_int32
	FmaxInternalSampleRate     OpusT_opus_int32
	FminInternalSampleRate     OpusT_opus_int32
	FdesiredInternalSampleRate OpusT_opus_int32
	FpayloadSize_ms            int32
	FbitRate                   OpusT_opus_int32
	FpacketLossPercentage      int32
	Fcomplexity                int32
	FuseInBandFEC              int32
	FuseDRED                   int32
	FLBRR_coded                int32
	FuseDTX                    int32
	FuseCBR                    int32
	FmaxBits                   int32
	FtoMono                    int32
	FopusCanSwitch             int32
	FreducedDependency         int32
	FinternalSampleRate        OpusT_opus_int32
	FallowBandwidthSwitch      int32
	FinWBmodeWithoutVariableLP int32
	FstereoWidth_Q14           int32
	FswitchReady               int32
	FsignalType                int32
	Foffset                    int32
}

type OpusT_silk_DecControlStruct = struct {
	FnChannelsAPI       OpusT_opus_int32
	FnChannelsInternal  OpusT_opus_int32
	FAPI_sampleRate     OpusT_opus_int32
	FinternalSampleRate OpusT_opus_int32
	FpayloadSize_ms     int32
	FprevPitchLag       int32
	Fenable_deep_plc    int32
}

type OpusT_silk_TOC_struct = struct {
	FVADFlag       int32
	FVADFlags      [3]int32
	FinbandFECFlag int32
}

type OpusT_time_t = int64

type OpusT_clockid_t = int32

type timespec = struct {
	Ftv_sec  OpusT_time_t
	Ftv_nsec int64
}

type OpusT_pthread_t = uintptr

type OpusT_pthread_once_t = int32

type OpusT_pthread_key_t = uint32

type OpusT_pthread_spinlock_t = int32

type OpusT_pthread_mutexattr_t = struct {
	F__attr uint32
}

type OpusT_pthread_condattr_t = struct {
	F__attr uint32
}

type OpusT_pthread_barrierattr_t = struct {
	F__attr uint32
}

type OpusT_pthread_rwlockattr_t = struct {
	F__attr [2]uint32
}

type OpusT_sigset_t = struct {
	F__bits [16]uint64
}

type __sigset_t = OpusT_sigset_t

type OpusT_pthread_attr_t = struct {
	F__u struct {
		F__vi [0][14]int32
		F__s  [0][7]uint64
		F__i  [14]int32
	}
}

type OpusT_pthread_mutex_t = struct {
	F__u struct {
		F__vi [0][10]int32
		F__p  [0][5]uintptr
		F__i  [10]int32
	}
}

type OpusT_pthread_cond_t = struct {
	F__u struct {
		F__vi [0][12]int32
		F__p  [0][6]uintptr
		F__i  [12]int32
	}
}

type OpusT_pthread_rwlock_t = struct {
	F__u struct {
		F__vi [0][14]int32
		F__p  [0][7]uintptr
		F__i  [14]int32
	}
}

type OpusT_pthread_barrier_t = struct {
	F__u struct {
		F__vi [0][8]int32
		F__p  [0][4]uintptr
		F__i  [8]int32
	}
}

type OpusT_pid_t = int32

type sched_param = struct {
	Fsched_priority int32
	F__reserved1    int32
	F__reserved2    [2]struct {
		F__reserved1 OpusT_time_t
		F__reserved2 int64
	}
	F__reserved3 int32
}

type OpusT_cpu_set_t = struct {
	F__bits [16]uint64
}

type OpusT_timer_t = uintptr

type OpusT_clock_t = int64

type tm = struct {
	Ftm_sec    int32
	Ftm_min    int32
	Ftm_hour   int32
	Ftm_mday   int32
	Ftm_mon    int32
	Ftm_year   int32
	Ftm_wday   int32
	Ftm_yday   int32
	Ftm_isdst  int32
	Ftm_gmtoff int64
	Ftm_zone   uintptr
}

type itimerspec = struct {
	Fit_interval timespec
	Fit_value    timespec
}

type __ptcb = struct {
	F__f    uintptr
	F__x    uintptr
	F__next uintptr
}

type cpu_set_t = struct {
	F__bits [16]uint64
}

type OpusT_opus_ccgo_pseudostack_state = struct {
	Fscratch_ptr  uintptr
	Fglobal_stack uintptr
}

type OpusT_silk_resampler_state_struct = struct {
	FsIIR [6]OpusT_opus_int32
	FsFIR struct {
		Fi16 [0][36]OpusT_opus_int16
		Fi32 [36]OpusT_opus_int32
	}
	FdelayBuf           [96]OpusT_opus_int16
	Fresampler_function int32
	FbatchSize          int32
	FinvRatio_Q16       OpusT_opus_int32
	FFIR_Order          int32
	FFIR_Fracs          int32
	FFs_in_kHz          int32
	FFs_out_kHz         int32
	FinputDelay         int32
	FCoefs              uintptr
}

type _silk_resampler_state_struct = OpusT_silk_resampler_state_struct

type OpusT_silk_nsq_state = struct {
	Fxq               [640]OpusT_opus_int16
	FsLTP_shp_Q14     [640]OpusT_opus_int32
	FsLPC_Q14         [96]OpusT_opus_int32
	FsAR2_Q14         [24]OpusT_opus_int32
	FsLF_AR_shp_Q14   OpusT_opus_int32
	FsDiff_shp_Q14    OpusT_opus_int32
	FlagPrev          int32
	FsLTP_buf_idx     int32
	FsLTP_shp_buf_idx int32
	Frand_seed        OpusT_opus_int32
	Fprev_gain_Q16    OpusT_opus_int32
	Frewhite_flag     int32
}

type OpusT_silk_VAD_state = struct {
	FAnaState        [2]OpusT_opus_int32
	FAnaState1       [2]OpusT_opus_int32
	FAnaState2       [2]OpusT_opus_int32
	FXnrgSubfr       [4]OpusT_opus_int32
	FNrgRatioSmth_Q8 [4]OpusT_opus_int32
	FHPstate         OpusT_opus_int16
	FNL              [4]OpusT_opus_int32
	Finv_NL          [4]OpusT_opus_int32
	FNoiseLevelBias  [4]OpusT_opus_int32
	Fcounter         OpusT_opus_int32
}

type OpusT_silk_LP_state = struct {
	FIn_LP_State         [2]OpusT_opus_int32
	Ftransition_frame_no OpusT_opus_int32
	Fmode                int32
	Fsaved_fs_kHz        OpusT_opus_int32
}

type OpusT_silk_NLSF_CB_struct = struct {
	FnVectors            OpusT_opus_int16
	Forder               OpusT_opus_int16
	FquantStepSize_Q16   OpusT_opus_int16
	FinvQuantStepSize_Q6 OpusT_opus_int16
	FCB1_NLSF_Q8         uintptr
	FCB1_Wght_Q9         uintptr
	FCB1_iCDF            uintptr
	Fpred_Q8             uintptr
	Fec_sel              uintptr
	Fec_iCDF             uintptr
	Fec_Rates_Q5         uintptr
	FdeltaMin_Q15        uintptr
}

type OpusT_stereo_enc_state = struct {
	Fpred_prev_Q13   [2]OpusT_opus_int16
	FsMid            [2]OpusT_opus_int16
	FsSide           [2]OpusT_opus_int16
	Fmid_side_amp_Q0 [4]OpusT_opus_int32
	Fsmth_width_Q14  OpusT_opus_int16
	Fwidth_prev_Q14  OpusT_opus_int16
	Fsilent_side_len OpusT_opus_int16
	FpredIx          [3][2][3]OpusT_opus_int8
	Fmid_only_flags  [3]OpusT_opus_int8
}

type OpusT_stereo_dec_state = struct {
	Fpred_prev_Q13 [2]OpusT_opus_int16
	FsMid          [2]OpusT_opus_int16
	FsSide         [2]OpusT_opus_int16
}

type OpusT_SideInfoIndices = struct {
	FGainsIndices      [4]OpusT_opus_int8
	FLTPIndex          [4]OpusT_opus_int8
	FNLSFIndices       [17]OpusT_opus_int8
	FlagIndex          OpusT_opus_int16
	FcontourIndex      OpusT_opus_int8
	FsignalType        OpusT_opus_int8
	FquantOffsetType   OpusT_opus_int8
	FNLSFInterpCoef_Q2 OpusT_opus_int8
	FPERIndex          OpusT_opus_int8
	FLTP_scaleIndex    OpusT_opus_int8
	FSeed              OpusT_opus_int8
}

type OpusT_silk_encoder_state = struct {
	FIn_HP_State                   [2]OpusT_opus_int32
	Fvariable_HP_smth1_Q15         OpusT_opus_int32
	Fvariable_HP_smth2_Q15         OpusT_opus_int32
	FsLP                           OpusT_silk_LP_state
	FsVAD                          OpusT_silk_VAD_state
	FsNSQ                          OpusT_silk_nsq_state
	Fprev_NLSFq_Q15                [16]OpusT_opus_int16
	Fspeech_activity_Q8            int32
	Fallow_bandwidth_switch        int32
	FLBRRprevLastGainIndex         OpusT_opus_int8
	FprevSignalType                OpusT_opus_int8
	FprevLag                       int32
	Fpitch_LPC_win_length          int32
	Fmax_pitch_lag                 int32
	FAPI_fs_Hz                     OpusT_opus_int32
	Fprev_API_fs_Hz                OpusT_opus_int32
	FmaxInternal_fs_Hz             int32
	FminInternal_fs_Hz             int32
	FdesiredInternal_fs_Hz         int32
	Ffs_kHz                        int32
	Fnb_subfr                      int32
	Fframe_length                  int32
	Fsubfr_length                  int32
	Fltp_mem_length                int32
	Fla_pitch                      int32
	Fla_shape                      int32
	FshapeWinLength                int32
	FTargetRate_bps                OpusT_opus_int32
	FPacketSize_ms                 int32
	FPacketLoss_perc               int32
	FframeCounter                  OpusT_opus_int32
	FComplexity                    int32
	FnStatesDelayedDecision        int32
	FuseInterpolatedNLSFs          int32
	FshapingLPCOrder               int32
	FpredictLPCOrder               int32
	FpitchEstimationComplexity     int32
	FpitchEstimationLPCOrder       int32
	FpitchEstimationThreshold_Q16  OpusT_opus_int32
	Fsum_log_gain_Q7               OpusT_opus_int32
	FNLSF_MSVQ_Survivors           int32
	Ffirst_frame_after_reset       int32
	Fcontrolled_since_last_payload int32
	Fwarping_Q16                   int32
	FuseCBR                        int32
	FprefillFlag                   int32
	Fpitch_lag_low_bits_iCDF       uintptr
	Fpitch_contour_iCDF            uintptr
	FpsNLSF_CB                     uintptr
	Finput_quality_bands_Q15       [4]int32
	Finput_tilt_Q15                int32
	FSNR_dB_Q7                     int32
	FVAD_flags                     [3]OpusT_opus_int8
	FLBRR_flag                     OpusT_opus_int8
	FLBRR_flags                    [3]int32
	Findices                       OpusT_SideInfoIndices
	Fpulses                        [320]OpusT_opus_int8
	Farch                          int32
	FinputBuf                      [322]OpusT_opus_int16
	FinputBufIx                    int32
	FnFramesPerPacket              int32
	FnFramesEncoded                int32
	FnChannelsAPI                  int32
	FnChannelsInternal             int32
	FchannelNb                     int32
	Fframes_since_onset            int32
	Fec_prevSignalType             int32
	Fec_prevLagIndex               OpusT_opus_int16
	Fresampler_state               OpusT_silk_resampler_state_struct
	FuseDTX                        int32
	FinDTX                         int32
	FnoSpeechCounter               int32
	FuseInBandFEC                  int32
	FLBRR_enabled                  int32
	FLBRR_GainIncreases            int32
	Findices_LBRR                  [3]OpusT_SideInfoIndices
	Fpulses_LBRR                   [3][320]OpusT_opus_int8
}

type OpusT_silk_PLC_struct = struct {
	FpitchL_Q8         OpusT_opus_int32
	FLTPCoef_Q14       [5]OpusT_opus_int16
	FprevLPC_Q12       [16]OpusT_opus_int16
	Flast_frame_lost   int32
	Frand_seed         OpusT_opus_int32
	FrandScale_Q14     OpusT_opus_int16
	Fconc_energy       OpusT_opus_int32
	Fconc_energy_shift int32
	FprevLTP_scale_Q14 OpusT_opus_int16
	FprevGain_Q16      [2]OpusT_opus_int32
	Ffs_kHz            int32
	Fnb_subfr          int32
	Fsubfr_length      int32
	Fenable_deep_plc   int32
}

type OpusT_silk_CNG_struct = struct {
	FCNG_exc_buf_Q14   [320]OpusT_opus_int32
	FCNG_smth_NLSF_Q15 [16]OpusT_opus_int16
	FCNG_synth_state   [16]OpusT_opus_int32
	FCNG_smth_Gain_Q16 OpusT_opus_int32
	Frand_seed         OpusT_opus_int32
	Ffs_kHz            int32
}

type OpusT_silk_decoder_state = struct {
	Fprev_gain_Q16           OpusT_opus_int32
	Fexc_Q14                 [320]OpusT_opus_int32
	FsLPC_Q14_buf            [16]OpusT_opus_int32
	FoutBuf                  [480]OpusT_opus_int16
	FlagPrev                 int32
	FLastGainIndex           OpusT_opus_int8
	Ffs_kHz                  int32
	Ffs_API_hz               OpusT_opus_int32
	Fnb_subfr                int32
	Fframe_length            int32
	Fsubfr_length            int32
	Fltp_mem_length          int32
	FLPC_order               int32
	FprevNLSF_Q15            [16]OpusT_opus_int16
	Ffirst_frame_after_reset int32
	Fpitch_lag_low_bits_iCDF uintptr
	Fpitch_contour_iCDF      uintptr
	FnFramesDecoded          int32
	FnFramesPerPacket        int32
	Fec_prevSignalType       int32
	Fec_prevLagIndex         OpusT_opus_int16
	FVAD_flags               [3]int32
	FLBRR_flag               int32
	FLBRR_flags              [3]int32
	Fresampler_state         OpusT_silk_resampler_state_struct
	FpsNLSF_CB               uintptr
	Findices                 OpusT_SideInfoIndices
	FsCNG                    OpusT_silk_CNG_struct
	FlossCnt                 int32
	FprevSignalType          int32
	Farch                    int32
	FsPLC                    OpusT_silk_PLC_struct
}

type OpusT_silk_decoder_control = struct {
	FpitchL        [4]int32
	FGains_Q16     [4]OpusT_opus_int32
	FPredCoef_Q12  [2][16]OpusT_opus_int16
	FLTPCoef_Q14   [20]OpusT_opus_int16
	FLTP_scale_Q14 int32
}

var log2_x_norm_coeff1 = [8]float32{
	0: float32(1),
	1: float32(0.8888888955116272),
	2: float32(0.8),
	3: float32(0.7272727489471436),
	4: float32(0.6666666865348816),
	5: float32(0.6153846383094788),
	6: float32(0.5714285969734192),
	7: float32(0.5333333611488342),
}
var log2_y_norm_coeff1 = [8]float32{
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

type OpusDecoder = struct {
	Fcelt_dec_offset      int32
	Fsilk_dec_offset      int32
	Fchannels             int32
	FFs                   OpusT_opus_int32
	FDecControl           OpusT_silk_DecControlStruct
	Fdecode_gain          int32
	Fcomplexity           int32
	Fignore_extensions    int32
	Farch                 int32
	Fstream_channels      int32
	Fbandwidth            int32
	Fmode                 int32
	Fprev_mode            int32
	Fframe_size           int32
	Fprev_redundancy      int32
	Flast_packet_duration int32
	Fsoftclip_mem         [2]OpusT_opus_val16
	FrangeFinal           OpusT_opus_uint32
}

func validate_opus_decoder(tls *libc.TLS, st uintptr) {
	if !((*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fchannels == int32(1) || (*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fchannels == int32(2)) {
		Opus_celt_fatal(tls, __ccgo_ts, __ccgo_ts+57, int32(99))
	}
	if !((*OpusT_OpusDecoder)(unsafe.Pointer(st)).FFs == int32(48000) || (*OpusT_OpusDecoder)(unsafe.Pointer(st)).FFs == int32(24000) || (*OpusT_OpusDecoder)(unsafe.Pointer(st)).FFs == int32(16000) || (*OpusT_OpusDecoder)(unsafe.Pointer(st)).FFs == int32(12000) || (*OpusT_OpusDecoder)(unsafe.Pointer(st)).FFs == int32(8000)) {
		Opus_celt_fatal(tls, __ccgo_ts+79, __ccgo_ts+57, int32(103))
	}
	if !((*OpusT_OpusDecoder)(unsafe.Pointer(st)).FDecControl.FAPI_sampleRate == (*OpusT_OpusDecoder)(unsafe.Pointer(st)).FFs) {
		Opus_celt_fatal(tls, __ccgo_ts+188, __ccgo_ts+57, int32(105))
	}
	if !((*OpusT_OpusDecoder)(unsafe.Pointer(st)).FDecControl.FinternalSampleRate == 0 || (*OpusT_OpusDecoder)(unsafe.Pointer(st)).FDecControl.FinternalSampleRate == int32(16000) || (*OpusT_OpusDecoder)(unsafe.Pointer(st)).FDecControl.FinternalSampleRate == int32(12000) || (*OpusT_OpusDecoder)(unsafe.Pointer(st)).FDecControl.FinternalSampleRate == int32(8000)) {
		Opus_celt_fatal(tls, __ccgo_ts+246, __ccgo_ts+57, int32(106))
	}
	if !((*OpusT_OpusDecoder)(unsafe.Pointer(st)).FDecControl.FnChannelsAPI == (*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fchannels) {
		Opus_celt_fatal(tls, __ccgo_ts+440, __ccgo_ts+57, int32(107))
	}
	if !((*OpusT_OpusDecoder)(unsafe.Pointer(st)).FDecControl.FnChannelsInternal == 0 || (*OpusT_OpusDecoder)(unsafe.Pointer(st)).FDecControl.FnChannelsInternal == int32(1) || (*OpusT_OpusDecoder)(unsafe.Pointer(st)).FDecControl.FnChannelsInternal == int32(2)) {
		Opus_celt_fatal(tls, __ccgo_ts+502, __ccgo_ts+57, int32(108))
	}
	if !((*OpusT_OpusDecoder)(unsafe.Pointer(st)).FDecControl.FpayloadSize_ms == 0 || (*OpusT_OpusDecoder)(unsafe.Pointer(st)).FDecControl.FpayloadSize_ms == int32(10) || (*OpusT_OpusDecoder)(unsafe.Pointer(st)).FDecControl.FpayloadSize_ms == int32(20) || (*OpusT_OpusDecoder)(unsafe.Pointer(st)).FDecControl.FpayloadSize_ms == int32(40) || (*OpusT_OpusDecoder)(unsafe.Pointer(st)).FDecControl.FpayloadSize_ms == int32(60)) {
		Opus_celt_fatal(tls, __ccgo_ts+640, __ccgo_ts+57, int32(109))
	}
	if !((*OpusT_OpusDecoder)(unsafe.Pointer(st)).Farch >= int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+849, __ccgo_ts+57, int32(111))
	}
	if !((*OpusT_OpusDecoder)(unsafe.Pointer(st)).Farch <= int32(OPUS_ARCHMASK)) {
		Opus_celt_fatal(tls, __ccgo_ts+881, __ccgo_ts+57, int32(112))
	}
	if !((*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fstream_channels == int32(1) || (*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fstream_channels == int32(2)) {
		Opus_celt_fatal(tls, __ccgo_ts+925, __ccgo_ts+57, int32(114))
	}
}

func Opus_opus_decoder_get_size(tls *libc.TLS, channels int32) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var alignment uint32
	var celtDecSizeBytes, ret, v1 int32
	var _ /* silkDecSizeBytes at bp+0 */ int32
	_, _, _, _ = alignment, celtDecSizeBytes, ret, v1
	if channels < int32(1) || channels > int32(2) {
		return 0
	}
	ret = Opus_silk_Get_Decoder_Size(tls, bp)
	if ret != 0 {
		return 0
	}
	alignment = uint32(uint64(uintptr(uint32(0)) + 8))
	v1 = int32((uint32(*(*int32)(unsafe.Pointer(bp))) + alignment - uint32(1)) / alignment * alignment)
	goto _2
_2:
	*(*int32)(unsafe.Pointer(bp)) = v1
	celtDecSizeBytes = Opus_celt_decoder_get_size(tls, channels)
	alignment = uint32(uint64(uintptr(uint32(0)) + 8))
	v1 = int32((uint32(int32(100)) + alignment - uint32(1)) / alignment * alignment)
	goto _4
_4:
	return v1 + *(*int32)(unsafe.Pointer(bp)) + celtDecSizeBytes
}

func Opus_opus_decoder_init(tls *libc.TLS, st uintptr, Fs OpusT_opus_int32, channels int32) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var alignment uint32
	var celt_dec, silk_dec uintptr
	var ret, v1 int32
	var _ /* silkDecSizeBytes at bp+0 */ int32
	_, _, _, _, _ = alignment, celt_dec, ret, silk_dec, v1
	if Fs != int32(48000) && Fs != int32(24000) && Fs != int32(16000) && Fs != int32(12000) && Fs != int32(8000) || channels != int32(1) && channels != int32(2) {
		return -int32(1)
	}
	libc.Xmemset(tls, st, 0, uint64(uint32(Opus_opus_decoder_get_size(tls, channels)))*uint64(1))
	/* Initialize SILK decoder */
	ret = Opus_silk_Get_Decoder_Size(tls, bp)
	if ret != 0 {
		return -int32(3)
	}
	alignment = uint32(uint64(uintptr(uint32(0)) + 8))
	v1 = int32((uint32(*(*int32)(unsafe.Pointer(bp))) + alignment - uint32(1)) / alignment * alignment)
	goto _2
_2:
	*(*int32)(unsafe.Pointer(bp)) = v1
	alignment = uint32(uint64(uintptr(uint32(0)) + 8))
	v1 = int32((uint32(int32(100)) + alignment - uint32(1)) / alignment * alignment)
	goto _4
_4:
	(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fsilk_dec_offset = v1
	(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fcelt_dec_offset = (*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fsilk_dec_offset + *(*int32)(unsafe.Pointer(bp))
	silk_dec = st + uintptr((*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fsilk_dec_offset)
	celt_dec = st + uintptr((*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fcelt_dec_offset)
	v1 = channels
	(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fchannels = v1
	(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fstream_channels = v1
	(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fcomplexity = 0
	(*OpusT_OpusDecoder)(unsafe.Pointer(st)).FFs = Fs
	(*OpusT_OpusDecoder)(unsafe.Pointer(st)).FDecControl.FAPI_sampleRate = (*OpusT_OpusDecoder)(unsafe.Pointer(st)).FFs
	(*OpusT_OpusDecoder)(unsafe.Pointer(st)).FDecControl.FnChannelsAPI = (*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fchannels
	/* Reset decoder */
	ret = Opus_silk_InitDecoder(tls, silk_dec)
	if ret != 0 {
		return -int32(3)
	}
	/* Initialize CELT decoder */
	ret = Opus_celt_decoder_init(tls, celt_dec, Fs, channels)
	if ret != OPUS_OK {
		return -int32(3)
	}
	_ = int32(0) == int32(0)
	Opus_opus_custom_decoder_ctl(tls, celt_dec, int32(CELT_SET_SIGNALLING_REQUEST), libc.VaList(bp+16, int32(0)))
	(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fprev_mode = 0
	(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fframe_size = Fs / int32(400)
	v1 = 0
	goto _7
_7:
	(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Farch = v1
	return OPUS_OK
}

func Opus_opus_decoder_create(tls *libc.TLS, Fs OpusT_opus_int32, channels int32, error1 uintptr) (r uintptr) {
	var ret int32
	var st, v1 uintptr
	_, _, _ = ret, st, v1
	if Fs != int32(48000) && Fs != int32(24000) && Fs != int32(16000) && Fs != int32(12000) && Fs != int32(8000) || channels != int32(1) && channels != int32(2) {
		if error1 != 0 {
			*(*int32)(unsafe.Pointer(error1)) = -int32(1)
		}
		return uintptr(uint32(0))
	}
	v1 = libc.Xmalloc(tls, uint64(uint32(Opus_opus_decoder_get_size(tls, channels))))
	goto _2
_2:
	st = v1
	if st == uintptr(uint32(0)) {
		if error1 != 0 {
			*(*int32)(unsafe.Pointer(error1)) = -int32(7)
		}
		return uintptr(uint32(0))
	}
	ret = Opus_opus_decoder_init(tls, st, Fs, channels)
	if error1 != 0 {
		*(*int32)(unsafe.Pointer(error1)) = ret
	}
	if ret != OPUS_OK {
		libc.Xfree(tls, st)
		st = uintptr(uint32(0))
	}
	return st
}

func smooth_fade(tls *libc.TLS, in1 uintptr, in2 uintptr, out uintptr, overlap int32, channels int32, window uintptr, Fs OpusT_opus_int32) {
	var c, i, inc int32
	var w OpusT_celt_coef
	_, _, _, _ = c, i, inc, w
	inc = int32(48000) / Fs
	c = 0
	for {
		if !(c < channels) {
			break
		}
		i = 0
		for {
			if !(i < overlap) {
				break
			}
			w = OpusT_celt_coef(*(*OpusT_celt_coef)(unsafe.Pointer(window + uintptr(i*inc)*4)) * *(*OpusT_celt_coef)(unsafe.Pointer(window + uintptr(i*inc)*4)))
			*(*OpusT_opus_res)(unsafe.Pointer(out + uintptr(i*channels+c)*4)) = OpusT_celt_coef(w**(*OpusT_opus_res)(unsafe.Pointer(in2 + uintptr(i*channels+c)*4))) + float32((float32(1)-w)**(*OpusT_opus_res)(unsafe.Pointer(in1 + uintptr(i*channels+c)*4)))
			goto _2
		_2:
			;
			i = i + 1
		}
		goto _1
	_1:
		;
		c = c + 1
	}
}

func opus_packet_get_mode(tls *libc.TLS, data uintptr) (r int32) {
	var mode int32
	_ = mode
	if int32(*(*uint8)(unsafe.Pointer(data)))&int32(0x80) != 0 {
		mode = int32(MODE_CELT_ONLY)
	} else {
		if int32(*(*uint8)(unsafe.Pointer(data)))&int32(0x60) == int32(0x60) {
			mode = int32(MODE_HYBRID)
		} else {
			mode = int32(MODE_SILK_ONLY)
		}
	}
	return mode
}

func opus_decode_frame(tls *libc.TLS, st1 uintptr, data uintptr, len1 OpusT_opus_int32, pcm uintptr, frame_size int32, decode_fec int32) (r int32) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var F10, F20, F2_5, F5, audiosize, bandwidth, c, celt_accum, celt_frame_size, celt_ret, celt_to_silk, decoded_samples, endband, first_frame, i, lost_flag, mode, pcm_silk_size, pcm_too_small, pcm_transition_celt_size, pcm_transition_silk_size, redundancy, redundancy_bytes, redundant_audio_size, ret, silk_ret, start_band, transition, v31, v32 int32
	var _saved_stack, celt_dec, pcm_ptr, pcm_silk, pcm_transition, pcm_transition_celt, pcm_transition_silk, redundant_audio, silk_dec, st, window, v1, v10, v11, v13, v15, v17, v19, v21, v3, v5, v6, v8 uintptr
	var frac, v175, v176 float32
	var gain, x1 OpusT_opus_val32
	var integer OpusT_opus_int32
	var v111 bool
	var _ /* celt_mode at bp+80 */ uintptr
	var _ /* dec at bp+8 */ OpusT_ec_dec
	var _ /* redundant_rng at bp+68 */ OpusT_opus_uint32
	var _ /* res at bp+0 */ struct {
		Fi [0]OpusT_opus_uint32
		Ff float32
	}
	var _ /* silence at bp+72 */ [2]uint8
	var _ /* silk_frame_size at bp+64 */ OpusT_opus_int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = F10, F20, F2_5, F5, _saved_stack, audiosize, bandwidth, c, celt_accum, celt_dec, celt_frame_size, celt_ret, celt_to_silk, decoded_samples, endband, first_frame, frac, gain, i, integer, lost_flag, mode, pcm_ptr, pcm_silk, pcm_silk_size, pcm_too_small, pcm_transition, pcm_transition_celt, pcm_transition_celt_size, pcm_transition_silk, pcm_transition_silk_size, redundancy, redundancy_bytes, redundant_audio, redundant_audio_size, ret, silk_dec, silk_ret, st, start_band, transition, window, x1, v1, v10, v11, v111, v13, v15, v17, v175, v176, v19, v21, v3, v31, v32, v5, v6, v8
	silk_ret = 0
	celt_ret = 0
	pcm_transition = uintptr(uint32(0))
	transition = 0
	redundancy = 0
	redundancy_bytes = 0
	celt_to_silk = 0
	*(*OpusT_opus_uint32)(unsafe.Pointer(bp + 68)) = uint32(0)
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
	silk_dec = st1 + uintptr((*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fsilk_dec_offset)
	celt_dec = st1 + uintptr((*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fcelt_dec_offset)
	F20 = (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).FFs / int32(50)
	F10 = F20 >> int32(1)
	F5 = F10 >> int32(1)
	F2_5 = F5 >> int32(1)
	if frame_size < F2_5 {
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			goto _28
		_28:
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v3 = st
		goto _30
	_30:
		(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
		return -int32(2)
	}
	/* Limit frame_size to avoid excessive stack allocations. */
	if frame_size < (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).FFs/int32(25)*int32(3) {
		v31 = frame_size
	} else {
		v31 = (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).FFs / int32(25) * int32(3)
	}
	frame_size = v31
	/* Payloads of 1 (2 including ToC) or 0 trigger the PLC/DTX */
	if len1 <= int32(1) {
		data = uintptr(uint32(0))
		/* In that case, don't conceal more than what the ToC says */
		if frame_size < (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fframe_size {
			v31 = frame_size
		} else {
			v31 = (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fframe_size
		}
		frame_size = v31
	}
	if data != uintptr(uint32(0)) {
		audiosize = (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fframe_size
		mode = (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fmode
		bandwidth = (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fbandwidth
		Opus_ec_dec_init(tls, bp+8, data, uint32(len1))
	} else {
		audiosize = frame_size
		/* Run PLC using last used mode (CELT if we ended with CELT redundancy) */
		if (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fprev_redundancy != 0 {
			v31 = int32(MODE_CELT_ONLY)
		} else {
			v31 = (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fprev_mode
		}
		mode = v31
		bandwidth = 0
		if mode == 0 {
			/* If we haven't got any packet yet, all we can do is return zeros */
			i = 0
			for {
				if !(i < audiosize*(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels) {
					break
				}
				*(*OpusT_opus_res)(unsafe.Pointer(pcm + uintptr(i)*4)) = float32(0)
				goto _34
			_34:
				;
				i = i + 1
			}
			st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
			if !(st != 0) {
				v1 = libc.Xmalloc(tls, uint64(16))
				goto _36
			_36:
				st = v1
				if st != 0 {
					libc.Xmemset(tls, st, 0, uint64(16))
				}
				libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
			}
			v3 = st
			goto _38
		_38:
			(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
			return audiosize
		}
		/* Avoids trying to run the PLC on sizes other than 2.5 (CELT), 5 (CELT),
		   10, or 20 (e.g. 12.5 or 30 ms). */
		if audiosize > F20 {
			for cond := true; cond; cond = audiosize > 0 {
				if audiosize < F20 {
					v31 = audiosize
				} else {
					v31 = F20
				}
				ret = opus_decode_frame(tls, st1, uintptr(uint32(0)), 0, pcm, v31, 0)
				if ret < 0 {
					st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
					if !(st != 0) {
						v1 = libc.Xmalloc(tls, uint64(16))
						goto _41
					_41:
						st = v1
						if st != 0 {
							libc.Xmemset(tls, st, 0, uint64(16))
						}
						libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
					}
					v3 = st
					goto _43
				_43:
					(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
					return ret
				}
				pcm = pcm + uintptr(ret*(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels)*4
				audiosize = audiosize - ret
			}
			st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
			if !(st != 0) {
				v1 = libc.Xmalloc(tls, uint64(16))
				goto _45
			_45:
				st = v1
				if st != 0 {
					libc.Xmemset(tls, st, 0, uint64(16))
				}
				libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
			}
			v3 = st
			goto _47
		_47:
			(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
			return frame_size
		} else {
			if audiosize < F20 {
				if audiosize > F10 {
					audiosize = F10
				} else {
					if mode != int32(MODE_SILK_ONLY) && audiosize > F5 && audiosize < F10 {
						audiosize = F5
					}
				}
			}
		}
	}
	/* In fixed-point, we can tell CELT to do the accumulation on top of the
	   SILK PCM buffer. This saves some stack space. */
	celt_accum = libc.BoolInt32(mode != int32(MODE_CELT_ONLY))
	pcm_transition_silk_size = ALLOC_NONE
	pcm_transition_celt_size = ALLOC_NONE
	if data != uintptr(uint32(0)) && (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fprev_mode > 0 && (mode == int32(MODE_CELT_ONLY) && (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fprev_mode != int32(MODE_CELT_ONLY) && !((*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fprev_redundancy != 0) || mode != int32(MODE_CELT_ONLY) && (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fprev_mode == int32(MODE_CELT_ONLY)) {
		transition = int32(1)
		/* Decide where to allocate the stack memory for pcm_transition */
		if mode == int32(MODE_CELT_ONLY) {
			pcm_transition_celt_size = F5 * (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels
		} else {
			pcm_transition_silk_size = F5 * (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels
		}
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _49
	_49:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _51
_51:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _53
	_53:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	goto _55
_55:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _57
	_57:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	goto _59
_59:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _61
	_61:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _63
_63:
	if !(int64(int32(uint64(uint32(pcm_transition_celt_size))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+57, int32(386))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _65
	_65:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _67
_67:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(uint64(uint32(pcm_transition_celt_size)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _69
	_69:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _71
_71:
	pcm_transition_celt = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(pcm_transition_celt_size))*(uint64(4)/uint64(1)))
	if transition != 0 && mode == int32(MODE_CELT_ONLY) {
		pcm_transition = pcm_transition_celt
		if F5 < audiosize {
			v31 = F5
		} else {
			v31 = audiosize
		}
		opus_decode_frame(tls, st1, uintptr(uint32(0)), 0, pcm_transition, v31, 0)
	}
	if audiosize > frame_size {
		/*fprintf(stderr, "PCM buffer too small: %d vs %d (mode = %d)\n", audiosize, frame_size, mode);*/
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			goto _74
		_74:
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v3 = st
		goto _76
	_76:
		(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
		return -int32(1)
	} else {
		frame_size = audiosize
	}
	/* SILK processing */
	if mode != int32(MODE_CELT_ONLY) {
		pcm_silk_size = ALLOC_NONE
		pcm_too_small = libc.BoolInt32(frame_size < F10)
		if pcm_too_small != 0 {
			pcm_silk_size = F10 * (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels
		}
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			goto _78
		_78:
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v3 = st
		goto _80
	_80:
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v5 = libc.Xmalloc(tls, uint64(16))
			goto _82
		_82:
			st = v5
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v6 = st
		goto _84
	_84:
		*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v8 = libc.Xmalloc(tls, uint64(16))
			goto _86
		_86:
			st = v8
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v10 = st
		goto _88
	_88:
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v11 = libc.Xmalloc(tls, uint64(16))
			goto _90
		_90:
			st = v11
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v13 = st
		goto _92
	_92:
		if !(int64(int32(uint64(uint32(pcm_silk_size))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
			Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+57, int32(412))
		}
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v15 = libc.Xmalloc(tls, uint64(16))
			goto _94
		_94:
			st = v15
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v17 = st
		goto _96
	_96:
		*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(uint64(uint32(pcm_silk_size)) * (uint64(4) / uint64(1)))
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v19 = libc.Xmalloc(tls, uint64(16))
			goto _98
		_98:
			st = v19
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v21 = st
		goto _100
	_100:
		pcm_silk = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(pcm_silk_size))*(uint64(4)/uint64(1)))
		if pcm_too_small != 0 {
			pcm_ptr = pcm_silk
		} else {
			pcm_ptr = pcm
		}
		if (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fprev_mode == int32(MODE_CELT_ONLY) {
			Opus_silk_ResetDecoder(tls, silk_dec)
		}
		/* The SILK PLC cannot produce frames of less than 10 ms */
		if int32(10) > int32(1000)*audiosize/(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).FFs {
			v31 = int32(10)
		} else {
			v31 = int32(1000) * audiosize / (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).FFs
		}
		(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).FDecControl.FpayloadSize_ms = v31
		if data != uintptr(uint32(0)) {
			(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).FDecControl.FnChannelsInternal = (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fstream_channels
			if mode == int32(MODE_SILK_ONLY) {
				if bandwidth == int32(OPUS_BANDWIDTH_NARROWBAND) {
					(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).FDecControl.FinternalSampleRate = int32(8000)
				} else {
					if bandwidth == int32(OPUS_BANDWIDTH_MEDIUMBAND) {
						(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).FDecControl.FinternalSampleRate = int32(12000)
					} else {
						if bandwidth == int32(OPUS_BANDWIDTH_WIDEBAND) {
							(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).FDecControl.FinternalSampleRate = int32(16000)
						} else {
							(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).FDecControl.FinternalSampleRate = int32(16000)
							if !(int32(0) != 0) {
								Opus_celt_fatal(tls, __ccgo_ts+1017, __ccgo_ts+57, int32(436))
							}
						}
					}
				}
			} else {
				/* Hybrid mode */
				(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).FDecControl.FinternalSampleRate = int32(16000)
			}
		}
		(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).FDecControl.Fenable_deep_plc = libc.BoolInt32((*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fcomplexity >= int32(5))
		if data == uintptr(uint32(0)) {
			v31 = int32(1)
		} else {
			v31 = int32(2) * libc.BoolInt32(!!(decode_fec != 0))
		}
		lost_flag = v31
		decoded_samples = 0
		for cond := true; cond; cond = decoded_samples < frame_size {
			/* Call SILK decoder */
			first_frame = libc.BoolInt32(decoded_samples == 0)
			silk_ret = Opus_silk_Decode(tls, silk_dec, st1+16, lost_flag, first_frame, bp+8, pcm_ptr, bp+64, (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Farch)
			if silk_ret != 0 {
				if lost_flag != 0 {
					/* PLC failure should not be fatal */
					*(*OpusT_opus_int32)(unsafe.Pointer(bp + 64)) = frame_size
					i = 0
					for {
						if !(i < frame_size*(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels) {
							break
						}
						*(*OpusT_opus_res)(unsafe.Pointer(pcm_ptr + uintptr(i)*4)) = float32(0)
						goto _103
					_103:
						;
						i = i + 1
					}
				} else {
					st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
					if !(st != 0) {
						v1 = libc.Xmalloc(tls, uint64(16))
						goto _105
					_105:
						st = v1
						if st != 0 {
							libc.Xmemset(tls, st, 0, uint64(16))
						}
						libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
					}
					v3 = st
					goto _107
				_107:
					(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
					return -int32(3)
				}
			}
			pcm_ptr = pcm_ptr + uintptr(*(*OpusT_opus_int32)(unsafe.Pointer(bp + 64))*(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels)*4
			decoded_samples = decoded_samples + *(*OpusT_opus_int32)(unsafe.Pointer(bp + 64))
		}
		if pcm_too_small != 0 {
			libc.Xmemcpy(tls, pcm, pcm_silk, uint64(uint32(frame_size*(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels))*uint64(4)+uint64(0*((int64(pcm)-int64(pcm_silk))/4)))
		}
	}
	start_band = 0
	if v111 = !(decode_fec != 0) && mode != int32(MODE_CELT_ONLY) && data != uintptr(uint32(0)); v111 {
		v1 = bp + 8
		v31 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
		goto _110
	_110:
	}
	if v111 && v31+int32(17)+int32(20)*libc.BoolInt32(mode == int32(MODE_HYBRID)) <= int32(8)*len1 {
		/* Check if we have a redundant 0-8 kHz band */
		if mode == int32(MODE_HYBRID) {
			redundancy = Opus_ec_dec_bit_logp(tls, bp+8, uint32(12))
		} else {
			redundancy = int32(1)
		}
		if redundancy != 0 {
			celt_to_silk = Opus_ec_dec_bit_logp(tls, bp+8, uint32(1))
			/* redundancy_bytes will be at least two, in the non-hybrid
			   case due to the ec_tell() check above */
			if mode == int32(MODE_HYBRID) {
				v31 = int32(Opus_ec_dec_uint(tls, bp+8, uint32(256))) + int32(2)
			} else {
				v1 = bp + 8
				v32 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
				goto _115
			_115:
				v31 = len1 - (v32+int32(7))>>int32(3)
			}
			redundancy_bytes = v31
			len1 = len1 - redundancy_bytes
			/* This is a sanity check. It should never happen for a valid
			   packet, so the exact behaviour is not normative. */
			v1 = bp + 8
			v31 = (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, (*OpusT_ec_ctx)(unsafe.Pointer(v1)).Frng))
			goto _118
		_118:
			if len1*int32(8) < v31 {
				len1 = 0
				redundancy_bytes = 0
				redundancy = 0
			}
			/* Shrink decoder because of raw bits */
			(*(*OpusT_ec_dec)(unsafe.Pointer(bp + 8))).Fstorage -= uint32(redundancy_bytes)
		}
	}
	if mode != int32(MODE_CELT_ONLY) {
		start_band = int32(17)
	}
	if redundancy != 0 {
		transition = 0
		pcm_transition_silk_size = ALLOC_NONE
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _120
	_120:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _122
_122:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _124
	_124:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	goto _126
_126:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _128
	_128:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	goto _130
_130:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _132
	_132:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _134
_134:
	if !(int64(int32(uint64(uint32(pcm_transition_silk_size))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+57, int32(538))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _136
	_136:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _138
_138:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(uint64(uint32(pcm_transition_silk_size)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _140
	_140:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _142
_142:
	pcm_transition_silk = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(pcm_transition_silk_size))*(uint64(4)/uint64(1)))
	if transition != 0 && mode != int32(MODE_CELT_ONLY) {
		pcm_transition = pcm_transition_silk
		if F5 < audiosize {
			v31 = F5
		} else {
			v31 = audiosize
		}
		opus_decode_frame(tls, st1, uintptr(uint32(0)), 0, pcm_transition, v31, 0)
	}
	if bandwidth != 0 {
		endband = int32(21)
		switch bandwidth {
		case int32(OPUS_BANDWIDTH_NARROWBAND):
			endband = int32(13)
		case int32(OPUS_BANDWIDTH_MEDIUMBAND):
			fallthrough
		case int32(OPUS_BANDWIDTH_WIDEBAND):
			endband = int32(17)
		case int32(OPUS_BANDWIDTH_SUPERWIDEBAND):
			endband = int32(19)
		case int32(OPUS_BANDWIDTH_FULLBAND):
			endband = int32(21)
		default:
			if !(int32(0) != 0) {
				Opus_celt_fatal(tls, __ccgo_ts+1017, __ccgo_ts+57, int32(567))
			}
			break
		}
		_ = endband == int32(0)
		if !(Opus_opus_custom_decoder_ctl(tls, celt_dec, int32(CELT_SET_END_BAND_REQUEST), libc.VaList(bp+96, endband)) == int32(OPUS_OK)) {
			Opus_celt_fatal(tls, __ccgo_ts+1037, __ccgo_ts+57, int32(570))
		}
	}
	_ = (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fstream_channels == int32(0)
	if !(Opus_opus_custom_decoder_ctl(tls, celt_dec, int32(CELT_SET_CHANNELS_REQUEST), libc.VaList(bp+96, (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fstream_channels)) == int32(OPUS_OK)) {
		Opus_celt_fatal(tls, __ccgo_ts+1172, __ccgo_ts+57, int32(572))
	}
	/* Only allocation memory for redundancy if/when needed */
	if redundancy != 0 {
		v31 = F5 * (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels
	} else {
		v31 = ALLOC_NONE
	}
	redundant_audio_size = v31
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _146
	_146:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _148
_148:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _150
	_150:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	goto _152
_152:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _154
	_154:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	goto _156
_156:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _158
	_158:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _160
_160:
	if !(int64(int32(uint64(uint32(redundant_audio_size))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+57, int32(576))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _162
	_162:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _164
_164:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(uint64(uint32(redundant_audio_size)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _166
	_166:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _168
_168:
	redundant_audio = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(redundant_audio_size))*(uint64(4)/uint64(1)))
	/* 5 ms redundant frame for CELT->SILK*/
	if redundancy != 0 && celt_to_silk != 0 {
		/* If the previous frame did not use CELT (the first redundancy frame in
		   a transition from SILK may have been lost) then the CELT decoder is
		   stale at this point and the redundancy audio is not useful, however
		   the final range is still needed (for testing), so the redundancy is
		   always decoded but the decoded audio may not be used */
		_ = int32(0) == int32(0)
		if !(Opus_opus_custom_decoder_ctl(tls, celt_dec, int32(CELT_SET_START_BAND_REQUEST), libc.VaList(bp+96, int32(0))) == int32(OPUS_OK)) {
			Opus_celt_fatal(tls, __ccgo_ts+1331, __ccgo_ts+57, int32(586))
		}
		Opus_celt_decode_with_ec(tls, celt_dec, data+uintptr(len1), redundancy_bytes, redundant_audio, F5, uintptr(uint32(0)), 0)
		if !(Opus_opus_custom_decoder_ctl(tls, celt_dec, int32(OPUS_GET_FINAL_RANGE_REQUEST), libc.VaList(bp+96, bp+68+uintptr((OpusT___predefined_ptrdiff_t(bp+68)-int64(bp+68))/4)*4)) == int32(OPUS_OK)) {
			Opus_celt_fatal(tls, __ccgo_ts+1454, __ccgo_ts+57, int32(589))
		}
	}
	/* MUST be after PLC */
	_ = start_band == int32(0)
	if !(Opus_opus_custom_decoder_ctl(tls, celt_dec, int32(CELT_SET_START_BAND_REQUEST), libc.VaList(bp+96, start_band)) == int32(OPUS_OK)) {
		Opus_celt_fatal(tls, __ccgo_ts+1599, __ccgo_ts+57, int32(593))
	}
	if mode != int32(MODE_SILK_ONLY) {
		if F20 < frame_size {
			v31 = F20
		} else {
			v31 = frame_size
		}
		celt_frame_size = v31
		/* Make sure to discard any previous CELT state */
		if mode != (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fprev_mode && (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fprev_mode > 0 && !((*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fprev_redundancy != 0) {
			if !(Opus_opus_custom_decoder_ctl(tls, celt_dec, int32(OPUS_RESET_STATE), 0) == int32(OPUS_OK)) {
				Opus_celt_fatal(tls, __ccgo_ts+1740, __ccgo_ts+57, int32(604))
			}
		}
		/* Decode CELT */
		if decode_fec != 0 {
			v1 = uintptr(uint32(0))
		} else {
			v1 = data
		}
		celt_ret = Opus_celt_decode_with_ec_dred(tls, celt_dec, v1, len1, pcm, celt_frame_size, bp+8, celt_accum)
		Opus_opus_custom_decoder_ctl(tls, celt_dec, int32(OPUS_GET_FINAL_RANGE_REQUEST), libc.VaList(bp+96, st1+96+uintptr((OpusT___predefined_ptrdiff_t(st1+96)-int64(st1+96))/4)*4))
	} else {
		*(*[2]uint8)(unsafe.Pointer(bp + 72)) = [2]uint8{
			0: uint8(0xFF),
			1: uint8(0xFF),
		}
		if !(celt_accum != 0) {
			i = 0
			for {
				if !(i < frame_size*(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels) {
					break
				}
				*(*OpusT_opus_res)(unsafe.Pointer(pcm + uintptr(i)*4)) = float32(0)
				goto _171
			_171:
				;
				i = i + 1
			}
		}
		/* For hybrid -> SILK transitions, we let the CELT MDCT
		   do a fade-out by decoding a silence frame */
		if (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fprev_mode == int32(MODE_HYBRID) && !(redundancy != 0 && celt_to_silk != 0 && (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fprev_redundancy != 0) {
			_ = int32(0) == int32(0)
			if !(Opus_opus_custom_decoder_ctl(tls, celt_dec, int32(CELT_SET_START_BAND_REQUEST), libc.VaList(bp+96, int32(0))) == int32(OPUS_OK)) {
				Opus_celt_fatal(tls, __ccgo_ts+1331, __ccgo_ts+57, int32(624))
			}
			Opus_celt_decode_with_ec(tls, celt_dec, bp+72, int32(2), pcm, F2_5, uintptr(uint32(0)), celt_accum)
		}
		(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).FrangeFinal = (*(*OpusT_ec_dec)(unsafe.Pointer(bp + 8))).Frng
	}
	if !(Opus_opus_custom_decoder_ctl(tls, celt_dec, int32(CELT_GET_MODE_REQUEST), libc.VaList(bp+96, bp+80+uintptr((OpusT___predefined_ptrdiff_t(bp+80)-int64(bp+80))/8)*8)) == int32(OPUS_OK)) {
		Opus_celt_fatal(tls, __ccgo_ts+1811, __ccgo_ts+57, int32(632))
	}
	window = (*OpusT_OpusCustomMode)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(bp + 80)))).Fwindow
	/* 5 ms redundant frame for SILK->CELT */
	if redundancy != 0 && !(celt_to_silk != 0) {
		if !(Opus_opus_custom_decoder_ctl(tls, celt_dec, int32(OPUS_RESET_STATE), 0) == int32(OPUS_OK)) {
			Opus_celt_fatal(tls, __ccgo_ts+1740, __ccgo_ts+57, int32(639))
		}
		_ = int32(0) == int32(0)
		if !(Opus_opus_custom_decoder_ctl(tls, celt_dec, int32(CELT_SET_START_BAND_REQUEST), libc.VaList(bp+96, int32(0))) == int32(OPUS_OK)) {
			Opus_celt_fatal(tls, __ccgo_ts+1331, __ccgo_ts+57, int32(640))
		}
		Opus_celt_decode_with_ec(tls, celt_dec, data+uintptr(len1), redundancy_bytes, redundant_audio, F5, uintptr(uint32(0)), 0)
		if !(Opus_opus_custom_decoder_ctl(tls, celt_dec, int32(OPUS_GET_FINAL_RANGE_REQUEST), libc.VaList(bp+96, bp+68+uintptr((OpusT___predefined_ptrdiff_t(bp+68)-int64(bp+68))/4)*4)) == int32(OPUS_OK)) {
			Opus_celt_fatal(tls, __ccgo_ts+1454, __ccgo_ts+57, int32(643))
		}
		smooth_fade(tls, pcm+uintptr((*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels*(frame_size-F2_5))*4, redundant_audio+uintptr((*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels*F2_5)*4, pcm+uintptr((*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels*(frame_size-F2_5))*4, F2_5, (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels, window, (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).FFs)
	}
	/* 5ms redundant frame for CELT->SILK; ignore if the previous frame did not
	   use CELT (the first redundancy frame in a transition from SILK may have
	   been lost) */
	if redundancy != 0 && celt_to_silk != 0 && ((*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fprev_mode != int32(MODE_SILK_ONLY) || (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fprev_redundancy != 0) {
		c = 0
		for {
			if !(c < (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels) {
				break
			}
			i = 0
			for {
				if !(i < F2_5) {
					break
				}
				*(*OpusT_opus_res)(unsafe.Pointer(pcm + uintptr((*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels*i+c)*4)) = *(*OpusT_opus_res)(unsafe.Pointer(redundant_audio + uintptr((*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels*i+c)*4))
				goto _173
			_173:
				;
				i = i + 1
			}
			goto _172
		_172:
			;
			c = c + 1
		}
		smooth_fade(tls, redundant_audio+uintptr((*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels*F2_5)*4, pcm+uintptr((*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels*F2_5)*4, pcm+uintptr((*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels*F2_5)*4, F2_5, (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels, window, (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).FFs)
	}
	if transition != 0 {
		if audiosize >= F5 {
			i = 0
			for {
				if !(i < (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels*F2_5) {
					break
				}
				*(*OpusT_opus_res)(unsafe.Pointer(pcm + uintptr(i)*4)) = *(*OpusT_opus_res)(unsafe.Pointer(pcm_transition + uintptr(i)*4))
				goto _174
			_174:
				;
				i = i + 1
			}
			smooth_fade(tls, pcm_transition+uintptr((*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels*F2_5)*4, pcm+uintptr((*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels*F2_5)*4, pcm+uintptr((*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels*F2_5)*4, F2_5, (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels, window, (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).FFs)
		} else {
			/* Not enough time to do a clean transition, but we do it anyway
			   This will not preserve amplitude perfectly and may introduce
			   a bit of temporal aliasing, but it shouldn't be too bad and
			   that's pretty much the best we can do. In any case, generating this
			   transition it pretty silly in the first place */
			smooth_fade(tls, pcm_transition, pcm, pcm, F2_5, (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels, window, (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).FFs)
		}
	}
	if (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fdecode_gain != 0 {
		v175 = float32(float32(0.000648814081) * float32((*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fdecode_gain))
		integer = int32(libc.Xfloor(tls, float64(v175)))
		if integer < -int32(50) {
			v176 = float32(0)
			goto _177
		}
		frac = v175 - float32(integer)
		*(*float32)(unsafe.Pointer(bp)) = float32(0.9999999403953552) + float32(frac*(float32(0.6931530833244324)+float32(frac*(float32(0.24015361070632935)+float32(frac*(float32(0.05582631751894951)+float32(frac*(float32(0.00898933969438076)+float32(frac*float32(0.0018775766948238015))))))))))
		*(*OpusT_opus_uint32)(unsafe.Pointer(bp)) = uint32(int32(*(*OpusT_opus_uint32)(unsafe.Pointer(bp)))+int32(uint32(integer)<<int32(23))) & uint32(0x7fffffff)
		v176 = *(*float32)(unsafe.Pointer(bp))
		goto _177
	_177:
		gain = v176
		i = 0
		for {
			if !(i < frame_size*(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels) {
				break
			}
			x1 = OpusT_opus_res(*(*OpusT_opus_res)(unsafe.Pointer(pcm + uintptr(i)*4)) * gain)
			*(*OpusT_opus_res)(unsafe.Pointer(pcm + uintptr(i)*4)) = x1
			goto _178
		_178:
			;
			i = i + 1
		}
	}
	if len1 <= int32(1) {
		(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).FrangeFinal = uint32(0)
	} else {
		*(*OpusT_opus_uint32)(unsafe.Pointer(st1 + 96)) ^= *(*OpusT_opus_uint32)(unsafe.Pointer(bp + 68))
	}
	(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fprev_mode = mode
	(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fprev_redundancy = libc.BoolInt32(redundancy != 0 && !(celt_to_silk != 0))
	if celt_ret >= 0 {
		v31 = 0
		goto _180
	_180:
		if v31 != 0 {
		}
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _182
	_182:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _184
_184:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
	if celt_ret < 0 {
		v31 = celt_ret
	} else {
		v31 = audiosize
	}
	return v31
}

func Opus_opus_decode_native(tls *libc.TLS, st uintptr, data uintptr, len1 OpusT_opus_int32, pcm uintptr, frame_size int32, decode_fec int32, self_delimited int32, packet_offset uintptr, soft_clip int32, dred uintptr, dred_offset OpusT_opus_int32) (r int32) {
	bp := tls.Alloc(208)
	defer tls.Free(208)
	var count, duration_copy, i, nb_samples, packet_bandwidth, packet_frame_size, packet_mode, packet_stream_channels, pcm_count, ret, ret1, ret2, v1 int32
	var v8 OpusT_opus_val16
	var _ /* iter at bp+120 */ OpusT_OpusExtensionIterator
	var _ /* offset at bp+0 */ int32
	var _ /* padding at bp+104 */ uintptr
	var _ /* padding_len at bp+112 */ OpusT_opus_int32
	var _ /* size at bp+6 */ [48]OpusT_opus_int16
	var _ /* toc at bp+4 */ uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = count, duration_copy, i, nb_samples, packet_bandwidth, packet_frame_size, packet_mode, packet_stream_channels, pcm_count, ret, ret1, ret2, v1, v8
	validate_opus_decoder(tls, st)
	if decode_fec < 0 || decode_fec > int32(1) {
		return -int32(1)
	}
	/* For FEC/PLC, frame_size has to be to have a multiple of 2.5 ms */
	if (decode_fec != 0 || len1 == 0 || data == uintptr(uint32(0))) && frame_size%((*OpusT_OpusDecoder)(unsafe.Pointer(st)).FFs/int32(400)) != 0 {
		return -int32(1)
	}
	_ = dred
	_ = dred_offset
	if len1 == 0 || data == uintptr(uint32(0)) {
		pcm_count = 0
		for cond := true; cond; cond = pcm_count < frame_size {
			ret = opus_decode_frame(tls, st, uintptr(uint32(0)), 0, pcm+uintptr(pcm_count*(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fchannels)*4, frame_size-pcm_count, 0)
			if ret < 0 {
				return ret
			}
			pcm_count = pcm_count + ret
		}
		if !(pcm_count == frame_size) {
			Opus_celt_fatal(tls, __ccgo_ts+1955, __ccgo_ts+57, int32(773))
		}
		v1 = 0
		goto _2
	_2:
		if v1 != 0 {
		}
		(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Flast_packet_duration = pcm_count
		return pcm_count
	} else {
		if len1 < 0 {
			return -int32(1)
		}
	}
	packet_mode = opus_packet_get_mode(tls, data)
	packet_bandwidth = Opus_opus_packet_get_bandwidth(tls, data)
	packet_frame_size = Opus_opus_packet_get_samples_per_frame(tls, data, (*OpusT_OpusDecoder)(unsafe.Pointer(st)).FFs)
	packet_stream_channels = Opus_opus_packet_get_nb_channels(tls, data)
	count = Opus_opus_packet_parse_impl(tls, data, len1, self_delimited, bp+4, uintptr(uint32(0)), bp+6, bp, packet_offset, bp+104, bp+112)
	if (*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fignore_extensions != 0 {
		*(*uintptr)(unsafe.Pointer(bp + 104)) = uintptr(uint32(0))
		*(*OpusT_opus_int32)(unsafe.Pointer(bp + 112)) = 0
	}
	if count < 0 {
		return count
	}
	Opus_opus_extension_iterator_init(tls, bp+120, *(*uintptr)(unsafe.Pointer(bp + 104)), *(*OpusT_opus_int32)(unsafe.Pointer(bp + 112)), count)
	data = data + uintptr(*(*int32)(unsafe.Pointer(bp)))
	if decode_fec != 0 {
		/* If no FEC can be present, run the PLC (recursive call) */
		if frame_size < packet_frame_size || packet_mode == int32(MODE_CELT_ONLY) || (*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fmode == int32(MODE_CELT_ONLY) {
			return Opus_opus_decode_native(tls, st, uintptr(uint32(0)), 0, pcm, frame_size, 0, 0, uintptr(uint32(0)), soft_clip, uintptr(uint32(0)), 0)
		}
		/* Otherwise, run the PLC on everything except the size for which we might have FEC */
		duration_copy = (*OpusT_OpusDecoder)(unsafe.Pointer(st)).Flast_packet_duration
		if frame_size-packet_frame_size != 0 {
			ret1 = Opus_opus_decode_native(tls, st, uintptr(uint32(0)), 0, pcm, frame_size-packet_frame_size, 0, 0, uintptr(uint32(0)), soft_clip, uintptr(uint32(0)), 0)
			if ret1 < 0 {
				(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Flast_packet_duration = duration_copy
				return ret1
			}
			if !(ret1 == frame_size-packet_frame_size) {
				Opus_celt_fatal(tls, __ccgo_ts+1997, __ccgo_ts+57, int32(815))
			}
		}
		/* Complete with FEC */
		(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fmode = packet_mode
		(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fbandwidth = packet_bandwidth
		(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fframe_size = packet_frame_size
		(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fstream_channels = packet_stream_channels
		ret1 = opus_decode_frame(tls, st, data, int32((*(*[48]OpusT_opus_int16)(unsafe.Pointer(bp + 6)))[0]), pcm+uintptr((*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fchannels*(frame_size-packet_frame_size))*4, packet_frame_size, int32(1))
		if ret1 < 0 {
			return ret1
		} else {
			v1 = 0
			goto _4
		_4:
			if v1 != 0 {
			}
			(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Flast_packet_duration = frame_size
			return frame_size
		}
	}
	if count*packet_frame_size > frame_size {
		return -int32(2)
	}
	/* Update the state as the last step to avoid updating it on an invalid packet */
	(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fmode = packet_mode
	(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fbandwidth = packet_bandwidth
	(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fframe_size = packet_frame_size
	(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fstream_channels = packet_stream_channels
	nb_samples = 0
	i = 0
	for {
		if !(i < count) {
			break
		}
		ret2 = opus_decode_frame(tls, st, data, int32((*(*[48]OpusT_opus_int16)(unsafe.Pointer(bp + 6)))[i]), pcm+uintptr(nb_samples*(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fchannels)*4, frame_size-nb_samples, 0)
		if ret2 < 0 {
			return ret2
		}
		if !(ret2 == packet_frame_size) {
			Opus_celt_fatal(tls, __ccgo_ts+2049, __ccgo_ts+57, int32(865))
		}
		data = data + uintptr((*(*[48]OpusT_opus_int16)(unsafe.Pointer(bp + 6)))[i])
		nb_samples = nb_samples + ret2
		goto _5
	_5:
		;
		i = i + 1
	}
	(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Flast_packet_duration = nb_samples
	v1 = 0
	goto _7
_7:
	if v1 != 0 {
	}
	if soft_clip != 0 {
		Opus_opus_pcm_soft_clip_impl(tls, pcm, nb_samples, (*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fchannels, st+88, (*OpusT_OpusDecoder)(unsafe.Pointer(st)).Farch)
	} else {
		v8 = float32(0)
		*(*OpusT_opus_val16)(unsafe.Pointer(st + 88 + 1*4)) = v8
		*(*OpusT_opus_val16)(unsafe.Pointer(st + 88)) = v8
	}
	return nb_samples
}

func Opus_opus_decode(tls *libc.TLS, st1 uintptr, data uintptr, len1 OpusT_opus_int32, pcm uintptr, frame_size int32, decode_fec int32) (r int32) {
	var _saved_stack, out, st, v1, v10, v11, v13, v15, v17, v19, v21, v3, v5, v6, v8 uintptr
	var nb_samples, ret, v31 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = _saved_stack, nb_samples, out, ret, st, v1, v10, v11, v13, v15, v17, v19, v21, v3, v31, v5, v6, v8
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
	if frame_size <= 0 {
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			goto _28
		_28:
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v3 = st
		goto _30
	_30:
		(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
		return -int32(1)
	}
	if data != uintptr(uint32(0)) && len1 > 0 && !(decode_fec != 0) {
		nb_samples = Opus_opus_decoder_get_nb_samples(tls, st1, data, len1)
		if nb_samples > 0 {
			if frame_size < nb_samples {
				v31 = frame_size
			} else {
				v31 = nb_samples
			}
			frame_size = v31
		} else {
			return -int32(4)
		}
	}
	if !((*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels == int32(1) || (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels == int32(2)) {
		Opus_celt_fatal(tls, __ccgo_ts, __ccgo_ts+57, int32(917))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _33
	_33:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _35
_35:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _37
	_37:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	goto _39
_39:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _41
	_41:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	goto _43
_43:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _45
	_45:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _47
_47:
	if !(int64(int32(uint64(uint32(frame_size*(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+57, int32(918))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _49
	_49:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _51
_51:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(uint64(uint32(frame_size*(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _53
	_53:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _55
_55:
	out = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(frame_size*(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels))*(uint64(4)/uint64(1)))
	ret = Opus_opus_decode_native(tls, st1, data, len1, out, frame_size, decode_fec, 0, uintptr(uint32(0)), int32(OPTIONAL_CLIP), uintptr(uint32(0)), 0)
	if ret > 0 {
		_ = (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Farch
		Opus_celt_float2int16_c(tls, out, pcm, ret*(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels)
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _57
	_57:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _59
_59:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
	return ret
}

func Opus_opus_decode24(tls *libc.TLS, st1 uintptr, data uintptr, len1 OpusT_opus_int32, pcm uintptr, frame_size int32, decode_fec int32) (r int32) {
	var _saved_stack, out, st, v1, v10, v11, v13, v15, v17, v19, v21, v3, v5, v6, v8 uintptr
	var i, nb_samples, ret, v31 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = _saved_stack, i, nb_samples, out, ret, st, v1, v10, v11, v13, v15, v17, v19, v21, v3, v31, v5, v6, v8
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
	if frame_size <= 0 {
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			goto _28
		_28:
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v3 = st
		goto _30
	_30:
		(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
		return -int32(1)
	}
	if data != uintptr(uint32(0)) && len1 > 0 && !(decode_fec != 0) {
		nb_samples = Opus_opus_decoder_get_nb_samples(tls, st1, data, len1)
		if nb_samples > 0 {
			if frame_size < nb_samples {
				v31 = frame_size
			} else {
				v31 = nb_samples
			}
			frame_size = v31
		} else {
			return -int32(4)
		}
	}
	if !((*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels == int32(1) || (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels == int32(2)) {
		Opus_celt_fatal(tls, __ccgo_ts, __ccgo_ts+57, int32(966))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _33
	_33:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _35
_35:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _37
	_37:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	goto _39
_39:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _41
	_41:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	goto _43
_43:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _45
	_45:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _47
_47:
	if !(int64(int32(uint64(uint32(frame_size*(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+57, int32(967))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _49
	_49:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _51
_51:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(uint64(uint32(frame_size*(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _53
	_53:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _55
_55:
	out = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(frame_size*(*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels))*(uint64(4)/uint64(1)))
	ret = Opus_opus_decode_native(tls, st1, data, len1, out, frame_size, decode_fec, 0, uintptr(uint32(0)), 0, uintptr(uint32(0)), 0)
	if ret > 0 {
		nb_samples = ret * (*OpusT_OpusDecoder)(unsafe.Pointer(st1)).Fchannels
		i = 0
		for {
			if !(i < nb_samples) {
				break
			}
			*(*OpusT_opus_int32)(unsafe.Pointer(pcm + uintptr(i)*4)) = int32(libc.Xlrintf(tls, float32(float32(float32(32768)*float32(256))**(*OpusT_opus_res)(unsafe.Pointer(out + uintptr(i)*4)))))
			goto _56
		_56:
			;
			i = i + 1
		}
	}
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
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
	return ret
}

func Opus_opus_decode_float(tls *libc.TLS, st uintptr, data uintptr, len1 OpusT_opus_int32, pcm uintptr, frame_size int32, decode_fec int32) (r int32) {
	if frame_size <= 0 {
		return -int32(1)
	}
	return Opus_opus_decode_native(tls, st, data, len1, pcm, frame_size, decode_fec, 0, uintptr(uint32(0)), 0, uintptr(uint32(0)), 0)
}

func Opus_opus_decoder_ctl(tls *libc.TLS, st uintptr, request int32, va uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var ap OpusT_va_list
	var celt_dec, silk_dec, value, value10, value12, value2, value3, value4, value5, value6, value8 uintptr
	var ret int32
	var value1, value11, value7, value9 OpusT_opus_int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = ap, celt_dec, ret, silk_dec, value, value1, value10, value11, value12, value2, value3, value4, value5, value6, value7, value8, value9
	ret = OPUS_OK
	silk_dec = st + uintptr((*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fsilk_dec_offset)
	celt_dec = st + uintptr((*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fcelt_dec_offset)
	ap = va
	switch request {
	case int32(OPUS_GET_BANDWIDTH_REQUEST):
		value = libc.VaUintptr(&ap)
		if !(value != 0) {
			goto bad_arg
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(value)) = (*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fbandwidth
	case int32(OPUS_SET_COMPLEXITY_REQUEST):
		value1 = libc.VaInt32(&ap)
		if value1 < 0 || value1 > int32(10) {
			goto bad_arg
		}
		(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fcomplexity = value1
		_ = value1 == int32(0)
		Opus_opus_custom_decoder_ctl(tls, celt_dec, int32(OPUS_SET_COMPLEXITY_REQUEST), libc.VaList(bp+8, value1))
	case int32(OPUS_GET_COMPLEXITY_REQUEST):
		value2 = libc.VaUintptr(&ap)
		if !(value2 != 0) {
			goto bad_arg
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(value2)) = (*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fcomplexity
	case int32(OPUS_GET_FINAL_RANGE_REQUEST):
		value3 = libc.VaUintptr(&ap)
		if !(value3 != 0) {
			goto bad_arg
		}
		*(*OpusT_opus_uint32)(unsafe.Pointer(value3)) = (*OpusT_OpusDecoder)(unsafe.Pointer(st)).FrangeFinal
	case int32(OPUS_RESET_STATE):
		libc.Xmemset(tls, st+60, 0, (uint64(100)-uint64(int64(st+60)-int64(st)))*uint64(1))
		Opus_opus_custom_decoder_ctl(tls, celt_dec, int32(OPUS_RESET_STATE), 0)
		Opus_silk_ResetDecoder(tls, silk_dec)
		(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fstream_channels = (*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fchannels
		(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fframe_size = (*OpusT_OpusDecoder)(unsafe.Pointer(st)).FFs / int32(400)
	case int32(OPUS_GET_SAMPLE_RATE_REQUEST):
		value4 = libc.VaUintptr(&ap)
		if !(value4 != 0) {
			goto bad_arg
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(value4)) = (*OpusT_OpusDecoder)(unsafe.Pointer(st)).FFs
	case int32(OPUS_GET_PITCH_REQUEST):
		value5 = libc.VaUintptr(&ap)
		if !(value5 != 0) {
			goto bad_arg
		}
		if (*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fprev_mode == int32(MODE_CELT_ONLY) {
			ret = Opus_opus_custom_decoder_ctl(tls, celt_dec, int32(OPUS_GET_PITCH_REQUEST), libc.VaList(bp+8, value5+uintptr((int64(value5)-int64(value5))/4)*4))
		} else {
			*(*OpusT_opus_int32)(unsafe.Pointer(value5)) = (*OpusT_OpusDecoder)(unsafe.Pointer(st)).FDecControl.FprevPitchLag
		}
	case int32(OPUS_GET_GAIN_REQUEST):
		value6 = libc.VaUintptr(&ap)
		if !(value6 != 0) {
			goto bad_arg
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(value6)) = (*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fdecode_gain
	case int32(OPUS_SET_GAIN_REQUEST):
		value7 = libc.VaInt32(&ap)
		if value7 < -int32(32768) || value7 > int32(32767) {
			goto bad_arg
		}
		(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fdecode_gain = value7
	case int32(OPUS_GET_LAST_PACKET_DURATION_REQUEST):
		value8 = libc.VaUintptr(&ap)
		if !(value8 != 0) {
			goto bad_arg
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(value8)) = (*OpusT_OpusDecoder)(unsafe.Pointer(st)).Flast_packet_duration
	case int32(OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST):
		value9 = libc.VaInt32(&ap)
		if value9 < 0 || value9 > int32(1) {
			goto bad_arg
		}
		_ = value9 == int32(0)
		ret = Opus_opus_custom_decoder_ctl(tls, celt_dec, int32(OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST), libc.VaList(bp+8, value9))
	case int32(OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST):
		value10 = libc.VaUintptr(&ap)
		if !(value10 != 0) {
			goto bad_arg
		}
		ret = Opus_opus_custom_decoder_ctl(tls, celt_dec, int32(OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST), libc.VaList(bp+8, value10+uintptr((int64(value10)-int64(value10))/4)*4))
	case int32(OPUS_SET_IGNORE_EXTENSIONS_REQUEST):
		value11 = libc.VaInt32(&ap)
		if value11 < 0 || value11 > int32(1) {
			goto bad_arg
		}
		(*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fignore_extensions = value11
	case int32(OPUS_GET_IGNORE_EXTENSIONS_REQUEST):
		value12 = libc.VaUintptr(&ap)
		if !(value12 != 0) {
			goto bad_arg
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(value12)) = (*OpusT_OpusDecoder)(unsafe.Pointer(st)).Fignore_extensions
	default:
		/*fprintf(stderr, "unknown opus_decoder_ctl() request: %d", request);*/
		ret = -int32(5)
		break
	}
	_ = ap
	return ret
	goto bad_arg
bad_arg:
	;
	_ = ap
	return -int32(1)
}

func Opus_opus_decoder_destroy(tls *libc.TLS, st uintptr) {
	libc.Xfree(tls, st)
}

func Opus_opus_packet_get_bandwidth(tls *libc.TLS, data uintptr) (r int32) {
	var bandwidth, v1 int32
	_, _ = bandwidth, v1
	if int32(*(*uint8)(unsafe.Pointer(data)))&int32(0x80) != 0 {
		bandwidth = int32(OPUS_BANDWIDTH_MEDIUMBAND) + int32(*(*uint8)(unsafe.Pointer(data)))>>int32(5)&int32(0x3)
		if bandwidth == int32(OPUS_BANDWIDTH_MEDIUMBAND) {
			bandwidth = int32(OPUS_BANDWIDTH_NARROWBAND)
		}
	} else {
		if int32(*(*uint8)(unsafe.Pointer(data)))&int32(0x60) == int32(0x60) {
			if int32(*(*uint8)(unsafe.Pointer(data)))&int32(0x10) != 0 {
				v1 = int32(OPUS_BANDWIDTH_FULLBAND)
			} else {
				v1 = int32(OPUS_BANDWIDTH_SUPERWIDEBAND)
			}
			bandwidth = v1
		} else {
			bandwidth = int32(OPUS_BANDWIDTH_NARROWBAND) + int32(*(*uint8)(unsafe.Pointer(data)))>>int32(5)&int32(0x3)
		}
	}
	return bandwidth
}

func Opus_opus_packet_get_nb_channels(tls *libc.TLS, data uintptr) (r int32) {
	var v1 int32
	_ = v1
	if int32(*(*uint8)(unsafe.Pointer(data)))&int32(0x4) != 0 {
		v1 = int32(2)
	} else {
		v1 = int32(1)
	}
	return v1
}

func Opus_opus_packet_get_nb_frames(tls *libc.TLS, packet uintptr, len1 OpusT_opus_int32) (r int32) {
	var count int32
	_ = count
	if len1 < int32(1) {
		return -int32(1)
	}
	count = int32(*(*uint8)(unsafe.Pointer(packet))) & int32(0x3)
	if count == 0 {
		return int32(1)
	} else {
		if count != int32(3) {
			return int32(2)
		} else {
			if len1 < int32(2) {
				return -int32(4)
			} else {
				return int32(*(*uint8)(unsafe.Pointer(packet + 1))) & int32(0x3F)
			}
		}
	}
	return r
}

func Opus_opus_packet_get_nb_samples(tls *libc.TLS, packet uintptr, len1 OpusT_opus_int32, Fs OpusT_opus_int32) (r int32) {
	var count, samples int32
	_, _ = count, samples
	count = Opus_opus_packet_get_nb_frames(tls, packet, len1)
	if count < 0 {
		return count
	}
	samples = count * Opus_opus_packet_get_samples_per_frame(tls, packet, Fs)
	/* Can't have more than 120 ms */
	if samples*int32(25) > Fs*int32(3) {
		return -int32(4)
	} else {
		return samples
	}
	return r
}

func Opus_opus_packet_has_lbrr(tls *libc.TLS, packet uintptr, len1 OpusT_opus_int32) (r int32) {
	bp := tls.Alloc(480)
	defer tls.Free(480)
	var lbrr, nb_frames, packet_frame_size, packet_mode, packet_stream_channels, ret int32
	var _ /* frames at bp+0 */ [48]uintptr
	var _ /* size at bp+384 */ [48]OpusT_opus_int16
	_, _, _, _, _, _ = lbrr, nb_frames, packet_frame_size, packet_mode, packet_stream_channels, ret
	nb_frames = int32(1)
	packet_mode = opus_packet_get_mode(tls, packet)
	if packet_mode == int32(MODE_CELT_ONLY) {
		return 0
	}
	packet_frame_size = Opus_opus_packet_get_samples_per_frame(tls, packet, int32(48000))
	if packet_frame_size > int32(960) {
		nb_frames = packet_frame_size / int32(960)
	}
	packet_stream_channels = Opus_opus_packet_get_nb_channels(tls, packet)
	ret = Opus_opus_packet_parse(tls, packet, len1, uintptr(uint32(0)), bp, bp+384, uintptr(uint32(0)))
	if ret <= 0 {
		return ret
	}
	if int32((*(*[48]OpusT_opus_int16)(unsafe.Pointer(bp + 384)))[0]) == 0 {
		return 0
	}
	lbrr = int32(*(*uint8)(unsafe.Pointer((*(*[48]uintptr)(unsafe.Pointer(bp)))[0]))) >> (int32(7) - nb_frames) & int32(0x1)
	if packet_stream_channels == int32(2) {
		lbrr = libc.BoolInt32(lbrr != 0 || int32(*(*uint8)(unsafe.Pointer((*(*[48]uintptr)(unsafe.Pointer(bp)))[0])))>>(int32(6)-int32(2)*nb_frames)&int32(0x1) != 0)
	}
	return lbrr
}

func Opus_opus_decoder_get_nb_samples(tls *libc.TLS, dec uintptr, packet uintptr, len1 OpusT_opus_int32) (r int32) {
	return Opus_opus_packet_get_nb_samples(tls, packet, len1, (*OpusT_OpusDecoder)(unsafe.Pointer(dec)).FFs)
}

type OpusDREDDecoder = struct {
	Floaded int32
	Farch   int32
	Fmagic  OpusT_opus_uint32
}

func Opus_opus_dred_decoder_get_size(tls *libc.TLS) (r int32) {
	return int32(12)
}

func Opus_opus_dred_decoder_init(tls *libc.TLS, dec uintptr) (r int32) {
	var ret, v1 int32
	_, _ = ret, v1
	ret = 0
	(*OpusT_OpusDREDDecoder)(unsafe.Pointer(dec)).Floaded = 0
	v1 = 0
	goto _2
_2:
	(*OpusT_OpusDREDDecoder)(unsafe.Pointer(dec)).Farch = v1
	/* To make sure nobody forgets to init, use a magic number. */
	(*OpusT_OpusDREDDecoder)(unsafe.Pointer(dec)).Fmagic = uint32(0xD8EDDEC0)
	if ret == 0 {
		v1 = OPUS_OK
	} else {
		v1 = -int32(5)
	}
	return v1
}

func Opus_opus_dred_decoder_create(tls *libc.TLS, error1 uintptr) (r uintptr) {
	var dec, v1 uintptr
	var ret int32
	_, _, _ = dec, ret, v1
	v1 = libc.Xmalloc(tls, uint64(uint32(Opus_opus_dred_decoder_get_size(tls))))
	goto _2
_2:
	dec = v1
	if dec == uintptr(uint32(0)) {
		if error1 != 0 {
			*(*int32)(unsafe.Pointer(error1)) = -int32(7)
		}
		return uintptr(uint32(0))
	}
	ret = Opus_opus_dred_decoder_init(tls, dec)
	if error1 != 0 {
		*(*int32)(unsafe.Pointer(error1)) = ret
	}
	if ret != OPUS_OK {
		libc.Xfree(tls, dec)
		dec = uintptr(uint32(0))
	}
	return dec
}

func Opus_opus_dred_decoder_destroy(tls *libc.TLS, dec uintptr) {
	if dec != 0 {
		(*OpusT_OpusDREDDecoder)(unsafe.Pointer(dec)).Fmagic = uint32(0xDE57801D)
	}
	libc.Xfree(tls, dec)
}

func Opus_opus_dred_decoder_ctl(tls *libc.TLS, dred_dec uintptr, request int32, va uintptr) (r int32) {
	_ = dred_dec
	_ = request
	return -int32(5)
}

func Opus_opus_dred_get_size(tls *libc.TLS) (r int32) {
	return 0
}

func Opus_opus_dred_alloc(tls *libc.TLS, error1 uintptr) (r uintptr) {
	if error1 != 0 {
		*(*int32)(unsafe.Pointer(error1)) = -int32(5)
	}
	return uintptr(uint32(0))
}

func Opus_opus_dred_free(tls *libc.TLS, dec uintptr) {
	_ = dec
}

func Opus_opus_dred_parse(tls *libc.TLS, dred_dec uintptr, dred uintptr, data uintptr, len1 OpusT_opus_int32, max_dred_samples OpusT_opus_int32, sampling_rate OpusT_opus_int32, dred_end uintptr, defer_processing int32) (r int32) {
	_ = dred_dec
	_ = dred
	_ = data
	_ = len1
	_ = max_dred_samples
	_ = sampling_rate
	_ = defer_processing
	_ = dred_end
	return -int32(5)
}

func Opus_opus_dred_process(tls *libc.TLS, dred_dec uintptr, src uintptr, dst uintptr) (r int32) {
	_ = dred_dec
	_ = src
	_ = dst
	return -int32(5)
}

func Opus_opus_decoder_dred_decode(tls *libc.TLS, st uintptr, dred uintptr, dred_offset OpusT_opus_int32, pcm uintptr, frame_size OpusT_opus_int32) (r int32) {
	_ = st
	_ = dred
	_ = dred_offset
	_ = pcm
	_ = frame_size
	return -int32(5)
}

func Opus_opus_decoder_dred_decode24(tls *libc.TLS, st uintptr, dred uintptr, dred_offset OpusT_opus_int32, pcm uintptr, frame_size OpusT_opus_int32) (r int32) {
	_ = st
	_ = dred
	_ = dred_offset
	_ = pcm
	_ = frame_size
	return -int32(5)
}

func Opus_opus_decoder_dred_decode_float(tls *libc.TLS, st uintptr, dred uintptr, dred_offset OpusT_opus_int32, pcm uintptr, frame_size OpusT_opus_int32) (r int32) {
	_ = st
	_ = dred
	_ = dred_offset
	_ = pcm
	_ = frame_size
	return -int32(5)
}

const COEF_ONE2 = "1.0f"
const OPUS_MULTISTREAM_GET_DECODER_STATE_REQUEST = 5122
const OPUS_MULTISTREAM_GET_ENCODER_STATE_REQUEST = 5120

type OpusT_OpusMSEncoder = struct {
	Flayout            OpusT_ChannelLayout
	Farch              int32
	Flfe_stream        int32
	Fapplication       int32
	FFs                OpusT_opus_int32
	Fvariable_duration int32
	Fmapping_type      OpusT_MappingType
	Fbitrate_bps       OpusT_opus_int32
}

type OpusT_OpusMSDecoder = struct {
	Flayout OpusT_ChannelLayout
}

var trim_icdf2 = [11]uint8{
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
var spread_icdf2 = [4]uint8{
	0: uint8(25),
	1: uint8(23),
	2: uint8(2),
}
var tapset_icdf2 = [3]uint8{
	0: uint8(2),
	1: uint8(1),
}

/* Copyright (C) 2007 Jean-Marc Valin

   File: os_support.h
   This is the (tiny) OS abstraction layer. Aside from math.h, this is the
   only place where system headers are allowed.

   Redistribution and use in source and binary forms, with or without
   modification, are permitted provided that the following conditions are
   met:

   1. Redistributions of source code must retain the above copyright notice,
   this list of conditions and the following disclaimer.

   2. Redistributions in binary form must reproduce the above copyright
   notice, this list of conditions and the following disclaimer in the
   documentation and/or other materials provided with the distribution.

   THIS SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR
   IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
   OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
   DISCLAIMED. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT,
   INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
   (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
   SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
   HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
   STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN
   ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
   POSSIBILITY OF SUCH DAMAGE.
*/

func Opus_validate_layout(tls *libc.TLS, layout uintptr) (r int32) {
	var i, max_channel int32
	_, _ = i, max_channel
	max_channel = (*OpusT_ChannelLayout)(unsafe.Pointer(layout)).Fnb_streams + (*OpusT_ChannelLayout)(unsafe.Pointer(layout)).Fnb_coupled_streams
	if max_channel > int32(255) {
		return 0
	}
	i = 0
	for {
		if !(i < (*OpusT_ChannelLayout)(unsafe.Pointer(layout)).Fnb_channels) {
			break
		}
		if int32(*(*uint8)(unsafe.Pointer(layout + 12 + uintptr(i)))) >= max_channel && int32(*(*uint8)(unsafe.Pointer(layout + 12 + uintptr(i)))) != int32(255) {
			return 0
		}
		goto _1
	_1:
		;
		i = i + 1
	}
	return int32(1)
}

func Opus_get_left_channel(tls *libc.TLS, layout uintptr, stream_id int32, prev int32) (r int32) {
	var i, v1 int32
	_, _ = i, v1
	if prev < 0 {
		v1 = 0
	} else {
		v1 = prev + int32(1)
	}
	i = v1
	for {
		if !(i < (*OpusT_ChannelLayout)(unsafe.Pointer(layout)).Fnb_channels) {
			break
		}
		if int32(*(*uint8)(unsafe.Pointer(layout + 12 + uintptr(i)))) == stream_id*int32(2) {
			return i
		}
		goto _2
	_2:
		;
		i = i + 1
	}
	return -int32(1)
}

func Opus_get_right_channel(tls *libc.TLS, layout uintptr, stream_id int32, prev int32) (r int32) {
	var i, v1 int32
	_, _ = i, v1
	if prev < 0 {
		v1 = 0
	} else {
		v1 = prev + int32(1)
	}
	i = v1
	for {
		if !(i < (*OpusT_ChannelLayout)(unsafe.Pointer(layout)).Fnb_channels) {
			break
		}
		if int32(*(*uint8)(unsafe.Pointer(layout + 12 + uintptr(i)))) == stream_id*int32(2)+int32(1) {
			return i
		}
		goto _2
	_2:
		;
		i = i + 1
	}
	return -int32(1)
}

func Opus_get_mono_channel(tls *libc.TLS, layout uintptr, stream_id int32, prev int32) (r int32) {
	var i, v1 int32
	_, _ = i, v1
	if prev < 0 {
		v1 = 0
	} else {
		v1 = prev + int32(1)
	}
	i = v1
	for {
		if !(i < (*OpusT_ChannelLayout)(unsafe.Pointer(layout)).Fnb_channels) {
			break
		}
		if int32(*(*uint8)(unsafe.Pointer(layout + 12 + uintptr(i)))) == stream_id+(*OpusT_ChannelLayout)(unsafe.Pointer(layout)).Fnb_coupled_streams {
			return i
		}
		goto _2
	_2:
		;
		i = i + 1
	}
	return -int32(1)
}

var trim_icdf3 = [11]uint8{
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
var spread_icdf3 = [4]uint8{
	0: uint8(25),
	1: uint8(23),
	2: uint8(2),
}
var tapset_icdf3 = [3]uint8{
	0: uint8(2),
	1: uint8(1),
}

/* Copyright (C) 2007 Jean-Marc Valin

   File: os_support.h
   This is the (tiny) OS abstraction layer. Aside from math.h, this is the
   only place where system headers are allowed.

   Redistribution and use in source and binary forms, with or without
   modification, are permitted provided that the following conditions are
   met:

   1. Redistributions of source code must retain the above copyright notice,
   this list of conditions and the following disclaimer.

   2. Redistributions in binary form must reproduce the above copyright
   notice, this list of conditions and the following disclaimer in the
   documentation and/or other materials provided with the distribution.

   THIS SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR
   IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
   OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
   DISCLAIMED. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT,
   INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
   (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
   SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
   HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
   STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN
   ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
   POSSIBILITY OF SUCH DAMAGE.
*/

/* DECODER */

func validate_ms_decoder(tls *libc.TLS, st uintptr) {
	Opus_validate_layout(tls, st)
}

func Opus_opus_multistream_decoder_get_size(tls *libc.TLS, nb_streams int32, nb_coupled_streams int32) (r OpusT_opus_int32) {
	var alignment uint32
	var coupled_size, mono_size, v1, v3, v5 int32
	_, _, _, _, _, _ = alignment, coupled_size, mono_size, v1, v3, v5
	if nb_streams < int32(1) || nb_coupled_streams > nb_streams || nb_coupled_streams < 0 {
		return 0
	}
	coupled_size = Opus_opus_decoder_get_size(tls, int32(2))
	mono_size = Opus_opus_decoder_get_size(tls, int32(1))
	alignment = uint32(uint64(uintptr(uint32(0)) + 8))
	v1 = int32((uint32(int32(268)) + alignment - uint32(1)) / alignment * alignment)
	goto _2
_2:
	alignment = uint32(uint64(uintptr(uint32(0)) + 8))
	v3 = int32((uint32(coupled_size) + alignment - uint32(1)) / alignment * alignment)
	goto _4
_4:
	alignment = uint32(uint64(uintptr(uint32(0)) + 8))
	v5 = int32((uint32(mono_size) + alignment - uint32(1)) / alignment * alignment)
	goto _6
_6:
	return v1 + nb_coupled_streams*v3 + (nb_streams-nb_coupled_streams)*v5
}

func Opus_opus_multistream_decoder_init(tls *libc.TLS, st uintptr, Fs OpusT_opus_int32, channels int32, streams int32, coupled_streams int32, mapping uintptr) (r int32) {
	var alignment uint32
	var coupled_size, i1, mono_size, ret, v2 int32
	var ptr uintptr
	_, _, _, _, _, _, _ = alignment, coupled_size, i1, mono_size, ptr, ret, v2
	if channels > int32(255) || channels < int32(1) || coupled_streams > streams || streams < int32(1) || coupled_streams < 0 || streams > int32(255)-coupled_streams {
		return -int32(1)
	}
	(*OpusT_OpusMSDecoder)(unsafe.Pointer(st)).Flayout.Fnb_channels = channels
	(*OpusT_OpusMSDecoder)(unsafe.Pointer(st)).Flayout.Fnb_streams = streams
	(*OpusT_OpusMSDecoder)(unsafe.Pointer(st)).Flayout.Fnb_coupled_streams = coupled_streams
	i1 = 0
	for {
		if !(i1 < (*OpusT_OpusMSDecoder)(unsafe.Pointer(st)).Flayout.Fnb_channels) {
			break
		}
		*(*uint8)(unsafe.Pointer(st + 12 + uintptr(i1))) = *(*uint8)(unsafe.Pointer(mapping + uintptr(i1)))
		goto _1
	_1:
		;
		i1 = i1 + 1
	}
	if !(Opus_validate_layout(tls, st) != 0) {
		return -int32(1)
	}
	alignment = uint32(uint64(uintptr(uint32(0)) + 8))
	v2 = int32((uint32(int32(268)) + alignment - uint32(1)) / alignment * alignment)
	goto _3
_3:
	ptr = st + uintptr(v2)
	coupled_size = Opus_opus_decoder_get_size(tls, int32(2))
	mono_size = Opus_opus_decoder_get_size(tls, int32(1))
	i1 = 0
	for {
		if !(i1 < (*OpusT_OpusMSDecoder)(unsafe.Pointer(st)).Flayout.Fnb_coupled_streams) {
			break
		}
		ret = Opus_opus_decoder_init(tls, ptr, Fs, int32(2))
		if ret != OPUS_OK {
			return ret
		}
		alignment = uint32(uint64(uintptr(uint32(0)) + 8))
		v2 = int32((uint32(coupled_size) + alignment - uint32(1)) / alignment * alignment)
		goto _6
	_6:
		ptr = ptr + uintptr(v2)
		goto _4
	_4:
		;
		i1 = i1 + 1
	}
	for {
		if !(i1 < (*OpusT_OpusMSDecoder)(unsafe.Pointer(st)).Flayout.Fnb_streams) {
			break
		}
		ret = Opus_opus_decoder_init(tls, ptr, Fs, int32(1))
		if ret != OPUS_OK {
			return ret
		}
		alignment = uint32(uint64(uintptr(uint32(0)) + 8))
		v2 = int32((uint32(mono_size) + alignment - uint32(1)) / alignment * alignment)
		goto _9
	_9:
		ptr = ptr + uintptr(v2)
		goto _7
	_7:
		;
		i1 = i1 + 1
	}
	return OPUS_OK
}

func Opus_opus_multistream_decoder_create(tls *libc.TLS, Fs OpusT_opus_int32, channels int32, streams int32, coupled_streams int32, mapping uintptr, error1 uintptr) (r uintptr) {
	var ret int32
	var st, v1 uintptr
	_, _, _ = ret, st, v1
	if channels > int32(255) || channels < int32(1) || coupled_streams > streams || streams < int32(1) || coupled_streams < 0 || streams > int32(255)-coupled_streams {
		if error1 != 0 {
			*(*int32)(unsafe.Pointer(error1)) = -int32(1)
		}
		return uintptr(uint32(0))
	}
	v1 = libc.Xmalloc(tls, uint64(uint32(Opus_opus_multistream_decoder_get_size(tls, streams, coupled_streams))))
	goto _2
_2:
	st = v1
	if st == uintptr(uint32(0)) {
		if error1 != 0 {
			*(*int32)(unsafe.Pointer(error1)) = -int32(7)
		}
		return uintptr(uint32(0))
	}
	ret = Opus_opus_multistream_decoder_init(tls, st, Fs, channels, streams, coupled_streams, mapping)
	if error1 != 0 {
		*(*int32)(unsafe.Pointer(error1)) = ret
	}
	if ret != OPUS_OK {
		libc.Xfree(tls, st)
		st = uintptr(uint32(0))
	}
	return st
}

func opus_multistream_packet_validate(tls *libc.TLS, data uintptr, len1 OpusT_opus_int32, nb_streams int32, Fs OpusT_opus_int32) (r int32) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var count, s, samples, tmp_samples int32
	var _ /* packet_offset at bp+100 */ OpusT_opus_int32
	var _ /* size at bp+2 */ [48]OpusT_opus_int16
	var _ /* toc at bp+0 */ uint8
	_, _, _, _ = count, s, samples, tmp_samples
	samples = 0
	s = 0
	for {
		if !(s < nb_streams) {
			break
		}
		if len1 <= 0 {
			return -int32(4)
		}
		count = Opus_opus_packet_parse_impl(tls, data, len1, libc.BoolInt32(s != nb_streams-int32(1)), bp, uintptr(uint32(0)), bp+2, uintptr(uint32(0)), bp+100, uintptr(uint32(0)), uintptr(uint32(0)))
		if count < 0 {
			return count
		}
		tmp_samples = Opus_opus_packet_get_nb_samples(tls, data, *(*OpusT_opus_int32)(unsafe.Pointer(bp + 100)), Fs)
		if s != 0 && samples != tmp_samples {
			return -int32(4)
		}
		samples = tmp_samples
		data = data + uintptr(*(*OpusT_opus_int32)(unsafe.Pointer(bp + 100)))
		len1 = len1 - *(*OpusT_opus_int32)(unsafe.Pointer(bp + 100))
		goto _1
	_1:
		;
		s = s + 1
	}
	return samples
}

type OpusT___ccgo_fp__Xopus_multistream_decode_native_4 = func(*libc.TLS, uintptr, int32, int32, uintptr, int32, int32, uintptr)

func Opus_opus_multistream_decode_native(tls *libc.TLS, st1 uintptr, data uintptr, len1 OpusT_opus_int32, pcm uintptr, __ccgo_fp_copy_channel_out OpusT_opus_copy_channel_out_func, frame_size int32, decode_fec int32, soft_clip int32, user_data uintptr) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var _saved_stack, buf, dec, ptr, st, v1, v10, v11, v13, v15, v17, v19, v21, v3, v5, v6, v8 uintptr
	var alignment uint32
	var c, chan1, chan11, coupled_size, do_plc, mono_size, prev, prev1, ret, ret1, s, v31, v56, v75 int32
	var _ /* Fs at bp+0 */ OpusT_opus_int32
	var _ /* packet_offset at bp+4 */ OpusT_opus_int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = _saved_stack, alignment, buf, c, chan1, chan11, coupled_size, dec, do_plc, mono_size, prev, prev1, ptr, ret, ret1, s, st, v1, v10, v11, v13, v15, v17, v19, v21, v3, v31, v5, v56, v6, v75, v8
	do_plc = 0
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
	validate_ms_decoder(tls, st1)
	if frame_size <= 0 {
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			goto _28
		_28:
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v3 = st
		goto _30
	_30:
		(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
		return -int32(1)
	}
	/* Limit frame_size to avoid excessive stack allocations. */
	if !(Opus_opus_multistream_decoder_ctl(tls, st1, int32(OPUS_GET_SAMPLE_RATE_REQUEST), libc.VaList(bp+16, bp+uintptr((OpusT___predefined_ptrdiff_t(bp)-int64(bp))/4)*4)) == int32(OPUS_OK)) {
		Opus_celt_fatal(tls, __ccgo_ts+2090, __ccgo_ts+2200, int32(206))
	}
	if frame_size < *(*OpusT_opus_int32)(unsafe.Pointer(bp))/int32(25)*int32(3) {
		v31 = frame_size
	} else {
		v31 = *(*OpusT_opus_int32)(unsafe.Pointer(bp)) / int32(25) * int32(3)
	}
	frame_size = v31
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _33
	_33:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _35
_35:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _37
	_37:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	goto _39
_39:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(4)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(4)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _41
	_41:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	goto _43
_43:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _45
	_45:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _47
_47:
	if !(int64(int32(uint64(uint32(int32(2)*frame_size))*(uint64(4)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+2200, int32(208))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _49
	_49:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _51
_51:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(uint64(uint32(int32(2)*frame_size)) * (uint64(4) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _53
	_53:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _55
_55:
	buf = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(int32(2)*frame_size))*(uint64(4)/uint64(1)))
	alignment = uint32(uint64(uintptr(uint32(0)) + 8))
	v31 = int32((uint32(int32(268)) + alignment - uint32(1)) / alignment * alignment)
	goto _57
_57:
	ptr = st1 + uintptr(v31)
	coupled_size = Opus_opus_decoder_get_size(tls, int32(2))
	mono_size = Opus_opus_decoder_get_size(tls, int32(1))
	if len1 == 0 {
		do_plc = int32(1)
	}
	if len1 < 0 {
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			goto _59
		_59:
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v3 = st
		goto _61
	_61:
		(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
		return -int32(1)
	}
	if !(do_plc != 0) && len1 < int32(2)*(*OpusT_OpusMSDecoder)(unsafe.Pointer(st1)).Flayout.Fnb_streams-int32(1) {
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			goto _63
		_63:
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v3 = st
		goto _65
	_65:
		(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
		return -int32(4)
	}
	if !(do_plc != 0) {
		ret = opus_multistream_packet_validate(tls, data, len1, (*OpusT_OpusMSDecoder)(unsafe.Pointer(st1)).Flayout.Fnb_streams, *(*OpusT_opus_int32)(unsafe.Pointer(bp)))
		if ret < 0 {
			st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
			if !(st != 0) {
				v1 = libc.Xmalloc(tls, uint64(16))
				goto _67
			_67:
				st = v1
				if st != 0 {
					libc.Xmemset(tls, st, 0, uint64(16))
				}
				libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
			}
			v3 = st
			goto _69
		_69:
			(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
			return ret
		} else {
			if ret > frame_size {
				st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
				if !(st != 0) {
					v1 = libc.Xmalloc(tls, uint64(16))
					goto _71
				_71:
					st = v1
					if st != 0 {
						libc.Xmemset(tls, st, 0, uint64(16))
					}
					libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
				}
				v3 = st
				goto _73
			_73:
				(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
				return -int32(2)
			}
		}
	}
	s = 0
	for {
		if !(s < (*OpusT_OpusMSDecoder)(unsafe.Pointer(st1)).Flayout.Fnb_streams) {
			break
		}
		dec = ptr
		if s < (*OpusT_OpusMSDecoder)(unsafe.Pointer(st1)).Flayout.Fnb_coupled_streams {
			alignment = uint32(uint64(uintptr(uint32(0)) + 8))
			v56 = int32((uint32(coupled_size) + alignment - uint32(1)) / alignment * alignment)
			goto _77
		_77:
			v31 = v56
		} else {
			alignment = uint32(uint64(uintptr(uint32(0)) + 8))
			v75 = int32((uint32(mono_size) + alignment - uint32(1)) / alignment * alignment)
			goto _79
		_79:
			v31 = v75
		}
		ptr = ptr + uintptr(v31)
		if !(do_plc != 0) && len1 <= 0 {
			st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
			if !(st != 0) {
				v1 = libc.Xmalloc(tls, uint64(16))
				goto _81
			_81:
				st = v1
				if st != 0 {
					libc.Xmemset(tls, st, 0, uint64(16))
				}
				libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
			}
			v3 = st
			goto _83
		_83:
			(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
			return -int32(3)
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(bp + 4)) = 0
		ret1 = Opus_opus_decode_native(tls, dec, data, len1, buf, frame_size, decode_fec, libc.BoolInt32(s != (*OpusT_OpusMSDecoder)(unsafe.Pointer(st1)).Flayout.Fnb_streams-int32(1)), bp+4, soft_clip, uintptr(uint32(0)), 0)
		if !(do_plc != 0) {
			data = data + uintptr(*(*OpusT_opus_int32)(unsafe.Pointer(bp + 4)))
			len1 = len1 - *(*OpusT_opus_int32)(unsafe.Pointer(bp + 4))
		}
		if ret1 <= 0 {
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
			return ret1
		}
		frame_size = ret1
		if s < (*OpusT_OpusMSDecoder)(unsafe.Pointer(st1)).Flayout.Fnb_coupled_streams {
			prev = -int32(1)
			/* Copy "left" audio to the channel(s) where it belongs */
			for {
				v31 = Opus_get_left_channel(tls, st1, s, prev)
				chan1 = v31
				if !(v31 != -int32(1)) {
					break
				}
				(*(*func(*libc.TLS, uintptr, int32, int32, uintptr, int32, int32, uintptr))(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_copy_channel_out})))(tls, pcm, (*OpusT_OpusMSDecoder)(unsafe.Pointer(st1)).Flayout.Fnb_channels, chan1, buf, int32(2), frame_size, user_data)
				prev = chan1
			}
			prev = -int32(1)
			/* Copy "right" audio to the channel(s) where it belongs */
			for {
				v31 = Opus_get_right_channel(tls, st1, s, prev)
				chan1 = v31
				if !(v31 != -int32(1)) {
					break
				}
				(*(*func(*libc.TLS, uintptr, int32, int32, uintptr, int32, int32, uintptr))(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_copy_channel_out})))(tls, pcm, (*OpusT_OpusMSDecoder)(unsafe.Pointer(st1)).Flayout.Fnb_channels, chan1, buf+uintptr(1)*4, int32(2), frame_size, user_data)
				prev = chan1
			}
		} else {
			prev1 = -int32(1)
			/* Copy audio to the channel(s) where it belongs */
			for {
				v31 = Opus_get_mono_channel(tls, st1, s, prev1)
				chan11 = v31
				if !(v31 != -int32(1)) {
					break
				}
				(*(*func(*libc.TLS, uintptr, int32, int32, uintptr, int32, int32, uintptr))(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_copy_channel_out})))(tls, pcm, (*OpusT_OpusMSDecoder)(unsafe.Pointer(st1)).Flayout.Fnb_channels, chan11, buf, int32(1), frame_size, user_data)
				prev1 = chan11
			}
		}
		goto _74
	_74:
		;
		s = s + 1
	}
	/* Handle muted channels */
	c = 0
	for {
		if !(c < (*OpusT_OpusMSDecoder)(unsafe.Pointer(st1)).Flayout.Fnb_channels) {
			break
		}
		if int32(*(*uint8)(unsafe.Pointer(st1 + 12 + uintptr(c)))) == int32(255) {
			(*(*func(*libc.TLS, uintptr, int32, int32, uintptr, int32, int32, uintptr))(unsafe.Pointer(&struct{ uintptr }{__ccgo_fp_copy_channel_out})))(tls, pcm, (*OpusT_OpusMSDecoder)(unsafe.Pointer(st1)).Flayout.Fnb_channels, c, uintptr(uint32(0)), 0, frame_size, user_data)
		}
		goto _91
	_91:
		;
		c = c + 1
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _93
	_93:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _95
_95:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
	return frame_size
}

func opus_copy_channel_out_float(tls *libc.TLS, dst uintptr, dst_stride int32, dst_channel int32, src uintptr, src_stride int32, frame_size int32, user_data uintptr) {
	var float_dst uintptr
	var i OpusT_opus_int32
	_, _ = float_dst, i
	_ = user_data
	float_dst = dst
	if src != uintptr(uint32(0)) {
		i = 0
		for {
			if !(i < frame_size) {
				break
			}
			*(*float32)(unsafe.Pointer(float_dst + uintptr(i*dst_stride+dst_channel)*4)) = *(*OpusT_opus_res)(unsafe.Pointer(src + uintptr(i*src_stride)*4))
			goto _1
		_1:
			;
			i = i + 1
		}
	} else {
		i = 0
		for {
			if !(i < frame_size) {
				break
			}
			*(*float32)(unsafe.Pointer(float_dst + uintptr(i*dst_stride+dst_channel)*4)) = float32(0)
			goto _2
		_2:
			;
			i = i + 1
		}
	}
}

func opus_copy_channel_out_short(tls *libc.TLS, dst uintptr, dst_stride int32, dst_channel int32, src uintptr, src_stride int32, frame_size int32, user_data uintptr) {
	var i OpusT_opus_int32
	var short_dst uintptr
	var v2, v3, v4 float32
	var v5 OpusT_opus_int16
	_, _, _, _, _, _ = i, short_dst, v2, v3, v4, v5
	_ = user_data
	short_dst = dst
	if src != uintptr(uint32(0)) {
		i = 0
		for {
			if !(i < frame_size) {
				break
			}
			v2 = *(*OpusT_opus_res)(unsafe.Pointer(src + uintptr(i*src_stride)*4))
			v2 = float32(v2 * float32(32768))
			if v2 > float32(-int32(32768)) {
				v3 = v2
			} else {
				v3 = float32(-int32(32768))
			}
			v2 = v3
			if v2 < float32(int32(32767)) {
				v4 = v2
			} else {
				v4 = float32(int32(32767))
			}
			v2 = v4
			v5 = int16(libc.Xlrintf(tls, v2))
			goto _6
		_6:
			*(*OpusT_opus_int16)(unsafe.Pointer(short_dst + uintptr(i*dst_stride+dst_channel)*2)) = v5
			goto _1
		_1:
			;
			i = i + 1
		}
	} else {
		i = 0
		for {
			if !(i < frame_size) {
				break
			}
			*(*OpusT_opus_int16)(unsafe.Pointer(short_dst + uintptr(i*dst_stride+dst_channel)*2)) = 0
			goto _7
		_7:
			;
			i = i + 1
		}
	}
}

func opus_copy_channel_out_int24(tls *libc.TLS, dst uintptr, dst_stride int32, dst_channel int32, src uintptr, src_stride int32, frame_size int32, user_data uintptr) {
	var i OpusT_opus_int32
	var short_dst uintptr
	_, _ = i, short_dst
	_ = user_data
	short_dst = dst
	if src != uintptr(uint32(0)) {
		i = 0
		for {
			if !(i < frame_size) {
				break
			}
			*(*OpusT_opus_int32)(unsafe.Pointer(short_dst + uintptr(i*dst_stride+dst_channel)*4)) = int32(libc.Xlrintf(tls, float32(float32(float32(32768)*float32(256))**(*OpusT_opus_res)(unsafe.Pointer(src + uintptr(i*src_stride)*4)))))
			goto _1
		_1:
			;
			i = i + 1
		}
	} else {
		i = 0
		for {
			if !(i < frame_size) {
				break
			}
			*(*OpusT_opus_int32)(unsafe.Pointer(short_dst + uintptr(i*dst_stride+dst_channel)*4)) = 0
			goto _2
		_2:
			;
			i = i + 1
		}
	}
}

func Opus_opus_multistream_decode(tls *libc.TLS, st uintptr, data uintptr, len1 OpusT_opus_int32, pcm uintptr, frame_size int32, decode_fec int32) (r int32) {
	return Opus_opus_multistream_decode_native(tls, st, data, len1, pcm, __ccgo_fp(opus_copy_channel_out_short), frame_size, decode_fec, int32(OPTIONAL_CLIP), uintptr(uint32(0)))
}

func Opus_opus_multistream_decode24(tls *libc.TLS, st uintptr, data uintptr, len1 OpusT_opus_int32, pcm uintptr, frame_size int32, decode_fec int32) (r int32) {
	return Opus_opus_multistream_decode_native(tls, st, data, len1, pcm, __ccgo_fp(opus_copy_channel_out_int24), frame_size, decode_fec, 0, uintptr(uint32(0)))
}

func Opus_opus_multistream_decode_float(tls *libc.TLS, st uintptr, data uintptr, len1 OpusT_opus_int32, pcm uintptr, frame_size int32, decode_fec int32) (r int32) {
	return Opus_opus_multistream_decode_native(tls, st, data, len1, pcm, __ccgo_fp(opus_copy_channel_out_float), frame_size, decode_fec, 0, uintptr(uint32(0)))
}

func Opus_opus_multistream_decoder_ctl_va_list(tls *libc.TLS, st uintptr, request int32, ap OpusT_va_list) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var alignment uint32
	var coupled_size, mono_size, ret, s, s1, s2, s3, v1 int32
	var dec, dec1, dec2, dec3, ptr, value, value1, value2 uintptr
	var stream_id, value3 OpusT_opus_int32
	var _ /* tmp at bp+0 */ OpusT_opus_uint32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = alignment, coupled_size, dec, dec1, dec2, dec3, mono_size, ptr, ret, s, s1, s2, s3, stream_id, value, value1, value2, value3, v1
	ret = OPUS_OK
	coupled_size = Opus_opus_decoder_get_size(tls, int32(2))
	mono_size = Opus_opus_decoder_get_size(tls, int32(1))
	alignment = uint32(uint64(uintptr(uint32(0)) + 8))
	v1 = int32((uint32(int32(268)) + alignment - uint32(1)) / alignment * alignment)
	goto _2
_2:
	ptr = st + uintptr(v1)
	switch request {
	case int32(OPUS_GET_BANDWIDTH_REQUEST):
		fallthrough
	case int32(OPUS_GET_SAMPLE_RATE_REQUEST):
		fallthrough
	case int32(OPUS_GET_GAIN_REQUEST):
		fallthrough
	case int32(OPUS_GET_LAST_PACKET_DURATION_REQUEST):
		fallthrough
	case int32(OPUS_GET_PHASE_INVERSION_DISABLED_REQUEST):
		fallthrough
	case int32(OPUS_GET_COMPLEXITY_REQUEST):
		/* For int32* GET params, just query the first stream */
		value = libc.VaUintptr(&ap)
		dec = ptr
		ret = Opus_opus_decoder_ctl(tls, dec, request, libc.VaList(bp+16, value))
	case int32(OPUS_GET_FINAL_RANGE_REQUEST):
		value1 = libc.VaUintptr(&ap)
		if !(value1 != 0) {
			goto bad_arg
		}
		*(*OpusT_opus_uint32)(unsafe.Pointer(value1)) = uint32(0)
		s = 0
		for {
			if !(s < (*OpusT_OpusMSDecoder)(unsafe.Pointer(st)).Flayout.Fnb_streams) {
				break
			}
			dec1 = ptr
			if s < (*OpusT_OpusMSDecoder)(unsafe.Pointer(st)).Flayout.Fnb_coupled_streams {
				alignment = uint32(uint64(uintptr(uint32(0)) + 8))
				v1 = int32((uint32(coupled_size) + alignment - uint32(1)) / alignment * alignment)
				goto _5
			_5:
				ptr = ptr + uintptr(v1)
			} else {
				alignment = uint32(uint64(uintptr(uint32(0)) + 8))
				v1 = int32((uint32(mono_size) + alignment - uint32(1)) / alignment * alignment)
				goto _7
			_7:
				ptr = ptr + uintptr(v1)
			}
			ret = Opus_opus_decoder_ctl(tls, dec1, request, libc.VaList(bp+16, bp))
			if ret != OPUS_OK {
				break
			}
			*(*OpusT_opus_uint32)(unsafe.Pointer(value1)) ^= *(*OpusT_opus_uint32)(unsafe.Pointer(bp))
			goto _3
		_3:
			;
			s = s + 1
		}
	case int32(OPUS_RESET_STATE):
		s1 = 0
		for {
			if !(s1 < (*OpusT_OpusMSDecoder)(unsafe.Pointer(st)).Flayout.Fnb_streams) {
				break
			}
			dec2 = ptr
			if s1 < (*OpusT_OpusMSDecoder)(unsafe.Pointer(st)).Flayout.Fnb_coupled_streams {
				alignment = uint32(uint64(uintptr(uint32(0)) + 8))
				v1 = int32((uint32(coupled_size) + alignment - uint32(1)) / alignment * alignment)
				goto _10
			_10:
				ptr = ptr + uintptr(v1)
			} else {
				alignment = uint32(uint64(uintptr(uint32(0)) + 8))
				v1 = int32((uint32(mono_size) + alignment - uint32(1)) / alignment * alignment)
				goto _12
			_12:
				ptr = ptr + uintptr(v1)
			}
			ret = Opus_opus_decoder_ctl(tls, dec2, int32(OPUS_RESET_STATE), 0)
			if ret != OPUS_OK {
				break
			}
			goto _8
		_8:
			;
			s1 = s1 + 1
		}
	case int32(OPUS_MULTISTREAM_GET_DECODER_STATE_REQUEST):
		stream_id = libc.VaInt32(&ap)
		if stream_id < 0 || stream_id >= (*OpusT_OpusMSDecoder)(unsafe.Pointer(st)).Flayout.Fnb_streams {
			goto bad_arg
		}
		value2 = libc.VaUintptr(&ap)
		if !(value2 != 0) {
			goto bad_arg
		}
		s2 = 0
		for {
			if !(s2 < stream_id) {
				break
			}
			if s2 < (*OpusT_OpusMSDecoder)(unsafe.Pointer(st)).Flayout.Fnb_coupled_streams {
				alignment = uint32(uint64(uintptr(uint32(0)) + 8))
				v1 = int32((uint32(coupled_size) + alignment - uint32(1)) / alignment * alignment)
				goto _15
			_15:
				ptr = ptr + uintptr(v1)
			} else {
				alignment = uint32(uint64(uintptr(uint32(0)) + 8))
				v1 = int32((uint32(mono_size) + alignment - uint32(1)) / alignment * alignment)
				goto _17
			_17:
				ptr = ptr + uintptr(v1)
			}
			goto _13
		_13:
			;
			s2 = s2 + 1
		}
		*(*uintptr)(unsafe.Pointer(value2)) = ptr
	case int32(OPUS_SET_GAIN_REQUEST):
		fallthrough
	case int32(OPUS_SET_COMPLEXITY_REQUEST):
		fallthrough
	case int32(OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST):
		/* This works for int32 params */
		value3 = libc.VaInt32(&ap)
		s3 = 0
		for {
			if !(s3 < (*OpusT_OpusMSDecoder)(unsafe.Pointer(st)).Flayout.Fnb_streams) {
				break
			}
			dec3 = ptr
			if s3 < (*OpusT_OpusMSDecoder)(unsafe.Pointer(st)).Flayout.Fnb_coupled_streams {
				alignment = uint32(uint64(uintptr(uint32(0)) + 8))
				v1 = int32((uint32(coupled_size) + alignment - uint32(1)) / alignment * alignment)
				goto _20
			_20:
				ptr = ptr + uintptr(v1)
			} else {
				alignment = uint32(uint64(uintptr(uint32(0)) + 8))
				v1 = int32((uint32(mono_size) + alignment - uint32(1)) / alignment * alignment)
				goto _22
			_22:
				ptr = ptr + uintptr(v1)
			}
			ret = Opus_opus_decoder_ctl(tls, dec3, request, libc.VaList(bp+16, value3))
			if ret != OPUS_OK {
				break
			}
			goto _18
		_18:
			;
			s3 = s3 + 1
		}
	default:
		ret = -int32(5)
		break
	}
	return ret
	goto bad_arg
bad_arg:
	;
	return -int32(1)
	return r
}

func Opus_opus_multistream_decoder_ctl(tls *libc.TLS, st uintptr, request int32, va uintptr) (r int32) {
	var ap OpusT_va_list
	var ret int32
	_, _ = ap, ret
	ap = va
	ret = Opus_opus_multistream_decoder_ctl_va_list(tls, st, request, ap)
	_ = ap
	return ret
}

func Opus_opus_multistream_decoder_destroy(tls *libc.TLS, st uintptr) {
	libc.Xfree(tls, st)
}

const OPUS_PROJECTION_GET_DEMIXING_MATRIX_GAIN_REQUEST = 6001
const OPUS_PROJECTION_GET_DEMIXING_MATRIX_REQUEST = 6005
const OPUS_PROJECTION_GET_DEMIXING_MATRIX_SIZE_REQUEST = 6003

var trim_icdf4 = [11]uint8{
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
var spread_icdf4 = [4]uint8{
	0: uint8(25),
	1: uint8(23),
	2: uint8(2),
}
var tapset_icdf4 = [3]uint8{
	0: uint8(2),
	1: uint8(1),
}

type OpusT_MappingMatrix = struct {
	Frows int32
	Fcols int32
	Fgain int32
}

func Opus_mapping_matrix_get_size(tls *libc.TLS, rows int32, cols int32) (r OpusT_opus_int32) {
	var alignment uint32
	var size OpusT_opus_int32
	var v1, v3 int32
	_, _, _, _ = alignment, size, v1, v3
	/* Mapping Matrix must only support up to 255 channels in or out.
	 * Additionally, the total cell count must be <= 65004 octets in order
	 * for the matrix to be stored in an OGG header.
	 */
	if rows > int32(255) || cols > int32(255) {
		return 0
	}
	size = int32(uint64(uint32(rows*cols)) * uint64(2))
	if size > int32(65004) {
		return 0
	}
	alignment = uint32(uint64(uintptr(uint32(0)) + 8))
	v1 = int32((uint32(int32(12)) + alignment - uint32(1)) / alignment * alignment)
	goto _2
_2:
	alignment = uint32(uint64(uintptr(uint32(0)) + 8))
	v3 = int32((uint32(size) + alignment - uint32(1)) / alignment * alignment)
	goto _4
_4:
	return v1 + v3
}

func Opus_mapping_matrix_get_data(tls *libc.TLS, matrix uintptr) (r uintptr) {
	var alignment uint32
	var v1 int32
	_, _ = alignment, v1
	/* void* cast avoids clang -Wcast-align warning */
	alignment = uint32(uint64(uintptr(uint32(0)) + 8))
	v1 = int32((uint32(int32(12)) + alignment - uint32(1)) / alignment * alignment)
	goto _2
_2:
	return matrix + uintptr(v1)
}

func Opus_mapping_matrix_init(tls *libc.TLS, matrix uintptr, rows int32, cols int32, gain int32, data uintptr, data_size OpusT_opus_int32) {
	var alignment uint32
	var i1, v1, v3 int32
	var ptr uintptr
	_, _, _, _, _ = alignment, i1, ptr, v1, v3
	_ = data_size
	alignment = uint32(uint64(uintptr(uint32(0)) + 8))
	v1 = int32((uint32(data_size) + alignment - uint32(1)) / alignment * alignment)
	goto _2
_2:
	alignment = uint32(uint64(uintptr(uint32(0)) + 8))
	v3 = int32((uint32(int32(uint64(uint32(rows*cols))*uint64(2))) + alignment - uint32(1)) / alignment * alignment)
	goto _4
_4:
	if !(v1 == v3) {
		Opus_celt_fatal(tls, __ccgo_ts+2234, __ccgo_ts+2312, int32(72))
	}
	(*OpusT_MappingMatrix)(unsafe.Pointer(matrix)).Frows = rows
	(*OpusT_MappingMatrix)(unsafe.Pointer(matrix)).Fcols = cols
	(*OpusT_MappingMatrix)(unsafe.Pointer(matrix)).Fgain = gain
	ptr = Opus_mapping_matrix_get_data(tls, matrix)
	i1 = 0
	for {
		if !(i1 < rows*cols) {
			break
		}
		*(*OpusT_opus_int16)(unsafe.Pointer(ptr + uintptr(i1)*2)) = *(*OpusT_opus_int16)(unsafe.Pointer(data + uintptr(i1)*2))
		goto _5
	_5:
		;
		i1 = i1 + 1
	}
}

func Opus_mapping_matrix_multiply_channel_in_float(tls *libc.TLS, matrix uintptr, input uintptr, input_rows int32, output uintptr, output_row int32, output_rows int32, frame_size int32) {
	var col, i int32
	var matrix_data uintptr
	var tmp float32
	_, _, _, _ = col, i, matrix_data, tmp
	if !(input_rows <= (*OpusT_MappingMatrix)(unsafe.Pointer(matrix)).Fcols && output_rows <= (*OpusT_MappingMatrix)(unsafe.Pointer(matrix)).Frows) {
		Opus_celt_fatal(tls, __ccgo_ts+2336, __ccgo_ts+2312, int32(98))
	}
	matrix_data = Opus_mapping_matrix_get_data(tls, matrix)
	i = 0
	for {
		if !(i < frame_size) {
			break
		}
		tmp = float32(0)
		col = 0
		for {
			if !(col < input_rows) {
				break
			}
			tmp = tmp + float32(float32(*(*OpusT_opus_int16)(unsafe.Pointer(matrix_data + uintptr((*OpusT_MappingMatrix)(unsafe.Pointer(matrix)).Frows*col+output_row)*2)))**(*float32)(unsafe.Pointer(input + uintptr(input_rows*i+col)*4)))
			goto _2
		_2:
			;
			col = col + 1
		}
		*(*OpusT_opus_res)(unsafe.Pointer(output + uintptr(output_rows*i)*4)) = float32(float32(1) / float32(32768) * tmp)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func Opus_mapping_matrix_multiply_channel_out_float(tls *libc.TLS, matrix uintptr, input uintptr, input_row int32, input_rows int32, output uintptr, output_rows int32, frame_size int32) {
	var i, row int32
	var input_sample, tmp float32
	var matrix_data uintptr
	_, _, _, _, _ = i, input_sample, matrix_data, row, tmp
	if !(input_rows <= (*OpusT_MappingMatrix)(unsafe.Pointer(matrix)).Fcols && output_rows <= (*OpusT_MappingMatrix)(unsafe.Pointer(matrix)).Frows) {
		Opus_celt_fatal(tls, __ccgo_ts+2336, __ccgo_ts+2312, int32(130))
	}
	matrix_data = Opus_mapping_matrix_get_data(tls, matrix)
	i = 0
	for {
		if !(i < frame_size) {
			break
		}
		input_sample = *(*OpusT_opus_res)(unsafe.Pointer(input + uintptr(input_rows*i)*4))
		row = 0
		for {
			if !(row < output_rows) {
				break
			}
			tmp = float32(float32(float32(1)/float32(32768)*float32(*(*OpusT_opus_int16)(unsafe.Pointer(matrix_data + uintptr((*OpusT_MappingMatrix)(unsafe.Pointer(matrix)).Frows*input_row+row)*2)))) * input_sample)
			*(*float32)(unsafe.Pointer(output + uintptr(output_rows*i+row)*4)) += tmp
			goto _2
		_2:
			;
			row = row + 1
		}
		goto _1
	_1:
		;
		i = i + 1
	}
}

func Opus_mapping_matrix_multiply_channel_in_short(tls *libc.TLS, matrix uintptr, input uintptr, input_rows int32, output uintptr, output_row int32, output_rows int32, frame_size int32) {
	var col, i int32
	var matrix_data uintptr
	var tmp OpusT_opus_val32
	_, _, _, _ = col, i, matrix_data, tmp
	if !(input_rows <= (*OpusT_MappingMatrix)(unsafe.Pointer(matrix)).Fcols && output_rows <= (*OpusT_MappingMatrix)(unsafe.Pointer(matrix)).Frows) {
		Opus_celt_fatal(tls, __ccgo_ts+2336, __ccgo_ts+2312, int32(161))
	}
	matrix_data = Opus_mapping_matrix_get_data(tls, matrix)
	i = 0
	for {
		if !(i < frame_size) {
			break
		}
		tmp = float32(0)
		col = 0
		for {
			if !(col < input_rows) {
				break
			}
			tmp = tmp + OpusT_opus_val32(int32(*(*OpusT_opus_int16)(unsafe.Pointer(matrix_data + uintptr((*OpusT_MappingMatrix)(unsafe.Pointer(matrix)).Frows*col+output_row)*2)))*int32(*(*OpusT_opus_int16)(unsafe.Pointer(input + uintptr(input_rows*i+col)*2))))
			goto _2
		_2:
			;
			col = col + 1
		}
		*(*OpusT_opus_res)(unsafe.Pointer(output + uintptr(output_rows*i)*4)) = OpusT_opus_res(float32(1) / float32(float32(32768)*float32(32768)) * tmp)
		goto _1
	_1:
		;
		i = i + 1
	}
}

func Opus_mapping_matrix_multiply_channel_out_short(tls *libc.TLS, matrix uintptr, input uintptr, input_row int32, input_rows int32, output uintptr, output_rows int32, frame_size int32) {
	var i, row int32
	var input_sample, tmp OpusT_opus_int32
	var matrix_data, v8 uintptr
	var v2, v3, v4 float32
	var v5 OpusT_opus_int16
	_, _, _, _, _, _, _, _, _, _ = i, input_sample, matrix_data, row, tmp, v2, v3, v4, v5, v8
	if !(input_rows <= (*OpusT_MappingMatrix)(unsafe.Pointer(matrix)).Fcols && output_rows <= (*OpusT_MappingMatrix)(unsafe.Pointer(matrix)).Frows) {
		Opus_celt_fatal(tls, __ccgo_ts+2336, __ccgo_ts+2312, int32(206))
	}
	matrix_data = Opus_mapping_matrix_get_data(tls, matrix)
	i = 0
	for {
		if !(i < frame_size) {
			break
		}
		v2 = *(*OpusT_opus_res)(unsafe.Pointer(input + uintptr(input_rows*i)*4))
		v2 = float32(v2 * float32(32768))
		if v2 > float32(-int32(32768)) {
			v3 = v2
		} else {
			v3 = float32(-int32(32768))
		}
		v2 = v3
		if v2 < float32(int32(32767)) {
			v4 = v2
		} else {
			v4 = float32(int32(32767))
		}
		v2 = v4
		v5 = int16(libc.Xlrintf(tls, v2))
		goto _6
	_6:
		input_sample = int32(v5)
		row = 0
		for {
			if !(row < output_rows) {
				break
			}
			tmp = int32(*(*OpusT_opus_int16)(unsafe.Pointer(matrix_data + uintptr((*OpusT_MappingMatrix)(unsafe.Pointer(matrix)).Frows*input_row+row)*2))) * input_sample
			v8 = output + uintptr(output_rows*i+row)*2
			*(*OpusT_opus_int16)(unsafe.Pointer(v8)) = OpusT_opus_int16(int32(*(*OpusT_opus_int16)(unsafe.Pointer(v8))) + (tmp+int32(16384))>>int32(15))
			goto _7
		_7:
			;
			row = row + 1
		}
		goto _1
	_1:
		;
		i = i + 1
	}
}

func Opus_mapping_matrix_multiply_channel_in_int24(tls *libc.TLS, matrix uintptr, input uintptr, input_rows int32, output uintptr, output_row int32, output_rows int32, frame_size int32) {
	var col, i int32
	var matrix_data uintptr
	var tmp OpusT_opus_val64
	_, _, _, _ = col, i, matrix_data, tmp
	if !(input_rows <= (*OpusT_MappingMatrix)(unsafe.Pointer(matrix)).Fcols && output_rows <= (*OpusT_MappingMatrix)(unsafe.Pointer(matrix)).Frows) {
		Opus_celt_fatal(tls, __ccgo_ts+2336, __ccgo_ts+2312, int32(236))
	}
	matrix_data = Opus_mapping_matrix_get_data(tls, matrix)
	i = 0
	for {
		if !(i < frame_size) {
			break
		}
		tmp = float32(0)
		col = 0
		for {
			if !(col < input_rows) {
				break
			}
			tmp = tmp + OpusT_opus_val64(float32(*(*OpusT_opus_int16)(unsafe.Pointer(matrix_data + uintptr((*OpusT_MappingMatrix)(unsafe.Pointer(matrix)).Frows*col+output_row)*2)))*float32(*(*OpusT_opus_int32)(unsafe.Pointer(input + uintptr(input_rows*i+col)*4))))
			goto _2
		_2:
			;
			col = col + 1
		}
		*(*OpusT_opus_res)(unsafe.Pointer(output + uintptr(output_rows*i)*4)) = float32(float32(1) / float32(32768) / float32(256) * float32(float32(1)/float32(32768)*tmp))
		goto _1
	_1:
		;
		i = i + 1
	}
}

func Opus_mapping_matrix_multiply_channel_out_int24(tls *libc.TLS, matrix uintptr, input uintptr, input_row int32, input_rows int32, output uintptr, output_rows int32, frame_size int32) {
	var i, row int32
	var input_sample OpusT_opus_int32
	var matrix_data, v3 uintptr
	var tmp OpusT_opus_int64
	_, _, _, _, _, _ = i, input_sample, matrix_data, row, tmp, v3
	if !(input_rows <= (*OpusT_MappingMatrix)(unsafe.Pointer(matrix)).Fcols && output_rows <= (*OpusT_MappingMatrix)(unsafe.Pointer(matrix)).Frows) {
		Opus_celt_fatal(tls, __ccgo_ts+2336, __ccgo_ts+2312, int32(271))
	}
	matrix_data = Opus_mapping_matrix_get_data(tls, matrix)
	i = 0
	for {
		if !(i < frame_size) {
			break
		}
		input_sample = int32(libc.Xlrintf(tls, float32(float32(float32(32768)*float32(256))**(*OpusT_opus_res)(unsafe.Pointer(input + uintptr(input_rows*i)*4)))))
		row = 0
		for {
			if !(row < output_rows) {
				break
			}
			tmp = int64(*(*OpusT_opus_int16)(unsafe.Pointer(matrix_data + uintptr((*OpusT_MappingMatrix)(unsafe.Pointer(matrix)).Frows*input_row+row)*2))) * int64(input_sample)
			v3 = output + uintptr(output_rows*i+row)*4
			*(*OpusT_opus_int32)(unsafe.Pointer(v3)) = OpusT_opus_int32(int64(*(*OpusT_opus_int32)(unsafe.Pointer(v3))) + (tmp+int64(16384))>>int32(15))
			goto _2
		_2:
			;
			row = row + 1
		}
		goto _1
	_1:
		;
		i = i + 1
	}
}

const CELT_SIG_SCALE2 = "32768.f"

var log2_x_norm_coeff2 = [8]float32{
	0: float32(1),
	1: float32(0.8888888955116272),
	2: float32(0.8),
	3: float32(0.7272727489471436),
	4: float32(0.6666666865348816),
	5: float32(0.6153846383094788),
	6: float32(0.5714285969734192),
	7: float32(0.5333333611488342),
}
var log2_y_norm_coeff2 = [8]float32{
	1: float32(0.1699250042438507),
	2: float32(0.32192808389663696),
	3: float32(0.45943161845207214),
	4: float32(0.5849624872207642),
	5: float32(0.7004396915435791),
	6: float32(0.8073549270629883),
	7: float32(0.9068905711174011),
}

var trim_icdf5 = [11]uint8{
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
var spread_icdf5 = [4]uint8{
	0: uint8(25),
	1: uint8(23),
	2: uint8(2),
}
var tapset_icdf5 = [3]uint8{
	0: uint8(2),
	1: uint8(1),
}

type OpusT_OpusProjectionDecoder = struct {
	Fdemixing_matrix_size_in_bytes OpusT_opus_int32
}

type OpusProjectionDecoder = struct {
	Fdemixing_matrix_size_in_bytes OpusT_opus_int32
}

func opus_projection_copy_channel_out_float(tls *libc.TLS, dst uintptr, dst_stride int32, dst_channel int32, src uintptr, src_stride int32, frame_size int32, user_data uintptr) {
	var float_dst, matrix uintptr
	_, _ = float_dst, matrix
	float_dst = dst
	matrix = user_data
	if dst_channel == 0 {
		libc.Xmemset(tls, float_dst, 0, uint64(uint32(frame_size*dst_stride))*uint64(4))
	}
	if src != uintptr(uint32(0)) {
		Opus_mapping_matrix_multiply_channel_out_float(tls, matrix, src, dst_channel, src_stride, float_dst, dst_stride, frame_size)
	}
}

func opus_projection_copy_channel_out_short(tls *libc.TLS, dst uintptr, dst_stride int32, dst_channel int32, src uintptr, src_stride int32, frame_size int32, user_data uintptr) {
	var matrix, short_dst uintptr
	_, _ = matrix, short_dst
	short_dst = dst
	matrix = user_data
	if dst_channel == 0 {
		libc.Xmemset(tls, short_dst, 0, uint64(uint32(frame_size*dst_stride))*uint64(2))
	}
	if src != uintptr(uint32(0)) {
		Opus_mapping_matrix_multiply_channel_out_short(tls, matrix, src, dst_channel, src_stride, short_dst, dst_stride, frame_size)
	}
}

func opus_projection_copy_channel_out_int24(tls *libc.TLS, dst uintptr, dst_stride int32, dst_channel int32, src uintptr, src_stride int32, frame_size int32, user_data uintptr) {
	var matrix, short_dst uintptr
	_, _ = matrix, short_dst
	short_dst = dst
	matrix = user_data
	if dst_channel == 0 {
		libc.Xmemset(tls, short_dst, 0, uint64(uint32(frame_size*dst_stride))*uint64(4))
	}
	if src != uintptr(uint32(0)) {
		Opus_mapping_matrix_multiply_channel_out_int24(tls, matrix, src, dst_channel, src_stride, short_dst, dst_stride, frame_size)
	}
}

func get_dec_demixing_matrix(tls *libc.TLS, st uintptr) (r uintptr) {
	var alignment uint32
	var v1 int32
	_, _ = alignment, v1
	/* void* cast avoids clang -Wcast-align warning */
	alignment = uint32(uint64(uintptr(uint32(0)) + 8))
	v1 = int32((uint32(int32(4)) + alignment - uint32(1)) / alignment * alignment)
	goto _2
_2:
	return st + uintptr(v1)
}

func get_multistream_decoder(tls *libc.TLS, st uintptr) (r uintptr) {
	var alignment uint32
	var v1 int32
	_, _ = alignment, v1
	/* void* cast avoids clang -Wcast-align warning */
	alignment = uint32(uint64(uintptr(uint32(0)) + 8))
	v1 = int32((uint32(int32(uint64(4)+uint64(uint32((*OpusT_OpusProjectionDecoder)(unsafe.Pointer(st)).Fdemixing_matrix_size_in_bytes)))) + alignment - uint32(1)) / alignment * alignment)
	goto _2
_2:
	return st + uintptr(v1)
}

func Opus_opus_projection_decoder_get_size(tls *libc.TLS, channels int32, streams int32, coupled_streams int32) (r OpusT_opus_int32) {
	var alignment uint32
	var decoder_size, matrix_size OpusT_opus_int32
	var v1 int32
	_, _, _, _ = alignment, decoder_size, matrix_size, v1
	matrix_size = Opus_mapping_matrix_get_size(tls, streams+coupled_streams, channels)
	if !(matrix_size != 0) {
		return 0
	}
	decoder_size = Opus_opus_multistream_decoder_get_size(tls, streams, coupled_streams)
	if !(decoder_size != 0) {
		return 0
	}
	alignment = uint32(uint64(uintptr(uint32(0)) + 8))
	v1 = int32((uint32(int32(4)) + alignment - uint32(1)) / alignment * alignment)
	goto _2
_2:
	return v1 + matrix_size + decoder_size
}

func Opus_opus_projection_decoder_init(tls *libc.TLS, st1 uintptr, Fs OpusT_opus_int32, channels int32, streams int32, coupled_streams int32, demixing_matrix uintptr, demixing_matrix_size OpusT_opus_int32) (r int32) {
	bp := tls.Alloc(256)
	defer tls.Free(256)
	var _saved_stack, buf, st, v1, v10, v11, v13, v15, v17, v19, v21, v3, v5, v6, v8 uintptr
	var expected_matrix_size OpusT_opus_int32
	var i, nb_input_streams, ret, s int32
	var _ /* mapping at bp+0 */ [255]uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = _saved_stack, buf, expected_matrix_size, i, nb_input_streams, ret, s, st, v1, v10, v11, v13, v15, v17, v19, v21, v3, v5, v6, v8
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
	/* Verify supplied matrix size. */
	nb_input_streams = streams + coupled_streams
	expected_matrix_size = int32(uint64(uint32(nb_input_streams*channels)) * uint64(2))
	if expected_matrix_size != demixing_matrix_size {
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			goto _28
		_28:
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v3 = st
		goto _30
	_30:
		(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
		return -int32(1)
	}
	/* Convert demixing matrix input into internal format. */
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _32
	_32:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _34
_34:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v5 = libc.Xmalloc(tls, uint64(16))
		goto _36
	_36:
		st = v5
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v6 = st
	goto _38
_38:
	*(*uintptr)(unsafe.Pointer(v3 + 8)) += uintptr((uint64(uint32(2)) - uint64(int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v6)).Fglobal_stack))) & (uint64(uint32(2)) - uint64(uint32(1))))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v8 = libc.Xmalloc(tls, uint64(16))
		goto _40
	_40:
		st = v8
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v10 = st
	goto _42
_42:
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v11 = libc.Xmalloc(tls, uint64(16))
		goto _44
	_44:
		st = v11
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v13 = st
	goto _46
_46:
	if !(int64(int32(uint64(uint32(nb_input_streams*channels))*(uint64(2)/uint64(1)))) <= int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v10)).Fscratch_ptr+uintptr(GLOBAL_STACK_SIZE))-int64((*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v13)).Fglobal_stack)) {
		Opus_celt_fatal(tls, __ccgo_ts+996, __ccgo_ts+2412, int32(167))
	}
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v15 = libc.Xmalloc(tls, uint64(16))
		goto _48
	_48:
		st = v15
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v17 = st
	goto _50
_50:
	*(*uintptr)(unsafe.Pointer(v17 + 8)) += uintptr(uint64(uint32(nb_input_streams*channels)) * (uint64(2) / uint64(1)))
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v19 = libc.Xmalloc(tls, uint64(16))
		goto _52
	_52:
		st = v19
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v21 = st
	goto _54
_54:
	buf = (*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v21)).Fglobal_stack - uintptr(uint64(uint32(nb_input_streams*channels))*(uint64(2)/uint64(1)))
	i = 0
	for {
		if !(i < nb_input_streams*channels) {
			break
		}
		s = int32(*(*uint8)(unsafe.Pointer(demixing_matrix + uintptr(int32(2)*i+int32(1)))))<<int32(8) | int32(*(*uint8)(unsafe.Pointer(demixing_matrix + uintptr(int32(2)*i))))
		s = s&int32(0xFFFF) ^ int32(0x8000) - int32(0x8000)
		*(*OpusT_opus_int16)(unsafe.Pointer(buf + uintptr(i)*2)) = int16(s)
		goto _55
	_55:
		;
		i = i + 1
	}
	/* Assign demixing matrix. */
	(*OpusT_OpusProjectionDecoder)(unsafe.Pointer(st1)).Fdemixing_matrix_size_in_bytes = Opus_mapping_matrix_get_size(tls, channels, nb_input_streams)
	if !((*OpusT_OpusProjectionDecoder)(unsafe.Pointer(st1)).Fdemixing_matrix_size_in_bytes != 0) {
		st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
		if !(st != 0) {
			v1 = libc.Xmalloc(tls, uint64(16))
			goto _57
		_57:
			st = v1
			if st != 0 {
				libc.Xmemset(tls, st, 0, uint64(16))
			}
			libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
		}
		v3 = st
		goto _59
	_59:
		(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
		return -int32(1)
	}
	Opus_mapping_matrix_init(tls, get_dec_demixing_matrix(tls, st1), channels, nb_input_streams, 0, buf, demixing_matrix_size)
	/* Set trivial mapping so each input channel pairs with a matrix column. */
	i = 0
	for {
		if !(i < channels) {
			break
		}
		(*(*[255]uint8)(unsafe.Pointer(bp)))[i] = uint8(i)
		goto _60
	_60:
		;
		i = i + 1
	}
	ret = Opus_opus_multistream_decoder_init(tls, get_multistream_decoder(tls, st1), Fs, channels, streams, coupled_streams, bp)
	st = libc.Xpthread_getspecific(tls, uint32(0x6f707573))
	if !(st != 0) {
		v1 = libc.Xmalloc(tls, uint64(16))
		goto _62
	_62:
		st = v1
		if st != 0 {
			libc.Xmemset(tls, st, 0, uint64(16))
		}
		libc.Xpthread_setspecific(tls, uint32(0x6f707573), st)
	}
	v3 = st
	goto _64
_64:
	(*OpusT_opus_ccgo_pseudostack_state)(unsafe.Pointer(v3)).Fglobal_stack = _saved_stack
	return ret
}

func Opus_opus_projection_decoder_create(tls *libc.TLS, Fs OpusT_opus_int32, channels int32, streams int32, coupled_streams int32, demixing_matrix uintptr, demixing_matrix_size OpusT_opus_int32, error1 uintptr) (r uintptr) {
	var ret, size1 int32
	var st, v1 uintptr
	_, _, _, _ = ret, size1, st, v1
	/* Allocate space for the projection decoder. */
	size1 = Opus_opus_projection_decoder_get_size(tls, channels, streams, coupled_streams)
	if !(size1 != 0) {
		if error1 != 0 {
			*(*int32)(unsafe.Pointer(error1)) = -int32(7)
		}
		return uintptr(uint32(0))
	}
	v1 = libc.Xmalloc(tls, uint64(uint32(size1)))
	goto _2
_2:
	st = v1
	if !(st != 0) {
		if error1 != 0 {
			*(*int32)(unsafe.Pointer(error1)) = -int32(7)
		}
		return uintptr(uint32(0))
	}
	/* Initialize projection decoder with provided settings. */
	ret = Opus_opus_projection_decoder_init(tls, st, Fs, channels, streams, coupled_streams, demixing_matrix, demixing_matrix_size)
	if ret != OPUS_OK {
		libc.Xfree(tls, st)
		st = uintptr(uint32(0))
	}
	if error1 != 0 {
		*(*int32)(unsafe.Pointer(error1)) = ret
	}
	return st
}

func Opus_opus_projection_decode(tls *libc.TLS, st uintptr, data uintptr, len1 OpusT_opus_int32, pcm uintptr, frame_size int32, decode_fec int32) (r int32) {
	return Opus_opus_multistream_decode_native(tls, get_multistream_decoder(tls, st), data, len1, pcm, __ccgo_fp(opus_projection_copy_channel_out_short), frame_size, decode_fec, int32(OPTIONAL_CLIP), get_dec_demixing_matrix(tls, st))
}

func Opus_opus_projection_decode24(tls *libc.TLS, st uintptr, data uintptr, len1 OpusT_opus_int32, pcm uintptr, frame_size int32, decode_fec int32) (r int32) {
	return Opus_opus_multistream_decode_native(tls, get_multistream_decoder(tls, st), data, len1, pcm, __ccgo_fp(opus_projection_copy_channel_out_int24), frame_size, decode_fec, 0, get_dec_demixing_matrix(tls, st))
}

func Opus_opus_projection_decode_float(tls *libc.TLS, st uintptr, data uintptr, len1 OpusT_opus_int32, pcm uintptr, frame_size int32, decode_fec int32) (r int32) {
	return Opus_opus_multistream_decode_native(tls, get_multistream_decoder(tls, st), data, len1, pcm, __ccgo_fp(opus_projection_copy_channel_out_float), frame_size, decode_fec, 0, get_dec_demixing_matrix(tls, st))
}

func Opus_opus_projection_decoder_ctl(tls *libc.TLS, st uintptr, request int32, va uintptr) (r int32) {
	var ap OpusT_va_list
	var ret int32
	_, _ = ap, ret
	ret = OPUS_OK
	ap = va
	ret = Opus_opus_multistream_decoder_ctl_va_list(tls, get_multistream_decoder(tls, st), request, ap)
	_ = ap
	return ret
}

func Opus_opus_projection_decoder_destroy(tls *libc.TLS, st uintptr) {
	libc.Xfree(tls, st)
}

var trim_icdf6 = [11]uint8{
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
var spread_icdf6 = [4]uint8{
	0: uint8(25),
	1: uint8(23),
	2: uint8(2),
}
var tapset_icdf6 = [3]uint8{
	0: uint8(2),
	1: uint8(1),
}

// C documentation
//
//	/* Given an extension payload (i.e., excluding the initial ID byte), advance
//	    data to the next extension and return the length of the remaining
//	    extensions.
//	   N.B., a "Repeat These Extensions" extension (ID==2) does not advance past
//	    the repeated extension payloads.
//	   That requires higher-level logic. */
func skip_extension_payload(tls *libc.TLS, pdata uintptr, len1 OpusT_opus_int32, pheader_size uintptr, id_byte int32, trailing_short_len OpusT_opus_int32) (r OpusT_opus_int32) {
	var L, id int32
	var bytes, header_size, lacing OpusT_opus_int32
	var data, v1 uintptr
	_, _, _, _, _, _, _ = L, bytes, data, header_size, id, lacing, v1
	data = *(*uintptr)(unsafe.Pointer(pdata))
	header_size = 0
	id = id_byte >> int32(1)
	L = id_byte & int32(1)
	if id == 0 && L == int32(1) || id == int32(2) {
		/* Nothing to do. */
	} else {
		if id > 0 && id < int32(32) {
			if len1 < L {
				return -int32(1)
			}
			data = data + uintptr(L)
			len1 = len1 - L
		} else {
			if L == 0 {
				if len1 < trailing_short_len {
					return -int32(1)
				}
				data = data + uintptr(len1-trailing_short_len)
				len1 = trailing_short_len
			} else {
				bytes = 0
				for cond := true; cond; cond = lacing == int32(255) {
					if len1 < int32(1) {
						return -int32(1)
					}
					v1 = data
					data = data + 1
					lacing = int32(*(*uint8)(unsafe.Pointer(v1)))
					bytes = bytes + lacing
					header_size = header_size + 1
					len1 = len1 - (lacing + int32(1))
				}
				if len1 < 0 {
					return -int32(1)
				}
				data = data + uintptr(bytes)
			}
		}
	}
	*(*uintptr)(unsafe.Pointer(pdata)) = data
	*(*OpusT_opus_int32)(unsafe.Pointer(pheader_size)) = header_size
	return len1
}

// C documentation
//
//	/* Given an extension, advance data to the next extension and return the
//	   length of the remaining extensions.
//	   N.B., a "Repeat These Extensions" extension (ID==2) only advances past the
//	    extension ID byte.
//	   Higher-level logic is required to skip the extension payloads that come
//	    after it.*/
func skip_extension(tls *libc.TLS, pdata uintptr, len1 OpusT_opus_int32, pheader_size uintptr) (r OpusT_opus_int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var id_byte int32
	var v1 uintptr
	var _ /* data at bp+0 */ uintptr
	_, _ = id_byte, v1
	if len1 == 0 {
		*(*OpusT_opus_int32)(unsafe.Pointer(pheader_size)) = 0
		return 0
	}
	if len1 < int32(1) {
		return -int32(1)
	}
	*(*uintptr)(unsafe.Pointer(bp)) = *(*uintptr)(unsafe.Pointer(pdata))
	v1 = *(*uintptr)(unsafe.Pointer(bp))
	*(*uintptr)(unsafe.Pointer(bp)) = *(*uintptr)(unsafe.Pointer(bp)) + 1
	id_byte = int32(*(*uint8)(unsafe.Pointer(v1)))
	len1 = len1 - 1
	len1 = skip_extension_payload(tls, bp, len1, pheader_size, id_byte, 0)
	if len1 >= 0 {
		*(*uintptr)(unsafe.Pointer(pdata)) = *(*uintptr)(unsafe.Pointer(bp))
		*(*OpusT_opus_int32)(unsafe.Pointer(pheader_size)) = *(*OpusT_opus_int32)(unsafe.Pointer(pheader_size)) + 1
	}
	return len1
}

func Opus_opus_extension_iterator_init(tls *libc.TLS, iter uintptr, data uintptr, len1 OpusT_opus_int32, nb_frames OpusT_opus_int32) {
	var v1, v2 uintptr
	var v4 OpusT_opus_int32
	var v6 int32
	_, _, _, _ = v1, v2, v4, v6
	if !(len1 >= int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+2445, __ccgo_ts+2472, int32(122))
	}
	if !(data != uintptr(uint32(0)) || len1 == 0) {
		Opus_celt_fatal(tls, __ccgo_ts+2492, __ccgo_ts+2472, int32(123))
	}
	if !(nb_frames >= 0 && nb_frames <= int32(48)) {
		Opus_celt_fatal(tls, __ccgo_ts+2535, __ccgo_ts+2472, int32(124))
	}
	v2 = data
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fdata = v2
	v1 = v2
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_data = v1
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_data = v1
	v1 = uintptr(uint32(0))
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fsrc_data = v1
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Flast_long = v1
	v4 = len1
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Flen1 = v4
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_len = v4
	v4 = int32(0)
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fsrc_len = v4
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_len = v4
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Ftrailing_short_len = 0
	v6 = nb_frames
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fnb_frames = v6
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fframe_max = v6
	v6 = int32(0)
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_frame = v6
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_frame = v6
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_l = uint8(0)
}

// C documentation
//
//	/* Reset the iterator so it can start iterating again from the first
//	    extension. */
func Opus_opus_extension_iterator_reset(tls *libc.TLS, iter uintptr) {
	var v1 uintptr
	var v2 int32
	_, _ = v1, v2
	v1 = (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fdata
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_data = v1
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_data = v1
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Flast_long = uintptr(uint32(0))
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_len = (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Flen1
	v2 = int32(0)
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_frame = v2
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_frame = v2
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Ftrailing_short_len = 0
}

// C documentation
//
//	/* Tell the iterator not to return any extensions for frames of index
//	    frame_max or larger.
//	   This can allow it to stop iterating early if these extensions are not
//	    needed. */
func Opus_opus_extension_iterator_set_frame_max(tls *libc.TLS, iter uintptr, frame_max int32) {
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fframe_max = frame_max
}

// C documentation
//
//	/* Return the next repeated extension.
//	   The return value is non-zero if one is found, negative on error, or 0 if we
//	    have finished repeating extensions. */
func opus_extension_iterator_next_repeat(tls *libc.TLS, iter uintptr, ext uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var curr_data0 uintptr
	var repeat_id_byte int32
	var _ /* header_size at bp+0 */ OpusT_opus_int32
	_, _ = curr_data0, repeat_id_byte
	if !((*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_frame > int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+2587, __ccgo_ts+2472, int32(160))
	}
	for {
		if !((*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_frame < (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fnb_frames) {
			break
		}
		for (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fsrc_len > 0 {
			repeat_id_byte = int32(*(*uint8)(unsafe.Pointer((*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fsrc_data)))
			(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fsrc_len = skip_extension(tls, iter+32, (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fsrc_len, bp)
			/* We skipped this extension earlier, so it should not fail now. */
			if !((*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fsrc_len >= int32(0)) {
				Opus_celt_fatal(tls, __ccgo_ts+2628, __ccgo_ts+2472, int32(169))
			}
			/* Don't repeat padding or frame separators with a 0 increment. */
			if repeat_id_byte <= int32(3) {
				continue
			}
			/* If the "Repeat These Extensions" extension had L == 0 and this
			   is the last repeated long extension, then force decoding the
			   payload with L = 0. */
			if int32((*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_l) == 0 && (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_frame+int32(1) >= (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fnb_frames && (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fsrc_data == (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Flast_long {
				repeat_id_byte = repeat_id_byte & ^int32(1)
			}
			curr_data0 = (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_data
			(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_len = skip_extension_payload(tls, iter+8, (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_len, bp, repeat_id_byte, (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Ftrailing_short_len)
			if (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_len < 0 {
				return -int32(4)
			}
			if !(int64((*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_data)-int64((*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fdata) == int64((*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Flen1-(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_len)) {
				Opus_celt_fatal(tls, __ccgo_ts+2665, __ccgo_ts+2472, int32(187))
			}
			/* If we were asked to stop at frame_max, skip extensions for later
			   frames. */
			if (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_frame >= (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fframe_max {
				continue
			}
			if ext != uintptr(uint32(0)) {
				(*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Fid = repeat_id_byte >> int32(1)
				(*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Fframe = (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_frame
				(*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Fdata = curr_data0 + uintptr(*(*OpusT_opus_int32)(unsafe.Pointer(bp)))
				(*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Flen1 = int32(int64((*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_data) - int64(curr_data0) - int64(*(*OpusT_opus_int32)(unsafe.Pointer(bp))))
			}
			return int32(1)
		}
		/* We finished repeating the extensions for this frame. */
		(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fsrc_data = (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_data
		(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fsrc_len = (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_len
		goto _1
	_1:
		;
		(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_frame = (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_frame + 1
	}
	/* We finished repeating extensions. */
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_data = (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_data
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Flast_long = uintptr(uint32(0))
	/* If L == 0, advance the frame number to handle the case where we did
	   not consume all of the data with an L == 0 long extension. */
	if int32((*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_l) == 0 {
		(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_frame = (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_frame + 1
		/* Ignore additional padding if this was already the last frame. */
		if (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_frame >= (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fnb_frames {
			(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_len = 0
		}
	}
	(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_frame = 0
	return 0
}

// C documentation
//
//	/* Return the next extension (excluding real padding, separators, and repeat
//	    indicators, but including the repeated extensions) in bitstream order.
//	   Due to the extension repetition mechanism, extensions are not necessarily
//	    returned in frame order. */
func Opus_opus_extension_iterator_next(tls *libc.TLS, iter uintptr, ext uintptr) (r int32) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var L, id, ret, ret1 int32
	var curr_data0 uintptr
	var _ /* header_size at bp+0 */ OpusT_opus_int32
	_, _, _, _, _ = L, curr_data0, id, ret, ret1
	if (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_len < 0 {
		return -int32(4)
	}
	if (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_frame > 0 {
		/* We are in the process of repeating some extensions. */
		ret = opus_extension_iterator_next_repeat(tls, iter, ext)
		if ret != 0 {
			return ret
		}
	}
	/* Checking this here allows opus_extension_iterator_set_frame_max() to be
	   called at any point. */
	if (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_frame >= (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fframe_max {
		return 0
	}
	for (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_len > 0 {
		curr_data0 = (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_data
		id = int32(*(*uint8)(unsafe.Pointer(curr_data0))) >> int32(1)
		L = int32(*(*uint8)(unsafe.Pointer(curr_data0))) & int32(1)
		(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_len = skip_extension(tls, iter+8, (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_len, bp)
		if (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_len < 0 {
			return -int32(4)
		}
		if !(int64((*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_data)-int64((*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fdata) == int64((*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Flen1-(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_len)) {
			Opus_celt_fatal(tls, __ccgo_ts+2665, __ccgo_ts+2472, int32(255))
		}
		if id == int32(1) {
			if L == 0 {
				(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_frame = (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_frame + 1
			} else {
				/* A frame increment of 0 is a no-op. */
				if !(*(*uint8)(unsafe.Pointer(curr_data0 + 1)) != 0) {
					continue
				}
				*(*int32)(unsafe.Pointer(iter + 68)) += int32(*(*uint8)(unsafe.Pointer(curr_data0 + 1)))
			}
			if (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_frame >= (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fnb_frames {
				(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_len = -int32(1)
				return -int32(4)
			}
			/* If we were asked to stop at frame_max, skip extensions for later
			   frames. */
			if (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_frame >= (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fframe_max {
				(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_len = 0
			}
			(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_data = (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_data
			(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Flast_long = uintptr(uint32(0))
			(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Ftrailing_short_len = 0
		} else {
			if id == int32(2) {
				(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_l = uint8(L)
				(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_frame = (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_frame + int32(1)
				(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_len = int32(int64(curr_data0) - int64((*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_data))
				(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fsrc_data = (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_data
				(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fsrc_len = (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Frepeat_len
				ret1 = opus_extension_iterator_next_repeat(tls, iter, ext)
				if ret1 != 0 {
					return ret1
				}
			} else {
				if id > int32(2) {
					/* Update the location of the last long extension.
					   This lets us know when we need to modify the last L flag if we
					    repeat these extensions with L=0. */
					if id >= int32(32) {
						(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Flast_long = (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_data
						(*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Ftrailing_short_len = 0
					} else {
						*(*OpusT_opus_int32)(unsafe.Pointer(iter + 56)) += L
					}
					if ext != uintptr(uint32(0)) {
						(*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Fid = id
						(*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Fframe = (*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_frame
						(*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Fdata = curr_data0 + uintptr(*(*OpusT_opus_int32)(unsafe.Pointer(bp)))
						(*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Flen1 = int32(int64((*OpusT_OpusExtensionIterator)(unsafe.Pointer(iter)).Fcurr_data) - int64(curr_data0) - int64(*(*OpusT_opus_int32)(unsafe.Pointer(bp))))
					}
					return int32(1)
				}
			}
		}
	}
	return 0
}

func Opus_opus_extension_iterator_find(tls *libc.TLS, iter uintptr, ext uintptr, id int32) (r int32) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var ret int32
	var _ /* curr_ext at bp+0 */ OpusT_opus_extension_data
	_ = ret
	for {
		ret = Opus_opus_extension_iterator_next(tls, iter, bp)
		if ret <= 0 {
			return ret
		}
		if (*(*OpusT_opus_extension_data)(unsafe.Pointer(bp))).Fid == id {
			*(*OpusT_opus_extension_data)(unsafe.Pointer(ext)) = *(*OpusT_opus_extension_data)(unsafe.Pointer(bp))
			return ret
		}
		goto _1
	_1:
	}
	return r
}

// C documentation
//
//	/* Count the number of extensions, excluding real padding, separators, and
//	    repeat indicators, but including the repeated extensions. */
func Opus_opus_packet_extensions_count(tls *libc.TLS, data uintptr, len1 OpusT_opus_int32, nb_frames int32) (r OpusT_opus_int32) {
	bp := tls.Alloc(80)
	defer tls.Free(80)
	var count int32
	var _ /* iter at bp+0 */ OpusT_OpusExtensionIterator
	_ = count
	Opus_opus_extension_iterator_init(tls, bp, data, len1, nb_frames)
	count = 0
	for {
		if !(Opus_opus_extension_iterator_next(tls, bp, uintptr(uint32(0))) > 0) {
			break
		}
		goto _1
	_1:
		;
		count = count + 1
	}
	return count
}

// C documentation
//
//	/* Count the number of extensions for each frame, excluding real padding and
//	    separators and repeat indicators, but including the repeated extensions. */
func Opus_opus_packet_extensions_count_ext(tls *libc.TLS, data uintptr, len1 OpusT_opus_int32, nb_frame_exts uintptr, nb_frames int32) (r OpusT_opus_int32) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var count int32
	var _ /* ext at bp+80 */ OpusT_opus_extension_data
	var _ /* iter at bp+0 */ OpusT_OpusExtensionIterator
	_ = count
	Opus_opus_extension_iterator_init(tls, bp, data, len1, nb_frames)
	libc.Xmemset(tls, nb_frame_exts, 0, uint64(uint32(nb_frames))*uint64(4))
	count = 0
	for {
		if !(Opus_opus_extension_iterator_next(tls, bp, bp+80) > 0) {
			break
		}
		*(*OpusT_opus_int32)(unsafe.Pointer(nb_frame_exts + uintptr((*(*OpusT_opus_extension_data)(unsafe.Pointer(bp + 80))).Fframe)*4)) = *(*OpusT_opus_int32)(unsafe.Pointer(nb_frame_exts + uintptr((*(*OpusT_opus_extension_data)(unsafe.Pointer(bp + 80))).Fframe)*4)) + 1
		goto _1
	_1:
		;
		count = count + 1
	}
	return count
}

// C documentation
//
//	/* Extract extensions from Opus padding (excluding real padding, separators,
//	    and repeat indicators, but including the repeated extensions) in bitstream
//	    order.
//	   Due to the extension repetition mechanism, extensions are not necessarily
//	    returned in frame order. */
func Opus_opus_packet_extensions_parse(tls *libc.TLS, data uintptr, len1 OpusT_opus_int32, extensions uintptr, nb_extensions uintptr, nb_frames int32) (r OpusT_opus_int32) {
	bp := tls.Alloc(112)
	defer tls.Free(112)
	var count, ret int32
	var _ /* ext at bp+80 */ OpusT_opus_extension_data
	var _ /* iter at bp+0 */ OpusT_OpusExtensionIterator
	_, _ = count, ret
	if !(nb_extensions != uintptr(uint32(0))) {
		Opus_celt_fatal(tls, __ccgo_ts+2742, __ccgo_ts+2472, int32(365))
	}
	if !(extensions != uintptr(uint32(0)) || *(*OpusT_opus_int32)(unsafe.Pointer(nb_extensions)) == 0) {
		Opus_celt_fatal(tls, __ccgo_ts+2782, __ccgo_ts+2472, int32(366))
	}
	Opus_opus_extension_iterator_init(tls, bp, data, len1, nb_frames)
	count = 0
	for {
		ret = Opus_opus_extension_iterator_next(tls, bp, bp+80)
		if ret <= 0 {
			break
		}
		if count == *(*OpusT_opus_int32)(unsafe.Pointer(nb_extensions)) {
			return -int32(2)
		}
		*(*OpusT_opus_extension_data)(unsafe.Pointer(extensions + uintptr(count)*24)) = *(*OpusT_opus_extension_data)(unsafe.Pointer(bp + 80))
		goto _1
	_1:
		;
		count = count + 1
	}
	*(*OpusT_opus_int32)(unsafe.Pointer(nb_extensions)) = count
	return ret
}

// C documentation
//
//	/* Extract extensions from Opus padding (excluding real padding, separators,
//	    and repeat indicators, but including the repeated extensions) in frame
//	    order.
//	   nb_frame_exts must be filled with the output of
//	    opus_packet_extensions_count_ext(). */
func Opus_opus_packet_extensions_parse_ext(tls *libc.TLS, data uintptr, len1 OpusT_opus_int32, extensions uintptr, nb_extensions uintptr, nb_frame_exts uintptr, nb_frames int32) (r OpusT_opus_int32) {
	bp := tls.Alloc(304)
	defer tls.Free(304)
	var count, prev_total, ret, total int32
	var idx, v3 OpusT_opus_int32
	var v4 uintptr
	var _ /* ext at bp+80 */ OpusT_opus_extension_data
	var _ /* iter at bp+0 */ OpusT_OpusExtensionIterator
	var _ /* nb_frames_cum at bp+104 */ [49]OpusT_opus_int32
	_, _, _, _, _, _, _ = count, idx, prev_total, ret, total, v3, v4
	if !(nb_extensions != uintptr(uint32(0))) {
		Opus_celt_fatal(tls, __ccgo_ts+2742, __ccgo_ts+2472, int32(395))
	}
	if !(extensions != uintptr(uint32(0)) || *(*OpusT_opus_int32)(unsafe.Pointer(nb_extensions)) == 0) {
		Opus_celt_fatal(tls, __ccgo_ts+2782, __ccgo_ts+2472, int32(396))
	}
	if !(nb_frames <= int32(48)) {
		Opus_celt_fatal(tls, __ccgo_ts+2842, __ccgo_ts+2472, int32(397))
	}
	/* Convert the frame extension count array to a cumulative sum. */
	prev_total = 0
	count = 0
	for {
		if !(count < nb_frames) {
			break
		}
		total = *(*OpusT_opus_int32)(unsafe.Pointer(nb_frame_exts + uintptr(count)*4)) + prev_total
		(*(*[49]OpusT_opus_int32)(unsafe.Pointer(bp + 104)))[count] = prev_total
		prev_total = total
		goto _1
	_1:
		;
		count = count + 1
	}
	(*(*[49]OpusT_opus_int32)(unsafe.Pointer(bp + 104)))[count] = prev_total
	Opus_opus_extension_iterator_init(tls, bp, data, len1, nb_frames)
	count = 0
	for {
		ret = Opus_opus_extension_iterator_next(tls, bp, bp+80)
		if ret <= 0 {
			break
		}
		v4 = bp + 104 + uintptr((*(*OpusT_opus_extension_data)(unsafe.Pointer(bp + 80))).Fframe)*4
		v3 = *(*OpusT_opus_int32)(unsafe.Pointer(v4))
		*(*OpusT_opus_int32)(unsafe.Pointer(v4)) = *(*OpusT_opus_int32)(unsafe.Pointer(v4)) + 1
		idx = v3
		if idx >= *(*OpusT_opus_int32)(unsafe.Pointer(nb_extensions)) {
			return -int32(2)
		}
		if !(idx < (*(*[49]OpusT_opus_int32)(unsafe.Pointer(bp + 104)))[(*(*OpusT_opus_extension_data)(unsafe.Pointer(bp + 80))).Fframe+int32(1)]) {
			Opus_celt_fatal(tls, __ccgo_ts+2876, __ccgo_ts+2472, int32(416))
		}
		*(*OpusT_opus_extension_data)(unsafe.Pointer(extensions + uintptr(idx)*24)) = *(*OpusT_opus_extension_data)(unsafe.Pointer(bp + 80))
		goto _2
	_2:
		;
		count = count + 1
	}
	*(*OpusT_opus_int32)(unsafe.Pointer(nb_extensions)) = count
	return ret
}

func write_extension_payload(tls *libc.TLS, data uintptr, len1 OpusT_opus_int32, pos OpusT_opus_int32, ext uintptr, last int32) (r int32) {
	var j, length_bytes OpusT_opus_int32
	_, _ = j, length_bytes
	if !((*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Fid >= int32(3) && (*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Fid <= int32(127)) {
		Opus_celt_fatal(tls, __ccgo_ts+2929, __ccgo_ts+2472, int32(425))
	}
	if (*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Fid < int32(32) {
		if (*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Flen1 < 0 || (*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Flen1 > int32(1) {
			return -int32(1)
		}
		if (*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Flen1 > 0 {
			if len1-pos < (*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Flen1 {
				return -int32(2)
			}
			if data != 0 {
				*(*uint8)(unsafe.Pointer(data + uintptr(pos))) = *(*uint8)(unsafe.Pointer((*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Fdata))
			}
			pos = pos + 1
		}
	} else {
		if (*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Flen1 < 0 {
			return -int32(1)
		}
		length_bytes = int32(1) + (*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Flen1/int32(255)
		if last != 0 {
			length_bytes = 0
		}
		if len1-pos < length_bytes+(*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Flen1 {
			return -int32(2)
		}
		if !(last != 0) {
			j = 0
			for {
				if !(j < (*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Flen1/int32(255)) {
					break
				}
				if data != 0 {
					*(*uint8)(unsafe.Pointer(data + uintptr(pos))) = uint8(255)
				}
				pos = pos + 1
				goto _1
			_1:
				;
				j = j + 1
			}
			if data != 0 {
				*(*uint8)(unsafe.Pointer(data + uintptr(pos))) = uint8((*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Flen1 % int32(255))
			}
			pos = pos + 1
		}
		if data != 0 {
			libc.Xmemcpy(tls, data+uintptr(pos), (*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Fdata, uint64(uint32((*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Flen1))*uint64(1)+uint64(0*(OpusT___predefined_ptrdiff_t(data+uintptr(pos))-int64((*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Fdata))))
		}
		pos = pos + (*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Flen1
	}
	return pos
}

func write_extension(tls *libc.TLS, data uintptr, len1 OpusT_opus_int32, pos OpusT_opus_int32, ext uintptr, last int32) (r int32) {
	var v1 int32
	_ = v1
	if len1-pos < int32(1) {
		return -int32(2)
	}
	if !((*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Fid >= int32(3) && (*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Fid <= int32(127)) {
		Opus_celt_fatal(tls, __ccgo_ts+2929, __ccgo_ts+2472, int32(465))
	}
	if data != 0 {
		if (*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Fid < int32(32) {
			v1 = (*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Flen1
		} else {
			v1 = libc.BoolInt32(!(last != 0))
		}
		*(*uint8)(unsafe.Pointer(data + uintptr(pos))) = uint8((*OpusT_opus_extension_data)(unsafe.Pointer(ext)).Fid<<int32(1) + v1)
	}
	pos = pos + 1
	return write_extension_payload(tls, data, len1, pos, ext, last)
}

func Opus_opus_packet_extensions_generate(tls *libc.TLS, data uintptr, len1 OpusT_opus_int32, extensions uintptr, nb_extensions OpusT_opus_int32, nb_frames int32, pad int32) (r OpusT_opus_int32) {
	bp := tls.Alloc(192)
	defer tls.Free(192)
	var curr_frame, diff, f, g, g1, j, j1, last, nb_repeated, repeat_count, v3 int32
	var frame_min_idx, frame_repeat_idx [48]OpusT_opus_int32
	var i, last_long_idx, padding, pos, written OpusT_opus_int32
	var _ /* frame_max_idx at bp+0 */ [48]OpusT_opus_int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = curr_frame, diff, f, frame_min_idx, frame_repeat_idx, g, g1, i, j, j1, last, last_long_idx, nb_repeated, padding, pos, repeat_count, written, v3
	curr_frame = 0
	pos = 0
	written = 0
	if !(len1 >= int32(0)) {
		Opus_celt_fatal(tls, __ccgo_ts+2445, __ccgo_ts+2472, int32(484))
	}
	if nb_frames > int32(48) {
		return -int32(1)
	}
	/* Do a little work up-front to make this O(nb_extensions) instead of
	   O(nb_extensions*nb_frames) so long as the extensions are in frame
	   order (without requiring that they be in frame order). */
	f = 0
	for {
		if !(f < nb_frames) {
			break
		}
		frame_min_idx[f] = nb_extensions
		goto _1
	_1:
		;
		f = f + 1
	}
	libc.Xmemset(tls, bp, 0, uint64(uint32(nb_frames))*uint64(4))
	i = 0
	for {
		if !(i < nb_extensions) {
			break
		}
		f = (*(*OpusT_opus_extension_data)(unsafe.Pointer(extensions + uintptr(i)*24))).Fframe
		if f < 0 || f >= nb_frames {
			return -int32(1)
		}
		if (*(*OpusT_opus_extension_data)(unsafe.Pointer(extensions + uintptr(i)*24))).Fid < int32(3) || (*(*OpusT_opus_extension_data)(unsafe.Pointer(extensions + uintptr(i)*24))).Fid > int32(127) {
			return -int32(1)
		}
		if frame_min_idx[f] < i {
			v3 = frame_min_idx[f]
		} else {
			v3 = i
		}
		frame_min_idx[f] = v3
		if (*(*[48]OpusT_opus_int32)(unsafe.Pointer(bp)))[f] > i+int32(1) {
			v3 = (*(*[48]OpusT_opus_int32)(unsafe.Pointer(bp)))[f]
		} else {
			v3 = i + int32(1)
		}
		(*(*[48]OpusT_opus_int32)(unsafe.Pointer(bp)))[f] = v3
		goto _2
	_2:
		;
		i = i + 1
	}
	f = 0
	for {
		if !(f < nb_frames) {
			break
		}
		frame_repeat_idx[f] = frame_min_idx[f]
		goto _5
	_5:
		;
		f = f + 1
	}
	f = 0
	for {
		if !(f < nb_frames) {
			break
		}
		repeat_count = 0
		last_long_idx = -int32(1)
		if f+int32(1) < nb_frames {
			i = frame_min_idx[f]
			for {
				if !(i < (*(*[48]OpusT_opus_int32)(unsafe.Pointer(bp)))[f]) {
					break
				}
				if (*(*OpusT_opus_extension_data)(unsafe.Pointer(extensions + uintptr(i)*24))).Fframe == f {
					/* Test if we can repeat this extension in future frames. */
					g = f + int32(1)
					for {
						if !(g < nb_frames) {
							break
						}
						if frame_repeat_idx[g] >= (*(*[48]OpusT_opus_int32)(unsafe.Pointer(bp)))[g] {
							break
						}
						if !((*(*OpusT_opus_extension_data)(unsafe.Pointer(extensions + uintptr(frame_repeat_idx[g])*24))).Fframe == g) {
							Opus_celt_fatal(tls, __ccgo_ts+2978, __ccgo_ts+2472, int32(518))
						}
						if (*(*OpusT_opus_extension_data)(unsafe.Pointer(extensions + uintptr(frame_repeat_idx[g])*24))).Fid != (*(*OpusT_opus_extension_data)(unsafe.Pointer(extensions + uintptr(i)*24))).Fid {
							break
						}
						if (*(*OpusT_opus_extension_data)(unsafe.Pointer(extensions + uintptr(frame_repeat_idx[g])*24))).Fid < int32(32) && (*(*OpusT_opus_extension_data)(unsafe.Pointer(extensions + uintptr(frame_repeat_idx[g])*24))).Flen1 != (*(*OpusT_opus_extension_data)(unsafe.Pointer(extensions + uintptr(i)*24))).Flen1 {
							break
						}
						goto _8
					_8:
						;
						g = g + 1
					}
					if g < nb_frames {
						break
					}
					/* We can! */
					/* If this is a long extension, save the index of the last
					   instance, so we can modify its L flag. */
					if (*(*OpusT_opus_extension_data)(unsafe.Pointer(extensions + uintptr(i)*24))).Fid >= int32(32) {
						last_long_idx = frame_repeat_idx[nb_frames-int32(1)]
					}
					/* Using the repeat mechanism almost always makes the
					    encoding smaller (or at least no larger).
					   However, there's one case where that might not be true: if
					    the last repeated long extension in the last frame was
					    previously the last extension, but using the repeat
					    mechanism makes that no longer true (because there are other
					    non-repeated extensions in earlier frames that must now be
					    coded after it), and coding its length requires more bytes
					    than the repeat mechanism saves.
					   This can only be true if its length is at least 255 bytes
					    (although sometimes it requires even more).
					   Currently we do not check for that, and just always use the
					    repeat mechanism if we can.
					   See git history for code that does the check. */
					/* Advance the repeat pointers. */
					g = f + int32(1)
					for {
						if !(g < nb_frames) {
							break
						}
						j = frame_repeat_idx[g] + int32(1)
						for {
							if !(j < (*(*[48]OpusT_opus_int32)(unsafe.Pointer(bp)))[g] && (*(*OpusT_opus_extension_data)(unsafe.Pointer(extensions + uintptr(j)*24))).Fframe != g) {
								break
							}
							goto _10
						_10:
							;
							j = j + 1
						}
						frame_repeat_idx[g] = j
						goto _9
					_9:
						;
						g = g + 1
					}
					repeat_count = repeat_count + 1
					/* Point the repeat pointer for this frame to the current
					   extension, so we know when to trigger the repeats. */
					frame_repeat_idx[f] = i
				}
				goto _7
			_7:
				;
				i = i + 1
			}
		}
		i = frame_min_idx[f]
		for {
			if !(i < (*(*[48]OpusT_opus_int32)(unsafe.Pointer(bp)))[f]) {
				break
			}
			if (*(*OpusT_opus_extension_data)(unsafe.Pointer(extensions + uintptr(i)*24))).Fframe == f {
				/* Insert separator when needed. */
				if f != curr_frame {
					diff = f - curr_frame
					if len1-pos < int32(2) {
						return -int32(2)
					}
					if diff == int32(1) {
						if data != 0 {
							*(*uint8)(unsafe.Pointer(data + uintptr(pos))) = uint8(0x02)
						}
						pos = pos + 1
					} else {
						if data != 0 {
							*(*uint8)(unsafe.Pointer(data + uintptr(pos))) = uint8(0x03)
						}
						pos = pos + 1
						if data != 0 {
							*(*uint8)(unsafe.Pointer(data + uintptr(pos))) = uint8(diff)
						}
						pos = pos + 1
					}
					curr_frame = f
				}
				pos = write_extension(tls, data, len1, pos, extensions+uintptr(i)*24, libc.BoolInt32(written == nb_extensions-int32(1)))
				if pos < 0 {
					return pos
				}
				written = written + 1
				if repeat_count > 0 && frame_repeat_idx[f] == i {
					/* Add the repeat indicator. */
					nb_repeated = repeat_count * (nb_frames - (f + int32(1)))
					last = libc.BoolInt32(written+nb_repeated == nb_extensions || last_long_idx < 0 && i+int32(1) >= (*(*[48]OpusT_opus_int32)(unsafe.Pointer(bp)))[f])
					if len1-pos < int32(1) {
						return -int32(2)
					}
					if data != 0 {
						*(*uint8)(unsafe.Pointer(data + uintptr(pos))) = uint8(int32(0x04) + libc.BoolInt32(!(last != 0)))
					}
					pos = pos + 1
					g1 = f + int32(1)
					for {
						if !(g1 < nb_frames) {
							break
						}
						j1 = frame_min_idx[g1]
						for {
							if !(j1 < frame_repeat_idx[g1]) {
								break
							}
							if (*(*OpusT_opus_extension_data)(unsafe.Pointer(extensions + uintptr(j1)*24))).Fframe == g1 {
								pos = write_extension_payload(tls, data, len1, pos, extensions+uintptr(j1)*24, libc.BoolInt32(last != 0 && j1 == last_long_idx))
								if pos < 0 {
									return pos
								}
								written = written + 1
							}
							goto _13
						_13:
							;
							j1 = j1 + 1
						}
						frame_min_idx[g1] = j1
						goto _12
					_12:
						;
						g1 = g1 + 1
					}
					if last != 0 {
						curr_frame = curr_frame + 1
					}
				}
			}
			goto _11
		_11:
			;
			i = i + 1
		}
		goto _6
	_6:
		;
		f = f + 1
	}
	if !(written == nb_extensions) {
		Opus_celt_fatal(tls, __ccgo_ts+3039, __ccgo_ts+2472, int32(624))
	}
	/* If we need to pad, just prepend 0x01 bytes. Even better would be to fill the
	   end with zeros, but that requires checking that turning the last extension into
	   an L=1 case still fits. */
	if pad != 0 && pos < len1 {
		padding = len1 - pos
		if data != 0 {
			libc.Xmemmove(tls, data+uintptr(padding), data, uint64(uint32(pos))*uint64(1)+uint64(0*(int64(data+uintptr(padding))-int64(data))))
			i = 0
			for {
				if !(i < padding) {
					break
				}
				*(*uint8)(unsafe.Pointer(data + uintptr(i))) = uint8(0x01)
				goto _14
			_14:
				;
				i = i + 1
			}
		}
		pos = pos + padding
	}
	return pos
}

const BUFSIZ = 1024
