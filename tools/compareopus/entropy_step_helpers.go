//go:build compareopus && cgo

package main

/*
#include "entdec.h"
#include "laplace.h"
#include "main.h"
// Copy fields explicitly; do not assume the translated Go/C layouts match.
static unsigned entropy_step(unsigned *s, unsigned char *data, int op,
                             unsigned a, unsigned b, unsigned c, const opus_uint16 *table, opus_int32 *out) {
 ec_dec dec = {0};
 dec.buf = data;
 dec.storage=s[0]; dec.end_offs=s[1]; dec.end_window=s[2];
 dec.nend_bits=(int)s[3]; dec.nbits_total=(int)s[4]; dec.offs=s[5];
 dec.rng=s[6]; dec.val=s[7]; dec.ext=s[8]; dec.rem=(int)s[9]; dec.error=(int)s[10];
 unsigned result=0;
 switch(op) {
 case 0: ec_dec_update(&dec,a,b,c); break;
 case 1: result=ec_dec_bit_logp(&dec,a); break;
 case 2: result=ec_dec_bits(&dec,a); break;
 case 3: result=ec_dec_uint(&dec,a); break;
 case 4: result=ec_dec_icdf16(&dec,table,a); break;
 case 5: result=ec_laplace_decode(&dec,a,(int)b); break;
 case 6: result=ec_laplace_decode_p0(&dec,(opus_uint16)a,(opus_uint16)b); break;
 case 7: silk_stereo_decode_pred(&dec,out); break;
 case 8: { int flag; silk_stereo_decode_mid_only(&dec,&flag); result=flag; break; }
 case 9: ec_dec_init(&dec,data,a); break;
 }
 s[0]=dec.storage; s[1]=dec.end_offs; s[2]=dec.end_window;
 s[3]=dec.nend_bits; s[4]=dec.nbits_total; s[5]=dec.offs;
 s[6]=dec.rng; s[7]=dec.val; s[8]=dec.ext; s[9]=dec.rem; s[10]=dec.error;
 return result;
}
*/
import "C"

import (
	"github.com/kazzmir/opus-go/opuscc"
	"unsafe"
)

func nativeEntropyStep(dec *opuscc.OpusT_ec_dec, data []byte, op int, a, b, c uint32, tables ...[]uint16) uint32 {
	value, _ := nativeEntropyStepOutput(dec, data, op, a, b, c, tables...)
	return value
}

func nativeEntropyStepOutput(dec *opuscc.OpusT_ec_dec, data []byte, op int, a, b, c uint32, tables ...[]uint16) (uint32, [2]int32) {
	var out [2]C.opus_int32
	var table *C.opus_uint16
	if len(tables) != 0 {
		table = (*C.opus_uint16)(unsafe.Pointer(unsafe.SliceData(tables[0])))
	}
	s := [11]C.uint{C.uint(dec.Fstorage), C.uint(dec.Fend_offs), C.uint(dec.Fend_window), C.uint(dec.Fnend_bits), C.uint(dec.Fnbits_total), C.uint(dec.Foffs), C.uint(dec.Frng), C.uint(dec.Fval), C.uint(dec.Fext), C.uint(dec.Frem), C.uint(dec.Ferror1)}
	result := C.entropy_step(&s[0], (*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.int(op), C.uint(a), C.uint(b), C.uint(c), table, &out[0])
	dec.Fstorage = uint32(s[0])
	dec.Fend_offs = uint32(s[1])
	dec.Fend_window = uint32(s[2])
	dec.Fnend_bits = int32(s[3])
	dec.Fnbits_total = int32(s[4])
	dec.Foffs = uint32(s[5])
	dec.Frng = uint32(s[6])
	dec.Fval = uint32(s[7])
	dec.Fext = uint32(s[8])
	dec.Frem = int32(s[9])
	dec.Ferror1 = int32(s[10])
	return uint32(result), [2]int32{int32(out[0]), int32(out[1])}
}
