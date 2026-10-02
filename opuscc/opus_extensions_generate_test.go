package opuscc

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
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
	if got := Opus_opus_packet_extensions_generate(nil, 0, 2000, &exts[0], 6, 6, 1); got != 2000 {
		t.Fatal("descriptor owner/sizing", got)
	}
	if exts[5].Fdata == nil || *exts[5].Fdata != 5 {
		t.Fatal("payload owner")
	}
	if Opus_opus_packet_extensions_generate(nil, 0, 0, nil, 0, 0, 0) != 0 || Opus_opus_packet_extensions_generate(nil, 0, 0, nil, 100, 49, 0) != -1 {
		t.Fatal("early validation")
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
			tls := libc.NewTLS()
			defer tls.Close()
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
			out := libc.Xmalloc(tls, uint64(capacity+16))
			defer libc.Xfree(tls, out)
			libc.Xmemset(tls, out, 0xa5, uint64(capacity+16))
			if got := Opus_opus_packet_extensions_generate(tls, out, capacity, unsafe.SliceData(extensions), count, frames, pad); got != wantRet {
				t.Fatalf("return %d, want %d", got, wantRet)
			}
			if got := Opus_opus_packet_extensions_generate(tls, 0, capacity, unsafe.SliceData(extensions), count, frames, pad); got != wantDry {
				t.Fatalf("sizing return %d, want %d", got, wantDry)
			}
			// Include the untouched suffix and 16 guard bytes, even on errors: the C
			// API may write a partial extension before reporting a short buffer.
			hash := uint64(14695981039346656037)
			for _, b := range unsafe.Slice((*byte)(unsafe.Pointer(out)), int(capacity)+16) {
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
