//go:build compareopus && cgo

package main

/*
int opus_limit2_checkwithin1_c(float *samples, int cnt);
*/
import "C"

import "unsafe"

func nativeLimit2(samples *float32, count int32) int32 {
	return int32(C.opus_limit2_checkwithin1_c((*C.float)(unsafe.Pointer(samples)), C.int(count)))
}
