//go:build compareopus && cgo

package main

/*
#include <setjmp.h>
static _Thread_local jmp_buf celt_validation_jump;
void comparison_celt_validator_fatal(const char *str,const char *file,int line) {longjmp(celt_validation_jump,1);}
#define celt_fatal comparison_celt_validator_fatal
#define VAR_ARRAYS 1
#define OPUS_BUILD 1
#define ENABLE_ASSERTIONS 1
#define OPUS_DISABLE_INTRINSICS 1
#define validate_celt_decoder comparison_celt_validate
#define celt_decoder_get_size comparison_celt_size
#define opus_custom_decoder_get_size comparison_custom_size
#define opus_custom_decoder_create comparison_custom_create
#define celt_decoder_init comparison_celt_init
#define opus_custom_decoder_init comparison_custom_init
#define opus_custom_decoder_destroy comparison_custom_destroy
#define celt_synthesis comparison_celt_synthesis
#define celt_decode_with_ec_dred comparison_celt_decode_dred
#define celt_decode_with_ec comparison_celt_decode_ec
#define opus_custom_decode comparison_custom_decode
#define opus_custom_decode24 comparison_custom_decode24
#define opus_custom_decode_float comparison_custom_decode_float
#define opus_custom_decoder_ctl comparison_custom_ctl
#define resampling_factor comparison_celt_resampling
#define pitch_downsample compare_pitch_downsample
#define pitch_search compare_pitch_search
#include "../../../opus/celt/celt_decoder.c"
static int native_plc_pitch(float *left,float *right,int channels) {float *data[2]={left,right};return celt_plc_pitch_search(NULL,data,channels,0);}
#define comb_filter comparison_celt_comb_filter
#define init_caps comparison_celt_init_caps
#define tf_select_table comparison_celt_tf_select_table
#define opus_strerror comparison_celt_strerror
#define opus_get_version_string comparison_celt_version
#include "../../../opus/celt/celt.c"
#define opus_custom_mode_create comparison_mode_create
#define opus_custom_mode_destroy comparison_mode_destroy
#include "../../../opus/celt/modes.c"
static int native_mode_lookup(int Fs,int frame,int *v) {int error=99;CELTMode *mode=opus_custom_mode_create(Fs,frame,&error);if(mode){v[0]=mode->Fs;v[1]=mode->overlap;v[2]=mode->nbEBands;v[3]=mode->effEBands;v[4]=mode->shortMdctSize;v[5]=mode->nbShortMdcts;v[6]=mode->maxLM;}return error;}
#undef opus_custom_mode_create
#undef opus_custom_mode_destroy
static void native_mode_pulse_rate(int band,int LM,int bits,int pulses,int *v) {CELTMode *m=comparison_mode_create(48000,960,NULL);v[0]=bits2pulses(m,band,LM,bits);v[1]=pulses2bits(m,band,LM,pulses);}
static void native_mode_remaining_table(int op,void *out,size_t *v) {
 CELTMode *m=comparison_mode_create(48000,960,NULL);const void *p;
 v[0]=sizeof(CELTMode);
 switch(op) {
 case 0:p=m->logN;v[1]=offsetof(CELTMode,logN);v[2]=m->nbEBands*sizeof(opus_int16);break;
 case 1:p=m->cache.index;v[0]=sizeof(PulseCache);v[1]=offsetof(PulseCache,index);v[2]=(m->maxLM+2)*m->nbEBands*sizeof(opus_int16);break;
 case 2:p=m->cache.bits;v[0]=sizeof(PulseCache);v[1]=offsetof(PulseCache,bits);v[2]=m->cache.size;break;
 default:p=m->eBands;v[1]=offsetof(CELTMode,eBands);v[2]=(m->nbEBands+1)*sizeof(opus_int16);break;
 }
 memcpy(out,p,v[2]);
}
static void native_mode_tables(float *window,unsigned char *vectors,unsigned char *caps,size_t *layout) {
 CELTMode *m=comparison_mode_create(48000,960,NULL);memcpy(window,m->window,m->overlap*sizeof(float));memcpy(vectors,m->allocVectors,m->nbAllocVectors*m->nbEBands);memcpy(caps,m->cache.caps,(m->maxLM+1)*2*m->nbEBands);
 layout[0]=sizeof(CELTMode);layout[1]=offsetof(CELTMode,allocVectors);layout[2]=offsetof(CELTMode,window);layout[3]=sizeof(PulseCache);layout[4]=offsetof(PulseCache,caps);
}
static int native_celt_state(unsigned char *data,size_t size,int op,int channels,int rate,int overlap,int bands,int eff) {
 CELTMode mode={0};mode.overlap=overlap;mode.nbEBands=bands;mode.effEBands=eff;CELTDecoder *st=size?malloc(size):NULL;if(size)memcpy(st,data,size);
 int result;if(setjmp(celt_validation_jump))result=-99;else if(op==0){st->mode=&mode;result=comparison_custom_ctl(st,OPUS_RESET_STATE);}else if(op==1)result=comparison_custom_init(st,&mode,channels);else result=comparison_celt_init(st,rate,channels);
 if(size){st->mode=NULL;memcpy(data,st,size);free(st);}return result;
}
static int native_custom_ctl(unsigned char *data,size_t size,int request,int value,int alias,unsigned *output) {
 CELTDecoder *st=malloc(size);memcpy(st,data,size);st->mode=comparison_mode_create(48000,960,NULL);unsigned out=77;const CELTMode *mode=NULL;void *p=alias==-2?NULL:alias>=0?(void*)((char*)st+alias):(void*)&out;
 int result;if(setjmp(celt_validation_jump))result=-99;else switch(request) {
 case OPUS_SET_COMPLEXITY_REQUEST:case CELT_SET_START_BAND_REQUEST:case CELT_SET_END_BAND_REQUEST:case CELT_SET_CHANNELS_REQUEST:case CELT_SET_SIGNALLING_REQUEST:case OPUS_SET_PHASE_INVERSION_DISABLED_REQUEST:result=comparison_custom_ctl(st,request,value);break;
 case CELT_GET_MODE_REQUEST:result=comparison_custom_ctl(st,request,alias==-2?NULL:alias>=0?(const CELTMode**)((char*)st+alias):&mode);out=mode==st->mode;break;
 case OPUS_RESET_STATE:result=comparison_custom_ctl(st,request);break;
 default:result=comparison_custom_ctl(st,request,p);break;
 }
 st->mode=NULL;memcpy(data,st,size);free(st);*output=out;return result;
}
static void native_tf(unsigned *s,unsigned char *data,int start,int end,int transient,int *out,int LM) {
 ec_dec dec={0};dec.buf=data;dec.storage=s[0];dec.end_offs=s[1];dec.end_window=s[2];dec.nend_bits=s[3];dec.nbits_total=s[4];dec.offs=s[5];dec.rng=s[6];dec.val=s[7];dec.ext=s[8];dec.rem=s[9];dec.error=s[10];
 tf_decode(start,end,transient,out,LM,&dec);
 s[0]=dec.storage;s[1]=dec.end_offs;s[2]=dec.end_window;s[3]=dec.nend_bits;s[4]=dec.nbits_total;s[5]=dec.offs;s[6]=dec.rng;s[7]=dec.val;s[8]=dec.ext;s[9]=dec.rem;s[10]=dec.error;
}
static int native_custom_size(int overlap,int bands,int channels) {CELTMode mode={0};mode.overlap=overlap;mode.nbEBands=bands;return comparison_custom_size(&mode,channels);}
static int native_celt_validation(const int *v,int badmode) {
 CELTDecoder st={0};st.mode=badmode?NULL:opus_custom_mode_create(48000,960,NULL);
 st.overlap=v[0];st.end=v[1];st.channels=v[2];st.stream_channels=v[3];st.downsample=v[4];st.start=v[5];st.arch=v[6];st.last_pitch_index=v[7];st.postfilter_period=v[8];st.postfilter_period_old=v[9];st.postfilter_tapset=v[10];st.postfilter_tapset_old=v[11];
 if(setjmp(celt_validation_jump)) return 1;
 comparison_celt_validate(&st);return 0;
}
*/
import "C"
import (
	"github.com/kazzmir/opus-go/opuscc"
	"unsafe"
)

func nativeCustomCtl(data []byte, request, value, alias int32) (int32, uint32) {
	var out C.uint
	r := C.native_custom_ctl((*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.size_t(len(data)), C.int(request), C.int(value), C.int(alias), &out)
	return int32(r), uint32(out)
}

func nativeModePulseRate(band, LM, bits, pulses int32) (int32, int32) {
	var v [2]C.int
	C.native_mode_pulse_rate(C.int(band), C.int(LM), C.int(bits), C.int(pulses), &v[0])
	return int32(v[0]), int32(v[1])
}

func nativeModeRemainingTable(op int, data unsafe.Pointer) [3]uint64 {
	var v [3]C.size_t
	C.native_mode_remaining_table(C.int(op), data, &v[0])
	return [3]uint64{uint64(v[0]), uint64(v[1]), uint64(v[2])}
}

func nativeModeTables(window []float32, vectors, caps []byte) [5]uint64 {
	var v [5]C.size_t
	C.native_mode_tables((*C.float)(unsafe.Pointer(unsafe.SliceData(window))), (*C.uchar)(unsafe.Pointer(unsafe.SliceData(vectors))), (*C.uchar)(unsafe.Pointer(unsafe.SliceData(caps))), &v[0])
	var out [5]uint64
	for i := range out {
		out[i] = uint64(v[i])
	}
	return out
}

func nativePLCPitchSearch(left, right []float32, channels int32) int32 {
	return int32(C.native_plc_pitch((*C.float)(unsafe.Pointer(unsafe.SliceData(left))), (*C.float)(unsafe.Pointer(unsafe.SliceData(right))), C.int(channels)))
}

func nativeCeltState(data []byte, op, channels, rate, overlap, bands, eff int32) int32 {
	return int32(C.native_celt_state((*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.size_t(len(data)), C.int(op), C.int(channels), C.int(rate), C.int(overlap), C.int(bands), C.int(eff)))
}

func nativeCustomMode(rate, frame int32) (int32, [7]int32) {
	var v [7]C.int
	code := C.native_mode_lookup(C.int(rate), C.int(frame), &v[0])
	var out [7]int32
	for i := range v {
		out[i] = int32(v[i])
	}
	return int32(code), out
}

func nativeTFDecode(dec *opuscc.OpusT_ec_dec, data []byte, start, end, transient int32, out []int32, LM int32) {
	s := [11]C.uint{C.uint(dec.Fstorage), C.uint(dec.Fend_offs), C.uint(dec.Fend_window), C.uint(dec.Fnend_bits), C.uint(dec.Fnbits_total), C.uint(dec.Foffs), C.uint(dec.Frng), C.uint(dec.Fval), C.uint(dec.Fext), C.uint(dec.Frem), C.uint(dec.Ferror1)}
	C.native_tf(&s[0], (*C.uchar)(unsafe.Pointer(unsafe.SliceData(data))), C.int(start), C.int(end), C.int(transient), (*C.int)(unsafe.Pointer(unsafe.SliceData(out))), C.int(LM))
	dec.Fstorage = uint32(s[0])
	dec.Fend_offs = uint32(s[1])
	dec.Fend_window = uint32(s[2])
	dec.Fnend_bits = int32(s[3])
	dec.Fnbits_total = int32(s[4])
	dec.Foffs = uint32(s[5])
	dec.Frng = uint32(s[6])
	dec.Fval = uint32(s[7])
	dec.Fext = uint32(s[8])
	dec.Frem = int32(s[9])
	dec.Ferror1 = int32(s[10])
}

func nativeCustomDecoderSize(overlap, bands, channels int32) int32 {
	return int32(C.native_custom_size(C.int(overlap), C.int(bands), C.int(channels)))
}

func nativeCeltValidation(st *opuscc.OpusT_OpusCustomDecoder) bool {
	mode, _ := opuscc.Opus_opus_custom_mode_create(nil, 48000, 960)
	badmode := int32(0)
	if st.Fmode != mode {
		badmode = 1
	}
	v := [12]int32{st.Foverlap, st.Fend, st.Fchannels, st.Fstream_channels, st.Fdownsample, st.Fstart, st.Farch, st.Flast_pitch_index, st.Fpostfilter_period, st.Fpostfilter_period_old, st.Fpostfilter_tapset, st.Fpostfilter_tapset_old}
	return C.native_celt_validation((*C.int)(unsafe.Pointer(&v[0])), C.int(badmode)) != 0
}
