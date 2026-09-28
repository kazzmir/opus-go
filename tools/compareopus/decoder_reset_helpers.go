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
*/
import "C"
import "github.com/kazzmir/opus-go/opuscc"

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
