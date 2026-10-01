//go:build compareopus && cgo

package main

/*
#include <string.h>
#define silk_resampler_init compare_resampler_init
#define silk_resampler compare_resampler_driver
#include "../../../opus/silk/resampler.c"
#define silk_decoder_set_fs compare_decoder_set_fs
#include "../../../opus/silk/decoder_set_fs.c"
static void native_resampler_layout(size_t *v){v[0]=sizeof(silk_resampler_state_struct);v[1]=offsetof(silk_resampler_state_struct,Coefs);}
static void native_decoder_fs_blank(silk_decoder_state *s) {memset(s,0xa5,sizeof(*s));}
static const unsigned char *native_fs_contour(int id) {switch(id) {case 1:return silk_pitch_contour_NB_iCDF;case 2:return silk_pitch_contour_10_ms_NB_iCDF;case 3:return silk_pitch_contour_iCDF;case 4:return silk_pitch_contour_10_ms_iCDF;default:return NULL;}}
static const unsigned char *native_fs_lag(int id) {switch(id) {case 1:return silk_uniform4_iCDF;case 2:return silk_uniform6_iCDF;case 3:return silk_uniform8_iCDF;default:return NULL;}}
static const silk_NLSF_CB_struct *native_fs_codebook(int id) {switch(id) {case 1:return &silk_NLSF_CB_NB_MB;case 2:return &silk_NLSF_CB_WB;default:return NULL;}}
static void native_fs_set_tables(silk_decoder_state *s,int contour,int lag,int cb) {s->pitch_contour_iCDF=native_fs_contour(contour);s->pitch_lag_low_bits_iCDF=native_fs_lag(lag);s->psNLSF_CB=native_fs_codebook(cb);}
static void native_fs_get_tables(const silk_decoder_state *s,int *ids) {
 ids[0]=ids[1]=ids[2]=-1;for(int i=0;i<5;i++) if(s->pitch_contour_iCDF==native_fs_contour(i)) ids[0]=i;
 for(int i=0;i<4;i++) if(s->pitch_lag_low_bits_iCDF==native_fs_lag(i)) ids[1]=i;
 for(int i=0;i<3;i++) if(s->psNLSF_CB==native_fs_codebook(i)) ids[2]=i;
}
static int native_fs_remainder(silk_decoder_state after,const silk_decoder_state *before) {
 after.subfr_length=before->subfr_length;after.fs_API_hz=before->fs_API_hz;after.resampler_state=before->resampler_state;
 after.pitch_contour_iCDF=before->pitch_contour_iCDF;after.ltp_mem_length=before->ltp_mem_length;after.LPC_order=before->LPC_order;after.psNLSF_CB=before->psNLSF_CB;after.pitch_lag_low_bits_iCDF=before->pitch_lag_low_bits_iCDF;
 after.first_frame_after_reset=before->first_frame_after_reset;after.lagPrev=before->lagPrev;after.LastGainIndex=before->LastGainIndex;after.prevSignalType=before->prevSignalType;
 memcpy(after.outBuf,before->outBuf,sizeof(after.outBuf));memcpy(after.sLPC_Q14_buf,before->sLPC_Q14_buf,sizeof(after.sLPC_Q14_buf));after.fs_kHz=before->fs_kHz;after.frame_length=before->frame_length;
 return memcmp(&after,before,sizeof(after))==0;
}
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

func nativeResamplerLayout() [2]uint64 {
	var v [2]C.size_t
	C.native_resampler_layout(&v[0])
	return [2]uint64{uint64(v[0]), uint64(v[1])}
}

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
	g.FCoefs = resamplerCoefPointers[int(C.native_resampler_coef_id(c))]
	return g
}

func resamplerStateToC(g *opuscc.OpusT_silk_resampler_state_struct) C.silk_resampler_state_struct {
	var c C.silk_resampler_state_struct
	for i := range g.FsIIR {
		c.sIIR[i] = C.opus_int32(g.FsIIR[i])
	}
	copy(unsafe.Slice((*int32)(unsafe.Pointer(&c.sFIR)), 36), g.FsFIR.Fi32[:])
	for i := range g.FdelayBuf {
		c.delayBuf[i] = C.opus_int16(g.FdelayBuf[i])
	}
	c.resampler_function = C.int(g.Fresampler_function)
	c.batchSize = C.int(g.FbatchSize)
	c.invRatio_Q16 = C.opus_int32(g.FinvRatio_Q16)
	c.FIR_Order = C.int(g.FFIR_Order)
	c.FIR_Fracs = C.int(g.FFIR_Fracs)
	c.Fs_in_kHz = C.int(g.FFs_in_kHz)
	c.Fs_out_kHz = C.int(g.FFs_out_kHz)
	c.inputDelay = C.int(g.FinputDelay)
	for id, p := range resamplerCoefPointers {
		if g.FCoefs == p {
			c.Coefs = C.native_resampler_coefs(C.int(id))
			break
		}
	}
	return c
}

func nativeResamplerDriver(g *opuscc.OpusT_silk_resampler_state_struct, out, in []int16) int32 {
	c := resamplerStateToC(g)
	ret := C.compare_resampler_driver(&c, (*C.opus_int16)(unsafe.Pointer(unsafe.SliceData(out))), (*C.opus_int16)(unsafe.Pointer(unsafe.SliceData(in))), C.opus_int32(len(in)))
	*g = resamplerStateFromC(&c)
	return int32(ret)
}

var fsContourTables = []uintptr{0, uintptr(unsafe.Pointer(&opuscc.Opus_silk_pitch_contour_NB_iCDF)), uintptr(unsafe.Pointer(&opuscc.Opus_silk_pitch_contour_10_ms_NB_iCDF)), uintptr(unsafe.Pointer(&opuscc.Opus_silk_pitch_contour_iCDF)), uintptr(unsafe.Pointer(&opuscc.Opus_silk_pitch_contour_10_ms_iCDF))}
var fsLagTables = []uintptr{0, uintptr(unsafe.Pointer(&opuscc.Opus_silk_uniform4_iCDF)), uintptr(unsafe.Pointer(&opuscc.Opus_silk_uniform6_iCDF)), uintptr(unsafe.Pointer(&opuscc.Opus_silk_uniform8_iCDF))}
var fsCodebooks = []uintptr{0, uintptr(unsafe.Pointer(&opuscc.Opus_silk_NLSF_CB_NB_MB)), uintptr(unsafe.Pointer(&opuscc.Opus_silk_NLSF_CB_WB))}

func fsTableID(table []uintptr, p uintptr) C.int {
	for i, v := range table {
		if v == p {
			return C.int(i)
		}
	}
	panic("unknown decoder table")
}

func nativeDecoderSetFS(g *opuscc.OpusT_silk_decoder_state, rate, api int32) (int32, bool) {
	var c C.silk_decoder_state
	C.native_decoder_fs_blank(&c)
	c.nb_subfr = C.int(g.Fnb_subfr)
	c.subfr_length = C.int(g.Fsubfr_length)
	c.frame_length = C.int(g.Fframe_length)
	c.fs_kHz = C.int(g.Ffs_kHz)
	c.fs_API_hz = C.int(g.Ffs_API_hz)
	c.ltp_mem_length = C.int(g.Fltp_mem_length)
	c.LPC_order = C.int(g.FLPC_order)
	c.first_frame_after_reset = C.int(g.Ffirst_frame_after_reset)
	c.lagPrev = C.int(g.FlagPrev)
	c.LastGainIndex = C.opus_int8(g.FLastGainIndex)
	c.prevSignalType = C.int(g.FprevSignalType)
	c.resampler_state = resamplerStateToC(&g.Fresampler_state)
	copy(unsafe.Slice((*int16)(unsafe.Pointer(&c.outBuf[0])), len(g.FoutBuf)), g.FoutBuf[:])
	copy(unsafe.Slice((*int32)(unsafe.Pointer(&c.sLPC_Q14_buf[0])), len(g.FsLPC_Q14_buf)), g.FsLPC_Q14_buf[:])
	C.native_fs_set_tables(&c, fsTableID(fsContourTables, g.Fpitch_contour_iCDF), fsTableID(fsLagTables, g.Fpitch_lag_low_bits_iCDF), fsTableID(fsCodebooks, g.FpsNLSF_CB))
	before := c
	ret := C.compare_decoder_set_fs(&c, C.int(rate), C.opus_int32(api))
	unchanged := C.native_fs_remainder(c, &before) != 0
	g.Fsubfr_length = int32(c.subfr_length)
	g.Fframe_length = int32(c.frame_length)
	g.Ffs_kHz = int32(c.fs_kHz)
	g.Ffs_API_hz = int32(c.fs_API_hz)
	g.Fltp_mem_length = int32(c.ltp_mem_length)
	g.FLPC_order = int32(c.LPC_order)
	g.Ffirst_frame_after_reset = int32(c.first_frame_after_reset)
	g.FlagPrev = int32(c.lagPrev)
	g.FLastGainIndex = int8(c.LastGainIndex)
	g.FprevSignalType = int32(c.prevSignalType)
	g.Fresampler_state = resamplerStateFromC(&c.resampler_state)
	copy(g.FoutBuf[:], unsafe.Slice((*int16)(unsafe.Pointer(&c.outBuf[0])), len(g.FoutBuf)))
	copy(g.FsLPC_Q14_buf[:], unsafe.Slice((*int32)(unsafe.Pointer(&c.sLPC_Q14_buf[0])), len(g.FsLPC_Q14_buf)))
	var ids [3]C.int
	C.native_fs_get_tables(&c, &ids[0])
	g.Fpitch_contour_iCDF = fsContourTables[int(ids[0])]
	g.Fpitch_lag_low_bits_iCDF = fsLagTables[int(ids[1])]
	g.FpsNLSF_CB = fsCodebooks[int(ids[2])]
	return int32(ret), unchanged
}

func nativeResamplerInit(in, out, enc int32) (opuscc.OpusT_silk_resampler_state_struct, int32) {
	var ret C.int
	c := C.native_resampler_init(C.int(in), C.int(out), C.int(enc), &ret)
	return resamplerStateFromC(&c), int32(ret)
}
