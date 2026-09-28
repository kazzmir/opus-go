//go:build compareopus && cgo

package main

/*
#include "main.h"
*/
import "C"
import "unsafe"

func nativeNLSFStabilize(values, delta []int16) {
	C.silk_NLSF_stabilize((*C.opus_int16)(unsafe.Pointer(&values[0])), (*C.opus_int16)(unsafe.Pointer(&delta[0])), C.int(len(values)))
}
