//go:build compareopus && cgo

package main

/*
#include "main.h"
*/
import "C"
import (
	"github.com/kazzmir/opus-go/opuscc"
	"unsafe"
)

func nativeStereoMS(state *opuscc.OpusT_stereo_dec_state, mid, side []int16, predictors *[2]int32, fs int32) {
	var cs C.stereo_dec_state
	for i := 0; i < 2; i++ {
		cs.pred_prev_Q13[i] = C.opus_int16(state.Fpred_prev_Q13[i])
		cs.sMid[i] = C.opus_int16(state.FsMid[i])
		cs.sSide[i] = C.opus_int16(state.FsSide[i])
	}
	C.silk_stereo_MS_to_LR(&cs, (*C.opus_int16)(unsafe.Pointer(&mid[0])), (*C.opus_int16)(unsafe.Pointer(&side[0])), (*C.opus_int32)(unsafe.Pointer(predictors)), C.int(fs), C.int(len(mid)-2))
	for i := 0; i < 2; i++ {
		state.Fpred_prev_Q13[i] = int16(cs.pred_prev_Q13[i])
		state.FsMid[i] = int16(cs.sMid[i])
		state.FsSide[i] = int16(cs.sSide[i])
	}
}
