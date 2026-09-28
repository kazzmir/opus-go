//go:build compareopus && cgo

package main

/*
#include "resampler_private.h"
*/
import "C"
import "unsafe"

func nativeUp2(state *[6]int32, output, input []int16) {
	C.silk_resampler_private_up2_HQ((*C.opus_int32)(unsafe.Pointer(state)), (*C.opus_int16)(unsafe.Pointer(&output[0])), (*C.opus_int16)(unsafe.Pointer(&input[0])), C.opus_int32(len(input)))
}
