//go:build compareopus && cgo

package main

/*
#define VAR_ARRAYS 1
#define OPUS_DISABLE_INTRINSICS 1
#define exp_rotation compare_vq_rotation
#define op_pvq_search_c compare_vq_search
#define alg_quant compare_vq_quant
#define alg_unquant compare_vq_unquant
#define renormalise_vector compare_vq_renormalise
#define stereo_itheta compare_vq_itheta
#include "../../../opus/celt/vq.c"
#include "entdec.h"
static unsigned decode_alg(unsigned char *data,unsigned size,float *out,int n,int k,
 int spread,int B,float gain,unsigned *s) {
 ec_dec dec={0};ec_dec_init(&dec,data,size);
 unsigned mask=alg_unquant(out,n,k,spread,B,&dec,gain);
 s[0]=dec.storage; s[1]=dec.end_offs; s[2]=dec.end_window;
 s[3]=dec.nend_bits; s[4]=dec.nbits_total; s[5]=dec.offs;
 s[6]=dec.rng; s[7]=dec.val; s[8]=dec.ext; s[9]=dec.rem; s[10]=dec.error;
 return mask;
}
*/
import "C"

import (
	"github.com/kazzmir/opus-go/opuscc"
	"unsafe"
)

func nativePVQSearch(x []float32, pulses []int32, k, n int32) float32 {
	return float32(C.compare_vq_search((*C.float)(unsafe.Pointer(unsafe.SliceData(x))), (*C.int)(unsafe.Pointer(unsafe.SliceData(pulses))), C.int(k), C.int(n), 0))
}

func nativeAlgUnquant(data []byte, out []float32, n, k, spread, B int32, gain float32) (uint32, opuscc.OpusT_ec_dec) {
	var s [11]C.uint
	mask := C.decode_alg((*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.uint(len(data)), (*C.float)(unsafe.Pointer(&out[0])), C.int(n), C.int(k), C.int(spread), C.int(B), C.float(gain), &s[0])
	return uint32(mask), opuscc.OpusT_ec_dec{Fstorage: uint32(s[0]), Fend_offs: uint32(s[1]), Fend_window: uint32(s[2]), Fnend_bits: int32(s[3]), Fnbits_total: int32(s[4]), Foffs: uint32(s[5]), Frng: uint32(s[6]), Fval: uint32(s[7]), Fext: uint32(s[8]), Frem: int32(s[9]), Ferror1: int32(s[10])}
}
