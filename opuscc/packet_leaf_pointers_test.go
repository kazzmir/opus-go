package opuscc

import "testing"

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
