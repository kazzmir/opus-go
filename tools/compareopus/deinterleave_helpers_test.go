//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"math/rand"
	"slices"
	"testing"
)

func TestBandContextLayoutAgainstC(t *testing.T) {
	if nativeBandContextLayout() != opuscc.CompareBandContextLayout() {
		t.Fatal(nativeBandContextLayout(), opuscc.CompareBandContextLayout())
	}
}

func TestThetaAgainstC(t *testing.T) {
	for _, N := range []int32{2, 4, 16} {
		for _, B := range []int32{1, 2} {
			for _, LM := range []int32{-1, 0, 3} {
				for _, stereo := range []int32{0, 1} {
					for _, budget := range []int32{0, 16, 64, 160} {
						for _, intensity := range []int32{0, 1} {
							for alias := int32(0); alias < 4; alias++ {
								cfg := [12]int32{N, B, B, LM, stereo, budget, 15, intensity, 24, 300, alias % 2, alias}
								data := []byte{17, 255, 88, 1, 192, 0, 77, 43}
								var ge opuscc.OpusT_ec_ctx
								opuscc.Opus_ec_dec_init(nil, &ge, &data[0], uint32(len(data)))
								ce := ge
								g := opuscc.CompareTheta(&ge, cfg)
								c := nativeTheta(&ce, data, cfg)
								if g != c || ge != ce {
									t.Fatal(cfg, g, c, ge, ce)
								}
							}
						}
					}
				}
			}
		}
	}
}

func TestMonoBandAgainstC(t *testing.T) {
	for LM := int32(0); LM <= 3; LM++ {
		N := int32(1) << LM
		for _, B := range []int32{1, N} {
			for _, tf := range []int32{-1, 0, 1} {
				if tf > 0 && B < 2 {
					continue
				}
				for _, budget := range []int32{0, 8, 24, 80} {
					cfg := [8]int32{N, B, LM, budget, 1, tf, 300, (1 << B) - 1}
					data := []byte{17, 255, 88, 1, 192, 0, 77, 43}
					var ge opuscc.OpusT_ec_ctx
					opuscc.Opus_ec_dec_init(nil, &ge, &data[0], uint32(len(data)))
					ce := ge
					x, cx, low, clow := make([]float32, N+2), make([]float32, N+2), make([]float32, N+2), make([]float32, N+2)
					for i := range x {
						x[i] = float32(i+1) / 17
						cx[i] = x[i]
						low[i] = 77
						clow[i] = 77
					}
					gm, gs := opuscc.CompareMonoBand(&ge, &x[1], &low[1], cfg)
					cm, cs := nativeMonoBand(&ce, data, cx[1:len(cx)-1], clow[1:len(clow)-1], cfg)
					if gm != cm || gs != cs || ge != ce || !sameFloatBits(x, cx) || !sameFloatBits(low, clow) {
						t.Fatal(cfg, gm, cm, gs, cs, x, cx, low, clow, ge, ce)
					}
				}
			}
		}
	}
}

func TestQuantN1AgainstC(t *testing.T) {
	for encode := int32(0); encode <= 1; encode++ {
		for resynth := int32(0); resynth <= 1; resynth++ {
			for _, remaining := range []int32{-1, 0, 7, 8, 15, 16, 24} {
				for _, capacity := range []int{0, 1, 8} {
					for _, y := range []int32{-1, 0, 1} {
						for _, low := range []int32{-1, 0, 1, 2} {
							gb := make([]byte, max(capacity, 1))
							for i := range gb {
								gb[i] = byte(0x81 + i)
							}
							cb := slices.Clone(gb)
							var g opuscc.OpusT_ec_ctx
							if encode != 0 {
								opuscc.Opus_ec_enc_init(nil, &g, &gb[0], uint32(capacity))
							} else {
								opuscc.Opus_ec_dec_init(nil, &g, &gb[0], uint32(capacity))
							}
							c := g
							gv := []float32{-0.375, math.Float32frombits(0x80000000), 77}
							cv := slices.Clone(gv)
							gr, cr := remaining, remaining
							var yp, lp *float32
							if y >= 0 {
								yp = &gv[y]
							}
							if low >= 0 {
								lp = &gv[low]
							}
							mask := opuscc.CompareQuantN1(encode, resynth, &gr, &g, &gv[0], yp, lp)
							if encode != 0 {
								opuscc.Opus_ec_enc_done(nil, &g)
							}
							native := nativeQuantN1(&c, cb, cv, encode, resynth, &cr, y, low)
							g.Fbuf = nil
							c.Fbuf = nil
							if mask != native || gr != cr || g != c || !slices.Equal(gb, cb) || !sameFloatBits(gv, cv) {
								t.Fatal(encode, resynth, remaining, capacity, y, low, g, c, gv, cv)
							}
						}
					}
				}
			}
		}
	}
}

func TestAntiCollapseAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(9171))
	bands := []int16{0, 1, 3, 8, 16}
	for lm := int32(0); lm <= 3; lm++ {
		for channels := int32(0); channels <= 2; channels++ {
			for encode := int32(0); encode <= 1; encode++ {
				for trial := 0; trial < 20; trial++ {
					size := int32(16) << lm
					g := make([]float32, max(channels, 1)*size+2)
					for i := range g {
						g[i] = float32(rng.NormFloat64() * .3)
					}
					g[0] = 77
					g[len(g)-1] = 88
					c := slices.Clone(g)
					energy := make([]float32, max(channels, 1)*4)
					p1, p2 := make([]float32, 8), make([]float32, 8)
					for i := range energy {
						energy[i] = float32(rng.Float64()*60 - 30)
					}
					for i := range p1 {
						p1[i] = float32(rng.Float64()*20 - 10)
						p2[i] = float32(rng.Float64()*20 - 10)
					}
					pulse := []int32{0, 3, 50, 1000}
					masks := make([]byte, max(channels*4, 1))
					rng.Read(masks)
					seed := rng.Uint32()
					start, end := int32(trial%4), int32(4)
					opuscc.Opus_anti_collapse(nil, &bands[0], 4, &g[1], &masks[0], lm, channels, size, start, end, &energy[0], &p1[0], &p2[0], &pulse[0], seed, encode, 0)
					nativeAntiCollapse(bands, 4, c[1:len(c)-1], masks, lm, channels, size, start, end, energy, p1, p2, pulse, seed, encode)
					if !sameFloatBits(g, c) {
						for i := range g {
							if math.Float32bits(g[i]) != math.Float32bits(c[i]) {
								t.Fatal(lm, channels, encode, trial, i, g[i], c[i])
							}
						}
					}
				}
			}
		}
	}
}

func TestSpreadingAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(28179))
	bands := []int16{0, 1, 4, 13, 32}
	for _, mult := range []int32{1, 2, 4, 8} {
		for channels := int32(0); channels <= 2; channels++ {
			for end := int32(1); end <= 4; end++ {
				for update := int32(0); update <= 1; update++ {
					if channels == 0 && update != 0 {
						continue
					}
					for trial := 0; trial < 12; trial++ {
						x := make([]float32, max(channels, 1)*32*mult)
						for i := range x {
							x[i] = float32(rng.NormFloat64() * .1)
						}
						if trial == 0 {
							clear(x)
						}
						if trial == 1 {
							for i := range x {
								x[i] = float32(math.Inf(1))
							}
						}
						weights := []int32{1, 2, 3, 4}
						if trial == 2 {
							clear(weights)
						}
						g := []int32{int32(rng.Intn(800)), int32(rng.Intn(50)), int32(rng.Intn(3))}
						c := slices.Clone(g)
						alias := [][3]int32{{0, 1, 2}, {0, 0, 0}, {0, 1, 0}, {0, 1, 1}}[trial%4]
						last := int32(trial % 4)
						result := extensionCollectionResult(func() int32 {
							return opuscc.Opus_spreading_decision(nil, &bands[0], 4, 32, &x[0], &g[alias[0]], last, &g[alias[1]], &g[alias[2]], update, end, channels, mult, &weights[0])
						})
						native := nativeSpreading(bands, 4, 32, x, c, alias[0], alias[1], alias[2], last, update, end, channels, mult, weights)
						if result != native || !slices.Equal(g, c) {
							t.Fatal(mult, channels, end, update, trial, result, native, g, c)
						}
					}
				}
			}
		}
	}
	g, c := []int32{77, 88, 99}, []int32{77, 88, 99}
	got := extensionCollectionResult(func() int32 {
		return opuscc.Opus_spreading_decision(nil, nil, 0, 0, nil, &g[0], 0, &g[1], &g[2], 0, 0, 1, 1, nil)
	})
	native := nativeSpreading(nil, 0, 0, nil, c, 0, 1, 2, 0, 0, 0, 1, 1, nil)
	if got != native || !slices.Equal(g, c) {
		t.Fatal("assert", got, native, g, c)
	}
}

func TestDeinterleaveAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(503))
	for _, stride := range []int32{1, 2, 3, 4, 8, 16} {
		for _, n0 := range []int32{0, 1, 2, 3, 7, 22, 37, 64} {
			for h := int32(0); h <= 1; h++ {
				if h != 0 && (stride == 1 || stride == 3) {
					continue
				}
				n := n0 * stride
				g := make([]float32, n+2)
				for i := range g {
					g[i] = math.Float32frombits(r.Uint32())
				}
				special := []uint32{0, 0x80000000, 0x7f800001, 0x7fc12345, 0xff800001, 0x7f800000, 0xff800000, 1, 0x80000001}
				for i, bits := range special {
					if i < int(n) {
						g[i+1] = math.Float32frombits(bits)
					}
				}
				c := slices.Clone(g)
				opuscc.CompareDeinterleaveHadamard(&g[1], n0, stride, h)
				nativeDeinterleaveHadamard(c[1:], n0, stride, h)
				for i := range g {
					if math.Float32bits(g[i]) != math.Float32bits(c[i]) {
						t.Fatalf("n0=%d stride=%d h=%d i=%d", n0, stride, h, i)
					}
				}
			}
		}
	}
}
