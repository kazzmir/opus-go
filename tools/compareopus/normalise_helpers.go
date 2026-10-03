//go:build compareopus && cgo

package main

/*
#include <math.h>
// The linked native library is a floating-point libopus build.
void renormalise_vector(float *X, int N, float gain, int arch);
// Scalar floating-point formula from celt/vq.c and pitch.h, without the
// native build's OPUS_X86_PRESUME_SSE inner-product substitution.
static void scalar_renormalise(float *x, int n, float gain) {
    float xy = 0;
    int i;
    for (i = 0; i < n; i++) xy = xy + x[i]*x[i];
    float energy = 1e-15f + xy;
    float g = (1.f / (float)sqrt((double)energy)) * gain;
    for (i = 0; i < n; i++) x[i] = g*x[i];
}
// Noise loop from celt_decoder.c with the scalar normalization above, not
// the linked library's potentially presumed-SSE inner product.
static unsigned scalar_plc_noise(unsigned seed,const short *bands,float *x,int N,int start,int end,int LM,int C) {for(int c=0;c<C;c++)for(int i=start;i<end;i++){int offset=N*c+(bands[i]<<LM);int length=(bands[i+1]-bands[i])<<LM;for(int j=0;j<length;j++){seed=1664525U*seed+1013904223U;x[offset+j]=(float)((int)seed>>20);}if(length>0)scalar_renormalise(x+offset,length,1);}return seed;}
*/
import "C"

import "unsafe"

func nativeCeltPLCNoise(seed uint32, bands *int16, spectrum *float32, N, start, end, LM, channels int32) uint32 {
	return uint32(C.scalar_plc_noise(C.uint(seed), (*C.short)(unsafe.Pointer(bands)), (*C.float)(unsafe.Pointer(spectrum)), C.int(N), C.int(start), C.int(end), C.int(LM), C.int(channels)))
}
func nativeRenormalise(x []float32, gain float32) {
	// The native build may presume SSE even with arch=0.
	C.renormalise_vector((*C.float)(unsafe.Pointer(&x[0])), C.int(len(x)), C.float(gain), 0)
}

func scalarCRenormalise(x []float32, gain float32) {
	C.scalar_renormalise((*C.float)(unsafe.Pointer(&x[0])), C.int(len(x)), C.float(gain))
}
