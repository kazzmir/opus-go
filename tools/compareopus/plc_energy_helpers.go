//go:build compareopus && cgo

package main

/*
// Include the actual C implementation to reach its static energy helper.
// Rename external definitions so the linked libopus retains its own symbols.
#define VAR_ARRAYS 1
#define silk_PLC_Reset comparison_PLC_Reset
#define silk_PLC comparison_PLC
#define silk_PLC_glue_frames comparison_PLC_glue_frames
#include "PLC.c"
static void plc_update(int *p,int *d,int *c) {
 silk_decoder_state dec={0};silk_decoder_control ctrl={0};
 dec.fs_kHz=d[0];dec.subfr_length=d[1];dec.nb_subfr=d[2];dec.LPC_order=d[3];dec.indices.signalType=d[4];dec.prevSignalType=d[5];dec.frame_length=d[6];dec.lossCnt=d[7];
 silk_PLC_struct *plc=&dec.sPLC;
 plc->pitchL_Q8=p[0];for(int i=0;i<5;i++) plc->LTPCoef_Q14[i]=p[1+i];for(int i=0;i<16;i++) plc->prevLPC_Q12[i]=p[6+i];
 plc->last_frame_lost=p[22];plc->rand_seed=p[23];plc->randScale_Q14=p[24];plc->conc_energy=p[25];plc->conc_energy_shift=p[26];plc->prevLTP_scale_Q14=p[27];plc->prevGain_Q16[0]=p[28];plc->prevGain_Q16[1]=p[29];plc->fs_kHz=p[30];plc->nb_subfr=p[31];plc->subfr_length=p[32];
 for(int i=0;i<20;i++) ctrl.LTPCoef_Q14[i]=c[i];for(int i=0;i<16;i++) ctrl.PredCoef_Q12[1][i]=c[20+i];for(int i=0;i<4;i++) {ctrl.Gains_Q16[i]=c[36+i];ctrl.pitchL[i]=c[40+i];}ctrl.LTP_scale_Q14=c[44];
 silk_PLC_update(&dec,&ctrl);
 p[0]=plc->pitchL_Q8;for(int i=0;i<5;i++) p[1+i]=plc->LTPCoef_Q14[i];for(int i=0;i<16;i++) p[6+i]=plc->prevLPC_Q12[i];p[22]=plc->last_frame_lost;p[23]=plc->rand_seed;p[24]=plc->randScale_Q14;p[25]=plc->conc_energy;p[26]=plc->conc_energy_shift;p[27]=plc->prevLTP_scale_Q14;p[28]=plc->prevGain_Q16[0];p[29]=plc->prevGain_Q16[1];p[30]=plc->fs_kHz;p[31]=plc->nb_subfr;p[32]=plc->subfr_length;d[5]=dec.prevSignalType;
}
static void plc_glue(int *p,int loss,short *frame,int length) {
 silk_decoder_state d={0};d.lossCnt=loss;d.sPLC.conc_energy=p[0];d.sPLC.conc_energy_shift=p[1];d.sPLC.last_frame_lost=p[2];
 silk_PLC_glue_frames(&d,frame,length);
 p[0]=d.sPLC.conc_energy;p[1]=d.sPLC.conc_energy_shift;p[2]=d.sPLC.last_frame_lost;
}
static void plc_energy(int *out, const opus_int32 *exc, const opus_int32 *gains, int n, int nb) {
 silk_PLC_energy(&out[0],&out[1],&out[2],&out[3],exc,gains,n,nb);
}
*/
import "C"
import (
	"github.com/kazzmir/opus-go/opuscc"
	"unsafe"
)

func nativePLCUpdate(dec *opuscc.OpusT_silk_decoder_state, ctrl *opuscc.OpusT_silk_decoder_control) {
	plc := &dec.FsPLC
	p := [33]int32{0: plc.FpitchL_Q8, 22: plc.Flast_frame_lost, 23: plc.Frand_seed, 24: int32(plc.FrandScale_Q14), 25: plc.Fconc_energy, 26: plc.Fconc_energy_shift, 27: int32(plc.FprevLTP_scale_Q14), 28: plc.FprevGain_Q16[0], 29: plc.FprevGain_Q16[1], 30: plc.Ffs_kHz, 31: plc.Fnb_subfr, 32: plc.Fsubfr_length}
	for i, v := range plc.FLTPCoef_Q14 {
		p[1+i] = int32(v)
	}
	for i, v := range plc.FprevLPC_Q12 {
		p[6+i] = int32(v)
	}
	d := [8]int32{dec.Ffs_kHz, dec.Fsubfr_length, dec.Fnb_subfr, dec.FLPC_order, int32(dec.Findices.FsignalType), dec.FprevSignalType, dec.Fframe_length, dec.FlossCnt}
	var c [45]int32
	for i, v := range ctrl.FLTPCoef_Q14 {
		c[i] = int32(v)
	}
	for i, v := range ctrl.FPredCoef_Q12[1] {
		c[20+i] = int32(v)
	}
	for i, v := range ctrl.FGains_Q16 {
		c[36+i] = v
	}
	copy(c[40:44], ctrl.FpitchL[:])
	c[44] = ctrl.FLTP_scale_Q14
	C.plc_update((*C.int)(unsafe.Pointer(&p[0])), (*C.int)(unsafe.Pointer(&d[0])), (*C.int)(unsafe.Pointer(&c[0])))
	plc.FpitchL_Q8 = p[0]
	for i := range plc.FLTPCoef_Q14 {
		plc.FLTPCoef_Q14[i] = int16(p[1+i])
	}
	for i := range plc.FprevLPC_Q12 {
		plc.FprevLPC_Q12[i] = int16(p[6+i])
	}
	plc.Flast_frame_lost = p[22]
	plc.Frand_seed = p[23]
	plc.FrandScale_Q14 = int16(p[24])
	plc.Fconc_energy = p[25]
	plc.Fconc_energy_shift = p[26]
	plc.FprevLTP_scale_Q14 = int16(p[27])
	plc.FprevGain_Q16 = [2]int32{p[28], p[29]}
	plc.Ffs_kHz = p[30]
	plc.Fnb_subfr = p[31]
	plc.Fsubfr_length = p[32]
	dec.FprevSignalType = d[5]
}

func nativePLCGlue(dec *opuscc.OpusT_silk_decoder_state, frame []int16) {
	p := [3]int32{dec.FsPLC.Fconc_energy, dec.FsPLC.Fconc_energy_shift, dec.FsPLC.Flast_frame_lost}
	C.plc_glue((*C.int)(unsafe.Pointer(&p[0])), C.int(dec.FlossCnt), (*C.short)(unsafe.Pointer(unsafe.SliceData(frame))), C.int(len(frame)))
	dec.FsPLC.Fconc_energy = p[0]
	dec.FsPLC.Fconc_energy_shift = p[1]
	dec.FsPLC.Flast_frame_lost = p[2]
}

func nativePLCEnergy(exc []int32, gains *[2]int32, n, nb int32) [4]int32 {
	var out [4]int32
	C.plc_energy((*C.int)(unsafe.Pointer(&out[0])), (*C.opus_int32)(unsafe.Pointer(&exc[0])), (*C.opus_int32)(unsafe.Pointer(gains)), C.int(n), C.int(nb))
	return out
}
