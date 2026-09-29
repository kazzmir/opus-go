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

func TestFineEnergyAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(3602))
	mode := opuscc.OpusT_OpusCustomMode{FnbEBands: 21}
	for channels := int32(1); channels <= 2; channels++ {
		for trial := 0; trial < 1000; trial++ {
			data := make([]byte, 64)
			rng.Read(data)
			var g opuscc.OpusT_ec_dec
			opuscc.Opus_ec_dec_init(nil, &g, &data[0], uint32(len(data)))
			g.Fnbits_total = int32(g.Fstorage*8) + int32(bits.Len32(g.Frng)) - int32(trial%129)
			c := g
			energy := make([]float32, 21*channels)
			for i := range energy {
				energy[i] = float32(rng.Intn(512)-256) / 16
			}
			want := slices.Clone(energy)
			extra, prev := make([]int32, 21), make([]int32, 21)
			for i := range extra {
				extra[i] = int32(rng.Intn(16) - 1)
				prev[i] = int32(rng.Intn(15))
			}
			if trial%2 == 0 {
				prev = nil
			}
			eb, pb := slices.Clone(extra), slices.Clone(prev)
			start, end := int32(trial%5), int32(21-trial%4)
			opuscc.Opus_unquant_fine_energy(nil, &mode, start, end, &energy[0], unsafe.SliceData(prev), &extra[0], &g, channels)
			nativeEnergyDecode(&c, data, want, 21, start, end, channels, 1, 0, 0, prev, extra)
			if !slices.Equal(energy, want) || g != c || !slices.Equal(extra, eb) || !slices.Equal(prev, pb) {
				t.Fatalf("channels=%d trial=%d Go=%v C=%v state=%+v C=%+v", channels, trial, energy, want, g, c)
			}
			runtime.KeepAlive(data)
		}
	}
}
