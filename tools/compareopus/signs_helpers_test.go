//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"runtime"
	"slices"
	"testing"
	"unsafe"
)

func TestSignsAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(3404))
	for signal := int32(0); signal <= 2; signal++ {
		for offset := int32(0); offset <= 1; offset++ {
			for _, length := range []int32{0, 7, 8, 16, 24, 120, 320} {
				for trial := 0; trial < 64; trial++ {
					blocks := (length + 8) >> 4
					n := int(blocks * 16)
					sums := make([]int32, blocks)
					for i := range sums {
						sums[i] = int32(trial)
						if i%3 == 1 {
							sums[i] = 0
						}
						if i%3 == 2 {
							sums[i] = int32(trial + 32)
						}
					}
					sumBefore := slices.Clone(sums)
					g := make([]int16, n+2)
					g[0] = 123
					g[n+1] = 456
					for i := 1; i <= n; i++ {
						g[i] = int16(rng.Intn(32768))
						if i%4 == 0 {
							g[i] = 0
						}
						if i%13 == 0 {
							g[i] = -32768
						}
					}
					c := slices.Clone(g)
					data := make([]byte, []int{0, 1, 3, 64}[trial%4])
					rng.Read(data)
					var gd opuscc.OpusT_ec_dec
					opuscc.Opus_ec_dec_init(nil, &gd, unsafe.SliceData(data), uint32(len(data)))
					opuscc.Opus_silk_decode_signs(nil, &gd, &g[1], length, signal, offset, unsafe.SliceData(sums))
					cd := nativeSignsDecode(data, c[1:n+1], length, signal, offset, sums)
					gd.Fbuf = nil
					if !slices.Equal(g, c) || gd != cd || !slices.Equal(sums, sumBefore) {
						t.Fatalf("signal=%d offset=%d length=%d trial=%d Go=%v C=%v state=%+v C=%+v", signal, offset, length, trial, g, c, gd, cd)
					}
					if g[0] != 123 || g[n+1] != 456 {
						t.Fatal("pulse guard changed")
					}
					runtime.KeepAlive(data)
				}
			}
		}
	}
}
