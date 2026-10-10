//go:build compareopus && cgo

package main

/*
#cgo CFLAGS: -I${SRCDIR}/../../../opus/silk/float
#include "SigProc_FLP.h"
*/
import "C"
import "unsafe"

func nativeScaleFLP(data []float32, gain float32, size int32) {
	C.silk_scale_vector_FLP((*C.silk_float)(unsafe.Pointer(unsafe.SliceData(data))), C.silk_float(gain), C.opus_int(size))
}

func nativeScaleCopyFLP(out, in []float32, gain float32, size int32) {
	C.silk_scale_copy_vector_FLP((*C.silk_float)(unsafe.Pointer(unsafe.SliceData(out))), (*C.silk_float)(unsafe.Pointer(unsafe.SliceData(in))), C.silk_float(gain), C.opus_int(size))
}
