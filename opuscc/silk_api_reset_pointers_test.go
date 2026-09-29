package opuscc

import (
	"testing"
	"unsafe"
)

func TestSilkAPIResetPointers(t *testing.T) {
	for _, init := range []bool{false, true} {
		var owner struct {
			before uint64
			state  OpusT_silk_decoder
			after  uint64
		}
		owner.before = 123
		owner.after = 456
		raw := unsafe.Slice((*byte)(unsafe.Pointer(&owner.state)), int(unsafe.Sizeof(owner.state)))
		for i := range raw {
			raw[i] = 0xa5
		}
		owner.state.FnChannelsAPI = 2
		owner.state.FnChannelsInternal = 1
		var want OpusT_silk_decoder
		want.FnChannelsAPI = 2
		want.FnChannelsInternal = 1
		for i := range want.Fchannel_state {
			Opus_silk_reset_decoder(nil, &want.Fchannel_state[i])
		}
		var ret int32
		if init {
			ret = Opus_silk_InitDecoder(nil, &owner.state)
		} else {
			ret = Opus_silk_ResetDecoder(nil, &owner.state)
		}
		if ret != 0 || owner.state != want || owner.before != 123 || owner.after != 456 {
			t.Fatalf("init=%v ret=%d", init, ret)
		}
	}
}
