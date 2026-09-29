//go:build compareopus && cgo

package main

/*
#define VAR_ARRAYS 1
#define celt_lcg_rand compare_hadamard_lcg_rand
#define hysteresis_decision compare_hadamard_hysteresis
#define bitexact_cos compare_hadamard_cos
#define bitexact_log2tan compare_hadamard_log2tan
#define compute_band_energies compare_hadamard_energies
#define normalise_bands compare_hadamard_normalise
#define denormalise_bands compare_hadamard_denormalise
#define anti_collapse compare_hadamard_anti_collapse
#define spreading_decision compare_hadamard_spreading
#define haar1 compare_hadamard_haar1
#define quant_all_bands compare_hadamard_quant_all_bands
#include "../../../opus/celt/bands.c"
static void compare_interleave(float *x,int n0,int stride,int hadamard) {
 interleave_hadamard(x,n0,stride,hadamard);
}
static void compare_deinterleave(float *x,int n0,int stride,int hadamard) {
 deinterleave_hadamard(x,n0,stride,hadamard);
}
*/
import "C"
import "unsafe"

func nativeInterleaveHadamard(x []float32, n0, stride, hadamard int32) {
	C.compare_interleave((*C.float)(unsafe.Pointer(unsafe.SliceData(x))), C.int(n0), C.int(stride), C.int(hadamard))
}

func nativeDeinterleaveHadamard(x []float32, n0, stride, hadamard int32) {
	C.compare_deinterleave((*C.float)(unsafe.Pointer(unsafe.SliceData(x))), C.int(n0), C.int(stride), C.int(hadamard))
}
