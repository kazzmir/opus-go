//go:build compareopus && cgo

package main

/*
void _celt_lpc(float *lpc, const float *ac, int p);
*/
import "C"
import "unsafe"

func nativeLPC(output, ac []float32) {
	C._celt_lpc((*C.float)(unsafe.Pointer(&output[0])), (*C.float)(unsafe.Pointer(&ac[0])), C.int(len(output)))
}
