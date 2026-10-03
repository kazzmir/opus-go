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
static void compare_plc_lag(float *ac) {ac[0]*=1.0001f;for(int i=1;i<=24;i++)ac[i]-=ac[i]*(.008f*.008f)*i*i;}
static void compare_plc_decay(float *a,const float *b,int bands,int start,int end,int C,int loss) {float decay=loss==0?1.5f:.5f;int c=0;do {for(int i=start;i<end;i++)a[c*bands+i]=MAXG(b[c*bands+i],a[c*bands+i]-decay);}while(++c<C);}
*/
import "C"
import "unsafe"

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
