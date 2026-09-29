//go:build compareopus && cgo

package main

/*
#include <opus.h>
*/
import "C"
import "unsafe"

func nativeSoftClip(pcm, mem []float32, n, channels int32) {
	C.opus_pcm_soft_clip((*C.float)(unsafe.Pointer(unsafe.SliceData(pcm))), C.int(n), C.int(channels), (*C.float)(unsafe.Pointer(unsafe.SliceData(mem))))
}
