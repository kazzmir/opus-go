//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math"
	"math/rand"
	"slices"
	"testing"
)

func TestSoftClipAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(704))
	for _, channels := range []int32{1, 2, 3, 8} {
		gm, cm := make([]float32, channels+2), make([]float32, channels+2)
		gm[0] = 123
		cm[0] = 123
		gm[channels+1] = 456
		cm[channels+1] = 456
		for pattern := 0; pattern < 8; pattern++ {
			for _, n := range []int32{0, 1, 2, 3, 7, 80, 480} {
				g := make([]float32, n*channels+2)
				g[0] = 77
				g[len(g)-1] = 88
				for i := int32(0); i < n; i++ {
					for ch := int32(0); ch < channels; ch++ {
						v := (r.Float32() - 0.5) * 6
						switch pattern {
						case 1:
							v = 0.1 + float32(i)*2/float32(n)
						case 2:
							v = -0.1 - float32(i)*2/float32(n)
						case 3:
							v = (r.Float32() - 0.5) * 1.5
						case 4:
							if i%2 == 0 {
								v = 2
							} else {
								v = -2
							}
						case 5:
							v = math.Float32frombits([]uint32{0, 0x80000000, 0x7fc12345, 0x7f800000, 0xff800000}[int(i)%5])
						case 6:
							v = 0
							if i == n/2 {
								v = 3
							}
						case 7:
							v = 1.25
						}
						g[1+i*channels+ch] = v
					}
				}
				c := slices.Clone(g)
				opuscc.Opus_opus_pcm_soft_clip(nil, &g[1], n, channels, &gm[1])
				nativeSoftClip(c[1:], cm[1:], n, channels)
				for i := range g {
					if math.Float32bits(g[i]) != math.Float32bits(c[i]) {
						t.Fatalf("C=%d pattern=%d N=%d i=%d Go=%g C=%g bits=%x/%x", channels, pattern, n, i, g[i], c[i], math.Float32bits(g[i]), math.Float32bits(c[i]))
					}
				}
				for i := range gm {
					if math.Float32bits(gm[i]) != math.Float32bits(cm[i]) {
						t.Fatalf("state C=%d pattern=%d N=%d i=%d", channels, pattern, n, i)
					}
				}
			}
		}
	}
}
