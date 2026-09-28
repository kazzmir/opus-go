//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"math/rand"
	"runtime"
	"testing"
	"unsafe"
)

func TestShellDecodeAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(2026))
	for _, n := range []int{0, 1, 3, 16, 64} {
		for total := int32(0); total <= 16; total++ {
			for trial := 0; trial < 100; trial++ {
				data := make([]byte, n)
				rng.Read(data)
				var state opuscc.OpusT_ec_dec
				var out [16]int16
				opuscc.Opus_ec_dec_init(nil, uintptr(unsafe.Pointer(&state)), uintptr(unsafe.Pointer(unsafe.SliceData(data))), uint32(n))
				opuscc.Opus_silk_shell_decoder(nil, &out, &state, total)
				c, cs := nativeShellDecode(data, total)
				state.Fbuf = 0 // The bridge compares every numeric field, not buffer addresses.
				if out != c || state != cs {
					t.Fatalf("n=%d total=%d trial=%d Go=%v C=%v state=%+v C=%+v", n, total, trial, out, c, state, cs)
				}
				runtime.KeepAlive(data)
			}
		}
	}
}
