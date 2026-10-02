package opuscc

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"
	"unsafe"
)

func TestExtensionGenerateDescriptorsPointers(t *testing.T) {
	exts := func() []OpusT_opus_extension_data {
		out := make([]OpusT_opus_extension_data, 6)
		for i := range out {
			p := make([]byte, 300)
			for j := range p {
				p[j] = byte(i + j)
			}
			out[i] = OpusT_opus_extension_data{Fid: 32, Fframe: int32(i), Flen1: 300, Fdata: &p[0]}
		}
		return out
	}()
	entropyInitGrowStack(12)
	runtime.GC()
	if got := Opus_opus_packet_extensions_generate(nil, nil, 2000, &exts[0], 6, 6, 1); got != 2000 {
		t.Fatal("descriptor owner/sizing", got)
	}
	if exts[5].Fdata == nil || *exts[5].Fdata != 5 {
		t.Fatal("payload owner")
	}
	if Opus_opus_packet_extensions_generate(nil, nil, 0, nil, 0, 0, 0) != 0 || Opus_opus_packet_extensions_generate(nil, nil, 0, nil, 100, 49, 0) != -1 {
		t.Fatal("early validation")
	}
}

func TestExtensionGenerateOutputPointers(t *testing.T) {
	for _, capacity := range []int32{0, 1, 2, 8, 32, 500, 2000} {
		exts := make([]OpusT_opus_extension_data, 6)
		for i := range exts {
			p := make([]byte, 300)
			for j := range p {
				p[j] = byte(i + j)
			}
			exts[i] = OpusT_opus_extension_data{Fid: 32, Fframe: int32(i), Flen1: 300, Fdata: &p[0]}
		}
		out := make([]byte, capacity+2)
		out[0] = 77
		out[capacity+1] = 88
		entropyInitGrowStack(12)
		runtime.GC()
		Opus_opus_packet_extensions_generate(nil, &out[1], capacity, &exts[0], 6, 6, 1)
		if out[0] != 77 || out[capacity+1] != 88 {
			t.Fatal("generator guards", capacity)
		}
	}
	// A header writes before its aliased payload is read.
	out := []byte{77, 88}
	ext := OpusT_opus_extension_data{Fid: 3, Flen1: 1, Fdata: &out[0]}
	if r := Opus_opus_packet_extensions_generate(nil, &out[0], 2, &ext, 1, 1, 0); r != 2 || out[0] != 7 || out[1] != 7 {
		t.Fatal("payload alias", r, out)
	}
	// Numeric descriptor bytes remain live after the header store (no pointer poisoning).
	ext = OpusT_opus_extension_data{Fid: 112}
	panicked := false
	func() {
		defer func() { panicked = recover() != nil }()
		Opus_opus_packet_extensions_generate(nil, (*byte)(unsafe.Pointer(&ext.Fid)), 1, &ext, 1, 1, 0)
	}()
	if !panicked || ext.Fid != 224 {
		t.Fatal("descriptor reload", panicked, ext.Fid)
	}
}

func TestPacketExtensionsGenerateCReference(t *testing.T) {
	// Retained C generator covers repeated short/long extensions (including
	// interleaved frame order), separators, 255-byte length coding, padding,
	// exact/short buffers, NULL-output sizing, and invalid arguments.
	f, err := os.Open("testdata/extensions_generate_ref.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for line := 1; scanner.Scan(); line++ {
		input := strings.NewReader(scanner.Text())
		var name string
		var frames, pad, capacity, wantRet, wantDry, count int32
		var wantHash uint64
		if _, err := fmt.Fscan(input, &name, &frames, &pad, &capacity, &wantRet, &wantDry, &wantHash, &count); err != nil {
			t.Fatal(err)
		}
		t.Run(fmt.Sprintf("%s/%d", name, line), func(t *testing.T) {
			extensions := make([]OpusT_opus_extension_data, count)
			payload := make([]byte, int(count+1)*600)
			for i := range extensions {
				var id, frame, length, seed int32
				if _, err := fmt.Fscan(input, &id, &frame, &length, &seed); err != nil {
					t.Fatal(err)
				}
				bytes := payload[i*600 : (i+1)*600]
				p := unsafe.SliceData(bytes)
				for j := range bytes {
					bytes[j] = byte(seed + int32(j)*17)
				}
				extensions[i] = OpusT_opus_extension_data{Fid: id, Fframe: frame, Flen1: length, Fdata: p}
			}
			out := make([]byte, int(capacity)+16)
			for i := range out {
				out[i] = 0xa5
			}
			if got := Opus_opus_packet_extensions_generate(nil, unsafe.SliceData(out), capacity, unsafe.SliceData(extensions), count, frames, pad); got != wantRet {
				t.Fatalf("return %d, want %d", got, wantRet)
			}
			if got := Opus_opus_packet_extensions_generate(nil, nil, capacity, unsafe.SliceData(extensions), count, frames, pad); got != wantDry {
				t.Fatalf("sizing return %d, want %d", got, wantDry)
			}
			// Include the untouched suffix and 16 guard bytes, even on errors: the C
			// API may write a partial extension before reporting a short buffer.
			hash := uint64(14695981039346656037)
			for _, b := range out {
				hash ^= uint64(b)
				hash *= 1099511628211
			}
			if hash != wantHash {
				t.Fatalf("output hash %d, want %d", hash, wantHash)
			}
		})
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
}
