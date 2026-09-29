//go:build compareopus

package opuscc

import "unsafe"

// CompareHybridFolding exposes the internal helper only in comparison builds.
func CompareHybridFolding(bands []int16, norm, norm2 []float32, start, M, dual int32) {
	special_hybrid_folding(nil, unsafe.SliceData(bands), unsafe.SliceData(norm), unsafe.SliceData(norm2), start, M, dual)
}
