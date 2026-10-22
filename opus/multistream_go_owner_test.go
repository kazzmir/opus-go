package opus

import (
	"runtime"
	"testing"
	"unsafe"
	"weak"

	libc "github.com/kazzmir/opus-go/libcshim"
	"github.com/kazzmir/opus-go/opuscc"
)

//go:noinline
func closedMultistreamGoOwner(t *testing.T) weak.Pointer[opuscc.OpusMSDecoder] {
	d, err := NewMultistreamDecoder(48000, 3, 2, 1, []byte{0, 1, 2})
	if err != nil {
		t.Fatal(err)
	}
	defer d.Close()
	if libc.RegisteredGoAllocation[opuscc.OpusMSDecoder](d.tls, uintptr(unsafe.Pointer(d.ms))) != nil {
		t.Fatal("multistream state registered in runtime TLS")
	}
	pcm := make([]float32, 360)
	runtime.GC()
	if n, err := d.DecodeF32(nil, pcm, 120, false); err != nil || n != 120 {
		t.Fatal("Go-owned multistream PLC", n, err)
	}
	w := weak.Make(d.ms)
	if err := d.Close(); err != nil || d.ms != nil || d.tls != nil {
		t.Fatal("close must drop Go owner and runtime context", err)
	}
	return w
}

func TestMultistreamGoOwnerClosePointers(t *testing.T) {
	w := closedMultistreamGoOwner(t)
	for trial := 0; trial < 10; trial++ {
		runtime.GC()
	}
	if w.Value() != nil {
		t.Fatal("closed multistream wrapper retains Go state")
	}
}
