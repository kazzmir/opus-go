//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"slices"
	"testing"
	"unsafe"
)

func TestEntropyEncInitAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(1002))
	for _, n := range []int{0, 1, 2, 7, 32, 256} {
		for trial := 0; trial < 32; trial++ {
			buf := make([]byte, n)
			r.Read(buf)
			original := slices.Clone(buf)
			g := opuscc.OpusT_ec_enc{Fstorage: r.Uint32(), Fend_offs: r.Uint32(), Fend_window: r.Uint32(), Fnend_bits: int32(r.Uint32()), Fnbits_total: int32(r.Uint32()), Foffs: r.Uint32(), Frng: r.Uint32(), Fval: r.Uint32(), Fext: r.Uint32(), Frem: int32(r.Uint32()), Ferror1: -1}
			c := g
			nativeEncoderStep(&c, buf, 0, uint32(n), 0)
			opuscc.Opus_ec_enc_init(nil, &g, unsafe.SliceData(buf), uint32(n))
			if g != c || !slices.Equal(buf, original) {
				t.Fatal(n, trial, g, c)
			}
		}
	}
}
