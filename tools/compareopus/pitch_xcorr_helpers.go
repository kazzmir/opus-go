//go:build compareopus && cgo

package main

/*
#define VAR_ARRAYS 1
#define _celt_autocorr comparison_autocorr
#define _celt_lpc comparison_lpc_unused
#define pitch_downsample compare_pitch_downsample
#define celt_pitch_xcorr_c compare_pitch_xcorr
#define pitch_search compare_pitch_search
#define remove_doubling compare_remove_doubling
#include "../../../opus/celt/pitch.c"
static void native_downsample(float *left,float *right,float *out,int n,int C,int factor) {float *channels[2]={left,right};compare_pitch_downsample(channels,out,n,C,factor,0);}
*/
import "C"
import "unsafe"

func nativePitchDownsample(left, right, out []float32, n, channels, factor int32) {
	C.native_downsample((*C.float)(unsafe.Pointer(unsafe.SliceData(left))), (*C.float)(unsafe.Pointer(unsafe.SliceData(right))), (*C.float)(unsafe.Pointer(unsafe.SliceData(out))), C.int(n), C.int(channels), C.int(factor))
}

func nativeScalarPitchXCorr(x, y, out []float32, n, lags int32) {
	C.compare_pitch_xcorr((*C.float)(unsafe.Pointer(unsafe.SliceData(x))), (*C.float)(unsafe.Pointer(unsafe.SliceData(y))), (*C.float)(unsafe.Pointer(&out[0])), C.int(n), C.int(lags), 0)
}
