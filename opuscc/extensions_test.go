package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

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
}
