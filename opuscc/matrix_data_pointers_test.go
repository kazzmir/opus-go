package opuscc

import (
	"runtime"
	"testing"
	"unsafe"
)

//go:noinline
func ownedMatrixData() *int16 {
	var owner mappingTestBuffer
	owner.Header = OpusT_MappingMatrix{Frows: 6, Fcols: 6}
	for i := range owner.Data {
		owner.Data[i] = int16(i*997 - 17000)
	}
	return Opus_mapping_matrix_get_data(nil, &owner.Header)
}

func TestMatrixDataPointers(t *testing.T) {
	var owner mappingTestBuffer
	if got := Opus_mapping_matrix_get_data(nil, &owner.Header); got != &owner.Data[0] {
		t.Fatal("alignment")
	}
	p := ownedMatrixData()
	entropyInitGrowStack(128)
	runtime.GC()
	for i, v := range unsafe.Slice(p, 36) {
		if v != int16(i*997-17000) {
			t.Fatal("owner lost", i, v)
		}
	}
}
