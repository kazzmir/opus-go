//go:build compareopus && cgo

package main

/*
#define VAR_ARRAYS 1
#define silk_resampler_private_IIR_FIR compare_IIR_FIR_driver
#include "../../../opus/silk/resampler_private_IIR_FIR.c"
static void compare_iir_driver(int *iir,int *fir,int batch,int step,short *out,const short *in,int len) {
 silk_resampler_state_struct s={0};
 memcpy(s.sIIR,iir,sizeof(s.sIIR)); memcpy(s.sFIR.i32,fir,sizeof(s.sFIR.i32));
 s.batchSize=batch; s.invRatio_Q16=step;
 compare_IIR_FIR_driver(&s,out,in,len);
 memcpy(iir,s.sIIR,sizeof(s.sIIR)); memcpy(fir,s.sFIR.i32,sizeof(s.sFIR.i32));
}
static int compare_iir_interpol(short *out,short *in,int limit,int step) {
 return silk_resampler_private_IIR_FIR_INTERPOL(out,in,limit,step)-out;
}
*/
import "C"
import (
	"github.com/kazzmir/opus-go/opuscc"
	"unsafe"
)

func nativeIIRDriver(s *opuscc.OpusT_silk_resampler_state_struct, out, in []int16) {
	C.compare_iir_driver((*C.int)(unsafe.Pointer(&s.FsIIR[0])), (*C.int)(unsafe.Pointer(&s.FsFIR.Fi32[0])), C.int(s.FbatchSize), C.int(s.FinvRatio_Q16), (*C.short)(unsafe.Pointer(unsafe.SliceData(out))), (*C.short)(unsafe.Pointer(unsafe.SliceData(in))), C.int(len(in)))
}

func nativeIIRInterpol(out, in []int16, limit, step int32) int32 {
	return int32(C.compare_iir_interpol((*C.short)(unsafe.Pointer(&out[0])), (*C.short)(unsafe.Pointer(&in[0])), C.int(limit), C.int(step)))
}
