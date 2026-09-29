//go:build compareopus && cgo

package main

/*
#include <opus.h>
static int decoder_duration(const unsigned char *p,int len,int rate) {
 int err;OpusDecoder *d=opus_decoder_create(rate,2,&err);
 if(!d) return err;
 int result=opus_decoder_get_nb_samples(d,p,len);opus_decoder_destroy(d);return result;
}
*/
import "C"
import "unsafe"

func nativeDecoderDuration(packet []byte, n, rate int32) int32 {
	return int32(C.decoder_duration((*C.uchar)(unsafe.Pointer(unsafe.SliceData(packet))), C.int(n), C.int(rate)))
}

func nativePacketDuration(packet []byte, n, rate int32) int32 {
	return int32(C.opus_packet_get_nb_samples((*C.uchar)(unsafe.Pointer(unsafe.SliceData(packet))), C.opus_int32(n), C.opus_int32(rate)))
}
