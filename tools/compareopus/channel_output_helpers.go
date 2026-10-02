//go:build compareopus && cgo

package main

/*
#include <stdlib.h>
static _Thread_local int channels_factory_fail;
static void *channels_factory_alloc(size_t size){return channels_factory_fail?NULL:calloc(1,size);}
static _Thread_local int channels_free_calls,channels_free_matches;
static _Thread_local void *channels_free_expected;
static void channels_factory_free(void *p){channels_free_calls++;channels_free_matches=p==channels_free_expected;free(p);}
#define OVERRIDE_OPUS_FREE 1
#define opus_free channels_factory_free
#define OVERRIDE_OPUS_ALLOC 1
#define opus_alloc channels_factory_alloc
#define VAR_ARRAYS 1
#define ENABLE_ASSERTIONS 1
#define opus_decoder_init validation_decoder_init
#define opus_decoder_get_size validation_decoder_get_size
#define opus_decoder_ctl validation_decoder_ctl
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
extern void validation_restore_decoder_mode(void *);
void comparison_restore_ms_modes(void *base,int *streams,int *coupled) {OpusMSDecoder *st=base;*streams=st->layout.nb_streams;*coupled=st->layout.nb_coupled_streams;char *ptr=(char*)st+align(sizeof(*st));for(int i=0;i<*streams;i++){validation_restore_decoder_mode(ptr);ptr+=align(validation_decoder_get_size(i<*coupled?2:1));}}
static int native_ms_ctl(unsigned char *data,size_t size,int request,int value,int alias,unsigned *output) {
 OpusMSDecoder *st=malloc(size);memcpy(st,data,size);int streams=st->layout.nb_streams,coupled=st->layout.nb_coupled_streams;char *ptr=(char*)st+align(sizeof(*st));for(int i=0;i<streams;i++){validation_restore_decoder_mode(ptr);ptr+=align(validation_decoder_get_size(i<coupled?2:1));}
 unsigned out=77;OpusDecoder *decoder=NULL;void *p=alias==-2?NULL:alias>=0?(void*)((char*)st+alias):(void*)&out;int result;
 switch(request) {
 case OPUS_SET_GAIN_REQUEST:case OPUS_SET_COMPLEXITY_REQUEST:case OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST:result=comparison_channels_ctl(st,request,value);break;
 case OPUS_MULTISTREAM_GET_DECODER_STATE_REQUEST:result=comparison_channels_ctl(st,request,value,alias==-2?NULL:&decoder);if(decoder)out=(unsigned)((char*)decoder-(char*)st);break;
 case OPUS_RESET_STATE:result=comparison_channels_ctl(st,request);break;
 default:result=comparison_channels_ctl(st,request,p);break;
 }
 comparison_normalize_ms_modes(st,streams,coupled);memcpy(data,st,size);free(st);*output=out;return result;
}
static int native_ms_destroy(int null) {void *p=null?NULL:malloc(sizeof(OpusMSDecoder));channels_free_calls=0;channels_free_expected=p;comparison_channels_destroy(p);return channels_free_calls==1&&channels_free_matches;}
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

func nativeMSCtl(data []byte, request, value, alias int32) (int32, uint32) {
	var out C.uint
	r := C.native_ms_ctl((*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.size_t(len(data)), C.int(request), C.int(value), C.int(alias), &out)
	return int32(r), uint32(out)
}

func nativeMSDestroy(null bool) bool {
	var n C.int
	if null {
		n = 1
	}
	return C.native_ms_destroy(n) != 0
}

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
