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
*/
import "C"

import "unsafe"

func nativeRenormalise(x []float32, gain float32) {
	// The native build may presume SSE even with arch=0.
	C.renormalise_vector((*C.float)(unsafe.Pointer(&x[0])), C.int(len(x)), C.float(gain), 0)
}

func scalarCRenormalise(x []float32, gain float32) {
	C.scalar_renormalise((*C.float)(unsafe.Pointer(&x[0])), C.int(len(x)), C.float(gain))
}
