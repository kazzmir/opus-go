//go:build compareopus && cgo

package main

/*
#include <string.h>
#define silk_resampler_init compare_resampler_init
#define silk_resampler compare_resampler_driver
#include "../../../opus/silk/resampler.c"
static silk_resampler_state_struct native_resampler_init(int in,int out,int enc,int *ret) {
 silk_resampler_state_struct s;memset(&s,0xa5,sizeof(s));
 *ret=compare_resampler_init(&s,in,out,enc);return s;
}
static const opus_int16 *native_resampler_coefs(int id) {
 switch(id) {
 case 1:return silk_Resampler_3_4_COEFS;case 2:return silk_Resampler_2_3_COEFS;
 case 3:return silk_Resampler_1_2_COEFS;case 4:return silk_Resampler_1_3_COEFS;
 case 5:return silk_Resampler_1_4_COEFS;case 6:return silk_Resampler_1_6_COEFS;
 default:return NULL;
 }
}
static int native_resampler_coef_id(const silk_resampler_state_struct *s) {
 for(int i=0;i<=6;i++) if(s->Coefs==native_resampler_coefs(i)) return i;
 return -1;
}
*/
import "C"
import (
	"github.com/kazzmir/opus-go/opuscc"
	"unsafe"
)

var resamplerCoefPointers = []*int16{nil, &opuscc.Opus_silk_Resampler_3_4_COEFS[0], &opuscc.Opus_silk_Resampler_2_3_COEFS[0], &opuscc.Opus_silk_Resampler_1_2_COEFS[0], &opuscc.Opus_silk_Resampler_1_3_COEFS[0], &opuscc.Opus_silk_Resampler_1_4_COEFS[0], &opuscc.Opus_silk_Resampler_1_6_COEFS[0]}

func resamplerStateFromC(c *C.silk_resampler_state_struct) opuscc.OpusT_silk_resampler_state_struct {
	var g opuscc.OpusT_silk_resampler_state_struct
	for i := range g.FsIIR {
		g.FsIIR[i] = int32(c.sIIR[i])
	}
	copy(g.FsFIR.Fi32[:], unsafe.Slice((*int32)(unsafe.Pointer(&c.sFIR)), 36))
	for i := range g.FdelayBuf {
		g.FdelayBuf[i] = int16(c.delayBuf[i])
	}
	g.Fresampler_function = int32(c.resampler_function)
	g.FbatchSize = int32(c.batchSize)
	g.FinvRatio_Q16 = int32(c.invRatio_Q16)
	g.FFIR_Order = int32(c.FIR_Order)
	g.FFIR_Fracs = int32(c.FIR_Fracs)
	g.FFs_in_kHz = int32(c.Fs_in_kHz)
	g.FFs_out_kHz = int32(c.Fs_out_kHz)
	g.FinputDelay = int32(c.inputDelay)
	g.FCoefs = uintptr(unsafe.Pointer(resamplerCoefPointers[int(C.native_resampler_coef_id(c))]))
	return g
}

func nativeResamplerInit(in, out, enc int32) (opuscc.OpusT_silk_resampler_state_struct, int32) {
	var ret C.int
	c := C.native_resampler_init(C.int(in), C.int(out), C.int(enc), &ret)
	return resamplerStateFromC(&c), int32(ret)
}
