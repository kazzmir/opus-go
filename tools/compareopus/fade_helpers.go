//go:build compareopus && cgo

package main

/*
#include "arch.h"
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

func nativeFade(a, b, out, window []float32, n, channels, rate int32) {
	C.compare_fade((*C.float)(unsafe.Pointer(unsafe.SliceData(a))), (*C.float)(unsafe.Pointer(unsafe.SliceData(b))), (*C.float)(unsafe.Pointer(unsafe.SliceData(out))), C.int(n), C.int(channels), (*C.float)(unsafe.Pointer(unsafe.SliceData(window))), C.int(rate))
}
