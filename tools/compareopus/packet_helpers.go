//go:build compareopus && cgo

package main

/*
#include <opus.h>
#include "../../../opus/celt/arch.h"
#include "../../../opus/celt/float_cast.h"
static void native_int24_pcm(const float *input,int *output,int count) {int i;for(i=0;i<count;i++) output[i]=RES2INT24(input[i]);}
int opus_packet_parse_impl(const unsigned char *,opus_int32,int,unsigned char *,const unsigned char *[48],opus_int16 [48],int *,opus_int32 *,const unsigned char **,opus_int32 *);
static void native_parse_packet(const unsigned char *data,int length,int self,int mask,short *size,int *frameOffsets,int *info,int public_api) {
 const unsigned char *frames[48],*padding=data;unsigned char toc=0xa5;int payload=-77,packet=-78,padlen=-79;int i;
 for(i=0;i<48;i++) frames[i]=data;
 int result;
 if(public_api) result=opus_packet_parse(data,length,mask&1?&toc:NULL,mask&2?frames:NULL,mask&4?size:NULL,mask&8?&payload:NULL);
 else result=opus_packet_parse_impl(data,length,self,mask&1?&toc:NULL,mask&2?frames:NULL,mask&4?size:NULL,mask&8?&payload:NULL,mask&16?&packet:NULL,mask&32?&padding:NULL,mask&32?&padlen:NULL);
 for(i=0;i<48;i++) frameOffsets[i]=frames[i]?(int)(frames[i]-data):-1;
 info[0]=result;info[1]=toc;info[2]=payload;info[3]=packet;info[4]=padding?(int)(padding-data):-1;info[5]=padlen;
}
*/
import "C"

import "unsafe"

type packetParseResult struct {
	Count                                int32
	Toc                                  byte
	Sizes                                [50]int16
	Frames                               [48]int32
	Payload, Packet, Padding, PaddingLen int32
}

func nativeInt24PCM(input []float32) []int32 {
	output := make([]int32, len(input))
	C.native_int24_pcm((*C.float)(unsafe.Pointer(unsafe.SliceData(input))), (*C.int)(unsafe.Pointer(unsafe.SliceData(output))), C.int(len(input)))
	return output
}

func nativePacketParse(data *byte, length, self, mask, public int32) packetParseResult {
	r := packetParseResult{}
	for i := range r.Sizes {
		r.Sizes[i] = 1234
	}
	var info [6]C.int
	C.native_parse_packet((*C.uchar)(unsafe.Pointer(data)), C.int(length), C.int(self), C.int(mask), (*C.short)(unsafe.Pointer(&r.Sizes[1])), (*C.int)(unsafe.Pointer(&r.Frames[0])), &info[0], C.int(public))
	r.Count = int32(info[0])
	r.Toc = byte(info[1])
	r.Payload = int32(info[2])
	r.Packet = int32(info[3])
	r.Padding = int32(info[4])
	r.PaddingLen = int32(info[5])
	return r
}

// These small bridges let opt-in tests compare the real C implementation;
// Go does not support importing C directly in a _test.go file.
func nativePacketLBRR(data *byte, length int32) int32 {
	return int32(C.opus_packet_has_lbrr((*C.uchar)(unsafe.Pointer(data)), C.opus_int32(length)))
}

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
