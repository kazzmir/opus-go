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
*/
import "C"
import "unsafe"

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
