//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
)

func TestPLCEnergyAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(3504))
	for _, n := range []int32{1, 19, 20, 40, 60, 80} {
		for _, nb := range []int32{2, 4} {
			for trial := 0; trial < 500; trial++ {
				exc := make([]int32, n*nb)
				gains := [2]int32{1024, 2048}
				for i := range exc {
					switch trial % 4 {
					case 0:
						exc[i] = int32(rng.Uint32())
					case 1:
						exc[i] = int32(rng.Intn(1<<24)) - (1 << 23)
					case 2:
						exc[i] = 1 << 28
					}
				}
				if trial%4 == 0 {
					gains = [2]int32{int32(rng.Uint32()), int32(rng.Uint32())}
				}
				if trial%4 == 2 {
					gains = [2]int32{1 << 20, 1 << 20}
				}
				before := slices.Clone(exc)
				gb := gains
				got := opuscc.ComparePLCEnergy(exc, gains, n, nb)
				want := nativePLCEnergy(exc, &gains, n, nb)
				if got != want || !slices.Equal(exc, before) || gains != gb {
					t.Fatalf("n=%d nb=%d trial=%d gains=%v Go=%v C=%v", n, nb, trial, gains, got, want)
				}
			}
		}
	}
}
