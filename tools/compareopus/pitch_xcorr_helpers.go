//go:build compareopus && cgo

package main

/*
#define VAR_ARRAYS 1
#define pitch_downsample compare_pitch_downsample
#define celt_pitch_xcorr_c compare_pitch_xcorr
#define pitch_search compare_pitch_search
#define remove_doubling compare_remove_doubling
#include "../../../opus/celt/pitch.c"
*/
import "C"
import "unsafe"

func nativeScalarPitchXCorr(x, y, out []float32, n, lags int32) {
	C.compare_pitch_xcorr((*C.float)(unsafe.Pointer(unsafe.SliceData(x))), (*C.float)(unsafe.Pointer(unsafe.SliceData(y))), (*C.float)(unsafe.Pointer(&out[0])), C.int(n), C.int(lags), 0)
}
