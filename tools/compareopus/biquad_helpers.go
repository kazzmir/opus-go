//go:build compareopus && cgo

package main

/*
#include <opus_types.h>
void silk_biquad_alt_stride1(const opus_int16 *in, const opus_int32 *B_Q28,
    const opus_int32 *A_Q28, opus_int32 *S, opus_int16 *out, const opus_int32 len);
void silk_biquad_alt_stride2_c(const opus_int16 *in, const opus_int32 *B_Q28,
    const opus_int32 *A_Q28, opus_int32 *S, opus_int16 *out, const opus_int32 len);
*/
import "C"

import "unsafe"

func nativeBiquad1(input, output []int16, b *[3]int32, a *[2]int32, state *[2]int32) {
	C.silk_biquad_alt_stride1((*C.opus_int16)(unsafe.Pointer(&input[0])),
		(*C.opus_int32)(unsafe.Pointer(b)), (*C.opus_int32)(unsafe.Pointer(a)),
		(*C.opus_int32)(unsafe.Pointer(state)), (*C.opus_int16)(unsafe.Pointer(&output[0])), C.opus_int32(len(input)))
}

func nativeBiquad2(input, output []int16, b *[3]int32, a *[2]int32, state *[4]int32) {
	C.silk_biquad_alt_stride2_c((*C.opus_int16)(unsafe.Pointer(&input[0])),
		(*C.opus_int32)(unsafe.Pointer(b)), (*C.opus_int32)(unsafe.Pointer(a)),
		(*C.opus_int32)(unsafe.Pointer(state)), (*C.opus_int16)(unsafe.Pointer(&output[0])), C.opus_int32(len(input)/2))
}
