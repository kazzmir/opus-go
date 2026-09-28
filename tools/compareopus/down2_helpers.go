//go:build compareopus && cgo

package main

/*
#include <opus_types.h>
void silk_resampler_down2(opus_int32 *S, opus_int16 *out, const opus_int16 *in, opus_int32 inLen);
*/
import "C"

import "unsafe"

func nativeDown2(state *[2]int32, output, input []int16) {
	C.silk_resampler_down2((*C.opus_int32)(unsafe.Pointer(state)), (*C.opus_int16)(unsafe.Pointer(&output[0])), (*C.opus_int16)(unsafe.Pointer(&input[0])), C.opus_int32(len(input)))
}
