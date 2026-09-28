//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"github.com/kazzmir/opus-go/opusccenc"
	"math"
	"math/rand"
	"testing"
)

func TestFloatPCMAgainstC(t *testing.T) {
	input := []float32{float32(math.Inf(-1)), float32(math.Inf(1))}
	// All int16 rounding midpoints, plus values on either side.
	for i := -32769; i <= 32768; i++ {
		x := (float32(i) + 0.5) / 32768
		input = append(input, x, math.Nextafter32(x, float32(math.Inf(-1))), math.Nextafter32(x, float32(math.Inf(1))))
	}
	rng := rand.New(rand.NewSource(2026))
	for i := 0; i < 10000; i++ {
		input = append(input, 4*rng.Float32()-2)
	}
	g, c := make([]int16, len(input)), make([]int16, len(input))
	opuscc.Opus_celt_float2int16_c(nil, &input[0], &g[0], int32(len(input)))
	enc := make([]int16, len(input))
	opusccenc.Opus_celt_float2int16_c(nil, &input[0], &enc[0], int32(len(input)))
	nativeFloatPCM(input, c)
	for i := range g {
		if g[i] != c[i] || enc[i] != c[i] {
			t.Fatalf("i=%d input=%g decoder=%d encoder=%d C=%d", i, input[i], g[i], enc[i], c[i])
		}
	}
	// The encoder copy preserves C's MAX32 NaN semantics.
	nan := []float32{float32(math.NaN())}
	opusccenc.Opus_celt_float2int16_c(nil, &nan[0], &enc[0], 1)
	nativeFloatPCM(nan, c)
	if enc[0] != c[0] {
		t.Fatalf("NaN: encoder=%d C=%d", enc[0], c[0])
	}
}
