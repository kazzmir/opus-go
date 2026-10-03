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
