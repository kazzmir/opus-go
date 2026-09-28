//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"runtime"
	"testing"
	"unsafe"
)

func TestLaplaceDecodeAgainstC(t *testing.T) {
	for _, fs := range []uint32{1, 6000, 16000, 32735} {
		for _, decay := range []uint32{1, 4096, 11456} {
			for fm := uint32(0); fm < 32768; fm++ {
				g := opuscc.OpusT_ec_dec{Frng: 1 << 31, Fval: (32767 - fm) << 16, Fnbits_total: 33}
				c := g
				gv := opuscc.Opus_ec_laplace_decode(nil, &g, fs, int32(decay))
				cv := int32(nativeEntropyStep(&c, nil, 5, fs, decay, 0))
				if gv != cv || g != c {
					t.Fatalf("fs=%d decay=%d fm=%d Go=%d C=%d state=%+v C=%+v", fs, decay, fm, gv, cv, g, c)
				}
			}
		}
	}
	rng := rand.New(rand.NewSource(2026))
	for _, n := range []int{0, 1, 3, 64} {
		data := make([]byte, n)
		rng.Read(data)
		var g opuscc.OpusT_ec_dec
		opuscc.Opus_ec_dec_init(nil, uintptr(unsafe.Pointer(&g)), uintptr(unsafe.Pointer(unsafe.SliceData(data))), uint32(n))
		c := g
		for step := 0; step < 1000; step++ {
			fs, decay := uint32(1+rng.Intn(32735)), uint32(1+rng.Intn(11456))
			gv := opuscc.Opus_ec_laplace_decode(nil, &g, fs, int32(decay))
			cv := int32(nativeEntropyStep(&c, data, 5, fs, decay, 0))
			if gv != cv || g != c {
				t.Fatalf("n=%d step=%d Go=%d C=%d", n, step, gv, cv)
			}
		}
		runtime.KeepAlive(data)
	}
}
