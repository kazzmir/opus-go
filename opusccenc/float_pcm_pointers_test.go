package opusccenc

import (
	"math"
	"runtime"
	"slices"
	"testing"
)

func flpVectorBitsEqual(a, b []float32) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
			return false
		}
	}
	return true
}

func TestScaleCopyVectorFLPPointers(t *testing.T) {
	for _, size := range []int32{0, 1, 3, 4, 5, 8, 31, 65537, -1, -65536} {
		for _, shift := range []int32{-1, 0, 1} {
			length := max(size, size&0xFFFC)
			data := make([]float32, length+3)
			for i := range data {
				data[i] = float32(i%19-9) / 7
			}
			want := slices.Clone(data)
			for i := int32(0); i < length; i++ {
				want[1+shift+i] = float32(.75 * want[1+i])
			}
			runtime.GC()
			silkScaleCopyVectorFLP(&data[1+shift], &data[1], .75, size)
			if !flpVectorBitsEqual(data, want) {
				t.Fatal("scale-copy alias/guards", size, shift)
			}
		}
	}
	silkScaleCopyVectorFLP(nil, nil, 1, 0)
	silkScaleCopyVectorFLP(nil, nil, 1, -65536)
}

func TestScaleVectorFLPPointers(t *testing.T) {
	for _, size := range []int32{0, 1, 3, 4, 5, 8, 31, 65537, -1, -65536} {
		length := max(size, size&0xFFFC)
		data := make([]float32, length+2)
		for i := range data {
			data[i] = float32(i%19-9) / 7
		}
		want := slices.Clone(data)
		for i := int32(0); i < length; i++ {
			want[1+i] = float32(want[1+i] * -.75)
		}
		runtime.GC()
		silkScaleVectorFLP(&data[1], -.75, size)
		if !flpVectorBitsEqual(data, want) {
			t.Fatal("scale/guards", size)
		}
	}
	for _, gain := range []float32{0, float32(math.Copysign(0, -1)), .75} {
		data := []float32{float32(math.Copysign(0, -1)), 0, float32(math.Inf(1)), float32(math.Inf(-1)), math.Float32frombits(0x7fc12345)}
		want := slices.Clone(data)
		for i := range want {
			want[i] = float32(want[i] * gain)
		}
		silkScaleVectorFLP(&data[0], gain, int32(len(data)))
		if !flpVectorBitsEqual(data, want) {
			t.Fatal("scale IEEE", gain)
		}
	}
	silkScaleVectorFLP(nil, 1, 0)
	silkScaleVectorFLP(nil, 1, -65536)
}

func TestFloatPCMPointers(t *testing.T) {
	input := []float32{-2, -1, -2.5 / 32768, -1.5 / 32768, -0.5 / 32768, 0, 0.5 / 32768, 1.5 / 32768, 2.5 / 32768, 1, 2, float32(math.NaN()), float32(math.Inf(-1)), float32(math.Inf(1))}
	want := []int16{-32768, -32768, -2, -2, 0, 0, 0, 2, 2, 32767, 32767, -32768, -32768, 32767}
	out := make([]int16, len(input)+2)
	out[0], out[len(out)-1] = 111, 222
	Opus_celt_float2int16_c(nil, &input[0], &out[1], int32(len(input)))
	if !slices.Equal(out[1:len(out)-1], want) || out[0] != 111 || out[len(out)-1] != 222 {
		t.Fatalf("got=%v want=%v", out, want)
	}
	Opus_celt_float2int16_c(nil, nil, nil, 0)
	Opus_celt_float2int16_c(nil, nil, nil, -1)
}
