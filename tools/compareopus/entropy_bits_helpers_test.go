//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"runtime"
	"testing"
	"unsafe"
)

func TestEntropyBitsAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, n := range []int{0, 1, 2, 3, 4, 8, 64} {
		for trial := 0; trial < 30; trial++ {
			data := make([]byte, n)
			rng.Read(data)
			var g opuscc.OpusT_ec_dec
			opuscc.Opus_ec_dec_init(nil, uintptr(unsafe.Pointer(&g)), uintptr(unsafe.Pointer(unsafe.SliceData(data))), uint32(n))
			c := g
			for step := 0; step < 100; step++ {
				width := uint32(step % 26)
				gv := opuscc.Opus_ec_dec_bits(nil, &g, width)
				cv := nativeEntropyStep(&c, data, 2, width, 0, 0)
				if gv != cv || g != c {
					t.Fatalf("n=%d trial=%d step=%d width=%d Go=%d C=%d state=%+v C=%+v", n, trial, step, width, gv, cv, g, c)
				}
				// Front and tail reads may overlap; neither stops at the other's offset.
				if step%3 == 0 {
					gb := opuscc.Opus_ec_dec_bit_logp(nil, &g, 1)
					cb := nativeEntropyStep(&c, data, 1, 1, 0, 0)
					if uint32(gb) != cb || g != c {
						t.Fatal("interleaved range decoding differs")
					}
				}
			}
			runtime.KeepAlive(data)
		}
	}
}
