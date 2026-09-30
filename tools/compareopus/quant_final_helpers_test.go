//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"slices"
	"testing"
	"unsafe"
)

func TestQuantFinalAgainstC(t *testing.T) {
	const bands = 8
	for _, channels := range []int32{1, 2} {
		for left := int32(0); left <= 20; left++ {
			for _, capacity := range []uint32{0, 1, 32} {
				for variant := 0; variant < 3; variant++ {
					g := make([]byte, capacity+2)
					for i := range g {
						g[i] = 77
					}
					c := slices.Clone(g)
					old := make([]float32, channels*bands)
					err := make([]float32, len(old))
					for i := range old {
						old[i] = float32(i) * .25
						err[i] = []float32{-.6, math.Float32frombits(0x80000000), 0, .25, math.Float32frombits(0x7fc01234), .5}[i%6]
					}
					co, ce := slices.Clone(old), slices.Clone(err)
					if variant == 1 {
						err = old
						ce = co
					}
					if variant == 2 {
						old = nil
						co = nil
					}
					quant := []int32{0, 1, 2, 3, 4, 5, 7, 8}
					priority := []int32{1, 0, 0, 1, 2, 0, 1, 0}
					m := opuscc.OpusT_OpusCustomMode{FnbEBands: bands}
					var ge opuscc.OpusT_ec_enc
					opuscc.Opus_ec_enc_init(nil, &ge, &g[1], capacity)
					ne := ge
					opuscc.Opus_quant_energy_finalise(nil, &m, 1, bands, unsafe.SliceData(old), unsafe.SliceData(err), unsafe.SliceData(quant), unsafe.SliceData(priority), left, &ge, channels)
					nativeEnergyEncode(&ne, c[1:], co, ce, bands, 1, bands, channels, quant, priority, 1, left)
					if ge != ne || !slices.Equal(g, c) || !sameFloatBits(old, co) || !sameFloatBits(err, ce) {
						t.Fatal(channels, left, capacity, variant, ge, ne, err, ce)
					}
				}
			}
		}
	}
}
