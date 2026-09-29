//go:build compareopus && cgo

package main

/*
#include <string.h>
#include "main.h"
static int native_decoder_reset(silk_decoder_state *s, int init) {
 memset(s,0xa5,sizeof(*s));
 return init ? silk_init_decoder(s) : silk_reset_decoder(s);
}
static int reset_remainder_zero(silk_decoder_state s) {
 s.first_frame_after_reset=0; s.prev_gain_Q16=0; s.arch=0;
 s.sCNG.rand_seed=0;
 s.sPLC.prevGain_Q16[0]=s.sPLC.prevGain_Q16[1]=0;
 s.sPLC.subfr_length=0; s.sPLC.nb_subfr=0;
 const unsigned char *p=(const unsigned char*)&s;
 for (unsigned i=0;i<sizeof(s);i++) if (p[i]) return 0;
 return 1;
}
#define VAR_ARRAYS 1
#define silk_LoadOSCEModels compare_api_LoadOSCEModels
#define silk_Get_Decoder_Size compare_api_Get_Decoder_Size
#define silk_ResetDecoder compare_api_ResetDecoder
#define silk_InitDecoder compare_api_InitDecoder
#define silk_Decode compare_api_Decode
#include "../../../opus/silk/dec_API.c"
static int native_api_reset(int init,int *meta) {
 silk_decoder s; silk_decoder_state ref; stereo_dec_state zero={0};
 memset(&s,0xa5,sizeof(s));s.nChannelsAPI=2;s.nChannelsInternal=1;s.prev_decode_only_middle=1;
 int ret=init?compare_api_InitDecoder(&s):compare_api_ResetDecoder(&s);
 native_decoder_reset(&ref,init);ref.arch=0;
 meta[0]=s.nChannelsAPI;meta[1]=s.nChannelsInternal;meta[2]=s.prev_decode_only_middle;
 meta[3]=memcmp(&s.sStereo,&zero,sizeof(zero))==0;meta[4]=1;
 for(int i=0;i<DECODER_NUM_CHANNELS;i++) {s.channel_state[i].arch=0; if(memcmp(&s.channel_state[i],&ref,sizeof(ref))) meta[4]=0;}
 return ret;
}
*/
import "C"
import "github.com/kazzmir/opus-go/opuscc"

func nativeSilkAPIReset(init bool) (int32, [5]int32) {
	var mode C.int
	if init {
		mode = 1
	}
	var meta [5]C.int
	ret := C.native_api_reset(mode, &meta[0])
	var result [5]int32
	for i := range result {
		result[i] = int32(meta[i])
	}
	return int32(ret), result
}

func nativeDecoderReset(init bool) (opuscc.OpusT_silk_decoder_state, int32, bool) {
	var c C.silk_decoder_state
	var mode C.int
	if init {
		mode = 1
	}
	result := C.native_decoder_reset(&c, mode)
	var g opuscc.OpusT_silk_decoder_state
	g.Ffirst_frame_after_reset = int32(c.first_frame_after_reset)
	g.Fprev_gain_Q16 = int32(c.prev_gain_Q16)
	g.FsCNG.Frand_seed = int32(c.sCNG.rand_seed)
	g.FsPLC.FprevGain_Q16 = [2]int32{int32(c.sPLC.prevGain_Q16[0]), int32(c.sPLC.prevGain_Q16[1])}
	g.FsPLC.Fsubfr_length = int32(c.sPLC.subfr_length)
	g.FsPLC.Fnb_subfr = int32(c.sPLC.nb_subfr)
	// Native CPU dispatch differs; Go intentionally keeps scalar arch == 0.
	return g, int32(result), C.reset_remainder_zero(c) != 0
}
