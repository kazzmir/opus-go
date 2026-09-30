package opuscc

import (
	"runtime"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestWritePayloadPointers(t *testing.T) {
	payload := make([]byte, 255)
	for i := range payload {
		payload[i] = byte(i)
	}
	out := make([]byte, 260)
	out[0] = 77
	out[259] = 88
	entropyInitGrowStack(12)
	runtime.GC()
	r := write_extension_payload(nil, &out[0], 259, 1, 32, 255, &payload[0], 0)
	if r != 258 || out[1] != 255 || out[2] != 0 || out[3] != 0 || out[257] != 254 || out[0] != 77 || out[259] != 88 {
		t.Fatal(r, out)
	}
	if write_extension_payload(nil, nil, 259, 1, 32, 255, nil, 0) != 258 {
		t.Fatal("size-only")
	}
	before := append([]byte(nil), out...)
	if write_extension_payload(nil, &out[0], 2, 1, 32, 255, &payload[0], 0) != -2 {
		t.Fatal("capacity")
	}
	for i := range out {
		if out[i] != before[i] {
			t.Fatal("failed write changed data")
		}
	}
}

func TestSkipExtensionPointers(t *testing.T) {
	p := (*byte)(nil)
	h := int32(77)
	if skip_extension(nil, &p, 0, &h) != 0 || h != 0 || p != nil {
		t.Fatal("empty")
	}
	h = 77
	if skip_extension(nil, &p, -1, &h) != -1 || h != 77 {
		t.Fatal("negative")
	}
	data := [6]byte{65, 2, 11, 12, 3, 99}
	p = &data[0]
	h = 77
	if r := skip_extension(nil, &p, 6, &h); r != 2 || h != 2 || p != &data[4] {
		t.Fatal(r, h, p)
	}
	p = &data[0]
	h = 77
	if r := skip_extension(nil, &p, 1, &h); r != -1 || h != 77 || p != &data[0] {
		t.Fatal("error committed", r, h, p)
	}
}

func TestSkipPayloadPointers(t *testing.T) {
	data := [8]byte{3, 1, 2, 3, 4, 5, 6, 7}
	p := &data[0]
	h := int32(77)
	if r := skip_extension_payload(nil, &p, 8, &h, 65, 0); r != 4 || p != &data[4] || h != 1 {
		t.Fatal(r, p, h)
	}
	before := p
	h = 77
	if r := skip_extension_payload(nil, &p, 1, &h, 65, 0); r != -1 || p != before || h != 77 {
		t.Fatal("error committed output")
	}
	var owner *byte
	func() {
		b := make([]byte, 8)
		b[0] = 1
		b[2] = 99
		p := &b[0]
		h := int32(0)
		skip_extension_payload(nil, &p, 8, &h, 65, 0)
		owner = p
	}()
	entropyInitGrowStack(12)
	runtime.GC()
	if *owner != 99 {
		t.Fatal("interior ownership")
	}
}

func TestRepeatedExtensionIterator(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	payload := []byte{'x'}
	extensions := []OpusT_opus_extension_data{
		{Fid: 3, Fframe: 0, Fdata: uintptr(unsafe.Pointer(&payload[0])), Flen1: 1},
		{Fid: 3, Fframe: 1, Fdata: uintptr(unsafe.Pointer(&payload[0])), Flen1: 1},
		{Fid: 3, Fframe: 2, Fdata: uintptr(unsafe.Pointer(&payload[0])), Flen1: 1},
	}
	packet := make([]byte, 32)
	length := Opus_opus_packet_extensions_generate(tls, uintptr(unsafe.Pointer(&packet[0])), int32(len(packet)), uintptr(unsafe.Pointer(&extensions[0])), int32(len(extensions)), 3, 1)
	if length <= 0 {
		t.Fatalf("generated extension packet length: got %d", length)
	}

	var iterator OpusT_OpusExtensionIterator
	Opus_opus_extension_iterator_init(tls, uintptr(unsafe.Pointer(&iterator)), uintptr(unsafe.Pointer(&packet[0])), length, 3)
	for frame := int32(0); frame < 3; frame++ {
		var extension OpusT_opus_extension_data
		if got := Opus_opus_extension_iterator_next(tls, uintptr(unsafe.Pointer(&iterator)), uintptr(unsafe.Pointer(&extension))); got != 1 {
			t.Fatalf("frame %d iterator result: got %d, want 1; length=%d packet=% x", frame, got, length, packet[:length])
		}
		if extension.Fid != 3 || extension.Fframe != frame || extension.Flen1 != 1 || *(*byte)(unsafe.Pointer(extension.Fdata)) != 'x' {
			t.Fatalf("frame %d extension: %+v", frame, extension)
		}
	}
	if got := Opus_opus_extension_iterator_next(tls, uintptr(unsafe.Pointer(&iterator)), 0); got != 0 {
		t.Fatalf("iterator exhaustion: got %d, want 0", got)
	}

	Opus_opus_extension_iterator_init(tls, uintptr(unsafe.Pointer(&iterator)), uintptr(unsafe.Pointer(&packet[0])), length, 3)
	var found OpusT_opus_extension_data
	if got := Opus_opus_extension_iterator_find(tls, uintptr(unsafe.Pointer(&iterator)), uintptr(unsafe.Pointer(&found)), 3); got != 1 {
		t.Fatalf("find result: got %d, want 1", got)
	}
	if found.Fframe != 0 || found.Flen1 != 1 || *(*byte)(unsafe.Pointer(found.Fdata)) != 'x' {
		t.Fatalf("found extension: %+v", found)
	}
	if got := Opus_opus_packet_extensions_count(tls, uintptr(unsafe.Pointer(&packet[0])), length, 3); got != 3 {
		t.Fatalf("extension count: got %d, want 3", got)
	}
	parsed := make([]OpusT_opus_extension_data, 3)
	count := int32(len(parsed))
	if got := Opus_opus_packet_extensions_parse(tls, uintptr(unsafe.Pointer(&packet[0])), length, uintptr(unsafe.Pointer(&parsed[0])), uintptr(unsafe.Pointer(&count)), 3); got != 0 {
		t.Fatalf("parse result: got %d, want 0", got)
	}
	if count != 3 || parsed[2].Fid != 3 || parsed[2].Fframe != 2 || *(*byte)(unsafe.Pointer(parsed[2].Fdata)) != 'x' {
		t.Fatalf("parsed extensions: count=%d entries=%+v", count, parsed)
	}
	frameCounts := make([]OpusT_opus_int32, 3)
	if got := Opus_opus_packet_extensions_count_ext(tls, uintptr(unsafe.Pointer(&packet[0])), length, uintptr(unsafe.Pointer(&frameCounts[0])), 3); got != 3 {
		t.Fatalf("per-frame extension count: got %d, want 3", got)
	}
	if want := []OpusT_opus_int32{1, 1, 1}; frameCounts[0] != want[0] || frameCounts[1] != want[1] || frameCounts[2] != want[2] {
		t.Fatalf("per-frame extension counts: got %v, want %v", frameCounts, want)
	}
	ordered := make([]OpusT_opus_extension_data, 3)
	count = int32(len(ordered))
	if got := Opus_opus_packet_extensions_parse_ext(tls, uintptr(unsafe.Pointer(&packet[0])), length, uintptr(unsafe.Pointer(&ordered[0])), uintptr(unsafe.Pointer(&count)), uintptr(unsafe.Pointer(&frameCounts[0])), 3); got != 0 {
		t.Fatalf("frame-ordered parse result: got %d, want 0", got)
	}
	if count != 3 || ordered[0].Fframe != 0 || ordered[1].Fframe != 1 || ordered[2].Fframe != 2 {
		t.Fatalf("frame-ordered parsed extensions: count=%d entries=%+v", count, ordered)
	}
}
