//go:build compareopus && cgo

package main

/*
#include "entcode.h"
static opus_uint32 native_tell_frac(opus_uint32 range, int nbits) {
 ec_ctx ctx = {0};
 ctx.rng = range;
 ctx.nbits_total = nbits;
 return ec_tell_frac(&ctx);
}
*/
import "C"

func nativeTellFrac(rng uint32, nbits int32) uint32 {
	return uint32(C.native_tell_frac(C.opus_uint32(rng), C.int(nbits)))
}
