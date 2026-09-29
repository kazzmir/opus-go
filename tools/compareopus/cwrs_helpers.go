//go:build compareopus && cgo

package main

/*
#define VAR_ARRAYS 1
// Read the same native table, keeping the linked library's symbols distinct.
#define log2_frac comparison_log2_frac
#define get_required_bits comparison_get_required_bits
#define encode_pulses comparison_encode_pulses
#define decode_pulses comparison_decode_pulses
#include "cwrs.c"
static unsigned pvq_count(int n,int k) { return CELT_PVQ_V(n,k); }
static int encode_pvq_index(unsigned char *data,unsigned size,int n,int k,unsigned index) {
 ec_enc enc; ec_enc_init(&enc,data,size);
 ec_enc_uint(&enc,index,CELT_PVQ_V(n,k)); ec_enc_done(&enc); return enc.error;
}
static float decode_pvq(unsigned char *data,unsigned size,int *out,int n,int k,unsigned *s) {
 ec_dec dec={0}; ec_dec_init(&dec,data,size);
 float energy=comparison_decode_pulses(out,n,k,&dec);
 s[0]=dec.storage; s[1]=dec.end_offs; s[2]=dec.end_window;
 s[3]=dec.nend_bits; s[4]=dec.nbits_total; s[5]=dec.offs;
 s[6]=dec.rng; s[7]=dec.val; s[8]=dec.ext; s[9]=dec.rem; s[10]=dec.error;
 return energy;
}
*/
import "C"

import (
	"github.com/kazzmir/opus-go/opuscc"
	"unsafe"
)

func nativePVQCount(n, k int32) uint32 { return uint32(C.pvq_count(C.int(n), C.int(k))) }
func nativePVQIndex(data []byte, n, k int32, index uint32) int {
	return int(C.encode_pvq_index((*C.uchar)(unsafe.Pointer(&data[0])), C.uint(len(data)), C.int(n), C.int(k), C.uint(index)))
}
func nativePVQDecode(data []byte, out []int32, n, k int32) (float32, opuscc.OpusT_ec_dec) {
	var s [11]C.uint
	energy := C.decode_pvq((*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.uint(len(data)), (*C.int)(unsafe.Pointer(&out[0])), C.int(n), C.int(k), &s[0])
	return float32(energy), opuscc.OpusT_ec_dec{Fstorage: uint32(s[0]), Fend_offs: uint32(s[1]), Fend_window: uint32(s[2]), Fnend_bits: int32(s[3]), Fnbits_total: int32(s[4]), Foffs: uint32(s[5]), Frng: uint32(s[6]), Fval: uint32(s[7]), Fext: uint32(s[8]), Frem: int32(s[9]), Ferror1: int32(s[10])}
}
