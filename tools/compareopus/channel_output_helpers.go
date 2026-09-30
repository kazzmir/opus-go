//go:build compareopus && cgo

package main

/*
#define VAR_ARRAYS 1
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
static void native_channel_output(void *dst,int ds,int dc,const float *src,int ss,int n,int op) {
 if(op==0) opus_copy_channel_out_float(dst,ds,dc,src,ss,n,NULL);
 else if(op==1) opus_copy_channel_out_short(dst,ds,dc,src,ss,n,NULL);
 else opus_copy_channel_out_int24(dst,ds,dc,src,ss,n,NULL);
}
*/
import "C"
import "unsafe"

func nativeChannelOutput(dst unsafe.Pointer, ds, dc int32, src []float32, ss, n int32, operations ...int32) {
	var op int32
	if len(operations) > 0 {
		op = operations[0]
	}
	C.native_channel_output(dst, C.int(ds), C.int(dc), (*C.float)(unsafe.Pointer(unsafe.SliceData(src))), C.int(ss), C.int(n), C.int(op))
}
