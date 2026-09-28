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
static void plc_energy(int *out, const opus_int32 *exc, const opus_int32 *gains, int n, int nb) {
 silk_PLC_energy(&out[0],&out[1],&out[2],&out[3],exc,gains,n,nb);
}
*/
import "C"
import "unsafe"

func nativePLCEnergy(exc []int32, gains *[2]int32, n, nb int32) [4]int32 {
	var out [4]int32
	C.plc_energy((*C.int)(unsafe.Pointer(&out[0])), (*C.opus_int32)(unsafe.Pointer(&exc[0])), (*C.opus_int32)(unsafe.Pointer(gains)), C.int(n), C.int(nb))
	return out
}
