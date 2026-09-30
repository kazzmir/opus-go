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
static void native_mini_fixture(int n,int inverse,void *out) {
 mini_kiss_fft_cfg st=mini_kiss_fft_alloc(n,inverse,NULL,NULL);
 memcpy(out,st,sizeof(mini_kiss_fft_state)+(n-1)*sizeof(mini_kiss_fft_cpx));free(st);
}
static void native_mini_transform(void *state,const void *in,void *out,int stride,int op) {
 mini_kiss_fft_cfg st=state;
 if(op==0) kf_work(out,in,1,stride,st->factors,st);
 else if(op==1) mini_kiss_fft_stride(st,in,out,stride);
 else mini_kiss_fft(st,in,out);
}
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

func nativeMiniAlloc(n, inverse int32, mem *byte, length *uint64) bool {
	cap := C.size_t(*length)
	st := C.compare_factor_fft_alloc(C.int(n), C.int(inverse), unsafe.Pointer(mem), &cap)
	*length = uint64(cap)
	return st != nil
}

func nativeMiniFixture(n, inverse int32) *opuscc.OpusT_mini_kiss_fft_state {
	backing := make([]uint64, (264+8*int(n)+7)/8)
	st := (*opuscc.OpusT_mini_kiss_fft_state)(unsafe.Pointer(&backing[0]))
	C.native_mini_fixture(C.int(n), C.int(inverse), unsafe.Pointer(st))
	return st
}

func nativeMiniTransform(st *opuscc.OpusT_mini_kiss_fft_state, in, out []opuscc.OpusT_mini_kiss_fft_cpx, stride, op int32) {
	C.native_mini_transform(unsafe.Pointer(st), unsafe.Pointer(unsafe.SliceData(in)), unsafe.Pointer(unsafe.SliceData(out)), C.int(stride), C.int(op))
}

func nativeMiniButterfly(radix int32, out, tw []opuscc.OpusT_mini_kiss_fft_cpx, stride, m uint64, inverse int32) {
	C.native_mini_butterfly(C.int(radix), unsafe.Pointer(unsafe.SliceData(out)), unsafe.Pointer(unsafe.SliceData(tw)), C.size_t(len(tw)), C.size_t(stride), C.size_t(m), C.int(inverse))
}

func nativeFFTFactor(n int32, factors *[64]int32) int {
	return int(C.native_factor(C.int(n), (*C.int)(unsafe.Pointer(factors))))
}
