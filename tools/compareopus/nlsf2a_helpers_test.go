//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
)

func TestNLSF2AAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(3402))
	for _, d := range []int{10, 16} {
		for trial := 0; trial < 2000; trial++ {
			in := make([]int16, d)
			for i := range in {
				switch trial % 5 {
				case 0:
					in[i] = int16(rng.Intn(32768))
				case 1:
					in[i] = int16(2000 + i*10)
				case 2:
					in[i] = int16(i)
				case 3:
					in[i] = int16(32767 - d + i)
				case 4:
					in[i] = int16((i+1)*32768/(d+1) + rng.Intn(129) - 64)
				}
			}
			slices.Sort(in)
			before := slices.Clone(in)
			g, c := make([]int16, d), make([]int16, d)
			opuscc.Opus_silk_NLSF2A(nil, &g[0], &in[0], int32(d), 0)
			nativeNLSF2A(c, in)
			if !slices.Equal(g, c) || !slices.Equal(in, before) {
				t.Fatalf("d=%d trial=%d Go=%v C=%v input=%v", d, trial, g, c, in)
			}
			if trial%100 == 0 {
				alias := slices.Clone(in)
				nativeNLSF2A(alias, alias)
				if !slices.Equal(g, alias) {
					t.Fatal("C in-place output differs")
				}
			}
		}
	}
}
