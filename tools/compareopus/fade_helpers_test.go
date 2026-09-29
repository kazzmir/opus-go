//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"math/rand"
	"slices"
	"testing"
	"unsafe"
)

func TestFadeAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(701))
	for _, rate := range []int32{8000, 12000, 16000, 24000, 48000} {
		for _, channels := range []int32{1, 2, 4} {
			for _, n := range []int32{1, 3, rate / 400} {
				for _, alias := range []int{0, 1, 2, 3} {
					count := n * channels
					w := make([]float32, (n-1)*(48000/rate)+1)
					for i := range w {
						w[i] = r.Float32()
					}
					w[0] = 0
					w[len(w)-1] = 1
					a, b := make([]float32, count+2), make([]float32, count+2)
					for i := range a {
						a[i] = (r.Float32() - 0.5) * 4
						b[i] = (r.Float32() - 0.5) * 4
					}
					ca, cb := slices.Clone(a), slices.Clone(b)
					g, c := make([]float32, count+2), make([]float32, count+2)
					switch alias {
					case 1:
						g = a
						c = ca
					case 2:
						g = b
						c = cb
					case 3:
						g = a[1:]
						c = ca[1:]
					}
					opuscc.CompareSmoothFade(&a[0], &b[0], &g[0], n, channels, unsafe.SliceData(w), rate)
					nativeFade(ca, cb, c, w, n, channels, rate)
					for i := range g {
						if math.Float32bits(g[i]) != math.Float32bits(c[i]) {
							t.Fatalf("rate=%d C=%d N=%d alias=%d i=%d", rate, channels, n, alias, i)
						}
					}
				}
			}
		}
	}
}
