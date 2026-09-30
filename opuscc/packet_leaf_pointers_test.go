package opuscc

import (
	"runtime"
	"testing"
)

func TestPacketLBRRPointers(t *testing.T) {
	for _, tc := range []struct {
		toc, payload byte
		want         int32
	}{{0, 0x40, 1}, {0, 0, 0}, {4, 0x10, 1}, {0x10, 0x20, 1}, {0x18, 0x10, 1}, {0x1c, 1, 1}, {0x80, 0xff, 0}} {
		packet := [2]byte{tc.toc, tc.payload}
		entropyInitGrowStack(12)
		runtime.GC()
		if got := Opus_opus_packet_has_lbrr(nil, &packet[0], 2); got != tc.want {
			t.Fatal(tc, got)
		}
	}
	header := byte(0)
	if Opus_opus_packet_has_lbrr(nil, &header, 1) != 0 {
		t.Fatal("zero-size frame")
	}
	if Opus_opus_packet_has_lbrr(nil, &header, 0) != -4 {
		t.Fatal("empty SILK packet")
	}
	header = 0x80
	if Opus_opus_packet_has_lbrr(nil, &header, -1) != 0 {
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
		r := Opus_opus_packet_parse(nil, &packet[0], 5, &toc, &frames, &sizes, &offset)
		if r != 2 || sizes[0] != 1 || sizes[1] != 2 || offset != 2 || toc != 0x82 {
			t.Fatal(r, sizes, offset, toc)
		}
	}()
	entropyInitGrowStack(12)
	runtime.GC()
	if *frames[0] != 11 || *frames[1] != 12 {
		t.Fatal("returned frame ownership")
	}
	if r := Opus_opus_packet_parse(nil, nil, 0, nil, nil, &sizes, nil); r != -4 {
		t.Fatal("empty", r)
	}
	if r := Opus_opus_packet_parse(nil, nil, 0, nil, nil, nil, nil); r != -1 {
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
		r := Opus_opus_packet_parse_impl(nil, &packet[0], 9, 0, &toc, &frames, &size, &payload, &offset, &pad, &padLen)
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
	r := Opus_opus_packet_parse_impl(nil, &packet[0], 10, 1, nil, &frames, &size, &payload, &offset, nil, nil)
	if r != 2 || size[0] != 3 || size[1] != 3 || payload != 2 || offset != 8 {
		t.Fatal("self-delimited", r, size, payload, offset)
	}
	pad = &packet[0]
	padLen = 99
	if r := Opus_opus_packet_parse_impl(nil, nil, -1, 0, nil, nil, nil, nil, nil, &pad, &padLen); r != -1 || pad != nil || padLen != 0 {
		t.Fatal("error padding", r, pad, padLen)
	}
}

func TestPacketSizePointers(t *testing.T) {
	// Every representable Opus frame length, including the 251/252 boundary.
	// Check the C wire format independently as well as round-tripping it.
	for size := int32(0); size <= 1275; size++ {
		buf := [4]byte{0xa5, 0xcc, 0xcc, 0x5a}
		n := Opus_encode_size(nil, size, &buf[1])
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
		if got := parse_size(nil, &buf[1], n, &decoded[1]); got != n || decoded != [3]int16{1234, int16(size), 5678} {
			t.Fatalf("size %d: parsed n=%d output=%v", size, got, decoded)
		}
	}
	// The single-byte branch must not require a second accessible byte.
	var single byte
	if n := Opus_encode_size(nil, 251, &single); n != 1 || single != 251 {
		t.Fatalf("single byte encode: n=%d byte=%d", n, single)
	}
	for _, length := range []int32{-1, 0} {
		var size int16
		if n := parse_size(nil, nil, length, &size); n != -1 || size != -1 {
			t.Fatalf("length %d: n=%d size=%d", length, n, size)
		}
	}
	for first := 0; first < 256; first++ {
		b := byte(first)
		var size int16
		n := parse_size(nil, &b, 1, &size)
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
		if got := opus_packet_get_mode(nil, &b); got != want {
			t.Fatalf("TOC %#x: got %d, want %d", toc, got, want)
		}
	}
}
