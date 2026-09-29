//go:build compareopus && cgo

package main

/*
#include <opus.h>
*/
import "C"
import "unsafe"

func nativePacketDuration(packet []byte, n, rate int32) int32 {
	return int32(C.opus_packet_get_nb_samples((*C.uchar)(unsafe.Pointer(unsafe.SliceData(packet))), C.opus_int32(n), C.opus_int32(rate)))
}
