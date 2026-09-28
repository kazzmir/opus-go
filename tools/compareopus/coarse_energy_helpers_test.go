//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/bits"
	"math/rand"
	"runtime"
	"slices"
	"testing"
	"unsafe"
)

func TestCoarseEnergyAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(3601))
	mode := opuscc.OpusT_OpusCustomMode{FnbEBands: 22}
	for lm := int32(0); lm < 4; lm++ {
		for intra := int32(0); intra < 2; intra++ {
			for channels := int32(1); channels <= 2; channels++ {
				for _, remaining := range []int32{0, 1, 2, 14, 15, 40, 400} {
					for trial := 0; trial < 20; trial++ {
						data := make([]byte, 64)
						rng.Read(data)
						var g opuscc.OpusT_ec_dec
						opuscc.Opus_ec_dec_init(nil, uintptr(unsafe.Pointer(&g)), uintptr(unsafe.Pointer(&data[0])), uint32(len(data)))
						g.Fnbits_total = int32(g.Fstorage*8) + int32(bits.Len32(g.Frng)) - remaining
						c := g
						energy := make([]float32, 22*channels)
						for i := range energy {
							energy[i] = float32(rng.Intn(1000)-700) / 16
						}
						want := slices.Clone(energy)
						start := int32(trial % 4)
						end := int32(22 - trial%3)
						opuscc.Opus_unquant_coarse_energy(nil, &mode, start, end, &energy[0], intra, &g, channels, lm)
						nativeEnergyDecode(&c, data, want, 22, start, end, channels, 0, intra, lm, nil, nil)
						if !slices.Equal(energy, want) || g != c {
							t.Fatalf("LM=%d intra=%d channels=%d remaining=%d trial=%d Go=%v C=%v state=%+v C=%+v", lm, intra, channels, remaining, trial, energy, want, g, c)
						}
						runtime.KeepAlive(data)
					}
				}
			}
		}
	}
}
