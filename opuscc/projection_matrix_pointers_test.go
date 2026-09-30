package opuscc

import (
	"runtime"
	"testing"
	"unsafe"
)

func TestProjectionMatrixPointers(t *testing.T) {
	var matrix *OpusT_MappingMatrix
	func() {
		backing := make([]uint64, 64)
		header := (*OpusT_OpusProjectionDecoder)(unsafe.Pointer(&backing[0]))
		header.Fdemixing_matrix_size_in_bytes = 128
		matrix = get_dec_demixing_matrix(nil, header)
		if unsafe.Pointer(matrix) != unsafe.Add(unsafe.Pointer(header), 8) {
			t.Fatal("offset")
		}
		matrix.Frows = 3
		matrix.Fcols = 2
		matrix.Fgain = 77
	}()
	entropyInitGrowStack(12)
	runtime.GC()
	if matrix.Frows != 3 || matrix.Fcols != 2 || matrix.Fgain != 77 {
		t.Fatal(*matrix)
	}
	// The returned interior pointer retains coefficients in the same allocation.
	data := Opus_mapping_matrix_get_data(nil, matrix)
	*data = 1234
	runtime.GC()
	if *data != 1234 {
		t.Fatal("ownership")
	}
}
