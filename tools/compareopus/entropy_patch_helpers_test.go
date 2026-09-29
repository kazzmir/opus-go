//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"slices"
	"testing"
)

func TestEntropyPatchAgainstC(t *testing.T) {
	for nbits := uint32(0); nbits <= 8; nbits++ {
		threshold := uint32(1) << 31 >> nbits
		for _, mode := range []int{0, 1, 2, 3} {
			for value := uint32(0); value < 256; value++ {
				for _, rng := range []uint32{threshold - 1, threshold, threshold + 1} {
					g := []byte{77, byte(value ^ 0xa5), 88}
					c := slices.Clone(g)
					ge := opuscc.OpusT_ec_enc{Fbuf: &g[1], Fstorage: 1, Fend_window: 0xabcdef, Fnend_bits: 7, Fnbits_total: 39, Frng: rng, Fval: 0xa5b6c7d8, Fext: 3, Frem: -1, Ferror1: 9}
					switch mode {
					case 0:
						ge.Foffs = 1
					case 1:
						ge.Frem = int32(value ^ 0x55)
					case 2:
						ge.Frem = 0x7fffffff
					case 3:
						ge.Frem = -2
					}
					ce := ge
					opuscc.Opus_ec_enc_patch_initial_bits(nil, &ge, value, nbits)
					nativeEncoderStep(&ce, c[1:], 2, value, nbits)
					ce.Fbuf = ge.Fbuf
					if ge != ce || !slices.Equal(g, c) {
						t.Fatal(nbits, mode, value, rng, ge, ce, g, c)
					}
				}
			}
		}
	}
}
