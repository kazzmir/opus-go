//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"runtime"
	"slices"
	"testing"
)

var pvqTestPairs = [][2]int32{{2, 1}, {2, 128}, {3, 128}, {4, 32}, {5, 12}, {8, 8}, {16, 4}, {32, 3}, {64, 2}, {128, 1}, {176, 1}}

func TestCWRSAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(3703))
	for _, pair := range pvqTestPairs {
		n, k := pair[0], pair[1]
		count := nativePVQCount(n, k)
		trials := min(count, uint32(1024))
		for trial := uint32(0); trial < trials+3; trial++ {
			index := trial
			if count > 1024 {
				index = rng.Uint32() % count
			}
			if trial >= trials {
				index = []uint32{0, count / 2, count - 1}[trial-trials]
			}
			data := make([]byte, 64)
			if err := nativePVQIndex(data, n, k, index); err != 0 {
				t.Fatalf("C encoder error=%d", err)
			}
			var gd opuscc.OpusT_ec_dec
			opuscc.Opus_ec_dec_init(nil, &gd, &data[0], uint32(len(data)))
			g, c := make([]int32, n), make([]int32, n)
			ge := opuscc.Opus_decode_pulses(nil, &g[0], n, k, &gd)
			ce, cd := nativePVQDecode(data, c, n, k)
			gd.Fbuf = nil
			if ge != ce || !slices.Equal(g, c) || gd != cd {
				t.Fatalf("n=%d k=%d index=%d Go=%v C=%v energy=%g/%g state=%+v/%+v", n, k, index, g, c, ge, ce, gd, cd)
			}
			runtime.KeepAlive(data)
		}
	}
}
