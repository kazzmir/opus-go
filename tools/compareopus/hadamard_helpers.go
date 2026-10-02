//go:build compareopus && cgo

package main

/*
#include <setjmp.h>
static _Thread_local jmp_buf hadamard_jump;
void comparison_hadamard_fatal(const char *str,const char *file,int line){longjmp(hadamard_jump,1);}
#define celt_fatal comparison_hadamard_fatal
#define ENABLE_ASSERTIONS 1
#define FLOAT_APPROX 1
#define OPUS_DISABLE_INTRINSICS 1
static void scalar_anti_renormalise(float *x,int n,float gain,int arch);
#define renormalise_vector scalar_anti_renormalise
#define VAR_ARRAYS 1
#define celt_lcg_rand compare_hadamard_lcg_rand
#define hysteresis_decision compare_hadamard_hysteresis
#define bitexact_cos compare_hadamard_cos
#define bitexact_log2tan compare_hadamard_log2tan
#define compute_band_energies compare_hadamard_energies
#define normalise_bands compare_hadamard_normalise
#define denormalise_bands compare_hadamard_denormalise
#define anti_collapse compare_hadamard_anti_collapse
#define spreading_decision compare_hadamard_spreading
#define haar1 compare_hadamard_haar1
#define quant_all_bands compare_hadamard_quant_all_bands
#include "../../../opus/celt/bands.c"
// Use the scalar vq.c normalization path, not the linked build's presumed SSE.
static void scalar_anti_renormalise(float *x,int n,float gain,int arch) {
 float energy=EPSILON+celt_inner_prod_c(x,x,n);float g=celt_rsqrt(energy)*gain;for(int i=0;i<n;i++)x[i]=g*x[i];
}
static void native_band_ctx_layout(size_t *v) {v[0]=sizeof(struct band_ctx);v[1]=offsetof(struct band_ctx,m);v[2]=offsetof(struct band_ctx,ec);v[3]=offsetof(struct band_ctx,bandE);}
static unsigned native_stereo_band(unsigned *s,unsigned char *buf,float *x,float *y,float *low,const int *cfg,int *meta) {
 ec_ctx ec={0};ec.buf=buf;ec.storage=s[0];ec.end_offs=s[1];ec.end_window=s[2];ec.nend_bits=s[3];ec.nbits_total=s[4];ec.offs=s[5];ec.rng=s[6];ec.val=s[7];ec.ext=s[8];ec.rem=s[9];ec.error=s[10];struct band_ctx ctx={0};ctx.m=opus_custom_mode_create(48000,960,NULL);ctx.ec=&ec;ctx.resynth=cfg[4];ctx.tf_change=cfg[5];ctx.remaining_bits=cfg[6];ctx.seed=123456;ctx.intensity=cfg[8];
 unsigned mask=quant_band_stereo(&ctx,x,y,cfg[0],cfg[3],cfg[1],NULL,cfg[2],low,NULL,cfg[7]);meta[0]=ctx.remaining_bits;meta[1]=ctx.seed;
 s[0]=ec.storage;s[1]=ec.end_offs;s[2]=ec.end_window;s[3]=ec.nend_bits;s[4]=ec.nbits_total;s[5]=ec.offs;s[6]=ec.rng;s[7]=ec.val;s[8]=ec.ext;s[9]=ec.rem;s[10]=ec.error;return mask;
}
static unsigned native_mono_band(unsigned *s,unsigned char *buf,float *x,float *low,const int *cfg,int *meta) {
 ec_ctx ec={0};ec.buf=buf;ec.storage=s[0];ec.end_offs=s[1];ec.end_window=s[2];ec.nend_bits=s[3];ec.nbits_total=s[4];ec.offs=s[5];ec.rng=s[6];ec.val=s[7];ec.ext=s[8];ec.rem=s[9];ec.error=s[10];struct band_ctx ctx={0};ctx.m=opus_custom_mode_create(48000,960,NULL);ctx.ec=&ec;ctx.resynth=cfg[4];ctx.tf_change=cfg[5];ctx.remaining_bits=cfg[6];ctx.seed=123456;
 unsigned mask=quant_band(&ctx,x,cfg[0],cfg[3],cfg[1],NULL,cfg[2],low,1,NULL,cfg[7]);meta[0]=ctx.remaining_bits;meta[1]=ctx.seed;
 s[0]=ec.storage;s[1]=ec.end_offs;s[2]=ec.end_window;s[3]=ec.nend_bits;s[4]=ec.nbits_total;s[5]=ec.offs;s[6]=ec.rng;s[7]=ec.val;s[8]=ec.ext;s[9]=ec.rem;s[10]=ec.error;return mask;
}
static void native_theta(unsigned *s,unsigned char *buf,const int *cfg,int *out,float *x,float *y,int encode) {
 ec_ctx ec={0};ec.buf=buf;ec.storage=s[0];ec.end_offs=s[1];ec.end_window=s[2];ec.nend_bits=s[3];ec.nbits_total=s[4];ec.offs=s[5];ec.rng=s[6];ec.val=s[7];ec.ext=s[8];ec.rem=s[9];ec.error=s[10];
 opus_int16 log[1]={cfg[8]};CELTMode m={0};m.nbEBands=1;m.logN=log;struct band_ctx ctx={0};float energy[2]={.8f,1.2f};ctx.m=&m;ctx.ec=&ec;ctx.bandE=energy;ctx.encode=encode;ctx.intensity=cfg[7];ctx.remaining_bits=cfg[9];ctx.disable_inv=cfg[10];struct split_ctx split={0};int b=cfg[5],fill=cfg[6];int *bp=&b,*fp=&fill;
 if(cfg[11]==1)fp=bp;else if(cfg[11]==2){ctx.remaining_bits=b;bp=&ctx.remaining_bits;}else if(cfg[11]==3){split.itheta=b;split.qalloc=fill;bp=&split.itheta;fp=&split.qalloc;}
 compute_theta(&ctx,&split,x,y,cfg[0],bp,cfg[1],cfg[2],cfg[3],cfg[4],fp);
 out[0]=split.inv;out[1]=split.imid;out[2]=split.iside;out[3]=split.delta;out[4]=split.itheta;out[5]=split.qalloc;out[6]=*bp;out[7]=*fp;out[8]=ctx.remaining_bits;
 s[0]=ec.storage;s[1]=ec.end_offs;s[2]=ec.end_window;s[3]=ec.nend_bits;s[4]=ec.nbits_total;s[5]=ec.offs;s[6]=ec.rng;s[7]=ec.val;s[8]=ec.ext;s[9]=ec.rem;s[10]=ec.error;
}
static unsigned native_quant_n1(unsigned *s,unsigned char *buf,float *v,int encode,int resynth,int *remaining,int y,int low,int done) {
 ec_ctx ec={0};ec.buf=buf;ec.storage=s[0];ec.end_offs=s[1];ec.end_window=s[2];ec.nend_bits=s[3];ec.nbits_total=s[4];ec.offs=s[5];ec.rng=s[6];ec.val=s[7];ec.ext=s[8];ec.rem=s[9];ec.error=s[10];
 struct band_ctx ctx={0};ctx.encode=encode;ctx.resynth=resynth;ctx.remaining_bits=*remaining;ctx.ec=&ec;
 unsigned result=quant_band_n1(&ctx,v,y<0?NULL:v+y,low<0?NULL:v+low);if(done&&encode)ec_enc_done(&ec);*remaining=ctx.remaining_bits;
 s[0]=ec.storage;s[1]=ec.end_offs;s[2]=ec.end_window;s[3]=ec.nend_bits;s[4]=ec.nbits_total;s[5]=ec.offs;s[6]=ec.rng;s[7]=ec.val;s[8]=ec.ext;s[9]=ec.rem;s[10]=ec.error;return result;
}
static void native_anti_collapse(const short *bands,int nb,float *x,unsigned char *masks,int lm,int channels,int size,int start,int end,const float *energy,const float *p1,const float *p2,const int *pulses,unsigned seed,int encode) {
 CELTMode mode={0};mode.eBands=bands;mode.nbEBands=nb;anti_collapse(&mode,x,masks,lm,channels,size,start,end,energy,p1,p2,pulses,seed,encode,0);
}
static int native_spreading(const short *bands,int nb,int short_size,float *x,int *state,int avg,int hf,int tap,int last,int update,int end,int channels,int mult,const int *weights) {
 CELTMode m={0};m.eBands=bands;m.nbEBands=nb;m.shortMdctSize=short_size;
 if(setjmp(hadamard_jump))return -99;return spreading_decision(&m,x,state+avg,last,state+hf,state+tap,update,end,channels,mult,weights);
}
static void compare_interleave(float *x,int n0,int stride,int hadamard) {
 interleave_hadamard(x,n0,stride,hadamard);
}
static void compare_deinterleave(float *x,int n0,int stride,int hadamard) {
 deinterleave_hadamard(x,n0,stride,hadamard);
}
*/
import "C"
import "unsafe"
import "github.com/kazzmir/opus-go/opuscc"

func nativeBandContextLayout() [4]uint64 {
	var v [4]C.size_t
	C.native_band_ctx_layout(&v[0])
	return [4]uint64{uint64(v[0]), uint64(v[1]), uint64(v[2]), uint64(v[3])}
}

func nativeStereoBand(e *opuscc.OpusT_ec_ctx, buf []byte, x, y, low []float32, cfg [9]int32) (uint32, [2]uint32) {
	s := [11]C.uint{C.uint(e.Fstorage), C.uint(e.Fend_offs), C.uint(e.Fend_window), C.uint(e.Fnend_bits), C.uint(e.Fnbits_total), C.uint(e.Foffs), C.uint(e.Frng), C.uint(e.Fval), C.uint(e.Fext), C.uint(e.Frem), C.uint(e.Ferror1)}
	var c [9]C.int
	for i := range c {
		c[i] = C.int(cfg[i])
	}
	var meta [2]C.int
	mask := C.native_stereo_band(&s[0], (*C.uchar)(unsafe.Pointer(unsafe.SliceData(buf))), (*C.float)(unsafe.Pointer(unsafe.SliceData(x))), (*C.float)(unsafe.Pointer(unsafe.SliceData(y))), (*C.float)(unsafe.Pointer(unsafe.SliceData(low))), &c[0], &meta[0])
	e.Fstorage = uint32(s[0])
	e.Fend_offs = uint32(s[1])
	e.Fend_window = uint32(s[2])
	e.Fnend_bits = int32(s[3])
	e.Fnbits_total = int32(s[4])
	e.Foffs = uint32(s[5])
	e.Frng = uint32(s[6])
	e.Fval = uint32(s[7])
	e.Fext = uint32(s[8])
	e.Frem = int32(s[9])
	e.Ferror1 = int32(s[10])
	return uint32(mask), [2]uint32{uint32(meta[0]), uint32(meta[1])}
}

func nativeMonoBand(e *opuscc.OpusT_ec_ctx, buf []byte, x, low []float32, cfg [8]int32) (uint32, [2]uint32) {
	s := [11]C.uint{C.uint(e.Fstorage), C.uint(e.Fend_offs), C.uint(e.Fend_window), C.uint(e.Fnend_bits), C.uint(e.Fnbits_total), C.uint(e.Foffs), C.uint(e.Frng), C.uint(e.Fval), C.uint(e.Fext), C.uint(e.Frem), C.uint(e.Ferror1)}
	var c [8]C.int
	for i := range c {
		c[i] = C.int(cfg[i])
	}
	var meta [2]C.int
	mask := C.native_mono_band(&s[0], (*C.uchar)(unsafe.Pointer(unsafe.SliceData(buf))), (*C.float)(unsafe.Pointer(unsafe.SliceData(x))), (*C.float)(unsafe.Pointer(unsafe.SliceData(low))), &c[0], &meta[0])
	e.Fstorage = uint32(s[0])
	e.Fend_offs = uint32(s[1])
	e.Fend_window = uint32(s[2])
	e.Fnend_bits = int32(s[3])
	e.Fnbits_total = int32(s[4])
	e.Foffs = uint32(s[5])
	e.Frng = uint32(s[6])
	e.Fval = uint32(s[7])
	e.Fext = uint32(s[8])
	e.Frem = int32(s[9])
	e.Ferror1 = int32(s[10])
	return uint32(mask), [2]uint32{uint32(meta[0]), uint32(meta[1])}
}

func nativeTheta(e *opuscc.OpusT_ec_ctx, buf []byte, cfg [12]int32) [9]int32 {
	return nativeThetaSpectrum(e, buf, nil, nil, cfg, false)
}

func nativeThetaSpectrum(e *opuscc.OpusT_ec_ctx, buf []byte, x, y []float32, cfg [12]int32, encode bool) [9]int32 {
	s := [11]C.uint{C.uint(e.Fstorage), C.uint(e.Fend_offs), C.uint(e.Fend_window), C.uint(e.Fnend_bits), C.uint(e.Fnbits_total), C.uint(e.Foffs), C.uint(e.Frng), C.uint(e.Fval), C.uint(e.Fext), C.uint(e.Frem), C.uint(e.Ferror1)}
	var c [12]C.int
	for i := range c {
		c[i] = C.int(cfg[i])
	}
	var out [9]C.int
	var en C.int
	if encode {
		en = 1
	}
	C.native_theta(&s[0], (*C.uchar)(unsafe.Pointer(unsafe.SliceData(buf))), &c[0], &out[0], (*C.float)(unsafe.Pointer(unsafe.SliceData(x))), (*C.float)(unsafe.Pointer(unsafe.SliceData(y))), en)
	e.Fstorage = uint32(s[0])
	e.Fend_offs = uint32(s[1])
	e.Fend_window = uint32(s[2])
	e.Fnend_bits = int32(s[3])
	e.Fnbits_total = int32(s[4])
	e.Foffs = uint32(s[5])
	e.Frng = uint32(s[6])
	e.Fval = uint32(s[7])
	e.Fext = uint32(s[8])
	e.Frem = int32(s[9])
	e.Ferror1 = int32(s[10])
	var result [9]int32
	for i := range result {
		result[i] = int32(out[i])
	}
	return result
}

func nativeQuantN1(e *opuscc.OpusT_ec_ctx, buf []byte, v []float32, encode, resynth int32, remaining *int32, y, low int32) uint32 {
	s := [11]C.uint{C.uint(e.Fstorage), C.uint(e.Fend_offs), C.uint(e.Fend_window), C.uint(e.Fnend_bits), C.uint(e.Fnbits_total), C.uint(e.Foffs), C.uint(e.Frng), C.uint(e.Fval), C.uint(e.Fext), C.uint(e.Frem), C.uint(e.Ferror1)}
	result := C.native_quant_n1(&s[0], (*C.uchar)(unsafe.Pointer(unsafe.SliceData(buf))), (*C.float)(unsafe.Pointer(unsafe.SliceData(v))), C.int(encode), C.int(resynth), (*C.int)(unsafe.Pointer(remaining)), C.int(y), C.int(low), 1)
	e.Fstorage = uint32(s[0])
	e.Fend_offs = uint32(s[1])
	e.Fend_window = uint32(s[2])
	e.Fnend_bits = int32(s[3])
	e.Fnbits_total = int32(s[4])
	e.Foffs = uint32(s[5])
	e.Frng = uint32(s[6])
	e.Fval = uint32(s[7])
	e.Fext = uint32(s[8])
	e.Frem = int32(s[9])
	e.Ferror1 = int32(s[10])
	return uint32(result)
}

func nativeAntiCollapse(bands []int16, nb int32, x []float32, masks []byte, lm, channels, size, start, end int32, energy, p1, p2 []float32, pulses []int32, seed uint32, encode int32) {
	C.native_anti_collapse((*C.short)(unsafe.Pointer(unsafe.SliceData(bands))), C.int(nb), (*C.float)(unsafe.Pointer(unsafe.SliceData(x))), (*C.uchar)(unsafe.Pointer(unsafe.SliceData(masks))), C.int(lm), C.int(channels), C.int(size), C.int(start), C.int(end), (*C.float)(unsafe.Pointer(unsafe.SliceData(energy))), (*C.float)(unsafe.Pointer(unsafe.SliceData(p1))), (*C.float)(unsafe.Pointer(unsafe.SliceData(p2))), (*C.int)(unsafe.Pointer(unsafe.SliceData(pulses))), C.uint(seed), C.int(encode))
}

func nativeSpreading(bands []int16, nb, shortSize int32, x []float32, state []int32, avg, hf, tap, last, update, end, channels, mult int32, weights []int32) int32 {
	return int32(C.native_spreading((*C.short)(unsafe.Pointer(unsafe.SliceData(bands))), C.int(nb), C.int(shortSize), (*C.float)(unsafe.Pointer(unsafe.SliceData(x))), (*C.int)(unsafe.Pointer(unsafe.SliceData(state))), C.int(avg), C.int(hf), C.int(tap), C.int(last), C.int(update), C.int(end), C.int(channels), C.int(mult), (*C.int)(unsafe.Pointer(unsafe.SliceData(weights)))))
}

func nativeInterleaveHadamard(x []float32, n0, stride, hadamard int32) {
	C.compare_interleave((*C.float)(unsafe.Pointer(unsafe.SliceData(x))), C.int(n0), C.int(stride), C.int(hadamard))
}

func nativeDeinterleaveHadamard(x []float32, n0, stride, hadamard int32) {
	C.compare_deinterleave((*C.float)(unsafe.Pointer(unsafe.SliceData(x))), C.int(n0), C.int(stride), C.int(hadamard))
}
