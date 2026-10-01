//go:build compareopus && cgo

package main

/*
#define VAR_ARRAYS 1
#define OPUS_DISABLE_INTRINSICS 1
#define resampling_factor compare_comb_resampling_factor
#define comb_filter compare_comb_filter
#define init_caps compare_comb_init_caps
#define tf_select_table compare_comb_tf_select_table
#define opus_strerror compare_comb_strerror
#define opus_get_version_string compare_comb_version
#include "../../../opus/celt/celt.c"
// Scalar smooth_fade from src/opus_decoder.c, using the floating-point macros.
static void compare_fade(const float *a,const float *b,float *out,int n,int channels,const float *window,int rate) {
 int inc=48000/rate;
 for(int c=0;c<channels;c++) for(int i=0;i<n;i++) {
  celt_coef w=MULT_COEF(window[i*inc],window[i*inc]);
  out[i*channels+c]=ADD32(MULT_COEF_32(w,b[i*channels+c]),MULT_COEF_32(COEF_ONE-w,a[i*channels+c]));
 }
}
*/
import "C"
import "unsafe"

func nativeComb(y, x *float32, T0, T1, N int32, g0, g1 float32, tap0, tap1 int32, window *float32, overlap int32) {
	C.compare_comb_filter((*C.float)(unsafe.Pointer(y)), (*C.float)(unsafe.Pointer(x)), C.int(T0), C.int(T1), C.int(N), C.float(g0), C.float(g1), C.int(tap0), C.int(tap1), (*C.float)(unsafe.Pointer(window)), C.int(overlap), 0)
}

func nativeFade(a, b, out, window []float32, n, channels, rate int32) {
	C.compare_fade((*C.float)(unsafe.Pointer(unsafe.SliceData(a))), (*C.float)(unsafe.Pointer(unsafe.SliceData(b))), (*C.float)(unsafe.Pointer(unsafe.SliceData(out))), C.int(n), C.int(channels), (*C.float)(unsafe.Pointer(unsafe.SliceData(window))), C.int(rate))
}
