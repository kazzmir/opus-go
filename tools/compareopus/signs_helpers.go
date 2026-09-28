//go:build compareopus && cgo

package main

/*
#include "main.h"
static void decode_signs(unsigned char *data, unsigned size, opus_int16 *pulses, int length,
                         int signal, int offset, const int *sums, unsigned *s) {
 ec_dec dec = {0};
 ec_dec_init(&dec,data,size);
 silk_decode_signs(&dec,pulses,length,signal,offset,sums);
 s[0]=dec.storage; s[1]=dec.end_offs; s[2]=dec.end_window;
 s[3]=dec.nend_bits; s[4]=dec.nbits_total; s[5]=dec.offs;
 s[6]=dec.rng; s[7]=dec.val; s[8]=dec.ext; s[9]=dec.rem; s[10]=dec.error;
}
*/
import "C"

import (
	"github.com/kazzmir/opus-go/opuscc"
	"unsafe"
)

func nativeSignsDecode(data []byte, pulses []int16, length, signal, offset int32, sums []int32) opuscc.OpusT_ec_dec {
	var s [11]C.uint
	C.decode_signs((*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.uint(len(data)), (*C.opus_int16)(unsafe.Pointer(unsafe.SliceData(pulses))), C.int(length), C.int(signal), C.int(offset), (*C.int)(unsafe.Pointer(unsafe.SliceData(sums))), &s[0])
	return opuscc.OpusT_ec_dec{Fstorage: uint32(s[0]), Fend_offs: uint32(s[1]), Fend_window: uint32(s[2]), Fnend_bits: int32(s[3]), Fnbits_total: int32(s[4]), Foffs: uint32(s[5]), Frng: uint32(s[6]), Fval: uint32(s[7]), Fext: uint32(s[8]), Frem: int32(s[9]), Ferror1: int32(s[10])}
}
