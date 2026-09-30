//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"math/rand"
	"testing"
)

func TestFIRAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(710))
	for _, ord := range []int{3, 4, 5, 7, 8, 24} {
		for _, n := range []int{0, 1, 2, 3, 4, 5, 7, 8, 31, 64} {
			for trial := 0; trial < 12; trial++ {
				input := make([]float32, ord+n+2)
				coeff := make([]float32, ord)
				for i := range input {
					input[i] = float32(rng.NormFloat64() * 3)
				}
				for i := range coeff {
					coeff[i] = float32(rng.NormFloat64())
				}
				if trial == 0 {
					for i := range input {
						input[i] = math.Float32frombits(uint32(i%2) << 31)
					}
				}
				if trial == 1 {
					for i := range input {
						input[i] = math.Float32frombits(uint32(i + 1))
					}
				}
				before := append([]float32(nil), input...)
				goOut := make([]float32, n+2)
				for i := range goOut {
					goOut[i] = 77
				}
				cOut := append([]float32(nil), goOut...)
				opuscc.Opus_celt_fir_c(nil, &input[ord], &coeff[0], &goOut[1], int32(n), int32(ord), 0)
				nativeFIR(input, coeff, cOut[1:], int32(n), int32(ord))
				if !sameFloatBits(goOut, cOut) || !sameFloatBits(input, before) {
					t.Fatal(ord, n, trial, goOut, cOut)
				}
			}
		}
	}
	// Partial overlaps are permitted (only exact x==y is asserted against).
	for _, offset := range []int{3, 5} {
		g := make([]float32, 40)
		for i := range g {
			g[i] = float32(i) * 0.17
		}
		c := append([]float32(nil), g...)
		coeff := []float32{0.1, 0.2, 0.3, 0.4}
		opuscc.Opus_celt_fir_c(nil, &g[4], &coeff[0], &g[offset], 24, 4, 0)
		nativeFIR(c, coeff, c[offset:], 24, 4)
		if !sameFloatBits(g, c) {
			t.Fatal("overlap", offset, g, c)
		}
	}
}

func TestLPCAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, p := range []int{1, 4, 16, 24} {
		for trial := 0; trial < 100; trial++ {
			signal := make([]float32, 128)
			for i := range signal {
				signal[i] = float32(rng.NormFloat64())
			}
			ac := make([]float32, p+1)
			for lag := range ac {
				for i := lag; i < len(signal); i++ {
					ac[lag] += float32(signal[i] * signal[i-lag])
				}
			}
			if trial == 0 {
				clear(ac)
			}
			if trial == 1 {
				for i := range ac {
					ac[i] = 1
				}
			}
			g, c := make([]float32, p), make([]float32, p)
			opuscc.Opus__celt_lpc(nil, &g[0], &ac[0], int32(p))
			nativeLPC(c, ac)
			for i := range g {
				if math.Float32bits(g[i]) != math.Float32bits(c[i]) {
					t.Fatalf("p=%d trial=%d tap=%d Go=%g C=%g", p, trial, i, g[i], c[i])
				}
			}
		}
	}
}
