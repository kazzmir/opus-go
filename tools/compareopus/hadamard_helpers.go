//go:build compareopus && cgo

package main

/*
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
static unsigned native_quant_n1(unsigned *s,unsigned char *buf,float *v,int encode,int resynth,int *remaining,int y,int low,int done) {
 ec_ctx ec={0};ec.buf=buf;ec.storage=s[0];ec.end_offs=s[1];ec.end_window=s[2];ec.nend_bits=s[3];ec.nbits_total=s[4];ec.offs=s[5];ec.rng=s[6];ec.val=s[7];ec.ext=s[8];ec.rem=s[9];ec.error=s[10];
 struct band_ctx ctx={0};ctx.encode=encode;ctx.resynth=resynth;ctx.remaining_bits=*remaining;ctx.ec=&ec;
 unsigned result=quant_band_n1(&ctx,v,y<0?NULL:v+y,low<0?NULL:v+low);if(done&&encode)ec_enc_done(&ec);*remaining=ctx.remaining_bits;
 s[0]=ec.storage;s[1]=ec.end_offs;s[2]=ec.end_window;s[3]=ec.nend_bits;s[4]=ec.nbits_total;s[5]=ec.offs;s[6]=ec.rng;s[7]=ec.val;s[8]=ec.ext;s[9]=ec.rem;s[10]=ec.error;return result;
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

func nativeInterleaveHadamard(x []float32, n0, stride, hadamard int32) {
	C.compare_interleave((*C.float)(unsafe.Pointer(unsafe.SliceData(x))), C.int(n0), C.int(stride), C.int(hadamard))
}

func nativeDeinterleaveHadamard(x []float32, n0, stride, hadamard int32) {
	C.compare_deinterleave((*C.float)(unsafe.Pointer(unsafe.SliceData(x))), C.int(n0), C.int(stride), C.int(hadamard))
}
