//go:build compareopus && cgo

package main

/*
#include "main.h"
*/
import "C"
import "unsafe"

func nativeGainsDequant(output []int32, indices []int8, previous *int8, conditional int32) {
	C.silk_gains_dequant((*C.opus_int32)(unsafe.Pointer(&output[0])), (*C.opus_int8)(unsafe.Pointer(&indices[0])), (*C.opus_int8)(unsafe.Pointer(previous)), C.int(conditional), C.int(len(output)))
}
