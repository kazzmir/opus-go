//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"github.com/kazzmir/opus-go/opusccenc"
	"testing"
)

func TestVADInitAgainstC(t *testing.T) {
	want, cr := nativeVADInit()
	g := opuscc.OpusT_silk_VAD_state{Fcounter: 999, FHPstate: -123, FAnaState: [2]int32{456, -789}}
	if gr := opuscc.Opus_silk_VAD_Init(nil, &g); gr != cr || g != want {
		t.Fatalf("Go=%+v (%d) C=%+v (%d)", g, gr, want, cr)
	}
	if gr := opuscc.Opus_silk_VAD_Init(nil, &g); gr != cr || g != want {
		t.Fatal("reinitialization differs")
	}
	enc := opusccenc.OpusT_silk_VAD_state{Fcounter: 999, FHPstate: -123, FAnaState: [2]int32{456, -789}}
	for pass := 0; pass < 2; pass++ {
		if ret := opusccenc.Opus_silk_VAD_Init(nil, &enc); ret != cr || enc != opusccenc.OpusT_silk_VAD_state(want) {
			t.Fatalf("encoder VAD init differs from C: %+v", enc)
		}
	}
}
