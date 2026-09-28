//go:build compareopus && cgo

package main

/*
#include "main.h"
#include "tables.h"
static int encode_stereo_pred(unsigned char *data, unsigned size, const int *ix, int mid) {
 ec_enc enc;
 ec_enc_init(&enc,data,size);
 ec_enc_icdf(&enc,ix[0],silk_stereo_pred_joint_iCDF,8);
 ec_enc_icdf(&enc,ix[1],silk_uniform3_iCDF,8);
 ec_enc_icdf(&enc,ix[2],silk_uniform5_iCDF,8);
 ec_enc_icdf(&enc,ix[3],silk_uniform3_iCDF,8);
 ec_enc_icdf(&enc,ix[4],silk_uniform5_iCDF,8);
 ec_enc_icdf(&enc,mid,silk_stereo_only_code_mid_iCDF,8);
 ec_enc_done(&enc);
 return enc.error;
}
*/
import "C"
import "unsafe"

func nativeStereoPredEncode(data []byte, indices *[5]int32, mid int32) int {
	return int(C.encode_stereo_pred((*C.uchar)(unsafe.Pointer(&data[0])), C.uint(len(data)), (*C.int)(unsafe.Pointer(indices)), C.int(mid)))
}
