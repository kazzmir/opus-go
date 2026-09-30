//go:build compareopus && cgo

package main

/*
#define VAR_ARRAYS 1
#define opus_projection_decoder_get_size comparison_projection_size
#define opus_projection_decoder_init comparison_projection_init
#define opus_projection_decoder_create comparison_projection_create
#define opus_projection_decode comparison_projection_decode
#define opus_projection_decode24 comparison_projection_decode24
#define opus_projection_decode_float comparison_projection_decode_float
#define opus_projection_decoder_ctl comparison_projection_ctl
#define opus_projection_decoder_destroy comparison_projection_destroy
#include "../../../opus/src/opus_projection_decoder.c"
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
