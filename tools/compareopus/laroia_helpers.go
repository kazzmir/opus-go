//go:build compareopus && cgo

package main

/*
#include "main.h"
*/
import "C"
import "unsafe"

func nativeLaroia(output, input []int16) {
	C.silk_NLSF_VQ_weights_laroia((*C.opus_int16)(unsafe.Pointer(&output[0])), (*C.opus_int16)(unsafe.Pointer(&input[0])), C.int(len(input)))
}
