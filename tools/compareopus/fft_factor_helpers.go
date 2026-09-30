//go:build compareopus && cgo

package main

/*
#include <string.h>
#define kf_work compare_factor_work
#define kf_factor compare_factor_factor
#define mini_kiss_fft_alloc compare_factor_fft_alloc
#define mini_kiss_fft_stride compare_factor_fft_stride
#define mini_kiss_fft compare_factor_fft
#define mini_kiss_fftr_alloc compare_factor_fftr_alloc
#define mini_kiss_fftr compare_factor_fftr
#include "../../../opus/celt/mini_kfft.c"
static void native_mini_butterfly(int radix,void *out,const void *tw,size_t count,size_t stride,size_t m,int inverse) {
 mini_kiss_fft_cfg st=malloc(sizeof(mini_kiss_fft_state)+(count-1)*sizeof(mini_kiss_fft_cpx));
 st->inverse=inverse;memcpy(st->twiddles,tw,count*sizeof(mini_kiss_fft_cpx));
 switch(radix) {
 case 2:kf_bfly2(out,stride,st,(int)m);break;
 case 3:kf_bfly3(out,stride,st,m);break;
 case 4:kf_bfly4(out,stride,st,m);break;
 case 5:kf_bfly5(out,stride,st,(int)m);break;
 }
 free(st);
}
static int native_factor(int n,int *factors) {
 kf_factor(n,factors);
 int used=0;
 do {used+=2;} while(factors[used-1]>1);
 return used;
}
*/
import "C"
import "unsafe"
import "github.com/kazzmir/opus-go/opuscc"

func nativeMiniButterfly(radix int32, out, tw []opuscc.OpusT_mini_kiss_fft_cpx, stride, m uint64, inverse int32) {
	C.native_mini_butterfly(C.int(radix), unsafe.Pointer(unsafe.SliceData(out)), unsafe.Pointer(unsafe.SliceData(tw)), C.size_t(len(tw)), C.size_t(stride), C.size_t(m), C.int(inverse))
}

func nativeFFTFactor(n int32, factors *[64]int32) int {
	return int(C.native_factor(C.int(n), (*C.int)(unsafe.Pointer(factors))))
}
