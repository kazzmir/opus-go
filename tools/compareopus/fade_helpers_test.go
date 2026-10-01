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

func TestCombFilterAgainstC(t *testing.T) {
	for _, n := range []int32{0, 1, 4, 9, 32, 120} {
		for _, overlap := range []int32{0, 1, 4, 8} {
			if overlap > n {
				continue
			}
			for tap0 := int32(0); tap0 < 3; tap0++ {
				for tap1 := int32(0); tap1 < 3; tap1++ {
					for _, gain := range [][2]float32{{0, 0}, {.25, 0}, {0, .5}, {.25, .5}, {.25, .25}, {-.5, -.25}} {
						for _, offset := range []int{0, 2, -2, 140} {
							g := make([]float32, 400)
							for i := range g {
								g[i] = float32(math.Sin(float64(i) * .31))
							}
							c := append([]float32(nil), g...)
							win := make([]float32, overlap)
							for i := range win {
								win[i] = float32(i+1) / float32(overlap+1)
							}
							opuscc.Opus_comb_filter(nil, &g[80+offset], &g[80], 0, 21, n, gain[0], gain[1], tap0, tap1, unsafe.SliceData(win), overlap, 0)
							nativeComb(&c[80+offset], &c[80], 0, 21, n, gain[0], gain[1], tap0, tap1, unsafe.SliceData(win), overlap)
							if !sameFloatBits(g, c) {
								t.Fatal(n, overlap, tap0, tap1, gain, offset)
							}
						}
					}
				}
			}
		}
	}
	// An unchanged filter ignores its window and requested overlap.
	g := make([]float32, 96)
	for i := range g {
		g[i] = float32(i) * .013
	}
	c := append([]float32(nil), g...)
	opuscc.Opus_comb_filter(nil, &g[32], &g[32], 15, 15, 32, .25, .25, 2, 2, nil, 20, 0)
	nativeComb(&c[32], &c[32], 15, 15, 32, .25, .25, 2, 2, nil, 20)
	if !sameFloatBits(g, c) {
		t.Fatal("unchanged filter")
	}
}

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
