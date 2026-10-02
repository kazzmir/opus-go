//go:build compareopus && cgo

package main

/*
#define VAR_ARRAYS 1
#define silk_decode_core comparison_decode_core
#include "decode_core.c"
static int decoder_core(unsigned char *d,int ds,unsigned char *c,int cs,short *output,const short *pulses) {
 if(ds!=sizeof(silk_decoder_state)||cs!=sizeof(silk_decoder_control))return -98;
 silk_decoder_state dec;silk_decoder_control ctrl;memcpy(&dec,d,ds);memcpy(&ctrl,c,cs);comparison_decode_core(&dec,&ctrl,output,pulses,0);memcpy(d,&dec,ds);memcpy(c,&ctrl,cs);return 0;
}
static void decoder_core_history(unsigned char *d,int memory,int count,const short *frame) {
 silk_decoder_state dec;memcpy(&dec,d,sizeof(dec));silk_memcpy(&dec.outBuf[memory],frame,count*sizeof(short));memcpy(d,&dec,sizeof(dec));
}
*/
import "C"
import (
	"github.com/kazzmir/opus-go/opuscc"
	"unsafe"
)

// These byte-image fixtures must leave all embedded pointers nil.
func nativeDecodeCore(dec *opuscc.OpusT_silk_decoder_state, ctrl *opuscc.OpusT_silk_decoder_control, output, pulses []int16) int32 {
	d, c := make([]byte, int(unsafe.Sizeof(*dec))), make([]byte, int(unsafe.Sizeof(*ctrl)))
	copy(d, unsafe.Slice((*byte)(unsafe.Pointer(dec)), len(d)))
	copy(c, unsafe.Slice((*byte)(unsafe.Pointer(ctrl)), len(c)))
	r := int32(C.decoder_core((*C.uchar)(unsafe.Pointer(&d[0])), C.int(len(d)), (*C.uchar)(unsafe.Pointer(&c[0])), C.int(len(c)), (*C.short)(unsafe.Pointer(unsafe.SliceData(output))), (*C.short)(unsafe.Pointer(unsafe.SliceData(pulses)))))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(dec)), len(d)), d)
	copy(unsafe.Slice((*byte)(unsafe.Pointer(ctrl)), len(c)), c)
	return r
}
func nativeDecodeCoreHistory(dec *opuscc.OpusT_silk_decoder_state, frame []int16) {
	d := make([]byte, int(unsafe.Sizeof(*dec)))
	copy(d, unsafe.Slice((*byte)(unsafe.Pointer(dec)), len(d)))
	C.decoder_core_history((*C.uchar)(unsafe.Pointer(&d[0])), C.int(dec.Fltp_mem_length), C.int(2*dec.Fsubfr_length), (*C.short)(unsafe.Pointer(unsafe.SliceData(frame))))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(dec)), len(d)), d)
}
