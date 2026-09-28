package opuscc

import (
	"runtime"
	"testing"
	"unsafe"
)

func TestEntropyWritePointers(t *testing.T) {
	buffer := [3]byte{}
	// Fbuf is still an integer field; keep its backing storage live explicitly.
	defer func() { runtime.KeepAlive(&buffer) }()
	enc := OpusT_ec_enc{
		Fbuf:     uintptr(unsafe.Pointer(&buffer[0])),
		Fstorage: uint32(len(buffer)),
	}
	if got := ec_write_byte(nil, &enc, 0x123); got != 0 {
		t.Fatalf("front write: %d", got)
	}
	if got := ec_write_byte_at_end(nil, &enc, 0x245); got != 0 {
		t.Fatalf("back write: %d", got)
	}
	if got := ec_write_byte(nil, &enc, 0x67); got != 0 {
		t.Fatalf("last free byte: %d", got)
	}
	if want := [3]byte{0x23, 0x67, 0x45}; buffer != want {
		t.Fatalf("buffer: got %x, want %x", buffer, want)
	}
	before := enc
	if got := ec_write_byte(nil, &enc, 0xff); got != -1 {
		t.Fatalf("full front write: %d", got)
	}
	if got := ec_write_byte_at_end(nil, &enc, 0xff); got != -1 {
		t.Fatalf("full back write: %d", got)
	}
	if enc != before || buffer != [3]byte{0x23, 0x67, 0x45} {
		t.Fatal("failed write modified state or buffer")
	}
	empty := OpusT_ec_enc{}
	if ec_write_byte(nil, &empty, 1) != -1 || ec_write_byte_at_end(nil, &empty, 1) != -1 || empty != (OpusT_ec_enc{}) {
		t.Fatal("zero-capacity writer should fail without accessing a buffer")
	}
}

func TestStereoSplitPointers(t *testing.T) {
	// Interior pointers and sentinels check both ends of the requested span.
	x := [...]float32{99, 1, -2, 0, 0.125, 88}
	y := [...]float32{77, -1, 3, 0, 0.25, 66}
	wantX, wantY := x, y
	for i := 1; i < len(x)-1; i++ {
		l := float32(float32(0.70710678) * x[i])
		r := float32(float32(0.70710678) * y[i])
		wantX[i], wantY[i] = l+r, r-l
	}
	stereo_split(nil, &x[1], &y[1], 4)
	if x != wantX || y != wantY {
		t.Fatalf("got X=%v Y=%v, want X=%v Y=%v", x, y, wantX, wantY)
	}
	stereo_split(nil, nil, nil, 0)
}
