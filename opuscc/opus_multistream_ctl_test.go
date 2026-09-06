package opuscc

import (
	"bufio"
	"fmt"
	"os"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestMultistreamDecoderCtlCReference(t *testing.T) {
	// C-generated transcript covers one coupled and two mono streams, including
	// distinct per-stream state to detect first-stream queries and XOR mistakes.
	f, err := os.Open("testdata/multistream_ctl_ref.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	tls := libc.NewTLS()
	defer tls.Close()
	mapping := libc.Xmalloc(tls, 4)
	defer libc.Xfree(tls, mapping)
	copy(unsafe.Slice((*byte)(unsafe.Pointer(mapping)), 4), []byte{0, 1, 2, 3})
	st, err := Opus_opus_multistream_decoder_create(tls, 48000, 4, 3, 1, mapping)
	if err != nil {
		t.Fatal(err)
	}
	defer Opus_opus_multistream_decoder_destroy(tls, st)
	va := libc.Xmalloc(tls, 2*uint64(unsafe.Sizeof(uintptr(0))))
	defer libc.Xfree(tls, va)
	out := libc.Xmalloc(tls, uint64(unsafe.Sizeof(uintptr(0))))
	defer libc.Xfree(tls, out)
	var dec [3]uintptr
	scanner := bufio.NewScanner(f)
	for line := 1; scanner.Scan(); line++ {
		var op string
		var stream, request, arg, wantRet int32
		var wantValue int64
		if n, err := fmt.Sscanf(scanner.Text(), "%s %d %d %d %d %d", &op, &stream, &request, &arg, &wantRet, &wantValue); err != nil || n != 6 {
			t.Fatalf("line %d: %v", line, err)
		}
		switch op {
		case "seed":
			for i, p := range dec {
				d := (*OpusT_OpusDecoder)(unsafe.Pointer(p))
				d.FrangeFinal = 0x87654321 + uint32(i)*0x12345678
				d.Fbandwidth = 1101 + int32(i)
				d.Flast_packet_duration = 120 * (int32(i) + 1)
				d.Fdecode_gain = 100 + int32(i)
			}
			continue
		case "get":
			*(*int32)(unsafe.Pointer(out)) = -999
			libc.VaList(va, out)
		case "null":
			libc.VaList(va, uintptr(0))
		case "set":
			libc.VaList(va, arg)
		case "state":
			*(*uintptr)(unsafe.Pointer(out)) = 0
			libc.VaList(va, arg, out)
		case "state_null":
			libc.VaList(va, arg, uintptr(0))
		case "reset":
		default:
			t.Fatalf("unknown operation %s", op)
		}
		var ret int32
		if stream < 0 {
			ret = Opus_opus_multistream_decoder_ctl_va_list(tls, st, request, va)
		} else {
			ret = Opus_opus_decoder_ctl(tls, dec[stream], request, va)
		}
		if ret != wantRet {
			t.Fatalf("line %d (%s): return %d, want %d", line, scanner.Text(), ret, wantRet)
		}
		if op == "get" {
			if got := int64(*(*int32)(unsafe.Pointer(out))); got != wantValue {
				t.Fatalf("line %d: got %d, want %d", line, got, wantValue)
			}
		}
		if op == "state" {
			p := *(*uintptr)(unsafe.Pointer(out))
			offset := int64(0)
			if p != 0 {
				offset = int64(p - st)
			}
			// Each preceding stream occupies one decoder allocation.
			if ret == 0 {
				wantValue += int64(arg) * int64(decoderReferenceLayoutDelta())
			}
			if offset != wantValue {
				t.Fatalf("line %d: offset %d, want %d", line, offset, wantValue)
			}
			if ret == 0 {
				dec[arg] = p
			}
		}
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
}
