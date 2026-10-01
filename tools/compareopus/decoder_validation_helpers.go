//go:build compareopus && cgo

package main

/*
#include <setjmp.h>
#include <stdlib.h>
static _Thread_local int validation_factory_fail;
static void *validation_factory_alloc(size_t size){return validation_factory_fail?NULL:calloc(1,size);}
static _Thread_local int validation_free_calls,validation_free_matches;
static _Thread_local void *validation_free_expected;
static void validation_factory_free(void *p){validation_free_calls++;validation_free_matches=p==validation_free_expected;free(p);}
#define OVERRIDE_OPUS_FREE 1
#define opus_free validation_factory_free
#define OVERRIDE_OPUS_ALLOC 1
#define opus_alloc validation_factory_alloc
static _Thread_local jmp_buf validation_jump;
void comparison_validator_fatal(const char *str,const char *file,int line) {longjmp(validation_jump,1);}
#define celt_fatal comparison_validator_fatal
#define VAR_ARRAYS 1
#define OPUS_BUILD 1
#define ENABLE_ASSERTIONS 1
#define OPUS_DISABLE_INTRINSICS 1
#define celt_decoder_init comparison_celt_init
#define silk_init_decoder validation_silk_state_init
#define silk_reset_decoder validation_silk_state_reset
#define silk_Get_Decoder_Size validation_silk_api_size
#define silk_InitDecoder validation_silk_api_init
#define silk_ResetDecoder validation_silk_api_reset
#define silk_Decode validation_silk_api_decode
#define silk_LoadOSCEModels validation_silk_api_models
#define opus_decoder_get_size validation_decoder_get_size
#define opus_decoder_init validation_decoder_init
#define opus_decoder_create validation_decoder_create
#define opus_decode_native validation_decode_native
#define opus_decode validation_decode
#define opus_decode24 validation_decode24
#define opus_decode_float validation_decode_float
#define opus_decoder_ctl validation_decoder_ctl
#define opus_decoder_destroy validation_decoder_destroy
#define opus_packet_get_bandwidth validation_packet_bandwidth
#define opus_packet_get_nb_channels validation_packet_channels
#define opus_packet_get_nb_frames validation_packet_frames
#define opus_packet_get_nb_samples validation_packet_samples
#define opus_packet_has_lbrr validation_packet_lbrr
#define opus_decoder_get_nb_samples validation_decoder_samples
#define opus_dred_decoder_get_size validation_dred_decoder_size
#define opus_dred_decoder_init validation_dred_decoder_init
#define opus_dred_decoder_create validation_dred_decoder_create
#define opus_dred_decoder_destroy validation_dred_decoder_destroy
#define opus_dred_decoder_ctl validation_dred_decoder_ctl
#define opus_dred_get_size validation_dred_size
#define opus_dred_alloc validation_dred_alloc
#define opus_dred_free validation_dred_free
#define opus_dred_parse validation_dred_parse
#define opus_dred_process validation_dred_process
#define opus_decoder_dred_decode validation_dred_decode
#define opus_decoder_dred_decode24 validation_dred_decode24
#define opus_decoder_dred_decode_float validation_dred_decode_float
#include "../../../opus/src/opus_decoder.c"
#include "../../../opus/silk/init_decoder.c"
#include "../../../opus/silk/dec_API.c"
void validation_normalize_decoder_mode(void *decoder) {OpusDecoder *st=decoder;memset((char*)st+st->celt_dec_offset,0,sizeof(void*));}
static int native_opus_destroy(int null) {
 void *p=null?NULL:malloc(sizeof(OpusDecoder));validation_free_calls=0;validation_free_expected=p;validation_decoder_destroy(p);return validation_free_calls==1&&validation_free_matches;
}
static int native_opus_create_image(unsigned char *data,size_t size,int rate,int channels,int fail) {
 int error=99;validation_factory_fail=fail;OpusDecoder *st=validation_decoder_create(rate,channels,&error);validation_factory_fail=0;
 if(st){validation_normalize_decoder_mode(st);memcpy(data,st,size);validation_decoder_destroy(st);}return error;
}
static int native_opus_init_image(unsigned char *data,size_t size,int rate,int channels) {
 OpusDecoder *st=malloc(size);memcpy(st,data,size);int result=validation_decoder_init(st,rate,channels);
 if(result==OPUS_OK)memset((char*)st+st->celt_dec_offset,0,sizeof(void*));memcpy(data,st,size);free(st);return result;
}
static int native_opus_validation(const int *v) {
 OpusDecoder st={0};st.channels=v[0];st.Fs=v[1];st.DecControl.API_sampleRate=v[2];st.DecControl.internalSampleRate=v[3];st.DecControl.nChannelsAPI=v[4];st.DecControl.nChannelsInternal=v[5];st.DecControl.payloadSize_ms=v[6];st.arch=v[7];st.stream_channels=v[8];
 if(setjmp(validation_jump)) return 1;
 validate_opus_decoder(&st);return 0;
}
*/
import "C"
import (
	"github.com/kazzmir/opus-go/opuscc"
	"unsafe"
)

func nativeOpusDestroy(null bool) bool {
	var n C.int
	if null {
		n = 1
	}
	return C.native_opus_destroy(n) != 0
}

func nativeOpusCreateImage(data []byte, rate, channels int32, fail bool) int32 {
	f := C.int(0)
	if fail {
		f = 1
	}
	return int32(C.native_opus_create_image((*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.size_t(len(data)), C.int(rate), C.int(channels), f))
}

func nativeOpusInitImage(data []byte, rate, channels int32) int32 {
	return int32(C.native_opus_init_image((*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.size_t(len(data)), C.int(rate), C.int(channels)))
}
func nativeOpusSize(channels int32) int32 {
	return int32(C.validation_decoder_get_size(C.int(channels)))
}

func nativeOpusValidation(st *opuscc.OpusT_OpusDecoder) bool {
	dc := st.FDecControl
	v := [9]int32{st.Fchannels, st.FFs, dc.FAPI_sampleRate, dc.FinternalSampleRate, dc.FnChannelsAPI, dc.FnChannelsInternal, dc.FpayloadSize_ms, st.Farch, st.Fstream_channels}
	return C.native_opus_validation((*C.int)(unsafe.Pointer(&v[0]))) != 0
}
