package opuscc

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"runtime"
	"testing"
	"unsafe"
)

func TestMSPacketValidationPointers(t *testing.T) {
	for _, tc := range []struct {
		packet              []byte
		streams, rate, want int32
	}{
		{[]byte{0}, 1, 48000, 480}, {[]byte{0, 0, 0}, 2, 48000, 480}, {[]byte{0, 0, 0x80}, 2, 48000, -4}, {nil, 1, 48000, -4}, {nil, 0, 48000, 0}, {nil, -1, 48000, 0}, {[]byte{0, 0, 0}, 2, 8000, 80},
	} {
		entropyInitGrowStack(12)
		runtime.GC()
		got := opus_multistream_packet_validate(nil, unsafe.SliceData(tc.packet), int32(len(tc.packet)), tc.streams, tc.rate)
		if got != tc.want {
			t.Fatal(tc, got)
		}
	}
}

func TestMultistreamPacketValidateCReference(t *testing.T) {
	// C-generated results include self-delimited zero-length frames, multi-frame
	// packets, mismatched durations, missing streams and malformed framing.
	f, err := os.Open("testdata/multistream_validate_ref.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
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
		if got := opus_multistream_packet_validate(nil, unsafe.SliceData(bytes), int32(len(bytes)), streams, fs); got != want {
			t.Fatalf("line %d (%s): got %d, want %d", line, scanner.Text(), got, want)
		}
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
}
