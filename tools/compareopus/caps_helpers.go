//go:build compareopus && cgo

package main

/*
#define VAR_ARRAYS 1
#include "modes.h"
#include "celt.h"
static void caps(const opus_int16 *bands, const unsigned char *cache, int *out, int nb, int LM, int C) {
 CELTMode m = {0}; m.nbEBands=nb; m.eBands=bands; m.cache.caps=cache;
 init_caps(&m,out,LM,C);
}
*/
import "C"
import "unsafe"

func nativeCaps(bands []int16, cache []uint8, out []int32, lm, channels int32) {
	C.caps((*C.opus_int16)(unsafe.Pointer(&bands[0])), (*C.uchar)(unsafe.Pointer(&cache[0])), (*C.int)(unsafe.Pointer(&out[0])), C.int(len(out)), C.int(lm), C.int(channels))
}
