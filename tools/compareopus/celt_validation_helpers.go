//go:build compareopus && cgo

package main

/*
#include <setjmp.h>
static _Thread_local jmp_buf celt_validation_jump;
void comparison_celt_validator_fatal(const char *str,const char *file,int line) {longjmp(celt_validation_jump,1);}
#define celt_fatal comparison_celt_validator_fatal
#define VAR_ARRAYS 1
#define OPUS_BUILD 1
#define ENABLE_ASSERTIONS 1
#define OPUS_DISABLE_INTRINSICS 1
#define validate_celt_decoder comparison_celt_validate
#define celt_decoder_get_size comparison_celt_size
#define opus_custom_decoder_get_size comparison_custom_size
#define opus_custom_decoder_create comparison_custom_create
#define celt_decoder_init comparison_celt_init
#define opus_custom_decoder_init comparison_custom_init
#define opus_custom_decoder_destroy comparison_custom_destroy
#define celt_synthesis comparison_celt_synthesis
#define celt_decode_with_ec_dred comparison_celt_decode_dred
#define celt_decode_with_ec comparison_celt_decode_ec
#define opus_custom_decode comparison_custom_decode
#define opus_custom_decode24 comparison_custom_decode24
#define opus_custom_decode_float comparison_custom_decode_float
#define opus_custom_decoder_ctl comparison_custom_ctl
#include "../../../opus/celt/celt_decoder.c"
static int native_custom_size(int overlap,int bands,int channels) {CELTMode mode={0};mode.overlap=overlap;mode.nbEBands=bands;return comparison_custom_size(&mode,channels);}
static int native_celt_validation(const int *v,int badmode) {
 CELTDecoder st={0};st.mode=badmode?NULL:opus_custom_mode_create(48000,960,NULL);
 st.overlap=v[0];st.end=v[1];st.channels=v[2];st.stream_channels=v[3];st.downsample=v[4];st.start=v[5];st.arch=v[6];st.last_pitch_index=v[7];st.postfilter_period=v[8];st.postfilter_period_old=v[9];st.postfilter_tapset=v[10];st.postfilter_tapset_old=v[11];
 if(setjmp(celt_validation_jump)) return 1;
 comparison_celt_validate(&st);return 0;
}
*/
import "C"
import (
	"github.com/kazzmir/opus-go/opuscc"
	"unsafe"
)

func nativeCustomDecoderSize(overlap, bands, channels int32) int32 {
	return int32(C.native_custom_size(C.int(overlap), C.int(bands), C.int(channels)))
}

func nativeCeltValidation(st *opuscc.OpusT_OpusCustomDecoder) bool {
	mode, _ := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	badmode := int32(0)
	if st.Fmode != mode {
		badmode = 1
	}
	v := [12]int32{st.Foverlap, st.Fend, st.Fchannels, st.Fstream_channels, st.Fdownsample, st.Fstart, st.Farch, st.Flast_pitch_index, st.Fpostfilter_period, st.Fpostfilter_period_old, st.Fpostfilter_tapset, st.Fpostfilter_tapset_old}
	return C.native_celt_validation((*C.int)(unsafe.Pointer(&v[0])), C.int(badmode)) != 0
}
