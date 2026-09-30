//go:build compareopus && cgo

package main

/*
#include "modes.h"
#include "quant_bands.h"
static void energy_encode(unsigned *s,unsigned char *buf,float *old,float *error,int bands,int start,int end,int channels,int *prev,int *extra,int op,int bitsLeft) {
 CELTMode m={0};m.nbEBands=bands;
 ec_enc e={0};e.buf=buf;e.storage=s[0];e.end_offs=s[1];e.end_window=s[2];e.nend_bits=s[3];e.nbits_total=s[4];e.offs=s[5];e.rng=s[6];e.val=s[7];e.ext=s[8];e.rem=s[9];e.error=s[10];
 if(op==0) quant_fine_energy(&m,start,end,old,error,prev,extra,&e,channels);
 else quant_energy_finalise(&m,start,end,old,error,prev,extra,bitsLeft,&e,channels);
 s[0]=e.storage;s[1]=e.end_offs;s[2]=e.end_window;s[3]=e.nend_bits;s[4]=e.nbits_total;s[5]=e.offs;s[6]=e.rng;s[7]=e.val;s[8]=e.ext;s[9]=e.rem;s[10]=e.error;
}
*/
import "C"
import (
	"github.com/kazzmir/opus-go/opuscc"
	"unsafe"
)

func nativeEnergyEncode(e *opuscc.OpusT_ec_enc, buf []byte, old, err []float32, bands, start, end, channels int32, prev, extra []int32, options ...int32) {
	var op, bitsLeft int32
	if len(options) > 0 {
		op = options[0]
		bitsLeft = options[1]
	}
	s := [11]C.uint{C.uint(e.Fstorage), C.uint(e.Fend_offs), C.uint(e.Fend_window), C.uint(e.Fnend_bits), C.uint(e.Fnbits_total), C.uint(e.Foffs), C.uint(e.Frng), C.uint(e.Fval), C.uint(e.Fext), C.uint(e.Frem), C.uint(e.Ferror1)}
	C.energy_encode(&s[0], (*C.uchar)(unsafe.Pointer(unsafe.SliceData(buf))), (*C.float)(unsafe.Pointer(unsafe.SliceData(old))), (*C.float)(unsafe.Pointer(unsafe.SliceData(err))), C.int(bands), C.int(start), C.int(end), C.int(channels), (*C.int)(unsafe.Pointer(unsafe.SliceData(prev))), (*C.int)(unsafe.Pointer(unsafe.SliceData(extra))), C.int(op), C.int(bitsLeft))
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
}
