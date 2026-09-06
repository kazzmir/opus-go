package opuscc

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestMultistreamPacketValidateCReference(t *testing.T) {
	// C-generated results include self-delimited zero-length frames, multi-frame
	// packets, mismatched durations, missing streams and malformed framing.
	f, err := os.Open("testdata/multistream_validate_ref.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	tls := libc.NewTLS()
	defer tls.Close()
	data := libc.Xmalloc(tls, 64)
	defer libc.Xfree(tls, data)
	scanner := bufio.NewScanner(f)
	for line := 1; scanner.Scan(); line++ {
		var packet string
		var fs, streams, want int32
		if n, err := fmt.Sscanf(scanner.Text(), "%s %d %d %d", &packet, &fs, &streams, &want); err != nil || n != 4 {
			t.Fatalf("line %d: %v", line, err)
		}
		if packet == "-" {
			packet = ""
		}
		bytes, err := hex.DecodeString(packet)
		if err != nil {
			t.Fatal(err)
		}
		copy(unsafe.Slice((*byte)(unsafe.Pointer(data)), 64), bytes)
		if got := opus_multistream_packet_validate(tls, data, int32(len(bytes)), streams, fs); got != want {
			t.Fatalf("line %d (%s): got %d, want %d", line, scanner.Text(), got, want)
		}
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
}
