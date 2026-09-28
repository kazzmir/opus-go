//go:build compareopus && cgo

package main

/*
#include "SigProc_FIX.h"
*/
import "C"
import "unsafe"

func nativeNLSF2A(out, in []int16) {
	C.silk_NLSF2A((*C.opus_int16)(unsafe.Pointer(&out[0])), (*C.opus_int16)(unsafe.Pointer(&in[0])), C.int(len(in)), 0)
}
