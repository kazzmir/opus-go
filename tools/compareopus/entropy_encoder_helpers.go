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
#define ec_laplace_encode compare_encoder_laplace_encode
#define ec_laplace_decode compare_encoder_laplace_decode
#define ec_laplace_encode_p0 compare_encoder_laplace_encode_p0
#define ec_laplace_decode_p0 compare_encoder_laplace_decode_p0
#include "../../../opus/celt/laplace.c"
#define VAR_ARRAYS 1
#define log2_frac compare_encoder_log2_frac
#define get_required_bits compare_encoder_get_required_bits
#define encode_pulses compare_encoder_encode_pulses
#define decode_pulses compare_encoder_decode_pulses
#include "../../../opus/celt/cwrs.c"
#define silk_shell_encoder compare_encoder_shell_encoder
#define silk_shell_decoder compare_encoder_shell_decoder
#include "../../../opus/silk/shell_coder.c"
#define silk_encode_signs compare_encoder_encode_signs
#define silk_decode_signs compare_encoder_decode_signs
#include "../../../opus/silk/code_signs.c"
static int native_encoder_step(unsigned *s,unsigned char *buf,int op,unsigned a,unsigned b,unsigned c,const opus_uint16 *table,const int *sums) {
 int symbol=(int)a;
 ec_enc e={0};e.buf=buf;
 e.storage=s[0];e.end_offs=s[1];e.end_window=s[2];e.nend_bits=s[3];e.nbits_total=s[4];e.offs=s[5];
 e.rng=s[6];e.val=s[7];e.ext=s[8];e.rem=s[9];e.error=s[10];
 switch(op) {case 0:ec_enc_init(&e,buf,a);break;case 1:ec_enc_shrink(&e,a);break;case 2:ec_enc_patch_initial_bits(&e,a,b);break;case 3:ec_enc_carry_out(&e,(int)a);break;case 4:ec_enc_normalize(&e);break;case 5:ec_encode(&e,a,b,c);break;case 6:ec_encode_bin(&e,a,b,c);break;case 7:ec_enc_bit_logp(&e,(int)a,b);break;case 8:ec_enc_icdf16(&e,(int)a,table,b);break;case 9:ec_enc_bits(&e,a,b);break;case 10:ec_enc_done(&e);break;case 11:ec_enc_uint(&e,a,b);break;case 12:ec_enc_icdf(&e,(int)a,(const unsigned char *)table,b);break;case 13:ec_laplace_encode(&e,&symbol,b,(int)c);break;case 14:ec_laplace_encode_p0(&e,symbol,(opus_uint16)b,(opus_uint16)c);break;case 15:encode_pulses((const int *)table,(int)b,(int)c,&e);break;case 16:silk_shell_encoder(&e,(const int *)table);break;case 17:silk_encode_signs(&e,(const opus_int8 *)table,(int)a,(int)b,(int)c,sums);break;}
 s[0]=e.storage;s[1]=e.end_offs;s[2]=e.end_window;s[3]=e.nend_bits;s[4]=e.nbits_total;s[5]=e.offs;
 s[6]=e.rng;s[7]=e.val;s[8]=e.ext;s[9]=e.rem;s[10]=e.error;
 return symbol;
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
	nativeEncoderStepPointer(e, buf, op, a, b, total, unsafe.Pointer(unsafe.SliceData(table)))
}

func nativeEncoderStepPointer(e *opuscc.OpusT_ec_enc, buf []byte, op int, a, b, total uint32, table unsafe.Pointer, sumPointers ...*int32) int32 {
	var sums *int32
	if len(sumPointers) > 0 {
		sums = sumPointers[0]
	}
	s := [11]C.uint{C.uint(e.Fstorage), C.uint(e.Fend_offs), C.uint(e.Fend_window), C.uint(e.Fnend_bits), C.uint(e.Fnbits_total), C.uint(e.Foffs), C.uint(e.Frng), C.uint(e.Fval), C.uint(e.Fext), C.uint(e.Frem), C.uint(e.Ferror1)}
	result := C.native_encoder_step(&s[0], (*C.uchar)(unsafe.Pointer(unsafe.SliceData(buf))), C.int(op), C.uint(a), C.uint(b), C.uint(total), (*C.opus_uint16)(table), (*C.int)(unsafe.Pointer(sums)))
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
	return int32(result)
}
