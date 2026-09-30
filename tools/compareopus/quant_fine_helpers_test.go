//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"slices"
	"testing"
	"unsafe"
)

func sameFloatBits(a, b []float32) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Float32bits(a[i]) != math.Float32bits(b[i]) {
			return false
		}
	}
	return true
}

func TestQuantFineAgainstC(t *testing.T) {
	const bands = 8
	for _, channels := range []int32{1, 2} {
		for _, capacity := range []uint32{0, 1, 3, 32} {
			for _, alias := range []bool{false, true} {
				for _, withPrev := range []bool{false, true} {
					g := make([]byte, capacity+2)
					for i := range g {
						g[i] = 77
					}
					c := slices.Clone(g)
					old := make([]float32, channels*bands)
					err := make([]float32, len(old))
					for i := range old {
						old[i] = float32(i) * .25
						err[i] = []float32{-.6, -.5, -.25, math.Float32frombits(0x80000000), 0, .0625, .25, .5, .6}[i%9]
					}
					co, ce := slices.Clone(old), slices.Clone(err)
					if alias {
						err = old
						ce = co
					}
					extra := []int32{0, 1, 2, 3, 4, 5, 7, 8}
					var prev []int32
					if withPrev {
						prev = []int32{0, 0, 1, 2, 4, 5, 6, 7}
					}
					m := opuscc.OpusT_OpusCustomMode{FnbEBands: bands}
					var ge opuscc.OpusT_ec_enc
					opuscc.Opus_ec_enc_init(nil, &ge, &g[1], capacity)
					ne := ge
					opuscc.Opus_quant_fine_energy(nil, &m, 1, bands, unsafe.SliceData(old), unsafe.SliceData(err), unsafe.SliceData(prev), unsafe.SliceData(extra), &ge, channels)
					nativeEnergyEncode(&ne, c[1:], co, ce, bands, 1, bands, channels, prev, extra)
					if ge != ne || !slices.Equal(g, c) || !sameFloatBits(old, co) || !sameFloatBits(err, ce) {
						t.Fatal(channels, capacity, alias, withPrev, ge, ne, old, co, err, ce)
					}
				}
			}
		}
	}
}
