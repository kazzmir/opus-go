//go:build compareopus && cgo

package main

/*
#include "main.h"
#include "tables.h"
#include <stddef.h>
static void nlsf_cb_layout(size_t *v) {v[0]=sizeof(silk_NLSF_CB_struct);v[1]=offsetof(silk_NLSF_CB_struct,CB1_NLSF_Q8);v[2]=offsetof(silk_NLSF_CB_struct,CB1_Wght_Q9);v[3]=offsetof(silk_NLSF_CB_struct,CB1_iCDF);v[4]=offsetof(silk_NLSF_CB_struct,pred_Q8);v[5]=offsetof(silk_NLSF_CB_struct,ec_sel);v[6]=offsetof(silk_NLSF_CB_struct,ec_iCDF);v[7]=offsetof(silk_NLSF_CB_struct,ec_Rates_Q5);v[8]=offsetof(silk_NLSF_CB_struct,deltaMin_Q15);}
static void unpack_nlsf(opus_int16 *indices, opus_uint8 *pred, int wb, int index) {
 silk_NLSF_unpack(indices, pred, wb ? &silk_NLSF_CB_WB : &silk_NLSF_CB_NB_MB, index);
}
*/
import "C"
import "unsafe"

func nativeNLSFCodebookLayout() [9]uint64 {
	var v [9]C.size_t
	C.nlsf_cb_layout(&v[0])
	var out [9]uint64
	for i := range out {
		out[i] = uint64(v[i])
	}
	return out
}

func nativeNLSFUnpack(indices []int16, pred []uint8, wb bool, index int32) {
	var wide C.int
	if wb {
		wide = 1
	}
	C.unpack_nlsf((*C.opus_int16)(unsafe.Pointer(&indices[0])), (*C.opus_uint8)(unsafe.Pointer(&pred[0])), wide, C.int(index))
}
