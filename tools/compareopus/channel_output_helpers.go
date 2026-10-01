//go:build compareopus && cgo

package main

/*
#include <stdlib.h>
static _Thread_local int channels_factory_fail;
static void *channels_factory_alloc(size_t size){return channels_factory_fail?NULL:calloc(1,size);}
#define OVERRIDE_OPUS_ALLOC 1
#define opus_alloc channels_factory_alloc
#define VAR_ARRAYS 1
#define ENABLE_ASSERTIONS 1
#define opus_decoder_init validation_decoder_init
#define opus_decoder_get_size validation_decoder_get_size
#define opus_multistream_decoder_get_size comparison_channels_get_size
#define opus_multistream_decoder_init comparison_channels_init
#define opus_multistream_decoder_create comparison_channels_create
#define opus_multistream_decode_native comparison_channels_decode_native
#define opus_multistream_decode comparison_channels_decode
#define opus_multistream_decode24 comparison_channels_decode24
#define opus_multistream_decode_float comparison_channels_decode_float
#define opus_multistream_decoder_ctl_va_list comparison_channels_ctl_va
#define opus_multistream_decoder_ctl comparison_channels_ctl
#define opus_multistream_decoder_destroy comparison_channels_destroy
#include "../../../opus/src/opus_multistream_decoder.c"
extern void validation_normalize_decoder_mode(void *);
void comparison_normalize_ms_modes(void *base,int streams,int coupled) {char *ptr=(char*)base+align(sizeof(OpusMSDecoder));for(int i=0;i<streams;i++){validation_normalize_decoder_mode(ptr);ptr+=align(validation_decoder_get_size(i<coupled?2:1));}}
static int native_ms_create_image(unsigned char *data,size_t size,int rate,int channels,int streams,int coupled,const unsigned char *mapping,int fail) {
 int error=99;channels_factory_fail=fail;OpusMSDecoder *st=comparison_channels_create(rate,channels,streams,coupled,mapping,&error);channels_factory_fail=0;
 if(st){comparison_normalize_ms_modes(st,streams,coupled);memcpy(data,st,size);comparison_channels_destroy(st);}return error;
}
static int native_ms_init_image(unsigned char *data,size_t size,int rate,int channels,int streams,int coupled,const unsigned char *mapping,int mapping_offset) {
 OpusMSDecoder *st=malloc(size);memcpy(st,data,size);if(mapping_offset>=0)mapping=(unsigned char*)st+mapping_offset;
 int result=comparison_channels_init(st,rate,channels,streams,coupled,mapping);
 if(result==OPUS_OK){char *ptr=(char*)st+align(sizeof(*st));for(int i=0;i<streams;i++){validation_normalize_decoder_mode(ptr);ptr+=align(validation_decoder_get_size(i<coupled?2:1));}}
 memcpy(data,st,size);free(st);return result;
}
static int native_ms_packet_validate(const unsigned char *data,int length,int streams,int Fs) {return opus_multistream_packet_validate(data,length,streams,Fs);}
static int native_ms_validate(void *state) {OpusMSDecoder *st=state;validate_ms_decoder(st);return validate_layout(&st->layout);}
static void native_channel_output(void *dst,int ds,int dc,const float *src,int ss,int n,int op) {
 if(op==0) opus_copy_channel_out_float(dst,ds,dc,src,ss,n,NULL);
 else if(op==1) opus_copy_channel_out_short(dst,ds,dc,src,ss,n,NULL);
 else opus_copy_channel_out_int24(dst,ds,dc,src,ss,n,NULL);
}
*/
import "C"

import "unsafe"
import "github.com/kazzmir/opus-go/opuscc"

func nativeMSCreateImage(data []byte, rate, channels, streams, coupled int32, mapping []byte, fail bool) int32 {
	f := C.int(0)
	if fail {
		f = 1
	}
	return int32(C.native_ms_create_image((*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.size_t(len(data)), C.int(rate), C.int(channels), C.int(streams), C.int(coupled), (*C.uchar)(unsafe.Pointer(unsafe.SliceData(mapping))), f))
}

func nativeMSInitImage(data []byte, rate, channels, streams, coupled int32, mapping []byte, mappingOffset int32) int32 {
	return int32(C.native_ms_init_image((*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.size_t(len(data)), C.int(rate), C.int(channels), C.int(streams), C.int(coupled), (*C.uchar)(unsafe.Pointer(unsafe.SliceData(mapping))), C.int(mappingOffset)))
}
func nativeMSSize(streams, coupled int32) int32 {
	return int32(C.comparison_channels_get_size(C.int(streams), C.int(coupled)))
}

func nativeMSPacketValidation(data *byte, length, streams, Fs int32) int32 {
	return int32(C.native_ms_packet_validate((*C.uchar)(unsafe.Pointer(data)), C.int(length), C.int(streams), C.int(Fs)))
}

func nativeMSValidation(st *opuscc.OpusT_OpusMSDecoder) int32 {
	return int32(C.native_ms_validate(unsafe.Pointer(st)))
}

func nativeChannelOutput(dst unsafe.Pointer, ds, dc int32, src []float32, ss, n int32, operations ...int32) {
	var op int32
	if len(operations) > 0 {
		op = operations[0]
	}
	C.native_channel_output(dst, C.int(ds), C.int(dc), (*C.float)(unsafe.Pointer(unsafe.SliceData(src))), C.int(ss), C.int(n), C.int(op))
}
