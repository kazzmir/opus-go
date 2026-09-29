package opuscc

import (
	"testing"
	"unsafe"
)

func TestResamplerInitPointers(t *testing.T) {
	for _, enc := range []int32{0, 1} {
		for _, in := range []int32{8000, 12000, 16000, 24000, 48000} {
			for _, out := range []int32{8000, 12000, 16000, 24000, 48000} {
				if enc == 0 && in > 16000 || enc != 0 && out > 16000 {
					continue
				}
				var owner struct {
					before uint32
					state  OpusT_silk_resampler_state_struct
					after  uint32
				}
				owner.before = 123
				owner.after = 456
				raw := unsafe.Slice((*byte)(unsafe.Pointer(&owner.state)), int(unsafe.Sizeof(owner.state)))
				for i := range raw {
					raw[i] = 0xa5
				}
				if ret := Opus_silk_resampler_init(nil, &owner.state, in, out, enc); ret != 0 {
					t.Fatal(ret)
				}
				s := owner.state
				if owner.before != 123 || owner.after != 456 {
					t.Fatalf("overwrite %d -> %d enc=%d", in, out, enc)
				}
				if s.FFs_in_kHz != in/1000 || s.FFs_out_kHz != out/1000 || s.FbatchSize != in/100 {
					t.Fatal("rates")
				}
				if s.FsIIR != [6]int32{} || s.FsFIR.Fi32 != [36]int32{} || s.FdelayBuf != [96]int16{} {
					t.Fatal("history not reset")
				}
				if s.FinputDelay < 0 || s.FinputDelay > s.FFs_in_kHz {
					t.Fatal("delay")
				}
			}
		}
	}
}
