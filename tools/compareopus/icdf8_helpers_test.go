//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
	"unsafe"
)

func TestICDF8AgainstC(t *testing.T) {
	for _, table := range [][]byte{{0}, {255, 1, 0}, {240, 180, 100, 0}, {3, 2, 1, 0}} {
		ftb := uint32(8)
		if table[0] < 4 {
			ftb = 2
		}
		// Test both sides of every decision boundary with normalized representative ranges.
		for _, rng := range []uint32{1<<23 + 1, 1 << 24, 0x12345678, 1 << 31} {
			values := []uint32{0, rng - 1}
			for _, v := range table {
				threshold := (rng >> ftb) * uint32(v)
				values = append(values, threshold)
				if threshold > 0 {
					values = append(values, threshold-1)
				}
			}
			for _, value := range values {
				if value >= rng {
					continue
				}
				for _, storage := range []uint32{0, 1, 8} {
					data := []byte{0, 255, 128, 64, 17, 239, 123, 34}
					original := slices.Clone(data)
					originalTable := slices.Clone(table)
					ge := opuscc.OpusT_ec_dec{Fbuf: &data[0], Fstorage: storage, Frng: rng, Fval: value, Fnbits_total: 33, Frem: 23, Fext: 77}
					ce := ge
					got := opuscc.Opus_ec_dec_icdf(nil, &ge, &table[0], ftb)
					want, _ := nativeEntropyStepPointer(&ce, data, 10, ftb, 0, 0, unsafe.Pointer(&table[0]))
					if got != int32(want) || ge != ce || !slices.Equal(data, original) || !slices.Equal(table, originalTable) {
						t.Fatal(table, ftb, rng, value, storage, got, want, ge, ce)
					}
				}
			}
		}
	}
}
