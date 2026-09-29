//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
)

func TestCWRSIndexAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(1401))
	for _, pair := range pvqTestPairs {
		n, k := pair[0], pair[1]
		for trial := 0; trial < 200; trial++ {
			y := make([]int32, n)
			for p := int32(0); p < k; p++ {
				pos := r.Intn(int(n))
				sign := int32(1)
				if y[pos] < 0 || y[pos] == 0 && r.Intn(2) == 0 {
					sign = -1
				}
				y[pos] += sign
			}
			before := slices.Clone(y)
			got := opuscc.CompareCWRSIndex(n, &y[0])
			want := nativeCWRSIndex(n, y)
			if got != want || !slices.Equal(y, before) || got >= nativePVQCount(n, k) {
				t.Fatal(n, k, got, want, y)
			}
		}
	}
}
