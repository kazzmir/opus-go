//go:build compareopus && cgo

package main

/*
#include "main.h"
#include "tables.h"
static void decode_nlsf(opus_int16 *out, opus_int8 *indices, int wb) {
 silk_NLSF_decode(out, indices, wb ? &silk_NLSF_CB_WB : &silk_NLSF_CB_NB_MB);
}
*/
import "C"
import "unsafe"

func nativeNLSFDecode(out []int16, indices []int8, wb bool) {
	var wide C.int
	if wb {
		wide = 1
	}
	C.decode_nlsf((*C.opus_int16)(unsafe.Pointer(&out[0])), (*C.opus_int8)(unsafe.Pointer(&indices[0])), wide)
}
