//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"github.com/kazzmir/opus-go/opusccenc"
	"math/rand"
	"slices"
	"testing"
	"unsafe"
)

func TestLPCAnalysisAgainstC(t *testing.T) {
	rng := rand.New(rand.NewSource(3403))
	for _, d := range []int{6, 8, 10, 16, 24} {
		for _, extra := range []int{0, 1, 32, 320} {
			for trial := 0; trial < 100; trial++ {
				n := d + extra
				in := make([]int16, n)
				b := make([]int16, d)
				for i := range in {
					in[i] = int16(rng.Uint32())
					if trial == 0 {
						in[i] = -32768
					}
				}
				for i := range b {
					b[i] = int16(rng.Uint32())
					if trial == 0 {
						b[i] = -32768
					}
				}
				before, bc := slices.Clone(in), slices.Clone(b)
				g, c := make([]int16, n), make([]int16, n)
				opuscc.Opus_silk_LPC_analysis_filter(nil, &g[0], &in[0], &b[0], int32(n), int32(d), 0)
				nativeLPCAnalysis(c, in, b)
				e := make([]int16, n)
				opusccenc.Opus_silk_LPC_analysis_filter(nil, uintptr(unsafe.Pointer(&e[0])), uintptr(unsafe.Pointer(&in[0])), uintptr(unsafe.Pointer(&b[0])), int32(n), int32(d), 0)
				if !slices.Equal(e, c) {
					t.Fatalf("encoder d=%d n=%d trial=%d", d, n, trial)
				}
				if !slices.Equal(g, c) || !slices.Equal(in, before) || !slices.Equal(b, bc) {
					t.Fatalf("d=%d n=%d trial=%d Go=%v C=%v", d, n, trial, g, c)
				}
				// Coefficients may share output storage: later taps must remain live reads.
				g = make([]int16, n+d+2)
				copy(g[1+d:], b)
				c, e = slices.Clone(g), slices.Clone(g)
				opuscc.Opus_silk_LPC_analysis_filter(nil, &g[1], &in[0], &g[1+d], int32(n), int32(d), 0)
				opusccenc.Opus_silk_LPC_analysis_filter(nil, uintptr(unsafe.Pointer(&e[1])), uintptr(unsafe.Pointer(&in[0])), uintptr(unsafe.Pointer(&e[1+d])), int32(n), int32(d), 0)
				nativeLPCAnalysis(c[1:1+n], in, c[1+d:1+2*d])
				if !slices.Equal(g, c) || !slices.Equal(e, c) {
					t.Fatalf("coefficient/output overlap d=%d n=%d trial=%d", d, n, trial)
				}
				// The C API has no restrict qualifier: preserve forward store order when overlapping.
				for _, shift := range []int{-1, 0, 1} {
					g = make([]int16, n+2)
					copy(g[1:], in)
					c = slices.Clone(g)
					e = slices.Clone(g)
					opusccenc.Opus_silk_LPC_analysis_filter(nil, uintptr(unsafe.Pointer(&e[1+shift])), uintptr(unsafe.Pointer(&e[1])), uintptr(unsafe.Pointer(&b[0])), int32(n), int32(d), 0)
					opuscc.Opus_silk_LPC_analysis_filter(nil, &g[1+shift], &g[1], &b[0], int32(n), int32(d), 0)
					nativeLPCAnalysis(c[1+shift:1+shift+n], c[1:1+n], b)
					if !slices.Equal(g, c) || !slices.Equal(e, c) {
						t.Fatalf("overlap d=%d n=%d trial=%d shift=%d", d, n, trial, shift)
					}
				}
			}
		}
	}
}
