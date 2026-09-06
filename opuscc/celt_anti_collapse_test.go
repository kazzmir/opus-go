package opuscc

import (
	"testing"
	"unsafe"

	libc "github.com/kazzmir/opus-go/libcshim"
)

func TestAntiCollapseLocalExp2Union(t *testing.T) {
	tls := libc.NewTLS()
	defer tls.Close()

	eBands := [2]int16{0, 2}
	mode := OpusT_OpusCustomMode{
		FnbEBands: 1,
		FeBands:   uintptr(unsafe.Pointer(&eBands[0])),
	}
	x := [2]OpusT_celt_norm{0, 0}
	collapseMasks := [1]byte{0}
	logE := [1]OpusT_celt_glog{2.5}
	prev1LogE := [1]OpusT_celt_glog{0.25}
	prev2LogE := [1]OpusT_celt_glog{0.75}
	pulses := [1]int32{3}

	Opus_anti_collapse(
		tls,
		uintptr(unsafe.Pointer(&mode)),
		uintptr(unsafe.Pointer(&x[0])),
		uintptr(unsafe.Pointer(&collapseMasks[0])),
		0,
		1,
		2,
		0,
		1,
		uintptr(unsafe.Pointer(&logE[0])),
		uintptr(unsafe.Pointer(&prev1LogE[0])),
		uintptr(unsafe.Pointer(&prev2LogE[0])),
		uintptr(unsafe.Pointer(&pulses[0])),
		123456,
		0,
		0,
	)

	if got, want := x, [2]OpusT_celt_norm{0.7071068, 0.7071068}; got != want {
		t.Fatalf("repaired spectrum: got %v, want %v", got, want)
	}
}
