//go:build compareopus && cgo

package main

/*
void _celt_lpc(float *lpc, const float *ac, int p);
#define VAR_ARRAYS 1
#define _celt_lpc comparison_lpc_unused
#define celt_fir_c comparison_fir
#define celt_iir comparison_iir
#define _celt_autocorr comparison_autocorr
// Use the scalar pitch.c fixture, not libopus's SIMD-dispatched build.
#define celt_pitch_xcorr_c compare_pitch_xcorr
#include "../../../opus/celt/celt_lpc.c"
// Source-equivalent leaf from celt_decoder.c, using its actual MAXG macro.
static int compare_decode_packet_start(int loss,int skip) {if(loss==0)skip=0;return skip;}
static void compare_decode_view_offsets(int *v,int bands,int overlap,int channels,int N) {v[0]=DEC_PITCH_BUF_SIZE+overlap;v[1]=v[0]*channels;v[2]=v[1]+2*bands;v[3]=v[2]+2*bands;v[4]=v[3]+2*bands;v[5]=DEC_PITCH_BUF_SIZE-N;}
static void compare_decode_energy_merge_mono(float *e,int bands) {for(int i=0;i<bands;i++)e[i]=MAXG(e[i],e[bands+i]);}
// celt_decoder.c defines FRAME_NORMAL as 1.
static void compare_decode_postfilter_clamp(int *p,int minimum) {p[0]=IMAX(minimum,p[0]);p[1]=IMAX(minimum,p[1]);}
static void compare_decode_packet_finish(int *state) {state[0]=0;state[1]=0;state[2]=1;state[3]=0;}
static void compare_decode_recovery_band(float *e,const float *l,const float *p,int missing,float safety) {if(*e<MAXG(*l,*p)){float E0=*e,E1=*l,E2=*p;float slope=MAX32(E1-E0,HALF32(E2-E0));slope=MING(slope,2.f);E0-=MAX32(0,(1+missing)*slope);*e=MAX32(-20.f,E0);}else *e=MING(MING(*e,*l),*p);*e-=safety;}
static int compare_decode_recovery_safety(int loss,int LM,float *safety) {*safety=0;if(LM==0)*safety=1.5f;else if(LM==1)*safety=.5f;return IMIN(10,loss>>LM);}
static void compare_decode_recover_energy(float *e,const float *l,const float *p,int bands,int start,int end,int LM,int intra,int loss) {if(!intra&&loss){int c=0;do{float safety;int missing=compare_decode_recovery_safety(loss,LM,&safety);for(int i=start;i<end;i++)compare_decode_recovery_band(e+c*bands+i,l+c*bands+i,p+c*bands+i,missing,safety);}while(++c<2);}}
static void compare_decode_postfilter_finish(int *p,float *g,int period,float gain,int tapset,int LM) {p[1]=p[0];g[1]=g[0];p[3]=p[2];p[0]=period;g[0]=gain;p[2]=tapset;if(LM){p[1]=p[0];g[1]=g[0];p[3]=p[2];}}
#include "entdec.h"
#include "celt.h"
#include "bands.h"
static void compare_decode_header(unsigned *s,unsigned char *data,int op,int *a,float *gain) {
 ec_dec dec={0};dec.buf=data;dec.storage=s[0];dec.end_offs=s[1];dec.end_window=s[2];dec.nend_bits=(int)s[3];dec.nbits_total=(int)s[4];dec.offs=s[5];dec.rng=s[6];dec.val=s[7];dec.ext=s[8];dec.rem=(int)s[9];dec.error=(int)s[10];
 if(op==0){int total=a[0],tell=ec_tell(&dec),silence=0;if(tell>=total)silence=1;else if(tell==1)silence=ec_dec_bit_logp(&dec,15);if(silence){tell=total;dec.nbits_total+=tell-ec_tell(&dec);}a[1]=silence;a[2]=tell;}
 if(op==2){int LM=a[0],M=a[1],tell=a[2],total=a[3],transient=0,shortBlocks=0,intra=0;if(LM>0&&tell+3<=total){transient=ec_dec_bit_logp(&dec,3);tell=ec_tell(&dec);}if(transient)shortBlocks=M;if(tell+3<=total)intra=ec_dec_bit_logp(&dec,3);a[4]=transient;a[5]=shortBlocks;a[6]=intra;a[7]=tell;}
 if(op==6){a[1]=a[0]>0?ec_dec_bits(&dec,1):0;}
 if(op==5){int budget=(a[0]*8<<BITRES)-(int)ec_tell_frac(&dec)-1;int reserved=a[1]&&a[2]>=2&&budget>=((a[2]+2)<<BITRES)?1<<BITRES:0;a[3]=budget-reserved;a[4]=reserved;}
 if(op==4){a[2]=a[0]+(6<<BITRES)<=a[1]?ec_dec_icdf(&dec,trim_icdf,7):5;}
 if(op==3){int tell=ec_tell(&dec),spread=SPREAD_NORMAL;if(tell+4<=a[0])spread=ec_dec_icdf(&dec,spread_icdf,5);a[1]=spread;a[2]=tell;}
 if(op==1){int start=a[0],tell=a[1],total=a[2],pitch=0,tap=0;*gain=0;if(start==0&&tell+16<=total){if(ec_dec_bit_logp(&dec,1)){int octave=ec_dec_uint(&dec,6);pitch=(16<<octave)+ec_dec_bits(&dec,4+octave)-1;int qg=ec_dec_bits(&dec,3);if(ec_tell(&dec)+2<=total)tap=ec_dec_icdf(&dec,tapset_icdf,2);*gain=.09375f*(qg+1);}tell=ec_tell(&dec);}a[3]=pitch;a[4]=tap;a[5]=tell;}
 s[0]=dec.storage;s[1]=dec.end_offs;s[2]=dec.end_window;s[3]=dec.nend_bits;s[4]=dec.nbits_total;s[5]=dec.offs;s[6]=dec.rng;s[7]=dec.val;s[8]=dec.ext;s[9]=dec.rem;s[10]=dec.error;
}
static int compare_decode_packet_error(int *stateError,int nbits,unsigned rng,int error,int length) {ec_dec dec={0};dec.nbits_total=nbits;dec.rng=rng;dec.error=error;if(ec_tell(&dec)>8*length)return -3;if(dec.error)*stateError=1;return 0;}
static int compare_decode_boosts(unsigned *s,unsigned char *data,short *e,int *cap,int *out,int start,int end,int C,int LM,int total,int *tell) {
 ec_dec dec={0};dec.buf=data;dec.storage=s[0];dec.end_offs=s[1];dec.end_window=s[2];dec.nend_bits=(int)s[3];dec.nbits_total=(int)s[4];dec.offs=s[5];dec.rng=s[6];dec.val=s[7];dec.ext=s[8];dec.rem=(int)s[9];dec.error=(int)s[10];
 int logp=6;total<<=BITRES;*tell=ec_tell_frac(&dec);for(int i=start;i<end;i++){int width=C*(e[i+1]-e[i])<<LM;int quanta=IMIN(width<<BITRES,IMAX(6<<BITRES,width));int loop_logp=logp,boost=0;while(*tell+(loop_logp<<BITRES)<total&&boost<cap[i]){int flag=ec_dec_bit_logp(&dec,loop_logp);*tell=ec_tell_frac(&dec);if(!flag)break;boost+=quanta;total-=quanta;loop_logp=1;}out[i]=boost;if(boost>0)logp=IMAX(2,logp-1);}
 s[0]=dec.storage;s[1]=dec.end_offs;s[2]=dec.end_window;s[3]=dec.nend_bits;s[4]=dec.nbits_total;s[5]=dec.offs;s[6]=dec.rng;s[7]=dec.val;s[8]=dec.ext;s[9]=dec.rem;s[10]=dec.error;return total;
}
static void compare_decode_silence_energy(float *e,int bands,int channels) {for(int i=0;i<channels*bands;i++)e[i]=-28.f;}
static void compare_decode_history_move(float *h,int N,int length) {if(length>0)memmove(h,h+N,length*sizeof(float));}
static void compare_decode_energy_clear(float *e,float *l,float *p,int bands,int start,int end) {int c=0;do{for(int i=0;i<start;i++){e[c*bands+i]=0;l[c*bands+i]=p[c*bands+i]=-28.f;}for(int i=end;i<bands;i++){e[c*bands+i]=0;l[c*bands+i]=p[c*bands+i]=-28.f;}}while(++c<2);}
static void compare_decode_energy_background(float *b,const float *e,int bands,int loss,int M) {float increase=IMIN(160,loss+M)*.001f;for(int i=0;i<2*bands;i++)b[i]=MING(b[i]+increase,e[i]);}
static void compare_decode_energy_logs(const float *e,float *l,float *p,int bands,int transient) {if(!transient){memcpy(p,l,2*bands*sizeof(float));memcpy(l,e,2*bands*sizeof(float));}else for(int i=0;i<2*bands;i++)l[i]=MING(l[i],e[i]);}
static void compare_decode_energy_mono(float *energy,int bands) {if(bands>0)memcpy(energy+bands,energy,bands*sizeof(float));}
static int compare_plc_dispatch(int duration,int start,int skip) {return duration>=40||start!=0||skip!=0;}
static void compare_plc_finish(int *state,int loss,int LM,int frameType) {state[0]=IMIN(10000,loss+(1<<LM));state[1]=IMIN(10000,state[1]+(1<<LM));state[2]=frameType;}
static void compare_plc_lpc_history(float *memory,const float *h,int size,int N) {for(int i=0;i<CELT_LPC_ORDER;i++)memory[i]=h[size-N-1-i];}
static float compare_plc_extrapolate(float *h,const float *exc,int size,int period,int N,int overlap,int pitch,float fade,float decay) {float energy=0,attenuation=fade*decay;int offset=period-pitch;for(int i=0,j=0;i<N+overlap;i++,j++){if(j>=pitch){j-=pitch;attenuation*=decay;}h[size-N+i]=attenuation*exc[offset+j];float sample=h[size-period-N+offset+j];energy+=sample*sample;}return energy;}
static void compare_plc_history(float *exc,const float *h,int size,int period) {for(int i=0;i<period+CELT_LPC_ORDER;i++)exc[i]=h[size-period-CELT_LPC_ORDER+i];}
static void compare_plc_attenuate(float *x,const float *w,int length,int overlap,float s1) {float s2=0;for(int i=0;i<length;i++) {float tmp=x[i];s2+=tmp*tmp;}if(!(s1>.2f*s2)) {for(int i=0;i<length;i++)x[i]=0;}else if(s1<s2) {float ratio=celt_sqrt((s1+1)/(s2+1));for(int i=0;i<overlap;i++){float g=1-w[i]*(1-ratio);x[i]=g*x[i];}for(int i=overlap;i<length;i++)x[i]=ratio*x[i];}}
static float compare_plc_exc(const float *exc,int period,int length) {float e1=1,e2=1;int half=length>>1;for(int i=0;i<half;i++){float e=exc[period-half+i];e1+=e*e;e=exc[period-2*half+i];e2+=e*e;}e1=MIN32(e1,e2);return celt_sqrt(e1/e2);}
static void compare_plc_lag(float *ac) {ac[0]*=1.0001f;for(int i=1;i<=24;i++)ac[i]-=ac[i]*(.008f*.008f)*i*i;}
static void compare_plc_decay(float *a,const float *b,int bands,int start,int end,int C,int loss) {float decay=loss==0?1.5f:.5f;int c=0;do {for(int i=start;i<end;i++)a[c*bands+i]=MAXG(b[c*bands+i],a[c*bands+i]-decay);}while(++c<C);}
*/
import "C"
import "unsafe"
import "github.com/kazzmir/opus-go/opuscc"

func nativeCeltDecodePacketStart(loss, skip int32) int32 {
	return int32(C.compare_decode_packet_start(C.int(loss), C.int(skip)))
}
func nativeCeltDecodeViewOffsets(bands, overlap, channels, N int32) (v [6]int32) {
	C.compare_decode_view_offsets((*C.int)(unsafe.Pointer(&v[0])), C.int(bands), C.int(overlap), C.int(channels), C.int(N))
	return
}
func nativeCeltDecodeEnergyMergeMono(e *float32, bands int32) {
	C.compare_decode_energy_merge_mono((*C.float)(unsafe.Pointer(e)), C.int(bands))
}
func nativeCeltDecodePostfilterClamp(state *opuscc.OpusT_OpusCustomDecoder) {
	p := [2]C.int{C.int(state.Fpostfilter_period), C.int(state.Fpostfilter_period_old)}
	C.compare_decode_postfilter_clamp(&p[0], C.int(opuscc.COMBFILTER_MINPERIOD))
	state.Fpostfilter_period = int32(p[0])
	state.Fpostfilter_period_old = int32(p[1])
}
func nativeCeltDecodePacketFinish(state *opuscc.OpusT_OpusCustomDecoder) {
	v := [4]C.int{C.int(state.Floss_duration), C.int(state.Fplc_duration), C.int(state.Flast_frame_type), C.int(state.Fprefilter_and_fold)}
	C.compare_decode_packet_finish(&v[0])
	state.Floss_duration = int32(v[0])
	state.Fplc_duration = int32(v[1])
	state.Flast_frame_type = int32(v[2])
	state.Fprefilter_and_fold = int32(v[3])
}
func nativeCeltDecodeRecoverEnergy(e, l, p []float32, bands, start, end, LM, intra, loss int32) {
	C.compare_decode_recover_energy((*C.float)(unsafe.Pointer(unsafe.SliceData(e))), (*C.float)(unsafe.Pointer(unsafe.SliceData(l))), (*C.float)(unsafe.Pointer(unsafe.SliceData(p))), C.int(bands), C.int(start), C.int(end), C.int(LM), C.int(intra), C.int(loss))
}
func nativeCeltDecodeRecoveryBand(e, l, p *float32, missing int32, safety float32) {
	C.compare_decode_recovery_band((*C.float)(unsafe.Pointer(e)), (*C.float)(unsafe.Pointer(l)), (*C.float)(unsafe.Pointer(p)), C.int(missing), C.float(safety))
}
func nativeCeltDecodeRecoverySafety(loss, LM int32) (int32, float32) {
	var safety C.float
	m := C.compare_decode_recovery_safety(C.int(loss), C.int(LM), &safety)
	return int32(m), float32(safety)
}
func nativeCeltDecodePostfilterFinish(state *opuscc.OpusT_OpusCustomDecoder, period int32, gain float32, tapset, LM int32) {
	p := [4]C.int{C.int(state.Fpostfilter_period), C.int(state.Fpostfilter_period_old), C.int(state.Fpostfilter_tapset), C.int(state.Fpostfilter_tapset_old)}
	g := [2]C.float{C.float(state.Fpostfilter_gain), C.float(state.Fpostfilter_gain_old)}
	C.compare_decode_postfilter_finish(&p[0], &g[0], C.int(period), C.float(gain), C.int(tapset), C.int(LM))
	state.Fpostfilter_period = int32(p[0])
	state.Fpostfilter_period_old = int32(p[1])
	state.Fpostfilter_tapset = int32(p[2])
	state.Fpostfilter_tapset_old = int32(p[3])
	state.Fpostfilter_gain = float32(g[0])
	state.Fpostfilter_gain_old = float32(g[1])
}
func nativeCeltDecodeHeader(ec *opuscc.OpusT_ec_ctx, data []byte, op int32, a *[8]int32) float32 {
	s := [11]C.uint{C.uint(ec.Fstorage), C.uint(ec.Fend_offs), C.uint(ec.Fend_window), C.uint(ec.Fnend_bits), C.uint(ec.Fnbits_total), C.uint(ec.Foffs), C.uint(ec.Frng), C.uint(ec.Fval), C.uint(ec.Fext), C.uint(ec.Frem), C.uint(ec.Ferror1)}
	var gain C.float
	C.compare_decode_header(&s[0], (*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.int(op), (*C.int)(unsafe.Pointer(a)), &gain)
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
	return float32(gain)
}
func nativeCeltDecodePacketError(state *opuscc.OpusT_OpusCustomDecoder, ec *opuscc.OpusT_ec_ctx, length int32) int32 {
	e := C.int(state.Ferror1)
	r := C.compare_decode_packet_error(&e, C.int(ec.Fnbits_total), C.uint(ec.Frng), C.int(ec.Ferror1), C.int(length))
	state.Ferror1 = int32(e)
	return int32(r)
}
func nativeCeltDecodeBoosts(ec *opuscc.OpusT_ec_ctx, data []byte, bands []int16, cap, out []int32, start, end, channels, LM, total int32) (int32, int32) {
	s := [11]C.uint{C.uint(ec.Fstorage), C.uint(ec.Fend_offs), C.uint(ec.Fend_window), C.uint(ec.Fnend_bits), C.uint(ec.Fnbits_total), C.uint(ec.Foffs), C.uint(ec.Frng), C.uint(ec.Fval), C.uint(ec.Fext), C.uint(ec.Frem), C.uint(ec.Ferror1)}
	var tell C.int
	r := C.compare_decode_boosts(&s[0], (*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), (*C.short)(unsafe.Pointer(unsafe.SliceData(bands))), (*C.int)(unsafe.Pointer(unsafe.SliceData(cap))), (*C.int)(unsafe.Pointer(unsafe.SliceData(out))), C.int(start), C.int(end), C.int(channels), C.int(LM), C.int(total), &tell)
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
	return int32(r), int32(tell)
}
func nativeCeltDecodeSilenceEnergy(e *float32, bands, channels int32) {
	C.compare_decode_silence_energy((*C.float)(unsafe.Pointer(e)), C.int(bands), C.int(channels))
}
func nativeCeltDecodeHistoryMove(h *float32, N, length int32) {
	C.compare_decode_history_move((*C.float)(unsafe.Pointer(h)), C.int(N), C.int(length))
}
func nativeCeltDecodeEnergyClear(e, l, p *float32, bands, start, end int32) {
	C.compare_decode_energy_clear((*C.float)(unsafe.Pointer(e)), (*C.float)(unsafe.Pointer(l)), (*C.float)(unsafe.Pointer(p)), C.int(bands), C.int(start), C.int(end))
}
func nativeCeltDecodeEnergyBackground(b, e *float32, bands, loss, M int32) {
	C.compare_decode_energy_background((*C.float)(unsafe.Pointer(b)), (*C.float)(unsafe.Pointer(e)), C.int(bands), C.int(loss), C.int(M))
}
func nativeCeltDecodeEnergyLogs(e, l, p *float32, bands, transient int32) {
	C.compare_decode_energy_logs((*C.float)(unsafe.Pointer(e)), (*C.float)(unsafe.Pointer(l)), (*C.float)(unsafe.Pointer(p)), C.int(bands), C.int(transient))
}
func nativeCeltDecodeEnergyMono(energy *float32, bands int32) {
	C.compare_decode_energy_mono((*C.float)(unsafe.Pointer(energy)), C.int(bands))
}
func nativeCeltPLCDispatch(duration, start, skip int32) bool {
	return C.compare_plc_dispatch(C.int(duration), C.int(start), C.int(skip)) != 0
}
func nativeCeltPLCFinish(state *opuscc.OpusT_OpusCustomDecoder, loss, LM, frameType int32) {
	v := [3]C.int{C.int(state.Floss_duration), C.int(state.Fplc_duration), C.int(state.Flast_frame_type)}
	C.compare_plc_finish(&v[0], C.int(loss), C.int(LM), C.int(frameType))
	state.Floss_duration = int32(v[0])
	state.Fplc_duration = int32(v[1])
	state.Flast_frame_type = int32(v[2])
}
func nativeCeltPLCLPCHistory(memory *[24]float32, history *float32, size, N int32) {
	C.compare_plc_lpc_history((*C.float)(unsafe.Pointer(memory)), (*C.float)(unsafe.Pointer(history)), C.int(size), C.int(N))
}
func nativeCeltPLCExtrapolate(history, exc *float32, size, period, N, overlap, pitch int32, fade, decay float32) float32 {
	return float32(C.compare_plc_extrapolate((*C.float)(unsafe.Pointer(history)), (*C.float)(unsafe.Pointer(exc)), C.int(size), C.int(period), C.int(N), C.int(overlap), C.int(pitch), C.float(fade), C.float(decay)))
}
func nativeCeltPLCExcitationHistory(exc, history *float32, size, period int32) {
	C.compare_plc_history((*C.float)(unsafe.Pointer(exc)), (*C.float)(unsafe.Pointer(history)), C.int(size), C.int(period))
}
func nativeCeltPLCSynthesisAttenuate(output, window *float32, length, overlap int32, s1 float32) {
	C.compare_plc_attenuate((*C.float)(unsafe.Pointer(output)), (*C.float)(unsafe.Pointer(window)), C.int(length), C.int(overlap), C.float(s1))
}
func nativeCeltPLCExcitationDecay(exc *float32, period, length int32) float32 {
	return float32(C.compare_plc_exc((*C.float)(unsafe.Pointer(exc)), C.int(period), C.int(length)))
}
func nativeCeltPLCLagWindow(ac *[25]float32) { C.compare_plc_lag((*C.float)(unsafe.Pointer(ac))) }
func nativeCeltPLCDecay(a, b *float32, bands, start, end, channels, loss int32) {
	C.compare_plc_decay((*C.float)(unsafe.Pointer(a)), (*C.float)(unsafe.Pointer(b)), C.int(bands), C.int(start), C.int(end), C.int(channels), C.int(loss))
}

func nativeAutocorr(input, ac, window []float32, overlap, lag, n int32) int32 {
	return int32(C.comparison_autocorr((*C.float)(unsafe.Pointer(unsafe.SliceData(input))), (*C.float)(unsafe.Pointer(unsafe.SliceData(ac))), (*C.float)(unsafe.Pointer(unsafe.SliceData(window))), C.int(overlap), C.int(lag), C.int(n), 0))
}

func nativeIIR(input, coeff, out, mem []float32, N, ord int32) {
	C.comparison_iir((*C.float)(unsafe.Pointer(unsafe.SliceData(input))), (*C.float)(unsafe.Pointer(unsafe.SliceData(coeff))), (*C.float)(unsafe.Pointer(unsafe.SliceData(out))), C.int(N), C.int(ord), (*C.float)(unsafe.Pointer(unsafe.SliceData(mem))), 0)
}

func nativeFIR(input, coeff, out []float32, N, ord int32) {
	C.comparison_fir((*C.float)(unsafe.Pointer(&input[ord])), (*C.float)(unsafe.Pointer(unsafe.SliceData(coeff))), (*C.float)(unsafe.Pointer(unsafe.SliceData(out))), C.int(N), C.int(ord), 0)
}

func nativeLPC(output, ac []float32) {
	C._celt_lpc((*C.float)(unsafe.Pointer(&output[0])), (*C.float)(unsafe.Pointer(&ac[0])), C.int(len(output)))
}
