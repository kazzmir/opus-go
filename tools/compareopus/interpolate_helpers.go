//go:build compareopus && cgo

package main

/*
#include <opus_types.h>
void silk_interpolate(opus_int16 *xi, const opus_int16 *x0, const opus_int16 *x1, const int ifact_Q2, const int d);
*/
import "C"

import "unsafe"

func nativeInterpolate(out, a, b []int16, factor int32) {
	C.silk_interpolate((*C.opus_int16)(unsafe.Pointer(&out[0])), (*C.opus_int16)(unsafe.Pointer(&a[0])), (*C.opus_int16)(unsafe.Pointer(&b[0])), C.int(factor), C.int(len(out)))
}
