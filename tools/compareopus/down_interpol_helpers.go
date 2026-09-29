//go:build compareopus && cgo

package main

/*
#define VAR_ARRAYS 1
#define silk_resampler_private_down_FIR compare_down_FIR_driver
#include "../../../opus/silk/resampler_private_down_FIR.c"
static int compare_down_interpol(short *out,int *in,const short *coefs,int order,int fracs,int limit,int step) {
 return silk_resampler_private_down_FIR_INTERPOL(out,in,coefs,order,fracs,limit,step)-out;
}
*/
import "C"
import "unsafe"

func nativeDownInterpol(out []int16, in []int32, coefs []int16, order, fracs, limit, step int32) int32 {
	return int32(C.compare_down_interpol((*C.short)(unsafe.Pointer(&out[0])), (*C.int)(unsafe.Pointer(&in[0])), (*C.short)(unsafe.Pointer(&coefs[0])), C.int(order), C.int(fracs), C.int(limit), C.int(step)))
}
