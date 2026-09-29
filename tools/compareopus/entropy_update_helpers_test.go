//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"runtime"
	"testing"
	"unsafe"
)

func TestEntropyUpdateAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, n := range []int{0, 1, 2, 3, 8, 64} {
		for trial := 0; trial < 30; trial++ {
			data := make([]byte, n)
			rng.Read(data)
			var g opuscc.OpusT_ec_dec
			opuscc.Opus_ec_dec_init(nil, &g, unsafe.SliceData(data), uint32(n))
			for step := 0; step < 100; step++ {
				ft := []uint32{2, 13, 256, 32768}[step%4]
				symbol := opuscc.Opus_ec_decode(nil, &g, ft)
				fl, fh := symbol, symbol+1
				if step%3 == 0 {
					fl = 0
				}
				if step%5 == 0 {
					fh = ft
				}
				c := g
				opuscc.Opus_ec_dec_update(nil, &g, fl, fh, ft)
				nativeEntropyStep(&c, data, 0, fl, fh, ft)
				if g != c {
					t.Fatalf("n=%d trial=%d step=%d Go=%+v C=%+v", n, trial, step, g, c)
				}
			}
			// Also retain the fixture owner across the native comparison.
			runtime.KeepAlive(data)
		}
	}
}
