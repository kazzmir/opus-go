//go:build compareopus && cgo

package main

/*
#cgo CFLAGS: -I${SRCDIR}/../../../opus/silk
#include "main.h"
*/
import "C"

import (
	"github.com/kazzmir/opus-go/opuscc"
	"unsafe"
)

func nativeLPCutoff(state *opuscc.OpusT_silk_LP_state, frame []int16) {
	// Copy fields explicitly: do not assume Go/C structure layouts match.
	cs := C.silk_LP_state{}
	cs.In_LP_State[0], cs.In_LP_State[1] = C.opus_int32(state.FIn_LP_State[0]), C.opus_int32(state.FIn_LP_State[1])
	cs.transition_frame_no = C.opus_int32(state.Ftransition_frame_no)
	cs.mode = C.int(state.Fmode)
	cs.saved_fs_kHz = C.opus_int32(state.Fsaved_fs_kHz)
	C.silk_LP_variable_cutoff(&cs, (*C.opus_int16)(unsafe.Pointer(unsafe.SliceData(frame))), C.int(len(frame)))
	state.FIn_LP_State = [2]int32{int32(cs.In_LP_State[0]), int32(cs.In_LP_State[1])}
	state.Ftransition_frame_no = int32(cs.transition_frame_no)
	state.Fmode = int32(cs.mode)
	state.Fsaved_fs_kHz = int32(cs.saved_fs_kHz)
}
