//go:build compareopus && cgo

package main

/*
#include "main.h"
*/
import "C"
import "unsafe"

func nativeGainsQuant(indices []int8, gains []int32, previous *int8, conditional int32) {
	C.silk_gains_quant((*C.opus_int8)(unsafe.Pointer(unsafe.SliceData(indices))), (*C.opus_int32)(unsafe.Pointer(unsafe.SliceData(gains))), (*C.opus_int8)(unsafe.Pointer(previous)), C.int(conditional), C.int(len(gains)))
}
