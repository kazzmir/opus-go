//go:build compareopus && cgo

package main

/*
#include "entenc.h"
static void native_entropy_buffer(unsigned char *buf,unsigned cap,unsigned size,unsigned *s) {
 ec_enc e;ec_enc_init(&e,buf,cap);
 for(unsigned i=0;i<18;i++) ec_enc_uint(&e,i%17,17);
 ec_enc_bits(&e,0xa5b,12);
 for(unsigned i=0;i<3;i++) ec_enc_bits(&e,0x1234567+i,25);
 ec_enc_shrink(&e,size);ec_enc_done(&e);
 s[0]=e.storage;s[1]=e.end_offs;s[2]=e.end_window;s[3]=e.nend_bits;s[4]=e.nbits_total;s[5]=e.offs;
 s[6]=e.rng;s[7]=e.val;s[8]=e.ext;s[9]=e.rem;s[10]=e.error;
}
*/
import "C"
import "unsafe"

func nativeEntropyBuffer(buf []byte, size uint32) [11]uint32 {
	var state [11]C.uint
	C.native_entropy_buffer((*C.uchar)(unsafe.Pointer(unsafe.SliceData(buf))), C.uint(len(buf)), C.uint(size), &state[0])
	var result [11]uint32
	for i := range result {
		result[i] = uint32(state[i])
	}
	return result
}
