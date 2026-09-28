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

func TestPulseDecodeAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(3503))
	compare := func(data []byte, length, signal, offset int32) []int16 {
		t.Helper()
		n := int((length+15)/16) * 16
		g := make([]int16, n+2)
		for i := range g {
			g[i] = 123
		}
		c := slices.Clone(g)
		var gd opuscc.OpusT_ec_dec
		opuscc.Opus_ec_dec_init(nil, uintptr(unsafe.Pointer(&gd)), uintptr(unsafe.Pointer(unsafe.SliceData(data))), uint32(len(data)))
		opuscc.Opus_silk_decode_pulses(nil, &gd, &g[1], signal, offset, length)
		cd := nativePulseDecode(data, c[1:n+1], length, signal, offset)
		gd.Fbuf = 0
		if !slices.Equal(g, c) || gd != cd {
			t.Fatalf("length=%d signal=%d offset=%d bytes=%d Go=%v C=%v state=%+v C=%+v", length, signal, offset, len(data), g, c, gd, cd)
		}
		if g[0] != 123 || g[n+1] != 123 {
			t.Fatal("pulse guard changed")
		}
		runtime.KeepAlive(data)
		return g[1 : n+1]
	}
	for signal := int32(0); signal <= 2; signal++ {
		for offset := int32(0); offset <= 1; offset++ {
			for _, length := range []int32{16, 80, 120, 160, 240, 320} {
				for trial := 0; trial < 50; trial++ {
					data := make([]byte, []int{0, 1, 3, 64, 128}[trial%5])
					rng.Read(data)
					if trial == 4 {
						for i := range data {
							data[i] = 255
						}
					}
					compare(data, length, signal, offset)
				}
			}
			data := make([]byte, 256)
			if err := nativePulseEscape(data, signal, offset); err != 0 {
				t.Fatalf("C escape encoder error=%d", err)
			}
			got := compare(data, 16, signal, offset)
			for i, v := range got {
				want := int16(1023)
				if i == 0 {
					want = 17407
				}
				if i%2 == 0 {
					want = -want
				}
				if v != want {
					t.Fatalf("escape pulse[%d]=%d want=%d", i, v, want)
				}
			}
		}
	}
}
