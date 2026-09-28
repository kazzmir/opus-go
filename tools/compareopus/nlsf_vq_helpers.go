//go:build compareopus && cgo

package main

/*
#include "main.h"
*/
import "C"
import "unsafe"

func nativeNLSFVQ(output []int32, input []int16, codebook []uint8, weights []int16) {
	C.silk_NLSF_VQ((*C.opus_int32)(unsafe.Pointer(&output[0])), (*C.opus_int16)(unsafe.Pointer(&input[0])), (*C.opus_uint8)(unsafe.Pointer(&codebook[0])), (*C.opus_int16)(unsafe.Pointer(&weights[0])), C.int(len(output)), C.int(len(input)))
}
