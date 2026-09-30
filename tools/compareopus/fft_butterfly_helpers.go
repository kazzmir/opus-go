//go:build compareopus && cgo

package main

/*
#define VAR_ARRAYS 1
#define CUSTOM_MODES 1
#define opus_fft_alloc_arch_c compare_bfly_alloc_arch
#define opus_fft_alloc_twiddles compare_bfly_alloc_twiddles
#define opus_fft_alloc compare_bfly_alloc
#define opus_fft_free_arch_c compare_bfly_free_arch
#define opus_fft_free compare_bfly_free
#define opus_fft_impl compare_bfly_impl
#define opus_fft_c compare_bfly_fft
#define opus_ifft_c compare_bfly_ifft
#include "../../../opus/celt/kiss_fft.c"
static int native_fft_fixture(int n,int shift,short *factors,short *bitrev,void *tw,float *scale) {
 int total=n<<(shift>0?shift:0);
 kiss_fft_state *base=compare_bfly_alloc(total,NULL,NULL,0);
 if(!base) return 0;
 kiss_fft_state *st=shift>0?compare_bfly_alloc_twiddles(n,NULL,NULL,base,0):base;
 if(!st) {compare_bfly_free(base,0);return 0;}
 memcpy(factors,st->factors,sizeof(st->factors));memcpy(bitrev,st->bitrev,n*sizeof(*bitrev));memcpy(tw,base->twiddles,total*sizeof(kiss_twiddle_cpx));*scale=st->scale;
 if(st!=base) compare_bfly_free(st,0);
 compare_bfly_free(base,0);return 1;
}
static void native_fft_transform(int n,int shift,float scale,const short *factors,const short *bitrev,const void *tw,const void *in,void *out,int op) {
 kiss_fft_state st={0};st.nfft=n;st.shift=shift;st.scale=scale;memcpy(st.factors,factors,sizeof(st.factors));st.bitrev=bitrev;st.twiddles=tw;
 if(op==0) compare_bfly_impl(&st,out);
 else if(op==1) compare_bfly_fft(&st,in,out);
 else compare_bfly_ifft(&st,in,out);
}
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

func nativeFFTFixture(n, shift int32) (opuscc.OpusT_kiss_fft_state, []int16, []opuscc.OpusT_kiss_twiddle_cpx) {
	state := opuscc.OpusT_kiss_fft_state{Fnfft: n, Fshift: shift}
	bitrev := make([]int16, n)
	tw := make([]opuscc.OpusT_kiss_twiddle_cpx, int(n)<<max(shift, 0))
	ok := C.native_fft_fixture(C.int(n), C.int(shift), (*C.short)(unsafe.Pointer(&state.Ffactors[0])), (*C.short)(unsafe.Pointer(&bitrev[0])), unsafe.Pointer(&tw[0]), (*C.float)(unsafe.Pointer(&state.Fscale)))
	if ok == 0 {
		panic("native FFT fixture allocation")
	}
	return state, bitrev, tw
}

func nativeFFTTransform(state *opuscc.OpusT_kiss_fft_state, bitrev []int16, tw []opuscc.OpusT_kiss_twiddle_cpx, in, out []opuscc.OpusT_kiss_fft_cpx, op int32) {
	C.native_fft_transform(C.int(state.Fnfft), C.int(state.Fshift), C.float(state.Fscale), (*C.short)(unsafe.Pointer(&state.Ffactors[0])), (*C.short)(unsafe.Pointer(unsafe.SliceData(bitrev))), unsafe.Pointer(unsafe.SliceData(tw)), unsafe.Pointer(unsafe.SliceData(in)), unsafe.Pointer(unsafe.SliceData(out)), C.int(op))
}

func nativeFFTButterfly(radix int32, out []opuscc.OpusT_kiss_fft_cpx, tw []opuscc.OpusT_kiss_twiddle_cpx, stride uint64, m, N, mm int32) {
	C.native_bfly(C.int(radix), unsafe.Pointer(unsafe.SliceData(out)), unsafe.Pointer(unsafe.SliceData(tw)), C.size_t(stride), C.int(m), C.int(N), C.int(mm))
}
