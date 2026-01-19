// Code generated for linux/amd64 by 'ccgo --package-name opusccenc --prefix-external Opus_ --prefix-typename OpusT_ -o opusccenc/libopus_enc.go -I .. -I ../include -I ../src -I ../celt -I ../silk -I ../silk/float -include config_ccgo.h -DOPUS_BUILD -DOPUS_DISABLE_INTRINSICS -DNONTHREADSAFE_PSEUDOSTACK -UVAR_ARRAYS -UUSE_ALLOCA -U__SSE__ -U__SSE2__ -U__SSE3__ -U__SSSE3__ -U__AVX__ -U__AVX2__ -std=c99 -O2 -fno-builtin -ignore-asm-errors -ignore-vector-functions ccgo_support/math_lrint.c ../src/opus.c ../src/opus_decoder.c ../src/opus_encoder.c ../src/extensions.c ../src/opus_multistream.c ../src/opus_multistream_encoder.c ../src/opus_multistream_decoder.c ../src/repacketizer.c ../src/opus_projection_encoder.c ../src/opus_projection_decoder.c ../src/mapping_matrix.c ../src/analysis.c ../src/mlp.c ../src/mlp_data.c ../celt/bands.c ../celt/celt.c ../celt/celt_encoder.c ../celt/celt_decoder.c ../celt/cwrs.c ../celt/entcode.c ../celt/entdec.c ../celt/entenc.c ../celt/kiss_fft.c ../celt/laplace.c ../celt/mathops.c ../celt/mdct.c ../celt/modes.c ../celt/pitch.c ../celt/celt_lpc.c ../celt/quant_bands.c ../celt/rate.c ../celt/vq.c ../celt/mini_kfft.c ../silk/CNG.c ../silk/code_signs.c ../silk/init_decoder.c ../silk/decode_core.c ../silk/decode_frame.c ../silk/decode_parameters.c ../silk/decode_indices.c ../silk/decode_pulses.c ../silk/decoder_set_fs.c ../silk/dec_API.c ../silk/enc_API.c ../silk/encode_indices.c ../silk/encode_pulses.c ../silk/gain_quant.c ../silk/interpolate.c ../silk/LP_variable_cutoff.c ../silk/NLSF_decode.c ../silk/NSQ.c ../silk/NSQ_del_dec.c ../silk/PLC.c ../silk/shell_coder.c ../silk/tables_gain.c ../silk/tables_LTP.c ../silk/tables_NLSF_CB_NB_MB.c ../silk/tables_NLSF_CB_WB.c ../silk/tables_other.c ../silk/tables_pitch_lag.c ../silk/tables_pulses_per_block.c ../silk/VAD.c ../silk/control_audio_bandwidth.c ../silk/quant_LTP_gains.c ../silk/VQ_WMat_EC.c ../silk/HP_variable_cutoff.c ../silk/NLSF_encode.c ../silk/NLSF_VQ.c ../silk/NLSF_unpack.c ../silk/NLSF_del_dec_quant.c ../silk/process_NLSFs.c ../silk/stereo_LR_to_MS.c ../silk/stereo_MS_to_LR.c ../silk/check_control_input.c ../silk/control_SNR.c ../silk/init_encoder.c ../silk/control_codec.c ../silk/A2NLSF.c ../silk/ana_filt_bank_1.c ../silk/biquad_alt.c ../silk/bwexpander_32.c ../silk/bwexpander.c ../silk/debug.c ../silk/decode_pitch.c ../silk/inner_prod_aligned.c ../silk/lin2log.c ../silk/log2lin.c ../silk/LPC_analysis_filter.c ../silk/LPC_inv_pred_gain.c ../silk/table_LSF_cos.c ../silk/NLSF2A.c ../silk/NLSF_stabilize.c ../silk/NLSF_VQ_weights_laroia.c ../silk/pitch_est_tables.c ../silk/resampler.c ../silk/resampler_down2_3.c ../silk/resampler_down2.c ../silk/resampler_private_AR2.c ../silk/resampler_private_down_FIR.c ../silk/resampler_private_IIR_FIR.c ../silk/resampler_private_up2_HQ.c ../silk/resampler_rom.c ../silk/sigm_Q15.c ../silk/sort.c ../silk/sum_sqr_shift.c ../silk/stereo_decode_pred.c ../silk/stereo_encode_pred.c ../silk/stereo_find_predictor.c ../silk/stereo_quant_pred.c ../silk/LPC_fit.c ../silk/float/apply_sine_window_FLP.c ../silk/float/corrMatrix_FLP.c ../silk/float/encode_frame_FLP.c ../silk/float/find_LPC_FLP.c ../silk/float/find_LTP_FLP.c ../silk/float/find_pitch_lags_FLP.c ../silk/float/find_pred_coefs_FLP.c ../silk/float/LPC_analysis_filter_FLP.c ../silk/float/LTP_analysis_filter_FLP.c ../silk/float/LTP_scale_ctrl_FLP.c ../silk/float/noise_shape_analysis_FLP.c ../silk/float/process_gains_FLP.c ../silk/float/regularize_correlations_FLP.c ../silk/float/residual_energy_FLP.c ../silk/float/warped_autocorrelation_FLP.c ../silk/float/wrappers_FLP.c ../silk/float/autocorrelation_FLP.c ../silk/float/burg_modified_FLP.c ../silk/float/bwexpander_FLP.c ../silk/float/energy_FLP.c ../silk/float/inner_product_FLP.c ../silk/float/k2a_FLP.c ../silk/float/LPC_inv_pred_gain_FLP.c ../silk/float/pitch_analysis_core_FLP.c ../silk/float/scale_copy_vector_FLP.c ../silk/float/scale_vector_FLP.c ../silk/float/schur_FLP.c ../silk/float/sort_FLP.c', DO NOT EDIT.

package opusccenc

import (
	"reflect"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

var _ reflect.Type
var _ unsafe.Pointer
var _ = libc.NewTLS

const CCGO_TLS_PSEUDOSTACK = 1
const CPU_INFO_BY_ASM = 1
const DISABLE_DEBUG_FLOAT = 1
const ENABLE_HARDENING = 1
const ENABLE_RES24 = 1
const FLOAT_APPROX = 1
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
const INT16_MAX = 32767
const INT32_MAX = 2147483647
const INT8_MAX = 127
const INTPTR_MAX = 9223372036854775807
const INT_FAST16_MAX = 9223372036854775807
const INT_FAST32_MAX = 9223372036854775807
const INT_FAST8_MAX = 127
const INT_LEAST16_MAX = 32767
const INT_LEAST32_MAX = 2147483647
const INT_LEAST8_MAX = 127
const LT_OBJDIR = ".libs/"
const NONTHREADSAFE_PSEUDOSTACK = 1
const OPUS_DISABLE_INTRINSICS = 1
const OPUS_X86_PRESUME_SSE = 1
const OPUS_X86_PRESUME_SSE2 = 1
const PACKAGE_BUGREPORT = "opus@xiph.org"
const PACKAGE_NAME = "opus"
const PACKAGE_STRING = "opus 1.6.1"
const PACKAGE_TARNAME = "opus"
const PACKAGE_URL = ""
const PACKAGE_VERSION = "1.6.1"
const PTRDIFF_MAX = 9223372036854775807
const SIG_ATOMIC_MAX = 2147483647
const SIZE_MAX = 18446744073709551615
const STDC_HEADERS = 1
const UINT16_MAX = 65535
const UINT32_MAX = 4294967295
const UINT8_MAX = 255
const UINTPTR_MAX = 18446744073709551615
const UINT_FAST16_MAX = 18446744073709551615
const UINT_FAST32_MAX = 18446744073709551615
const UINT_FAST8_MAX = 255
const UINT_LEAST16_MAX = 65535
const UINT_LEAST32_MAX = 4294967295
const UINT_LEAST8_MAX = 255
const WCHAR_MAX = "__WCHAR_MAX"
const WCHAR_MIN = "__WCHAR_MIN"
const WINT_MAX = "4294967295u"
const WINT_MIN = "0u"
const _BITS_STDINT_INTN_H = 1
const _BITS_STDINT_LEAST_H = 1
const _BITS_STDINT_UINTN_H = 1
const _BITS_TIME64_H = 1
const _BITS_TYPESIZES_H = 1
const _BITS_TYPES_H = 1
const _BITS_WCHAR_H = 1
const _FEATURES_H = 1
const _FORTIFY_SOURCE = 3
const _LP64 = 1
const _STDC_PREDEF_H = 1
const _STDINT_H = 1
const _SYS_CDEFS_H = 1
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
const __BLKCNT64_T_TYPE = "__SQUAD_TYPE"
const __BLKCNT_T_TYPE = "__SYSCALL_SLONG_TYPE"
const __BLKSIZE_T_TYPE = "__SYSCALL_SLONG_TYPE"
const __BYTE_ORDER__ = "__ORDER_LITTLE_ENDIAN__"
const __CCGO__ = 1
const __CET__ = 3
const __CHAR_BIT__ = 8
const __CLOCKID_T_TYPE = "__S32_TYPE"
const __CLOCK_T_TYPE = "__SYSCALL_SLONG_TYPE"
const __CPU_MASK_TYPE = "__SYSCALL_ULONG_TYPE"
const __DADDR_T_TYPE = "__S32_TYPE"
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
const __DEV_T_TYPE = "__UQUAD_TYPE"
const __ELF__ = 1
const __FD_SETSIZE = 1024
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
const __FSBLKCNT64_T_TYPE = "__UQUAD_TYPE"
const __FSBLKCNT_T_TYPE = "__SYSCALL_ULONG_TYPE"
const __FSFILCNT64_T_TYPE = "__UQUAD_TYPE"
const __FSFILCNT_T_TYPE = "__SYSCALL_ULONG_TYPE"
const __FSWORD_T_TYPE = "__SYSCALL_SLONG_TYPE"
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
const __GID_T_TYPE = "__U32_TYPE"
const __GLIBC_MINOR__ = 39
const __GLIBC__ = 2
const __GNUC_EXECUTION_CHARSET_NAME = "UTF-8"
const __GNUC_MINOR__ = 3
const __GNUC_PATCHLEVEL__ = 0
const __GNUC_STDC_INLINE__ = 1
const __GNUC_WIDE_EXECUTION_CHARSET_NAME = "UTF-32LE"
const __GNUC__ = 13
const __GXX_ABI_VERSION = 1018
const __HAVE_GENERIC_SELECTION = 1
const __HAVE_SPECULATION_SAFE_VALUE = 1
const __ID_T_TYPE = "__U32_TYPE"
const __INO64_T_TYPE = "__UQUAD_TYPE"
const __INO_T_MATCHES_INO64_T = 1
const __INO_T_TYPE = "__SYSCALL_ULONG_TYPE"
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
const __KERNEL_OLD_TIMEVAL_MATCHES_TIMEVAL64 = 1
const __KEY_T_TYPE = "__S32_TYPE"
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
const __LDOUBLE_REDIRECTS_TO_FLOAT128_ABI = 0
const __LONG_DOUBLE_64__ = 1
const __LONG_LONG_MAX__ = 0x7fffffffffffffff
const __LONG_LONG_WIDTH__ = 64
const __LONG_MAX__ = 0x7fffffffffffffff
const __LONG_WIDTH__ = 64
const __LP64__ = 1
const __MMX_WITH_SSE__ = 1
const __MMX__ = 1
const __MODE_T_TYPE = "__U32_TYPE"
const __NLINK_T_TYPE = "__SYSCALL_ULONG_TYPE"
const __OFF64_T_TYPE = "__SQUAD_TYPE"
const __OFF_T_MATCHES_OFF64_T = 1
const __OFF_T_TYPE = "__SYSCALL_SLONG_TYPE"
const __OPTIMIZE__ = 1
const __ORDER_BIG_ENDIAN__ = 4321
const __ORDER_LITTLE_ENDIAN__ = 1234
const __ORDER_PDP_ENDIAN__ = 3412
const __PIC__ = 2
const __PID_T_TYPE = "__S32_TYPE"
const __PIE__ = 2
const __PRAGMA_REDEFINE_EXTNAME = 1
const __PRETTY_FUNCTION__ = "__func__"
const __PTRDIFF_MAX__ = 0x7fffffffffffffff
const __PTRDIFF_WIDTH__ = 64
const __REDIRECT_FORTIFY = "__REDIRECT"
const __REDIRECT_FORTIFY_NTH = "__REDIRECT_NTH"
const __RLIM64_T_TYPE = "__UQUAD_TYPE"
const __RLIM_T_MATCHES_RLIM64_T = 1
const __RLIM_T_TYPE = "__SYSCALL_ULONG_TYPE"
const __S32_TYPE = "int"
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
const __SLONG32_TYPE = "int"
const __SSE2_MATH__ = 1
const __SSE_MATH__ = 1
const __SSIZE_T_TYPE = "__SWORD_TYPE"
const __SSP_STRONG__ = 3
const __STATFS_MATCHES_STATFS64 = 1
const __STDC_HOSTED__ = 1
const __STDC_IEC_559_COMPLEX__ = 1
const __STDC_IEC_559__ = 1
const __STDC_IEC_60559_BFP__ = 201404
const __STDC_IEC_60559_COMPLEX__ = 201404
const __STDC_ISO_10646__ = 201706
const __STDC_VERSION__ = 199901
const __STDC__ = 1
const __STRICT_ANSI__ = 1
const __SUSECONDS64_T_TYPE = "__SQUAD_TYPE"
const __SUSECONDS_T_TYPE = "__SYSCALL_SLONG_TYPE"
const __SYSCALL_SLONG_TYPE = "__SLONGWORD_TYPE"
const __SYSCALL_ULONG_TYPE = "__ULONGWORD_TYPE"
const __SYSCALL_WORDSIZE = 64
const __TIME64_T_TYPE = "__TIME_T_TYPE"
const __TIMESIZE = "__WORDSIZE"
const __TIME_T_TYPE = "__SYSCALL_SLONG_TYPE"
const __UID_T_TYPE = "__U32_TYPE"
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
const __USECONDS_T_TYPE = "__U32_TYPE"
const __USE_EXTERN_INLINES = 1
const __VERSION__ = "13.3.0"
const __WCHAR_MAX = "__WCHAR_MAX__"
const __WCHAR_MAX__ = 0x7fffffff
const __WCHAR_MIN = "__WCHAR_MIN__"
const __WCHAR_TYPE__ = "int"
const __WCHAR_WIDTH__ = 32
const __WINT_MAX__ = 0xffffffff
const __WINT_MIN__ = 0
const __WINT_WIDTH__ = 32
const __WORDSIZE = 64
const __WORDSIZE_TIME64_COMPAT32 = 1
const __amd64 = 1
const __amd64__ = 1
const __code_model_small__ = 1
const __glibc_c99_flexarr_available = 1
const __gnu_linux__ = 1
const __k8 = 1
const __k8__ = 1
const __linux = 1
const __linux__ = 1
const __pic__ = 2
const __pie__ = 2
const __unix = 1
const __unix__ = 1
const __wur = "__attribute_warn_unused_result__"
const __x86_64 = 1
const __x86_64__ = 1
const restrict = "__restrict"

type OpusT___builtin_va_list = uintptr

type OpusT___predefined_size_t = uint64

type OpusT___predefined_wchar_t = int32

type OpusT___predefined_ptrdiff_t = int64

type OpusT___u_char = uint8

type OpusT___u_short = uint16

type OpusT___u_int = uint32

type OpusT___u_long = uint64

type OpusT___int8_t = int8

type OpusT___uint8_t = uint8

type OpusT___int16_t = int16

type OpusT___uint16_t = uint16

type OpusT___int32_t = int32

type OpusT___uint32_t = uint32

type OpusT___int64_t = int64

type OpusT___uint64_t = uint64

type OpusT___int_least8_t = int8

type OpusT___uint_least8_t = uint8

type OpusT___int_least16_t = int16

type OpusT___uint_least16_t = uint16

type OpusT___int_least32_t = int32

type OpusT___uint_least32_t = uint32

type OpusT___int_least64_t = int64

type OpusT___uint_least64_t = uint64

type OpusT___quad_t = int64

type OpusT___u_quad_t = uint64

type OpusT___intmax_t = int64

type OpusT___uintmax_t = uint64

type OpusT___dev_t = uint64

type OpusT___uid_t = uint32

type OpusT___gid_t = uint32

type OpusT___ino_t = uint64

type OpusT___ino64_t = uint64

type OpusT___mode_t = uint32

type OpusT___nlink_t = uint64

type OpusT___off_t = int64

type OpusT___off64_t = int64

type OpusT___pid_t = int32

type OpusT___fsid_t = struct {
	F__val [2]int32
}

type OpusT___clock_t = int64

type OpusT___rlim_t = uint64

type OpusT___rlim64_t = uint64

type OpusT___id_t = uint32

type OpusT___time_t = int64

type OpusT___useconds_t = uint32

type OpusT___suseconds_t = int64

type OpusT___suseconds64_t = int64

type OpusT___daddr_t = int32

type OpusT___key_t = int32

type OpusT___clockid_t = int32

type OpusT___timer_t = uintptr

type OpusT___blksize_t = int64

type OpusT___blkcnt_t = int64

type OpusT___blkcnt64_t = int64

type OpusT___fsblkcnt_t = uint64

type OpusT___fsblkcnt64_t = uint64

type OpusT___fsfilcnt_t = uint64

type OpusT___fsfilcnt64_t = uint64

type OpusT___fsword_t = int64

type OpusT___ssize_t = int64

type OpusT___syscall_slong_t = int64

type OpusT___syscall_ulong_t = uint64

type OpusT___loff_t = int64

type OpusT___caddr_t = uintptr

type OpusT___intptr_t = int64

type OpusT___socklen_t = uint32

type OpusT___sig_atomic_t = int32

type OpusT_int8_t = int8

type OpusT_int16_t = int16

type OpusT_int32_t = int32

type OpusT_int64_t = int64

type OpusT_uint8_t = uint8

type OpusT_uint16_t = uint16

type OpusT_uint32_t = uint32

type OpusT_uint64_t = uint64

type OpusT_int_least8_t = int8

type OpusT_int_least16_t = int16

type OpusT_int_least32_t = int32

type OpusT_int_least64_t = int64

type OpusT_uint_least8_t = uint8

type OpusT_uint_least16_t = uint16

type OpusT_uint_least32_t = uint32

type OpusT_uint_least64_t = uint64

type OpusT_int_fast8_t = int8

type OpusT_int_fast16_t = int64

type OpusT_int_fast32_t = int64

type OpusT_int_fast64_t = int64

type OpusT_uint_fast8_t = uint8

type OpusT_uint_fast16_t = uint64

type OpusT_uint_fast32_t = uint64

type OpusT_uint_fast64_t = uint64

type OpusT_intptr_t = int64

type OpusT_uintptr_t = uint64

type OpusT_intmax_t = int64

type OpusT_uintmax_t = uint64

/* Limits of integral types.  */

/* Minimum of signed integral types.  */
/* Maximum of signed integral types.  */

/* Maximum of unsigned integral types.  */

/* Minimum of signed integral types having a minimum size.  */
/* Maximum of signed integral types having a minimum size.  */

/* Maximum of unsigned integral types having a minimum size.  */

/* Minimum of fast signed integral types having a minimum size.  */
/* Maximum of fast signed integral types having a minimum size.  */

/* Maximum of fast unsigned integral types having a minimum size.  */

/* Values to test for integral types holding `void *' pointer.  */

/* Minimum for largest signed integral type.  */
/* Maximum for largest signed integral type.  */

/* Maximum for largest unsigned integral type.  */

/* Limits of other integer types.  */

/* Limits of `ptrdiff_t' type.  */

/* Limits of `sig_atomic_t'.  */

/* Limit of `size_t' type.  */

/* Limits of `wchar_t'.  */
/* These constants might also be defined in <wchar.h>.  */

/* Limits of `wint_t'.  */

/* Signed.  */

/* Unsigned.  */

/* Maximal type.  */

/*
 * Minimal implementations to satisfy ccgo when libopus expects C99 lrint/lrintf.
 *
 * We implement round-to-nearest with ties-to-even (the typical default rounding
 * mode used by lrint/lrintf) without relying on libm.
 */
