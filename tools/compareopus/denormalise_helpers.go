//go:build compareopus && cgo

package main

/*
#define VAR_ARRAYS 1
#include "modes.h"
#include "bands.h"
static void denormalise(const opus_int16 *bands, int nb, int size, const float *x, float *out,
 const float *energy, int start, int end, int M, int downsample, int silence) {
 CELTMode m = {0}; m.nbEBands=nb; m.shortMdctSize=size; m.eBands=bands;
 denormalise_bands(&m,x,out,energy,start,end,M,downsample,silence);
}
*/
import "C"
import "unsafe"

func nativeDenormalise(bands []int16, size int32, x, out, energy []float32, start, end, M, downsample, silence int32) {
	C.denormalise((*C.opus_int16)(unsafe.Pointer(&bands[0])), C.int(len(bands)-1), C.int(size), (*C.float)(unsafe.Pointer(unsafe.SliceData(x))), (*C.float)(unsafe.Pointer(&out[0])), (*C.float)(unsafe.Pointer(unsafe.SliceData(energy))), C.int(start), C.int(end), C.int(M), C.int(downsample), C.int(silence))
}
