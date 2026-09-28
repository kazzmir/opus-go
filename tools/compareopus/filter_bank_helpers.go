//go:build compareopus && cgo

package main

/*
#include <opus_types.h>
void silk_ana_filt_bank_1(const opus_int16 *in, opus_int32 *S, opus_int16 *outL, opus_int16 *outH, opus_int32 N);
*/
import "C"
import "unsafe"

func nativeFilterBank(input []int16, state *[2]int32, low, high []int16) {
	C.silk_ana_filt_bank_1((*C.opus_int16)(unsafe.Pointer(&input[0])), (*C.opus_int32)(unsafe.Pointer(state)), (*C.opus_int16)(unsafe.Pointer(&low[0])), (*C.opus_int16)(unsafe.Pointer(&high[0])), C.opus_int32(len(input)))
}
