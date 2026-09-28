//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"runtime"
	"testing"
	"unsafe"
)

func TestEntropyUintAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	totals := []uint32{2, 3, 6, 13, 255, 256, 257, 258, 511, 512, 513, 65535, 65536, 65537, 1 << 24, 0xffffffff}
	for _, n := range []int{0, 1, 2, 3, 8, 64} {
		for trial := 0; trial < 30; trial++ {
			data := make([]byte, n)
			rng.Read(data)
			var g opuscc.OpusT_ec_dec
			opuscc.Opus_ec_dec_init(nil, uintptr(unsafe.Pointer(&g)), uintptr(unsafe.Pointer(unsafe.SliceData(data))), uint32(n))
			c := g
			for step := 0; step < 200; step++ {
				ft := totals[step%len(totals)]
				if step%3 == 0 {
					ft = rng.Uint32()
					if ft < 2 {
						ft = 2
					}
				}
				gv := opuscc.Opus_ec_dec_uint(nil, &g, ft)
				cv := nativeEntropyStep(&c, data, 3, ft, 0, 0)
				if gv != cv || g != c {
					t.Fatalf("n=%d trial=%d step=%d ft=%d Go=%d C=%d state=%+v C=%+v", n, trial, step, ft, gv, cv, g, c)
				}
			}
			runtime.KeepAlive(data)
		}
	}
	for _, tail := range []uint32{0, 1} {
		g := opuscc.OpusT_ec_dec{Frng: 1 << 31, Fnbits_total: 33, Fnend_bits: 1, Fend_window: tail, Ferror1: 7}
		c := g
		gv := opuscc.Opus_ec_dec_uint(nil, &g, 257)
		cv := nativeEntropyStep(&c, nil, 3, 257, 0, 0)
		if gv != cv || g != c {
			t.Fatalf("clamp/error handling differs: tail=%d Go=%+v C=%+v", tail, g, c)
		}
	}
}
