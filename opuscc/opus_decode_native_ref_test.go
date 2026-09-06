package opuscc

import (
	"math"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

// C reference values from /tmp/opencode/nativeref.c, which calls the exported
// opus_decode_native directly (plus opus_packet_pad and the repacketizer for
// the padded / self-delimited packets). All packets come from the C encoder
// (deterministic LCG noise, 20 kbps WIDEBAND -> SILK WB), so every frame is
// pure SILK and the Go build is expected to be bit-exact against C.
//
// Covers the opus_decode_native locals: toc/offset/size[48]/padding/
// padding_len via opus_packet_parse_impl (single-frame, 3-frame 60 ms code-3
// packet, padded packet with real padding + extension iterator,
// OPUS_SET_IGNORE_EXTENSIONS zeroing, self-delimited parsing), the multi-frame
// decode loop with packet_offset out-param, the decode_fec early-PLC branch,
// bad-arg checks, and the int16 opus_decode wrapper (soft-clip path).

func TestOpusDecodeNativeCReference(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()
	setupResamplerPseudostack(tls)

	p20 := mustHex(t, "4ca8d089b659ccaf812f110c4dcf5534387e253fe1b09e8a9d148cbfec4e70417f5db5fca4")
	p60 := mustHex(t, "5cee8d90596d8af0f0d2ad3b3c4fb41ec5e54572e252746cfdafe15414893ad4b3266349870335f48e50e3e859b64abc86816b3d9197342c62362cfd6f21a58bf15633e95b984854c9436dd5a3507e116609440521cfe55117e844c4044549367c6f27cc6a3585667ae2ff9c998645f2ac7259b7b723aeac0a63469595f365f6154c81d6b4eeb380")
	pad := mustHex(t, "4f410ba8d089b659ccaf812f110c4dcf5534387e253fe1b09e8a9d148cbfec4e70417f5db5fca40000000000000000000000")
	selfdel := mustHex(t, "4c24a8d089b659ccaf812f110c4dcf5534387e253fe1b09e8a9d148cbfec4e70417f5db5fca4")

	pcm := make([]float32, 5760*2)
	pcmPtr := uintptr(unsafe.Pointer(unsafe.SliceData(pcm)))
	var po int32

	/* multi-frame 60 ms code-3 packet, packet_offset out-param */
	dec, err := Opus_opus_decoder_create(tls, 48000, 2)
	if err != nil || dec == 0 {
		t.Fatalf("decoder create: %v", err)
	}
	dp := (*OpusT_OpusDecoder)(unsafe.Pointer(dec))
	po = -1
	n := Opus_opus_decode_native(tls, dec, uintptr(unsafe.Pointer(&p60[0])), int32(len(p60)), pcmPtr, 2880, 0, 0, uintptr(unsafe.Pointer(&po)), 0, 0, 0)
	if n != 2880 {
		t.Fatalf("multi60: n=%d want 2880", n)
	}
	if got, want := fnv1aFloats(pcm[:2880*2]), uint32(0x4a1f8ba0); got != want {
		t.Fatalf("multi60: fnv=%08x want %08x", got, want)
	}
	if got, want := dp.FrangeFinal, uint32(0x1bb6c5bb); got != want {
		t.Fatalf("multi60: rangeFinal=%08x want %08x", got, want)
	}
	if got, want := dp.Flast_packet_duration, int32(2880); got != want {
		t.Fatalf("multi60: duration=%d want %d", got, want)
	}
	if got, want := po, int32(136); got != want {
		t.Fatalf("multi60: packet_offset=%d want %d", got, want)
	}
	if got, want := math.Float32bits(pcm[0]), uint32(0); got != want {
		t.Fatalf("multi60: pcm[0]=%08x want %08x", got, want)
	}
	Opus_opus_decoder_destroy(tls, dec)

	/* plain 20 ms, then the same packet padded with 13 bytes (real padding +
	   extension iterator), then padded again with ignore_extensions */
	dec, err = Opus_opus_decoder_create(tls, 48000, 2)
	if err != nil || dec == 0 {
		t.Fatalf("decoder create: %v", err)
	}
	dp = (*OpusT_OpusDecoder)(unsafe.Pointer(dec))
	po = -1
	n = Opus_opus_decode_native(tls, dec, uintptr(unsafe.Pointer(&p20[0])), int32(len(p20)), pcmPtr, 960, 0, 0, uintptr(unsafe.Pointer(&po)), 0, 0, 0)
	if n != 960 {
		t.Fatalf("plain20: n=%d want 960", n)
	}
	if got, want := fnv1aFloats(pcm[:960*2]), uint32(0x1ec618f7); got != want {
		t.Fatalf("plain20: fnv=%08x want %08x", got, want)
	}
	if got, want := dp.FrangeFinal, uint32(0x743dec00); got != want {
		t.Fatalf("plain20: rangeFinal=%08x want %08x", got, want)
	}
	if got, want := po, int32(37); got != want {
		t.Fatalf("plain20: packet_offset=%d want %d", got, want)
	}
	po = -1
	n = Opus_opus_decode_native(tls, dec, uintptr(unsafe.Pointer(&pad[0])), int32(len(pad)), pcmPtr, 960, 0, 0, uintptr(unsafe.Pointer(&po)), 0, 0, 0)
	if n != 960 {
		t.Fatalf("padded20: n=%d want 960", n)
	}
	if got, want := fnv1aFloats(pcm[:960*2]), uint32(0x1e06ca04); got != want {
		t.Fatalf("padded20: fnv=%08x want %08x", got, want)
	}
	if got, want := dp.FrangeFinal, uint32(0x743dec00); got != want {
		t.Fatalf("padded20: rangeFinal=%08x want %08x", got, want)
	}
	if got, want := po, int32(50); got != want {
		t.Fatalf("padded20: packet_offset=%d want %d", got, want)
	}
	if got, want := math.Float32bits(pcm[0]), uint32(0xbf086200); got != want {
		t.Fatalf("padded20: pcm[0]=%08x want %08x", got, want)
	}
	var slot [2]int32
	if got := Opus_opus_decoder_ctl(tls, dec, OPUS_SET_IGNORE_EXTENSIONS_REQUEST, libc.VaList(uintptr(unsafe.Pointer(&slot[0])), int32(1))); got != OPUS_OK {
		t.Fatalf("OPUS_SET_IGNORE_EXTENSIONS: %d", got)
	}
	po = -1
	n = Opus_opus_decode_native(tls, dec, uintptr(unsafe.Pointer(&pad[0])), int32(len(pad)), pcmPtr, 960, 0, 0, uintptr(unsafe.Pointer(&po)), 0, 0, 0)
	if n != 960 {
		t.Fatalf("ignoreext: n=%d want 960", n)
	}
	if got, want := fnv1aFloats(pcm[:960*2]), uint32(0x1e06ca04); got != want {
		t.Fatalf("ignoreext: fnv=%08x want %08x", got, want)
	}
	if got, want := po, int32(50); got != want {
		t.Fatalf("ignoreext: packet_offset=%d want %d", got, want)
	}
	Opus_opus_decoder_destroy(tls, dec)

	/* self-delimited packet */
	dec, err = Opus_opus_decoder_create(tls, 48000, 2)
	if err != nil || dec == 0 {
		t.Fatalf("decoder create: %v", err)
	}
	dp = (*OpusT_OpusDecoder)(unsafe.Pointer(dec))
	po = -1
	n = Opus_opus_decode_native(tls, dec, uintptr(unsafe.Pointer(&selfdel[0])), int32(len(selfdel)), pcmPtr, 960, 0, 1, uintptr(unsafe.Pointer(&po)), 0, 0, 0)
	if n != 960 {
		t.Fatalf("selfdel: n=%d want 960", n)
	}
	if got, want := fnv1aFloats(pcm[:960*2]), uint32(0x1ec618f7); got != want {
		t.Fatalf("selfdel: fnv=%08x want %08x", got, want)
	}
	if got, want := dp.FrangeFinal, uint32(0x743dec00); got != want {
		t.Fatalf("selfdel: rangeFinal=%08x want %08x", got, want)
	}
	if got, want := po, int32(38); got != want {
		t.Fatalf("selfdel: packet_offset=%d want %d", got, want)
	}
	Opus_opus_decoder_destroy(tls, dec)

	/* decode_fec early-PLC branch: frame_size < packet_frame_size */
	dec, err = Opus_opus_decoder_create(tls, 48000, 2)
	if err != nil || dec == 0 {
		t.Fatalf("decoder create: %v", err)
	}
	dp = (*OpusT_OpusDecoder)(unsafe.Pointer(dec))
	n = Opus_opus_decode_native(tls, dec, uintptr(unsafe.Pointer(&p20[0])), int32(len(p20)), pcmPtr, 960, 0, 0, 0, 0, 0, 0)
	if n != 960 {
		t.Fatalf("fecshort warmup: n=%d want 960", n)
	}
	n = Opus_opus_decode_native(tls, dec, uintptr(unsafe.Pointer(&p60[0])), int32(len(p60)), pcmPtr, 480, 1, 0, 0, 0, 0, 0)
	if n != 480 {
		t.Fatalf("fecshort: n=%d want 480", n)
	}
	if got, want := fnv1aFloats(pcm[:480*2]), uint32(0x3ee2e5c1); got != want {
		t.Fatalf("fecshort: fnv=%08x want %08x", got, want)
	}
	if got, want := dp.FrangeFinal, uint32(0); got != want {
		t.Fatalf("fecshort: rangeFinal=%08x want %08x", got, want)
	}
	if got, want := dp.Flast_packet_duration, int32(480); got != want {
		t.Fatalf("fecshort: duration=%d want %d", got, want)
	}
	if got, want := math.Float32bits(pcm[0]), uint32(0xbf086200); got != want {
		t.Fatalf("fecshort: pcm[0]=%08x want %08x", got, want)
	}
	Opus_opus_decoder_destroy(tls, dec)

	/* bad args */
	dec, err = Opus_opus_decoder_create(tls, 48000, 2)
	if err != nil || dec == 0 {
		t.Fatalf("decoder create: %v", err)
	}
	if got := Opus_opus_decode_native(tls, dec, uintptr(unsafe.Pointer(&p20[0])), int32(len(p20)), pcmPtr, 100, 1, 0, 0, 0, 0, 0); got != -1 {
		t.Fatalf("badsize: %d want -1", got)
	}
	if got := Opus_opus_decode_native(tls, dec, uintptr(unsafe.Pointer(&p20[0])), int32(len(p20)), pcmPtr, 960, 2, 0, 0, 0, 0, 0); got != -1 {
		t.Fatalf("badfec: %d want -1", got)
	}
	Opus_opus_decoder_destroy(tls, dec)

	/* int16 wrapper on the 60 ms packet (soft clip path) */
	dec, err = Opus_opus_decoder_create(tls, 48000, 2)
	if err != nil || dec == 0 {
		t.Fatalf("decoder create: %v", err)
	}
	dp = (*OpusT_OpusDecoder)(unsafe.Pointer(dec))
	pcm16 := make([]int16, 5760*2)
	n = Opus_opus_decode(tls, dec, uintptr(unsafe.Pointer(&p60[0])), int32(len(p60)), uintptr(unsafe.Pointer(unsafe.SliceData(pcm16))), 5760, 0)
	if n != 2880 {
		t.Fatalf("i16: n=%d want 2880", n)
	}
	if got, want := fnv1aInt16s(pcm16[:2880*2]), uint32(0x5467880c); got != want {
		t.Fatalf("i16: fnv=%08x want %08x", got, want)
	}
	if got, want := dp.FrangeFinal, uint32(0x1bb6c5bb); got != want {
		t.Fatalf("i16: rangeFinal=%08x want %08x", got, want)
	}
	if got, want := dp.Flast_packet_duration, int32(2880); got != want {
		t.Fatalf("i16: duration=%d want %d", got, want)
	}
	for j, want := range []int16{0, 0, 0, 0} {
		if got := pcm16[j]; got != want {
			t.Fatalf("i16: pcm16[%d]=%d want %d", j, got, want)
		}
	}
	Opus_opus_decoder_destroy(tls, dec)
}
