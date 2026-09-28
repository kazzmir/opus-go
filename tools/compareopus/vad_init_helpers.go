//go:build compareopus && cgo

package main

/*
#include "main.h"
*/
import "C"
import "github.com/kazzmir/opus-go/opuscc"

func nativeVADInit() (opuscc.OpusT_silk_VAD_state, int32) {
	var c C.silk_VAD_state
	ret := C.silk_VAD_Init(&c)
	var g opuscc.OpusT_silk_VAD_state
	for i := range g.FAnaState {
		g.FAnaState[i] = int32(c.AnaState[i])
		g.FAnaState1[i] = int32(c.AnaState1[i])
		g.FAnaState2[i] = int32(c.AnaState2[i])
	}
	for i := range g.FNL {
		g.FXnrgSubfr[i] = int32(c.XnrgSubfr[i])
		g.FNrgRatioSmth_Q8[i] = int32(c.NrgRatioSmth_Q8[i])
		g.FNL[i] = int32(c.NL[i])
		g.Finv_NL[i] = int32(c.inv_NL[i])
		g.FNoiseLevelBias[i] = int32(c.NoiseLevelBias[i])
	}
	g.FHPstate = int16(c.HPstate)
	g.Fcounter = int32(c.counter)
	return g, int32(ret)
}
