//go:build compareopus && cgo

package main

/*
#include "SigProc_FIX.h"
*/
import "C"
import "unsafe"

func nativeLPCAnalysis(out, in, b []int16) {
	C.silk_LPC_analysis_filter((*C.opus_int16)(unsafe.Pointer(&out[0])), (*C.opus_int16)(unsafe.Pointer(&in[0])), (*C.opus_int16)(unsafe.Pointer(&b[0])), C.opus_int32(len(in)), C.opus_int32(len(b)), 0)
}
