//go:build compareopus && cgo

package main

/*
#include <opus_types.h>
// Internal scalar routines exported by the local static libopus build.
void silk_bwexpander(opus_int16 *ar, const int d, opus_int32 chirp_Q16);
void silk_bwexpander_32(opus_int32 *ar, const int d, opus_int32 chirp_Q16);
*/
import "C"

import "unsafe"

func nativeBWExpander(ar []int16, chirp int32) {
	C.silk_bwexpander((*C.opus_int16)(unsafe.Pointer(&ar[0])), C.int(len(ar)), C.opus_int32(chirp))
}

func nativeBWExpander32(ar []int32, chirp int32) {
	C.silk_bwexpander_32((*C.opus_int32)(unsafe.Pointer(&ar[0])), C.int(len(ar)), C.opus_int32(chirp))
}
