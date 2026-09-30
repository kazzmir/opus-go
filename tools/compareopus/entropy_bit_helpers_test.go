//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"runtime"
	"testing"
	"unsafe"
)

func TestTFDecodeAgainstC(t *testing.T) {
	for _, n := range []int{0, 1, 2, 8, 32} {
		for LM := int32(0); LM < 4; LM++ {
			for transient := int32(0); transient < 2; transient++ {
				for _, start := range []int32{0, 2} {
					data := make([]byte, n)
					for i := range data {
						data[i] = byte(71*i + 123)
					}
					g := opuscc.OpusT_ec_dec{}
					opuscc.Opus_ec_dec_init(nil, &g, unsafe.SliceData(data), uint32(n))
					c := g
					goOut := [23]int32{}
					for i := range goOut {
						goOut[i] = 77
					}
					cOut := goOut
					opuscc.CompareTFDecode(start, 21, transient, &goOut[0], LM, &g)
					nativeTFDecode(&c, data, start, 21, transient, cOut[:], LM)
					if goOut != cOut || g != c {
						t.Fatal(n, LM, transient, start, goOut, cOut, g, c)
					}
				}
			}
		}
	}
}

func TestEntropyBitAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, n := range []int{0, 1, 2, 3, 8, 64} {
		for trial := 0; trial < 30; trial++ {
			data := make([]byte, n)
			rng.Read(data)
			if trial == 0 {
				clear(data)
			}
			if trial == 1 {
				for i := range data {
					data[i] = 255
				}
			}
			var g opuscc.OpusT_ec_dec
			opuscc.Opus_ec_dec_init(nil, &g, unsafe.SliceData(data), uint32(n))
			c := g
			for step := 0; step < 200; step++ {
				logp := uint32(1 + rng.Intn(15))
				gv := opuscc.Opus_ec_dec_bit_logp(nil, &g, logp)
				cv := nativeEntropyStep(&c, data, 1, logp, 0, 0)
				if uint32(gv) != cv || g != c {
					t.Fatalf("n=%d trial=%d step=%d logp=%d Go=%d C=%d state=%+v C=%+v", n, trial, step, logp, gv, cv, g, c)
				}
			}
			runtime.KeepAlive(data)
		}
	}
}
