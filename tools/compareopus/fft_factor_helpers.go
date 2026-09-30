//go:build compareopus && cgo

package main

/*
#define kf_work compare_factor_work
#define kf_factor compare_factor_factor
#define mini_kiss_fft_alloc compare_factor_fft_alloc
#define mini_kiss_fft_stride compare_factor_fft_stride
#define mini_kiss_fft compare_factor_fft
#define mini_kiss_fftr_alloc compare_factor_fftr_alloc
#define mini_kiss_fftr compare_factor_fftr
#include "../../../opus/celt/mini_kfft.c"
static int native_factor(int n,int *factors) {
 kf_factor(n,factors);
 int used=0;
 do {used+=2;} while(factors[used-1]>1);
 return used;
}
*/
import "C"
import "unsafe"

func nativeFFTFactor(n int32, factors *[64]int32) int {
	return int(C.native_factor(C.int(n), (*C.int)(unsafe.Pointer(factors))))
}
