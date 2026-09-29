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

func TestEntropyInitAgainstC(t *testing.T) {
	r := rand.New(rand.NewSource(804))
	for _, n := range []int{0, 1, 2, 3, 4, 5, 7, 32, 128} {
		for trial := 0; trial < 256; trial++ {
			data := make([]byte, n)
			r.Read(data)
			if n > 0 {
				data[0] = byte(trial)
			}
			original := slices.Clone(data)
			g := opuscc.OpusT_ec_dec{Fstorage: 0xa5a5a5a5, Fend_offs: 0xa5a5a5a5, Fend_window: 0xa5a5a5a5, Fnend_bits: -1515870811, Fnbits_total: -1515870811, Foffs: 0xa5a5a5a5, Frng: 0xa5a5a5a5, Fval: 0xa5a5a5a5, Fext: uint32(r.Uint64()), Frem: -1515870811, Ferror1: -1515870811}
			c := g
			nativeEntropyStep(&c, data, 9, uint32(n), 0, 0)
			opuscc.Opus_ec_dec_init(nil, &g, unsafe.SliceData(data), uint32(n))
			c.Fbuf = unsafe.SliceData(data)
			if g != c || !slices.Equal(data, original) {
				t.Fatalf("n=%d trial=%d\nGo=%+v\nC=%+v", n, trial, g, c)
			}
			runtime.KeepAlive(data)
		}
	}
}
