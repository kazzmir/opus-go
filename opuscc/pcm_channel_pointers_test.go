package opuscc

import (
	"math"
	"slices"
	"testing"
)

// From channel_float_pointers_test.go.
func TestChannelFloatPointers(t *testing.T) {
	opus_copy_channel_out_float(nil, 0, 0, nil, 0, 0)
	b := [7]float32{77, 1, 2, 3, 4, 5, 88}
	opus_copy_channel_out_float(&b[1], 1, 0, &b[1], 1, 5)
	opus_copy_channel_out_float(&b[1], 2, 1, nil, 1, 2)
	if b != [7]float32{77, 1, 0, 3, 0, 5, 88} {
		t.Fatal(b)
	}
	v := math.Float32frombits(0x7fc01234)
	var out float32
	opus_copy_channel_out_float(&out, 0, 0, &v, 0, 3)
	if math.Float32bits(out) != 0x7fc01234 {
		t.Fatal(out)
	}
	b = [7]float32{77, 1, 2, 3, 4, 5, 88}
	opus_copy_channel_out_float(&b[2], 1, 0, &b[1], 1, 4)
	if b != [7]float32{77, 1, 1, 1, 1, 1, 88} {
		t.Fatal("forward overlap", b)
	}
}

// From channel_int24_pointers_test.go.
func TestChannelInt24Pointers(t *testing.T) {
	opus_copy_channel_out_int24(nil, nil, 0, 0, nil, 0, 0)
	src := [8]float32{-2, -.5, -.5 / 8388608, .5 / 8388608, 1.5 / 8388608, 2.5 / 8388608, 2, 255}
	var dst [10]int32
	dst[0] = 77
	dst[9] = 88
	opus_copy_channel_out_int24(nil, &dst[1], 1, 0, &src[0], 1, 8)
	if dst != [10]int32{77, -16777216, -4194304, 0, 0, 2, 2, 16777216, 2139095040, 88} {
		t.Fatal(dst)
	}
	opus_copy_channel_out_int24(nil, &dst[1], 2, 1, nil, 0, 4)
	if dst[2] != 0 || dst[4] != 0 || dst[6] != 0 || dst[8] != 0 || dst[0] != 77 || dst[9] != 88 {
		t.Fatal(dst)
	}
}

// From channel_short_pointers_test.go.
func TestChannelShortPointers(t *testing.T) {
	opus_copy_channel_out_short(nil, nil, 0, 0, nil, 0, 0)
	src := [8]float32{-2, -.5, -.5 / 32768, .5 / 32768, 1.5 / 32768, 2.5 / 32768, 2, math.Float32frombits(0x7fc01234)}
	var dst [10]int16
	dst[0] = 77
	dst[9] = 88
	opus_copy_channel_out_short(nil, &dst[1], 1, 0, &src[0], 1, 8)
	if dst != [10]int16{77, -32768, -16384, 0, 0, 2, 2, 32767, -32768, 88} {
		t.Fatal(dst)
	}
	opus_copy_channel_out_short(nil, &dst[1], 2, 1, nil, 0, 4)
	if dst[2] != 0 || dst[4] != 0 || dst[6] != 0 || dst[8] != 0 || dst[0] != 77 || dst[9] != 88 {
		t.Fatal(dst)
	}
}

// From float_pcm_pointers_test.go.
func TestFloatPCMPointers(t *testing.T) {
	input := []float32{-2, -1, -32767.5 / 32768, -2.5 / 32768, -1.5 / 32768, -0.5 / 32768, 0, 0.5 / 32768, 1.5 / 32768, 2.5 / 32768, 32766.5 / 32768, 1, 2, float32(math.Inf(-1)), float32(math.Inf(1))}
	want := []int16{-32768, -32768, -32768, -2, -2, 0, 0, 0, 2, 2, 32766, 32767, 32767, -32768, 32767}
	before := slices.Clone(input)
	output := make([]int16, len(input)+2)
	output[0], output[len(output)-1] = 111, 222
	Opus_celt_float2int16_c(nil, &input[0], &output[1], int32(len(input)))
	if !slices.Equal(output[1:len(output)-1], want) || output[0] != 111 || output[len(output)-1] != 222 {
		t.Fatalf("got=%v want=%v", output, want)
	}
	if !slices.Equal(input, before) {
		t.Fatal("input changed")
	}
	Opus_celt_float2int16_c(nil, nil, nil, 0)
	Opus_celt_float2int16_c(nil, nil, nil, -1)
}

// From soft_clip_pointers_test.go.
func TestSoftClipPointers(t *testing.T) {
	x := []float32{77, 0.5, -0.5, 1.5, -1.5, 2, -2, 0, 0, 88}
	mem := [4]float32{123, 0, 0, 456}
	Opus_opus_pcm_soft_clip(nil, &x[1], 4, 2, &mem[1])
	for _, v := range x[1:9] {
		if v < -1 || v > 1 {
			t.Fatal("range", x)
		}
	}
	a := float32(0.25)
	a += float32(a * float32(2.4e-7))
	if mem != [4]float32{123, -a, a, 456} || x[0] != 77 || x[9] != 88 {
		t.Fatal("state/guards", mem, x)
	}
	var empty [2]float32
	Opus_opus_pcm_soft_clip(nil, &empty[0], 1, 2, &mem[1])
	if mem[1] != 0 || mem[2] != 0 {
		t.Fatal("history not cleared")
	}
	value := float32(2)
	state := float32(0.25)
	Opus_opus_pcm_soft_clip(nil, nil, 1, 1, &state)
	Opus_opus_pcm_soft_clip(nil, &value, 1, 1, nil)
	Opus_opus_pcm_soft_clip(nil, &value, 0, 1, &state)
	if value != 2 || state != 0.25 {
		t.Fatal("no-op")
	}
	z := math.Float32frombits(0x80000000)
	state = 0
	Opus_opus_pcm_soft_clip(nil, &z, 1, 1, &state)
	if math.Float32bits(z) != 0x80000000 {
		t.Fatal("signed zero")
	}
}

// From fade_pointers_test.go.
func TestFadePointers(t *testing.T) {
	smooth_fade(nil, nil, nil, 0, 2, nil, 48000)
	a := [4]float32{1, 2, 3, 4}
	b := [4]float32{5, 6, 7, 8}
	w := [3]float32{0, 99, 0.5}
	out := [6]float32{77, 0, 0, 0, 0, 88}
	smooth_fade(&a[0], &b[0], &out[1], 2, 2, &w[0], 24000)
	if out != [6]float32{77, 1, 2, 4, 5, 88} {
		t.Fatal(out)
	}
	smooth_fade(&a[0], &b[0], &a[0], 2, 2, &w[0], 24000)
	if a != [4]float32{1, 2, 4, 5} {
		t.Fatal("alias", a)
	}
}
