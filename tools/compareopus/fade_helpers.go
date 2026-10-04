//go:build compareopus && cgo

package main

/*
#define VAR_ARRAYS 1
#define OPUS_DISABLE_INTRINSICS 1
#define resampling_factor compare_comb_resampling_factor
#define comb_filter compare_comb_filter
#define init_caps compare_comb_init_caps
#define tf_select_table compare_comb_tf_select_table
#define opus_strerror compare_comb_strerror
#define opus_get_version_string compare_comb_version
#include "../../../opus/celt/celt.c"
static void compare_decode_postfilter(int *p,float *g,float *left,float *right,int channels,int N,int LM,int period,float gain,int tapset,const float *window,int overlap,int shortSize) {float *out[2]={left,right};int c=0;do{p[0]=IMAX(COMBFILTER_MINPERIOD,p[0]);p[1]=IMAX(COMBFILTER_MINPERIOD,p[1]);compare_comb_filter(out[c],out[c],p[1],p[0],shortSize,g[1],g[0],p[3],p[2],window,overlap,0);if(LM)compare_comb_filter(out[c]+shortSize,out[c]+shortSize,p[0],period,N-shortSize,g[0],gain,p[2],tapset,window,overlap,0);}while(++c<channels);}
// Scalar smooth_fade from src/opus_decoder.c, using the floating-point macros.
static void compare_fade(const float *a,const float *b,float *out,int n,int channels,const float *window,int rate) {
 int inc=48000/rate;
 for(int c=0;c<channels;c++) for(int i=0;i<n;i++) {
  celt_coef w=MULT_COEF(window[i*inc],window[i*inc]);
  out[i*channels+c]=ADD32(MULT_COEF_32(w,b[i*channels+c]),MULT_COEF_32(COEF_ONE-w,a[i*channels+c]));
 }
}
*/
import "C"
import "unsafe"
import "github.com/kazzmir/opus-go/opuscc"

func nativeCeltDecodePostfilter(state *opuscc.OpusT_OpusCustomDecoder, mode *opuscc.OpusT_OpusCustomMode, left, right *float32, channels, N, LM, period int32, gain float32, tapset, overlap int32) {
	p := [4]C.int{C.int(state.Fpostfilter_period), C.int(state.Fpostfilter_period_old), C.int(state.Fpostfilter_tapset), C.int(state.Fpostfilter_tapset_old)}
	g := [2]C.float{C.float(state.Fpostfilter_gain), C.float(state.Fpostfilter_gain_old)}
	C.compare_decode_postfilter(&p[0], &g[0], (*C.float)(unsafe.Pointer(left)), (*C.float)(unsafe.Pointer(right)), C.int(channels), C.int(N), C.int(LM), C.int(period), C.float(gain), C.int(tapset), (*C.float)(unsafe.Pointer(mode.Fwindow)), C.int(overlap), C.int(mode.FshortMdctSize))
	state.Fpostfilter_period = int32(p[0])
	state.Fpostfilter_period_old = int32(p[1])
}

func nativeComb(y, x *float32, T0, T1, N int32, g0, g1 float32, tap0, tap1 int32, window *float32, overlap int32) {
	C.compare_comb_filter((*C.float)(unsafe.Pointer(y)), (*C.float)(unsafe.Pointer(x)), C.int(T0), C.int(T1), C.int(N), C.float(g0), C.float(g1), C.int(tap0), C.int(tap1), (*C.float)(unsafe.Pointer(window)), C.int(overlap), 0)
}

func nativeFade(a, b, out, window []float32, n, channels, rate int32) {
	C.compare_fade((*C.float)(unsafe.Pointer(unsafe.SliceData(a))), (*C.float)(unsafe.Pointer(unsafe.SliceData(b))), (*C.float)(unsafe.Pointer(unsafe.SliceData(out))), C.int(n), C.int(channels), (*C.float)(unsafe.Pointer(unsafe.SliceData(window))), C.int(rate))
}
