//go:build compareopus && cgo

package main

/*
#cgo CFLAGS: -I${SRCDIR}/../../../opus/celt
#include "entdec.h"
static unsigned lookup(unsigned rng, unsigned val, unsigned param, int binary, unsigned *ext) {
    ec_dec dec = {0};
    dec.rng = rng;
    dec.val = val;
    unsigned symbol = binary ? ec_decode_bin(&dec, param) : ec_decode(&dec, param);
    *ext = dec.ext;
    return symbol;
}
*/
import "C"

func nativeRangeLookup(rng, val, param uint32, binary bool) (uint32, uint32) {
	var ext C.uint
	var bin C.int
	if binary {
		bin = 1
	}
	symbol := C.lookup(C.uint(rng), C.uint(val), C.uint(param), bin, &ext)
	return uint32(symbol), uint32(ext)
}
