//go:build compareopus && cgo

package main

/*
#define ec_enc_init compare_enc_init
#define ec_encode compare_encode
#define ec_encode_bin compare_encode_bin
#define ec_enc_bit_logp compare_enc_bit_logp
#define ec_enc_icdf compare_enc_icdf
#define ec_enc_icdf16 compare_enc_icdf16
#define ec_enc_uint compare_enc_uint
#define ec_enc_bits compare_enc_bits
#define ec_enc_patch_initial_bits compare_enc_patch_initial_bits
#define ec_enc_shrink compare_enc_shrink
#define ec_enc_done compare_enc_done
#include "../../../opus/celt/entenc.c"
static void native_encoder_step(unsigned *s,unsigned char *buf,int op,unsigned a,unsigned b,unsigned c,const opus_uint16 *table) {
 ec_enc e={0};e.buf=buf;
 e.storage=s[0];e.end_offs=s[1];e.end_window=s[2];e.nend_bits=s[3];e.nbits_total=s[4];e.offs=s[5];
 e.rng=s[6];e.val=s[7];e.ext=s[8];e.rem=s[9];e.error=s[10];
 switch(op) {case 0:ec_enc_init(&e,buf,a);break;case 1:ec_enc_shrink(&e,a);break;case 2:ec_enc_patch_initial_bits(&e,a,b);break;case 3:ec_enc_carry_out(&e,(int)a);break;case 4:ec_enc_normalize(&e);break;case 5:ec_encode(&e,a,b,c);break;case 6:ec_encode_bin(&e,a,b,c);break;case 7:ec_enc_bit_logp(&e,(int)a,b);break;case 8:ec_enc_icdf16(&e,(int)a,table,b);break;case 9:ec_enc_bits(&e,a,b);break;case 10:ec_enc_done(&e);break;}
 s[0]=e.storage;s[1]=e.end_offs;s[2]=e.end_window;s[3]=e.nend_bits;s[4]=e.nbits_total;s[5]=e.offs;
 s[6]=e.rng;s[7]=e.val;s[8]=e.ext;s[9]=e.rem;s[10]=e.error;
}
*/
import "C"
import (
	"github.com/kazzmir/opus-go/opuscc"
	"unsafe"
)

func nativeEncoderStep(e *opuscc.OpusT_ec_enc, buf []byte, op int, a, b uint32, totals ...uint32) {
	var total uint32
	if len(totals) > 0 {
		total = totals[0]
	}
	nativeEncoderStepTable(e, buf, op, a, b, total, nil)
}

func nativeEncoderStepTable(e *opuscc.OpusT_ec_enc, buf []byte, op int, a, b, total uint32, table []uint16) {
	s := [11]C.uint{C.uint(e.Fstorage), C.uint(e.Fend_offs), C.uint(e.Fend_window), C.uint(e.Fnend_bits), C.uint(e.Fnbits_total), C.uint(e.Foffs), C.uint(e.Frng), C.uint(e.Fval), C.uint(e.Fext), C.uint(e.Frem), C.uint(e.Ferror1)}
	C.native_encoder_step(&s[0], (*C.uchar)(unsafe.Pointer(unsafe.SliceData(buf))), C.int(op), C.uint(a), C.uint(b), C.uint(total), (*C.opus_uint16)(unsafe.Pointer(unsafe.SliceData(table))))
	e.Fbuf = unsafe.SliceData(buf)
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
