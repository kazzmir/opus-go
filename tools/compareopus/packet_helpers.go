//go:build compareopus && cgo

package main

/*
#include <opus.h>
*/
import "C"

import "unsafe"

// These small bridges let opt-in tests compare the real C implementation;
// Go does not support importing C directly in a _test.go file.
func nativePacketBandwidth(data *byte) int32 {
	return int32(C.opus_packet_get_bandwidth((*C.uchar)(unsafe.Pointer(data))))
}

func nativePacketChannels(data *byte) int32 {
	return int32(C.opus_packet_get_nb_channels((*C.uchar)(unsafe.Pointer(data))))
}

func nativePacketFrames(data *byte, length int32) int32 {
	return int32(C.opus_packet_get_nb_frames((*C.uchar)(unsafe.Pointer(data)), C.opus_int32(length)))
}

func nativePacketSamplesPerFrame(data *byte, rate int32) int32 {
	return int32(C.opus_packet_get_samples_per_frame((*C.uchar)(unsafe.Pointer(data)), C.opus_int32(rate)))
}
