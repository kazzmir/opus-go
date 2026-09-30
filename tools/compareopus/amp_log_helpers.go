//go:build compareopus && cgo

package main

/*
#include "modes.h"
#include "quant_bands.h"
static void native_amp_log(float *in,float *out,int bands,int eff,int end,int channels) {
 CELTMode m={0};m.nbEBands=bands;amp2Log2(&m,eff,end,in,out,channels);
}
*/
import "C"
import "unsafe"

func nativeAmpLog(in, out []float32, bands, eff, end, channels int32) {
	C.native_amp_log((*C.float)(unsafe.Pointer(unsafe.SliceData(in))), (*C.float)(unsafe.Pointer(unsafe.SliceData(out))), C.int(bands), C.int(eff), C.int(end), C.int(channels))
}
