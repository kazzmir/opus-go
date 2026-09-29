//go:build compareopus && cgo

package main

/*
#include <stdlib.h>
#include <string.h>
#include "arch.h"
#include "../../../opus/src/mapping_matrix.h"
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

func nativeMatrixInit(rows, cols, gain int32, data []int16) ([]int16, [5]int32) {
	out := make([]int16, len(data))
	var meta [5]int32
	C.compare_matrix_init(C.int(rows), C.int(cols), C.int(gain), (*C.short)(unsafe.Pointer(unsafe.SliceData(data))), (*C.short)(unsafe.Pointer(unsafe.SliceData(out))), (*C.int)(unsafe.Pointer(&meta[0])))
	return out, meta
}
