//go:build compareopus && cgo

package main

/*
#include <stdlib.h>
#include <string.h>
#include "arch.h"
#include "../../../opus/src/mapping_matrix.h"
static MappingMatrix *compare_new_matrix(int rows,int cols,const short *data) {
 MappingMatrix *m=(MappingMatrix*)calloc(1,mapping_matrix_get_size(rows,cols));
 mapping_matrix_init(m,rows,cols,0,data,2*rows*cols); return m;
}
static void compare_matrix_float(int rows,int cols,const short *data,const float *input,int col,int istride,float *output,int ostride,int frames) {
 MappingMatrix *m=compare_new_matrix(rows,cols,data);
 mapping_matrix_multiply_channel_out_float(m,input,col,istride,output,ostride,frames);free(m);
}
static void compare_matrix_short(int rows,int cols,const short *data,const float *input,int col,int istride,short *output,int ostride,int frames) {
 MappingMatrix *m=compare_new_matrix(rows,cols,data);
 mapping_matrix_multiply_channel_out_short(m,input,col,istride,output,ostride,frames);free(m);
}
static void compare_matrix_int24(int rows,int cols,const short *data,const float *input,int col,int istride,int *output,int ostride,int frames) {
 MappingMatrix *m=compare_new_matrix(rows,cols,data);
 mapping_matrix_multiply_channel_out_int24(m,input,col,istride,output,ostride,frames);free(m);
}
static void compare_matrix_init(int rows,int cols,int gain,const short *data,short *out,int *meta) {
 int size=mapping_matrix_get_size(rows,cols);
 MappingMatrix *m=(MappingMatrix*)calloc(1,size);
 mapping_matrix_init(m,rows,cols,gain,data,2*rows*cols);
 meta[0]=m->rows; meta[1]=m->cols; meta[2]=m->gain;
 meta[3]=(char*)mapping_matrix_get_data(m)-(char*)m; meta[4]=size;
 if(rows*cols) memcpy(out,mapping_matrix_get_data(m),2*rows*cols);
 free(m);
}
*/
import "C"
import "unsafe"

func nativeMatrixFloat(rows, cols int32, data []int16, input []float32, col, istride int32, output []float32, ostride, frames int32) {
	C.compare_matrix_float(C.int(rows), C.int(cols), (*C.short)(unsafe.Pointer(&data[0])), (*C.float)(unsafe.Pointer(unsafe.SliceData(input))), C.int(col), C.int(istride), (*C.float)(unsafe.Pointer(unsafe.SliceData(output))), C.int(ostride), C.int(frames))
}

func nativeMatrixShort(rows, cols int32, data []int16, input []float32, col, istride int32, output []int16, ostride, frames int32) {
	C.compare_matrix_short(C.int(rows), C.int(cols), (*C.short)(unsafe.Pointer(&data[0])), (*C.float)(unsafe.Pointer(unsafe.SliceData(input))), C.int(col), C.int(istride), (*C.short)(unsafe.Pointer(unsafe.SliceData(output))), C.int(ostride), C.int(frames))
}

func nativeMatrixInt24(rows, cols int32, data []int16, input []float32, col, istride int32, output []int32, ostride, frames int32) {
	C.compare_matrix_int24(C.int(rows), C.int(cols), (*C.short)(unsafe.Pointer(&data[0])), (*C.float)(unsafe.Pointer(unsafe.SliceData(input))), C.int(col), C.int(istride), (*C.int)(unsafe.Pointer(unsafe.SliceData(output))), C.int(ostride), C.int(frames))
}

func nativeMatrixInit(rows, cols, gain int32, data []int16) ([]int16, [5]int32) {
	out := make([]int16, len(data))
	var meta [5]int32
	C.compare_matrix_init(C.int(rows), C.int(cols), C.int(gain), (*C.short)(unsafe.Pointer(unsafe.SliceData(data))), (*C.short)(unsafe.Pointer(unsafe.SliceData(out))), (*C.int)(unsafe.Pointer(&meta[0])))
	return out, meta
}
