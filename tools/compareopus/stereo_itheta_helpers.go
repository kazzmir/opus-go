//go:build compareopus && cgo

package main

/*
#define VAR_ARRAYS 1
#define exp_rotation compare_theta_exp_rotation
#define op_pvq_search_c compare_theta_pvq_search
#define alg_quant compare_theta_alg_quant
#define alg_unquant compare_theta_alg_unquant
#define renormalise_vector compare_theta_renormalise
#define stereo_itheta compare_stereo_itheta
#include "../../../opus/celt/vq.c"
*/
import "C"
import "unsafe"

func nativeScalarStereoITheta(x, y []float32, stereo int32) int32 {
	return int32(C.compare_stereo_itheta((*C.float)(unsafe.Pointer(unsafe.SliceData(x))), (*C.float)(unsafe.Pointer(unsafe.SliceData(y))), C.int(stereo), C.int(len(x)), 0))
}
