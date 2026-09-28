//go:build compareopus && cgo

package main

import (
	"math"
	"math/rand"
	"slices"
	"testing"

	"github.com/kazzmir/opus-go/opuscc"
)

func TestLimiterAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	goPCM := make([]float32, 4096)
	for i := range goPCM {
		goPCM[i] = float32(rng.NormFloat64() * 3)
	}
	for i, bits := range []uint32{0x80000000, 0, 0x7f800000, 0xff800000, 0x7fc01234, 0xffc01234, 0x40000000, 0xc0000000} {
		goPCM[i] = math.Float32frombits(bits)
	}
	cPCM := slices.Clone(goPCM)
	goRet := opuscc.Opus_opus_limit2_checkwithin1_c(nil, &goPCM[0], int32(len(goPCM)))
	cRet := nativeLimit2(&cPCM[0], int32(len(cPCM)))
	if goRet != cRet {
		t.Fatalf("return: Go=%d C=%d", goRet, cRet)
	}
	for i := range goPCM {
		if math.Float32bits(goPCM[i]) != math.Float32bits(cPCM[i]) {
			t.Fatalf("index %d: Go=%08x C=%08x", i, math.Float32bits(goPCM[i]), math.Float32bits(cPCM[i]))
		}
	}
	for _, n := range []int32{-1, 0} {
		if goRet, cRet := opuscc.Opus_opus_limit2_checkwithin1_c(nil, nil, n), nativeLimit2(nil, n); goRet != cRet {
			t.Fatalf("count=%d: Go=%d C=%d", n, goRet, cRet)
		}
	}
}
