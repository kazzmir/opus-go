//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
	"unsafe"
)

func TestStereoIThetaAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(502))
	for _, n := range []int{0, 1, 2, 3, 4, 8, 16, 32, 64, 128, 176} {
		for _, stereo := range []int32{0, 1} {
			for _, scale := range []float32{0, 1e-20, 1e-10, 1e-9, 0.125, 1, 32768, 1e15} {
				for trial := 0; trial < 24; trial++ {
					x, y := make([]float32, n), make([]float32, n)
					for i := range x {
						x[i] = (r.Float32() - 0.5) * scale
						y[i] = (r.Float32() - 0.5) * scale
						if trial == 0 {
							y[i] = x[i]
						}
						if trial == 1 {
							y[i] = -x[i]
						}
						if trial == 2 {
							x[i] = 0
						}
						if trial == 3 {
							y[i] = 0
						}
					}
					oldX, oldY := slices.Clone(x), slices.Clone(y)
					got := opuscc.Opus_stereo_itheta(nil, unsafe.SliceData(x), unsafe.SliceData(y), stereo, int32(n), 0)
					want := nativeScalarStereoITheta(x, y, stereo)
					if got != want || !slices.Equal(x, oldX) || !slices.Equal(y, oldY) {
						t.Fatalf("n=%d stereo=%d scale=%g trial=%d Go=%d C=%d", n, stereo, scale, trial, got, want)
					}
				}
			}
		}
	}
}
