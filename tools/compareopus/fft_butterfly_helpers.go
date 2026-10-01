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
#define clt_mdct_init compare_mdct_init
#define clt_mdct_clear compare_mdct_clear
#define clt_mdct_forward_c compare_mdct_forward
#define clt_mdct_backward_c compare_mdct_backward
#include "../../../opus/celt/mdct.c"
static void native_fft_layout(size_t *layout) {layout[0]=sizeof(kiss_fft_state);layout[1]=offsetof(kiss_fft_state,bitrev);layout[2]=offsetof(kiss_fft_state,twiddles);layout[3]=offsetof(kiss_fft_state,arch_fft);}
static int native_mdct_fixture(int n,int shifts,float *trig,int *sizes,size_t *layout) {
 mdct_lookup l={0};if(!compare_mdct_init(&l,n,shifts,0))return 0;
 memcpy(trig,l.trig,(n-((n/2)>>shifts))*sizeof(float));for(int i=0;i<=shifts;i++)sizes[i]=l.kfft[i]->nfft;
 layout[0]=sizeof(l);layout[1]=offsetof(mdct_lookup,kfft);layout[2]=offsetof(mdct_lookup,trig);compare_mdct_clear(&l,0);return 1;
}
static void native_mdct_transform(int n,int shift,int fftshift,float scale,const short *factors,const short *bitrev,const void *tw,const float *trig,float *in,float *out,const float *window,int overlap,int stride,int op) {
 kiss_fft_state st={0};st.nfft=n>>2>>shift;st.shift=fftshift;st.scale=scale;memcpy(st.factors,factors,sizeof(st.factors));st.bitrev=bitrev;st.twiddles=tw;mdct_lookup l={0};l.n=n;l.maxshift=shift;l.kfft[shift]=&st;l.trig=trig;
 if(op==0)compare_mdct_forward(&l,in,out,window,overlap,shift,stride,0);else compare_mdct_backward(&l,in,out,window,overlap,shift,stride,0);
}
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

func nativeFFTLayout() [4]uint64 {
	var data [4]C.size_t
	C.native_fft_layout(&data[0])
	return [4]uint64{uint64(data[0]), uint64(data[1]), uint64(data[2]), uint64(data[3])}
}

func nativeMDCTLookup(n, shifts int32) ([]float32, []int32, [3]uint64) {
	trig := make([]float32, n-((n/2)>>shifts))
	sizes := make([]int32, shifts+1)
	var layout [3]C.size_t
	if C.native_mdct_fixture(C.int(n), C.int(shifts), (*C.float)(unsafe.Pointer(&trig[0])), (*C.int)(unsafe.Pointer(&sizes[0])), &layout[0]) == 0 {
		panic("MDCT fixture allocation")
	}
	return trig, sizes, [3]uint64{uint64(layout[0]), uint64(layout[1]), uint64(layout[2])}
}

func nativeMDCTTransform(l *opuscc.OpusT_mdct_lookup, bitrev []int16, tw []opuscc.OpusT_kiss_twiddle_cpx, trig, in, out, window []float32, overlap, shift, stride, op int32) {
	st := l.Fkfft[shift]
	factors := st.Ffactors // Only scalar array memory crosses cgo, not table owners.
	C.native_mdct_transform(C.int(l.Fn), C.int(shift), C.int(st.Fshift), C.float(st.Fscale), (*C.short)(unsafe.Pointer(&factors[0])), (*C.short)(unsafe.Pointer(unsafe.SliceData(bitrev))), unsafe.Pointer(unsafe.SliceData(tw)), (*C.float)(unsafe.Pointer(unsafe.SliceData(trig))), (*C.float)(unsafe.Pointer(unsafe.SliceData(in))), (*C.float)(unsafe.Pointer(unsafe.SliceData(out))), (*C.float)(unsafe.Pointer(unsafe.SliceData(window))), C.int(overlap), C.int(stride), C.int(op))
}

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
	factors := state.Ffactors
	C.native_fft_transform(C.int(state.Fnfft), C.int(state.Fshift), C.float(state.Fscale), (*C.short)(unsafe.Pointer(&factors[0])), (*C.short)(unsafe.Pointer(unsafe.SliceData(bitrev))), unsafe.Pointer(unsafe.SliceData(tw)), unsafe.Pointer(unsafe.SliceData(in)), unsafe.Pointer(unsafe.SliceData(out)), C.int(op))
}

func nativeFFTButterfly(radix int32, out []opuscc.OpusT_kiss_fft_cpx, tw []opuscc.OpusT_kiss_twiddle_cpx, stride uint64, m, N, mm int32) {
	C.native_bfly(C.int(radix), unsafe.Pointer(unsafe.SliceData(out)), unsafe.Pointer(unsafe.SliceData(tw)), C.size_t(stride), C.int(m), C.int(N), C.int(mm))
}
