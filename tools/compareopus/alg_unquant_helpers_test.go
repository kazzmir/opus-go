//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"math/rand"
	"runtime"
	"slices"
	"testing"
	"unsafe"
)

func TestAlgQuantAgainstC(t *testing.T) {
	for _, pair := range pvqTestPairs {
		n, k := pair[0], pair[1]
		for _, B := range []int32{1, 2, 4, 8} {
			if n%B != 0 {
				continue
			}
			for spread := int32(0); spread < 4; spread++ {
				for trial := 0; trial < 12; trial++ {
					g := make([]float32, n+2)
					g[0] = 77
					g[len(g)-1] = 88
					for i := int32(1); i <= n; i++ {
						g[i] = float32(math.Sin(float64(i*17+int32(trial)) * .31))
					}
					if trial == 0 {
						clear(g[1 : len(g)-1])
					}
					c := slices.Clone(g)
					size := []int{0, 1, 8, 64}[trial%4]
					gb := make([]byte, size+2)
					for i := range gb {
						gb[i] = 0xa5
					}
					cb := slices.Clone(gb)
					gdata, cdata := gb[1:len(gb)-1], cb[1:len(cb)-1]
					var ge, ce opuscc.OpusT_ec_enc
					opuscc.Opus_ec_enc_init(nil, &ge, unsafe.SliceData(gdata), uint32(size))
					nativeEncoderStep(&ce, cdata, 0, uint32(size), 0)
					gain := []float32{0, .25, 1, 1.75}[trial%4]
					resynth := []int32{0, 1, -1}[trial%3]
					gm := opuscc.Opus_alg_quant(nil, &g[1], n, k, spread, B, &ge, gain, resynth, 0)
					cm := nativeAlgQuant(&ce, cdata, c[1:], n, k, spread, B, gain, resynth)
					gs, cs := ge, ce
					gs.Fbuf = nil
					cs.Fbuf = nil
					if gm != cm || gs != cs || !slices.Equal(gb, cb) || !sameFloatBits(g, c) {
						t.Fatal(n, k, B, spread, trial, "quantization", gm, cm, gs, cs)
					}
					opuscc.Opus_ec_enc_done(nil, &ge)
					nativeEncoderStep(&ce, cdata, 10, 0, 0)
					gs, cs = ge, ce
					gs.Fbuf = nil
					cs.Fbuf = nil
					if gs != cs || !slices.Equal(gb, cb) {
						t.Fatal(n, k, B, spread, trial, "finalized bytes/state")
					}
				}
			}
		}
	}
}

func TestPVQSearchAgainstC(t *testing.T) {
	for _, n := range []int32{2, 3, 4, 5, 8, 16, 32, 64, 128, 176} {
		for _, k := range []int32{0, 1, 2, 3, 8, 16, 32, 128} {
			for trial := 0; trial < 20; trial++ {
				g := make([]float32, n+2)
				g[0] = 77
				g[len(g)-1] = 88
				for j := int32(1); j <= n; j++ {
					g[j] = float32(math.Sin(float64(j*17+int32(trial)) * 0.31))
				}
				if trial == 0 {
					clear(g[1 : len(g)-1])
				}
				if trial == 1 {
					for j := 1; j < len(g)-1; j++ {
						g[j] = 1e-20
					}
				}
				if trial == 2 {
					g[1] = float32(math.Inf(1))
				}
				if trial == 3 {
					g[1] = float32(math.NaN())
				}
				if trial == 4 {
					for j := 1; j < len(g)-1; j++ {
						g[j] = math.Float32frombits(uint32(j%2) << 31)
					}
				}
				c := append([]float32(nil), g...)
				gp := make([]int32, n+2)
				gp[0] = 77
				gp[len(gp)-1] = 88
				cp := append([]int32(nil), gp...)
				gy := opuscc.Opus_op_pvq_search_c(nil, &g[1], &gp[1], k, n, 0)
				cy := nativePVQSearch(c[1:], cp[1:], k, n)
				if math.Float32bits(gy) != math.Float32bits(cy) || !sameFloatBits(g, c) || !slices.Equal(gp, cp) {
					t.Fatal(n, k, trial, gy, cy, gp, cp)
				}
			}
		}
	}
}

func TestAlgUnquantAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(3704))
	for _, pair := range pvqTestPairs {
		n, k := pair[0], pair[1]
		for _, B := range []int32{1, 2, 4, 8} {
			if n%B != 0 {
				continue
			}
			for spread := int32(0); spread <= 3; spread++ {
				for trial := 0; trial < 24; trial++ {
					data := make([]byte, []int{0, 1, 64}[trial%3])
					rng.Read(data)
					gain := []float32{0, 0.25, 1, 1.75}[trial%4]
					var gd opuscc.OpusT_ec_dec
					opuscc.Opus_ec_dec_init(nil, &gd, unsafe.SliceData(data), uint32(len(data)))
					g, c := make([]float32, n), make([]float32, n)
					gm := opuscc.Opus_alg_unquant(nil, &g[0], n, k, spread, B, &gd, gain)
					cm, cd := nativeAlgUnquant(data, c, n, k, spread, B, gain)
					gd.Fbuf = nil
					if gm != cm || gd != cd {
						t.Fatalf("n=%d k=%d B=%d spread=%d trial=%d mask=%x/%x state=%+v/%+v", n, k, B, spread, trial, gm, cm, gd, cd)
					}
					for i := range g {
						if math.Float32bits(g[i]) != math.Float32bits(c[i]) {
							t.Fatalf("n=%d k=%d B=%d spread=%d trial=%d i=%d Go=%g C=%g", n, k, B, spread, trial, i, g[i], c[i])
						}
					}
					runtime.KeepAlive(data)
				}
			}
		}
	}
}
