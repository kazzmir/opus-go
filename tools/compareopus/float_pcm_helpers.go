//go:build compareopus && cgo

package main

/*
void celt_float2int16_c(const float *in, short *out, int cnt);
*/
import "C"
import "unsafe"

func nativeFloatPCM(input []float32, output []int16) {
	C.celt_float2int16_c((*C.float)(unsafe.Pointer(&input[0])), (*C.short)(unsafe.Pointer(&output[0])), C.int(len(input)))
}
