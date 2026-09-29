//go:build compareopus && cgo

package main

import (
	"github.com/kazzmir/opus-go/opuscc"
	"testing"
	"unsafe"
)

func TestSilkAPIResetAgainstC(t *testing.T) {
	for _, init := range []bool{false, true} {
		var g opuscc.OpusT_silk_decoder
		raw := unsafe.Slice((*byte)(unsafe.Pointer(&g)), int(unsafe.Sizeof(g)))
		for i := range raw {
			raw[i] = 0xa5
		}
		g.FnChannelsAPI = 2
		g.FnChannelsInternal = 1
		var ret int32
		if init {
			ret = opuscc.Opus_silk_InitDecoder(nil, &g)
		} else {
			ret = opuscc.Opus_silk_ResetDecoder(nil, &g)
		}
		cr, meta := nativeSilkAPIReset(init)
		channel, _, zero := nativeDecoderReset(init)
		want := opuscc.OpusT_silk_decoder{FnChannelsAPI: meta[0], FnChannelsInternal: meta[1], Fprev_decode_only_middle: meta[2]}
		for i := range want.Fchannel_state {
			want.Fchannel_state[i] = channel
		}
		if ret != cr || g != want || !zero || meta[3] != 1 || meta[4] != 1 {
			t.Fatalf("init=%v ret=%d/%d meta=%v", init, ret, cr, meta)
		}
	}
}
