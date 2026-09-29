//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"runtime"
	"slices"
	"testing"
	"unsafe"
)

func TestEntropyBufferAgainstC(t *testing.T) {
	for _, size := range []uint32{32, 48, 95, 96} {
		g := make([]byte, 96)
		for i := range g {
			g[i] = 0xa5
		}
		c := slices.Clone(g)
		var e opuscc.OpusT_ec_enc
		// Encoder entry points still take uintptr; keep their context and input stable.
		var pins runtime.Pinner
		pins.Pin(&e)
		pins.Pin(&g[0])
		opuscc.Opus_ec_enc_init(nil, uintptr(unsafe.Pointer(&e)), uintptr(unsafe.Pointer(&g[0])), uint32(len(g)))
		for i := uint32(0); i < 18; i++ {
			opuscc.Opus_ec_enc_uint(nil, uintptr(unsafe.Pointer(&e)), i%17, 17)
		}
		opuscc.Opus_ec_enc_bits(nil, uintptr(unsafe.Pointer(&e)), 0xa5b, 12)
		for i := uint32(0); i < 3; i++ {
			opuscc.Opus_ec_enc_bits(nil, uintptr(unsafe.Pointer(&e)), 0x1234567+i, 25)
		}
		opuscc.Opus_ec_enc_shrink(nil, uintptr(unsafe.Pointer(&e)), size)
		opuscc.Opus_ec_enc_done(nil, uintptr(unsafe.Pointer(&e)))
		want := nativeEntropyBuffer(c, size)
		got := [11]uint32{e.Fstorage, e.Fend_offs, e.Fend_window, uint32(e.Fnend_bits), uint32(e.Fnbits_total), e.Foffs, e.Frng, e.Fval, e.Fext, uint32(e.Frem), uint32(e.Ferror1)}
		pins.Unpin()
		if got != want || !slices.Equal(g, c) {
			t.Fatalf("size=%d state=%v PCM=%v", size, got != want, !slices.Equal(g, c))
		}
	}
}
