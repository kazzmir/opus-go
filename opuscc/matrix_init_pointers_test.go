package opuscc

import (
	"slices"
	"testing"
	"unsafe"
)

type mappingTestBuffer struct {
	Header  OpusT_MappingMatrix
	Padding [2]int16
	Data    [36]int16
	Guard   int16
}

func TestMatrixInitPointers(t *testing.T) {
	var owner mappingTestBuffer
	owner.Padding = [2]int16{123, 456}
	owner.Guard = 789
	data := [6]int16{-32768, 32767, 0, -1, 123, -456}
	if unsafe.Offsetof(owner.Data) != 16 {
		t.Fatal("coefficient alignment")
	}
	Opus_mapping_matrix_init(nil, &owner.Header, 2, 3, -17, &data[0], 12)
	if owner.Header != (OpusT_MappingMatrix{Frows: 2, Fcols: 3, Fgain: -17}) || !slices.Equal(owner.Data[:6], data[:]) || mappingMatrixData(&owner.Header) != &owner.Data[0] {
		t.Fatal("header/data")
	}
	// A source immediately before the destination must propagate forward.
	Opus_mapping_matrix_init(nil, &owner.Header, 2, 3, 0, &owner.Padding[1], 12)
	for _, v := range owner.Data[:6] {
		if v != 456 {
			t.Fatal("copy order", owner.Data)
		}
	}
	if owner.Padding != [2]int16{123, 456} || owner.Guard != 789 {
		t.Fatal("guards")
	}
	Opus_mapping_matrix_init(nil, &owner.Header, 0, 0, 31, nil, 0)
	if owner.Header.Fgain != 31 {
		t.Fatal("empty")
	}
}
