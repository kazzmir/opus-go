//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"math/rand"
	"runtime"
	"testing"
	"unsafe"
)

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
					opuscc.Opus_ec_dec_init(nil, uintptr(unsafe.Pointer(&gd)), uintptr(unsafe.Pointer(unsafe.SliceData(data))), uint32(len(data)))
					g, c := make([]float32, n), make([]float32, n)
					gm := opuscc.Opus_alg_unquant(nil, &g[0], n, k, spread, B, &gd, gain)
					cm, cd := nativeAlgUnquant(data, c, n, k, spread, B, gain)
					gd.Fbuf = 0
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
