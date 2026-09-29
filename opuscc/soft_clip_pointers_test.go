package opuscc

import (
	"math"
	"testing"
)

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
