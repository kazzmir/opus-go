//go:build compareopus && cgo

package main

/*
#include "SigProc_FIX.h"
*/
import "C"
import "unsafe"

func nativeInverseGain(a []int16) int32 {
	return int32(C.silk_LPC_inverse_pred_gain_c((*C.opus_int16)(unsafe.Pointer(&a[0])), C.int(len(a))))
}
