//go:build compareopus && cgo

package main

/*
#include "main.h"
#include "tables.h"
static void decode_pulses(unsigned char *data, unsigned size, opus_int16 *pulses, int length,
                          int signal, int offset, unsigned *s) {
 ec_dec dec = {0};
 ec_dec_init(&dec,data,size);
 silk_decode_pulses(&dec,pulses,signal,offset,length);
 s[0]=dec.storage; s[1]=dec.end_offs; s[2]=dec.end_window;
 s[3]=dec.nend_bits; s[4]=dec.nbits_total; s[5]=dec.offs;
 s[6]=dec.rng; s[7]=dec.val; s[8]=dec.ext; s[9]=dec.rem; s[10]=dec.error;
}
// A single shell block with the maximum ten escape/LSB stages.
static int encode_pulse_escape(unsigned char *data, unsigned size, int signal, int offset) {
 ec_enc enc;
 ec_enc_init(&enc,data,size);
 ec_enc_icdf(&enc,0,silk_rate_levels_iCDF[signal>>1],8);
 ec_enc_icdf(&enc,17,silk_pulses_per_block_iCDF[0],8);
 for (int i=1;i<10;i++) ec_enc_icdf(&enc,17,silk_pulses_per_block_iCDF[9],8);
 ec_enc_icdf(&enc,16,silk_pulses_per_block_iCDF[9]+1,8);
 int shell[16]={16};
 silk_shell_encoder(&enc,shell);
 for (int k=0;k<16;k++) for (int j=0;j<10;j++) ec_enc_icdf(&enc,1,silk_lsb_iCDF,8);
 unsigned char sign[2]={silk_sign_iCDF[7*(offset+2*signal)+6],0};
 for (int k=0;k<16;k++) ec_enc_icdf(&enc,k&1,sign,8);
 ec_enc_done(&enc);
 return enc.error;
}
*/
import "C"

import (
	"github.com/kazzmir/opus-go/opuscc"
	"unsafe"
)

func nativePulseDecode(data []byte, pulses []int16, length, signal, offset int32) opuscc.OpusT_ec_dec {
	var s [11]C.uint
	C.decode_pulses((*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.uint(len(data)), (*C.opus_int16)(unsafe.Pointer(&pulses[0])), C.int(length), C.int(signal), C.int(offset), &s[0])
	return opuscc.OpusT_ec_dec{Fstorage: uint32(s[0]), Fend_offs: uint32(s[1]), Fend_window: uint32(s[2]), Fnend_bits: int32(s[3]), Fnbits_total: int32(s[4]), Foffs: uint32(s[5]), Frng: uint32(s[6]), Fval: uint32(s[7]), Fext: uint32(s[8]), Frem: int32(s[9]), Ferror1: int32(s[10])}
}

func nativePulseEscape(data []byte, signal, offset int32) int {
	return int(C.encode_pulse_escape((*C.uchar)(unsafe.Pointer(&data[0])), C.uint(len(data)), C.int(signal), C.int(offset)))
}
