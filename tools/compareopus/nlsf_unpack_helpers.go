//go:build compareopus && cgo

package main

/*
#include "main.h"
#include "tables.h"
static void unpack_nlsf(opus_int16 *indices, opus_uint8 *pred, int wb, int index) {
 silk_NLSF_unpack(indices, pred, wb ? &silk_NLSF_CB_WB : &silk_NLSF_CB_NB_MB, index);
}
*/
import "C"
import "unsafe"

func nativeNLSFUnpack(indices []int16, pred []uint8, wb bool, index int32) {
	var wide C.int
	if wb {
		wide = 1
	}
	C.unpack_nlsf((*C.opus_int16)(unsafe.Pointer(&indices[0])), (*C.opus_uint8)(unsafe.Pointer(&pred[0])), wide, C.int(index))
}
