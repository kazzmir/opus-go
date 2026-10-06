// Code generated for linux/amd64 by 'ccgo --package-name opuscc --prefix-external Opus_ --prefix-typename OpusT_ -o opuscc/libopus.go -I .. -I ../include -I ../src -I ../celt -I ../silk -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ../src/opus.c ../src/opus_decoder.c ../src/opus_multistream.c ../src/opus_multistream_decoder.c ../src/mapping_matrix.c ../src/opus_projection_decoder.c ../src/extensions.c ../celt/celt.c ../celt/celt_lpc.c ../celt/kiss_fft.c ../celt/mathops.c ../celt/entdec.c ../celt/cwrs.c ../celt/celt_decoder.c ../celt/pitch.c ../celt/entenc.c ../celt/quant_bands.c ../celt/modes.c ../celt/vq.c ../celt/rate.c ../celt/entcode.c ../celt/bands.c ../celt/mdct.c ../celt/mini_kfft.c ../celt/laplace.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/stereo_MS_to_LR.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/LPC_fit.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c', DO NOT EDIT.

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
	Fframes            [48]*byte
	Flen1              [48]OpusT_opus_int16
	Fframesize         int32
	Fpaddings          [48]*byte
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
	Fbuf         *byte
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
	Fbuf         *byte
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
	Fbuf         *byte
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
	Fpriv         unsafe.Pointer
}

type OpusT_kiss_fft_state = struct {
	Fnfft     int32
	Fscale    OpusT_celt_coef
	Fshift    int32
	Ffactors  [16]OpusT_opus_int16
	Fbitrev   *int16
	Ftwiddles *OpusT_kiss_twiddle_cpx
	Farch_fft *OpusT_arch_fft_state
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

type OpusRepacketizer = OpusT_OpusRepacketizer

type OpusT_OpusExtensionIterator = struct {
	Fdata               *byte
	Fcurr_data          *byte
	Frepeat_data        *byte
	Flast_long          *byte
	Fsrc_data           *byte
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
	Fdata  *byte
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

func Opus_opus_pcm_soft_clip_impl(tls *libc.TLS, pcm *float32, N, C int32, declip_mem *float32, arch int32) {
	if C < 1 || N < 1 || pcm == nil || declip_mem == nil {
		return
	}
	_ = arch
	values, memory := unsafe.Slice(pcm, N*C), unsafe.Slice(declip_mem, C)
	allWithin := Opus_opus_limit2_checkwithin1_c(tls, pcm, N*C)
	for c := int32(0); c < C; c++ {
		x := values[c:]
		a := memory[c]
		// Continue the previous frame's non-linearity through the first crossing.
		for i := int32(0); i < N; i++ {
			if float32(x[i*C]*a) >= 0 {
				break
			}
			x[i*C] = x[i*C] + float32(float32(a*x[i*C])*x[i*C])
		}
		curr := int32(0)
		x0 := x[0]
		for {
			i := curr
			if allWithin != 0 {
				i = N
			} else {
				for i < N {
					if x[i*C] > 1 || x[i*C] < -1 {
						break
					}
					i++
				}
			}
			if i == N {
				a = 0
				break
			}
			start, end, peak := i, i, i
			maxval := float32(libc.Xfabs(tls, float64(x[i*C])))
			for start > 0 && float32(x[i*C]*x[(start-1)*C]) >= 0 {
				start--
			}
			for end < N && float32(x[i*C]*x[end*C]) >= 0 {
				v := float32(libc.Xfabs(tls, float64(x[end*C])))
				if v > maxval {
					maxval = v
					peak = end
				}
				end++
			}
			special := start == 0 && float32(x[i*C]*x[0]) >= 0
			a = (maxval - float32(1)) / float32(maxval*maxval)
			// Preserve the float32 boost and every multiply's rounding.
			a += float32(a * float32(2.4e-7))
			if x[i*C] > 0 {
				a = -a
			}
			for j := start; j < end; j++ {
				x[j*C] = x[j*C] + float32(float32(a*x[j*C])*x[j*C])
			}
			if special && peak >= 2 {
				offset := x0 - x[0]
				delta := offset / float32(peak)
				for j := curr; j < peak; j++ {
					offset -= delta
					x[j*C] += offset
					// C MIN/MAX comparisons preserve NaNs here.
					if x[j*C] > 1 {
						x[j*C] = 1
					}
					if x[j*C] < -1 {
						x[j*C] = -1
					}
				}
			}
			curr = end
			if curr == N {
				break
			}
		}
		memory[c] = a
	}
}

func Opus_opus_pcm_soft_clip(tls *libc.TLS, _x *float32, N int32, C int32, declip_mem *float32) {
	Opus_opus_pcm_soft_clip_impl(tls, _x, N, C, declip_mem, 0)
}

func Opus_encode_size(tls *libc.TLS, size int32, data *byte) (r int32) {
	if size < 252 {
		*data = byte(size)
		return 1
	}
	bytes := unsafe.Slice(data, 2)
	bytes[0] = byte(252 + (size & 3))
	bytes[1] = byte((size - int32(bytes[0])) >> 2)
	return 2
}

func parse_size(tls *libc.TLS, data *byte, len1 OpusT_opus_int32, size *OpusT_opus_int16) (r int32) {
	if len1 < 1 {
		*size = -1
		return -1
	}
	if *data < 252 {
		*size = int16(*data)
		return 1
	}
	if len1 < 2 {
		*size = -1
		return -1
	}
	bytes := unsafe.Slice(data, 2)
	*size = 4*int16(bytes[1]) + int16(bytes[0])
	return 2
}

func Opus_opus_packet_get_samples_per_frame(tls *libc.TLS, data *byte, Fs OpusT_opus_int32) (r int32) {
	toc := *data
	audiosize := (toc >> 3) & 3
	if toc&0x80 != 0 {
		return (Fs << audiosize) / 400
	}
	if toc&0x60 == 0x60 {
		if toc&0x08 != 0 {
			return Fs / 50
		}
		return Fs / 100
	}
	if audiosize == 3 {
		return Fs * 60 / 1000
	}
	return (Fs << audiosize) / 100
}

func Opus_opus_packet_parse_impl(tls *libc.TLS, packet *byte, length, selfDelimited int32, outToc *byte, frames *[48]*byte, size *[48]int16, payloadOffset, packetOffset *int32, padding **byte, paddingLen *int32) int32 {
	if padding != nil {
		*padding = nil
		*paddingLen = 0
	}
	if size == nil || length < 0 {
		return OPUS_BAD_ARG
	}
	if length == 0 {
		return OPUS_INVALID_PACKET
	}
	data := unsafe.Slice(packet, length)
	position := int32(1)
	toc := data[0]
	frameSize := Opus_opus_packet_get_samples_per_frame(tls, packet, 48000)
	length--
	lastSize := length
	count := int32(1)
	cbr := false
	pad := int32(0)
	// The typed base remains live even for zero-length frames at packet end.
	at := func() *byte { return (*byte)(unsafe.Add(unsafe.Pointer(packet), position)) }
	switch toc & 3 {
	case 0:
	case 1:
		count = 2
		cbr = true
		if selfDelimited == 0 {
			if length&1 != 0 {
				return OPUS_INVALID_PACKET
			}
			lastSize = length / 2
			size[0] = int16(lastSize)
		}
	case 2:
		count = 2
		bytes := parse_size(tls, at(), length, &size[0])
		length -= bytes
		if size[0] < 0 || int32(size[0]) > length {
			return OPUS_INVALID_PACKET
		}
		position += bytes
		lastSize = length - int32(size[0])
	default:
		if length < 1 {
			return OPUS_INVALID_PACKET
		}
		ch := data[position]
		position++
		count = int32(ch & 63)
		if count <= 0 || frameSize*count > 5760 {
			return OPUS_INVALID_PACKET
		}
		length--
		if ch&64 != 0 {
			for {
				if length <= 0 {
					return OPUS_INVALID_PACKET
				}
				p := int32(data[position])
				position++
				length--
				tmp := p
				if p == 255 {
					tmp = 254
				}
				length -= tmp
				pad += tmp
				if p != 255 {
					break
				}
			}
		}
		if length < 0 {
			return OPUS_INVALID_PACKET
		}
		cbr = ch&128 == 0
		if !cbr {
			lastSize = length
			for i := int32(0); i < count-1; i++ {
				bytes := parse_size(tls, at(), length, &size[i])
				length -= bytes
				if size[i] < 0 || int32(size[i]) > length {
					return OPUS_INVALID_PACKET
				}
				position += bytes
				lastSize -= bytes + int32(size[i])
			}
			if lastSize < 0 {
				return OPUS_INVALID_PACKET
			}
		} else if selfDelimited == 0 {
			lastSize = length / count
			if lastSize*count != length {
				return OPUS_INVALID_PACKET
			}
			for i := int32(0); i < count-1; i++ {
				size[i] = int16(lastSize)
			}
		}
	}
	if selfDelimited != 0 {
		bytes := parse_size(tls, at(), length, &size[count-1])
		length -= bytes
		if size[count-1] < 0 || int32(size[count-1]) > length {
			return OPUS_INVALID_PACKET
		}
		position += bytes
		if cbr {
			if int32(size[count-1])*count > length {
				return OPUS_INVALID_PACKET
			}
			for i := int32(0); i < count-1; i++ {
				size[i] = size[count-1]
			}
		} else if bytes+int32(size[count-1]) > lastSize {
			return OPUS_INVALID_PACKET
		}
	} else {
		if lastSize > 1275 {
			return OPUS_INVALID_PACKET
		}
		size[count-1] = int16(lastSize)
	}
	if payloadOffset != nil {
		*payloadOffset = position
	}
	for i := int32(0); i < count; i++ {
		if frames != nil {
			frames[i] = at()
		}
		position += int32(size[i])
	}
	if padding != nil {
		*padding = at()
		*paddingLen = pad
	}
	if packetOffset != nil {
		*packetOffset = pad + position
	}
	if outToc != nil {
		*outToc = toc
	}
	return count
}

// Remaining outer decode APIs still store packet/frame addresses as integers.
func opus_packet_parse_impl_legacy(tls *libc.TLS, data uintptr, length, selfDelimited int32, toc, frames, size, payloadOffset, packetOffset, padding, paddingLen uintptr) int32 {
	var framePointers [48]*byte
	var f *[48]*byte
	if frames != 0 {
		f = &framePointers
	}
	var pad *byte
	var p **byte
	if padding != 0 {
		p = &pad
	}
	r := Opus_opus_packet_parse_impl(tls, (*byte)(unsafe.Pointer(data)), length, selfDelimited, (*byte)(unsafe.Pointer(toc)), f, (*[48]int16)(unsafe.Pointer(size)), (*int32)(unsafe.Pointer(payloadOffset)), (*int32)(unsafe.Pointer(packetOffset)), p, (*int32)(unsafe.Pointer(paddingLen)))
	if r > 0 && frames != 0 {
		out := unsafe.Slice((*uintptr)(unsafe.Pointer(frames)), r)
		for i := range out {
			out[i] = uintptr(unsafe.Pointer(framePointers[i]))
		}
	}
	if padding != 0 {
		*(*uintptr)(unsafe.Pointer(padding)) = uintptr(unsafe.Pointer(pad))
	}
	return r
}

func Opus_opus_packet_parse(tls *libc.TLS, data *byte, length int32, toc *byte, frames *[48]*byte, size *[48]int16, payload *int32) int32 {
	return Opus_opus_packet_parse_impl(tls, data, length, 0, toc, frames, size, payload, nil, nil, nil)
}

const OPUS_BAD_ARG = -1
const OPUS_INVALID_PACKET = -4

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
	FeBands         *int16
	FmaxLM          int32
	FnbShortMdcts   int32
	FshortMdctSize  int32
	FnbAllocVectors int32
	FallocVectors   *byte
	FlogN           *int16
	Fwindow         *float32
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
	Fkfft     [4]*OpusT_kiss_fft_state
	Ftrig     *float32
}

type OpusT_PulseCache = struct {
	Fsize  int32
	Findex *int16
	Fbits  *byte
	Fcaps  *byte
}

type OpusCustomMode = OpusT_OpusCustomMode

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
	FCoefs              *int16
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
	FCB1_NLSF_Q8         *byte
	FCB1_Wght_Q9         *int16
	FCB1_iCDF            *byte
	Fpred_Q8             *byte
	Fec_sel              *byte
	Fec_iCDF             *byte
	Fec_Rates_Q5         *byte
	FdeltaMin_Q15        *int16
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
	Fpitch_lag_low_bits_iCDF       *byte
	Fpitch_contour_iCDF            *byte
	FpsNLSF_CB                     *OpusT_silk_NLSF_CB_struct
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
	Fpitch_lag_low_bits_iCDF *byte
	Fpitch_contour_iCDF      *byte
	FnFramesDecoded          int32
	FnFramesPerPacket        int32
	Fec_prevSignalType       int32
	Fec_prevLagIndex         OpusT_opus_int16
	FVAD_flags               [3]int32
	FLBRR_flag               int32
	FLBRR_flags              [3]int32
	Fresampler_state         OpusT_silk_resampler_state_struct
	FpsNLSF_CB               *OpusT_silk_NLSF_CB_struct
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

func validate_opus_decoder(tls *libc.TLS, st *OpusT_OpusDecoder) {
	if !(st.Fchannels == 1 || st.Fchannels == 2) {
		Opus_celt_fatal(tls, __ccgo_ts, __ccgo_ts+57, 99)
	}
	if !(st.FFs == 48000 || st.FFs == 24000 || st.FFs == 16000 || st.FFs == 12000 || st.FFs == 8000) {
		Opus_celt_fatal(tls, __ccgo_ts+79, __ccgo_ts+57, 103)
	}
	if st.FDecControl.FAPI_sampleRate != st.FFs {
		Opus_celt_fatal(tls, __ccgo_ts+188, __ccgo_ts+57, 105)
	}
	dc := &st.FDecControl
	if !(dc.FinternalSampleRate == 0 || dc.FinternalSampleRate == 16000 || dc.FinternalSampleRate == 12000 || dc.FinternalSampleRate == 8000) {
		Opus_celt_fatal(tls, __ccgo_ts+246, __ccgo_ts+57, 106)
	}
	if dc.FnChannelsAPI != st.Fchannels {
		Opus_celt_fatal(tls, __ccgo_ts+440, __ccgo_ts+57, 107)
	}
	if !(dc.FnChannelsInternal == 0 || dc.FnChannelsInternal == 1 || dc.FnChannelsInternal == 2) {
		Opus_celt_fatal(tls, __ccgo_ts+502, __ccgo_ts+57, 108)
	}
	if !(dc.FpayloadSize_ms == 0 || dc.FpayloadSize_ms == 10 || dc.FpayloadSize_ms == 20 || dc.FpayloadSize_ms == 40 || dc.FpayloadSize_ms == 60) {
		Opus_celt_fatal(tls, __ccgo_ts+640, __ccgo_ts+57, 109)
	}
	if st.Farch < 0 {
		Opus_celt_fatal(tls, __ccgo_ts+849, __ccgo_ts+57, 111)
	}
	if st.Farch > OPUS_ARCHMASK {
		Opus_celt_fatal(tls, __ccgo_ts+881, __ccgo_ts+57, 112)
	}
	if !(st.Fstream_channels == 1 || st.Fstream_channels == 2) {
		Opus_celt_fatal(tls, __ccgo_ts+925, __ccgo_ts+57, 114)
	}
}

// opusAlignSize8 retains this port's fixed 8-byte layout alignment and uint32
// wrapping before division/narrowing, including for negative numeric sizes.
func opusAlignSize8(size int32) int32 {
	return int32((uint32(size) + 7) / 8 * 8)
}

func Opus_opus_decoder_get_size(tls *libc.TLS, channels int32) int32 {
	if channels < 1 || channels > 2 {
		return 0
	}
	var silkSize int32
	if Opus_silk_Get_Decoder_Size(tls, &silkSize) != 0 {
		return 0
	}
	silkSize = opusAlignSize8(silkSize)
	celtSize := Opus_celt_decoder_get_size(tls, channels)
	return opusAlignSize8(100) + silkSize + celtSize
}

func Opus_opus_decoder_init(tls *libc.TLS, st *OpusT_OpusDecoder, Fs OpusT_opus_int32, channels int32) int32 {
	if Fs != 48000 && Fs != 24000 && Fs != 16000 && Fs != 12000 && Fs != 8000 || channels != 1 && channels != 2 {
		return OPUS_BAD_ARG
	}
	var silkSize int32
	ret := Opus_silk_Get_Decoder_Size(tls, &silkSize)
	if ret != 0 {
		return -3
	}
	align := func(n int32) int32 { return int32((uint32(n) + 7) &^ uint32(7)) }
	silkOffset := align(int32(unsafe.Sizeof(*st)))
	celtOffset := silkOffset + align(silkSize)
	celt := (*OpusT_OpusCustomDecoder)(unsafe.Add(unsafe.Pointer(st), celtOffset))
	// This is the sole typed pointer slot in the composite decoder layout.
	// Clear it through its typed slot before clearing the surrounding numeric storage.
	celt.Fmode = nil
	clear(unsafe.Slice((*byte)(unsafe.Pointer(st)), Opus_opus_decoder_get_size(tls, channels)))
	st.Fsilk_dec_offset = silkOffset
	st.Fcelt_dec_offset = celtOffset
	st.Fchannels = channels
	st.Fstream_channels = channels
	st.Fcomplexity = 0
	st.FFs = Fs
	st.FDecControl.FAPI_sampleRate = st.FFs
	st.FDecControl.FnChannelsAPI = st.Fchannels
	silk := (*OpusT_silk_decoder)(unsafe.Add(unsafe.Pointer(st), silkOffset))
	if Opus_silk_InitDecoder(tls, silk) != 0 {
		return -3
	}
	if Opus_celt_decoder_init(tls, celt, Fs, channels) != OPUS_OK {
		return -3
	}
	// CELT_SET_SIGNALLING has no validation or other side effects; avoid legacy varargs/TLS allocation.
	celt.Fsignalling = 0
	st.Fprev_mode = 0
	st.Fframe_size = Fs / 400
	st.Farch = 0
	return OPUS_OK
}

func Opus_opus_decoder_create_typed(tls *libc.TLS, Fs OpusT_opus_int32, channels int32) (*OpusT_OpusDecoder, error) {
	if Fs != 48000 && Fs != 24000 && Fs != 16000 && Fs != 12000 && Fs != 8000 || channels != 1 && channels != 2 {
		return nil, opusErrorFromCode(-1)
	}
	st := (*OpusT_OpusDecoder)(libc.XmallocPointer(tls, uint64(uint32(Opus_opus_decoder_get_size(tls, channels)))))
	if st == nil {
		return nil, opusErrorFromCode(-7)
	}
	if ret := Opus_opus_decoder_init(tls, st, Fs, channels); ret != OPUS_OK {
		libc.XfreePointer(tls, unsafe.Pointer(st))
		return nil, opusErrorFromCode(ret)
	}
	return st, nil
}

// Legacy exported creation ABI for the remaining integer-address decode/control callers.
func Opus_opus_decoder_create(tls *libc.TLS, Fs OpusT_opus_int32, channels int32) (uintptr, error) {
	st, err := Opus_opus_decoder_create_typed(tls, Fs, channels)
	return uintptr(unsafe.Pointer(st)), err
}

func smooth_fade(tls *libc.TLS, in1, in2, out *OpusT_opus_res, overlap, channels int32, window *OpusT_celt_coef, Fs OpusT_opus_int32) {
	inc := int32(48000) / Fs
	if overlap <= 0 || channels <= 0 {
		return
	}
	a, b := unsafe.Slice(in1, overlap*channels), unsafe.Slice(in2, overlap*channels)
	dst := unsafe.Slice(out, overlap*channels)
	win := unsafe.Slice(window, (overlap-1)*inc+1)
	// Keep channel-major stores, including when the buffers overlap.
	for c := int32(0); c < channels; c++ {
		for i := int32(0); i < overlap; i++ {
			w := float32(win[i*inc] * win[i*inc])
			idx := i*channels + c
			dst[idx] = float32(w*b[idx]) + float32((float32(1)-w)*a[idx])
		}
	}
}

func opus_packet_get_mode(tls *libc.TLS, data *byte) (r int32) {
	if *data&0x80 != 0 {
		return int32(MODE_CELT_ONLY)
	}
	if *data&0x60 == 0x60 {
		return int32(MODE_HYBRID)
	}
	return int32(MODE_SILK_ONLY)
}

func opusFrameAudioStorage(size int32) []float32 {
	if size == 0 {
		return nil
	}
	return make([]float32, size)
}

func opusFrameCeltState(decoder *OpusT_OpusDecoder) *OpusT_OpusCustomDecoder {
	return (*OpusT_OpusCustomDecoder)(unsafe.Add(unsafe.Pointer(decoder), decoder.Fcelt_dec_offset))
}

func opusFrameSilkPCM(base *float32, offset uintptr) *float32 {
	if offset == 0 {
		return base
	}
	index := offset / 4
	return &unsafe.Slice(base, index+1)[index]
}

func opusFrameSilkState(decoder *OpusT_OpusDecoder) *OpusT_silk_decoder {
	return (*OpusT_silk_decoder)(unsafe.Add(unsafe.Pointer(decoder), decoder.Fsilk_dec_offset))
}

//go:uintptrescapes
func opus_decode_frame(tls *libc.TLS, st1, data uintptr, len1 int32, pcm uintptr, frame_size, decode_fec int32) int32 {
	return opusDecodeFrame(tls, (*OpusT_OpusDecoder)(unsafe.Pointer(st1)), (*byte)(unsafe.Pointer(data)), len1, (*float32)(unsafe.Pointer(pcm)), frame_size, decode_fec)
}

func opusDecodeFrame(tls *libc.TLS, st1 *OpusT_OpusDecoder, data *byte, len1 int32, pcm *float32, frame_size, decode_fec int32) (r int32) {
	var F10, F20, F2_5, F5, audiosize, bandwidth, c, celt_accum, celt_frame_size, celt_ret, celt_to_silk, decoded_samples, endband, first_frame, i, lost_flag, mode, pcm_silk_size, pcm_too_small, pcm_transition_celt_size, pcm_transition_silk_size, redundancy, redundancy_bytes, redundant_audio_size, ret, silk_ret, start_band, transition, v31, v32 int32
	var window *float32
	var silk_dec *OpusT_silk_decoder
	var celt_dec *OpusT_OpusCustomDecoder
	var pcm_base, pcm_ptr *float32
	var pcm_offset uintptr
	var pcm_silk, pcm_transition_celt, pcm_transition_silk []float32
	var pcm_transition, redundant_audio *float32
	var frac, v175, v176 float32
	var gain, x1 OpusT_opus_val32
	var integer OpusT_opus_int32
	var v111 bool
	var res struct { /* union{opus_uint32 i; float f;} of inlined celt_exp2 */
		Fi [0]OpusT_opus_uint32
		Ff float32
	}
	var dec OpusT_ec_dec
	var silk_frame_size OpusT_opus_int32
	var redundant_rng OpusT_opus_uint32
	var silence [2]uint8
	var celt_mode *OpusT_OpusCustomMode
	decoder := st1
	pcm_owner := pcm
	silk_ret = 0
	celt_ret = 0
	pcm_transition = nil
	transition = 0
	redundancy = 0
	redundancy_bytes = 0
	celt_to_silk = 0
	redundant_rng = uint32(0)
	silk_dec = opusFrameSilkState(decoder)
	celt_dec = opusFrameCeltState(decoder)
	F20 = decoder.FFs / int32(50)
	F10 = F20 >> int32(1)
	F5 = F10 >> int32(1)
	F2_5 = F5 >> int32(1)
	if frame_size < F2_5 {
		return -int32(2)
	}
	/* Limit frame_size to avoid excessive stack allocations. */
	if frame_size < decoder.FFs/int32(25)*int32(3) {
		v31 = frame_size
	} else {
		v31 = decoder.FFs / int32(25) * int32(3)
	}
	frame_size = v31
	/* Payloads of 1 (2 including ToC) or 0 trigger the PLC/DTX */
	if len1 <= int32(1) {
		data = nil
		/* In that case, don't conceal more than what the ToC says */
		if frame_size < decoder.Fframe_size {
			v31 = frame_size
		} else {
			v31 = decoder.Fframe_size
		}
		frame_size = v31
	}
	if data != nil {
		audiosize = decoder.Fframe_size
		mode = decoder.Fmode
		bandwidth = decoder.Fbandwidth
		Opus_ec_dec_init(tls, &dec, data, uint32(len1))
	} else {
		audiosize = frame_size
		/* Run PLC using last used mode (CELT if we ended with CELT redundancy) */
		if decoder.Fprev_redundancy != 0 {
			v31 = int32(MODE_CELT_ONLY)
		} else {
			v31 = decoder.Fprev_mode
		}
		mode = v31
		bandwidth = 0
		if mode == 0 {
			/* If we haven't got any packet yet, all we can do is return zeros */
			i = 0
			for {
				if !(i < audiosize*decoder.Fchannels) {
					break
				}
				*opusFrameSilkPCM(pcm_owner, uintptr(i)*4) = float32(0)
				i = i + 1
			}
			return audiosize
		}
		/* Avoids trying to run the PLC on sizes other than 2.5 (CELT), 5 (CELT),
		   10, or 20 (e.g. 12.5 or 30 ms). */
		if audiosize > F20 {
			recursive_offset := uintptr(0)
			for cond := true; cond; cond = audiosize > 0 {
				if audiosize < F20 {
					v31 = audiosize
				} else {
					v31 = F20
				}
				ret = opusDecodeFrame(tls, st1, nil, 0, opusFrameSilkPCM(pcm, recursive_offset), v31, 0)
				if ret < 0 {
					return ret
				}
				recursive_offset += uintptr(ret*decoder.Fchannels) * 4
				audiosize = audiosize - ret
			}
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
	if data != nil && decoder.Fprev_mode > 0 && (mode == int32(MODE_CELT_ONLY) && decoder.Fprev_mode != int32(MODE_CELT_ONLY) && !(decoder.Fprev_redundancy != 0) || mode != int32(MODE_CELT_ONLY) && decoder.Fprev_mode == int32(MODE_CELT_ONLY)) {
		transition = int32(1)
		/* Decide where to allocate the stack memory for pcm_transition */
		if mode == int32(MODE_CELT_ONLY) {
			pcm_transition_celt_size = F5 * decoder.Fchannels
		} else {
			pcm_transition_silk_size = F5 * decoder.Fchannels
		}
	}
	pcm_transition_celt = opusFrameAudioStorage(pcm_transition_celt_size)
	if transition != 0 && mode == int32(MODE_CELT_ONLY) {
		pcm_transition = unsafe.SliceData(pcm_transition_celt)
		if F5 < audiosize {
			v31 = F5
		} else {
			v31 = audiosize
		}
		opusDecodeFrame(tls, st1, nil, 0, pcm_transition, v31, 0)
	}
	if audiosize > frame_size {
		return -int32(1)
	} else {
		frame_size = audiosize
	}
	/* SILK processing */
	if mode != int32(MODE_CELT_ONLY) {
		pcm_silk_size = ALLOC_NONE
		pcm_too_small = libc.BoolInt32(frame_size < F10)
		if pcm_too_small != 0 {
			pcm_silk_size = F10 * decoder.Fchannels
		}
		pcm_silk = opusFrameAudioStorage(pcm_silk_size)
		if pcm_too_small != 0 {
			pcm_base = unsafe.SliceData(pcm_silk)
		} else {
			pcm_base = pcm_owner
		}
		pcm_offset = 0
		if decoder.Fprev_mode == int32(MODE_CELT_ONLY) {
			Opus_silk_ResetDecoder(tls, silk_dec)
		}
		/* The SILK PLC cannot produce frames of less than 10 ms */
		if int32(10) > int32(1000)*audiosize/decoder.FFs {
			v31 = int32(10)
		} else {
			v31 = int32(1000) * audiosize / decoder.FFs
		}
		decoder.FDecControl.FpayloadSize_ms = v31
		if data != nil {
			decoder.FDecControl.FnChannelsInternal = decoder.Fstream_channels
			if mode == int32(MODE_SILK_ONLY) {
				if bandwidth == int32(OPUS_BANDWIDTH_NARROWBAND) {
					decoder.FDecControl.FinternalSampleRate = int32(8000)
				} else {
					if bandwidth == int32(OPUS_BANDWIDTH_MEDIUMBAND) {
						decoder.FDecControl.FinternalSampleRate = int32(12000)
					} else {
						if bandwidth == int32(OPUS_BANDWIDTH_WIDEBAND) {
							decoder.FDecControl.FinternalSampleRate = int32(16000)
						} else {
							decoder.FDecControl.FinternalSampleRate = int32(16000)
							if !(int32(0) != 0) {
								Opus_celt_fatal(tls, __ccgo_ts+1017, __ccgo_ts+57, int32(436))
							}
						}
					}
				}
			} else {
				/* Hybrid mode */
				decoder.FDecControl.FinternalSampleRate = int32(16000)
			}
		}
		decoder.FDecControl.Fenable_deep_plc = libc.BoolInt32(decoder.Fcomplexity >= int32(5))
		if data == nil {
			v31 = int32(1)
		} else {
			v31 = int32(2) * libc.BoolInt32(!!(decode_fec != 0))
		}
		lost_flag = v31
		decoded_samples = 0
		for cond := true; cond; cond = decoded_samples < frame_size {
			pcm_ptr = opusFrameSilkPCM(pcm_base, pcm_offset)
			/* Call SILK decoder */
			first_frame = libc.BoolInt32(decoded_samples == 0)
			silk_ret = silk_Decode(tls, silk_dec, &decoder.FDecControl, lost_flag, first_frame, &dec, pcm_ptr, &silk_frame_size, decoder.Farch)
			if silk_ret != 0 {
				if lost_flag != 0 {
					/* PLC failure should not be fatal */
					silk_frame_size = frame_size
					i = 0
					for {
						if !(i < frame_size*decoder.Fchannels) {
							break
						}
						*opusFrameSilkPCM(pcm_ptr, uintptr(i)*4) = float32(0)
						i = i + 1
					}
				} else {
					return -int32(3)
				}
			}
			pcm_offset += uintptr(silk_frame_size*decoder.Fchannels) * 4
			decoded_samples = decoded_samples + silk_frame_size
		}
		if pcm_too_small != 0 {
			copy(unsafe.Slice(pcm_owner, frame_size*decoder.Fchannels), pcm_silk[:frame_size*decoder.Fchannels])
		}
	}
	start_band = 0
	if v111 = !(decode_fec != 0) && mode != int32(MODE_CELT_ONLY) && data != nil; v111 {
		v31 = dec.Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, dec.Frng))
	}
	if v111 && v31+int32(17)+int32(20)*libc.BoolInt32(mode == int32(MODE_HYBRID)) <= int32(8)*len1 {
		/* Check if we have a redundant 0-8 kHz band */
		if mode == int32(MODE_HYBRID) {
			redundancy = Opus_ec_dec_bit_logp(tls, &dec, uint32(12))
		} else {
			redundancy = int32(1)
		}
		if redundancy != 0 {
			celt_to_silk = Opus_ec_dec_bit_logp(tls, &dec, uint32(1))
			/* redundancy_bytes will be at least two, in the non-hybrid
			   case due to the ec_tell() check above */
			if mode == int32(MODE_HYBRID) {
				v31 = int32(Opus_ec_dec_uint(tls, &dec, uint32(256))) + int32(2)
			} else {
				v32 = dec.Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, dec.Frng))
				v31 = len1 - (v32+int32(7))>>int32(3)
			}
			redundancy_bytes = v31
			len1 = len1 - redundancy_bytes
			/* This is a sanity check. It should never happen for a valid
			   packet, so the exact behaviour is not normative. */
			v31 = dec.Fnbits_total - (int32(4)*int32(CHAR_BIT) - libc.X__builtin_clz(tls, dec.Frng))
			if len1*int32(8) < v31 {
				len1 = 0
				redundancy_bytes = 0
				redundancy = 0
			}
			/* Shrink decoder because of raw bits */
			dec.Fstorage -= uint32(redundancy_bytes)
		}
	}
	if mode != int32(MODE_CELT_ONLY) {
		start_band = int32(17)
	}
	if redundancy != 0 {
		transition = 0
		pcm_transition_silk_size = ALLOC_NONE
	}
	pcm_transition_silk = opusFrameAudioStorage(pcm_transition_silk_size)
	if transition != 0 && mode != int32(MODE_CELT_ONLY) {
		pcm_transition = unsafe.SliceData(pcm_transition_silk)
		if F5 < audiosize {
			v31 = F5
		} else {
			v31 = audiosize
		}
		opusDecodeFrame(tls, st1, nil, 0, pcm_transition, v31, 0)
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
		if !(Opus_opus_custom_decoder_ctl_typed(tls, celt_dec, CELT_SET_END_BAND_REQUEST, OpusDecoderCtlArgs{Value: endband}) == int32(OPUS_OK)) {
			Opus_celt_fatal(tls, __ccgo_ts+1037, __ccgo_ts+57, int32(570))
		}
	}
	_ = decoder.Fstream_channels == int32(0)
	if !(Opus_opus_custom_decoder_ctl_typed(tls, celt_dec, CELT_SET_CHANNELS_REQUEST, OpusDecoderCtlArgs{Value: decoder.Fstream_channels}) == int32(OPUS_OK)) {
		Opus_celt_fatal(tls, __ccgo_ts+1172, __ccgo_ts+57, int32(572))
	}
	/* Only allocation memory for redundancy if/when needed */
	if redundancy != 0 {
		v31 = F5 * decoder.Fchannels
	} else {
		v31 = ALLOC_NONE
	}
	redundant_audio_size = v31
	redundant_audio = unsafe.SliceData(opusFrameAudioStorage(redundant_audio_size))
	/* 5 ms redundant frame for CELT->SILK*/
	if redundancy != 0 && celt_to_silk != 0 {
		/* If the previous frame did not use CELT (the first redundancy frame in
		   a transition from SILK may have been lost) then the CELT decoder is
		   stale at this point and the redundancy audio is not useful, however
		   the final range is still needed (for testing), so the redundancy is
		   always decoded but the decoded audio may not be used */
		_ = int32(0) == int32(0)
		if !(Opus_opus_custom_decoder_ctl_typed(tls, celt_dec, CELT_SET_START_BAND_REQUEST, OpusDecoderCtlArgs{Value: 0}) == int32(OPUS_OK)) {
			Opus_celt_fatal(tls, __ccgo_ts+1331, __ccgo_ts+57, int32(586))
		}
		opusFrameCeltRedundant(tls, celt_dec, data, len1, redundancy_bytes, redundant_audio, F5)
		if !(Opus_opus_custom_decoder_ctl_typed(tls, celt_dec, OPUS_GET_FINAL_RANGE_REQUEST, OpusDecoderCtlArgs{U32: &redundant_rng}) == int32(OPUS_OK)) {
			Opus_celt_fatal(tls, __ccgo_ts+1454, __ccgo_ts+57, int32(589))
		}
	}
	/* MUST be after PLC */
	_ = start_band == int32(0)
	if !(Opus_opus_custom_decoder_ctl_typed(tls, celt_dec, CELT_SET_START_BAND_REQUEST, OpusDecoderCtlArgs{Value: start_band}) == int32(OPUS_OK)) {
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
		if mode != decoder.Fprev_mode && decoder.Fprev_mode > 0 && !(decoder.Fprev_redundancy != 0) {
			if !(Opus_opus_custom_decoder_ctl_typed(tls, celt_dec, OPUS_RESET_STATE, OpusDecoderCtlArgs{}) == int32(OPUS_OK)) {
				Opus_celt_fatal(tls, __ccgo_ts+1740, __ccgo_ts+57, int32(604))
			}
		}
		/* Decode CELT */
		celt_ret = opusFrameCelt(tls, celt_dec, data, len1, pcm_owner, celt_frame_size, &dec, decode_fec, celt_accum)
		Opus_opus_custom_decoder_ctl_typed(tls, celt_dec, OPUS_GET_FINAL_RANGE_REQUEST, OpusDecoderCtlArgs{U32: &decoder.FrangeFinal})
	} else {
		silence = [2]uint8{
			0: uint8(0xFF),
			1: uint8(0xFF),
		}
		if !(celt_accum != 0) {
			i = 0
			for {
				if !(i < frame_size*decoder.Fchannels) {
					break
				}
				*opusFrameSilkPCM(pcm_owner, uintptr(i)*4) = float32(0)
				i = i + 1
			}
		}
		/* For hybrid -> SILK transitions, we let the CELT MDCT
		   do a fade-out by decoding a silence frame */
		if decoder.Fprev_mode == int32(MODE_HYBRID) && !(redundancy != 0 && celt_to_silk != 0 && decoder.Fprev_redundancy != 0) {
			_ = int32(0) == int32(0)
			if !(Opus_opus_custom_decoder_ctl_typed(tls, celt_dec, CELT_SET_START_BAND_REQUEST, OpusDecoderCtlArgs{Value: 0}) == int32(OPUS_OK)) {
				Opus_celt_fatal(tls, __ccgo_ts+1331, __ccgo_ts+57, int32(624))
			}
			opusFrameCeltSilence(tls, celt_dec, &silence, pcm_owner, F2_5, celt_accum)
		}
		decoder.FrangeFinal = dec.Frng
	}
	if !(Opus_opus_custom_decoder_ctl_typed(tls, celt_dec, CELT_GET_MODE_REQUEST, OpusDecoderCtlArgs{Mode: &celt_mode}) == int32(OPUS_OK)) {
		Opus_celt_fatal(tls, __ccgo_ts+1811, __ccgo_ts+57, int32(632))
	}
	window = celt_mode.Fwindow
	/* 5 ms redundant frame for SILK->CELT */
	if redundancy != 0 && !(celt_to_silk != 0) {
		if !(Opus_opus_custom_decoder_ctl_typed(tls, celt_dec, OPUS_RESET_STATE, OpusDecoderCtlArgs{}) == int32(OPUS_OK)) {
			Opus_celt_fatal(tls, __ccgo_ts+1740, __ccgo_ts+57, int32(639))
		}
		_ = int32(0) == int32(0)
		if !(Opus_opus_custom_decoder_ctl_typed(tls, celt_dec, CELT_SET_START_BAND_REQUEST, OpusDecoderCtlArgs{Value: 0}) == int32(OPUS_OK)) {
			Opus_celt_fatal(tls, __ccgo_ts+1331, __ccgo_ts+57, int32(640))
		}
		opusFrameCeltRedundant(tls, celt_dec, data, len1, redundancy_bytes, redundant_audio, F5)
		if !(Opus_opus_custom_decoder_ctl_typed(tls, celt_dec, OPUS_GET_FINAL_RANGE_REQUEST, OpusDecoderCtlArgs{U32: &redundant_rng}) == int32(OPUS_OK)) {
			Opus_celt_fatal(tls, __ccgo_ts+1454, __ccgo_ts+57, int32(643))
		}
		smooth_fade(tls, opusFrameSilkPCM(pcm_owner, uintptr(decoder.Fchannels*(frame_size-F2_5))*4), opusFrameSilkPCM(redundant_audio, uintptr(decoder.Fchannels*F2_5)*4), opusFrameSilkPCM(pcm_owner, uintptr(decoder.Fchannels*(frame_size-F2_5))*4), F2_5, decoder.Fchannels, window, decoder.FFs)
	}
	/* 5ms redundant frame for CELT->SILK; ignore if the previous frame did not
	   use CELT (the first redundancy frame in a transition from SILK may have
	   been lost) */
	if redundancy != 0 && celt_to_silk != 0 && (decoder.Fprev_mode != int32(MODE_SILK_ONLY) || decoder.Fprev_redundancy != 0) {
		c = 0
		for {
			if !(c < decoder.Fchannels) {
				break
			}
			i = 0
			for {
				if !(i < F2_5) {
					break
				}
				*opusFrameSilkPCM(pcm_owner, uintptr(decoder.Fchannels*i+c)*4) = *opusFrameSilkPCM(redundant_audio, uintptr(decoder.Fchannels*i+c)*4)
				i = i + 1
			}
			c = c + 1
		}
		smooth_fade(tls, opusFrameSilkPCM(redundant_audio, uintptr(decoder.Fchannels*F2_5)*4), opusFrameSilkPCM(pcm_owner, uintptr(decoder.Fchannels*F2_5)*4), opusFrameSilkPCM(pcm_owner, uintptr(decoder.Fchannels*F2_5)*4), F2_5, decoder.Fchannels, window, decoder.FFs)
	}
	if transition != 0 {
		if audiosize >= F5 {
			i = 0
			for {
				if !(i < decoder.Fchannels*F2_5) {
					break
				}
				*opusFrameSilkPCM(pcm_owner, uintptr(i)*4) = *opusFrameSilkPCM(pcm_transition, uintptr(i)*4)
				i = i + 1
			}
			smooth_fade(tls, opusFrameSilkPCM(pcm_transition, uintptr(decoder.Fchannels*F2_5)*4), opusFrameSilkPCM(pcm_owner, uintptr(decoder.Fchannels*F2_5)*4), opusFrameSilkPCM(pcm_owner, uintptr(decoder.Fchannels*F2_5)*4), F2_5, decoder.Fchannels, window, decoder.FFs)
		} else {
			/* Not enough time to do a clean transition, but we do it anyway
			   This will not preserve amplitude perfectly and may introduce
			   a bit of temporal aliasing, but it shouldn't be too bad and
			   that's pretty much the best we can do. In any case, generating this
			   transition it pretty silly in the first place */
			smooth_fade(tls, pcm_transition, pcm_owner, pcm_owner, F2_5, decoder.Fchannels, window, decoder.FFs)
		}
	}
	if decoder.Fdecode_gain != 0 {
		v175 = float32(float32(0.000648814081) * float32(decoder.Fdecode_gain))
		integer = int32(libc.Xfloor(tls, float64(v175)))
		if integer < -int32(50) {
			v176 = float32(0)
			goto _177
		}
		frac = v175 - float32(integer)
		res.Ff = float32(0.9999999403953552) + float32(frac*(float32(0.6931530833244324)+float32(frac*(float32(0.24015361070632935)+float32(frac*(float32(0.05582631751894951)+float32(frac*(float32(0.00898933969438076)+float32(frac*float32(0.0018775766948238015))))))))))
		*(*OpusT_opus_uint32)(unsafe.Pointer(&res)) = uint32(int32(*(*OpusT_opus_uint32)(unsafe.Pointer(&res)))+int32(uint32(integer)<<int32(23))) & uint32(0x7fffffff)
		v176 = res.Ff
	_177:
		gain = v176
		i = 0
		for {
			if !(i < frame_size*decoder.Fchannels) {
				break
			}
			x1 = OpusT_opus_res(*opusFrameSilkPCM(pcm_owner, uintptr(i)*4) * gain)
			*opusFrameSilkPCM(pcm_owner, uintptr(i)*4) = x1
			i = i + 1
		}
	}
	if len1 <= int32(1) {
		decoder.FrangeFinal = uint32(0)
	} else {
		decoder.FrangeFinal ^= redundant_rng
	}
	decoder.Fprev_mode = mode
	decoder.Fprev_redundancy = libc.BoolInt32(redundancy != 0 && !(celt_to_silk != 0))
	if celt_ret >= 0 {
		v31 = 0
		if v31 != 0 {
		}
	}
	if celt_ret < 0 {
		v31 = celt_ret
	} else {
		v31 = audiosize
	}
	return v31
}

// The public parser preserves C EOF pointers. Native decode only consumes padding
// when its length is nonzero, so derive that view from numeric descriptors instead.
func opusNativeParsePacket(tls *libc.TLS, data *byte, length, selfDelimited int32, toc *byte, size *[48]int16, offset, packetOffset *int32, padding **byte, paddingLength *int32) int32 {
	var consumed int32
	output := packetOffset
	if output == nil {
		output = &consumed
	}
	count := Opus_opus_packet_parse_impl(tls, data, length, selfDelimited, toc, nil, size, offset, output, nil, nil)
	if count < 0 {
		return count
	}
	position := *offset
	for i := int32(0); i < count; i++ {
		position += int32(size[i])
	}
	*paddingLength = *output - position
	*padding = nil
	if *paddingLength != 0 {
		*padding = &unsafe.Slice(data, position+*paddingLength)[position]
	}
	return count
}

func opusNativePayload(packet *byte, offset uintptr, length int32) *byte {
	if length <= 1 {
		return nil
	}
	return &unsafe.Slice(packet, offset+uintptr(length))[offset]
}

func opusNativePacketFrame(tls *libc.TLS, decoder *OpusT_OpusDecoder, data *byte, length int32, pcm *float32, count, total int32) int32 {
	return opusDecodeFrame(tls, decoder, data, length, opusFrameSilkPCM(pcm, uintptr(count*decoder.Fchannels)*4), total-count, 0)
}

func opusNativeFECFrame(tls *libc.TLS, decoder *OpusT_OpusDecoder, data *byte, length int32, pcm *float32, total, packetFrame int32) int32 {
	return opusDecodeFrame(tls, decoder, data, length, opusFrameSilkPCM(pcm, uintptr(decoder.Fchannels*(total-packetFrame))*4), packetFrame, 1)
}

func opusNativePLCFrame(tls *libc.TLS, decoder *OpusT_OpusDecoder, pcm *float32, count, total int32) int32 {
	return opusDecodeFrame(tls, decoder, nil, 0, opusFrameSilkPCM(pcm, uintptr(count*decoder.Fchannels)*4), total-count, 0)
}

//go:uintptrescapes
func Opus_opus_decode_native(tls *libc.TLS, st, data uintptr, len1 int32, pcm uintptr, frame_size, decode_fec, self_delimited int32, packet_offset uintptr, soft_clip int32, dred uintptr, dred_offset int32) int32 {
	return opusDecodeNative(tls, (*OpusT_OpusDecoder)(unsafe.Pointer(st)), (*byte)(unsafe.Pointer(data)), len1, (*float32)(unsafe.Pointer(pcm)), frame_size, decode_fec, self_delimited, (*int32)(unsafe.Pointer(packet_offset)), soft_clip)
}

func opusDecodeNative(tls *libc.TLS, decoder *OpusT_OpusDecoder, data *byte, len1 int32, pcm *float32, frame_size, decode_fec, self_delimited int32, packet_offset *int32, soft_clip int32) (r int32) {
	var count, duration_copy, i, nb_samples, packet_bandwidth, packet_frame_size, packet_mode, packet_stream_channels, pcm_count, ret, ret1, ret2, v1 int32
	var v8 OpusT_opus_val16
	var toc uint8
	var size [48]OpusT_opus_int16 /* 48 x 2.5 ms = 120 ms */
	var offset int32
	var padding *byte
	var padding_len OpusT_opus_int32
	var iter OpusT_OpusExtensionIterator
	validate_opus_decoder(tls, decoder)
	if decode_fec < 0 || decode_fec > int32(1) {
		return -int32(1)
	}
	/* For FEC/PLC, frame_size has to be to have a multiple of 2.5 ms */
	if (decode_fec != 0 || len1 == 0 || data == nil) && frame_size%(decoder.FFs/int32(400)) != 0 {
		return -int32(1)
	}
	if len1 == 0 || data == nil {
		pcm_count = 0
		for cond := true; cond; cond = pcm_count < frame_size {
			ret = opusNativePLCFrame(tls, decoder, pcm, pcm_count, frame_size)
			if ret < 0 {
				return ret
			}
			pcm_count = pcm_count + ret
		}
		if !(pcm_count == frame_size) {
			Opus_celt_fatal(tls, __ccgo_ts+1955, __ccgo_ts+57, int32(773))
		}
		v1 = 0
		if v1 != 0 {
		}
		decoder.Flast_packet_duration = pcm_count
		return pcm_count
	} else {
		if len1 < 0 {
			return -int32(1)
		}
	}
	packet_mode = opus_packet_get_mode(tls, data)
	packet_bandwidth = Opus_opus_packet_get_bandwidth(tls, data)
	packet_frame_size = Opus_opus_packet_get_samples_per_frame(tls, data, decoder.FFs)
	packet_stream_channels = Opus_opus_packet_get_nb_channels(tls, data)
	count = opusNativeParsePacket(tls, data, len1, self_delimited, &toc, &size, &offset, packet_offset, &padding, &padding_len)
	if decoder.Fignore_extensions != 0 {
		padding = nil
		padding_len = 0
	}
	if count < 0 {
		return count
	}
	Opus_opus_extension_iterator_init(tls, &iter, padding, padding_len, count)
	payload_offset := uintptr(offset)
	if decode_fec != 0 {
		/* If no FEC can be present, run the PLC (recursive call) */
		if frame_size < packet_frame_size || packet_mode == int32(MODE_CELT_ONLY) || decoder.Fmode == int32(MODE_CELT_ONLY) {
			return opusDecodeNative(tls, decoder, nil, 0, pcm, frame_size, 0, 0, nil, soft_clip)
		}
		/* Otherwise, run the PLC on everything except the size for which we might have FEC */
		duration_copy = decoder.Flast_packet_duration
		if frame_size-packet_frame_size != 0 {
			ret1 = opusDecodeNative(tls, decoder, nil, 0, pcm, frame_size-packet_frame_size, 0, 0, nil, soft_clip)
			if ret1 < 0 {
				decoder.Flast_packet_duration = duration_copy
				return ret1
			}
			if !(ret1 == frame_size-packet_frame_size) {
				Opus_celt_fatal(tls, __ccgo_ts+1997, __ccgo_ts+57, int32(815))
			}
		}
		/* Complete with FEC */
		decoder.Fmode = packet_mode
		decoder.Fbandwidth = packet_bandwidth
		decoder.Fframe_size = packet_frame_size
		decoder.Fstream_channels = packet_stream_channels
		ret1 = opusNativeFECFrame(tls, decoder, opusNativePayload(data, payload_offset, int32(size[0])), int32(size[0]), pcm, frame_size, packet_frame_size)
		if ret1 < 0 {
			return ret1
		} else {
			v1 = 0
			if v1 != 0 {
			}
			decoder.Flast_packet_duration = frame_size
			return frame_size
		}
	}
	if count*packet_frame_size > frame_size {
		return -int32(2)
	}
	/* Update the state as the last step to avoid updating it on an invalid packet */
	decoder.Fmode = packet_mode
	decoder.Fbandwidth = packet_bandwidth
	decoder.Fframe_size = packet_frame_size
	decoder.Fstream_channels = packet_stream_channels
	nb_samples = 0
	i = 0
	for {
		if !(i < count) {
			break
		}
		ret2 = opusNativePacketFrame(tls, decoder, opusNativePayload(data, payload_offset, int32(size[i])), int32(size[i]), pcm, nb_samples, frame_size)
		if ret2 < 0 {
			return ret2
		}
		if !(ret2 == packet_frame_size) {
			Opus_celt_fatal(tls, __ccgo_ts+2049, __ccgo_ts+57, int32(865))
		}
		payload_offset += uintptr(size[i])
		nb_samples = nb_samples + ret2
		i = i + 1
	}
	decoder.Flast_packet_duration = nb_samples
	v1 = 0
	if v1 != 0 {
	}
	if soft_clip != 0 {
		Opus_opus_pcm_soft_clip_impl(tls, pcm, nb_samples, decoder.Fchannels, &decoder.Fsoftclip_mem[0], decoder.Farch)
	} else {
		v8 = float32(0)
		decoder.Fsoftclip_mem[1] = v8
		decoder.Fsoftclip_mem[0] = v8
	}
	return nb_samples
}

//go:uintptrescapes
func Opus_opus_decode(tls *libc.TLS, st1, data uintptr, length int32, pcm uintptr, frameSize, fec int32) int32 {
	return opusDecodeInt16(tls, (*OpusT_OpusDecoder)(unsafe.Pointer(st1)), (*byte)(unsafe.Pointer(data)), length, (*int16)(unsafe.Pointer(pcm)), frameSize, fec)
}

func opusDecodeInt16(tls *libc.TLS, decoder *OpusT_OpusDecoder, data *byte, length int32, pcm *int16, frameSize, fec int32) int32 {
	if frameSize <= 0 {
		return -1
	}
	if data != nil && length > 0 && fec == 0 {
		samples := Opus_opus_decoder_get_nb_samples(tls, decoder, data, length)
		if samples <= 0 {
			return -4
		}
		if frameSize >= samples {
			frameSize = samples
		}
	}
	if !(decoder.Fchannels == 1 || decoder.Fchannels == 2) {
		Opus_celt_fatal(tls, __ccgo_ts, __ccgo_ts+57, 917)
	}
	out := opusFrameAudioStorage(frameSize * decoder.Fchannels)
	result := opusDecodeNative(tls, decoder, data, length, unsafe.SliceData(out), frameSize, fec, 0, nil, OPTIONAL_CLIP)
	if result > 0 {
		_ = decoder.Farch
		Opus_celt_float2int16_c(tls, unsafe.SliceData(out), pcm, result*decoder.Fchannels)
	}
	return result
}

//go:uintptrescapes
func Opus_opus_decode24(tls *libc.TLS, st1, data uintptr, length int32, pcm uintptr, frameSize, fec int32) int32 {
	return opusDecodeInt24(tls, (*OpusT_OpusDecoder)(unsafe.Pointer(st1)), (*byte)(unsafe.Pointer(data)), length, (*int32)(unsafe.Pointer(pcm)), frameSize, fec)
}

func opusDecodeInt24(tls *libc.TLS, decoder *OpusT_OpusDecoder, data *byte, length int32, pcm *int32, frameSize, fec int32) int32 {
	if frameSize <= 0 {
		return -1
	}
	if data != nil && length > 0 && fec == 0 {
		samples := Opus_opus_decoder_get_nb_samples(tls, decoder, data, length)
		if samples <= 0 {
			return -4
		}
		if frameSize >= samples {
			frameSize = samples
		}
	}
	if !(decoder.Fchannels == 1 || decoder.Fchannels == 2) {
		Opus_celt_fatal(tls, __ccgo_ts, __ccgo_ts+57, 966)
	}
	out := opusFrameAudioStorage(frameSize * decoder.Fchannels)
	result := opusDecodeNative(tls, decoder, data, length, unsafe.SliceData(out), frameSize, fec, 0, nil, 0)
	if result > 0 {
		opusDecodeInt24PCM(tls, unsafe.SliceData(out), pcm, result*decoder.Fchannels)
	}
	return result
}

//go:uintptrescapes
func Opus_opus_decode_float(tls *libc.TLS, st, data uintptr, length int32, pcm uintptr, frameSize, fec int32) int32 {
	return opusDecodeFloat(tls, (*OpusT_OpusDecoder)(unsafe.Pointer(st)), (*byte)(unsafe.Pointer(data)), length, (*float32)(unsafe.Pointer(pcm)), frameSize, fec)
}

func opusDecodeInt24PCM(tls *libc.TLS, input *float32, output *int32, count int32) {
	if count <= 0 {
		return
	}
	src, dst := unsafe.Slice(input, count), unsafe.Slice(output, count)
	for i := int32(0); i < count; i++ {
		dst[i] = int32(libc.Xlrintf(tls, float32(float32(float32(32768)*float32(256))*src[i])))
	}
}

func opusDecodeFloat(tls *libc.TLS, decoder *OpusT_OpusDecoder, data *byte, length int32, pcm *float32, frameSize, fec int32) int32 {
	if frameSize <= 0 {
		return -1
	}
	return opusDecodeNative(tls, decoder, data, length, pcm, frameSize, fec, 0, nil, 0)
}

// Legacy varargs boundary; forwarding stays typed and allocates no TLS scratch.
func Opus_opus_decoder_ctl(tls *libc.TLS, st uintptr, request int32, va uintptr) int32 {
	return Opus_opus_decoder_ctl_typed(tls, (*OpusT_OpusDecoder)(unsafe.Pointer(st)), request, opusCtlLegacyArgs(request, va))
}

// The legacy ABI frees an address key without reconstructing a Go pointer.
func Opus_opus_decoder_destroy(tls *libc.TLS, st uintptr) {
	libc.Xfree(tls, st)
}

func Opus_opus_decoder_destroy_typed(tls *libc.TLS, st *OpusT_OpusDecoder) {
	libc.XfreePointer(tls, unsafe.Pointer(st))
}

func Opus_opus_packet_get_bandwidth(tls *libc.TLS, data *byte) (r int32) {
	toc := *data
	if toc&0x80 != 0 {
		bandwidth := int32(OPUS_BANDWIDTH_MEDIUMBAND) + int32((toc>>5)&3)
		if bandwidth == int32(OPUS_BANDWIDTH_MEDIUMBAND) {
			return int32(OPUS_BANDWIDTH_NARROWBAND)
		}
		return bandwidth
	}
	if toc&0x60 == 0x60 {
		if toc&0x10 != 0 {
			return int32(OPUS_BANDWIDTH_FULLBAND)
		}
		return int32(OPUS_BANDWIDTH_SUPERWIDEBAND)
	}
	return int32(OPUS_BANDWIDTH_NARROWBAND) + int32((toc>>5)&3)
}

func Opus_opus_packet_get_nb_channels(tls *libc.TLS, data *byte) (r int32) {
	if *data&4 != 0 {
		return 2
	}
	return 1
}

func Opus_opus_packet_get_nb_frames(tls *libc.TLS, packet *byte, len1 OpusT_opus_int32) (r int32) {
	if len1 < 1 {
		return -1 // OPUS_BAD_ARG
	}
	count := *packet & 3
	if count == 0 {
		return 1
	}
	if count != 3 {
		return 2
	}
	if len1 < 2 {
		return -4 // OPUS_INVALID_PACKET
	}
	return int32(unsafe.Slice(packet, 2)[1] & 0x3f)
}

func Opus_opus_packet_get_nb_samples(tls *libc.TLS, packet *byte, len1 OpusT_opus_int32, Fs OpusT_opus_int32) int32 {
	count := Opus_opus_packet_get_nb_frames(tls, packet, len1)
	if count < 0 {
		return count
	}
	samples := count * Opus_opus_packet_get_samples_per_frame(tls, packet, Fs)
	// Can't have more than 120 ms.
	if samples*25 > Fs*3 {
		return -4
	}
	return samples
}

func Opus_opus_packet_has_lbrr(tls *libc.TLS, packet *byte, length int32) int32 {
	// C reads the TOC before checking length and skips parsing CELT-only packets.
	if opus_packet_get_mode(tls, packet) == MODE_CELT_ONLY {
		return 0
	}
	nbFrames := int32(1)
	frameSize := Opus_opus_packet_get_samples_per_frame(tls, packet, 48000)
	if frameSize > 960 {
		nbFrames = frameSize / 960
	}
	channels := Opus_opus_packet_get_nb_channels(tls, packet)
	var frames [48]*byte
	var sizes [48]int16
	ret := Opus_opus_packet_parse(tls, packet, length, nil, &frames, &sizes, nil)
	if ret <= 0 {
		return ret
	}
	if sizes[0] == 0 {
		return 0
	}
	lbrr := int32(*frames[0]) >> (7 - nbFrames) & 1
	if channels == 2 {
		lbrr = libc.BoolInt32(lbrr != 0 || int32(*frames[0])>>(6-2*nbFrames)&1 != 0)
	}
	return lbrr
}

func Opus_opus_decoder_get_nb_samples(tls *libc.TLS, dec *OpusT_OpusDecoder, packet *byte, len1 OpusT_opus_int32) int32 {
	return Opus_opus_packet_get_nb_samples(tls, packet, len1, dec.FFs)
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

func Opus_opus_dred_decoder_create(tls *libc.TLS) (uintptr, error) {
	var dec, v1 uintptr
	var ret int32
	_, _, _ = dec, ret, v1
	v1 = libc.Xmalloc(tls, uint64(uint32(Opus_opus_dred_decoder_get_size(tls))))
	dec = v1
	if dec == uintptr(uint32(0)) {
		return uintptr(uint32(0)), opusErrorFromCode(-int32(7))
	}
	ret = Opus_opus_dred_decoder_init(tls, dec)
	if ret != OPUS_OK {
		libc.Xfree(tls, dec)
		dec = uintptr(uint32(0))
		return uintptr(uint32(0)), opusErrorFromCode(ret)
	}
	return dec, nil
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

func Opus_opus_dred_alloc(tls *libc.TLS) (uintptr, error) {
	return uintptr(uint32(0)), opusErrorFromCode(-int32(5))
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

func Opus_validate_layout(tls *libc.TLS, channelLayout *OpusT_ChannelLayout) (r int32) {
	var i, max_channel int32
	_, _ = i, max_channel
	max_channel = channelLayout.Fnb_streams + channelLayout.Fnb_coupled_streams
	if max_channel > int32(255) {
		return 0
	}
	i = 0
	for {
		if !(i < channelLayout.Fnb_channels) {
			break
		}
		if int32(channelLayout.Fmapping[i]) >= max_channel && int32(channelLayout.Fmapping[i]) != int32(255) {
			return 0
		}
		i = i + 1
	}
	return int32(1)
}

func Opus_get_left_channel(tls *libc.TLS, layout *OpusT_ChannelLayout, stream_id int32, prev int32) (r int32) {
	var i int32
	if prev >= 0 {
		i = prev + 1
	}
	for ; i < layout.Fnb_channels; i++ {
		if int32(layout.Fmapping[i]) == stream_id*2 {
			return i
		}
	}
	return -1
}

func Opus_get_right_channel(tls *libc.TLS, layout *OpusT_ChannelLayout, stream_id int32, prev int32) (r int32) {
	var i int32
	if prev >= 0 {
		i = prev + 1
	}
	for ; i < layout.Fnb_channels; i++ {
		if int32(layout.Fmapping[i]) == stream_id*2+1 {
			return i
		}
	}
	return -1
}

func Opus_get_mono_channel(tls *libc.TLS, layout *OpusT_ChannelLayout, stream_id int32, prev int32) (r int32) {
	var i int32
	if prev >= 0 {
		i = prev + 1
	}
	for ; i < layout.Fnb_channels; i++ {
		if int32(layout.Fmapping[i]) == stream_id+layout.Fnb_coupled_streams {
			return i
		}
	}
	return -1
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

func validate_ms_decoder(tls *libc.TLS, st *OpusT_OpusMSDecoder) {
	// Preserve C's ignored layout result: this helper does not assert on it.
	Opus_validate_layout(tls, &st.Flayout)
}

func Opus_opus_multistream_decoder_get_size(tls *libc.TLS, nb_streams, nb_coupled_streams int32) int32 {
	if nb_streams < 1 || nb_coupled_streams > nb_streams || nb_coupled_streams < 0 {
		return 0
	}
	coupledSize := Opus_opus_decoder_get_size(tls, 2)
	monoSize := Opus_opus_decoder_get_size(tls, 1)
	return opusAlignSize8(268) + nb_coupled_streams*opusAlignSize8(coupledSize) + (nb_streams-nb_coupled_streams)*opusAlignSize8(monoSize)
}

func Opus_opus_multistream_decoder_init(tls *libc.TLS, st *OpusT_OpusMSDecoder, Fs OpusT_opus_int32, channels, streams, coupled int32, mapping *byte) int32 {
	if channels > 255 || channels < 1 || coupled > streams || streams < 1 || coupled < 0 || streams > 255-coupled {
		return OPUS_BAD_ARG
	}
	layout := &st.Flayout
	layout.Fnb_channels = channels
	layout.Fnb_streams = streams
	layout.Fnb_coupled_streams = coupled
	src := unsafe.Slice(mapping, channels)
	// Preserve forward stores when mapping aliases the layout's own storage.
	for i := int32(0); i < layout.Fnb_channels; i++ {
		layout.Fmapping[i] = src[i]
	}
	if Opus_validate_layout(tls, layout) == 0 {
		return OPUS_BAD_ARG
	}
	offset := int((unsafe.Sizeof(*st) + 7) &^ uintptr(7))
	coupledSize := Opus_opus_decoder_get_size(tls, 2)
	monoSize := Opus_opus_decoder_get_size(tls, 1)
	for i := int32(0); i < layout.Fnb_streams; i++ {
		ch, size := int32(1), monoSize
		if i < layout.Fnb_coupled_streams {
			ch, size = 2, coupledSize
		}
		ptr := unsafe.Add(unsafe.Pointer(st), offset)
		if ret := Opus_opus_decoder_init(tls, (*OpusT_OpusDecoder)(ptr), Fs, ch); ret != OPUS_OK {
			return ret
		}
		// Do not materialize an unused end pointer on an exactly-sized allocation.
		offset += int((uint32(size) + 7) &^ uint32(7))
	}
	return OPUS_OK
}

func Opus_opus_multistream_decoder_create_typed(tls *libc.TLS, Fs OpusT_opus_int32, channels, streams, coupled int32, mapping *byte) (*OpusT_OpusMSDecoder, error) {
	if channels > 255 || channels < 1 || coupled > streams || streams < 1 || coupled < 0 || streams > 255-coupled {
		return nil, opusErrorFromCode(-1)
	}
	st := (*OpusT_OpusMSDecoder)(libc.XmallocPointer(tls, uint64(uint32(Opus_opus_multistream_decoder_get_size(tls, streams, coupled)))))
	if st == nil {
		return nil, opusErrorFromCode(-7)
	}
	if ret := Opus_opus_multistream_decoder_init(tls, st, Fs, channels, streams, coupled, mapping); ret != OPUS_OK {
		libc.XfreePointer(tls, unsafe.Pointer(st))
		return nil, opusErrorFromCode(ret)
	}
	return st, nil
}

// Legacy exported creation ABI for integer-address outer decode/control callers.
func Opus_opus_multistream_decoder_create(tls *libc.TLS, Fs OpusT_opus_int32, channels, streams, coupled int32, mapping uintptr) (uintptr, error) {
	st, err := Opus_opus_multistream_decoder_create_typed(tls, Fs, channels, streams, coupled, (*byte)(unsafe.Pointer(mapping)))
	return uintptr(unsafe.Pointer(st)), err
}

func opus_multistream_packet_validate(tls *libc.TLS, data *byte, length, streams, Fs int32) int32 {
	var toc byte
	var sizes [48]int16
	var packetOffset int32
	samples := int32(0)
	for s := int32(0); s < streams; s++ {
		if length <= 0 {
			return OPUS_INVALID_PACKET
		}
		count := Opus_opus_packet_parse_impl(tls, data, length, libc.BoolInt32(s != streams-1), &toc, nil, &sizes, nil, &packetOffset, nil, nil)
		if count < 0 {
			return count
		}
		nextSamples := Opus_opus_packet_get_nb_samples(tls, data, packetOffset, Fs)
		if s != 0 && samples != nextSamples {
			return OPUS_INVALID_PACKET
		}
		samples = nextSamples
		length -= packetOffset
		if s+1 < streams {
			data = (*byte)(unsafe.Add(unsafe.Pointer(data), packetOffset))
		}
	}
	return samples
}

type OpusT___ccgo_fp__Xopus_multistream_decode_native_4 = func(*libc.TLS, uintptr, int32, int32, uintptr, int32, int32, uintptr)

type opusMSChannelCopy func(*libc.TLS, unsafe.Pointer, int32, int32, *float32, int32, int32)

func opusMSCopyFloat(tls *libc.TLS, dst unsafe.Pointer, ds, dc int32, src *float32, ss, n int32) {
	opus_copy_channel_out_float(tls, (*float32)(dst), ds, dc, src, ss, n)
}

func opusMSCopyShort(tls *libc.TLS, dst unsafe.Pointer, ds, dc int32, src *float32, ss, n int32) {
	opus_copy_channel_out_short(tls, (*int16)(dst), ds, dc, src, ss, n)
}

func opusMSCopyInt24(tls *libc.TLS, dst unsafe.Pointer, ds, dc int32, src *float32, ss, n int32) {
	opus_copy_channel_out_int24(tls, (*int32)(dst), ds, dc, src, ss, n)
}

func opusMSBindCopy(callback, user uintptr) opusMSChannelCopy {
	switch callback {
	case __ccgo_fp(opus_copy_channel_out_float_legacy):
		return opusMSCopyFloat
	case __ccgo_fp(opus_copy_channel_out_short_legacy):
		return opusMSCopyShort
	case __ccgo_fp(opus_copy_channel_out_int24_legacy):
		return opusMSCopyInt24
	default:
		legacy := *(*OpusT___ccgo_fp__Xopus_multistream_decode_native_4)(unsafe.Pointer(&struct{ uintptr }{callback}))
		return opusMSBindLegacyCopy(legacy, user)
	}
}

func opusMSBindLegacyCopy(legacy OpusT___ccgo_fp__Xopus_multistream_decode_native_4, user uintptr) opusMSChannelCopy {
	return func(tls *libc.TLS, dst unsafe.Pointer, ds, dc int32, src *float32, ss, n int32) {
		opusMSInvokeLegacy(tls, legacy, uintptr(dst), ds, dc, uintptr(unsafe.Pointer(src)), ss, n, user)
	}
}

//go:uintptrescapes
func opusMSInvokeLegacy(tls *libc.TLS, callback OpusT___ccgo_fp__Xopus_multistream_decode_native_4, dst uintptr, ds, dc int32, src uintptr, ss, n int32, user uintptr) {
	callback(tls, dst, ds, dc, src, ss, n, user)
}

func opusMSDecoderAt(decoder *OpusT_OpusMSDecoder, offset uintptr) *OpusT_OpusDecoder {
	return (*OpusT_OpusDecoder)(unsafe.Add(unsafe.Pointer(decoder), offset))
}

func opusMSDecodeChild(tls *libc.TLS, decoder *OpusT_OpusDecoder, data *byte, length int32, pcm *float32, frame, fec, selfDelimited int32, packetOffset *int32, softClip int32) int32 {
	return opusDecodeNative(tls, decoder, data, length, pcm, frame, fec, selfDelimited, packetOffset, softClip)
}

//go:uintptrescapes
func Opus_opus_multistream_decode_native(tls *libc.TLS, st1, data uintptr, len1 int32, pcm uintptr, callback uintptr, frame_size, decode_fec, soft_clip int32, user_data uintptr) int32 {
	return opusMSDecodeNative(tls, (*OpusT_OpusMSDecoder)(unsafe.Pointer(st1)), (*byte)(unsafe.Pointer(data)), len1, unsafe.Pointer(pcm), opusMSBindCopy(callback, user_data), frame_size, decode_fec, soft_clip)
}

func opusMSPacketAt(data *byte, offset uintptr, length int32) *byte {
	if length <= 0 {
		return nil
	}
	return (*byte)(unsafe.Add(unsafe.Pointer(data), offset))
}

func opusMSDecodeNative(tls *libc.TLS, decoder *OpusT_OpusMSDecoder, data *byte, len1 int32, pcm unsafe.Pointer, copyChannel opusMSChannelCopy, frame_size, decode_fec, soft_clip int32) (r int32) {
	var position uintptr
	var scratch struct{ Fs, packetOffset int32 } // typed CTL and child outputs need no opaque allocation
	var dec *OpusT_OpusDecoder
	var buf *float32
	var ptr uintptr
	var alignment uint32
	var c, chan1, chan11, coupled_size, do_plc, mono_size, prev, prev1, ret, ret1, s, v31, v56, v75 int32
	validate_ms_decoder(tls, decoder)
	if frame_size <= 0 {
		return -1
	}
	/* Limit frame_size to avoid excessive stack allocations. */
	if !(Opus_opus_multistream_decoder_ctl_typed(tls, decoder, OPUS_GET_SAMPLE_RATE_REQUEST, OpusDecoderCtlArgs{I32: &scratch.Fs}) == int32(OPUS_OK)) {
		Opus_celt_fatal(tls, __ccgo_ts+2090, __ccgo_ts+2200, int32(206))
	}
	if frame_size < scratch.Fs/int32(25)*int32(3) {
		v31 = frame_size
	} else {
		v31 = scratch.Fs / int32(25) * int32(3)
	}
	frame_size = v31
	audio := opusFrameAudioStorage(2 * frame_size)
	buf = unsafe.SliceData(audio)
	alignment = uint32(uint64(uintptr(uint32(0)) + 8))
	v31 = int32((uint32(int32(268)) + alignment - uint32(1)) / alignment * alignment)
	ptr = uintptr(v31) // numeric byte offset; only consumed children become pointers
	coupled_size = Opus_opus_decoder_get_size(tls, int32(2))
	mono_size = Opus_opus_decoder_get_size(tls, int32(1))
	if len1 == 0 {
		do_plc = int32(1)
	}
	if len1 < 0 {
		return -1
	}
	if do_plc == 0 && len1 < 2*decoder.Flayout.Fnb_streams-1 {
		return -4
	}
	if !(do_plc != 0) {
		ret = opus_multistream_packet_validate(tls, data, len1, decoder.Flayout.Fnb_streams, scratch.Fs)
		if ret < 0 {
			return ret
		}
		if ret > frame_size {
			return -2
		}
	}
	s = 0
	for {
		if !(s < decoder.Flayout.Fnb_streams) {
			break
		}
		dec = opusMSDecoderAt(decoder, ptr)
		if s < decoder.Flayout.Fnb_coupled_streams {
			alignment = uint32(uint64(uintptr(uint32(0)) + 8))
			v56 = int32((uint32(coupled_size) + alignment - uint32(1)) / alignment * alignment)
			v31 = v56
		} else {
			alignment = uint32(uint64(uintptr(uint32(0)) + 8))
			v75 = int32((uint32(mono_size) + alignment - uint32(1)) / alignment * alignment)
			v31 = v75
		}
		ptr = ptr + uintptr(v31)
		if do_plc == 0 && len1 <= 0 {
			return -3
		}
		scratch.packetOffset = 0
		ret1 = opusMSDecodeChild(tls, dec, opusMSPacketAt(data, position, len1), len1, buf, frame_size, decode_fec, libc.BoolInt32(s != decoder.Flayout.Fnb_streams-int32(1)), &scratch.packetOffset, soft_clip)
		if !(do_plc != 0) {
			position += uintptr(scratch.packetOffset)
			len1 = len1 - scratch.packetOffset
		}
		if ret1 <= 0 {
			return ret1
		}
		frame_size = ret1
		if s < decoder.Flayout.Fnb_coupled_streams {
			prev = -int32(1)
			/* Copy "left" audio to the channel(s) where it belongs */
			for {
				v31 = Opus_get_left_channel(tls, &decoder.Flayout, s, prev)
				chan1 = v31
				if !(v31 != -int32(1)) {
					break
				}
				copyChannel(tls, pcm, decoder.Flayout.Fnb_channels, chan1, buf, 2, frame_size)
				prev = chan1
			}
			prev = -int32(1)
			/* Copy "right" audio to the channel(s) where it belongs */
			for {
				v31 = Opus_get_right_channel(tls, &decoder.Flayout, s, prev)
				chan1 = v31
				if !(v31 != -int32(1)) {
					break
				}
				copyChannel(tls, pcm, decoder.Flayout.Fnb_channels, chan1, (*float32)(unsafe.Add(unsafe.Pointer(buf), 4)), 2, frame_size)
				prev = chan1
			}
		} else {
			prev1 = -int32(1)
			/* Copy audio to the channel(s) where it belongs */
			for {
				v31 = Opus_get_mono_channel(tls, &decoder.Flayout, s, prev1)
				chan11 = v31
				if !(v31 != -int32(1)) {
					break
				}
				copyChannel(tls, pcm, decoder.Flayout.Fnb_channels, chan11, buf, 1, frame_size)
				prev1 = chan11
			}
		}
		s = s + 1
	}
	/* Handle muted channels */
	c = 0
	for {
		if !(c < decoder.Flayout.Fnb_channels) {
			break
		}
		if int32(decoder.Flayout.Fmapping[c]) == int32(255) {
			copyChannel(tls, pcm, decoder.Flayout.Fnb_channels, c, nil, 0, frame_size)
		}
		c = c + 1
	}
	return frame_size
}

func opus_copy_channel_out_float(tls *libc.TLS, dst *float32, dstStride, dstChannel int32, src *OpusT_opus_res, srcStride, frames int32) {
	if frames <= 0 {
		return
	}
	out := unsafe.Slice(dst, (frames-1)*dstStride+dstChannel+1)
	var in []float32
	if src != nil {
		in = unsafe.Slice(src, (frames-1)*srcStride+1)
	}
	for i := int32(0); i < frames; i++ {
		value := float32(0)
		if in != nil {
			value = in[i*srcStride]
		}
		out[i*dstStride+dstChannel] = value
	}
}

// Adapter for the remaining uintptr-valued multistream callback ABI.
//
//go:uintptrescapes
func opus_copy_channel_out_float_legacy(tls *libc.TLS, dst uintptr, dstStride, dstChannel int32, src uintptr, srcStride, frames int32, userData uintptr) {
	opus_copy_channel_out_float(tls, (*float32)(unsafe.Pointer(dst)), dstStride, dstChannel, (*OpusT_opus_res)(unsafe.Pointer(src)), srcStride, frames)
}

func opus_copy_channel_out_short(tls *libc.TLS, dst *int16, dstStride, dstChannel int32, src *OpusT_opus_res, srcStride, frames int32) {
	if frames <= 0 {
		return
	}
	out := unsafe.Slice(dst, (frames-1)*dstStride+dstChannel+1)
	var in []float32
	if src != nil {
		in = unsafe.Slice(src, (frames-1)*srcStride+1)
	}
	for i := int32(0); i < frames; i++ {
		value := int16(0)
		if in != nil {
			x := float32(in[i*srcStride] * 32768)
			// C MIN/MAX comparisons map NaN to the lower bound, unlike Go min/max.
			if !(x > -32768) {
				x = -32768
			}
			if !(x < 32767) {
				x = 32767
			}
			value = int16(libc.Xlrintf(tls, x))
		}
		out[i*dstStride+dstChannel] = value
	}
}

//go:uintptrescapes
func opus_copy_channel_out_short_legacy(tls *libc.TLS, dst uintptr, dstStride, dstChannel int32, src uintptr, srcStride, frames int32, userData uintptr) {
	opus_copy_channel_out_short(tls, (*int16)(unsafe.Pointer(dst)), dstStride, dstChannel, (*OpusT_opus_res)(unsafe.Pointer(src)), srcStride, frames)
}

func opus_copy_channel_out_int24(tls *libc.TLS, dst *int32, dstStride, dstChannel int32, src *OpusT_opus_res, srcStride, frames int32) {
	if frames <= 0 {
		return
	}
	out := unsafe.Slice(dst, (frames-1)*dstStride+dstChannel+1)
	var in []float32
	if src != nil {
		in = unsafe.Slice(src, (frames-1)*srcStride+1)
	}
	for i := int32(0); i < frames; i++ {
		value := int32(0)
		if in != nil {
			value = int32(libc.Xlrintf(tls, float32(float32(32768*256)*in[i*srcStride])))
		}
		out[i*dstStride+dstChannel] = value
	}
}

//go:uintptrescapes
func opus_copy_channel_out_int24_legacy(tls *libc.TLS, dst uintptr, dstStride, dstChannel int32, src uintptr, srcStride, frames int32, userData uintptr) {
	opus_copy_channel_out_int24(tls, (*int32)(unsafe.Pointer(dst)), dstStride, dstChannel, (*OpusT_opus_res)(unsafe.Pointer(src)), srcStride, frames)
}

//go:uintptrescapes
func Opus_opus_multistream_decode(tls *libc.TLS, st, data uintptr, length int32, pcm uintptr, frame, fec int32) int32 {
	return opusMSDecodeShort(tls, (*OpusT_OpusMSDecoder)(unsafe.Pointer(st)), (*byte)(unsafe.Pointer(data)), length, (*int16)(unsafe.Pointer(pcm)), frame, fec)
}

func opusMSDecodeShort(tls *libc.TLS, decoder *OpusT_OpusMSDecoder, data *byte, length int32, pcm *int16, frame, fec int32) int32 {
	return opusMSDecodeNative(tls, decoder, data, length, unsafe.Pointer(pcm), opusMSCopyShort, frame, fec, OPTIONAL_CLIP)
}

//go:uintptrescapes
func Opus_opus_multistream_decode24(tls *libc.TLS, st, data uintptr, length int32, pcm uintptr, frame, fec int32) int32 {
	return opusMSDecodeInt24(tls, (*OpusT_OpusMSDecoder)(unsafe.Pointer(st)), (*byte)(unsafe.Pointer(data)), length, (*int32)(unsafe.Pointer(pcm)), frame, fec)
}

func opusMSDecodeInt24(tls *libc.TLS, decoder *OpusT_OpusMSDecoder, data *byte, length int32, pcm *int32, frame, fec int32) int32 {
	return opusMSDecodeNative(tls, decoder, data, length, unsafe.Pointer(pcm), opusMSCopyInt24, frame, fec, 0)
}

//go:uintptrescapes
func Opus_opus_multistream_decode_float(tls *libc.TLS, st, data uintptr, length int32, pcm uintptr, frame, fec int32) int32 {
	return opusMSDecodeFloat(tls, (*OpusT_OpusMSDecoder)(unsafe.Pointer(st)), (*byte)(unsafe.Pointer(data)), length, (*float32)(unsafe.Pointer(pcm)), frame, fec)
}

func opusMSDecodeFloat(tls *libc.TLS, decoder *OpusT_OpusMSDecoder, data *byte, length int32, pcm *float32, frame, fec int32) int32 {
	return opusMSDecodeNative(tls, decoder, data, length, unsafe.Pointer(pcm), opusMSCopyFloat, frame, fec, 0)
}

func Opus_opus_multistream_decoder_ctl_va_list(tls *libc.TLS, st uintptr, request int32, ap OpusT_va_list) int32 {
	decoder := (*OpusT_OpusMSDecoder)(unsafe.Pointer(st))
	return Opus_opus_multistream_decoder_ctl_typed(tls, decoder, request, msCtlLegacyArgs(decoder, request, ap))
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

func Opus_opus_multistream_decoder_destroy_typed(tls *libc.TLS, st *OpusT_OpusMSDecoder) {
	// All component decoders belong to the same allocation; free only the base.
	libc.XfreePointer(tls, unsafe.Pointer(st))
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

func Opus_mapping_matrix_get_size(tls *libc.TLS, rows int32, cols int32) int32 {
	/* Mapping Matrix must only support up to 255 channels in or out.
	 * Additionally, the total cell count must be <= 65004 octets in order
	 * for the matrix to be stored in an OGG header.
	 */
	if rows > int32(255) || cols > int32(255) {
		return 0
	}
	size := int32(uint64(uint32(rows*cols)) * uint64(2))
	if size > int32(65004) {
		return 0
	}
	return opusAlignSize8(12) + opusAlignSize8(size)
}

// MappingMatrix is followed by int16 coefficients after its 8-byte-aligned
// header (16 bytes in this build). The pointer must belong to that full allocation.
func mappingMatrixData(matrix *OpusT_MappingMatrix) *int16 {
	return (*int16)(unsafe.Add(unsafe.Pointer(matrix), 16))
}

func Opus_mapping_matrix_get_data(tls *libc.TLS, matrix *OpusT_MappingMatrix) *int16 {
	return mappingMatrixData(matrix)
}

func Opus_mapping_matrix_init(tls *libc.TLS, matrix *OpusT_MappingMatrix, rows, cols, gain int32, data *int16, data_size OpusT_opus_int32) {
	if (uint32(data_size)+7)&^uint32(7) != (uint32(rows*cols)*2+7)&^uint32(7) {
		Opus_celt_fatal(tls, __ccgo_ts+2234, __ccgo_ts+2312, 72)
	}
	matrix.Frows, matrix.Fcols, matrix.Fgain = rows, cols, gain
	if rows*cols == 0 {
		return
	}
	dst := unsafe.Slice(mappingMatrixData(matrix), rows*cols)
	src := unsafe.Slice(data, rows*cols)
	// Preserve the C forward-copy order, including overlapping coefficients.
	for i := range dst {
		dst[i] = src[i]
	}
}

func Opus_mapping_matrix_multiply_channel_in_float(tls *libc.TLS, matrix *OpusT_MappingMatrix, input *float32, input_rows int32, output *OpusT_opus_res, output_row, output_rows, frame_size int32) {
	if !(input_rows <= matrix.Fcols && output_rows <= matrix.Frows) {
		Opus_celt_fatal(tls, __ccgo_ts+2336, __ccgo_ts+2312, 98)
	}
	if frame_size <= 0 {
		return
	}
	data := unsafe.Slice(Opus_mapping_matrix_get_data(tls, matrix), matrix.Frows*matrix.Fcols)
	src := unsafe.Slice(input, frame_size*input_rows)
	dst := unsafe.Slice(output, (frame_size-1)*output_rows+1)
	for i := int32(0); i < frame_size; i++ {
		tmp := float32(0)
		for col := int32(0); col < input_rows; col++ {
			tmp += float32(float32(data[matrix.Frows*col+output_row]) * src[input_rows*i+col])
		}
		// Store only after reading this frame, preserving overlapping-buffer behavior.
		dst[output_rows*i] = float32(float32(1.0/32768) * tmp)
	}
}

func Opus_mapping_matrix_multiply_channel_out_float(tls *libc.TLS, matrix *OpusT_MappingMatrix, input *OpusT_opus_res, input_row, input_rows int32, output *float32, output_rows, frame_size int32) {
	if !(input_rows <= matrix.Fcols && output_rows <= matrix.Frows) {
		Opus_celt_fatal(tls, __ccgo_ts+2336, __ccgo_ts+2312, 130)
	}
	if frame_size <= 0 || output_rows == 0 {
		return
	}
	data := unsafe.Slice(mappingMatrixData(matrix), matrix.Frows*matrix.Fcols)
	src := unsafe.Slice(input, (frame_size-1)*input_rows+1)
	dst := unsafe.Slice(output, frame_size*output_rows)
	for i := int32(0); i < frame_size; i++ {
		sample := src[input_rows*i]
		for row := int32(0); row < output_rows; row++ {
			factor := float32(float32(1.0/32768) * float32(data[matrix.Frows*input_row+row]))
			tmp := float32(factor * sample)
			dst[output_rows*i+row] += tmp
		}
	}
}

func Opus_mapping_matrix_multiply_channel_in_short(tls *libc.TLS, matrix *OpusT_MappingMatrix, input *int16, input_rows int32, output *OpusT_opus_res, output_row, output_rows, frame_size int32) {
	if !(input_rows <= matrix.Fcols && output_rows <= matrix.Frows) {
		Opus_celt_fatal(tls, __ccgo_ts+2336, __ccgo_ts+2312, 161)
	}
	if frame_size <= 0 {
		return
	}
	data := unsafe.Slice(Opus_mapping_matrix_get_data(tls, matrix), matrix.Frows*matrix.Fcols)
	src := unsafe.Slice(input, frame_size*input_rows)
	dst := unsafe.Slice(output, (frame_size-1)*output_rows+1)
	for i := int32(0); i < frame_size; i++ {
		tmp := float32(0)
		for col := int32(0); col < input_rows; col++ {
			// C promotes both int16 operands to int before converting the product.
			tmp += float32(int32(data[matrix.Frows*col+output_row]) * int32(src[input_rows*i+col]))
		}
		dst[output_rows*i] = float32(float32(1.0/(32768*32768)) * tmp)
	}
}

func Opus_mapping_matrix_multiply_channel_out_short(tls *libc.TLS, matrix *OpusT_MappingMatrix, input *OpusT_opus_res, input_row, input_rows int32, output *int16, output_rows, frame_size int32) {
	if !(input_rows <= matrix.Fcols && output_rows <= matrix.Frows) {
		Opus_celt_fatal(tls, __ccgo_ts+2336, __ccgo_ts+2312, 206)
	}
	if frame_size <= 0 || output_rows == 0 {
		return
	}
	data := unsafe.Slice(mappingMatrixData(matrix), matrix.Frows*matrix.Fcols)
	src := unsafe.Slice(input, (frame_size-1)*input_rows+1)
	dst := unsafe.Slice(output, frame_size*output_rows)
	for i := int32(0); i < frame_size; i++ {
		scaled := float32(src[input_rows*i] * 32768)
		// Match MAX32 then MIN32, including NaN -> -32768.
		if !(scaled > -32768) {
			scaled = -32768
		}
		if !(scaled < 32767) {
			scaled = 32767
		}
		sample := int32(int16(libc.Xlrintf(tls, scaled)))
		for row := int32(0); row < output_rows; row++ {
			tmp := int32(data[matrix.Frows*input_row+row]) * sample
			idx := output_rows*i + row
			dst[idx] = int16(int32(dst[idx]) + ((tmp + 16384) >> 15))
		}
	}
}

func Opus_mapping_matrix_multiply_channel_in_int24(tls *libc.TLS, matrix *OpusT_MappingMatrix, input *int32, input_rows int32, output *OpusT_opus_res, output_row, output_rows, frame_size int32) {
	if !(input_rows <= matrix.Fcols && output_rows <= matrix.Frows) {
		Opus_celt_fatal(tls, __ccgo_ts+2336, __ccgo_ts+2312, 236)
	}
	if frame_size <= 0 {
		return
	}
	data := unsafe.Slice(Opus_mapping_matrix_get_data(tls, matrix), matrix.Frows*matrix.Fcols)
	src := unsafe.Slice(input, frame_size*input_rows)
	dst := unsafe.Slice(output, (frame_size-1)*output_rows+1)
	for i := int32(0); i < frame_size; i++ {
		// opus_val64 is float32 in this build, not a widened accumulator.
		tmp := float32(0)
		for col := int32(0); col < input_rows; col++ {
			tmp += float32(float32(data[matrix.Frows*col+output_row]) * float32(src[input_rows*i+col]))
		}
		// Preserve the two rounded scales in INT24TORES((1/32768.f)*tmp).
		dst[output_rows*i] = float32(float32(1.0/32768/256) * float32(float32(1.0/32768)*tmp))
	}
}

func Opus_mapping_matrix_multiply_channel_out_int24(tls *libc.TLS, matrix *OpusT_MappingMatrix, input *OpusT_opus_res, input_row, input_rows int32, output *int32, output_rows, frame_size int32) {
	if !(input_rows <= matrix.Fcols && output_rows <= matrix.Frows) {
		Opus_celt_fatal(tls, __ccgo_ts+2336, __ccgo_ts+2312, 271)
	}
	if frame_size <= 0 || output_rows == 0 {
		return
	}
	data := unsafe.Slice(mappingMatrixData(matrix), matrix.Frows*matrix.Fcols)
	src := unsafe.Slice(input, (frame_size-1)*input_rows+1)
	dst := unsafe.Slice(output, frame_size*output_rows)
	for i := int32(0); i < frame_size; i++ {
		sample := int32(libc.Xlrintf(tls, float32(float32(32768*256)*src[input_rows*i])))
		for row := int32(0); row < output_rows; row++ {
			tmp := int64(data[matrix.Frows*input_row+row]) * int64(sample)
			idx := output_rows*i + row
			// Accumulate in 64 bits, then narrow; do not clip to 24 bits.
			dst[idx] = int32(int64(dst[idx]) + ((tmp + 16384) >> 15))
		}
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

func opus_projection_copy_channel_out_float(tls *libc.TLS, dst *float32, dstStride, dstChannel int32, src *OpusT_opus_res, srcStride, frames int32, matrix *OpusT_MappingMatrix) {
	if frames <= 0 {
		return
	}
	// Clear before reading any aliased source samples, matching the C callback.
	if dstChannel == 0 {
		clear(unsafe.Slice(dst, frames*dstStride))
	}
	if src != nil {
		Opus_mapping_matrix_multiply_channel_out_float(tls, matrix, src, dstChannel, srcStride, dst, dstStride, frames)
	}
}

func opus_projection_copy_channel_out_short(tls *libc.TLS, dst *int16, dstStride, dstChannel int32, src *OpusT_opus_res, srcStride, frames int32, matrix *OpusT_MappingMatrix) {
	if frames <= 0 {
		return
	}
	if dstChannel == 0 {
		clear(unsafe.Slice(dst, frames*dstStride))
	}
	if src != nil {
		Opus_mapping_matrix_multiply_channel_out_short(tls, matrix, src, dstChannel, srcStride, dst, dstStride, frames)
	}
}

func opus_projection_copy_channel_out_int24(tls *libc.TLS, dst *int32, dstStride, dstChannel int32, src *OpusT_opus_res, srcStride, frames int32, matrix *OpusT_MappingMatrix) {
	if frames <= 0 {
		return
	}
	if dstChannel == 0 {
		clear(unsafe.Slice(dst, frames*dstStride))
	}
	if src != nil {
		Opus_mapping_matrix_multiply_channel_out_int24(tls, matrix, src, dstChannel, srcStride, dst, dstStride, frames)
	}
}

// The header must belong to the complete projection decoder backing allocation.
func get_dec_demixing_matrix(tls *libc.TLS, st *OpusT_OpusProjectionDecoder) *OpusT_MappingMatrix {
	return (*OpusT_MappingMatrix)(unsafe.Add(unsafe.Pointer(st), 8))
}

// st belongs to the full header/matrix/multistream backing allocation.
func get_multistream_decoder(tls *libc.TLS, st *OpusT_OpusProjectionDecoder) *OpusT_OpusMSDecoder {
	offset := int32((uint32(st.Fdemixing_matrix_size_in_bytes) + 4 + 7) / 8 * 8)
	return (*OpusT_OpusMSDecoder)(unsafe.Add(unsafe.Pointer(st), uintptr(offset)))
}

func Opus_opus_projection_decoder_get_size(tls *libc.TLS, channels, streams, coupled_streams int32) int32 {
	matrixSize := Opus_mapping_matrix_get_size(tls, streams+coupled_streams, channels)
	if matrixSize == 0 {
		return 0
	}
	decoderSize := Opus_opus_multistream_decoder_get_size(tls, streams, coupled_streams)
	if decoderSize == 0 {
		return 0
	}
	return opusAlignSize8(4) + matrixSize + decoderSize
}

func Opus_opus_projection_decoder_init(tls *libc.TLS, st *OpusT_OpusProjectionDecoder, Fs OpusT_opus_int32, channels, streams, coupled int32, matrix *byte, matrixBytes OpusT_opus_int32) int32 {
	inputs := streams + coupled
	count := inputs * channels
	expected := count * 2
	if expected != matrixBytes {
		return OPUS_BAD_ARG
	}
	// Snapshot all coefficients before writing state, preserving aliased input.
	src := unsafe.Slice(matrix, matrixBytes)
	coefficients := make([]int16, count)
	for i := int32(0); i < count; i++ {
		s := int32(src[2*i+1])<<8 | int32(src[2*i])
		coefficients[i] = int16(((s & 0xffff) ^ 0x8000) - 0x8000)
	}
	st.Fdemixing_matrix_size_in_bytes = Opus_mapping_matrix_get_size(tls, channels, inputs)
	if st.Fdemixing_matrix_size_in_bytes == 0 {
		return OPUS_BAD_ARG
	}
	Opus_mapping_matrix_init(tls, get_dec_demixing_matrix(tls, st), channels, inputs, 0, unsafe.SliceData(coefficients), matrixBytes)
	var mapping [255]byte
	for i := int32(0); i < channels; i++ {
		mapping[i] = byte(i)
	}
	return Opus_opus_multistream_decoder_init(tls, get_multistream_decoder(tls, st), Fs, channels, streams, coupled, &mapping[0])
}

func Opus_opus_projection_decoder_create_typed(tls *libc.TLS, Fs OpusT_opus_int32, channels, streams, coupled int32, matrix *byte, matrixBytes OpusT_opus_int32) (*OpusT_OpusProjectionDecoder, error) {
	// Projection creation checks size/allocation before initialization arguments.
	size := Opus_opus_projection_decoder_get_size(tls, channels, streams, coupled)
	if size == 0 {
		return nil, opusErrorFromCode(-7)
	}
	st := (*OpusT_OpusProjectionDecoder)(libc.XmallocPointer(tls, uint64(uint32(size))))
	if st == nil {
		return nil, opusErrorFromCode(-7)
	}
	if ret := Opus_opus_projection_decoder_init(tls, st, Fs, channels, streams, coupled, matrix, matrixBytes); ret != OPUS_OK {
		libc.XfreePointer(tls, unsafe.Pointer(st))
		return nil, opusErrorFromCode(ret)
	}
	return st, nil
}

// Legacy exported creation ABI for the remaining integer-address projection APIs.
func Opus_opus_projection_decoder_create(tls *libc.TLS, Fs OpusT_opus_int32, channels, streams, coupled int32, matrix uintptr, matrixBytes OpusT_opus_int32) (uintptr, error) {
	st, err := Opus_opus_projection_decoder_create_typed(tls, Fs, channels, streams, coupled, (*byte)(unsafe.Pointer(matrix)), matrixBytes)
	return uintptr(unsafe.Pointer(st)), err
}

//go:uintptrescapes
func Opus_opus_projection_decode(tls *libc.TLS, st, data uintptr, length int32, pcm uintptr, frame, fec int32) int32 {
	return opusProjectionDecodeShort(tls, (*OpusT_OpusProjectionDecoder)(unsafe.Pointer(st)), (*byte)(unsafe.Pointer(data)), length, (*int16)(unsafe.Pointer(pcm)), frame, fec)
}

func opusProjectionShortCopy(matrix *OpusT_MappingMatrix) opusMSChannelCopy {
	return func(tls *libc.TLS, dst unsafe.Pointer, ds, dc int32, src *float32, ss, n int32) {
		opus_projection_copy_channel_out_short(tls, (*int16)(dst), ds, dc, src, ss, n, matrix)
	}
}

func opusProjectionDecodeShort(tls *libc.TLS, decoder *OpusT_OpusProjectionDecoder, data *byte, length int32, pcm *int16, frame, fec int32) int32 {
	ms := get_multistream_decoder(tls, decoder)
	matrix := get_dec_demixing_matrix(tls, decoder)
	return opusMSDecodeNative(tls, ms, data, length, unsafe.Pointer(pcm), opusProjectionShortCopy(matrix), frame, fec, OPTIONAL_CLIP)
}

//go:uintptrescapes
func Opus_opus_projection_decode24(tls *libc.TLS, st, data uintptr, length int32, pcm uintptr, frame, fec int32) int32 {
	return opusProjectionDecodeInt24(tls, (*OpusT_OpusProjectionDecoder)(unsafe.Pointer(st)), (*byte)(unsafe.Pointer(data)), length, (*int32)(unsafe.Pointer(pcm)), frame, fec)
}

func opusProjectionInt24Copy(matrix *OpusT_MappingMatrix) opusMSChannelCopy {
	return func(tls *libc.TLS, dst unsafe.Pointer, ds, dc int32, src *float32, ss, n int32) {
		opus_projection_copy_channel_out_int24(tls, (*int32)(dst), ds, dc, src, ss, n, matrix)
	}
}

func opusProjectionDecodeInt24(tls *libc.TLS, decoder *OpusT_OpusProjectionDecoder, data *byte, length int32, pcm *int32, frame, fec int32) int32 {
	ms := get_multistream_decoder(tls, decoder)
	matrix := get_dec_demixing_matrix(tls, decoder)
	return opusMSDecodeNative(tls, ms, data, length, unsafe.Pointer(pcm), opusProjectionInt24Copy(matrix), frame, fec, 0)
}

//go:uintptrescapes
func Opus_opus_projection_decode_float(tls *libc.TLS, st, data uintptr, length int32, pcm uintptr, frame, fec int32) int32 {
	return opusProjectionDecodeFloat(tls, (*OpusT_OpusProjectionDecoder)(unsafe.Pointer(st)), (*byte)(unsafe.Pointer(data)), length, (*float32)(unsafe.Pointer(pcm)), frame, fec)
}

func opusProjectionFloatCopy(matrix *OpusT_MappingMatrix) opusMSChannelCopy {
	return func(tls *libc.TLS, dst unsafe.Pointer, ds, dc int32, src *float32, ss, n int32) {
		opus_projection_copy_channel_out_float(tls, (*float32)(dst), ds, dc, src, ss, n, matrix)
	}
}

func opusProjectionDecodeFloat(tls *libc.TLS, decoder *OpusT_OpusProjectionDecoder, data *byte, length int32, pcm *float32, frame, fec int32) int32 {
	ms := get_multistream_decoder(tls, decoder)
	matrix := get_dec_demixing_matrix(tls, decoder)
	return opusMSDecodeNative(tls, ms, data, length, unsafe.Pointer(pcm), opusProjectionFloatCopy(matrix), frame, fec, 0)
}

func Opus_opus_projection_decoder_ctl(tls *libc.TLS, st uintptr, request int32, va uintptr) int32 {
	decoder := (*OpusT_OpusProjectionDecoder)(unsafe.Pointer(st))
	ms := get_multistream_decoder(tls, decoder)
	return Opus_opus_projection_decoder_ctl_typed(tls, decoder, request, msCtlLegacyArgs(ms, request, va))
}

func Opus_opus_projection_decoder_destroy(tls *libc.TLS, st uintptr) {
	libc.Xfree(tls, st)
}

func Opus_opus_projection_decoder_destroy_typed(tls *libc.TLS, st *OpusT_OpusProjectionDecoder) {
	// Demixing coefficients and multistream decoders are interiors, not owners.
	libc.XfreePointer(tls, unsafe.Pointer(st))
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
func skip_extension_payload(tls *libc.TLS, pdata **byte, length int32, headerSize *int32, idByte, trailingShort int32) int32 {
	data := *pdata
	header := int32(0)
	id, L := idByte>>1, idByte&1
	if (id == 0 && L == 1) || id == 2 {
	} else if id > 0 && id < 32 {
		if length < L {
			return -1
		}
		data = (*byte)(unsafe.Add(unsafe.Pointer(data), L))
		length -= L
	} else if L == 0 {
		if length < trailingShort {
			return -1
		}
		data = (*byte)(unsafe.Add(unsafe.Pointer(data), length-trailingShort))
		length = trailingShort
	} else {
		bytes := int32(0)
		for {
			if length < 1 {
				return -1
			}
			lacing := int32(*data)
			data = (*byte)(unsafe.Add(unsafe.Pointer(data), 1))
			bytes += lacing
			header++
			length -= lacing + 1
			if lacing != 255 {
				break
			}
		}
		if length < 0 {
			return -1
		}
		data = (*byte)(unsafe.Add(unsafe.Pointer(data), bytes))
	}
	*pdata = data
	*headerSize = header
	return length
}

// C documentation
//
//	/* Given an extension, advance data to the next extension and return the
//	   length of the remaining extensions.
//	   N.B., a "Repeat These Extensions" extension (ID==2) only advances past the
//	    extension ID byte.
//	   Higher-level logic is required to skip the extension payloads that come
//	    after it.*/
func skip_extension(tls *libc.TLS, pdata **byte, length int32, header *int32) int32 {
	if length == 0 {
		*header = 0
		return 0
	}
	if length < 1 {
		return -1
	}
	data := *pdata
	id := int32(*data)
	data = (*byte)(unsafe.Add(unsafe.Pointer(data), 1))
	result := skip_extension_payload(tls, &data, length-1, header, id, 0)
	if result >= 0 {
		*pdata = data
		*header += 1
	}
	return result
}

func Opus_opus_extension_iterator_init(tls *libc.TLS, iter *OpusT_OpusExtensionIterator, data *byte, length, frames int32) {
	if length < 0 {
		Opus_celt_fatal(tls, __ccgo_ts+2445, __ccgo_ts+2472, 122)
	}
	if data == nil && length != 0 {
		Opus_celt_fatal(tls, __ccgo_ts+2492, __ccgo_ts+2472, 123)
	}
	if frames < 0 || frames > 48 {
		Opus_celt_fatal(tls, __ccgo_ts+2535, __ccgo_ts+2472, 124)
	}
	iter.Fdata = data
	iter.Fcurr_data = data
	iter.Frepeat_data = data
	iter.Fsrc_data = nil
	iter.Flast_long = nil
	iter.Flen1 = length
	iter.Fcurr_len = length
	iter.Fsrc_len = 0
	iter.Frepeat_len = 0
	iter.Ftrailing_short_len = 0
	iter.Fnb_frames = frames
	iter.Fframe_max = frames
	iter.Fcurr_frame = 0
	iter.Frepeat_frame = 0
	iter.Frepeat_l = 0
}

func extensionPointerDifference(a, b *byte) int64 {
	// Only compare addresses; never reconstruct a pointer from this difference.
	return int64(uintptr(unsafe.Pointer(a))) - int64(uintptr(unsafe.Pointer(b)))
}

// C documentation
//
//	/* Reset the iterator so it can start iterating again from the first
//	    extension. */
func Opus_opus_extension_iterator_reset(tls *libc.TLS, iter *OpusT_OpusExtensionIterator) {
	iter.Fcurr_data = iter.Fdata
	iter.Frepeat_data = iter.Fdata
	iter.Flast_long = nil
	iter.Fcurr_len = iter.Flen1
	iter.Fcurr_frame = 0
	iter.Frepeat_frame = 0
	iter.Ftrailing_short_len = 0
}

// C documentation
//
//	/* Tell the iterator not to return any extensions for frames of index
//	    frame_max or larger.
//	   This can allow it to stop iterating early if these extensions are not
//	    needed. */
func Opus_opus_extension_iterator_set_frame_max(tls *libc.TLS, iter *OpusT_OpusExtensionIterator, frame_max int32) {
	iter.Fframe_max = frame_max
}

// C documentation
//
//	/* Return the next repeated extension.
//	   The return value is non-zero if one is found, negative on error, or 0 if we
//	    have finished repeating extensions. */
func opus_extension_iterator_next_repeat(tls *libc.TLS, iter *OpusT_OpusExtensionIterator, ext *OpusT_opus_extension_data) int32 {
	var header int32
	if iter.Frepeat_frame <= 0 {
		Opus_celt_fatal(tls, __ccgo_ts+2587, __ccgo_ts+2472, 160)
	}
	for ; iter.Frepeat_frame < iter.Fnb_frames; iter.Frepeat_frame++ {
		for iter.Fsrc_len > 0 {
			idByte := int32(*iter.Fsrc_data)
			iter.Fsrc_len = skip_extension(tls, &iter.Fsrc_data, iter.Fsrc_len, &header)
			// This source was already skipped successfully when the repeat was recorded.
			if iter.Fsrc_len < 0 {
				Opus_celt_fatal(tls, __ccgo_ts+2628, __ccgo_ts+2472, 169)
			}
			if idByte <= 3 {
				continue
			}
			if iter.Frepeat_l == 0 && iter.Frepeat_frame+1 >= iter.Fnb_frames && iter.Fsrc_data == iter.Flast_long {
				idByte &= ^int32(1)
			}
			start := iter.Fcurr_data
			iter.Fcurr_len = skip_extension_payload(tls, &iter.Fcurr_data, iter.Fcurr_len, &header, idByte, iter.Ftrailing_short_len)
			if iter.Fcurr_len < 0 {
				return OPUS_INVALID_PACKET
			}
			if extensionPointerDifference(iter.Fcurr_data, iter.Fdata) != int64(iter.Flen1-iter.Fcurr_len) {
				Opus_celt_fatal(tls, __ccgo_ts+2665, __ccgo_ts+2472, 187)
			}
			if iter.Frepeat_frame >= iter.Fframe_max {
				continue
			}
			if ext != nil {
				ext.Fid = idByte >> 1
				ext.Fframe = iter.Frepeat_frame
				ext.Fdata = (*byte)(unsafe.Add(unsafe.Pointer(start), header))
				ext.Flen1 = int32(extensionPointerDifference(iter.Fcurr_data, start) - int64(header))
			}
			return 1
		}
		iter.Fsrc_data = iter.Frepeat_data
		iter.Fsrc_len = iter.Frepeat_len
	}
	iter.Frepeat_data = iter.Fcurr_data
	iter.Flast_long = nil
	if iter.Frepeat_l == 0 {
		iter.Fcurr_frame++
		if iter.Fcurr_frame >= iter.Fnb_frames {
			iter.Fcurr_len = 0
		}
	}
	iter.Frepeat_frame = 0
	return 0
}

// C documentation
//
//	/* Return the next extension (excluding real padding, separators, and repeat
//	    indicators, but including the repeated extensions) in bitstream order.
//	   Due to the extension repetition mechanism, extensions are not necessarily
//	    returned in frame order. */
func Opus_opus_extension_iterator_next(tls *libc.TLS, iter *OpusT_OpusExtensionIterator, ext *OpusT_opus_extension_data) int32 {
	var header int32
	if iter.Fcurr_len < 0 {
		return OPUS_INVALID_PACKET
	}
	if iter.Frepeat_frame > 0 {
		if ret := opus_extension_iterator_next_repeat(tls, iter, ext); ret != 0 {
			return ret
		}
	}
	// Frame limits can change between calls, including during a repeat.
	if iter.Fcurr_frame >= iter.Fframe_max {
		return 0
	}
	for iter.Fcurr_len > 0 {
		start := iter.Fcurr_data
		id := int32(*start) >> 1
		l := int32(*start) & 1
		iter.Fcurr_len = skip_extension(tls, &iter.Fcurr_data, iter.Fcurr_len, &header)
		if iter.Fcurr_len < 0 {
			return OPUS_INVALID_PACKET
		}
		if extensionPointerDifference(iter.Fcurr_data, iter.Fdata) != int64(iter.Flen1-iter.Fcurr_len) {
			Opus_celt_fatal(tls, __ccgo_ts+2665, __ccgo_ts+2472, 255)
		}
		if id == 1 {
			if l == 0 {
				iter.Fcurr_frame++
			} else {
				increment := *(*byte)(unsafe.Add(unsafe.Pointer(start), 1))
				if increment == 0 {
					continue
				}
				iter.Fcurr_frame += int32(increment)
			}
			if iter.Fcurr_frame >= iter.Fnb_frames {
				iter.Fcurr_len = -1
				return OPUS_INVALID_PACKET
			}
			if iter.Fcurr_frame >= iter.Fframe_max {
				iter.Fcurr_len = 0
			}
			iter.Frepeat_data = iter.Fcurr_data
			iter.Flast_long = nil
			iter.Ftrailing_short_len = 0
		} else if id == 2 {
			iter.Frepeat_l = byte(l)
			iter.Frepeat_frame = iter.Fcurr_frame + 1
			iter.Frepeat_len = int32(extensionPointerDifference(start, iter.Frepeat_data))
			iter.Fsrc_data = iter.Frepeat_data
			iter.Fsrc_len = iter.Frepeat_len
			if ret := opus_extension_iterator_next_repeat(tls, iter, ext); ret != 0 {
				return ret
			}
		} else if id > 2 {
			if id >= 32 {
				iter.Flast_long = iter.Fcurr_data
				iter.Ftrailing_short_len = 0
			} else {
				iter.Ftrailing_short_len += l
			}
			if ext != nil {
				ext.Fid = id
				ext.Fframe = iter.Fcurr_frame
				ext.Fdata = (*byte)(unsafe.Add(unsafe.Pointer(start), header))
				ext.Flen1 = int32(extensionPointerDifference(iter.Fcurr_data, start) - int64(header))
			}
			return 1
		}
	}
	return 0
}

func Opus_opus_extension_iterator_find(tls *libc.TLS, iter *OpusT_OpusExtensionIterator, ext *OpusT_opus_extension_data, id int32) int32 {
	var current OpusT_opus_extension_data
	for {
		ret := Opus_opus_extension_iterator_next(tls, iter, &current)
		if ret <= 0 {
			return ret
		}
		// Do not touch the output until a matching extension has been found.
		if current.Fid == id {
			*ext = current
			return ret
		}
	}
}

// C documentation
//
//	/* Count the number of extensions, excluding real padding, separators, and
//	    repeat indicators, but including the repeated extensions. */
func Opus_opus_packet_extensions_count(tls *libc.TLS, data *byte, length, frames int32) int32 {
	var iter OpusT_OpusExtensionIterator
	Opus_opus_extension_iterator_init(tls, &iter, data, length, frames)
	var count int32
	// As in C, malformed trailing data terminates counting rather than returning an error.
	for Opus_opus_extension_iterator_next(tls, &iter, nil) > 0 {
		count++
	}
	return count
}

// C documentation
//
//	/* Count the number of extensions for each frame, excluding real padding and
//	    separators and repeat indicators, but including the repeated extensions. */
func Opus_opus_packet_extensions_count_ext(tls *libc.TLS, data *byte, length int32, frameCounts *int32, frames int32) int32 {
	var iter OpusT_OpusExtensionIterator
	var ext OpusT_opus_extension_data
	Opus_opus_extension_iterator_init(tls, &iter, data, length, frames)
	counts := unsafe.Slice(frameCounts, frames)
	// Clear before reading extensions, including when counts aliases packet bytes.
	clear(counts)
	var count int32
	for Opus_opus_extension_iterator_next(tls, &iter, &ext) > 0 {
		counts[ext.Fframe]++
		count++
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
func Opus_opus_packet_extensions_parse(tls *libc.TLS, data *byte, length int32, extensions *OpusT_opus_extension_data, nbExtensions *int32, frames int32) int32 {
	if nbExtensions == nil {
		Opus_celt_fatal(tls, __ccgo_ts+2742, __ccgo_ts+2472, 365)
	}
	if extensions == nil && *nbExtensions != 0 {
		Opus_celt_fatal(tls, __ccgo_ts+2782, __ccgo_ts+2472, 366)
	}
	var iter OpusT_OpusExtensionIterator
	var ext OpusT_opus_extension_data
	Opus_opus_extension_iterator_init(tls, &iter, data, length, frames)
	var count int32
	for {
		ret := Opus_opus_extension_iterator_next(tls, &iter, &ext)
		if ret <= 0 {
			*nbExtensions = count
			return ret
		}
		// Read capacity each time: it may alias an output field. A full buffer leaves it unchanged.
		if count == *nbExtensions {
			return -2
		}
		*(*OpusT_opus_extension_data)(unsafe.Add(unsafe.Pointer(extensions), uintptr(count)*unsafe.Sizeof(ext))) = ext
		count++
	}
}

// C documentation
//
//	/* Extract extensions from Opus padding (excluding real padding, separators,
//	    and repeat indicators, but including the repeated extensions) in frame
//	    order.
//	   nb_frame_exts must be filled with the output of
//	    opus_packet_extensions_count_ext(). */
func Opus_opus_packet_extensions_parse_ext(tls *libc.TLS, data *byte, length int32, extensions *OpusT_opus_extension_data, nbExtensions, frameCounts *int32, frames int32) int32 {
	if nbExtensions == nil {
		Opus_celt_fatal(tls, __ccgo_ts+2742, __ccgo_ts+2472, 395)
	}
	if extensions == nil && *nbExtensions != 0 {
		Opus_celt_fatal(tls, __ccgo_ts+2782, __ccgo_ts+2472, 396)
	}
	if frames > 48 {
		Opus_celt_fatal(tls, __ccgo_ts+2842, __ccgo_ts+2472, 397)
	}
	// Snapshot prefix sums before iterator initialization or any output writes.
	// Negative frame counts reach the iterator's assertion after an empty prefix loop, as in C.
	counts := unsafe.Slice(frameCounts, max(frames, 0))
	var cumulative [49]int32
	var total int32
	for i := int32(0); i < frames; i++ {
		cumulative[i] = total
		total += counts[i]
	}
	cumulative[max(frames, 0)] = total
	var iter OpusT_OpusExtensionIterator
	var ext OpusT_opus_extension_data
	Opus_opus_extension_iterator_init(tls, &iter, data, length, frames)
	var count int32
	for {
		ret := Opus_opus_extension_iterator_next(tls, &iter, &ext)
		if ret <= 0 {
			*nbExtensions = count
			return ret
		}
		idx := cumulative[ext.Fframe]
		cumulative[ext.Fframe]++
		if idx >= *nbExtensions {
			return -2
		}
		if idx >= cumulative[ext.Fframe+1] {
			Opus_celt_fatal(tls, __ccgo_ts+2876, __ccgo_ts+2472, 416)
		}
		*(*OpusT_opus_extension_data)(unsafe.Add(unsafe.Pointer(extensions), uintptr(idx)*unsafe.Sizeof(ext))) = ext
		count++
	}
}

func write_extension_payload(tls *libc.TLS, data *byte, capacity, pos, id, length int32, payload *byte, last int32) int32 {
	if id < 3 || id > 127 {
		Opus_celt_fatal(tls, __ccgo_ts+2929, __ccgo_ts+2472, 425)
	}
	if id < 32 {
		if length < 0 || length > 1 {
			return -1
		}
		if length > 0 {
			if capacity-pos < length {
				return -2
			}
			if data != nil {
				unsafe.Slice(data, capacity)[pos] = *payload
			}
			pos++
		}
		return pos
	}
	if length < 0 {
		return -1
	}
	lengthBytes := int32(1) + length/255
	if last != 0 {
		lengthBytes = 0
	}
	if capacity-pos < lengthBytes+length {
		return -2
	}
	var out []byte
	if data != nil {
		out = unsafe.Slice(data, capacity)
	}
	if last == 0 {
		for j := int32(0); j < length/255; j++ {
			if data != nil {
				out[pos] = 255
			}
			pos++
		}
		if data != nil {
			out[pos] = byte(length % 255)
		}
		pos++
	}
	if data != nil {
		copy(out[pos:pos+length], unsafe.Slice(payload, length))
	}
	return pos + length
}

func write_extension_payload_record(tls *libc.TLS, data *byte, capacity, pos int32, e *OpusT_opus_extension_data, last int32) int32 {
	if e.Fid < 3 || e.Fid > 127 {
		Opus_celt_fatal(tls, __ccgo_ts+2929, __ccgo_ts+2472, 425)
	}
	if e.Fid < 32 {
		if e.Flen1 < 0 || e.Flen1 > 1 {
			return -1
		}
		if e.Flen1 > 0 {
			if capacity-pos < e.Flen1 {
				return -2
			}
			if data != nil {
				unsafe.Slice(data, capacity)[pos] = *e.Fdata
			}
			pos++
		}
		return pos
	}
	if e.Flen1 < 0 {
		return -1
	}
	lengthBytes := int32(1) + e.Flen1/255
	if last != 0 {
		lengthBytes = 0
	}
	if capacity-pos < lengthBytes+e.Flen1 {
		return -2
	}
	if last == 0 {
		for j := int32(0); j < e.Flen1/255; j++ {
			if data != nil {
				unsafe.Slice(data, capacity)[pos] = 255
			}
			pos++
		}
		if data != nil {
			unsafe.Slice(data, capacity)[pos] = byte(e.Flen1 % 255)
		}
		pos++
	}
	if data != nil {
		copy(unsafe.Slice(data, capacity)[pos:pos+e.Flen1], unsafe.Slice(e.Fdata, e.Flen1))
	}
	return pos + e.Flen1
}

func write_extension(tls *libc.TLS, data *byte, capacity, pos, id, length int32, payload *byte, last int32) int32 {
	if capacity-pos < 1 {
		return -2
	}
	if id < 3 || id > 127 {
		Opus_celt_fatal(tls, __ccgo_ts+2929, __ccgo_ts+2472, 465)
	}
	if data != nil {
		L := length
		if id >= 32 {
			L = libc.BoolInt32(last == 0)
		}
		unsafe.Slice(data, capacity)[pos] = byte(id<<1 + L)
	}
	return write_extension_payload(tls, data, capacity, pos+1, id, length, payload, last)
}

func write_extension_record(tls *libc.TLS, data *byte, capacity, pos int32, e *OpusT_opus_extension_data, last int32) int32 {
	if capacity-pos < 1 {
		return -2
	}
	if e.Fid < 3 || e.Fid > 127 {
		Opus_celt_fatal(tls, __ccgo_ts+2929, __ccgo_ts+2472, 465)
	}
	if data != nil {
		L := e.Flen1
		if e.Fid >= 32 {
			L = libc.BoolInt32(last == 0)
		}
		unsafe.Slice(data, capacity)[pos] = byte(e.Fid<<1 + L)
	}
	// The header can alias numeric descriptor bytes; reload before payload validation.
	return write_extension_payload_record(tls, data, capacity, pos+1, e, last)
}

func Opus_opus_packet_extensions_generate(tls *libc.TLS, data *byte, len1 OpusT_opus_int32, extensions *OpusT_opus_extension_data, nb_extensions OpusT_opus_int32, nb_frames int32, pad int32) (r OpusT_opus_int32) {
	var curr_frame, diff, f, g, g1, j, j1, last, nb_repeated, repeat_count, v3 int32
	var frame_min_idx, frame_repeat_idx [48]OpusT_opus_int32
	var i, last_long_idx, padding, pos, written OpusT_opus_int32
	var frame_max_idx [48]OpusT_opus_int32
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
	exts := unsafe.Slice(extensions, max(nb_extensions, 0))
	/* Do a little work up-front to make this O(nb_extensions) instead of
	   O(nb_extensions*nb_frames) so long as the extensions are in frame
	   order (without requiring that they be in frame order). */
	f = 0
	for {
		if !(f < nb_frames) {
			break
		}
		frame_min_idx[f] = nb_extensions
		f = f + 1
	}
	i = 0
	for {
		if !(i < nb_extensions) {
			break
		}
		f = exts[i].Fframe
		if f < 0 || f >= nb_frames {
			return -int32(1)
		}
		if exts[i].Fid < int32(3) || exts[i].Fid > int32(127) {
			return -int32(1)
		}
		if frame_min_idx[f] < i {
			v3 = frame_min_idx[f]
		} else {
			v3 = i
		}
		frame_min_idx[f] = v3
		if frame_max_idx[f] > i+int32(1) {
			v3 = frame_max_idx[f]
		} else {
			v3 = i + int32(1)
		}
		frame_max_idx[f] = v3
		i = i + 1
	}
	f = 0
	for {
		if !(f < nb_frames) {
			break
		}
		frame_repeat_idx[f] = frame_min_idx[f]
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
				if !(i < frame_max_idx[f]) {
					break
				}
				if exts[i].Fframe == f {
					/* Test if we can repeat this extension in future frames. */
					g = f + int32(1)
					for {
						if !(g < nb_frames) {
							break
						}
						if frame_repeat_idx[g] >= frame_max_idx[g] {
							break
						}
						if !(exts[frame_repeat_idx[g]].Fframe == g) {
							Opus_celt_fatal(tls, __ccgo_ts+2978, __ccgo_ts+2472, int32(518))
						}
						if exts[frame_repeat_idx[g]].Fid != exts[i].Fid {
							break
						}
						if exts[frame_repeat_idx[g]].Fid < int32(32) && exts[frame_repeat_idx[g]].Flen1 != exts[i].Flen1 {
							break
						}
						g = g + 1
					}
					if g < nb_frames {
						break
					}
					/* We can! */
					/* If this is a long extension, save the index of the last
					   instance, so we can modify its L flag. */
					if exts[i].Fid >= int32(32) {
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
							if !(j < frame_max_idx[g] && exts[j].Fframe != g) {
								break
							}
							j = j + 1
						}
						frame_repeat_idx[g] = j
						g = g + 1
					}
					repeat_count = repeat_count + 1
					/* Point the repeat pointer for this frame to the current
					   extension, so we know when to trigger the repeats. */
					frame_repeat_idx[f] = i
				}
				i = i + 1
			}
		}
		i = frame_min_idx[f]
		for {
			if !(i < frame_max_idx[f]) {
				break
			}
			if exts[i].Fframe == f {
				/* Insert separator when needed. */
				if f != curr_frame {
					diff = f - curr_frame
					if len1-pos < int32(2) {
						return -int32(2)
					}
					if diff == int32(1) {
						if data != nil {
							unsafe.Slice(data, len1)[pos] = 0x02
						}
						pos++
					} else {
						if data != nil {
							unsafe.Slice(data, len1)[pos] = 0x03
						}
						pos++
						if data != nil {
							unsafe.Slice(data, len1)[pos] = byte(diff)
						}
						pos++
					}
					curr_frame = f
				}
				pos = write_extension_record(tls, data, len1, pos, &exts[i], libc.BoolInt32(written == nb_extensions-int32(1)))
				if pos < 0 {
					return pos
				}
				written = written + 1
				if repeat_count > 0 && frame_repeat_idx[f] == i {
					/* Add the repeat indicator. */
					nb_repeated = repeat_count * (nb_frames - (f + int32(1)))
					last = libc.BoolInt32(written+nb_repeated == nb_extensions || last_long_idx < 0 && i+int32(1) >= frame_max_idx[f])
					if len1-pos < int32(1) {
						return -int32(2)
					}
					if data != nil {
						unsafe.Slice(data, len1)[pos] = uint8(int32(0x04) + libc.BoolInt32(!(last != 0)))
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
							if exts[j1].Fframe == g1 {
								pos = write_extension_payload_record(tls, data, len1, pos, &exts[j1], libc.BoolInt32(last != 0 && j1 == last_long_idx))
								if pos < 0 {
									return pos
								}
								written = written + 1
							}
							j1 = j1 + 1
						}
						frame_min_idx[g1] = j1
						g1 = g1 + 1
					}
					if last != 0 {
						curr_frame = curr_frame + 1
					}
				}
			}
			i = i + 1
		}
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
		if data != nil {
			out := unsafe.Slice(data, len1)
			copy(out[padding:padding+pos], out[:pos])
			for i := int32(0); i < padding; i++ {
				out[i] = 0x01
			}
		}
		pos = pos + padding
	}
	return pos
}

const BUFSIZ = 1024
