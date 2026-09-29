//go:build compareopus && cgo

package main

/*
#define VAR_ARRAYS 1
#define silk_resampler_private_down_FIR compare_down_FIR_driver
#include "../../../opus/silk/resampler_private_down_FIR.c"
static void compare_down_driver(int *iir,int *fir,int batch,int step,int order,int fracs,const short *coefs,short *out,const short *in,int len) {
 silk_resampler_state_struct s={0};
 memcpy(s.sIIR,iir,sizeof(s.sIIR)); memcpy(s.sFIR.i32,fir,sizeof(s.sFIR.i32));
 s.batchSize=batch; s.invRatio_Q16=step; s.FIR_Order=order; s.FIR_Fracs=fracs; s.Coefs=coefs;
 compare_down_FIR_driver(&s,out,in,len);
 memcpy(iir,s.sIIR,sizeof(s.sIIR)); memcpy(fir,s.sFIR.i32,sizeof(s.sFIR.i32));
}
static int compare_down_interpol(short *out,int *in,const short *coefs,int order,int fracs,int limit,int step) {
 return silk_resampler_private_down_FIR_INTERPOL(out,in,coefs,order,fracs,limit,step)-out;
}
*/
import "C"
import (
	"github.com/kazzmir/opus-go/opuscc"
	"unsafe"
)

func nativeDownDriver(s *opuscc.OpusT_silk_resampler_state_struct, coefs, out, in []int16) {
	C.compare_down_driver((*C.int)(unsafe.Pointer(&s.FsIIR[0])), (*C.int)(unsafe.Pointer(&s.FsFIR.Fi32[0])), C.int(s.FbatchSize), C.int(s.FinvRatio_Q16), C.int(s.FFIR_Order), C.int(s.FFIR_Fracs), (*C.short)(unsafe.Pointer(&coefs[0])), (*C.short)(unsafe.Pointer(unsafe.SliceData(out))), (*C.short)(unsafe.Pointer(unsafe.SliceData(in))), C.int(len(in)))
}

func nativeDownInterpol(out []int16, in []int32, coefs []int16, order, fracs, limit, step int32) int32 {
	return int32(C.compare_down_interpol((*C.short)(unsafe.Pointer(&out[0])), (*C.int)(unsafe.Pointer(&in[0])), (*C.short)(unsafe.Pointer(&coefs[0])), C.int(order), C.int(fracs), C.int(limit), C.int(step)))
}
