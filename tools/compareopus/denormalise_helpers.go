//go:build compareopus && cgo

package main

/*
#define VAR_ARRAYS 1
#include "modes.h"
#include "bands.h"
static void band_energies(const opus_int16 *bands,int nb,int size,const float *freq,float *out,int end,int C,int LM) {
 CELTMode m={0};m.nbEBands=nb;m.shortMdctSize=size;m.eBands=bands;
 compute_band_energies(&m,freq,out,end,C,LM,0);
}
static void normalise(const opus_int16 *bands,int nb,int size,const float *freq,float *out,const float *energy,int end,int C,int M) {
 CELTMode m={0};m.nbEBands=nb;m.shortMdctSize=size;m.eBands=bands;
 normalise_bands(&m,freq,out,energy,end,C,M);
}
static void denormalise(const opus_int16 *bands, int nb, int size, const float *x, float *out,
 const float *energy, int start, int end, int M, int downsample, int silence) {
 CELTMode m = {0}; m.nbEBands=nb; m.shortMdctSize=size; m.eBands=bands;
 denormalise_bands(&m,x,out,energy,start,end,M,downsample,silence);
}
*/
import "C"
import "unsafe"

func nativeBandEnergies(bands []int16, nb, size int32, freq, out []float32, end, channels, LM int32) {
	C.band_energies((*C.opus_int16)(unsafe.Pointer(unsafe.SliceData(bands))), C.int(nb), C.int(size), (*C.float)(unsafe.Pointer(unsafe.SliceData(freq))), (*C.float)(unsafe.Pointer(unsafe.SliceData(out))), C.int(end), C.int(channels), C.int(LM))
}

func nativeNormalise(bands []int16, nb, size int32, freq, out, energy []float32, end, channels, M int32) {
	C.normalise((*C.opus_int16)(unsafe.Pointer(unsafe.SliceData(bands))), C.int(nb), C.int(size), (*C.float)(unsafe.Pointer(unsafe.SliceData(freq))), (*C.float)(unsafe.Pointer(unsafe.SliceData(out))), (*C.float)(unsafe.Pointer(unsafe.SliceData(energy))), C.int(end), C.int(channels), C.int(M))
}

func nativeDenormalise(bands []int16, size int32, x, out, energy []float32, start, end, M, downsample, silence int32) {
	C.denormalise((*C.opus_int16)(unsafe.Pointer(&bands[0])), C.int(len(bands)-1), C.int(size), (*C.float)(unsafe.Pointer(unsafe.SliceData(x))), (*C.float)(unsafe.Pointer(&out[0])), (*C.float)(unsafe.Pointer(unsafe.SliceData(energy))), C.int(start), C.int(end), C.int(M), C.int(downsample), C.int(silence))
}
