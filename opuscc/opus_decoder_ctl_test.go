package opuscc

import (
	"bufio"
	"fmt"
	"os"
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestOpusDecoderCtlCReference(t *testing.T) {
	// Replay the C-generated transcript; generator and build command are in testdata.
	f, err := os.Open("testdata/decoder_ctl_ref.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	tls := libc.NewTLS()
	defer tls.Close()
	st, err := Opus_opus_decoder_create_typed(48000, 2)
	if err != nil {
		t.Fatal(err)
	}
	defer Opus_opus_decoder_destroy_typed(st)
	va := libc.XmallocPointer(tls, uint64(unsafe.Sizeof(uintptr(0))))
	defer libc.XfreePointer(tls, va)
	out := libc.XmallocPointer(tls, 4)
	defer libc.XfreePointer(tls, out)
	d := st
	celt := (*OpusT_OpusCustomDecoder)(unsafe.Add(unsafe.Pointer(st), d.Fcelt_dec_offset))
	scanner := bufio.NewScanner(f)
	for line := 1; scanner.Scan(); line++ {
		var op string
		var request, arg, wantRet, wantValue int32
		if n, err := fmt.Sscanf(scanner.Text(), "%s %d %d %d %d", &op, &request, &arg, &wantRet, &wantValue); err != nil || n != 5 {
			t.Fatalf("line %d: %v", line, err)
		}
		switch op {
		case "seed":
			d.FDecControl.FprevPitchLag = 77
			d.Fprev_mode = MODE_SILK_ONLY
			d.Fbandwidth = OPUS_BANDWIDTH_WIDEBAND
			d.Flast_packet_duration = 960
			d.FrangeFinal = 123456789
			continue
		case "celt":
			d.Fprev_mode = MODE_CELT_ONLY
			continue
		case "get":
			*(*int32)(out) = -999
			*(*unsafe.Pointer)(va) = out
		case "null":
			*(*unsafe.Pointer)(va) = nil
		case "set":
			*(*uintptr)(va) = uintptr(uint32(arg))
		case "reset":
		default:
			t.Fatalf("unknown operation %q", op)
		}
		ret := Opus_opus_decoder_ctl(tls, d, request, libc.NewVaCursor(va))
		if ret != wantRet {
			t.Fatalf("line %d (%s): return %d, want %d", line, scanner.Text(), ret, wantRet)
		}
		if op == "get" && *(*int32)(out) != wantValue {
			t.Fatalf("line %d: output %d, want %d", line, *(*int32)(out), wantValue)
		}
		if request == OPUS_SET_COMPLEXITY_REQUEST && celt.Fcomplexity != d.Fcomplexity {
			t.Fatalf("line %d: complexity not forwarded", line)
		}
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
}
