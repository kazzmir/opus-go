//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
)

func TestInverseGainAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(3401))
	stable, unstable := 0, 0
	for _, order := range []int{1, 2, 10, 16, 24} {
		for trial := 0; trial < 3000; trial++ {
			a := make([]int16, order)
			for i := range a {
				switch trial % 4 {
				case 0:
					a[i] = int16(rng.Intn(257) - 128)
				case 1:
					a[i] = int16(rng.Intn(8193) - 4096)
				case 2:
					a[i] = int16(rng.Uint32())
				}
			}
			if trial%4 == 3 {
				a[trial%order] = int16(trial * 31)
			}
			before := slices.Clone(a)
			got := opuscc.Opus_silk_LPC_inverse_pred_gain_c(nil, &a[0], int32(order))
			want := nativeInverseGain(a)
			if got != want || !slices.Equal(a, before) {
				t.Fatalf("order=%d trial=%d Go=%d C=%d a=%v", order, trial, got, want, a)
			}
			if got == 0 {
				unstable++
			} else {
				stable++
			}
		}
	}
	if stable == 0 || unstable == 0 {
		t.Fatal("missing stability coverage")
	}
}
