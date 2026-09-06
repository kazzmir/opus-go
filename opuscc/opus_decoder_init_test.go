package opuscc

import (
	"fmt"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

// C expectations: /tmp/opencode/decoder_init_ref.c includes
// ../opus/src/opus_decoder.c and links ../opus/.libs/libopus.a.
// Each call starts with 0xa5-filled storage. Architecture dispatch is
// intentionally not compared: the Go port disables SIMD.
func TestOpusDecoderInitCReference(t *testing.T) {
	for _, tc := range []struct{ fs, frame int32 }{
		{48000, 120}, {24000, 60}, {16000, 40}, {12000, 30}, {8000, 20},
		{44100, 0}, {0, 0}, {-1, 0},
	} {
		for ch := int32(0); ch <= 3; ch++ {
			t.Run(fmt.Sprintf("%d/%d", tc.fs, ch), func(t *testing.T) {
				tls := libc.NewTLS()
				defer tls.Close()
				size := Opus_opus_decoder_get_size(tls, 2)
				p := libc.Xmalloc(tls, uint64(size))
				defer libc.Xfree(tls, p)
				memory := unsafe.Slice((*byte)(unsafe.Pointer(p)), int(size))
				for i := range memory {
					memory[i] = 0xa5
				}
				valid := tc.frame != 0 && (ch == 1 || ch == 2)
				ret := Opus_opus_decoder_init(tls, p, tc.fs, ch)
				if !valid {
					if ret != -1 {
						t.Fatalf("return = %d, want -1", ret)
					}
					for i, b := range memory {
						if b != 0xa5 {
							t.Fatalf("invalid init changed byte %d", i)
						}
					}
					return
				}
				for pass := 0; pass < 2; pass++ {
					if ret != 0 {
						t.Fatalf("return = %d, want 0", ret)
					}
					d := (*OpusT_OpusDecoder)(unsafe.Pointer(p))
					got := []int32{d.Fsilk_dec_offset, d.Fcelt_dec_offset, d.Fchannels, d.Fstream_channels, d.FFs, d.FDecControl.FAPI_sampleRate, d.FDecControl.FnChannelsAPI, d.Fframe_size, d.Fdecode_gain, d.Fcomplexity, d.Fignore_extensions, d.Fbandwidth, d.Fmode, d.Fprev_mode, d.Fprev_redundancy, d.Flast_packet_duration, int32(d.FrangeFinal)}
					// The C fixture uses an 8808-byte amd64 SILK decoder.
					celtOffset := int32(104) + int32((unsafe.Sizeof(OpusT_silk_decoder{})+7)&^7)
					want := []int32{104, celtOffset, ch, ch, tc.fs, tc.fs, ch, tc.frame, 0, 0, 0, 0, 0, 0, 0, 0, 0}
					for i, v := range got {
						if v != want[i] {
							t.Fatalf("pass %d field %d = %d, want %d", pass, i, v, want[i])
						}
					}
					// The reference explicitly calls CELT_SET_SIGNALLING(0).
					celt := (*OpusT_OpusCustomDecoder)(unsafe.Pointer(p + uintptr(d.Fcelt_dec_offset)))
					if celt.Fsignalling != 0 {
						t.Fatalf("signalling = %d", celt.Fsignalling)
					}
					if pass == 0 {
						d.Fdecode_gain = 123
						d.Fprev_mode = 1002
						celt.Fsignalling = 1
						ret = Opus_opus_decoder_init(tls, p, tc.fs, ch)
					}
				}
			})
		}
	}
}
