//go:build compareopus && cgo

package main

/*
#include <opus_types.h>
void silk_insertion_sort_increasing(opus_int32 *a, int *idx, const int L, const int K);
void silk_insertion_sort_increasing_all_values_int16(opus_int16 *a, const int L);
*/
import "C"

import "unsafe"

func nativeSort32(a []int32, idx []int32) {
	C.silk_insertion_sort_increasing((*C.opus_int32)(unsafe.Pointer(&a[0])), (*C.int)(unsafe.Pointer(&idx[0])), C.int(len(a)), C.int(len(idx)))
}

func nativeSort16(a []int16) {
	C.silk_insertion_sort_increasing_all_values_int16((*C.opus_int16)(unsafe.Pointer(&a[0])), C.int(len(a)))
}
