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
#define silk_decode_parameters compare_silk_decode_parameters
#include "../../../opus/silk/decode_parameters.c"
#define silk_decode_indices compare_silk_decode_indices
#include "../../../opus/silk/decode_indices.c"
static void native_silk_indices(void *state,unsigned *s,unsigned char *data,int frame,int lbrr,int cond) {
 silk_decoder_state *st=state;const unsigned char *low=st->pitch_lag_low_bits_iCDF,*contour=st->pitch_contour_iCDF;
 st->psNLSF_CB=st->fs_kHz==16?&silk_NLSF_CB_WB:&silk_NLSF_CB_NB_MB;
 st->pitch_lag_low_bits_iCDF=st->fs_kHz==16?silk_uniform8_iCDF:st->fs_kHz==12?silk_uniform6_iCDF:silk_uniform4_iCDF;
 st->pitch_contour_iCDF=st->fs_kHz==8?(st->nb_subfr==4?silk_pitch_contour_NB_iCDF:silk_pitch_contour_10_ms_NB_iCDF):(st->nb_subfr==4?silk_pitch_contour_iCDF:silk_pitch_contour_10_ms_iCDF);
 ec_dec dec={0};dec.buf=data;dec.storage=s[0];dec.end_offs=s[1];dec.end_window=s[2];dec.nend_bits=s[3];dec.nbits_total=s[4];dec.offs=s[5];dec.rng=s[6];dec.val=s[7];dec.ext=s[8];dec.rem=s[9];dec.error=s[10];
 silk_decode_indices(st,&dec,frame,lbrr,cond);st->psNLSF_CB=NULL;st->pitch_lag_low_bits_iCDF=low;st->pitch_contour_iCDF=contour;
 s[0]=dec.storage;s[1]=dec.end_offs;s[2]=dec.end_window;s[3]=dec.nend_bits;s[4]=dec.nbits_total;s[5]=dec.offs;s[6]=dec.rng;s[7]=dec.val;s[8]=dec.ext;s[9]=dec.rem;s[10]=dec.error;
}
static void native_silk_parameters(void *state,void *control,int cond) {
 silk_decoder_state *st=state;st->psNLSF_CB=st->fs_kHz==16?&silk_NLSF_CB_WB:&silk_NLSF_CB_NB_MB;
 silk_decode_parameters(st,control,cond);st->psNLSF_CB=NULL;
}
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
#define silk_CNG compare_whole_cng
#define silk_CNG_Reset compare_whole_cng_reset
#define silk_CNG_exc compare_whole_cng_exc
#include "../../../opus/silk/CNG.c"
static void native_whole_cng(void *state,void *control,short *frame,int length) {silk_CNG(state,control,frame,length);}
*/
import "C"
import "github.com/kazzmir/opus-go/opuscc"
import "unsafe"

func nativeWholeCNG(st *opuscc.OpusT_silk_decoder_state, control *opuscc.OpusT_silk_decoder_control, frame []int16) {
	C.native_whole_cng(unsafe.Pointer(st), unsafe.Pointer(control), (*C.short)(unsafe.Pointer(unsafe.SliceData(frame))), C.int(len(frame)))
}

func nativeSilkIndices(st *opuscc.OpusT_silk_decoder_state, e *opuscc.OpusT_ec_dec, data []byte, frame, lbrr, cond int32) {
	s := [11]C.uint{C.uint(e.Fstorage), C.uint(e.Fend_offs), C.uint(e.Fend_window), C.uint(e.Fnend_bits), C.uint(e.Fnbits_total), C.uint(e.Foffs), C.uint(e.Frng), C.uint(e.Fval), C.uint(e.Fext), C.uint(e.Frem), C.uint(e.Ferror1)}
	C.native_silk_indices(unsafe.Pointer(st), &s[0], (*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.int(frame), C.int(lbrr), C.int(cond))
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

func nativeSilkParameters(st *opuscc.OpusT_silk_decoder_state, control *opuscc.OpusT_silk_decoder_control, cond int32) {
	C.native_silk_parameters(unsafe.Pointer(st), unsafe.Pointer(control), C.int(cond))
}

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
