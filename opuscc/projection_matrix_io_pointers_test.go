package opuscc

import (
	"math"
	"runtime"
	"slices"
	"testing"
	"unsafe"
)

// From matrix_data_pointers_test.go.
//
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

// From matrix_float_pointers_test.go.
func TestMatrixFloatPointers(t *testing.T) {
	var owner mappingTestBuffer
	coefs := [6]int16{16384, -32768, 32767, 8192, -16384, 4096}
	Opus_mapping_matrix_init(nil, &owner.Header, 2, 3, 999, &coefs[0], 12)
	Opus_mapping_matrix_multiply_channel_out_float(nil, &owner.Header, nil, 0, 1, nil, 2, 0)
	in := [3]float32{0.5, 123, -0.25}
	out := [6]float32{77, 1, 2, 3, 4, 88}
	Opus_mapping_matrix_multiply_channel_out_float(nil, &owner.Header, &in[0], 2, 2, &out[1], 2, 2)
	if out != [6]float32{77, 0.75, 2.0625, 3.125, 3.96875, 88} {
		t.Fatal(out)
	}
	// Read each frame's sample before writing any rows, even when aliased.
	alias := [4]float32{0.5, 0.25, 1, 2}
	Opus_mapping_matrix_multiply_channel_out_float(nil, &owner.Header, &alias[0], 0, 1, &alias[0], 2, 2)
	if alias != [4]float32{0.75, -0.25, 0.875, 2.25} {
		t.Fatal("alias", alias)
	}
}

// From matrix_int24_pointers_test.go.
func TestMatrixInt24Pointers(t *testing.T) {
	var owner mappingTestBuffer
	coef := int16(-32768)
	Opus_mapping_matrix_init(nil, &owner.Header, 1, 1, 0, &coef, 2)
	Opus_mapping_matrix_multiply_channel_out_int24(nil, &owner.Header, nil, 0, 1, nil, 1, 0)
	in := []float32{1.25, -1.25, 0.5 / 8388608, 1.5 / 8388608, -1.5 / 8388608, 255, -256}
	samples := []int32{10485760, -10485760, 0, 2, -2, 2139095040, -2147483648}
	out := make([]int32, len(in)+2)
	for i := range out {
		out[i] = 2147483640
	}
	out[0] = 77
	out[len(out)-1] = 88
	Opus_mapping_matrix_multiply_channel_out_int24(nil, &owner.Header, &in[0], 0, 1, &out[1], 1, int32(len(in)))
	for i, sample := range samples {
		want := int32(int64(2147483640) - int64(sample))
		if out[i+1] != want {
			t.Fatalf("i=%d got=%d want=%d", i, out[i+1], want)
		}
	}
	if out[0] != 77 || out[len(out)-1] != 88 {
		t.Fatal("guards")
	}
	coef = 16384
	Opus_mapping_matrix_init(nil, &owner.Header, 1, 1, 0, &coef, 2)
	in = []float32{1.0 / 8388608, -1.0 / 8388608}
	var rounded [2]int32
	Opus_mapping_matrix_multiply_channel_out_int24(nil, &owner.Header, &in[0], 0, 1, &rounded[0], 1, 2)
	if rounded != [2]int32{1, 0} {
		t.Fatal("Q15 half rounding", rounded)
	}
}

// From matrix_short_pointers_test.go.
func TestMatrixShortPointers(t *testing.T) {
	var owner mappingTestBuffer
	coef := int16(-32768)
	Opus_mapping_matrix_init(nil, &owner.Header, 1, 1, 0, &coef, 2)
	Opus_mapping_matrix_multiply_channel_out_short(nil, &owner.Header, nil, 0, 1, nil, 1, 0)
	in := []float32{float32(math.NaN()), float32(math.Inf(1)), float32(math.Inf(-1)), 0.5 / 32768, 1.5 / 32768, -1.5 / 32768, 1.2, -1.2}
	out := make([]int16, len(in)+2)
	for i := range out {
		out[i] = 10
	}
	out[0] = 77
	out[len(out)-1] = 88
	Opus_mapping_matrix_multiply_channel_out_short(nil, &owner.Header, &in[0], 0, 1, &out[1], 1, int32(len(in)))
	want := []int16{77, -32758, -32757, -32758, 10, 8, 12, -32757, -32758, 88}
	if !slices.Equal(out, want) {
		t.Fatal(out)
	}
}

// From matrix_in_float_pointers_test.go.
func TestMatrixInFloatPointers(t *testing.T) {
	var owner mappingTestBuffer
	data := [6]int16{16384, 8192, -16384, 4096, 0, 32767}
	Opus_mapping_matrix_init(nil, &owner.Header, 2, 3, 0, &data[0], 12)
	in := [6]float32{1, 2, 3, 4, 5, 6}
	out := [5]float32{77, 9, 8, 7, 88}
	Opus_mapping_matrix_multiply_channel_in_float(nil, &owner.Header, &in[0], 3, &out[1], 0, 2, 2)
	if out != [5]float32{77, -0.5, 8, -0.5, 88} {
		t.Fatal(out)
	}
	Opus_mapping_matrix_multiply_channel_in_float(nil, &owner.Header, nil, 3, nil, 0, 2, 0)
	Opus_mapping_matrix_multiply_channel_in_float(nil, &owner.Header, nil, 0, &out[1], 0, 0, 3)
	if out[1] != 0 || out[0] != 77 || out[4] != 88 {
		t.Fatal("empty dot product", out)
	}
}

// From matrix_in_int24_pointers_test.go.
func TestMatrixInInt24Pointers(t *testing.T) {
	var owner mappingTestBuffer
	data := [6]int16{16384, 8192, -16384, 4096, 0, 32767}
	Opus_mapping_matrix_init(nil, &owner.Header, 2, 3, 0, &data[0], 12)
	in := [6]int32{10000 * 256, 20000 * 256, 30000 * 256, 4000 * 256, 5000 * 256, 6000 * 256}
	out := [5]float32{77, 9, 8, 7, 88}
	Opus_mapping_matrix_multiply_channel_in_int24(nil, &owner.Header, &in[0], 3, &out[1], 0, 2, 2)
	if out != [5]float32{77, -5000.0 / 32768, 8, -500.0 / 32768, 88} {
		t.Fatal(out)
	}
	Opus_mapping_matrix_multiply_channel_in_int24(nil, &owner.Header, nil, 3, nil, 0, 2, 0)
	Opus_mapping_matrix_multiply_channel_in_int24(nil, &owner.Header, nil, 0, &out[1], 0, 0, 3)
	if out[1] != 0 || out[0] != 77 || out[4] != 88 {
		t.Fatal("empty dot product", out)
	}
}

// From matrix_in_short_pointers_test.go.
func TestMatrixInShortPointers(t *testing.T) {
	var owner mappingTestBuffer
	data := [6]int16{16384, 8192, -16384, 4096, 0, 32767}
	Opus_mapping_matrix_init(nil, &owner.Header, 2, 3, 0, &data[0], 12)
	in := [6]int16{10000, 20000, 30000, 4000, 5000, 6000}
	out := [5]float32{77, 9, 8, 7, 88}
	Opus_mapping_matrix_multiply_channel_in_short(nil, &owner.Header, &in[0], 3, &out[1], 0, 2, 2)
	if out != [5]float32{77, -5000.0 / 32768, 8, -500.0 / 32768, 88} {
		t.Fatal(out)
	}
	Opus_mapping_matrix_multiply_channel_in_short(nil, &owner.Header, nil, 3, nil, 0, 2, 0)
	Opus_mapping_matrix_multiply_channel_in_short(nil, &owner.Header, nil, 0, &out[1], 0, 0, 3)
	if out[1] != 0 || out[0] != 77 || out[4] != 88 {
		t.Fatal("empty dot product", out)
	}
}

// From matrix_init_pointers_test.go.
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
