//go:build compareopus && cgo

package main

/*
void exp_rotation(float *X, int len, int dir, int stride, int K, int spread);
*/
import "C"
import "unsafe"

func nativeExpRotation(x []float32, dir, stride, k, spread int32) {
	C.exp_rotation((*C.float)(unsafe.Pointer(&x[0])), C.int(len(x)), C.int(dir), C.int(stride), C.int(k), C.int(spread))
}
