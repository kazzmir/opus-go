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

func TestICDF16AgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, n := range []int{0, 1, 3, 16, 64} {
		for _, ftb := range []uint32{1, 8, 15, 16} {
			for trial := 0; trial < 30; trial++ {
				total := 1 << ftb
				table := []uint16{uint16(total - 1), uint16(total / 2), 0}
				if ftb == 1 {
					table = []uint16{1, 0}
				}
				if trial%2 == 1 && ftb > 1 {
					table = []uint16{uint16(total - 1), uint16(total/2 + 1), 1, 0}
				}
				original := slices.Clone(table)
				data := make([]byte, n)
				rng.Read(data)
				var g opuscc.OpusT_ec_dec
				opuscc.Opus_ec_dec_init(nil, &g, unsafe.SliceData(data), uint32(n))
				c := g
				for step := 0; step < 100; step++ {
					gv := opuscc.Opus_ec_dec_icdf16(nil, &g, &table[0], ftb)
					cv := nativeEntropyStep(&c, data, 4, ftb, 0, 0, table)
					if uint32(gv) != cv || g != c {
						t.Fatalf("n=%d ftb=%d trial=%d step=%d Go=%d C=%d state=%+v C=%+v", n, ftb, trial, step, gv, cv, g, c)
					}
				}
				if !slices.Equal(table, original) {
					t.Fatal("table changed")
				}
				runtime.KeepAlive(data)
			}
		}
	}
}
