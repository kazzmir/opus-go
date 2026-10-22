package opuscc

import (
	"runtime"
	"testing"
)

// From packet_duration_pointers_test.go.
func TestPacketDurationPointers(t *testing.T) {
	for _, n := range []int32{-2, -1, 0} {
		if got := Opus_opus_packet_get_nb_samples(nil, n, 48000); got != -1 {
			t.Fatal(n, got)
		}
	}
	toc := byte(3)
	if got := Opus_opus_packet_get_nb_samples(&toc, 1, 48000); got != -4 {
		t.Fatal(got)
	}
	packet := [2]byte{0x83, 48}
	if got := Opus_opus_packet_get_nb_samples(&packet[0], 2, 48000); got != 5760 {
		t.Fatal(got)
	}
	packet[1] = 49
	if got := Opus_opus_packet_get_nb_samples(&packet[0], 2, 48000); got != -4 {
		t.Fatal(got)
	}
	packet[1] = 0
	if got := Opus_opus_packet_get_nb_samples(&packet[0], 2, 48000); got != 0 {
		t.Fatal(got)
	}
}

// From packet_header_pointers_test.go.
func TestPacketHeaderPointers(t *testing.T) {
	// RFC 6716's 32 TOC configurations, expressed independently of bit masks.
	bandwidths := [32]int32{
		OPUS_BANDWIDTH_NARROWBAND, OPUS_BANDWIDTH_NARROWBAND, OPUS_BANDWIDTH_NARROWBAND, OPUS_BANDWIDTH_NARROWBAND,
		OPUS_BANDWIDTH_MEDIUMBAND, OPUS_BANDWIDTH_MEDIUMBAND, OPUS_BANDWIDTH_MEDIUMBAND, OPUS_BANDWIDTH_MEDIUMBAND,
		OPUS_BANDWIDTH_WIDEBAND, OPUS_BANDWIDTH_WIDEBAND, OPUS_BANDWIDTH_WIDEBAND, OPUS_BANDWIDTH_WIDEBAND,
		OPUS_BANDWIDTH_SUPERWIDEBAND, OPUS_BANDWIDTH_SUPERWIDEBAND, OPUS_BANDWIDTH_FULLBAND, OPUS_BANDWIDTH_FULLBAND,
		OPUS_BANDWIDTH_NARROWBAND, OPUS_BANDWIDTH_NARROWBAND, OPUS_BANDWIDTH_NARROWBAND, OPUS_BANDWIDTH_NARROWBAND,
		OPUS_BANDWIDTH_WIDEBAND, OPUS_BANDWIDTH_WIDEBAND, OPUS_BANDWIDTH_WIDEBAND, OPUS_BANDWIDTH_WIDEBAND,
		OPUS_BANDWIDTH_SUPERWIDEBAND, OPUS_BANDWIDTH_SUPERWIDEBAND, OPUS_BANDWIDTH_SUPERWIDEBAND, OPUS_BANDWIDTH_SUPERWIDEBAND,
		OPUS_BANDWIDTH_FULLBAND, OPUS_BANDWIDTH_FULLBAND, OPUS_BANDWIDTH_FULLBAND, OPUS_BANDWIDTH_FULLBAND,
	}
	// Frame durations in units of 2.5 ms.
	durations := [32]int32{4, 8, 16, 24, 4, 8, 16, 24, 4, 8, 16, 24, 4, 8, 4, 8, 1, 2, 4, 8, 1, 2, 4, 8, 1, 2, 4, 8, 1, 2, 4, 8}
	for config := 0; config < 32; config++ {
		for stereo := 0; stereo < 2; stereo++ {
			for code := 0; code < 4; code++ {
				toc := byte(config*8 + stereo*4 + code)
				if got := Opus_opus_packet_get_bandwidth(&toc); got != bandwidths[config] {
					t.Fatalf("TOC %#x bandwidth: %d", toc, got)
				}
				if got := Opus_opus_packet_get_nb_channels(&toc); got != int32(stereo+1) {
					t.Fatalf("TOC %#x channels: %d", toc, got)
				}
				for _, rate := range []int32{8000, 12000, 16000, 24000, 48000} {
					want := rate * durations[config] / 400
					if got := Opus_opus_packet_get_samples_per_frame(&toc, rate); got != want {
						t.Fatalf("TOC %#x rate %d: got %d samples, want %d", toc, rate, got, want)
					}
				}
				wantFrames := [4]int32{1, 2, 2, -4}[code]
				if got := Opus_opus_packet_get_nb_frames(&toc, 1); got != wantFrames {
					t.Fatalf("TOC %#x one-byte frames: %d", toc, got)
				}
				for second := 0; second < 256; second++ {
					packet := [2]byte{toc, byte(second)}
					want := wantFrames
					if code == 3 {
						want = int32(second % 64)
					}
					if got := Opus_opus_packet_get_nb_frames(&packet[0], 2); got != want {
						t.Fatalf("packet %x: got %d frames, want %d", packet, got, want)
					}
				}
			}
		}
	}
	for _, length := range []int32{-1, 0} {
		if got := Opus_opus_packet_get_nb_frames(nil, length); got != -1 {
			t.Fatalf("length %d: got %d, want OPUS_BAD_ARG", length, got)
		}
	}
}

// From packet_leaf_pointers_test.go.
func TestPacketLBRRPointers(t *testing.T) {
	for _, tc := range []struct {
		toc, payload byte
		want         int32
	}{{0, 0x40, 1}, {0, 0, 0}, {4, 0x10, 1}, {0x10, 0x20, 1}, {0x18, 0x10, 1}, {0x1c, 1, 1}, {0x80, 0xff, 0}} {
		packet := [2]byte{tc.toc, tc.payload}
		entropyInitGrowStack(12)
		runtime.GC()
		if got := Opus_opus_packet_has_lbrr(&packet[0], 2); got != tc.want {
			t.Fatal(tc, got)
		}
	}
	header := byte(0)
	if Opus_opus_packet_has_lbrr(&header, 1) != 0 {
		t.Fatal("zero-size frame")
	}
	if Opus_opus_packet_has_lbrr(&header, 0) != -4 {
		t.Fatal("empty SILK packet")
	}
	header = 0x80
	if Opus_opus_packet_has_lbrr(&header, -1) != 0 {
		t.Fatal("CELT short circuit")
	}
}

func TestPacketParsePointers(t *testing.T) {
	var frames [48]*byte
	var sizes [48]int16
	var toc byte
	offset := int32(-1)
	func() {
		packet := []byte{0x82, 1, 11, 12, 13}
		r := Opus_opus_packet_parse(&packet[0], 5, &toc, &frames, &sizes, &offset)
		if r != 2 || sizes[0] != 1 || sizes[1] != 2 || offset != 2 || toc != 0x82 {
			t.Fatal(r, sizes, offset, toc)
		}
	}()
	entropyInitGrowStack(12)
	runtime.GC()
	if *frames[0] != 11 || *frames[1] != 12 {
		t.Fatal("returned frame ownership")
	}
	if r := Opus_opus_packet_parse(nil, 0, nil, nil, &sizes, nil); r != -4 {
		t.Fatal("empty", r)
	}
	if r := Opus_opus_packet_parse(nil, 0, nil, nil, nil, nil); r != -1 {
		t.Fatal("bad size", r)
	}
}

func TestPacketParseImplPointers(t *testing.T) {
	var frames [48]*byte
	var size [48]int16
	var toc byte
	var payload, offset, padLen int32
	var pad *byte
	func() {
		packet := []byte{0x83, 0xc2, 2, 1, 0xaa, 0xbb, 0xcc, 0xdd, 0xee}
		r := Opus_opus_packet_parse_impl(&packet[0], 9, 0, &toc, &frames, &size, &payload, &offset, &pad, &padLen)
		if r != 2 || size[0] != 1 || size[1] != 2 || payload != 4 || offset != 9 || padLen != 2 || toc != 0x83 {
			t.Fatal(r, size, payload, offset, padLen, toc)
		}
	}()
	entropyInitGrowStack(12)
	runtime.GC()
	if *frames[0] != 0xaa || *frames[1] != 0xbb || *pad != 0xdd {
		t.Fatal("packet ownership")
	}
	packet := []byte{0x81, 3, 11, 12, 13, 14, 15, 16, 99, 99}
	r := Opus_opus_packet_parse_impl(&packet[0], 10, 1, nil, &frames, &size, &payload, &offset, nil, nil)
	if r != 2 || size[0] != 3 || size[1] != 3 || payload != 2 || offset != 8 {
		t.Fatal("self-delimited", r, size, payload, offset)
	}
	pad = &packet[0]
	padLen = 99
	if r := Opus_opus_packet_parse_impl(nil, -1, 0, nil, nil, nil, nil, nil, &pad, &padLen); r != -1 || pad != nil || padLen != 0 {
		t.Fatal("error padding", r, pad, padLen)
	}
}

func TestPacketSizePointers(t *testing.T) {
	// Every representable Opus frame length, including the 251/252 boundary.
	// Check the C wire format independently as well as round-tripping it.
	for size := int32(0); size <= 1275; size++ {
		buf := [4]byte{0xa5, 0xcc, 0xcc, 0x5a}
		n := Opus_encode_size(size, &buf[1])
		wantN := int32(1)
		want := [4]byte{0xa5, byte(size), 0xcc, 0x5a}
		if size >= 252 {
			wantN = 2
			want[1] = byte(252 + size%4)
			want[2] = byte((size - int32(want[1])) / 4)
		}
		if n != wantN || buf != want {
			t.Fatalf("size %d: n=%d bytes=%x, want n=%d bytes=%x", size, n, buf, wantN, want)
		}
		decoded := [3]int16{1234, -99, 5678}
		if got := parse_size(&buf[1], n, &decoded[1]); got != n || decoded != [3]int16{1234, int16(size), 5678} {
			t.Fatalf("size %d: parsed n=%d output=%v", size, got, decoded)
		}
	}
	// The single-byte branch must not require a second accessible byte.
	var single byte
	if n := Opus_encode_size(251, &single); n != 1 || single != 251 {
		t.Fatalf("single byte encode: n=%d byte=%d", n, single)
	}
	for _, length := range []int32{-1, 0} {
		var size int16
		if n := parse_size(nil, length, &size); n != -1 || size != -1 {
			t.Fatalf("length %d: n=%d size=%d", length, n, size)
		}
	}
	for first := 0; first < 256; first++ {
		b := byte(first)
		var size int16
		n := parse_size(&b, 1, &size)
		if first < 252 {
			if n != 1 || size != int16(first) {
				t.Fatalf("single byte %d: n=%d size=%d", first, n, size)
			}
		} else if n != -1 || size != -1 {
			t.Fatalf("truncated two-byte size %d: n=%d size=%d", first, n, size)
		}
	}
}

func TestPacketModePointers(t *testing.T) {
	for toc := 0; toc < 256; toc++ {
		// RFC 6716 TOC configuration ranges: SILK 0..11, hybrid 12..15,
		// CELT 16..31. Stereo and frame-count bits do not affect the mode.
		config := toc >> 3
		want := int32(MODE_SILK_ONLY)
		if config >= 16 {
			want = int32(MODE_CELT_ONLY)
		} else if config >= 12 {
			want = int32(MODE_HYBRID)
		}
		b := byte(toc)
		if got := opus_packet_get_mode(&b); got != want {
			t.Fatalf("TOC %#x: got %d, want %d", toc, got, want)
		}
	}
}

// From decoder_duration_pointers_test.go.
func TestDecoderDurationPointers(t *testing.T) {
	for _, rate := range []int32{8000, 12000, 16000, 24000, 48000} {
		dec := OpusT_OpusDecoder{FFs: rate}
		before := dec
		for _, n := range []int32{-1, 0} {
			if got := Opus_opus_decoder_get_nb_samples(&dec, nil, n); got != -1 {
				t.Fatal(got)
			}
		}
		p := [2]byte{0x83, 48}
		if got := Opus_opus_decoder_get_nb_samples(&dec, &p[0], 2); got != rate*120/1000 {
			t.Fatal(rate, got)
		}
		p[1] = 49
		if got := Opus_opus_decoder_get_nb_samples(&dec, &p[0], 2); got != -4 {
			t.Fatal(rate, got)
		}
		if dec != before {
			t.Fatal("state modified")
		}
	}
}
