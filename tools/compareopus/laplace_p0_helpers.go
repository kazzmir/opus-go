//go:build compareopus && cgo

package main

/*
#include "laplace.h"
static int encode_p0(unsigned char *data, unsigned size, const int *values, int n, opus_uint16 p0, opus_uint16 decay) {
 ec_enc enc;
 ec_enc_init(&enc,data,size);
 for (int i=0;i<n;i++) ec_laplace_encode_p0(&enc,values[i],p0,decay);
 ec_enc_done(&enc);
 return enc.error;
}
*/
import "C"
import "unsafe"

func nativeLaplaceP0Encode(data []byte, values []int32, p0, decay uint16) int {
	// A zero-width symbol interval makes the C encoder normalize forever.
	if p0 > 32768 || decay >= 32768 {
		return -1
	}
	for _, value := range values {
		if (p0 == 0 && value == 0) || (p0 == 32768 && value != 0) || (p0 == 32767 && value < 0) {
			return -1
		}
	}
	return int(C.encode_p0((*C.uchar)(unsafe.Pointer(&data[0])), C.uint(len(data)), (*C.int)(unsafe.Pointer(&values[0])), C.int(len(values)), C.opus_uint16(p0), C.opus_uint16(decay)))
}
