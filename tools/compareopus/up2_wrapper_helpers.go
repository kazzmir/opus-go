//go:build compareopus && cgo

package main

/*
#include <string.h>
#include "resampler_private.h"
static int up2_wrapper(opus_int32 *iir, opus_int16 *out, const opus_int16 *in, int n) {
 silk_resampler_state_struct s, before;
 memset(&s,0xa5,sizeof(s));
 memcpy(s.sIIR,iir,sizeof(s.sIIR));
 memcpy(&before,&s,sizeof(s));
 silk_resampler_private_up2_HQ_wrapper(&s,out,in,n);
 memcpy(iir,s.sIIR,sizeof(s.sIIR));
 memcpy(s.sIIR,before.sIIR,sizeof(s.sIIR));
 return memcmp(&s,&before,sizeof(s))==0;
}
*/
import "C"
import "unsafe"

func nativeUp2Wrapper(iir *[6]int32, out, in []int16) bool {
	return C.up2_wrapper((*C.opus_int32)(unsafe.Pointer(iir)), (*C.opus_int16)(unsafe.Pointer(unsafe.SliceData(out))), (*C.opus_int16)(unsafe.Pointer(unsafe.SliceData(in))), C.int(len(in))) != 0
}
