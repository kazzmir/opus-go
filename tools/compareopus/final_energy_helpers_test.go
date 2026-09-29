//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"runtime"
	"slices"
	"testing"
	"unsafe"
)

func TestFinalEnergyAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(3603))
	mode := opuscc.OpusT_OpusCustomMode{FnbEBands: 21}
	for channels := int32(1); channels <= 2; channels++ {
		for trial := 0; trial < 1000; trial++ {
			data := make([]byte, []int{0, 1, 16}[trial%3])
			rng.Read(data)
			var g opuscc.OpusT_ec_dec
			opuscc.Opus_ec_dec_init(nil, &g, unsafe.SliceData(data), uint32(len(data)))
			c := g
			energy := make([]float32, 21*channels)
			for i := range energy {
				energy[i] = float32(rng.Intn(512)-256) / 16
			}
			if trial%2 == 0 {
				energy = nil
			}
			want := slices.Clone(energy)
			quant, priority := make([]int32, 21), make([]int32, 21)
			for i := range quant {
				quant[i] = int32(rng.Intn(10))
				priority[i] = int32(rng.Intn(2))
			}
			qb, pb := slices.Clone(quant), slices.Clone(priority)
			start, end := int32(trial%5), int32(21-trial%4)
			budget := int32(trial%48 - 2)
			opuscc.Opus_unquant_energy_finalise(nil, &mode, start, end, unsafe.SliceData(energy), &quant[0], &priority[0], budget, &g, channels)
			nativeEnergyDecode(&c, data, want, 21, start, end, channels, 2, budget, 0, quant, priority)
			if !slices.Equal(energy, want) || g != c || !slices.Equal(quant, qb) || !slices.Equal(priority, pb) {
				t.Fatalf("channels=%d trial=%d Go=%v C=%v state=%+v C=%+v", channels, trial, energy, want, g, c)
			}
			runtime.KeepAlive(data)
		}
	}
}
