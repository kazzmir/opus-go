//go:build compareopus

package opuscc

func CompareInterleaveHadamard(x *float32, n0, stride, hadamard int32) {
	interleave_hadamard(nil, x, n0, stride, hadamard)
}

func CompareDeinterleaveHadamard(x *float32, n0, stride, hadamard int32) {
	deinterleave_hadamard(nil, x, n0, stride, hadamard)
}
