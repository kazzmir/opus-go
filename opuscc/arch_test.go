package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

// decoderReferenceLayoutDelta adjusts amd64 C fixture sizes for the two
// pointer-containing subdecoders. Opus aligns their allocations to 8 bytes;
// the sample buffers and the outer decoder header have fixed-width layouts.
func decoderReferenceLayoutDelta() int32 {
	align8 := func(n uintptr) int32 { return int32((n + 7) &^ 7) }
	return align8(unsafe.Sizeof(OpusT_silk_decoder{})) - 8808 +
		align8(unsafe.Sizeof(OpusT_OpusCustomDecoder{})) - 120
}

// These tests derive layouts from Go types rather than amd64 C reference sizes.
func TestArchPacketFramePointers(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	packet := [3]byte{0x81, 0x12, 0x34} // Two equal-sized CELT frames.
	storage := struct {
		frames [2]uintptr
		guard  [2]uintptr
	}{guard: [2]uintptr{123, 456}}
	var sizes [2]int16
	n := Opus_opus_packet_parse(tls, uintptr(unsafe.Pointer(&packet[0])), int32(len(packet)), 0, uintptr(unsafe.Pointer(&storage.frames)), uintptr(unsafe.Pointer(&sizes)), 0)
	if n != 2 {
		t.Fatalf("frame count = %d", n)
	}
	for i := range storage.frames {
		if storage.frames[i] != uintptr(unsafe.Pointer(&packet[i+1])) || sizes[i] != 1 {
			t.Fatalf("frame %d: wrong pointer or size", i)
		}
	}
	if storage.guard != [2]uintptr{123, 456} {
		t.Fatal("frame pointers overwrite adjacent memory")
	}
}

func TestArchDeemphasisPointers(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)
	left, right := [2]float32{32768, 65536}, [2]float32{-32768, -65536}
	input := [2]uintptr{uintptr(unsafe.Pointer(&left)), uintptr(unsafe.Pointer(&right))}
	var coef [4]float32
	for _, accum := range []int32{0, 1} {
		var output [4]float32
		var mem [2]float32
		deemphasis(tls, uintptr(unsafe.Pointer(&input)), uintptr(unsafe.Pointer(&output)), 2, 2, 1, uintptr(unsafe.Pointer(&coef)), uintptr(unsafe.Pointer(&mem)), accum)
		if output != [4]float32{1, -1, 2, -2} {
			t.Fatalf("accum %d: output = %v", accum, output)
		}
	}
}

func TestArchSilkDecoderLayout(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	var size int32
	Opus_silk_Get_Decoder_Size(tls, uintptr(unsafe.Pointer(&size)))
	if size != int32(unsafe.Sizeof(OpusT_silk_decoder{})) {
		t.Fatalf("decoder size = %d", size)
	}
	storage := struct {
		decoder OpusT_silk_decoder
		guard   [64]byte
	}{}
	for i := range storage.guard {
		storage.guard[i] = 0xa5
	}
	for _, init := range []func(*libc.TLS, uintptr) int32{Opus_silk_InitDecoder, Opus_silk_ResetDecoder} {
		storage.decoder.FsStereo.Fpred_prev_Q13 = [2]int16{123, 456}
		if ret := init(tls, uintptr(unsafe.Pointer(&storage.decoder))); ret != 0 {
			t.Fatalf("init/reset = %d", ret)
		}
		for i, ch := range storage.decoder.Fchannel_state {
			if ch.Fprev_gain_Q16 != 65536 || ch.Ffirst_frame_after_reset != 1 {
				t.Fatalf("channel %d not initialized", i)
			}
		}
		if storage.decoder.FsStereo != (OpusT_stereo_dec_state{}) {
			t.Fatal("stereo state not cleared")
		}
		for _, b := range storage.guard {
			if b != 0xa5 {
				t.Fatal("init/reset overwrites adjacent memory")
			}
		}
	}
	ch := &storage.decoder.Fchannel_state[0]
	ch.Fnb_subfr = MAX_NB_SUBFR
	if ret := Opus_silk_decoder_set_fs(tls, uintptr(unsafe.Pointer(ch)), 16, 48000); ret != 0 {
		t.Fatalf("set_fs = %d", ret)
	}
	var want OpusT_silk_resampler_state_struct
	if ret := Opus_silk_resampler_init(tls, uintptr(unsafe.Pointer(&want)), 16000, 48000, 0); ret != 0 {
		t.Fatalf("resampler init = %d", ret)
	}
	if ch.Fresampler_state != want {
		t.Fatal("resampler initialized at wrong offset")
	}
}
