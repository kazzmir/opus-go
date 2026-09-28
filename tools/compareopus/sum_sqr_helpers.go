//go:build compareopus && cgo

package main

/*
#include <opus_types.h>
void silk_sum_sqr_shift(opus_int32 *energy, int *shift, const opus_int16 *x, int len);
*/
import "C"
import "unsafe"

func nativeSumSqr(input []int16) (int32, int32) {
	var energy C.opus_int32
	var shift C.int
	C.silk_sum_sqr_shift(&energy, &shift, (*C.opus_int16)(unsafe.Pointer(unsafe.SliceData(input))), C.int(len(input)))
	return int32(energy), int32(shift)
}
