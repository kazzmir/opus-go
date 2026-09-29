//go:build compareopus && cgo

package main

/*
#define VAR_ARRAYS 1
#define silk_resampler_private_IIR_FIR compare_IIR_FIR_driver
#include "../../../opus/silk/resampler_private_IIR_FIR.c"
static int compare_iir_interpol(short *out,short *in,int limit,int step) {
 return silk_resampler_private_IIR_FIR_INTERPOL(out,in,limit,step)-out;
}
*/
import "C"
import "unsafe"

func nativeIIRInterpol(out, in []int16, limit, step int32) int32 {
	return int32(C.compare_iir_interpol((*C.short)(unsafe.Pointer(&out[0])), (*C.short)(unsafe.Pointer(&in[0])), C.int(limit), C.int(step)))
}
