//go:build compareopus && cgo

package main

/*
#include "modes.h"
#include "quant_bands.h"
static void energy_decode(unsigned *s, unsigned char *data, float *energy, int bands,
 int start, int end, int channels, int op, int a, int b, int *quant, int *priority) {
 CELTMode m = {0}; m.nbEBands=bands;
 ec_dec dec = {0}; dec.buf=data;
 dec.storage=s[0]; dec.end_offs=s[1]; dec.end_window=s[2];
 dec.nend_bits=(int)s[3]; dec.nbits_total=(int)s[4]; dec.offs=s[5];
 dec.rng=s[6]; dec.val=s[7]; dec.ext=s[8]; dec.rem=(int)s[9]; dec.error=(int)s[10];
 switch(op) {
 case 0: unquant_coarse_energy(&m,start,end,energy,a,&dec,channels,b); break;
 case 1: unquant_fine_energy(&m,start,end,energy,quant,priority,&dec,channels); break;
 case 2: unquant_energy_finalise(&m,start,end,energy,quant,priority,a,&dec,channels); break;
 }
 s[0]=dec.storage; s[1]=dec.end_offs; s[2]=dec.end_window;
 s[3]=dec.nend_bits; s[4]=dec.nbits_total; s[5]=dec.offs;
 s[6]=dec.rng; s[7]=dec.val; s[8]=dec.ext; s[9]=dec.rem; s[10]=dec.error;
}
*/
import "C"

import (
	"github.com/kazzmir/opus-go/opuscc"
	"unsafe"
)

func nativeEnergyDecode(dec *opuscc.OpusT_ec_dec, data []byte, energy []float32, bands, start, end, channels, op, a, b int32, quant, priority []int32) {
	s := [11]C.uint{C.uint(dec.Fstorage), C.uint(dec.Fend_offs), C.uint(dec.Fend_window), C.uint(dec.Fnend_bits), C.uint(dec.Fnbits_total), C.uint(dec.Foffs), C.uint(dec.Frng), C.uint(dec.Fval), C.uint(dec.Fext), C.uint(dec.Frem), C.uint(dec.Ferror1)}
	C.energy_decode(&s[0], (*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), (*C.float)(unsafe.Pointer(unsafe.SliceData(energy))), C.int(bands), C.int(start), C.int(end), C.int(channels), C.int(op), C.int(a), C.int(b), (*C.int)(unsafe.Pointer(unsafe.SliceData(quant))), (*C.int)(unsafe.Pointer(unsafe.SliceData(priority))))
	dec.Fstorage = uint32(s[0])
	dec.Fend_offs = uint32(s[1])
	dec.Fend_window = uint32(s[2])
	dec.Fnend_bits = int32(s[3])
	dec.Fnbits_total = int32(s[4])
	dec.Foffs = uint32(s[5])
	dec.Frng = uint32(s[6])
	dec.Fval = uint32(s[7])
	dec.Fext = uint32(s[8])
	dec.Frem = int32(s[9])
	dec.Ferror1 = int32(s[10])
}
