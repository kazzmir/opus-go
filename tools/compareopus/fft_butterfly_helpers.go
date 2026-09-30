//go:build compareopus && cgo

package main

/*
#define VAR_ARRAYS 1
#define opus_fft_alloc_arch_c compare_bfly_alloc_arch
#define opus_fft_alloc_twiddles compare_bfly_alloc_twiddles
#define opus_fft_alloc compare_bfly_alloc
#define opus_fft_free_arch_c compare_bfly_free_arch
#define opus_fft_free compare_bfly_free
#define opus_fft_impl compare_bfly_impl
#define opus_fft_c compare_bfly_fft
#define opus_ifft_c compare_bfly_ifft
#include "../../../opus/celt/kiss_fft.c"
static void native_bfly(int radix,void *out,const void *tw,size_t stride,int m,int N,int mm) {
 kiss_fft_state st={0}; st.twiddles=(const kiss_twiddle_cpx *)tw;
 switch(radix) {
 case 2:kf_bfly2((kiss_fft_cpx *)out,m,N);break;
 case 4:kf_bfly4((kiss_fft_cpx *)out,stride,&st,m,N,mm);break;
 case 3:kf_bfly3((kiss_fft_cpx *)out,stride,&st,m,N,mm);break;
 case 5:kf_bfly5((kiss_fft_cpx *)out,stride,&st,m,N,mm);break;
 }
}
*/
import "C"
import (
	"github.com/kazzmir/opus-go/opuscc"
	"unsafe"
)

func nativeFFTButterfly(radix int32, out []opuscc.OpusT_kiss_fft_cpx, tw []opuscc.OpusT_kiss_twiddle_cpx, stride uint64, m, N, mm int32) {
	C.native_bfly(C.int(radix), unsafe.Pointer(unsafe.SliceData(out)), unsafe.Pointer(unsafe.SliceData(tw)), C.size_t(stride), C.int(m), C.int(N), C.int(mm))
}
