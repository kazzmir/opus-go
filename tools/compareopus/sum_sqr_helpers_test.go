//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"testing"
	"unsafe"
)

func TestSumSqrAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, n := range []int{0, 1, 2, 3, 15, 16, 17, 120, 321, 960, 4096} {
		for trial := 0; trial < 100; trial++ {
			input := make([]int16, n)
			for i := range input {
				switch trial {
				case 0:
					input[i] = 0
				case 1:
					input[i] = -32768
				case 2:
					input[i] = 32767
				default:
					input[i] = int16(rng.Uint32())
				}
			}
			var energy, shift int32
			opuscc.Opus_silk_sum_sqr_shift(nil, &energy, &shift, unsafe.SliceData(input), int32(n))
			ce, cs := nativeSumSqr(input)
			if energy != ce || shift != cs {
				t.Fatalf("n=%d trial=%d Go=(%d,%d) C=(%d,%d)", n, trial, energy, shift, ce, cs)
			}
			if energy < 0 || energy >= (1<<29) {
				t.Fatalf("missing two headroom bits: %d", energy)
			}
		}
	}
}
