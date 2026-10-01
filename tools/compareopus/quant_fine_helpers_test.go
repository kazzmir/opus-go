//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"math/bits"
	"math/rand"
	"slices"
	"testing"
	"unsafe"
)

func TestCoarseEnergyDriverAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(117679))
	// Full-band output initializes every C candidate-error slot. Partial-band
	// intra copies include indeterminate C scratch, which is not a defined oracle.
	for _, nb := range []int32{3, 21} {
		for channels := int32(1); channels <= 2; channels++ {
			for lm := int32(0); lm <= 3; lm++ {
				for _, budget := range []uint32{0, 1, 3, 15, 30, 120, 400} {
					for trial := 0; trial < 16; trial++ {
						n := nb * channels
						g := make([]float32, 3*n+3)
						for i := range g {
							g[i] = float32(rng.Float64()*32 - 26)
						}
						g[0] = 77
						g[len(g)-1] = 88
						energy, old, errors := int32(1), 1+n, 1+2*n
						delay := int32(1 + 3*n)
						g[delay] = float32(trial * 4)
						if trial%4 == 1 {
							old = energy
						}
						if trial%4 == 2 {
							errors = old
						}
						if trial%4 == 3 {
							delay = old
						}
						c := slices.Clone(g)
						capacity := []int{0, 1, 4, 64}[trial%4]
						if trial == 9 || trial == 13 {
							capacity = 64
						} // Seeded prefix plus two-pass rollback space.
						gb := make([]byte, max(capacity, 1))
						for i := range gb {
							gb[i] = 99
						}
						var ge opuscc.OpusT_ec_enc
						opuscc.Opus_ec_enc_init(nil, &ge, &gb[0], uint32(capacity))
						if trial >= 8 {
							for i := 0; i < 6; i++ {
								opuscc.Opus_ec_enc_uint(nil, &ge, uint32(i*17), 256)
							}
							opuscc.Opus_ec_enc_bits(nil, &ge, 5, 3)
						}
						cb := slices.Clone(gb)
						ce := ge
						two := int32(trial % 2)
						force := int32((trial / 2) % 2)
						lfe := int32((trial / 4) % 2)
						available := int32(capacity)
						if trial%3 == 0 {
							available = 64
						}
						loss := int32(trial * 3)
						eff := nb - int32(trial%2)
						mode := opuscc.OpusT_OpusCustomMode{FnbEBands: nb}
						opuscc.Opus_quant_coarse_energy(nil, &mode, 0, nb, eff, &g[energy], &g[old], budget, &g[errors], &ge, channels, lm, available, force, &g[delay], two, loss, lfe)
						opuscc.Opus_ec_enc_done(nil, &ge)
						nativeCoarseEnergyDriver(&ce, cb, &c[energy], &c[old], &c[errors], nb, 0, nb, eff, budget, channels, lm, available, force, &c[delay], two, loss, lfe)
						ge.Fbuf = nil
						ce.Fbuf = nil
						if ge != ce || !slices.Equal(gb, cb) || !sameFloatBits(g, c) {
							t.Fatal(nb, channels, lm, budget, trial, ge, ce, g, c)
						}
					}
				}
			}
		}
	}
}

func TestCoarseEnergyImplAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(39831))
	for _, nb := range []int32{3, 24} {
		for channels := int32(0); channels <= 2; channels++ {
			for lm := int32(0); lm <= 3; lm++ {
				for intra := int32(0); intra <= 1; intra++ {
					for _, budget := range []int32{0, 1, 2, 3, 14, 15, 16, 24, 30, 120, 400} {
						for trial := 0; trial < 6; trial++ {
							n := nb * max(channels, 1)
							g := make([]float32, 3*n+2)
							for i := range g {
								g[i] = float32(rng.Float64()*40 - 30)
							}
							g[0] = 77
							g[len(g)-1] = 88
							c := slices.Clone(g)
							energy, old, errors := int32(1), 1+n, 1+2*n
							if trial%3 == 1 {
								old = energy
							}
							if trial%3 == 2 {
								errors = old
							}
							capacity := []int{0, 1, 4, 32}[trial%4]
							gb := make([]byte, max(capacity, 1))
							for i := range gb {
								gb[i] = 99
							}
							cb := slices.Clone(gb)
							var ge opuscc.OpusT_ec_enc
							opuscc.Opus_ec_enc_init(nil, &ge, &gb[0], uint32(capacity))
							ce := ge
							start, end := int32(trial%2), nb
							if trial == 5 {
								end = start
							}
							decay := float32(16)
							if trial == 3 {
								decay = 0
							}
							lfe := int32(trial % 2)
							tell := int32(ge.Fnbits_total) - int32(32-bits.LeadingZeros32(ge.Frng))
							result := opuscc.CompareCoarseEnergyImpl(nb, start, end, &g[energy], &g[old], budget, tell, &g[errors], &ge, channels, lm, intra, decay, lfe)
							opuscc.Opus_ec_enc_done(nil, &ge)
							native := nativeCoarseEnergyImpl(&ce, cb, &c[energy], &c[old], &c[errors], nb, start, end, budget, tell, channels, lm, intra, decay, lfe)
							ge.Fbuf = nil
							ce.Fbuf = nil
							if result != native || ge != ce || !slices.Equal(gb, cb) || !sameFloatBits(g, c) {
								t.Fatal(nb, channels, lm, intra, budget, trial, result, native, ge, ce)
							}
						}
					}
				}
			}
		}
	}
}

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
