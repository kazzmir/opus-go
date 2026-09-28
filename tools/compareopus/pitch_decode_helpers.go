//go:build compareopus && cgo

package main

/*
#include "main.h"
*/
import "C"
import "unsafe"

func nativePitchDecode(lag int16, contour int8, output []int32, fs int32) {
	C.silk_decode_pitch(C.opus_int16(lag), C.opus_int8(contour), (*C.int)(unsafe.Pointer(&output[0])), C.int(fs), C.int(len(output)))
}
