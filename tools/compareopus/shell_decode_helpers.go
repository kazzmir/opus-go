//go:build compareopus && cgo

package main

/*
#include "main.h"
static void decode_shell(unsigned char *data, unsigned size, int total, opus_int16 *out, unsigned *s) {
 ec_dec dec = {0};
 ec_dec_init(&dec,data,size);
 silk_shell_decoder(out,&dec,total);
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

func nativeShellDecode(data []byte, total int32) ([16]int16, opuscc.OpusT_ec_dec) {
	var out [16]int16
	var s [11]C.uint
	C.decode_shell((*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.uint(len(data)), C.int(total), (*C.opus_int16)(unsafe.Pointer(&out[0])), &s[0])
	return out, opuscc.OpusT_ec_dec{Fstorage: uint32(s[0]), Fend_offs: uint32(s[1]), Fend_window: uint32(s[2]), Fnend_bits: int32(s[3]), Fnbits_total: int32(s[4]), Foffs: uint32(s[5]), Frng: uint32(s[6]), Fval: uint32(s[7]), Fext: uint32(s[8]), Frem: int32(s[9]), Ferror1: int32(s[10])}
}
