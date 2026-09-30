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
static size_t native_projection_matrix(void *base,int *fields) {
 MappingMatrix *matrix=get_dec_demixing_matrix((OpusProjectionDecoder *)base);
 fields[0]=matrix->rows;fields[1]=matrix->cols;fields[2]=matrix->gain;
 return (char *)matrix-(char *)base;
}
*/
import "C"
import "unsafe"

func nativeProjectionMatrix(base unsafe.Pointer) (uintptr, [3]int32) {
	var fields [3]C.int
	offset := C.native_projection_matrix(base, &fields[0])
	return uintptr(offset), [3]int32{int32(fields[0]), int32(fields[1]), int32(fields[2])}
}
