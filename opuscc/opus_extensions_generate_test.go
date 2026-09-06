package opuscc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

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
			extMemory := libc.Xmalloc(tls, uint64(count+1)*uint64(unsafe.Sizeof(OpusT_opus_extension_data{})))
			defer libc.Xfree(tls, extMemory)
			extensions := unsafe.Slice((*OpusT_opus_extension_data)(unsafe.Pointer(extMemory)), int(count))
			payload := libc.Xmalloc(tls, uint64(count+1)*600)
			defer libc.Xfree(tls, payload)
			for i := range extensions {
				var id, frame, length, seed int32
				if _, err := fmt.Fscan(input, &id, &frame, &length, &seed); err != nil {
					t.Fatal(err)
				}
				p := payload + uintptr(i*600)
				bytes := unsafe.Slice((*byte)(unsafe.Pointer(p)), 600)
				for j := range bytes {
					bytes[j] = byte(seed + int32(j)*17)
				}
				extensions[i] = OpusT_opus_extension_data{Fid: id, Fframe: frame, Flen1: length, Fdata: p}
			}
			out := libc.Xmalloc(tls, uint64(capacity+16))
			defer libc.Xfree(tls, out)
			libc.Xmemset(tls, out, 0xa5, uint64(capacity+16))
			if got := Opus_opus_packet_extensions_generate(tls, out, capacity, extMemory, count, frames, pad); got != wantRet {
				t.Fatalf("return %d, want %d", got, wantRet)
			}
			if got := Opus_opus_packet_extensions_generate(tls, 0, capacity, extMemory, count, frames, pad); got != wantDry {
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
