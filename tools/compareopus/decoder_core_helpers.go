//go:build compareopus && cgo

package main

/*
#define VAR_ARRAYS 1
#define silk_decode_core comparison_decode_core
#include "decode_core.c"
static int decoder_core(unsigned char *d,int ds,unsigned char *c,int cs,short *output,const short *pulses,int alias) {
 if(ds!=sizeof(silk_decoder_state)||cs!=sizeof(silk_decoder_control))return -98;
 silk_decoder_state dec;silk_decoder_control ctrl;memcpy(&dec,d,ds);memcpy(&ctrl,c,cs);comparison_decode_core(&dec,&ctrl,alias==1?dec.outBuf:alias==2?(short*)pulses:output,pulses,0);memcpy(d,&dec,ds);memcpy(c,&ctrl,cs);return 0;
}
static int decoder_core_residual(int excitation,int prediction) {return silk_ADD_LSHIFT32(excitation,prediction,1);}
static void decoder_core_ltp(int *history,const short *samples,int index,int memory,int lag,int gain,int scale) {
 for(int i=0;i<lag+LTP_ORDER/2;i++) {int at=index-i-1;history[at]=scale?silk_SMULWW(gain,history[at]):silk_SMULWB(gain,samples[memory-i-1]);}
}
static void decoder_core_coefficients(unsigned char *c,int order,int k,short *snapshot,short *a,short *b) {
 silk_decoder_control ctrl;memcpy(&ctrl,c,sizeof(ctrl));const short *live=ctrl.PredCoef_Q12[k>>1];silk_memcpy(snapshot,live,order*sizeof(short));memcpy(a,live,MAX_LPC_ORDER*sizeof(short));memcpy(b,&ctrl.LTPCoef_Q14[k*LTP_ORDER],LTP_ORDER*sizeof(short));
}
static int decoder_core_excitation(unsigned char *d,const short *pulses,int offset) {
 silk_decoder_state dec;memcpy(&dec,d,sizeof(dec));int seed=dec.indices.Seed;
 for(int i=0;i<dec.frame_length;i++){seed=silk_RAND(seed);dec.exc_Q14[i]=silk_LSHIFT((int)pulses[i],14);if(dec.exc_Q14[i]>0)dec.exc_Q14[i]-=QUANT_LEVEL_ADJUST_Q10<<4;else if(dec.exc_Q14[i]<0)dec.exc_Q14[i]+=QUANT_LEVEL_ADJUST_Q10<<4;dec.exc_Q14[i]+=offset<<4;if(seed<0)dec.exc_Q14[i]=-dec.exc_Q14[i];seed=silk_ADD32_ovflw(seed,pulses[i]);}memcpy(d,&dec,sizeof(dec));return seed;
}
static int decoder_core_transition(unsigned char *d,unsigned char *c,int k) {
 silk_decoder_state dec;silk_decoder_control ctrl;memcpy(&dec,d,sizeof(dec));memcpy(&ctrl,c,sizeof(ctrl));int active=dec.lossCnt && dec.prevSignalType==TYPE_VOICED && dec.indices.signalType!=TYPE_VOICED && k<MAX_NB_SUBFR/2;
 if(active){opus_int16 *b=&ctrl.LTPCoef_Q14[k*LTP_ORDER];silk_memset(b,0,LTP_ORDER*sizeof(short));b[LTP_ORDER/2]=SILK_FIX_CONST(0.25,14);ctrl.pitchL[k]=dec.lagPrev;}memcpy(d,&dec,sizeof(dec));memcpy(c,&ctrl,sizeof(ctrl));return active;
}
static void decoder_frame_history(unsigned char *d,const short *frame) {
 silk_decoder_state dec;memcpy(&dec,d,sizeof(dec));int move=dec.ltp_mem_length-dec.frame_length;silk_memmove(dec.outBuf,&dec.outBuf[dec.frame_length],move*sizeof(short));silk_memcpy(&dec.outBuf[move],frame,dec.frame_length*sizeof(short));memcpy(d,&dec,sizeof(dec));
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
	return nativeDecodeCoreAlias(dec, ctrl, output, pulses, 0)
}
func nativeDecodeCoreAlias(dec *opuscc.OpusT_silk_decoder_state, ctrl *opuscc.OpusT_silk_decoder_control, output, pulses []int16, alias int32) int32 {
	d, c := make([]byte, int(unsafe.Sizeof(*dec))), make([]byte, int(unsafe.Sizeof(*ctrl)))
	copy(d, unsafe.Slice((*byte)(unsafe.Pointer(dec)), len(d)))
	copy(c, unsafe.Slice((*byte)(unsafe.Pointer(ctrl)), len(c)))
	r := int32(C.decoder_core((*C.uchar)(unsafe.Pointer(&d[0])), C.int(len(d)), (*C.uchar)(unsafe.Pointer(&c[0])), C.int(len(c)), (*C.short)(unsafe.Pointer(unsafe.SliceData(output))), (*C.short)(unsafe.Pointer(unsafe.SliceData(pulses))), C.int(alias)))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(dec)), len(d)), d)
	copy(unsafe.Slice((*byte)(unsafe.Pointer(ctrl)), len(c)), c)
	return r
}
func nativeDecodeCoreResidual(excitation, prediction int32) int32 {
	return int32(C.decoder_core_residual(C.int(excitation), C.int(prediction)))
}
func nativeDecodeCoreLTP(history []int32, samples []int16, index, memory, lag, gain, scale int32) {
	C.decoder_core_ltp((*C.int)(unsafe.Pointer(unsafe.SliceData(history))), (*C.short)(unsafe.Pointer(unsafe.SliceData(samples))), C.int(index), C.int(memory), C.int(lag), C.int(gain), C.int(scale))
}
func nativeDecodeCoreCoefficients(ctrl *opuscc.OpusT_silk_decoder_control, order, k int32, snapshot *[16]int16) ([16]int16, [5]int16) {
	c := make([]byte, int(unsafe.Sizeof(*ctrl)))
	copy(c, unsafe.Slice((*byte)(unsafe.Pointer(ctrl)), len(c)))
	var a [16]int16
	var b [5]int16
	C.decoder_core_coefficients((*C.uchar)(unsafe.Pointer(&c[0])), C.int(order), C.int(k), (*C.short)(unsafe.Pointer(snapshot)), (*C.short)(unsafe.Pointer(&a)), (*C.short)(unsafe.Pointer(&b)))
	return a, b
}
func nativeDecodeCoreExcitation(dec *opuscc.OpusT_silk_decoder_state, pulses []int16, offset int32) int32 {
	d := make([]byte, int(unsafe.Sizeof(*dec)))
	copy(d, unsafe.Slice((*byte)(unsafe.Pointer(dec)), len(d)))
	r := int32(C.decoder_core_excitation((*C.uchar)(unsafe.Pointer(&d[0])), (*C.short)(unsafe.Pointer(unsafe.SliceData(pulses))), C.int(offset)))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(dec)), len(d)), d)
	return r
}
func nativeDecodeCoreTransition(dec *opuscc.OpusT_silk_decoder_state, ctrl *opuscc.OpusT_silk_decoder_control, k int32) bool {
	d, c := make([]byte, int(unsafe.Sizeof(*dec))), make([]byte, int(unsafe.Sizeof(*ctrl)))
	copy(d, unsafe.Slice((*byte)(unsafe.Pointer(dec)), len(d)))
	copy(c, unsafe.Slice((*byte)(unsafe.Pointer(ctrl)), len(c)))
	r := C.decoder_core_transition((*C.uchar)(unsafe.Pointer(&d[0])), (*C.uchar)(unsafe.Pointer(&c[0])), C.int(k))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(dec)), len(d)), d)
	copy(unsafe.Slice((*byte)(unsafe.Pointer(ctrl)), len(c)), c)
	return r != 0
}
func nativeDecodeFrameHistory(dec *opuscc.OpusT_silk_decoder_state, frame []int16) {
	d := make([]byte, int(unsafe.Sizeof(*dec)))
	copy(d, unsafe.Slice((*byte)(unsafe.Pointer(dec)), len(d)))
	C.decoder_frame_history((*C.uchar)(unsafe.Pointer(&d[0])), (*C.short)(unsafe.Pointer(unsafe.SliceData(frame))))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(dec)), len(d)), d)
}
func nativeDecodeCoreHistory(dec *opuscc.OpusT_silk_decoder_state, frame []int16) {
	d := make([]byte, int(unsafe.Sizeof(*dec)))
	copy(d, unsafe.Slice((*byte)(unsafe.Pointer(dec)), len(d)))
	C.decoder_core_history((*C.uchar)(unsafe.Pointer(&d[0])), C.int(dec.Fltp_mem_length), C.int(2*dec.Fsubfr_length), (*C.short)(unsafe.Pointer(unsafe.SliceData(frame))))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(dec)), len(d)), d)
}
