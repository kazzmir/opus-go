//go:build compareopus && cgo

package main

/*
#include <stdlib.h>
static _Thread_local int projection_factory_fail;
static void *projection_factory_alloc(size_t size){return projection_factory_fail?NULL:calloc(1,size);}
static _Thread_local int projection_free_calls,projection_free_matches;
static _Thread_local void *projection_free_expected;
static void projection_factory_free(void *p){projection_free_calls++;projection_free_matches=p==projection_free_expected;free(p);}
#define OVERRIDE_OPUS_FREE 1
#define opus_free projection_factory_free
#define OVERRIDE_OPUS_ALLOC 1
#define opus_alloc projection_factory_alloc
#define VAR_ARRAYS 1
#define opus_multistream_decoder_init comparison_channels_init
#define opus_multistream_decoder_get_size comparison_channels_get_size
#define opus_projection_decoder_get_size comparison_projection_size
#define opus_projection_decoder_init comparison_projection_init
#define opus_projection_decoder_create comparison_projection_create
#define opus_projection_decode comparison_projection_decode
#define opus_projection_decode24 comparison_projection_decode24
#define opus_projection_decode_float comparison_projection_decode_float
#define opus_projection_decoder_ctl comparison_projection_ctl
#define opus_projection_decoder_destroy comparison_projection_destroy
#include "../../../opus/src/opus_projection_decoder.c"
extern void comparison_normalize_ms_modes(void *,int,int);
static int native_projection_destroy(int null) {void *p=null?NULL:malloc(sizeof(OpusProjectionDecoder));projection_free_calls=0;projection_free_expected=p;comparison_projection_destroy(p);return projection_free_calls==1&&projection_free_matches;}
static int native_projection_create_image(unsigned char *data,size_t size,int rate,int channels,int streams,int coupled,unsigned char *matrix,int bytes,int fail) {
 int error=99;projection_factory_fail=fail;OpusProjectionDecoder *st=comparison_projection_create(rate,channels,streams,coupled,matrix,bytes,&error);projection_factory_fail=0;
 if(st){comparison_normalize_ms_modes(get_multistream_decoder(st),streams,coupled);memcpy(data,st,size);comparison_projection_destroy(st);}return error;
}
static int native_projection_init_image(unsigned char *data,size_t size,int rate,int channels,int streams,int coupled,unsigned char *matrix,int bytes,int offset) {
 OpusProjectionDecoder *st=malloc(size);memcpy(st,data,size);if(offset>=0)matrix=(unsigned char*)st+offset;
 int result=comparison_projection_init(st,rate,channels,streams,coupled,matrix,bytes);if(result==OPUS_OK)comparison_normalize_ms_modes(get_multistream_decoder(st),streams,coupled);
 memcpy(data,st,size);free(st);return result;
}
static void native_projection_output(void *dst,int ds,int dc,const float *src,int ss,int n,void *matrix,int op) {
 if(op==0) opus_projection_copy_channel_out_float(dst,ds,dc,src,ss,n,matrix);
 else if(op==1) opus_projection_copy_channel_out_short(dst,ds,dc,src,ss,n,matrix);
 else opus_projection_copy_channel_out_int24(dst,ds,dc,src,ss,n,matrix);
}
static size_t native_projection_multistream(void *base) {return (char *)get_multistream_decoder((OpusProjectionDecoder *)base)-(char *)base;}
static size_t native_projection_matrix(void *base,int *fields) {
 MappingMatrix *matrix=get_dec_demixing_matrix((OpusProjectionDecoder *)base);
 fields[0]=matrix->rows;fields[1]=matrix->cols;fields[2]=matrix->gain;
 return (char *)matrix-(char *)base;
}
*/
import "C"
import "unsafe"

func nativeProjectionDestroy(null bool) bool {
	var n C.int
	if null {
		n = 1
	}
	return C.native_projection_destroy(n) != 0
}

func nativeProjectionCreateImage(data []byte, rate, channels, streams, coupled int32, matrix []byte, bytes int32, fail bool) int32 {
	f := C.int(0)
	if fail {
		f = 1
	}
	return int32(C.native_projection_create_image((*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.size_t(len(data)), C.int(rate), C.int(channels), C.int(streams), C.int(coupled), (*C.uchar)(unsafe.Pointer(unsafe.SliceData(matrix))), C.int(bytes), f))
}

func nativeProjectionInitImage(data []byte, rate, channels, streams, coupled int32, matrix []byte, bytes, offset int32) int32 {
	return int32(C.native_projection_init_image((*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.size_t(len(data)), C.int(rate), C.int(channels), C.int(streams), C.int(coupled), (*C.uchar)(unsafe.Pointer(unsafe.SliceData(matrix))), C.int(bytes), C.int(offset)))
}
func nativeProjectionSize(channels, streams, coupled int32) int32 {
	return int32(C.comparison_projection_size(C.int(channels), C.int(streams), C.int(coupled)))
}

func nativeProjectionOutput(dst unsafe.Pointer, ds, dc int32, src []float32, ss, n int32, matrix unsafe.Pointer, op int32) {
	C.native_projection_output(dst, C.int(ds), C.int(dc), (*C.float)(unsafe.Pointer(unsafe.SliceData(src))), C.int(ss), C.int(n), matrix, C.int(op))
}

func nativeProjectionMultistream(base unsafe.Pointer) uintptr {
	return uintptr(C.native_projection_multistream(base))
}

func nativeProjectionMatrix(base unsafe.Pointer) (uintptr, [3]int32) {
	var fields [3]C.int
	offset := C.native_projection_matrix(base, &fields[0])
	return uintptr(offset), [3]int32{int32(fields[0]), int32(fields[1]), int32(fields[2])}
}
