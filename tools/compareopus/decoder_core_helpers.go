//go:build compareopus && cgo

package main

/*
#define VAR_ARRAYS 1
#define silk_decode_core comparison_decode_core
#include "decode_core.c"
#define silk_decode_frame comparison_decode_frame
#include "decode_frame.c"
static void decoder_api_lbrr(unsigned *s,unsigned char *buf,int *flags,int frames,int flag) {
 ec_dec ec={0};ec.buf=buf;ec.storage=s[0];ec.end_offs=s[1];ec.end_window=s[2];ec.nend_bits=s[3];ec.nbits_total=s[4];ec.offs=s[5];ec.rng=s[6];ec.val=s[7];ec.ext=s[8];ec.rem=s[9];ec.error=s[10];memset(flags,0,3*sizeof(int));if(flag){if(frames==1)flags[0]=1;else {int symbol=ec_dec_icdf(&ec,silk_LBRR_flags_iCDF_ptr[frames-2],8)+1;for(int i=0;i<frames;i++)flags[i]=(symbol>>i)&1;}}
 s[0]=ec.storage;s[1]=ec.end_offs;s[2]=ec.end_window;s[3]=ec.nend_bits;s[4]=ec.nbits_total;s[5]=ec.offs;s[6]=ec.rng;s[7]=ec.val;s[8]=ec.ext;s[9]=ec.rem;s[10]=ec.error;
}
static int decoder_api_resampler(unsigned char *dst,const unsigned char *src,int size,short *pred,short *side,int api,int internal,int oldapi,int oldinternal) {
 if(size!=sizeof(silk_resampler_state_struct))return -98;
 if(api==2&&internal==2&&(oldapi==1||oldinternal==1)){memset(pred,0,2*sizeof(short));memset(side,0,2*sizeof(short));memcpy(dst,src,sizeof(silk_resampler_state_struct));}return 0;
}
static void decoder_api_packet(int *frames,int channels,int flag) {if(flag)for(int n=0;n<channels;n++)frames[n]=0;}
static int decoder_frame(unsigned char *d,int ds,short *output,int *count,int lost,int cond,unsigned *s,unsigned char *buf) {
 if(ds!=sizeof(silk_decoder_state))return -98;
 silk_decoder_state dec,bindings;memcpy(&dec,d,ds);silk_init_decoder(&bindings);bindings.nb_subfr=dec.nb_subfr;silk_decoder_set_fs(&bindings,dec.fs_kHz,dec.fs_kHz*1000);
 dec.psNLSF_CB=bindings.psNLSF_CB;dec.pitch_lag_low_bits_iCDF=bindings.pitch_lag_low_bits_iCDF;dec.pitch_contour_iCDF=bindings.pitch_contour_iCDF;dec.resampler_state.Coefs=bindings.resampler_state.Coefs;
 ec_dec ec={0};ec.buf=buf;ec.storage=s[0];ec.end_offs=s[1];ec.end_window=s[2];ec.nend_bits=s[3];ec.nbits_total=s[4];ec.offs=s[5];ec.rng=s[6];ec.val=s[7];ec.ext=s[8];ec.rem=s[9];ec.error=s[10];
 int ret=comparison_decode_frame(&dec,&ec,output,count,lost,cond,0);
 s[0]=ec.storage;s[1]=ec.end_offs;s[2]=ec.end_window;s[3]=ec.nend_bits;s[4]=ec.nbits_total;s[5]=ec.offs;s[6]=ec.rng;s[7]=ec.val;s[8]=ec.ext;s[9]=ec.rem;s[10]=ec.error;
 dec.psNLSF_CB=0;dec.pitch_lag_low_bits_iCDF=0;dec.pitch_contour_iCDF=0;dec.resampler_state.Coefs=0;memcpy(d,&dec,ds);return ret;
}
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
static void decoder_frame_finish(unsigned char *d,unsigned char *c,int *count,int length,int alias) {silk_decoder_state dec;silk_decoder_control ctrl;memcpy(&dec,d,sizeof(dec));memcpy(&ctrl,c,sizeof(ctrl));int *output=alias==1?&dec.lagPrev:alias==2?&ctrl.pitchL[dec.nb_subfr-1]:count;dec.lagPrev=ctrl.pitchL[dec.nb_subfr-1];*output=length;*count=*output;memcpy(d,&dec,sizeof(dec));memcpy(c,&ctrl,sizeof(ctrl));}
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
func nativeDecodeAPILBRR(dec *opuscc.OpusT_silk_decoder_state, ec *opuscc.OpusT_ec_ctx, buf []byte) {
	s := [11]C.uint{C.uint(ec.Fstorage), C.uint(ec.Fend_offs), C.uint(ec.Fend_window), C.uint(ec.Fnend_bits), C.uint(ec.Fnbits_total), C.uint(ec.Foffs), C.uint(ec.Frng), C.uint(ec.Fval), C.uint(ec.Fext), C.uint(ec.Frem), C.uint(ec.Ferror1)}
	C.decoder_api_lbrr(&s[0], (*C.uchar)(unsafe.Pointer(unsafe.SliceData(buf))), (*C.int)(unsafe.Pointer(&dec.FLBRR_flags[0])), C.int(dec.FnFramesPerPacket), C.int(dec.FLBRR_flag))
	ec.Fstorage = uint32(s[0])
	ec.Fend_offs = uint32(s[1])
	ec.Fend_window = uint32(s[2])
	ec.Fnend_bits = int32(s[3])
	ec.Fnbits_total = int32(s[4])
	ec.Foffs = uint32(s[5])
	ec.Frng = uint32(s[6])
	ec.Fval = uint32(s[7])
	ec.Fext = uint32(s[8])
	ec.Frem = int32(s[9])
	ec.Ferror1 = int32(s[10])
}
func nativeDecodeAPIStartStereo(dec *opuscc.OpusT_silk_decoder, control *opuscc.OpusT_silk_DecControlStruct) int32 {
	src, dst := dec.Fchannel_state[0].Fresampler_state, dec.Fchannel_state[1].Fresampler_state
	src.FCoefs = nil
	dst.FCoefs = nil
	size := int(unsafe.Sizeof(src))
	s, d := make([]byte, size), make([]byte, size)
	copy(s, unsafe.Slice((*byte)(unsafe.Pointer(&src)), size))
	copy(d, unsafe.Slice((*byte)(unsafe.Pointer(&dst)), size))
	r := int32(C.decoder_api_resampler((*C.uchar)(unsafe.Pointer(&d[0])), (*C.uchar)(unsafe.Pointer(&s[0])), C.int(size), (*C.short)(unsafe.Pointer(&dec.FsStereo.Fpred_prev_Q13[0])), (*C.short)(unsafe.Pointer(&dec.FsStereo.FsSide[0])), C.int(control.FnChannelsAPI), C.int(control.FnChannelsInternal), C.int(dec.FnChannelsAPI), C.int(dec.FnChannelsInternal)))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&dst)), size), d)
	dst.FCoefs = dec.Fchannel_state[1].Fresampler_state.FCoefs
	if control.FnChannelsAPI == 2 && control.FnChannelsInternal == 2 && (dec.FnChannelsAPI == 1 || dec.FnChannelsInternal == 1) {
		dst.FCoefs = dec.Fchannel_state[0].Fresampler_state.FCoefs
	}
	dec.Fchannel_state[1].Fresampler_state = dst
	return r
}
func nativeDecodeAPIPacketStart(frames *[2]int32, channels, flag int32) {
	C.decoder_api_packet((*C.int)(unsafe.Pointer(frames)), C.int(channels), C.int(flag))
}
func nativeDecodeFrame(dec *opuscc.OpusT_silk_decoder_state, ec *opuscc.OpusT_ec_ctx, buf []byte, output []int16, count *int32, lost, cond int32) int32 {
	// Export numeric images only. Native table pointers are rebound on the C stack
	// and stripped before return; embedded Go pointers never cross the byte bridge.
	numeric := *dec
	numeric.FpsNLSF_CB = nil
	numeric.Fpitch_lag_low_bits_iCDF = nil
	numeric.Fpitch_contour_iCDF = nil
	numeric.Fresampler_state.FCoefs = nil
	d := make([]byte, int(unsafe.Sizeof(numeric)))
	copy(d, unsafe.Slice((*byte)(unsafe.Pointer(&numeric)), len(d)))
	s := [11]C.uint{}
	if ec != nil {
		s = [11]C.uint{C.uint(ec.Fstorage), C.uint(ec.Fend_offs), C.uint(ec.Fend_window), C.uint(ec.Fnend_bits), C.uint(ec.Fnbits_total), C.uint(ec.Foffs), C.uint(ec.Frng), C.uint(ec.Fval), C.uint(ec.Fext), C.uint(ec.Frem), C.uint(ec.Ferror1)}
	}
	ret := int32(C.decoder_frame((*C.uchar)(unsafe.Pointer(&d[0])), C.int(len(d)), (*C.short)(unsafe.Pointer(unsafe.SliceData(output))), (*C.int)(unsafe.Pointer(count)), C.int(lost), C.int(cond), &s[0], (*C.uchar)(unsafe.Pointer(unsafe.SliceData(buf)))))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(&numeric)), len(d)), d)
	numeric.FpsNLSF_CB = dec.FpsNLSF_CB
	numeric.Fpitch_lag_low_bits_iCDF = dec.Fpitch_lag_low_bits_iCDF
	numeric.Fpitch_contour_iCDF = dec.Fpitch_contour_iCDF
	numeric.Fresampler_state.FCoefs = dec.Fresampler_state.FCoefs
	*dec = numeric
	if ec != nil {
		ec.Fstorage = uint32(s[0])
		ec.Fend_offs = uint32(s[1])
		ec.Fend_window = uint32(s[2])
		ec.Fnend_bits = int32(s[3])
		ec.Fnbits_total = int32(s[4])
		ec.Foffs = uint32(s[5])
		ec.Frng = uint32(s[6])
		ec.Fval = uint32(s[7])
		ec.Fext = uint32(s[8])
		ec.Frem = int32(s[9])
		ec.Ferror1 = int32(s[10])
	}
	return ret
}

func nativeDecodeFrameFinish(dec *opuscc.OpusT_silk_decoder_state, ctrl *opuscc.OpusT_silk_decoder_control, count *int32, length, alias int32) {
	d, c := make([]byte, int(unsafe.Sizeof(*dec))), make([]byte, int(unsafe.Sizeof(*ctrl)))
	copy(d, unsafe.Slice((*byte)(unsafe.Pointer(dec)), len(d)))
	copy(c, unsafe.Slice((*byte)(unsafe.Pointer(ctrl)), len(c)))
	C.decoder_frame_finish((*C.uchar)(unsafe.Pointer(&d[0])), (*C.uchar)(unsafe.Pointer(&c[0])), (*C.int)(unsafe.Pointer(count)), C.int(length), C.int(alias))
	copy(unsafe.Slice((*byte)(unsafe.Pointer(dec)), len(d)), d)
	copy(unsafe.Slice((*byte)(unsafe.Pointer(ctrl)), len(c)), c)
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
