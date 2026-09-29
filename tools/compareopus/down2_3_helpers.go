//go:build compareopus && cgo

package main

/*
#include "SigProc_FIX.h"
*/
import "C"
import "unsafe"

func nativeDownTwoThirds(s *[6]int32, out, in []int16) {
	C.silk_resampler_down2_3((*C.opus_int32)(unsafe.Pointer(&s[0])), (*C.opus_int16)(unsafe.Pointer(unsafe.SliceData(out))), (*C.opus_int16)(unsafe.Pointer(unsafe.SliceData(in))), C.opus_int32(len(in)))
}
