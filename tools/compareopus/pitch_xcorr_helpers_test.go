//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"math/rand"
	"testing"
	"unsafe"
)

func TestPitchXCorrAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(501))
	for _, n := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 17, 64, 127, 480} {
		for lags := 1; lags <= 33; lags++ {
			if lags > 3 && n < 3 {
				continue
			}
			x := make([]float32, n)
			y := make([]float32, 0)
			if n > 0 {
				y = make([]float32, n+lags-1)
			}
			for trial := 0; trial < 12; trial++ {
				for i := range x {
					x[i] = (r.Float32() - 0.5) * float32(uint32(1)<<uint(trial))
				}
				for i := range y {
					y[i] = (r.Float32() - 0.5) * float32(uint32(1)<<uint(trial))
				}
				g, c := make([]float32, lags+2), make([]float32, lags+2)
				g[0] = 123
				c[0] = 123
				g[lags+1] = 456
				c[lags+1] = 456
				opuscc.Opus_celt_pitch_xcorr_c(nil, unsafe.SliceData(x), unsafe.SliceData(y), &g[1], int32(n), int32(lags), 0)
				nativeScalarPitchXCorr(x, y, c[1:], int32(n), int32(lags))
				for i := range g {
					if math.Float32bits(g[i]) != math.Float32bits(c[i]) {
						t.Fatalf("n=%d lags=%d trial=%d i=%d Go=%g C=%g", n, lags, trial, i, g[i], c[i])
					}
				}
			}
		}
	}
}
