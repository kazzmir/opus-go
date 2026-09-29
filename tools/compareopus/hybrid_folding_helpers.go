//go:build compareopus && cgo

package main

/*
#include <string.h>
#include "opus_types.h"
// Scalar C reference for the static special_hybrid_folding in celt/bands.c.
// The mode is reduced to the only field that helper reads: eBands.
static void hybrid_folding(const opus_int16 *eBands, float *norm, float *norm2,
                          int start, int M, int dual_stereo) {
 int n1=M*(eBands[start+1]-eBands[start]);
 int n2=M*(eBands[start+2]-eBands[start+1]);
 memcpy(&norm[n1],&norm[2*n1-n2],(n2-n1)*sizeof(*norm));
 if (dual_stereo) memcpy(&norm2[n1],&norm2[2*n1-n2],(n2-n1)*sizeof(*norm2));
}
*/
import "C"
import "unsafe"

func nativeHybridFolding(bands []int16, norm, norm2 []float32, start, M, dual int32) {
	C.hybrid_folding((*C.opus_int16)(unsafe.Pointer(&bands[0])), (*C.float)(unsafe.Pointer(&norm[0])), (*C.float)(unsafe.Pointer(unsafe.SliceData(norm2))), C.int(start), C.int(M), C.int(dual))
}
