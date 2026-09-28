//go:build compareopus && cgo

package main

/*
#include "main.h"
*/
import "C"
import "unsafe"

func nativeLPCFit(out []int16, input []int32, qout, qin int32) {
	C.silk_LPC_fit((*C.opus_int16)(unsafe.Pointer(&out[0])), (*C.opus_int32)(unsafe.Pointer(&input[0])), C.int(qout), C.int(qin), C.int(len(input)))
}
